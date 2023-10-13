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

func resourceWantempSystemVirtualWanLinkHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkHealthCheckCreate,
		Read:   resourceWantempSystemVirtualWanLinkHealthCheckRead,
		Update: resourceWantempSystemVirtualWanLinkHealthCheckUpdate,
		Delete: resourceWantempSystemVirtualWanLinkHealthCheckDelete,

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
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"members": &schema.Schema{
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
			"probe_packets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"probe_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"recoverytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
						"packetloss_threshold": &schema.Schema{
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
			"sla_pass_log_period": &schema.Schema{
				Type:     schema.TypeInt,
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
			},
			"update_static_route": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceWantempSystemVirtualWanLinkHealthCheckCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemVirtualWanLinkHealthCheck(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkHealthCheck resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemVirtualWanLinkHealthCheck(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkHealthCheck resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWantempSystemVirtualWanLinkHealthCheckRead(d, m)
}

func resourceWantempSystemVirtualWanLinkHealthCheckUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemVirtualWanLinkHealthCheck(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkHealthCheck resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLinkHealthCheck(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkHealthCheck resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWantempSystemVirtualWanLinkHealthCheckRead(d, m)
}

func resourceWantempSystemVirtualWanLinkHealthCheckDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemVirtualWanLinkHealthCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLinkHealthCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkHealthCheckRead(d *schema.ResourceData, m interface{}) error {
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
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	paradict["wanprof"] = wanprof

	o, err := c.ReadWantempSystemVirtualWanLinkHealthCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkHealthCheck resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLinkHealthCheck(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkHealthCheck resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkHealthCheckDynamicServerWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckAddrModeWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckFailtimeWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHaPriorityWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpAgentWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpGetWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpMatchWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckIntervalWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckMembersWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckNameWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckPacketSizeWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckPasswordWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckPortWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProbePacketsWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProtocolWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSecurityModeWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckServerWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaWsvha(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaIdWsvha(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsvha(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsvha(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsvha(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsvha(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-PacketlossThreshold")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaIdWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsvha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemVirtualWanLinkHealthCheck(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_dynamic_server", flattenWantempSystemVirtualWanLinkHealthCheckDynamicServerWsvha(o["_dynamic-server"], d, "_dynamic_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["_dynamic-server"], "WantempSystemVirtualWanLinkHealthCheck-DynamicServer"); ok {
			if err = d.Set("_dynamic_server", vv); err != nil {
				return fmt.Errorf("Error reading _dynamic_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _dynamic_server: %v", err)
		}
	}

	if err = d.Set("addr_mode", flattenWantempSystemVirtualWanLinkHealthCheckAddrModeWsvha(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "WantempSystemVirtualWanLinkHealthCheck-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsvha(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "WantempSystemVirtualWanLinkHealthCheck-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("failtime", flattenWantempSystemVirtualWanLinkHealthCheckFailtimeWsvha(o["failtime"], d, "failtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["failtime"], "WantempSystemVirtualWanLinkHealthCheck-Failtime"); ok {
			if err = d.Set("failtime", vv); err != nil {
				return fmt.Errorf("Error reading failtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenWantempSystemVirtualWanLinkHealthCheckHaPriorityWsvha(o["ha-priority"], d, "ha_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-priority"], "WantempSystemVirtualWanLinkHealthCheck-HaPriority"); ok {
			if err = d.Set("ha_priority", vv); err != nil {
				return fmt.Errorf("Error reading ha_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenWantempSystemVirtualWanLinkHealthCheckHttpAgentWsvha(o["http-agent"], d, "http_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-agent"], "WantempSystemVirtualWanLinkHealthCheck-HttpAgent"); ok {
			if err = d.Set("http_agent", vv); err != nil {
				return fmt.Errorf("Error reading http_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", flattenWantempSystemVirtualWanLinkHealthCheckHttpGetWsvha(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "WantempSystemVirtualWanLinkHealthCheck-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenWantempSystemVirtualWanLinkHealthCheckHttpMatchWsvha(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "WantempSystemVirtualWanLinkHealthCheck-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsvha(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "WantempSystemVirtualWanLinkHealthCheck-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("interval", flattenWantempSystemVirtualWanLinkHealthCheckIntervalWsvha(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "WantempSystemVirtualWanLinkHealthCheck-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("members", flattenWantempSystemVirtualWanLinkHealthCheckMembersWsvha(o["members"], d, "members")); err != nil {
		if vv, ok := fortiAPIPatch(o["members"], "WantempSystemVirtualWanLinkHealthCheck-Members"); ok {
			if err = d.Set("members", vv); err != nil {
				return fmt.Errorf("Error reading members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading members: %v", err)
		}
	}

	if err = d.Set("name", flattenWantempSystemVirtualWanLinkHealthCheckNameWsvha(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WantempSystemVirtualWanLinkHealthCheck-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenWantempSystemVirtualWanLinkHealthCheckPacketSizeWsvha(o["packet-size"], d, "packet_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-size"], "WantempSystemVirtualWanLinkHealthCheck-PacketSize"); ok {
			if err = d.Set("packet_size", vv); err != nil {
				return fmt.Errorf("Error reading packet_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("password", flattenWantempSystemVirtualWanLinkHealthCheckPasswordWsvha(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "WantempSystemVirtualWanLinkHealthCheck-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenWantempSystemVirtualWanLinkHealthCheckPortWsvha(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "WantempSystemVirtualWanLinkHealthCheck-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("probe_packets", flattenWantempSystemVirtualWanLinkHealthCheckProbePacketsWsvha(o["probe-packets"], d, "probe_packets")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-packets"], "WantempSystemVirtualWanLinkHealthCheck-ProbePackets"); ok {
			if err = d.Set("probe_packets", vv); err != nil {
				return fmt.Errorf("Error reading probe_packets: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_packets: %v", err)
		}
	}

	if err = d.Set("probe_timeout", flattenWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsvha(o["probe-timeout"], d, "probe_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-timeout"], "WantempSystemVirtualWanLinkHealthCheck-ProbeTimeout"); ok {
			if err = d.Set("probe_timeout", vv); err != nil {
				return fmt.Errorf("Error reading probe_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWantempSystemVirtualWanLinkHealthCheckProtocolWsvha(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WantempSystemVirtualWanLinkHealthCheck-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsvha(o["recoverytime"], d, "recoverytime")); err != nil {
		if vv, ok := fortiAPIPatch(o["recoverytime"], "WantempSystemVirtualWanLinkHealthCheck-Recoverytime"); ok {
			if err = d.Set("recoverytime", vv); err != nil {
				return fmt.Errorf("Error reading recoverytime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenWantempSystemVirtualWanLinkHealthCheckSecurityModeWsvha(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "WantempSystemVirtualWanLinkHealthCheck-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", flattenWantempSystemVirtualWanLinkHealthCheckServerWsvha(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "WantempSystemVirtualWanLinkHealthCheck-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenWantempSystemVirtualWanLinkHealthCheckSlaWsvha(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemVirtualWanLinkHealthCheck-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenWantempSystemVirtualWanLinkHealthCheckSlaWsvha(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemVirtualWanLinkHealthCheck-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_fail_log_period", flattenWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsvha(o["sla-fail-log-period"], d, "sla_fail_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-fail-log-period"], "WantempSystemVirtualWanLinkHealthCheck-SlaFailLogPeriod"); ok {
			if err = d.Set("sla_fail_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_fail_log_period: %v", err)
		}
	}

	if err = d.Set("sla_pass_log_period", flattenWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsvha(o["sla-pass-log-period"], d, "sla_pass_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-pass-log-period"], "WantempSystemVirtualWanLinkHealthCheck-SlaPassLogPeriod"); ok {
			if err = d.Set("sla_pass_log_period", vv); err != nil {
				return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_pass_log_period: %v", err)
		}
	}

	if err = d.Set("threshold_alert_jitter", flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsvha(o["threshold-alert-jitter"], d, "threshold_alert_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-jitter"], "WantempSystemVirtualWanLinkHealthCheck-ThresholdAlertJitter"); ok {
			if err = d.Set("threshold_alert_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_alert_latency", flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsvha(o["threshold-alert-latency"], d, "threshold_alert_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-latency"], "WantempSystemVirtualWanLinkHealthCheck-ThresholdAlertLatency"); ok {
			if err = d.Set("threshold_alert_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_latency: %v", err)
		}
	}

	if err = d.Set("threshold_alert_packetloss", flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsvha(o["threshold-alert-packetloss"], d, "threshold_alert_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-alert-packetloss"], "WantempSystemVirtualWanLinkHealthCheck-ThresholdAlertPacketloss"); ok {
			if err = d.Set("threshold_alert_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_alert_packetloss: %v", err)
		}
	}

	if err = d.Set("threshold_warning_jitter", flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsvha(o["threshold-warning-jitter"], d, "threshold_warning_jitter")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-jitter"], "WantempSystemVirtualWanLinkHealthCheck-ThresholdWarningJitter"); ok {
			if err = d.Set("threshold_warning_jitter", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_jitter: %v", err)
		}
	}

	if err = d.Set("threshold_warning_latency", flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsvha(o["threshold-warning-latency"], d, "threshold_warning_latency")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-latency"], "WantempSystemVirtualWanLinkHealthCheck-ThresholdWarningLatency"); ok {
			if err = d.Set("threshold_warning_latency", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_latency: %v", err)
		}
	}

	if err = d.Set("threshold_warning_packetloss", flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsvha(o["threshold-warning-packetloss"], d, "threshold_warning_packetloss")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-warning-packetloss"], "WantempSystemVirtualWanLinkHealthCheck-ThresholdWarningPacketloss"); ok {
			if err = d.Set("threshold_warning_packetloss", vv); err != nil {
				return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_warning_packetloss: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsvha(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-cascade-interface"], "WantempSystemVirtualWanLinkHealthCheck-UpdateCascadeInterface"); ok {
			if err = d.Set("update_cascade_interface", vv); err != nil {
				return fmt.Errorf("Error reading update_cascade_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsvha(o["update-static-route"], d, "update_static_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-static-route"], "WantempSystemVirtualWanLinkHealthCheck-UpdateStaticRoute"); ok {
			if err = d.Set("update_static_route", vv); err != nil {
				return fmt.Errorf("Error reading update_static_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkHealthCheckFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkHealthCheckDynamicServerWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckAddrModeWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckFailtimeWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHaPriorityWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpAgentWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpGetWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpMatchWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckIntervalWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckMembersWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckNameWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPacketSizeWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPasswordWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPortWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProbePacketsWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProtocolWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSecurityModeWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckServerWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaIdWsvha(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsvha(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsvha(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsvha(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsvha(d, i["packetloss_threshold"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaIdWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsvha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLinkHealthCheck(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_dynamic_server"); ok || d.HasChange("_dynamic_server") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckDynamicServerWsvha(d, v, "_dynamic_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_dynamic-server"] = t
		}
	}

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckAddrModeWsvha(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsvha(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok || d.HasChange("failtime") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckFailtimeWsvha(d, v, "failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok || d.HasChange("ha_priority") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckHaPriorityWsvha(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok || d.HasChange("http_agent") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckHttpAgentWsvha(d, v, "http_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckHttpGetWsvha(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckHttpMatchWsvha(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsvha(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckIntervalWsvha(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckMembersWsvha(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckNameWsvha(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok || d.HasChange("packet_size") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckPacketSizeWsvha(d, v, "packet_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckPasswordWsvha(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckPortWsvha(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("probe_packets"); ok || d.HasChange("probe_packets") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckProbePacketsWsvha(d, v, "probe_packets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-packets"] = t
		}
	}

	if v, ok := d.GetOk("probe_timeout"); ok || d.HasChange("probe_timeout") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsvha(d, v, "probe_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-timeout"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckProtocolWsvha(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok || d.HasChange("recoverytime") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsvha(d, v, "recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSecurityModeWsvha(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckServerWsvha(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaWsvha(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_fail_log_period"); ok || d.HasChange("sla_fail_log_period") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsvha(d, v, "sla_fail_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-fail-log-period"] = t
		}
	}

	if v, ok := d.GetOk("sla_pass_log_period"); ok || d.HasChange("sla_pass_log_period") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsvha(d, v, "sla_pass_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-pass-log-period"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_jitter"); ok || d.HasChange("threshold_alert_jitter") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsvha(d, v, "threshold_alert_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_latency"); ok || d.HasChange("threshold_alert_latency") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsvha(d, v, "threshold_alert_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_alert_packetloss"); ok || d.HasChange("threshold_alert_packetloss") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsvha(d, v, "threshold_alert_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-alert-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_jitter"); ok || d.HasChange("threshold_warning_jitter") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsvha(d, v, "threshold_warning_jitter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-jitter"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_latency"); ok || d.HasChange("threshold_warning_latency") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsvha(d, v, "threshold_warning_latency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-latency"] = t
		}
	}

	if v, ok := d.GetOk("threshold_warning_packetloss"); ok || d.HasChange("threshold_warning_packetloss") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsvha(d, v, "threshold_warning_packetloss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-warning-packetloss"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok || d.HasChange("update_cascade_interface") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsvha(d, v, "update_cascade_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok || d.HasChange("update_static_route") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsvha(d, v, "update_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	return &obj, nil
}
