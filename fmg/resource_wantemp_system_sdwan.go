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
				Computed: true,
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
				Computed: true,
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
							Computed: true,
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
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mos_codec": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
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
										Computed: true,
									},
									"latency_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
										Computed: true,
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
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"system_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"health_check_fortiguard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"class_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"detect_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"diffservcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_match_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_request_domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"embed_measured_health": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"failtime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ftp_file": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ftp_mode": &schema.Schema{
							Type:     schema.TypeString,
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
						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mos_codec": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"packet_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"probe_count": &schema.Schema{
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
						"quality_measured_method": &schema.Schema{
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
						"target_name": &schema.Schema{
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
						},
						"update_static_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
				Computed: true,
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
							Computed: true,
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
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority_in_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_out_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"transport_group": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"service_id": &schema.Schema{
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
				Computed: true,
			},
			"neighbor_hold_down_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"agent_exclusive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_forward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_forward_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp_reverse": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"dst6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"end_src_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
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
							Computed: true,
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
							Computed: true,
						},
						"load_balance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"minimum_sla_meet_members": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
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
							Computed: true,
						},
						"shortcut": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shortcut_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shortcut_stickiness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"sla_stickiness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"standalone_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_src_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tie_break": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tos_mask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"use_shortcut_sla": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"users": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"zone_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"speedtest_bypass_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advpn_health_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"advpn_select": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"minimum_sla_meet_members": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
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
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
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

func flattenWantempSystemSdwanAppPerfLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanDuplicationDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {
			v := flattenWantempSystemSdwanDuplicationDstaddr6(i["dstaddr6"], d, pre_append)
			tmp["dstaddr6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Dstaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := i["dstintf"]; ok {
			v := flattenWantempSystemSdwanDuplicationDstintf(i["dstintf"], d, pre_append)
			tmp["dstintf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Dstintf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanDuplicationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := i["packet-de-duplication"]; ok {
			v := flattenWantempSystemSdwanDuplicationPacketDeDuplication(i["packet-de-duplication"], d, pre_append)
			tmp["packet_de_duplication"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-PacketDeDuplication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := i["packet-duplication"]; ok {
			v := flattenWantempSystemSdwanDuplicationPacketDuplication(i["packet-duplication"], d, pre_append)
			tmp["packet_duplication"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-PacketDuplication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenWantempSystemSdwanDuplicationService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenWantempSystemSdwanDuplicationServiceId(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-ServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_match_service"
		if _, ok := i["sla-match-service"]; ok {
			v := flattenWantempSystemSdwanDuplicationSlaMatchService(i["sla-match-service"], d, pre_append)
			tmp["sla_match_service"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-SlaMatchService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenWantempSystemSdwanDuplicationSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {
			v := flattenWantempSystemSdwanDuplicationSrcaddr6(i["srcaddr6"], d, pre_append)
			tmp["srcaddr6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Srcaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := i["srcintf"]; ok {
			v := flattenWantempSystemSdwanDuplicationSrcintf(i["srcintf"], d, pre_append)
			tmp["srcintf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Duplication-Srcintf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanDuplicationDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationPacketDeDuplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationPacketDuplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationSlaMatchService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanDuplicationSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanDuplicationMaxNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanFailDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheck(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckDynamicServer(i["_dynamic-server"], d, pre_append)
			tmp["_dynamic_server"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DynamicServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenWantempSystemSdwanHealthCheckClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := i["detect-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDetectMode(i["detect-mode"], d, pre_append)
			tmp["detect_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DetectMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDiffservcode(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := i["dns-match-ip"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDnsMatchIp(i["dns-match-ip"], d, pre_append)
			tmp["dns_match_ip"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DnsMatchIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := i["dns-request-domain"]; ok {
			v := flattenWantempSystemSdwanHealthCheckDnsRequestDomain(i["dns-request-domain"], d, pre_append)
			tmp["dns_request_domain"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-DnsRequestDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := i["embed-measured-health"]; ok {
			v := flattenWantempSystemSdwanHealthCheckEmbedMeasuredHealth(i["embed-measured-health"], d, pre_append)
			tmp["embed_measured_health"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-EmbedMeasuredHealth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFailtime(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := i["ftp-file"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFtpFile(i["ftp-file"], d, pre_append)
			tmp["ftp_file"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-FtpFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := i["ftp-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFtpMode(i["ftp-mode"], d, pre_append)
			tmp["ftp_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-FtpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHttpAgent(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHttpGet(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenWantempSystemSdwanHealthCheckHttpMatch(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenWantempSystemSdwanHealthCheckInterval(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenWantempSystemSdwanHealthCheckMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := i["mos-codec"]; ok {
			v := flattenWantempSystemSdwanHealthCheckMosCodec(i["mos-codec"], d, pre_append)
			tmp["mos_codec"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-MosCodec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemSdwanHealthCheckName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenWantempSystemSdwanHealthCheckPacketSize(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenWantempSystemSdwanHealthCheckPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := i["probe-count"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProbeCount(i["probe-count"], d, pre_append)
			tmp["probe_count"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ProbeCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProbePackets(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProbeTimeout(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemSdwanHealthCheckProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := i["quality-measured-method"]; ok {
			v := flattenWantempSystemSdwanHealthCheckQualityMeasuredMethod(i["quality-measured-method"], d, pre_append)
			tmp["quality_measured_method"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-QualityMeasuredMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenWantempSystemSdwanHealthCheckRecoverytime(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSecurityMode(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenWantempSystemSdwanHealthCheckServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaFailLogPeriod(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := i["sla-id-redistribute"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaIdRedistribute(i["sla-id-redistribute"], d, pre_append)
			tmp["sla_id_redistribute"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SlaIdRedistribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPassLogPeriod(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := i["system-dns"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSystemDns(i["system-dns"], d, pre_append)
			tmp["system_dns"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-SystemDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdAlertJitter(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdAlertLatency(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdAlertPacketloss(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdWarningJitter(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdWarningLatency(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenWantempSystemSdwanHealthCheckThresholdWarningPacketloss(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenWantempSystemSdwanHealthCheckUpdateCascadeInterface(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenWantempSystemSdwanHealthCheckUpdateStaticRoute(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-UpdateStaticRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenWantempSystemSdwanHealthCheckUser(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-User")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenWantempSystemSdwanHealthCheckVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheck-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckDynamicServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanHealthCheckDetectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckEmbedMeasuredHealth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFtpFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWantempSystemSdwanHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckMosCodec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbeCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaMosThreshold(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPriorityInSla(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckSlaPriorityOutSla(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheck-Sla-PriorityOutSla")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSlaMosThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityInSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityOutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaIdRedistribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSystemDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguard(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckFortiguardAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := i["detect-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardDetectMode(i["detect-mode"], d, pre_append)
			tmp["detect_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-DetectMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardDiffservcode(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := i["dns-match-ip"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardDnsMatchIp(i["dns-match-ip"], d, pre_append)
			tmp["dns_match_ip"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-DnsMatchIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := i["dns-request-domain"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardDnsRequestDomain(i["dns-request-domain"], d, pre_append)
			tmp["dns_request_domain"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-DnsRequestDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := i["embed-measured-health"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(i["embed-measured-health"], d, pre_append)
			tmp["embed_measured_health"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-EmbedMeasuredHealth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardFailtime(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := i["ftp-file"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardFtpFile(i["ftp-file"], d, pre_append)
			tmp["ftp_file"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-FtpFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := i["ftp-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardFtpMode(i["ftp-mode"], d, pre_append)
			tmp["ftp_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-FtpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardHttpAgent(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardHttpGet(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardHttpMatch(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardInterval(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := i["mos-codec"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardMosCodec(i["mos-codec"], d, pre_append)
			tmp["mos_codec"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-MosCodec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardPacketSize(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := i["probe-count"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardProbeCount(i["probe-count"], d, pre_append)
			tmp["probe_count"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ProbeCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardProbePackets(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardProbeTimeout(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := i["quality-measured-method"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(i["quality-measured-method"], d, pre_append)
			tmp["quality_measured_method"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-QualityMeasuredMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardRecoverytime(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSecurityMode(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := i["sla-id-redistribute"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaIdRedistribute(i["sla-id-redistribute"], d, pre_append)
			tmp["sla_id_redistribute"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-SlaIdRedistribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := i["system-dns"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSystemDns(i["system-dns"], d, pre_append)
			tmp["system_dns"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-SystemDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_name"
		if _, ok := i["target-name"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardTargetName(i["target-name"], d, pre_append)
			tmp["target_name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-TargetName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardThresholdAlertJitter(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardThresholdAlertLatency(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardThresholdWarningJitter(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardThresholdWarningLatency(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardUpdateStaticRoute(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-UpdateStaticRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardUser(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-User")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-HealthCheckFortiguard-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckFortiguardAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckFortiguardDetectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardFtpFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardFtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckFortiguardMosCodec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardProbeCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardProbePackets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckFortiguardSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaMosThreshold(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaPriorityInSla(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenWantempSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwanHealthCheckFortiguard-Sla-PriorityOutSla")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaMosThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaPriorityInSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaIdRedistribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardSystemDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardTargetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckFortiguardVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanLoadBalanceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanMembersDynamicMember(i["_dynamic-member"], d, pre_append)
			tmp["_dynamic_member"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-DynamicMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWantempSystemSdwanMembersComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenWantempSystemSdwanMembersCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemSdwanMembersGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenWantempSystemSdwanMembersGateway6(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {
			v := flattenWantempSystemSdwanMembersIngressSpilloverThreshold(i["ingress-spillover-threshold"], d, pre_append)
			tmp["ingress_spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-IngressSpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenWantempSystemSdwanMembersInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_source"
		if _, ok := i["preferred-source"]; ok {
			v := flattenWantempSystemSdwanMembersPreferredSource(i["preferred-source"], d, pre_append)
			tmp["preferred_source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-PreferredSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenWantempSystemSdwanMembersPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenWantempSystemSdwanMembersPriorityInSla(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenWantempSystemSdwanMembersPriorityOutSla(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-PriorityOutSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := i["priority6"]; ok {
			v := flattenWantempSystemSdwanMembersPriority6(i["priority6"], d, pre_append)
			tmp["priority6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Priority6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {
			v := flattenWantempSystemSdwanMembersSeqNum(i["seq-num"], d, pre_append)
			tmp["seq_num"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-SeqNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenWantempSystemSdwanMembersSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenWantempSystemSdwanMembersSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {
			v := flattenWantempSystemSdwanMembersSpilloverThreshold(i["spillover-threshold"], d, pre_append)
			tmp["spillover_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-SpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemSdwanMembersStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport_group"
		if _, ok := i["transport-group"]; ok {
			v := flattenWantempSystemSdwanMembersTransportGroup(i["transport-group"], d, pre_append)
			tmp["transport_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-TransportGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {
			v := flattenWantempSystemSdwanMembersVolumeRatio(i["volume-ratio"], d, pre_append)
			tmp["volume_ratio"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-VolumeRatio")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenWantempSystemSdwanMembersWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Weight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := i["zone"]; ok {
			v := flattenWantempSystemSdwanMembersZone(i["zone"], d, pre_append)
			tmp["zone"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Members-Zone")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanMembersDynamicMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenWantempSystemSdwanMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanMembersPreferredSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriorityInSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriorityOutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersPriority6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersTransportGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanMembersZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanNeighborHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWantempSystemSdwanNeighborIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenWantempSystemSdwanNeighborMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenWantempSystemSdwanNeighborMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenWantempSystemSdwanNeighborMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemSdwanNeighborRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenWantempSystemSdwanNeighborServiceId(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-ServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := i["sla-id"]; ok {
			v := flattenWantempSystemSdwanNeighborSlaId(i["sla-id"], d, pre_append)
			tmp["sla_id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Neighbor-SlaId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanNeighborMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanNeighborMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanNeighborSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborHoldBootTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborHoldDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanServiceAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_exclusive"
		if _, ok := i["agent-exclusive"]; ok {
			v := flattenWantempSystemSdwanServiceAgentExclusive(i["agent-exclusive"], d, pre_append)
			tmp["agent_exclusive"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-AgentExclusive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := i["bandwidth-weight"]; ok {
			v := flattenWantempSystemSdwanServiceBandwidthWeight(i["bandwidth-weight"], d, pre_append)
			tmp["bandwidth_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-BandwidthWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenWantempSystemSdwanServiceComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenWantempSystemSdwanServiceDefault(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := i["dscp-forward"]; ok {
			v := flattenWantempSystemSdwanServiceDscpForward(i["dscp-forward"], d, pre_append)
			tmp["dscp_forward"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpForward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := i["dscp-forward-tag"]; ok {
			v := flattenWantempSystemSdwanServiceDscpForwardTag(i["dscp-forward-tag"], d, pre_append)
			tmp["dscp_forward_tag"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpForwardTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := i["dscp-reverse"]; ok {
			v := flattenWantempSystemSdwanServiceDscpReverse(i["dscp-reverse"], d, pre_append)
			tmp["dscp_reverse"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := i["dscp-reverse-tag"]; ok {
			v := flattenWantempSystemSdwanServiceDscpReverseTag(i["dscp-reverse-tag"], d, pre_append)
			tmp["dscp_reverse_tag"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DscpReverseTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenWantempSystemSdwanServiceDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := i["dst-negate"]; ok {
			v := flattenWantempSystemSdwanServiceDstNegate(i["dst-negate"], d, pre_append)
			tmp["dst_negate"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-DstNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenWantempSystemSdwanServiceDst6(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {
			v := flattenWantempSystemSdwanServiceEndPort(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_src_port"
		if _, ok := i["end-src-port"]; ok {
			v := flattenWantempSystemSdwanServiceEndSrcPort(i["end-src-port"], d, pre_append)
			tmp["end_src_port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-EndSrcPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenWantempSystemSdwanServiceGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenWantempSystemSdwanServiceGroups(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := i["hash-mode"]; ok {
			v := flattenWantempSystemSdwanServiceHashMode(i["hash-mode"], d, pre_append)
			tmp["hash_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-HashMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenWantempSystemSdwanServiceHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := i["hold-down-time"]; ok {
			v := flattenWantempSystemSdwanServiceHoldDownTime(i["hold-down-time"], d, pre_append)
			tmp["hold_down_time"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-HoldDownTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanServiceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := i["input-device"]; ok {
			v := flattenWantempSystemSdwanServiceInputDevice(i["input-device"], d, pre_append)
			tmp["input_device"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InputDevice")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := i["input-device-negate"]; ok {
			v := flattenWantempSystemSdwanServiceInputDeviceNegate(i["input-device-negate"], d, pre_append)
			tmp["input_device_negate"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InputDeviceNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_zone"
		if _, ok := i["input-zone"]; ok {
			v := flattenWantempSystemSdwanServiceInputZone(i["input-zone"], d, pre_append)
			tmp["input_zone"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InputZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := i["internet-service"]; ok {
			v := flattenWantempSystemSdwanServiceInternetService(i["internet-service"], d, pre_append)
			tmp["internet_service"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := i["internet-service-app-ctrl"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceAppCtrl(i["internet-service-app-ctrl"], d, pre_append)
			tmp["internet_service_app_ctrl"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceAppCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_category"
		if _, ok := i["internet-service-app-ctrl-category"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceAppCtrlCategory(i["internet-service-app-ctrl-category"], d, pre_append)
			tmp["internet_service_app_ctrl_category"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceAppCtrlCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := i["internet-service-app-ctrl-group"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceAppCtrlGroup(i["internet-service-app-ctrl-group"], d, pre_append)
			tmp["internet_service_app_ctrl_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceAppCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := i["internet-service-custom"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceCustom(i["internet-service-custom"], d, pre_append)
			tmp["internet_service_custom"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceCustom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := i["internet-service-custom-group"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceCustomGroup(i["internet-service-custom-group"], d, pre_append)
			tmp["internet_service_custom_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceCustomGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := i["internet-service-group"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceGroup(i["internet-service-group"], d, pre_append)
			tmp["internet_service_group"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := i["internet-service-name"]; ok {
			v := flattenWantempSystemSdwanServiceInternetServiceName(i["internet-service-name"], d, pre_append)
			tmp["internet_service_name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-InternetServiceName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := i["jitter-weight"]; ok {
			v := flattenWantempSystemSdwanServiceJitterWeight(i["jitter-weight"], d, pre_append)
			tmp["jitter_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-JitterWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := i["latency-weight"]; ok {
			v := flattenWantempSystemSdwanServiceLatencyWeight(i["latency-weight"], d, pre_append)
			tmp["latency_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LatencyWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenWantempSystemSdwanServiceLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := i["link-cost-threshold"]; ok {
			v := flattenWantempSystemSdwanServiceLinkCostThreshold(i["link-cost-threshold"], d, pre_append)
			tmp["link_cost_threshold"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LinkCostThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balance"
		if _, ok := i["load-balance"]; ok {
			v := flattenWantempSystemSdwanServiceLoadBalance(i["load-balance"], d, pre_append)
			tmp["load_balance"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-LoadBalance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenWantempSystemSdwanServiceMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenWantempSystemSdwanServiceMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemSdwanServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := i["packet-loss-weight"]; ok {
			v := flattenWantempSystemSdwanServicePacketLossWeight(i["packet-loss-weight"], d, pre_append)
			tmp["packet_loss_weight"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PacketLossWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := i["passive-measurement"]; ok {
			v := flattenWantempSystemSdwanServicePassiveMeasurement(i["passive-measurement"], d, pre_append)
			tmp["passive_measurement"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PassiveMeasurement")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := i["priority-members"]; ok {
			v := flattenWantempSystemSdwanServicePriorityMembers(i["priority-members"], d, pre_append)
			tmp["priority_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PriorityMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := i["priority-zone"]; ok {
			v := flattenWantempSystemSdwanServicePriorityZone(i["priority-zone"], d, pre_append)
			tmp["priority_zone"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-PriorityZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenWantempSystemSdwanServiceProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := i["quality-link"]; ok {
			v := flattenWantempSystemSdwanServiceQualityLink(i["quality-link"], d, pre_append)
			tmp["quality_link"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-QualityLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenWantempSystemSdwanServiceRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			v := flattenWantempSystemSdwanServiceShortcut(i["shortcut"], d, pre_append)
			tmp["shortcut"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Shortcut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_priority"
		if _, ok := i["shortcut-priority"]; ok {
			v := flattenWantempSystemSdwanServiceShortcutPriority(i["shortcut-priority"], d, pre_append)
			tmp["shortcut_priority"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-ShortcutPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_stickiness"
		if _, ok := i["shortcut-stickiness"]; ok {
			v := flattenWantempSystemSdwanServiceShortcutStickiness(i["shortcut-stickiness"], d, pre_append)
			tmp["shortcut_stickiness"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-ShortcutStickiness")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenWantempSystemSdwanServiceRouteTag(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenWantempSystemSdwanServiceSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := i["sla-compare-method"]; ok {
			v := flattenWantempSystemSdwanServiceSlaCompareMethod(i["sla-compare-method"], d, pre_append)
			tmp["sla_compare_method"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-SlaCompareMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_stickiness"
		if _, ok := i["sla-stickiness"]; ok {
			v := flattenWantempSystemSdwanServiceSlaStickiness(i["sla-stickiness"], d, pre_append)
			tmp["sla_stickiness"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-SlaStickiness")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenWantempSystemSdwanServiceSrc(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := i["src-negate"]; ok {
			v := flattenWantempSystemSdwanServiceSrcNegate(i["src-negate"], d, pre_append)
			tmp["src_negate"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-SrcNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenWantempSystemSdwanServiceSrc6(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := i["standalone-action"]; ok {
			v := flattenWantempSystemSdwanServiceStandaloneAction(i["standalone-action"], d, pre_append)
			tmp["standalone_action"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-StandaloneAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenWantempSystemSdwanServiceStartPort(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-StartPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_src_port"
		if _, ok := i["start-src-port"]; ok {
			v := flattenWantempSystemSdwanServiceStartSrcPort(i["start-src-port"], d, pre_append)
			tmp["start_src_port"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-StartSrcPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenWantempSystemSdwanServiceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := i["tie-break"]; ok {
			v := flattenWantempSystemSdwanServiceTieBreak(i["tie-break"], d, pre_append)
			tmp["tie_break"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-TieBreak")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenWantempSystemSdwanServiceTos(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenWantempSystemSdwanServiceTosMask(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-TosMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := i["use-shortcut-sla"]; ok {
			v := flattenWantempSystemSdwanServiceUseShortcutSla(i["use-shortcut-sla"], d, pre_append)
			tmp["use_shortcut_sla"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-UseShortcutSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			v := flattenWantempSystemSdwanServiceUsers(i["users"], d, pre_append)
			tmp["users"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-Users")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone_mode"
		if _, ok := i["zone-mode"]; ok {
			v := flattenWantempSystemSdwanServiceZoneMode(i["zone-mode"], d, pre_append)
			tmp["zone_mode"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Service-ZoneMode")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanServiceAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceAgentExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceDstNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceDst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceEndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceEndSrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceHashMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInputZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanServiceInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenWantempSystemSdwanServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceLoadBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePassiveMeasurement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServicePriorityZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceQualityLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcutPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceShortcutStickiness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWantempSystemSdwanServiceSlaHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwanService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenWantempSystemSdwanServiceSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "WantempSystemSdwanService-Sla-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSlaStickiness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceSrc6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStartSrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTieBreak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceUseShortcutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanServiceUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanServiceZoneMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanSpeedtestBypassRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_health_check"
		if _, ok := i["advpn-health-check"]; ok {
			v := flattenWantempSystemSdwanZoneAdvpnHealthCheck(i["advpn-health-check"], d, pre_append)
			tmp["advpn_health_check"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-AdvpnHealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_select"
		if _, ok := i["advpn-select"]; ok {
			v := flattenWantempSystemSdwanZoneAdvpnSelect(i["advpn-select"], d, pre_append)
			tmp["advpn_select"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-AdvpnSelect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenWantempSystemSdwanZoneMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenWantempSystemSdwanZoneName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := i["service-sla-tie-break"]; ok {
			v := flattenWantempSystemSdwanZoneServiceSlaTieBreak(i["service-sla-tie-break"], d, pre_append)
			tmp["service_sla_tie_break"] = fortiAPISubPartPatch(v, "WantempSystemSdwan-Zone-ServiceSlaTieBreak")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWantempSystemSdwanZoneAdvpnHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenWantempSystemSdwanZoneAdvpnSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZoneMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZoneName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanZoneServiceSlaTieBreak(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("app_perf_log_period", flattenWantempSystemSdwanAppPerfLogPeriod(o["app-perf-log-period"], d, "app_perf_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-perf-log-period"], "WantempSystemSdwan-AppPerfLogPeriod"); ok {
			if err = d.Set("app_perf_log_period", vv); err != nil {
				return fmt.Errorf("Error reading app_perf_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_perf_log_period: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("duplication", flattenWantempSystemSdwanDuplication(o["duplication"], d, "duplication")); err != nil {
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
			if err = d.Set("duplication", flattenWantempSystemSdwanDuplication(o["duplication"], d, "duplication")); err != nil {
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

	if err = d.Set("duplication_max_num", flattenWantempSystemSdwanDuplicationMaxNum(o["duplication-max-num"], d, "duplication_max_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-max-num"], "WantempSystemSdwan-DuplicationMaxNum"); ok {
			if err = d.Set("duplication_max_num", vv); err != nil {
				return fmt.Errorf("Error reading duplication_max_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_max_num: %v", err)
		}
	}

	if err = d.Set("fail_alert_interfaces", flattenWantempSystemSdwanFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-alert-interfaces"], "WantempSystemSdwan-FailAlertInterfaces"); ok {
			if err = d.Set("fail_alert_interfaces", vv); err != nil {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenWantempSystemSdwanFailDetect(o["fail-detect"], d, "fail_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect"], "WantempSystemSdwan-FailDetect"); ok {
			if err = d.Set("fail_detect", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenWantempSystemSdwanHealthCheck(o["health-check"], d, "health_check")); err != nil {
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
			if err = d.Set("health_check", flattenWantempSystemSdwanHealthCheck(o["health-check"], d, "health_check")); err != nil {
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

	if isImportTable() {
		if err = d.Set("health_check_fortiguard", flattenWantempSystemSdwanHealthCheckFortiguard(o["health-check-fortiguard"], d, "health_check_fortiguard")); err != nil {
			if vv, ok := fortiAPIPatch(o["health-check-fortiguard"], "WantempSystemSdwan-HealthCheckFortiguard"); ok {
				if err = d.Set("health_check_fortiguard", vv); err != nil {
					return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check_fortiguard"); ok {
			if err = d.Set("health_check_fortiguard", flattenWantempSystemSdwanHealthCheckFortiguard(o["health-check-fortiguard"], d, "health_check_fortiguard")); err != nil {
				if vv, ok := fortiAPIPatch(o["health-check-fortiguard"], "WantempSystemSdwan-HealthCheckFortiguard"); ok {
					if err = d.Set("health_check_fortiguard", vv); err != nil {
						return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("load_balance_mode", flattenWantempSystemSdwanLoadBalanceMode(o["load-balance-mode"], d, "load_balance_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-mode"], "WantempSystemSdwan-LoadBalanceMode"); ok {
			if err = d.Set("load_balance_mode", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenWantempSystemSdwanMembers(o["members"], d, "members")); err != nil {
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
			if err = d.Set("members", flattenWantempSystemSdwanMembers(o["members"], d, "members")); err != nil {
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
		if err = d.Set("neighbor", flattenWantempSystemSdwanNeighbor(o["neighbor"], d, "neighbor")); err != nil {
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
			if err = d.Set("neighbor", flattenWantempSystemSdwanNeighbor(o["neighbor"], d, "neighbor")); err != nil {
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

	if err = d.Set("neighbor_hold_boot_time", flattenWantempSystemSdwanNeighborHoldBootTime(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-boot-time"], "WantempSystemSdwan-NeighborHoldBootTime"); ok {
			if err = d.Set("neighbor_hold_boot_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenWantempSystemSdwanNeighborHoldDown(o["neighbor-hold-down"], d, "neighbor_hold_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down"], "WantempSystemSdwan-NeighborHoldDown"); ok {
			if err = d.Set("neighbor_hold_down", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenWantempSystemSdwanNeighborHoldDownTime(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down-time"], "WantempSystemSdwan-NeighborHoldDownTime"); ok {
			if err = d.Set("neighbor_hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if err = d.Set("option", flattenWantempSystemSdwanOption(o["option"], d, "option")); err != nil {
		if vv, ok := fortiAPIPatch(o["option"], "WantempSystemSdwan-Option"); ok {
			if err = d.Set("option", vv); err != nil {
				return fmt.Errorf("Error reading option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenWantempSystemSdwanService(o["service"], d, "service")); err != nil {
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
			if err = d.Set("service", flattenWantempSystemSdwanService(o["service"], d, "service")); err != nil {
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

	if err = d.Set("speedtest_bypass_routing", flattenWantempSystemSdwanSpeedtestBypassRouting(o["speedtest-bypass-routing"], d, "speedtest_bypass_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["speedtest-bypass-routing"], "WantempSystemSdwan-SpeedtestBypassRouting"); ok {
			if err = d.Set("speedtest_bypass_routing", vv); err != nil {
				return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
		}
	}

	if err = d.Set("status", flattenWantempSystemSdwanStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WantempSystemSdwan-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("zone", flattenWantempSystemSdwanZone(o["zone"], d, "zone")); err != nil {
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
			if err = d.Set("zone", flattenWantempSystemSdwanZone(o["zone"], d, "zone")); err != nil {
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

func expandWantempSystemSdwanAppPerfLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dstaddr"], _ = expandWantempSystemSdwanDuplicationDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr6"], _ = expandWantempSystemSdwanDuplicationDstaddr6(d, i["dstaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstintf"], _ = expandWantempSystemSdwanDuplicationDstintf(d, i["dstintf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanDuplicationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-de-duplication"], _ = expandWantempSystemSdwanDuplicationPacketDeDuplication(d, i["packet_de_duplication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-duplication"], _ = expandWantempSystemSdwanDuplicationPacketDuplication(d, i["packet_duplication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandWantempSystemSdwanDuplicationService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandWantempSystemSdwanDuplicationServiceId(d, i["service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_match_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-match-service"], _ = expandWantempSystemSdwanDuplicationSlaMatchService(d, i["sla_match_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandWantempSystemSdwanDuplicationSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr6"], _ = expandWantempSystemSdwanDuplicationSrcaddr6(d, i["srcaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcintf"], _ = expandWantempSystemSdwanDuplicationSrcintf(d, i["srcintf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanDuplicationDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationPacketDeDuplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationPacketDuplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationSlaMatchService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanDuplicationSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanDuplicationMaxNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanFailDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_dynamic-server"], _ = expandWantempSystemSdwanHealthCheckDynamicServer(d, i["_dynamic_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandWantempSystemSdwanHealthCheckAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandWantempSystemSdwanHealthCheckClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detect-mode"], _ = expandWantempSystemSdwanHealthCheckDetectMode(d, i["detect_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandWantempSystemSdwanHealthCheckDiffservcode(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-match-ip"], _ = expandWantempSystemSdwanHealthCheckDnsMatchIp(d, i["dns_match_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-request-domain"], _ = expandWantempSystemSdwanHealthCheckDnsRequestDomain(d, i["dns_request_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embed-measured-health"], _ = expandWantempSystemSdwanHealthCheckEmbedMeasuredHealth(d, i["embed_measured_health"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandWantempSystemSdwanHealthCheckFailtime(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-file"], _ = expandWantempSystemSdwanHealthCheckFtpFile(d, i["ftp_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-mode"], _ = expandWantempSystemSdwanHealthCheckFtpMode(d, i["ftp_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandWantempSystemSdwanHealthCheckHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandWantempSystemSdwanHealthCheckHttpAgent(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandWantempSystemSdwanHealthCheckHttpGet(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandWantempSystemSdwanHealthCheckHttpMatch(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandWantempSystemSdwanHealthCheckInterval(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandWantempSystemSdwanHealthCheckMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-codec"], _ = expandWantempSystemSdwanHealthCheckMosCodec(d, i["mos_codec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemSdwanHealthCheckName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandWantempSystemSdwanHealthCheckPacketSize(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandWantempSystemSdwanHealthCheckPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandWantempSystemSdwanHealthCheckPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-count"], _ = expandWantempSystemSdwanHealthCheckProbeCount(d, i["probe_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandWantempSystemSdwanHealthCheckProbePackets(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandWantempSystemSdwanHealthCheckProbeTimeout(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemSdwanHealthCheckProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-measured-method"], _ = expandWantempSystemSdwanHealthCheckQualityMeasuredMethod(d, i["quality_measured_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandWantempSystemSdwanHealthCheckRecoverytime(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandWantempSystemSdwanHealthCheckSecurityMode(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandWantempSystemSdwanHealthCheckServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemSdwanHealthCheckSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandWantempSystemSdwanHealthCheckSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id-redistribute"], _ = expandWantempSystemSdwanHealthCheckSlaIdRedistribute(d, i["sla_id_redistribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandWantempSystemSdwanHealthCheckSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemSdwanHealthCheckSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemSdwanHealthCheckSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["system-dns"], _ = expandWantempSystemSdwanHealthCheckSystemDns(d, i["system_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandWantempSystemSdwanHealthCheckThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandWantempSystemSdwanHealthCheckThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandWantempSystemSdwanHealthCheckThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandWantempSystemSdwanHealthCheckThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandWantempSystemSdwanHealthCheckThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandWantempSystemSdwanHealthCheckThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandWantempSystemSdwanHealthCheckUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandWantempSystemSdwanHealthCheckUpdateStaticRoute(d, i["update_static_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandWantempSystemSdwanHealthCheckUser(d, i["user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandWantempSystemSdwanHealthCheckVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckDynamicServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanHealthCheckDetectMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDnsMatchIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckEmbedMeasuredHealth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFtpFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckMosCodec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbeCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbePackets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProbeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckQualityMeasuredMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemSdwanHealthCheckSlaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaJitterThreshold(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemSdwanHealthCheckSlaLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaMosThreshold(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemSdwanHealthCheckSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandWantempSystemSdwanHealthCheckSlaPriorityInSla(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandWantempSystemSdwanHealthCheckSlaPriorityOutSla(d, i["priority_out_sla"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSlaMosThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityInSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityOutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaIdRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSystemDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-mode"], _ = expandWantempSystemSdwanHealthCheckFortiguardAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandWantempSystemSdwanHealthCheckFortiguardClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detect-mode"], _ = expandWantempSystemSdwanHealthCheckFortiguardDetectMode(d, i["detect_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandWantempSystemSdwanHealthCheckFortiguardDiffservcode(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-match-ip"], _ = expandWantempSystemSdwanHealthCheckFortiguardDnsMatchIp(d, i["dns_match_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-request-domain"], _ = expandWantempSystemSdwanHealthCheckFortiguardDnsRequestDomain(d, i["dns_request_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embed-measured-health"], _ = expandWantempSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(d, i["embed_measured_health"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandWantempSystemSdwanHealthCheckFortiguardFailtime(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-file"], _ = expandWantempSystemSdwanHealthCheckFortiguardFtpFile(d, i["ftp_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-mode"], _ = expandWantempSystemSdwanHealthCheckFortiguardFtpMode(d, i["ftp_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandWantempSystemSdwanHealthCheckFortiguardHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandWantempSystemSdwanHealthCheckFortiguardHttpAgent(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandWantempSystemSdwanHealthCheckFortiguardHttpGet(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandWantempSystemSdwanHealthCheckFortiguardHttpMatch(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandWantempSystemSdwanHealthCheckFortiguardInterval(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandWantempSystemSdwanHealthCheckFortiguardMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-codec"], _ = expandWantempSystemSdwanHealthCheckFortiguardMosCodec(d, i["mos_codec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandWantempSystemSdwanHealthCheckFortiguardPacketSize(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandWantempSystemSdwanHealthCheckFortiguardPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandWantempSystemSdwanHealthCheckFortiguardPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-count"], _ = expandWantempSystemSdwanHealthCheckFortiguardProbeCount(d, i["probe_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandWantempSystemSdwanHealthCheckFortiguardProbePackets(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandWantempSystemSdwanHealthCheckFortiguardProbeTimeout(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemSdwanHealthCheckFortiguardProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-measured-method"], _ = expandWantempSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(d, i["quality_measured_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandWantempSystemSdwanHealthCheckFortiguardRecoverytime(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandWantempSystemSdwanHealthCheckFortiguardSecurityMode(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandWantempSystemSdwanHealthCheckFortiguardServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemSdwanHealthCheckFortiguardSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id-redistribute"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaIdRedistribute(d, i["sla_id_redistribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemSdwanHealthCheckFortiguardSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemSdwanHealthCheckFortiguardSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["system-dns"], _ = expandWantempSystemSdwanHealthCheckFortiguardSystemDns(d, i["system_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target-name"], _ = expandWantempSystemSdwanHealthCheckFortiguardTargetName(d, i["target_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandWantempSystemSdwanHealthCheckFortiguardThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandWantempSystemSdwanHealthCheckFortiguardThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandWantempSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandWantempSystemSdwanHealthCheckFortiguardThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandWantempSystemSdwanHealthCheckFortiguardThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandWantempSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandWantempSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandWantempSystemSdwanHealthCheckFortiguardUpdateStaticRoute(d, i["update_static_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandWantempSystemSdwanHealthCheckFortiguardUser(d, i["user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandWantempSystemSdwanHealthCheckFortiguardVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckFortiguardDetectMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardDnsMatchIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardFtpFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardFtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckFortiguardMosCodec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckFortiguardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardProbeCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardProbePackets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardProbeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaJitterThreshold(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaMosThreshold(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaPriorityInSla(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandWantempSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(d, i["priority_out_sla"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaMosThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaPriorityInSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaIdRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardSystemDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardTargetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckFortiguardVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanLoadBalanceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_dynamic-member"], _ = expandWantempSystemSdwanMembersDynamicMember(d, i["_dynamic_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWantempSystemSdwanMembersComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandWantempSystemSdwanMembersCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemSdwanMembersGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway6"], _ = expandWantempSystemSdwanMembersGateway6(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-spillover-threshold"], _ = expandWantempSystemSdwanMembersIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandWantempSystemSdwanMembersInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-source"], _ = expandWantempSystemSdwanMembersPreferredSource(d, i["preferred_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandWantempSystemSdwanMembersPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandWantempSystemSdwanMembersPriorityInSla(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandWantempSystemSdwanMembersPriorityOutSla(d, i["priority_out_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority6"], _ = expandWantempSystemSdwanMembersPriority6(d, i["priority6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq-num"], _ = expandWantempSystemSdwanMembersSeqNum(d, i["seq_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandWantempSystemSdwanMembersSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandWantempSystemSdwanMembersSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spillover-threshold"], _ = expandWantempSystemSdwanMembersSpilloverThreshold(d, i["spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemSdwanMembersStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transport-group"], _ = expandWantempSystemSdwanMembersTransportGroup(d, i["transport_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["volume-ratio"], _ = expandWantempSystemSdwanMembersVolumeRatio(d, i["volume_ratio"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandWantempSystemSdwanMembersWeight(d, i["weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["zone"], _ = expandWantempSystemSdwanMembersZone(d, i["zone"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanMembersDynamicMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanMembersPreferredSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriorityInSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriorityOutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersPriority6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersTransportGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanMembersZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemSdwanNeighborHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWantempSystemSdwanNeighborIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandWantempSystemSdwanNeighborMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandWantempSystemSdwanNeighborMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandWantempSystemSdwanNeighborMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemSdwanNeighborRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandWantempSystemSdwanNeighborServiceId(d, i["service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id"], _ = expandWantempSystemSdwanNeighborSlaId(d, i["sla_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanNeighborMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanNeighborMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanNeighborSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborHoldBootTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborHoldDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["addr-mode"], _ = expandWantempSystemSdwanServiceAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_exclusive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["agent-exclusive"], _ = expandWantempSystemSdwanServiceAgentExclusive(d, i["agent_exclusive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-weight"], _ = expandWantempSystemSdwanServiceBandwidthWeight(d, i["bandwidth_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandWantempSystemSdwanServiceComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandWantempSystemSdwanServiceDefault(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward"], _ = expandWantempSystemSdwanServiceDscpForward(d, i["dscp_forward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward-tag"], _ = expandWantempSystemSdwanServiceDscpForwardTag(d, i["dscp_forward_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse"], _ = expandWantempSystemSdwanServiceDscpReverse(d, i["dscp_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse-tag"], _ = expandWantempSystemSdwanServiceDscpReverseTag(d, i["dscp_reverse_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandWantempSystemSdwanServiceDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-negate"], _ = expandWantempSystemSdwanServiceDstNegate(d, i["dst_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst6"], _ = expandWantempSystemSdwanServiceDst6(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-port"], _ = expandWantempSystemSdwanServiceEndPort(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_src_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-src-port"], _ = expandWantempSystemSdwanServiceEndSrcPort(d, i["end_src_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandWantempSystemSdwanServiceGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandWantempSystemSdwanServiceGroups(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hash-mode"], _ = expandWantempSystemSdwanServiceHashMode(d, i["hash_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandWantempSystemSdwanServiceHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hold-down-time"], _ = expandWantempSystemSdwanServiceHoldDownTime(d, i["hold_down_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device"], _ = expandWantempSystemSdwanServiceInputDevice(d, i["input_device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device-negate"], _ = expandWantempSystemSdwanServiceInputDeviceNegate(d, i["input_device_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-zone"], _ = expandWantempSystemSdwanServiceInputZone(d, i["input_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service"], _ = expandWantempSystemSdwanServiceInternetService(d, i["internet_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl"], _ = expandWantempSystemSdwanServiceInternetServiceAppCtrl(d, i["internet_service_app_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-category"], _ = expandWantempSystemSdwanServiceInternetServiceAppCtrlCategory(d, i["internet_service_app_ctrl_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-group"], _ = expandWantempSystemSdwanServiceInternetServiceAppCtrlGroup(d, i["internet_service_app_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom"], _ = expandWantempSystemSdwanServiceInternetServiceCustom(d, i["internet_service_custom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom-group"], _ = expandWantempSystemSdwanServiceInternetServiceCustomGroup(d, i["internet_service_custom_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-group"], _ = expandWantempSystemSdwanServiceInternetServiceGroup(d, i["internet_service_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-name"], _ = expandWantempSystemSdwanServiceInternetServiceName(d, i["internet_service_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-weight"], _ = expandWantempSystemSdwanServiceJitterWeight(d, i["jitter_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-weight"], _ = expandWantempSystemSdwanServiceLatencyWeight(d, i["latency_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandWantempSystemSdwanServiceLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-threshold"], _ = expandWantempSystemSdwanServiceLinkCostThreshold(d, i["link_cost_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balance"], _ = expandWantempSystemSdwanServiceLoadBalance(d, i["load_balance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandWantempSystemSdwanServiceMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandWantempSystemSdwanServiceMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemSdwanServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-weight"], _ = expandWantempSystemSdwanServicePacketLossWeight(d, i["packet_loss_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passive-measurement"], _ = expandWantempSystemSdwanServicePassiveMeasurement(d, i["passive_measurement"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-members"], _ = expandWantempSystemSdwanServicePriorityMembers(d, i["priority_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-zone"], _ = expandWantempSystemSdwanServicePriorityZone(d, i["priority_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandWantempSystemSdwanServiceProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-link"], _ = expandWantempSystemSdwanServiceQualityLink(d, i["quality_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandWantempSystemSdwanServiceRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut"], _ = expandWantempSystemSdwanServiceShortcut(d, i["shortcut"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut-priority"], _ = expandWantempSystemSdwanServiceShortcutPriority(d, i["shortcut_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_stickiness"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut-stickiness"], _ = expandWantempSystemSdwanServiceShortcutStickiness(d, i["shortcut_stickiness"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandWantempSystemSdwanServiceRouteTag(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandWantempSystemSdwanServiceSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-compare-method"], _ = expandWantempSystemSdwanServiceSlaCompareMethod(d, i["sla_compare_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_stickiness"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-stickiness"], _ = expandWantempSystemSdwanServiceSlaStickiness(d, i["sla_stickiness"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandWantempSystemSdwanServiceSrc(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-negate"], _ = expandWantempSystemSdwanServiceSrcNegate(d, i["src_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src6"], _ = expandWantempSystemSdwanServiceSrc6(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["standalone-action"], _ = expandWantempSystemSdwanServiceStandaloneAction(d, i["standalone_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandWantempSystemSdwanServiceStartPort(d, i["start_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_src_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-src-port"], _ = expandWantempSystemSdwanServiceStartSrcPort(d, i["start_src_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandWantempSystemSdwanServiceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tie-break"], _ = expandWantempSystemSdwanServiceTieBreak(d, i["tie_break"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandWantempSystemSdwanServiceTos(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandWantempSystemSdwanServiceTosMask(d, i["tos_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["use-shortcut-sla"], _ = expandWantempSystemSdwanServiceUseShortcutSla(d, i["use_shortcut_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["users"], _ = expandWantempSystemSdwanServiceUsers(d, i["users"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["zone-mode"], _ = expandWantempSystemSdwanServiceZoneMode(d, i["zone_mode"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanServiceAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceAgentExclusive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceBandwidthWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpForwardTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDscpReverseTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceDstNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceDst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceEndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceEndSrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceHashMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInputZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceJitterWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLatencyWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLinkCostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceLoadBalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePacketLossWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePassiveMeasurement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServicePriorityZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceQualityLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcutPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceShortcutStickiness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["health-check"], _ = expandWantempSystemSdwanServiceSlaHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandWantempSystemSdwanServiceSlaId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanServiceSlaHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaCompareMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSlaStickiness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceSrc6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceStandaloneAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStartSrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTieBreak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceUseShortcutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanServiceUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanServiceZoneMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanSpeedtestBypassRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advpn-health-check"], _ = expandWantempSystemSdwanZoneAdvpnHealthCheck(d, i["advpn_health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advpn-select"], _ = expandWantempSystemSdwanZoneAdvpnSelect(d, i["advpn_select"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandWantempSystemSdwanZoneMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandWantempSystemSdwanZoneName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-sla-tie-break"], _ = expandWantempSystemSdwanZoneServiceSlaTieBreak(d, i["service_sla_tie_break"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWantempSystemSdwanZoneAdvpnHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandWantempSystemSdwanZoneAdvpnSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZoneMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZoneName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanZoneServiceSlaTieBreak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_perf_log_period"); ok || d.HasChange("app_perf_log_period") {
		t, err := expandWantempSystemSdwanAppPerfLogPeriod(d, v, "app_perf_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-perf-log-period"] = t
		}
	}

	if v, ok := d.GetOk("duplication"); ok || d.HasChange("duplication") {
		t, err := expandWantempSystemSdwanDuplication(d, v, "duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication"] = t
		}
	}

	if v, ok := d.GetOk("duplication_max_num"); ok || d.HasChange("duplication_max_num") {
		t, err := expandWantempSystemSdwanDuplicationMaxNum(d, v, "duplication_max_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-max-num"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok || d.HasChange("fail_alert_interfaces") {
		t, err := expandWantempSystemSdwanFailAlertInterfaces(d, v, "fail_alert_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok || d.HasChange("fail_detect") {
		t, err := expandWantempSystemSdwanFailDetect(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemSdwanHealthCheck(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_fortiguard"); ok || d.HasChange("health_check_fortiguard") {
		t, err := expandWantempSystemSdwanHealthCheckFortiguard(d, v, "health_check_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_mode"); ok || d.HasChange("load_balance_mode") {
		t, err := expandWantempSystemSdwanLoadBalanceMode(d, v, "load_balance_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-mode"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandWantempSystemSdwanMembers(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandWantempSystemSdwanNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_boot_time"); ok || d.HasChange("neighbor_hold_boot_time") {
		t, err := expandWantempSystemSdwanNeighborHoldBootTime(d, v, "neighbor_hold_boot_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-boot-time"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok || d.HasChange("neighbor_hold_down") {
		t, err := expandWantempSystemSdwanNeighborHoldDown(d, v, "neighbor_hold_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down_time"); ok || d.HasChange("neighbor_hold_down_time") {
		t, err := expandWantempSystemSdwanNeighborHoldDownTime(d, v, "neighbor_hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandWantempSystemSdwanOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandWantempSystemSdwanService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("speedtest_bypass_routing"); ok || d.HasChange("speedtest_bypass_routing") {
		t, err := expandWantempSystemSdwanSpeedtestBypassRouting(d, v, "speedtest_bypass_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtest-bypass-routing"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWantempSystemSdwanStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		t, err := expandWantempSystemSdwanZone(d, v, "zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
