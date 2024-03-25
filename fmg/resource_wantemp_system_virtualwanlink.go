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
			"fail_alert_interfaces": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
							Type:     schema.TypeString,
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
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
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

func flattenWantempSystemVirtualWanLinkFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkFailDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheck(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkHealthCheckDynamicServer(i["_dynamic-server"], d, pre_append)
			tmp["_dynamic_server"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-DynamicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckDiffservcode(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckFailtime(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHttpAgent(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHttpGet(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckHttpMatch(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := i["internet-service-id"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckInternetServiceId(i["internet-service-id"], d, pre_append)
			tmp["internet_service_id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-InternetServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckInterval(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckPacketSize(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckProbePackets(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckProbeTimeout(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckRecoverytime(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSecurityMode(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitter(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatency(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitter(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatency(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckUpdateStaticRoute(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-HealthCheck-UpdateStaticRoute")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkHealthCheckDynamicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkHealthCheckName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkHealthCheck-Sla-PacketlossThreshold")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkLoadBalanceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkMembersDynamicMember(i["_dynamic-member"], d, pre_append)
			tmp["_dynamic_member"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-DynamicMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersGateway6(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(i["ingress-spillover-threshold"], d, pre_append)
			tmp["ingress_spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-IngressSpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSeqNum(i["seq-num"], d, pre_append)
			tmp["seq_num"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-SeqNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersSpilloverThreshold(i["spillover-threshold"], d, pre_append)
			tmp["spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-SpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersVolumeRatio(i["volume-ratio"], d, pre_append)
			tmp["volume_ratio"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-VolumeRatio")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkMembersWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Members-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkMembersDynamicMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWantempSystemVirtualWanLinkMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkNeighborHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := i["sla-id"]; ok {
			v := flattenWantempSystemVirtualWanLinkNeighborSlaId(i["sla-id"], d, pre_append)
			tmp["sla_id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Neighbor-SlaId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborHoldBootTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborHoldDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkServiceAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := i["bandwidth-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceBandwidthWeight(i["bandwidth-weight"], d, pre_append)
			tmp["bandwidth_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-BandwidthWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDefault(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := i["dscp-forward"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpForward(i["dscp-forward"], d, pre_append)
			tmp["dscp_forward"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpForward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := i["dscp-forward-tag"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpForwardTag(i["dscp-forward-tag"], d, pre_append)
			tmp["dscp_forward_tag"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpForwardTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := i["dscp-reverse"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpReverse(i["dscp-reverse"], d, pre_append)
			tmp["dscp_reverse"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := i["dscp-reverse-tag"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDscpReverseTag(i["dscp-reverse-tag"], d, pre_append)
			tmp["dscp_reverse_tag"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DscpReverseTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := i["dst-negate"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDstNegate(i["dst-negate"], d, pre_append)
			tmp["dst_negate"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-DstNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceDst6(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceEndPort(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceGroups(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := i["hold-down-time"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceHoldDownTime(i["hold-down-time"], d, pre_append)
			tmp["hold_down_time"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-HoldDownTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := i["input-device"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInputDevice(i["input-device"], d, pre_append)
			tmp["input_device"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InputDevice")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := i["input-device-negate"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInputDeviceNegate(i["input-device-negate"], d, pre_append)
			tmp["input_device_negate"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InputDeviceNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := i["internet-service"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetService(i["internet-service"], d, pre_append)
			tmp["internet_service"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl"
		if _, ok := i["internet-service-ctrl"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrl(i["internet-service-ctrl"], d, pre_append)
			tmp["internet_service_ctrl"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl_group"
		if _, ok := i["internet-service-ctrl-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup(i["internet-service-ctrl-group"], d, pre_append)
			tmp["internet_service_ctrl_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := i["internet-service-app-ctrl"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl(i["internet-service-app-ctrl"], d, pre_append)
			tmp["internet_service_app_ctrl"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceAppCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := i["internet-service-app-ctrl-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(i["internet-service-app-ctrl-group"], d, pre_append)
			tmp["internet_service_app_ctrl_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceAppCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := i["internet-service-custom"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCustom(i["internet-service-custom"], d, pre_append)
			tmp["internet_service_custom"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCustom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := i["internet-service-custom-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup(i["internet-service-custom-group"], d, pre_append)
			tmp["internet_service_custom_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceCustomGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := i["internet-service-group"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceGroup(i["internet-service-group"], d, pre_append)
			tmp["internet_service_group"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := i["internet-service-id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceInternetServiceId(i["internet-service-id"], d, pre_append)
			tmp["internet_service_id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-InternetServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := i["jitter-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceJitterWeight(i["jitter-weight"], d, pre_append)
			tmp["jitter_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-JitterWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := i["latency-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceLatencyWeight(i["latency-weight"], d, pre_append)
			tmp["latency_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-LatencyWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := i["link-cost-threshold"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceLinkCostThreshold(i["link-cost-threshold"], d, pre_append)
			tmp["link_cost_threshold"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-LinkCostThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := i["packet-loss-weight"]; ok {
			v := flattenWantempSystemVirtualWanLinkServicePacketLossWeight(i["packet-loss-weight"], d, pre_append)
			tmp["packet_loss_weight"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-PacketLossWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := i["priority-members"]; ok {
			v := flattenWantempSystemVirtualWanLinkServicePriorityMembers(i["priority-members"], d, pre_append)
			tmp["priority_members"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-PriorityMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := i["quality-link"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceQualityLink(i["quality-link"], d, pre_append)
			tmp["quality_link"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-QualityLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceRouteTag(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := i["sla-compare-method"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSlaCompareMethod(i["sla-compare-method"], d, pre_append)
			tmp["sla_compare_method"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-SlaCompareMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSrc(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := i["src-negate"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSrcNegate(i["src-negate"], d, pre_append)
			tmp["src_negate"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-SrcNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSrc6(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := i["standalone-action"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceStandaloneAction(i["standalone-action"], d, pre_append)
			tmp["standalone_action"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-StandaloneAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceStartPort(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-StartPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceTos(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceTosMask(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-TosMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceUsers(i["users"], d, pre_append)
			tmp["users"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLink-Service-Users")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkServiceAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceDstNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceDst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceEndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceQualityLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemVirtualWanLinkServiceSlaHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemVirtualWanLinkServiceSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemVirtualWanLinkService-Sla-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemVirtualWanLinkServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceSrc6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkServiceUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemVirtualWanLinkStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("fail_alert_interfaces", flattenWantempSystemVirtualWanLinkFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-alert-interfaces"], "WantempSystemVirtualWanLink-FailAlertInterfaces"); ok {
			if err = d.Set("fail_alert_interfaces", vv); err != nil {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenWantempSystemVirtualWanLinkFailDetect(o["fail-detect"], d, "fail_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect"], "WantempSystemVirtualWanLink-FailDetect"); ok {
			if err = d.Set("fail_detect", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkHealthCheck(o["health-check"], d, "health_check")); err != nil {
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
			if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkHealthCheck(o["health-check"], d, "health_check")); err != nil {
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

	if err = d.Set("load_balance_mode", flattenWantempSystemVirtualWanLinkLoadBalanceMode(o["load-balance-mode"], d, "load_balance_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-mode"], "WantempSystemVirtualWanLink-LoadBalanceMode"); ok {
			if err = d.Set("load_balance_mode", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenWantempSystemVirtualWanLinkMembers(o["members"], d, "members")); err != nil {
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
			if err = d.Set("members", flattenWantempSystemVirtualWanLinkMembers(o["members"], d, "members")); err != nil {
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
		if err = d.Set("neighbor", flattenWantempSystemVirtualWanLinkNeighbor(o["neighbor"], d, "neighbor")); err != nil {
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
			if err = d.Set("neighbor", flattenWantempSystemVirtualWanLinkNeighbor(o["neighbor"], d, "neighbor")); err != nil {
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

	if err = d.Set("neighbor_hold_boot_time", flattenWantempSystemVirtualWanLinkNeighborHoldBootTime(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-boot-time"], "WantempSystemVirtualWanLink-NeighborHoldBootTime"); ok {
			if err = d.Set("neighbor_hold_boot_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenWantempSystemVirtualWanLinkNeighborHoldDown(o["neighbor-hold-down"], d, "neighbor_hold_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down"], "WantempSystemVirtualWanLink-NeighborHoldDown"); ok {
			if err = d.Set("neighbor_hold_down", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenWantempSystemVirtualWanLinkNeighborHoldDownTime(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down-time"], "WantempSystemVirtualWanLink-NeighborHoldDownTime"); ok {
			if err = d.Set("neighbor_hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenWantempSystemVirtualWanLinkService(o["service"], d, "service")); err != nil {
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
			if err = d.Set("service", flattenWantempSystemVirtualWanLinkService(o["service"], d, "service")); err != nil {
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

	if err = d.Set("status", flattenWantempSystemVirtualWanLinkStatus(o["status"], d, "status")); err != nil {
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

func expandWantempSystemVirtualWanLinkFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkFailDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_dynamic-server"], _ = expandWantempSystemVirtualWanLinkHealthCheckDynamicServer(d, i["_dynamic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandWantempSystemVirtualWanLinkHealthCheckAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandWantempSystemVirtualWanLinkHealthCheckDiffservcode(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandWantempSystemVirtualWanLinkHealthCheckFailtime(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandWantempSystemVirtualWanLinkHealthCheckHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandWantempSystemVirtualWanLinkHealthCheckHttpAgent(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandWantempSystemVirtualWanLinkHealthCheckHttpGet(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandWantempSystemVirtualWanLinkHealthCheckHttpMatch(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-id"], _ = expandWantempSystemVirtualWanLinkHealthCheckInternetServiceId(d, i["internet_service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandWantempSystemVirtualWanLinkHealthCheckInterval(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandWantempSystemVirtualWanLinkHealthCheckMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemVirtualWanLinkHealthCheckName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandWantempSystemVirtualWanLinkHealthCheckPacketSize(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandWantempSystemVirtualWanLinkHealthCheckPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandWantempSystemVirtualWanLinkHealthCheckPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandWantempSystemVirtualWanLinkHealthCheckProbePackets(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandWantempSystemVirtualWanLinkHealthCheckProbeTimeout(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemVirtualWanLinkHealthCheckProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandWantempSystemVirtualWanLinkHealthCheckRecoverytime(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandWantempSystemVirtualWanLinkHealthCheckSecurityMode(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandWantempSystemVirtualWanLinkHealthCheckServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemVirtualWanLinkHealthCheckSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandWantempSystemVirtualWanLinkHealthCheckUpdateStaticRoute(d, i["update_static_route"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckDynamicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProbePackets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProbeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkLoadBalanceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_dynamic-member"], _ = expandWantempSystemVirtualWanLinkMembersDynamicMember(d, i["_dynamic_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWantempSystemVirtualWanLinkMembersComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandWantempSystemVirtualWanLinkMembersCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemVirtualWanLinkMembersGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway6"], _ = expandWantempSystemVirtualWanLinkMembersGateway6(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-spillover-threshold"], _ = expandWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandWantempSystemVirtualWanLinkMembersInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandWantempSystemVirtualWanLinkMembersPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq-num"], _ = expandWantempSystemVirtualWanLinkMembersSeqNum(d, i["seq_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemVirtualWanLinkMembersSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemVirtualWanLinkMembersSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spillover-threshold"], _ = expandWantempSystemVirtualWanLinkMembersSpilloverThreshold(d, i["spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemVirtualWanLinkMembersStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["volume-ratio"], _ = expandWantempSystemVirtualWanLinkMembersVolumeRatio(d, i["volume_ratio"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandWantempSystemVirtualWanLinkMembersWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkMembersDynamicMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkMembersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkNeighborHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWantempSystemVirtualWanLinkNeighborIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandWantempSystemVirtualWanLinkNeighborMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemVirtualWanLinkNeighborRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id"], _ = expandWantempSystemVirtualWanLinkNeighborSlaId(d, i["sla_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborHoldBootTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborHoldDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-mode"], _ = expandWantempSystemVirtualWanLinkServiceAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-weight"], _ = expandWantempSystemVirtualWanLinkServiceBandwidthWeight(d, i["bandwidth_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandWantempSystemVirtualWanLinkServiceDefault(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward"], _ = expandWantempSystemVirtualWanLinkServiceDscpForward(d, i["dscp_forward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward-tag"], _ = expandWantempSystemVirtualWanLinkServiceDscpForwardTag(d, i["dscp_forward_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse"], _ = expandWantempSystemVirtualWanLinkServiceDscpReverse(d, i["dscp_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse-tag"], _ = expandWantempSystemVirtualWanLinkServiceDscpReverseTag(d, i["dscp_reverse_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandWantempSystemVirtualWanLinkServiceDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-negate"], _ = expandWantempSystemVirtualWanLinkServiceDstNegate(d, i["dst_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst6"], _ = expandWantempSystemVirtualWanLinkServiceDst6(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-port"], _ = expandWantempSystemVirtualWanLinkServiceEndPort(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemVirtualWanLinkServiceGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandWantempSystemVirtualWanLinkServiceGroups(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkServiceHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hold-down-time"], _ = expandWantempSystemVirtualWanLinkServiceHoldDownTime(d, i["hold_down_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemVirtualWanLinkServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device"], _ = expandWantempSystemVirtualWanLinkServiceInputDevice(d, i["input_device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device-negate"], _ = expandWantempSystemVirtualWanLinkServiceInputDeviceNegate(d, i["input_device_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service"], _ = expandWantempSystemVirtualWanLinkServiceInternetService(d, i["internet_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-ctrl"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCtrl(d, i["internet_service_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-ctrl-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup(d, i["internet_service_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl(d, i["internet_service_app_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d, i["internet_service_app_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCustom(d, i["internet_service_custom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup(d, i["internet_service_custom_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-group"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceGroup(d, i["internet_service_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-id"], _ = expandWantempSystemVirtualWanLinkServiceInternetServiceId(d, i["internet_service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-weight"], _ = expandWantempSystemVirtualWanLinkServiceJitterWeight(d, i["jitter_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-weight"], _ = expandWantempSystemVirtualWanLinkServiceLatencyWeight(d, i["latency_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemVirtualWanLinkServiceLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-threshold"], _ = expandWantempSystemVirtualWanLinkServiceLinkCostThreshold(d, i["link_cost_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandWantempSystemVirtualWanLinkServiceMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandWantempSystemVirtualWanLinkServiceMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemVirtualWanLinkServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-weight"], _ = expandWantempSystemVirtualWanLinkServicePacketLossWeight(d, i["packet_loss_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-members"], _ = expandWantempSystemVirtualWanLinkServicePriorityMembers(d, i["priority_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemVirtualWanLinkServiceProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-link"], _ = expandWantempSystemVirtualWanLinkServiceQualityLink(d, i["quality_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemVirtualWanLinkServiceRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandWantempSystemVirtualWanLinkServiceRouteTag(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemVirtualWanLinkServiceSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-compare-method"], _ = expandWantempSystemVirtualWanLinkServiceSlaCompareMethod(d, i["sla_compare_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandWantempSystemVirtualWanLinkServiceSrc(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-negate"], _ = expandWantempSystemVirtualWanLinkServiceSrcNegate(d, i["src_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src6"], _ = expandWantempSystemVirtualWanLinkServiceSrc6(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["standalone-action"], _ = expandWantempSystemVirtualWanLinkServiceStandaloneAction(d, i["standalone_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandWantempSystemVirtualWanLinkServiceStartPort(d, i["start_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemVirtualWanLinkServiceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandWantempSystemVirtualWanLinkServiceTos(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandWantempSystemVirtualWanLinkServiceTosMask(d, i["tos_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["users"], _ = expandWantempSystemVirtualWanLinkServiceUsers(d, i["users"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkServiceAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceBandwidthWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpForwardTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDscpReverseTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceDstNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceDst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceEndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInputDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCtrlGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceJitterWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLatencyWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceLinkCostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServicePacketLossWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceQualityLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemVirtualWanLinkServiceSlaHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemVirtualWanLinkServiceSlaId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSlaCompareMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceSrc6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkServiceStandaloneAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkServiceUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemVirtualWanLinkStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLink(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fail_alert_interfaces"); ok || d.HasChange("fail_alert_interfaces") {
		t, err := expandWantempSystemVirtualWanLinkFailAlertInterfaces(d, v, "fail_alert_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok || d.HasChange("fail_detect") {
		t, err := expandWantempSystemVirtualWanLinkFailDetect(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheck(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_mode"); ok || d.HasChange("load_balance_mode") {
		t, err := expandWantempSystemVirtualWanLinkLoadBalanceMode(d, v, "load_balance_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-mode"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandWantempSystemVirtualWanLinkMembers(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandWantempSystemVirtualWanLinkNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_boot_time"); ok || d.HasChange("neighbor_hold_boot_time") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHoldBootTime(d, v, "neighbor_hold_boot_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-boot-time"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok || d.HasChange("neighbor_hold_down") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHoldDown(d, v, "neighbor_hold_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down_time"); ok || d.HasChange("neighbor_hold_down_time") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHoldDownTime(d, v, "neighbor_hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandWantempSystemVirtualWanLinkService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemVirtualWanLinkStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
