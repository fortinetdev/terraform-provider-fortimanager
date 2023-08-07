// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemSdwanService() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanServiceCreate,
		Read:   resourceWantempSystemSdwanServiceRead,
		Update: resourceWantempSystemSdwanServiceUpdate,
		Delete: resourceWantempSystemSdwanServiceDelete,

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
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_exclusive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bandwidth_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_forward_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp_reverse_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hash_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hold_down_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"input_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"input_device_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"input_zone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_app_ctrl": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"internet_service_app_ctrl_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"internet_service_app_ctrl_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"jitter_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"link_cost_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"link_cost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_sla_meet_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"packet_loss_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"passive_measurement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority_members": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"quality_link": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shortcut": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shortcut_stickiness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sla": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"sla_compare_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"standalone_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tie_break": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_shortcut_sla": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": &schema.Schema{
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

func resourceWantempSystemSdwanServiceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanService(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanService resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemSdwanService(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanService resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceWantempSystemSdwanServiceRead(d, m)
}

func resourceWantempSystemSdwanServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanService(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanService resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemSdwanService(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceWantempSystemSdwanServiceRead(d, m)
}

func resourceWantempSystemSdwanServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemSdwanService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemSdwanService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanService resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanServiceAddrModeWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceAgentExclusiveWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceBandwidthWeightWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDefaultWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpForwardWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpForwardTagWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpReverseWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpReverseTagWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDstWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDstNegateWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDst6Wsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceEndPortWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceGatewayWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceGroupsWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHashModeWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHealthCheckWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHoldDownTimeWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceIdWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputDeviceWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputDeviceNegateWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputZoneWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlGroupWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceCustomWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceCustomGroupWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceGroupWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceNameWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceJitterWeightWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLatencyWeightWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLinkCostFactorWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLinkCostThresholdWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceMinimumSlaMeetMembersWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceModeWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceNameWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePacketLossWeightWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePassiveMeasurementWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePriorityMembersWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePriorityZoneWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceProtocolWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceQualityLinkWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceRoleWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcutWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcutStickinessWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceRouteTagWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaWsssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenWantempSystemSdwanServiceSlaHealthCheckWsssa(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwanService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanServiceSlaIdWsssa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanService-Sla-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanServiceSlaHealthCheckWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaIdWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaCompareMethodWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrcWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrcNegateWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrc6Wsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStandaloneActionWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStartPortWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStatusWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTieBreakWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTosWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTosMaskWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceUseShortcutSlaWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceUsersWsssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwanService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenWantempSystemSdwanServiceAddrModeWsssa(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "WantempSystemSdwanService-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("agent_exclusive", flattenWantempSystemSdwanServiceAgentExclusiveWsssa(o["agent-exclusive"], d, "agent_exclusive")); err != nil {
		if vv, ok := fortiAPIPatch(o["agent-exclusive"], "WantempSystemSdwanService-AgentExclusive"); ok {
			if err = d.Set("agent_exclusive", vv); err != nil {
				return fmt.Errorf("Error reading agent_exclusive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agent_exclusive: %v", err)
		}
	}

	if err = d.Set("bandwidth_weight", flattenWantempSystemSdwanServiceBandwidthWeightWsssa(o["bandwidth-weight"], d, "bandwidth_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-weight"], "WantempSystemSdwanService-BandwidthWeight"); ok {
			if err = d.Set("bandwidth_weight", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_weight: %v", err)
		}
	}

	if err = d.Set("default", flattenWantempSystemSdwanServiceDefaultWsssa(o["default"], d, "default")); err != nil {
		if vv, ok := fortiAPIPatch(o["default"], "WantempSystemSdwanService-Default"); ok {
			if err = d.Set("default", vv); err != nil {
				return fmt.Errorf("Error reading default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("dscp_forward", flattenWantempSystemSdwanServiceDscpForwardWsssa(o["dscp-forward"], d, "dscp_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-forward"], "WantempSystemSdwanService-DscpForward"); ok {
			if err = d.Set("dscp_forward", vv); err != nil {
				return fmt.Errorf("Error reading dscp_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_forward: %v", err)
		}
	}

	if err = d.Set("dscp_forward_tag", flattenWantempSystemSdwanServiceDscpForwardTagWsssa(o["dscp-forward-tag"], d, "dscp_forward_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-forward-tag"], "WantempSystemSdwanService-DscpForwardTag"); ok {
			if err = d.Set("dscp_forward_tag", vv); err != nil {
				return fmt.Errorf("Error reading dscp_forward_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_forward_tag: %v", err)
		}
	}

	if err = d.Set("dscp_reverse", flattenWantempSystemSdwanServiceDscpReverseWsssa(o["dscp-reverse"], d, "dscp_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-reverse"], "WantempSystemSdwanService-DscpReverse"); ok {
			if err = d.Set("dscp_reverse", vv); err != nil {
				return fmt.Errorf("Error reading dscp_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_reverse: %v", err)
		}
	}

	if err = d.Set("dscp_reverse_tag", flattenWantempSystemSdwanServiceDscpReverseTagWsssa(o["dscp-reverse-tag"], d, "dscp_reverse_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-reverse-tag"], "WantempSystemSdwanService-DscpReverseTag"); ok {
			if err = d.Set("dscp_reverse_tag", vv); err != nil {
				return fmt.Errorf("Error reading dscp_reverse_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_reverse_tag: %v", err)
		}
	}

	if err = d.Set("dst", flattenWantempSystemSdwanServiceDstWsssa(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "WantempSystemSdwanService-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dst_negate", flattenWantempSystemSdwanServiceDstNegateWsssa(o["dst-negate"], d, "dst_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-negate"], "WantempSystemSdwanService-DstNegate"); ok {
			if err = d.Set("dst_negate", vv); err != nil {
				return fmt.Errorf("Error reading dst_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("dst6", flattenWantempSystemSdwanServiceDst6Wsssa(o["dst6"], d, "dst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst6"], "WantempSystemSdwanService-Dst6"); ok {
			if err = d.Set("dst6", vv); err != nil {
				return fmt.Errorf("Error reading dst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst6: %v", err)
		}
	}

	if err = d.Set("end_port", flattenWantempSystemSdwanServiceEndPortWsssa(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "WantempSystemSdwanService-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWantempSystemSdwanServiceGatewayWsssa(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "WantempSystemSdwanService-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("groups", flattenWantempSystemSdwanServiceGroupsWsssa(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "WantempSystemSdwanService-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("hash_mode", flattenWantempSystemSdwanServiceHashModeWsssa(o["hash-mode"], d, "hash_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-mode"], "WantempSystemSdwanService-HashMode"); ok {
			if err = d.Set("hash_mode", vv); err != nil {
				return fmt.Errorf("Error reading hash_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_mode: %v", err)
		}
	}

	if err = d.Set("health_check", flattenWantempSystemSdwanServiceHealthCheckWsssa(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemSdwanService-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("hold_down_time", flattenWantempSystemSdwanServiceHoldDownTimeWsssa(o["hold-down-time"], d, "hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["hold-down-time"], "WantempSystemSdwanService-HoldDownTime"); ok {
			if err = d.Set("hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hold_down_time: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWantempSystemSdwanServiceIdWsssa(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WantempSystemSdwanService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("input_device", flattenWantempSystemSdwanServiceInputDeviceWsssa(o["input-device"], d, "input_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-device"], "WantempSystemSdwanService-InputDevice"); ok {
			if err = d.Set("input_device", vv); err != nil {
				return fmt.Errorf("Error reading input_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_device: %v", err)
		}
	}

	if err = d.Set("input_device_negate", flattenWantempSystemSdwanServiceInputDeviceNegateWsssa(o["input-device-negate"], d, "input_device_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-device-negate"], "WantempSystemSdwanService-InputDeviceNegate"); ok {
			if err = d.Set("input_device_negate", vv); err != nil {
				return fmt.Errorf("Error reading input_device_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_device_negate: %v", err)
		}
	}

	if err = d.Set("input_zone", flattenWantempSystemSdwanServiceInputZoneWsssa(o["input-zone"], d, "input_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-zone"], "WantempSystemSdwanService-InputZone"); ok {
			if err = d.Set("input_zone", vv); err != nil {
				return fmt.Errorf("Error reading input_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_zone: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenWantempSystemSdwanServiceInternetServiceWsssa(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "WantempSystemSdwanService-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl", flattenWantempSystemSdwanServiceInternetServiceAppCtrlWsssa(o["internet-service-app-ctrl"], d, "internet_service_app_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl"], "WantempSystemSdwanService-InternetServiceAppCtrl"); ok {
			if err = d.Set("internet_service_app_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_category", flattenWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWsssa(o["internet-service-app-ctrl-category"], d, "internet_service_app_ctrl_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl-category"], "WantempSystemSdwanService-InternetServiceAppCtrlCategory"); ok {
			if err = d.Set("internet_service_app_ctrl_category", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl_category: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_group", flattenWantempSystemSdwanServiceInternetServiceAppCtrlGroupWsssa(o["internet-service-app-ctrl-group"], d, "internet_service_app_ctrl_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl-group"], "WantempSystemSdwanService-InternetServiceAppCtrlGroup"); ok {
			if err = d.Set("internet_service_app_ctrl_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl_group: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenWantempSystemSdwanServiceInternetServiceCustomWsssa(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "WantempSystemSdwanService-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenWantempSystemSdwanServiceInternetServiceCustomGroupWsssa(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "WantempSystemSdwanService-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenWantempSystemSdwanServiceInternetServiceGroupWsssa(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "WantempSystemSdwanService-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenWantempSystemSdwanServiceInternetServiceNameWsssa(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "WantempSystemSdwanService-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("jitter_weight", flattenWantempSystemSdwanServiceJitterWeightWsssa(o["jitter-weight"], d, "jitter_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-weight"], "WantempSystemSdwanService-JitterWeight"); ok {
			if err = d.Set("jitter_weight", vv); err != nil {
				return fmt.Errorf("Error reading jitter_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_weight: %v", err)
		}
	}

	if err = d.Set("latency_weight", flattenWantempSystemSdwanServiceLatencyWeightWsssa(o["latency-weight"], d, "latency_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-weight"], "WantempSystemSdwanService-LatencyWeight"); ok {
			if err = d.Set("latency_weight", vv); err != nil {
				return fmt.Errorf("Error reading latency_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_weight: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenWantempSystemSdwanServiceLinkCostFactorWsssa(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "WantempSystemSdwanService-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("link_cost_threshold", flattenWantempSystemSdwanServiceLinkCostThresholdWsssa(o["link-cost-threshold"], d, "link_cost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-threshold"], "WantempSystemSdwanService-LinkCostThreshold"); ok {
			if err = d.Set("link_cost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_threshold: %v", err)
		}
	}

	if err = d.Set("minimum_sla_meet_members", flattenWantempSystemSdwanServiceMinimumSlaMeetMembersWsssa(o["minimum-sla-meet-members"], d, "minimum_sla_meet_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-sla-meet-members"], "WantempSystemSdwanService-MinimumSlaMeetMembers"); ok {
			if err = d.Set("minimum_sla_meet_members", vv); err != nil {
				return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
		}
	}

	if err = d.Set("mode", flattenWantempSystemSdwanServiceModeWsssa(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "WantempSystemSdwanService-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("name", flattenWantempSystemSdwanServiceNameWsssa(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WantempSystemSdwanService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_loss_weight", flattenWantempSystemSdwanServicePacketLossWeightWsssa(o["packet-loss-weight"], d, "packet_loss_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-weight"], "WantempSystemSdwanService-PacketLossWeight"); ok {
			if err = d.Set("packet_loss_weight", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_weight: %v", err)
		}
	}

	if err = d.Set("passive_measurement", flattenWantempSystemSdwanServicePassiveMeasurementWsssa(o["passive-measurement"], d, "passive_measurement")); err != nil {
		if vv, ok := fortiAPIPatch(o["passive-measurement"], "WantempSystemSdwanService-PassiveMeasurement"); ok {
			if err = d.Set("passive_measurement", vv); err != nil {
				return fmt.Errorf("Error reading passive_measurement: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passive_measurement: %v", err)
		}
	}

	if err = d.Set("priority_members", flattenWantempSystemSdwanServicePriorityMembersWsssa(o["priority-members"], d, "priority_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-members"], "WantempSystemSdwanService-PriorityMembers"); ok {
			if err = d.Set("priority_members", vv); err != nil {
				return fmt.Errorf("Error reading priority_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_members: %v", err)
		}
	}

	if err = d.Set("priority_zone", flattenWantempSystemSdwanServicePriorityZoneWsssa(o["priority-zone"], d, "priority_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-zone"], "WantempSystemSdwanService-PriorityZone"); ok {
			if err = d.Set("priority_zone", vv); err != nil {
				return fmt.Errorf("Error reading priority_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_zone: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWantempSystemSdwanServiceProtocolWsssa(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WantempSystemSdwanService-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_link", flattenWantempSystemSdwanServiceQualityLinkWsssa(o["quality-link"], d, "quality_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["quality-link"], "WantempSystemSdwanService-QualityLink"); ok {
			if err = d.Set("quality_link", vv); err != nil {
				return fmt.Errorf("Error reading quality_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quality_link: %v", err)
		}
	}

	if err = d.Set("role", flattenWantempSystemSdwanServiceRoleWsssa(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "WantempSystemSdwanService-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("shortcut", flattenWantempSystemSdwanServiceShortcutWsssa(o["shortcut"], d, "shortcut")); err != nil {
		if vv, ok := fortiAPIPatch(o["shortcut"], "WantempSystemSdwanService-Shortcut"); ok {
			if err = d.Set("shortcut", vv); err != nil {
				return fmt.Errorf("Error reading shortcut: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shortcut: %v", err)
		}
	}

	if err = d.Set("shortcut_stickiness", flattenWantempSystemSdwanServiceShortcutStickinessWsssa(o["shortcut-stickiness"], d, "shortcut_stickiness")); err != nil {
		if vv, ok := fortiAPIPatch(o["shortcut-stickiness"], "WantempSystemSdwanService-ShortcutStickiness"); ok {
			if err = d.Set("shortcut_stickiness", vv); err != nil {
				return fmt.Errorf("Error reading shortcut_stickiness: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shortcut_stickiness: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenWantempSystemSdwanServiceRouteTagWsssa(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "WantempSystemSdwanService-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenWantempSystemSdwanServiceSlaWsssa(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemSdwanService-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenWantempSystemSdwanServiceSlaWsssa(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemSdwanService-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_compare_method", flattenWantempSystemSdwanServiceSlaCompareMethodWsssa(o["sla-compare-method"], d, "sla_compare_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-compare-method"], "WantempSystemSdwanService-SlaCompareMethod"); ok {
			if err = d.Set("sla_compare_method", vv); err != nil {
				return fmt.Errorf("Error reading sla_compare_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_compare_method: %v", err)
		}
	}

	if err = d.Set("src", flattenWantempSystemSdwanServiceSrcWsssa(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "WantempSystemSdwanService-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("src_negate", flattenWantempSystemSdwanServiceSrcNegateWsssa(o["src-negate"], d, "src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-negate"], "WantempSystemSdwanService-SrcNegate"); ok {
			if err = d.Set("src_negate", vv); err != nil {
				return fmt.Errorf("Error reading src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_negate: %v", err)
		}
	}

	if err = d.Set("src6", flattenWantempSystemSdwanServiceSrc6Wsssa(o["src6"], d, "src6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src6"], "WantempSystemSdwanService-Src6"); ok {
			if err = d.Set("src6", vv); err != nil {
				return fmt.Errorf("Error reading src6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src6: %v", err)
		}
	}

	if err = d.Set("standalone_action", flattenWantempSystemSdwanServiceStandaloneActionWsssa(o["standalone-action"], d, "standalone_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["standalone-action"], "WantempSystemSdwanService-StandaloneAction"); ok {
			if err = d.Set("standalone_action", vv); err != nil {
				return fmt.Errorf("Error reading standalone_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading standalone_action: %v", err)
		}
	}

	if err = d.Set("start_port", flattenWantempSystemSdwanServiceStartPortWsssa(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "WantempSystemSdwanService-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemSdwanServiceStatusWsssa(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemSdwanService-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tie_break", flattenWantempSystemSdwanServiceTieBreakWsssa(o["tie-break"], d, "tie_break")); err != nil {
		if vv, ok := fortiAPIPatch(o["tie-break"], "WantempSystemSdwanService-TieBreak"); ok {
			if err = d.Set("tie_break", vv); err != nil {
				return fmt.Errorf("Error reading tie_break: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tie_break: %v", err)
		}
	}

	if err = d.Set("tos", flattenWantempSystemSdwanServiceTosWsssa(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "WantempSystemSdwanService-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenWantempSystemSdwanServiceTosMaskWsssa(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "WantempSystemSdwanService-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("use_shortcut_sla", flattenWantempSystemSdwanServiceUseShortcutSlaWsssa(o["use-shortcut-sla"], d, "use_shortcut_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-shortcut-sla"], "WantempSystemSdwanService-UseShortcutSla"); ok {
			if err = d.Set("use_shortcut_sla", vv); err != nil {
				return fmt.Errorf("Error reading use_shortcut_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_shortcut_sla: %v", err)
		}
	}

	if err = d.Set("users", flattenWantempSystemSdwanServiceUsersWsssa(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "WantempSystemSdwanService-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanServiceAddrModeWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceAgentExclusiveWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceBandwidthWeightWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDefaultWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpForwardWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpForwardTagWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpReverseWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpReverseTagWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDstWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDstNegateWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDst6Wsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceEndPortWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceGatewayWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceGroupsWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHashModeWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHealthCheckWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHoldDownTimeWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceIdWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputDeviceWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputDeviceNegateWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputZoneWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlGroupWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceCustomWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceCustomGroupWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceGroupWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceNameWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceJitterWeightWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLatencyWeightWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLinkCostFactorWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLinkCostThresholdWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceMinimumSlaMeetMembersWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceModeWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceNameWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePacketLossWeightWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePassiveMeasurementWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePriorityMembersWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePriorityZoneWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceProtocolWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceQualityLinkWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceRoleWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcutWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcutStickinessWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceRouteTagWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandWantempSystemSdwanServiceSlaHealthCheckWsssa(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanServiceSlaIdWsssa(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanServiceSlaHealthCheckWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaIdWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaCompareMethodWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrcWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrcNegateWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrc6Wsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStandaloneActionWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStartPortWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStatusWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTieBreakWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTosWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTosMaskWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceUseShortcutSlaWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceUsersWsssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwanService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandWantempSystemSdwanServiceAddrModeWsssa(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("agent_exclusive"); ok || d.HasChange("agent_exclusive") {
		t, err := expandWantempSystemSdwanServiceAgentExclusiveWsssa(d, v, "agent_exclusive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-exclusive"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_weight"); ok || d.HasChange("bandwidth_weight") {
		t, err := expandWantempSystemSdwanServiceBandwidthWeightWsssa(d, v, "bandwidth_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-weight"] = t
		}
	}

	if v, ok := d.GetOk("default"); ok || d.HasChange("default") {
		t, err := expandWantempSystemSdwanServiceDefaultWsssa(d, v, "default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward"); ok || d.HasChange("dscp_forward") {
		t, err := expandWantempSystemSdwanServiceDscpForwardWsssa(d, v, "dscp_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward_tag"); ok || d.HasChange("dscp_forward_tag") {
		t, err := expandWantempSystemSdwanServiceDscpForwardTagWsssa(d, v, "dscp_forward_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward-tag"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse"); ok || d.HasChange("dscp_reverse") {
		t, err := expandWantempSystemSdwanServiceDscpReverseWsssa(d, v, "dscp_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse_tag"); ok || d.HasChange("dscp_reverse_tag") {
		t, err := expandWantempSystemSdwanServiceDscpReverseTagWsssa(d, v, "dscp_reverse_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse-tag"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandWantempSystemSdwanServiceDstWsssa(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dst_negate"); ok || d.HasChange("dst_negate") {
		t, err := expandWantempSystemSdwanServiceDstNegateWsssa(d, v, "dst_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-negate"] = t
		}
	}

	if v, ok := d.GetOk("dst6"); ok || d.HasChange("dst6") {
		t, err := expandWantempSystemSdwanServiceDst6Wsssa(d, v, "dst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst6"] = t
		}
	}

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandWantempSystemSdwanServiceEndPortWsssa(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandWantempSystemSdwanServiceGatewayWsssa(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandWantempSystemSdwanServiceGroupsWsssa(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("hash_mode"); ok || d.HasChange("hash_mode") {
		t, err := expandWantempSystemSdwanServiceHashModeWsssa(d, v, "hash_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-mode"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemSdwanServiceHealthCheckWsssa(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("hold_down_time"); ok || d.HasChange("hold_down_time") {
		t, err := expandWantempSystemSdwanServiceHoldDownTimeWsssa(d, v, "hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWantempSystemSdwanServiceIdWsssa(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok || d.HasChange("input_device") {
		t, err := expandWantempSystemSdwanServiceInputDeviceWsssa(d, v, "input_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	}

	if v, ok := d.GetOk("input_device_negate"); ok || d.HasChange("input_device_negate") {
		t, err := expandWantempSystemSdwanServiceInputDeviceNegateWsssa(d, v, "input_device_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device-negate"] = t
		}
	}

	if v, ok := d.GetOk("input_zone"); ok || d.HasChange("input_zone") {
		t, err := expandWantempSystemSdwanServiceInputZoneWsssa(d, v, "input_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-zone"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandWantempSystemSdwanServiceInternetServiceWsssa(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl"); ok || d.HasChange("internet_service_app_ctrl") {
		t, err := expandWantempSystemSdwanServiceInternetServiceAppCtrlWsssa(d, v, "internet_service_app_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_category"); ok || d.HasChange("internet_service_app_ctrl_category") {
		t, err := expandWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWsssa(d, v, "internet_service_app_ctrl_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-category"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_group"); ok || d.HasChange("internet_service_app_ctrl_group") {
		t, err := expandWantempSystemSdwanServiceInternetServiceAppCtrlGroupWsssa(d, v, "internet_service_app_ctrl_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandWantempSystemSdwanServiceInternetServiceCustomWsssa(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandWantempSystemSdwanServiceInternetServiceCustomGroupWsssa(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandWantempSystemSdwanServiceInternetServiceGroupWsssa(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandWantempSystemSdwanServiceInternetServiceNameWsssa(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("jitter_weight"); ok || d.HasChange("jitter_weight") {
		t, err := expandWantempSystemSdwanServiceJitterWeightWsssa(d, v, "jitter_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-weight"] = t
		}
	}

	if v, ok := d.GetOk("latency_weight"); ok || d.HasChange("latency_weight") {
		t, err := expandWantempSystemSdwanServiceLatencyWeightWsssa(d, v, "latency_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-weight"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandWantempSystemSdwanServiceLinkCostFactorWsssa(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_threshold"); ok || d.HasChange("link_cost_threshold") {
		t, err := expandWantempSystemSdwanServiceLinkCostThresholdWsssa(d, v, "link_cost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("minimum_sla_meet_members"); ok || d.HasChange("minimum_sla_meet_members") {
		t, err := expandWantempSystemSdwanServiceMinimumSlaMeetMembersWsssa(d, v, "minimum_sla_meet_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-sla-meet-members"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandWantempSystemSdwanServiceModeWsssa(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWantempSystemSdwanServiceNameWsssa(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_weight"); ok || d.HasChange("packet_loss_weight") {
		t, err := expandWantempSystemSdwanServicePacketLossWeightWsssa(d, v, "packet_loss_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-weight"] = t
		}
	}

	if v, ok := d.GetOk("passive_measurement"); ok || d.HasChange("passive_measurement") {
		t, err := expandWantempSystemSdwanServicePassiveMeasurementWsssa(d, v, "passive_measurement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-measurement"] = t
		}
	}

	if v, ok := d.GetOk("priority_members"); ok || d.HasChange("priority_members") {
		t, err := expandWantempSystemSdwanServicePriorityMembersWsssa(d, v, "priority_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-members"] = t
		}
	}

	if v, ok := d.GetOk("priority_zone"); ok || d.HasChange("priority_zone") {
		t, err := expandWantempSystemSdwanServicePriorityZoneWsssa(d, v, "priority_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-zone"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWantempSystemSdwanServiceProtocolWsssa(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_link"); ok || d.HasChange("quality_link") {
		t, err := expandWantempSystemSdwanServiceQualityLinkWsssa(d, v, "quality_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-link"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandWantempSystemSdwanServiceRoleWsssa(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("shortcut"); ok || d.HasChange("shortcut") {
		t, err := expandWantempSystemSdwanServiceShortcutWsssa(d, v, "shortcut")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortcut"] = t
		}
	}

	if v, ok := d.GetOk("shortcut_stickiness"); ok || d.HasChange("shortcut_stickiness") {
		t, err := expandWantempSystemSdwanServiceShortcutStickinessWsssa(d, v, "shortcut_stickiness")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortcut-stickiness"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandWantempSystemSdwanServiceRouteTagWsssa(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandWantempSystemSdwanServiceSlaWsssa(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_compare_method"); ok || d.HasChange("sla_compare_method") {
		t, err := expandWantempSystemSdwanServiceSlaCompareMethodWsssa(d, v, "sla_compare_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-compare-method"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandWantempSystemSdwanServiceSrcWsssa(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("src_negate"); ok || d.HasChange("src_negate") {
		t, err := expandWantempSystemSdwanServiceSrcNegateWsssa(d, v, "src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-negate"] = t
		}
	}

	if v, ok := d.GetOk("src6"); ok || d.HasChange("src6") {
		t, err := expandWantempSystemSdwanServiceSrc6Wsssa(d, v, "src6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src6"] = t
		}
	}

	if v, ok := d.GetOk("standalone_action"); ok || d.HasChange("standalone_action") {
		t, err := expandWantempSystemSdwanServiceStandaloneActionWsssa(d, v, "standalone_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-action"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandWantempSystemSdwanServiceStartPortWsssa(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemSdwanServiceStatusWsssa(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tie_break"); ok || d.HasChange("tie_break") {
		t, err := expandWantempSystemSdwanServiceTieBreakWsssa(d, v, "tie_break")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tie-break"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandWantempSystemSdwanServiceTosWsssa(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandWantempSystemSdwanServiceTosMaskWsssa(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("use_shortcut_sla"); ok || d.HasChange("use_shortcut_sla") {
		t, err := expandWantempSystemSdwanServiceUseShortcutSlaWsssa(d, v, "use_shortcut_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-shortcut-sla"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandWantempSystemSdwanServiceUsersWsssa(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	return &obj, nil
}
