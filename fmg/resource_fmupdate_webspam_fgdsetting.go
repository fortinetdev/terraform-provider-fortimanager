// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the FortiGuard run parameters.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateWebSpamFgdSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateWebSpamFgdSettingUpdate,
		Read:   resourceFmupdateWebSpamFgdSettingRead,
		Update: resourceFmupdateWebSpamFgdSettingUpdate,
		Delete: resourceFmupdateWebSpamFgdSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"as_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"as_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"av_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av2_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"av2_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av2_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eventlog_query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_pull_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fq_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fq_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fq_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iot_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iot_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iotv_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkd_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_client_worker": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_log_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_unrated_site": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"restrict_as1_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_as2_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_as4_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_av_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_av2_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_fq_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_iots_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restrict_wf_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"stat_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stat_log_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stat_sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"update_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wf_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wf_dn_cache_expire_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wf_dn_cache_max_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wf_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wf_preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateWebSpamFgdSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateWebSpamFgdSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSetting resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFmupdateWebSpamFgdSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateWebSpamFgdSetting")

	return resourceFmupdateWebSpamFgdSettingRead(d, m)
}

func resourceFmupdateWebSpamFgdSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteFmupdateWebSpamFgdSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateWebSpamFgdSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateWebSpamFgdSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateWebSpamFgdSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamFgdSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateWebSpamFgdSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamFgdSetting resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateWebSpamFgdSettingAsCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAsPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAvCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAvLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAvPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAv2Cache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAv2Log(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAv2Preload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingEventlogQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFgdPullInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFqCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFqLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFqPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotvPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingLinkdLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingMaxClientWorker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingMaxLogQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingMaxUnratedSite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAs1Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAs2Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAs4Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAvDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAv2Dbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictFqDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictIotsDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictWfDbver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := i["servlist"]; ok {
		result["servlist"] = flattenFmupdateWebSpamFgdSettingServerOverrideServlist(i["servlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateWebSpamFgdSettingServerOverrideStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlist(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-ServiceType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingStatLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingStatLogInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingStatSyncInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingUpdateLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfDnCacheExpireTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfDnCacheMaxNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfPreload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateWebSpamFgdSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("as_cache", flattenFmupdateWebSpamFgdSettingAsCache(o["as-cache"], d, "as_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-cache"], "FmupdateWebSpamFgdSetting-AsCache"); ok {
			if err = d.Set("as_cache", vv); err != nil {
				return fmt.Errorf("Error reading as_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_cache: %v", err)
		}
	}

	if err = d.Set("as_log", flattenFmupdateWebSpamFgdSettingAsLog(o["as-log"], d, "as_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-log"], "FmupdateWebSpamFgdSetting-AsLog"); ok {
			if err = d.Set("as_log", vv); err != nil {
				return fmt.Errorf("Error reading as_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_log: %v", err)
		}
	}

	if err = d.Set("as_preload", flattenFmupdateWebSpamFgdSettingAsPreload(o["as-preload"], d, "as_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-preload"], "FmupdateWebSpamFgdSetting-AsPreload"); ok {
			if err = d.Set("as_preload", vv); err != nil {
				return fmt.Errorf("Error reading as_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_preload: %v", err)
		}
	}

	if err = d.Set("av_cache", flattenFmupdateWebSpamFgdSettingAvCache(o["av-cache"], d, "av_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-cache"], "FmupdateWebSpamFgdSetting-AvCache"); ok {
			if err = d.Set("av_cache", vv); err != nil {
				return fmt.Errorf("Error reading av_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_cache: %v", err)
		}
	}

	if err = d.Set("av_log", flattenFmupdateWebSpamFgdSettingAvLog(o["av-log"], d, "av_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-log"], "FmupdateWebSpamFgdSetting-AvLog"); ok {
			if err = d.Set("av_log", vv); err != nil {
				return fmt.Errorf("Error reading av_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_log: %v", err)
		}
	}

	if err = d.Set("av_preload", flattenFmupdateWebSpamFgdSettingAvPreload(o["av-preload"], d, "av_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-preload"], "FmupdateWebSpamFgdSetting-AvPreload"); ok {
			if err = d.Set("av_preload", vv); err != nil {
				return fmt.Errorf("Error reading av_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_preload: %v", err)
		}
	}

	if err = d.Set("av2_cache", flattenFmupdateWebSpamFgdSettingAv2Cache(o["av2-cache"], d, "av2_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-cache"], "FmupdateWebSpamFgdSetting-Av2Cache"); ok {
			if err = d.Set("av2_cache", vv); err != nil {
				return fmt.Errorf("Error reading av2_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_cache: %v", err)
		}
	}

	if err = d.Set("av2_log", flattenFmupdateWebSpamFgdSettingAv2Log(o["av2-log"], d, "av2_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-log"], "FmupdateWebSpamFgdSetting-Av2Log"); ok {
			if err = d.Set("av2_log", vv); err != nil {
				return fmt.Errorf("Error reading av2_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_log: %v", err)
		}
	}

	if err = d.Set("av2_preload", flattenFmupdateWebSpamFgdSettingAv2Preload(o["av2-preload"], d, "av2_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-preload"], "FmupdateWebSpamFgdSetting-Av2Preload"); ok {
			if err = d.Set("av2_preload", vv); err != nil {
				return fmt.Errorf("Error reading av2_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_preload: %v", err)
		}
	}

	if err = d.Set("eventlog_query", flattenFmupdateWebSpamFgdSettingEventlogQuery(o["eventlog-query"], d, "eventlog_query")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventlog-query"], "FmupdateWebSpamFgdSetting-EventlogQuery"); ok {
			if err = d.Set("eventlog_query", vv); err != nil {
				return fmt.Errorf("Error reading eventlog_query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventlog_query: %v", err)
		}
	}

	if err = d.Set("fgd_pull_interval", flattenFmupdateWebSpamFgdSettingFgdPullInterval(o["fgd-pull-interval"], d, "fgd_pull_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-pull-interval"], "FmupdateWebSpamFgdSetting-FgdPullInterval"); ok {
			if err = d.Set("fgd_pull_interval", vv); err != nil {
				return fmt.Errorf("Error reading fgd_pull_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_pull_interval: %v", err)
		}
	}

	if err = d.Set("fq_cache", flattenFmupdateWebSpamFgdSettingFqCache(o["fq-cache"], d, "fq_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-cache"], "FmupdateWebSpamFgdSetting-FqCache"); ok {
			if err = d.Set("fq_cache", vv); err != nil {
				return fmt.Errorf("Error reading fq_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_cache: %v", err)
		}
	}

	if err = d.Set("fq_log", flattenFmupdateWebSpamFgdSettingFqLog(o["fq-log"], d, "fq_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-log"], "FmupdateWebSpamFgdSetting-FqLog"); ok {
			if err = d.Set("fq_log", vv); err != nil {
				return fmt.Errorf("Error reading fq_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_log: %v", err)
		}
	}

	if err = d.Set("fq_preload", flattenFmupdateWebSpamFgdSettingFqPreload(o["fq-preload"], d, "fq_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-preload"], "FmupdateWebSpamFgdSetting-FqPreload"); ok {
			if err = d.Set("fq_preload", vv); err != nil {
				return fmt.Errorf("Error reading fq_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_preload: %v", err)
		}
	}

	if err = d.Set("iot_cache", flattenFmupdateWebSpamFgdSettingIotCache(o["iot-cache"], d, "iot_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-cache"], "FmupdateWebSpamFgdSetting-IotCache"); ok {
			if err = d.Set("iot_cache", vv); err != nil {
				return fmt.Errorf("Error reading iot_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_cache: %v", err)
		}
	}

	if err = d.Set("iot_log", flattenFmupdateWebSpamFgdSettingIotLog(o["iot-log"], d, "iot_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-log"], "FmupdateWebSpamFgdSetting-IotLog"); ok {
			if err = d.Set("iot_log", vv); err != nil {
				return fmt.Errorf("Error reading iot_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_log: %v", err)
		}
	}

	if err = d.Set("iot_preload", flattenFmupdateWebSpamFgdSettingIotPreload(o["iot-preload"], d, "iot_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-preload"], "FmupdateWebSpamFgdSetting-IotPreload"); ok {
			if err = d.Set("iot_preload", vv); err != nil {
				return fmt.Errorf("Error reading iot_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_preload: %v", err)
		}
	}

	if err = d.Set("iotv_preload", flattenFmupdateWebSpamFgdSettingIotvPreload(o["iotv-preload"], d, "iotv_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["iotv-preload"], "FmupdateWebSpamFgdSetting-IotvPreload"); ok {
			if err = d.Set("iotv_preload", vv); err != nil {
				return fmt.Errorf("Error reading iotv_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iotv_preload: %v", err)
		}
	}

	if err = d.Set("linkd_log", flattenFmupdateWebSpamFgdSettingLinkdLog(o["linkd-log"], d, "linkd_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["linkd-log"], "FmupdateWebSpamFgdSetting-LinkdLog"); ok {
			if err = d.Set("linkd_log", vv); err != nil {
				return fmt.Errorf("Error reading linkd_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linkd_log: %v", err)
		}
	}

	if err = d.Set("max_client_worker", flattenFmupdateWebSpamFgdSettingMaxClientWorker(o["max-client-worker"], d, "max_client_worker")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-client-worker"], "FmupdateWebSpamFgdSetting-MaxClientWorker"); ok {
			if err = d.Set("max_client_worker", vv); err != nil {
				return fmt.Errorf("Error reading max_client_worker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_client_worker: %v", err)
		}
	}

	if err = d.Set("max_log_quota", flattenFmupdateWebSpamFgdSettingMaxLogQuota(o["max-log-quota"], d, "max_log_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-quota"], "FmupdateWebSpamFgdSetting-MaxLogQuota"); ok {
			if err = d.Set("max_log_quota", vv); err != nil {
				return fmt.Errorf("Error reading max_log_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_quota: %v", err)
		}
	}

	if err = d.Set("max_unrated_site", flattenFmupdateWebSpamFgdSettingMaxUnratedSite(o["max-unrated-site"], d, "max_unrated_site")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-unrated-site"], "FmupdateWebSpamFgdSetting-MaxUnratedSite"); ok {
			if err = d.Set("max_unrated_site", vv); err != nil {
				return fmt.Errorf("Error reading max_unrated_site: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_unrated_site: %v", err)
		}
	}

	if err = d.Set("restrict_as1_dbver", flattenFmupdateWebSpamFgdSettingRestrictAs1Dbver(o["restrict-as1-dbver"], d, "restrict_as1_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as1-dbver"], "FmupdateWebSpamFgdSetting-RestrictAs1Dbver"); ok {
			if err = d.Set("restrict_as1_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as1_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as1_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_as2_dbver", flattenFmupdateWebSpamFgdSettingRestrictAs2Dbver(o["restrict-as2-dbver"], d, "restrict_as2_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as2-dbver"], "FmupdateWebSpamFgdSetting-RestrictAs2Dbver"); ok {
			if err = d.Set("restrict_as2_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as2_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as2_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_as4_dbver", flattenFmupdateWebSpamFgdSettingRestrictAs4Dbver(o["restrict-as4-dbver"], d, "restrict_as4_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as4-dbver"], "FmupdateWebSpamFgdSetting-RestrictAs4Dbver"); ok {
			if err = d.Set("restrict_as4_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as4_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as4_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_av_dbver", flattenFmupdateWebSpamFgdSettingRestrictAvDbver(o["restrict-av-dbver"], d, "restrict_av_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-av-dbver"], "FmupdateWebSpamFgdSetting-RestrictAvDbver"); ok {
			if err = d.Set("restrict_av_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_av_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_av_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_av2_dbver", flattenFmupdateWebSpamFgdSettingRestrictAv2Dbver(o["restrict-av2-dbver"], d, "restrict_av2_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-av2-dbver"], "FmupdateWebSpamFgdSetting-RestrictAv2Dbver"); ok {
			if err = d.Set("restrict_av2_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_av2_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_av2_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_fq_dbver", flattenFmupdateWebSpamFgdSettingRestrictFqDbver(o["restrict-fq-dbver"], d, "restrict_fq_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-fq-dbver"], "FmupdateWebSpamFgdSetting-RestrictFqDbver"); ok {
			if err = d.Set("restrict_fq_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_fq_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_fq_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_iots_dbver", flattenFmupdateWebSpamFgdSettingRestrictIotsDbver(o["restrict-iots-dbver"], d, "restrict_iots_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-iots-dbver"], "FmupdateWebSpamFgdSetting-RestrictIotsDbver"); ok {
			if err = d.Set("restrict_iots_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_iots_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_iots_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_wf_dbver", flattenFmupdateWebSpamFgdSettingRestrictWfDbver(o["restrict-wf-dbver"], d, "restrict_wf_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-wf-dbver"], "FmupdateWebSpamFgdSetting-RestrictWfDbver"); ok {
			if err = d.Set("restrict_wf_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_wf_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_wf_dbver: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_override", flattenFmupdateWebSpamFgdSettingServerOverride(o["server-override"], d, "server_override")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-override"], "FmupdateWebSpamFgdSetting-ServerOverride"); ok {
				if err = d.Set("server_override", vv); err != nil {
					return fmt.Errorf("Error reading server_override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_override"); ok {
			if err = d.Set("server_override", flattenFmupdateWebSpamFgdSettingServerOverride(o["server-override"], d, "server_override")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-override"], "FmupdateWebSpamFgdSetting-ServerOverride"); ok {
					if err = d.Set("server_override", vv); err != nil {
						return fmt.Errorf("Error reading server_override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_override: %v", err)
				}
			}
		}
	}

	if err = d.Set("stat_log", flattenFmupdateWebSpamFgdSettingStatLog(o["stat-log"], d, "stat_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-log"], "FmupdateWebSpamFgdSetting-StatLog"); ok {
			if err = d.Set("stat_log", vv); err != nil {
				return fmt.Errorf("Error reading stat_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_log: %v", err)
		}
	}

	if err = d.Set("stat_log_interval", flattenFmupdateWebSpamFgdSettingStatLogInterval(o["stat-log-interval"], d, "stat_log_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-log-interval"], "FmupdateWebSpamFgdSetting-StatLogInterval"); ok {
			if err = d.Set("stat_log_interval", vv); err != nil {
				return fmt.Errorf("Error reading stat_log_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_log_interval: %v", err)
		}
	}

	if err = d.Set("stat_sync_interval", flattenFmupdateWebSpamFgdSettingStatSyncInterval(o["stat-sync-interval"], d, "stat_sync_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-sync-interval"], "FmupdateWebSpamFgdSetting-StatSyncInterval"); ok {
			if err = d.Set("stat_sync_interval", vv); err != nil {
				return fmt.Errorf("Error reading stat_sync_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_sync_interval: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenFmupdateWebSpamFgdSettingUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "FmupdateWebSpamFgdSetting-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("update_log", flattenFmupdateWebSpamFgdSettingUpdateLog(o["update-log"], d, "update_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-log"], "FmupdateWebSpamFgdSetting-UpdateLog"); ok {
			if err = d.Set("update_log", vv); err != nil {
				return fmt.Errorf("Error reading update_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_log: %v", err)
		}
	}

	if err = d.Set("wf_cache", flattenFmupdateWebSpamFgdSettingWfCache(o["wf-cache"], d, "wf_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-cache"], "FmupdateWebSpamFgdSetting-WfCache"); ok {
			if err = d.Set("wf_cache", vv); err != nil {
				return fmt.Errorf("Error reading wf_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_cache: %v", err)
		}
	}

	if err = d.Set("wf_dn_cache_expire_time", flattenFmupdateWebSpamFgdSettingWfDnCacheExpireTime(o["wf-dn-cache-expire-time"], d, "wf_dn_cache_expire_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-dn-cache-expire-time"], "FmupdateWebSpamFgdSetting-WfDnCacheExpireTime"); ok {
			if err = d.Set("wf_dn_cache_expire_time", vv); err != nil {
				return fmt.Errorf("Error reading wf_dn_cache_expire_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_dn_cache_expire_time: %v", err)
		}
	}

	if err = d.Set("wf_dn_cache_max_number", flattenFmupdateWebSpamFgdSettingWfDnCacheMaxNumber(o["wf-dn-cache-max-number"], d, "wf_dn_cache_max_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-dn-cache-max-number"], "FmupdateWebSpamFgdSetting-WfDnCacheMaxNumber"); ok {
			if err = d.Set("wf_dn_cache_max_number", vv); err != nil {
				return fmt.Errorf("Error reading wf_dn_cache_max_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_dn_cache_max_number: %v", err)
		}
	}

	if err = d.Set("wf_log", flattenFmupdateWebSpamFgdSettingWfLog(o["wf-log"], d, "wf_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-log"], "FmupdateWebSpamFgdSetting-WfLog"); ok {
			if err = d.Set("wf_log", vv); err != nil {
				return fmt.Errorf("Error reading wf_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_log: %v", err)
		}
	}

	if err = d.Set("wf_preload", flattenFmupdateWebSpamFgdSettingWfPreload(o["wf-preload"], d, "wf_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-preload"], "FmupdateWebSpamFgdSetting-WfPreload"); ok {
			if err = d.Set("wf_preload", vv); err != nil {
				return fmt.Errorf("Error reading wf_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_preload: %v", err)
		}
	}

	return nil
}

func flattenFmupdateWebSpamFgdSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateWebSpamFgdSettingAsCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAsPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAvCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAvLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAvPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAv2Cache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAv2Log(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAv2Preload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingEventlogQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFgdPullInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFqCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFqLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFqPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotvPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingLinkdLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingMaxClientWorker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingMaxLogQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingMaxUnratedSite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAs1Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAs2Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAs4Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAvDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAv2Dbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictFqDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictIotsDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictWfDbver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideServlist(d, i["servlist"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["servlist"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateWebSpamFgdSettingServerOverrideStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-type"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(d, i["service_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingStatLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingStatLogInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingStatSyncInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingUpdateLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfDnCacheExpireTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfDnCacheMaxNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfPreload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateWebSpamFgdSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("as_cache"); ok || d.HasChange("as_cache") {
		t, err := expandFmupdateWebSpamFgdSettingAsCache(d, v, "as_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-cache"] = t
		}
	}

	if v, ok := d.GetOk("as_log"); ok || d.HasChange("as_log") {
		t, err := expandFmupdateWebSpamFgdSettingAsLog(d, v, "as_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-log"] = t
		}
	}

	if v, ok := d.GetOk("as_preload"); ok || d.HasChange("as_preload") {
		t, err := expandFmupdateWebSpamFgdSettingAsPreload(d, v, "as_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-preload"] = t
		}
	}

	if v, ok := d.GetOk("av_cache"); ok || d.HasChange("av_cache") {
		t, err := expandFmupdateWebSpamFgdSettingAvCache(d, v, "av_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-cache"] = t
		}
	}

	if v, ok := d.GetOk("av_log"); ok || d.HasChange("av_log") {
		t, err := expandFmupdateWebSpamFgdSettingAvLog(d, v, "av_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-log"] = t
		}
	}

	if v, ok := d.GetOk("av_preload"); ok || d.HasChange("av_preload") {
		t, err := expandFmupdateWebSpamFgdSettingAvPreload(d, v, "av_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-preload"] = t
		}
	}

	if v, ok := d.GetOk("av2_cache"); ok || d.HasChange("av2_cache") {
		t, err := expandFmupdateWebSpamFgdSettingAv2Cache(d, v, "av2_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-cache"] = t
		}
	}

	if v, ok := d.GetOk("av2_log"); ok || d.HasChange("av2_log") {
		t, err := expandFmupdateWebSpamFgdSettingAv2Log(d, v, "av2_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-log"] = t
		}
	}

	if v, ok := d.GetOk("av2_preload"); ok || d.HasChange("av2_preload") {
		t, err := expandFmupdateWebSpamFgdSettingAv2Preload(d, v, "av2_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-preload"] = t
		}
	}

	if v, ok := d.GetOk("eventlog_query"); ok || d.HasChange("eventlog_query") {
		t, err := expandFmupdateWebSpamFgdSettingEventlogQuery(d, v, "eventlog_query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventlog-query"] = t
		}
	}

	if v, ok := d.GetOk("fgd_pull_interval"); ok || d.HasChange("fgd_pull_interval") {
		t, err := expandFmupdateWebSpamFgdSettingFgdPullInterval(d, v, "fgd_pull_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-pull-interval"] = t
		}
	}

	if v, ok := d.GetOk("fq_cache"); ok || d.HasChange("fq_cache") {
		t, err := expandFmupdateWebSpamFgdSettingFqCache(d, v, "fq_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-cache"] = t
		}
	}

	if v, ok := d.GetOk("fq_log"); ok || d.HasChange("fq_log") {
		t, err := expandFmupdateWebSpamFgdSettingFqLog(d, v, "fq_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-log"] = t
		}
	}

	if v, ok := d.GetOk("fq_preload"); ok || d.HasChange("fq_preload") {
		t, err := expandFmupdateWebSpamFgdSettingFqPreload(d, v, "fq_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-preload"] = t
		}
	}

	if v, ok := d.GetOk("iot_cache"); ok || d.HasChange("iot_cache") {
		t, err := expandFmupdateWebSpamFgdSettingIotCache(d, v, "iot_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-cache"] = t
		}
	}

	if v, ok := d.GetOk("iot_log"); ok || d.HasChange("iot_log") {
		t, err := expandFmupdateWebSpamFgdSettingIotLog(d, v, "iot_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-log"] = t
		}
	}

	if v, ok := d.GetOk("iot_preload"); ok || d.HasChange("iot_preload") {
		t, err := expandFmupdateWebSpamFgdSettingIotPreload(d, v, "iot_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-preload"] = t
		}
	}

	if v, ok := d.GetOk("iotv_preload"); ok || d.HasChange("iotv_preload") {
		t, err := expandFmupdateWebSpamFgdSettingIotvPreload(d, v, "iotv_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iotv-preload"] = t
		}
	}

	if v, ok := d.GetOk("linkd_log"); ok || d.HasChange("linkd_log") {
		t, err := expandFmupdateWebSpamFgdSettingLinkdLog(d, v, "linkd_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkd-log"] = t
		}
	}

	if v, ok := d.GetOk("max_client_worker"); ok || d.HasChange("max_client_worker") {
		t, err := expandFmupdateWebSpamFgdSettingMaxClientWorker(d, v, "max_client_worker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-client-worker"] = t
		}
	}

	if v, ok := d.GetOk("max_log_quota"); ok || d.HasChange("max_log_quota") {
		t, err := expandFmupdateWebSpamFgdSettingMaxLogQuota(d, v, "max_log_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-quota"] = t
		}
	}

	if v, ok := d.GetOk("max_unrated_site"); ok || d.HasChange("max_unrated_site") {
		t, err := expandFmupdateWebSpamFgdSettingMaxUnratedSite(d, v, "max_unrated_site")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-unrated-site"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as1_dbver"); ok || d.HasChange("restrict_as1_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAs1Dbver(d, v, "restrict_as1_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as1-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as2_dbver"); ok || d.HasChange("restrict_as2_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAs2Dbver(d, v, "restrict_as2_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as2-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as4_dbver"); ok || d.HasChange("restrict_as4_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAs4Dbver(d, v, "restrict_as4_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as4-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_av_dbver"); ok || d.HasChange("restrict_av_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAvDbver(d, v, "restrict_av_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-av-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_av2_dbver"); ok || d.HasChange("restrict_av2_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAv2Dbver(d, v, "restrict_av2_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-av2-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_fq_dbver"); ok || d.HasChange("restrict_fq_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictFqDbver(d, v, "restrict_fq_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-fq-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_iots_dbver"); ok || d.HasChange("restrict_iots_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictIotsDbver(d, v, "restrict_iots_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-iots-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_wf_dbver"); ok || d.HasChange("restrict_wf_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictWfDbver(d, v, "restrict_wf_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-wf-dbver"] = t
		}
	}

	if v, ok := d.GetOk("server_override"); ok || d.HasChange("server_override") {
		t, err := expandFmupdateWebSpamFgdSettingServerOverride(d, v, "server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-override"] = t
		}
	}

	if v, ok := d.GetOk("stat_log"); ok || d.HasChange("stat_log") {
		t, err := expandFmupdateWebSpamFgdSettingStatLog(d, v, "stat_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-log"] = t
		}
	}

	if v, ok := d.GetOk("stat_log_interval"); ok || d.HasChange("stat_log_interval") {
		t, err := expandFmupdateWebSpamFgdSettingStatLogInterval(d, v, "stat_log_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-log-interval"] = t
		}
	}

	if v, ok := d.GetOk("stat_sync_interval"); ok || d.HasChange("stat_sync_interval") {
		t, err := expandFmupdateWebSpamFgdSettingStatSyncInterval(d, v, "stat_sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-sync-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandFmupdateWebSpamFgdSettingUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_log"); ok || d.HasChange("update_log") {
		t, err := expandFmupdateWebSpamFgdSettingUpdateLog(d, v, "update_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-log"] = t
		}
	}

	if v, ok := d.GetOk("wf_cache"); ok || d.HasChange("wf_cache") {
		t, err := expandFmupdateWebSpamFgdSettingWfCache(d, v, "wf_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-cache"] = t
		}
	}

	if v, ok := d.GetOk("wf_dn_cache_expire_time"); ok || d.HasChange("wf_dn_cache_expire_time") {
		t, err := expandFmupdateWebSpamFgdSettingWfDnCacheExpireTime(d, v, "wf_dn_cache_expire_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-dn-cache-expire-time"] = t
		}
	}

	if v, ok := d.GetOk("wf_dn_cache_max_number"); ok || d.HasChange("wf_dn_cache_max_number") {
		t, err := expandFmupdateWebSpamFgdSettingWfDnCacheMaxNumber(d, v, "wf_dn_cache_max_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-dn-cache-max-number"] = t
		}
	}

	if v, ok := d.GetOk("wf_log"); ok || d.HasChange("wf_log") {
		t, err := expandFmupdateWebSpamFgdSettingWfLog(d, v, "wf_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-log"] = t
		}
	}

	if v, ok := d.GetOk("wf_preload"); ok || d.HasChange("wf_preload") {
		t, err := expandFmupdateWebSpamFgdSettingWfPreload(d, v, "wf_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-preload"] = t
		}
	}

	return &obj, nil
}
