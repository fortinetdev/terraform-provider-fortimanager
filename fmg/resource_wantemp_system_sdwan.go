// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure redundant Internet connections with multiple outbound links and health-check profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemSdwan() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanUpdate,
		Read:   resourceWantempSystemSdwanRead,
		Update: resourceWantempSystemSdwanUpdate,
		Delete: resourceWantempSystemSdwanDelete,

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
			"app_perf_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstintf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_de_duplication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_duplication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sla_match_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcintf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"duplication_max_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ingress_spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"preferred_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"seq_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"volume_ratio": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"zone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"minimum_sla_meet_members": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
					},
				},
			},
			"speedtest_bypass_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_sla_tie_break": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWantempSystemSdwanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwan(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwan resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemSdwan(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WantempSystemSdwan")

	return resourceWantempSystemSdwanRead(d, m)
}

func resourceWantempSystemSdwanDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemSdwan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemSdwan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwan resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanAppPerfLogPeriodWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenWantempSystemSdwanDuplicationDstaddrWssa(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {
			v := flattenWantempSystemSdwanDuplicationDstaddr6Wssa(i["dstaddr6"], d, pre_append)
			tmp["dstaddr6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Dstaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := i["dstintf"]; ok {
			v := flattenWantempSystemSdwanDuplicationDstintfWssa(i["dstintf"], d, pre_append)
			tmp["dstintf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Dstintf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanDuplicationIdWssa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := i["packet-de-duplication"]; ok {
			v := flattenWantempSystemSdwanDuplicationPacketDeDuplicationWssa(i["packet-de-duplication"], d, pre_append)
			tmp["packet_de_duplication"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-PacketDeDuplication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := i["packet-duplication"]; ok {
			v := flattenWantempSystemSdwanDuplicationPacketDuplicationWssa(i["packet-duplication"], d, pre_append)
			tmp["packet_duplication"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-PacketDuplication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenWantempSystemSdwanDuplicationServiceWssa(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenWantempSystemSdwanDuplicationServiceIdWssa(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-ServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_match_service"
		if _, ok := i["sla-match-service"]; ok {
			v := flattenWantempSystemSdwanDuplicationSlaMatchServiceWssa(i["sla-match-service"], d, pre_append)
			tmp["sla_match_service"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-SlaMatchService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWantempSystemSdwanDuplicationSrcaddrWssa(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {
			v := flattenWantempSystemSdwanDuplicationSrcaddr6Wssa(i["srcaddr6"], d, pre_append)
			tmp["srcaddr6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Srcaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := i["srcintf"]; ok {
			v := flattenWantempSystemSdwanDuplicationSrcintfWssa(i["srcintf"], d, pre_append)
			tmp["srcintf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Srcintf")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanDuplicationDstaddrWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationDstaddr6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationDstintfWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationPacketDeDuplicationWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationPacketDuplicationWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationServiceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationServiceIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationSlaMatchServiceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationSrcaddrWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationSrcaddr6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationSrcintfWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationMaxNumWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanFailAlertInterfacesWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanFailDetectWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckDynamicServerWssa(i["_dynamic-server"], d, pre_append)
			tmp["_dynamic_server"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DynamicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckAddrModeWssa(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenWantempSystemSdwanHealthCheckClassIdWssa(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := i["detect-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDetectModeWssa(i["detect-mode"], d, pre_append)
			tmp["detect_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DetectMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDiffservcodeWssa(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := i["dns-match-ip"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDnsMatchIpWssa(i["dns-match-ip"], d, pre_append)
			tmp["dns_match_ip"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DnsMatchIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := i["dns-request-domain"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDnsRequestDomainWssa(i["dns-request-domain"], d, pre_append)
			tmp["dns_request_domain"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DnsRequestDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := i["embed-measured-health"]; ok {
			v := flattenWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssa(i["embed-measured-health"], d, pre_append)
			tmp["embed_measured_health"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-EmbedMeasuredHealth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFailtimeWssa(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := i["ftp-file"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFtpFileWssa(i["ftp-file"], d, pre_append)
			tmp["ftp_file"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-FtpFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := i["ftp-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFtpModeWssa(i["ftp-mode"], d, pre_append)
			tmp["ftp_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-FtpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHaPriorityWssa(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHttpAgentWssa(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHttpGetWssa(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHttpMatchWssa(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenWantempSystemSdwanHealthCheckIntervalWssa(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenWantempSystemSdwanHealthCheckMembersWssa(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := i["mos-codec"]; ok {
			v := flattenWantempSystemSdwanHealthCheckMosCodecWssa(i["mos-codec"], d, pre_append)
			tmp["mos_codec"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-MosCodec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemSdwanHealthCheckNameWssa(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenWantempSystemSdwanHealthCheckPacketSizeWssa(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenWantempSystemSdwanHealthCheckPasswordWssa(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenWantempSystemSdwanHealthCheckPortWssa(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := i["probe-count"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProbeCountWssa(i["probe-count"], d, pre_append)
			tmp["probe_count"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ProbeCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProbePacketsWssa(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProbeTimeoutWssa(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProtocolWssa(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := i["quality-measured-method"]; ok {
			v := flattenWantempSystemSdwanHealthCheckQualityMeasuredMethodWssa(i["quality-measured-method"], d, pre_append)
			tmp["quality_measured_method"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-QualityMeasuredMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenWantempSystemSdwanHealthCheckRecoverytimeWssa(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSecurityModeWssa(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenWantempSystemSdwanHealthCheckServerWssa(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaWssa(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaFailLogPeriodWssa(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := i["sla-id-redistribute"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaIdRedistributeWssa(i["sla-id-redistribute"], d, pre_append)
			tmp["sla_id_redistribute"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SlaIdRedistribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPassLogPeriodWssa(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSourceWssa(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSource6Wssa(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := i["system-dns"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSystemDnsWssa(i["system-dns"], d, pre_append)
			tmp["system_dns"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SystemDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdAlertJitterWssa(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdAlertLatencyWssa(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssa(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdWarningJitterWssa(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdWarningLatencyWssa(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssa(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssa(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenWantempSystemSdwanHealthCheckUpdateStaticRouteWssa(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-UpdateStaticRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenWantempSystemSdwanHealthCheckUserWssa(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-User")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenWantempSystemSdwanHealthCheckVrfWssa(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Vrf")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckDynamicServerWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckAddrModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckClassIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDetectModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDiffservcodeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDnsMatchIpWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDnsRequestDomainWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFailtimeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFtpFileWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFtpModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHaPriorityWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHttpAgentWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWantempSystemSdwanHealthCheckHttpGetWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHttpMatchWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckIntervalWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckMembersWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckMosCodecWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckNameWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckPacketSizeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckPasswordWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckPortWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbeCountWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbePacketsWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbeTimeoutWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProtocolWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckQualityMeasuredMethodWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckRecoverytimeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSecurityModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckServerWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSlaWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckSlaIdWssa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaJitterThresholdWssa(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaLatencyThresholdWssa(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaLinkCostFactorWssa(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaMosThresholdWssa(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssa(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPriorityInSlaWssa(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssa(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PriorityOutSla")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckSlaIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaJitterThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLatencyThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLinkCostFactorWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSlaMosThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityInSlaWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaFailLogPeriodWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaIdRedistributeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPassLogPeriodWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSourceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSource6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSystemDnsWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertJitterWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertLatencyWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningJitterWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningLatencyWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUpdateStaticRouteWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUserWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckVrfWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanLoadBalanceModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanMembersDynamicMemberWssa(i["_dynamic-member"], d, pre_append)
			tmp["_dynamic_member"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-DynamicMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWantempSystemSdwanMembersCommentWssa(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenWantempSystemSdwanMembersCostWssa(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemSdwanMembersGatewayWssa(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenWantempSystemSdwanMembersGateway6Wssa(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {
			v := flattenWantempSystemSdwanMembersIngressSpilloverThresholdWssa(i["ingress-spillover-threshold"], d, pre_append)
			tmp["ingress_spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-IngressSpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenWantempSystemSdwanMembersInterfaceWssa(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_source"
		if _, ok := i["preferred-source"]; ok {
			v := flattenWantempSystemSdwanMembersPreferredSourceWssa(i["preferred-source"], d, pre_append)
			tmp["preferred_source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-PreferredSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenWantempSystemSdwanMembersPriorityWssa(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := i["priority6"]; ok {
			v := flattenWantempSystemSdwanMembersPriority6Wssa(i["priority6"], d, pre_append)
			tmp["priority6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Priority6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {
			v := flattenWantempSystemSdwanMembersSeqNumWssa(i["seq-num"], d, pre_append)
			tmp["seq_num"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-SeqNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemSdwanMembersSourceWssa(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemSdwanMembersSource6Wssa(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {
			v := flattenWantempSystemSdwanMembersSpilloverThresholdWssa(i["spillover-threshold"], d, pre_append)
			tmp["spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-SpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemSdwanMembersStatusWssa(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {
			v := flattenWantempSystemSdwanMembersVolumeRatioWssa(i["volume-ratio"], d, pre_append)
			tmp["volume_ratio"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-VolumeRatio")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenWantempSystemSdwanMembersWeightWssa(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Weight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := i["zone"]; ok {
			v := flattenWantempSystemSdwanMembersZoneWssa(i["zone"], d, pre_append)
			tmp["zone"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Zone")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanMembersDynamicMemberWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersCommentWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersCostWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGatewayWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGateway6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersIngressSpilloverThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersInterfaceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPreferredSourceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriorityWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSeqNumWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSourceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSpilloverThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersStatusWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersVolumeRatioWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersWeightWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersZoneWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanNeighborHealthCheckWssa(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWantempSystemSdwanNeighborIpWssa(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenWantempSystemSdwanNeighborMemberWssa(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenWantempSystemSdwanNeighborMinimumSlaMeetMembersWssa(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenWantempSystemSdwanNeighborModeWssa(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemSdwanNeighborRoleWssa(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := i["sla-id"]; ok {
			v := flattenWantempSystemSdwanNeighborSlaIdWssa(i["sla-id"], d, pre_append)
			tmp["sla_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-SlaId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanNeighborHealthCheckWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborIpWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborMemberWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborMinimumSlaMeetMembersWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborRoleWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborSlaIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborHoldBootTimeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborHoldDownWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborHoldDownTimeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanServiceAddrModeWssa(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_exclusive"
		if _, ok := i["agent-exclusive"]; ok {
			v := flattenWantempSystemSdwanServiceAgentExclusiveWssa(i["agent-exclusive"], d, pre_append)
			tmp["agent_exclusive"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-AgentExclusive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := i["bandwidth-weight"]; ok {
			v := flattenWantempSystemSdwanServiceBandwidthWeightWssa(i["bandwidth-weight"], d, pre_append)
			tmp["bandwidth_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-BandwidthWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenWantempSystemSdwanServiceDefaultWssa(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := i["dscp-forward"]; ok {
			v := flattenWantempSystemSdwanServiceDscpForwardWssa(i["dscp-forward"], d, pre_append)
			tmp["dscp_forward"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpForward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := i["dscp-forward-tag"]; ok {
			v := flattenWantempSystemSdwanServiceDscpForwardTagWssa(i["dscp-forward-tag"], d, pre_append)
			tmp["dscp_forward_tag"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpForwardTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := i["dscp-reverse"]; ok {
			v := flattenWantempSystemSdwanServiceDscpReverseWssa(i["dscp-reverse"], d, pre_append)
			tmp["dscp_reverse"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := i["dscp-reverse-tag"]; ok {
			v := flattenWantempSystemSdwanServiceDscpReverseTagWssa(i["dscp-reverse-tag"], d, pre_append)
			tmp["dscp_reverse_tag"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpReverseTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenWantempSystemSdwanServiceDstWssa(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := i["dst-negate"]; ok {
			v := flattenWantempSystemSdwanServiceDstNegateWssa(i["dst-negate"], d, pre_append)
			tmp["dst_negate"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DstNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenWantempSystemSdwanServiceDst6Wssa(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {
			v := flattenWantempSystemSdwanServiceEndPortWssa(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemSdwanServiceGatewayWssa(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenWantempSystemSdwanServiceGroupsWssa(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := i["hash-mode"]; ok {
			v := flattenWantempSystemSdwanServiceHashModeWssa(i["hash-mode"], d, pre_append)
			tmp["hash_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-HashMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenWantempSystemSdwanServiceHealthCheckWssa(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := i["hold-down-time"]; ok {
			v := flattenWantempSystemSdwanServiceHoldDownTimeWssa(i["hold-down-time"], d, pre_append)
			tmp["hold_down_time"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-HoldDownTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanServiceIdWssa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := i["input-device"]; ok {
			v := flattenWantempSystemSdwanServiceInputDeviceWssa(i["input-device"], d, pre_append)
			tmp["input_device"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InputDevice")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := i["input-device-negate"]; ok {
			v := flattenWantempSystemSdwanServiceInputDeviceNegateWssa(i["input-device-negate"], d, pre_append)
			tmp["input_device_negate"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InputDeviceNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_zone"
		if _, ok := i["input-zone"]; ok {
			v := flattenWantempSystemSdwanServiceInputZoneWssa(i["input-zone"], d, pre_append)
			tmp["input_zone"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InputZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := i["internet-service"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceWssa(i["internet-service"], d, pre_append)
			tmp["internet_service"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := i["internet-service-app-ctrl"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceAppCtrlWssa(i["internet-service-app-ctrl"], d, pre_append)
			tmp["internet_service_app_ctrl"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceAppCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_category"
		if _, ok := i["internet-service-app-ctrl-category"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWssa(i["internet-service-app-ctrl-category"], d, pre_append)
			tmp["internet_service_app_ctrl_category"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceAppCtrlCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := i["internet-service-app-ctrl-group"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceAppCtrlGroupWssa(i["internet-service-app-ctrl-group"], d, pre_append)
			tmp["internet_service_app_ctrl_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceAppCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := i["internet-service-custom"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceCustomWssa(i["internet-service-custom"], d, pre_append)
			tmp["internet_service_custom"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceCustom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := i["internet-service-custom-group"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceCustomGroupWssa(i["internet-service-custom-group"], d, pre_append)
			tmp["internet_service_custom_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceCustomGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := i["internet-service-group"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceGroupWssa(i["internet-service-group"], d, pre_append)
			tmp["internet_service_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := i["internet-service-name"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceNameWssa(i["internet-service-name"], d, pre_append)
			tmp["internet_service_name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := i["jitter-weight"]; ok {
			v := flattenWantempSystemSdwanServiceJitterWeightWssa(i["jitter-weight"], d, pre_append)
			tmp["jitter_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-JitterWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := i["latency-weight"]; ok {
			v := flattenWantempSystemSdwanServiceLatencyWeightWssa(i["latency-weight"], d, pre_append)
			tmp["latency_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LatencyWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemSdwanServiceLinkCostFactorWssa(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := i["link-cost-threshold"]; ok {
			v := flattenWantempSystemSdwanServiceLinkCostThresholdWssa(i["link-cost-threshold"], d, pre_append)
			tmp["link_cost_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LinkCostThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenWantempSystemSdwanServiceMinimumSlaMeetMembersWssa(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenWantempSystemSdwanServiceModeWssa(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemSdwanServiceNameWssa(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := i["packet-loss-weight"]; ok {
			v := flattenWantempSystemSdwanServicePacketLossWeightWssa(i["packet-loss-weight"], d, pre_append)
			tmp["packet_loss_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PacketLossWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := i["passive-measurement"]; ok {
			v := flattenWantempSystemSdwanServicePassiveMeasurementWssa(i["passive-measurement"], d, pre_append)
			tmp["passive_measurement"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PassiveMeasurement")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := i["priority-members"]; ok {
			v := flattenWantempSystemSdwanServicePriorityMembersWssa(i["priority-members"], d, pre_append)
			tmp["priority_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PriorityMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := i["priority-zone"]; ok {
			v := flattenWantempSystemSdwanServicePriorityZoneWssa(i["priority-zone"], d, pre_append)
			tmp["priority_zone"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PriorityZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemSdwanServiceProtocolWssa(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := i["quality-link"]; ok {
			v := flattenWantempSystemSdwanServiceQualityLinkWssa(i["quality-link"], d, pre_append)
			tmp["quality_link"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-QualityLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemSdwanServiceRoleWssa(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			v := flattenWantempSystemSdwanServiceShortcutWssa(i["shortcut"], d, pre_append)
			tmp["shortcut"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Shortcut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_stickiness"
		if _, ok := i["shortcut-stickiness"]; ok {
			v := flattenWantempSystemSdwanServiceShortcutStickinessWssa(i["shortcut-stickiness"], d, pre_append)
			tmp["shortcut_stickiness"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-ShortcutStickiness")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenWantempSystemSdwanServiceRouteTagWssa(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemSdwanServiceSlaWssa(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := i["sla-compare-method"]; ok {
			v := flattenWantempSystemSdwanServiceSlaCompareMethodWssa(i["sla-compare-method"], d, pre_append)
			tmp["sla_compare_method"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-SlaCompareMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenWantempSystemSdwanServiceSrcWssa(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := i["src-negate"]; ok {
			v := flattenWantempSystemSdwanServiceSrcNegateWssa(i["src-negate"], d, pre_append)
			tmp["src_negate"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-SrcNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenWantempSystemSdwanServiceSrc6Wssa(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := i["standalone-action"]; ok {
			v := flattenWantempSystemSdwanServiceStandaloneActionWssa(i["standalone-action"], d, pre_append)
			tmp["standalone_action"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-StandaloneAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenWantempSystemSdwanServiceStartPortWssa(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-StartPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemSdwanServiceStatusWssa(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := i["tie-break"]; ok {
			v := flattenWantempSystemSdwanServiceTieBreakWssa(i["tie-break"], d, pre_append)
			tmp["tie_break"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-TieBreak")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenWantempSystemSdwanServiceTosWssa(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenWantempSystemSdwanServiceTosMaskWssa(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-TosMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := i["use-shortcut-sla"]; ok {
			v := flattenWantempSystemSdwanServiceUseShortcutSlaWssa(i["use-shortcut-sla"], d, pre_append)
			tmp["use_shortcut_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-UseShortcutSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			v := flattenWantempSystemSdwanServiceUsersWssa(i["users"], d, pre_append)
			tmp["users"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Users")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanServiceAddrModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceAgentExclusiveWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceBandwidthWeightWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDefaultWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpForwardWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpForwardTagWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpReverseWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpReverseTagWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDstWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDstNegateWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDst6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceEndPortWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceGatewayWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceGroupsWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHashModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHealthCheckWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHoldDownTimeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputDeviceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputDeviceNegateWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputZoneWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlGroupWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceCustomWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceCustomGroupWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceGroupWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceNameWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceJitterWeightWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLatencyWeightWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLinkCostFactorWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLinkCostThresholdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceMinimumSlaMeetMembersWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceModeWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceNameWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePacketLossWeightWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePassiveMeasurementWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePriorityMembersWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePriorityZoneWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceProtocolWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceQualityLinkWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceRoleWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcutWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcutStickinessWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceRouteTagWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanServiceSlaHealthCheckWssa(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwanService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanServiceSlaIdWssa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanService-Sla-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanServiceSlaHealthCheckWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaIdWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaCompareMethodWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrcWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrcNegateWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrc6Wssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStandaloneActionWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStartPortWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStatusWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTieBreakWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTosWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTosMaskWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceUseShortcutSlaWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceUsersWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanSpeedtestBypassRoutingWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanStatusWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZoneWssa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemSdwanZoneNameWssa(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := i["service-sla-tie-break"]; ok {
			v := flattenWantempSystemSdwanZoneServiceSlaTieBreakWssa(i["service-sla-tie-break"], d, pre_append)
			tmp["service_sla_tie_break"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-ServiceSlaTieBreak")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanZoneNameWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZoneServiceSlaTieBreakWssa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("app_perf_log_period", flattenWantempSystemSdwanAppPerfLogPeriodWssa(o["app-perf-log-period"], d, "app_perf_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-perf-log-period"], "WantempSystemSdwan-AppPerfLogPeriod"); ok {
			if err = d.Set("app_perf_log_period", vv); err != nil {
				return fmt.Errorf("Error reading app_perf_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_perf_log_period: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("duplication", flattenWantempSystemSdwanDuplicationWssa(o["duplication"], d, "duplication")); err != nil {
			if vv, ok := fortiAPIPatch(o["duplication"], "WantempSystemSdwan-Duplication"); ok {
				if err = d.Set("duplication", vv); err != nil {
					return fmt.Errorf("Error reading duplication: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading duplication: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("duplication"); ok {
			if err = d.Set("duplication", flattenWantempSystemSdwanDuplicationWssa(o["duplication"], d, "duplication")); err != nil {
				if vv, ok := fortiAPIPatch(o["duplication"], "WantempSystemSdwan-Duplication"); ok {
					if err = d.Set("duplication", vv); err != nil {
						return fmt.Errorf("Error reading duplication: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading duplication: %v", err)
				}
			}
		}
	}

	if err = d.Set("duplication_max_num", flattenWantempSystemSdwanDuplicationMaxNumWssa(o["duplication-max-num"], d, "duplication_max_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-max-num"], "WantempSystemSdwan-DuplicationMaxNum"); ok {
			if err = d.Set("duplication_max_num", vv); err != nil {
				return fmt.Errorf("Error reading duplication_max_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_max_num: %v", err)
		}
	}

	if err = d.Set("fail_alert_interfaces", flattenWantempSystemSdwanFailAlertInterfacesWssa(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-alert-interfaces"], "WantempSystemSdwan-FailAlertInterfaces"); ok {
			if err = d.Set("fail_alert_interfaces", vv); err != nil {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenWantempSystemSdwanFailDetectWssa(o["fail-detect"], d, "fail_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect"], "WantempSystemSdwan-FailDetect"); ok {
			if err = d.Set("fail_detect", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenWantempSystemSdwanHealthCheckWssa(o["health-check"], d, "health_check")); err != nil {
			if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemSdwan-HealthCheck"); ok {
				if err = d.Set("health_check", vv); err != nil {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check"); ok {
			if err = d.Set("health_check", flattenWantempSystemSdwanHealthCheckWssa(o["health-check"], d, "health_check")); err != nil {
				if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemSdwan-HealthCheck"); ok {
					if err = d.Set("health_check", vv); err != nil {
						return fmt.Errorf("Error reading health_check: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			}
		}
	}

	if err = d.Set("load_balance_mode", flattenWantempSystemSdwanLoadBalanceModeWssa(o["load-balance-mode"], d, "load_balance_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-mode"], "WantempSystemSdwan-LoadBalanceMode"); ok {
			if err = d.Set("load_balance_mode", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenWantempSystemSdwanMembersWssa(o["members"], d, "members")); err != nil {
			if vv, ok := fortiAPIPatch(o["members"], "WantempSystemSdwan-Members"); ok {
				if err = d.Set("members", vv); err != nil {
					return fmt.Errorf("Error reading members: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenWantempSystemSdwanMembersWssa(o["members"], d, "members")); err != nil {
				if vv, ok := fortiAPIPatch(o["members"], "WantempSystemSdwan-Members"); ok {
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
		if err = d.Set("neighbor", flattenWantempSystemSdwanNeighborWssa(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "WantempSystemSdwan-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenWantempSystemSdwanNeighborWssa(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "WantempSystemSdwan-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if err = d.Set("neighbor_hold_boot_time", flattenWantempSystemSdwanNeighborHoldBootTimeWssa(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-boot-time"], "WantempSystemSdwan-NeighborHoldBootTime"); ok {
			if err = d.Set("neighbor_hold_boot_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenWantempSystemSdwanNeighborHoldDownWssa(o["neighbor-hold-down"], d, "neighbor_hold_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down"], "WantempSystemSdwan-NeighborHoldDown"); ok {
			if err = d.Set("neighbor_hold_down", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenWantempSystemSdwanNeighborHoldDownTimeWssa(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down-time"], "WantempSystemSdwan-NeighborHoldDownTime"); ok {
			if err = d.Set("neighbor_hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenWantempSystemSdwanServiceWssa(o["service"], d, "service")); err != nil {
			if vv, ok := fortiAPIPatch(o["service"], "WantempSystemSdwan-Service"); ok {
				if err = d.Set("service", vv); err != nil {
					return fmt.Errorf("Error reading service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenWantempSystemSdwanServiceWssa(o["service"], d, "service")); err != nil {
				if vv, ok := fortiAPIPatch(o["service"], "WantempSystemSdwan-Service"); ok {
					if err = d.Set("service", vv); err != nil {
						return fmt.Errorf("Error reading service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("speedtest_bypass_routing", flattenWantempSystemSdwanSpeedtestBypassRoutingWssa(o["speedtest-bypass-routing"], d, "speedtest_bypass_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["speedtest-bypass-routing"], "WantempSystemSdwan-SpeedtestBypassRouting"); ok {
			if err = d.Set("speedtest_bypass_routing", vv); err != nil {
				return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemSdwanStatusWssa(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemSdwan-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("zone", flattenWantempSystemSdwanZoneWssa(o["zone"], d, "zone")); err != nil {
			if vv, ok := fortiAPIPatch(o["zone"], "WantempSystemSdwan-Zone"); ok {
				if err = d.Set("zone", vv); err != nil {
					return fmt.Errorf("Error reading zone: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("zone"); ok {
			if err = d.Set("zone", flattenWantempSystemSdwanZoneWssa(o["zone"], d, "zone")); err != nil {
				if vv, ok := fortiAPIPatch(o["zone"], "WantempSystemSdwan-Zone"); ok {
					if err = d.Set("zone", vv); err != nil {
						return fmt.Errorf("Error reading zone: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading zone: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWantempSystemSdwanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanAppPerfLogPeriodWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandWantempSystemSdwanDuplicationDstaddrWssa(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr6"], _ = expandWantempSystemSdwanDuplicationDstaddr6Wssa(d, i["dstaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstintf"], _ = expandWantempSystemSdwanDuplicationDstintfWssa(d, i["dstintf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanDuplicationIdWssa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-de-duplication"], _ = expandWantempSystemSdwanDuplicationPacketDeDuplicationWssa(d, i["packet_de_duplication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-duplication"], _ = expandWantempSystemSdwanDuplicationPacketDuplicationWssa(d, i["packet_duplication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandWantempSystemSdwanDuplicationServiceWssa(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandWantempSystemSdwanDuplicationServiceIdWssa(d, i["service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_match_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-match-service"], _ = expandWantempSystemSdwanDuplicationSlaMatchServiceWssa(d, i["sla_match_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWantempSystemSdwanDuplicationSrcaddrWssa(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr6"], _ = expandWantempSystemSdwanDuplicationSrcaddr6Wssa(d, i["srcaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcintf"], _ = expandWantempSystemSdwanDuplicationSrcintfWssa(d, i["srcintf"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanDuplicationDstaddrWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationDstaddr6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationDstintfWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationPacketDeDuplicationWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationPacketDuplicationWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationServiceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationServiceIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationSlaMatchServiceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationSrcaddrWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationSrcaddr6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationSrcintfWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationMaxNumWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanFailAlertInterfacesWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanFailDetectWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_dynamic-server"], _ = expandWantempSystemSdwanHealthCheckDynamicServerWssa(d, i["_dynamic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandWantempSystemSdwanHealthCheckAddrModeWssa(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandWantempSystemSdwanHealthCheckClassIdWssa(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detect-mode"], _ = expandWantempSystemSdwanHealthCheckDetectModeWssa(d, i["detect_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandWantempSystemSdwanHealthCheckDiffservcodeWssa(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-match-ip"], _ = expandWantempSystemSdwanHealthCheckDnsMatchIpWssa(d, i["dns_match_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-request-domain"], _ = expandWantempSystemSdwanHealthCheckDnsRequestDomainWssa(d, i["dns_request_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embed-measured-health"], _ = expandWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssa(d, i["embed_measured_health"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandWantempSystemSdwanHealthCheckFailtimeWssa(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-file"], _ = expandWantempSystemSdwanHealthCheckFtpFileWssa(d, i["ftp_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-mode"], _ = expandWantempSystemSdwanHealthCheckFtpModeWssa(d, i["ftp_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandWantempSystemSdwanHealthCheckHaPriorityWssa(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandWantempSystemSdwanHealthCheckHttpAgentWssa(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandWantempSystemSdwanHealthCheckHttpGetWssa(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandWantempSystemSdwanHealthCheckHttpMatchWssa(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandWantempSystemSdwanHealthCheckIntervalWssa(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandWantempSystemSdwanHealthCheckMembersWssa(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-codec"], _ = expandWantempSystemSdwanHealthCheckMosCodecWssa(d, i["mos_codec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemSdwanHealthCheckNameWssa(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandWantempSystemSdwanHealthCheckPacketSizeWssa(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandWantempSystemSdwanHealthCheckPasswordWssa(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandWantempSystemSdwanHealthCheckPortWssa(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-count"], _ = expandWantempSystemSdwanHealthCheckProbeCountWssa(d, i["probe_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandWantempSystemSdwanHealthCheckProbePacketsWssa(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandWantempSystemSdwanHealthCheckProbeTimeoutWssa(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemSdwanHealthCheckProtocolWssa(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-measured-method"], _ = expandWantempSystemSdwanHealthCheckQualityMeasuredMethodWssa(d, i["quality_measured_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandWantempSystemSdwanHealthCheckRecoverytimeWssa(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandWantempSystemSdwanHealthCheckSecurityModeWssa(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandWantempSystemSdwanHealthCheckServerWssa(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemSdwanHealthCheckSlaWssa(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandWantempSystemSdwanHealthCheckSlaFailLogPeriodWssa(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id-redistribute"], _ = expandWantempSystemSdwanHealthCheckSlaIdRedistributeWssa(d, i["sla_id_redistribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandWantempSystemSdwanHealthCheckSlaPassLogPeriodWssa(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemSdwanHealthCheckSourceWssa(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemSdwanHealthCheckSource6Wssa(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["system-dns"], _ = expandWantempSystemSdwanHealthCheckSystemDnsWssa(d, i["system_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandWantempSystemSdwanHealthCheckThresholdAlertJitterWssa(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandWantempSystemSdwanHealthCheckThresholdAlertLatencyWssa(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssa(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandWantempSystemSdwanHealthCheckThresholdWarningJitterWssa(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandWantempSystemSdwanHealthCheckThresholdWarningLatencyWssa(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssa(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssa(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandWantempSystemSdwanHealthCheckUpdateStaticRouteWssa(d, i["update_static_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandWantempSystemSdwanHealthCheckUserWssa(d, i["user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandWantempSystemSdwanHealthCheckVrfWssa(d, i["vrf"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckDynamicServerWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckAddrModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckClassIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDetectModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDiffservcodeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDnsMatchIpWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDnsRequestDomainWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckEmbedMeasuredHealthWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFailtimeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFtpFileWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFtpModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHaPriorityWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpAgentWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpGetWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpMatchWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckIntervalWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckMembersWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckMosCodecWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckNameWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckPacketSizeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckPasswordWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckPortWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbeCountWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbePacketsWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbeTimeoutWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProtocolWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckQualityMeasuredMethodWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckRecoverytimeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSecurityModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckServerWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSlaWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemSdwanHealthCheckSlaIdWssa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaJitterThresholdWssa(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaLatencyThresholdWssa(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemSdwanHealthCheckSlaLinkCostFactorWssa(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaMosThresholdWssa(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssa(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandWantempSystemSdwanHealthCheckSlaPriorityInSlaWssa(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssa(d, i["priority_out_sla"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckSlaIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaJitterThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLatencyThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLinkCostFactorWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSlaMosThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPacketlossThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityInSlaWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityOutSlaWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaFailLogPeriodWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaIdRedistributeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPassLogPeriodWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSourceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSource6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSystemDnsWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertJitterWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertLatencyWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertPacketlossWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningJitterWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningLatencyWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningPacketlossWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUpdateCascadeInterfaceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUpdateStaticRouteWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUserWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckVrfWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanLoadBalanceModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_dynamic-member"], _ = expandWantempSystemSdwanMembersDynamicMemberWssa(d, i["_dynamic_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWantempSystemSdwanMembersCommentWssa(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandWantempSystemSdwanMembersCostWssa(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemSdwanMembersGatewayWssa(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway6"], _ = expandWantempSystemSdwanMembersGateway6Wssa(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-spillover-threshold"], _ = expandWantempSystemSdwanMembersIngressSpilloverThresholdWssa(d, i["ingress_spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandWantempSystemSdwanMembersInterfaceWssa(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-source"], _ = expandWantempSystemSdwanMembersPreferredSourceWssa(d, i["preferred_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandWantempSystemSdwanMembersPriorityWssa(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority6"], _ = expandWantempSystemSdwanMembersPriority6Wssa(d, i["priority6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq-num"], _ = expandWantempSystemSdwanMembersSeqNumWssa(d, i["seq_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemSdwanMembersSourceWssa(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemSdwanMembersSource6Wssa(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spillover-threshold"], _ = expandWantempSystemSdwanMembersSpilloverThresholdWssa(d, i["spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemSdwanMembersStatusWssa(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["volume-ratio"], _ = expandWantempSystemSdwanMembersVolumeRatioWssa(d, i["volume_ratio"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandWantempSystemSdwanMembersWeightWssa(d, i["weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["zone"], _ = expandWantempSystemSdwanMembersZoneWssa(d, i["zone"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanMembersDynamicMemberWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersCommentWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersCostWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGatewayWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersIngressSpilloverThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersInterfaceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPreferredSourceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriorityWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSeqNumWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSourceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSpilloverThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersStatusWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersVolumeRatioWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersWeightWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersZoneWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemSdwanNeighborHealthCheckWssa(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWantempSystemSdwanNeighborIpWssa(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandWantempSystemSdwanNeighborMemberWssa(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandWantempSystemSdwanNeighborMinimumSlaMeetMembersWssa(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandWantempSystemSdwanNeighborModeWssa(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemSdwanNeighborRoleWssa(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id"], _ = expandWantempSystemSdwanNeighborSlaIdWssa(d, i["sla_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanNeighborHealthCheckWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborIpWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborMemberWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborMinimumSlaMeetMembersWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborRoleWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborSlaIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborHoldBootTimeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborHoldDownWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborHoldDownTimeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-mode"], _ = expandWantempSystemSdwanServiceAddrModeWssa(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_exclusive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["agent-exclusive"], _ = expandWantempSystemSdwanServiceAgentExclusiveWssa(d, i["agent_exclusive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-weight"], _ = expandWantempSystemSdwanServiceBandwidthWeightWssa(d, i["bandwidth_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandWantempSystemSdwanServiceDefaultWssa(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward"], _ = expandWantempSystemSdwanServiceDscpForwardWssa(d, i["dscp_forward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward-tag"], _ = expandWantempSystemSdwanServiceDscpForwardTagWssa(d, i["dscp_forward_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse"], _ = expandWantempSystemSdwanServiceDscpReverseWssa(d, i["dscp_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse-tag"], _ = expandWantempSystemSdwanServiceDscpReverseTagWssa(d, i["dscp_reverse_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandWantempSystemSdwanServiceDstWssa(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-negate"], _ = expandWantempSystemSdwanServiceDstNegateWssa(d, i["dst_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst6"], _ = expandWantempSystemSdwanServiceDst6Wssa(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-port"], _ = expandWantempSystemSdwanServiceEndPortWssa(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemSdwanServiceGatewayWssa(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandWantempSystemSdwanServiceGroupsWssa(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hash-mode"], _ = expandWantempSystemSdwanServiceHashModeWssa(d, i["hash_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandWantempSystemSdwanServiceHealthCheckWssa(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hold-down-time"], _ = expandWantempSystemSdwanServiceHoldDownTimeWssa(d, i["hold_down_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanServiceIdWssa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device"], _ = expandWantempSystemSdwanServiceInputDeviceWssa(d, i["input_device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device-negate"], _ = expandWantempSystemSdwanServiceInputDeviceNegateWssa(d, i["input_device_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-zone"], _ = expandWantempSystemSdwanServiceInputZoneWssa(d, i["input_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service"], _ = expandWantempSystemSdwanServiceInternetServiceWssa(d, i["internet_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl"], _ = expandWantempSystemSdwanServiceInternetServiceAppCtrlWssa(d, i["internet_service_app_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-category"], _ = expandWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWssa(d, i["internet_service_app_ctrl_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-group"], _ = expandWantempSystemSdwanServiceInternetServiceAppCtrlGroupWssa(d, i["internet_service_app_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom"], _ = expandWantempSystemSdwanServiceInternetServiceCustomWssa(d, i["internet_service_custom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom-group"], _ = expandWantempSystemSdwanServiceInternetServiceCustomGroupWssa(d, i["internet_service_custom_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-group"], _ = expandWantempSystemSdwanServiceInternetServiceGroupWssa(d, i["internet_service_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-name"], _ = expandWantempSystemSdwanServiceInternetServiceNameWssa(d, i["internet_service_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-weight"], _ = expandWantempSystemSdwanServiceJitterWeightWssa(d, i["jitter_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-weight"], _ = expandWantempSystemSdwanServiceLatencyWeightWssa(d, i["latency_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemSdwanServiceLinkCostFactorWssa(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-threshold"], _ = expandWantempSystemSdwanServiceLinkCostThresholdWssa(d, i["link_cost_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandWantempSystemSdwanServiceMinimumSlaMeetMembersWssa(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandWantempSystemSdwanServiceModeWssa(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemSdwanServiceNameWssa(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-weight"], _ = expandWantempSystemSdwanServicePacketLossWeightWssa(d, i["packet_loss_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passive-measurement"], _ = expandWantempSystemSdwanServicePassiveMeasurementWssa(d, i["passive_measurement"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-members"], _ = expandWantempSystemSdwanServicePriorityMembersWssa(d, i["priority_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-zone"], _ = expandWantempSystemSdwanServicePriorityZoneWssa(d, i["priority_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemSdwanServiceProtocolWssa(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-link"], _ = expandWantempSystemSdwanServiceQualityLinkWssa(d, i["quality_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemSdwanServiceRoleWssa(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut"], _ = expandWantempSystemSdwanServiceShortcutWssa(d, i["shortcut"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_stickiness"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut-stickiness"], _ = expandWantempSystemSdwanServiceShortcutStickinessWssa(d, i["shortcut_stickiness"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandWantempSystemSdwanServiceRouteTagWssa(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemSdwanServiceSlaWssa(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-compare-method"], _ = expandWantempSystemSdwanServiceSlaCompareMethodWssa(d, i["sla_compare_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandWantempSystemSdwanServiceSrcWssa(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-negate"], _ = expandWantempSystemSdwanServiceSrcNegateWssa(d, i["src_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src6"], _ = expandWantempSystemSdwanServiceSrc6Wssa(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["standalone-action"], _ = expandWantempSystemSdwanServiceStandaloneActionWssa(d, i["standalone_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandWantempSystemSdwanServiceStartPortWssa(d, i["start_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemSdwanServiceStatusWssa(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tie-break"], _ = expandWantempSystemSdwanServiceTieBreakWssa(d, i["tie_break"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandWantempSystemSdwanServiceTosWssa(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandWantempSystemSdwanServiceTosMaskWssa(d, i["tos_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["use-shortcut-sla"], _ = expandWantempSystemSdwanServiceUseShortcutSlaWssa(d, i["use_shortcut_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["users"], _ = expandWantempSystemSdwanServiceUsersWssa(d, i["users"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanServiceAddrModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceAgentExclusiveWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceBandwidthWeightWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDefaultWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpForwardWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpForwardTagWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpReverseWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpReverseTagWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDstWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDstNegateWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDst6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceEndPortWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceGatewayWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceGroupsWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHashModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHealthCheckWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHoldDownTimeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputDeviceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputDeviceNegateWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputZoneWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlCategoryWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlGroupWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceCustomWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceCustomGroupWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceGroupWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceNameWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceJitterWeightWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLatencyWeightWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLinkCostFactorWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLinkCostThresholdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceMinimumSlaMeetMembersWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceModeWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceNameWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePacketLossWeightWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePassiveMeasurementWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePriorityMembersWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePriorityZoneWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceProtocolWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceQualityLinkWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceRoleWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcutWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcutStickinessWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceRouteTagWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemSdwanServiceSlaHealthCheckWssa(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanServiceSlaIdWssa(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanServiceSlaHealthCheckWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaIdWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaCompareMethodWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrcWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrcNegateWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrc6Wssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStandaloneActionWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStartPortWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStatusWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTieBreakWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTosWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTosMaskWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceUseShortcutSlaWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceUsersWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanSpeedtestBypassRoutingWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanStatusWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZoneWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemSdwanZoneNameWssa(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-sla-tie-break"], _ = expandWantempSystemSdwanZoneServiceSlaTieBreakWssa(d, i["service_sla_tie_break"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanZoneNameWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZoneServiceSlaTieBreakWssa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_perf_log_period"); ok || d.HasChange("app_perf_log_period") {
		t, err := expandWantempSystemSdwanAppPerfLogPeriodWssa(d, v, "app_perf_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-perf-log-period"] = t
		}
	}

	if v, ok := d.GetOk("duplication"); ok || d.HasChange("duplication") {
		t, err := expandWantempSystemSdwanDuplicationWssa(d, v, "duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication"] = t
		}
	}

	if v, ok := d.GetOk("duplication_max_num"); ok || d.HasChange("duplication_max_num") {
		t, err := expandWantempSystemSdwanDuplicationMaxNumWssa(d, v, "duplication_max_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-max-num"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok || d.HasChange("fail_alert_interfaces") {
		t, err := expandWantempSystemSdwanFailAlertInterfacesWssa(d, v, "fail_alert_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok || d.HasChange("fail_detect") {
		t, err := expandWantempSystemSdwanFailDetectWssa(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemSdwanHealthCheckWssa(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_mode"); ok || d.HasChange("load_balance_mode") {
		t, err := expandWantempSystemSdwanLoadBalanceModeWssa(d, v, "load_balance_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-mode"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandWantempSystemSdwanMembersWssa(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandWantempSystemSdwanNeighborWssa(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_boot_time"); ok || d.HasChange("neighbor_hold_boot_time") {
		t, err := expandWantempSystemSdwanNeighborHoldBootTimeWssa(d, v, "neighbor_hold_boot_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-boot-time"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok || d.HasChange("neighbor_hold_down") {
		t, err := expandWantempSystemSdwanNeighborHoldDownWssa(d, v, "neighbor_hold_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down_time"); ok || d.HasChange("neighbor_hold_down_time") {
		t, err := expandWantempSystemSdwanNeighborHoldDownTimeWssa(d, v, "neighbor_hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandWantempSystemSdwanServiceWssa(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("speedtest_bypass_routing"); ok || d.HasChange("speedtest_bypass_routing") {
		t, err := expandWantempSystemSdwanSpeedtestBypassRoutingWssa(d, v, "speedtest_bypass_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtest-bypass-routing"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemSdwanStatusWssa(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		t, err := expandWantempSystemSdwanZoneWssa(d, v, "zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
