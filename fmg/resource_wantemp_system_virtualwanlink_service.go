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

func resourceWantempSystemVirtualWanLinkService() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkServiceCreate,
		Read:   resourceWantempSystemVirtualWanLinkServiceRead,
		Update: resourceWantempSystemVirtualWanLinkServiceUpdate,
		Delete: resourceWantempSystemVirtualWanLinkServiceDelete,

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
				ForceNew: true,
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
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_ctrl": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"internet_service_ctrl_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_app_ctrl": &schema.Schema{
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
			"internet_service_id": &schema.Schema{
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
			"member": &schema.Schema{
				Type:     schema.TypeString,
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
			"priority_members": &schema.Schema{
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
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_mask": &schema.Schema{
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

func resourceWantempSystemVirtualWanLinkServiceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemVirtualWanLinkService(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkService resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemVirtualWanLinkService(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkService resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemVirtualWanLinkServiceRead(d, m)
}

func resourceWantempSystemVirtualWanLinkServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemVirtualWanLinkService(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkService resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLinkService(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemVirtualWanLinkServiceRead(d, m)
}

func resourceWantempSystemVirtualWanLinkServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemVirtualWanLinkService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLinkService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemVirtualWanLinkService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLinkService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkService resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkServiceAddrMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceBandwidthWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpForward2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpForwardTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpReverse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpReverseTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceDstNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDst62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceEndPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceHoldDownTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInputDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInputDeviceNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCustom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceJitterWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLatencyWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLinkCostFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLinkCostThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServicePacketLossWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServicePriorityMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceQualityLink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceRole2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceRouteTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSla2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkServiceSlaHealthCheck2edl(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSlaId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkService-Sla-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkServiceSlaHealthCheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaCompareMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceSrcNegate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrc62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceStandaloneAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStartPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceTosMask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceUsers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectWantempSystemVirtualWanLinkService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenWantempSystemVirtualWanLinkServiceAddrMode2edl(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "WantempSystemVirtualWanLinkService-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("bandwidth_weight", flattenWantempSystemVirtualWanLinkServiceBandwidthWeight2edl(o["bandwidth-weight"], d, "bandwidth_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-weight"], "WantempSystemVirtualWanLinkService-BandwidthWeight"); ok {
			if err = d.Set("bandwidth_weight", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_weight: %v", err)
		}
	}

	if err = d.Set("default", flattenWantempSystemVirtualWanLinkServiceDefault2edl(o["default"], d, "default")); err != nil {
		if vv, ok := fortiAPIPatch(o["default"], "WantempSystemVirtualWanLinkService-Default"); ok {
			if err = d.Set("default", vv); err != nil {
				return fmt.Errorf("Error reading default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("dscp_forward", flattenWantempSystemVirtualWanLinkServiceDscpForward2edl(o["dscp-forward"], d, "dscp_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-forward"], "WantempSystemVirtualWanLinkService-DscpForward"); ok {
			if err = d.Set("dscp_forward", vv); err != nil {
				return fmt.Errorf("Error reading dscp_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_forward: %v", err)
		}
	}

	if err = d.Set("dscp_forward_tag", flattenWantempSystemVirtualWanLinkServiceDscpForwardTag2edl(o["dscp-forward-tag"], d, "dscp_forward_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-forward-tag"], "WantempSystemVirtualWanLinkService-DscpForwardTag"); ok {
			if err = d.Set("dscp_forward_tag", vv); err != nil {
				return fmt.Errorf("Error reading dscp_forward_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_forward_tag: %v", err)
		}
	}

	if err = d.Set("dscp_reverse", flattenWantempSystemVirtualWanLinkServiceDscpReverse2edl(o["dscp-reverse"], d, "dscp_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-reverse"], "WantempSystemVirtualWanLinkService-DscpReverse"); ok {
			if err = d.Set("dscp_reverse", vv); err != nil {
				return fmt.Errorf("Error reading dscp_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_reverse: %v", err)
		}
	}

	if err = d.Set("dscp_reverse_tag", flattenWantempSystemVirtualWanLinkServiceDscpReverseTag2edl(o["dscp-reverse-tag"], d, "dscp_reverse_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-reverse-tag"], "WantempSystemVirtualWanLinkService-DscpReverseTag"); ok {
			if err = d.Set("dscp_reverse_tag", vv); err != nil {
				return fmt.Errorf("Error reading dscp_reverse_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_reverse_tag: %v", err)
		}
	}

	if err = d.Set("dst", flattenWantempSystemVirtualWanLinkServiceDst2edl(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "WantempSystemVirtualWanLinkService-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dst_negate", flattenWantempSystemVirtualWanLinkServiceDstNegate2edl(o["dst-negate"], d, "dst_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-negate"], "WantempSystemVirtualWanLinkService-DstNegate"); ok {
			if err = d.Set("dst_negate", vv); err != nil {
				return fmt.Errorf("Error reading dst_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("dst6", flattenWantempSystemVirtualWanLinkServiceDst62edl(o["dst6"], d, "dst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst6"], "WantempSystemVirtualWanLinkService-Dst6"); ok {
			if err = d.Set("dst6", vv); err != nil {
				return fmt.Errorf("Error reading dst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst6: %v", err)
		}
	}

	if err = d.Set("end_port", flattenWantempSystemVirtualWanLinkServiceEndPort2edl(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "WantempSystemVirtualWanLinkService-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWantempSystemVirtualWanLinkServiceGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "WantempSystemVirtualWanLinkService-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("groups", flattenWantempSystemVirtualWanLinkServiceGroups2edl(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "WantempSystemVirtualWanLinkService-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkServiceHealthCheck2edl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemVirtualWanLinkService-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("hold_down_time", flattenWantempSystemVirtualWanLinkServiceHoldDownTime2edl(o["hold-down-time"], d, "hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["hold-down-time"], "WantempSystemVirtualWanLinkService-HoldDownTime"); ok {
			if err = d.Set("hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hold_down_time: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWantempSystemVirtualWanLinkServiceId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WantempSystemVirtualWanLinkService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("input_device", flattenWantempSystemVirtualWanLinkServiceInputDevice2edl(o["input-device"], d, "input_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-device"], "WantempSystemVirtualWanLinkService-InputDevice"); ok {
			if err = d.Set("input_device", vv); err != nil {
				return fmt.Errorf("Error reading input_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_device: %v", err)
		}
	}

	if err = d.Set("input_device_negate", flattenWantempSystemVirtualWanLinkServiceInputDeviceNegate2edl(o["input-device-negate"], d, "input_device_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["input-device-negate"], "WantempSystemVirtualWanLinkService-InputDeviceNegate"); ok {
			if err = d.Set("input_device_negate", vv); err != nil {
				return fmt.Errorf("Error reading input_device_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input_device_negate: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenWantempSystemVirtualWanLinkServiceInternetService2edl(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "WantempSystemVirtualWanLinkService-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_ctrl", flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrl2edl(o["internet-service-ctrl"], d, "internet_service_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-ctrl"], "WantempSystemVirtualWanLinkService-InternetServiceCtrl"); ok {
			if err = d.Set("internet_service_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_ctrl: %v", err)
		}
	}

	if err = d.Set("internet_service_ctrl_group", flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup2edl(o["internet-service-ctrl-group"], d, "internet_service_ctrl_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-ctrl-group"], "WantempSystemVirtualWanLinkService-InternetServiceCtrlGroup"); ok {
			if err = d.Set("internet_service_ctrl_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_ctrl_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_ctrl_group: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl", flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl2edl(o["internet-service-app-ctrl"], d, "internet_service_app_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl"], "WantempSystemVirtualWanLinkService-InternetServiceAppCtrl"); ok {
			if err = d.Set("internet_service_app_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_group", flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup2edl(o["internet-service-app-ctrl-group"], d, "internet_service_app_ctrl_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-app-ctrl-group"], "WantempSystemVirtualWanLinkService-InternetServiceAppCtrlGroup"); ok {
			if err = d.Set("internet_service_app_ctrl_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_app_ctrl_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_app_ctrl_group: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenWantempSystemVirtualWanLinkServiceInternetServiceCustom2edl(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "WantempSystemVirtualWanLinkService-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup2edl(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "WantempSystemVirtualWanLinkService-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenWantempSystemVirtualWanLinkServiceInternetServiceGroup2edl(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "WantempSystemVirtualWanLinkService-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenWantempSystemVirtualWanLinkServiceInternetServiceId2edl(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "WantempSystemVirtualWanLinkService-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("jitter_weight", flattenWantempSystemVirtualWanLinkServiceJitterWeight2edl(o["jitter-weight"], d, "jitter_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-weight"], "WantempSystemVirtualWanLinkService-JitterWeight"); ok {
			if err = d.Set("jitter_weight", vv); err != nil {
				return fmt.Errorf("Error reading jitter_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_weight: %v", err)
		}
	}

	if err = d.Set("latency_weight", flattenWantempSystemVirtualWanLinkServiceLatencyWeight2edl(o["latency-weight"], d, "latency_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-weight"], "WantempSystemVirtualWanLinkService-LatencyWeight"); ok {
			if err = d.Set("latency_weight", vv); err != nil {
				return fmt.Errorf("Error reading latency_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_weight: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenWantempSystemVirtualWanLinkServiceLinkCostFactor2edl(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "WantempSystemVirtualWanLinkService-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("link_cost_threshold", flattenWantempSystemVirtualWanLinkServiceLinkCostThreshold2edl(o["link-cost-threshold"], d, "link_cost_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-threshold"], "WantempSystemVirtualWanLinkService-LinkCostThreshold"); ok {
			if err = d.Set("link_cost_threshold", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_threshold: %v", err)
		}
	}

	if err = d.Set("member", flattenWantempSystemVirtualWanLinkServiceMember2edl(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "WantempSystemVirtualWanLinkService-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("mode", flattenWantempSystemVirtualWanLinkServiceMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "WantempSystemVirtualWanLinkService-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("name", flattenWantempSystemVirtualWanLinkServiceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WantempSystemVirtualWanLinkService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_loss_weight", flattenWantempSystemVirtualWanLinkServicePacketLossWeight2edl(o["packet-loss-weight"], d, "packet_loss_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-loss-weight"], "WantempSystemVirtualWanLinkService-PacketLossWeight"); ok {
			if err = d.Set("packet_loss_weight", vv); err != nil {
				return fmt.Errorf("Error reading packet_loss_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_loss_weight: %v", err)
		}
	}

	if err = d.Set("priority_members", flattenWantempSystemVirtualWanLinkServicePriorityMembers2edl(o["priority-members"], d, "priority_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-members"], "WantempSystemVirtualWanLinkService-PriorityMembers"); ok {
			if err = d.Set("priority_members", vv); err != nil {
				return fmt.Errorf("Error reading priority_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_members: %v", err)
		}
	}

	if err = d.Set("protocol", flattenWantempSystemVirtualWanLinkServiceProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "WantempSystemVirtualWanLinkService-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_link", flattenWantempSystemVirtualWanLinkServiceQualityLink2edl(o["quality-link"], d, "quality_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["quality-link"], "WantempSystemVirtualWanLinkService-QualityLink"); ok {
			if err = d.Set("quality_link", vv); err != nil {
				return fmt.Errorf("Error reading quality_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quality_link: %v", err)
		}
	}

	if err = d.Set("role", flattenWantempSystemVirtualWanLinkServiceRole2edl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "WantempSystemVirtualWanLinkService-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenWantempSystemVirtualWanLinkServiceRouteTag2edl(o["route-tag"], d, "route_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-tag"], "WantempSystemVirtualWanLinkService-RouteTag"); ok {
			if err = d.Set("route_tag", vv); err != nil {
				return fmt.Errorf("Error reading route_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenWantempSystemVirtualWanLinkServiceSla2edl(o["sla"], d, "sla")); err != nil {
			if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemVirtualWanLinkService-Sla"); ok {
				if err = d.Set("sla", vv); err != nil {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenWantempSystemVirtualWanLinkServiceSla2edl(o["sla"], d, "sla")); err != nil {
				if vv, ok := fortiAPIPatch(o["sla"], "WantempSystemVirtualWanLinkService-Sla"); ok {
					if err = d.Set("sla", vv); err != nil {
						return fmt.Errorf("Error reading sla: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_compare_method", flattenWantempSystemVirtualWanLinkServiceSlaCompareMethod2edl(o["sla-compare-method"], d, "sla_compare_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-compare-method"], "WantempSystemVirtualWanLinkService-SlaCompareMethod"); ok {
			if err = d.Set("sla_compare_method", vv); err != nil {
				return fmt.Errorf("Error reading sla_compare_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_compare_method: %v", err)
		}
	}

	if err = d.Set("src", flattenWantempSystemVirtualWanLinkServiceSrc2edl(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "WantempSystemVirtualWanLinkService-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("src_negate", flattenWantempSystemVirtualWanLinkServiceSrcNegate2edl(o["src-negate"], d, "src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-negate"], "WantempSystemVirtualWanLinkService-SrcNegate"); ok {
			if err = d.Set("src_negate", vv); err != nil {
				return fmt.Errorf("Error reading src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_negate: %v", err)
		}
	}

	if err = d.Set("src6", flattenWantempSystemVirtualWanLinkServiceSrc62edl(o["src6"], d, "src6")); err != nil {
		if vv, ok := fortiAPIPatch(o["src6"], "WantempSystemVirtualWanLinkService-Src6"); ok {
			if err = d.Set("src6", vv); err != nil {
				return fmt.Errorf("Error reading src6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src6: %v", err)
		}
	}

	if err = d.Set("standalone_action", flattenWantempSystemVirtualWanLinkServiceStandaloneAction2edl(o["standalone-action"], d, "standalone_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["standalone-action"], "WantempSystemVirtualWanLinkService-StandaloneAction"); ok {
			if err = d.Set("standalone_action", vv); err != nil {
				return fmt.Errorf("Error reading standalone_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading standalone_action: %v", err)
		}
	}

	if err = d.Set("start_port", flattenWantempSystemVirtualWanLinkServiceStartPort2edl(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "WantempSystemVirtualWanLinkService-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemVirtualWanLinkServiceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemVirtualWanLinkService-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tos", flattenWantempSystemVirtualWanLinkServiceTos2edl(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "WantempSystemVirtualWanLinkService-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenWantempSystemVirtualWanLinkServiceTosMask2edl(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "WantempSystemVirtualWanLinkService-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("users", flattenWantempSystemVirtualWanLinkServiceUsers2edl(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "WantempSystemVirtualWanLinkService-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkServiceAddrMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceBandwidthWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpForward2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpForwardTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpReverse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpReverseTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceDstNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDst62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceEndPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceHoldDownTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInputDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInputDeviceNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCtrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCustom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceJitterWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLatencyWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLinkCostFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLinkCostThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServicePacketLossWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServicePriorityMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceQualityLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceRole2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceRouteTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSla2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkServiceSlaHealthCheck2edl(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemVirtualWanLinkServiceSlaId2edl(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaHealthCheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaCompareMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceSrcNegate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrc62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceStandaloneAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStartPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceTosMask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceUsers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectWantempSystemVirtualWanLinkService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandWantempSystemVirtualWanLinkServiceAddrMode2edl(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_weight"); ok || d.HasChange("bandwidth_weight") {
		t, err := expandWantempSystemVirtualWanLinkServiceBandwidthWeight2edl(d, v, "bandwidth_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-weight"] = t
		}
	}

	if v, ok := d.GetOk("default"); ok || d.HasChange("default") {
		t, err := expandWantempSystemVirtualWanLinkServiceDefault2edl(d, v, "default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward"); ok || d.HasChange("dscp_forward") {
		t, err := expandWantempSystemVirtualWanLinkServiceDscpForward2edl(d, v, "dscp_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward_tag"); ok || d.HasChange("dscp_forward_tag") {
		t, err := expandWantempSystemVirtualWanLinkServiceDscpForwardTag2edl(d, v, "dscp_forward_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward-tag"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse"); ok || d.HasChange("dscp_reverse") {
		t, err := expandWantempSystemVirtualWanLinkServiceDscpReverse2edl(d, v, "dscp_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse_tag"); ok || d.HasChange("dscp_reverse_tag") {
		t, err := expandWantempSystemVirtualWanLinkServiceDscpReverseTag2edl(d, v, "dscp_reverse_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse-tag"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandWantempSystemVirtualWanLinkServiceDst2edl(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dst_negate"); ok || d.HasChange("dst_negate") {
		t, err := expandWantempSystemVirtualWanLinkServiceDstNegate2edl(d, v, "dst_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-negate"] = t
		}
	}

	if v, ok := d.GetOk("dst6"); ok || d.HasChange("dst6") {
		t, err := expandWantempSystemVirtualWanLinkServiceDst62edl(d, v, "dst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst6"] = t
		}
	}

	if v, ok := d.GetOk("end_port"); ok || d.HasChange("end_port") {
		t, err := expandWantempSystemVirtualWanLinkServiceEndPort2edl(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandWantempSystemVirtualWanLinkServiceGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandWantempSystemVirtualWanLinkServiceGroups2edl(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemVirtualWanLinkServiceHealthCheck2edl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("hold_down_time"); ok || d.HasChange("hold_down_time") {
		t, err := expandWantempSystemVirtualWanLinkServiceHoldDownTime2edl(d, v, "hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWantempSystemVirtualWanLinkServiceId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok || d.HasChange("input_device") {
		t, err := expandWantempSystemVirtualWanLinkServiceInputDevice2edl(d, v, "input_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	}

	if v, ok := d.GetOk("input_device_negate"); ok || d.HasChange("input_device_negate") {
		t, err := expandWantempSystemVirtualWanLinkServiceInputDeviceNegate2edl(d, v, "input_device_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok || d.HasChange("internet_service") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetService2edl(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_ctrl"); ok || d.HasChange("internet_service_ctrl") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceCtrl2edl(d, v, "internet_service_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_ctrl_group"); ok || d.HasChange("internet_service_ctrl_group") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup2edl(d, v, "internet_service_ctrl_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-ctrl-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl"); ok || d.HasChange("internet_service_app_ctrl") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl2edl(d, v, "internet_service_app_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_group"); ok || d.HasChange("internet_service_app_ctrl_group") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup2edl(d, v, "internet_service_app_ctrl_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceCustom2edl(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup2edl(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceGroup2edl(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandWantempSystemVirtualWanLinkServiceInternetServiceId2edl(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("jitter_weight"); ok || d.HasChange("jitter_weight") {
		t, err := expandWantempSystemVirtualWanLinkServiceJitterWeight2edl(d, v, "jitter_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-weight"] = t
		}
	}

	if v, ok := d.GetOk("latency_weight"); ok || d.HasChange("latency_weight") {
		t, err := expandWantempSystemVirtualWanLinkServiceLatencyWeight2edl(d, v, "latency_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-weight"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandWantempSystemVirtualWanLinkServiceLinkCostFactor2edl(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_threshold"); ok || d.HasChange("link_cost_threshold") {
		t, err := expandWantempSystemVirtualWanLinkServiceLinkCostThreshold2edl(d, v, "link_cost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandWantempSystemVirtualWanLinkServiceMember2edl(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandWantempSystemVirtualWanLinkServiceMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWantempSystemVirtualWanLinkServiceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_weight"); ok || d.HasChange("packet_loss_weight") {
		t, err := expandWantempSystemVirtualWanLinkServicePacketLossWeight2edl(d, v, "packet_loss_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-weight"] = t
		}
	}

	if v, ok := d.GetOk("priority_members"); ok || d.HasChange("priority_members") {
		t, err := expandWantempSystemVirtualWanLinkServicePriorityMembers2edl(d, v, "priority_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-members"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandWantempSystemVirtualWanLinkServiceProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_link"); ok || d.HasChange("quality_link") {
		t, err := expandWantempSystemVirtualWanLinkServiceQualityLink2edl(d, v, "quality_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-link"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandWantempSystemVirtualWanLinkServiceRole2edl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok || d.HasChange("route_tag") {
		t, err := expandWantempSystemVirtualWanLinkServiceRouteTag2edl(d, v, "route_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok || d.HasChange("sla") {
		t, err := expandWantempSystemVirtualWanLinkServiceSla2edl(d, v, "sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	}

	if v, ok := d.GetOk("sla_compare_method"); ok || d.HasChange("sla_compare_method") {
		t, err := expandWantempSystemVirtualWanLinkServiceSlaCompareMethod2edl(d, v, "sla_compare_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-compare-method"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandWantempSystemVirtualWanLinkServiceSrc2edl(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("src_negate"); ok || d.HasChange("src_negate") {
		t, err := expandWantempSystemVirtualWanLinkServiceSrcNegate2edl(d, v, "src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-negate"] = t
		}
	}

	if v, ok := d.GetOk("src6"); ok || d.HasChange("src6") {
		t, err := expandWantempSystemVirtualWanLinkServiceSrc62edl(d, v, "src6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src6"] = t
		}
	}

	if v, ok := d.GetOk("standalone_action"); ok || d.HasChange("standalone_action") {
		t, err := expandWantempSystemVirtualWanLinkServiceStandaloneAction2edl(d, v, "standalone_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-action"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok || d.HasChange("start_port") {
		t, err := expandWantempSystemVirtualWanLinkServiceStartPort2edl(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemVirtualWanLinkServiceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandWantempSystemVirtualWanLinkServiceTos2edl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandWantempSystemVirtualWanLinkServiceTosMask2edl(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandWantempSystemVirtualWanLinkServiceUsers2edl(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	return &obj, nil
}
