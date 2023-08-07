// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemSdwanHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanHealthCheckCreate,
		Read:   resourceWantempSystemSdwanHealthCheckRead,
		Update: resourceWantempSystemSdwanHealthCheckUpdate,
		Delete: resourceWantempSystemSdwanHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"wanprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_dynamic_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"detect_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_match_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_request_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"embed_measured_health": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ftp_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ftp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mos_codec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"packet_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"probe_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"probe_packets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quality_measured_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recoverytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sla": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"jitter_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"latency_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"link_cost_factor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mos_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"packetloss_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_in_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_out_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"sla_fail_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sla_id_redistribute": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sla_pass_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold_alert_jitter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_alert_latency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_alert_packetloss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_warning_jitter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_warning_latency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"threshold_warning_packetloss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"update_cascade_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_static_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWantempSystemSdwanHealthCheckCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemSdwanHealthCheck(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanHealthCheck resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemSdwanHealthCheck(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanHealthCheck resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWantempSystemSdwanHealthCheckRead(d, m)
}

func resourceWantempSystemSdwanHealthCheckUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemSdwanHealthCheck(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanHealthCheck resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemSdwanHealthCheck(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanHealthCheck resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWantempSystemSdwanHealthCheckRead(d, m)
}

func resourceWantempSystemSdwanHealthCheckDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	err = c.DeleteWantempSystemSdwanHealthCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanHealthCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanHealthCheckRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	if wanprof == "" {
		wanprof = importOptionChecking(m.(*FortiClient).Cfg, "wanprof")
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	paradict["wanprof"] = wanprof

	o, err := c.ReadWantempSystemSdwanHealthCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanHealthCheck resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanHealthCheck(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanHealthCheck resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanHealthCheckDynamicServerWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckAddrModeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckClassIdWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDetectModeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDiffservcodeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDnsMatchIpWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDnsRequestDomainWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFailtimeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFtpFileWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFtpModeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHaPriorityWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHttpAgentWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWantempSystemSdwanHealthCheckHttpGetWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHttpMatchWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckIntervalWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckMembersWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckMosCodecWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckNameWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckPacketSizeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckPasswordWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckPortWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbeCountWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbePacketsWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbeTimeoutWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProtocolWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckQualityMeasuredMethodWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckRecoverytimeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSecurityModeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckServerWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSlaWssha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckSlaIdWssha(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaJitterThresholdWssha(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaLatencyThresholdWssha(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaLinkCostFactorWssha(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaMosThresholdWssha(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssha(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPriorityInSlaWssha(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssha(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PriorityOutSla")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckSlaIdWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaJitterThresholdWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLatencyThresholdWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLinkCostFactorWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSlaMosThresholdWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityInSlaWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaFailLogPeriodWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaIdRedistributeWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPassLogPeriodWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSourceWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSource6Wssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSystemDnsWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertJitterWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertLatencyWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningJitterWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningLatencyWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUpdateStaticRouteWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUserWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckVrfWssha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwanHealthCheck(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_dynamic_server", flattenWantempSystemSdwanHealthCheckDynamicServerWssha(o["_dynamic-server"], d, "_dynamic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dynamic-server"], "WantempSystemSdwanHealthCheck-DynamicServer"); ok {
			if err = d.Set("_dynamic_server", vv); err != nil {
				return fmt.Errorf("Error reading _dynamic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dynamic_server: %v", err)
		}
	}

	if err = d.Set("addr_mode", flattenWantempSystemSdwanHealthCheckAddrModeWssha(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "WantempSystemSdwanHealthCheck-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("class_id", flattenWantempSystemSdwanHealthCheckClassIdWssha(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "WantempSystemSdwanHealthCheck-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("detect_mode", flattenWantempSystemSdwanHealthCheckDetectModeWssha(o["detect-mode"], d, "detect_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-mode"], "WantempSystemSdwanHealthCheck-DetectMode"); ok {
			if err = d.Set("detect_mode", vv); err != nil {
				return fmt.Errorf("Error reading detect_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_mode: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenWantempSystemSdwanHealthCheckDiffservcodeWssha(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "WantempSystemSdwanHealthCheck-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dns_match_ip", flattenWantempSystemSdwanHealthCheckDnsMatchIpWssha(o["dns-match-ip"], d, "dns_match_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-match-ip"], "WantempSystemSdwanHealthCheck-DnsMatchIp"); ok {
			if err = d.Set("dns_match_ip", vv); err != nil {
				return fmt.Errorf("Error reading dns_match_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenWantempSystemSdwanHealthCheckDnsRequestDomainWssha(o["dns-request-domain"], d, "dns_request_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-request-domain"], "WantempSystemSdwanHealthCheck-DnsRequestDomain"); ok {
			if err = d.Set("dns_request_domain", vv); err != nil {
				return fmt.Errorf("Error reading dns_request_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("embed_measured_health", flattenWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssha(o["embed-measured-health"], d, "embed_measured_health")); err != nil {
		if vv, ok := fortiAPIPatch(o["embed-measured-health"], "WantempSystemSdwanHealthCheck-EmbedMeasuredHealth"); ok {
			if err = d.Set("embed_measured_health", vv); err != nil {
				return fmt.Errorf("Error reading embed_measured_health: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading embed_measured_health: %v", err)
		}
	}

	if err = d.Set("failtime", flattenWantempSystemSdwanHealthCheckFailtimeWssha(o["failtime"], d, "failtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["failtime"], "WantempSystemSdwanHealthCheck-Failtime"); ok {
			if err = d.Set("failtime", vv); err != nil {
				return fmt.Errorf("Error reading failtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("ftp_file", flattenWantempSystemSdwanHealthCheckFtpFileWssha(o["ftp-file"], d, "ftp_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-file"], "WantempSystemSdwanHealthCheck-FtpFile"); ok {
			if err = d.Set("ftp_file", vv); err != nil {
				return fmt.Errorf("Error reading ftp_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_file: %v", err)
		}
	}

	if err = d.Set("ftp_mode", flattenWantempSystemSdwanHealthCheckFtpModeWssha(o["ftp-mode"], d, "ftp_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-mode"], "WantempSystemSdwanHealthCheck-FtpMode"); ok {
			if err = d.Set("ftp_mode", vv); err != nil {
				return fmt.Errorf("Error reading ftp_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_mode: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenWantempSystemSdwanHealthCheckHaPriorityWssha(o["ha-priority"], d, "ha_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-priority"], "WantempSystemSdwanHealthCheck-HaPriority"); ok {
			if err = d.Set("ha_priority", vv); err != nil {
				return fmt.Errorf("Error reading ha_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenWantempSystemSdwanHealthCheckHttpAgentWssha(o["http-agent"], d, "http_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-agent"], "WantempSystemSdwanHealthCheck-HttpAgent"); ok {
			if err = d.Set("http_agent", vv); err != nil {
				return fmt.Errorf("Error reading http_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", flattenWantempSystemSdwanHealthCheckHttpGetWssha(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "WantempSystemSdwanHealthCheck-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenWantempSystemSdwanHealthCheckHttpMatchWssha(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "WantempSystemSdwanHealthCheck-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", flattenWantempSystemSdwanHealthCheckIntervalWssha(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "WantempSystemSdwanHealthCheck-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("members", flattenWantempSystemSdwanHealthCheckMembersWssha(o["members"], d, "members")); err != nil {
		if vv, ok := fortiAPIPatch(o["members"], "WantempSystemSdwanHealthCheck-Members"); ok {
			if err = d.Set("members", vv); err != nil {
				return fmt.Errorf("Error reading members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading members: %v", err)
		}
	}

	if err = d.Set("mos_codec", flattenWantempSystemSdwanHealthCheckMosCodecWssha(o["mos-codec"], d, "mos_codec")); err != nil {
		if vv, ok := fortiAPIPatch(o["mos-codec"], "WantempSystemSdwanHealthCheck-MosCodec"); ok {
			if err = d.Set("mos_codec", vv); err != nil {
				return fmt.Errorf("Error reading mos_codec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mos_codec: %v", err)
		}
	}

	if err = d.Set("name", flattenWantempSystemSdwanHealthCheckNameWssha(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WantempSystemSdwanHealthCheck-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenWantempSystemSdwanHealthCheckPacketSizeWssha(o["packet-size"], d, "packet_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-size"], "WantempSystemSdwanHealthCheck-PacketSize"); ok {
			if err = d.Set("packet_size", vv); err != nil {
				return fmt.Errorf("Error reading packet_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("password", flattenWantempSystemSdwanHealthCheckPasswordWssha(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "WantempSystemSdwanHealthCheck-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenWantempSystemSdwanHealthCheckPortWssha(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "WantempSystemSdwanHealthCheck-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("probe_count", flattenWantempSystemSdwanHealthCheckProbeCountWssha(o["probe-count"], d, "probe_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-count"], "WantempSystemSdwanHealthCheck-ProbeCount"); ok {
			if err = d.Set("probe_count", vv); err != nil {
				return fmt.Errorf("Error reading probe_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_count: %v", err)
		}
	}

	if err = d.Set("probe_packets", flattenWantempSystemSdwanHealthCheckProbePacketsWssha(o["probe-packets"], d, "probe_packets")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-packets"], "WantempSystemSdwanHealthCheck-ProbePackets"); ok {
			if err = d.Set("probe_packets", vv); err != nil {
				return fmt.Errorf("Error reading probe_packets: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_packets: %v", err)
		}
	}

	if err = d.Set("probe_timeout", flattenWantempSystemSdwanHealthCheckProbeTimeoutWssha(o["probe-timeout"], d, "probe_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-timeout"], "WantempSystemSdwanHealthCheck-ProbeTimeout"); ok {
			if err = d.Set("probe_timeout", vv); err != nil {
				return fmt.Errorf("Error reading probe_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWantempSystemSdwanHealthCheckProtocolWssha(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WantempSystemSdwanHealthCheck-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_measured_method", flattenWantempSystemSdwanHealthCheckQualityMeasuredMethodWssha(o["quality-measured-method"], d, "quality_measured_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["quality-measured-method"], "WantempSystemSdwanHealthCheck-QualityMeasuredMethod"); ok {
			if err = d.Set("quality_measured_method", vv); err != nil {
				return fmt.Errorf("Error reading quality_measured_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quality_measured_method: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenWantempSystemSdwanHealthCheckRecoverytimeWssha(o["recoverytime"], d, "recoverytime")); err != nil {
		if vv, ok := fortiAPIPatch(o["recoverytime"], "WantempSystemSdwanHealthCheck-Recoverytime"); ok {
			if err = d.Set("recoverytime", vv); err != nil {
				return fmt.Errorf("Error reading recoverytime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenWantempSystemSdwanHealthCheckSecurityModeWssha(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "WantempSystemSdwanHealthCheck-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", flattenWantempSystemSdwanHealthCheckServerWssha(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "WantempSystemSdwanHealthCheck-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenWantempSystemSdwanHealthCheckSlaWssha(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemSdwanHealthCheck-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenWantempSystemSdwanHealthCheckSlaWssha(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemSdwanHealthCheck-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_fail_log_period", flattenWantempSystemSdwanHealthCheckSlaFailLogPeriodWssha(o["sla-fail-log-period"], d, "sla_fail_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-fail-log-period"], "WantempSystemSdwanHealthCheck-SlaFailLogPeriod"); ok {
			if err = d.Set("sla_fail_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
		}
	}

	if err = d.Set("sla_id_redistribute", flattenWantempSystemSdwanHealthCheckSlaIdRedistributeWssha(o["sla-id-redistribute"], d, "sla_id_redistribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-id-redistribute"], "WantempSystemSdwanHealthCheck-SlaIdRedistribute"); ok {
			if err = d.Set("sla_id_redistribute", vv); err != nil {
				return fmt.Errorf("Error reading sla_id_redistribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_id_redistribute: %v", err)
		}
	}

	if err = d.Set("sla_pass_log_period", flattenWantempSystemSdwanHealthCheckSlaPassLogPeriodWssha(o["sla-pass-log-period"], d, "sla_pass_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-pass-log-period"], "WantempSystemSdwanHealthCheck-SlaPassLogPeriod"); ok {
			if err = d.Set("sla_pass_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
		}
	}

	if err = d.Set("source", flattenWantempSystemSdwanHealthCheckSourceWssha(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "WantempSystemSdwanHealthCheck-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenWantempSystemSdwanHealthCheckSource6Wssha(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "WantempSystemSdwanHealthCheck-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("system_dns", flattenWantempSystemSdwanHealthCheckSystemDnsWssha(o["system-dns"], d, "system_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-dns"], "WantempSystemSdwanHealthCheck-SystemDns"); ok {
			if err = d.Set("system_dns", vv); err != nil {
				return fmt.Errorf("Error reading system_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_dns: %v", err)
		}
	}

	if err = d.Set("threshold_alert_jitter", flattenWantempSystemSdwanHealthCheckThresholdAlertJitterWssha(o["threshold-alert-jitter"], d, "threshold_alert_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-jitter"], "WantempSystemSdwanHealthCheck-ThresholdAlertJitter"); ok {
			if err = d.Set("threshold_alert_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_alert_latency", flattenWantempSystemSdwanHealthCheckThresholdAlertLatencyWssha(o["threshold-alert-latency"], d, "threshold_alert_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-latency"], "WantempSystemSdwanHealthCheck-ThresholdAlertLatency"); ok {
			if err = d.Set("threshold_alert_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
		}
	}

	if err = d.Set("threshold_alert_packetloss", flattenWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssha(o["threshold-alert-packetloss"], d, "threshold_alert_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-packetloss"], "WantempSystemSdwanHealthCheck-ThresholdAlertPacketloss"); ok {
			if err = d.Set("threshold_alert_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
		}
	}

	if err = d.Set("threshold_warning_jitter", flattenWantempSystemSdwanHealthCheckThresholdWarningJitterWssha(o["threshold-warning-jitter"], d, "threshold_warning_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-jitter"], "WantempSystemSdwanHealthCheck-ThresholdWarningJitter"); ok {
			if err = d.Set("threshold_warning_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_warning_latency", flattenWantempSystemSdwanHealthCheckThresholdWarningLatencyWssha(o["threshold-warning-latency"], d, "threshold_warning_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-latency"], "WantempSystemSdwanHealthCheck-ThresholdWarningLatency"); ok {
			if err = d.Set("threshold_warning_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
		}
	}

	if err = d.Set("threshold_warning_packetloss", flattenWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssha(o["threshold-warning-packetloss"], d, "threshold_warning_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-packetloss"], "WantempSystemSdwanHealthCheck-ThresholdWarningPacketloss"); ok {
			if err = d.Set("threshold_warning_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssha(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-cascade-interface"], "WantempSystemSdwanHealthCheck-UpdateCascadeInterface"); ok {
			if err = d.Set("update_cascade_interface", vv); err != nil {
				return fmt.Errorf("Error reading update_cascade_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenWantempSystemSdwanHealthCheckUpdateStaticRouteWssha(o["update-static-route"], d, "update_static_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-static-route"], "WantempSystemSdwanHealthCheck-UpdateStaticRoute"); ok {
			if err = d.Set("update_static_route", vv); err != nil {
				return fmt.Errorf("Error reading update_static_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("user", flattenWantempSystemSdwanHealthCheckUserWssha(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "WantempSystemSdwanHealthCheck-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("vrf", flattenWantempSystemSdwanHealthCheckVrfWssha(o["vrf"], d, "vrf")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf"], "WantempSystemSdwanHealthCheck-Vrf"); ok {
			if err = d.Set("vrf", vv); err != nil {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanHealthCheckFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanHealthCheckDynamicServerWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckAddrModeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckClassIdWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDetectModeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDiffservcodeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDnsMatchIpWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDnsRequestDomainWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFailtimeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFtpFileWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFtpModeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHaPriorityWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpAgentWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpGetWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpMatchWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckIntervalWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckMembersWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckMosCodecWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckNameWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckPacketSizeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckPasswordWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckPortWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbeCountWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbePacketsWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbeTimeoutWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProtocolWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckQualityMeasuredMethodWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckRecoverytimeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSecurityModeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckServerWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSlaWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemSdwanHealthCheckSlaIdWssha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaJitterThresholdWssha(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaLatencyThresholdWssha(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemSdwanHealthCheckSlaLinkCostFactorWssha(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaMosThresholdWssha(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssha(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandWantempSystemSdwanHealthCheckSlaPriorityInSlaWssha(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssha(d, i["priority_out_sla"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckSlaIdWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaJitterThresholdWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLatencyThresholdWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLinkCostFactorWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSlaMosThresholdWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityInSlaWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaFailLogPeriodWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaIdRedistributeWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPassLogPeriodWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSourceWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSource6Wssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSystemDnsWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertJitterWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertLatencyWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningJitterWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningLatencyWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUpdateStaticRouteWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUserWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckVrfWssha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwanHealthCheck(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dynamic_server"); ok || d.HasChange("_dynamic_server") {
		t, err := expandWantempSystemSdwanHealthCheckDynamicServerWssha(d, v, "_dynamic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dynamic-server"] = t
		}
	}

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandWantempSystemSdwanHealthCheckAddrModeWssha(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandWantempSystemSdwanHealthCheckClassIdWssha(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("detect_mode"); ok || d.HasChange("detect_mode") {
		t, err := expandWantempSystemSdwanHealthCheckDetectModeWssha(d, v, "detect_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-mode"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandWantempSystemSdwanHealthCheckDiffservcodeWssha(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dns_match_ip"); ok || d.HasChange("dns_match_ip") {
		t, err := expandWantempSystemSdwanHealthCheckDnsMatchIpWssha(d, v, "dns_match_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok || d.HasChange("dns_request_domain") {
		t, err := expandWantempSystemSdwanHealthCheckDnsRequestDomainWssha(d, v, "dns_request_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("embed_measured_health"); ok || d.HasChange("embed_measured_health") {
		t, err := expandWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssha(d, v, "embed_measured_health")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embed-measured-health"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok || d.HasChange("failtime") {
		t, err := expandWantempSystemSdwanHealthCheckFailtimeWssha(d, v, "failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("ftp_file"); ok || d.HasChange("ftp_file") {
		t, err := expandWantempSystemSdwanHealthCheckFtpFileWssha(d, v, "ftp_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-file"] = t
		}
	}

	if v, ok := d.GetOk("ftp_mode"); ok || d.HasChange("ftp_mode") {
		t, err := expandWantempSystemSdwanHealthCheckFtpModeWssha(d, v, "ftp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-mode"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok || d.HasChange("ha_priority") {
		t, err := expandWantempSystemSdwanHealthCheckHaPriorityWssha(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok || d.HasChange("http_agent") {
		t, err := expandWantempSystemSdwanHealthCheckHttpAgentWssha(d, v, "http_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandWantempSystemSdwanHealthCheckHttpGetWssha(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandWantempSystemSdwanHealthCheckHttpMatchWssha(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandWantempSystemSdwanHealthCheckIntervalWssha(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandWantempSystemSdwanHealthCheckMembersWssha(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("mos_codec"); ok || d.HasChange("mos_codec") {
		t, err := expandWantempSystemSdwanHealthCheckMosCodecWssha(d, v, "mos_codec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mos-codec"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWantempSystemSdwanHealthCheckNameWssha(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok || d.HasChange("packet_size") {
		t, err := expandWantempSystemSdwanHealthCheckPacketSizeWssha(d, v, "packet_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandWantempSystemSdwanHealthCheckPasswordWssha(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandWantempSystemSdwanHealthCheckPortWssha(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("probe_count"); ok || d.HasChange("probe_count") {
		t, err := expandWantempSystemSdwanHealthCheckProbeCountWssha(d, v, "probe_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-count"] = t
		}
	}

	if v, ok := d.GetOk("probe_packets"); ok || d.HasChange("probe_packets") {
		t, err := expandWantempSystemSdwanHealthCheckProbePacketsWssha(d, v, "probe_packets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-packets"] = t
		}
	}

	if v, ok := d.GetOk("probe_timeout"); ok || d.HasChange("probe_timeout") {
		t, err := expandWantempSystemSdwanHealthCheckProbeTimeoutWssha(d, v, "probe_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-timeout"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWantempSystemSdwanHealthCheckProtocolWssha(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_measured_method"); ok || d.HasChange("quality_measured_method") {
		t, err := expandWantempSystemSdwanHealthCheckQualityMeasuredMethodWssha(d, v, "quality_measured_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-measured-method"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok || d.HasChange("recoverytime") {
		t, err := expandWantempSystemSdwanHealthCheckRecoverytimeWssha(d, v, "recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandWantempSystemSdwanHealthCheckSecurityModeWssha(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandWantempSystemSdwanHealthCheckServerWssha(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandWantempSystemSdwanHealthCheckSlaWssha(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_fail_log_period"); ok || d.HasChange("sla_fail_log_period") {
		t, err := expandWantempSystemSdwanHealthCheckSlaFailLogPeriodWssha(d, v, "sla_fail_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-fail-log-period"] = t
		}
	}

	if v, ok := d.GetOk("sla_id_redistribute"); ok || d.HasChange("sla_id_redistribute") {
		t, err := expandWantempSystemSdwanHealthCheckSlaIdRedistributeWssha(d, v, "sla_id_redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-id-redistribute"] = t
		}
	}

	if v, ok := d.GetOk("sla_pass_log_period"); ok || d.HasChange("sla_pass_log_period") {
		t, err := expandWantempSystemSdwanHealthCheckSlaPassLogPeriodWssha(d, v, "sla_pass_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-pass-log-period"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandWantempSystemSdwanHealthCheckSourceWssha(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok || d.HasChange("source6") {
		t, err := expandWantempSystemSdwanHealthCheckSource6Wssha(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("system_dns"); ok || d.HasChange("system_dns") {
		t, err := expandWantempSystemSdwanHealthCheckSystemDnsWssha(d, v, "system_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-dns"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_jitter"); ok || d.HasChange("threshold_alert_jitter") {
		t, err := expandWantempSystemSdwanHealthCheckThresholdAlertJitterWssha(d, v, "threshold_alert_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_latency"); ok || d.HasChange("threshold_alert_latency") {
		t, err := expandWantempSystemSdwanHealthCheckThresholdAlertLatencyWssha(d, v, "threshold_alert_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_packetloss"); ok || d.HasChange("threshold_alert_packetloss") {
		t, err := expandWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssha(d, v, "threshold_alert_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_jitter"); ok || d.HasChange("threshold_warning_jitter") {
		t, err := expandWantempSystemSdwanHealthCheckThresholdWarningJitterWssha(d, v, "threshold_warning_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_latency"); ok || d.HasChange("threshold_warning_latency") {
		t, err := expandWantempSystemSdwanHealthCheckThresholdWarningLatencyWssha(d, v, "threshold_warning_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_packetloss"); ok || d.HasChange("threshold_warning_packetloss") {
		t, err := expandWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssha(d, v, "threshold_warning_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok || d.HasChange("update_cascade_interface") {
		t, err := expandWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssha(d, v, "update_cascade_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok || d.HasChange("update_static_route") {
		t, err := expandWantempSystemSdwanHealthCheckUpdateStaticRouteWssha(d, v, "update_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandWantempSystemSdwanHealthCheckUserWssha(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		t, err := expandWantempSystemSdwanHealthCheckVrfWssha(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
