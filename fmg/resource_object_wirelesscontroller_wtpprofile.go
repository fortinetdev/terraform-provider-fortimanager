// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerWtpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWtpProfileCreate,
		Read:   resourceObjectWirelessControllerWtpProfileRead,
		Update: resourceObjectWirelessControllerWtpProfileUpdate,
		Delete: resourceObjectWirelessControllerWtpProfileDelete,

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
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_handoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apcfg_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ble_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"control_message_offload": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"deny_mac_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dtls_in_kernel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"energy_efficient_ethernet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_info_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frequency_handoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"handoff_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"handoff_rssi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"handoff_sta_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_fragment_preventing": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"lan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_esl_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_esl_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port1_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port1_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port2_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port2_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port3_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port3_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port4_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port4_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port5_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port5_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port6_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port6_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port7_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port7_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port8_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port8_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"lbs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aeroscout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_ap_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mmu_report": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mu_factor": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mu_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"aeroscout_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ekahau_blink_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ekahau_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"erc_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"erc_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortipresence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_ble": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_frequency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortipresence_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortipresence_project": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_rogue": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_secret": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"fortipresence_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_unassoc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"station_locate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"led_schedules": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"led_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_passwd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"login_passwd_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_local_platform_str": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ddscan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"poe_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"amsdu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_bufsize": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_chan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_ctl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_beacon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_probe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"beacon_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"call_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"channel_bonding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_utilization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coexistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"darrp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dtim": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"frag_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"iperf_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"powersave_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rts_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"short_guard_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wids_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"radio_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"amsdu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_bufsize": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_chan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_ctl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_beacon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_probe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"beacon_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"call_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"channel_bonding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_utilization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coexistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"darrp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dtim": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"frag_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"iperf_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"powersave_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rts_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"short_guard_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wids_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"radio_3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"amsdu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_bufsize": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_chan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_ctl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_beacon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_probe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"beacon_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"call_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"channel_bonding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_utilization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coexistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"darrp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dtim": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"frag_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"iperf_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"powersave_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rts_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"short_guard_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wids_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"radio_4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"amsdu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_bufsize": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_chan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_ctl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_beacon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_probe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"beacon_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"call_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"channel": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"channel_bonding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_utilization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coexistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"darrp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dtim": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"frag_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"iperf_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_distance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"powersave_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rts_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"short_guard_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_optimize": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vaps": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wids_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"snmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_acl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"split_tunneling_acl_local_ap_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_acl_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tun_mtu_downlink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tun_mtu_uplink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wan_port_mode": &schema.Schema{
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

func resourceObjectWirelessControllerWtpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerWtpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWtpProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerWtpProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWtpProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerWtpProfileRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerWtpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerWtpProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerWtpProfileRead(d, m)
}

func resourceObjectWirelessControllerWtpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerWtpProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWtpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWtpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerWtpProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWtpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWtpProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileApCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileApHandoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileApcfgProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileBleProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileControlMessageOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileDenyMacList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerWtpProfileDenyMacListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerWtpProfile-DenyMacList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectWirelessControllerWtpProfileDenyMacListMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerWtpProfile-DenyMacList-Mac")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerWtpProfileDenyMacListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileDenyMacListMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileDtlsInKernel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileDtlsPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileEnergyEfficientEthernet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileExtInfoEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileFrequencyHandoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileHandoffRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileHandoffRssi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileHandoffStaThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileIpFragmentPreventing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileLan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := i["port-esl-mode"]; ok {
		result["port_esl_mode"] = flattenObjectWirelessControllerWtpProfileLanPortEslMode(i["port-esl-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := i["port-esl-ssid"]; ok {
		result["port_esl_ssid"] = flattenObjectWirelessControllerWtpProfileLanPortEslSsid(i["port-esl-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_mode"
	if _, ok := i["port-mode"]; ok {
		result["port_mode"] = flattenObjectWirelessControllerWtpProfileLanPortMode(i["port-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_ssid"
	if _, ok := i["port-ssid"]; ok {
		result["port_ssid"] = flattenObjectWirelessControllerWtpProfileLanPortSsid(i["port-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_mode"
	if _, ok := i["port1-mode"]; ok {
		result["port1_mode"] = flattenObjectWirelessControllerWtpProfileLanPort1Mode(i["port1-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := i["port1-ssid"]; ok {
		result["port1_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort1Ssid(i["port1-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_mode"
	if _, ok := i["port2-mode"]; ok {
		result["port2_mode"] = flattenObjectWirelessControllerWtpProfileLanPort2Mode(i["port2-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := i["port2-ssid"]; ok {
		result["port2_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort2Ssid(i["port2-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_mode"
	if _, ok := i["port3-mode"]; ok {
		result["port3_mode"] = flattenObjectWirelessControllerWtpProfileLanPort3Mode(i["port3-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := i["port3-ssid"]; ok {
		result["port3_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort3Ssid(i["port3-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_mode"
	if _, ok := i["port4-mode"]; ok {
		result["port4_mode"] = flattenObjectWirelessControllerWtpProfileLanPort4Mode(i["port4-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := i["port4-ssid"]; ok {
		result["port4_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort4Ssid(i["port4-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_mode"
	if _, ok := i["port5-mode"]; ok {
		result["port5_mode"] = flattenObjectWirelessControllerWtpProfileLanPort5Mode(i["port5-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := i["port5-ssid"]; ok {
		result["port5_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort5Ssid(i["port5-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_mode"
	if _, ok := i["port6-mode"]; ok {
		result["port6_mode"] = flattenObjectWirelessControllerWtpProfileLanPort6Mode(i["port6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := i["port6-ssid"]; ok {
		result["port6_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort6Ssid(i["port6-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_mode"
	if _, ok := i["port7-mode"]; ok {
		result["port7_mode"] = flattenObjectWirelessControllerWtpProfileLanPort7Mode(i["port7-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := i["port7-ssid"]; ok {
		result["port7_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort7Ssid(i["port7-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_mode"
	if _, ok := i["port8-mode"]; ok {
		result["port8_mode"] = flattenObjectWirelessControllerWtpProfileLanPort8Mode(i["port8-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := i["port8-ssid"]; ok {
		result["port8_ssid"] = flattenObjectWirelessControllerWtpProfileLanPort8Ssid(i["port8-ssid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfileLanPortEslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPortEslSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPortSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort1Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort2Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort3Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort3Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort4Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort4Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort5Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort5Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort6Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort7Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort7Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort8Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLanPort8Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "aeroscout"
	if _, ok := i["aeroscout"]; ok {
		result["aeroscout"] = flattenObjectWirelessControllerWtpProfileLbsAeroscout(i["aeroscout"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_ap_mac"
	if _, ok := i["aeroscout-ap-mac"]; ok {
		result["aeroscout_ap_mac"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutApMac(i["aeroscout-ap-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mmu_report"
	if _, ok := i["aeroscout-mmu-report"]; ok {
		result["aeroscout_mmu_report"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport(i["aeroscout-mmu-report"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mu"
	if _, ok := i["aeroscout-mu"]; ok {
		result["aeroscout_mu"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutMu(i["aeroscout-mu"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mu_factor"
	if _, ok := i["aeroscout-mu-factor"]; ok {
		result["aeroscout_mu_factor"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor(i["aeroscout-mu-factor"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mu_timeout"
	if _, ok := i["aeroscout-mu-timeout"]; ok {
		result["aeroscout_mu_timeout"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout(i["aeroscout-mu-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_server_ip"
	if _, ok := i["aeroscout-server-ip"]; ok {
		result["aeroscout_server_ip"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerIp(i["aeroscout-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_server_port"
	if _, ok := i["aeroscout-server-port"]; ok {
		result["aeroscout_server_port"] = flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerPort(i["aeroscout-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "ekahau_blink_mode"
	if _, ok := i["ekahau-blink-mode"]; ok {
		result["ekahau_blink_mode"] = flattenObjectWirelessControllerWtpProfileLbsEkahauBlinkMode(i["ekahau-blink-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ekahau_tag"
	if _, ok := i["ekahau-tag"]; ok {
		result["ekahau_tag"] = flattenObjectWirelessControllerWtpProfileLbsEkahauTag(i["ekahau-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "erc_server_ip"
	if _, ok := i["erc-server-ip"]; ok {
		result["erc_server_ip"] = flattenObjectWirelessControllerWtpProfileLbsErcServerIp(i["erc-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "erc_server_port"
	if _, ok := i["erc-server-port"]; ok {
		result["erc_server_port"] = flattenObjectWirelessControllerWtpProfileLbsErcServerPort(i["erc-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence"
	if _, ok := i["fortipresence"]; ok {
		result["fortipresence"] = flattenObjectWirelessControllerWtpProfileLbsFortipresence(i["fortipresence"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_ble"
	if _, ok := i["fortipresence-ble"]; ok {
		result["fortipresence_ble"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceBle(i["fortipresence-ble"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_frequency"
	if _, ok := i["fortipresence-frequency"]; ok {
		result["fortipresence_frequency"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceFrequency(i["fortipresence-frequency"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_port"
	if _, ok := i["fortipresence-port"]; ok {
		result["fortipresence_port"] = flattenObjectWirelessControllerWtpProfileLbsFortipresencePort(i["fortipresence-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_project"
	if _, ok := i["fortipresence-project"]; ok {
		result["fortipresence_project"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceProject(i["fortipresence-project"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_rogue"
	if _, ok := i["fortipresence-rogue"]; ok {
		result["fortipresence_rogue"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceRogue(i["fortipresence-rogue"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_secret"
	if _, ok := i["fortipresence-secret"]; ok {
		result["fortipresence_secret"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceSecret(i["fortipresence-secret"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_server"
	if _, ok := i["fortipresence-server"]; ok {
		result["fortipresence_server"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceServer(i["fortipresence-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_unassoc"
	if _, ok := i["fortipresence-unassoc"]; ok {
		result["fortipresence_unassoc"] = flattenObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc(i["fortipresence-unassoc"], d, pre_append)
	}

	pre_append = pre + ".0." + "station_locate"
	if _, ok := i["station-locate"]; ok {
		result["station_locate"] = flattenObjectWirelessControllerWtpProfileLbsStationLocate(i["station-locate"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutApMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsAeroscoutServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsEkahauBlinkMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsEkahauTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsErcServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsErcServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceBle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresencePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceRogue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLbsStationLocate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLedSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLedState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLldp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileLoginPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatform(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "_local_platform_str"
	if _, ok := i["_local_platform_str"]; ok {
		result["_local_platform_str"] = flattenObjectWirelessControllerWtpProfilePlatformLocalPlatformStr(i["_local_platform_str"], d, pre_append)
	}

	pre_append = pre + ".0." + "ddscan"
	if _, ok := i["ddscan"]; ok {
		result["ddscan"] = flattenObjectWirelessControllerWtpProfilePlatformDdscan(i["ddscan"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectWirelessControllerWtpProfilePlatformMode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenObjectWirelessControllerWtpProfilePlatformType(i["type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfilePlatformLocalPlatformStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatformDdscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatformMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatformType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePoeMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenObjectWirelessControllerWtpProfileRadio1AirtimeFairness(i["airtime-fairness"], d, pre_append)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenObjectWirelessControllerWtpProfileRadio1Amsdu(i["amsdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferChan(i["ap-sniffer-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferData(i["ap-sniffer-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenObjectWirelessControllerWtpProfileRadio1AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenObjectWirelessControllerWtpProfileRadio1AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenObjectWirelessControllerWtpProfileRadio1AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenObjectWirelessControllerWtpProfileRadio1AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenObjectWirelessControllerWtpProfileRadio1Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenObjectWirelessControllerWtpProfileRadio1Band5GType(i["band-5g-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenObjectWirelessControllerWtpProfileRadio1BandwidthCapacity(i["bandwidth-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenObjectWirelessControllerWtpProfileRadio1BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenObjectWirelessControllerWtpProfileRadio1BssColor(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio1CallAdmissionControl(i["call-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenObjectWirelessControllerWtpProfileRadio1CallCapacity(i["call-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenObjectWirelessControllerWtpProfileRadio1Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenObjectWirelessControllerWtpProfileRadio1ChannelBonding(i["channel-bonding"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenObjectWirelessControllerWtpProfileRadio1ChannelUtilization(i["channel-utilization"], d, pre_append)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenObjectWirelessControllerWtpProfileRadio1Coexistence(i["coexistence"], d, pre_append)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenObjectWirelessControllerWtpProfileRadio1Darrp(i["darrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenObjectWirelessControllerWtpProfileRadio1Drma(i["drma"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenObjectWirelessControllerWtpProfileRadio1DrmaSensitivity(i["drma-sensitivity"], d, pre_append)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenObjectWirelessControllerWtpProfileRadio1Dtim(i["dtim"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenObjectWirelessControllerWtpProfileRadio1FragThreshold(i["frag-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenObjectWirelessControllerWtpProfileRadio1IperfProtocol(i["iperf-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenObjectWirelessControllerWtpProfileRadio1IperfServerPort(i["iperf-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenObjectWirelessControllerWtpProfileRadio1MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenObjectWirelessControllerWtpProfileRadio1MaxDistance(i["max-distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectWirelessControllerWtpProfileRadio1Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenObjectWirelessControllerWtpProfileRadio1PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenObjectWirelessControllerWtpProfileRadio1PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenObjectWirelessControllerWtpProfileRadio1PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenObjectWirelessControllerWtpProfileRadio1PowersaveOptimize(i["powersave-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenObjectWirelessControllerWtpProfileRadio1ProtectionMode(i["protection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenObjectWirelessControllerWtpProfileRadio1RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenObjectWirelessControllerWtpProfileRadio1RtsThreshold(i["rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenObjectWirelessControllerWtpProfileRadio1SamBssid(i["sam-bssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenObjectWirelessControllerWtpProfileRadio1SamCaptivePortal(i["sam-captive-portal"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenObjectWirelessControllerWtpProfileRadio1SamPassword(i["sam-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenObjectWirelessControllerWtpProfileRadio1SamReportIntv(i["sam-report-intv"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenObjectWirelessControllerWtpProfileRadio1SamSecurityType(i["sam-security-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_server"
	if _, ok := i["sam-server"]; ok {
		result["sam_server"] = flattenObjectWirelessControllerWtpProfileRadio1SamServer(i["sam-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenObjectWirelessControllerWtpProfileRadio1SamSsid(i["sam-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenObjectWirelessControllerWtpProfileRadio1SamTest(i["sam-test"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenObjectWirelessControllerWtpProfileRadio1SamUsername(i["sam-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenObjectWirelessControllerWtpProfileRadio1ShortGuardInterval(i["short-guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenObjectWirelessControllerWtpProfileRadio1SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenObjectWirelessControllerWtpProfileRadio1TransmitOptimize(i["transmit-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenObjectWirelessControllerWtpProfileRadio1VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenObjectWirelessControllerWtpProfileRadio1Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenObjectWirelessControllerWtpProfileRadio1Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenObjectWirelessControllerWtpProfileRadio1Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenObjectWirelessControllerWtpProfileRadio1Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenObjectWirelessControllerWtpProfileRadio1Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenObjectWirelessControllerWtpProfileRadio1Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenObjectWirelessControllerWtpProfileRadio1Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenObjectWirelessControllerWtpProfileRadio1Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenObjectWirelessControllerWtpProfileRadio1Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenObjectWirelessControllerWtpProfileRadio1WidsProfile(i["wids-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenObjectWirelessControllerWtpProfileRadio1ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfileRadio1AirtimeFairness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Amsdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Band5GType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1BssColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1CallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio1ChannelBonding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ChannelUtilization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Coexistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Darrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Drma(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Dtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1FragThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1IperfProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1IperfServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1MaxDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio1ProtectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1RtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamBssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio1SamReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamSecurityType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SamUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1TransmitOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio1VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1WidsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio1ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenObjectWirelessControllerWtpProfileRadio2AirtimeFairness(i["airtime-fairness"], d, pre_append)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenObjectWirelessControllerWtpProfileRadio2Amsdu(i["amsdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferChan(i["ap-sniffer-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferData(i["ap-sniffer-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenObjectWirelessControllerWtpProfileRadio2AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenObjectWirelessControllerWtpProfileRadio2AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenObjectWirelessControllerWtpProfileRadio2AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenObjectWirelessControllerWtpProfileRadio2AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenObjectWirelessControllerWtpProfileRadio2Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenObjectWirelessControllerWtpProfileRadio2Band5GType(i["band-5g-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenObjectWirelessControllerWtpProfileRadio2BandwidthCapacity(i["bandwidth-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenObjectWirelessControllerWtpProfileRadio2BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenObjectWirelessControllerWtpProfileRadio2BssColor(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio2CallAdmissionControl(i["call-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenObjectWirelessControllerWtpProfileRadio2CallCapacity(i["call-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenObjectWirelessControllerWtpProfileRadio2Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenObjectWirelessControllerWtpProfileRadio2ChannelBonding(i["channel-bonding"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenObjectWirelessControllerWtpProfileRadio2ChannelUtilization(i["channel-utilization"], d, pre_append)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenObjectWirelessControllerWtpProfileRadio2Coexistence(i["coexistence"], d, pre_append)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenObjectWirelessControllerWtpProfileRadio2Darrp(i["darrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenObjectWirelessControllerWtpProfileRadio2Drma(i["drma"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenObjectWirelessControllerWtpProfileRadio2DrmaSensitivity(i["drma-sensitivity"], d, pre_append)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenObjectWirelessControllerWtpProfileRadio2Dtim(i["dtim"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenObjectWirelessControllerWtpProfileRadio2FragThreshold(i["frag-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenObjectWirelessControllerWtpProfileRadio2IperfProtocol(i["iperf-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenObjectWirelessControllerWtpProfileRadio2IperfServerPort(i["iperf-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenObjectWirelessControllerWtpProfileRadio2MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenObjectWirelessControllerWtpProfileRadio2MaxDistance(i["max-distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectWirelessControllerWtpProfileRadio2Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenObjectWirelessControllerWtpProfileRadio2PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenObjectWirelessControllerWtpProfileRadio2PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenObjectWirelessControllerWtpProfileRadio2PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenObjectWirelessControllerWtpProfileRadio2PowersaveOptimize(i["powersave-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenObjectWirelessControllerWtpProfileRadio2ProtectionMode(i["protection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenObjectWirelessControllerWtpProfileRadio2RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenObjectWirelessControllerWtpProfileRadio2RtsThreshold(i["rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenObjectWirelessControllerWtpProfileRadio2SamBssid(i["sam-bssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenObjectWirelessControllerWtpProfileRadio2SamCaptivePortal(i["sam-captive-portal"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenObjectWirelessControllerWtpProfileRadio2SamPassword(i["sam-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenObjectWirelessControllerWtpProfileRadio2SamReportIntv(i["sam-report-intv"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenObjectWirelessControllerWtpProfileRadio2SamSecurityType(i["sam-security-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_server"
	if _, ok := i["sam-server"]; ok {
		result["sam_server"] = flattenObjectWirelessControllerWtpProfileRadio2SamServer(i["sam-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenObjectWirelessControllerWtpProfileRadio2SamSsid(i["sam-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenObjectWirelessControllerWtpProfileRadio2SamTest(i["sam-test"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenObjectWirelessControllerWtpProfileRadio2SamUsername(i["sam-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenObjectWirelessControllerWtpProfileRadio2ShortGuardInterval(i["short-guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenObjectWirelessControllerWtpProfileRadio2SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenObjectWirelessControllerWtpProfileRadio2TransmitOptimize(i["transmit-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenObjectWirelessControllerWtpProfileRadio2VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenObjectWirelessControllerWtpProfileRadio2Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenObjectWirelessControllerWtpProfileRadio2Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenObjectWirelessControllerWtpProfileRadio2Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenObjectWirelessControllerWtpProfileRadio2Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenObjectWirelessControllerWtpProfileRadio2Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenObjectWirelessControllerWtpProfileRadio2Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenObjectWirelessControllerWtpProfileRadio2Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenObjectWirelessControllerWtpProfileRadio2Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenObjectWirelessControllerWtpProfileRadio2Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenObjectWirelessControllerWtpProfileRadio2WidsProfile(i["wids-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenObjectWirelessControllerWtpProfileRadio2ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfileRadio2AirtimeFairness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Amsdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Band5GType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2BssColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2CallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio2ChannelBonding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ChannelUtilization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Coexistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Darrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Drma(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Dtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2FragThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2IperfProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2IperfServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2MaxDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio2ProtectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2RtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamBssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio2SamReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamSecurityType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SamUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2TransmitOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio2VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2WidsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio2ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenObjectWirelessControllerWtpProfileRadio3AirtimeFairness(i["airtime-fairness"], d, pre_append)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenObjectWirelessControllerWtpProfileRadio3Amsdu(i["amsdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferChan(i["ap-sniffer-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferData(i["ap-sniffer-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenObjectWirelessControllerWtpProfileRadio3AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenObjectWirelessControllerWtpProfileRadio3AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenObjectWirelessControllerWtpProfileRadio3AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenObjectWirelessControllerWtpProfileRadio3AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenObjectWirelessControllerWtpProfileRadio3Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenObjectWirelessControllerWtpProfileRadio3Band5GType(i["band-5g-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenObjectWirelessControllerWtpProfileRadio3BandwidthCapacity(i["bandwidth-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenObjectWirelessControllerWtpProfileRadio3BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenObjectWirelessControllerWtpProfileRadio3BssColor(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio3CallAdmissionControl(i["call-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenObjectWirelessControllerWtpProfileRadio3CallCapacity(i["call-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenObjectWirelessControllerWtpProfileRadio3Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenObjectWirelessControllerWtpProfileRadio3ChannelBonding(i["channel-bonding"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenObjectWirelessControllerWtpProfileRadio3ChannelUtilization(i["channel-utilization"], d, pre_append)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenObjectWirelessControllerWtpProfileRadio3Coexistence(i["coexistence"], d, pre_append)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenObjectWirelessControllerWtpProfileRadio3Darrp(i["darrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenObjectWirelessControllerWtpProfileRadio3Drma(i["drma"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenObjectWirelessControllerWtpProfileRadio3DrmaSensitivity(i["drma-sensitivity"], d, pre_append)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenObjectWirelessControllerWtpProfileRadio3Dtim(i["dtim"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenObjectWirelessControllerWtpProfileRadio3FragThreshold(i["frag-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenObjectWirelessControllerWtpProfileRadio3IperfProtocol(i["iperf-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenObjectWirelessControllerWtpProfileRadio3IperfServerPort(i["iperf-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenObjectWirelessControllerWtpProfileRadio3MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenObjectWirelessControllerWtpProfileRadio3MaxDistance(i["max-distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectWirelessControllerWtpProfileRadio3Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenObjectWirelessControllerWtpProfileRadio3PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenObjectWirelessControllerWtpProfileRadio3PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenObjectWirelessControllerWtpProfileRadio3PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenObjectWirelessControllerWtpProfileRadio3PowersaveOptimize(i["powersave-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenObjectWirelessControllerWtpProfileRadio3ProtectionMode(i["protection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenObjectWirelessControllerWtpProfileRadio3RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenObjectWirelessControllerWtpProfileRadio3RtsThreshold(i["rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenObjectWirelessControllerWtpProfileRadio3SamBssid(i["sam-bssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenObjectWirelessControllerWtpProfileRadio3SamCaptivePortal(i["sam-captive-portal"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenObjectWirelessControllerWtpProfileRadio3SamPassword(i["sam-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenObjectWirelessControllerWtpProfileRadio3SamReportIntv(i["sam-report-intv"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenObjectWirelessControllerWtpProfileRadio3SamSecurityType(i["sam-security-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_server"
	if _, ok := i["sam-server"]; ok {
		result["sam_server"] = flattenObjectWirelessControllerWtpProfileRadio3SamServer(i["sam-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenObjectWirelessControllerWtpProfileRadio3SamSsid(i["sam-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenObjectWirelessControllerWtpProfileRadio3SamTest(i["sam-test"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenObjectWirelessControllerWtpProfileRadio3SamUsername(i["sam-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenObjectWirelessControllerWtpProfileRadio3ShortGuardInterval(i["short-guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenObjectWirelessControllerWtpProfileRadio3SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenObjectWirelessControllerWtpProfileRadio3TransmitOptimize(i["transmit-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenObjectWirelessControllerWtpProfileRadio3VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenObjectWirelessControllerWtpProfileRadio3Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenObjectWirelessControllerWtpProfileRadio3Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenObjectWirelessControllerWtpProfileRadio3Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenObjectWirelessControllerWtpProfileRadio3Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenObjectWirelessControllerWtpProfileRadio3Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenObjectWirelessControllerWtpProfileRadio3Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenObjectWirelessControllerWtpProfileRadio3Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenObjectWirelessControllerWtpProfileRadio3Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenObjectWirelessControllerWtpProfileRadio3Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenObjectWirelessControllerWtpProfileRadio3WidsProfile(i["wids-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenObjectWirelessControllerWtpProfileRadio3ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfileRadio3AirtimeFairness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Amsdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Band5GType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3BssColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3CallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio3ChannelBonding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ChannelUtilization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Coexistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Darrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Drma(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Dtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3FragThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3IperfProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3IperfServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3MaxDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio3ProtectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3RtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamBssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio3SamReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamSecurityType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SamUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3TransmitOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio3VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3WidsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio3ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenObjectWirelessControllerWtpProfileRadio4AirtimeFairness(i["airtime-fairness"], d, pre_append)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenObjectWirelessControllerWtpProfileRadio4Amsdu(i["amsdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferChan(i["ap-sniffer-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferData(i["ap-sniffer-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenObjectWirelessControllerWtpProfileRadio4AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenObjectWirelessControllerWtpProfileRadio4AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenObjectWirelessControllerWtpProfileRadio4AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenObjectWirelessControllerWtpProfileRadio4AutoPowerTarget(i["auto-power-target"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenObjectWirelessControllerWtpProfileRadio4Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenObjectWirelessControllerWtpProfileRadio4Band5GType(i["band-5g-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenObjectWirelessControllerWtpProfileRadio4BandwidthCapacity(i["bandwidth-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenObjectWirelessControllerWtpProfileRadio4BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenObjectWirelessControllerWtpProfileRadio4BssColor(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenObjectWirelessControllerWtpProfileRadio4CallAdmissionControl(i["call-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenObjectWirelessControllerWtpProfileRadio4CallCapacity(i["call-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenObjectWirelessControllerWtpProfileRadio4Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenObjectWirelessControllerWtpProfileRadio4ChannelBonding(i["channel-bonding"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenObjectWirelessControllerWtpProfileRadio4ChannelUtilization(i["channel-utilization"], d, pre_append)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenObjectWirelessControllerWtpProfileRadio4Coexistence(i["coexistence"], d, pre_append)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenObjectWirelessControllerWtpProfileRadio4Darrp(i["darrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenObjectWirelessControllerWtpProfileRadio4Drma(i["drma"], d, pre_append)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenObjectWirelessControllerWtpProfileRadio4DrmaSensitivity(i["drma-sensitivity"], d, pre_append)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenObjectWirelessControllerWtpProfileRadio4Dtim(i["dtim"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenObjectWirelessControllerWtpProfileRadio4FragThreshold(i["frag-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenObjectWirelessControllerWtpProfileRadio4IperfProtocol(i["iperf-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenObjectWirelessControllerWtpProfileRadio4IperfServerPort(i["iperf-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenObjectWirelessControllerWtpProfileRadio4MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenObjectWirelessControllerWtpProfileRadio4MaxDistance(i["max-distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectWirelessControllerWtpProfileRadio4Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenObjectWirelessControllerWtpProfileRadio4PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenObjectWirelessControllerWtpProfileRadio4PowerMode(i["power-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenObjectWirelessControllerWtpProfileRadio4PowerValue(i["power-value"], d, pre_append)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenObjectWirelessControllerWtpProfileRadio4PowersaveOptimize(i["powersave-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenObjectWirelessControllerWtpProfileRadio4ProtectionMode(i["protection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenObjectWirelessControllerWtpProfileRadio4RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenObjectWirelessControllerWtpProfileRadio4RtsThreshold(i["rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenObjectWirelessControllerWtpProfileRadio4SamBssid(i["sam-bssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenObjectWirelessControllerWtpProfileRadio4SamCaptivePortal(i["sam-captive-portal"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenObjectWirelessControllerWtpProfileRadio4SamPassword(i["sam-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenObjectWirelessControllerWtpProfileRadio4SamReportIntv(i["sam-report-intv"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenObjectWirelessControllerWtpProfileRadio4SamSecurityType(i["sam-security-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_server"
	if _, ok := i["sam-server"]; ok {
		result["sam_server"] = flattenObjectWirelessControllerWtpProfileRadio4SamServer(i["sam-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenObjectWirelessControllerWtpProfileRadio4SamSsid(i["sam-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenObjectWirelessControllerWtpProfileRadio4SamTest(i["sam-test"], d, pre_append)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenObjectWirelessControllerWtpProfileRadio4SamUsername(i["sam-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenObjectWirelessControllerWtpProfileRadio4ShortGuardInterval(i["short-guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenObjectWirelessControllerWtpProfileRadio4SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenObjectWirelessControllerWtpProfileRadio4TransmitOptimize(i["transmit-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenObjectWirelessControllerWtpProfileRadio4VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap1"
	if _, ok := i["vap1"]; ok {
		result["vap1"] = flattenObjectWirelessControllerWtpProfileRadio4Vap1(i["vap1"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap2"
	if _, ok := i["vap2"]; ok {
		result["vap2"] = flattenObjectWirelessControllerWtpProfileRadio4Vap2(i["vap2"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap3"
	if _, ok := i["vap3"]; ok {
		result["vap3"] = flattenObjectWirelessControllerWtpProfileRadio4Vap3(i["vap3"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap4"
	if _, ok := i["vap4"]; ok {
		result["vap4"] = flattenObjectWirelessControllerWtpProfileRadio4Vap4(i["vap4"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap5"
	if _, ok := i["vap5"]; ok {
		result["vap5"] = flattenObjectWirelessControllerWtpProfileRadio4Vap5(i["vap5"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap6"
	if _, ok := i["vap6"]; ok {
		result["vap6"] = flattenObjectWirelessControllerWtpProfileRadio4Vap6(i["vap6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap7"
	if _, ok := i["vap7"]; ok {
		result["vap7"] = flattenObjectWirelessControllerWtpProfileRadio4Vap7(i["vap7"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap8"
	if _, ok := i["vap8"]; ok {
		result["vap8"] = flattenObjectWirelessControllerWtpProfileRadio4Vap8(i["vap8"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenObjectWirelessControllerWtpProfileRadio4Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenObjectWirelessControllerWtpProfileRadio4WidsProfile(i["wids-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenObjectWirelessControllerWtpProfileRadio4ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectWirelessControllerWtpProfileRadio4AirtimeFairness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Amsdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Band5GType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4BssColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4CallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Channel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio4ChannelBonding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ChannelUtilization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Coexistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Darrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Drma(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Dtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4FragThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4IperfProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4IperfServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4MaxDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4PowerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4PowerValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio4ProtectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4RtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamBssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio4SamReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamSecurityType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SamUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4TransmitOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWtpProfileRadio4VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vap8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4Vaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4WidsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileRadio4ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileSnmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAcl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := i["dest-ip"]; ok {
			v := flattenObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp(i["dest-ip"], d, pre_append)
			tmp["dest_ip"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerWtpProfile-SplitTunnelingAcl-DestIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWirelessControllerWtpProfileSplitTunnelingAclId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerWtpProfile-SplitTunnelingAcl-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileSplitTunnelingAclPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileTunMtuDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileTunMtuUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfileWanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWtpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allowaccess", flattenObjectWirelessControllerWtpProfileAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ObjectWirelessControllerWtpProfile-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("ap_country", flattenObjectWirelessControllerWtpProfileApCountry(o["ap-country"], d, "ap_country")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-country"], "ObjectWirelessControllerWtpProfile-ApCountry"); ok {
			if err = d.Set("ap_country", vv); err != nil {
				return fmt.Errorf("Error reading ap_country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_country: %v", err)
		}
	}

	if err = d.Set("ap_handoff", flattenObjectWirelessControllerWtpProfileApHandoff(o["ap-handoff"], d, "ap_handoff")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-handoff"], "ObjectWirelessControllerWtpProfile-ApHandoff"); ok {
			if err = d.Set("ap_handoff", vv); err != nil {
				return fmt.Errorf("Error reading ap_handoff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_handoff: %v", err)
		}
	}

	if err = d.Set("apcfg_profile", flattenObjectWirelessControllerWtpProfileApcfgProfile(o["apcfg-profile"], d, "apcfg_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["apcfg-profile"], "ObjectWirelessControllerWtpProfile-ApcfgProfile"); ok {
			if err = d.Set("apcfg_profile", vv); err != nil {
				return fmt.Errorf("Error reading apcfg_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apcfg_profile: %v", err)
		}
	}

	if err = d.Set("ble_profile", flattenObjectWirelessControllerWtpProfileBleProfile(o["ble-profile"], d, "ble_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-profile"], "ObjectWirelessControllerWtpProfile-BleProfile"); ok {
			if err = d.Set("ble_profile", vv); err != nil {
				return fmt.Errorf("Error reading ble_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_profile: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWirelessControllerWtpProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerWtpProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("control_message_offload", flattenObjectWirelessControllerWtpProfileControlMessageOffload(o["control-message-offload"], d, "control_message_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["control-message-offload"], "ObjectWirelessControllerWtpProfile-ControlMessageOffload"); ok {
			if err = d.Set("control_message_offload", vv); err != nil {
				return fmt.Errorf("Error reading control_message_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading control_message_offload: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("deny_mac_list", flattenObjectWirelessControllerWtpProfileDenyMacList(o["deny-mac-list"], d, "deny_mac_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["deny-mac-list"], "ObjectWirelessControllerWtpProfile-DenyMacList"); ok {
				if err = d.Set("deny_mac_list", vv); err != nil {
					return fmt.Errorf("Error reading deny_mac_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading deny_mac_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("deny_mac_list"); ok {
			if err = d.Set("deny_mac_list", flattenObjectWirelessControllerWtpProfileDenyMacList(o["deny-mac-list"], d, "deny_mac_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["deny-mac-list"], "ObjectWirelessControllerWtpProfile-DenyMacList"); ok {
					if err = d.Set("deny_mac_list", vv); err != nil {
						return fmt.Errorf("Error reading deny_mac_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading deny_mac_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("dtls_in_kernel", flattenObjectWirelessControllerWtpProfileDtlsInKernel(o["dtls-in-kernel"], d, "dtls_in_kernel")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-in-kernel"], "ObjectWirelessControllerWtpProfile-DtlsInKernel"); ok {
			if err = d.Set("dtls_in_kernel", vv); err != nil {
				return fmt.Errorf("Error reading dtls_in_kernel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_in_kernel: %v", err)
		}
	}

	if err = d.Set("dtls_policy", flattenObjectWirelessControllerWtpProfileDtlsPolicy(o["dtls-policy"], d, "dtls_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["dtls-policy"], "ObjectWirelessControllerWtpProfile-DtlsPolicy"); ok {
			if err = d.Set("dtls_policy", vv); err != nil {
				return fmt.Errorf("Error reading dtls_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dtls_policy: %v", err)
		}
	}

	if err = d.Set("energy_efficient_ethernet", flattenObjectWirelessControllerWtpProfileEnergyEfficientEthernet(o["energy-efficient-ethernet"], d, "energy_efficient_ethernet")); err != nil {
		if vv, ok := fortiAPIPatch(o["energy-efficient-ethernet"], "ObjectWirelessControllerWtpProfile-EnergyEfficientEthernet"); ok {
			if err = d.Set("energy_efficient_ethernet", vv); err != nil {
				return fmt.Errorf("Error reading energy_efficient_ethernet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading energy_efficient_ethernet: %v", err)
		}
	}

	if err = d.Set("ext_info_enable", flattenObjectWirelessControllerWtpProfileExtInfoEnable(o["ext-info-enable"], d, "ext_info_enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-info-enable"], "ObjectWirelessControllerWtpProfile-ExtInfoEnable"); ok {
			if err = d.Set("ext_info_enable", vv); err != nil {
				return fmt.Errorf("Error reading ext_info_enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_info_enable: %v", err)
		}
	}

	if err = d.Set("frequency_handoff", flattenObjectWirelessControllerWtpProfileFrequencyHandoff(o["frequency-handoff"], d, "frequency_handoff")); err != nil {
		if vv, ok := fortiAPIPatch(o["frequency-handoff"], "ObjectWirelessControllerWtpProfile-FrequencyHandoff"); ok {
			if err = d.Set("frequency_handoff", vv); err != nil {
				return fmt.Errorf("Error reading frequency_handoff: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frequency_handoff: %v", err)
		}
	}

	if err = d.Set("handoff_roaming", flattenObjectWirelessControllerWtpProfileHandoffRoaming(o["handoff-roaming"], d, "handoff_roaming")); err != nil {
		if vv, ok := fortiAPIPatch(o["handoff-roaming"], "ObjectWirelessControllerWtpProfile-HandoffRoaming"); ok {
			if err = d.Set("handoff_roaming", vv); err != nil {
				return fmt.Errorf("Error reading handoff_roaming: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading handoff_roaming: %v", err)
		}
	}

	if err = d.Set("handoff_rssi", flattenObjectWirelessControllerWtpProfileHandoffRssi(o["handoff-rssi"], d, "handoff_rssi")); err != nil {
		if vv, ok := fortiAPIPatch(o["handoff-rssi"], "ObjectWirelessControllerWtpProfile-HandoffRssi"); ok {
			if err = d.Set("handoff_rssi", vv); err != nil {
				return fmt.Errorf("Error reading handoff_rssi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading handoff_rssi: %v", err)
		}
	}

	if err = d.Set("handoff_sta_thresh", flattenObjectWirelessControllerWtpProfileHandoffStaThresh(o["handoff-sta-thresh"], d, "handoff_sta_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["handoff-sta-thresh"], "ObjectWirelessControllerWtpProfile-HandoffStaThresh"); ok {
			if err = d.Set("handoff_sta_thresh", vv); err != nil {
				return fmt.Errorf("Error reading handoff_sta_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading handoff_sta_thresh: %v", err)
		}
	}

	if err = d.Set("ip_fragment_preventing", flattenObjectWirelessControllerWtpProfileIpFragmentPreventing(o["ip-fragment-preventing"], d, "ip_fragment_preventing")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-fragment-preventing"], "ObjectWirelessControllerWtpProfile-IpFragmentPreventing"); ok {
			if err = d.Set("ip_fragment_preventing", vv); err != nil {
				return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan", flattenObjectWirelessControllerWtpProfileLan(o["lan"], d, "lan")); err != nil {
			if vv, ok := fortiAPIPatch(o["lan"], "ObjectWirelessControllerWtpProfile-Lan"); ok {
				if err = d.Set("lan", vv); err != nil {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan"); ok {
			if err = d.Set("lan", flattenObjectWirelessControllerWtpProfileLan(o["lan"], d, "lan")); err != nil {
				if vv, ok := fortiAPIPatch(o["lan"], "ObjectWirelessControllerWtpProfile-Lan"); ok {
					if err = d.Set("lan", vv); err != nil {
						return fmt.Errorf("Error reading lan: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("lbs", flattenObjectWirelessControllerWtpProfileLbs(o["lbs"], d, "lbs")); err != nil {
			if vv, ok := fortiAPIPatch(o["lbs"], "ObjectWirelessControllerWtpProfile-Lbs"); ok {
				if err = d.Set("lbs", vv); err != nil {
					return fmt.Errorf("Error reading lbs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lbs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lbs"); ok {
			if err = d.Set("lbs", flattenObjectWirelessControllerWtpProfileLbs(o["lbs"], d, "lbs")); err != nil {
				if vv, ok := fortiAPIPatch(o["lbs"], "ObjectWirelessControllerWtpProfile-Lbs"); ok {
					if err = d.Set("lbs", vv); err != nil {
						return fmt.Errorf("Error reading lbs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lbs: %v", err)
				}
			}
		}
	}

	if err = d.Set("led_schedules", flattenObjectWirelessControllerWtpProfileLedSchedules(o["led-schedules"], d, "led_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["led-schedules"], "ObjectWirelessControllerWtpProfile-LedSchedules"); ok {
			if err = d.Set("led_schedules", vv); err != nil {
				return fmt.Errorf("Error reading led_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading led_schedules: %v", err)
		}
	}

	if err = d.Set("led_state", flattenObjectWirelessControllerWtpProfileLedState(o["led-state"], d, "led_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["led-state"], "ObjectWirelessControllerWtpProfile-LedState"); ok {
			if err = d.Set("led_state", vv); err != nil {
				return fmt.Errorf("Error reading led_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading led_state: %v", err)
		}
	}

	if err = d.Set("lldp", flattenObjectWirelessControllerWtpProfileLldp(o["lldp"], d, "lldp")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp"], "ObjectWirelessControllerWtpProfile-Lldp"); ok {
			if err = d.Set("lldp", vv); err != nil {
				return fmt.Errorf("Error reading lldp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp: %v", err)
		}
	}

	if err = d.Set("login_passwd", flattenObjectWirelessControllerWtpProfileLoginPasswd(o["login-passwd"], d, "login_passwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-passwd"], "ObjectWirelessControllerWtpProfile-LoginPasswd"); ok {
			if err = d.Set("login_passwd", vv); err != nil {
				return fmt.Errorf("Error reading login_passwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_passwd: %v", err)
		}
	}

	if err = d.Set("login_passwd_change", flattenObjectWirelessControllerWtpProfileLoginPasswdChange(o["login-passwd-change"], d, "login_passwd_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-passwd-change"], "ObjectWirelessControllerWtpProfile-LoginPasswdChange"); ok {
			if err = d.Set("login_passwd_change", vv); err != nil {
				return fmt.Errorf("Error reading login_passwd_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_passwd_change: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenObjectWirelessControllerWtpProfileMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "ObjectWirelessControllerWtpProfile-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerWtpProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerWtpProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("platform", flattenObjectWirelessControllerWtpProfilePlatform(o["platform"], d, "platform")); err != nil {
			if vv, ok := fortiAPIPatch(o["platform"], "ObjectWirelessControllerWtpProfile-Platform"); ok {
				if err = d.Set("platform", vv); err != nil {
					return fmt.Errorf("Error reading platform: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading platform: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("platform"); ok {
			if err = d.Set("platform", flattenObjectWirelessControllerWtpProfilePlatform(o["platform"], d, "platform")); err != nil {
				if vv, ok := fortiAPIPatch(o["platform"], "ObjectWirelessControllerWtpProfile-Platform"); ok {
					if err = d.Set("platform", vv); err != nil {
						return fmt.Errorf("Error reading platform: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading platform: %v", err)
				}
			}
		}
	}

	if err = d.Set("poe_mode", flattenObjectWirelessControllerWtpProfilePoeMode(o["poe-mode"], d, "poe_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-mode"], "ObjectWirelessControllerWtpProfile-PoeMode"); ok {
			if err = d.Set("poe_mode", vv); err != nil {
				return fmt.Errorf("Error reading poe_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radio_1", flattenObjectWirelessControllerWtpProfileRadio1(o["radio-1"], d, "radio_1")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-1"], "ObjectWirelessControllerWtpProfile-Radio1"); ok {
				if err = d.Set("radio_1", vv); err != nil {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenObjectWirelessControllerWtpProfileRadio1(o["radio-1"], d, "radio_1")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-1"], "ObjectWirelessControllerWtpProfile-Radio1"); ok {
					if err = d.Set("radio_1", vv); err != nil {
						return fmt.Errorf("Error reading radio_1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_2", flattenObjectWirelessControllerWtpProfileRadio2(o["radio-2"], d, "radio_2")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-2"], "ObjectWirelessControllerWtpProfile-Radio2"); ok {
				if err = d.Set("radio_2", vv); err != nil {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenObjectWirelessControllerWtpProfileRadio2(o["radio-2"], d, "radio_2")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-2"], "ObjectWirelessControllerWtpProfile-Radio2"); ok {
					if err = d.Set("radio_2", vv); err != nil {
						return fmt.Errorf("Error reading radio_2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_3", flattenObjectWirelessControllerWtpProfileRadio3(o["radio-3"], d, "radio_3")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-3"], "ObjectWirelessControllerWtpProfile-Radio3"); ok {
				if err = d.Set("radio_3", vv); err != nil {
					return fmt.Errorf("Error reading radio_3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_3"); ok {
			if err = d.Set("radio_3", flattenObjectWirelessControllerWtpProfileRadio3(o["radio-3"], d, "radio_3")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-3"], "ObjectWirelessControllerWtpProfile-Radio3"); ok {
					if err = d.Set("radio_3", vv); err != nil {
						return fmt.Errorf("Error reading radio_3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_3: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_4", flattenObjectWirelessControllerWtpProfileRadio4(o["radio-4"], d, "radio_4")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-4"], "ObjectWirelessControllerWtpProfile-Radio4"); ok {
				if err = d.Set("radio_4", vv); err != nil {
					return fmt.Errorf("Error reading radio_4: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_4"); ok {
			if err = d.Set("radio_4", flattenObjectWirelessControllerWtpProfileRadio4(o["radio-4"], d, "radio_4")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-4"], "ObjectWirelessControllerWtpProfile-Radio4"); ok {
					if err = d.Set("radio_4", vv); err != nil {
						return fmt.Errorf("Error reading radio_4: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_4: %v", err)
				}
			}
		}
	}

	if err = d.Set("snmp", flattenObjectWirelessControllerWtpProfileSnmp(o["snmp"], d, "snmp")); err != nil {
		if vv, ok := fortiAPIPatch(o["snmp"], "ObjectWirelessControllerWtpProfile-Snmp"); ok {
			if err = d.Set("snmp", vv); err != nil {
				return fmt.Errorf("Error reading snmp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading snmp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_tunneling_acl", flattenObjectWirelessControllerWtpProfileSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
			if vv, ok := fortiAPIPatch(o["split-tunneling-acl"], "ObjectWirelessControllerWtpProfile-SplitTunnelingAcl"); ok {
				if err = d.Set("split_tunneling_acl", vv); err != nil {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_acl"); ok {
			if err = d.Set("split_tunneling_acl", flattenObjectWirelessControllerWtpProfileSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
				if vv, ok := fortiAPIPatch(o["split-tunneling-acl"], "ObjectWirelessControllerWtpProfile-SplitTunnelingAcl"); ok {
					if err = d.Set("split_tunneling_acl", vv); err != nil {
						return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			}
		}
	}

	if err = d.Set("split_tunneling_acl_local_ap_subnet", flattenObjectWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(o["split-tunneling-acl-local-ap-subnet"], d, "split_tunneling_acl_local_ap_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-acl-local-ap-subnet"], "ObjectWirelessControllerWtpProfile-SplitTunnelingAclLocalApSubnet"); ok {
			if err = d.Set("split_tunneling_acl_local_ap_subnet", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_path", flattenObjectWirelessControllerWtpProfileSplitTunnelingAclPath(o["split-tunneling-acl-path"], d, "split_tunneling_acl_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-acl-path"], "ObjectWirelessControllerWtpProfile-SplitTunnelingAclPath"); ok {
			if err = d.Set("split_tunneling_acl_path", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
		}
	}

	if err = d.Set("tun_mtu_downlink", flattenObjectWirelessControllerWtpProfileTunMtuDownlink(o["tun-mtu-downlink"], d, "tun_mtu_downlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["tun-mtu-downlink"], "ObjectWirelessControllerWtpProfile-TunMtuDownlink"); ok {
			if err = d.Set("tun_mtu_downlink", vv); err != nil {
				return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
		}
	}

	if err = d.Set("tun_mtu_uplink", flattenObjectWirelessControllerWtpProfileTunMtuUplink(o["tun-mtu-uplink"], d, "tun_mtu_uplink")); err != nil {
		if vv, ok := fortiAPIPatch(o["tun-mtu-uplink"], "ObjectWirelessControllerWtpProfile-TunMtuUplink"); ok {
			if err = d.Set("tun_mtu_uplink", vv); err != nil {
				return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
		}
	}

	if err = d.Set("wan_port_mode", flattenObjectWirelessControllerWtpProfileWanPortMode(o["wan-port-mode"], d, "wan_port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wan-port-mode"], "ObjectWirelessControllerWtpProfile-WanPortMode"); ok {
			if err = d.Set("wan_port_mode", vv); err != nil {
				return fmt.Errorf("Error reading wan_port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wan_port_mode: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWtpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWtpProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileApCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileApHandoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileApcfgProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileBleProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileControlMessageOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileDenyMacList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectWirelessControllerWtpProfileDenyMacListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandObjectWirelessControllerWtpProfileDenyMacListMac(d, i["mac"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileDenyMacListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileDenyMacListMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileDtlsInKernel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileDtlsPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileEnergyEfficientEthernet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileExtInfoEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileFrequencyHandoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileHandoffRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileHandoffRssi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileHandoffStaThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileIpFragmentPreventing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-esl-mode"], _ = expandObjectWirelessControllerWtpProfileLanPortEslMode(d, i["port_esl_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-esl-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPortEslSsid(d, i["port_esl_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-mode"], _ = expandObjectWirelessControllerWtpProfileLanPortMode(d, i["port_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPortSsid(d, i["port_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port1_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort1Mode(d, i["port1_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort1Ssid(d, i["port1_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port2_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort2Mode(d, i["port2_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort2Ssid(d, i["port2_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port3_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort3Mode(d, i["port3_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort3Ssid(d, i["port3_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port4_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort4Mode(d, i["port4_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort4Ssid(d, i["port4_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port5_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort5Mode(d, i["port5_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort5Ssid(d, i["port5_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port6_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort6Mode(d, i["port6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort6Ssid(d, i["port6_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port7_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort7Mode(d, i["port7_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort7Ssid(d, i["port7_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port8_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-mode"], _ = expandObjectWirelessControllerWtpProfileLanPort8Mode(d, i["port8_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-ssid"], _ = expandObjectWirelessControllerWtpProfileLanPort8Ssid(d, i["port8_ssid"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileLanPortEslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPortEslSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPortSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort1Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort2Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort3Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort3Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort4Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort4Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort5Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort5Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort6Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort7Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort7Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort8Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLanPort8Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "aeroscout"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscout(d, i["aeroscout"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_ap_mac"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-ap-mac"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutApMac(d, i["aeroscout_ap_mac"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mmu_report"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mmu-report"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport(d, i["aeroscout_mmu_report"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mu"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutMu(d, i["aeroscout_mu"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mu_factor"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu-factor"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor(d, i["aeroscout_mu_factor"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mu_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu-timeout"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout(d, i["aeroscout_mu_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-server-ip"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutServerIp(d, i["aeroscout_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-server-port"], _ = expandObjectWirelessControllerWtpProfileLbsAeroscoutServerPort(d, i["aeroscout_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "ekahau_blink_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ekahau-blink-mode"], _ = expandObjectWirelessControllerWtpProfileLbsEkahauBlinkMode(d, i["ekahau_blink_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ekahau_tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ekahau-tag"], _ = expandObjectWirelessControllerWtpProfileLbsEkahauTag(d, i["ekahau_tag"], pre_append)
	}
	pre_append = pre + ".0." + "erc_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["erc-server-ip"], _ = expandObjectWirelessControllerWtpProfileLbsErcServerIp(d, i["erc_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "erc_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["erc-server-port"], _ = expandObjectWirelessControllerWtpProfileLbsErcServerPort(d, i["erc_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresence(d, i["fortipresence"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_ble"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-ble"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceBle(d, i["fortipresence_ble"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_frequency"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-frequency"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceFrequency(d, i["fortipresence_frequency"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-port"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresencePort(d, i["fortipresence_port"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_project"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-project"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceProject(d, i["fortipresence_project"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_rogue"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-rogue"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceRogue(d, i["fortipresence_rogue"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_secret"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-secret"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceSecret(d, i["fortipresence_secret"], pre_append)
	} else {
		result["fortipresence-secret"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "fortipresence_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-server"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceServer(d, i["fortipresence_server"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_unassoc"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-unassoc"], _ = expandObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc(d, i["fortipresence_unassoc"], pre_append)
	}
	pre_append = pre + ".0." + "station_locate"
	if _, ok := d.GetOk(pre_append); ok {
		result["station-locate"], _ = expandObjectWirelessControllerWtpProfileLbsStationLocate(d, i["station_locate"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutApMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMmuReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMuFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutMuTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsAeroscoutServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsEkahauBlinkMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsEkahauTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsErcServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsErcServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceBle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresencePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceRogue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsFortipresenceUnassoc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLbsStationLocate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLedSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLedState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLldp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileLoginPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "_local_platform_str"
	if _, ok := d.GetOk(pre_append); ok {
		result["_local_platform_str"], _ = expandObjectWirelessControllerWtpProfilePlatformLocalPlatformStr(d, i["_local_platform_str"], pre_append)
	}
	pre_append = pre + ".0." + "ddscan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ddscan"], _ = expandObjectWirelessControllerWtpProfilePlatformDdscan(d, i["ddscan"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandObjectWirelessControllerWtpProfilePlatformMode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok {
		result["type"], _ = expandObjectWirelessControllerWtpProfilePlatformType(d, i["type"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfilePlatformLocalPlatformStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatformDdscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatformMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatformType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePoeMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandObjectWirelessControllerWtpProfileRadio1AirtimeFairness(d, i["airtime_fairness"], pre_append)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandObjectWirelessControllerWtpProfileRadio1Amsdu(d, i["amsdu"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferChan(d, i["ap_sniffer_chan"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferData(d, i["ap_sniffer_data"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandObjectWirelessControllerWtpProfileRadio1AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandObjectWirelessControllerWtpProfileRadio1AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandObjectWirelessControllerWtpProfileRadio1AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandObjectWirelessControllerWtpProfileRadio1AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandObjectWirelessControllerWtpProfileRadio1Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandObjectWirelessControllerWtpProfileRadio1Band5GType(d, i["band_5g_type"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio1BandwidthCapacity(d, i["bandwidth_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandObjectWirelessControllerWtpProfileRadio1BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandObjectWirelessControllerWtpProfileRadio1BssColor(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio1CallAdmissionControl(d, i["call_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio1CallCapacity(d, i["call_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandObjectWirelessControllerWtpProfileRadio1Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandObjectWirelessControllerWtpProfileRadio1ChannelBonding(d, i["channel_bonding"], pre_append)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandObjectWirelessControllerWtpProfileRadio1ChannelUtilization(d, i["channel_utilization"], pre_append)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandObjectWirelessControllerWtpProfileRadio1Coexistence(d, i["coexistence"], pre_append)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandObjectWirelessControllerWtpProfileRadio1Darrp(d, i["darrp"], pre_append)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandObjectWirelessControllerWtpProfileRadio1Drma(d, i["drma"], pre_append)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandObjectWirelessControllerWtpProfileRadio1DrmaSensitivity(d, i["drma_sensitivity"], pre_append)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandObjectWirelessControllerWtpProfileRadio1Dtim(d, i["dtim"], pre_append)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio1FragThreshold(d, i["frag_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandObjectWirelessControllerWtpProfileRadio1IperfProtocol(d, i["iperf_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandObjectWirelessControllerWtpProfileRadio1IperfServerPort(d, i["iperf_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandObjectWirelessControllerWtpProfileRadio1MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandObjectWirelessControllerWtpProfileRadio1MaxDistance(d, i["max_distance"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandObjectWirelessControllerWtpProfileRadio1Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandObjectWirelessControllerWtpProfileRadio1PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandObjectWirelessControllerWtpProfileRadio1PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandObjectWirelessControllerWtpProfileRadio1PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio1PowersaveOptimize(d, i["powersave_optimize"], pre_append)
	} else {
		result["powersave-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandObjectWirelessControllerWtpProfileRadio1ProtectionMode(d, i["protection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandObjectWirelessControllerWtpProfileRadio1RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio1RtsThreshold(d, i["rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandObjectWirelessControllerWtpProfileRadio1SamBssid(d, i["sam_bssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandObjectWirelessControllerWtpProfileRadio1SamCaptivePortal(d, i["sam_captive_portal"], pre_append)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandObjectWirelessControllerWtpProfileRadio1SamPassword(d, i["sam_password"], pre_append)
	} else {
		result["sam-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandObjectWirelessControllerWtpProfileRadio1SamReportIntv(d, i["sam_report_intv"], pre_append)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandObjectWirelessControllerWtpProfileRadio1SamSecurityType(d, i["sam_security_type"], pre_append)
	}
	pre_append = pre + ".0." + "sam_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server"], _ = expandObjectWirelessControllerWtpProfileRadio1SamServer(d, i["sam_server"], pre_append)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandObjectWirelessControllerWtpProfileRadio1SamSsid(d, i["sam_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandObjectWirelessControllerWtpProfileRadio1SamTest(d, i["sam_test"], pre_append)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandObjectWirelessControllerWtpProfileRadio1SamUsername(d, i["sam_username"], pre_append)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandObjectWirelessControllerWtpProfileRadio1ShortGuardInterval(d, i["short_guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandObjectWirelessControllerWtpProfileRadio1SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio1TransmitOptimize(d, i["transmit_optimize"], pre_append)
	} else {
		result["transmit-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandObjectWirelessControllerWtpProfileRadio1VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap1"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap2"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap3"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap4"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap5"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap6"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap7"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap8"], _ = expandObjectWirelessControllerWtpProfileRadio1Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandObjectWirelessControllerWtpProfileRadio1Vaps(d, i["vaps"], pre_append)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandObjectWirelessControllerWtpProfileRadio1WidsProfile(d, i["wids_profile"], pre_append)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandObjectWirelessControllerWtpProfileRadio1ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileRadio1AirtimeFairness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Amsdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Band5GType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1BssColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1CallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio1ChannelBonding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ChannelUtilization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Coexistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Darrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Drma(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Dtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1FragThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1IperfProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1IperfServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1MaxDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio1ProtectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1RtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamBssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamSecurityType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SamUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1TransmitOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio1VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1WidsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio1ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandObjectWirelessControllerWtpProfileRadio2AirtimeFairness(d, i["airtime_fairness"], pre_append)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandObjectWirelessControllerWtpProfileRadio2Amsdu(d, i["amsdu"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferChan(d, i["ap_sniffer_chan"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferData(d, i["ap_sniffer_data"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandObjectWirelessControllerWtpProfileRadio2AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandObjectWirelessControllerWtpProfileRadio2AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandObjectWirelessControllerWtpProfileRadio2AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandObjectWirelessControllerWtpProfileRadio2AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandObjectWirelessControllerWtpProfileRadio2Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandObjectWirelessControllerWtpProfileRadio2Band5GType(d, i["band_5g_type"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio2BandwidthCapacity(d, i["bandwidth_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandObjectWirelessControllerWtpProfileRadio2BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandObjectWirelessControllerWtpProfileRadio2BssColor(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio2CallAdmissionControl(d, i["call_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio2CallCapacity(d, i["call_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandObjectWirelessControllerWtpProfileRadio2Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandObjectWirelessControllerWtpProfileRadio2ChannelBonding(d, i["channel_bonding"], pre_append)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandObjectWirelessControllerWtpProfileRadio2ChannelUtilization(d, i["channel_utilization"], pre_append)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandObjectWirelessControllerWtpProfileRadio2Coexistence(d, i["coexistence"], pre_append)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandObjectWirelessControllerWtpProfileRadio2Darrp(d, i["darrp"], pre_append)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandObjectWirelessControllerWtpProfileRadio2Drma(d, i["drma"], pre_append)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandObjectWirelessControllerWtpProfileRadio2DrmaSensitivity(d, i["drma_sensitivity"], pre_append)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandObjectWirelessControllerWtpProfileRadio2Dtim(d, i["dtim"], pre_append)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio2FragThreshold(d, i["frag_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandObjectWirelessControllerWtpProfileRadio2IperfProtocol(d, i["iperf_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandObjectWirelessControllerWtpProfileRadio2IperfServerPort(d, i["iperf_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandObjectWirelessControllerWtpProfileRadio2MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandObjectWirelessControllerWtpProfileRadio2MaxDistance(d, i["max_distance"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandObjectWirelessControllerWtpProfileRadio2Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandObjectWirelessControllerWtpProfileRadio2PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandObjectWirelessControllerWtpProfileRadio2PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandObjectWirelessControllerWtpProfileRadio2PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio2PowersaveOptimize(d, i["powersave_optimize"], pre_append)
	} else {
		result["powersave-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandObjectWirelessControllerWtpProfileRadio2ProtectionMode(d, i["protection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandObjectWirelessControllerWtpProfileRadio2RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio2RtsThreshold(d, i["rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandObjectWirelessControllerWtpProfileRadio2SamBssid(d, i["sam_bssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandObjectWirelessControllerWtpProfileRadio2SamCaptivePortal(d, i["sam_captive_portal"], pre_append)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandObjectWirelessControllerWtpProfileRadio2SamPassword(d, i["sam_password"], pre_append)
	} else {
		result["sam-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandObjectWirelessControllerWtpProfileRadio2SamReportIntv(d, i["sam_report_intv"], pre_append)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandObjectWirelessControllerWtpProfileRadio2SamSecurityType(d, i["sam_security_type"], pre_append)
	}
	pre_append = pre + ".0." + "sam_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server"], _ = expandObjectWirelessControllerWtpProfileRadio2SamServer(d, i["sam_server"], pre_append)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandObjectWirelessControllerWtpProfileRadio2SamSsid(d, i["sam_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandObjectWirelessControllerWtpProfileRadio2SamTest(d, i["sam_test"], pre_append)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandObjectWirelessControllerWtpProfileRadio2SamUsername(d, i["sam_username"], pre_append)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandObjectWirelessControllerWtpProfileRadio2ShortGuardInterval(d, i["short_guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandObjectWirelessControllerWtpProfileRadio2SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio2TransmitOptimize(d, i["transmit_optimize"], pre_append)
	} else {
		result["transmit-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandObjectWirelessControllerWtpProfileRadio2VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap1"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap2"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap3"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap4"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap5"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap6"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap7"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap8"], _ = expandObjectWirelessControllerWtpProfileRadio2Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandObjectWirelessControllerWtpProfileRadio2Vaps(d, i["vaps"], pre_append)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandObjectWirelessControllerWtpProfileRadio2WidsProfile(d, i["wids_profile"], pre_append)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandObjectWirelessControllerWtpProfileRadio2ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileRadio2AirtimeFairness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Amsdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Band5GType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2BssColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2CallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio2ChannelBonding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ChannelUtilization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Coexistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Darrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Drma(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Dtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2FragThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2IperfProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2IperfServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2MaxDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio2ProtectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2RtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamBssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamSecurityType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SamUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2TransmitOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio2VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2WidsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio2ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandObjectWirelessControllerWtpProfileRadio3AirtimeFairness(d, i["airtime_fairness"], pre_append)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandObjectWirelessControllerWtpProfileRadio3Amsdu(d, i["amsdu"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferChan(d, i["ap_sniffer_chan"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferData(d, i["ap_sniffer_data"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandObjectWirelessControllerWtpProfileRadio3AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandObjectWirelessControllerWtpProfileRadio3AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandObjectWirelessControllerWtpProfileRadio3AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandObjectWirelessControllerWtpProfileRadio3AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandObjectWirelessControllerWtpProfileRadio3Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandObjectWirelessControllerWtpProfileRadio3Band5GType(d, i["band_5g_type"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio3BandwidthCapacity(d, i["bandwidth_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandObjectWirelessControllerWtpProfileRadio3BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandObjectWirelessControllerWtpProfileRadio3BssColor(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio3CallAdmissionControl(d, i["call_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio3CallCapacity(d, i["call_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandObjectWirelessControllerWtpProfileRadio3Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandObjectWirelessControllerWtpProfileRadio3ChannelBonding(d, i["channel_bonding"], pre_append)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandObjectWirelessControllerWtpProfileRadio3ChannelUtilization(d, i["channel_utilization"], pre_append)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandObjectWirelessControllerWtpProfileRadio3Coexistence(d, i["coexistence"], pre_append)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandObjectWirelessControllerWtpProfileRadio3Darrp(d, i["darrp"], pre_append)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandObjectWirelessControllerWtpProfileRadio3Drma(d, i["drma"], pre_append)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandObjectWirelessControllerWtpProfileRadio3DrmaSensitivity(d, i["drma_sensitivity"], pre_append)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandObjectWirelessControllerWtpProfileRadio3Dtim(d, i["dtim"], pre_append)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio3FragThreshold(d, i["frag_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandObjectWirelessControllerWtpProfileRadio3IperfProtocol(d, i["iperf_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandObjectWirelessControllerWtpProfileRadio3IperfServerPort(d, i["iperf_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandObjectWirelessControllerWtpProfileRadio3MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandObjectWirelessControllerWtpProfileRadio3MaxDistance(d, i["max_distance"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandObjectWirelessControllerWtpProfileRadio3Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandObjectWirelessControllerWtpProfileRadio3PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandObjectWirelessControllerWtpProfileRadio3PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandObjectWirelessControllerWtpProfileRadio3PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio3PowersaveOptimize(d, i["powersave_optimize"], pre_append)
	} else {
		result["powersave-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandObjectWirelessControllerWtpProfileRadio3ProtectionMode(d, i["protection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandObjectWirelessControllerWtpProfileRadio3RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio3RtsThreshold(d, i["rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandObjectWirelessControllerWtpProfileRadio3SamBssid(d, i["sam_bssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandObjectWirelessControllerWtpProfileRadio3SamCaptivePortal(d, i["sam_captive_portal"], pre_append)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandObjectWirelessControllerWtpProfileRadio3SamPassword(d, i["sam_password"], pre_append)
	} else {
		result["sam-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandObjectWirelessControllerWtpProfileRadio3SamReportIntv(d, i["sam_report_intv"], pre_append)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandObjectWirelessControllerWtpProfileRadio3SamSecurityType(d, i["sam_security_type"], pre_append)
	}
	pre_append = pre + ".0." + "sam_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server"], _ = expandObjectWirelessControllerWtpProfileRadio3SamServer(d, i["sam_server"], pre_append)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandObjectWirelessControllerWtpProfileRadio3SamSsid(d, i["sam_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandObjectWirelessControllerWtpProfileRadio3SamTest(d, i["sam_test"], pre_append)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandObjectWirelessControllerWtpProfileRadio3SamUsername(d, i["sam_username"], pre_append)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandObjectWirelessControllerWtpProfileRadio3ShortGuardInterval(d, i["short_guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandObjectWirelessControllerWtpProfileRadio3SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio3TransmitOptimize(d, i["transmit_optimize"], pre_append)
	} else {
		result["transmit-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandObjectWirelessControllerWtpProfileRadio3VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap1"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap2"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap3"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap4"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap5"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap6"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap7"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap8"], _ = expandObjectWirelessControllerWtpProfileRadio3Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandObjectWirelessControllerWtpProfileRadio3Vaps(d, i["vaps"], pre_append)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandObjectWirelessControllerWtpProfileRadio3WidsProfile(d, i["wids_profile"], pre_append)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandObjectWirelessControllerWtpProfileRadio3ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileRadio3AirtimeFairness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Amsdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Band5GType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3BssColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3CallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio3ChannelBonding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ChannelUtilization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Coexistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Darrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Drma(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Dtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3FragThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3IperfProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3IperfServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3MaxDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio3ProtectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3RtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamBssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamSecurityType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SamUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3TransmitOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio3VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3WidsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio3ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandObjectWirelessControllerWtpProfileRadio4AirtimeFairness(d, i["airtime_fairness"], pre_append)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandObjectWirelessControllerWtpProfileRadio4Amsdu(d, i["amsdu"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferChan(d, i["ap_sniffer_chan"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferData(d, i["ap_sniffer_data"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandObjectWirelessControllerWtpProfileRadio4AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandObjectWirelessControllerWtpProfileRadio4AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandObjectWirelessControllerWtpProfileRadio4AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandObjectWirelessControllerWtpProfileRadio4AutoPowerTarget(d, i["auto_power_target"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandObjectWirelessControllerWtpProfileRadio4Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandObjectWirelessControllerWtpProfileRadio4Band5GType(d, i["band_5g_type"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio4BandwidthCapacity(d, i["bandwidth_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandObjectWirelessControllerWtpProfileRadio4BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandObjectWirelessControllerWtpProfileRadio4BssColor(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandObjectWirelessControllerWtpProfileRadio4CallAdmissionControl(d, i["call_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandObjectWirelessControllerWtpProfileRadio4CallCapacity(d, i["call_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandObjectWirelessControllerWtpProfileRadio4Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandObjectWirelessControllerWtpProfileRadio4ChannelBonding(d, i["channel_bonding"], pre_append)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandObjectWirelessControllerWtpProfileRadio4ChannelUtilization(d, i["channel_utilization"], pre_append)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandObjectWirelessControllerWtpProfileRadio4Coexistence(d, i["coexistence"], pre_append)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandObjectWirelessControllerWtpProfileRadio4Darrp(d, i["darrp"], pre_append)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandObjectWirelessControllerWtpProfileRadio4Drma(d, i["drma"], pre_append)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandObjectWirelessControllerWtpProfileRadio4DrmaSensitivity(d, i["drma_sensitivity"], pre_append)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandObjectWirelessControllerWtpProfileRadio4Dtim(d, i["dtim"], pre_append)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio4FragThreshold(d, i["frag_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandObjectWirelessControllerWtpProfileRadio4IperfProtocol(d, i["iperf_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandObjectWirelessControllerWtpProfileRadio4IperfServerPort(d, i["iperf_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandObjectWirelessControllerWtpProfileRadio4MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandObjectWirelessControllerWtpProfileRadio4MaxDistance(d, i["max_distance"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandObjectWirelessControllerWtpProfileRadio4Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandObjectWirelessControllerWtpProfileRadio4PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandObjectWirelessControllerWtpProfileRadio4PowerMode(d, i["power_mode"], pre_append)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandObjectWirelessControllerWtpProfileRadio4PowerValue(d, i["power_value"], pre_append)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio4PowersaveOptimize(d, i["powersave_optimize"], pre_append)
	} else {
		result["powersave-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandObjectWirelessControllerWtpProfileRadio4ProtectionMode(d, i["protection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandObjectWirelessControllerWtpProfileRadio4RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandObjectWirelessControllerWtpProfileRadio4RtsThreshold(d, i["rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandObjectWirelessControllerWtpProfileRadio4SamBssid(d, i["sam_bssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandObjectWirelessControllerWtpProfileRadio4SamCaptivePortal(d, i["sam_captive_portal"], pre_append)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandObjectWirelessControllerWtpProfileRadio4SamPassword(d, i["sam_password"], pre_append)
	} else {
		result["sam-password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandObjectWirelessControllerWtpProfileRadio4SamReportIntv(d, i["sam_report_intv"], pre_append)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandObjectWirelessControllerWtpProfileRadio4SamSecurityType(d, i["sam_security_type"], pre_append)
	}
	pre_append = pre + ".0." + "sam_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server"], _ = expandObjectWirelessControllerWtpProfileRadio4SamServer(d, i["sam_server"], pre_append)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandObjectWirelessControllerWtpProfileRadio4SamSsid(d, i["sam_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandObjectWirelessControllerWtpProfileRadio4SamTest(d, i["sam_test"], pre_append)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandObjectWirelessControllerWtpProfileRadio4SamUsername(d, i["sam_username"], pre_append)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandObjectWirelessControllerWtpProfileRadio4ShortGuardInterval(d, i["short_guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandObjectWirelessControllerWtpProfileRadio4SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandObjectWirelessControllerWtpProfileRadio4TransmitOptimize(d, i["transmit_optimize"], pre_append)
	} else {
		result["transmit-optimize"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandObjectWirelessControllerWtpProfileRadio4VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vap1"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap1"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap1(d, i["vap1"], pre_append)
	}
	pre_append = pre + ".0." + "vap2"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap2"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap2(d, i["vap2"], pre_append)
	}
	pre_append = pre + ".0." + "vap3"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap3"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap3(d, i["vap3"], pre_append)
	}
	pre_append = pre + ".0." + "vap4"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap4"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap4(d, i["vap4"], pre_append)
	}
	pre_append = pre + ".0." + "vap5"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap5"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap5(d, i["vap5"], pre_append)
	}
	pre_append = pre + ".0." + "vap6"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap6"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap6(d, i["vap6"], pre_append)
	}
	pre_append = pre + ".0." + "vap7"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap7"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap7(d, i["vap7"], pre_append)
	}
	pre_append = pre + ".0." + "vap8"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap8"], _ = expandObjectWirelessControllerWtpProfileRadio4Vap8(d, i["vap8"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandObjectWirelessControllerWtpProfileRadio4Vaps(d, i["vaps"], pre_append)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandObjectWirelessControllerWtpProfileRadio4WidsProfile(d, i["wids_profile"], pre_append)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandObjectWirelessControllerWtpProfileRadio4ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append)
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileRadio4AirtimeFairness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Amsdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Band5GType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4BssColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4CallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio4ChannelBonding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ChannelUtilization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Coexistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Darrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Drma(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Dtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4FragThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4IperfProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4IperfServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4MaxDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4PowerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4PowerValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio4ProtectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4RtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamBssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamSecurityType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SamUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4TransmitOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWtpProfileRadio4VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vap8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4WidsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileRadio4ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileSnmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dest-ip"], _ = expandObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp(d, i["dest_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectWirelessControllerWtpProfileSplitTunnelingAclId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAclDestIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAclId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileSplitTunnelingAclPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileTunMtuDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileTunMtuUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfileWanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWtpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandObjectWirelessControllerWtpProfileAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("ap_country"); ok {
		t, err := expandObjectWirelessControllerWtpProfileApCountry(d, v, "ap_country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-country"] = t
		}
	}

	if v, ok := d.GetOk("ap_handoff"); ok {
		t, err := expandObjectWirelessControllerWtpProfileApHandoff(d, v, "ap_handoff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-handoff"] = t
		}
	}

	if v, ok := d.GetOk("apcfg_profile"); ok {
		t, err := expandObjectWirelessControllerWtpProfileApcfgProfile(d, v, "apcfg_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apcfg-profile"] = t
		}
	}

	if v, ok := d.GetOk("ble_profile"); ok {
		t, err := expandObjectWirelessControllerWtpProfileBleProfile(d, v, "ble_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-profile"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectWirelessControllerWtpProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("control_message_offload"); ok {
		t, err := expandObjectWirelessControllerWtpProfileControlMessageOffload(d, v, "control_message_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-message-offload"] = t
		}
	}

	if v, ok := d.GetOk("deny_mac_list"); ok {
		t, err := expandObjectWirelessControllerWtpProfileDenyMacList(d, v, "deny_mac_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny-mac-list"] = t
		}
	}

	if v, ok := d.GetOk("dtls_in_kernel"); ok {
		t, err := expandObjectWirelessControllerWtpProfileDtlsInKernel(d, v, "dtls_in_kernel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-in-kernel"] = t
		}
	}

	if v, ok := d.GetOk("dtls_policy"); ok {
		t, err := expandObjectWirelessControllerWtpProfileDtlsPolicy(d, v, "dtls_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-policy"] = t
		}
	}

	if v, ok := d.GetOk("energy_efficient_ethernet"); ok {
		t, err := expandObjectWirelessControllerWtpProfileEnergyEfficientEthernet(d, v, "energy_efficient_ethernet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["energy-efficient-ethernet"] = t
		}
	}

	if v, ok := d.GetOk("ext_info_enable"); ok {
		t, err := expandObjectWirelessControllerWtpProfileExtInfoEnable(d, v, "ext_info_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-info-enable"] = t
		}
	}

	if v, ok := d.GetOk("frequency_handoff"); ok {
		t, err := expandObjectWirelessControllerWtpProfileFrequencyHandoff(d, v, "frequency_handoff")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frequency-handoff"] = t
		}
	}

	if v, ok := d.GetOk("handoff_roaming"); ok {
		t, err := expandObjectWirelessControllerWtpProfileHandoffRoaming(d, v, "handoff_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-roaming"] = t
		}
	}

	if v, ok := d.GetOk("handoff_rssi"); ok {
		t, err := expandObjectWirelessControllerWtpProfileHandoffRssi(d, v, "handoff_rssi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-rssi"] = t
		}
	}

	if v, ok := d.GetOk("handoff_sta_thresh"); ok {
		t, err := expandObjectWirelessControllerWtpProfileHandoffStaThresh(d, v, "handoff_sta_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-sta-thresh"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_preventing"); ok {
		t, err := expandObjectWirelessControllerWtpProfileIpFragmentPreventing(d, v, "ip_fragment_preventing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-preventing"] = t
		}
	}

	if v, ok := d.GetOk("lan"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLan(d, v, "lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan"] = t
		}
	}

	if v, ok := d.GetOk("lbs"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLbs(d, v, "lbs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lbs"] = t
		}
	}

	if v, ok := d.GetOk("led_schedules"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLedSchedules(d, v, "led_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-schedules"] = t
		}
	}

	if v, ok := d.GetOk("led_state"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLedState(d, v, "led_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-state"] = t
		}
	}

	if v, ok := d.GetOk("lldp"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLldp(d, v, "lldp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLoginPasswd(d, v, "login_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_change"); ok {
		t, err := expandObjectWirelessControllerWtpProfileLoginPasswdChange(d, v, "login_passwd_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok {
		t, err := expandObjectWirelessControllerWtpProfileMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerWtpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("platform"); ok {
		t, err := expandObjectWirelessControllerWtpProfilePlatform(d, v, "platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform"] = t
		}
	}

	if v, ok := d.GetOk("poe_mode"); ok {
		t, err := expandObjectWirelessControllerWtpProfilePoeMode(d, v, "poe_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-mode"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok {
		t, err := expandObjectWirelessControllerWtpProfileRadio1(d, v, "radio_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok {
		t, err := expandObjectWirelessControllerWtpProfileRadio2(d, v, "radio_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	if v, ok := d.GetOk("radio_3"); ok {
		t, err := expandObjectWirelessControllerWtpProfileRadio3(d, v, "radio_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-3"] = t
		}
	}

	if v, ok := d.GetOk("radio_4"); ok {
		t, err := expandObjectWirelessControllerWtpProfileRadio4(d, v, "radio_4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-4"] = t
		}
	}

	if v, ok := d.GetOk("snmp"); ok {
		t, err := expandObjectWirelessControllerWtpProfileSnmp(d, v, "snmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl"); ok {
		t, err := expandObjectWirelessControllerWtpProfileSplitTunnelingAcl(d, v, "split_tunneling_acl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok {
		t, err := expandObjectWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(d, v, "split_tunneling_acl_local_ap_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-local-ap-subnet"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_path"); ok {
		t, err := expandObjectWirelessControllerWtpProfileSplitTunnelingAclPath(d, v, "split_tunneling_acl_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-path"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_downlink"); ok {
		t, err := expandObjectWirelessControllerWtpProfileTunMtuDownlink(d, v, "tun_mtu_downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-downlink"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_uplink"); ok {
		t, err := expandObjectWirelessControllerWtpProfileTunMtuUplink(d, v, "tun_mtu_uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-uplink"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_mode"); ok {
		t, err := expandObjectWirelessControllerWtpProfileWanPortMode(d, v, "wan_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-mode"] = t
		}
	}

	return &obj, nil
}
