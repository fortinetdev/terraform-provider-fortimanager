// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU attributes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpu() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuUpdate,
		Read:   resourceObjectSystemNpuRead,
		Update: resourceObjectSystemNpuUpdate,
		Delete: resourceObjectSystemNpuDelete,

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
			"background_sse_scan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stats_update_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"udp_keepalive_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"capwap_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_management_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_management_cpu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_qos_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dos_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"npu_dos_meter_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"npu_dos_synproxy_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"npu_dos_tpe_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"double_level_mcast_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dse_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dsw_dts_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"step": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dsw_queue_dts_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"queue_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fastpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fp_anomaly": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capwap_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"esp_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gre_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gtpu_plen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_frag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_ihl_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_len_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_opt_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optlsrr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optrr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optsecurity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optssrr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optstream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_opttimestamp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_proto_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_ttlzero_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_unknopt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_ver_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_daddr_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_exthdr_len_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_exthdr_order_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_ihl_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optendpid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_opthomeaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optinvld": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optjumbo": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optnsap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optralert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_opttunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_plen_zero": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_proto_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_saddr_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_unknopt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_ver_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nvgre_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sctp_clen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sctp_crc_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sctp_l4len_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_fin_noack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_fin_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_hlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_hlenvsl4len_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_no_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_plen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_syn_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_syn_fin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_winnuke": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_hlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_len_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_plen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udplite_cover_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udplite_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uesp_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknproto_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vxlan_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"gtp_enhanced_cpu_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp_enhanced_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hash_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hash_tbl_spread": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_shortcut_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hpe": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"arp_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"enable_shaper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"esp_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"high_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmp_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip_frag_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip_others_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"l2_others_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pri_type_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sctp_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcpfin_rst_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcpsyn_ack_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcpsyn_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"udp_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"htab_dedi_queue_nr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"htab_msg_queue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htx_gtse_quota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htx_icmp_csum_chk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hw_ha_scan_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"inbound_dscp_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbound_dscp_copy_port": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_reassembly": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"intf_shaping_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iph_rsvd_re_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool_overload_high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ippool_overload_low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_dec_subengine_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_enc_subengine_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_host_dfclr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_inbound_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_local_uesp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_ob_np_sel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_over_vlink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isf_np_queues": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"lag_out_port_select": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_session_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mcast_session_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcast_session_counting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcast_session_counting6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"napi_break_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nat46_force_ipv4_packet_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"np_queues": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_type": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"queue": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ip_protocol": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"queue": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ip_service": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"queue": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"sport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cos0": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos5": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos7": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp0": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp10": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp11": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp12": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp13": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp14": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp15": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp16": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp17": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp18": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp19": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp20": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp21": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp22": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp23": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp24": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp25": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp26": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp27": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp28": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp29": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp30": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp31": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp32": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp33": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp34": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp35": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp36": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp37": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp38": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp39": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp40": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp41": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp42": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp43": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp44": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp45": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp46": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp47": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp48": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp49": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp5": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp50": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp51": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp52": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp53": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp54": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp55": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp56": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp57": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp58": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp59": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp60": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp61": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp62": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp63": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp7": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp8": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp9": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"scheduler": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"np6_cps_optimization_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pba_eim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_policy_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_session_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_offload_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_cpu_map": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_core": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"port_npu_map": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"npu_group_index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"port_path_option": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports_using_npu": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"priority_protocol": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bgp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"slbc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"process_icmp_by_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prp_port_in": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prp_port_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qtm_buf_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recover_np6_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_acct_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"session_denied_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sse_backpressure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_clear_text_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_esp_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_eh_hash": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"computation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_ip_lower_16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_ip_upper_16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"netmask_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_ip_lower_16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_ip_upper_16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sw_np_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_np_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_rst_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_timeout_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"close_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fin_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"syn_sent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"syn_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_idle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"time_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"udp_timeout_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"udp_idle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"uesp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_lookup_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectSystemNpuUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpu(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpu resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpu(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpu resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpu")

	return resourceObjectSystemNpuRead(d, m)
}

func resourceObjectSystemNpuDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpu(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpu resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpu(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpu resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpu(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpu resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuBackgroundSseScanOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "scan"
	if _, ok := i["scan"]; ok {
		result["scan"] = flattenObjectSystemNpuBackgroundSseScanScanOsna(i["scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "stats_update_interval"
	if _, ok := i["stats-update-interval"]; ok {
		result["stats_update_interval"] = flattenObjectSystemNpuBackgroundSseScanStatsUpdateIntervalOsna(i["stats-update-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_keepalive_interval"
	if _, ok := i["udp-keepalive-interval"]; ok {
		result["udp_keepalive_interval"] = flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveIntervalOsna(i["udp-keepalive-interval"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuBackgroundSseScanScanOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsUpdateIntervalOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveIntervalOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuCapwapOffloadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDedicatedManagementAffinityOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDedicatedManagementCpuOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDefaultQosTypeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "npu_dos_meter_mode"
	if _, ok := i["npu-dos-meter-mode"]; ok {
		result["npu_dos_meter_mode"] = flattenObjectSystemNpuDosOptionsNpuDosMeterModeOsna(i["npu-dos-meter-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "npu_dos_synproxy_mode"
	if _, ok := i["npu-dos-synproxy-mode"]; ok {
		result["npu_dos_synproxy_mode"] = flattenObjectSystemNpuDosOptionsNpuDosSynproxyModeOsna(i["npu-dos-synproxy-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "npu_dos_tpe_mode"
	if _, ok := i["npu-dos-tpe-mode"]; ok {
		result["npu_dos_tpe_mode"] = flattenObjectSystemNpuDosOptionsNpuDosTpeModeOsna(i["npu-dos-tpe-mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuDosOptionsNpuDosMeterModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsNpuDosSynproxyModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsNpuDosTpeModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDoubleLevelMcastOffloadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDseTimeoutOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileActionOsna(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_limit"
		if _, ok := i["min-limit"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileMinLimitOsna(i["min-limit"], d, pre_append)
			tmp["min_limit"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-MinLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := i["profile-id"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileProfileIdOsna(i["profile-id"], d, pre_append)
			tmp["profile_id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-ProfileId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "step"
		if _, ok := i["step"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileStepOsna(i["step"], d, pre_append)
			tmp["step"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-Step")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuDswDtsProfileActionOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileMinLimitOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileProfileIdOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileStepOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iport"
		if _, ok := i["iport"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileIportOsna(i["iport"], d, pre_append)
			tmp["iport"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-Iport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileNameOsna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oport"
		if _, ok := i["oport"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileOportOsna(i["oport"], d, pre_append)
			tmp["oport"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-Oport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := i["profile-id"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileProfileIdOsna(i["profile-id"], d, pre_append)
			tmp["profile_id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-ProfileId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue_select"
		if _, ok := i["queue-select"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileQueueSelectOsna(i["queue-select"], d, pre_append)
			tmp["queue_select"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-QueueSelect")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuDswQueueDtsProfileIportOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileNameOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileOportOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileProfileIdOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileQueueSelectOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFastpathOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "capwap_minlen_err"
	if _, ok := i["capwap-minlen-err"]; ok {
		result["capwap_minlen_err"] = flattenObjectSystemNpuFpAnomalyCapwapMinlenErrOsna(i["capwap-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "esp_minlen_err"
	if _, ok := i["esp-minlen-err"]; ok {
		result["esp_minlen_err"] = flattenObjectSystemNpuFpAnomalyEspMinlenErrOsna(i["esp-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "gre_csum_err"
	if _, ok := i["gre-csum-err"]; ok {
		result["gre_csum_err"] = flattenObjectSystemNpuFpAnomalyGreCsumErrOsna(i["gre-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "gtpu_plen_err"
	if _, ok := i["gtpu-plen-err"]; ok {
		result["gtpu_plen_err"] = flattenObjectSystemNpuFpAnomalyGtpuPlenErrOsna(i["gtpu-plen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_csum_err"
	if _, ok := i["icmp-csum-err"]; ok {
		result["icmp_csum_err"] = flattenObjectSystemNpuFpAnomalyIcmpCsumErrOsna(i["icmp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_frag"
	if _, ok := i["icmp-frag"]; ok {
		result["icmp_frag"] = flattenObjectSystemNpuFpAnomalyIcmpFragOsna(i["icmp-frag"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_land"
	if _, ok := i["icmp-land"]; ok {
		result["icmp_land"] = flattenObjectSystemNpuFpAnomalyIcmpLandOsna(i["icmp-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_minlen_err"
	if _, ok := i["icmp-minlen-err"]; ok {
		result["icmp_minlen_err"] = flattenObjectSystemNpuFpAnomalyIcmpMinlenErrOsna(i["icmp-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_csum_err"
	if _, ok := i["ipv4-csum-err"]; ok {
		result["ipv4_csum_err"] = flattenObjectSystemNpuFpAnomalyIpv4CsumErrOsna(i["ipv4-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_ihl_err"
	if _, ok := i["ipv4-ihl-err"]; ok {
		result["ipv4_ihl_err"] = flattenObjectSystemNpuFpAnomalyIpv4IhlErrOsna(i["ipv4-ihl-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_land"
	if _, ok := i["ipv4-land"]; ok {
		result["ipv4_land"] = flattenObjectSystemNpuFpAnomalyIpv4LandOsna(i["ipv4-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_len_err"
	if _, ok := i["ipv4-len-err"]; ok {
		result["ipv4_len_err"] = flattenObjectSystemNpuFpAnomalyIpv4LenErrOsna(i["ipv4-len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_opt_err"
	if _, ok := i["ipv4-opt-err"]; ok {
		result["ipv4_opt_err"] = flattenObjectSystemNpuFpAnomalyIpv4OptErrOsna(i["ipv4-opt-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optlsrr"
	if _, ok := i["ipv4-optlsrr"]; ok {
		result["ipv4_optlsrr"] = flattenObjectSystemNpuFpAnomalyIpv4OptlsrrOsna(i["ipv4-optlsrr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optrr"
	if _, ok := i["ipv4-optrr"]; ok {
		result["ipv4_optrr"] = flattenObjectSystemNpuFpAnomalyIpv4OptrrOsna(i["ipv4-optrr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optsecurity"
	if _, ok := i["ipv4-optsecurity"]; ok {
		result["ipv4_optsecurity"] = flattenObjectSystemNpuFpAnomalyIpv4OptsecurityOsna(i["ipv4-optsecurity"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optssrr"
	if _, ok := i["ipv4-optssrr"]; ok {
		result["ipv4_optssrr"] = flattenObjectSystemNpuFpAnomalyIpv4OptssrrOsna(i["ipv4-optssrr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optstream"
	if _, ok := i["ipv4-optstream"]; ok {
		result["ipv4_optstream"] = flattenObjectSystemNpuFpAnomalyIpv4OptstreamOsna(i["ipv4-optstream"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_opttimestamp"
	if _, ok := i["ipv4-opttimestamp"]; ok {
		result["ipv4_opttimestamp"] = flattenObjectSystemNpuFpAnomalyIpv4OpttimestampOsna(i["ipv4-opttimestamp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_proto_err"
	if _, ok := i["ipv4-proto-err"]; ok {
		result["ipv4_proto_err"] = flattenObjectSystemNpuFpAnomalyIpv4ProtoErrOsna(i["ipv4-proto-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_ttlzero_err"
	if _, ok := i["ipv4-ttlzero-err"]; ok {
		result["ipv4_ttlzero_err"] = flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErrOsna(i["ipv4-ttlzero-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_unknopt"
	if _, ok := i["ipv4-unknopt"]; ok {
		result["ipv4_unknopt"] = flattenObjectSystemNpuFpAnomalyIpv4UnknoptOsna(i["ipv4-unknopt"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_ver_err"
	if _, ok := i["ipv4-ver-err"]; ok {
		result["ipv4_ver_err"] = flattenObjectSystemNpuFpAnomalyIpv4VerErrOsna(i["ipv4-ver-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_daddr_err"
	if _, ok := i["ipv6-daddr-err"]; ok {
		result["ipv6_daddr_err"] = flattenObjectSystemNpuFpAnomalyIpv6DaddrErrOsna(i["ipv6-daddr-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_exthdr_len_err"
	if _, ok := i["ipv6-exthdr-len-err"]; ok {
		result["ipv6_exthdr_len_err"] = flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErrOsna(i["ipv6-exthdr-len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_exthdr_order_err"
	if _, ok := i["ipv6-exthdr-order-err"]; ok {
		result["ipv6_exthdr_order_err"] = flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErrOsna(i["ipv6-exthdr-order-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_ihl_err"
	if _, ok := i["ipv6-ihl-err"]; ok {
		result["ipv6_ihl_err"] = flattenObjectSystemNpuFpAnomalyIpv6IhlErrOsna(i["ipv6-ihl-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_land"
	if _, ok := i["ipv6-land"]; ok {
		result["ipv6_land"] = flattenObjectSystemNpuFpAnomalyIpv6LandOsna(i["ipv6-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optendpid"
	if _, ok := i["ipv6-optendpid"]; ok {
		result["ipv6_optendpid"] = flattenObjectSystemNpuFpAnomalyIpv6OptendpidOsna(i["ipv6-optendpid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_opthomeaddr"
	if _, ok := i["ipv6-opthomeaddr"]; ok {
		result["ipv6_opthomeaddr"] = flattenObjectSystemNpuFpAnomalyIpv6OpthomeaddrOsna(i["ipv6-opthomeaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optinvld"
	if _, ok := i["ipv6-optinvld"]; ok {
		result["ipv6_optinvld"] = flattenObjectSystemNpuFpAnomalyIpv6OptinvldOsna(i["ipv6-optinvld"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optjumbo"
	if _, ok := i["ipv6-optjumbo"]; ok {
		result["ipv6_optjumbo"] = flattenObjectSystemNpuFpAnomalyIpv6OptjumboOsna(i["ipv6-optjumbo"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optnsap"
	if _, ok := i["ipv6-optnsap"]; ok {
		result["ipv6_optnsap"] = flattenObjectSystemNpuFpAnomalyIpv6OptnsapOsna(i["ipv6-optnsap"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optralert"
	if _, ok := i["ipv6-optralert"]; ok {
		result["ipv6_optralert"] = flattenObjectSystemNpuFpAnomalyIpv6OptralertOsna(i["ipv6-optralert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_opttunnel"
	if _, ok := i["ipv6-opttunnel"]; ok {
		result["ipv6_opttunnel"] = flattenObjectSystemNpuFpAnomalyIpv6OpttunnelOsna(i["ipv6-opttunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_plen_zero"
	if _, ok := i["ipv6-plen-zero"]; ok {
		result["ipv6_plen_zero"] = flattenObjectSystemNpuFpAnomalyIpv6PlenZeroOsna(i["ipv6-plen-zero"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_proto_err"
	if _, ok := i["ipv6-proto-err"]; ok {
		result["ipv6_proto_err"] = flattenObjectSystemNpuFpAnomalyIpv6ProtoErrOsna(i["ipv6-proto-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_saddr_err"
	if _, ok := i["ipv6-saddr-err"]; ok {
		result["ipv6_saddr_err"] = flattenObjectSystemNpuFpAnomalyIpv6SaddrErrOsna(i["ipv6-saddr-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_unknopt"
	if _, ok := i["ipv6-unknopt"]; ok {
		result["ipv6_unknopt"] = flattenObjectSystemNpuFpAnomalyIpv6UnknoptOsna(i["ipv6-unknopt"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_ver_err"
	if _, ok := i["ipv6-ver-err"]; ok {
		result["ipv6_ver_err"] = flattenObjectSystemNpuFpAnomalyIpv6VerErrOsna(i["ipv6-ver-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "nvgre_minlen_err"
	if _, ok := i["nvgre-minlen-err"]; ok {
		result["nvgre_minlen_err"] = flattenObjectSystemNpuFpAnomalyNvgreMinlenErrOsna(i["nvgre-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_clen_err"
	if _, ok := i["sctp-clen-err"]; ok {
		result["sctp_clen_err"] = flattenObjectSystemNpuFpAnomalySctpClenErrOsna(i["sctp-clen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_crc_err"
	if _, ok := i["sctp-crc-err"]; ok {
		result["sctp_crc_err"] = flattenObjectSystemNpuFpAnomalySctpCrcErrOsna(i["sctp-crc-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_l4len_err"
	if _, ok := i["sctp-l4len-err"]; ok {
		result["sctp_l4len_err"] = flattenObjectSystemNpuFpAnomalySctpL4LenErrOsna(i["sctp-l4len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_csum_err"
	if _, ok := i["tcp-csum-err"]; ok {
		result["tcp_csum_err"] = flattenObjectSystemNpuFpAnomalyTcpCsumErrOsna(i["tcp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin_noack"
	if _, ok := i["tcp-fin-noack"]; ok {
		result["tcp_fin_noack"] = flattenObjectSystemNpuFpAnomalyTcpFinNoackOsna(i["tcp-fin-noack"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin_only"
	if _, ok := i["tcp-fin-only"]; ok {
		result["tcp_fin_only"] = flattenObjectSystemNpuFpAnomalyTcpFinOnlyOsna(i["tcp-fin-only"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_hlen_err"
	if _, ok := i["tcp-hlen-err"]; ok {
		result["tcp_hlen_err"] = flattenObjectSystemNpuFpAnomalyTcpHlenErrOsna(i["tcp-hlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_hlenvsl4len_err"
	if _, ok := i["tcp-hlenvsl4len-err"]; ok {
		result["tcp_hlenvsl4len_err"] = flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErrOsna(i["tcp-hlenvsl4len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_land"
	if _, ok := i["tcp-land"]; ok {
		result["tcp_land"] = flattenObjectSystemNpuFpAnomalyTcpLandOsna(i["tcp-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_no_flag"
	if _, ok := i["tcp-no-flag"]; ok {
		result["tcp_no_flag"] = flattenObjectSystemNpuFpAnomalyTcpNoFlagOsna(i["tcp-no-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_plen_err"
	if _, ok := i["tcp-plen-err"]; ok {
		result["tcp_plen_err"] = flattenObjectSystemNpuFpAnomalyTcpPlenErrOsna(i["tcp-plen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn_data"
	if _, ok := i["tcp-syn-data"]; ok {
		result["tcp_syn_data"] = flattenObjectSystemNpuFpAnomalyTcpSynDataOsna(i["tcp-syn-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn_fin"
	if _, ok := i["tcp-syn-fin"]; ok {
		result["tcp_syn_fin"] = flattenObjectSystemNpuFpAnomalyTcpSynFinOsna(i["tcp-syn-fin"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_winnuke"
	if _, ok := i["tcp-winnuke"]; ok {
		result["tcp_winnuke"] = flattenObjectSystemNpuFpAnomalyTcpWinnukeOsna(i["tcp-winnuke"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_csum_err"
	if _, ok := i["udp-csum-err"]; ok {
		result["udp_csum_err"] = flattenObjectSystemNpuFpAnomalyUdpCsumErrOsna(i["udp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_hlen_err"
	if _, ok := i["udp-hlen-err"]; ok {
		result["udp_hlen_err"] = flattenObjectSystemNpuFpAnomalyUdpHlenErrOsna(i["udp-hlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_land"
	if _, ok := i["udp-land"]; ok {
		result["udp_land"] = flattenObjectSystemNpuFpAnomalyUdpLandOsna(i["udp-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_len_err"
	if _, ok := i["udp-len-err"]; ok {
		result["udp_len_err"] = flattenObjectSystemNpuFpAnomalyUdpLenErrOsna(i["udp-len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_plen_err"
	if _, ok := i["udp-plen-err"]; ok {
		result["udp_plen_err"] = flattenObjectSystemNpuFpAnomalyUdpPlenErrOsna(i["udp-plen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udplite_cover_err"
	if _, ok := i["udplite-cover-err"]; ok {
		result["udplite_cover_err"] = flattenObjectSystemNpuFpAnomalyUdpliteCoverErrOsna(i["udplite-cover-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udplite_csum_err"
	if _, ok := i["udplite-csum-err"]; ok {
		result["udplite_csum_err"] = flattenObjectSystemNpuFpAnomalyUdpliteCsumErrOsna(i["udplite-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "uesp_minlen_err"
	if _, ok := i["uesp-minlen-err"]; ok {
		result["uesp_minlen_err"] = flattenObjectSystemNpuFpAnomalyUespMinlenErrOsna(i["uesp-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknproto_minlen_err"
	if _, ok := i["unknproto-minlen-err"]; ok {
		result["unknproto_minlen_err"] = flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErrOsna(i["unknproto-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "vxlan_minlen_err"
	if _, ok := i["vxlan-minlen-err"]; ok {
		result["vxlan_minlen_err"] = flattenObjectSystemNpuFpAnomalyVxlanMinlenErrOsna(i["vxlan-minlen-err"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuFpAnomalyCapwapMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyEspMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGreCsumErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGtpuPlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpCsumErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpFragOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpLandOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4CsumErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4IhlErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4LandOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4LenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptlsrrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptrrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptsecurityOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptssrrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptstreamOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OpttimestampOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4ProtoErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4UnknoptOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4VerErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6DaddrErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6IhlErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6LandOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OptendpidOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OpthomeaddrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OptinvldOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OptjumboOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OptnsapOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OptralertOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6OpttunnelOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6PlenZeroOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ProtoErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6SaddrErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6UnknoptOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6VerErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyNvgreMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpClenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpCrcErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpL4LenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpCsumErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinNoackOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinOnlyOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpLandOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpNoFlagOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpPlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynDataOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynFinOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpWinnukeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpCsumErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpHlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLandOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpPlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCoverErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCsumErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUespMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyVxlanMinlenErrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuGtpEnhancedCpuRangeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuGtpEnhancedModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuGtpSupportOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHashConfigOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHashTblSpreadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHostShortcutModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all_protocol"
	if _, ok := i["all-protocol"]; ok {
		result["all_protocol"] = flattenObjectSystemNpuHpeAllProtocolOsna(i["all-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "arp_max"
	if _, ok := i["arp-max"]; ok {
		result["arp_max"] = flattenObjectSystemNpuHpeArpMaxOsna(i["arp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "enable_shaper"
	if _, ok := i["enable-shaper"]; ok {
		result["enable_shaper"] = flattenObjectSystemNpuHpeEnableShaperOsna(i["enable-shaper"], d, pre_append)
	}

	pre_append = pre + ".0." + "esp_max"
	if _, ok := i["esp-max"]; ok {
		result["esp_max"] = flattenObjectSystemNpuHpeEspMaxOsna(i["esp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "high_priority"
	if _, ok := i["high-priority"]; ok {
		result["high_priority"] = flattenObjectSystemNpuHpeHighPriorityOsna(i["high-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_max"
	if _, ok := i["icmp-max"]; ok {
		result["icmp_max"] = flattenObjectSystemNpuHpeIcmpMaxOsna(i["icmp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_frag_max"
	if _, ok := i["ip-frag-max"]; ok {
		result["ip_frag_max"] = flattenObjectSystemNpuHpeIpFragMaxOsna(i["ip-frag-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_others_max"
	if _, ok := i["ip-others-max"]; ok {
		result["ip_others_max"] = flattenObjectSystemNpuHpeIpOthersMaxOsna(i["ip-others-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "l2_others_max"
	if _, ok := i["l2-others-max"]; ok {
		result["l2_others_max"] = flattenObjectSystemNpuHpeL2OthersMaxOsna(i["l2-others-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "pri_type_max"
	if _, ok := i["pri-type-max"]; ok {
		result["pri_type_max"] = flattenObjectSystemNpuHpePriTypeMaxOsna(i["pri-type-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_max"
	if _, ok := i["sctp-max"]; ok {
		result["sctp_max"] = flattenObjectSystemNpuHpeSctpMaxOsna(i["sctp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_max"
	if _, ok := i["tcp-max"]; ok {
		result["tcp_max"] = flattenObjectSystemNpuHpeTcpMaxOsna(i["tcp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcpfin_rst_max"
	if _, ok := i["tcpfin-rst-max"]; ok {
		result["tcpfin_rst_max"] = flattenObjectSystemNpuHpeTcpfinRstMaxOsna(i["tcpfin-rst-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcpsyn_ack_max"
	if _, ok := i["tcpsyn-ack-max"]; ok {
		result["tcpsyn_ack_max"] = flattenObjectSystemNpuHpeTcpsynAckMaxOsna(i["tcpsyn-ack-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcpsyn_max"
	if _, ok := i["tcpsyn-max"]; ok {
		result["tcpsyn_max"] = flattenObjectSystemNpuHpeTcpsynMaxOsna(i["tcpsyn-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_max"
	if _, ok := i["udp-max"]; ok {
		result["udp_max"] = flattenObjectSystemNpuHpeUdpMaxOsna(i["udp-max"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuHpeAllProtocolOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeArpMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEnableShaperOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEspMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeHighPriorityOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIcmpMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIpFragMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIpOthersMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeL2OthersMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpePriTypeMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeSctpMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpfinRstMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpsynAckMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpsynMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeUdpMaxOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtabDediQueueNrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtabMsgQueueOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtxGtseQuotaOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtxIcmpCsumChkOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHwHaScanIntervalOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuInboundDscpCopyOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuInboundDscpCopyPortOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemNpuIpReassemblyOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_timeout"
	if _, ok := i["max-timeout"]; ok {
		result["max_timeout"] = flattenObjectSystemNpuIpReassemblyMaxTimeoutOsna(i["max-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_timeout"
	if _, ok := i["min-timeout"]; ok {
		result["min_timeout"] = flattenObjectSystemNpuIpReassemblyMinTimeoutOsna(i["min-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectSystemNpuIpReassemblyStatusOsna(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuIpReassemblyMaxTimeoutOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpReassemblyMinTimeoutOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpReassemblyStatusOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIntfShapingOffloadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIphRsvdReCksumOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIppoolOverloadHighOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIppoolOverloadLowOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecDecSubengineMaskOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecEncSubengineMaskOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecHostDfclrOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecInboundCacheOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecLocalUespPortOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecMtuOverrideOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecObNpSelOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecOverVlinkOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cos0"
	if _, ok := i["cos0"]; ok {
		result["cos0"] = flattenObjectSystemNpuIsfNpQueuesCos0Osna(i["cos0"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos1"
	if _, ok := i["cos1"]; ok {
		result["cos1"] = flattenObjectSystemNpuIsfNpQueuesCos1Osna(i["cos1"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos2"
	if _, ok := i["cos2"]; ok {
		result["cos2"] = flattenObjectSystemNpuIsfNpQueuesCos2Osna(i["cos2"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos3"
	if _, ok := i["cos3"]; ok {
		result["cos3"] = flattenObjectSystemNpuIsfNpQueuesCos3Osna(i["cos3"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos4"
	if _, ok := i["cos4"]; ok {
		result["cos4"] = flattenObjectSystemNpuIsfNpQueuesCos4Osna(i["cos4"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos5"
	if _, ok := i["cos5"]; ok {
		result["cos5"] = flattenObjectSystemNpuIsfNpQueuesCos5Osna(i["cos5"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos6"
	if _, ok := i["cos6"]; ok {
		result["cos6"] = flattenObjectSystemNpuIsfNpQueuesCos6Osna(i["cos6"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos7"
	if _, ok := i["cos7"]; ok {
		result["cos7"] = flattenObjectSystemNpuIsfNpQueuesCos7Osna(i["cos7"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuIsfNpQueuesCos0Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos1Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos2Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos3Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos4Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos5Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos6Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueuesCos7Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuLagOutPortSelectOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMaxSessionTimeoutOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMcastSessionAccountingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMcastSessionCountingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMcastSessionCounting6Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNapiBreakIntervalOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNat46ForceIpv4PacketForwardingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ethernet_type"
	if _, ok := i["ethernet-type"]; ok {
		result["ethernet_type"] = flattenObjectSystemNpuNpQueuesEthernetTypeOsna(i["ethernet-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := i["ip-protocol"]; ok {
		result["ip_protocol"] = flattenObjectSystemNpuNpQueuesIpProtocolOsna(i["ip-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_service"
	if _, ok := i["ip-service"]; ok {
		result["ip_service"] = flattenObjectSystemNpuNpQueuesIpServiceOsna(i["ip-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenObjectSystemNpuNpQueuesProfileOsna(i["profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "scheduler"
	if _, ok := i["scheduler"]; ok {
		result["scheduler"] = flattenObjectSystemNpuNpQueuesSchedulerOsna(i["scheduler"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpQueuesEthernetTypeOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesEthernetTypeNameOsna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeQueueOsna(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeTypeOsna(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeWeightOsna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesEthernetTypeNameOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeQueueOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeTypeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeWeightOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesIpProtocolNameOsna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolProtocolOsna(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolQueueOsna(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolWeightOsna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpProtocolNameOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolProtocolOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolQueueOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolWeightOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := i["dport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceDportOsna(i["dport"], d, pre_append)
			tmp["dport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Dport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceNameOsna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceProtocolOsna(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceQueueOsna(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := i["sport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceSportOsna(i["sport"], d, pre_append)
			tmp["sport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Sport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceWeightOsna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpServiceDportOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceNameOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceProtocolOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceQueueOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceSportOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceWeightOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := i["cos0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos0Osna(i["cos0"], d, pre_append)
			tmp["cos0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := i["cos1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos1Osna(i["cos1"], d, pre_append)
			tmp["cos1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := i["cos2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos2Osna(i["cos2"], d, pre_append)
			tmp["cos2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := i["cos3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos3Osna(i["cos3"], d, pre_append)
			tmp["cos3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := i["cos4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos4Osna(i["cos4"], d, pre_append)
			tmp["cos4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := i["cos5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos5Osna(i["cos5"], d, pre_append)
			tmp["cos5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := i["cos6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos6Osna(i["cos6"], d, pre_append)
			tmp["cos6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := i["cos7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos7Osna(i["cos7"], d, pre_append)
			tmp["cos7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := i["dscp0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp0Osna(i["dscp0"], d, pre_append)
			tmp["dscp0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := i["dscp1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp1Osna(i["dscp1"], d, pre_append)
			tmp["dscp1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := i["dscp10"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp10Osna(i["dscp10"], d, pre_append)
			tmp["dscp10"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp10")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := i["dscp11"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp11Osna(i["dscp11"], d, pre_append)
			tmp["dscp11"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp11")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := i["dscp12"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp12Osna(i["dscp12"], d, pre_append)
			tmp["dscp12"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := i["dscp13"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp13Osna(i["dscp13"], d, pre_append)
			tmp["dscp13"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp13")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := i["dscp14"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp14Osna(i["dscp14"], d, pre_append)
			tmp["dscp14"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp14")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := i["dscp15"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp15Osna(i["dscp15"], d, pre_append)
			tmp["dscp15"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp15")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := i["dscp16"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp16Osna(i["dscp16"], d, pre_append)
			tmp["dscp16"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp16")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := i["dscp17"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp17Osna(i["dscp17"], d, pre_append)
			tmp["dscp17"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp17")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := i["dscp18"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp18Osna(i["dscp18"], d, pre_append)
			tmp["dscp18"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp18")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := i["dscp19"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp19Osna(i["dscp19"], d, pre_append)
			tmp["dscp19"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp19")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := i["dscp2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp2Osna(i["dscp2"], d, pre_append)
			tmp["dscp2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := i["dscp20"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp20Osna(i["dscp20"], d, pre_append)
			tmp["dscp20"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp20")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := i["dscp21"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp21Osna(i["dscp21"], d, pre_append)
			tmp["dscp21"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp21")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := i["dscp22"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp22Osna(i["dscp22"], d, pre_append)
			tmp["dscp22"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp22")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := i["dscp23"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp23Osna(i["dscp23"], d, pre_append)
			tmp["dscp23"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp23")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := i["dscp24"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp24Osna(i["dscp24"], d, pre_append)
			tmp["dscp24"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp24")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := i["dscp25"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp25Osna(i["dscp25"], d, pre_append)
			tmp["dscp25"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp25")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := i["dscp26"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp26Osna(i["dscp26"], d, pre_append)
			tmp["dscp26"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp26")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := i["dscp27"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp27Osna(i["dscp27"], d, pre_append)
			tmp["dscp27"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp27")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := i["dscp28"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp28Osna(i["dscp28"], d, pre_append)
			tmp["dscp28"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp28")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := i["dscp29"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp29Osna(i["dscp29"], d, pre_append)
			tmp["dscp29"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp29")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := i["dscp3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp3Osna(i["dscp3"], d, pre_append)
			tmp["dscp3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := i["dscp30"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp30Osna(i["dscp30"], d, pre_append)
			tmp["dscp30"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp30")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := i["dscp31"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp31Osna(i["dscp31"], d, pre_append)
			tmp["dscp31"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp31")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := i["dscp32"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp32Osna(i["dscp32"], d, pre_append)
			tmp["dscp32"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp32")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := i["dscp33"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp33Osna(i["dscp33"], d, pre_append)
			tmp["dscp33"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp33")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := i["dscp34"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp34Osna(i["dscp34"], d, pre_append)
			tmp["dscp34"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := i["dscp35"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp35Osna(i["dscp35"], d, pre_append)
			tmp["dscp35"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp35")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := i["dscp36"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp36Osna(i["dscp36"], d, pre_append)
			tmp["dscp36"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp36")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := i["dscp37"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp37Osna(i["dscp37"], d, pre_append)
			tmp["dscp37"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp37")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := i["dscp38"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp38Osna(i["dscp38"], d, pre_append)
			tmp["dscp38"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp38")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := i["dscp39"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp39Osna(i["dscp39"], d, pre_append)
			tmp["dscp39"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp39")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := i["dscp4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp4Osna(i["dscp4"], d, pre_append)
			tmp["dscp4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := i["dscp40"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp40Osna(i["dscp40"], d, pre_append)
			tmp["dscp40"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp40")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := i["dscp41"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp41Osna(i["dscp41"], d, pre_append)
			tmp["dscp41"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp41")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := i["dscp42"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp42Osna(i["dscp42"], d, pre_append)
			tmp["dscp42"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp42")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := i["dscp43"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp43Osna(i["dscp43"], d, pre_append)
			tmp["dscp43"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp43")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := i["dscp44"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp44Osna(i["dscp44"], d, pre_append)
			tmp["dscp44"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := i["dscp45"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp45Osna(i["dscp45"], d, pre_append)
			tmp["dscp45"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp45")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := i["dscp46"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp46Osna(i["dscp46"], d, pre_append)
			tmp["dscp46"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := i["dscp47"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp47Osna(i["dscp47"], d, pre_append)
			tmp["dscp47"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp47")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := i["dscp48"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp48Osna(i["dscp48"], d, pre_append)
			tmp["dscp48"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp48")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := i["dscp49"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp49Osna(i["dscp49"], d, pre_append)
			tmp["dscp49"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp49")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := i["dscp5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp5Osna(i["dscp5"], d, pre_append)
			tmp["dscp5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := i["dscp50"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp50Osna(i["dscp50"], d, pre_append)
			tmp["dscp50"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp50")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := i["dscp51"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp51Osna(i["dscp51"], d, pre_append)
			tmp["dscp51"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp51")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := i["dscp52"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp52Osna(i["dscp52"], d, pre_append)
			tmp["dscp52"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp52")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := i["dscp53"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp53Osna(i["dscp53"], d, pre_append)
			tmp["dscp53"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp53")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := i["dscp54"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp54Osna(i["dscp54"], d, pre_append)
			tmp["dscp54"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp54")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := i["dscp55"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp55Osna(i["dscp55"], d, pre_append)
			tmp["dscp55"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp55")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := i["dscp56"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp56Osna(i["dscp56"], d, pre_append)
			tmp["dscp56"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp56")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := i["dscp57"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp57Osna(i["dscp57"], d, pre_append)
			tmp["dscp57"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp57")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := i["dscp58"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp58Osna(i["dscp58"], d, pre_append)
			tmp["dscp58"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp58")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := i["dscp59"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp59Osna(i["dscp59"], d, pre_append)
			tmp["dscp59"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp59")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := i["dscp6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp6Osna(i["dscp6"], d, pre_append)
			tmp["dscp6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := i["dscp60"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp60Osna(i["dscp60"], d, pre_append)
			tmp["dscp60"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp60")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := i["dscp61"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp61Osna(i["dscp61"], d, pre_append)
			tmp["dscp61"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp61")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := i["dscp62"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp62Osna(i["dscp62"], d, pre_append)
			tmp["dscp62"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp62")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := i["dscp63"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp63Osna(i["dscp63"], d, pre_append)
			tmp["dscp63"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp63")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := i["dscp7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp7Osna(i["dscp7"], d, pre_append)
			tmp["dscp7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := i["dscp8"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp8Osna(i["dscp8"], d, pre_append)
			tmp["dscp8"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp8")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := i["dscp9"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp9Osna(i["dscp9"], d, pre_append)
			tmp["dscp9"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp9")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileIdOsna(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileTypeOsna(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileWeightOsna(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesProfileCos0Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos1Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos2Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos3Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos4Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos5Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos6Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos7Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp0Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp1Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp10Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp11Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp12Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp13Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp14Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp15Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp16Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp17Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp18Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp19Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp2Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp20Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp21Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp22Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp23Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp24Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp25Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp26Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp27Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp28Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp29Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp3Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp30Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp31Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp32Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp33Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp34Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp35Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp36Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp37Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp38Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp39Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp4Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp40Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp41Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp42Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp43Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp44Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp45Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp46Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp47Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp48Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp49Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp5Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp50Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp51Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp52Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp53Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp54Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp55Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp56Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp57Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp58Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp59Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp6Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp60Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp61Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp62Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp63Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp7Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp8Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp9Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileIdOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileTypeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileWeightOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesSchedulerOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerModeOsna(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerNameOsna(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesSchedulerModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesSchedulerNameOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNp6CpsOptimizationModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPbaEimOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPerPolicyAccountingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPerSessionAccountingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPolicyOffloadLevelOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortCpuMapOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cpu_core"
		if _, ok := i["cpu-core"]; ok {
			v := flattenObjectSystemNpuPortCpuMapCpuCoreOsna(i["cpu-core"], d, pre_append)
			tmp["cpu_core"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortCpuMap-CpuCore")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectSystemNpuPortCpuMapInterfaceOsna(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortCpuMap-Interface")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuPortCpuMapCpuCoreOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortCpuMapInterfaceOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortNpuMapOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectSystemNpuPortNpuMapInterfaceOsna(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortNpuMap-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "npu_group_index"
		if _, ok := i["npu-group-index"]; ok {
			v := flattenObjectSystemNpuPortNpuMapNpuGroupIndexOsna(i["npu-group-index"], d, pre_append)
			tmp["npu_group_index"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortNpuMap-NpuGroupIndex")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuPortNpuMapInterfaceOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortNpuMapNpuGroupIndexOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortPathOptionOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports_using_npu"
	if _, ok := i["ports-using-npu"]; ok {
		result["ports_using_npu"] = flattenObjectSystemNpuPortPathOptionPortsUsingNpuOsna(i["ports-using-npu"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuPortPathOptionPortsUsingNpuOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemNpuPriorityProtocolOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bfd"
	if _, ok := i["bfd"]; ok {
		result["bfd"] = flattenObjectSystemNpuPriorityProtocolBfdOsna(i["bfd"], d, pre_append)
	}

	pre_append = pre + ".0." + "bgp"
	if _, ok := i["bgp"]; ok {
		result["bgp"] = flattenObjectSystemNpuPriorityProtocolBgpOsna(i["bgp"], d, pre_append)
	}

	pre_append = pre + ".0." + "slbc"
	if _, ok := i["slbc"]; ok {
		result["slbc"] = flattenObjectSystemNpuPriorityProtocolSlbcOsna(i["slbc"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuPriorityProtocolBfdOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolBgpOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolSlbcOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuProcessIcmpByHostOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPrpPortInOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPrpPortOutOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuQosModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuQtmBufModeOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuRdpOffloadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuRecoverNp6LinkOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSessionAcctIntervalOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSessionDeniedOffloadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseBackpressureOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuStripClearTextPaddingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuStripEspPaddingOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "computation"
	if _, ok := i["computation"]; ok {
		result["computation"] = flattenObjectSystemNpuSwEhHashComputationOsna(i["computation"], d, pre_append)
	}

	pre_append = pre + ".0." + "destination_ip_lower_16"
	if _, ok := i["destination-ip-lower-16"]; ok {
		result["destination_ip_lower_16"] = flattenObjectSystemNpuSwEhHashDestinationIpLower16Osna(i["destination-ip-lower-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "destination_ip_upper_16"
	if _, ok := i["destination-ip-upper-16"]; ok {
		result["destination_ip_upper_16"] = flattenObjectSystemNpuSwEhHashDestinationIpUpper16Osna(i["destination-ip-upper-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "destination_port"
	if _, ok := i["destination-port"]; ok {
		result["destination_port"] = flattenObjectSystemNpuSwEhHashDestinationPortOsna(i["destination-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := i["ip-protocol"]; ok {
		result["ip_protocol"] = flattenObjectSystemNpuSwEhHashIpProtocolOsna(i["ip-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "netmask_length"
	if _, ok := i["netmask-length"]; ok {
		result["netmask_length"] = flattenObjectSystemNpuSwEhHashNetmaskLengthOsna(i["netmask-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_ip_lower_16"
	if _, ok := i["source-ip-lower-16"]; ok {
		result["source_ip_lower_16"] = flattenObjectSystemNpuSwEhHashSourceIpLower16Osna(i["source-ip-lower-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_ip_upper_16"
	if _, ok := i["source-ip-upper-16"]; ok {
		result["source_ip_upper_16"] = flattenObjectSystemNpuSwEhHashSourceIpUpper16Osna(i["source-ip-upper-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_port"
	if _, ok := i["source-port"]; ok {
		result["source_port"] = flattenObjectSystemNpuSwEhHashSourcePortOsna(i["source-port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuSwEhHashComputationOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationIpLower16Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationIpUpper16Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationPortOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashIpProtocolOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashNetmaskLengthOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourceIpLower16Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourceIpUpper16Osna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourcePortOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwNpBandwidthOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwitchNpHashOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpRstTimeoutOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "close_wait"
		if _, ok := i["close-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileCloseWaitOsna(i["close-wait"], d, pre_append)
			tmp["close_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-CloseWait")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fin_wait"
		if _, ok := i["fin-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileFinWaitOsna(i["fin-wait"], d, pre_append)
			tmp["fin_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-FinWait")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileIdOsna(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_sent"
		if _, ok := i["syn-sent"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileSynSentOsna(i["syn-sent"], d, pre_append)
			tmp["syn_sent"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-SynSent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_wait"
		if _, ok := i["syn-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileSynWaitOsna(i["syn-wait"], d, pre_append)
			tmp["syn_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-SynWait")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_idle"
		if _, ok := i["tcp-idle"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileTcpIdleOsna(i["tcp-idle"], d, pre_append)
			tmp["tcp_idle"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-TcpIdle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_wait"
		if _, ok := i["time-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileTimeWaitOsna(i["time-wait"], d, pre_append)
			tmp["time_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-TimeWait")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuTcpTimeoutProfileCloseWaitOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileFinWaitOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileIdOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileSynSentOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileSynWaitOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileTcpIdleOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileTimeWaitOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUdpTimeoutProfileOsna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuUdpTimeoutProfileIdOsna(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-UdpTimeoutProfile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "udp_idle"
		if _, ok := i["udp-idle"]; ok {
			v := flattenObjectSystemNpuUdpTimeoutProfileUdpIdleOsna(i["udp-idle"], d, pre_append)
			tmp["udp_idle"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-UdpTimeoutProfile-UdpIdle")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemNpuUdpTimeoutProfileIdOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUdpTimeoutProfileUdpIdleOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUespOffloadOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuVlanLookupCacheOsna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpu(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("background_sse_scan", flattenObjectSystemNpuBackgroundSseScanOsna(o["background-sse-scan"], d, "background_sse_scan")); err != nil {
			if vv, ok := fortiAPIPatch(o["background-sse-scan"], "ObjectSystemNpu-BackgroundSseScan"); ok {
				if err = d.Set("background_sse_scan", vv); err != nil {
					return fmt.Errorf("Error reading background_sse_scan: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading background_sse_scan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("background_sse_scan"); ok {
			if err = d.Set("background_sse_scan", flattenObjectSystemNpuBackgroundSseScanOsna(o["background-sse-scan"], d, "background_sse_scan")); err != nil {
				if vv, ok := fortiAPIPatch(o["background-sse-scan"], "ObjectSystemNpu-BackgroundSseScan"); ok {
					if err = d.Set("background_sse_scan", vv); err != nil {
						return fmt.Errorf("Error reading background_sse_scan: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading background_sse_scan: %v", err)
				}
			}
		}
	}

	if err = d.Set("capwap_offload", flattenObjectSystemNpuCapwapOffloadOsna(o["capwap-offload"], d, "capwap_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["capwap-offload"], "ObjectSystemNpu-CapwapOffload"); ok {
			if err = d.Set("capwap_offload", vv); err != nil {
				return fmt.Errorf("Error reading capwap_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capwap_offload: %v", err)
		}
	}

	if err = d.Set("dedicated_management_affinity", flattenObjectSystemNpuDedicatedManagementAffinityOsna(o["dedicated-management-affinity"], d, "dedicated_management_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["dedicated-management-affinity"], "ObjectSystemNpu-DedicatedManagementAffinity"); ok {
			if err = d.Set("dedicated_management_affinity", vv); err != nil {
				return fmt.Errorf("Error reading dedicated_management_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dedicated_management_affinity: %v", err)
		}
	}

	if err = d.Set("dedicated_management_cpu", flattenObjectSystemNpuDedicatedManagementCpuOsna(o["dedicated-management-cpu"], d, "dedicated_management_cpu")); err != nil {
		if vv, ok := fortiAPIPatch(o["dedicated-management-cpu"], "ObjectSystemNpu-DedicatedManagementCpu"); ok {
			if err = d.Set("dedicated_management_cpu", vv); err != nil {
				return fmt.Errorf("Error reading dedicated_management_cpu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dedicated_management_cpu: %v", err)
		}
	}

	if err = d.Set("default_qos_type", flattenObjectSystemNpuDefaultQosTypeOsna(o["default-qos-type"], d, "default_qos_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-qos-type"], "ObjectSystemNpu-DefaultQosType"); ok {
			if err = d.Set("default_qos_type", vv); err != nil {
				return fmt.Errorf("Error reading default_qos_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_qos_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dos_options", flattenObjectSystemNpuDosOptionsOsna(o["dos-options"], d, "dos_options")); err != nil {
			if vv, ok := fortiAPIPatch(o["dos-options"], "ObjectSystemNpu-DosOptions"); ok {
				if err = d.Set("dos_options", vv); err != nil {
					return fmt.Errorf("Error reading dos_options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dos_options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dos_options"); ok {
			if err = d.Set("dos_options", flattenObjectSystemNpuDosOptionsOsna(o["dos-options"], d, "dos_options")); err != nil {
				if vv, ok := fortiAPIPatch(o["dos-options"], "ObjectSystemNpu-DosOptions"); ok {
					if err = d.Set("dos_options", vv); err != nil {
						return fmt.Errorf("Error reading dos_options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dos_options: %v", err)
				}
			}
		}
	}

	if err = d.Set("double_level_mcast_offload", flattenObjectSystemNpuDoubleLevelMcastOffloadOsna(o["double-level-mcast-offload"], d, "double_level_mcast_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["double-level-mcast-offload"], "ObjectSystemNpu-DoubleLevelMcastOffload"); ok {
			if err = d.Set("double_level_mcast_offload", vv); err != nil {
				return fmt.Errorf("Error reading double_level_mcast_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading double_level_mcast_offload: %v", err)
		}
	}

	if err = d.Set("dse_timeout", flattenObjectSystemNpuDseTimeoutOsna(o["dse-timeout"], d, "dse_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dse-timeout"], "ObjectSystemNpu-DseTimeout"); ok {
			if err = d.Set("dse_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dse_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dse_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dsw_dts_profile", flattenObjectSystemNpuDswDtsProfileOsna(o["dsw-dts-profile"], d, "dsw_dts_profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["dsw-dts-profile"], "ObjectSystemNpu-DswDtsProfile"); ok {
				if err = d.Set("dsw_dts_profile", vv); err != nil {
					return fmt.Errorf("Error reading dsw_dts_profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dsw_dts_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dsw_dts_profile"); ok {
			if err = d.Set("dsw_dts_profile", flattenObjectSystemNpuDswDtsProfileOsna(o["dsw-dts-profile"], d, "dsw_dts_profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["dsw-dts-profile"], "ObjectSystemNpu-DswDtsProfile"); ok {
					if err = d.Set("dsw_dts_profile", vv); err != nil {
						return fmt.Errorf("Error reading dsw_dts_profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dsw_dts_profile: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dsw_queue_dts_profile", flattenObjectSystemNpuDswQueueDtsProfileOsna(o["dsw-queue-dts-profile"], d, "dsw_queue_dts_profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["dsw-queue-dts-profile"], "ObjectSystemNpu-DswQueueDtsProfile"); ok {
				if err = d.Set("dsw_queue_dts_profile", vv); err != nil {
					return fmt.Errorf("Error reading dsw_queue_dts_profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dsw_queue_dts_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dsw_queue_dts_profile"); ok {
			if err = d.Set("dsw_queue_dts_profile", flattenObjectSystemNpuDswQueueDtsProfileOsna(o["dsw-queue-dts-profile"], d, "dsw_queue_dts_profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["dsw-queue-dts-profile"], "ObjectSystemNpu-DswQueueDtsProfile"); ok {
					if err = d.Set("dsw_queue_dts_profile", vv); err != nil {
						return fmt.Errorf("Error reading dsw_queue_dts_profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dsw_queue_dts_profile: %v", err)
				}
			}
		}
	}

	if err = d.Set("fastpath", flattenObjectSystemNpuFastpathOsna(o["fastpath"], d, "fastpath")); err != nil {
		if vv, ok := fortiAPIPatch(o["fastpath"], "ObjectSystemNpu-Fastpath"); ok {
			if err = d.Set("fastpath", vv); err != nil {
				return fmt.Errorf("Error reading fastpath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fastpath: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fp_anomaly", flattenObjectSystemNpuFpAnomalyOsna(o["fp-anomaly"], d, "fp_anomaly")); err != nil {
			if vv, ok := fortiAPIPatch(o["fp-anomaly"], "ObjectSystemNpu-FpAnomaly"); ok {
				if err = d.Set("fp_anomaly", vv); err != nil {
					return fmt.Errorf("Error reading fp_anomaly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fp_anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fp_anomaly"); ok {
			if err = d.Set("fp_anomaly", flattenObjectSystemNpuFpAnomalyOsna(o["fp-anomaly"], d, "fp_anomaly")); err != nil {
				if vv, ok := fortiAPIPatch(o["fp-anomaly"], "ObjectSystemNpu-FpAnomaly"); ok {
					if err = d.Set("fp_anomaly", vv); err != nil {
						return fmt.Errorf("Error reading fp_anomaly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fp_anomaly: %v", err)
				}
			}
		}
	}

	if err = d.Set("gtp_enhanced_cpu_range", flattenObjectSystemNpuGtpEnhancedCpuRangeOsna(o["gtp-enhanced-cpu-range"], d, "gtp_enhanced_cpu_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-enhanced-cpu-range"], "ObjectSystemNpu-GtpEnhancedCpuRange"); ok {
			if err = d.Set("gtp_enhanced_cpu_range", vv); err != nil {
				return fmt.Errorf("Error reading gtp_enhanced_cpu_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_enhanced_cpu_range: %v", err)
		}
	}

	if err = d.Set("gtp_enhanced_mode", flattenObjectSystemNpuGtpEnhancedModeOsna(o["gtp-enhanced-mode"], d, "gtp_enhanced_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-enhanced-mode"], "ObjectSystemNpu-GtpEnhancedMode"); ok {
			if err = d.Set("gtp_enhanced_mode", vv); err != nil {
				return fmt.Errorf("Error reading gtp_enhanced_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_enhanced_mode: %v", err)
		}
	}

	if err = d.Set("gtp_support", flattenObjectSystemNpuGtpSupportOsna(o["gtp-support"], d, "gtp_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-support"], "ObjectSystemNpu-GtpSupport"); ok {
			if err = d.Set("gtp_support", vv); err != nil {
				return fmt.Errorf("Error reading gtp_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_support: %v", err)
		}
	}

	if err = d.Set("hash_config", flattenObjectSystemNpuHashConfigOsna(o["hash-config"], d, "hash_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-config"], "ObjectSystemNpu-HashConfig"); ok {
			if err = d.Set("hash_config", vv); err != nil {
				return fmt.Errorf("Error reading hash_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_config: %v", err)
		}
	}

	if err = d.Set("hash_tbl_spread", flattenObjectSystemNpuHashTblSpreadOsna(o["hash-tbl-spread"], d, "hash_tbl_spread")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-tbl-spread"], "ObjectSystemNpu-HashTblSpread"); ok {
			if err = d.Set("hash_tbl_spread", vv); err != nil {
				return fmt.Errorf("Error reading hash_tbl_spread: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_tbl_spread: %v", err)
		}
	}

	if err = d.Set("host_shortcut_mode", flattenObjectSystemNpuHostShortcutModeOsna(o["host-shortcut-mode"], d, "host_shortcut_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-shortcut-mode"], "ObjectSystemNpu-HostShortcutMode"); ok {
			if err = d.Set("host_shortcut_mode", vv); err != nil {
				return fmt.Errorf("Error reading host_shortcut_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_shortcut_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hpe", flattenObjectSystemNpuHpeOsna(o["hpe"], d, "hpe")); err != nil {
			if vv, ok := fortiAPIPatch(o["hpe"], "ObjectSystemNpu-Hpe"); ok {
				if err = d.Set("hpe", vv); err != nil {
					return fmt.Errorf("Error reading hpe: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hpe: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hpe"); ok {
			if err = d.Set("hpe", flattenObjectSystemNpuHpeOsna(o["hpe"], d, "hpe")); err != nil {
				if vv, ok := fortiAPIPatch(o["hpe"], "ObjectSystemNpu-Hpe"); ok {
					if err = d.Set("hpe", vv); err != nil {
						return fmt.Errorf("Error reading hpe: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hpe: %v", err)
				}
			}
		}
	}

	if err = d.Set("htab_dedi_queue_nr", flattenObjectSystemNpuHtabDediQueueNrOsna(o["htab-dedi-queue-nr"], d, "htab_dedi_queue_nr")); err != nil {
		if vv, ok := fortiAPIPatch(o["htab-dedi-queue-nr"], "ObjectSystemNpu-HtabDediQueueNr"); ok {
			if err = d.Set("htab_dedi_queue_nr", vv); err != nil {
				return fmt.Errorf("Error reading htab_dedi_queue_nr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htab_dedi_queue_nr: %v", err)
		}
	}

	if err = d.Set("htab_msg_queue", flattenObjectSystemNpuHtabMsgQueueOsna(o["htab-msg-queue"], d, "htab_msg_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["htab-msg-queue"], "ObjectSystemNpu-HtabMsgQueue"); ok {
			if err = d.Set("htab_msg_queue", vv); err != nil {
				return fmt.Errorf("Error reading htab_msg_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htab_msg_queue: %v", err)
		}
	}

	if err = d.Set("htx_gtse_quota", flattenObjectSystemNpuHtxGtseQuotaOsna(o["htx-gtse-quota"], d, "htx_gtse_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["htx-gtse-quota"], "ObjectSystemNpu-HtxGtseQuota"); ok {
			if err = d.Set("htx_gtse_quota", vv); err != nil {
				return fmt.Errorf("Error reading htx_gtse_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htx_gtse_quota: %v", err)
		}
	}

	if err = d.Set("htx_icmp_csum_chk", flattenObjectSystemNpuHtxIcmpCsumChkOsna(o["htx-icmp-csum-chk"], d, "htx_icmp_csum_chk")); err != nil {
		if vv, ok := fortiAPIPatch(o["htx-icmp-csum-chk"], "ObjectSystemNpu-HtxIcmpCsumChk"); ok {
			if err = d.Set("htx_icmp_csum_chk", vv); err != nil {
				return fmt.Errorf("Error reading htx_icmp_csum_chk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htx_icmp_csum_chk: %v", err)
		}
	}

	if err = d.Set("hw_ha_scan_interval", flattenObjectSystemNpuHwHaScanIntervalOsna(o["hw-ha-scan-interval"], d, "hw_ha_scan_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-ha-scan-interval"], "ObjectSystemNpu-HwHaScanInterval"); ok {
			if err = d.Set("hw_ha_scan_interval", vv); err != nil {
				return fmt.Errorf("Error reading hw_ha_scan_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_ha_scan_interval: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenObjectSystemNpuInboundDscpCopyOsna(o["inbound-dscp-copy"], d, "inbound_dscp_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy"], "ObjectSystemNpu-InboundDscpCopy"); ok {
			if err = d.Set("inbound_dscp_copy", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy_port", flattenObjectSystemNpuInboundDscpCopyPortOsna(o["inbound-dscp-copy-port"], d, "inbound_dscp_copy_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy-port"], "ObjectSystemNpu-InboundDscpCopyPort"); ok {
			if err = d.Set("inbound_dscp_copy_port", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_reassembly", flattenObjectSystemNpuIpReassemblyOsna(o["ip-reassembly"], d, "ip_reassembly")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-reassembly"], "ObjectSystemNpu-IpReassembly"); ok {
				if err = d.Set("ip_reassembly", vv); err != nil {
					return fmt.Errorf("Error reading ip_reassembly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_reassembly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_reassembly"); ok {
			if err = d.Set("ip_reassembly", flattenObjectSystemNpuIpReassemblyOsna(o["ip-reassembly"], d, "ip_reassembly")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-reassembly"], "ObjectSystemNpu-IpReassembly"); ok {
					if err = d.Set("ip_reassembly", vv); err != nil {
						return fmt.Errorf("Error reading ip_reassembly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_reassembly: %v", err)
				}
			}
		}
	}

	if err = d.Set("intf_shaping_offload", flattenObjectSystemNpuIntfShapingOffloadOsna(o["intf-shaping-offload"], d, "intf_shaping_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf-shaping-offload"], "ObjectSystemNpu-IntfShapingOffload"); ok {
			if err = d.Set("intf_shaping_offload", vv); err != nil {
				return fmt.Errorf("Error reading intf_shaping_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf_shaping_offload: %v", err)
		}
	}

	if err = d.Set("iph_rsvd_re_cksum", flattenObjectSystemNpuIphRsvdReCksumOsna(o["iph-rsvd-re-cksum"], d, "iph_rsvd_re_cksum")); err != nil {
		if vv, ok := fortiAPIPatch(o["iph-rsvd-re-cksum"], "ObjectSystemNpu-IphRsvdReCksum"); ok {
			if err = d.Set("iph_rsvd_re_cksum", vv); err != nil {
				return fmt.Errorf("Error reading iph_rsvd_re_cksum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iph_rsvd_re_cksum: %v", err)
		}
	}

	if err = d.Set("ippool_overload_high", flattenObjectSystemNpuIppoolOverloadHighOsna(o["ippool-overload-high"], d, "ippool_overload_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool-overload-high"], "ObjectSystemNpu-IppoolOverloadHigh"); ok {
			if err = d.Set("ippool_overload_high", vv); err != nil {
				return fmt.Errorf("Error reading ippool_overload_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool_overload_high: %v", err)
		}
	}

	if err = d.Set("ippool_overload_low", flattenObjectSystemNpuIppoolOverloadLowOsna(o["ippool-overload-low"], d, "ippool_overload_low")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool-overload-low"], "ObjectSystemNpu-IppoolOverloadLow"); ok {
			if err = d.Set("ippool_overload_low", vv); err != nil {
				return fmt.Errorf("Error reading ippool_overload_low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool_overload_low: %v", err)
		}
	}

	if err = d.Set("ipsec_dec_subengine_mask", flattenObjectSystemNpuIpsecDecSubengineMaskOsna(o["ipsec-dec-subengine-mask"], d, "ipsec_dec_subengine_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-dec-subengine-mask"], "ObjectSystemNpu-IpsecDecSubengineMask"); ok {
			if err = d.Set("ipsec_dec_subengine_mask", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_dec_subengine_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_dec_subengine_mask: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_subengine_mask", flattenObjectSystemNpuIpsecEncSubengineMaskOsna(o["ipsec-enc-subengine-mask"], d, "ipsec_enc_subengine_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-enc-subengine-mask"], "ObjectSystemNpu-IpsecEncSubengineMask"); ok {
			if err = d.Set("ipsec_enc_subengine_mask", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_enc_subengine_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_enc_subengine_mask: %v", err)
		}
	}

	if err = d.Set("ipsec_host_dfclr", flattenObjectSystemNpuIpsecHostDfclrOsna(o["ipsec-host-dfclr"], d, "ipsec_host_dfclr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-host-dfclr"], "ObjectSystemNpu-IpsecHostDfclr"); ok {
			if err = d.Set("ipsec_host_dfclr", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_host_dfclr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_host_dfclr: %v", err)
		}
	}

	if err = d.Set("ipsec_inbound_cache", flattenObjectSystemNpuIpsecInboundCacheOsna(o["ipsec-inbound-cache"], d, "ipsec_inbound_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-inbound-cache"], "ObjectSystemNpu-IpsecInboundCache"); ok {
			if err = d.Set("ipsec_inbound_cache", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_inbound_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_inbound_cache: %v", err)
		}
	}

	if err = d.Set("ipsec_local_uesp_port", flattenObjectSystemNpuIpsecLocalUespPortOsna(o["ipsec-local-uesp-port"], d, "ipsec_local_uesp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-local-uesp-port"], "ObjectSystemNpu-IpsecLocalUespPort"); ok {
			if err = d.Set("ipsec_local_uesp_port", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_local_uesp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_local_uesp_port: %v", err)
		}
	}

	if err = d.Set("ipsec_mtu_override", flattenObjectSystemNpuIpsecMtuOverrideOsna(o["ipsec-mtu-override"], d, "ipsec_mtu_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-mtu-override"], "ObjectSystemNpu-IpsecMtuOverride"); ok {
			if err = d.Set("ipsec_mtu_override", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_mtu_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_mtu_override: %v", err)
		}
	}

	if err = d.Set("ipsec_ob_np_sel", flattenObjectSystemNpuIpsecObNpSelOsna(o["ipsec-ob-np-sel"], d, "ipsec_ob_np_sel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-ob-np-sel"], "ObjectSystemNpu-IpsecObNpSel"); ok {
			if err = d.Set("ipsec_ob_np_sel", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_ob_np_sel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_ob_np_sel: %v", err)
		}
	}

	if err = d.Set("ipsec_over_vlink", flattenObjectSystemNpuIpsecOverVlinkOsna(o["ipsec-over-vlink"], d, "ipsec_over_vlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-over-vlink"], "ObjectSystemNpu-IpsecOverVlink"); ok {
			if err = d.Set("ipsec_over_vlink", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_over_vlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_over_vlink: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("isf_np_queues", flattenObjectSystemNpuIsfNpQueuesOsna(o["isf-np-queues"], d, "isf_np_queues")); err != nil {
			if vv, ok := fortiAPIPatch(o["isf-np-queues"], "ObjectSystemNpu-IsfNpQueues"); ok {
				if err = d.Set("isf_np_queues", vv); err != nil {
					return fmt.Errorf("Error reading isf_np_queues: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading isf_np_queues: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("isf_np_queues"); ok {
			if err = d.Set("isf_np_queues", flattenObjectSystemNpuIsfNpQueuesOsna(o["isf-np-queues"], d, "isf_np_queues")); err != nil {
				if vv, ok := fortiAPIPatch(o["isf-np-queues"], "ObjectSystemNpu-IsfNpQueues"); ok {
					if err = d.Set("isf_np_queues", vv); err != nil {
						return fmt.Errorf("Error reading isf_np_queues: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading isf_np_queues: %v", err)
				}
			}
		}
	}

	if err = d.Set("lag_out_port_select", flattenObjectSystemNpuLagOutPortSelectOsna(o["lag-out-port-select"], d, "lag_out_port_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["lag-out-port-select"], "ObjectSystemNpu-LagOutPortSelect"); ok {
			if err = d.Set("lag_out_port_select", vv); err != nil {
				return fmt.Errorf("Error reading lag_out_port_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lag_out_port_select: %v", err)
		}
	}

	if err = d.Set("max_session_timeout", flattenObjectSystemNpuMaxSessionTimeoutOsna(o["max-session-timeout"], d, "max_session_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-session-timeout"], "ObjectSystemNpu-MaxSessionTimeout"); ok {
			if err = d.Set("max_session_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_session_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_session_timeout: %v", err)
		}
	}

	if err = d.Set("mcast_session_accounting", flattenObjectSystemNpuMcastSessionAccountingOsna(o["mcast-session-accounting"], d, "mcast_session_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-session-accounting"], "ObjectSystemNpu-McastSessionAccounting"); ok {
			if err = d.Set("mcast_session_accounting", vv); err != nil {
				return fmt.Errorf("Error reading mcast_session_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_session_accounting: %v", err)
		}
	}

	if err = d.Set("mcast_session_counting", flattenObjectSystemNpuMcastSessionCountingOsna(o["mcast-session-counting"], d, "mcast_session_counting")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-session-counting"], "ObjectSystemNpu-McastSessionCounting"); ok {
			if err = d.Set("mcast_session_counting", vv); err != nil {
				return fmt.Errorf("Error reading mcast_session_counting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_session_counting: %v", err)
		}
	}

	if err = d.Set("mcast_session_counting6", flattenObjectSystemNpuMcastSessionCounting6Osna(o["mcast-session-counting6"], d, "mcast_session_counting6")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-session-counting6"], "ObjectSystemNpu-McastSessionCounting6"); ok {
			if err = d.Set("mcast_session_counting6", vv); err != nil {
				return fmt.Errorf("Error reading mcast_session_counting6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_session_counting6: %v", err)
		}
	}

	if err = d.Set("napi_break_interval", flattenObjectSystemNpuNapiBreakIntervalOsna(o["napi-break-interval"], d, "napi_break_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["napi-break-interval"], "ObjectSystemNpu-NapiBreakInterval"); ok {
			if err = d.Set("napi_break_interval", vv); err != nil {
				return fmt.Errorf("Error reading napi_break_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading napi_break_interval: %v", err)
		}
	}

	if err = d.Set("nat46_force_ipv4_packet_forwarding", flattenObjectSystemNpuNat46ForceIpv4PacketForwardingOsna(o["nat46-force-ipv4-packet-forwarding"], d, "nat46_force_ipv4_packet_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46-force-ipv4-packet-forwarding"], "ObjectSystemNpu-Nat46ForceIpv4PacketForwarding"); ok {
			if err = d.Set("nat46_force_ipv4_packet_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("np_queues", flattenObjectSystemNpuNpQueuesOsna(o["np-queues"], d, "np_queues")); err != nil {
			if vv, ok := fortiAPIPatch(o["np-queues"], "ObjectSystemNpu-NpQueues"); ok {
				if err = d.Set("np_queues", vv); err != nil {
					return fmt.Errorf("Error reading np_queues: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading np_queues: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("np_queues"); ok {
			if err = d.Set("np_queues", flattenObjectSystemNpuNpQueuesOsna(o["np-queues"], d, "np_queues")); err != nil {
				if vv, ok := fortiAPIPatch(o["np-queues"], "ObjectSystemNpu-NpQueues"); ok {
					if err = d.Set("np_queues", vv); err != nil {
						return fmt.Errorf("Error reading np_queues: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading np_queues: %v", err)
				}
			}
		}
	}

	if err = d.Set("np6_cps_optimization_mode", flattenObjectSystemNpuNp6CpsOptimizationModeOsna(o["np6-cps-optimization-mode"], d, "np6_cps_optimization_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["np6-cps-optimization-mode"], "ObjectSystemNpu-Np6CpsOptimizationMode"); ok {
			if err = d.Set("np6_cps_optimization_mode", vv); err != nil {
				return fmt.Errorf("Error reading np6_cps_optimization_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np6_cps_optimization_mode: %v", err)
		}
	}

	if err = d.Set("pba_eim", flattenObjectSystemNpuPbaEimOsna(o["pba-eim"], d, "pba_eim")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-eim"], "ObjectSystemNpu-PbaEim"); ok {
			if err = d.Set("pba_eim", vv); err != nil {
				return fmt.Errorf("Error reading pba_eim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_eim: %v", err)
		}
	}

	if err = d.Set("per_policy_accounting", flattenObjectSystemNpuPerPolicyAccountingOsna(o["per-policy-accounting"], d, "per_policy_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-policy-accounting"], "ObjectSystemNpu-PerPolicyAccounting"); ok {
			if err = d.Set("per_policy_accounting", vv); err != nil {
				return fmt.Errorf("Error reading per_policy_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_policy_accounting: %v", err)
		}
	}

	if err = d.Set("per_session_accounting", flattenObjectSystemNpuPerSessionAccountingOsna(o["per-session-accounting"], d, "per_session_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-session-accounting"], "ObjectSystemNpu-PerSessionAccounting"); ok {
			if err = d.Set("per_session_accounting", vv); err != nil {
				return fmt.Errorf("Error reading per_session_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_session_accounting: %v", err)
		}
	}

	if err = d.Set("policy_offload_level", flattenObjectSystemNpuPolicyOffloadLevelOsna(o["policy-offload-level"], d, "policy_offload_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload-level"], "ObjectSystemNpu-PolicyOffloadLevel"); ok {
			if err = d.Set("policy_offload_level", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload_level: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port_cpu_map", flattenObjectSystemNpuPortCpuMapOsna(o["port-cpu-map"], d, "port_cpu_map")); err != nil {
			if vv, ok := fortiAPIPatch(o["port-cpu-map"], "ObjectSystemNpu-PortCpuMap"); ok {
				if err = d.Set("port_cpu_map", vv); err != nil {
					return fmt.Errorf("Error reading port_cpu_map: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading port_cpu_map: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_cpu_map"); ok {
			if err = d.Set("port_cpu_map", flattenObjectSystemNpuPortCpuMapOsna(o["port-cpu-map"], d, "port_cpu_map")); err != nil {
				if vv, ok := fortiAPIPatch(o["port-cpu-map"], "ObjectSystemNpu-PortCpuMap"); ok {
					if err = d.Set("port_cpu_map", vv); err != nil {
						return fmt.Errorf("Error reading port_cpu_map: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading port_cpu_map: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("port_npu_map", flattenObjectSystemNpuPortNpuMapOsna(o["port-npu-map"], d, "port_npu_map")); err != nil {
			if vv, ok := fortiAPIPatch(o["port-npu-map"], "ObjectSystemNpu-PortNpuMap"); ok {
				if err = d.Set("port_npu_map", vv); err != nil {
					return fmt.Errorf("Error reading port_npu_map: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading port_npu_map: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_npu_map"); ok {
			if err = d.Set("port_npu_map", flattenObjectSystemNpuPortNpuMapOsna(o["port-npu-map"], d, "port_npu_map")); err != nil {
				if vv, ok := fortiAPIPatch(o["port-npu-map"], "ObjectSystemNpu-PortNpuMap"); ok {
					if err = d.Set("port_npu_map", vv); err != nil {
						return fmt.Errorf("Error reading port_npu_map: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading port_npu_map: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("port_path_option", flattenObjectSystemNpuPortPathOptionOsna(o["port-path-option"], d, "port_path_option")); err != nil {
			if vv, ok := fortiAPIPatch(o["port-path-option"], "ObjectSystemNpu-PortPathOption"); ok {
				if err = d.Set("port_path_option", vv); err != nil {
					return fmt.Errorf("Error reading port_path_option: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading port_path_option: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_path_option"); ok {
			if err = d.Set("port_path_option", flattenObjectSystemNpuPortPathOptionOsna(o["port-path-option"], d, "port_path_option")); err != nil {
				if vv, ok := fortiAPIPatch(o["port-path-option"], "ObjectSystemNpu-PortPathOption"); ok {
					if err = d.Set("port_path_option", vv); err != nil {
						return fmt.Errorf("Error reading port_path_option: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading port_path_option: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("priority_protocol", flattenObjectSystemNpuPriorityProtocolOsna(o["priority-protocol"], d, "priority_protocol")); err != nil {
			if vv, ok := fortiAPIPatch(o["priority-protocol"], "ObjectSystemNpu-PriorityProtocol"); ok {
				if err = d.Set("priority_protocol", vv); err != nil {
					return fmt.Errorf("Error reading priority_protocol: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading priority_protocol: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("priority_protocol"); ok {
			if err = d.Set("priority_protocol", flattenObjectSystemNpuPriorityProtocolOsna(o["priority-protocol"], d, "priority_protocol")); err != nil {
				if vv, ok := fortiAPIPatch(o["priority-protocol"], "ObjectSystemNpu-PriorityProtocol"); ok {
					if err = d.Set("priority_protocol", vv); err != nil {
						return fmt.Errorf("Error reading priority_protocol: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading priority_protocol: %v", err)
				}
			}
		}
	}

	if err = d.Set("process_icmp_by_host", flattenObjectSystemNpuProcessIcmpByHostOsna(o["process-icmp-by-host"], d, "process_icmp_by_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["process-icmp-by-host"], "ObjectSystemNpu-ProcessIcmpByHost"); ok {
			if err = d.Set("process_icmp_by_host", vv); err != nil {
				return fmt.Errorf("Error reading process_icmp_by_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading process_icmp_by_host: %v", err)
		}
	}

	if err = d.Set("prp_port_in", flattenObjectSystemNpuPrpPortInOsna(o["prp-port-in"], d, "prp_port_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["prp-port-in"], "ObjectSystemNpu-PrpPortIn"); ok {
			if err = d.Set("prp_port_in", vv); err != nil {
				return fmt.Errorf("Error reading prp_port_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prp_port_in: %v", err)
		}
	}

	if err = d.Set("prp_port_out", flattenObjectSystemNpuPrpPortOutOsna(o["prp-port-out"], d, "prp_port_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["prp-port-out"], "ObjectSystemNpu-PrpPortOut"); ok {
			if err = d.Set("prp_port_out", vv); err != nil {
				return fmt.Errorf("Error reading prp_port_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prp_port_out: %v", err)
		}
	}

	if err = d.Set("qos_mode", flattenObjectSystemNpuQosModeOsna(o["qos-mode"], d, "qos_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-mode"], "ObjectSystemNpu-QosMode"); ok {
			if err = d.Set("qos_mode", vv); err != nil {
				return fmt.Errorf("Error reading qos_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_mode: %v", err)
		}
	}

	if err = d.Set("qtm_buf_mode", flattenObjectSystemNpuQtmBufModeOsna(o["qtm-buf-mode"], d, "qtm_buf_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["qtm-buf-mode"], "ObjectSystemNpu-QtmBufMode"); ok {
			if err = d.Set("qtm_buf_mode", vv); err != nil {
				return fmt.Errorf("Error reading qtm_buf_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qtm_buf_mode: %v", err)
		}
	}

	if err = d.Set("rdp_offload", flattenObjectSystemNpuRdpOffloadOsna(o["rdp-offload"], d, "rdp_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdp-offload"], "ObjectSystemNpu-RdpOffload"); ok {
			if err = d.Set("rdp_offload", vv); err != nil {
				return fmt.Errorf("Error reading rdp_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdp_offload: %v", err)
		}
	}

	if err = d.Set("recover_np6_link", flattenObjectSystemNpuRecoverNp6LinkOsna(o["recover-np6-link"], d, "recover_np6_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["recover-np6-link"], "ObjectSystemNpu-RecoverNp6Link"); ok {
			if err = d.Set("recover_np6_link", vv); err != nil {
				return fmt.Errorf("Error reading recover_np6_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recover_np6_link: %v", err)
		}
	}

	if err = d.Set("session_acct_interval", flattenObjectSystemNpuSessionAcctIntervalOsna(o["session-acct-interval"], d, "session_acct_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-acct-interval"], "ObjectSystemNpu-SessionAcctInterval"); ok {
			if err = d.Set("session_acct_interval", vv); err != nil {
				return fmt.Errorf("Error reading session_acct_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_acct_interval: %v", err)
		}
	}

	if err = d.Set("session_denied_offload", flattenObjectSystemNpuSessionDeniedOffloadOsna(o["session-denied-offload"], d, "session_denied_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-denied-offload"], "ObjectSystemNpu-SessionDeniedOffload"); ok {
			if err = d.Set("session_denied_offload", vv); err != nil {
				return fmt.Errorf("Error reading session_denied_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_denied_offload: %v", err)
		}
	}

	if err = d.Set("sse_backpressure", flattenObjectSystemNpuSseBackpressureOsna(o["sse-backpressure"], d, "sse_backpressure")); err != nil {
		if vv, ok := fortiAPIPatch(o["sse-backpressure"], "ObjectSystemNpu-SseBackpressure"); ok {
			if err = d.Set("sse_backpressure", vv); err != nil {
				return fmt.Errorf("Error reading sse_backpressure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sse_backpressure: %v", err)
		}
	}

	if err = d.Set("strip_clear_text_padding", flattenObjectSystemNpuStripClearTextPaddingOsna(o["strip-clear-text-padding"], d, "strip_clear_text_padding")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-clear-text-padding"], "ObjectSystemNpu-StripClearTextPadding"); ok {
			if err = d.Set("strip_clear_text_padding", vv); err != nil {
				return fmt.Errorf("Error reading strip_clear_text_padding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_clear_text_padding: %v", err)
		}
	}

	if err = d.Set("strip_esp_padding", flattenObjectSystemNpuStripEspPaddingOsna(o["strip-esp-padding"], d, "strip_esp_padding")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-esp-padding"], "ObjectSystemNpu-StripEspPadding"); ok {
			if err = d.Set("strip_esp_padding", vv); err != nil {
				return fmt.Errorf("Error reading strip_esp_padding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_esp_padding: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sw_eh_hash", flattenObjectSystemNpuSwEhHashOsna(o["sw-eh-hash"], d, "sw_eh_hash")); err != nil {
			if vv, ok := fortiAPIPatch(o["sw-eh-hash"], "ObjectSystemNpu-SwEhHash"); ok {
				if err = d.Set("sw_eh_hash", vv); err != nil {
					return fmt.Errorf("Error reading sw_eh_hash: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sw_eh_hash: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sw_eh_hash"); ok {
			if err = d.Set("sw_eh_hash", flattenObjectSystemNpuSwEhHashOsna(o["sw-eh-hash"], d, "sw_eh_hash")); err != nil {
				if vv, ok := fortiAPIPatch(o["sw-eh-hash"], "ObjectSystemNpu-SwEhHash"); ok {
					if err = d.Set("sw_eh_hash", vv); err != nil {
						return fmt.Errorf("Error reading sw_eh_hash: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sw_eh_hash: %v", err)
				}
			}
		}
	}

	if err = d.Set("sw_np_bandwidth", flattenObjectSystemNpuSwNpBandwidthOsna(o["sw-np-bandwidth"], d, "sw_np_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-np-bandwidth"], "ObjectSystemNpu-SwNpBandwidth"); ok {
			if err = d.Set("sw_np_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading sw_np_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_np_bandwidth: %v", err)
		}
	}

	if err = d.Set("switch_np_hash", flattenObjectSystemNpuSwitchNpHashOsna(o["switch-np-hash"], d, "switch_np_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-np-hash"], "ObjectSystemNpu-SwitchNpHash"); ok {
			if err = d.Set("switch_np_hash", vv); err != nil {
				return fmt.Errorf("Error reading switch_np_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_np_hash: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timeout", flattenObjectSystemNpuTcpRstTimeoutOsna(o["tcp-rst-timeout"], d, "tcp_rst_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst-timeout"], "ObjectSystemNpu-TcpRstTimeout"); ok {
			if err = d.Set("tcp_rst_timeout", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tcp_timeout_profile", flattenObjectSystemNpuTcpTimeoutProfileOsna(o["tcp-timeout-profile"], d, "tcp_timeout_profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["tcp-timeout-profile"], "ObjectSystemNpu-TcpTimeoutProfile"); ok {
				if err = d.Set("tcp_timeout_profile", vv); err != nil {
					return fmt.Errorf("Error reading tcp_timeout_profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tcp_timeout_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tcp_timeout_profile"); ok {
			if err = d.Set("tcp_timeout_profile", flattenObjectSystemNpuTcpTimeoutProfileOsna(o["tcp-timeout-profile"], d, "tcp_timeout_profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["tcp-timeout-profile"], "ObjectSystemNpu-TcpTimeoutProfile"); ok {
					if err = d.Set("tcp_timeout_profile", vv); err != nil {
						return fmt.Errorf("Error reading tcp_timeout_profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tcp_timeout_profile: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("udp_timeout_profile", flattenObjectSystemNpuUdpTimeoutProfileOsna(o["udp-timeout-profile"], d, "udp_timeout_profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["udp-timeout-profile"], "ObjectSystemNpu-UdpTimeoutProfile"); ok {
				if err = d.Set("udp_timeout_profile", vv); err != nil {
					return fmt.Errorf("Error reading udp_timeout_profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading udp_timeout_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("udp_timeout_profile"); ok {
			if err = d.Set("udp_timeout_profile", flattenObjectSystemNpuUdpTimeoutProfileOsna(o["udp-timeout-profile"], d, "udp_timeout_profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["udp-timeout-profile"], "ObjectSystemNpu-UdpTimeoutProfile"); ok {
					if err = d.Set("udp_timeout_profile", vv); err != nil {
						return fmt.Errorf("Error reading udp_timeout_profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading udp_timeout_profile: %v", err)
				}
			}
		}
	}

	if err = d.Set("uesp_offload", flattenObjectSystemNpuUespOffloadOsna(o["uesp-offload"], d, "uesp_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["uesp-offload"], "ObjectSystemNpu-UespOffload"); ok {
			if err = d.Set("uesp_offload", vv); err != nil {
				return fmt.Errorf("Error reading uesp_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uesp_offload: %v", err)
		}
	}

	if err = d.Set("vlan_lookup_cache", flattenObjectSystemNpuVlanLookupCacheOsna(o["vlan-lookup-cache"], d, "vlan_lookup_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-lookup-cache"], "ObjectSystemNpu-VlanLookupCache"); ok {
			if err = d.Set("vlan_lookup_cache", vv); err != nil {
				return fmt.Errorf("Error reading vlan_lookup_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_lookup_cache: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuBackgroundSseScanOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan"], _ = expandObjectSystemNpuBackgroundSseScanScanOsna(d, i["scan"], pre_append)
	}
	pre_append = pre + ".0." + "stats_update_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["stats-update-interval"], _ = expandObjectSystemNpuBackgroundSseScanStatsUpdateIntervalOsna(d, i["stats_update_interval"], pre_append)
	}
	pre_append = pre + ".0." + "udp_keepalive_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-keepalive-interval"], _ = expandObjectSystemNpuBackgroundSseScanUdpKeepaliveIntervalOsna(d, i["udp_keepalive_interval"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuBackgroundSseScanScanOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsUpdateIntervalOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpKeepaliveIntervalOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuCapwapOffloadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDedicatedManagementAffinityOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDedicatedManagementCpuOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDefaultQosTypeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "npu_dos_meter_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["npu-dos-meter-mode"], _ = expandObjectSystemNpuDosOptionsNpuDosMeterModeOsna(d, i["npu_dos_meter_mode"], pre_append)
	}
	pre_append = pre + ".0." + "npu_dos_synproxy_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["npu-dos-synproxy-mode"], _ = expandObjectSystemNpuDosOptionsNpuDosSynproxyModeOsna(d, i["npu_dos_synproxy_mode"], pre_append)
	}
	pre_append = pre + ".0." + "npu_dos_tpe_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["npu-dos-tpe-mode"], _ = expandObjectSystemNpuDosOptionsNpuDosTpeModeOsna(d, i["npu_dos_tpe_mode"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuDosOptionsNpuDosMeterModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsNpuDosSynproxyModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsNpuDosTpeModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDoubleLevelMcastOffloadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDseTimeoutOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectSystemNpuDswDtsProfileActionOsna(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_limit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-limit"], _ = expandObjectSystemNpuDswDtsProfileMinLimitOsna(d, i["min_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["profile-id"], _ = expandObjectSystemNpuDswDtsProfileProfileIdOsna(d, i["profile_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "step"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["step"], _ = expandObjectSystemNpuDswDtsProfileStepOsna(d, i["step"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuDswDtsProfileActionOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileMinLimitOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileProfileIdOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileStepOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["iport"], _ = expandObjectSystemNpuDswQueueDtsProfileIportOsna(d, i["iport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSystemNpuDswQueueDtsProfileNameOsna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oport"], _ = expandObjectSystemNpuDswQueueDtsProfileOportOsna(d, i["oport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["profile-id"], _ = expandObjectSystemNpuDswQueueDtsProfileProfileIdOsna(d, i["profile_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue_select"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue-select"], _ = expandObjectSystemNpuDswQueueDtsProfileQueueSelectOsna(d, i["queue_select"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuDswQueueDtsProfileIportOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileNameOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileOportOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileProfileIdOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileQueueSelectOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFastpathOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "capwap_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["capwap-minlen-err"], _ = expandObjectSystemNpuFpAnomalyCapwapMinlenErrOsna(d, i["capwap_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "esp_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["esp-minlen-err"], _ = expandObjectSystemNpuFpAnomalyEspMinlenErrOsna(d, i["esp_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "gre_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["gre-csum-err"], _ = expandObjectSystemNpuFpAnomalyGreCsumErrOsna(d, i["gre_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "gtpu_plen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["gtpu-plen-err"], _ = expandObjectSystemNpuFpAnomalyGtpuPlenErrOsna(d, i["gtpu_plen_err"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["icmp-csum-err"], _ = expandObjectSystemNpuFpAnomalyIcmpCsumErrOsna(d, i["icmp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_frag"
	if _, ok := d.GetOk(pre_append); ok {
		result["icmp-frag"], _ = expandObjectSystemNpuFpAnomalyIcmpFragOsna(d, i["icmp_frag"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_land"
	if _, ok := d.GetOk(pre_append); ok {
		result["icmp-land"], _ = expandObjectSystemNpuFpAnomalyIcmpLandOsna(d, i["icmp_land"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["icmp-minlen-err"], _ = expandObjectSystemNpuFpAnomalyIcmpMinlenErrOsna(d, i["icmp_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-csum-err"], _ = expandObjectSystemNpuFpAnomalyIpv4CsumErrOsna(d, i["ipv4_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_ihl_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-ihl-err"], _ = expandObjectSystemNpuFpAnomalyIpv4IhlErrOsna(d, i["ipv4_ihl_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_land"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-land"], _ = expandObjectSystemNpuFpAnomalyIpv4LandOsna(d, i["ipv4_land"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_len_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-len-err"], _ = expandObjectSystemNpuFpAnomalyIpv4LenErrOsna(d, i["ipv4_len_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_opt_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-opt-err"], _ = expandObjectSystemNpuFpAnomalyIpv4OptErrOsna(d, i["ipv4_opt_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optlsrr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-optlsrr"], _ = expandObjectSystemNpuFpAnomalyIpv4OptlsrrOsna(d, i["ipv4_optlsrr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optrr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-optrr"], _ = expandObjectSystemNpuFpAnomalyIpv4OptrrOsna(d, i["ipv4_optrr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optsecurity"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-optsecurity"], _ = expandObjectSystemNpuFpAnomalyIpv4OptsecurityOsna(d, i["ipv4_optsecurity"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optssrr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-optssrr"], _ = expandObjectSystemNpuFpAnomalyIpv4OptssrrOsna(d, i["ipv4_optssrr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optstream"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-optstream"], _ = expandObjectSystemNpuFpAnomalyIpv4OptstreamOsna(d, i["ipv4_optstream"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_opttimestamp"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-opttimestamp"], _ = expandObjectSystemNpuFpAnomalyIpv4OpttimestampOsna(d, i["ipv4_opttimestamp"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_proto_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-proto-err"], _ = expandObjectSystemNpuFpAnomalyIpv4ProtoErrOsna(d, i["ipv4_proto_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_ttlzero_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-ttlzero-err"], _ = expandObjectSystemNpuFpAnomalyIpv4TtlzeroErrOsna(d, i["ipv4_ttlzero_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_unknopt"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-unknopt"], _ = expandObjectSystemNpuFpAnomalyIpv4UnknoptOsna(d, i["ipv4_unknopt"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_ver_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv4-ver-err"], _ = expandObjectSystemNpuFpAnomalyIpv4VerErrOsna(d, i["ipv4_ver_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_daddr_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-daddr-err"], _ = expandObjectSystemNpuFpAnomalyIpv6DaddrErrOsna(d, i["ipv6_daddr_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_exthdr_len_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-exthdr-len-err"], _ = expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErrOsna(d, i["ipv6_exthdr_len_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_exthdr_order_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-exthdr-order-err"], _ = expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErrOsna(d, i["ipv6_exthdr_order_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_ihl_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-ihl-err"], _ = expandObjectSystemNpuFpAnomalyIpv6IhlErrOsna(d, i["ipv6_ihl_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_land"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-land"], _ = expandObjectSystemNpuFpAnomalyIpv6LandOsna(d, i["ipv6_land"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optendpid"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-optendpid"], _ = expandObjectSystemNpuFpAnomalyIpv6OptendpidOsna(d, i["ipv6_optendpid"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_opthomeaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-opthomeaddr"], _ = expandObjectSystemNpuFpAnomalyIpv6OpthomeaddrOsna(d, i["ipv6_opthomeaddr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optinvld"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-optinvld"], _ = expandObjectSystemNpuFpAnomalyIpv6OptinvldOsna(d, i["ipv6_optinvld"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optjumbo"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-optjumbo"], _ = expandObjectSystemNpuFpAnomalyIpv6OptjumboOsna(d, i["ipv6_optjumbo"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optnsap"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-optnsap"], _ = expandObjectSystemNpuFpAnomalyIpv6OptnsapOsna(d, i["ipv6_optnsap"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optralert"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-optralert"], _ = expandObjectSystemNpuFpAnomalyIpv6OptralertOsna(d, i["ipv6_optralert"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_opttunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-opttunnel"], _ = expandObjectSystemNpuFpAnomalyIpv6OpttunnelOsna(d, i["ipv6_opttunnel"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_plen_zero"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-plen-zero"], _ = expandObjectSystemNpuFpAnomalyIpv6PlenZeroOsna(d, i["ipv6_plen_zero"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_proto_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-proto-err"], _ = expandObjectSystemNpuFpAnomalyIpv6ProtoErrOsna(d, i["ipv6_proto_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_saddr_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-saddr-err"], _ = expandObjectSystemNpuFpAnomalyIpv6SaddrErrOsna(d, i["ipv6_saddr_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_unknopt"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-unknopt"], _ = expandObjectSystemNpuFpAnomalyIpv6UnknoptOsna(d, i["ipv6_unknopt"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_ver_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipv6-ver-err"], _ = expandObjectSystemNpuFpAnomalyIpv6VerErrOsna(d, i["ipv6_ver_err"], pre_append)
	}
	pre_append = pre + ".0." + "nvgre_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["nvgre-minlen-err"], _ = expandObjectSystemNpuFpAnomalyNvgreMinlenErrOsna(d, i["nvgre_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_clen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["sctp-clen-err"], _ = expandObjectSystemNpuFpAnomalySctpClenErrOsna(d, i["sctp_clen_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_crc_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["sctp-crc-err"], _ = expandObjectSystemNpuFpAnomalySctpCrcErrOsna(d, i["sctp_crc_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_l4len_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["sctp-l4len-err"], _ = expandObjectSystemNpuFpAnomalySctpL4LenErrOsna(d, i["sctp_l4len_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-csum-err"], _ = expandObjectSystemNpuFpAnomalyTcpCsumErrOsna(d, i["tcp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin_noack"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-fin-noack"], _ = expandObjectSystemNpuFpAnomalyTcpFinNoackOsna(d, i["tcp_fin_noack"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin_only"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-fin-only"], _ = expandObjectSystemNpuFpAnomalyTcpFinOnlyOsna(d, i["tcp_fin_only"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_hlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-hlen-err"], _ = expandObjectSystemNpuFpAnomalyTcpHlenErrOsna(d, i["tcp_hlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_hlenvsl4len_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-hlenvsl4len-err"], _ = expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErrOsna(d, i["tcp_hlenvsl4len_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_land"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-land"], _ = expandObjectSystemNpuFpAnomalyTcpLandOsna(d, i["tcp_land"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_no_flag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-no-flag"], _ = expandObjectSystemNpuFpAnomalyTcpNoFlagOsna(d, i["tcp_no_flag"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_plen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-plen-err"], _ = expandObjectSystemNpuFpAnomalyTcpPlenErrOsna(d, i["tcp_plen_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-syn-data"], _ = expandObjectSystemNpuFpAnomalyTcpSynDataOsna(d, i["tcp_syn_data"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn_fin"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-syn-fin"], _ = expandObjectSystemNpuFpAnomalyTcpSynFinOsna(d, i["tcp_syn_fin"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_winnuke"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-winnuke"], _ = expandObjectSystemNpuFpAnomalyTcpWinnukeOsna(d, i["tcp_winnuke"], pre_append)
	}
	pre_append = pre + ".0." + "udp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-csum-err"], _ = expandObjectSystemNpuFpAnomalyUdpCsumErrOsna(d, i["udp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "udp_hlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-hlen-err"], _ = expandObjectSystemNpuFpAnomalyUdpHlenErrOsna(d, i["udp_hlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "udp_land"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-land"], _ = expandObjectSystemNpuFpAnomalyUdpLandOsna(d, i["udp_land"], pre_append)
	}
	pre_append = pre + ".0." + "udp_len_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-len-err"], _ = expandObjectSystemNpuFpAnomalyUdpLenErrOsna(d, i["udp_len_err"], pre_append)
	}
	pre_append = pre + ".0." + "udp_plen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-plen-err"], _ = expandObjectSystemNpuFpAnomalyUdpPlenErrOsna(d, i["udp_plen_err"], pre_append)
	}
	pre_append = pre + ".0." + "udplite_cover_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["udplite-cover-err"], _ = expandObjectSystemNpuFpAnomalyUdpliteCoverErrOsna(d, i["udplite_cover_err"], pre_append)
	}
	pre_append = pre + ".0." + "udplite_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["udplite-csum-err"], _ = expandObjectSystemNpuFpAnomalyUdpliteCsumErrOsna(d, i["udplite_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "uesp_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["uesp-minlen-err"], _ = expandObjectSystemNpuFpAnomalyUespMinlenErrOsna(d, i["uesp_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "unknproto_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["unknproto-minlen-err"], _ = expandObjectSystemNpuFpAnomalyUnknprotoMinlenErrOsna(d, i["unknproto_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "vxlan_minlen_err"
	if _, ok := d.GetOk(pre_append); ok {
		result["vxlan-minlen-err"], _ = expandObjectSystemNpuFpAnomalyVxlanMinlenErrOsna(d, i["vxlan_minlen_err"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuFpAnomalyCapwapMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyEspMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGreCsumErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGtpuPlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpCsumErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpFragOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpLandOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4CsumErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4IhlErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4LandOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4LenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptlsrrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptrrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptsecurityOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptssrrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptstreamOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OpttimestampOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4ProtoErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4TtlzeroErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4UnknoptOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4VerErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6DaddrErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6IhlErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6LandOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OptendpidOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OpthomeaddrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OptinvldOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OptjumboOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OptnsapOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OptralertOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6OpttunnelOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6PlenZeroOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ProtoErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6SaddrErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6UnknoptOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6VerErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyNvgreMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpClenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpCrcErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpL4LenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpCsumErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinNoackOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinOnlyOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpLandOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpNoFlagOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpPlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynDataOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynFinOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpWinnukeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpCsumErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpHlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLandOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpPlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCoverErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCsumErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUespMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUnknprotoMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyVxlanMinlenErrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuGtpEnhancedCpuRangeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuGtpEnhancedModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuGtpSupportOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHashConfigOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHashTblSpreadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHostShortcutModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["all-protocol"], _ = expandObjectSystemNpuHpeAllProtocolOsna(d, i["all_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "arp_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["arp-max"], _ = expandObjectSystemNpuHpeArpMaxOsna(d, i["arp_max"], pre_append)
	}
	pre_append = pre + ".0." + "enable_shaper"
	if _, ok := d.GetOk(pre_append); ok {
		result["enable-shaper"], _ = expandObjectSystemNpuHpeEnableShaperOsna(d, i["enable_shaper"], pre_append)
	}
	pre_append = pre + ".0." + "esp_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["esp-max"], _ = expandObjectSystemNpuHpeEspMaxOsna(d, i["esp_max"], pre_append)
	}
	pre_append = pre + ".0." + "high_priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["high-priority"], _ = expandObjectSystemNpuHpeHighPriorityOsna(d, i["high_priority"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["icmp-max"], _ = expandObjectSystemNpuHpeIcmpMaxOsna(d, i["icmp_max"], pre_append)
	}
	pre_append = pre + ".0." + "ip_frag_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip-frag-max"], _ = expandObjectSystemNpuHpeIpFragMaxOsna(d, i["ip_frag_max"], pre_append)
	}
	pre_append = pre + ".0." + "ip_others_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip-others-max"], _ = expandObjectSystemNpuHpeIpOthersMaxOsna(d, i["ip_others_max"], pre_append)
	}
	pre_append = pre + ".0." + "l2_others_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["l2-others-max"], _ = expandObjectSystemNpuHpeL2OthersMaxOsna(d, i["l2_others_max"], pre_append)
	}
	pre_append = pre + ".0." + "pri_type_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["pri-type-max"], _ = expandObjectSystemNpuHpePriTypeMaxOsna(d, i["pri_type_max"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["sctp-max"], _ = expandObjectSystemNpuHpeSctpMaxOsna(d, i["sctp_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-max"], _ = expandObjectSystemNpuHpeTcpMaxOsna(d, i["tcp_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcpfin_rst_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcpfin-rst-max"], _ = expandObjectSystemNpuHpeTcpfinRstMaxOsna(d, i["tcpfin_rst_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcpsyn_ack_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcpsyn-ack-max"], _ = expandObjectSystemNpuHpeTcpsynAckMaxOsna(d, i["tcpsyn_ack_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcpsyn_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcpsyn-max"], _ = expandObjectSystemNpuHpeTcpsynMaxOsna(d, i["tcpsyn_max"], pre_append)
	}
	pre_append = pre + ".0." + "udp_max"
	if _, ok := d.GetOk(pre_append); ok {
		result["udp-max"], _ = expandObjectSystemNpuHpeUdpMaxOsna(d, i["udp_max"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuHpeAllProtocolOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeArpMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEnableShaperOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEspMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeHighPriorityOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIcmpMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIpFragMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIpOthersMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeL2OthersMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpePriTypeMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeSctpMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpfinRstMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpsynAckMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpsynMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeUdpMaxOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtabDediQueueNrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtabMsgQueueOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtxGtseQuotaOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtxIcmpCsumChkOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHwHaScanIntervalOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuInboundDscpCopyOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuInboundDscpCopyPortOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemNpuIpReassemblyOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-timeout"], _ = expandObjectSystemNpuIpReassemblyMaxTimeoutOsna(d, i["max_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "min_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["min-timeout"], _ = expandObjectSystemNpuIpReassemblyMinTimeoutOsna(d, i["min_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectSystemNpuIpReassemblyStatusOsna(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuIpReassemblyMaxTimeoutOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpReassemblyMinTimeoutOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpReassemblyStatusOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIntfShapingOffloadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIphRsvdReCksumOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIppoolOverloadHighOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIppoolOverloadLowOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecDecSubengineMaskOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecEncSubengineMaskOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecHostDfclrOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecInboundCacheOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecLocalUespPortOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecMtuOverrideOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecObNpSelOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecOverVlinkOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cos0"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos0"], _ = expandObjectSystemNpuIsfNpQueuesCos0Osna(d, i["cos0"], pre_append)
	}
	pre_append = pre + ".0." + "cos1"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos1"], _ = expandObjectSystemNpuIsfNpQueuesCos1Osna(d, i["cos1"], pre_append)
	}
	pre_append = pre + ".0." + "cos2"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos2"], _ = expandObjectSystemNpuIsfNpQueuesCos2Osna(d, i["cos2"], pre_append)
	}
	pre_append = pre + ".0." + "cos3"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos3"], _ = expandObjectSystemNpuIsfNpQueuesCos3Osna(d, i["cos3"], pre_append)
	}
	pre_append = pre + ".0." + "cos4"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos4"], _ = expandObjectSystemNpuIsfNpQueuesCos4Osna(d, i["cos4"], pre_append)
	}
	pre_append = pre + ".0." + "cos5"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos5"], _ = expandObjectSystemNpuIsfNpQueuesCos5Osna(d, i["cos5"], pre_append)
	}
	pre_append = pre + ".0." + "cos6"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos6"], _ = expandObjectSystemNpuIsfNpQueuesCos6Osna(d, i["cos6"], pre_append)
	}
	pre_append = pre + ".0." + "cos7"
	if _, ok := d.GetOk(pre_append); ok {
		result["cos7"], _ = expandObjectSystemNpuIsfNpQueuesCos7Osna(d, i["cos7"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuIsfNpQueuesCos0Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos1Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos2Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos3Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos4Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos5Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos6Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueuesCos7Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuLagOutPortSelectOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMaxSessionTimeoutOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMcastSessionAccountingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMcastSessionCountingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMcastSessionCounting6Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNapiBreakIntervalOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNat46ForceIpv4PacketForwardingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ethernet_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["ethernet-type"], _ = expandObjectSystemNpuNpQueuesEthernetTypeOsna(d, i["ethernet_type"], pre_append)
	} else {
		result["ethernet-type"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip-protocol"], _ = expandObjectSystemNpuNpQueuesIpProtocolOsna(d, i["ip_protocol"], pre_append)
	} else {
		result["ip-protocol"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip_service"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip-service"], _ = expandObjectSystemNpuNpQueuesIpServiceOsna(d, i["ip_service"], pre_append)
	} else {
		result["ip-service"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile"], _ = expandObjectSystemNpuNpQueuesProfileOsna(d, i["profile"], pre_append)
	} else {
		result["profile"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "scheduler"
	if _, ok := d.GetOk(pre_append); ok {
		result["scheduler"], _ = expandObjectSystemNpuNpQueuesSchedulerOsna(d, i["scheduler"], pre_append)
	} else {
		result["scheduler"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesEthernetTypeNameOsna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesEthernetTypeQueueOsna(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesEthernetTypeTypeOsna(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesEthernetTypeWeightOsna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeNameOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeQueueOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeTypeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeWeightOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpProtocolNameOsna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpProtocolProtocolOsna(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpProtocolQueueOsna(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpProtocolWeightOsna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolNameOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolProtocolOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolQueueOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolWeightOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dport"], _ = expandObjectSystemNpuNpQueuesIpServiceDportOsna(d, i["dport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpServiceNameOsna(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpServiceProtocolOsna(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpServiceQueueOsna(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sport"], _ = expandObjectSystemNpuNpQueuesIpServiceSportOsna(d, i["sport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpServiceWeightOsna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpServiceDportOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceNameOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceProtocolOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceQueueOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceSportOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceWeightOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos0"], _ = expandObjectSystemNpuNpQueuesProfileCos0Osna(d, i["cos0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos1"], _ = expandObjectSystemNpuNpQueuesProfileCos1Osna(d, i["cos1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos2"], _ = expandObjectSystemNpuNpQueuesProfileCos2Osna(d, i["cos2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos3"], _ = expandObjectSystemNpuNpQueuesProfileCos3Osna(d, i["cos3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos4"], _ = expandObjectSystemNpuNpQueuesProfileCos4Osna(d, i["cos4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos5"], _ = expandObjectSystemNpuNpQueuesProfileCos5Osna(d, i["cos5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos6"], _ = expandObjectSystemNpuNpQueuesProfileCos6Osna(d, i["cos6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos7"], _ = expandObjectSystemNpuNpQueuesProfileCos7Osna(d, i["cos7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp0"], _ = expandObjectSystemNpuNpQueuesProfileDscp0Osna(d, i["dscp0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp1"], _ = expandObjectSystemNpuNpQueuesProfileDscp1Osna(d, i["dscp1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp10"], _ = expandObjectSystemNpuNpQueuesProfileDscp10Osna(d, i["dscp10"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp11"], _ = expandObjectSystemNpuNpQueuesProfileDscp11Osna(d, i["dscp11"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp12"], _ = expandObjectSystemNpuNpQueuesProfileDscp12Osna(d, i["dscp12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp13"], _ = expandObjectSystemNpuNpQueuesProfileDscp13Osna(d, i["dscp13"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp14"], _ = expandObjectSystemNpuNpQueuesProfileDscp14Osna(d, i["dscp14"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp15"], _ = expandObjectSystemNpuNpQueuesProfileDscp15Osna(d, i["dscp15"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp16"], _ = expandObjectSystemNpuNpQueuesProfileDscp16Osna(d, i["dscp16"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp17"], _ = expandObjectSystemNpuNpQueuesProfileDscp17Osna(d, i["dscp17"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp18"], _ = expandObjectSystemNpuNpQueuesProfileDscp18Osna(d, i["dscp18"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp19"], _ = expandObjectSystemNpuNpQueuesProfileDscp19Osna(d, i["dscp19"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp2"], _ = expandObjectSystemNpuNpQueuesProfileDscp2Osna(d, i["dscp2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp20"], _ = expandObjectSystemNpuNpQueuesProfileDscp20Osna(d, i["dscp20"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp21"], _ = expandObjectSystemNpuNpQueuesProfileDscp21Osna(d, i["dscp21"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp22"], _ = expandObjectSystemNpuNpQueuesProfileDscp22Osna(d, i["dscp22"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp23"], _ = expandObjectSystemNpuNpQueuesProfileDscp23Osna(d, i["dscp23"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp24"], _ = expandObjectSystemNpuNpQueuesProfileDscp24Osna(d, i["dscp24"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp25"], _ = expandObjectSystemNpuNpQueuesProfileDscp25Osna(d, i["dscp25"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp26"], _ = expandObjectSystemNpuNpQueuesProfileDscp26Osna(d, i["dscp26"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp27"], _ = expandObjectSystemNpuNpQueuesProfileDscp27Osna(d, i["dscp27"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp28"], _ = expandObjectSystemNpuNpQueuesProfileDscp28Osna(d, i["dscp28"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp29"], _ = expandObjectSystemNpuNpQueuesProfileDscp29Osna(d, i["dscp29"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp3"], _ = expandObjectSystemNpuNpQueuesProfileDscp3Osna(d, i["dscp3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp30"], _ = expandObjectSystemNpuNpQueuesProfileDscp30Osna(d, i["dscp30"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp31"], _ = expandObjectSystemNpuNpQueuesProfileDscp31Osna(d, i["dscp31"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp32"], _ = expandObjectSystemNpuNpQueuesProfileDscp32Osna(d, i["dscp32"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp33"], _ = expandObjectSystemNpuNpQueuesProfileDscp33Osna(d, i["dscp33"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp34"], _ = expandObjectSystemNpuNpQueuesProfileDscp34Osna(d, i["dscp34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp35"], _ = expandObjectSystemNpuNpQueuesProfileDscp35Osna(d, i["dscp35"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp36"], _ = expandObjectSystemNpuNpQueuesProfileDscp36Osna(d, i["dscp36"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp37"], _ = expandObjectSystemNpuNpQueuesProfileDscp37Osna(d, i["dscp37"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp38"], _ = expandObjectSystemNpuNpQueuesProfileDscp38Osna(d, i["dscp38"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp39"], _ = expandObjectSystemNpuNpQueuesProfileDscp39Osna(d, i["dscp39"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp4"], _ = expandObjectSystemNpuNpQueuesProfileDscp4Osna(d, i["dscp4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp40"], _ = expandObjectSystemNpuNpQueuesProfileDscp40Osna(d, i["dscp40"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp41"], _ = expandObjectSystemNpuNpQueuesProfileDscp41Osna(d, i["dscp41"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp42"], _ = expandObjectSystemNpuNpQueuesProfileDscp42Osna(d, i["dscp42"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp43"], _ = expandObjectSystemNpuNpQueuesProfileDscp43Osna(d, i["dscp43"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp44"], _ = expandObjectSystemNpuNpQueuesProfileDscp44Osna(d, i["dscp44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp45"], _ = expandObjectSystemNpuNpQueuesProfileDscp45Osna(d, i["dscp45"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp46"], _ = expandObjectSystemNpuNpQueuesProfileDscp46Osna(d, i["dscp46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp47"], _ = expandObjectSystemNpuNpQueuesProfileDscp47Osna(d, i["dscp47"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp48"], _ = expandObjectSystemNpuNpQueuesProfileDscp48Osna(d, i["dscp48"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp49"], _ = expandObjectSystemNpuNpQueuesProfileDscp49Osna(d, i["dscp49"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp5"], _ = expandObjectSystemNpuNpQueuesProfileDscp5Osna(d, i["dscp5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp50"], _ = expandObjectSystemNpuNpQueuesProfileDscp50Osna(d, i["dscp50"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp51"], _ = expandObjectSystemNpuNpQueuesProfileDscp51Osna(d, i["dscp51"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp52"], _ = expandObjectSystemNpuNpQueuesProfileDscp52Osna(d, i["dscp52"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp53"], _ = expandObjectSystemNpuNpQueuesProfileDscp53Osna(d, i["dscp53"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp54"], _ = expandObjectSystemNpuNpQueuesProfileDscp54Osna(d, i["dscp54"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp55"], _ = expandObjectSystemNpuNpQueuesProfileDscp55Osna(d, i["dscp55"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp56"], _ = expandObjectSystemNpuNpQueuesProfileDscp56Osna(d, i["dscp56"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp57"], _ = expandObjectSystemNpuNpQueuesProfileDscp57Osna(d, i["dscp57"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp58"], _ = expandObjectSystemNpuNpQueuesProfileDscp58Osna(d, i["dscp58"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp59"], _ = expandObjectSystemNpuNpQueuesProfileDscp59Osna(d, i["dscp59"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp6"], _ = expandObjectSystemNpuNpQueuesProfileDscp6Osna(d, i["dscp6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp60"], _ = expandObjectSystemNpuNpQueuesProfileDscp60Osna(d, i["dscp60"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp61"], _ = expandObjectSystemNpuNpQueuesProfileDscp61Osna(d, i["dscp61"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp62"], _ = expandObjectSystemNpuNpQueuesProfileDscp62Osna(d, i["dscp62"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp63"], _ = expandObjectSystemNpuNpQueuesProfileDscp63Osna(d, i["dscp63"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp7"], _ = expandObjectSystemNpuNpQueuesProfileDscp7Osna(d, i["dscp7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp8"], _ = expandObjectSystemNpuNpQueuesProfileDscp8Osna(d, i["dscp8"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp9"], _ = expandObjectSystemNpuNpQueuesProfileDscp9Osna(d, i["dscp9"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectSystemNpuNpQueuesProfileIdOsna(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesProfileTypeOsna(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesProfileWeightOsna(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesProfileCos0Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos1Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos2Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos3Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos4Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos5Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos6Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos7Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp0Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp1Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp10Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp11Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp12Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp13Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp14Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp15Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp16Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp17Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp18Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp19Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp2Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp20Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp21Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp22Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp23Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp24Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp25Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp26Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp27Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp28Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp29Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp3Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp30Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp31Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp32Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp33Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp34Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp35Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp36Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp37Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp38Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp39Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp4Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp40Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp41Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp42Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp43Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp44Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp45Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp46Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp47Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp48Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp49Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp5Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp50Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp51Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp52Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp53Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp54Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp55Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp56Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp57Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp58Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp59Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp6Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp60Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp61Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp62Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp63Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp7Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp8Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp9Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileIdOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileTypeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileWeightOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesSchedulerOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mode"], _ = expandObjectSystemNpuNpQueuesSchedulerModeOsna(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesSchedulerNameOsna(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesSchedulerModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesSchedulerNameOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNp6CpsOptimizationModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPbaEimOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPerPolicyAccountingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPerSessionAccountingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPolicyOffloadLevelOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortCpuMapOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cpu_core"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cpu-core"], _ = expandObjectSystemNpuPortCpuMapCpuCoreOsna(d, i["cpu_core"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectSystemNpuPortCpuMapInterfaceOsna(d, i["interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuPortCpuMapCpuCoreOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortCpuMapInterfaceOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortNpuMapOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectSystemNpuPortNpuMapInterfaceOsna(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "npu_group_index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["npu-group-index"], _ = expandObjectSystemNpuPortNpuMapNpuGroupIndexOsna(d, i["npu_group_index"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuPortNpuMapInterfaceOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortNpuMapNpuGroupIndexOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortPathOptionOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports_using_npu"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports-using-npu"], _ = expandObjectSystemNpuPortPathOptionPortsUsingNpuOsna(d, i["ports_using_npu"], pre_append)
	} else {
		result["ports-using-npu"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectSystemNpuPortPathOptionPortsUsingNpuOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemNpuPriorityProtocolOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bfd"
	if _, ok := d.GetOk(pre_append); ok {
		result["bfd"], _ = expandObjectSystemNpuPriorityProtocolBfdOsna(d, i["bfd"], pre_append)
	}
	pre_append = pre + ".0." + "bgp"
	if _, ok := d.GetOk(pre_append); ok {
		result["bgp"], _ = expandObjectSystemNpuPriorityProtocolBgpOsna(d, i["bgp"], pre_append)
	}
	pre_append = pre + ".0." + "slbc"
	if _, ok := d.GetOk(pre_append); ok {
		result["slbc"], _ = expandObjectSystemNpuPriorityProtocolSlbcOsna(d, i["slbc"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuPriorityProtocolBfdOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolBgpOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolSlbcOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuProcessIcmpByHostOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPrpPortInOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPrpPortOutOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuQosModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuQtmBufModeOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuRdpOffloadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuRecoverNp6LinkOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSessionAcctIntervalOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSessionDeniedOffloadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseBackpressureOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuStripClearTextPaddingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuStripEspPaddingOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "computation"
	if _, ok := d.GetOk(pre_append); ok {
		result["computation"], _ = expandObjectSystemNpuSwEhHashComputationOsna(d, i["computation"], pre_append)
	}
	pre_append = pre + ".0." + "destination_ip_lower_16"
	if _, ok := d.GetOk(pre_append); ok {
		result["destination-ip-lower-16"], _ = expandObjectSystemNpuSwEhHashDestinationIpLower16Osna(d, i["destination_ip_lower_16"], pre_append)
	}
	pre_append = pre + ".0." + "destination_ip_upper_16"
	if _, ok := d.GetOk(pre_append); ok {
		result["destination-ip-upper-16"], _ = expandObjectSystemNpuSwEhHashDestinationIpUpper16Osna(d, i["destination_ip_upper_16"], pre_append)
	}
	pre_append = pre + ".0." + "destination_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["destination-port"], _ = expandObjectSystemNpuSwEhHashDestinationPortOsna(d, i["destination_port"], pre_append)
	}
	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip-protocol"], _ = expandObjectSystemNpuSwEhHashIpProtocolOsna(d, i["ip_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "netmask_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["netmask-length"], _ = expandObjectSystemNpuSwEhHashNetmaskLengthOsna(d, i["netmask_length"], pre_append)
	}
	pre_append = pre + ".0." + "source_ip_lower_16"
	if _, ok := d.GetOk(pre_append); ok {
		result["source-ip-lower-16"], _ = expandObjectSystemNpuSwEhHashSourceIpLower16Osna(d, i["source_ip_lower_16"], pre_append)
	}
	pre_append = pre + ".0." + "source_ip_upper_16"
	if _, ok := d.GetOk(pre_append); ok {
		result["source-ip-upper-16"], _ = expandObjectSystemNpuSwEhHashSourceIpUpper16Osna(d, i["source_ip_upper_16"], pre_append)
	}
	pre_append = pre + ".0." + "source_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["source-port"], _ = expandObjectSystemNpuSwEhHashSourcePortOsna(d, i["source_port"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuSwEhHashComputationOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationIpLower16Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationIpUpper16Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationPortOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashIpProtocolOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashNetmaskLengthOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourceIpLower16Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourceIpUpper16Osna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourcePortOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwNpBandwidthOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwitchNpHashOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpRstTimeoutOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "close_wait"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["close-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileCloseWaitOsna(d, i["close_wait"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fin_wait"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fin-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileFinWaitOsna(d, i["fin_wait"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectSystemNpuTcpTimeoutProfileIdOsna(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_sent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["syn-sent"], _ = expandObjectSystemNpuTcpTimeoutProfileSynSentOsna(d, i["syn_sent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_wait"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["syn-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileSynWaitOsna(d, i["syn_wait"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_idle"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tcp-idle"], _ = expandObjectSystemNpuTcpTimeoutProfileTcpIdleOsna(d, i["tcp_idle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_wait"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["time-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileTimeWaitOsna(d, i["time_wait"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuTcpTimeoutProfileCloseWaitOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileFinWaitOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileIdOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileSynSentOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileSynWaitOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileTcpIdleOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileTimeWaitOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUdpTimeoutProfileOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectSystemNpuUdpTimeoutProfileIdOsna(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "udp_idle"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["udp-idle"], _ = expandObjectSystemNpuUdpTimeoutProfileUdpIdleOsna(d, i["udp_idle"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuUdpTimeoutProfileIdOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUdpTimeoutProfileUdpIdleOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUespOffloadOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuVlanLookupCacheOsna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpu(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("background_sse_scan"); ok {
		t, err := expandObjectSystemNpuBackgroundSseScanOsna(d, v, "background_sse_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["background-sse-scan"] = t
		}
	}

	if v, ok := d.GetOk("capwap_offload"); ok {
		t, err := expandObjectSystemNpuCapwapOffloadOsna(d, v, "capwap_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capwap-offload"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_management_affinity"); ok {
		t, err := expandObjectSystemNpuDedicatedManagementAffinityOsna(d, v, "dedicated_management_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-management-affinity"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_management_cpu"); ok {
		t, err := expandObjectSystemNpuDedicatedManagementCpuOsna(d, v, "dedicated_management_cpu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-management-cpu"] = t
		}
	}

	if v, ok := d.GetOk("default_qos_type"); ok {
		t, err := expandObjectSystemNpuDefaultQosTypeOsna(d, v, "default_qos_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-qos-type"] = t
		}
	}

	if v, ok := d.GetOk("dos_options"); ok {
		t, err := expandObjectSystemNpuDosOptionsOsna(d, v, "dos_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dos-options"] = t
		}
	}

	if v, ok := d.GetOk("double_level_mcast_offload"); ok {
		t, err := expandObjectSystemNpuDoubleLevelMcastOffloadOsna(d, v, "double_level_mcast_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["double-level-mcast-offload"] = t
		}
	}

	if v, ok := d.GetOk("dse_timeout"); ok {
		t, err := expandObjectSystemNpuDseTimeoutOsna(d, v, "dse_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dse-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dsw_dts_profile"); ok {
		t, err := expandObjectSystemNpuDswDtsProfileOsna(d, v, "dsw_dts_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsw-dts-profile"] = t
		}
	}

	if v, ok := d.GetOk("dsw_queue_dts_profile"); ok {
		t, err := expandObjectSystemNpuDswQueueDtsProfileOsna(d, v, "dsw_queue_dts_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsw-queue-dts-profile"] = t
		}
	}

	if v, ok := d.GetOk("fastpath"); ok {
		t, err := expandObjectSystemNpuFastpathOsna(d, v, "fastpath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fastpath"] = t
		}
	}

	if v, ok := d.GetOk("fp_anomaly"); ok {
		t, err := expandObjectSystemNpuFpAnomalyOsna(d, v, "fp_anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp-anomaly"] = t
		}
	}

	if v, ok := d.GetOk("gtp_enhanced_cpu_range"); ok {
		t, err := expandObjectSystemNpuGtpEnhancedCpuRangeOsna(d, v, "gtp_enhanced_cpu_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-enhanced-cpu-range"] = t
		}
	}

	if v, ok := d.GetOk("gtp_enhanced_mode"); ok {
		t, err := expandObjectSystemNpuGtpEnhancedModeOsna(d, v, "gtp_enhanced_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-enhanced-mode"] = t
		}
	}

	if v, ok := d.GetOk("gtp_support"); ok {
		t, err := expandObjectSystemNpuGtpSupportOsna(d, v, "gtp_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-support"] = t
		}
	}

	if v, ok := d.GetOk("hash_config"); ok {
		t, err := expandObjectSystemNpuHashConfigOsna(d, v, "hash_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-config"] = t
		}
	}

	if v, ok := d.GetOk("hash_tbl_spread"); ok {
		t, err := expandObjectSystemNpuHashTblSpreadOsna(d, v, "hash_tbl_spread")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-tbl-spread"] = t
		}
	}

	if v, ok := d.GetOk("host_shortcut_mode"); ok {
		t, err := expandObjectSystemNpuHostShortcutModeOsna(d, v, "host_shortcut_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-shortcut-mode"] = t
		}
	}

	if v, ok := d.GetOk("hpe"); ok {
		t, err := expandObjectSystemNpuHpeOsna(d, v, "hpe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hpe"] = t
		}
	}

	if v, ok := d.GetOk("htab_dedi_queue_nr"); ok {
		t, err := expandObjectSystemNpuHtabDediQueueNrOsna(d, v, "htab_dedi_queue_nr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htab-dedi-queue-nr"] = t
		}
	}

	if v, ok := d.GetOk("htab_msg_queue"); ok {
		t, err := expandObjectSystemNpuHtabMsgQueueOsna(d, v, "htab_msg_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htab-msg-queue"] = t
		}
	}

	if v, ok := d.GetOk("htx_gtse_quota"); ok {
		t, err := expandObjectSystemNpuHtxGtseQuotaOsna(d, v, "htx_gtse_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htx-gtse-quota"] = t
		}
	}

	if v, ok := d.GetOk("htx_icmp_csum_chk"); ok {
		t, err := expandObjectSystemNpuHtxIcmpCsumChkOsna(d, v, "htx_icmp_csum_chk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htx-icmp-csum-chk"] = t
		}
	}

	if v, ok := d.GetOk("hw_ha_scan_interval"); ok {
		t, err := expandObjectSystemNpuHwHaScanIntervalOsna(d, v, "hw_ha_scan_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-ha-scan-interval"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok {
		t, err := expandObjectSystemNpuInboundDscpCopyOsna(d, v, "inbound_dscp_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy_port"); ok {
		t, err := expandObjectSystemNpuInboundDscpCopyPortOsna(d, v, "inbound_dscp_copy_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy-port"] = t
		}
	}

	if v, ok := d.GetOk("ip_reassembly"); ok {
		t, err := expandObjectSystemNpuIpReassemblyOsna(d, v, "ip_reassembly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-reassembly"] = t
		}
	}

	if v, ok := d.GetOk("intf_shaping_offload"); ok {
		t, err := expandObjectSystemNpuIntfShapingOffloadOsna(d, v, "intf_shaping_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf-shaping-offload"] = t
		}
	}

	if v, ok := d.GetOk("iph_rsvd_re_cksum"); ok {
		t, err := expandObjectSystemNpuIphRsvdReCksumOsna(d, v, "iph_rsvd_re_cksum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iph-rsvd-re-cksum"] = t
		}
	}

	if v, ok := d.GetOk("ippool_overload_high"); ok {
		t, err := expandObjectSystemNpuIppoolOverloadHighOsna(d, v, "ippool_overload_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool-overload-high"] = t
		}
	}

	if v, ok := d.GetOk("ippool_overload_low"); ok {
		t, err := expandObjectSystemNpuIppoolOverloadLowOsna(d, v, "ippool_overload_low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool-overload-low"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_dec_subengine_mask"); ok {
		t, err := expandObjectSystemNpuIpsecDecSubengineMaskOsna(d, v, "ipsec_dec_subengine_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-dec-subengine-mask"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_enc_subengine_mask"); ok {
		t, err := expandObjectSystemNpuIpsecEncSubengineMaskOsna(d, v, "ipsec_enc_subengine_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-enc-subengine-mask"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_host_dfclr"); ok {
		t, err := expandObjectSystemNpuIpsecHostDfclrOsna(d, v, "ipsec_host_dfclr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-host-dfclr"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_inbound_cache"); ok {
		t, err := expandObjectSystemNpuIpsecInboundCacheOsna(d, v, "ipsec_inbound_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-inbound-cache"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_local_uesp_port"); ok {
		t, err := expandObjectSystemNpuIpsecLocalUespPortOsna(d, v, "ipsec_local_uesp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-local-uesp-port"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_mtu_override"); ok {
		t, err := expandObjectSystemNpuIpsecMtuOverrideOsna(d, v, "ipsec_mtu_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-mtu-override"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_ob_np_sel"); ok {
		t, err := expandObjectSystemNpuIpsecObNpSelOsna(d, v, "ipsec_ob_np_sel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-ob-np-sel"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_over_vlink"); ok {
		t, err := expandObjectSystemNpuIpsecOverVlinkOsna(d, v, "ipsec_over_vlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-over-vlink"] = t
		}
	}

	if v, ok := d.GetOk("isf_np_queues"); ok {
		t, err := expandObjectSystemNpuIsfNpQueuesOsna(d, v, "isf_np_queues")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isf-np-queues"] = t
		}
	}

	if v, ok := d.GetOk("lag_out_port_select"); ok {
		t, err := expandObjectSystemNpuLagOutPortSelectOsna(d, v, "lag_out_port_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lag-out-port-select"] = t
		}
	}

	if v, ok := d.GetOk("max_session_timeout"); ok {
		t, err := expandObjectSystemNpuMaxSessionTimeoutOsna(d, v, "max_session_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-session-timeout"] = t
		}
	}

	if v, ok := d.GetOk("mcast_session_accounting"); ok {
		t, err := expandObjectSystemNpuMcastSessionAccountingOsna(d, v, "mcast_session_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-session-accounting"] = t
		}
	}

	if v, ok := d.GetOk("mcast_session_counting"); ok {
		t, err := expandObjectSystemNpuMcastSessionCountingOsna(d, v, "mcast_session_counting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-session-counting"] = t
		}
	}

	if v, ok := d.GetOk("mcast_session_counting6"); ok {
		t, err := expandObjectSystemNpuMcastSessionCounting6Osna(d, v, "mcast_session_counting6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-session-counting6"] = t
		}
	}

	if v, ok := d.GetOk("napi_break_interval"); ok {
		t, err := expandObjectSystemNpuNapiBreakIntervalOsna(d, v, "napi_break_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["napi-break-interval"] = t
		}
	}

	if v, ok := d.GetOk("nat46_force_ipv4_packet_forwarding"); ok {
		t, err := expandObjectSystemNpuNat46ForceIpv4PacketForwardingOsna(d, v, "nat46_force_ipv4_packet_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46-force-ipv4-packet-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("np_queues"); ok {
		t, err := expandObjectSystemNpuNpQueuesOsna(d, v, "np_queues")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-queues"] = t
		}
	}

	if v, ok := d.GetOk("np6_cps_optimization_mode"); ok {
		t, err := expandObjectSystemNpuNp6CpsOptimizationModeOsna(d, v, "np6_cps_optimization_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np6-cps-optimization-mode"] = t
		}
	}

	if v, ok := d.GetOk("pba_eim"); ok {
		t, err := expandObjectSystemNpuPbaEimOsna(d, v, "pba_eim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-eim"] = t
		}
	}

	if v, ok := d.GetOk("per_policy_accounting"); ok {
		t, err := expandObjectSystemNpuPerPolicyAccountingOsna(d, v, "per_policy_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy-accounting"] = t
		}
	}

	if v, ok := d.GetOk("per_session_accounting"); ok {
		t, err := expandObjectSystemNpuPerSessionAccountingOsna(d, v, "per_session_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-session-accounting"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload_level"); ok {
		t, err := expandObjectSystemNpuPolicyOffloadLevelOsna(d, v, "policy_offload_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload-level"] = t
		}
	}

	if v, ok := d.GetOk("port_cpu_map"); ok {
		t, err := expandObjectSystemNpuPortCpuMapOsna(d, v, "port_cpu_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-cpu-map"] = t
		}
	}

	if v, ok := d.GetOk("port_npu_map"); ok {
		t, err := expandObjectSystemNpuPortNpuMapOsna(d, v, "port_npu_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-npu-map"] = t
		}
	}

	if v, ok := d.GetOk("port_path_option"); ok {
		t, err := expandObjectSystemNpuPortPathOptionOsna(d, v, "port_path_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-path-option"] = t
		}
	}

	if v, ok := d.GetOk("priority_protocol"); ok {
		t, err := expandObjectSystemNpuPriorityProtocolOsna(d, v, "priority_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-protocol"] = t
		}
	}

	if v, ok := d.GetOk("process_icmp_by_host"); ok {
		t, err := expandObjectSystemNpuProcessIcmpByHostOsna(d, v, "process_icmp_by_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["process-icmp-by-host"] = t
		}
	}

	if v, ok := d.GetOk("prp_port_in"); ok {
		t, err := expandObjectSystemNpuPrpPortInOsna(d, v, "prp_port_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-port-in"] = t
		}
	}

	if v, ok := d.GetOk("prp_port_out"); ok {
		t, err := expandObjectSystemNpuPrpPortOutOsna(d, v, "prp_port_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-port-out"] = t
		}
	}

	if v, ok := d.GetOk("qos_mode"); ok {
		t, err := expandObjectSystemNpuQosModeOsna(d, v, "qos_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-mode"] = t
		}
	}

	if v, ok := d.GetOk("qtm_buf_mode"); ok {
		t, err := expandObjectSystemNpuQtmBufModeOsna(d, v, "qtm_buf_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qtm-buf-mode"] = t
		}
	}

	if v, ok := d.GetOk("rdp_offload"); ok {
		t, err := expandObjectSystemNpuRdpOffloadOsna(d, v, "rdp_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdp-offload"] = t
		}
	}

	if v, ok := d.GetOk("recover_np6_link"); ok {
		t, err := expandObjectSystemNpuRecoverNp6LinkOsna(d, v, "recover_np6_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recover-np6-link"] = t
		}
	}

	if v, ok := d.GetOk("session_acct_interval"); ok {
		t, err := expandObjectSystemNpuSessionAcctIntervalOsna(d, v, "session_acct_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-acct-interval"] = t
		}
	}

	if v, ok := d.GetOk("session_denied_offload"); ok {
		t, err := expandObjectSystemNpuSessionDeniedOffloadOsna(d, v, "session_denied_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-denied-offload"] = t
		}
	}

	if v, ok := d.GetOk("sse_backpressure"); ok {
		t, err := expandObjectSystemNpuSseBackpressureOsna(d, v, "sse_backpressure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sse-backpressure"] = t
		}
	}

	if v, ok := d.GetOk("strip_clear_text_padding"); ok {
		t, err := expandObjectSystemNpuStripClearTextPaddingOsna(d, v, "strip_clear_text_padding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-clear-text-padding"] = t
		}
	}

	if v, ok := d.GetOk("strip_esp_padding"); ok {
		t, err := expandObjectSystemNpuStripEspPaddingOsna(d, v, "strip_esp_padding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-esp-padding"] = t
		}
	}

	if v, ok := d.GetOk("sw_eh_hash"); ok {
		t, err := expandObjectSystemNpuSwEhHashOsna(d, v, "sw_eh_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-eh-hash"] = t
		}
	}

	if v, ok := d.GetOk("sw_np_bandwidth"); ok {
		t, err := expandObjectSystemNpuSwNpBandwidthOsna(d, v, "sw_np_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-np-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("switch_np_hash"); ok {
		t, err := expandObjectSystemNpuSwitchNpHashOsna(d, v, "switch_np_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-np-hash"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst_timeout"); ok {
		t, err := expandObjectSystemNpuTcpRstTimeoutOsna(d, v, "tcp_rst_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst-timeout"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timeout_profile"); ok {
		t, err := expandObjectSystemNpuTcpTimeoutProfileOsna(d, v, "tcp_timeout_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timeout-profile"] = t
		}
	}

	if v, ok := d.GetOk("udp_timeout_profile"); ok {
		t, err := expandObjectSystemNpuUdpTimeoutProfileOsna(d, v, "udp_timeout_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-timeout-profile"] = t
		}
	}

	if v, ok := d.GetOk("uesp_offload"); ok {
		t, err := expandObjectSystemNpuUespOffloadOsna(d, v, "uesp_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uesp-offload"] = t
		}
	}

	if v, ok := d.GetOk("vlan_lookup_cache"); ok {
		t, err := expandObjectSystemNpuVlanLookupCacheOsna(d, v, "vlan_lookup_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-lookup-cache"] = t
		}
	}

	return &obj, nil
}
