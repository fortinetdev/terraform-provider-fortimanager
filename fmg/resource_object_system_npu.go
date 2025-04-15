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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"scan_stale": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"scan_vt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"stats_qual_access": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"stats_qual_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"stats_update_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"udp_keepalive_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"udp_qual_access": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"udp_qual_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"capwap_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_lacp_queue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"default_tcp_refresh_dir": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_udp_refresh_dir": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dos_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"step": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						},
						"oport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capwap_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"esp_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gre_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gtpu_plen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"sctp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nvgre_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sctp_clen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sctp_crc_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sctp_l4len_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"unknproto_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vxlan_minlen_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
			},
			"hash_ipv6_sel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
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
						"enable_queue_shaper": &schema.Schema{
							Type:     schema.TypeString,
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
						"exception_code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fragment_with_sess": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fragment_without_session": &schema.Schema{
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
						"queue_shaper_max": &schema.Schema{
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
			"icmp_error_rate_ctrl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmpv4_error_bucket_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmpv4_error_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmpv4_error_rate_limit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmpv6_error_bucket_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmpv6_error_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmpv6_error_rate_limit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"icmp_rate_ctrl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_v4_bucket_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"icmp_v4_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmp_v6_bucket_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmp_v6_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"inbound_dscp_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Computed: true,
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
			"ip_fragment_offload": &schema.Schema{
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
			},
			"ippool_overload_low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipsec_sts_timeout": &schema.Schema{
				Type:     schema.TypeString,
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
			"ipv4_session_quota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_session_quota_high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv4_session_quota_low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv6_prefix_session_quota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_prefix_session_quota_high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv6_prefix_session_quota_low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_throughput_msg_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipt_sts_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipt_throughput_msg_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isf_np_queues": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"isf_np_rx_tr_distr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lag_out_port_select": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_receive_unit": &schema.Schema{
				Type:     schema.TypeInt,
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
			},
			"mcast_session_counting6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"napi_break_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nat46_force_ipv4_packet_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"np_queues": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"queue": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"queue": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"sport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
			"npu_group_effective_scope": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"npu_tcam": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"df": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ethertype": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_tag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"frag_off": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_buf_cnt": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_iv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_l3_flags": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_l4_flags": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_pkt_ctrl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_pri": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_pri_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_tv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ihl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip4_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip6_fl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ipver": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd10": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd11": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd8": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd9": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"slink": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"smac_change": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"src_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"src_prio": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"src_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"svid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tcp_ack": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_cwr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_ece": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_fin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_push": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_rst": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_syn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_urg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_prio": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tgt_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tos": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tvid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vdid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"dbg_dump": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mask": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"df": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ethertype": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_tag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"frag_off": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_buf_cnt": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_iv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_l3_flags": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_l4_flags": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_pkt_ctrl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_pri": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"gen_pri_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_tv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ihl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip4_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip6_fl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ipver": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd10": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd11": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd8": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"l4_wd9": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"slink": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"smac_change": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"src_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"src_prio": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"src_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcport": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"svid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tcp_ack": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_cwr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_ece": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_fin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_push": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_rst": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_syn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_urg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_prio": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tgt_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tos": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tvid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vdid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"mir_act": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlif": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"oid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pri_act": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": &schema.Schema{
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
						"sact": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"act": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bmproc": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"bmproc_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"df_lif": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"df_lif_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dfr": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"dfr_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dmac_skip": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"dmac_skip_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dosen": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"dosen_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"espff_proc": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"espff_proc_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"etype_pid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"etype_pid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"frag_proc": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"frag_proc_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fwd": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"fwd_lif": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"fwd_lif_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fwd_tvid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"fwd_tvid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fwd_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"icpen": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"icpen_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"igmp_mld_snp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"igmp_mld_snp_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"learn": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"learn_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"m_srh_ctrl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"m_srh_ctrl_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mac_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mss": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mss_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pleen": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"pleen_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prio_pid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"prio_pid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"promis": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"promis_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rfsh": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"rfsh_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"smac_skip": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"smac_skip_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tp_smchk_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tp_smchk": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tpe_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tpe_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vdm": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vdm_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vdom_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vdom_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"x_mode": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"x_mode_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"tact": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"act": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fmtuv4_s": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"fmtuv4_s_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fmtuv6_s": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"fmtuv6_s_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"lnkid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"lnkid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mac_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mss_t": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mss_t_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mtuv4": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mtuv4_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mtuv6": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"mtuv6_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"slif_act": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"slif_act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sublnkid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"sublnkid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgtv_act": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tgtv_act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tlif_act": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tlif_act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tpeid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"tpeid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"v6fe": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"v6fe_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vep_en_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vep_slid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"vep_slid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vep_en": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"xlt_lif": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"xlt_lif_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"xlt_vid": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"xlt_vid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"nss_threads_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pba_eim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pba_port_select_mode": &schema.Schema{
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
			"ple_non_syn_tcp_action": &schema.Schema{
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
						},
					},
				},
			},
			"port_path_option": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
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
				Computed: true,
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
			"prp_session_clear_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_icmp_by_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prp_port_in": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prp_port_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"rps_mode": &schema.Schema{
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
			"shaping_stats": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spa_port_select_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_ipsec_engines": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sse_backpressure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sse_ha_scan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gap": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_session_cnt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
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
				Computed: true,
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
			"sw_tr_hash": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"draco15": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_udp_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"switch_np_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_rst_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tunnel_over_vlink": &schema.Schema{
				Type:     schema.TypeString,
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
						},
						"fin_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"syn_sent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"syn_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tcp_idle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"time_wait": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						},
						"udp_idle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"uesp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ull_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_lookup_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_offload": &schema.Schema{
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpu(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpu resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpu(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemNpu(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSystemNpu(mkey, paradict)
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

func flattenObjectSystemNpuBackgroundSseScan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "scan"
	if _, ok := i["scan"]; ok {
		result["scan"] = flattenObjectSystemNpuBackgroundSseScanScan(i["scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_stale"
	if _, ok := i["scan-stale"]; ok {
		result["scan_stale"] = flattenObjectSystemNpuBackgroundSseScanScanStale(i["scan-stale"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_vt"
	if _, ok := i["scan-vt"]; ok {
		result["scan_vt"] = flattenObjectSystemNpuBackgroundSseScanScanVt(i["scan-vt"], d, pre_append)
	}

	pre_append = pre + ".0." + "stats_qual_access"
	if _, ok := i["stats-qual-access"]; ok {
		result["stats_qual_access"] = flattenObjectSystemNpuBackgroundSseScanStatsQualAccess(i["stats-qual-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "stats_qual_duration"
	if _, ok := i["stats-qual-duration"]; ok {
		result["stats_qual_duration"] = flattenObjectSystemNpuBackgroundSseScanStatsQualDuration(i["stats-qual-duration"], d, pre_append)
	}

	pre_append = pre + ".0." + "stats_update_interval"
	if _, ok := i["stats-update-interval"]; ok {
		result["stats_update_interval"] = flattenObjectSystemNpuBackgroundSseScanStatsUpdateInterval(i["stats-update-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_keepalive_interval"
	if _, ok := i["udp-keepalive-interval"]; ok {
		result["udp_keepalive_interval"] = flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(i["udp-keepalive-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_qual_access"
	if _, ok := i["udp-qual-access"]; ok {
		result["udp_qual_access"] = flattenObjectSystemNpuBackgroundSseScanUdpQualAccess(i["udp-qual-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_qual_duration"
	if _, ok := i["udp-qual-duration"]; ok {
		result["udp_qual_duration"] = flattenObjectSystemNpuBackgroundSseScanUdpQualDuration(i["udp-qual-duration"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuBackgroundSseScanScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanScanStale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanScanVt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsQualAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsQualDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpQualAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpQualDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuCapwapOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDedicatedLacpQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDedicatedManagementAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDedicatedManagementCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDefaultQosType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDefaultTcpRefreshDir(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDefaultUdpRefreshDir(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "npu_dos_meter_mode"
	if _, ok := i["npu-dos-meter-mode"]; ok {
		result["npu_dos_meter_mode"] = flattenObjectSystemNpuDosOptionsNpuDosMeterMode(i["npu-dos-meter-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "npu_dos_synproxy_mode"
	if _, ok := i["npu-dos-synproxy-mode"]; ok {
		result["npu_dos_synproxy_mode"] = flattenObjectSystemNpuDosOptionsNpuDosSynproxyMode(i["npu-dos-synproxy-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "npu_dos_tpe_mode"
	if _, ok := i["npu-dos-tpe-mode"]; ok {
		result["npu_dos_tpe_mode"] = flattenObjectSystemNpuDosOptionsNpuDosTpeMode(i["npu-dos-tpe-mode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuDosOptionsNpuDosMeterMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsNpuDosSynproxyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsNpuDosTpeMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDoubleLevelMcastOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDseTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuDswDtsProfileAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_limit"
		if _, ok := i["min-limit"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileMinLimit(i["min-limit"], d, pre_append)
			tmp["min_limit"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-MinLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := i["profile-id"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileProfileId(i["profile-id"], d, pre_append)
			tmp["profile_id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-ProfileId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "step"
		if _, ok := i["step"]; ok {
			v := flattenObjectSystemNpuDswDtsProfileStep(i["step"], d, pre_append)
			tmp["step"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswDtsProfile-Step")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuDswDtsProfileAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileMinLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileStep(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuDswQueueDtsProfileIport(i["iport"], d, pre_append)
			tmp["iport"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-Iport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oport"
		if _, ok := i["oport"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileOport(i["oport"], d, pre_append)
			tmp["oport"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-Oport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := i["profile-id"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileProfileId(i["profile-id"], d, pre_append)
			tmp["profile_id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-ProfileId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue_select"
		if _, ok := i["queue-select"]; ok {
			v := flattenObjectSystemNpuDswQueueDtsProfileQueueSelect(i["queue-select"], d, pre_append)
			tmp["queue_select"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-DswQueueDtsProfile-QueueSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuDswQueueDtsProfileIport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileOport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileQueueSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFastpath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "capwap_minlen_err"
	if _, ok := i["capwap-minlen-err"]; ok {
		result["capwap_minlen_err"] = flattenObjectSystemNpuFpAnomalyCapwapMinlenErr(i["capwap-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "esp_minlen_err"
	if _, ok := i["esp-minlen-err"]; ok {
		result["esp_minlen_err"] = flattenObjectSystemNpuFpAnomalyEspMinlenErr(i["esp-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "gre_csum_err"
	if _, ok := i["gre-csum-err"]; ok {
		result["gre_csum_err"] = flattenObjectSystemNpuFpAnomalyGreCsumErr(i["gre-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "gtpu_plen_err"
	if _, ok := i["gtpu-plen-err"]; ok {
		result["gtpu_plen_err"] = flattenObjectSystemNpuFpAnomalyGtpuPlenErr(i["gtpu-plen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_csum_err"
	if _, ok := i["icmp-csum-err"]; ok {
		result["icmp_csum_err"] = flattenObjectSystemNpuFpAnomalyIcmpCsumErr(i["icmp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_frag"
	if _, ok := i["icmp-frag"]; ok {
		result["icmp_frag"] = flattenObjectSystemNpuFpAnomalyIcmpFrag(i["icmp-frag"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_land"
	if _, ok := i["icmp-land"]; ok {
		result["icmp_land"] = flattenObjectSystemNpuFpAnomalyIcmpLand(i["icmp-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_minlen_err"
	if _, ok := i["icmp-minlen-err"]; ok {
		result["icmp_minlen_err"] = flattenObjectSystemNpuFpAnomalyIcmpMinlenErr(i["icmp-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_csum_err"
	if _, ok := i["ipv4-csum-err"]; ok {
		result["ipv4_csum_err"] = flattenObjectSystemNpuFpAnomalyIpv4CsumErr(i["ipv4-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_ihl_err"
	if _, ok := i["ipv4-ihl-err"]; ok {
		result["ipv4_ihl_err"] = flattenObjectSystemNpuFpAnomalyIpv4IhlErr(i["ipv4-ihl-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_land"
	if _, ok := i["ipv4-land"]; ok {
		result["ipv4_land"] = flattenObjectSystemNpuFpAnomalyIpv4Land(i["ipv4-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_len_err"
	if _, ok := i["ipv4-len-err"]; ok {
		result["ipv4_len_err"] = flattenObjectSystemNpuFpAnomalyIpv4LenErr(i["ipv4-len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_opt_err"
	if _, ok := i["ipv4-opt-err"]; ok {
		result["ipv4_opt_err"] = flattenObjectSystemNpuFpAnomalyIpv4OptErr(i["ipv4-opt-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optlsrr"
	if _, ok := i["ipv4-optlsrr"]; ok {
		result["ipv4_optlsrr"] = flattenObjectSystemNpuFpAnomalyIpv4Optlsrr(i["ipv4-optlsrr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optrr"
	if _, ok := i["ipv4-optrr"]; ok {
		result["ipv4_optrr"] = flattenObjectSystemNpuFpAnomalyIpv4Optrr(i["ipv4-optrr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optsecurity"
	if _, ok := i["ipv4-optsecurity"]; ok {
		result["ipv4_optsecurity"] = flattenObjectSystemNpuFpAnomalyIpv4Optsecurity(i["ipv4-optsecurity"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optssrr"
	if _, ok := i["ipv4-optssrr"]; ok {
		result["ipv4_optssrr"] = flattenObjectSystemNpuFpAnomalyIpv4Optssrr(i["ipv4-optssrr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_optstream"
	if _, ok := i["ipv4-optstream"]; ok {
		result["ipv4_optstream"] = flattenObjectSystemNpuFpAnomalyIpv4Optstream(i["ipv4-optstream"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_opttimestamp"
	if _, ok := i["ipv4-opttimestamp"]; ok {
		result["ipv4_opttimestamp"] = flattenObjectSystemNpuFpAnomalyIpv4Opttimestamp(i["ipv4-opttimestamp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_proto_err"
	if _, ok := i["ipv4-proto-err"]; ok {
		result["ipv4_proto_err"] = flattenObjectSystemNpuFpAnomalyIpv4ProtoErr(i["ipv4-proto-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_ttlzero_err"
	if _, ok := i["ipv4-ttlzero-err"]; ok {
		result["ipv4_ttlzero_err"] = flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErr(i["ipv4-ttlzero-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_unknopt"
	if _, ok := i["ipv4-unknopt"]; ok {
		result["ipv4_unknopt"] = flattenObjectSystemNpuFpAnomalyIpv4Unknopt(i["ipv4-unknopt"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv4_ver_err"
	if _, ok := i["ipv4-ver-err"]; ok {
		result["ipv4_ver_err"] = flattenObjectSystemNpuFpAnomalyIpv4VerErr(i["ipv4-ver-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_daddr_err"
	if _, ok := i["ipv6-daddr-err"]; ok {
		result["ipv6_daddr_err"] = flattenObjectSystemNpuFpAnomalyIpv6DaddrErr(i["ipv6-daddr-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_exthdr_len_err"
	if _, ok := i["ipv6-exthdr-len-err"]; ok {
		result["ipv6_exthdr_len_err"] = flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(i["ipv6-exthdr-len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_exthdr_order_err"
	if _, ok := i["ipv6-exthdr-order-err"]; ok {
		result["ipv6_exthdr_order_err"] = flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(i["ipv6-exthdr-order-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_ihl_err"
	if _, ok := i["ipv6-ihl-err"]; ok {
		result["ipv6_ihl_err"] = flattenObjectSystemNpuFpAnomalyIpv6IhlErr(i["ipv6-ihl-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_land"
	if _, ok := i["ipv6-land"]; ok {
		result["ipv6_land"] = flattenObjectSystemNpuFpAnomalyIpv6Land(i["ipv6-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optendpid"
	if _, ok := i["ipv6-optendpid"]; ok {
		result["ipv6_optendpid"] = flattenObjectSystemNpuFpAnomalyIpv6Optendpid(i["ipv6-optendpid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_opthomeaddr"
	if _, ok := i["ipv6-opthomeaddr"]; ok {
		result["ipv6_opthomeaddr"] = flattenObjectSystemNpuFpAnomalyIpv6Opthomeaddr(i["ipv6-opthomeaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optinvld"
	if _, ok := i["ipv6-optinvld"]; ok {
		result["ipv6_optinvld"] = flattenObjectSystemNpuFpAnomalyIpv6Optinvld(i["ipv6-optinvld"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optjumbo"
	if _, ok := i["ipv6-optjumbo"]; ok {
		result["ipv6_optjumbo"] = flattenObjectSystemNpuFpAnomalyIpv6Optjumbo(i["ipv6-optjumbo"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optnsap"
	if _, ok := i["ipv6-optnsap"]; ok {
		result["ipv6_optnsap"] = flattenObjectSystemNpuFpAnomalyIpv6Optnsap(i["ipv6-optnsap"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_optralert"
	if _, ok := i["ipv6-optralert"]; ok {
		result["ipv6_optralert"] = flattenObjectSystemNpuFpAnomalyIpv6Optralert(i["ipv6-optralert"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_opttunnel"
	if _, ok := i["ipv6-opttunnel"]; ok {
		result["ipv6_opttunnel"] = flattenObjectSystemNpuFpAnomalyIpv6Opttunnel(i["ipv6-opttunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_plen_zero"
	if _, ok := i["ipv6-plen-zero"]; ok {
		result["ipv6_plen_zero"] = flattenObjectSystemNpuFpAnomalyIpv6PlenZero(i["ipv6-plen-zero"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_proto_err"
	if _, ok := i["ipv6-proto-err"]; ok {
		result["ipv6_proto_err"] = flattenObjectSystemNpuFpAnomalyIpv6ProtoErr(i["ipv6-proto-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_saddr_err"
	if _, ok := i["ipv6-saddr-err"]; ok {
		result["ipv6_saddr_err"] = flattenObjectSystemNpuFpAnomalyIpv6SaddrErr(i["ipv6-saddr-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_unknopt"
	if _, ok := i["ipv6-unknopt"]; ok {
		result["ipv6_unknopt"] = flattenObjectSystemNpuFpAnomalyIpv6Unknopt(i["ipv6-unknopt"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipv6_ver_err"
	if _, ok := i["ipv6-ver-err"]; ok {
		result["ipv6_ver_err"] = flattenObjectSystemNpuFpAnomalyIpv6VerErr(i["ipv6-ver-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_csum_err"
	if _, ok := i["sctp-csum-err"]; ok {
		result["sctp_csum_err"] = flattenObjectSystemNpuFpAnomalySctpCsumErr(i["sctp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "nvgre_minlen_err"
	if _, ok := i["nvgre-minlen-err"]; ok {
		result["nvgre_minlen_err"] = flattenObjectSystemNpuFpAnomalyNvgreMinlenErr(i["nvgre-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_clen_err"
	if _, ok := i["sctp-clen-err"]; ok {
		result["sctp_clen_err"] = flattenObjectSystemNpuFpAnomalySctpClenErr(i["sctp-clen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_crc_err"
	if _, ok := i["sctp-crc-err"]; ok {
		result["sctp_crc_err"] = flattenObjectSystemNpuFpAnomalySctpCrcErr(i["sctp-crc-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_l4len_err"
	if _, ok := i["sctp-l4len-err"]; ok {
		result["sctp_l4len_err"] = flattenObjectSystemNpuFpAnomalySctpL4LenErr(i["sctp-l4len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_csum_err"
	if _, ok := i["tcp-csum-err"]; ok {
		result["tcp_csum_err"] = flattenObjectSystemNpuFpAnomalyTcpCsumErr(i["tcp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin_noack"
	if _, ok := i["tcp-fin-noack"]; ok {
		result["tcp_fin_noack"] = flattenObjectSystemNpuFpAnomalyTcpFinNoack(i["tcp-fin-noack"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin_only"
	if _, ok := i["tcp-fin-only"]; ok {
		result["tcp_fin_only"] = flattenObjectSystemNpuFpAnomalyTcpFinOnly(i["tcp-fin-only"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_hlen_err"
	if _, ok := i["tcp-hlen-err"]; ok {
		result["tcp_hlen_err"] = flattenObjectSystemNpuFpAnomalyTcpHlenErr(i["tcp-hlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_hlenvsl4len_err"
	if _, ok := i["tcp-hlenvsl4len-err"]; ok {
		result["tcp_hlenvsl4len_err"] = flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(i["tcp-hlenvsl4len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_land"
	if _, ok := i["tcp-land"]; ok {
		result["tcp_land"] = flattenObjectSystemNpuFpAnomalyTcpLand(i["tcp-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_no_flag"
	if _, ok := i["tcp-no-flag"]; ok {
		result["tcp_no_flag"] = flattenObjectSystemNpuFpAnomalyTcpNoFlag(i["tcp-no-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_plen_err"
	if _, ok := i["tcp-plen-err"]; ok {
		result["tcp_plen_err"] = flattenObjectSystemNpuFpAnomalyTcpPlenErr(i["tcp-plen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn_data"
	if _, ok := i["tcp-syn-data"]; ok {
		result["tcp_syn_data"] = flattenObjectSystemNpuFpAnomalyTcpSynData(i["tcp-syn-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn_fin"
	if _, ok := i["tcp-syn-fin"]; ok {
		result["tcp_syn_fin"] = flattenObjectSystemNpuFpAnomalyTcpSynFin(i["tcp-syn-fin"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_winnuke"
	if _, ok := i["tcp-winnuke"]; ok {
		result["tcp_winnuke"] = flattenObjectSystemNpuFpAnomalyTcpWinnuke(i["tcp-winnuke"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_csum_err"
	if _, ok := i["udp-csum-err"]; ok {
		result["udp_csum_err"] = flattenObjectSystemNpuFpAnomalyUdpCsumErr(i["udp-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_hlen_err"
	if _, ok := i["udp-hlen-err"]; ok {
		result["udp_hlen_err"] = flattenObjectSystemNpuFpAnomalyUdpHlenErr(i["udp-hlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_land"
	if _, ok := i["udp-land"]; ok {
		result["udp_land"] = flattenObjectSystemNpuFpAnomalyUdpLand(i["udp-land"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_len_err"
	if _, ok := i["udp-len-err"]; ok {
		result["udp_len_err"] = flattenObjectSystemNpuFpAnomalyUdpLenErr(i["udp-len-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_plen_err"
	if _, ok := i["udp-plen-err"]; ok {
		result["udp_plen_err"] = flattenObjectSystemNpuFpAnomalyUdpPlenErr(i["udp-plen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udplite_cover_err"
	if _, ok := i["udplite-cover-err"]; ok {
		result["udplite_cover_err"] = flattenObjectSystemNpuFpAnomalyUdpliteCoverErr(i["udplite-cover-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "udplite_csum_err"
	if _, ok := i["udplite-csum-err"]; ok {
		result["udplite_csum_err"] = flattenObjectSystemNpuFpAnomalyUdpliteCsumErr(i["udplite-csum-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "uesp_minlen_err"
	if _, ok := i["uesp-minlen-err"]; ok {
		result["uesp_minlen_err"] = flattenObjectSystemNpuFpAnomalyUespMinlenErr(i["uesp-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknproto_minlen_err"
	if _, ok := i["unknproto-minlen-err"]; ok {
		result["unknproto_minlen_err"] = flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErr(i["unknproto-minlen-err"], d, pre_append)
	}

	pre_append = pre + ".0." + "vxlan_minlen_err"
	if _, ok := i["vxlan-minlen-err"]; ok {
		result["vxlan_minlen_err"] = flattenObjectSystemNpuFpAnomalyVxlanMinlenErr(i["vxlan-minlen-err"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuFpAnomalyCapwapMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyEspMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGreCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyGtpuPlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpFrag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpLand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIcmpMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4CsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4IhlErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Land(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4LenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4OptErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optlsrr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optrr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optsecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optssrr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Optstream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Opttimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4ProtoErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4TtlzeroErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4Unknopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv4VerErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6DaddrErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6IhlErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Land(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optendpid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Opthomeaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optinvld(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optjumbo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optnsap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Optralert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Opttunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6PlenZero(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6ProtoErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6SaddrErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6Unknopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyIpv6VerErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyNvgreMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpClenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpCrcErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalySctpL4LenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinNoack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpFinOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpLand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpNoFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpPlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpSynFin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyTcpWinnuke(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpHlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpLenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpPlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCoverErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUdpliteCsumErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUespMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyUnknprotoMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuFpAnomalyVxlanMinlenErr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuGtpEnhancedCpuRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuGtpEnhancedMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuGtpSupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHashConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHashIpv6Sel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHashTblSpread(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHostShortcutMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpe(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all_protocol"
	if _, ok := i["all-protocol"]; ok {
		result["all_protocol"] = flattenObjectSystemNpuHpeAllProtocol(i["all-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "arp_max"
	if _, ok := i["arp-max"]; ok {
		result["arp_max"] = flattenObjectSystemNpuHpeArpMax(i["arp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "enable_queue_shaper"
	if _, ok := i["enable-queue-shaper"]; ok {
		result["enable_queue_shaper"] = flattenObjectSystemNpuHpeEnableQueueShaper(i["enable-queue-shaper"], d, pre_append)
	}

	pre_append = pre + ".0." + "enable_shaper"
	if _, ok := i["enable-shaper"]; ok {
		result["enable_shaper"] = flattenObjectSystemNpuHpeEnableShaper(i["enable-shaper"], d, pre_append)
	}

	pre_append = pre + ".0." + "esp_max"
	if _, ok := i["esp-max"]; ok {
		result["esp_max"] = flattenObjectSystemNpuHpeEspMax(i["esp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "exception_code"
	if _, ok := i["exception-code"]; ok {
		result["exception_code"] = flattenObjectSystemNpuHpeExceptionCode(i["exception-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "fragment_with_sess"
	if _, ok := i["fragment-with-sess"]; ok {
		result["fragment_with_sess"] = flattenObjectSystemNpuHpeFragmentWithSess(i["fragment-with-sess"], d, pre_append)
	}

	pre_append = pre + ".0." + "fragment_without_session"
	if _, ok := i["fragment-without-session"]; ok {
		result["fragment_without_session"] = flattenObjectSystemNpuHpeFragmentWithoutSession(i["fragment-without-session"], d, pre_append)
	}

	pre_append = pre + ".0." + "high_priority"
	if _, ok := i["high-priority"]; ok {
		result["high_priority"] = flattenObjectSystemNpuHpeHighPriority(i["high-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_max"
	if _, ok := i["icmp-max"]; ok {
		result["icmp_max"] = flattenObjectSystemNpuHpeIcmpMax(i["icmp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_frag_max"
	if _, ok := i["ip-frag-max"]; ok {
		result["ip_frag_max"] = flattenObjectSystemNpuHpeIpFragMax(i["ip-frag-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_others_max"
	if _, ok := i["ip-others-max"]; ok {
		result["ip_others_max"] = flattenObjectSystemNpuHpeIpOthersMax(i["ip-others-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "l2_others_max"
	if _, ok := i["l2-others-max"]; ok {
		result["l2_others_max"] = flattenObjectSystemNpuHpeL2OthersMax(i["l2-others-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "pri_type_max"
	if _, ok := i["pri-type-max"]; ok {
		result["pri_type_max"] = flattenObjectSystemNpuHpePriTypeMax(i["pri-type-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "queue_shaper_max"
	if _, ok := i["queue-shaper-max"]; ok {
		result["queue_shaper_max"] = flattenObjectSystemNpuHpeQueueShaperMax(i["queue-shaper-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "sctp_max"
	if _, ok := i["sctp-max"]; ok {
		result["sctp_max"] = flattenObjectSystemNpuHpeSctpMax(i["sctp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_max"
	if _, ok := i["tcp-max"]; ok {
		result["tcp_max"] = flattenObjectSystemNpuHpeTcpMax(i["tcp-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcpfin_rst_max"
	if _, ok := i["tcpfin-rst-max"]; ok {
		result["tcpfin_rst_max"] = flattenObjectSystemNpuHpeTcpfinRstMax(i["tcpfin-rst-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcpsyn_ack_max"
	if _, ok := i["tcpsyn-ack-max"]; ok {
		result["tcpsyn_ack_max"] = flattenObjectSystemNpuHpeTcpsynAckMax(i["tcpsyn-ack-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcpsyn_max"
	if _, ok := i["tcpsyn-max"]; ok {
		result["tcpsyn_max"] = flattenObjectSystemNpuHpeTcpsynMax(i["tcpsyn-max"], d, pre_append)
	}

	pre_append = pre + ".0." + "udp_max"
	if _, ok := i["udp-max"]; ok {
		result["udp_max"] = flattenObjectSystemNpuHpeUdpMax(i["udp-max"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuHpeAllProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeArpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEnableQueueShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEnableShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeEspMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeExceptionCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeFragmentWithSess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeFragmentWithoutSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeHighPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIcmpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIpFragMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeIpOthersMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeL2OthersMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpePriTypeMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeQueueShaperMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeSctpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpfinRstMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpsynAckMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeTcpsynMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHpeUdpMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtabDediQueueNr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtabMsgQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtxGtseQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHtxIcmpCsumChk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuHwHaScanInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpErrorRateCtrl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "icmpv4_error_bucket_size"
	if _, ok := i["icmpv4-error-bucket-size"]; ok {
		result["icmpv4_error_bucket_size"] = flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorBucketSize(i["icmpv4-error-bucket-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmpv4_error_rate"
	if _, ok := i["icmpv4-error-rate"]; ok {
		result["icmpv4_error_rate"] = flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRate(i["icmpv4-error-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmpv4_error_rate_limit"
	if _, ok := i["icmpv4-error-rate-limit"]; ok {
		result["icmpv4_error_rate_limit"] = flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRateLimit(i["icmpv4-error-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmpv6_error_bucket_size"
	if _, ok := i["icmpv6-error-bucket-size"]; ok {
		result["icmpv6_error_bucket_size"] = flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorBucketSize(i["icmpv6-error-bucket-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmpv6_error_rate"
	if _, ok := i["icmpv6-error-rate"]; ok {
		result["icmpv6_error_rate"] = flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRate(i["icmpv6-error-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmpv6_error_rate_limit"
	if _, ok := i["icmpv6-error-rate-limit"]; ok {
		result["icmpv6_error_rate_limit"] = flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRateLimit(i["icmpv6-error-rate-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorBucketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorBucketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "icmp_v4_bucket_size"
	if _, ok := i["icmp-v4-bucket-size"]; ok {
		result["icmp_v4_bucket_size"] = flattenObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize(i["icmp-v4-bucket-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_v4_rate"
	if _, ok := i["icmp-v4-rate"]; ok {
		result["icmp_v4_rate"] = flattenObjectSystemNpuIcmpRateCtrlIcmpV4Rate(i["icmp-v4-rate"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_v6_bucket_size"
	if _, ok := i["icmp-v6-bucket-size"]; ok {
		result["icmp_v6_bucket_size"] = flattenObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize(i["icmp-v6-bucket-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp_v6_rate"
	if _, ok := i["icmp-v6-rate"]; ok {
		result["icmp_v6_rate"] = flattenObjectSystemNpuIcmpRateCtrlIcmpV6Rate(i["icmp-v6-rate"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV4Rate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV6Rate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuInboundDscpCopy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuInboundDscpCopyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemNpuIpReassembly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_timeout"
	if _, ok := i["max-timeout"]; ok {
		result["max_timeout"] = flattenObjectSystemNpuIpReassemblyMaxTimeout(i["max-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_timeout"
	if _, ok := i["min-timeout"]; ok {
		result["min_timeout"] = flattenObjectSystemNpuIpReassemblyMinTimeout(i["min-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectSystemNpuIpReassemblyStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuIpReassemblyMaxTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpReassemblyMinTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpReassemblyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIntfShapingOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpFragmentOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIphRsvdReCksum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIppoolOverloadHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIppoolOverloadLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecStsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecDecSubengineMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecEncSubengineMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecHostDfclr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecInboundCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecLocalUespPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecMtuOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecObNpSel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecOverVlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpv4SessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpv4SessionQuotaHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpv4SessionQuotaLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpv6PrefixSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpv6PrefixSessionQuotaHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpv6PrefixSessionQuotaLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpsecThroughputMsgFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIptStsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIptThroughputMsgFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIsfNpQueues(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cos0"
	if _, ok := i["cos0"]; ok {
		result["cos0"] = flattenObjectSystemNpuIsfNpQueuesCos0(i["cos0"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos1"
	if _, ok := i["cos1"]; ok {
		result["cos1"] = flattenObjectSystemNpuIsfNpQueuesCos1(i["cos1"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos2"
	if _, ok := i["cos2"]; ok {
		result["cos2"] = flattenObjectSystemNpuIsfNpQueuesCos2(i["cos2"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos3"
	if _, ok := i["cos3"]; ok {
		result["cos3"] = flattenObjectSystemNpuIsfNpQueuesCos3(i["cos3"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos4"
	if _, ok := i["cos4"]; ok {
		result["cos4"] = flattenObjectSystemNpuIsfNpQueuesCos4(i["cos4"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos5"
	if _, ok := i["cos5"]; ok {
		result["cos5"] = flattenObjectSystemNpuIsfNpQueuesCos5(i["cos5"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos6"
	if _, ok := i["cos6"]; ok {
		result["cos6"] = flattenObjectSystemNpuIsfNpQueuesCos6(i["cos6"], d, pre_append)
	}

	pre_append = pre + ".0." + "cos7"
	if _, ok := i["cos7"]; ok {
		result["cos7"] = flattenObjectSystemNpuIsfNpQueuesCos7(i["cos7"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuIsfNpQueuesCos0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpQueuesCos7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSystemNpuIsfNpRxTrDistr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuLagOutPortSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMaxReceiveUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMaxSessionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMcastSessionAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMcastSessionCounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuMcastSessionCounting6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNapiBreakInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNat46ForceIpv4PacketForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueues(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ethernet_type"
	if _, ok := i["ethernet-type"]; ok {
		result["ethernet_type"] = flattenObjectSystemNpuNpQueuesEthernetType(i["ethernet-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := i["ip-protocol"]; ok {
		result["ip_protocol"] = flattenObjectSystemNpuNpQueuesIpProtocol(i["ip-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_service"
	if _, ok := i["ip-service"]; ok {
		result["ip_service"] = flattenObjectSystemNpuNpQueuesIpService(i["ip-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenObjectSystemNpuNpQueuesProfile(i["profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "scheduler"
	if _, ok := i["scheduler"]; ok {
		result["scheduler"] = flattenObjectSystemNpuNpQueuesScheduler(i["scheduler"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpQueuesEthernetType(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesEthernetTypeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeQueue(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesEthernetTypeWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-EthernetType-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesEthernetTypeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesEthernetTypeWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocol(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesIpProtocolName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolQueue(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpProtocolWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpProtocol-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpProtocolName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpProtocolWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesIpServiceDport(i["dport"], d, pre_append)
			tmp["dport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Dport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceQueue(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := i["sport"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceSport(i["sport"], d, pre_append)
			tmp["sport"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Sport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesIpServiceWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-IpService-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesIpServiceDport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesIpServiceWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesProfileCos0(i["cos0"], d, pre_append)
			tmp["cos0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := i["cos1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos1(i["cos1"], d, pre_append)
			tmp["cos1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := i["cos2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos2(i["cos2"], d, pre_append)
			tmp["cos2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := i["cos3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos3(i["cos3"], d, pre_append)
			tmp["cos3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := i["cos4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos4(i["cos4"], d, pre_append)
			tmp["cos4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := i["cos5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos5(i["cos5"], d, pre_append)
			tmp["cos5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := i["cos6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos6(i["cos6"], d, pre_append)
			tmp["cos6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := i["cos7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileCos7(i["cos7"], d, pre_append)
			tmp["cos7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Cos7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := i["dscp0"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp0(i["dscp0"], d, pre_append)
			tmp["dscp0"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := i["dscp1"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp1(i["dscp1"], d, pre_append)
			tmp["dscp1"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := i["dscp10"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp10(i["dscp10"], d, pre_append)
			tmp["dscp10"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp10")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := i["dscp11"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp11(i["dscp11"], d, pre_append)
			tmp["dscp11"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp11")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := i["dscp12"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp12(i["dscp12"], d, pre_append)
			tmp["dscp12"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := i["dscp13"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp13(i["dscp13"], d, pre_append)
			tmp["dscp13"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp13")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := i["dscp14"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp14(i["dscp14"], d, pre_append)
			tmp["dscp14"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp14")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := i["dscp15"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp15(i["dscp15"], d, pre_append)
			tmp["dscp15"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp15")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := i["dscp16"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp16(i["dscp16"], d, pre_append)
			tmp["dscp16"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp16")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := i["dscp17"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp17(i["dscp17"], d, pre_append)
			tmp["dscp17"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp17")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := i["dscp18"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp18(i["dscp18"], d, pre_append)
			tmp["dscp18"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp18")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := i["dscp19"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp19(i["dscp19"], d, pre_append)
			tmp["dscp19"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp19")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := i["dscp2"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp2(i["dscp2"], d, pre_append)
			tmp["dscp2"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := i["dscp20"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp20(i["dscp20"], d, pre_append)
			tmp["dscp20"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp20")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := i["dscp21"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp21(i["dscp21"], d, pre_append)
			tmp["dscp21"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp21")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := i["dscp22"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp22(i["dscp22"], d, pre_append)
			tmp["dscp22"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp22")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := i["dscp23"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp23(i["dscp23"], d, pre_append)
			tmp["dscp23"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp23")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := i["dscp24"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp24(i["dscp24"], d, pre_append)
			tmp["dscp24"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp24")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := i["dscp25"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp25(i["dscp25"], d, pre_append)
			tmp["dscp25"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp25")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := i["dscp26"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp26(i["dscp26"], d, pre_append)
			tmp["dscp26"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp26")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := i["dscp27"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp27(i["dscp27"], d, pre_append)
			tmp["dscp27"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp27")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := i["dscp28"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp28(i["dscp28"], d, pre_append)
			tmp["dscp28"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp28")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := i["dscp29"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp29(i["dscp29"], d, pre_append)
			tmp["dscp29"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp29")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := i["dscp3"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp3(i["dscp3"], d, pre_append)
			tmp["dscp3"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := i["dscp30"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp30(i["dscp30"], d, pre_append)
			tmp["dscp30"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp30")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := i["dscp31"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp31(i["dscp31"], d, pre_append)
			tmp["dscp31"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp31")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := i["dscp32"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp32(i["dscp32"], d, pre_append)
			tmp["dscp32"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp32")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := i["dscp33"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp33(i["dscp33"], d, pre_append)
			tmp["dscp33"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp33")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := i["dscp34"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp34(i["dscp34"], d, pre_append)
			tmp["dscp34"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := i["dscp35"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp35(i["dscp35"], d, pre_append)
			tmp["dscp35"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp35")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := i["dscp36"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp36(i["dscp36"], d, pre_append)
			tmp["dscp36"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp36")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := i["dscp37"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp37(i["dscp37"], d, pre_append)
			tmp["dscp37"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp37")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := i["dscp38"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp38(i["dscp38"], d, pre_append)
			tmp["dscp38"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp38")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := i["dscp39"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp39(i["dscp39"], d, pre_append)
			tmp["dscp39"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp39")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := i["dscp4"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp4(i["dscp4"], d, pre_append)
			tmp["dscp4"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := i["dscp40"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp40(i["dscp40"], d, pre_append)
			tmp["dscp40"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp40")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := i["dscp41"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp41(i["dscp41"], d, pre_append)
			tmp["dscp41"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp41")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := i["dscp42"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp42(i["dscp42"], d, pre_append)
			tmp["dscp42"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp42")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := i["dscp43"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp43(i["dscp43"], d, pre_append)
			tmp["dscp43"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp43")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := i["dscp44"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp44(i["dscp44"], d, pre_append)
			tmp["dscp44"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := i["dscp45"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp45(i["dscp45"], d, pre_append)
			tmp["dscp45"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp45")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := i["dscp46"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp46(i["dscp46"], d, pre_append)
			tmp["dscp46"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := i["dscp47"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp47(i["dscp47"], d, pre_append)
			tmp["dscp47"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp47")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := i["dscp48"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp48(i["dscp48"], d, pre_append)
			tmp["dscp48"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp48")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := i["dscp49"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp49(i["dscp49"], d, pre_append)
			tmp["dscp49"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp49")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := i["dscp5"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp5(i["dscp5"], d, pre_append)
			tmp["dscp5"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := i["dscp50"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp50(i["dscp50"], d, pre_append)
			tmp["dscp50"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp50")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := i["dscp51"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp51(i["dscp51"], d, pre_append)
			tmp["dscp51"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp51")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := i["dscp52"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp52(i["dscp52"], d, pre_append)
			tmp["dscp52"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp52")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := i["dscp53"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp53(i["dscp53"], d, pre_append)
			tmp["dscp53"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp53")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := i["dscp54"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp54(i["dscp54"], d, pre_append)
			tmp["dscp54"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp54")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := i["dscp55"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp55(i["dscp55"], d, pre_append)
			tmp["dscp55"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp55")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := i["dscp56"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp56(i["dscp56"], d, pre_append)
			tmp["dscp56"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp56")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := i["dscp57"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp57(i["dscp57"], d, pre_append)
			tmp["dscp57"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp57")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := i["dscp58"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp58(i["dscp58"], d, pre_append)
			tmp["dscp58"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp58")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := i["dscp59"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp59(i["dscp59"], d, pre_append)
			tmp["dscp59"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp59")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := i["dscp6"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp6(i["dscp6"], d, pre_append)
			tmp["dscp6"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := i["dscp60"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp60(i["dscp60"], d, pre_append)
			tmp["dscp60"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp60")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := i["dscp61"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp61(i["dscp61"], d, pre_append)
			tmp["dscp61"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp61")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := i["dscp62"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp62(i["dscp62"], d, pre_append)
			tmp["dscp62"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp62")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := i["dscp63"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp63(i["dscp63"], d, pre_append)
			tmp["dscp63"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp63")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := i["dscp7"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp7(i["dscp7"], d, pre_append)
			tmp["dscp7"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := i["dscp8"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp8(i["dscp8"], d, pre_append)
			tmp["dscp8"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp8")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := i["dscp9"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileDscp9(i["dscp9"], d, pre_append)
			tmp["dscp9"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Dscp9")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSystemNpuNpQueuesProfileWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Profile-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesProfileCos0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp11(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp13(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp14(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp15(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp16(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp17(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp18(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp19(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp20(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp21(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp22(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp23(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp24(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp25(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp26(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp27(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp28(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp29(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp30(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp31(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp32(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp33(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp35(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp36(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp37(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp38(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp39(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp40(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp41(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp42(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp43(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp44(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp45(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp47(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp48(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp49(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp50(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp51(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp52(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp53(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp54(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp55(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp56(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp57(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp58(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp59(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp60(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesScheduler(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuNpQueuesSchedulerMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpQueuesSchedulerName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpuNpQueues-Scheduler-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpQueuesSchedulerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesSchedulerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNp6CpsOptimizationMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuGroupEffectiveScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := i["data"]; ok {
			v := flattenObjectSystemNpuNpuTcamData(i["data"], d, pre_append)
			tmp["data"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Data")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dbg_dump"
		if _, ok := i["dbg-dump"]; ok {
			v := flattenObjectSystemNpuNpuTcamDbgDump(i["dbg-dump"], d, pre_append)
			tmp["dbg_dump"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-DbgDump")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mask"
		if _, ok := i["mask"]; ok {
			v := flattenObjectSystemNpuNpuTcamMask(i["mask"], d, pre_append)
			tmp["mask"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Mask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mir_act"
		if _, ok := i["mir-act"]; ok {
			v := flattenObjectSystemNpuNpuTcamMirAct(i["mir-act"], d, pre_append)
			tmp["mir_act"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-MirAct")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemNpuNpuTcamName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oid"
		if _, ok := i["oid"]; ok {
			v := flattenObjectSystemNpuNpuTcamOid(i["oid"], d, pre_append)
			tmp["oid"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Oid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pri_act"
		if _, ok := i["pri-act"]; ok {
			v := flattenObjectSystemNpuNpuTcamPriAct(i["pri-act"], d, pre_append)
			tmp["pri_act"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-PriAct")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sact"
		if _, ok := i["sact"]; ok {
			v := flattenObjectSystemNpuNpuTcamSact(i["sact"], d, pre_append)
			tmp["sact"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Sact")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tact"
		if _, ok := i["tact"]; ok {
			v := flattenObjectSystemNpuNpuTcamTact(i["tact"], d, pre_append)
			tmp["tact"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Tact")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemNpuNpuTcamType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vid"
		if _, ok := i["vid"]; ok {
			v := flattenObjectSystemNpuNpuTcamVid(i["vid"], d, pre_append)
			tmp["vid"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-NpuTcam-Vid")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuNpuTcamData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := i["df"]; ok {
		result["df"] = flattenObjectSystemNpuNpuTcamDataDf(i["df"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstip"
	if _, ok := i["dstip"]; ok {
		result["dstip"] = flattenObjectSystemNpuNpuTcamDataDstip(i["dstip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstipv6"
	if _, ok := i["dstipv6"]; ok {
		result["dstipv6"] = flattenObjectSystemNpuNpuTcamDataDstipv6(i["dstipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstmac"
	if _, ok := i["dstmac"]; ok {
		result["dstmac"] = flattenObjectSystemNpuNpuTcamDataDstmac(i["dstmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstport"
	if _, ok := i["dstport"]; ok {
		result["dstport"] = flattenObjectSystemNpuNpuTcamDataDstport(i["dstport"], d, pre_append)
	}

	pre_append = pre + ".0." + "ethertype"
	if _, ok := i["ethertype"]; ok {
		result["ethertype"] = flattenObjectSystemNpuNpuTcamDataEthertype(i["ethertype"], d, pre_append)
	}

	pre_append = pre + ".0." + "ext_tag"
	if _, ok := i["ext-tag"]; ok {
		result["ext_tag"] = flattenObjectSystemNpuNpuTcamDataExtTag(i["ext-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_off"
	if _, ok := i["frag-off"]; ok {
		result["frag_off"] = flattenObjectSystemNpuNpuTcamDataFragOff(i["frag-off"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := i["gen-buf-cnt"]; ok {
		result["gen_buf_cnt"] = flattenObjectSystemNpuNpuTcamDataGenBufCnt(i["gen-buf-cnt"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_iv"
	if _, ok := i["gen-iv"]; ok {
		result["gen_iv"] = flattenObjectSystemNpuNpuTcamDataGenIv(i["gen-iv"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := i["gen-l3-flags"]; ok {
		result["gen_l3_flags"] = flattenObjectSystemNpuNpuTcamDataGenL3Flags(i["gen-l3-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := i["gen-l4-flags"]; ok {
		result["gen_l4_flags"] = flattenObjectSystemNpuNpuTcamDataGenL4Flags(i["gen-l4-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := i["gen-pkt-ctrl"]; ok {
		result["gen_pkt_ctrl"] = flattenObjectSystemNpuNpuTcamDataGenPktCtrl(i["gen-pkt-ctrl"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri"
	if _, ok := i["gen-pri"]; ok {
		result["gen_pri"] = flattenObjectSystemNpuNpuTcamDataGenPri(i["gen-pri"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := i["gen-pri-v"]; ok {
		result["gen_pri_v"] = flattenObjectSystemNpuNpuTcamDataGenPriV(i["gen-pri-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_tv"
	if _, ok := i["gen-tv"]; ok {
		result["gen_tv"] = flattenObjectSystemNpuNpuTcamDataGenTv(i["gen-tv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ihl"
	if _, ok := i["ihl"]; ok {
		result["ihl"] = flattenObjectSystemNpuNpuTcamDataIhl(i["ihl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip4_id"
	if _, ok := i["ip4-id"]; ok {
		result["ip4_id"] = flattenObjectSystemNpuNpuTcamDataIp4Id(i["ip4-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := i["ip6-fl"]; ok {
		result["ip6_fl"] = flattenObjectSystemNpuNpuTcamDataIp6Fl(i["ip6-fl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipver"
	if _, ok := i["ipver"]; ok {
		result["ipver"] = flattenObjectSystemNpuNpuTcamDataIpver(i["ipver"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := i["l4-wd10"]; ok {
		result["l4_wd10"] = flattenObjectSystemNpuNpuTcamDataL4Wd10(i["l4-wd10"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := i["l4-wd11"]; ok {
		result["l4_wd11"] = flattenObjectSystemNpuNpuTcamDataL4Wd11(i["l4-wd11"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := i["l4-wd8"]; ok {
		result["l4_wd8"] = flattenObjectSystemNpuNpuTcamDataL4Wd8(i["l4-wd8"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := i["l4-wd9"]; ok {
		result["l4_wd9"] = flattenObjectSystemNpuNpuTcamDataL4Wd9(i["l4-wd9"], d, pre_append)
	}

	pre_append = pre + ".0." + "mf"
	if _, ok := i["mf"]; ok {
		result["mf"] = flattenObjectSystemNpuNpuTcamDataMf(i["mf"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectSystemNpuNpuTcamDataProtocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "slink"
	if _, ok := i["slink"]; ok {
		result["slink"] = flattenObjectSystemNpuNpuTcamDataSlink(i["slink"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_change"
	if _, ok := i["smac-change"]; ok {
		result["smac_change"] = flattenObjectSystemNpuNpuTcamDataSmacChange(i["smac-change"], d, pre_append)
	}

	pre_append = pre + ".0." + "sp"
	if _, ok := i["sp"]; ok {
		result["sp"] = flattenObjectSystemNpuNpuTcamDataSp(i["sp"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_cfi"
	if _, ok := i["src-cfi"]; ok {
		result["src_cfi"] = flattenObjectSystemNpuNpuTcamDataSrcCfi(i["src-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_prio"
	if _, ok := i["src-prio"]; ok {
		result["src_prio"] = flattenObjectSystemNpuNpuTcamDataSrcPrio(i["src-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_updt"
	if _, ok := i["src-updt"]; ok {
		result["src_updt"] = flattenObjectSystemNpuNpuTcamDataSrcUpdt(i["src-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcip"
	if _, ok := i["srcip"]; ok {
		result["srcip"] = flattenObjectSystemNpuNpuTcamDataSrcip(i["srcip"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcipv6"
	if _, ok := i["srcipv6"]; ok {
		result["srcipv6"] = flattenObjectSystemNpuNpuTcamDataSrcipv6(i["srcipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcmac"
	if _, ok := i["srcmac"]; ok {
		result["srcmac"] = flattenObjectSystemNpuNpuTcamDataSrcmac(i["srcmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcport"
	if _, ok := i["srcport"]; ok {
		result["srcport"] = flattenObjectSystemNpuNpuTcamDataSrcport(i["srcport"], d, pre_append)
	}

	pre_append = pre + ".0." + "svid"
	if _, ok := i["svid"]; ok {
		result["svid"] = flattenObjectSystemNpuNpuTcamDataSvid(i["svid"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := i["tcp-ack"]; ok {
		result["tcp_ack"] = flattenObjectSystemNpuNpuTcamDataTcpAck(i["tcp-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := i["tcp-cwr"]; ok {
		result["tcp_cwr"] = flattenObjectSystemNpuNpuTcamDataTcpCwr(i["tcp-cwr"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := i["tcp-ece"]; ok {
		result["tcp_ece"] = flattenObjectSystemNpuNpuTcamDataTcpEce(i["tcp-ece"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := i["tcp-fin"]; ok {
		result["tcp_fin"] = flattenObjectSystemNpuNpuTcamDataTcpFin(i["tcp-fin"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_push"
	if _, ok := i["tcp-push"]; ok {
		result["tcp_push"] = flattenObjectSystemNpuNpuTcamDataTcpPush(i["tcp-push"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := i["tcp-rst"]; ok {
		result["tcp_rst"] = flattenObjectSystemNpuNpuTcamDataTcpRst(i["tcp-rst"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := i["tcp-syn"]; ok {
		result["tcp_syn"] = flattenObjectSystemNpuNpuTcamDataTcpSyn(i["tcp-syn"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := i["tcp-urg"]; ok {
		result["tcp_urg"] = flattenObjectSystemNpuNpuTcamDataTcpUrg(i["tcp-urg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := i["tgt-cfi"]; ok {
		result["tgt_cfi"] = flattenObjectSystemNpuNpuTcamDataTgtCfi(i["tgt-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := i["tgt-prio"]; ok {
		result["tgt_prio"] = flattenObjectSystemNpuNpuTcamDataTgtPrio(i["tgt-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := i["tgt-updt"]; ok {
		result["tgt_updt"] = flattenObjectSystemNpuNpuTcamDataTgtUpdt(i["tgt-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_v"
	if _, ok := i["tgt-v"]; ok {
		result["tgt_v"] = flattenObjectSystemNpuNpuTcamDataTgtV(i["tgt-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tos"
	if _, ok := i["tos"]; ok {
		result["tos"] = flattenObjectSystemNpuNpuTcamDataTos(i["tos"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp"
	if _, ok := i["tp"]; ok {
		result["tp"] = flattenObjectSystemNpuNpuTcamDataTp(i["tp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ttl"
	if _, ok := i["ttl"]; ok {
		result["ttl"] = flattenObjectSystemNpuNpuTcamDataTtl(i["ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "tvid"
	if _, ok := i["tvid"]; ok {
		result["tvid"] = flattenObjectSystemNpuNpuTcamDataTvid(i["tvid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdid"
	if _, ok := i["vdid"]; ok {
		result["vdid"] = flattenObjectSystemNpuNpuTcamDataVdid(i["vdid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamDataDf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataEthertype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataExtTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataFragOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenBufCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenIv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenL3Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenL4Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPktCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPriV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenTv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIhl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIp4Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIp6Fl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIpver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd11(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataMf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSmacChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcCfi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcPrio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcUpdt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSvid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpAck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpCwr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpEce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpFin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpPush(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpUrg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtCfi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtPrio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtUpdt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTvid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataVdid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDbgDump(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMask(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := i["df"]; ok {
		result["df"] = flattenObjectSystemNpuNpuTcamMaskDf(i["df"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstip"
	if _, ok := i["dstip"]; ok {
		result["dstip"] = flattenObjectSystemNpuNpuTcamMaskDstip(i["dstip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstipv6"
	if _, ok := i["dstipv6"]; ok {
		result["dstipv6"] = flattenObjectSystemNpuNpuTcamMaskDstipv6(i["dstipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstmac"
	if _, ok := i["dstmac"]; ok {
		result["dstmac"] = flattenObjectSystemNpuNpuTcamMaskDstmac(i["dstmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstport"
	if _, ok := i["dstport"]; ok {
		result["dstport"] = flattenObjectSystemNpuNpuTcamMaskDstport(i["dstport"], d, pre_append)
	}

	pre_append = pre + ".0." + "ethertype"
	if _, ok := i["ethertype"]; ok {
		result["ethertype"] = flattenObjectSystemNpuNpuTcamMaskEthertype(i["ethertype"], d, pre_append)
	}

	pre_append = pre + ".0." + "ext_tag"
	if _, ok := i["ext-tag"]; ok {
		result["ext_tag"] = flattenObjectSystemNpuNpuTcamMaskExtTag(i["ext-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_off"
	if _, ok := i["frag-off"]; ok {
		result["frag_off"] = flattenObjectSystemNpuNpuTcamMaskFragOff(i["frag-off"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := i["gen-buf-cnt"]; ok {
		result["gen_buf_cnt"] = flattenObjectSystemNpuNpuTcamMaskGenBufCnt(i["gen-buf-cnt"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_iv"
	if _, ok := i["gen-iv"]; ok {
		result["gen_iv"] = flattenObjectSystemNpuNpuTcamMaskGenIv(i["gen-iv"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := i["gen-l3-flags"]; ok {
		result["gen_l3_flags"] = flattenObjectSystemNpuNpuTcamMaskGenL3Flags(i["gen-l3-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := i["gen-l4-flags"]; ok {
		result["gen_l4_flags"] = flattenObjectSystemNpuNpuTcamMaskGenL4Flags(i["gen-l4-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := i["gen-pkt-ctrl"]; ok {
		result["gen_pkt_ctrl"] = flattenObjectSystemNpuNpuTcamMaskGenPktCtrl(i["gen-pkt-ctrl"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri"
	if _, ok := i["gen-pri"]; ok {
		result["gen_pri"] = flattenObjectSystemNpuNpuTcamMaskGenPri(i["gen-pri"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := i["gen-pri-v"]; ok {
		result["gen_pri_v"] = flattenObjectSystemNpuNpuTcamMaskGenPriV(i["gen-pri-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_tv"
	if _, ok := i["gen-tv"]; ok {
		result["gen_tv"] = flattenObjectSystemNpuNpuTcamMaskGenTv(i["gen-tv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ihl"
	if _, ok := i["ihl"]; ok {
		result["ihl"] = flattenObjectSystemNpuNpuTcamMaskIhl(i["ihl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip4_id"
	if _, ok := i["ip4-id"]; ok {
		result["ip4_id"] = flattenObjectSystemNpuNpuTcamMaskIp4Id(i["ip4-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := i["ip6-fl"]; ok {
		result["ip6_fl"] = flattenObjectSystemNpuNpuTcamMaskIp6Fl(i["ip6-fl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipver"
	if _, ok := i["ipver"]; ok {
		result["ipver"] = flattenObjectSystemNpuNpuTcamMaskIpver(i["ipver"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := i["l4-wd10"]; ok {
		result["l4_wd10"] = flattenObjectSystemNpuNpuTcamMaskL4Wd10(i["l4-wd10"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := i["l4-wd11"]; ok {
		result["l4_wd11"] = flattenObjectSystemNpuNpuTcamMaskL4Wd11(i["l4-wd11"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := i["l4-wd8"]; ok {
		result["l4_wd8"] = flattenObjectSystemNpuNpuTcamMaskL4Wd8(i["l4-wd8"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := i["l4-wd9"]; ok {
		result["l4_wd9"] = flattenObjectSystemNpuNpuTcamMaskL4Wd9(i["l4-wd9"], d, pre_append)
	}

	pre_append = pre + ".0." + "mf"
	if _, ok := i["mf"]; ok {
		result["mf"] = flattenObjectSystemNpuNpuTcamMaskMf(i["mf"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectSystemNpuNpuTcamMaskProtocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "slink"
	if _, ok := i["slink"]; ok {
		result["slink"] = flattenObjectSystemNpuNpuTcamMaskSlink(i["slink"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_change"
	if _, ok := i["smac-change"]; ok {
		result["smac_change"] = flattenObjectSystemNpuNpuTcamMaskSmacChange(i["smac-change"], d, pre_append)
	}

	pre_append = pre + ".0." + "sp"
	if _, ok := i["sp"]; ok {
		result["sp"] = flattenObjectSystemNpuNpuTcamMaskSp(i["sp"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_cfi"
	if _, ok := i["src-cfi"]; ok {
		result["src_cfi"] = flattenObjectSystemNpuNpuTcamMaskSrcCfi(i["src-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_prio"
	if _, ok := i["src-prio"]; ok {
		result["src_prio"] = flattenObjectSystemNpuNpuTcamMaskSrcPrio(i["src-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_updt"
	if _, ok := i["src-updt"]; ok {
		result["src_updt"] = flattenObjectSystemNpuNpuTcamMaskSrcUpdt(i["src-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcip"
	if _, ok := i["srcip"]; ok {
		result["srcip"] = flattenObjectSystemNpuNpuTcamMaskSrcip(i["srcip"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcipv6"
	if _, ok := i["srcipv6"]; ok {
		result["srcipv6"] = flattenObjectSystemNpuNpuTcamMaskSrcipv6(i["srcipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcmac"
	if _, ok := i["srcmac"]; ok {
		result["srcmac"] = flattenObjectSystemNpuNpuTcamMaskSrcmac(i["srcmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcport"
	if _, ok := i["srcport"]; ok {
		result["srcport"] = flattenObjectSystemNpuNpuTcamMaskSrcport(i["srcport"], d, pre_append)
	}

	pre_append = pre + ".0." + "svid"
	if _, ok := i["svid"]; ok {
		result["svid"] = flattenObjectSystemNpuNpuTcamMaskSvid(i["svid"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := i["tcp-ack"]; ok {
		result["tcp_ack"] = flattenObjectSystemNpuNpuTcamMaskTcpAck(i["tcp-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := i["tcp-cwr"]; ok {
		result["tcp_cwr"] = flattenObjectSystemNpuNpuTcamMaskTcpCwr(i["tcp-cwr"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := i["tcp-ece"]; ok {
		result["tcp_ece"] = flattenObjectSystemNpuNpuTcamMaskTcpEce(i["tcp-ece"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := i["tcp-fin"]; ok {
		result["tcp_fin"] = flattenObjectSystemNpuNpuTcamMaskTcpFin(i["tcp-fin"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_push"
	if _, ok := i["tcp-push"]; ok {
		result["tcp_push"] = flattenObjectSystemNpuNpuTcamMaskTcpPush(i["tcp-push"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := i["tcp-rst"]; ok {
		result["tcp_rst"] = flattenObjectSystemNpuNpuTcamMaskTcpRst(i["tcp-rst"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := i["tcp-syn"]; ok {
		result["tcp_syn"] = flattenObjectSystemNpuNpuTcamMaskTcpSyn(i["tcp-syn"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := i["tcp-urg"]; ok {
		result["tcp_urg"] = flattenObjectSystemNpuNpuTcamMaskTcpUrg(i["tcp-urg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := i["tgt-cfi"]; ok {
		result["tgt_cfi"] = flattenObjectSystemNpuNpuTcamMaskTgtCfi(i["tgt-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := i["tgt-prio"]; ok {
		result["tgt_prio"] = flattenObjectSystemNpuNpuTcamMaskTgtPrio(i["tgt-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := i["tgt-updt"]; ok {
		result["tgt_updt"] = flattenObjectSystemNpuNpuTcamMaskTgtUpdt(i["tgt-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_v"
	if _, ok := i["tgt-v"]; ok {
		result["tgt_v"] = flattenObjectSystemNpuNpuTcamMaskTgtV(i["tgt-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tos"
	if _, ok := i["tos"]; ok {
		result["tos"] = flattenObjectSystemNpuNpuTcamMaskTos(i["tos"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp"
	if _, ok := i["tp"]; ok {
		result["tp"] = flattenObjectSystemNpuNpuTcamMaskTp(i["tp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ttl"
	if _, ok := i["ttl"]; ok {
		result["ttl"] = flattenObjectSystemNpuNpuTcamMaskTtl(i["ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "tvid"
	if _, ok := i["tvid"]; ok {
		result["tvid"] = flattenObjectSystemNpuNpuTcamMaskTvid(i["tvid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdid"
	if _, ok := i["vdid"]; ok {
		result["vdid"] = flattenObjectSystemNpuNpuTcamMaskVdid(i["vdid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamMaskDf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskEthertype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskExtTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskFragOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenBufCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenIv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenL3Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenL4Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPktCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPriV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenTv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIhl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIp4Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIp6Fl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIpver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd11(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskMf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSmacChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcCfi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcPrio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcUpdt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSvid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpAck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpCwr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpEce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpFin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpPush(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpUrg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtCfi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtPrio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtUpdt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTvid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskVdid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMirAct(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlif"
	if _, ok := i["vlif"]; ok {
		result["vlif"] = flattenObjectSystemNpuNpuTcamMirActVlif(i["vlif"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamMirActVlif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamPriAct(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenObjectSystemNpuNpuTcamPriActPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "weight"
	if _, ok := i["weight"]; ok {
		result["weight"] = flattenObjectSystemNpuNpuTcamPriActWeight(i["weight"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamPriActPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamPriActWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSact(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := i["act"]; ok {
		result["act"] = flattenObjectSystemNpuNpuTcamSactAct(i["act"], d, pre_append)
	}

	pre_append = pre + ".0." + "act_v"
	if _, ok := i["act-v"]; ok {
		result["act_v"] = flattenObjectSystemNpuNpuTcamSactActV(i["act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "bmproc"
	if _, ok := i["bmproc"]; ok {
		result["bmproc"] = flattenObjectSystemNpuNpuTcamSactBmproc(i["bmproc"], d, pre_append)
	}

	pre_append = pre + ".0." + "bmproc_v"
	if _, ok := i["bmproc-v"]; ok {
		result["bmproc_v"] = flattenObjectSystemNpuNpuTcamSactBmprocV(i["bmproc-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "df_lif"
	if _, ok := i["df-lif"]; ok {
		result["df_lif"] = flattenObjectSystemNpuNpuTcamSactDfLif(i["df-lif"], d, pre_append)
	}

	pre_append = pre + ".0." + "df_lif_v"
	if _, ok := i["df-lif-v"]; ok {
		result["df_lif_v"] = flattenObjectSystemNpuNpuTcamSactDfLifV(i["df-lif-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "dfr"
	if _, ok := i["dfr"]; ok {
		result["dfr"] = flattenObjectSystemNpuNpuTcamSactDfr(i["dfr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dfr_v"
	if _, ok := i["dfr-v"]; ok {
		result["dfr_v"] = flattenObjectSystemNpuNpuTcamSactDfrV(i["dfr-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "dmac_skip"
	if _, ok := i["dmac-skip"]; ok {
		result["dmac_skip"] = flattenObjectSystemNpuNpuTcamSactDmacSkip(i["dmac-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dmac_skip_v"
	if _, ok := i["dmac-skip-v"]; ok {
		result["dmac_skip_v"] = flattenObjectSystemNpuNpuTcamSactDmacSkipV(i["dmac-skip-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "dosen"
	if _, ok := i["dosen"]; ok {
		result["dosen"] = flattenObjectSystemNpuNpuTcamSactDosen(i["dosen"], d, pre_append)
	}

	pre_append = pre + ".0." + "dosen_v"
	if _, ok := i["dosen-v"]; ok {
		result["dosen_v"] = flattenObjectSystemNpuNpuTcamSactDosenV(i["dosen-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "espff_proc"
	if _, ok := i["espff-proc"]; ok {
		result["espff_proc"] = flattenObjectSystemNpuNpuTcamSactEspffProc(i["espff-proc"], d, pre_append)
	}

	pre_append = pre + ".0." + "espff_proc_v"
	if _, ok := i["espff-proc-v"]; ok {
		result["espff_proc_v"] = flattenObjectSystemNpuNpuTcamSactEspffProcV(i["espff-proc-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "etype_pid"
	if _, ok := i["etype-pid"]; ok {
		result["etype_pid"] = flattenObjectSystemNpuNpuTcamSactEtypePid(i["etype-pid"], d, pre_append)
	}

	pre_append = pre + ".0." + "etype_pid_v"
	if _, ok := i["etype-pid-v"]; ok {
		result["etype_pid_v"] = flattenObjectSystemNpuNpuTcamSactEtypePidV(i["etype-pid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_proc"
	if _, ok := i["frag-proc"]; ok {
		result["frag_proc"] = flattenObjectSystemNpuNpuTcamSactFragProc(i["frag-proc"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_proc_v"
	if _, ok := i["frag-proc-v"]; ok {
		result["frag_proc_v"] = flattenObjectSystemNpuNpuTcamSactFragProcV(i["frag-proc-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd"
	if _, ok := i["fwd"]; ok {
		result["fwd"] = flattenObjectSystemNpuNpuTcamSactFwd(i["fwd"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_lif"
	if _, ok := i["fwd-lif"]; ok {
		result["fwd_lif"] = flattenObjectSystemNpuNpuTcamSactFwdLif(i["fwd-lif"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_lif_v"
	if _, ok := i["fwd-lif-v"]; ok {
		result["fwd_lif_v"] = flattenObjectSystemNpuNpuTcamSactFwdLifV(i["fwd-lif-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_tvid"
	if _, ok := i["fwd-tvid"]; ok {
		result["fwd_tvid"] = flattenObjectSystemNpuNpuTcamSactFwdTvid(i["fwd-tvid"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_tvid_v"
	if _, ok := i["fwd-tvid-v"]; ok {
		result["fwd_tvid_v"] = flattenObjectSystemNpuNpuTcamSactFwdTvidV(i["fwd-tvid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_v"
	if _, ok := i["fwd-v"]; ok {
		result["fwd_v"] = flattenObjectSystemNpuNpuTcamSactFwdV(i["fwd-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "icpen"
	if _, ok := i["icpen"]; ok {
		result["icpen"] = flattenObjectSystemNpuNpuTcamSactIcpen(i["icpen"], d, pre_append)
	}

	pre_append = pre + ".0." + "icpen_v"
	if _, ok := i["icpen-v"]; ok {
		result["icpen_v"] = flattenObjectSystemNpuNpuTcamSactIcpenV(i["icpen-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "igmp_mld_snp"
	if _, ok := i["igmp-mld-snp"]; ok {
		result["igmp_mld_snp"] = flattenObjectSystemNpuNpuTcamSactIgmpMldSnp(i["igmp-mld-snp"], d, pre_append)
	}

	pre_append = pre + ".0." + "igmp_mld_snp_v"
	if _, ok := i["igmp-mld-snp-v"]; ok {
		result["igmp_mld_snp_v"] = flattenObjectSystemNpuNpuTcamSactIgmpMldSnpV(i["igmp-mld-snp-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "learn"
	if _, ok := i["learn"]; ok {
		result["learn"] = flattenObjectSystemNpuNpuTcamSactLearn(i["learn"], d, pre_append)
	}

	pre_append = pre + ".0." + "learn_v"
	if _, ok := i["learn-v"]; ok {
		result["learn_v"] = flattenObjectSystemNpuNpuTcamSactLearnV(i["learn-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "m_srh_ctrl"
	if _, ok := i["m-srh-ctrl"]; ok {
		result["m_srh_ctrl"] = flattenObjectSystemNpuNpuTcamSactMSrhCtrl(i["m-srh-ctrl"], d, pre_append)
	}

	pre_append = pre + ".0." + "m_srh_ctrl_v"
	if _, ok := i["m-srh-ctrl-v"]; ok {
		result["m_srh_ctrl_v"] = flattenObjectSystemNpuNpuTcamSactMSrhCtrlV(i["m-srh-ctrl-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id"
	if _, ok := i["mac-id"]; ok {
		result["mac_id"] = flattenObjectSystemNpuNpuTcamSactMacId(i["mac-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := i["mac-id-v"]; ok {
		result["mac_id_v"] = flattenObjectSystemNpuNpuTcamSactMacIdV(i["mac-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss"
	if _, ok := i["mss"]; ok {
		result["mss"] = flattenObjectSystemNpuNpuTcamSactMss(i["mss"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss_v"
	if _, ok := i["mss-v"]; ok {
		result["mss_v"] = flattenObjectSystemNpuNpuTcamSactMssV(i["mss-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "pleen"
	if _, ok := i["pleen"]; ok {
		result["pleen"] = flattenObjectSystemNpuNpuTcamSactPleen(i["pleen"], d, pre_append)
	}

	pre_append = pre + ".0." + "pleen_v"
	if _, ok := i["pleen-v"]; ok {
		result["pleen_v"] = flattenObjectSystemNpuNpuTcamSactPleenV(i["pleen-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "prio_pid"
	if _, ok := i["prio-pid"]; ok {
		result["prio_pid"] = flattenObjectSystemNpuNpuTcamSactPrioPid(i["prio-pid"], d, pre_append)
	}

	pre_append = pre + ".0." + "prio_pid_v"
	if _, ok := i["prio-pid-v"]; ok {
		result["prio_pid_v"] = flattenObjectSystemNpuNpuTcamSactPrioPidV(i["prio-pid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "promis"
	if _, ok := i["promis"]; ok {
		result["promis"] = flattenObjectSystemNpuNpuTcamSactPromis(i["promis"], d, pre_append)
	}

	pre_append = pre + ".0." + "promis_v"
	if _, ok := i["promis-v"]; ok {
		result["promis_v"] = flattenObjectSystemNpuNpuTcamSactPromisV(i["promis-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "rfsh"
	if _, ok := i["rfsh"]; ok {
		result["rfsh"] = flattenObjectSystemNpuNpuTcamSactRfsh(i["rfsh"], d, pre_append)
	}

	pre_append = pre + ".0." + "rfsh_v"
	if _, ok := i["rfsh-v"]; ok {
		result["rfsh_v"] = flattenObjectSystemNpuNpuTcamSactRfshV(i["rfsh-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_skip"
	if _, ok := i["smac-skip"]; ok {
		result["smac_skip"] = flattenObjectSystemNpuNpuTcamSactSmacSkip(i["smac-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_skip_v"
	if _, ok := i["smac-skip-v"]; ok {
		result["smac_skip_v"] = flattenObjectSystemNpuNpuTcamSactSmacSkipV(i["smac-skip-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp_smchk_v"
	if _, ok := i["tp-smchk-v"]; ok {
		result["tp_smchk_v"] = flattenObjectSystemNpuNpuTcamSactTpSmchkV(i["tp-smchk-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp_smchk"
	if _, ok := i["tp_smchk"]; ok {
		result["tp_smchk"] = flattenObjectSystemNpuNpuTcamSactTpSmchk(i["tp_smchk"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpe_id"
	if _, ok := i["tpe-id"]; ok {
		result["tpe_id"] = flattenObjectSystemNpuNpuTcamSactTpeId(i["tpe-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpe_id_v"
	if _, ok := i["tpe-id-v"]; ok {
		result["tpe_id_v"] = flattenObjectSystemNpuNpuTcamSactTpeIdV(i["tpe-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdm"
	if _, ok := i["vdm"]; ok {
		result["vdm"] = flattenObjectSystemNpuNpuTcamSactVdm(i["vdm"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdm_v"
	if _, ok := i["vdm-v"]; ok {
		result["vdm_v"] = flattenObjectSystemNpuNpuTcamSactVdmV(i["vdm-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom_id"
	if _, ok := i["vdom-id"]; ok {
		result["vdom_id"] = flattenObjectSystemNpuNpuTcamSactVdomId(i["vdom-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom_id_v"
	if _, ok := i["vdom-id-v"]; ok {
		result["vdom_id_v"] = flattenObjectSystemNpuNpuTcamSactVdomIdV(i["vdom-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "x_mode"
	if _, ok := i["x-mode"]; ok {
		result["x_mode"] = flattenObjectSystemNpuNpuTcamSactXMode(i["x-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "x_mode_v"
	if _, ok := i["x-mode-v"]; ok {
		result["x_mode_v"] = flattenObjectSystemNpuNpuTcamSactXModeV(i["x-mode-v"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamSactAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactActV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactBmproc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactBmprocV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfLif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfLifV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfrV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDmacSkip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDmacSkipV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDosen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDosenV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEspffProc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEspffProcV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEtypePid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEtypePidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFragProc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFragProcV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdLif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdLifV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdTvid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdTvidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIcpen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIcpenV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIgmpMldSnp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIgmpMldSnpV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactLearn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactLearnV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMSrhCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMSrhCtrlV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMacId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMacIdV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMssV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPleen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPleenV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPrioPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPrioPidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPromis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPromisV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactRfsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactRfshV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactSmacSkip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactSmacSkipV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpSmchkV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpSmchk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpeIdV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdmV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdomId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdomIdV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactXMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactXModeV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTact(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := i["act"]; ok {
		result["act"] = flattenObjectSystemNpuNpuTcamTactAct(i["act"], d, pre_append)
	}

	pre_append = pre + ".0." + "act_v"
	if _, ok := i["act-v"]; ok {
		result["act_v"] = flattenObjectSystemNpuNpuTcamTactActV(i["act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv4_s"
	if _, ok := i["fmtuv4-s"]; ok {
		result["fmtuv4_s"] = flattenObjectSystemNpuNpuTcamTactFmtuv4S(i["fmtuv4-s"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv4_s_v"
	if _, ok := i["fmtuv4-s-v"]; ok {
		result["fmtuv4_s_v"] = flattenObjectSystemNpuNpuTcamTactFmtuv4SV(i["fmtuv4-s-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv6_s"
	if _, ok := i["fmtuv6-s"]; ok {
		result["fmtuv6_s"] = flattenObjectSystemNpuNpuTcamTactFmtuv6S(i["fmtuv6-s"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv6_s_v"
	if _, ok := i["fmtuv6-s-v"]; ok {
		result["fmtuv6_s_v"] = flattenObjectSystemNpuNpuTcamTactFmtuv6SV(i["fmtuv6-s-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "lnkid"
	if _, ok := i["lnkid"]; ok {
		result["lnkid"] = flattenObjectSystemNpuNpuTcamTactLnkid(i["lnkid"], d, pre_append)
	}

	pre_append = pre + ".0." + "lnkid_v"
	if _, ok := i["lnkid-v"]; ok {
		result["lnkid_v"] = flattenObjectSystemNpuNpuTcamTactLnkidV(i["lnkid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id"
	if _, ok := i["mac-id"]; ok {
		result["mac_id"] = flattenObjectSystemNpuNpuTcamTactMacId(i["mac-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := i["mac-id-v"]; ok {
		result["mac_id_v"] = flattenObjectSystemNpuNpuTcamTactMacIdV(i["mac-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss_t"
	if _, ok := i["mss-t"]; ok {
		result["mss_t"] = flattenObjectSystemNpuNpuTcamTactMssT(i["mss-t"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss_t_v"
	if _, ok := i["mss-t-v"]; ok {
		result["mss_t_v"] = flattenObjectSystemNpuNpuTcamTactMssTV(i["mss-t-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv4"
	if _, ok := i["mtuv4"]; ok {
		result["mtuv4"] = flattenObjectSystemNpuNpuTcamTactMtuv4(i["mtuv4"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv4_v"
	if _, ok := i["mtuv4-v"]; ok {
		result["mtuv4_v"] = flattenObjectSystemNpuNpuTcamTactMtuv4V(i["mtuv4-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv6"
	if _, ok := i["mtuv6"]; ok {
		result["mtuv6"] = flattenObjectSystemNpuNpuTcamTactMtuv6(i["mtuv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv6_v"
	if _, ok := i["mtuv6-v"]; ok {
		result["mtuv6_v"] = flattenObjectSystemNpuNpuTcamTactMtuv6V(i["mtuv6-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "slif_act"
	if _, ok := i["slif-act"]; ok {
		result["slif_act"] = flattenObjectSystemNpuNpuTcamTactSlifAct(i["slif-act"], d, pre_append)
	}

	pre_append = pre + ".0." + "slif_act_v"
	if _, ok := i["slif-act-v"]; ok {
		result["slif_act_v"] = flattenObjectSystemNpuNpuTcamTactSlifActV(i["slif-act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "sublnkid"
	if _, ok := i["sublnkid"]; ok {
		result["sublnkid"] = flattenObjectSystemNpuNpuTcamTactSublnkid(i["sublnkid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sublnkid_v"
	if _, ok := i["sublnkid-v"]; ok {
		result["sublnkid_v"] = flattenObjectSystemNpuNpuTcamTactSublnkidV(i["sublnkid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgtv_act"
	if _, ok := i["tgtv-act"]; ok {
		result["tgtv_act"] = flattenObjectSystemNpuNpuTcamTactTgtvAct(i["tgtv-act"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgtv_act_v"
	if _, ok := i["tgtv-act-v"]; ok {
		result["tgtv_act_v"] = flattenObjectSystemNpuNpuTcamTactTgtvActV(i["tgtv-act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tlif_act"
	if _, ok := i["tlif-act"]; ok {
		result["tlif_act"] = flattenObjectSystemNpuNpuTcamTactTlifAct(i["tlif-act"], d, pre_append)
	}

	pre_append = pre + ".0." + "tlif_act_v"
	if _, ok := i["tlif-act-v"]; ok {
		result["tlif_act_v"] = flattenObjectSystemNpuNpuTcamTactTlifActV(i["tlif-act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpeid"
	if _, ok := i["tpeid"]; ok {
		result["tpeid"] = flattenObjectSystemNpuNpuTcamTactTpeid(i["tpeid"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpeid_v"
	if _, ok := i["tpeid-v"]; ok {
		result["tpeid_v"] = flattenObjectSystemNpuNpuTcamTactTpeidV(i["tpeid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "v6fe"
	if _, ok := i["v6fe"]; ok {
		result["v6fe"] = flattenObjectSystemNpuNpuTcamTactV6Fe(i["v6fe"], d, pre_append)
	}

	pre_append = pre + ".0." + "v6fe_v"
	if _, ok := i["v6fe-v"]; ok {
		result["v6fe_v"] = flattenObjectSystemNpuNpuTcamTactV6FeV(i["v6fe-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_en_v"
	if _, ok := i["vep-en-v"]; ok {
		result["vep_en_v"] = flattenObjectSystemNpuNpuTcamTactVepEnV(i["vep-en-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_slid"
	if _, ok := i["vep-slid"]; ok {
		result["vep_slid"] = flattenObjectSystemNpuNpuTcamTactVepSlid(i["vep-slid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_slid_v"
	if _, ok := i["vep-slid-v"]; ok {
		result["vep_slid_v"] = flattenObjectSystemNpuNpuTcamTactVepSlidV(i["vep-slid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_en"
	if _, ok := i["vep_en"]; ok {
		result["vep_en"] = flattenObjectSystemNpuNpuTcamTactVepEn(i["vep_en"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_lif"
	if _, ok := i["xlt-lif"]; ok {
		result["xlt_lif"] = flattenObjectSystemNpuNpuTcamTactXltLif(i["xlt-lif"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_lif_v"
	if _, ok := i["xlt-lif-v"]; ok {
		result["xlt_lif_v"] = flattenObjectSystemNpuNpuTcamTactXltLifV(i["xlt-lif-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_vid"
	if _, ok := i["xlt-vid"]; ok {
		result["xlt_vid"] = flattenObjectSystemNpuNpuTcamTactXltVid(i["xlt-vid"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_vid_v"
	if _, ok := i["xlt-vid-v"]; ok {
		result["xlt_vid_v"] = flattenObjectSystemNpuNpuTcamTactXltVidV(i["xlt-vid-v"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamTactAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactActV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv4S(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv4SV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv6S(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv6SV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactLnkid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactLnkidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMacId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMacIdV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMssT(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMssTV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv4V(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv6V(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSlifAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSlifActV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSublnkid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSublnkidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTgtvAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTgtvActV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTlifAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTlifActV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTpeid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTpeidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactV6Fe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactV6FeV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepEnV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepSlid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepSlidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepEn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltLif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltLifV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltVid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltVidV(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamVid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNssThreadsOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPbaEim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPbaPortSelectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPerPolicyAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPerSessionAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPleNonSynTcpAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPolicyOffloadLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortCpuMap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuPortCpuMapCpuCore(i["cpu-core"], d, pre_append)
			tmp["cpu_core"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortCpuMap-CpuCore")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectSystemNpuPortCpuMapInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortCpuMap-Interface")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuPortCpuMapCpuCore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortCpuMapInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortNpuMap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuPortNpuMapInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortNpuMap-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "npu_group_index"
		if _, ok := i["npu-group-index"]; ok {
			v := flattenObjectSystemNpuPortNpuMapNpuGroupIndex(i["npu-group-index"], d, pre_append)
			tmp["npu_group_index"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-PortNpuMap-NpuGroupIndex")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuPortNpuMapInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortNpuMapNpuGroupIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortPathOption(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports_using_npu"
	if _, ok := i["ports-using-npu"]; ok {
		result["ports_using_npu"] = flattenObjectSystemNpuPortPathOptionPortsUsingNpu(i["ports-using-npu"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuPortPathOptionPortsUsingNpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemNpuPriorityProtocol(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bfd"
	if _, ok := i["bfd"]; ok {
		result["bfd"] = flattenObjectSystemNpuPriorityProtocolBfd(i["bfd"], d, pre_append)
	}

	pre_append = pre + ".0." + "bgp"
	if _, ok := i["bgp"]; ok {
		result["bgp"] = flattenObjectSystemNpuPriorityProtocolBgp(i["bgp"], d, pre_append)
	}

	pre_append = pre + ".0." + "slbc"
	if _, ok := i["slbc"]; ok {
		result["slbc"] = flattenObjectSystemNpuPriorityProtocolSlbc(i["slbc"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuPriorityProtocolBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolBgp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPriorityProtocolSlbc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPrpSessionClearMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuProcessIcmpByHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPrpPortIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectSystemNpuPrpPortOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectSystemNpuQosMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuQtmBufMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuRdpOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuRecoverNp6Link(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuRpsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSessionAcctInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSessionDeniedOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuShapingStats(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSpaPortSelectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSplitIpsecEngines(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseBackpressure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseHaScan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "gap"
	if _, ok := i["gap"]; ok {
		result["gap"] = flattenObjectSystemNpuSseHaScanGap(i["gap"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_session_cnt"
	if _, ok := i["max-session-cnt"]; ok {
		result["max_session_cnt"] = flattenObjectSystemNpuSseHaScanMaxSessionCnt(i["max-session-cnt"], d, pre_append)
	}

	pre_append = pre + ".0." + "min_duration"
	if _, ok := i["min-duration"]; ok {
		result["min_duration"] = flattenObjectSystemNpuSseHaScanMinDuration(i["min-duration"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuSseHaScanGap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseHaScanMaxSessionCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseHaScanMinDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuStripClearTextPadding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuStripEspPadding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHash(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "computation"
	if _, ok := i["computation"]; ok {
		result["computation"] = flattenObjectSystemNpuSwEhHashComputation(i["computation"], d, pre_append)
	}

	pre_append = pre + ".0." + "destination_ip_lower_16"
	if _, ok := i["destination-ip-lower-16"]; ok {
		result["destination_ip_lower_16"] = flattenObjectSystemNpuSwEhHashDestinationIpLower16(i["destination-ip-lower-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "destination_ip_upper_16"
	if _, ok := i["destination-ip-upper-16"]; ok {
		result["destination_ip_upper_16"] = flattenObjectSystemNpuSwEhHashDestinationIpUpper16(i["destination-ip-upper-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "destination_port"
	if _, ok := i["destination-port"]; ok {
		result["destination_port"] = flattenObjectSystemNpuSwEhHashDestinationPort(i["destination-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := i["ip-protocol"]; ok {
		result["ip_protocol"] = flattenObjectSystemNpuSwEhHashIpProtocol(i["ip-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "netmask_length"
	if _, ok := i["netmask-length"]; ok {
		result["netmask_length"] = flattenObjectSystemNpuSwEhHashNetmaskLength(i["netmask-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_ip_lower_16"
	if _, ok := i["source-ip-lower-16"]; ok {
		result["source_ip_lower_16"] = flattenObjectSystemNpuSwEhHashSourceIpLower16(i["source-ip-lower-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_ip_upper_16"
	if _, ok := i["source-ip-upper-16"]; ok {
		result["source_ip_upper_16"] = flattenObjectSystemNpuSwEhHashSourceIpUpper16(i["source-ip-upper-16"], d, pre_append)
	}

	pre_append = pre + ".0." + "source_port"
	if _, ok := i["source-port"]; ok {
		result["source_port"] = flattenObjectSystemNpuSwEhHashSourcePort(i["source-port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuSwEhHashComputation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationIpLower16(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationIpUpper16(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashDestinationPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashIpProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashNetmaskLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourceIpLower16(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourceIpUpper16(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwEhHashSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwNpBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwTrHash(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "draco15"
	if _, ok := i["draco15"]; ok {
		result["draco15"] = flattenObjectSystemNpuSwTrHashDraco15(i["draco15"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_udp_port"
	if _, ok := i["tcp-udp-port"]; ok {
		result["tcp_udp_port"] = flattenObjectSystemNpuSwTrHashTcpUdpPort(i["tcp-udp-port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuSwTrHashDraco15(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwTrHashTcpUdpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwitchNpHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpRstTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTunnelOverVlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuTcpTimeoutProfileCloseWait(i["close-wait"], d, pre_append)
			tmp["close_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-CloseWait")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fin_wait"
		if _, ok := i["fin-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileFinWait(i["fin-wait"], d, pre_append)
			tmp["fin_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-FinWait")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_sent"
		if _, ok := i["syn-sent"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileSynSent(i["syn-sent"], d, pre_append)
			tmp["syn_sent"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-SynSent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_wait"
		if _, ok := i["syn-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileSynWait(i["syn-wait"], d, pre_append)
			tmp["syn_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-SynWait")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_idle"
		if _, ok := i["tcp-idle"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileTcpIdle(i["tcp-idle"], d, pre_append)
			tmp["tcp_idle"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-TcpIdle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_wait"
		if _, ok := i["time-wait"]; ok {
			v := flattenObjectSystemNpuTcpTimeoutProfileTimeWait(i["time-wait"], d, pre_append)
			tmp["time_wait"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-TcpTimeoutProfile-TimeWait")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuTcpTimeoutProfileCloseWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileFinWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileSynSent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileSynWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileTcpIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileTimeWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUdpTimeoutProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemNpuUdpTimeoutProfileId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-UdpTimeoutProfile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "udp_idle"
		if _, ok := i["udp-idle"]; ok {
			v := flattenObjectSystemNpuUdpTimeoutProfileUdpIdle(i["udp-idle"], d, pre_append)
			tmp["udp_idle"] = fortiAPISubPartPatch(v, "ObjectSystemNpu-UdpTimeoutProfile-UdpIdle")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemNpuUdpTimeoutProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUdpTimeoutProfileUdpIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUespOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUllPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuVlanLookupCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuVxlanOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
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
		if err = d.Set("background_sse_scan", flattenObjectSystemNpuBackgroundSseScan(o["background-sse-scan"], d, "background_sse_scan")); err != nil {
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
			if err = d.Set("background_sse_scan", flattenObjectSystemNpuBackgroundSseScan(o["background-sse-scan"], d, "background_sse_scan")); err != nil {
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

	if err = d.Set("capwap_offload", flattenObjectSystemNpuCapwapOffload(o["capwap-offload"], d, "capwap_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["capwap-offload"], "ObjectSystemNpu-CapwapOffload"); ok {
			if err = d.Set("capwap_offload", vv); err != nil {
				return fmt.Errorf("Error reading capwap_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capwap_offload: %v", err)
		}
	}

	if err = d.Set("dedicated_lacp_queue", flattenObjectSystemNpuDedicatedLacpQueue(o["dedicated-lacp-queue"], d, "dedicated_lacp_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["dedicated-lacp-queue"], "ObjectSystemNpu-DedicatedLacpQueue"); ok {
			if err = d.Set("dedicated_lacp_queue", vv); err != nil {
				return fmt.Errorf("Error reading dedicated_lacp_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dedicated_lacp_queue: %v", err)
		}
	}

	if err = d.Set("dedicated_management_affinity", flattenObjectSystemNpuDedicatedManagementAffinity(o["dedicated-management-affinity"], d, "dedicated_management_affinity")); err != nil {
		if vv, ok := fortiAPIPatch(o["dedicated-management-affinity"], "ObjectSystemNpu-DedicatedManagementAffinity"); ok {
			if err = d.Set("dedicated_management_affinity", vv); err != nil {
				return fmt.Errorf("Error reading dedicated_management_affinity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dedicated_management_affinity: %v", err)
		}
	}

	if err = d.Set("dedicated_management_cpu", flattenObjectSystemNpuDedicatedManagementCpu(o["dedicated-management-cpu"], d, "dedicated_management_cpu")); err != nil {
		if vv, ok := fortiAPIPatch(o["dedicated-management-cpu"], "ObjectSystemNpu-DedicatedManagementCpu"); ok {
			if err = d.Set("dedicated_management_cpu", vv); err != nil {
				return fmt.Errorf("Error reading dedicated_management_cpu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dedicated_management_cpu: %v", err)
		}
	}

	if err = d.Set("default_qos_type", flattenObjectSystemNpuDefaultQosType(o["default-qos-type"], d, "default_qos_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-qos-type"], "ObjectSystemNpu-DefaultQosType"); ok {
			if err = d.Set("default_qos_type", vv); err != nil {
				return fmt.Errorf("Error reading default_qos_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_qos_type: %v", err)
		}
	}

	if err = d.Set("default_tcp_refresh_dir", flattenObjectSystemNpuDefaultTcpRefreshDir(o["default-tcp-refresh-dir"], d, "default_tcp_refresh_dir")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-tcp-refresh-dir"], "ObjectSystemNpu-DefaultTcpRefreshDir"); ok {
			if err = d.Set("default_tcp_refresh_dir", vv); err != nil {
				return fmt.Errorf("Error reading default_tcp_refresh_dir: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_tcp_refresh_dir: %v", err)
		}
	}

	if err = d.Set("default_udp_refresh_dir", flattenObjectSystemNpuDefaultUdpRefreshDir(o["default-udp-refresh-dir"], d, "default_udp_refresh_dir")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-udp-refresh-dir"], "ObjectSystemNpu-DefaultUdpRefreshDir"); ok {
			if err = d.Set("default_udp_refresh_dir", vv); err != nil {
				return fmt.Errorf("Error reading default_udp_refresh_dir: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_udp_refresh_dir: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dos_options", flattenObjectSystemNpuDosOptions(o["dos-options"], d, "dos_options")); err != nil {
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
			if err = d.Set("dos_options", flattenObjectSystemNpuDosOptions(o["dos-options"], d, "dos_options")); err != nil {
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

	if err = d.Set("double_level_mcast_offload", flattenObjectSystemNpuDoubleLevelMcastOffload(o["double-level-mcast-offload"], d, "double_level_mcast_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["double-level-mcast-offload"], "ObjectSystemNpu-DoubleLevelMcastOffload"); ok {
			if err = d.Set("double_level_mcast_offload", vv); err != nil {
				return fmt.Errorf("Error reading double_level_mcast_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading double_level_mcast_offload: %v", err)
		}
	}

	if err = d.Set("dse_timeout", flattenObjectSystemNpuDseTimeout(o["dse-timeout"], d, "dse_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["dse-timeout"], "ObjectSystemNpu-DseTimeout"); ok {
			if err = d.Set("dse_timeout", vv); err != nil {
				return fmt.Errorf("Error reading dse_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dse_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dsw_dts_profile", flattenObjectSystemNpuDswDtsProfile(o["dsw-dts-profile"], d, "dsw_dts_profile")); err != nil {
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
			if err = d.Set("dsw_dts_profile", flattenObjectSystemNpuDswDtsProfile(o["dsw-dts-profile"], d, "dsw_dts_profile")); err != nil {
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
		if err = d.Set("dsw_queue_dts_profile", flattenObjectSystemNpuDswQueueDtsProfile(o["dsw-queue-dts-profile"], d, "dsw_queue_dts_profile")); err != nil {
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
			if err = d.Set("dsw_queue_dts_profile", flattenObjectSystemNpuDswQueueDtsProfile(o["dsw-queue-dts-profile"], d, "dsw_queue_dts_profile")); err != nil {
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

	if err = d.Set("fastpath", flattenObjectSystemNpuFastpath(o["fastpath"], d, "fastpath")); err != nil {
		if vv, ok := fortiAPIPatch(o["fastpath"], "ObjectSystemNpu-Fastpath"); ok {
			if err = d.Set("fastpath", vv); err != nil {
				return fmt.Errorf("Error reading fastpath: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fastpath: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fp_anomaly", flattenObjectSystemNpuFpAnomaly(o["fp-anomaly"], d, "fp_anomaly")); err != nil {
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
			if err = d.Set("fp_anomaly", flattenObjectSystemNpuFpAnomaly(o["fp-anomaly"], d, "fp_anomaly")); err != nil {
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

	if err = d.Set("gtp_enhanced_cpu_range", flattenObjectSystemNpuGtpEnhancedCpuRange(o["gtp-enhanced-cpu-range"], d, "gtp_enhanced_cpu_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-enhanced-cpu-range"], "ObjectSystemNpu-GtpEnhancedCpuRange"); ok {
			if err = d.Set("gtp_enhanced_cpu_range", vv); err != nil {
				return fmt.Errorf("Error reading gtp_enhanced_cpu_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_enhanced_cpu_range: %v", err)
		}
	}

	if err = d.Set("gtp_enhanced_mode", flattenObjectSystemNpuGtpEnhancedMode(o["gtp-enhanced-mode"], d, "gtp_enhanced_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-enhanced-mode"], "ObjectSystemNpu-GtpEnhancedMode"); ok {
			if err = d.Set("gtp_enhanced_mode", vv); err != nil {
				return fmt.Errorf("Error reading gtp_enhanced_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_enhanced_mode: %v", err)
		}
	}

	if err = d.Set("gtp_support", flattenObjectSystemNpuGtpSupport(o["gtp-support"], d, "gtp_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["gtp-support"], "ObjectSystemNpu-GtpSupport"); ok {
			if err = d.Set("gtp_support", vv); err != nil {
				return fmt.Errorf("Error reading gtp_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gtp_support: %v", err)
		}
	}

	if err = d.Set("hash_config", flattenObjectSystemNpuHashConfig(o["hash-config"], d, "hash_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-config"], "ObjectSystemNpu-HashConfig"); ok {
			if err = d.Set("hash_config", vv); err != nil {
				return fmt.Errorf("Error reading hash_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_config: %v", err)
		}
	}

	if err = d.Set("hash_ipv6_sel", flattenObjectSystemNpuHashIpv6Sel(o["hash-ipv6-sel"], d, "hash_ipv6_sel")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-ipv6-sel"], "ObjectSystemNpu-HashIpv6Sel"); ok {
			if err = d.Set("hash_ipv6_sel", vv); err != nil {
				return fmt.Errorf("Error reading hash_ipv6_sel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_ipv6_sel: %v", err)
		}
	}

	if err = d.Set("hash_tbl_spread", flattenObjectSystemNpuHashTblSpread(o["hash-tbl-spread"], d, "hash_tbl_spread")); err != nil {
		if vv, ok := fortiAPIPatch(o["hash-tbl-spread"], "ObjectSystemNpu-HashTblSpread"); ok {
			if err = d.Set("hash_tbl_spread", vv); err != nil {
				return fmt.Errorf("Error reading hash_tbl_spread: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hash_tbl_spread: %v", err)
		}
	}

	if err = d.Set("host_shortcut_mode", flattenObjectSystemNpuHostShortcutMode(o["host-shortcut-mode"], d, "host_shortcut_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-shortcut-mode"], "ObjectSystemNpu-HostShortcutMode"); ok {
			if err = d.Set("host_shortcut_mode", vv); err != nil {
				return fmt.Errorf("Error reading host_shortcut_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_shortcut_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hpe", flattenObjectSystemNpuHpe(o["hpe"], d, "hpe")); err != nil {
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
			if err = d.Set("hpe", flattenObjectSystemNpuHpe(o["hpe"], d, "hpe")); err != nil {
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

	if err = d.Set("htab_dedi_queue_nr", flattenObjectSystemNpuHtabDediQueueNr(o["htab-dedi-queue-nr"], d, "htab_dedi_queue_nr")); err != nil {
		if vv, ok := fortiAPIPatch(o["htab-dedi-queue-nr"], "ObjectSystemNpu-HtabDediQueueNr"); ok {
			if err = d.Set("htab_dedi_queue_nr", vv); err != nil {
				return fmt.Errorf("Error reading htab_dedi_queue_nr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htab_dedi_queue_nr: %v", err)
		}
	}

	if err = d.Set("htab_msg_queue", flattenObjectSystemNpuHtabMsgQueue(o["htab-msg-queue"], d, "htab_msg_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["htab-msg-queue"], "ObjectSystemNpu-HtabMsgQueue"); ok {
			if err = d.Set("htab_msg_queue", vv); err != nil {
				return fmt.Errorf("Error reading htab_msg_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htab_msg_queue: %v", err)
		}
	}

	if err = d.Set("htx_gtse_quota", flattenObjectSystemNpuHtxGtseQuota(o["htx-gtse-quota"], d, "htx_gtse_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["htx-gtse-quota"], "ObjectSystemNpu-HtxGtseQuota"); ok {
			if err = d.Set("htx_gtse_quota", vv); err != nil {
				return fmt.Errorf("Error reading htx_gtse_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htx_gtse_quota: %v", err)
		}
	}

	if err = d.Set("htx_icmp_csum_chk", flattenObjectSystemNpuHtxIcmpCsumChk(o["htx-icmp-csum-chk"], d, "htx_icmp_csum_chk")); err != nil {
		if vv, ok := fortiAPIPatch(o["htx-icmp-csum-chk"], "ObjectSystemNpu-HtxIcmpCsumChk"); ok {
			if err = d.Set("htx_icmp_csum_chk", vv); err != nil {
				return fmt.Errorf("Error reading htx_icmp_csum_chk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading htx_icmp_csum_chk: %v", err)
		}
	}

	if err = d.Set("hw_ha_scan_interval", flattenObjectSystemNpuHwHaScanInterval(o["hw-ha-scan-interval"], d, "hw_ha_scan_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-ha-scan-interval"], "ObjectSystemNpu-HwHaScanInterval"); ok {
			if err = d.Set("hw_ha_scan_interval", vv); err != nil {
				return fmt.Errorf("Error reading hw_ha_scan_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_ha_scan_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("icmp_error_rate_ctrl", flattenObjectSystemNpuIcmpErrorRateCtrl(o["icmp-error-rate-ctrl"], d, "icmp_error_rate_ctrl")); err != nil {
			if vv, ok := fortiAPIPatch(o["icmp-error-rate-ctrl"], "ObjectSystemNpu-IcmpErrorRateCtrl"); ok {
				if err = d.Set("icmp_error_rate_ctrl", vv); err != nil {
					return fmt.Errorf("Error reading icmp_error_rate_ctrl: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icmp_error_rate_ctrl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icmp_error_rate_ctrl"); ok {
			if err = d.Set("icmp_error_rate_ctrl", flattenObjectSystemNpuIcmpErrorRateCtrl(o["icmp-error-rate-ctrl"], d, "icmp_error_rate_ctrl")); err != nil {
				if vv, ok := fortiAPIPatch(o["icmp-error-rate-ctrl"], "ObjectSystemNpu-IcmpErrorRateCtrl"); ok {
					if err = d.Set("icmp_error_rate_ctrl", vv); err != nil {
						return fmt.Errorf("Error reading icmp_error_rate_ctrl: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icmp_error_rate_ctrl: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("icmp_rate_ctrl", flattenObjectSystemNpuIcmpRateCtrl(o["icmp-rate-ctrl"], d, "icmp_rate_ctrl")); err != nil {
			if vv, ok := fortiAPIPatch(o["icmp-rate-ctrl"], "ObjectSystemNpu-IcmpRateCtrl"); ok {
				if err = d.Set("icmp_rate_ctrl", vv); err != nil {
					return fmt.Errorf("Error reading icmp_rate_ctrl: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icmp_rate_ctrl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icmp_rate_ctrl"); ok {
			if err = d.Set("icmp_rate_ctrl", flattenObjectSystemNpuIcmpRateCtrl(o["icmp-rate-ctrl"], d, "icmp_rate_ctrl")); err != nil {
				if vv, ok := fortiAPIPatch(o["icmp-rate-ctrl"], "ObjectSystemNpu-IcmpRateCtrl"); ok {
					if err = d.Set("icmp_rate_ctrl", vv); err != nil {
						return fmt.Errorf("Error reading icmp_rate_ctrl: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icmp_rate_ctrl: %v", err)
				}
			}
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenObjectSystemNpuInboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy"], "ObjectSystemNpu-InboundDscpCopy"); ok {
			if err = d.Set("inbound_dscp_copy", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy_port", flattenObjectSystemNpuInboundDscpCopyPort(o["inbound-dscp-copy-port"], d, "inbound_dscp_copy_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["inbound-dscp-copy-port"], "ObjectSystemNpu-InboundDscpCopyPort"); ok {
			if err = d.Set("inbound_dscp_copy_port", vv); err != nil {
				return fmt.Errorf("Error reading inbound_dscp_copy_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inbound_dscp_copy_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_reassembly", flattenObjectSystemNpuIpReassembly(o["ip-reassembly"], d, "ip_reassembly")); err != nil {
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
			if err = d.Set("ip_reassembly", flattenObjectSystemNpuIpReassembly(o["ip-reassembly"], d, "ip_reassembly")); err != nil {
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

	if err = d.Set("intf_shaping_offload", flattenObjectSystemNpuIntfShapingOffload(o["intf-shaping-offload"], d, "intf_shaping_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf-shaping-offload"], "ObjectSystemNpu-IntfShapingOffload"); ok {
			if err = d.Set("intf_shaping_offload", vv); err != nil {
				return fmt.Errorf("Error reading intf_shaping_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf_shaping_offload: %v", err)
		}
	}

	if err = d.Set("ip_fragment_offload", flattenObjectSystemNpuIpFragmentOffload(o["ip-fragment-offload"], d, "ip_fragment_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-fragment-offload"], "ObjectSystemNpu-IpFragmentOffload"); ok {
			if err = d.Set("ip_fragment_offload", vv); err != nil {
				return fmt.Errorf("Error reading ip_fragment_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_fragment_offload: %v", err)
		}
	}

	if err = d.Set("iph_rsvd_re_cksum", flattenObjectSystemNpuIphRsvdReCksum(o["iph-rsvd-re-cksum"], d, "iph_rsvd_re_cksum")); err != nil {
		if vv, ok := fortiAPIPatch(o["iph-rsvd-re-cksum"], "ObjectSystemNpu-IphRsvdReCksum"); ok {
			if err = d.Set("iph_rsvd_re_cksum", vv); err != nil {
				return fmt.Errorf("Error reading iph_rsvd_re_cksum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iph_rsvd_re_cksum: %v", err)
		}
	}

	if err = d.Set("ippool_overload_high", flattenObjectSystemNpuIppoolOverloadHigh(o["ippool-overload-high"], d, "ippool_overload_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool-overload-high"], "ObjectSystemNpu-IppoolOverloadHigh"); ok {
			if err = d.Set("ippool_overload_high", vv); err != nil {
				return fmt.Errorf("Error reading ippool_overload_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool_overload_high: %v", err)
		}
	}

	if err = d.Set("ippool_overload_low", flattenObjectSystemNpuIppoolOverloadLow(o["ippool-overload-low"], d, "ippool_overload_low")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool-overload-low"], "ObjectSystemNpu-IppoolOverloadLow"); ok {
			if err = d.Set("ippool_overload_low", vv); err != nil {
				return fmt.Errorf("Error reading ippool_overload_low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool_overload_low: %v", err)
		}
	}

	if err = d.Set("ipsec_sts_timeout", flattenObjectSystemNpuIpsecStsTimeout(o["ipsec-STS-timeout"], d, "ipsec_sts_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-STS-timeout"], "ObjectSystemNpu-IpsecStsTimeout"); ok {
			if err = d.Set("ipsec_sts_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_sts_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_sts_timeout: %v", err)
		}
	}

	if err = d.Set("ipsec_dec_subengine_mask", flattenObjectSystemNpuIpsecDecSubengineMask(o["ipsec-dec-subengine-mask"], d, "ipsec_dec_subengine_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-dec-subengine-mask"], "ObjectSystemNpu-IpsecDecSubengineMask"); ok {
			if err = d.Set("ipsec_dec_subengine_mask", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_dec_subengine_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_dec_subengine_mask: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_subengine_mask", flattenObjectSystemNpuIpsecEncSubengineMask(o["ipsec-enc-subengine-mask"], d, "ipsec_enc_subengine_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-enc-subengine-mask"], "ObjectSystemNpu-IpsecEncSubengineMask"); ok {
			if err = d.Set("ipsec_enc_subengine_mask", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_enc_subengine_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_enc_subengine_mask: %v", err)
		}
	}

	if err = d.Set("ipsec_host_dfclr", flattenObjectSystemNpuIpsecHostDfclr(o["ipsec-host-dfclr"], d, "ipsec_host_dfclr")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-host-dfclr"], "ObjectSystemNpu-IpsecHostDfclr"); ok {
			if err = d.Set("ipsec_host_dfclr", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_host_dfclr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_host_dfclr: %v", err)
		}
	}

	if err = d.Set("ipsec_inbound_cache", flattenObjectSystemNpuIpsecInboundCache(o["ipsec-inbound-cache"], d, "ipsec_inbound_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-inbound-cache"], "ObjectSystemNpu-IpsecInboundCache"); ok {
			if err = d.Set("ipsec_inbound_cache", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_inbound_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_inbound_cache: %v", err)
		}
	}

	if err = d.Set("ipsec_local_uesp_port", flattenObjectSystemNpuIpsecLocalUespPort(o["ipsec-local-uesp-port"], d, "ipsec_local_uesp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-local-uesp-port"], "ObjectSystemNpu-IpsecLocalUespPort"); ok {
			if err = d.Set("ipsec_local_uesp_port", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_local_uesp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_local_uesp_port: %v", err)
		}
	}

	if err = d.Set("ipsec_mtu_override", flattenObjectSystemNpuIpsecMtuOverride(o["ipsec-mtu-override"], d, "ipsec_mtu_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-mtu-override"], "ObjectSystemNpu-IpsecMtuOverride"); ok {
			if err = d.Set("ipsec_mtu_override", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_mtu_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_mtu_override: %v", err)
		}
	}

	if err = d.Set("ipsec_ob_np_sel", flattenObjectSystemNpuIpsecObNpSel(o["ipsec-ob-np-sel"], d, "ipsec_ob_np_sel")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-ob-np-sel"], "ObjectSystemNpu-IpsecObNpSel"); ok {
			if err = d.Set("ipsec_ob_np_sel", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_ob_np_sel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_ob_np_sel: %v", err)
		}
	}

	if err = d.Set("ipsec_over_vlink", flattenObjectSystemNpuIpsecOverVlink(o["ipsec-over-vlink"], d, "ipsec_over_vlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-over-vlink"], "ObjectSystemNpu-IpsecOverVlink"); ok {
			if err = d.Set("ipsec_over_vlink", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_over_vlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_over_vlink: %v", err)
		}
	}

	if err = d.Set("ipv4_session_quota", flattenObjectSystemNpuIpv4SessionQuota(o["ipv4-session-quota"], d, "ipv4_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-session-quota"], "ObjectSystemNpu-Ipv4SessionQuota"); ok {
			if err = d.Set("ipv4_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_session_quota: %v", err)
		}
	}

	if err = d.Set("ipv4_session_quota_high", flattenObjectSystemNpuIpv4SessionQuotaHigh(o["ipv4-session-quota-high"], d, "ipv4_session_quota_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-session-quota-high"], "ObjectSystemNpu-Ipv4SessionQuotaHigh"); ok {
			if err = d.Set("ipv4_session_quota_high", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_session_quota_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_session_quota_high: %v", err)
		}
	}

	if err = d.Set("ipv4_session_quota_low", flattenObjectSystemNpuIpv4SessionQuotaLow(o["ipv4-session-quota-low"], d, "ipv4_session_quota_low")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-session-quota-low"], "ObjectSystemNpu-Ipv4SessionQuotaLow"); ok {
			if err = d.Set("ipv4_session_quota_low", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_session_quota_low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_session_quota_low: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix_session_quota", flattenObjectSystemNpuIpv6PrefixSessionQuota(o["ipv6-prefix-session-quota"], d, "ipv6_prefix_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-prefix-session-quota"], "ObjectSystemNpu-Ipv6PrefixSessionQuota"); ok {
			if err = d.Set("ipv6_prefix_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_prefix_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_prefix_session_quota: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix_session_quota_high", flattenObjectSystemNpuIpv6PrefixSessionQuotaHigh(o["ipv6-prefix-session-quota-high"], d, "ipv6_prefix_session_quota_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-prefix-session-quota-high"], "ObjectSystemNpu-Ipv6PrefixSessionQuotaHigh"); ok {
			if err = d.Set("ipv6_prefix_session_quota_high", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_prefix_session_quota_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_prefix_session_quota_high: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix_session_quota_low", flattenObjectSystemNpuIpv6PrefixSessionQuotaLow(o["ipv6-prefix-session-quota-low"], d, "ipv6_prefix_session_quota_low")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-prefix-session-quota-low"], "ObjectSystemNpu-Ipv6PrefixSessionQuotaLow"); ok {
			if err = d.Set("ipv6_prefix_session_quota_low", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_prefix_session_quota_low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_prefix_session_quota_low: %v", err)
		}
	}

	if err = d.Set("ipsec_throughput_msg_frequency", flattenObjectSystemNpuIpsecThroughputMsgFrequency(o["ipsec-throughput-msg-frequency"], d, "ipsec_throughput_msg_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-throughput-msg-frequency"], "ObjectSystemNpu-IpsecThroughputMsgFrequency"); ok {
			if err = d.Set("ipsec_throughput_msg_frequency", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_throughput_msg_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_throughput_msg_frequency: %v", err)
		}
	}

	if err = d.Set("ipt_sts_timeout", flattenObjectSystemNpuIptStsTimeout(o["ipt-STS-timeout"], d, "ipt_sts_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipt-STS-timeout"], "ObjectSystemNpu-IptStsTimeout"); ok {
			if err = d.Set("ipt_sts_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ipt_sts_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipt_sts_timeout: %v", err)
		}
	}

	if err = d.Set("ipt_throughput_msg_frequency", flattenObjectSystemNpuIptThroughputMsgFrequency(o["ipt-throughput-msg-frequency"], d, "ipt_throughput_msg_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipt-throughput-msg-frequency"], "ObjectSystemNpu-IptThroughputMsgFrequency"); ok {
			if err = d.Set("ipt_throughput_msg_frequency", vv); err != nil {
				return fmt.Errorf("Error reading ipt_throughput_msg_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipt_throughput_msg_frequency: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("isf_np_queues", flattenObjectSystemNpuIsfNpQueues(o["isf-np-queues"], d, "isf_np_queues")); err != nil {
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
			if err = d.Set("isf_np_queues", flattenObjectSystemNpuIsfNpQueues(o["isf-np-queues"], d, "isf_np_queues")); err != nil {
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

	if err = d.Set("isf_np_rx_tr_distr", flattenObjectSystemNpuIsfNpRxTrDistr(o["isf-np-rx-tr-distr"], d, "isf_np_rx_tr_distr")); err != nil {
		if vv, ok := fortiAPIPatch(o["isf-np-rx-tr-distr"], "ObjectSystemNpu-IsfNpRxTrDistr"); ok {
			if err = d.Set("isf_np_rx_tr_distr", vv); err != nil {
				return fmt.Errorf("Error reading isf_np_rx_tr_distr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isf_np_rx_tr_distr: %v", err)
		}
	}

	if err = d.Set("lag_out_port_select", flattenObjectSystemNpuLagOutPortSelect(o["lag-out-port-select"], d, "lag_out_port_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["lag-out-port-select"], "ObjectSystemNpu-LagOutPortSelect"); ok {
			if err = d.Set("lag_out_port_select", vv); err != nil {
				return fmt.Errorf("Error reading lag_out_port_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lag_out_port_select: %v", err)
		}
	}

	if err = d.Set("max_receive_unit", flattenObjectSystemNpuMaxReceiveUnit(o["max-receive-unit"], d, "max_receive_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-receive-unit"], "ObjectSystemNpu-MaxReceiveUnit"); ok {
			if err = d.Set("max_receive_unit", vv); err != nil {
				return fmt.Errorf("Error reading max_receive_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_receive_unit: %v", err)
		}
	}

	if err = d.Set("max_session_timeout", flattenObjectSystemNpuMaxSessionTimeout(o["max-session-timeout"], d, "max_session_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-session-timeout"], "ObjectSystemNpu-MaxSessionTimeout"); ok {
			if err = d.Set("max_session_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_session_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_session_timeout: %v", err)
		}
	}

	if err = d.Set("mcast_session_accounting", flattenObjectSystemNpuMcastSessionAccounting(o["mcast-session-accounting"], d, "mcast_session_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-session-accounting"], "ObjectSystemNpu-McastSessionAccounting"); ok {
			if err = d.Set("mcast_session_accounting", vv); err != nil {
				return fmt.Errorf("Error reading mcast_session_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_session_accounting: %v", err)
		}
	}

	if err = d.Set("mcast_session_counting", flattenObjectSystemNpuMcastSessionCounting(o["mcast-session-counting"], d, "mcast_session_counting")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-session-counting"], "ObjectSystemNpu-McastSessionCounting"); ok {
			if err = d.Set("mcast_session_counting", vv); err != nil {
				return fmt.Errorf("Error reading mcast_session_counting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_session_counting: %v", err)
		}
	}

	if err = d.Set("mcast_session_counting6", flattenObjectSystemNpuMcastSessionCounting6(o["mcast-session-counting6"], d, "mcast_session_counting6")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-session-counting6"], "ObjectSystemNpu-McastSessionCounting6"); ok {
			if err = d.Set("mcast_session_counting6", vv); err != nil {
				return fmt.Errorf("Error reading mcast_session_counting6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_session_counting6: %v", err)
		}
	}

	if err = d.Set("napi_break_interval", flattenObjectSystemNpuNapiBreakInterval(o["napi-break-interval"], d, "napi_break_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["napi-break-interval"], "ObjectSystemNpu-NapiBreakInterval"); ok {
			if err = d.Set("napi_break_interval", vv); err != nil {
				return fmt.Errorf("Error reading napi_break_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading napi_break_interval: %v", err)
		}
	}

	if err = d.Set("nat46_force_ipv4_packet_forwarding", flattenObjectSystemNpuNat46ForceIpv4PacketForwarding(o["nat46-force-ipv4-packet-forwarding"], d, "nat46_force_ipv4_packet_forwarding")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat46-force-ipv4-packet-forwarding"], "ObjectSystemNpu-Nat46ForceIpv4PacketForwarding"); ok {
			if err = d.Set("nat46_force_ipv4_packet_forwarding", vv); err != nil {
				return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("np_queues", flattenObjectSystemNpuNpQueues(o["np-queues"], d, "np_queues")); err != nil {
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
			if err = d.Set("np_queues", flattenObjectSystemNpuNpQueues(o["np-queues"], d, "np_queues")); err != nil {
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

	if err = d.Set("np6_cps_optimization_mode", flattenObjectSystemNpuNp6CpsOptimizationMode(o["np6-cps-optimization-mode"], d, "np6_cps_optimization_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["np6-cps-optimization-mode"], "ObjectSystemNpu-Np6CpsOptimizationMode"); ok {
			if err = d.Set("np6_cps_optimization_mode", vv); err != nil {
				return fmt.Errorf("Error reading np6_cps_optimization_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading np6_cps_optimization_mode: %v", err)
		}
	}

	if err = d.Set("npu_group_effective_scope", flattenObjectSystemNpuNpuGroupEffectiveScope(o["npu-group-effective-scope"], d, "npu_group_effective_scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-group-effective-scope"], "ObjectSystemNpu-NpuGroupEffectiveScope"); ok {
			if err = d.Set("npu_group_effective_scope", vv); err != nil {
				return fmt.Errorf("Error reading npu_group_effective_scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_group_effective_scope: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("npu_tcam", flattenObjectSystemNpuNpuTcam(o["npu-tcam"], d, "npu_tcam")); err != nil {
			if vv, ok := fortiAPIPatch(o["npu-tcam"], "ObjectSystemNpu-NpuTcam"); ok {
				if err = d.Set("npu_tcam", vv); err != nil {
					return fmt.Errorf("Error reading npu_tcam: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading npu_tcam: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("npu_tcam"); ok {
			if err = d.Set("npu_tcam", flattenObjectSystemNpuNpuTcam(o["npu-tcam"], d, "npu_tcam")); err != nil {
				if vv, ok := fortiAPIPatch(o["npu-tcam"], "ObjectSystemNpu-NpuTcam"); ok {
					if err = d.Set("npu_tcam", vv); err != nil {
						return fmt.Errorf("Error reading npu_tcam: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading npu_tcam: %v", err)
				}
			}
		}
	}

	if err = d.Set("nss_threads_option", flattenObjectSystemNpuNssThreadsOption(o["nss-threads-option"], d, "nss_threads_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["nss-threads-option"], "ObjectSystemNpu-NssThreadsOption"); ok {
			if err = d.Set("nss_threads_option", vv); err != nil {
				return fmt.Errorf("Error reading nss_threads_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nss_threads_option: %v", err)
		}
	}

	if err = d.Set("pba_eim", flattenObjectSystemNpuPbaEim(o["pba-eim"], d, "pba_eim")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-eim"], "ObjectSystemNpu-PbaEim"); ok {
			if err = d.Set("pba_eim", vv); err != nil {
				return fmt.Errorf("Error reading pba_eim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_eim: %v", err)
		}
	}

	if err = d.Set("pba_port_select_mode", flattenObjectSystemNpuPbaPortSelectMode(o["pba-port-select-mode"], d, "pba_port_select_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-port-select-mode"], "ObjectSystemNpu-PbaPortSelectMode"); ok {
			if err = d.Set("pba_port_select_mode", vv); err != nil {
				return fmt.Errorf("Error reading pba_port_select_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_port_select_mode: %v", err)
		}
	}

	if err = d.Set("per_policy_accounting", flattenObjectSystemNpuPerPolicyAccounting(o["per-policy-accounting"], d, "per_policy_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-policy-accounting"], "ObjectSystemNpu-PerPolicyAccounting"); ok {
			if err = d.Set("per_policy_accounting", vv); err != nil {
				return fmt.Errorf("Error reading per_policy_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_policy_accounting: %v", err)
		}
	}

	if err = d.Set("per_session_accounting", flattenObjectSystemNpuPerSessionAccounting(o["per-session-accounting"], d, "per_session_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-session-accounting"], "ObjectSystemNpu-PerSessionAccounting"); ok {
			if err = d.Set("per_session_accounting", vv); err != nil {
				return fmt.Errorf("Error reading per_session_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_session_accounting: %v", err)
		}
	}

	if err = d.Set("ple_non_syn_tcp_action", flattenObjectSystemNpuPleNonSynTcpAction(o["ple-non-syn-tcp-action"], d, "ple_non_syn_tcp_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["ple-non-syn-tcp-action"], "ObjectSystemNpu-PleNonSynTcpAction"); ok {
			if err = d.Set("ple_non_syn_tcp_action", vv); err != nil {
				return fmt.Errorf("Error reading ple_non_syn_tcp_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ple_non_syn_tcp_action: %v", err)
		}
	}

	if err = d.Set("policy_offload_level", flattenObjectSystemNpuPolicyOffloadLevel(o["policy-offload-level"], d, "policy_offload_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload-level"], "ObjectSystemNpu-PolicyOffloadLevel"); ok {
			if err = d.Set("policy_offload_level", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload_level: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port_cpu_map", flattenObjectSystemNpuPortCpuMap(o["port-cpu-map"], d, "port_cpu_map")); err != nil {
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
			if err = d.Set("port_cpu_map", flattenObjectSystemNpuPortCpuMap(o["port-cpu-map"], d, "port_cpu_map")); err != nil {
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
		if err = d.Set("port_npu_map", flattenObjectSystemNpuPortNpuMap(o["port-npu-map"], d, "port_npu_map")); err != nil {
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
			if err = d.Set("port_npu_map", flattenObjectSystemNpuPortNpuMap(o["port-npu-map"], d, "port_npu_map")); err != nil {
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
		if err = d.Set("port_path_option", flattenObjectSystemNpuPortPathOption(o["port-path-option"], d, "port_path_option")); err != nil {
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
			if err = d.Set("port_path_option", flattenObjectSystemNpuPortPathOption(o["port-path-option"], d, "port_path_option")); err != nil {
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
		if err = d.Set("priority_protocol", flattenObjectSystemNpuPriorityProtocol(o["priority-protocol"], d, "priority_protocol")); err != nil {
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
			if err = d.Set("priority_protocol", flattenObjectSystemNpuPriorityProtocol(o["priority-protocol"], d, "priority_protocol")); err != nil {
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

	if err = d.Set("prp_session_clear_mode", flattenObjectSystemNpuPrpSessionClearMode(o["prp-session-clear-mode"], d, "prp_session_clear_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["prp-session-clear-mode"], "ObjectSystemNpu-PrpSessionClearMode"); ok {
			if err = d.Set("prp_session_clear_mode", vv); err != nil {
				return fmt.Errorf("Error reading prp_session_clear_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prp_session_clear_mode: %v", err)
		}
	}

	if err = d.Set("process_icmp_by_host", flattenObjectSystemNpuProcessIcmpByHost(o["process-icmp-by-host"], d, "process_icmp_by_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["process-icmp-by-host"], "ObjectSystemNpu-ProcessIcmpByHost"); ok {
			if err = d.Set("process_icmp_by_host", vv); err != nil {
				return fmt.Errorf("Error reading process_icmp_by_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading process_icmp_by_host: %v", err)
		}
	}

	if err = d.Set("prp_port_in", flattenObjectSystemNpuPrpPortIn(o["prp-port-in"], d, "prp_port_in")); err != nil {
		if vv, ok := fortiAPIPatch(o["prp-port-in"], "ObjectSystemNpu-PrpPortIn"); ok {
			if err = d.Set("prp_port_in", vv); err != nil {
				return fmt.Errorf("Error reading prp_port_in: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prp_port_in: %v", err)
		}
	}

	if err = d.Set("prp_port_out", flattenObjectSystemNpuPrpPortOut(o["prp-port-out"], d, "prp_port_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["prp-port-out"], "ObjectSystemNpu-PrpPortOut"); ok {
			if err = d.Set("prp_port_out", vv); err != nil {
				return fmt.Errorf("Error reading prp_port_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prp_port_out: %v", err)
		}
	}

	if err = d.Set("qos_mode", flattenObjectSystemNpuQosMode(o["qos-mode"], d, "qos_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-mode"], "ObjectSystemNpu-QosMode"); ok {
			if err = d.Set("qos_mode", vv); err != nil {
				return fmt.Errorf("Error reading qos_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_mode: %v", err)
		}
	}

	if err = d.Set("qtm_buf_mode", flattenObjectSystemNpuQtmBufMode(o["qtm-buf-mode"], d, "qtm_buf_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["qtm-buf-mode"], "ObjectSystemNpu-QtmBufMode"); ok {
			if err = d.Set("qtm_buf_mode", vv); err != nil {
				return fmt.Errorf("Error reading qtm_buf_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qtm_buf_mode: %v", err)
		}
	}

	if err = d.Set("rdp_offload", flattenObjectSystemNpuRdpOffload(o["rdp-offload"], d, "rdp_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["rdp-offload"], "ObjectSystemNpu-RdpOffload"); ok {
			if err = d.Set("rdp_offload", vv); err != nil {
				return fmt.Errorf("Error reading rdp_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rdp_offload: %v", err)
		}
	}

	if err = d.Set("recover_np6_link", flattenObjectSystemNpuRecoverNp6Link(o["recover-np6-link"], d, "recover_np6_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["recover-np6-link"], "ObjectSystemNpu-RecoverNp6Link"); ok {
			if err = d.Set("recover_np6_link", vv); err != nil {
				return fmt.Errorf("Error reading recover_np6_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recover_np6_link: %v", err)
		}
	}

	if err = d.Set("rps_mode", flattenObjectSystemNpuRpsMode(o["rps-mode"], d, "rps_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rps-mode"], "ObjectSystemNpu-RpsMode"); ok {
			if err = d.Set("rps_mode", vv); err != nil {
				return fmt.Errorf("Error reading rps_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rps_mode: %v", err)
		}
	}

	if err = d.Set("session_acct_interval", flattenObjectSystemNpuSessionAcctInterval(o["session-acct-interval"], d, "session_acct_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-acct-interval"], "ObjectSystemNpu-SessionAcctInterval"); ok {
			if err = d.Set("session_acct_interval", vv); err != nil {
				return fmt.Errorf("Error reading session_acct_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_acct_interval: %v", err)
		}
	}

	if err = d.Set("session_denied_offload", flattenObjectSystemNpuSessionDeniedOffload(o["session-denied-offload"], d, "session_denied_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-denied-offload"], "ObjectSystemNpu-SessionDeniedOffload"); ok {
			if err = d.Set("session_denied_offload", vv); err != nil {
				return fmt.Errorf("Error reading session_denied_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_denied_offload: %v", err)
		}
	}

	if err = d.Set("shaping_stats", flattenObjectSystemNpuShapingStats(o["shaping-stats"], d, "shaping_stats")); err != nil {
		if vv, ok := fortiAPIPatch(o["shaping-stats"], "ObjectSystemNpu-ShapingStats"); ok {
			if err = d.Set("shaping_stats", vv); err != nil {
				return fmt.Errorf("Error reading shaping_stats: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shaping_stats: %v", err)
		}
	}

	if err = d.Set("spa_port_select_mode", flattenObjectSystemNpuSpaPortSelectMode(o["spa-port-select-mode"], d, "spa_port_select_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["spa-port-select-mode"], "ObjectSystemNpu-SpaPortSelectMode"); ok {
			if err = d.Set("spa_port_select_mode", vv); err != nil {
				return fmt.Errorf("Error reading spa_port_select_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spa_port_select_mode: %v", err)
		}
	}

	if err = d.Set("split_ipsec_engines", flattenObjectSystemNpuSplitIpsecEngines(o["split-ipsec-engines"], d, "split_ipsec_engines")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-ipsec-engines"], "ObjectSystemNpu-SplitIpsecEngines"); ok {
			if err = d.Set("split_ipsec_engines", vv); err != nil {
				return fmt.Errorf("Error reading split_ipsec_engines: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_ipsec_engines: %v", err)
		}
	}

	if err = d.Set("sse_backpressure", flattenObjectSystemNpuSseBackpressure(o["sse-backpressure"], d, "sse_backpressure")); err != nil {
		if vv, ok := fortiAPIPatch(o["sse-backpressure"], "ObjectSystemNpu-SseBackpressure"); ok {
			if err = d.Set("sse_backpressure", vv); err != nil {
				return fmt.Errorf("Error reading sse_backpressure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sse_backpressure: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sse_ha_scan", flattenObjectSystemNpuSseHaScan(o["sse-ha-scan"], d, "sse_ha_scan")); err != nil {
			if vv, ok := fortiAPIPatch(o["sse-ha-scan"], "ObjectSystemNpu-SseHaScan"); ok {
				if err = d.Set("sse_ha_scan", vv); err != nil {
					return fmt.Errorf("Error reading sse_ha_scan: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sse_ha_scan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sse_ha_scan"); ok {
			if err = d.Set("sse_ha_scan", flattenObjectSystemNpuSseHaScan(o["sse-ha-scan"], d, "sse_ha_scan")); err != nil {
				if vv, ok := fortiAPIPatch(o["sse-ha-scan"], "ObjectSystemNpu-SseHaScan"); ok {
					if err = d.Set("sse_ha_scan", vv); err != nil {
						return fmt.Errorf("Error reading sse_ha_scan: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sse_ha_scan: %v", err)
				}
			}
		}
	}

	if err = d.Set("strip_clear_text_padding", flattenObjectSystemNpuStripClearTextPadding(o["strip-clear-text-padding"], d, "strip_clear_text_padding")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-clear-text-padding"], "ObjectSystemNpu-StripClearTextPadding"); ok {
			if err = d.Set("strip_clear_text_padding", vv); err != nil {
				return fmt.Errorf("Error reading strip_clear_text_padding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_clear_text_padding: %v", err)
		}
	}

	if err = d.Set("strip_esp_padding", flattenObjectSystemNpuStripEspPadding(o["strip-esp-padding"], d, "strip_esp_padding")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-esp-padding"], "ObjectSystemNpu-StripEspPadding"); ok {
			if err = d.Set("strip_esp_padding", vv); err != nil {
				return fmt.Errorf("Error reading strip_esp_padding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_esp_padding: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sw_eh_hash", flattenObjectSystemNpuSwEhHash(o["sw-eh-hash"], d, "sw_eh_hash")); err != nil {
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
			if err = d.Set("sw_eh_hash", flattenObjectSystemNpuSwEhHash(o["sw-eh-hash"], d, "sw_eh_hash")); err != nil {
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

	if err = d.Set("sw_np_bandwidth", flattenObjectSystemNpuSwNpBandwidth(o["sw-np-bandwidth"], d, "sw_np_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-np-bandwidth"], "ObjectSystemNpu-SwNpBandwidth"); ok {
			if err = d.Set("sw_np_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading sw_np_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_np_bandwidth: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sw_tr_hash", flattenObjectSystemNpuSwTrHash(o["sw-tr-hash"], d, "sw_tr_hash")); err != nil {
			if vv, ok := fortiAPIPatch(o["sw-tr-hash"], "ObjectSystemNpu-SwTrHash"); ok {
				if err = d.Set("sw_tr_hash", vv); err != nil {
					return fmt.Errorf("Error reading sw_tr_hash: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sw_tr_hash: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sw_tr_hash"); ok {
			if err = d.Set("sw_tr_hash", flattenObjectSystemNpuSwTrHash(o["sw-tr-hash"], d, "sw_tr_hash")); err != nil {
				if vv, ok := fortiAPIPatch(o["sw-tr-hash"], "ObjectSystemNpu-SwTrHash"); ok {
					if err = d.Set("sw_tr_hash", vv); err != nil {
						return fmt.Errorf("Error reading sw_tr_hash: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sw_tr_hash: %v", err)
				}
			}
		}
	}

	if err = d.Set("switch_np_hash", flattenObjectSystemNpuSwitchNpHash(o["switch-np-hash"], d, "switch_np_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-np-hash"], "ObjectSystemNpu-SwitchNpHash"); ok {
			if err = d.Set("switch_np_hash", vv); err != nil {
				return fmt.Errorf("Error reading switch_np_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_np_hash: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timeout", flattenObjectSystemNpuTcpRstTimeout(o["tcp-rst-timeout"], d, "tcp_rst_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst-timeout"], "ObjectSystemNpu-TcpRstTimeout"); ok {
			if err = d.Set("tcp_rst_timeout", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst_timeout: %v", err)
		}
	}

	if err = d.Set("tunnel_over_vlink", flattenObjectSystemNpuTunnelOverVlink(o["tunnel-over-vlink"], d, "tunnel_over_vlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-over-vlink"], "ObjectSystemNpu-TunnelOverVlink"); ok {
			if err = d.Set("tunnel_over_vlink", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_over_vlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_over_vlink: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tcp_timeout_profile", flattenObjectSystemNpuTcpTimeoutProfile(o["tcp-timeout-profile"], d, "tcp_timeout_profile")); err != nil {
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
			if err = d.Set("tcp_timeout_profile", flattenObjectSystemNpuTcpTimeoutProfile(o["tcp-timeout-profile"], d, "tcp_timeout_profile")); err != nil {
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
		if err = d.Set("udp_timeout_profile", flattenObjectSystemNpuUdpTimeoutProfile(o["udp-timeout-profile"], d, "udp_timeout_profile")); err != nil {
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
			if err = d.Set("udp_timeout_profile", flattenObjectSystemNpuUdpTimeoutProfile(o["udp-timeout-profile"], d, "udp_timeout_profile")); err != nil {
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

	if err = d.Set("uesp_offload", flattenObjectSystemNpuUespOffload(o["uesp-offload"], d, "uesp_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["uesp-offload"], "ObjectSystemNpu-UespOffload"); ok {
			if err = d.Set("uesp_offload", vv); err != nil {
				return fmt.Errorf("Error reading uesp_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uesp_offload: %v", err)
		}
	}

	if err = d.Set("ull_port_mode", flattenObjectSystemNpuUllPortMode(o["ull-port-mode"], d, "ull_port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ull-port-mode"], "ObjectSystemNpu-UllPortMode"); ok {
			if err = d.Set("ull_port_mode", vv); err != nil {
				return fmt.Errorf("Error reading ull_port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ull_port_mode: %v", err)
		}
	}

	if err = d.Set("vlan_lookup_cache", flattenObjectSystemNpuVlanLookupCache(o["vlan-lookup-cache"], d, "vlan_lookup_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-lookup-cache"], "ObjectSystemNpu-VlanLookupCache"); ok {
			if err = d.Set("vlan_lookup_cache", vv); err != nil {
				return fmt.Errorf("Error reading vlan_lookup_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_lookup_cache: %v", err)
		}
	}

	if err = d.Set("vxlan_offload", flattenObjectSystemNpuVxlanOffload(o["vxlan-offload"], d, "vxlan_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["vxlan-offload"], "ObjectSystemNpu-VxlanOffload"); ok {
			if err = d.Set("vxlan_offload", vv); err != nil {
				return fmt.Errorf("Error reading vxlan_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vxlan_offload: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuBackgroundSseScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan"], _ = expandObjectSystemNpuBackgroundSseScanScan(d, i["scan"], pre_append)
	}
	pre_append = pre + ".0." + "scan_stale"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-stale"], _ = expandObjectSystemNpuBackgroundSseScanScanStale(d, i["scan_stale"], pre_append)
	}
	pre_append = pre + ".0." + "scan_vt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["scan-vt"], _ = expandObjectSystemNpuBackgroundSseScanScanVt(d, i["scan_vt"], pre_append)
	}
	pre_append = pre + ".0." + "stats_qual_access"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stats-qual-access"], _ = expandObjectSystemNpuBackgroundSseScanStatsQualAccess(d, i["stats_qual_access"], pre_append)
	}
	pre_append = pre + ".0." + "stats_qual_duration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stats-qual-duration"], _ = expandObjectSystemNpuBackgroundSseScanStatsQualDuration(d, i["stats_qual_duration"], pre_append)
	}
	pre_append = pre + ".0." + "stats_update_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["stats-update-interval"], _ = expandObjectSystemNpuBackgroundSseScanStatsUpdateInterval(d, i["stats_update_interval"], pre_append)
	}
	pre_append = pre + ".0." + "udp_keepalive_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-keepalive-interval"], _ = expandObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(d, i["udp_keepalive_interval"], pre_append)
	}
	pre_append = pre + ".0." + "udp_qual_access"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-qual-access"], _ = expandObjectSystemNpuBackgroundSseScanUdpQualAccess(d, i["udp_qual_access"], pre_append)
	}
	pre_append = pre + ".0." + "udp_qual_duration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-qual-duration"], _ = expandObjectSystemNpuBackgroundSseScanUdpQualDuration(d, i["udp_qual_duration"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuBackgroundSseScanScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanScanStale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanScanVt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsQualAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsQualDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpQualAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpQualDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuCapwapOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDedicatedLacpQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDedicatedManagementAffinity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDedicatedManagementCpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDefaultQosType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDefaultTcpRefreshDir(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDefaultUdpRefreshDir(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "npu_dos_meter_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["npu-dos-meter-mode"], _ = expandObjectSystemNpuDosOptionsNpuDosMeterMode(d, i["npu_dos_meter_mode"], pre_append)
	}
	pre_append = pre + ".0." + "npu_dos_synproxy_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["npu-dos-synproxy-mode"], _ = expandObjectSystemNpuDosOptionsNpuDosSynproxyMode(d, i["npu_dos_synproxy_mode"], pre_append)
	}
	pre_append = pre + ".0." + "npu_dos_tpe_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["npu-dos-tpe-mode"], _ = expandObjectSystemNpuDosOptionsNpuDosTpeMode(d, i["npu_dos_tpe_mode"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuDosOptionsNpuDosMeterMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsNpuDosSynproxyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsNpuDosTpeMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDoubleLevelMcastOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDseTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectSystemNpuDswDtsProfileAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min-limit"], _ = expandObjectSystemNpuDswDtsProfileMinLimit(d, i["min_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["profile-id"], _ = expandObjectSystemNpuDswDtsProfileProfileId(d, i["profile_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "step"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["step"], _ = expandObjectSystemNpuDswDtsProfileStep(d, i["step"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuDswDtsProfileAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileMinLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileStep(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["iport"], _ = expandObjectSystemNpuDswQueueDtsProfileIport(d, i["iport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuDswQueueDtsProfileName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oport"], _ = expandObjectSystemNpuDswQueueDtsProfileOport(d, i["oport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["profile-id"], _ = expandObjectSystemNpuDswQueueDtsProfileProfileId(d, i["profile_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue-select"], _ = expandObjectSystemNpuDswQueueDtsProfileQueueSelect(d, i["queue_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuDswQueueDtsProfileIport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileOport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileQueueSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFastpath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "capwap_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["capwap-minlen-err"], _ = expandObjectSystemNpuFpAnomalyCapwapMinlenErr(d, i["capwap_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "esp_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["esp-minlen-err"], _ = expandObjectSystemNpuFpAnomalyEspMinlenErr(d, i["esp_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "gre_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gre-csum-err"], _ = expandObjectSystemNpuFpAnomalyGreCsumErr(d, i["gre_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "gtpu_plen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gtpu-plen-err"], _ = expandObjectSystemNpuFpAnomalyGtpuPlenErr(d, i["gtpu_plen_err"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-csum-err"], _ = expandObjectSystemNpuFpAnomalyIcmpCsumErr(d, i["icmp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_frag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-frag"], _ = expandObjectSystemNpuFpAnomalyIcmpFrag(d, i["icmp_frag"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_land"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-land"], _ = expandObjectSystemNpuFpAnomalyIcmpLand(d, i["icmp_land"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-minlen-err"], _ = expandObjectSystemNpuFpAnomalyIcmpMinlenErr(d, i["icmp_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-csum-err"], _ = expandObjectSystemNpuFpAnomalyIpv4CsumErr(d, i["ipv4_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_ihl_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-ihl-err"], _ = expandObjectSystemNpuFpAnomalyIpv4IhlErr(d, i["ipv4_ihl_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_land"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-land"], _ = expandObjectSystemNpuFpAnomalyIpv4Land(d, i["ipv4_land"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_len_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-len-err"], _ = expandObjectSystemNpuFpAnomalyIpv4LenErr(d, i["ipv4_len_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_opt_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-opt-err"], _ = expandObjectSystemNpuFpAnomalyIpv4OptErr(d, i["ipv4_opt_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optlsrr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-optlsrr"], _ = expandObjectSystemNpuFpAnomalyIpv4Optlsrr(d, i["ipv4_optlsrr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optrr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-optrr"], _ = expandObjectSystemNpuFpAnomalyIpv4Optrr(d, i["ipv4_optrr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optsecurity"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-optsecurity"], _ = expandObjectSystemNpuFpAnomalyIpv4Optsecurity(d, i["ipv4_optsecurity"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optssrr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-optssrr"], _ = expandObjectSystemNpuFpAnomalyIpv4Optssrr(d, i["ipv4_optssrr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_optstream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-optstream"], _ = expandObjectSystemNpuFpAnomalyIpv4Optstream(d, i["ipv4_optstream"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_opttimestamp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-opttimestamp"], _ = expandObjectSystemNpuFpAnomalyIpv4Opttimestamp(d, i["ipv4_opttimestamp"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_proto_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-proto-err"], _ = expandObjectSystemNpuFpAnomalyIpv4ProtoErr(d, i["ipv4_proto_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_ttlzero_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-ttlzero-err"], _ = expandObjectSystemNpuFpAnomalyIpv4TtlzeroErr(d, i["ipv4_ttlzero_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_unknopt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-unknopt"], _ = expandObjectSystemNpuFpAnomalyIpv4Unknopt(d, i["ipv4_unknopt"], pre_append)
	}
	pre_append = pre + ".0." + "ipv4_ver_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv4-ver-err"], _ = expandObjectSystemNpuFpAnomalyIpv4VerErr(d, i["ipv4_ver_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_daddr_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-daddr-err"], _ = expandObjectSystemNpuFpAnomalyIpv6DaddrErr(d, i["ipv6_daddr_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_exthdr_len_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-exthdr-len-err"], _ = expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(d, i["ipv6_exthdr_len_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_exthdr_order_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-exthdr-order-err"], _ = expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(d, i["ipv6_exthdr_order_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_ihl_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-ihl-err"], _ = expandObjectSystemNpuFpAnomalyIpv6IhlErr(d, i["ipv6_ihl_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_land"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-land"], _ = expandObjectSystemNpuFpAnomalyIpv6Land(d, i["ipv6_land"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optendpid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-optendpid"], _ = expandObjectSystemNpuFpAnomalyIpv6Optendpid(d, i["ipv6_optendpid"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_opthomeaddr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-opthomeaddr"], _ = expandObjectSystemNpuFpAnomalyIpv6Opthomeaddr(d, i["ipv6_opthomeaddr"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optinvld"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-optinvld"], _ = expandObjectSystemNpuFpAnomalyIpv6Optinvld(d, i["ipv6_optinvld"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optjumbo"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-optjumbo"], _ = expandObjectSystemNpuFpAnomalyIpv6Optjumbo(d, i["ipv6_optjumbo"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optnsap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-optnsap"], _ = expandObjectSystemNpuFpAnomalyIpv6Optnsap(d, i["ipv6_optnsap"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_optralert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-optralert"], _ = expandObjectSystemNpuFpAnomalyIpv6Optralert(d, i["ipv6_optralert"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_opttunnel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-opttunnel"], _ = expandObjectSystemNpuFpAnomalyIpv6Opttunnel(d, i["ipv6_opttunnel"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_plen_zero"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-plen-zero"], _ = expandObjectSystemNpuFpAnomalyIpv6PlenZero(d, i["ipv6_plen_zero"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_proto_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-proto-err"], _ = expandObjectSystemNpuFpAnomalyIpv6ProtoErr(d, i["ipv6_proto_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_saddr_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-saddr-err"], _ = expandObjectSystemNpuFpAnomalyIpv6SaddrErr(d, i["ipv6_saddr_err"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_unknopt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-unknopt"], _ = expandObjectSystemNpuFpAnomalyIpv6Unknopt(d, i["ipv6_unknopt"], pre_append)
	}
	pre_append = pre + ".0." + "ipv6_ver_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipv6-ver-err"], _ = expandObjectSystemNpuFpAnomalyIpv6VerErr(d, i["ipv6_ver_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sctp-csum-err"], _ = expandObjectSystemNpuFpAnomalySctpCsumErr(d, i["sctp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "nvgre_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nvgre-minlen-err"], _ = expandObjectSystemNpuFpAnomalyNvgreMinlenErr(d, i["nvgre_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_clen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sctp-clen-err"], _ = expandObjectSystemNpuFpAnomalySctpClenErr(d, i["sctp_clen_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_crc_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sctp-crc-err"], _ = expandObjectSystemNpuFpAnomalySctpCrcErr(d, i["sctp_crc_err"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_l4len_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sctp-l4len-err"], _ = expandObjectSystemNpuFpAnomalySctpL4LenErr(d, i["sctp_l4len_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-csum-err"], _ = expandObjectSystemNpuFpAnomalyTcpCsumErr(d, i["tcp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin_noack"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-fin-noack"], _ = expandObjectSystemNpuFpAnomalyTcpFinNoack(d, i["tcp_fin_noack"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin_only"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-fin-only"], _ = expandObjectSystemNpuFpAnomalyTcpFinOnly(d, i["tcp_fin_only"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_hlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-hlen-err"], _ = expandObjectSystemNpuFpAnomalyTcpHlenErr(d, i["tcp_hlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_hlenvsl4len_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-hlenvsl4len-err"], _ = expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(d, i["tcp_hlenvsl4len_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_land"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-land"], _ = expandObjectSystemNpuFpAnomalyTcpLand(d, i["tcp_land"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_no_flag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-no-flag"], _ = expandObjectSystemNpuFpAnomalyTcpNoFlag(d, i["tcp_no_flag"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_plen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-plen-err"], _ = expandObjectSystemNpuFpAnomalyTcpPlenErr(d, i["tcp_plen_err"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn_data"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-syn-data"], _ = expandObjectSystemNpuFpAnomalyTcpSynData(d, i["tcp_syn_data"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn_fin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-syn-fin"], _ = expandObjectSystemNpuFpAnomalyTcpSynFin(d, i["tcp_syn_fin"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_winnuke"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-winnuke"], _ = expandObjectSystemNpuFpAnomalyTcpWinnuke(d, i["tcp_winnuke"], pre_append)
	}
	pre_append = pre + ".0." + "udp_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-csum-err"], _ = expandObjectSystemNpuFpAnomalyUdpCsumErr(d, i["udp_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "udp_hlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-hlen-err"], _ = expandObjectSystemNpuFpAnomalyUdpHlenErr(d, i["udp_hlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "udp_land"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-land"], _ = expandObjectSystemNpuFpAnomalyUdpLand(d, i["udp_land"], pre_append)
	}
	pre_append = pre + ".0." + "udp_len_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-len-err"], _ = expandObjectSystemNpuFpAnomalyUdpLenErr(d, i["udp_len_err"], pre_append)
	}
	pre_append = pre + ".0." + "udp_plen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-plen-err"], _ = expandObjectSystemNpuFpAnomalyUdpPlenErr(d, i["udp_plen_err"], pre_append)
	}
	pre_append = pre + ".0." + "udplite_cover_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udplite-cover-err"], _ = expandObjectSystemNpuFpAnomalyUdpliteCoverErr(d, i["udplite_cover_err"], pre_append)
	}
	pre_append = pre + ".0." + "udplite_csum_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udplite-csum-err"], _ = expandObjectSystemNpuFpAnomalyUdpliteCsumErr(d, i["udplite_csum_err"], pre_append)
	}
	pre_append = pre + ".0." + "uesp_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["uesp-minlen-err"], _ = expandObjectSystemNpuFpAnomalyUespMinlenErr(d, i["uesp_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "unknproto_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknproto-minlen-err"], _ = expandObjectSystemNpuFpAnomalyUnknprotoMinlenErr(d, i["unknproto_minlen_err"], pre_append)
	}
	pre_append = pre + ".0." + "vxlan_minlen_err"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vxlan-minlen-err"], _ = expandObjectSystemNpuFpAnomalyVxlanMinlenErr(d, i["vxlan_minlen_err"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuFpAnomalyCapwapMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyEspMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGreCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyGtpuPlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpFrag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpLand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIcmpMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4CsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4IhlErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Land(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4LenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4OptErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optlsrr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optrr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optsecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optssrr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Optstream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Opttimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4ProtoErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4TtlzeroErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4Unknopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv4VerErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6DaddrErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrLenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ExthdrOrderErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6IhlErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Land(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optendpid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Opthomeaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optinvld(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optjumbo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optnsap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Optralert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Opttunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6PlenZero(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6ProtoErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6SaddrErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6Unknopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyIpv6VerErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyNvgreMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpClenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpCrcErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalySctpL4LenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinNoack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpFinOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpHlenvsl4LenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpLand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpNoFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpPlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpSynFin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyTcpWinnuke(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpHlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpLenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpPlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCoverErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUdpliteCsumErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUespMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyUnknprotoMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuFpAnomalyVxlanMinlenErr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuGtpEnhancedCpuRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuGtpEnhancedMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuGtpSupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHashConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHashIpv6Sel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHashTblSpread(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHostShortcutMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["all-protocol"], _ = expandObjectSystemNpuHpeAllProtocol(d, i["all_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "arp_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["arp-max"], _ = expandObjectSystemNpuHpeArpMax(d, i["arp_max"], pre_append)
	}
	pre_append = pre + ".0." + "enable_queue_shaper"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["enable-queue-shaper"], _ = expandObjectSystemNpuHpeEnableQueueShaper(d, i["enable_queue_shaper"], pre_append)
	}
	pre_append = pre + ".0." + "enable_shaper"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["enable-shaper"], _ = expandObjectSystemNpuHpeEnableShaper(d, i["enable_shaper"], pre_append)
	}
	pre_append = pre + ".0." + "esp_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["esp-max"], _ = expandObjectSystemNpuHpeEspMax(d, i["esp_max"], pre_append)
	}
	pre_append = pre + ".0." + "exception_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["exception-code"], _ = expandObjectSystemNpuHpeExceptionCode(d, i["exception_code"], pre_append)
	}
	pre_append = pre + ".0." + "fragment_with_sess"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fragment-with-sess"], _ = expandObjectSystemNpuHpeFragmentWithSess(d, i["fragment_with_sess"], pre_append)
	}
	pre_append = pre + ".0." + "fragment_without_session"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fragment-without-session"], _ = expandObjectSystemNpuHpeFragmentWithoutSession(d, i["fragment_without_session"], pre_append)
	}
	pre_append = pre + ".0." + "high_priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["high-priority"], _ = expandObjectSystemNpuHpeHighPriority(d, i["high_priority"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-max"], _ = expandObjectSystemNpuHpeIcmpMax(d, i["icmp_max"], pre_append)
	}
	pre_append = pre + ".0." + "ip_frag_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-frag-max"], _ = expandObjectSystemNpuHpeIpFragMax(d, i["ip_frag_max"], pre_append)
	}
	pre_append = pre + ".0." + "ip_others_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-others-max"], _ = expandObjectSystemNpuHpeIpOthersMax(d, i["ip_others_max"], pre_append)
	}
	pre_append = pre + ".0." + "l2_others_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l2-others-max"], _ = expandObjectSystemNpuHpeL2OthersMax(d, i["l2_others_max"], pre_append)
	}
	pre_append = pre + ".0." + "pri_type_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pri-type-max"], _ = expandObjectSystemNpuHpePriTypeMax(d, i["pri_type_max"], pre_append)
	}
	pre_append = pre + ".0." + "queue_shaper_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["queue-shaper-max"], _ = expandObjectSystemNpuHpeQueueShaperMax(d, i["queue_shaper_max"], pre_append)
	}
	pre_append = pre + ".0." + "sctp_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sctp-max"], _ = expandObjectSystemNpuHpeSctpMax(d, i["sctp_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-max"], _ = expandObjectSystemNpuHpeTcpMax(d, i["tcp_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcpfin_rst_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcpfin-rst-max"], _ = expandObjectSystemNpuHpeTcpfinRstMax(d, i["tcpfin_rst_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcpsyn_ack_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcpsyn-ack-max"], _ = expandObjectSystemNpuHpeTcpsynAckMax(d, i["tcpsyn_ack_max"], pre_append)
	}
	pre_append = pre + ".0." + "tcpsyn_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcpsyn-max"], _ = expandObjectSystemNpuHpeTcpsynMax(d, i["tcpsyn_max"], pre_append)
	}
	pre_append = pre + ".0." + "udp_max"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["udp-max"], _ = expandObjectSystemNpuHpeUdpMax(d, i["udp_max"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuHpeAllProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeArpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEnableQueueShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEnableShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeEspMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeExceptionCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeFragmentWithSess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeFragmentWithoutSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeHighPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIcmpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIpFragMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeIpOthersMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeL2OthersMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpePriTypeMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeQueueShaperMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeSctpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpfinRstMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpsynAckMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeTcpsynMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHpeUdpMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtabDediQueueNr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtabMsgQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtxGtseQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHtxIcmpCsumChk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuHwHaScanInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "icmpv4_error_bucket_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmpv4-error-bucket-size"], _ = expandObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorBucketSize(d, i["icmpv4_error_bucket_size"], pre_append)
	}
	pre_append = pre + ".0." + "icmpv4_error_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmpv4-error-rate"], _ = expandObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRate(d, i["icmpv4_error_rate"], pre_append)
	}
	pre_append = pre + ".0." + "icmpv4_error_rate_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmpv4-error-rate-limit"], _ = expandObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRateLimit(d, i["icmpv4_error_rate_limit"], pre_append)
	}
	pre_append = pre + ".0." + "icmpv6_error_bucket_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmpv6-error-bucket-size"], _ = expandObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorBucketSize(d, i["icmpv6_error_bucket_size"], pre_append)
	}
	pre_append = pre + ".0." + "icmpv6_error_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmpv6-error-rate"], _ = expandObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRate(d, i["icmpv6_error_rate"], pre_append)
	}
	pre_append = pre + ".0." + "icmpv6_error_rate_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmpv6-error-rate-limit"], _ = expandObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRateLimit(d, i["icmpv6_error_rate_limit"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorBucketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrlIcmpv4ErrorRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorBucketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpErrorRateCtrlIcmpv6ErrorRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "icmp_v4_bucket_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-v4-bucket-size"], _ = expandObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize(d, i["icmp_v4_bucket_size"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_v4_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-v4-rate"], _ = expandObjectSystemNpuIcmpRateCtrlIcmpV4Rate(d, i["icmp_v4_rate"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_v6_bucket_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-v6-bucket-size"], _ = expandObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize(d, i["icmp_v6_bucket_size"], pre_append)
	}
	pre_append = pre + ".0." + "icmp_v6_rate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icmp-v6-rate"], _ = expandObjectSystemNpuIcmpRateCtrlIcmpV6Rate(d, i["icmp_v6_rate"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV4Rate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV6Rate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuInboundDscpCopy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuInboundDscpCopyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemNpuIpReassembly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-timeout"], _ = expandObjectSystemNpuIpReassemblyMaxTimeout(d, i["max_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "min_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-timeout"], _ = expandObjectSystemNpuIpReassemblyMinTimeout(d, i["min_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectSystemNpuIpReassemblyStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuIpReassemblyMaxTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpReassemblyMinTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpReassemblyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIntfShapingOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpFragmentOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIphRsvdReCksum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIppoolOverloadHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIppoolOverloadLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecStsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecDecSubengineMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecEncSubengineMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecHostDfclr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecInboundCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecLocalUespPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecMtuOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecObNpSel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecOverVlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpv4SessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpv4SessionQuotaHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpv4SessionQuotaLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpv6PrefixSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpv6PrefixSessionQuotaHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpv6PrefixSessionQuotaLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpsecThroughputMsgFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIptStsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIptThroughputMsgFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIsfNpQueues(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cos0"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos0"], _ = expandObjectSystemNpuIsfNpQueuesCos0(d, i["cos0"], pre_append)
	}
	pre_append = pre + ".0." + "cos1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos1"], _ = expandObjectSystemNpuIsfNpQueuesCos1(d, i["cos1"], pre_append)
	}
	pre_append = pre + ".0." + "cos2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos2"], _ = expandObjectSystemNpuIsfNpQueuesCos2(d, i["cos2"], pre_append)
	}
	pre_append = pre + ".0." + "cos3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos3"], _ = expandObjectSystemNpuIsfNpQueuesCos3(d, i["cos3"], pre_append)
	}
	pre_append = pre + ".0." + "cos4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos4"], _ = expandObjectSystemNpuIsfNpQueuesCos4(d, i["cos4"], pre_append)
	}
	pre_append = pre + ".0." + "cos5"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos5"], _ = expandObjectSystemNpuIsfNpQueuesCos5(d, i["cos5"], pre_append)
	}
	pre_append = pre + ".0." + "cos6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos6"], _ = expandObjectSystemNpuIsfNpQueuesCos6(d, i["cos6"], pre_append)
	}
	pre_append = pre + ".0." + "cos7"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cos7"], _ = expandObjectSystemNpuIsfNpQueuesCos7(d, i["cos7"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuIsfNpQueuesCos0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpQueuesCos7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuIsfNpRxTrDistr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuLagOutPortSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMaxReceiveUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMaxSessionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMcastSessionAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMcastSessionCounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuMcastSessionCounting6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNapiBreakInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNat46ForceIpv4PacketForwarding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueues(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ethernet_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectSystemNpuNpQueuesEthernetType(d, i["ethernet_type"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ethernet-type"] = t
		}
	}
	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectSystemNpuNpQueuesIpProtocol(d, i["ip_protocol"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip-protocol"] = t
		}
	}
	pre_append = pre + ".0." + "ip_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectSystemNpuNpQueuesIpService(d, i["ip_service"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["ip-service"] = t
		}
	}
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectSystemNpuNpQueuesProfile(d, i["profile"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["profile"] = t
		}
	}
	pre_append = pre + ".0." + "scheduler"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectSystemNpuNpQueuesScheduler(d, i["scheduler"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["scheduler"] = t
		}
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesEthernetType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemNpuNpQueuesEthernetTypeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesEthernetTypeQueue(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesEthernetTypeType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesEthernetTypeWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesEthernetTypeWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpProtocolName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpProtocolProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpProtocolQueue(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpProtocolWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpProtocolWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dport"], _ = expandObjectSystemNpuNpQueuesIpServiceDport(d, i["dport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesIpServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectSystemNpuNpQueuesIpServiceProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandObjectSystemNpuNpQueuesIpServiceQueue(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sport"], _ = expandObjectSystemNpuNpQueuesIpServiceSport(d, i["sport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesIpServiceWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesIpServiceDport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceSport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesIpServiceWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos0"], _ = expandObjectSystemNpuNpQueuesProfileCos0(d, i["cos0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos1"], _ = expandObjectSystemNpuNpQueuesProfileCos1(d, i["cos1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos2"], _ = expandObjectSystemNpuNpQueuesProfileCos2(d, i["cos2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos3"], _ = expandObjectSystemNpuNpQueuesProfileCos3(d, i["cos3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos4"], _ = expandObjectSystemNpuNpQueuesProfileCos4(d, i["cos4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos5"], _ = expandObjectSystemNpuNpQueuesProfileCos5(d, i["cos5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos6"], _ = expandObjectSystemNpuNpQueuesProfileCos6(d, i["cos6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos7"], _ = expandObjectSystemNpuNpQueuesProfileCos7(d, i["cos7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp0"], _ = expandObjectSystemNpuNpQueuesProfileDscp0(d, i["dscp0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp1"], _ = expandObjectSystemNpuNpQueuesProfileDscp1(d, i["dscp1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp10"], _ = expandObjectSystemNpuNpQueuesProfileDscp10(d, i["dscp10"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp11"], _ = expandObjectSystemNpuNpQueuesProfileDscp11(d, i["dscp11"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp12"], _ = expandObjectSystemNpuNpQueuesProfileDscp12(d, i["dscp12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp13"], _ = expandObjectSystemNpuNpQueuesProfileDscp13(d, i["dscp13"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp14"], _ = expandObjectSystemNpuNpQueuesProfileDscp14(d, i["dscp14"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp15"], _ = expandObjectSystemNpuNpQueuesProfileDscp15(d, i["dscp15"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp16"], _ = expandObjectSystemNpuNpQueuesProfileDscp16(d, i["dscp16"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp17"], _ = expandObjectSystemNpuNpQueuesProfileDscp17(d, i["dscp17"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp18"], _ = expandObjectSystemNpuNpQueuesProfileDscp18(d, i["dscp18"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp19"], _ = expandObjectSystemNpuNpQueuesProfileDscp19(d, i["dscp19"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp2"], _ = expandObjectSystemNpuNpQueuesProfileDscp2(d, i["dscp2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp20"], _ = expandObjectSystemNpuNpQueuesProfileDscp20(d, i["dscp20"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp21"], _ = expandObjectSystemNpuNpQueuesProfileDscp21(d, i["dscp21"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp22"], _ = expandObjectSystemNpuNpQueuesProfileDscp22(d, i["dscp22"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp23"], _ = expandObjectSystemNpuNpQueuesProfileDscp23(d, i["dscp23"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp24"], _ = expandObjectSystemNpuNpQueuesProfileDscp24(d, i["dscp24"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp25"], _ = expandObjectSystemNpuNpQueuesProfileDscp25(d, i["dscp25"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp26"], _ = expandObjectSystemNpuNpQueuesProfileDscp26(d, i["dscp26"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp27"], _ = expandObjectSystemNpuNpQueuesProfileDscp27(d, i["dscp27"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp28"], _ = expandObjectSystemNpuNpQueuesProfileDscp28(d, i["dscp28"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp29"], _ = expandObjectSystemNpuNpQueuesProfileDscp29(d, i["dscp29"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp3"], _ = expandObjectSystemNpuNpQueuesProfileDscp3(d, i["dscp3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp30"], _ = expandObjectSystemNpuNpQueuesProfileDscp30(d, i["dscp30"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp31"], _ = expandObjectSystemNpuNpQueuesProfileDscp31(d, i["dscp31"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp32"], _ = expandObjectSystemNpuNpQueuesProfileDscp32(d, i["dscp32"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp33"], _ = expandObjectSystemNpuNpQueuesProfileDscp33(d, i["dscp33"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp34"], _ = expandObjectSystemNpuNpQueuesProfileDscp34(d, i["dscp34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp35"], _ = expandObjectSystemNpuNpQueuesProfileDscp35(d, i["dscp35"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp36"], _ = expandObjectSystemNpuNpQueuesProfileDscp36(d, i["dscp36"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp37"], _ = expandObjectSystemNpuNpQueuesProfileDscp37(d, i["dscp37"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp38"], _ = expandObjectSystemNpuNpQueuesProfileDscp38(d, i["dscp38"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp39"], _ = expandObjectSystemNpuNpQueuesProfileDscp39(d, i["dscp39"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp4"], _ = expandObjectSystemNpuNpQueuesProfileDscp4(d, i["dscp4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp40"], _ = expandObjectSystemNpuNpQueuesProfileDscp40(d, i["dscp40"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp41"], _ = expandObjectSystemNpuNpQueuesProfileDscp41(d, i["dscp41"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp42"], _ = expandObjectSystemNpuNpQueuesProfileDscp42(d, i["dscp42"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp43"], _ = expandObjectSystemNpuNpQueuesProfileDscp43(d, i["dscp43"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp44"], _ = expandObjectSystemNpuNpQueuesProfileDscp44(d, i["dscp44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp45"], _ = expandObjectSystemNpuNpQueuesProfileDscp45(d, i["dscp45"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp46"], _ = expandObjectSystemNpuNpQueuesProfileDscp46(d, i["dscp46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp47"], _ = expandObjectSystemNpuNpQueuesProfileDscp47(d, i["dscp47"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp48"], _ = expandObjectSystemNpuNpQueuesProfileDscp48(d, i["dscp48"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp49"], _ = expandObjectSystemNpuNpQueuesProfileDscp49(d, i["dscp49"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp5"], _ = expandObjectSystemNpuNpQueuesProfileDscp5(d, i["dscp5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp50"], _ = expandObjectSystemNpuNpQueuesProfileDscp50(d, i["dscp50"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp51"], _ = expandObjectSystemNpuNpQueuesProfileDscp51(d, i["dscp51"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp52"], _ = expandObjectSystemNpuNpQueuesProfileDscp52(d, i["dscp52"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp53"], _ = expandObjectSystemNpuNpQueuesProfileDscp53(d, i["dscp53"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp54"], _ = expandObjectSystemNpuNpQueuesProfileDscp54(d, i["dscp54"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp55"], _ = expandObjectSystemNpuNpQueuesProfileDscp55(d, i["dscp55"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp56"], _ = expandObjectSystemNpuNpQueuesProfileDscp56(d, i["dscp56"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp57"], _ = expandObjectSystemNpuNpQueuesProfileDscp57(d, i["dscp57"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp58"], _ = expandObjectSystemNpuNpQueuesProfileDscp58(d, i["dscp58"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp59"], _ = expandObjectSystemNpuNpQueuesProfileDscp59(d, i["dscp59"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp6"], _ = expandObjectSystemNpuNpQueuesProfileDscp6(d, i["dscp6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp60"], _ = expandObjectSystemNpuNpQueuesProfileDscp60(d, i["dscp60"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp61"], _ = expandObjectSystemNpuNpQueuesProfileDscp61(d, i["dscp61"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp62"], _ = expandObjectSystemNpuNpQueuesProfileDscp62(d, i["dscp62"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp63"], _ = expandObjectSystemNpuNpQueuesProfileDscp63(d, i["dscp63"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp7"], _ = expandObjectSystemNpuNpQueuesProfileDscp7(d, i["dscp7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp8"], _ = expandObjectSystemNpuNpQueuesProfileDscp8(d, i["dscp8"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp9"], _ = expandObjectSystemNpuNpQueuesProfileDscp9(d, i["dscp9"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemNpuNpQueuesProfileId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpQueuesProfileType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectSystemNpuNpQueuesProfileWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesProfileCos0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp13(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp14(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp17(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp18(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp19(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp20(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp21(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp22(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp23(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp24(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp25(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp26(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp27(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp28(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp29(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp30(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp31(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp32(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp33(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp35(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp36(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp37(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp38(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp39(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp40(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp41(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp42(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp43(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp44(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp45(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp47(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp48(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp49(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp50(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp51(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp52(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp53(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp54(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp55(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp56(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp57(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp58(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp59(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp60(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp61(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp62(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp63(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesScheduler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandObjectSystemNpuNpQueuesSchedulerMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpQueuesSchedulerName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpQueuesSchedulerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesSchedulerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNp6CpsOptimizationMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuGroupEffectiveScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSystemNpuNpuTcamData(d, i["data"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["data"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dbg_dump"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dbg-dump"], _ = expandObjectSystemNpuNpuTcamDbgDump(d, i["dbg_dump"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSystemNpuNpuTcamMask(d, i["mask"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["mask"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mir_act"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSystemNpuNpuTcamMirAct(d, i["mir_act"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["mir-act"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemNpuNpuTcamName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oid"], _ = expandObjectSystemNpuNpuTcamOid(d, i["oid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pri_act"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSystemNpuNpuTcamPriAct(d, i["pri_act"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["pri-act"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sact"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSystemNpuNpuTcamSact(d, i["sact"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sact"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tact"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSystemNpuNpuTcamTact(d, i["tact"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["tact"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemNpuNpuTcamType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vid"], _ = expandObjectSystemNpuNpuTcamVid(d, i["vid"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df"], _ = expandObjectSystemNpuNpuTcamDataDf(d, i["df"], pre_append)
	}
	pre_append = pre + ".0." + "dstip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstip"], _ = expandObjectSystemNpuNpuTcamDataDstip(d, i["dstip"], pre_append)
	}
	pre_append = pre + ".0." + "dstipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstipv6"], _ = expandObjectSystemNpuNpuTcamDataDstipv6(d, i["dstipv6"], pre_append)
	}
	pre_append = pre + ".0." + "dstmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstmac"], _ = expandObjectSystemNpuNpuTcamDataDstmac(d, i["dstmac"], pre_append)
	}
	pre_append = pre + ".0." + "dstport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstport"], _ = expandObjectSystemNpuNpuTcamDataDstport(d, i["dstport"], pre_append)
	}
	pre_append = pre + ".0." + "ethertype"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ethertype"], _ = expandObjectSystemNpuNpuTcamDataEthertype(d, i["ethertype"], pre_append)
	}
	pre_append = pre + ".0." + "ext_tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ext-tag"], _ = expandObjectSystemNpuNpuTcamDataExtTag(d, i["ext_tag"], pre_append)
	}
	pre_append = pre + ".0." + "frag_off"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-off"], _ = expandObjectSystemNpuNpuTcamDataFragOff(d, i["frag_off"], pre_append)
	}
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-buf-cnt"], _ = expandObjectSystemNpuNpuTcamDataGenBufCnt(d, i["gen_buf_cnt"], pre_append)
	}
	pre_append = pre + ".0." + "gen_iv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-iv"], _ = expandObjectSystemNpuNpuTcamDataGenIv(d, i["gen_iv"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l3-flags"], _ = expandObjectSystemNpuNpuTcamDataGenL3Flags(d, i["gen_l3_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l4-flags"], _ = expandObjectSystemNpuNpuTcamDataGenL4Flags(d, i["gen_l4_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pkt-ctrl"], _ = expandObjectSystemNpuNpuTcamDataGenPktCtrl(d, i["gen_pkt_ctrl"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri"], _ = expandObjectSystemNpuNpuTcamDataGenPri(d, i["gen_pri"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri-v"], _ = expandObjectSystemNpuNpuTcamDataGenPriV(d, i["gen_pri_v"], pre_append)
	}
	pre_append = pre + ".0." + "gen_tv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-tv"], _ = expandObjectSystemNpuNpuTcamDataGenTv(d, i["gen_tv"], pre_append)
	}
	pre_append = pre + ".0." + "ihl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ihl"], _ = expandObjectSystemNpuNpuTcamDataIhl(d, i["ihl"], pre_append)
	}
	pre_append = pre + ".0." + "ip4_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip4-id"], _ = expandObjectSystemNpuNpuTcamDataIp4Id(d, i["ip4_id"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-fl"], _ = expandObjectSystemNpuNpuTcamDataIp6Fl(d, i["ip6_fl"], pre_append)
	}
	pre_append = pre + ".0." + "ipver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipver"], _ = expandObjectSystemNpuNpuTcamDataIpver(d, i["ipver"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd10"], _ = expandObjectSystemNpuNpuTcamDataL4Wd10(d, i["l4_wd10"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd11"], _ = expandObjectSystemNpuNpuTcamDataL4Wd11(d, i["l4_wd11"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd8"], _ = expandObjectSystemNpuNpuTcamDataL4Wd8(d, i["l4_wd8"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd9"], _ = expandObjectSystemNpuNpuTcamDataL4Wd9(d, i["l4_wd9"], pre_append)
	}
	pre_append = pre + ".0." + "mf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mf"], _ = expandObjectSystemNpuNpuTcamDataMf(d, i["mf"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectSystemNpuNpuTcamDataProtocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "slink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slink"], _ = expandObjectSystemNpuNpuTcamDataSlink(d, i["slink"], pre_append)
	}
	pre_append = pre + ".0." + "smac_change"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-change"], _ = expandObjectSystemNpuNpuTcamDataSmacChange(d, i["smac_change"], pre_append)
	}
	pre_append = pre + ".0." + "sp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sp"], _ = expandObjectSystemNpuNpuTcamDataSp(d, i["sp"], pre_append)
	}
	pre_append = pre + ".0." + "src_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-cfi"], _ = expandObjectSystemNpuNpuTcamDataSrcCfi(d, i["src_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "src_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-prio"], _ = expandObjectSystemNpuNpuTcamDataSrcPrio(d, i["src_prio"], pre_append)
	}
	pre_append = pre + ".0." + "src_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-updt"], _ = expandObjectSystemNpuNpuTcamDataSrcUpdt(d, i["src_updt"], pre_append)
	}
	pre_append = pre + ".0." + "srcip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcip"], _ = expandObjectSystemNpuNpuTcamDataSrcip(d, i["srcip"], pre_append)
	}
	pre_append = pre + ".0." + "srcipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcipv6"], _ = expandObjectSystemNpuNpuTcamDataSrcipv6(d, i["srcipv6"], pre_append)
	}
	pre_append = pre + ".0." + "srcmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcmac"], _ = expandObjectSystemNpuNpuTcamDataSrcmac(d, i["srcmac"], pre_append)
	}
	pre_append = pre + ".0." + "srcport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcport"], _ = expandObjectSystemNpuNpuTcamDataSrcport(d, i["srcport"], pre_append)
	}
	pre_append = pre + ".0." + "svid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["svid"], _ = expandObjectSystemNpuNpuTcamDataSvid(d, i["svid"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ack"], _ = expandObjectSystemNpuNpuTcamDataTcpAck(d, i["tcp_ack"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-cwr"], _ = expandObjectSystemNpuNpuTcamDataTcpCwr(d, i["tcp_cwr"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ece"], _ = expandObjectSystemNpuNpuTcamDataTcpEce(d, i["tcp_ece"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-fin"], _ = expandObjectSystemNpuNpuTcamDataTcpFin(d, i["tcp_fin"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_push"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-push"], _ = expandObjectSystemNpuNpuTcamDataTcpPush(d, i["tcp_push"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-rst"], _ = expandObjectSystemNpuNpuTcamDataTcpRst(d, i["tcp_rst"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-syn"], _ = expandObjectSystemNpuNpuTcamDataTcpSyn(d, i["tcp_syn"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-urg"], _ = expandObjectSystemNpuNpuTcamDataTcpUrg(d, i["tcp_urg"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-cfi"], _ = expandObjectSystemNpuNpuTcamDataTgtCfi(d, i["tgt_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-prio"], _ = expandObjectSystemNpuNpuTcamDataTgtPrio(d, i["tgt_prio"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-updt"], _ = expandObjectSystemNpuNpuTcamDataTgtUpdt(d, i["tgt_updt"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-v"], _ = expandObjectSystemNpuNpuTcamDataTgtV(d, i["tgt_v"], pre_append)
	}
	pre_append = pre + ".0." + "tos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tos"], _ = expandObjectSystemNpuNpuTcamDataTos(d, i["tos"], pre_append)
	}
	pre_append = pre + ".0." + "tp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp"], _ = expandObjectSystemNpuNpuTcamDataTp(d, i["tp"], pre_append)
	}
	pre_append = pre + ".0." + "ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ttl"], _ = expandObjectSystemNpuNpuTcamDataTtl(d, i["ttl"], pre_append)
	}
	pre_append = pre + ".0." + "tvid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tvid"], _ = expandObjectSystemNpuNpuTcamDataTvid(d, i["tvid"], pre_append)
	}
	pre_append = pre + ".0." + "vdid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdid"], _ = expandObjectSystemNpuNpuTcamDataVdid(d, i["vdid"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamDataDf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstipv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataEthertype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataExtTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataFragOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenBufCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenIv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenL3Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenL4Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPktCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPriV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenTv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIhl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIp4Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIp6Fl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIpver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataMf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSmacChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcCfi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcPrio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcUpdt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcipv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSvid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpAck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpCwr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpEce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpFin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpPush(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpUrg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtCfi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtPrio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtUpdt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTvid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataVdid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDbgDump(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df"], _ = expandObjectSystemNpuNpuTcamMaskDf(d, i["df"], pre_append)
	}
	pre_append = pre + ".0." + "dstip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstip"], _ = expandObjectSystemNpuNpuTcamMaskDstip(d, i["dstip"], pre_append)
	}
	pre_append = pre + ".0." + "dstipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstipv6"], _ = expandObjectSystemNpuNpuTcamMaskDstipv6(d, i["dstipv6"], pre_append)
	}
	pre_append = pre + ".0." + "dstmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstmac"], _ = expandObjectSystemNpuNpuTcamMaskDstmac(d, i["dstmac"], pre_append)
	}
	pre_append = pre + ".0." + "dstport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstport"], _ = expandObjectSystemNpuNpuTcamMaskDstport(d, i["dstport"], pre_append)
	}
	pre_append = pre + ".0." + "ethertype"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ethertype"], _ = expandObjectSystemNpuNpuTcamMaskEthertype(d, i["ethertype"], pre_append)
	}
	pre_append = pre + ".0." + "ext_tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ext-tag"], _ = expandObjectSystemNpuNpuTcamMaskExtTag(d, i["ext_tag"], pre_append)
	}
	pre_append = pre + ".0." + "frag_off"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-off"], _ = expandObjectSystemNpuNpuTcamMaskFragOff(d, i["frag_off"], pre_append)
	}
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-buf-cnt"], _ = expandObjectSystemNpuNpuTcamMaskGenBufCnt(d, i["gen_buf_cnt"], pre_append)
	}
	pre_append = pre + ".0." + "gen_iv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-iv"], _ = expandObjectSystemNpuNpuTcamMaskGenIv(d, i["gen_iv"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l3-flags"], _ = expandObjectSystemNpuNpuTcamMaskGenL3Flags(d, i["gen_l3_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l4-flags"], _ = expandObjectSystemNpuNpuTcamMaskGenL4Flags(d, i["gen_l4_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pkt-ctrl"], _ = expandObjectSystemNpuNpuTcamMaskGenPktCtrl(d, i["gen_pkt_ctrl"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri"], _ = expandObjectSystemNpuNpuTcamMaskGenPri(d, i["gen_pri"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri-v"], _ = expandObjectSystemNpuNpuTcamMaskGenPriV(d, i["gen_pri_v"], pre_append)
	}
	pre_append = pre + ".0." + "gen_tv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-tv"], _ = expandObjectSystemNpuNpuTcamMaskGenTv(d, i["gen_tv"], pre_append)
	}
	pre_append = pre + ".0." + "ihl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ihl"], _ = expandObjectSystemNpuNpuTcamMaskIhl(d, i["ihl"], pre_append)
	}
	pre_append = pre + ".0." + "ip4_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip4-id"], _ = expandObjectSystemNpuNpuTcamMaskIp4Id(d, i["ip4_id"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-fl"], _ = expandObjectSystemNpuNpuTcamMaskIp6Fl(d, i["ip6_fl"], pre_append)
	}
	pre_append = pre + ".0." + "ipver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipver"], _ = expandObjectSystemNpuNpuTcamMaskIpver(d, i["ipver"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd10"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd10(d, i["l4_wd10"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd11"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd11(d, i["l4_wd11"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd8"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd8(d, i["l4_wd8"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd9"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd9(d, i["l4_wd9"], pre_append)
	}
	pre_append = pre + ".0." + "mf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mf"], _ = expandObjectSystemNpuNpuTcamMaskMf(d, i["mf"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectSystemNpuNpuTcamMaskProtocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "slink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slink"], _ = expandObjectSystemNpuNpuTcamMaskSlink(d, i["slink"], pre_append)
	}
	pre_append = pre + ".0." + "smac_change"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-change"], _ = expandObjectSystemNpuNpuTcamMaskSmacChange(d, i["smac_change"], pre_append)
	}
	pre_append = pre + ".0." + "sp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sp"], _ = expandObjectSystemNpuNpuTcamMaskSp(d, i["sp"], pre_append)
	}
	pre_append = pre + ".0." + "src_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-cfi"], _ = expandObjectSystemNpuNpuTcamMaskSrcCfi(d, i["src_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "src_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-prio"], _ = expandObjectSystemNpuNpuTcamMaskSrcPrio(d, i["src_prio"], pre_append)
	}
	pre_append = pre + ".0." + "src_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-updt"], _ = expandObjectSystemNpuNpuTcamMaskSrcUpdt(d, i["src_updt"], pre_append)
	}
	pre_append = pre + ".0." + "srcip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcip"], _ = expandObjectSystemNpuNpuTcamMaskSrcip(d, i["srcip"], pre_append)
	}
	pre_append = pre + ".0." + "srcipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcipv6"], _ = expandObjectSystemNpuNpuTcamMaskSrcipv6(d, i["srcipv6"], pre_append)
	}
	pre_append = pre + ".0." + "srcmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcmac"], _ = expandObjectSystemNpuNpuTcamMaskSrcmac(d, i["srcmac"], pre_append)
	}
	pre_append = pre + ".0." + "srcport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcport"], _ = expandObjectSystemNpuNpuTcamMaskSrcport(d, i["srcport"], pre_append)
	}
	pre_append = pre + ".0." + "svid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["svid"], _ = expandObjectSystemNpuNpuTcamMaskSvid(d, i["svid"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ack"], _ = expandObjectSystemNpuNpuTcamMaskTcpAck(d, i["tcp_ack"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-cwr"], _ = expandObjectSystemNpuNpuTcamMaskTcpCwr(d, i["tcp_cwr"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ece"], _ = expandObjectSystemNpuNpuTcamMaskTcpEce(d, i["tcp_ece"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-fin"], _ = expandObjectSystemNpuNpuTcamMaskTcpFin(d, i["tcp_fin"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_push"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-push"], _ = expandObjectSystemNpuNpuTcamMaskTcpPush(d, i["tcp_push"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-rst"], _ = expandObjectSystemNpuNpuTcamMaskTcpRst(d, i["tcp_rst"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-syn"], _ = expandObjectSystemNpuNpuTcamMaskTcpSyn(d, i["tcp_syn"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-urg"], _ = expandObjectSystemNpuNpuTcamMaskTcpUrg(d, i["tcp_urg"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-cfi"], _ = expandObjectSystemNpuNpuTcamMaskTgtCfi(d, i["tgt_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-prio"], _ = expandObjectSystemNpuNpuTcamMaskTgtPrio(d, i["tgt_prio"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-updt"], _ = expandObjectSystemNpuNpuTcamMaskTgtUpdt(d, i["tgt_updt"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-v"], _ = expandObjectSystemNpuNpuTcamMaskTgtV(d, i["tgt_v"], pre_append)
	}
	pre_append = pre + ".0." + "tos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tos"], _ = expandObjectSystemNpuNpuTcamMaskTos(d, i["tos"], pre_append)
	}
	pre_append = pre + ".0." + "tp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp"], _ = expandObjectSystemNpuNpuTcamMaskTp(d, i["tp"], pre_append)
	}
	pre_append = pre + ".0." + "ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ttl"], _ = expandObjectSystemNpuNpuTcamMaskTtl(d, i["ttl"], pre_append)
	}
	pre_append = pre + ".0." + "tvid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tvid"], _ = expandObjectSystemNpuNpuTcamMaskTvid(d, i["tvid"], pre_append)
	}
	pre_append = pre + ".0." + "vdid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdid"], _ = expandObjectSystemNpuNpuTcamMaskVdid(d, i["vdid"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamMaskDf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstipv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskEthertype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskExtTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskFragOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenBufCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenIv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenL3Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenL4Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPktCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPriV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenTv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIhl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIp4Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIp6Fl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIpver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskMf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSmacChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcCfi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcPrio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcUpdt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcipv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSvid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpAck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpCwr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpEce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpFin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpPush(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpUrg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtCfi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtPrio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtUpdt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTvid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskVdid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMirAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlif"], _ = expandObjectSystemNpuNpuTcamMirActVlif(d, i["vlif"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamMirActVlif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamOid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamPriAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandObjectSystemNpuNpuTcamPriActPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "weight"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["weight"], _ = expandObjectSystemNpuNpuTcamPriActWeight(d, i["weight"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamPriActPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamPriActWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act"], _ = expandObjectSystemNpuNpuTcamSactAct(d, i["act"], pre_append)
	}
	pre_append = pre + ".0." + "act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act-v"], _ = expandObjectSystemNpuNpuTcamSactActV(d, i["act_v"], pre_append)
	}
	pre_append = pre + ".0." + "bmproc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bmproc"], _ = expandObjectSystemNpuNpuTcamSactBmproc(d, i["bmproc"], pre_append)
	}
	pre_append = pre + ".0." + "bmproc_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bmproc-v"], _ = expandObjectSystemNpuNpuTcamSactBmprocV(d, i["bmproc_v"], pre_append)
	}
	pre_append = pre + ".0." + "df_lif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df-lif"], _ = expandObjectSystemNpuNpuTcamSactDfLif(d, i["df_lif"], pre_append)
	}
	pre_append = pre + ".0." + "df_lif_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df-lif-v"], _ = expandObjectSystemNpuNpuTcamSactDfLifV(d, i["df_lif_v"], pre_append)
	}
	pre_append = pre + ".0." + "dfr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dfr"], _ = expandObjectSystemNpuNpuTcamSactDfr(d, i["dfr"], pre_append)
	}
	pre_append = pre + ".0." + "dfr_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dfr-v"], _ = expandObjectSystemNpuNpuTcamSactDfrV(d, i["dfr_v"], pre_append)
	}
	pre_append = pre + ".0." + "dmac_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dmac-skip"], _ = expandObjectSystemNpuNpuTcamSactDmacSkip(d, i["dmac_skip"], pre_append)
	}
	pre_append = pre + ".0." + "dmac_skip_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dmac-skip-v"], _ = expandObjectSystemNpuNpuTcamSactDmacSkipV(d, i["dmac_skip_v"], pre_append)
	}
	pre_append = pre + ".0." + "dosen"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dosen"], _ = expandObjectSystemNpuNpuTcamSactDosen(d, i["dosen"], pre_append)
	}
	pre_append = pre + ".0." + "dosen_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dosen-v"], _ = expandObjectSystemNpuNpuTcamSactDosenV(d, i["dosen_v"], pre_append)
	}
	pre_append = pre + ".0." + "espff_proc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["espff-proc"], _ = expandObjectSystemNpuNpuTcamSactEspffProc(d, i["espff_proc"], pre_append)
	}
	pre_append = pre + ".0." + "espff_proc_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["espff-proc-v"], _ = expandObjectSystemNpuNpuTcamSactEspffProcV(d, i["espff_proc_v"], pre_append)
	}
	pre_append = pre + ".0." + "etype_pid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["etype-pid"], _ = expandObjectSystemNpuNpuTcamSactEtypePid(d, i["etype_pid"], pre_append)
	}
	pre_append = pre + ".0." + "etype_pid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["etype-pid-v"], _ = expandObjectSystemNpuNpuTcamSactEtypePidV(d, i["etype_pid_v"], pre_append)
	}
	pre_append = pre + ".0." + "frag_proc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-proc"], _ = expandObjectSystemNpuNpuTcamSactFragProc(d, i["frag_proc"], pre_append)
	}
	pre_append = pre + ".0." + "frag_proc_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-proc-v"], _ = expandObjectSystemNpuNpuTcamSactFragProcV(d, i["frag_proc_v"], pre_append)
	}
	pre_append = pre + ".0." + "fwd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd"], _ = expandObjectSystemNpuNpuTcamSactFwd(d, i["fwd"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_lif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-lif"], _ = expandObjectSystemNpuNpuTcamSactFwdLif(d, i["fwd_lif"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_lif_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-lif-v"], _ = expandObjectSystemNpuNpuTcamSactFwdLifV(d, i["fwd_lif_v"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_tvid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-tvid"], _ = expandObjectSystemNpuNpuTcamSactFwdTvid(d, i["fwd_tvid"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_tvid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-tvid-v"], _ = expandObjectSystemNpuNpuTcamSactFwdTvidV(d, i["fwd_tvid_v"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-v"], _ = expandObjectSystemNpuNpuTcamSactFwdV(d, i["fwd_v"], pre_append)
	}
	pre_append = pre + ".0." + "icpen"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icpen"], _ = expandObjectSystemNpuNpuTcamSactIcpen(d, i["icpen"], pre_append)
	}
	pre_append = pre + ".0." + "icpen_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icpen-v"], _ = expandObjectSystemNpuNpuTcamSactIcpenV(d, i["icpen_v"], pre_append)
	}
	pre_append = pre + ".0." + "igmp_mld_snp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["igmp-mld-snp"], _ = expandObjectSystemNpuNpuTcamSactIgmpMldSnp(d, i["igmp_mld_snp"], pre_append)
	}
	pre_append = pre + ".0." + "igmp_mld_snp_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["igmp-mld-snp-v"], _ = expandObjectSystemNpuNpuTcamSactIgmpMldSnpV(d, i["igmp_mld_snp_v"], pre_append)
	}
	pre_append = pre + ".0." + "learn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["learn"], _ = expandObjectSystemNpuNpuTcamSactLearn(d, i["learn"], pre_append)
	}
	pre_append = pre + ".0." + "learn_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["learn-v"], _ = expandObjectSystemNpuNpuTcamSactLearnV(d, i["learn_v"], pre_append)
	}
	pre_append = pre + ".0." + "m_srh_ctrl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["m-srh-ctrl"], _ = expandObjectSystemNpuNpuTcamSactMSrhCtrl(d, i["m_srh_ctrl"], pre_append)
	}
	pre_append = pre + ".0." + "m_srh_ctrl_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["m-srh-ctrl-v"], _ = expandObjectSystemNpuNpuTcamSactMSrhCtrlV(d, i["m_srh_ctrl_v"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id"], _ = expandObjectSystemNpuNpuTcamSactMacId(d, i["mac_id"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id-v"], _ = expandObjectSystemNpuNpuTcamSactMacIdV(d, i["mac_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "mss"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss"], _ = expandObjectSystemNpuNpuTcamSactMss(d, i["mss"], pre_append)
	}
	pre_append = pre + ".0." + "mss_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss-v"], _ = expandObjectSystemNpuNpuTcamSactMssV(d, i["mss_v"], pre_append)
	}
	pre_append = pre + ".0." + "pleen"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pleen"], _ = expandObjectSystemNpuNpuTcamSactPleen(d, i["pleen"], pre_append)
	}
	pre_append = pre + ".0." + "pleen_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pleen-v"], _ = expandObjectSystemNpuNpuTcamSactPleenV(d, i["pleen_v"], pre_append)
	}
	pre_append = pre + ".0." + "prio_pid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["prio-pid"], _ = expandObjectSystemNpuNpuTcamSactPrioPid(d, i["prio_pid"], pre_append)
	}
	pre_append = pre + ".0." + "prio_pid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["prio-pid-v"], _ = expandObjectSystemNpuNpuTcamSactPrioPidV(d, i["prio_pid_v"], pre_append)
	}
	pre_append = pre + ".0." + "promis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["promis"], _ = expandObjectSystemNpuNpuTcamSactPromis(d, i["promis"], pre_append)
	}
	pre_append = pre + ".0." + "promis_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["promis-v"], _ = expandObjectSystemNpuNpuTcamSactPromisV(d, i["promis_v"], pre_append)
	}
	pre_append = pre + ".0." + "rfsh"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rfsh"], _ = expandObjectSystemNpuNpuTcamSactRfsh(d, i["rfsh"], pre_append)
	}
	pre_append = pre + ".0." + "rfsh_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rfsh-v"], _ = expandObjectSystemNpuNpuTcamSactRfshV(d, i["rfsh_v"], pre_append)
	}
	pre_append = pre + ".0." + "smac_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-skip"], _ = expandObjectSystemNpuNpuTcamSactSmacSkip(d, i["smac_skip"], pre_append)
	}
	pre_append = pre + ".0." + "smac_skip_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-skip-v"], _ = expandObjectSystemNpuNpuTcamSactSmacSkipV(d, i["smac_skip_v"], pre_append)
	}
	pre_append = pre + ".0." + "tp_smchk_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp-smchk-v"], _ = expandObjectSystemNpuNpuTcamSactTpSmchkV(d, i["tp_smchk_v"], pre_append)
	}
	pre_append = pre + ".0." + "tp_smchk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp_smchk"], _ = expandObjectSystemNpuNpuTcamSactTpSmchk(d, i["tp_smchk"], pre_append)
	}
	pre_append = pre + ".0." + "tpe_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpe-id"], _ = expandObjectSystemNpuNpuTcamSactTpeId(d, i["tpe_id"], pre_append)
	}
	pre_append = pre + ".0." + "tpe_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpe-id-v"], _ = expandObjectSystemNpuNpuTcamSactTpeIdV(d, i["tpe_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "vdm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdm"], _ = expandObjectSystemNpuNpuTcamSactVdm(d, i["vdm"], pre_append)
	}
	pre_append = pre + ".0." + "vdm_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdm-v"], _ = expandObjectSystemNpuNpuTcamSactVdmV(d, i["vdm_v"], pre_append)
	}
	pre_append = pre + ".0." + "vdom_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdom-id"], _ = expandObjectSystemNpuNpuTcamSactVdomId(d, i["vdom_id"], pre_append)
	}
	pre_append = pre + ".0." + "vdom_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdom-id-v"], _ = expandObjectSystemNpuNpuTcamSactVdomIdV(d, i["vdom_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "x_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["x-mode"], _ = expandObjectSystemNpuNpuTcamSactXMode(d, i["x_mode"], pre_append)
	}
	pre_append = pre + ".0." + "x_mode_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["x-mode-v"], _ = expandObjectSystemNpuNpuTcamSactXModeV(d, i["x_mode_v"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamSactAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactActV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactBmproc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactBmprocV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfLif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfLifV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfrV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDmacSkip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDmacSkipV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDosen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDosenV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEspffProc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEspffProcV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEtypePid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEtypePidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFragProc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFragProcV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdLif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdLifV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdTvid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdTvidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIcpen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIcpenV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIgmpMldSnp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIgmpMldSnpV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactLearn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactLearnV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMSrhCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMSrhCtrlV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMacId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMacIdV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMssV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPleen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPleenV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPrioPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPrioPidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPromis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPromisV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactRfsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactRfshV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactSmacSkip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactSmacSkipV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpSmchkV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpSmchk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpeIdV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdmV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdomId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdomIdV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactXMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactXModeV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act"], _ = expandObjectSystemNpuNpuTcamTactAct(d, i["act"], pre_append)
	}
	pre_append = pre + ".0." + "act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act-v"], _ = expandObjectSystemNpuNpuTcamTactActV(d, i["act_v"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv4_s"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv4-s"], _ = expandObjectSystemNpuNpuTcamTactFmtuv4S(d, i["fmtuv4_s"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv4_s_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv4-s-v"], _ = expandObjectSystemNpuNpuTcamTactFmtuv4SV(d, i["fmtuv4_s_v"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv6_s"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv6-s"], _ = expandObjectSystemNpuNpuTcamTactFmtuv6S(d, i["fmtuv6_s"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv6_s_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv6-s-v"], _ = expandObjectSystemNpuNpuTcamTactFmtuv6SV(d, i["fmtuv6_s_v"], pre_append)
	}
	pre_append = pre + ".0." + "lnkid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lnkid"], _ = expandObjectSystemNpuNpuTcamTactLnkid(d, i["lnkid"], pre_append)
	}
	pre_append = pre + ".0." + "lnkid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lnkid-v"], _ = expandObjectSystemNpuNpuTcamTactLnkidV(d, i["lnkid_v"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id"], _ = expandObjectSystemNpuNpuTcamTactMacId(d, i["mac_id"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id-v"], _ = expandObjectSystemNpuNpuTcamTactMacIdV(d, i["mac_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "mss_t"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss-t"], _ = expandObjectSystemNpuNpuTcamTactMssT(d, i["mss_t"], pre_append)
	}
	pre_append = pre + ".0." + "mss_t_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss-t-v"], _ = expandObjectSystemNpuNpuTcamTactMssTV(d, i["mss_t_v"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv4"], _ = expandObjectSystemNpuNpuTcamTactMtuv4(d, i["mtuv4"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv4_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv4-v"], _ = expandObjectSystemNpuNpuTcamTactMtuv4V(d, i["mtuv4_v"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv6"], _ = expandObjectSystemNpuNpuTcamTactMtuv6(d, i["mtuv6"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv6_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv6-v"], _ = expandObjectSystemNpuNpuTcamTactMtuv6V(d, i["mtuv6_v"], pre_append)
	}
	pre_append = pre + ".0." + "slif_act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slif-act"], _ = expandObjectSystemNpuNpuTcamTactSlifAct(d, i["slif_act"], pre_append)
	}
	pre_append = pre + ".0." + "slif_act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slif-act-v"], _ = expandObjectSystemNpuNpuTcamTactSlifActV(d, i["slif_act_v"], pre_append)
	}
	pre_append = pre + ".0." + "sublnkid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sublnkid"], _ = expandObjectSystemNpuNpuTcamTactSublnkid(d, i["sublnkid"], pre_append)
	}
	pre_append = pre + ".0." + "sublnkid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sublnkid-v"], _ = expandObjectSystemNpuNpuTcamTactSublnkidV(d, i["sublnkid_v"], pre_append)
	}
	pre_append = pre + ".0." + "tgtv_act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgtv-act"], _ = expandObjectSystemNpuNpuTcamTactTgtvAct(d, i["tgtv_act"], pre_append)
	}
	pre_append = pre + ".0." + "tgtv_act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgtv-act-v"], _ = expandObjectSystemNpuNpuTcamTactTgtvActV(d, i["tgtv_act_v"], pre_append)
	}
	pre_append = pre + ".0." + "tlif_act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tlif-act"], _ = expandObjectSystemNpuNpuTcamTactTlifAct(d, i["tlif_act"], pre_append)
	}
	pre_append = pre + ".0." + "tlif_act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tlif-act-v"], _ = expandObjectSystemNpuNpuTcamTactTlifActV(d, i["tlif_act_v"], pre_append)
	}
	pre_append = pre + ".0." + "tpeid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpeid"], _ = expandObjectSystemNpuNpuTcamTactTpeid(d, i["tpeid"], pre_append)
	}
	pre_append = pre + ".0." + "tpeid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpeid-v"], _ = expandObjectSystemNpuNpuTcamTactTpeidV(d, i["tpeid_v"], pre_append)
	}
	pre_append = pre + ".0." + "v6fe"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["v6fe"], _ = expandObjectSystemNpuNpuTcamTactV6Fe(d, i["v6fe"], pre_append)
	}
	pre_append = pre + ".0." + "v6fe_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["v6fe-v"], _ = expandObjectSystemNpuNpuTcamTactV6FeV(d, i["v6fe_v"], pre_append)
	}
	pre_append = pre + ".0." + "vep_en_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep-en-v"], _ = expandObjectSystemNpuNpuTcamTactVepEnV(d, i["vep_en_v"], pre_append)
	}
	pre_append = pre + ".0." + "vep_slid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep-slid"], _ = expandObjectSystemNpuNpuTcamTactVepSlid(d, i["vep_slid"], pre_append)
	}
	pre_append = pre + ".0." + "vep_slid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep-slid-v"], _ = expandObjectSystemNpuNpuTcamTactVepSlidV(d, i["vep_slid_v"], pre_append)
	}
	pre_append = pre + ".0." + "vep_en"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep_en"], _ = expandObjectSystemNpuNpuTcamTactVepEn(d, i["vep_en"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_lif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-lif"], _ = expandObjectSystemNpuNpuTcamTactXltLif(d, i["xlt_lif"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_lif_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-lif-v"], _ = expandObjectSystemNpuNpuTcamTactXltLifV(d, i["xlt_lif_v"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_vid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-vid"], _ = expandObjectSystemNpuNpuTcamTactXltVid(d, i["xlt_vid"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_vid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-vid-v"], _ = expandObjectSystemNpuNpuTcamTactXltVidV(d, i["xlt_vid_v"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamTactAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactActV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv4S(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv4SV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv6S(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv6SV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactLnkid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactLnkidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMacId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMacIdV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMssT(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMssTV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv4V(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv6V(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSlifAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSlifActV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSublnkid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSublnkidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTgtvAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTgtvActV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTlifAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTlifActV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTpeid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTpeidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactV6Fe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactV6FeV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepEnV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepSlid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepSlidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepEn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltLif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltLifV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltVid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltVidV(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamVid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNssThreadsOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPbaEim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPbaPortSelectMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPerPolicyAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPerSessionAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPleNonSynTcpAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPolicyOffloadLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortCpuMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cpu_core"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cpu-core"], _ = expandObjectSystemNpuPortCpuMapCpuCore(d, i["cpu_core"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandObjectSystemNpuPortCpuMapInterface(d, i["interface"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuPortCpuMapCpuCore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortCpuMapInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortNpuMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandObjectSystemNpuPortNpuMapInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "npu_group_index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["npu-group-index"], _ = expandObjectSystemNpuPortNpuMapNpuGroupIndex(d, i["npu_group_index"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuPortNpuMapInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortNpuMapNpuGroupIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortPathOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports_using_npu"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ports-using-npu"], _ = expandObjectSystemNpuPortPathOptionPortsUsingNpu(d, i["ports_using_npu"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuPortPathOptionPortsUsingNpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemNpuPriorityProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bfd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bfd"], _ = expandObjectSystemNpuPriorityProtocolBfd(d, i["bfd"], pre_append)
	}
	pre_append = pre + ".0." + "bgp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bgp"], _ = expandObjectSystemNpuPriorityProtocolBgp(d, i["bgp"], pre_append)
	}
	pre_append = pre + ".0." + "slbc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slbc"], _ = expandObjectSystemNpuPriorityProtocolSlbc(d, i["slbc"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuPriorityProtocolBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolBgp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPriorityProtocolSlbc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPrpSessionClearMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuProcessIcmpByHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPrpPortIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuPrpPortOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSystemNpuQosMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuQtmBufMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuRdpOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuRecoverNp6Link(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuRpsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSessionAcctInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSessionDeniedOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuShapingStats(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSpaPortSelectMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSplitIpsecEngines(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseBackpressure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseHaScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "gap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gap"], _ = expandObjectSystemNpuSseHaScanGap(d, i["gap"], pre_append)
	}
	pre_append = pre + ".0." + "max_session_cnt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-session-cnt"], _ = expandObjectSystemNpuSseHaScanMaxSessionCnt(d, i["max_session_cnt"], pre_append)
	}
	pre_append = pre + ".0." + "min_duration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min-duration"], _ = expandObjectSystemNpuSseHaScanMinDuration(d, i["min_duration"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuSseHaScanGap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseHaScanMaxSessionCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseHaScanMinDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuStripClearTextPadding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuStripEspPadding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "computation"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["computation"], _ = expandObjectSystemNpuSwEhHashComputation(d, i["computation"], pre_append)
	}
	pre_append = pre + ".0." + "destination_ip_lower_16"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["destination-ip-lower-16"], _ = expandObjectSystemNpuSwEhHashDestinationIpLower16(d, i["destination_ip_lower_16"], pre_append)
	}
	pre_append = pre + ".0." + "destination_ip_upper_16"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["destination-ip-upper-16"], _ = expandObjectSystemNpuSwEhHashDestinationIpUpper16(d, i["destination_ip_upper_16"], pre_append)
	}
	pre_append = pre + ".0." + "destination_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["destination-port"], _ = expandObjectSystemNpuSwEhHashDestinationPort(d, i["destination_port"], pre_append)
	}
	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip-protocol"], _ = expandObjectSystemNpuSwEhHashIpProtocol(d, i["ip_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "netmask_length"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["netmask-length"], _ = expandObjectSystemNpuSwEhHashNetmaskLength(d, i["netmask_length"], pre_append)
	}
	pre_append = pre + ".0." + "source_ip_lower_16"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["source-ip-lower-16"], _ = expandObjectSystemNpuSwEhHashSourceIpLower16(d, i["source_ip_lower_16"], pre_append)
	}
	pre_append = pre + ".0." + "source_ip_upper_16"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["source-ip-upper-16"], _ = expandObjectSystemNpuSwEhHashSourceIpUpper16(d, i["source_ip_upper_16"], pre_append)
	}
	pre_append = pre + ".0." + "source_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["source-port"], _ = expandObjectSystemNpuSwEhHashSourcePort(d, i["source_port"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuSwEhHashComputation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationIpLower16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationIpUpper16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashDestinationPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashIpProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashNetmaskLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourceIpLower16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourceIpUpper16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwEhHashSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwNpBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwTrHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "draco15"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["draco15"], _ = expandObjectSystemNpuSwTrHashDraco15(d, i["draco15"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_udp_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-udp-port"], _ = expandObjectSystemNpuSwTrHashTcpUdpPort(d, i["tcp_udp_port"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuSwTrHashDraco15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwTrHashTcpUdpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwitchNpHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpRstTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTunnelOverVlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "close_wait"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["close-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileCloseWait(d, i["close_wait"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fin_wait"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fin-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileFinWait(d, i["fin_wait"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemNpuTcpTimeoutProfileId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_sent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["syn-sent"], _ = expandObjectSystemNpuTcpTimeoutProfileSynSent(d, i["syn_sent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syn_wait"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["syn-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileSynWait(d, i["syn_wait"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tcp_idle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tcp-idle"], _ = expandObjectSystemNpuTcpTimeoutProfileTcpIdle(d, i["tcp_idle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time_wait"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["time-wait"], _ = expandObjectSystemNpuTcpTimeoutProfileTimeWait(d, i["time_wait"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuTcpTimeoutProfileCloseWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileFinWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileSynSent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileSynWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileTcpIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileTimeWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUdpTimeoutProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectSystemNpuUdpTimeoutProfileId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "udp_idle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["udp-idle"], _ = expandObjectSystemNpuUdpTimeoutProfileUdpIdle(d, i["udp_idle"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemNpuUdpTimeoutProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUdpTimeoutProfileUdpIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUespOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUllPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuVlanLookupCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuVxlanOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpu(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("background_sse_scan"); ok || d.HasChange("background_sse_scan") {
		t, err := expandObjectSystemNpuBackgroundSseScan(d, v, "background_sse_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["background-sse-scan"] = t
		}
	}

	if v, ok := d.GetOk("capwap_offload"); ok || d.HasChange("capwap_offload") {
		t, err := expandObjectSystemNpuCapwapOffload(d, v, "capwap_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capwap-offload"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_lacp_queue"); ok || d.HasChange("dedicated_lacp_queue") {
		t, err := expandObjectSystemNpuDedicatedLacpQueue(d, v, "dedicated_lacp_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-lacp-queue"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_management_affinity"); ok || d.HasChange("dedicated_management_affinity") {
		t, err := expandObjectSystemNpuDedicatedManagementAffinity(d, v, "dedicated_management_affinity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-management-affinity"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_management_cpu"); ok || d.HasChange("dedicated_management_cpu") {
		t, err := expandObjectSystemNpuDedicatedManagementCpu(d, v, "dedicated_management_cpu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-management-cpu"] = t
		}
	}

	if v, ok := d.GetOk("default_qos_type"); ok || d.HasChange("default_qos_type") {
		t, err := expandObjectSystemNpuDefaultQosType(d, v, "default_qos_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-qos-type"] = t
		}
	}

	if v, ok := d.GetOk("default_tcp_refresh_dir"); ok || d.HasChange("default_tcp_refresh_dir") {
		t, err := expandObjectSystemNpuDefaultTcpRefreshDir(d, v, "default_tcp_refresh_dir")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-tcp-refresh-dir"] = t
		}
	}

	if v, ok := d.GetOk("default_udp_refresh_dir"); ok || d.HasChange("default_udp_refresh_dir") {
		t, err := expandObjectSystemNpuDefaultUdpRefreshDir(d, v, "default_udp_refresh_dir")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-udp-refresh-dir"] = t
		}
	}

	if v, ok := d.GetOk("dos_options"); ok || d.HasChange("dos_options") {
		t, err := expandObjectSystemNpuDosOptions(d, v, "dos_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dos-options"] = t
		}
	}

	if v, ok := d.GetOk("double_level_mcast_offload"); ok || d.HasChange("double_level_mcast_offload") {
		t, err := expandObjectSystemNpuDoubleLevelMcastOffload(d, v, "double_level_mcast_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["double-level-mcast-offload"] = t
		}
	}

	if v, ok := d.GetOk("dse_timeout"); ok || d.HasChange("dse_timeout") {
		t, err := expandObjectSystemNpuDseTimeout(d, v, "dse_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dse-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dsw_dts_profile"); ok || d.HasChange("dsw_dts_profile") {
		t, err := expandObjectSystemNpuDswDtsProfile(d, v, "dsw_dts_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsw-dts-profile"] = t
		}
	}

	if v, ok := d.GetOk("dsw_queue_dts_profile"); ok || d.HasChange("dsw_queue_dts_profile") {
		t, err := expandObjectSystemNpuDswQueueDtsProfile(d, v, "dsw_queue_dts_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsw-queue-dts-profile"] = t
		}
	}

	if v, ok := d.GetOk("fastpath"); ok || d.HasChange("fastpath") {
		t, err := expandObjectSystemNpuFastpath(d, v, "fastpath")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fastpath"] = t
		}
	}

	if v, ok := d.GetOk("fp_anomaly"); ok || d.HasChange("fp_anomaly") {
		t, err := expandObjectSystemNpuFpAnomaly(d, v, "fp_anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp-anomaly"] = t
		}
	}

	if v, ok := d.GetOk("gtp_enhanced_cpu_range"); ok || d.HasChange("gtp_enhanced_cpu_range") {
		t, err := expandObjectSystemNpuGtpEnhancedCpuRange(d, v, "gtp_enhanced_cpu_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-enhanced-cpu-range"] = t
		}
	}

	if v, ok := d.GetOk("gtp_enhanced_mode"); ok || d.HasChange("gtp_enhanced_mode") {
		t, err := expandObjectSystemNpuGtpEnhancedMode(d, v, "gtp_enhanced_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-enhanced-mode"] = t
		}
	}

	if v, ok := d.GetOk("gtp_support"); ok || d.HasChange("gtp_support") {
		t, err := expandObjectSystemNpuGtpSupport(d, v, "gtp_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp-support"] = t
		}
	}

	if v, ok := d.GetOk("hash_config"); ok || d.HasChange("hash_config") {
		t, err := expandObjectSystemNpuHashConfig(d, v, "hash_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-config"] = t
		}
	}

	if v, ok := d.GetOk("hash_ipv6_sel"); ok || d.HasChange("hash_ipv6_sel") {
		t, err := expandObjectSystemNpuHashIpv6Sel(d, v, "hash_ipv6_sel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-ipv6-sel"] = t
		}
	}

	if v, ok := d.GetOk("hash_tbl_spread"); ok || d.HasChange("hash_tbl_spread") {
		t, err := expandObjectSystemNpuHashTblSpread(d, v, "hash_tbl_spread")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-tbl-spread"] = t
		}
	}

	if v, ok := d.GetOk("host_shortcut_mode"); ok || d.HasChange("host_shortcut_mode") {
		t, err := expandObjectSystemNpuHostShortcutMode(d, v, "host_shortcut_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-shortcut-mode"] = t
		}
	}

	if v, ok := d.GetOk("hpe"); ok || d.HasChange("hpe") {
		t, err := expandObjectSystemNpuHpe(d, v, "hpe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hpe"] = t
		}
	}

	if v, ok := d.GetOk("htab_dedi_queue_nr"); ok || d.HasChange("htab_dedi_queue_nr") {
		t, err := expandObjectSystemNpuHtabDediQueueNr(d, v, "htab_dedi_queue_nr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htab-dedi-queue-nr"] = t
		}
	}

	if v, ok := d.GetOk("htab_msg_queue"); ok || d.HasChange("htab_msg_queue") {
		t, err := expandObjectSystemNpuHtabMsgQueue(d, v, "htab_msg_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htab-msg-queue"] = t
		}
	}

	if v, ok := d.GetOk("htx_gtse_quota"); ok || d.HasChange("htx_gtse_quota") {
		t, err := expandObjectSystemNpuHtxGtseQuota(d, v, "htx_gtse_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htx-gtse-quota"] = t
		}
	}

	if v, ok := d.GetOk("htx_icmp_csum_chk"); ok || d.HasChange("htx_icmp_csum_chk") {
		t, err := expandObjectSystemNpuHtxIcmpCsumChk(d, v, "htx_icmp_csum_chk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["htx-icmp-csum-chk"] = t
		}
	}

	if v, ok := d.GetOk("hw_ha_scan_interval"); ok || d.HasChange("hw_ha_scan_interval") {
		t, err := expandObjectSystemNpuHwHaScanInterval(d, v, "hw_ha_scan_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-ha-scan-interval"] = t
		}
	}

	if v, ok := d.GetOk("icmp_error_rate_ctrl"); ok || d.HasChange("icmp_error_rate_ctrl") {
		t, err := expandObjectSystemNpuIcmpErrorRateCtrl(d, v, "icmp_error_rate_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-error-rate-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("icmp_rate_ctrl"); ok || d.HasChange("icmp_rate_ctrl") {
		t, err := expandObjectSystemNpuIcmpRateCtrl(d, v, "icmp_rate_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-rate-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok || d.HasChange("inbound_dscp_copy") {
		t, err := expandObjectSystemNpuInboundDscpCopy(d, v, "inbound_dscp_copy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy_port"); ok || d.HasChange("inbound_dscp_copy_port") {
		t, err := expandObjectSystemNpuInboundDscpCopyPort(d, v, "inbound_dscp_copy_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy-port"] = t
		}
	}

	if v, ok := d.GetOk("ip_reassembly"); ok || d.HasChange("ip_reassembly") {
		t, err := expandObjectSystemNpuIpReassembly(d, v, "ip_reassembly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-reassembly"] = t
		}
	}

	if v, ok := d.GetOk("intf_shaping_offload"); ok || d.HasChange("intf_shaping_offload") {
		t, err := expandObjectSystemNpuIntfShapingOffload(d, v, "intf_shaping_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf-shaping-offload"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_offload"); ok || d.HasChange("ip_fragment_offload") {
		t, err := expandObjectSystemNpuIpFragmentOffload(d, v, "ip_fragment_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-offload"] = t
		}
	}

	if v, ok := d.GetOk("iph_rsvd_re_cksum"); ok || d.HasChange("iph_rsvd_re_cksum") {
		t, err := expandObjectSystemNpuIphRsvdReCksum(d, v, "iph_rsvd_re_cksum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iph-rsvd-re-cksum"] = t
		}
	}

	if v, ok := d.GetOk("ippool_overload_high"); ok || d.HasChange("ippool_overload_high") {
		t, err := expandObjectSystemNpuIppoolOverloadHigh(d, v, "ippool_overload_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool-overload-high"] = t
		}
	}

	if v, ok := d.GetOk("ippool_overload_low"); ok || d.HasChange("ippool_overload_low") {
		t, err := expandObjectSystemNpuIppoolOverloadLow(d, v, "ippool_overload_low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool-overload-low"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_sts_timeout"); ok || d.HasChange("ipsec_sts_timeout") {
		t, err := expandObjectSystemNpuIpsecStsTimeout(d, v, "ipsec_sts_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-STS-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_dec_subengine_mask"); ok || d.HasChange("ipsec_dec_subengine_mask") {
		t, err := expandObjectSystemNpuIpsecDecSubengineMask(d, v, "ipsec_dec_subengine_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-dec-subengine-mask"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_enc_subengine_mask"); ok || d.HasChange("ipsec_enc_subengine_mask") {
		t, err := expandObjectSystemNpuIpsecEncSubengineMask(d, v, "ipsec_enc_subengine_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-enc-subengine-mask"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_host_dfclr"); ok || d.HasChange("ipsec_host_dfclr") {
		t, err := expandObjectSystemNpuIpsecHostDfclr(d, v, "ipsec_host_dfclr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-host-dfclr"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_inbound_cache"); ok || d.HasChange("ipsec_inbound_cache") {
		t, err := expandObjectSystemNpuIpsecInboundCache(d, v, "ipsec_inbound_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-inbound-cache"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_local_uesp_port"); ok || d.HasChange("ipsec_local_uesp_port") {
		t, err := expandObjectSystemNpuIpsecLocalUespPort(d, v, "ipsec_local_uesp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-local-uesp-port"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_mtu_override"); ok || d.HasChange("ipsec_mtu_override") {
		t, err := expandObjectSystemNpuIpsecMtuOverride(d, v, "ipsec_mtu_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-mtu-override"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_ob_np_sel"); ok || d.HasChange("ipsec_ob_np_sel") {
		t, err := expandObjectSystemNpuIpsecObNpSel(d, v, "ipsec_ob_np_sel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-ob-np-sel"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_over_vlink"); ok || d.HasChange("ipsec_over_vlink") {
		t, err := expandObjectSystemNpuIpsecOverVlink(d, v, "ipsec_over_vlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-over-vlink"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_session_quota"); ok || d.HasChange("ipv4_session_quota") {
		t, err := expandObjectSystemNpuIpv4SessionQuota(d, v, "ipv4_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_session_quota_high"); ok || d.HasChange("ipv4_session_quota_high") {
		t, err := expandObjectSystemNpuIpv4SessionQuotaHigh(d, v, "ipv4_session_quota_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-session-quota-high"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_session_quota_low"); ok || d.HasChange("ipv4_session_quota_low") {
		t, err := expandObjectSystemNpuIpv4SessionQuotaLow(d, v, "ipv4_session_quota_low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-session-quota-low"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix_session_quota"); ok || d.HasChange("ipv6_prefix_session_quota") {
		t, err := expandObjectSystemNpuIpv6PrefixSessionQuota(d, v, "ipv6_prefix_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix_session_quota_high"); ok || d.HasChange("ipv6_prefix_session_quota_high") {
		t, err := expandObjectSystemNpuIpv6PrefixSessionQuotaHigh(d, v, "ipv6_prefix_session_quota_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix-session-quota-high"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix_session_quota_low"); ok || d.HasChange("ipv6_prefix_session_quota_low") {
		t, err := expandObjectSystemNpuIpv6PrefixSessionQuotaLow(d, v, "ipv6_prefix_session_quota_low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix-session-quota-low"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_throughput_msg_frequency"); ok || d.HasChange("ipsec_throughput_msg_frequency") {
		t, err := expandObjectSystemNpuIpsecThroughputMsgFrequency(d, v, "ipsec_throughput_msg_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-throughput-msg-frequency"] = t
		}
	}

	if v, ok := d.GetOk("ipt_sts_timeout"); ok || d.HasChange("ipt_sts_timeout") {
		t, err := expandObjectSystemNpuIptStsTimeout(d, v, "ipt_sts_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipt-STS-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ipt_throughput_msg_frequency"); ok || d.HasChange("ipt_throughput_msg_frequency") {
		t, err := expandObjectSystemNpuIptThroughputMsgFrequency(d, v, "ipt_throughput_msg_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipt-throughput-msg-frequency"] = t
		}
	}

	if v, ok := d.GetOk("isf_np_queues"); ok || d.HasChange("isf_np_queues") {
		t, err := expandObjectSystemNpuIsfNpQueues(d, v, "isf_np_queues")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isf-np-queues"] = t
		}
	}

	if v, ok := d.GetOk("isf_np_rx_tr_distr"); ok || d.HasChange("isf_np_rx_tr_distr") {
		t, err := expandObjectSystemNpuIsfNpRxTrDistr(d, v, "isf_np_rx_tr_distr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isf-np-rx-tr-distr"] = t
		}
	}

	if v, ok := d.GetOk("lag_out_port_select"); ok || d.HasChange("lag_out_port_select") {
		t, err := expandObjectSystemNpuLagOutPortSelect(d, v, "lag_out_port_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lag-out-port-select"] = t
		}
	}

	if v, ok := d.GetOk("max_receive_unit"); ok || d.HasChange("max_receive_unit") {
		t, err := expandObjectSystemNpuMaxReceiveUnit(d, v, "max_receive_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-receive-unit"] = t
		}
	}

	if v, ok := d.GetOk("max_session_timeout"); ok || d.HasChange("max_session_timeout") {
		t, err := expandObjectSystemNpuMaxSessionTimeout(d, v, "max_session_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-session-timeout"] = t
		}
	}

	if v, ok := d.GetOk("mcast_session_accounting"); ok || d.HasChange("mcast_session_accounting") {
		t, err := expandObjectSystemNpuMcastSessionAccounting(d, v, "mcast_session_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-session-accounting"] = t
		}
	}

	if v, ok := d.GetOk("mcast_session_counting"); ok || d.HasChange("mcast_session_counting") {
		t, err := expandObjectSystemNpuMcastSessionCounting(d, v, "mcast_session_counting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-session-counting"] = t
		}
	}

	if v, ok := d.GetOk("mcast_session_counting6"); ok || d.HasChange("mcast_session_counting6") {
		t, err := expandObjectSystemNpuMcastSessionCounting6(d, v, "mcast_session_counting6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-session-counting6"] = t
		}
	}

	if v, ok := d.GetOk("napi_break_interval"); ok || d.HasChange("napi_break_interval") {
		t, err := expandObjectSystemNpuNapiBreakInterval(d, v, "napi_break_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["napi-break-interval"] = t
		}
	}

	if v, ok := d.GetOk("nat46_force_ipv4_packet_forwarding"); ok || d.HasChange("nat46_force_ipv4_packet_forwarding") {
		t, err := expandObjectSystemNpuNat46ForceIpv4PacketForwarding(d, v, "nat46_force_ipv4_packet_forwarding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46-force-ipv4-packet-forwarding"] = t
		}
	}

	if v, ok := d.GetOk("np_queues"); ok || d.HasChange("np_queues") {
		t, err := expandObjectSystemNpuNpQueues(d, v, "np_queues")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-queues"] = t
		}
	}

	if v, ok := d.GetOk("np6_cps_optimization_mode"); ok || d.HasChange("np6_cps_optimization_mode") {
		t, err := expandObjectSystemNpuNp6CpsOptimizationMode(d, v, "np6_cps_optimization_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np6-cps-optimization-mode"] = t
		}
	}

	if v, ok := d.GetOk("npu_group_effective_scope"); ok || d.HasChange("npu_group_effective_scope") {
		t, err := expandObjectSystemNpuNpuGroupEffectiveScope(d, v, "npu_group_effective_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-group-effective-scope"] = t
		}
	}

	if v, ok := d.GetOk("npu_tcam"); ok || d.HasChange("npu_tcam") {
		t, err := expandObjectSystemNpuNpuTcam(d, v, "npu_tcam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-tcam"] = t
		}
	}

	if v, ok := d.GetOk("nss_threads_option"); ok || d.HasChange("nss_threads_option") {
		t, err := expandObjectSystemNpuNssThreadsOption(d, v, "nss_threads_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nss-threads-option"] = t
		}
	}

	if v, ok := d.GetOk("pba_eim"); ok || d.HasChange("pba_eim") {
		t, err := expandObjectSystemNpuPbaEim(d, v, "pba_eim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-eim"] = t
		}
	}

	if v, ok := d.GetOk("pba_port_select_mode"); ok || d.HasChange("pba_port_select_mode") {
		t, err := expandObjectSystemNpuPbaPortSelectMode(d, v, "pba_port_select_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-port-select-mode"] = t
		}
	}

	if v, ok := d.GetOk("per_policy_accounting"); ok || d.HasChange("per_policy_accounting") {
		t, err := expandObjectSystemNpuPerPolicyAccounting(d, v, "per_policy_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy-accounting"] = t
		}
	}

	if v, ok := d.GetOk("per_session_accounting"); ok || d.HasChange("per_session_accounting") {
		t, err := expandObjectSystemNpuPerSessionAccounting(d, v, "per_session_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-session-accounting"] = t
		}
	}

	if v, ok := d.GetOk("ple_non_syn_tcp_action"); ok || d.HasChange("ple_non_syn_tcp_action") {
		t, err := expandObjectSystemNpuPleNonSynTcpAction(d, v, "ple_non_syn_tcp_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ple-non-syn-tcp-action"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload_level"); ok || d.HasChange("policy_offload_level") {
		t, err := expandObjectSystemNpuPolicyOffloadLevel(d, v, "policy_offload_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload-level"] = t
		}
	}

	if v, ok := d.GetOk("port_cpu_map"); ok || d.HasChange("port_cpu_map") {
		t, err := expandObjectSystemNpuPortCpuMap(d, v, "port_cpu_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-cpu-map"] = t
		}
	}

	if v, ok := d.GetOk("port_npu_map"); ok || d.HasChange("port_npu_map") {
		t, err := expandObjectSystemNpuPortNpuMap(d, v, "port_npu_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-npu-map"] = t
		}
	}

	if v, ok := d.GetOk("port_path_option"); ok || d.HasChange("port_path_option") {
		t, err := expandObjectSystemNpuPortPathOption(d, v, "port_path_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-path-option"] = t
		}
	}

	if v, ok := d.GetOk("priority_protocol"); ok || d.HasChange("priority_protocol") {
		t, err := expandObjectSystemNpuPriorityProtocol(d, v, "priority_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-protocol"] = t
		}
	}

	if v, ok := d.GetOk("prp_session_clear_mode"); ok || d.HasChange("prp_session_clear_mode") {
		t, err := expandObjectSystemNpuPrpSessionClearMode(d, v, "prp_session_clear_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-session-clear-mode"] = t
		}
	}

	if v, ok := d.GetOk("process_icmp_by_host"); ok || d.HasChange("process_icmp_by_host") {
		t, err := expandObjectSystemNpuProcessIcmpByHost(d, v, "process_icmp_by_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["process-icmp-by-host"] = t
		}
	}

	if v, ok := d.GetOk("prp_port_in"); ok || d.HasChange("prp_port_in") {
		t, err := expandObjectSystemNpuPrpPortIn(d, v, "prp_port_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-port-in"] = t
		}
	}

	if v, ok := d.GetOk("prp_port_out"); ok || d.HasChange("prp_port_out") {
		t, err := expandObjectSystemNpuPrpPortOut(d, v, "prp_port_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-port-out"] = t
		}
	}

	if v, ok := d.GetOk("qos_mode"); ok || d.HasChange("qos_mode") {
		t, err := expandObjectSystemNpuQosMode(d, v, "qos_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-mode"] = t
		}
	}

	if v, ok := d.GetOk("qtm_buf_mode"); ok || d.HasChange("qtm_buf_mode") {
		t, err := expandObjectSystemNpuQtmBufMode(d, v, "qtm_buf_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qtm-buf-mode"] = t
		}
	}

	if v, ok := d.GetOk("rdp_offload"); ok || d.HasChange("rdp_offload") {
		t, err := expandObjectSystemNpuRdpOffload(d, v, "rdp_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rdp-offload"] = t
		}
	}

	if v, ok := d.GetOk("recover_np6_link"); ok || d.HasChange("recover_np6_link") {
		t, err := expandObjectSystemNpuRecoverNp6Link(d, v, "recover_np6_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recover-np6-link"] = t
		}
	}

	if v, ok := d.GetOk("rps_mode"); ok || d.HasChange("rps_mode") {
		t, err := expandObjectSystemNpuRpsMode(d, v, "rps_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rps-mode"] = t
		}
	}

	if v, ok := d.GetOk("session_acct_interval"); ok || d.HasChange("session_acct_interval") {
		t, err := expandObjectSystemNpuSessionAcctInterval(d, v, "session_acct_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-acct-interval"] = t
		}
	}

	if v, ok := d.GetOk("session_denied_offload"); ok || d.HasChange("session_denied_offload") {
		t, err := expandObjectSystemNpuSessionDeniedOffload(d, v, "session_denied_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-denied-offload"] = t
		}
	}

	if v, ok := d.GetOk("shaping_stats"); ok || d.HasChange("shaping_stats") {
		t, err := expandObjectSystemNpuShapingStats(d, v, "shaping_stats")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaping-stats"] = t
		}
	}

	if v, ok := d.GetOk("spa_port_select_mode"); ok || d.HasChange("spa_port_select_mode") {
		t, err := expandObjectSystemNpuSpaPortSelectMode(d, v, "spa_port_select_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spa-port-select-mode"] = t
		}
	}

	if v, ok := d.GetOk("split_ipsec_engines"); ok || d.HasChange("split_ipsec_engines") {
		t, err := expandObjectSystemNpuSplitIpsecEngines(d, v, "split_ipsec_engines")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-ipsec-engines"] = t
		}
	}

	if v, ok := d.GetOk("sse_backpressure"); ok || d.HasChange("sse_backpressure") {
		t, err := expandObjectSystemNpuSseBackpressure(d, v, "sse_backpressure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sse-backpressure"] = t
		}
	}

	if v, ok := d.GetOk("sse_ha_scan"); ok || d.HasChange("sse_ha_scan") {
		t, err := expandObjectSystemNpuSseHaScan(d, v, "sse_ha_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sse-ha-scan"] = t
		}
	}

	if v, ok := d.GetOk("strip_clear_text_padding"); ok || d.HasChange("strip_clear_text_padding") {
		t, err := expandObjectSystemNpuStripClearTextPadding(d, v, "strip_clear_text_padding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-clear-text-padding"] = t
		}
	}

	if v, ok := d.GetOk("strip_esp_padding"); ok || d.HasChange("strip_esp_padding") {
		t, err := expandObjectSystemNpuStripEspPadding(d, v, "strip_esp_padding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-esp-padding"] = t
		}
	}

	if v, ok := d.GetOk("sw_eh_hash"); ok || d.HasChange("sw_eh_hash") {
		t, err := expandObjectSystemNpuSwEhHash(d, v, "sw_eh_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-eh-hash"] = t
		}
	}

	if v, ok := d.GetOk("sw_np_bandwidth"); ok || d.HasChange("sw_np_bandwidth") {
		t, err := expandObjectSystemNpuSwNpBandwidth(d, v, "sw_np_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-np-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("sw_tr_hash"); ok || d.HasChange("sw_tr_hash") {
		t, err := expandObjectSystemNpuSwTrHash(d, v, "sw_tr_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-tr-hash"] = t
		}
	}

	if v, ok := d.GetOk("switch_np_hash"); ok || d.HasChange("switch_np_hash") {
		t, err := expandObjectSystemNpuSwitchNpHash(d, v, "switch_np_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-np-hash"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst_timeout"); ok || d.HasChange("tcp_rst_timeout") {
		t, err := expandObjectSystemNpuTcpRstTimeout(d, v, "tcp_rst_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst-timeout"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_over_vlink"); ok || d.HasChange("tunnel_over_vlink") {
		t, err := expandObjectSystemNpuTunnelOverVlink(d, v, "tunnel_over_vlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-over-vlink"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timeout_profile"); ok || d.HasChange("tcp_timeout_profile") {
		t, err := expandObjectSystemNpuTcpTimeoutProfile(d, v, "tcp_timeout_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timeout-profile"] = t
		}
	}

	if v, ok := d.GetOk("udp_timeout_profile"); ok || d.HasChange("udp_timeout_profile") {
		t, err := expandObjectSystemNpuUdpTimeoutProfile(d, v, "udp_timeout_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-timeout-profile"] = t
		}
	}

	if v, ok := d.GetOk("uesp_offload"); ok || d.HasChange("uesp_offload") {
		t, err := expandObjectSystemNpuUespOffload(d, v, "uesp_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uesp-offload"] = t
		}
	}

	if v, ok := d.GetOk("ull_port_mode"); ok || d.HasChange("ull_port_mode") {
		t, err := expandObjectSystemNpuUllPortMode(d, v, "ull_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ull-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("vlan_lookup_cache"); ok || d.HasChange("vlan_lookup_cache") {
		t, err := expandObjectSystemNpuVlanLookupCache(d, v, "vlan_lookup_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-lookup-cache"] = t
		}
	}

	if v, ok := d.GetOk("vxlan_offload"); ok || d.HasChange("vxlan_offload") {
		t, err := expandObjectSystemNpuVxlanOffload(d, v, "vxlan_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vxlan-offload"] = t
		}
	}

	return &obj, nil
}
