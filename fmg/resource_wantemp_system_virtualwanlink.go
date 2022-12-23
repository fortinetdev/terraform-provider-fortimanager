// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemVirtualWanLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkUpdate,
		Read:   resourceWantempSystemVirtualWanLinkRead,
		Update: resourceWantempSystemVirtualWanLinkUpdate,
		Delete: resourceWantempSystemVirtualWanLinkDelete,

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
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"load_balance_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_dynamic_member": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ingress_spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"seq_num": &schema.Schema{
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
						"spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_ratio": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sla_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"neighbor_hold_boot_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"neighbor_hold_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"neighbor_hold_down_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
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
					},
				},
			},
			"status": &schema.Schema{
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

func resourceWantempSystemVirtualWanLinkUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemVirtualWanLink(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLink resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLink(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLink resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WantempSystemVirtualWanLink")

	return resourceWantempSystemVirtualWanLinkRead(d, m)
}

func resourceWantempSystemVirtualWanLinkDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemVirtualWanLink(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLink resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemVirtualWanLink(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLink resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLink(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLink resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkFailDetectWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckWsva(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dynamic_server"
		if _, ok := i["_dynamic-server"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckDynamicServerWsva(i["_dynamic-server"], d, pre_append)
			tmp["_dynamic_server"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-DynamicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckAddrModeWsva(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsva(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckFailtimeWsva(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHaPriorityWsva(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHttpAgentWsva(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHttpGetWsva(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHttpMatchWsva(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := i["internet-service-id"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsva(i["internet-service-id"], d, pre_append)
			tmp["internet_service_id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-InternetServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckIntervalWsva(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckMembersWsva(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckNameWsva(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckPacketSizeWsva(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckPasswordWsva(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckPortWsva(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckProbePacketsWsva(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsva(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckProtocolWsva(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsva(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSecurityModeWsva(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckServerWsva(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaWsva(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsva(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsva(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsva(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsva(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsva(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsva(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsva(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsva(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsva(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsva(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-UpdateStaticRoute")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkHealthCheckDynamicServerWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckAddrModeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckFailtimeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHaPriorityWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpAgentWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpGetWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpMatchWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckIntervalWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckMembersWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckNameWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckPacketSizeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckPasswordWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckPortWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProbePacketsWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProtocolWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSecurityModeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckServerWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaWsva(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaIdWsva(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsva(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsva(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsva(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsva(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-PacketlossThreshold")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaIdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkLoadBalanceModeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersWsva(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dynamic_member"
		if _, ok := i["_dynamic-member"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersDynamicMemberWsva(i["_dynamic-member"], d, pre_append)
			tmp["_dynamic_member"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-DynamicMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersCommentWsva(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersCostWsva(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersGatewayWsva(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersGateway6Wsva(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersIngressSpilloverThresholdWsva(i["ingress-spillover-threshold"], d, pre_append)
			tmp["ingress_spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-IngressSpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersInterfaceWsva(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersPriorityWsva(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSeqNumWsva(i["seq-num"], d, pre_append)
			tmp["seq_num"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-SeqNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSourceWsva(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSource6Wsva(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSpilloverThresholdWsva(i["spillover-threshold"], d, pre_append)
			tmp["spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-SpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersStatusWsva(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersVolumeRatioWsva(i["volume-ratio"], d, pre_append)
			tmp["volume_ratio"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-VolumeRatio")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersWeightWsva(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkMembersDynamicMemberWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersCommentWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersCostWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersGatewayWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersGateway6Wsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersIngressSpilloverThresholdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersInterfaceWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersPriorityWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSeqNumWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSourceWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSource6Wsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSpilloverThresholdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersStatusWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersVolumeRatioWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersWeightWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborWsva(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkNeighborHealthCheckWsva(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborIpWsva(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborMemberWsva(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborRoleWsva(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := i["sla-id"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborSlaIdWsva(i["sla-id"], d, pre_append)
			tmp["sla_id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-SlaId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkNeighborHealthCheckWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborIpWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborMemberWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborRoleWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborSlaIdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborHoldBootTimeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborHoldDownWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborHoldDownTimeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceWsva(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceAddrModeWsva(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := i["bandwidth-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceBandwidthWeightWsva(i["bandwidth-weight"], d, pre_append)
			tmp["bandwidth_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-BandwidthWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDefaultWsva(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := i["dscp-forward"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpForwardWsva(i["dscp-forward"], d, pre_append)
			tmp["dscp_forward"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpForward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := i["dscp-forward-tag"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpForwardTagWsva(i["dscp-forward-tag"], d, pre_append)
			tmp["dscp_forward_tag"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpForwardTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := i["dscp-reverse"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpReverseWsva(i["dscp-reverse"], d, pre_append)
			tmp["dscp_reverse"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := i["dscp-reverse-tag"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpReverseTagWsva(i["dscp-reverse-tag"], d, pre_append)
			tmp["dscp_reverse_tag"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpReverseTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDstWsva(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := i["dst-negate"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDstNegateWsva(i["dst-negate"], d, pre_append)
			tmp["dst_negate"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DstNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDst6Wsva(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceEndPortWsva(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceGatewayWsva(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceGroupsWsva(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceHealthCheckWsva(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := i["hold-down-time"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceHoldDownTimeWsva(i["hold-down-time"], d, pre_append)
			tmp["hold_down_time"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-HoldDownTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceIdWsva(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := i["input-device"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInputDeviceWsva(i["input-device"], d, pre_append)
			tmp["input_device"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InputDevice")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := i["input-device-negate"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInputDeviceNegateWsva(i["input-device-negate"], d, pre_append)
			tmp["input_device_negate"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InputDeviceNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := i["internet-service"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceWsva(i["internet-service"], d, pre_append)
			tmp["internet_service"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl"
		if _, ok := i["internet-service-ctrl"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlWsva(i["internet-service-ctrl"], d, pre_append)
			tmp["internet_service_ctrl"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl_group"
		if _, ok := i["internet-service-ctrl-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroupWsva(i["internet-service-ctrl-group"], d, pre_append)
			tmp["internet_service_ctrl_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := i["internet-service-app-ctrl"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlWsva(i["internet-service-app-ctrl"], d, pre_append)
			tmp["internet_service_app_ctrl"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceAppCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := i["internet-service-app-ctrl-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupWsva(i["internet-service-app-ctrl-group"], d, pre_append)
			tmp["internet_service_app_ctrl_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceAppCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := i["internet-service-custom"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomWsva(i["internet-service-custom"], d, pre_append)
			tmp["internet_service_custom"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCustom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := i["internet-service-custom-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomGroupWsva(i["internet-service-custom-group"], d, pre_append)
			tmp["internet_service_custom_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCustomGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := i["internet-service-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceGroupWsva(i["internet-service-group"], d, pre_append)
			tmp["internet_service_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := i["internet-service-id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceIdWsva(i["internet-service-id"], d, pre_append)
			tmp["internet_service_id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := i["jitter-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceJitterWeightWsva(i["jitter-weight"], d, pre_append)
			tmp["jitter_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-JitterWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := i["latency-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceLatencyWeightWsva(i["latency-weight"], d, pre_append)
			tmp["latency_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-LatencyWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceLinkCostFactorWsva(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := i["link-cost-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceLinkCostThresholdWsva(i["link-cost-threshold"], d, pre_append)
			tmp["link_cost_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-LinkCostThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceMemberWsva(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceModeWsva(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceNameWsva(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := i["packet-loss-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServicePacketLossWeightWsva(i["packet-loss-weight"], d, pre_append)
			tmp["packet_loss_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-PacketLossWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := i["priority-members"]; ok {
			v := flattenWantempSystemVirtualWanLinkServicePriorityMembersWsva(i["priority-members"], d, pre_append)
			tmp["priority_members"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-PriorityMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceProtocolWsva(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := i["quality-link"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceQualityLinkWsva(i["quality-link"], d, pre_append)
			tmp["quality_link"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-QualityLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceRoleWsva(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceRouteTagWsva(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSlaWsva(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := i["sla-compare-method"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSlaCompareMethodWsva(i["sla-compare-method"], d, pre_append)
			tmp["sla_compare_method"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-SlaCompareMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSrcWsva(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := i["src-negate"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSrcNegateWsva(i["src-negate"], d, pre_append)
			tmp["src_negate"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-SrcNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSrc6Wsva(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := i["standalone-action"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceStandaloneActionWsva(i["standalone-action"], d, pre_append)
			tmp["standalone_action"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-StandaloneAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceStartPortWsva(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-StartPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceStatusWsva(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceTosWsva(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceTosMaskWsva(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-TosMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceUsersWsva(i["users"], d, pre_append)
			tmp["users"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Users")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkServiceAddrModeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceBandwidthWeightWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDefaultWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpForwardWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpForwardTagWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpReverseWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpReverseTagWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDstWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDstNegateWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDst6Wsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceEndPortWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceGatewayWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceGroupsWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceHealthCheckWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceHoldDownTimeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceIdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInputDeviceWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInputDeviceNegateWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroupWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomGroupWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceGroupWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceIdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceJitterWeightWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLatencyWeightWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLinkCostFactorWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLinkCostThresholdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceMemberWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceModeWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceNameWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServicePacketLossWeightWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServicePriorityMembersWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceProtocolWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceQualityLinkWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceRoleWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceRouteTagWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaWsva(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkServiceSlaHealthCheckWsva(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSlaIdWsva(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkService-Sla-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkServiceSlaHealthCheckWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaIdWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaCompareMethodWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrcWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrcNegateWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrc6Wsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStandaloneActionWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStartPortWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStatusWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceTosWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceTosMaskWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceUsersWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkStatusWsva(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemVirtualWanLink(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fail_detect", flattenWantempSystemVirtualWanLinkFailDetectWsva(o["fail-detect"], d, "fail_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect"], "WantempSystemVirtualWanLink-FailDetect"); ok {
			if err = d.Set("fail_detect", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkHealthCheckWsva(o["health-check"], d, "health_check")); err != nil {
			if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemVirtualWanLink-HealthCheck"); ok {
				if err = d.Set("health_check", vv); err != nil {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check"); ok {
			if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkHealthCheckWsva(o["health-check"], d, "health_check")); err != nil {
				if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemVirtualWanLink-HealthCheck"); ok {
					if err = d.Set("health_check", vv); err != nil {
						return fmt.Errorf("Error reading health_check: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			}
		}
	}

	if err = d.Set("load_balance_mode", flattenWantempSystemVirtualWanLinkLoadBalanceModeWsva(o["load-balance-mode"], d, "load_balance_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-mode"], "WantempSystemVirtualWanLink-LoadBalanceMode"); ok {
			if err = d.Set("load_balance_mode", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenWantempSystemVirtualWanLinkMembersWsva(o["members"], d, "members")); err != nil {
			if vv, ok := fortiAPIPatch(o["members"], "WantempSystemVirtualWanLink-Members"); ok {
				if err = d.Set("members", vv); err != nil {
					return fmt.Errorf("Error reading members: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenWantempSystemVirtualWanLinkMembersWsva(o["members"], d, "members")); err != nil {
				if vv, ok := fortiAPIPatch(o["members"], "WantempSystemVirtualWanLink-Members"); ok {
					if err = d.Set("members", vv); err != nil {
						return fmt.Errorf("Error reading members: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenWantempSystemVirtualWanLinkNeighborWsva(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "WantempSystemVirtualWanLink-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenWantempSystemVirtualWanLinkNeighborWsva(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "WantempSystemVirtualWanLink-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if err = d.Set("neighbor_hold_boot_time", flattenWantempSystemVirtualWanLinkNeighborHoldBootTimeWsva(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-boot-time"], "WantempSystemVirtualWanLink-NeighborHoldBootTime"); ok {
			if err = d.Set("neighbor_hold_boot_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenWantempSystemVirtualWanLinkNeighborHoldDownWsva(o["neighbor-hold-down"], d, "neighbor_hold_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down"], "WantempSystemVirtualWanLink-NeighborHoldDown"); ok {
			if err = d.Set("neighbor_hold_down", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenWantempSystemVirtualWanLinkNeighborHoldDownTimeWsva(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down-time"], "WantempSystemVirtualWanLink-NeighborHoldDownTime"); ok {
			if err = d.Set("neighbor_hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenWantempSystemVirtualWanLinkServiceWsva(o["service"], d, "service")); err != nil {
			if vv, ok := fortiAPIPatch(o["service"], "WantempSystemVirtualWanLink-Service"); ok {
				if err = d.Set("service", vv); err != nil {
					return fmt.Errorf("Error reading service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenWantempSystemVirtualWanLinkServiceWsva(o["service"], d, "service")); err != nil {
				if vv, ok := fortiAPIPatch(o["service"], "WantempSystemVirtualWanLink-Service"); ok {
					if err = d.Set("service", vv); err != nil {
						return fmt.Errorf("Error reading service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenWantempSystemVirtualWanLinkStatusWsva(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemVirtualWanLink-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkFailDetectWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dynamic_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_dynamic-server"], _ = expandWantempSystemVirtualWanLinkHealthCheckDynamicServerWsva(d, i["_dynamic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandWantempSystemVirtualWanLinkHealthCheckAddrModeWsva(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsva(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandWantempSystemVirtualWanLinkHealthCheckFailtimeWsva(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandWantempSystemVirtualWanLinkHealthCheckHaPriorityWsva(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandWantempSystemVirtualWanLinkHealthCheckHttpAgentWsva(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandWantempSystemVirtualWanLinkHealthCheckHttpGetWsva(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandWantempSystemVirtualWanLinkHealthCheckHttpMatchWsva(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-id"], _ = expandWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsva(d, i["internet_service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandWantempSystemVirtualWanLinkHealthCheckIntervalWsva(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandWantempSystemVirtualWanLinkHealthCheckMembersWsva(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemVirtualWanLinkHealthCheckNameWsva(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandWantempSystemVirtualWanLinkHealthCheckPacketSizeWsva(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandWantempSystemVirtualWanLinkHealthCheckPasswordWsva(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandWantempSystemVirtualWanLinkHealthCheckPortWsva(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandWantempSystemVirtualWanLinkHealthCheckProbePacketsWsva(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsva(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemVirtualWanLinkHealthCheckProtocolWsva(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsva(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandWantempSystemVirtualWanLinkHealthCheckSecurityModeWsva(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandWantempSystemVirtualWanLinkHealthCheckServerWsva(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaWsva(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsva(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsva(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsva(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsva(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsva(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsva(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsva(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsva(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsva(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsva(d, i["update_static_route"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckDynamicServerWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckAddrModeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckDiffservcodeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckFailtimeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHaPriorityWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpAgentWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpGetWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpMatchWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckInternetServiceIdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckIntervalWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckMembersWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckNameWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPacketSizeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPasswordWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPortWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProbePacketsWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProbeTimeoutWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProtocolWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckRecoverytimeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSecurityModeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckServerWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaIdWsva(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsva(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsva(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsva(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsva(d, i["packetloss_threshold"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaIdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThresholdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThresholdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactorWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThresholdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriodWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriodWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitterWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatencyWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketlossWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitterWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatencyWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketlossWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterfaceWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckUpdateStaticRouteWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkLoadBalanceModeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_dynamic_member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_dynamic-member"], _ = expandWantempSystemVirtualWanLinkMembersDynamicMemberWsva(d, i["_dynamic_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWantempSystemVirtualWanLinkMembersCommentWsva(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandWantempSystemVirtualWanLinkMembersCostWsva(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemVirtualWanLinkMembersGatewayWsva(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway6"], _ = expandWantempSystemVirtualWanLinkMembersGateway6Wsva(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-spillover-threshold"], _ = expandWantempSystemVirtualWanLinkMembersIngressSpilloverThresholdWsva(d, i["ingress_spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandWantempSystemVirtualWanLinkMembersInterfaceWsva(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandWantempSystemVirtualWanLinkMembersPriorityWsva(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq-num"], _ = expandWantempSystemVirtualWanLinkMembersSeqNumWsva(d, i["seq_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemVirtualWanLinkMembersSourceWsva(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemVirtualWanLinkMembersSource6Wsva(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spillover-threshold"], _ = expandWantempSystemVirtualWanLinkMembersSpilloverThresholdWsva(d, i["spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemVirtualWanLinkMembersStatusWsva(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["volume-ratio"], _ = expandWantempSystemVirtualWanLinkMembersVolumeRatioWsva(d, i["volume_ratio"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandWantempSystemVirtualWanLinkMembersWeightWsva(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkMembersDynamicMemberWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersCommentWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersCostWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersGatewayWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersGateway6Wsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersIngressSpilloverThresholdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersInterfaceWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersPriorityWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSeqNumWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSourceWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSource6Wsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSpilloverThresholdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersStatusWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersVolumeRatioWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersWeightWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkNeighborHealthCheckWsva(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWantempSystemVirtualWanLinkNeighborIpWsva(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandWantempSystemVirtualWanLinkNeighborMemberWsva(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemVirtualWanLinkNeighborRoleWsva(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id"], _ = expandWantempSystemVirtualWanLinkNeighborSlaIdWsva(d, i["sla_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkNeighborHealthCheckWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborIpWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborMemberWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborRoleWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborSlaIdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborHoldBootTimeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborHoldDownWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborHoldDownTimeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandWantempSystemVirtualWanLinkServiceAddrModeWsva(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-weight"], _ = expandWantempSystemVirtualWanLinkServiceBandwidthWeightWsva(d, i["bandwidth_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandWantempSystemVirtualWanLinkServiceDefaultWsva(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward"], _ = expandWantempSystemVirtualWanLinkServiceDscpForwardWsva(d, i["dscp_forward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward-tag"], _ = expandWantempSystemVirtualWanLinkServiceDscpForwardTagWsva(d, i["dscp_forward_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse"], _ = expandWantempSystemVirtualWanLinkServiceDscpReverseWsva(d, i["dscp_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse-tag"], _ = expandWantempSystemVirtualWanLinkServiceDscpReverseTagWsva(d, i["dscp_reverse_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandWantempSystemVirtualWanLinkServiceDstWsva(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-negate"], _ = expandWantempSystemVirtualWanLinkServiceDstNegateWsva(d, i["dst_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst6"], _ = expandWantempSystemVirtualWanLinkServiceDst6Wsva(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-port"], _ = expandWantempSystemVirtualWanLinkServiceEndPortWsva(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemVirtualWanLinkServiceGatewayWsva(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandWantempSystemVirtualWanLinkServiceGroupsWsva(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkServiceHealthCheckWsva(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hold-down-time"], _ = expandWantempSystemVirtualWanLinkServiceHoldDownTimeWsva(d, i["hold_down_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemVirtualWanLinkServiceIdWsva(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device"], _ = expandWantempSystemVirtualWanLinkServiceInputDeviceWsva(d, i["input_device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device-negate"], _ = expandWantempSystemVirtualWanLinkServiceInputDeviceNegateWsva(d, i["input_device_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceWsva(d, i["internet_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-ctrl"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlWsva(d, i["internet_service_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-ctrl-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroupWsva(d, i["internet_service_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlWsva(d, i["internet_service_app_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupWsva(d, i["internet_service_app_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCustomWsva(d, i["internet_service_custom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCustomGroupWsva(d, i["internet_service_custom_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceGroupWsva(d, i["internet_service_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-id"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceIdWsva(d, i["internet_service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-weight"], _ = expandWantempSystemVirtualWanLinkServiceJitterWeightWsva(d, i["jitter_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-weight"], _ = expandWantempSystemVirtualWanLinkServiceLatencyWeightWsva(d, i["latency_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemVirtualWanLinkServiceLinkCostFactorWsva(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-threshold"], _ = expandWantempSystemVirtualWanLinkServiceLinkCostThresholdWsva(d, i["link_cost_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandWantempSystemVirtualWanLinkServiceMemberWsva(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandWantempSystemVirtualWanLinkServiceModeWsva(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemVirtualWanLinkServiceNameWsva(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-weight"], _ = expandWantempSystemVirtualWanLinkServicePacketLossWeightWsva(d, i["packet_loss_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-members"], _ = expandWantempSystemVirtualWanLinkServicePriorityMembersWsva(d, i["priority_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemVirtualWanLinkServiceProtocolWsva(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-link"], _ = expandWantempSystemVirtualWanLinkServiceQualityLinkWsva(d, i["quality_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemVirtualWanLinkServiceRoleWsva(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandWantempSystemVirtualWanLinkServiceRouteTagWsva(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemVirtualWanLinkServiceSlaWsva(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-compare-method"], _ = expandWantempSystemVirtualWanLinkServiceSlaCompareMethodWsva(d, i["sla_compare_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandWantempSystemVirtualWanLinkServiceSrcWsva(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-negate"], _ = expandWantempSystemVirtualWanLinkServiceSrcNegateWsva(d, i["src_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src6"], _ = expandWantempSystemVirtualWanLinkServiceSrc6Wsva(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["standalone-action"], _ = expandWantempSystemVirtualWanLinkServiceStandaloneActionWsva(d, i["standalone_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandWantempSystemVirtualWanLinkServiceStartPortWsva(d, i["start_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemVirtualWanLinkServiceStatusWsva(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandWantempSystemVirtualWanLinkServiceTosWsva(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandWantempSystemVirtualWanLinkServiceTosMaskWsva(d, i["tos_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["users"], _ = expandWantempSystemVirtualWanLinkServiceUsersWsva(d, i["users"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkServiceAddrModeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceBandwidthWeightWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDefaultWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpForwardWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpForwardTagWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpReverseWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpReverseTagWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDstWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDstNegateWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDst6Wsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceEndPortWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceGatewayWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceGroupsWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceHealthCheckWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceHoldDownTimeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceIdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInputDeviceWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInputDeviceNegateWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroupWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCustomWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCustomGroupWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceGroupWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceIdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceJitterWeightWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLatencyWeightWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLinkCostFactorWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLinkCostThresholdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceMemberWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceModeWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceNameWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServicePacketLossWeightWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServicePriorityMembersWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceProtocolWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceQualityLinkWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceRoleWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceRouteTagWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkServiceSlaHealthCheckWsva(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemVirtualWanLinkServiceSlaIdWsva(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaHealthCheckWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaIdWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaCompareMethodWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrcWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrcNegateWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrc6Wsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStandaloneActionWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStartPortWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStatusWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceTosWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceTosMaskWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceUsersWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkStatusWsva(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLink(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fail_detect"); ok || d.HasChange("fail_detect") {
		t, err := expandWantempSystemVirtualWanLinkFailDetectWsva(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckWsva(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_mode"); ok || d.HasChange("load_balance_mode") {
		t, err := expandWantempSystemVirtualWanLinkLoadBalanceModeWsva(d, v, "load_balance_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-mode"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandWantempSystemVirtualWanLinkMembersWsva(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandWantempSystemVirtualWanLinkNeighborWsva(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_boot_time"); ok || d.HasChange("neighbor_hold_boot_time") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHoldBootTimeWsva(d, v, "neighbor_hold_boot_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-boot-time"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok || d.HasChange("neighbor_hold_down") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHoldDownWsva(d, v, "neighbor_hold_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down_time"); ok || d.HasChange("neighbor_hold_down_time") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHoldDownTimeWsva(d, v, "neighbor_hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandWantempSystemVirtualWanLinkServiceWsva(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemVirtualWanLinkStatusWsva(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
