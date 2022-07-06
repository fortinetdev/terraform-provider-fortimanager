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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"linkd_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_client_worker": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"restrict_as2_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrict_as4_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrict_av_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrict_av2_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrict_fq_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrict_iots_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restrict_wf_dbver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
										Computed: true,
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

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateWebSpamFgdSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamFgdSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateWebSpamFgdSetting(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateWebSpamFgdSetting(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateWebSpamFgdSetting(adomv, mkey, nil)
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

func flattenFmupdateWebSpamFgdSettingAsCacheFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAsLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAsPreloadFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAvCacheFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAvLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAvPreloadFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAv2CacheFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAv2LogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingAv2PreloadFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingEventlogQueryFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFgdPullIntervalFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFqCacheFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFqLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingFqPreloadFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotCacheFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingIotPreloadFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingLinkdLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingMaxClientWorkerFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingMaxLogQuotaFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingMaxUnratedSiteFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAs1DbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAs2DbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAs4DbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAvDbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictAv2DbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictFqDbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictIotsDbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingRestrictWfDbverFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideFwfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := i["servlist"]; ok {
		result["servlist"] = flattenFmupdateWebSpamFgdSettingServerOverrideServlistFwfa(i["servlist"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFmupdateWebSpamFgdSettingServerOverrideStatusFwfa(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistFwfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfa(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfa(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfa(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := i["service-type"]; ok {
			v := flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfa(i["service-type"], d, pre_append)
			tmp["service_type"] = fortiAPISubPartPatch(v, "FmupdateWebSpamFgdSettingServerOverride-Servlist-ServiceType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingServerOverrideStatusFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingStatLogIntervalFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingStatSyncIntervalFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingUpdateIntervalFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingUpdateLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfCacheFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfDnCacheExpireTimeFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfDnCacheMaxNumberFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfLogFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamFgdSettingWfPreloadFwfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateWebSpamFgdSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("as_cache", flattenFmupdateWebSpamFgdSettingAsCacheFwfa(o["as-cache"], d, "as_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-cache"], "FmupdateWebSpamFgdSetting-AsCache"); ok {
			if err = d.Set("as_cache", vv); err != nil {
				return fmt.Errorf("Error reading as_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_cache: %v", err)
		}
	}

	if err = d.Set("as_log", flattenFmupdateWebSpamFgdSettingAsLogFwfa(o["as-log"], d, "as_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-log"], "FmupdateWebSpamFgdSetting-AsLog"); ok {
			if err = d.Set("as_log", vv); err != nil {
				return fmt.Errorf("Error reading as_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_log: %v", err)
		}
	}

	if err = d.Set("as_preload", flattenFmupdateWebSpamFgdSettingAsPreloadFwfa(o["as-preload"], d, "as_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["as-preload"], "FmupdateWebSpamFgdSetting-AsPreload"); ok {
			if err = d.Set("as_preload", vv); err != nil {
				return fmt.Errorf("Error reading as_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading as_preload: %v", err)
		}
	}

	if err = d.Set("av_cache", flattenFmupdateWebSpamFgdSettingAvCacheFwfa(o["av-cache"], d, "av_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-cache"], "FmupdateWebSpamFgdSetting-AvCache"); ok {
			if err = d.Set("av_cache", vv); err != nil {
				return fmt.Errorf("Error reading av_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_cache: %v", err)
		}
	}

	if err = d.Set("av_log", flattenFmupdateWebSpamFgdSettingAvLogFwfa(o["av-log"], d, "av_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-log"], "FmupdateWebSpamFgdSetting-AvLog"); ok {
			if err = d.Set("av_log", vv); err != nil {
				return fmt.Errorf("Error reading av_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_log: %v", err)
		}
	}

	if err = d.Set("av_preload", flattenFmupdateWebSpamFgdSettingAvPreloadFwfa(o["av-preload"], d, "av_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-preload"], "FmupdateWebSpamFgdSetting-AvPreload"); ok {
			if err = d.Set("av_preload", vv); err != nil {
				return fmt.Errorf("Error reading av_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_preload: %v", err)
		}
	}

	if err = d.Set("av2_cache", flattenFmupdateWebSpamFgdSettingAv2CacheFwfa(o["av2-cache"], d, "av2_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-cache"], "FmupdateWebSpamFgdSetting-Av2Cache"); ok {
			if err = d.Set("av2_cache", vv); err != nil {
				return fmt.Errorf("Error reading av2_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_cache: %v", err)
		}
	}

	if err = d.Set("av2_log", flattenFmupdateWebSpamFgdSettingAv2LogFwfa(o["av2-log"], d, "av2_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-log"], "FmupdateWebSpamFgdSetting-Av2Log"); ok {
			if err = d.Set("av2_log", vv); err != nil {
				return fmt.Errorf("Error reading av2_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_log: %v", err)
		}
	}

	if err = d.Set("av2_preload", flattenFmupdateWebSpamFgdSettingAv2PreloadFwfa(o["av2-preload"], d, "av2_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["av2-preload"], "FmupdateWebSpamFgdSetting-Av2Preload"); ok {
			if err = d.Set("av2_preload", vv); err != nil {
				return fmt.Errorf("Error reading av2_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av2_preload: %v", err)
		}
	}

	if err = d.Set("eventlog_query", flattenFmupdateWebSpamFgdSettingEventlogQueryFwfa(o["eventlog-query"], d, "eventlog_query")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventlog-query"], "FmupdateWebSpamFgdSetting-EventlogQuery"); ok {
			if err = d.Set("eventlog_query", vv); err != nil {
				return fmt.Errorf("Error reading eventlog_query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventlog_query: %v", err)
		}
	}

	if err = d.Set("fgd_pull_interval", flattenFmupdateWebSpamFgdSettingFgdPullIntervalFwfa(o["fgd-pull-interval"], d, "fgd_pull_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd-pull-interval"], "FmupdateWebSpamFgdSetting-FgdPullInterval"); ok {
			if err = d.Set("fgd_pull_interval", vv); err != nil {
				return fmt.Errorf("Error reading fgd_pull_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd_pull_interval: %v", err)
		}
	}

	if err = d.Set("fq_cache", flattenFmupdateWebSpamFgdSettingFqCacheFwfa(o["fq-cache"], d, "fq_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-cache"], "FmupdateWebSpamFgdSetting-FqCache"); ok {
			if err = d.Set("fq_cache", vv); err != nil {
				return fmt.Errorf("Error reading fq_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_cache: %v", err)
		}
	}

	if err = d.Set("fq_log", flattenFmupdateWebSpamFgdSettingFqLogFwfa(o["fq-log"], d, "fq_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-log"], "FmupdateWebSpamFgdSetting-FqLog"); ok {
			if err = d.Set("fq_log", vv); err != nil {
				return fmt.Errorf("Error reading fq_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_log: %v", err)
		}
	}

	if err = d.Set("fq_preload", flattenFmupdateWebSpamFgdSettingFqPreloadFwfa(o["fq-preload"], d, "fq_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["fq-preload"], "FmupdateWebSpamFgdSetting-FqPreload"); ok {
			if err = d.Set("fq_preload", vv); err != nil {
				return fmt.Errorf("Error reading fq_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fq_preload: %v", err)
		}
	}

	if err = d.Set("iot_cache", flattenFmupdateWebSpamFgdSettingIotCacheFwfa(o["iot-cache"], d, "iot_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-cache"], "FmupdateWebSpamFgdSetting-IotCache"); ok {
			if err = d.Set("iot_cache", vv); err != nil {
				return fmt.Errorf("Error reading iot_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_cache: %v", err)
		}
	}

	if err = d.Set("iot_log", flattenFmupdateWebSpamFgdSettingIotLogFwfa(o["iot-log"], d, "iot_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-log"], "FmupdateWebSpamFgdSetting-IotLog"); ok {
			if err = d.Set("iot_log", vv); err != nil {
				return fmt.Errorf("Error reading iot_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_log: %v", err)
		}
	}

	if err = d.Set("iot_preload", flattenFmupdateWebSpamFgdSettingIotPreloadFwfa(o["iot-preload"], d, "iot_preload")); err != nil {
		if vv, ok := fortiAPIPatch(o["iot-preload"], "FmupdateWebSpamFgdSetting-IotPreload"); ok {
			if err = d.Set("iot_preload", vv); err != nil {
				return fmt.Errorf("Error reading iot_preload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iot_preload: %v", err)
		}
	}

	if err = d.Set("linkd_log", flattenFmupdateWebSpamFgdSettingLinkdLogFwfa(o["linkd-log"], d, "linkd_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["linkd-log"], "FmupdateWebSpamFgdSetting-LinkdLog"); ok {
			if err = d.Set("linkd_log", vv); err != nil {
				return fmt.Errorf("Error reading linkd_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linkd_log: %v", err)
		}
	}

	if err = d.Set("max_client_worker", flattenFmupdateWebSpamFgdSettingMaxClientWorkerFwfa(o["max-client-worker"], d, "max_client_worker")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-client-worker"], "FmupdateWebSpamFgdSetting-MaxClientWorker"); ok {
			if err = d.Set("max_client_worker", vv); err != nil {
				return fmt.Errorf("Error reading max_client_worker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_client_worker: %v", err)
		}
	}

	if err = d.Set("max_log_quota", flattenFmupdateWebSpamFgdSettingMaxLogQuotaFwfa(o["max-log-quota"], d, "max_log_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-quota"], "FmupdateWebSpamFgdSetting-MaxLogQuota"); ok {
			if err = d.Set("max_log_quota", vv); err != nil {
				return fmt.Errorf("Error reading max_log_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_quota: %v", err)
		}
	}

	if err = d.Set("max_unrated_site", flattenFmupdateWebSpamFgdSettingMaxUnratedSiteFwfa(o["max-unrated-site"], d, "max_unrated_site")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-unrated-site"], "FmupdateWebSpamFgdSetting-MaxUnratedSite"); ok {
			if err = d.Set("max_unrated_site", vv); err != nil {
				return fmt.Errorf("Error reading max_unrated_site: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_unrated_site: %v", err)
		}
	}

	if err = d.Set("restrict_as1_dbver", flattenFmupdateWebSpamFgdSettingRestrictAs1DbverFwfa(o["restrict-as1-dbver"], d, "restrict_as1_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as1-dbver"], "FmupdateWebSpamFgdSetting-RestrictAs1Dbver"); ok {
			if err = d.Set("restrict_as1_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as1_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as1_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_as2_dbver", flattenFmupdateWebSpamFgdSettingRestrictAs2DbverFwfa(o["restrict-as2-dbver"], d, "restrict_as2_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as2-dbver"], "FmupdateWebSpamFgdSetting-RestrictAs2Dbver"); ok {
			if err = d.Set("restrict_as2_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as2_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as2_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_as4_dbver", flattenFmupdateWebSpamFgdSettingRestrictAs4DbverFwfa(o["restrict-as4-dbver"], d, "restrict_as4_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-as4-dbver"], "FmupdateWebSpamFgdSetting-RestrictAs4Dbver"); ok {
			if err = d.Set("restrict_as4_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_as4_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_as4_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_av_dbver", flattenFmupdateWebSpamFgdSettingRestrictAvDbverFwfa(o["restrict-av-dbver"], d, "restrict_av_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-av-dbver"], "FmupdateWebSpamFgdSetting-RestrictAvDbver"); ok {
			if err = d.Set("restrict_av_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_av_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_av_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_av2_dbver", flattenFmupdateWebSpamFgdSettingRestrictAv2DbverFwfa(o["restrict-av2-dbver"], d, "restrict_av2_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-av2-dbver"], "FmupdateWebSpamFgdSetting-RestrictAv2Dbver"); ok {
			if err = d.Set("restrict_av2_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_av2_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_av2_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_fq_dbver", flattenFmupdateWebSpamFgdSettingRestrictFqDbverFwfa(o["restrict-fq-dbver"], d, "restrict_fq_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-fq-dbver"], "FmupdateWebSpamFgdSetting-RestrictFqDbver"); ok {
			if err = d.Set("restrict_fq_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_fq_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_fq_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_iots_dbver", flattenFmupdateWebSpamFgdSettingRestrictIotsDbverFwfa(o["restrict-iots-dbver"], d, "restrict_iots_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-iots-dbver"], "FmupdateWebSpamFgdSetting-RestrictIotsDbver"); ok {
			if err = d.Set("restrict_iots_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_iots_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_iots_dbver: %v", err)
		}
	}

	if err = d.Set("restrict_wf_dbver", flattenFmupdateWebSpamFgdSettingRestrictWfDbverFwfa(o["restrict-wf-dbver"], d, "restrict_wf_dbver")); err != nil {
		if vv, ok := fortiAPIPatch(o["restrict-wf-dbver"], "FmupdateWebSpamFgdSetting-RestrictWfDbver"); ok {
			if err = d.Set("restrict_wf_dbver", vv); err != nil {
				return fmt.Errorf("Error reading restrict_wf_dbver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restrict_wf_dbver: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_override", flattenFmupdateWebSpamFgdSettingServerOverrideFwfa(o["server-override"], d, "server_override")); err != nil {
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
			if err = d.Set("server_override", flattenFmupdateWebSpamFgdSettingServerOverrideFwfa(o["server-override"], d, "server_override")); err != nil {
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

	if err = d.Set("stat_log_interval", flattenFmupdateWebSpamFgdSettingStatLogIntervalFwfa(o["stat-log-interval"], d, "stat_log_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-log-interval"], "FmupdateWebSpamFgdSetting-StatLogInterval"); ok {
			if err = d.Set("stat_log_interval", vv); err != nil {
				return fmt.Errorf("Error reading stat_log_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_log_interval: %v", err)
		}
	}

	if err = d.Set("stat_sync_interval", flattenFmupdateWebSpamFgdSettingStatSyncIntervalFwfa(o["stat-sync-interval"], d, "stat_sync_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stat-sync-interval"], "FmupdateWebSpamFgdSetting-StatSyncInterval"); ok {
			if err = d.Set("stat_sync_interval", vv); err != nil {
				return fmt.Errorf("Error reading stat_sync_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stat_sync_interval: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenFmupdateWebSpamFgdSettingUpdateIntervalFwfa(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "FmupdateWebSpamFgdSetting-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("update_log", flattenFmupdateWebSpamFgdSettingUpdateLogFwfa(o["update-log"], d, "update_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-log"], "FmupdateWebSpamFgdSetting-UpdateLog"); ok {
			if err = d.Set("update_log", vv); err != nil {
				return fmt.Errorf("Error reading update_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_log: %v", err)
		}
	}

	if err = d.Set("wf_cache", flattenFmupdateWebSpamFgdSettingWfCacheFwfa(o["wf-cache"], d, "wf_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-cache"], "FmupdateWebSpamFgdSetting-WfCache"); ok {
			if err = d.Set("wf_cache", vv); err != nil {
				return fmt.Errorf("Error reading wf_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_cache: %v", err)
		}
	}

	if err = d.Set("wf_dn_cache_expire_time", flattenFmupdateWebSpamFgdSettingWfDnCacheExpireTimeFwfa(o["wf-dn-cache-expire-time"], d, "wf_dn_cache_expire_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-dn-cache-expire-time"], "FmupdateWebSpamFgdSetting-WfDnCacheExpireTime"); ok {
			if err = d.Set("wf_dn_cache_expire_time", vv); err != nil {
				return fmt.Errorf("Error reading wf_dn_cache_expire_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_dn_cache_expire_time: %v", err)
		}
	}

	if err = d.Set("wf_dn_cache_max_number", flattenFmupdateWebSpamFgdSettingWfDnCacheMaxNumberFwfa(o["wf-dn-cache-max-number"], d, "wf_dn_cache_max_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-dn-cache-max-number"], "FmupdateWebSpamFgdSetting-WfDnCacheMaxNumber"); ok {
			if err = d.Set("wf_dn_cache_max_number", vv); err != nil {
				return fmt.Errorf("Error reading wf_dn_cache_max_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_dn_cache_max_number: %v", err)
		}
	}

	if err = d.Set("wf_log", flattenFmupdateWebSpamFgdSettingWfLogFwfa(o["wf-log"], d, "wf_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["wf-log"], "FmupdateWebSpamFgdSetting-WfLog"); ok {
			if err = d.Set("wf_log", vv); err != nil {
				return fmt.Errorf("Error reading wf_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wf_log: %v", err)
		}
	}

	if err = d.Set("wf_preload", flattenFmupdateWebSpamFgdSettingWfPreloadFwfa(o["wf-preload"], d, "wf_preload")); err != nil {
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

func expandFmupdateWebSpamFgdSettingAsCacheFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAsLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAsPreloadFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAvCacheFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAvLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAvPreloadFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAv2CacheFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAv2LogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingAv2PreloadFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingEventlogQueryFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFgdPullIntervalFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFqCacheFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFqLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingFqPreloadFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotCacheFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingIotPreloadFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingLinkdLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingMaxClientWorkerFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingMaxLogQuotaFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingMaxUnratedSiteFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAs1DbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAs2DbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAs4DbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAvDbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictAv2DbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictFqDbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictIotsDbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingRestrictWfDbverFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "servlist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["servlist"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistFwfa(d, i["servlist"], pre_append)
	} else {
		result["servlist"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFmupdateWebSpamFgdSettingServerOverrideStatusFwfa(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfa(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfa(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfa(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-type"], _ = expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfa(d, i["service_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIdFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIpFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistIp6Fwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistPortFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideServlistServiceTypeFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingServerOverrideStatusFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingStatLogIntervalFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingStatSyncIntervalFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingUpdateIntervalFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingUpdateLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfCacheFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfDnCacheExpireTimeFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfDnCacheMaxNumberFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfLogFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamFgdSettingWfPreloadFwfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateWebSpamFgdSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("as_cache"); ok || d.HasChange("as_cache") {
		t, err := expandFmupdateWebSpamFgdSettingAsCacheFwfa(d, v, "as_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-cache"] = t
		}
	}

	if v, ok := d.GetOk("as_log"); ok || d.HasChange("as_log") {
		t, err := expandFmupdateWebSpamFgdSettingAsLogFwfa(d, v, "as_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-log"] = t
		}
	}

	if v, ok := d.GetOk("as_preload"); ok || d.HasChange("as_preload") {
		t, err := expandFmupdateWebSpamFgdSettingAsPreloadFwfa(d, v, "as_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-preload"] = t
		}
	}

	if v, ok := d.GetOk("av_cache"); ok || d.HasChange("av_cache") {
		t, err := expandFmupdateWebSpamFgdSettingAvCacheFwfa(d, v, "av_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-cache"] = t
		}
	}

	if v, ok := d.GetOk("av_log"); ok || d.HasChange("av_log") {
		t, err := expandFmupdateWebSpamFgdSettingAvLogFwfa(d, v, "av_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-log"] = t
		}
	}

	if v, ok := d.GetOk("av_preload"); ok || d.HasChange("av_preload") {
		t, err := expandFmupdateWebSpamFgdSettingAvPreloadFwfa(d, v, "av_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-preload"] = t
		}
	}

	if v, ok := d.GetOk("av2_cache"); ok || d.HasChange("av2_cache") {
		t, err := expandFmupdateWebSpamFgdSettingAv2CacheFwfa(d, v, "av2_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-cache"] = t
		}
	}

	if v, ok := d.GetOk("av2_log"); ok || d.HasChange("av2_log") {
		t, err := expandFmupdateWebSpamFgdSettingAv2LogFwfa(d, v, "av2_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-log"] = t
		}
	}

	if v, ok := d.GetOk("av2_preload"); ok || d.HasChange("av2_preload") {
		t, err := expandFmupdateWebSpamFgdSettingAv2PreloadFwfa(d, v, "av2_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av2-preload"] = t
		}
	}

	if v, ok := d.GetOk("eventlog_query"); ok || d.HasChange("eventlog_query") {
		t, err := expandFmupdateWebSpamFgdSettingEventlogQueryFwfa(d, v, "eventlog_query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventlog-query"] = t
		}
	}

	if v, ok := d.GetOk("fgd_pull_interval"); ok || d.HasChange("fgd_pull_interval") {
		t, err := expandFmupdateWebSpamFgdSettingFgdPullIntervalFwfa(d, v, "fgd_pull_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-pull-interval"] = t
		}
	}

	if v, ok := d.GetOk("fq_cache"); ok || d.HasChange("fq_cache") {
		t, err := expandFmupdateWebSpamFgdSettingFqCacheFwfa(d, v, "fq_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-cache"] = t
		}
	}

	if v, ok := d.GetOk("fq_log"); ok || d.HasChange("fq_log") {
		t, err := expandFmupdateWebSpamFgdSettingFqLogFwfa(d, v, "fq_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-log"] = t
		}
	}

	if v, ok := d.GetOk("fq_preload"); ok || d.HasChange("fq_preload") {
		t, err := expandFmupdateWebSpamFgdSettingFqPreloadFwfa(d, v, "fq_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fq-preload"] = t
		}
	}

	if v, ok := d.GetOk("iot_cache"); ok || d.HasChange("iot_cache") {
		t, err := expandFmupdateWebSpamFgdSettingIotCacheFwfa(d, v, "iot_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-cache"] = t
		}
	}

	if v, ok := d.GetOk("iot_log"); ok || d.HasChange("iot_log") {
		t, err := expandFmupdateWebSpamFgdSettingIotLogFwfa(d, v, "iot_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-log"] = t
		}
	}

	if v, ok := d.GetOk("iot_preload"); ok || d.HasChange("iot_preload") {
		t, err := expandFmupdateWebSpamFgdSettingIotPreloadFwfa(d, v, "iot_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iot-preload"] = t
		}
	}

	if v, ok := d.GetOk("linkd_log"); ok || d.HasChange("linkd_log") {
		t, err := expandFmupdateWebSpamFgdSettingLinkdLogFwfa(d, v, "linkd_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkd-log"] = t
		}
	}

	if v, ok := d.GetOk("max_client_worker"); ok || d.HasChange("max_client_worker") {
		t, err := expandFmupdateWebSpamFgdSettingMaxClientWorkerFwfa(d, v, "max_client_worker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-client-worker"] = t
		}
	}

	if v, ok := d.GetOk("max_log_quota"); ok || d.HasChange("max_log_quota") {
		t, err := expandFmupdateWebSpamFgdSettingMaxLogQuotaFwfa(d, v, "max_log_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-quota"] = t
		}
	}

	if v, ok := d.GetOk("max_unrated_site"); ok || d.HasChange("max_unrated_site") {
		t, err := expandFmupdateWebSpamFgdSettingMaxUnratedSiteFwfa(d, v, "max_unrated_site")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-unrated-site"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as1_dbver"); ok || d.HasChange("restrict_as1_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAs1DbverFwfa(d, v, "restrict_as1_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as1-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as2_dbver"); ok || d.HasChange("restrict_as2_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAs2DbverFwfa(d, v, "restrict_as2_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as2-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_as4_dbver"); ok || d.HasChange("restrict_as4_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAs4DbverFwfa(d, v, "restrict_as4_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-as4-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_av_dbver"); ok || d.HasChange("restrict_av_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAvDbverFwfa(d, v, "restrict_av_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-av-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_av2_dbver"); ok || d.HasChange("restrict_av2_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictAv2DbverFwfa(d, v, "restrict_av2_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-av2-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_fq_dbver"); ok || d.HasChange("restrict_fq_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictFqDbverFwfa(d, v, "restrict_fq_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-fq-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_iots_dbver"); ok || d.HasChange("restrict_iots_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictIotsDbverFwfa(d, v, "restrict_iots_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-iots-dbver"] = t
		}
	}

	if v, ok := d.GetOk("restrict_wf_dbver"); ok || d.HasChange("restrict_wf_dbver") {
		t, err := expandFmupdateWebSpamFgdSettingRestrictWfDbverFwfa(d, v, "restrict_wf_dbver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restrict-wf-dbver"] = t
		}
	}

	if v, ok := d.GetOk("server_override"); ok || d.HasChange("server_override") {
		t, err := expandFmupdateWebSpamFgdSettingServerOverrideFwfa(d, v, "server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-override"] = t
		}
	}

	if v, ok := d.GetOk("stat_log_interval"); ok || d.HasChange("stat_log_interval") {
		t, err := expandFmupdateWebSpamFgdSettingStatLogIntervalFwfa(d, v, "stat_log_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-log-interval"] = t
		}
	}

	if v, ok := d.GetOk("stat_sync_interval"); ok || d.HasChange("stat_sync_interval") {
		t, err := expandFmupdateWebSpamFgdSettingStatSyncIntervalFwfa(d, v, "stat_sync_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stat-sync-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok || d.HasChange("update_interval") {
		t, err := expandFmupdateWebSpamFgdSettingUpdateIntervalFwfa(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("update_log"); ok || d.HasChange("update_log") {
		t, err := expandFmupdateWebSpamFgdSettingUpdateLogFwfa(d, v, "update_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-log"] = t
		}
	}

	if v, ok := d.GetOk("wf_cache"); ok || d.HasChange("wf_cache") {
		t, err := expandFmupdateWebSpamFgdSettingWfCacheFwfa(d, v, "wf_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-cache"] = t
		}
	}

	if v, ok := d.GetOk("wf_dn_cache_expire_time"); ok || d.HasChange("wf_dn_cache_expire_time") {
		t, err := expandFmupdateWebSpamFgdSettingWfDnCacheExpireTimeFwfa(d, v, "wf_dn_cache_expire_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-dn-cache-expire-time"] = t
		}
	}

	if v, ok := d.GetOk("wf_dn_cache_max_number"); ok || d.HasChange("wf_dn_cache_max_number") {
		t, err := expandFmupdateWebSpamFgdSettingWfDnCacheMaxNumberFwfa(d, v, "wf_dn_cache_max_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-dn-cache-max-number"] = t
		}
	}

	if v, ok := d.GetOk("wf_log"); ok || d.HasChange("wf_log") {
		t, err := expandFmupdateWebSpamFgdSettingWfLogFwfa(d, v, "wf_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-log"] = t
		}
	}

	if v, ok := d.GetOk("wf_preload"); ok || d.HasChange("wf_preload") {
		t, err := expandFmupdateWebSpamFgdSettingWfPreloadFwfa(d, v, "wf_preload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wf-preload"] = t
		}
	}

	return &obj, nil
}
