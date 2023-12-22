// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch devices that are managed by this FortiGate.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerManagedSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerManagedSwitchCreate,
		Read:   resourceObjectSwitchControllerManagedSwitchRead,
		Update: resourceObjectSwitchControllerManagedSwitchUpdate,
		Delete: resourceObjectSwitchControllerManagedSwitchDelete,

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
			"_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"command_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping_static_client": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"firmware_provision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"l3_discovered": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mclag_igmp_snooping_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_snmp_community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_snmp_sysinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_snmp_trap_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_snmp_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_detection_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"acl_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"aggregator_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"allowed_vlans": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"allowed_vlans_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"arp_inspection_trust": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authenticated_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bundle": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dhcp_snoop_option82_override": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"circuit_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vlan_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"dhcp_snoop_option82_trust": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"discard_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dsl_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"edge_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encrypted_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fec_capable": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fec_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flap_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"flap_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"flap_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flapguard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flow_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiswitch_acls": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"igmp_snooping_flood_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"igmps_flood_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmps_flood_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_tags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_source_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"isl_peer_device_sn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lacp_speed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"learning_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"link_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lldp_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lldp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"loop_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"loop_guard_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"matched_dpp_intf_tags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"matched_dpp_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_bundle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mcast_snooping_flood_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mclag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mclag_icl_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"media_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"member_withdrawal_behavior": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"min_bundle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"p2p_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_sample_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pause_meter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pause_meter_resume": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_max_power": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_mode_bt_cabable": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"poe_port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_port_power": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_port_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_pre_standard_detection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_standard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poe_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_owner": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_security_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_selection_criteria": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ptp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"qos_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"restricted_auth_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rpvst_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sample_direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sflow_counter_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sflow_sample_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sflow_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sticky_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_bpdu_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_bpdu_guard_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"stp_root_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trunk_member": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untagged_vlans": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ptp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_drop_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_red_probability": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"route_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload_mclag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload_router": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"router_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"switch_dhcp_opt43_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"tdr_supported": &schema.Schema{
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

func resourceObjectSwitchControllerManagedSwitchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerManagedSwitch(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitch resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerManagedSwitch(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitch resource: %v", err)
	}

	d.SetId(getStringKey(d, "switch_id"))

	return resourceObjectSwitchControllerManagedSwitchRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerManagedSwitch(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitch resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerManagedSwitch(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "switch_id"))

	return resourceObjectSwitchControllerManagedSwitchRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerManagedSwitch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerManagedSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerManagedSwitchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerManagedSwitch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerManagedSwitch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitch resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerManagedSwitchPlatform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchCustomCommand(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := i["command-entry"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchCustomCommandCommandEntry(i["command-entry"], d, pre_append)
			tmp["command_entry"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-CustomCommand-CommandEntry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchCustomCommandCommandName(i["command-name"], d, pre_append)
			tmp["command_name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-CustomCommand-CommandName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerManagedSwitchCustomCommandCommandEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient-Vlan")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchFirmwareProvision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchFirmwareProvisionVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchL3Discovered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchMclagIgmpSnoopingAware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchOverrideSnmpCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchOverrideSnmpSysinfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchOverrideSnmpUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPoeDetectionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPorts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if _, ok := i["access-mode"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsAccessMode(i["access-mode"], d, pre_append)
			tmp["access_mode"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-AccessMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acl_group"
		if _, ok := i["acl-group"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsAclGroup(i["acl-group"], d, pre_append)
			tmp["acl_group"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-AclGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := i["aggregator-mode"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsAggregatorMode(i["aggregator-mode"], d, pre_append)
			tmp["aggregator_mode"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-AggregatorMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if _, ok := i["allowed-vlans"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsAllowedVlans(i["allowed-vlans"], d, pre_append)
			tmp["allowed_vlans"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-AllowedVlans")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if _, ok := i["allowed-vlans-all"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsAllowedVlansAll(i["allowed-vlans-all"], d, pre_append)
			tmp["allowed_vlans_all"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-AllowedVlansAll")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := i["arp-inspection-trust"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsArpInspectionTrust(i["arp-inspection-trust"], d, pre_append)
			tmp["arp_inspection_trust"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-ArpInspectionTrust")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authenticated_port"
		if _, ok := i["authenticated-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsAuthenticatedPort(i["authenticated-port"], d, pre_append)
			tmp["authenticated_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-AuthenticatedPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if _, ok := i["bundle"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsBundle(i["bundle"], d, pre_append)
			tmp["bundle"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Bundle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_override"
		if _, ok := i["dhcp-snoop-option82-override"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(i["dhcp-snoop-option82-override"], d, pre_append)
			tmp["dhcp_snoop_option82_override"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-DhcpSnoopOption82Override")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := i["dhcp-snoop-option82-trust"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(i["dhcp-snoop-option82-trust"], d, pre_append)
			tmp["dhcp_snoop_option82_trust"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-DhcpSnoopOption82Trust")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := i["dhcp-snooping"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnooping(i["dhcp-snooping"], d, pre_append)
			tmp["dhcp_snooping"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-DhcpSnooping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if _, ok := i["discard-mode"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDiscardMode(i["discard-mode"], d, pre_append)
			tmp["discard_mode"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-DiscardMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dsl_profile"
		if _, ok := i["dsl-profile"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDslProfile(i["dsl-profile"], d, pre_append)
			tmp["dsl_profile"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-DslProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if _, ok := i["edge-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsEdgePort(i["edge-port"], d, pre_append)
			tmp["edge_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-EdgePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypted_port"
		if _, ok := i["encrypted-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsEncryptedPort(i["encrypted-port"], d, pre_append)
			tmp["encrypted_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-EncryptedPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if _, ok := i["fec-capable"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFecCapable(i["fec-capable"], d, pre_append)
			tmp["fec_capable"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FecCapable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if _, ok := i["fec-state"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFecState(i["fec-state"], d, pre_append)
			tmp["fec_state"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FecState")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_duration"
		if _, ok := i["flap-duration"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFlapDuration(i["flap-duration"], d, pre_append)
			tmp["flap_duration"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FlapDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_rate"
		if _, ok := i["flap-rate"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFlapRate(i["flap-rate"], d, pre_append)
			tmp["flap_rate"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FlapRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_timeout"
		if _, ok := i["flap-timeout"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFlapTimeout(i["flap-timeout"], d, pre_append)
			tmp["flap_timeout"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FlapTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flapguard"
		if _, ok := i["flapguard"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFlapguard(i["flapguard"], d, pre_append)
			tmp["flapguard"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Flapguard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if _, ok := i["flow-control"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFlowControl(i["flow-control"], d, pre_append)
			tmp["flow_control"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FlowControl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiswitch_acls"
		if _, ok := i["fortiswitch-acls"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsFortiswitchAcls(i["fortiswitch-acls"], d, pre_append)
			tmp["fortiswitch_acls"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-FortiswitchAcls")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping_flood_reports"
		if _, ok := i["igmp-snooping-flood-reports"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(i["igmp-snooping-flood-reports"], d, pre_append)
			tmp["igmp_snooping_flood_reports"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-IgmpSnoopingFloodReports")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := i["igmp-snooping"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsIgmpSnooping(i["igmp-snooping"], d, pre_append)
			tmp["igmp_snooping"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-IgmpSnooping")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if _, ok := i["igmps-flood-reports"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports(i["igmps-flood-reports"], d, pre_append)
			tmp["igmps_flood_reports"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-IgmpsFloodReports")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if _, ok := i["igmps-flood-traffic"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(i["igmps-flood-traffic"], d, pre_append)
			tmp["igmps_flood_traffic"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-IgmpsFloodTraffic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := i["interface-tags"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsInterfaceTags(i["interface-tags"], d, pre_append)
			tmp["interface_tags"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-InterfaceTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if _, ok := i["ip-source-guard"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsIpSourceGuard(i["ip-source-guard"], d, pre_append)
			tmp["ip_source_guard"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-IpSourceGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_sn"
		if _, ok := i["isl-peer-device-sn"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn(i["isl-peer-device-sn"], d, pre_append)
			tmp["isl_peer_device_sn"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-IslPeerDeviceSn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if _, ok := i["lacp-speed"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLacpSpeed(i["lacp-speed"], d, pre_append)
			tmp["lacp_speed"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LacpSpeed")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if _, ok := i["learning-limit"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLearningLimit(i["learning-limit"], d, pre_append)
			tmp["learning_limit"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LearningLimit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_status"
		if _, ok := i["link-status"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLinkStatus(i["link-status"], d, pre_append)
			tmp["link_status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LinkStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := i["lldp-profile"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLldpProfile(i["lldp-profile"], d, pre_append)
			tmp["lldp_profile"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LldpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if _, ok := i["lldp-status"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLldpStatus(i["lldp-status"], d, pre_append)
			tmp["lldp_status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LldpStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if _, ok := i["loop-guard"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLoopGuard(i["loop-guard"], d, pre_append)
			tmp["loop_guard"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LoopGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if _, ok := i["loop-guard-timeout"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout(i["loop-guard-timeout"], d, pre_append)
			tmp["loop_guard_timeout"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-LoopGuardTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_intf_tags"
		if _, ok := i["matched-dpp-intf-tags"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags(i["matched-dpp-intf-tags"], d, pre_append)
			tmp["matched_dpp_intf_tags"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MatchedDppIntfTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_policy"
		if _, ok := i["matched-dpp-policy"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy(i["matched-dpp-policy"], d, pre_append)
			tmp["matched_dpp_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MatchedDppPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if _, ok := i["max-bundle"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMaxBundle(i["max-bundle"], d, pre_append)
			tmp["max_bundle"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MaxBundle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_snooping_flood_traffic"
		if _, ok := i["mcast-snooping-flood-traffic"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(i["mcast-snooping-flood-traffic"], d, pre_append)
			tmp["mcast_snooping_flood_traffic"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-McastSnoopingFloodTraffic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if _, ok := i["mclag"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMclag(i["mclag"], d, pre_append)
			tmp["mclag"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Mclag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if _, ok := i["mclag-icl-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMclagIclPort(i["mclag-icl-port"], d, pre_append)
			tmp["mclag_icl_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MclagIclPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := i["media-type"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMediaType(i["media-type"], d, pre_append)
			tmp["media_type"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MediaType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if _, ok := i["member-withdrawal-behavior"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(i["member-withdrawal-behavior"], d, pre_append)
			tmp["member_withdrawal_behavior"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MemberWithdrawalBehavior")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if _, ok := i["min-bundle"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMinBundle(i["min-bundle"], d, pre_append)
			tmp["min_bundle"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-MinBundle")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if _, ok := i["p2p-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsP2PPort(i["p2p-port"], d, pre_append)
			tmp["p2p_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-P2PPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if _, ok := i["packet-sample-rate"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPacketSampleRate(i["packet-sample-rate"], d, pre_append)
			tmp["packet_sample_rate"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PacketSampleRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if _, ok := i["packet-sampler"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPacketSampler(i["packet-sampler"], d, pre_append)
			tmp["packet_sampler"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PacketSampler")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if _, ok := i["pause-meter"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPauseMeter(i["pause-meter"], d, pre_append)
			tmp["pause_meter"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PauseMeter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if _, ok := i["pause-meter-resume"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPauseMeterResume(i["pause-meter-resume"], d, pre_append)
			tmp["pause_meter_resume"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PauseMeterResume")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_max_power"
		if _, ok := i["poe-max-power"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoeMaxPower(i["poe-max-power"], d, pre_append)
			tmp["poe_max_power"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoeMaxPower")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_mode_bt_cabable"
		if _, ok := i["poe-mode-bt-cabable"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable(i["poe-mode-bt-cabable"], d, pre_append)
			tmp["poe_mode_bt_cabable"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoeModeBtCabable")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_mode"
		if _, ok := i["poe-port-mode"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoePortMode(i["poe-port-mode"], d, pre_append)
			tmp["poe_port_mode"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoePortMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_power"
		if _, ok := i["poe-port-power"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoePortPower(i["poe-port-power"], d, pre_append)
			tmp["poe_port_power"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoePortPower")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_priority"
		if _, ok := i["poe-port-priority"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoePortPriority(i["poe-port-priority"], d, pre_append)
			tmp["poe_port_priority"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoePortPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if _, ok := i["poe-pre-standard-detection"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection(i["poe-pre-standard-detection"], d, pre_append)
			tmp["poe_pre_standard_detection"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoePreStandardDetection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_standard"
		if _, ok := i["poe-standard"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoeStandard(i["poe-standard"], d, pre_append)
			tmp["poe_standard"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoeStandard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if _, ok := i["poe-status"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPoeStatus(i["poe-status"], d, pre_append)
			tmp["poe_status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PoeStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if _, ok := i["port-name"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPortName(i["port-name"], d, pre_append)
			tmp["port_name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PortName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if _, ok := i["port-owner"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPortOwner(i["port-owner"], d, pre_append)
			tmp["port_owner"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PortOwner")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_policy"
		if _, ok := i["port-policy"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPortPolicy(i["port-policy"], d, pre_append)
			tmp["port_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PortPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if _, ok := i["port-security-policy"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy(i["port-security-policy"], d, pre_append)
			tmp["port_security_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PortSecurityPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if _, ok := i["port-selection-criteria"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria(i["port-selection-criteria"], d, pre_append)
			tmp["port_selection_criteria"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PortSelectionCriteria")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_status"
		if _, ok := i["ptp-status"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsPtpStatus(i["ptp-status"], d, pre_append)
			tmp["ptp_status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-PtpStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := i["qos-policy"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsQosPolicy(i["qos-policy"], d, pre_append)
			tmp["qos_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-QosPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_auth_port"
		if _, ok := i["restricted-auth-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort(i["restricted-auth-port"], d, pre_append)
			tmp["restricted_auth_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-RestrictedAuthPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if _, ok := i["rpvst-port"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsRpvstPort(i["rpvst-port"], d, pre_append)
			tmp["rpvst_port"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-RpvstPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if _, ok := i["sample-direction"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsSampleDirection(i["sample-direction"], d, pre_append)
			tmp["sample_direction"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-SampleDirection")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if _, ok := i["sflow-counter-interval"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsSflowCounterInterval(i["sflow-counter-interval"], d, pre_append)
			tmp["sflow_counter_interval"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-SflowCounterInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sample_rate"
		if _, ok := i["sflow-sample-rate"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsSflowSampleRate(i["sflow-sample-rate"], d, pre_append)
			tmp["sflow_sample_rate"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-SflowSampleRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sampler"
		if _, ok := i["sflow-sampler"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsSflowSampler(i["sflow-sampler"], d, pre_append)
			tmp["sflow_sampler"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-SflowSampler")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if _, ok := i["sticky-mac"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsStickyMac(i["sticky-mac"], d, pre_append)
			tmp["sticky_mac"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-StickyMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if _, ok := i["stp-bpdu-guard"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuard(i["stp-bpdu-guard"], d, pre_append)
			tmp["stp_bpdu_guard"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-StpBpduGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if _, ok := i["stp-bpdu-guard-timeout"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(i["stp-bpdu-guard-timeout"], d, pre_append)
			tmp["stp_bpdu_guard_timeout"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-StpBpduGuardTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if _, ok := i["stp-root-guard"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsStpRootGuard(i["stp-root-guard"], d, pre_append)
			tmp["stp_root_guard"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-StpRootGuard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if _, ok := i["stp-state"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsStpState(i["stp-state"], d, pre_append)
			tmp["stp_state"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-StpState")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trunk_member"
		if _, ok := i["trunk-member"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsTrunkMember(i["trunk-member"], d, pre_append)
			tmp["trunk_member"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-TrunkMember")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := i["untagged-vlans"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsUntaggedVlans(i["untagged-vlans"], d, pre_append)
			tmp["untagged_vlans"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-UntaggedVlans")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-Ports-Vlan")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerManagedSwitchPortsAccessMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAclGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerManagedSwitchPortsAggregatorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAllowedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAllowedVlansAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsArpInspectionTrust(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAuthenticatedPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsBundle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDiscardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDslProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsEdgePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsEncryptedPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFecCapable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFecState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapguard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlowControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFortiswitchAcls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsInterfaceTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIpSourceGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLacpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLearningLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLinkStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLldpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLldpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLoopGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMaxBundle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMclag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMclagIclPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMediaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerManagedSwitchPortsMinBundle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsP2PPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPacketSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPacketSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPauseMeter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPauseMeterResume(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeMaxPower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePortPower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePortPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeStandard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortOwner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsQosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsRpvstPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSampleDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSflowCounterInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSflowSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStickyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpRootGuard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsTrunkMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsUntaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPtpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchQosDropPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchQosRedProbability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchRouteOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadMclag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadRouter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_ip"
		if _, ok := i["router-ip"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(i["router-ip"], d, pre_append)
			tmp["router_ip"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-RouteOffloadRouter-RouterIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitch-RouteOffloadRouter-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchSwitchDhcpOpt43Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchSwitchId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchTdrSupported(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerManagedSwitch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_platform", flattenObjectSwitchControllerManagedSwitchPlatform(o["_platform"], d, "_platform")); err != nil {
		if vv, ok := fortiAPIPatch(o["_platform"], "ObjectSwitchControllerManagedSwitch-Platform"); ok {
			if err = d.Set("_platform", vv); err != nil {
				return fmt.Errorf("Error reading _platform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _platform: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_command", flattenObjectSwitchControllerManagedSwitchCustomCommand(o["custom-command"], d, "custom_command")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-command"], "ObjectSwitchControllerManagedSwitch-CustomCommand"); ok {
				if err = d.Set("custom_command", vv); err != nil {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_command: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_command"); ok {
			if err = d.Set("custom_command", flattenObjectSwitchControllerManagedSwitchCustomCommand(o["custom-command"], d, "custom_command")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-command"], "ObjectSwitchControllerManagedSwitch-CustomCommand"); ok {
					if err = d.Set("custom_command", vv); err != nil {
						return fmt.Errorf("Error reading custom_command: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerManagedSwitchDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerManagedSwitch-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", flattenObjectSwitchControllerManagedSwitchDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server-access-list"], "ObjectSwitchControllerManagedSwitch-DhcpServerAccessList"); ok {
			if err = d.Set("dhcp_server_access_list", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_snooping_static_client", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client")); err != nil {
			if vv, ok := fortiAPIPatch(o["dhcp-snooping-static-client"], "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient"); ok {
				if err = d.Set("dhcp_snooping_static_client", vv); err != nil {
					return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_snooping_static_client"); ok {
			if err = d.Set("dhcp_snooping_static_client", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client")); err != nil {
				if vv, ok := fortiAPIPatch(o["dhcp-snooping-static-client"], "ObjectSwitchControllerManagedSwitch-DhcpSnoopingStaticClient"); ok {
					if err = d.Set("dhcp_snooping_static_client", vv); err != nil {
						return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
				}
			}
		}
	}

	if err = d.Set("firmware_provision", flattenObjectSwitchControllerManagedSwitchFirmwareProvision(o["firmware-provision"], d, "firmware_provision")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision"], "ObjectSwitchControllerManagedSwitch-FirmwareProvision"); ok {
			if err = d.Set("firmware_provision", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision: %v", err)
		}
	}

	if err = d.Set("firmware_provision_latest", flattenObjectSwitchControllerManagedSwitchFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-latest"], "ObjectSwitchControllerManagedSwitch-FirmwareProvisionLatest"); ok {
			if err = d.Set("firmware_provision_latest", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	if err = d.Set("firmware_provision_version", flattenObjectSwitchControllerManagedSwitchFirmwareProvisionVersion(o["firmware-provision-version"], d, "firmware_provision_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-version"], "ObjectSwitchControllerManagedSwitch-FirmwareProvisionVersion"); ok {
			if err = d.Set("firmware_provision_version", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_version: %v", err)
		}
	}

	if err = d.Set("l3_discovered", flattenObjectSwitchControllerManagedSwitchL3Discovered(o["l3-discovered"], d, "l3_discovered")); err != nil {
		if vv, ok := fortiAPIPatch(o["l3-discovered"], "ObjectSwitchControllerManagedSwitch-L3Discovered"); ok {
			if err = d.Set("l3_discovered", vv); err != nil {
				return fmt.Errorf("Error reading l3_discovered: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l3_discovered: %v", err)
		}
	}

	if err = d.Set("mclag_igmp_snooping_aware", flattenObjectSwitchControllerManagedSwitchMclagIgmpSnoopingAware(o["mclag-igmp-snooping-aware"], d, "mclag_igmp_snooping_aware")); err != nil {
		if vv, ok := fortiAPIPatch(o["mclag-igmp-snooping-aware"], "ObjectSwitchControllerManagedSwitch-MclagIgmpSnoopingAware"); ok {
			if err = d.Set("mclag_igmp_snooping_aware", vv); err != nil {
				return fmt.Errorf("Error reading mclag_igmp_snooping_aware: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mclag_igmp_snooping_aware: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerManagedSwitchName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerManagedSwitch-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_snmp_community", flattenObjectSwitchControllerManagedSwitchOverrideSnmpCommunity(o["override-snmp-community"], d, "override_snmp_community")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-community"], "ObjectSwitchControllerManagedSwitch-OverrideSnmpCommunity"); ok {
			if err = d.Set("override_snmp_community", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_community: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_community: %v", err)
		}
	}

	if err = d.Set("override_snmp_sysinfo", flattenObjectSwitchControllerManagedSwitchOverrideSnmpSysinfo(o["override-snmp-sysinfo"], d, "override_snmp_sysinfo")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-sysinfo"], "ObjectSwitchControllerManagedSwitch-OverrideSnmpSysinfo"); ok {
			if err = d.Set("override_snmp_sysinfo", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_sysinfo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_sysinfo: %v", err)
		}
	}

	if err = d.Set("override_snmp_trap_threshold", flattenObjectSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(o["override-snmp-trap-threshold"], d, "override_snmp_trap_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-trap-threshold"], "ObjectSwitchControllerManagedSwitch-OverrideSnmpTrapThreshold"); ok {
			if err = d.Set("override_snmp_trap_threshold", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_trap_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_trap_threshold: %v", err)
		}
	}

	if err = d.Set("override_snmp_user", flattenObjectSwitchControllerManagedSwitchOverrideSnmpUser(o["override-snmp-user"], d, "override_snmp_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-snmp-user"], "ObjectSwitchControllerManagedSwitch-OverrideSnmpUser"); ok {
			if err = d.Set("override_snmp_user", vv); err != nil {
				return fmt.Errorf("Error reading override_snmp_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_snmp_user: %v", err)
		}
	}

	if err = d.Set("poe_detection_type", flattenObjectSwitchControllerManagedSwitchPoeDetectionType(o["poe-detection-type"], d, "poe_detection_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-detection-type"], "ObjectSwitchControllerManagedSwitch-PoeDetectionType"); ok {
			if err = d.Set("poe_detection_type", vv); err != nil {
				return fmt.Errorf("Error reading poe_detection_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_detection_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ports", flattenObjectSwitchControllerManagedSwitchPorts(o["ports"], d, "ports")); err != nil {
			if vv, ok := fortiAPIPatch(o["ports"], "ObjectSwitchControllerManagedSwitch-Ports"); ok {
				if err = d.Set("ports", vv); err != nil {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ports"); ok {
			if err = d.Set("ports", flattenObjectSwitchControllerManagedSwitchPorts(o["ports"], d, "ports")); err != nil {
				if vv, ok := fortiAPIPatch(o["ports"], "ObjectSwitchControllerManagedSwitch-Ports"); ok {
					if err = d.Set("ports", vv); err != nil {
						return fmt.Errorf("Error reading ports: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("ptp_profile", flattenObjectSwitchControllerManagedSwitchPtpProfile(o["ptp-profile"], d, "ptp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptp-profile"], "ObjectSwitchControllerManagedSwitch-PtpProfile"); ok {
			if err = d.Set("ptp_profile", vv); err != nil {
				return fmt.Errorf("Error reading ptp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptp_profile: %v", err)
		}
	}

	if err = d.Set("ptp_status", flattenObjectSwitchControllerManagedSwitchPtpStatus(o["ptp-status"], d, "ptp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptp-status"], "ObjectSwitchControllerManagedSwitch-PtpStatus"); ok {
			if err = d.Set("ptp_status", vv); err != nil {
				return fmt.Errorf("Error reading ptp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptp_status: %v", err)
		}
	}

	if err = d.Set("qos_drop_policy", flattenObjectSwitchControllerManagedSwitchQosDropPolicy(o["qos-drop-policy"], d, "qos_drop_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-drop-policy"], "ObjectSwitchControllerManagedSwitch-QosDropPolicy"); ok {
			if err = d.Set("qos_drop_policy", vv); err != nil {
				return fmt.Errorf("Error reading qos_drop_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_drop_policy: %v", err)
		}
	}

	if err = d.Set("qos_red_probability", flattenObjectSwitchControllerManagedSwitchQosRedProbability(o["qos-red-probability"], d, "qos_red_probability")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-red-probability"], "ObjectSwitchControllerManagedSwitch-QosRedProbability"); ok {
			if err = d.Set("qos_red_probability", vv); err != nil {
				return fmt.Errorf("Error reading qos_red_probability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_red_probability: %v", err)
		}
	}

	if err = d.Set("route_offload", flattenObjectSwitchControllerManagedSwitchRouteOffload(o["route-offload"], d, "route_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-offload"], "ObjectSwitchControllerManagedSwitch-RouteOffload"); ok {
			if err = d.Set("route_offload", vv); err != nil {
				return fmt.Errorf("Error reading route_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_offload: %v", err)
		}
	}

	if err = d.Set("route_offload_mclag", flattenObjectSwitchControllerManagedSwitchRouteOffloadMclag(o["route-offload-mclag"], d, "route_offload_mclag")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-offload-mclag"], "ObjectSwitchControllerManagedSwitch-RouteOffloadMclag"); ok {
			if err = d.Set("route_offload_mclag", vv); err != nil {
				return fmt.Errorf("Error reading route_offload_mclag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_offload_mclag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("route_offload_router", flattenObjectSwitchControllerManagedSwitchRouteOffloadRouter(o["route-offload-router"], d, "route_offload_router")); err != nil {
			if vv, ok := fortiAPIPatch(o["route-offload-router"], "ObjectSwitchControllerManagedSwitch-RouteOffloadRouter"); ok {
				if err = d.Set("route_offload_router", vv); err != nil {
					return fmt.Errorf("Error reading route_offload_router: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route_offload_router: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_offload_router"); ok {
			if err = d.Set("route_offload_router", flattenObjectSwitchControllerManagedSwitchRouteOffloadRouter(o["route-offload-router"], d, "route_offload_router")); err != nil {
				if vv, ok := fortiAPIPatch(o["route-offload-router"], "ObjectSwitchControllerManagedSwitch-RouteOffloadRouter"); ok {
					if err = d.Set("route_offload_router", vv); err != nil {
						return fmt.Errorf("Error reading route_offload_router: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route_offload_router: %v", err)
				}
			}
		}
	}

	if err = d.Set("switch_dhcp_opt43_key", flattenObjectSwitchControllerManagedSwitchSwitchDhcpOpt43Key(o["switch-dhcp_opt43_key"], d, "switch_dhcp_opt43_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-dhcp_opt43_key"], "ObjectSwitchControllerManagedSwitch-SwitchDhcpOpt43Key"); ok {
			if err = d.Set("switch_dhcp_opt43_key", vv); err != nil {
				return fmt.Errorf("Error reading switch_dhcp_opt43_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_dhcp_opt43_key: %v", err)
		}
	}

	if err = d.Set("switch_id", flattenObjectSwitchControllerManagedSwitchSwitchId(o["switch-id"], d, "switch_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-id"], "ObjectSwitchControllerManagedSwitch-SwitchId"); ok {
			if err = d.Set("switch_id", vv); err != nil {
				return fmt.Errorf("Error reading switch_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	if err = d.Set("tdr_supported", flattenObjectSwitchControllerManagedSwitchTdrSupported(o["tdr-supported"], d, "tdr_supported")); err != nil {
		if vv, ok := fortiAPIPatch(o["tdr-supported"], "ObjectSwitchControllerManagedSwitch-TdrSupported"); ok {
			if err = d.Set("tdr_supported", vv); err != nil {
				return fmt.Errorf("Error reading tdr_supported: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tdr_supported: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerManagedSwitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerManagedSwitchPlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["command-entry"], _ = expandObjectSwitchControllerManagedSwitchCustomCommandCommandEntry(d, i["command_entry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["command-name"], _ = expandObjectSwitchControllerManagedSwitchCustomCommandCommandName(d, i["command_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerManagedSwitchCustomCommandCommandEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(d, i["vlan"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchFirmwareProvision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchFirmwareProvisionVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchL3Discovered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchMclagIgmpSnoopingAware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchOverrideSnmpCommunity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchOverrideSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchOverrideSnmpUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPoeDetectionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["access-mode"], _ = expandObjectSwitchControllerManagedSwitchPortsAccessMode(d, i["access_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acl-group"], _ = expandObjectSwitchControllerManagedSwitchPortsAclGroup(d, i["acl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["aggregator-mode"], _ = expandObjectSwitchControllerManagedSwitchPortsAggregatorMode(d, i["aggregator_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-vlans"], _ = expandObjectSwitchControllerManagedSwitchPortsAllowedVlans(d, i["allowed_vlans"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-vlans-all"], _ = expandObjectSwitchControllerManagedSwitchPortsAllowedVlansAll(d, i["allowed_vlans_all"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-inspection-trust"], _ = expandObjectSwitchControllerManagedSwitchPortsArpInspectionTrust(d, i["arp_inspection_trust"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authenticated_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authenticated-port"], _ = expandObjectSwitchControllerManagedSwitchPortsAuthenticatedPort(d, i["authenticated_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bundle"], _ = expandObjectSwitchControllerManagedSwitchPortsBundle(d, i["bundle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectSwitchControllerManagedSwitchPortsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d, i["dhcp_snoop_option82_override"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["dhcp-snoop-option82-override"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-snoop-option82-trust"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d, i["dhcp_snoop_option82_trust"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-snooping"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnooping(d, i["dhcp_snooping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["discard-mode"], _ = expandObjectSwitchControllerManagedSwitchPortsDiscardMode(d, i["discard_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dsl_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dsl-profile"], _ = expandObjectSwitchControllerManagedSwitchPortsDslProfile(d, i["dsl_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["edge-port"], _ = expandObjectSwitchControllerManagedSwitchPortsEdgePort(d, i["edge_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypted_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encrypted-port"], _ = expandObjectSwitchControllerManagedSwitchPortsEncryptedPort(d, i["encrypted_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fec-capable"], _ = expandObjectSwitchControllerManagedSwitchPortsFecCapable(d, i["fec_capable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fec-state"], _ = expandObjectSwitchControllerManagedSwitchPortsFecState(d, i["fec_state"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flap-duration"], _ = expandObjectSwitchControllerManagedSwitchPortsFlapDuration(d, i["flap_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flap-rate"], _ = expandObjectSwitchControllerManagedSwitchPortsFlapRate(d, i["flap_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flap-timeout"], _ = expandObjectSwitchControllerManagedSwitchPortsFlapTimeout(d, i["flap_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flapguard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flapguard"], _ = expandObjectSwitchControllerManagedSwitchPortsFlapguard(d, i["flapguard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flow-control"], _ = expandObjectSwitchControllerManagedSwitchPortsFlowControl(d, i["flow_control"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiswitch_acls"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiswitch-acls"], _ = expandObjectSwitchControllerManagedSwitchPortsFortiswitchAcls(d, i["fortiswitch_acls"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping_flood_reports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmp-snooping-flood-reports"], _ = expandObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(d, i["igmp_snooping_flood_reports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmp-snooping"], _ = expandObjectSwitchControllerManagedSwitchPortsIgmpSnooping(d, i["igmp_snooping"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmps-flood-reports"], _ = expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports(d, i["igmps_flood_reports"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["igmps-flood-traffic"], _ = expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(d, i["igmps_flood_traffic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-tags"], _ = expandObjectSwitchControllerManagedSwitchPortsInterfaceTags(d, i["interface_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-source-guard"], _ = expandObjectSwitchControllerManagedSwitchPortsIpSourceGuard(d, i["ip_source_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_sn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["isl-peer-device-sn"], _ = expandObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn(d, i["isl_peer_device_sn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lacp-speed"], _ = expandObjectSwitchControllerManagedSwitchPortsLacpSpeed(d, i["lacp_speed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["learning-limit"], _ = expandObjectSwitchControllerManagedSwitchPortsLearningLimit(d, i["learning_limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-status"], _ = expandObjectSwitchControllerManagedSwitchPortsLinkStatus(d, i["link_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lldp-profile"], _ = expandObjectSwitchControllerManagedSwitchPortsLldpProfile(d, i["lldp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lldp-status"], _ = expandObjectSwitchControllerManagedSwitchPortsLldpStatus(d, i["lldp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["loop-guard"], _ = expandObjectSwitchControllerManagedSwitchPortsLoopGuard(d, i["loop_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["loop-guard-timeout"], _ = expandObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout(d, i["loop_guard_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_intf_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["matched-dpp-intf-tags"], _ = expandObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags(d, i["matched_dpp_intf_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["matched-dpp-policy"], _ = expandObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy(d, i["matched_dpp_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max-bundle"], _ = expandObjectSwitchControllerManagedSwitchPortsMaxBundle(d, i["max_bundle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_snooping_flood_traffic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mcast-snooping-flood-traffic"], _ = expandObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(d, i["mcast_snooping_flood_traffic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mclag"], _ = expandObjectSwitchControllerManagedSwitchPortsMclag(d, i["mclag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mclag-icl-port"], _ = expandObjectSwitchControllerManagedSwitchPortsMclagIclPort(d, i["mclag_icl_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["media-type"], _ = expandObjectSwitchControllerManagedSwitchPortsMediaType(d, i["media_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member-withdrawal-behavior"], _ = expandObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(d, i["member_withdrawal_behavior"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandObjectSwitchControllerManagedSwitchPortsMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min-bundle"], _ = expandObjectSwitchControllerManagedSwitchPortsMinBundle(d, i["min_bundle"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandObjectSwitchControllerManagedSwitchPortsMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["p2p-port"], _ = expandObjectSwitchControllerManagedSwitchPortsP2PPort(d, i["p2p_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-sample-rate"], _ = expandObjectSwitchControllerManagedSwitchPortsPacketSampleRate(d, i["packet_sample_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-sampler"], _ = expandObjectSwitchControllerManagedSwitchPortsPacketSampler(d, i["packet_sampler"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pause-meter"], _ = expandObjectSwitchControllerManagedSwitchPortsPauseMeter(d, i["pause_meter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pause-meter-resume"], _ = expandObjectSwitchControllerManagedSwitchPortsPauseMeterResume(d, i["pause_meter_resume"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_max_power"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-max-power"], _ = expandObjectSwitchControllerManagedSwitchPortsPoeMaxPower(d, i["poe_max_power"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_mode_bt_cabable"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-mode-bt-cabable"], _ = expandObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable(d, i["poe_mode_bt_cabable"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-port-mode"], _ = expandObjectSwitchControllerManagedSwitchPortsPoePortMode(d, i["poe_port_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_power"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-port-power"], _ = expandObjectSwitchControllerManagedSwitchPortsPoePortPower(d, i["poe_port_power"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-port-priority"], _ = expandObjectSwitchControllerManagedSwitchPortsPoePortPriority(d, i["poe_port_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-pre-standard-detection"], _ = expandObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection(d, i["poe_pre_standard_detection"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_standard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-standard"], _ = expandObjectSwitchControllerManagedSwitchPortsPoeStandard(d, i["poe_standard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poe-status"], _ = expandObjectSwitchControllerManagedSwitchPortsPoeStatus(d, i["poe_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-name"], _ = expandObjectSwitchControllerManagedSwitchPortsPortName(d, i["port_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-owner"], _ = expandObjectSwitchControllerManagedSwitchPortsPortOwner(d, i["port_owner"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-policy"], _ = expandObjectSwitchControllerManagedSwitchPortsPortPolicy(d, i["port_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-security-policy"], _ = expandObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy(d, i["port_security_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-selection-criteria"], _ = expandObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria(d, i["port_selection_criteria"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ptp-status"], _ = expandObjectSwitchControllerManagedSwitchPortsPtpStatus(d, i["ptp_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qos-policy"], _ = expandObjectSwitchControllerManagedSwitchPortsQosPolicy(d, i["qos_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_auth_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-auth-port"], _ = expandObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort(d, i["restricted_auth_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rpvst-port"], _ = expandObjectSwitchControllerManagedSwitchPortsRpvstPort(d, i["rpvst_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sample-direction"], _ = expandObjectSwitchControllerManagedSwitchPortsSampleDirection(d, i["sample_direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sflow-counter-interval"], _ = expandObjectSwitchControllerManagedSwitchPortsSflowCounterInterval(d, i["sflow_counter_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sample_rate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sflow-sample-rate"], _ = expandObjectSwitchControllerManagedSwitchPortsSflowSampleRate(d, i["sflow_sample_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sampler"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sflow-sampler"], _ = expandObjectSwitchControllerManagedSwitchPortsSflowSampler(d, i["sflow_sampler"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSwitchControllerManagedSwitchPortsStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sticky-mac"], _ = expandObjectSwitchControllerManagedSwitchPortsStickyMac(d, i["sticky_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-bpdu-guard"], _ = expandObjectSwitchControllerManagedSwitchPortsStpBpduGuard(d, i["stp_bpdu_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-bpdu-guard-timeout"], _ = expandObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(d, i["stp_bpdu_guard_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-root-guard"], _ = expandObjectSwitchControllerManagedSwitchPortsStpRootGuard(d, i["stp_root_guard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["stp-state"], _ = expandObjectSwitchControllerManagedSwitchPortsStpState(d, i["stp_state"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trunk_member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["trunk-member"], _ = expandObjectSwitchControllerManagedSwitchPortsTrunkMember(d, i["trunk_member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSwitchControllerManagedSwitchPortsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["untagged-vlans"], _ = expandObjectSwitchControllerManagedSwitchPortsUntaggedVlans(d, i["untagged_vlans"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandObjectSwitchControllerManagedSwitchPortsVlan(d, i["vlan"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAccessMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAclGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerManagedSwitchPortsAggregatorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAllowedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAllowedVlansAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsArpInspectionTrust(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAuthenticatedPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsBundle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDiscardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDslProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsEdgePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsEncryptedPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFecCapable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFecState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlowControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFortiswitchAcls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsInterfaceTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIpSourceGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLacpSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLearningLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLinkStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLldpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLldpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLoopGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMaxBundle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMclag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMclagIclPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMediaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerManagedSwitchPortsMinBundle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsP2PPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPacketSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPacketSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPauseMeter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPauseMeterResume(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeMaxPower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePortPower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePortPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeStandard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortOwner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsQosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsRpvstPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSampleDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSflowCounterInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSflowSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSflowSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStickyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpBpduGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpRootGuard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsTrunkMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsUntaggedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPtpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchQosDropPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchQosRedProbability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchRouteOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchRouteOffloadMclag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["router-ip"], _ = expandObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(d, i["router_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchSwitchDhcpOpt43Key(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchSwitchId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchTdrSupported(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerManagedSwitch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_platform"); ok || d.HasChange("_platform") {
		t, err := expandObjectSwitchControllerManagedSwitchPlatform(d, v, "_platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_platform"] = t
		}
	}

	if v, ok := d.GetOk("custom_command"); ok || d.HasChange("custom_command") {
		t, err := expandObjectSwitchControllerManagedSwitchCustomCommand(d, v, "custom_command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-command"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerManagedSwitchDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok || d.HasChange("dhcp_server_access_list") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpServerAccessList(d, v, "dhcp_server_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-access-list"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_static_client"); ok || d.HasChange("dhcp_snooping_static_client") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d, v, "dhcp_snooping_static_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping-static-client"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision"); ok || d.HasChange("firmware_provision") {
		t, err := expandObjectSwitchControllerManagedSwitchFirmwareProvision(d, v, "firmware_provision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok || d.HasChange("firmware_provision_latest") {
		t, err := expandObjectSwitchControllerManagedSwitchFirmwareProvisionLatest(d, v, "firmware_provision_latest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_version"); ok || d.HasChange("firmware_provision_version") {
		t, err := expandObjectSwitchControllerManagedSwitchFirmwareProvisionVersion(d, v, "firmware_provision_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-version"] = t
		}
	}

	if v, ok := d.GetOk("l3_discovered"); ok || d.HasChange("l3_discovered") {
		t, err := expandObjectSwitchControllerManagedSwitchL3Discovered(d, v, "l3_discovered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-discovered"] = t
		}
	}

	if v, ok := d.GetOk("mclag_igmp_snooping_aware"); ok || d.HasChange("mclag_igmp_snooping_aware") {
		t, err := expandObjectSwitchControllerManagedSwitchMclagIgmpSnoopingAware(d, v, "mclag_igmp_snooping_aware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag-igmp-snooping-aware"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerManagedSwitchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_community"); ok || d.HasChange("override_snmp_community") {
		t, err := expandObjectSwitchControllerManagedSwitchOverrideSnmpCommunity(d, v, "override_snmp_community")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-community"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_sysinfo"); ok || d.HasChange("override_snmp_sysinfo") {
		t, err := expandObjectSwitchControllerManagedSwitchOverrideSnmpSysinfo(d, v, "override_snmp_sysinfo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_trap_threshold"); ok || d.HasChange("override_snmp_trap_threshold") {
		t, err := expandObjectSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(d, v, "override_snmp_trap_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-trap-threshold"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_user"); ok || d.HasChange("override_snmp_user") {
		t, err := expandObjectSwitchControllerManagedSwitchOverrideSnmpUser(d, v, "override_snmp_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-user"] = t
		}
	}

	if v, ok := d.GetOk("poe_detection_type"); ok || d.HasChange("poe_detection_type") {
		t, err := expandObjectSwitchControllerManagedSwitchPoeDetectionType(d, v, "poe_detection_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-detection-type"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandObjectSwitchControllerManagedSwitchPorts(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("ptp_profile"); ok || d.HasChange("ptp_profile") {
		t, err := expandObjectSwitchControllerManagedSwitchPtpProfile(d, v, "ptp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-profile"] = t
		}
	}

	if v, ok := d.GetOk("ptp_status"); ok || d.HasChange("ptp_status") {
		t, err := expandObjectSwitchControllerManagedSwitchPtpStatus(d, v, "ptp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-status"] = t
		}
	}

	if v, ok := d.GetOk("qos_drop_policy"); ok || d.HasChange("qos_drop_policy") {
		t, err := expandObjectSwitchControllerManagedSwitchQosDropPolicy(d, v, "qos_drop_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-drop-policy"] = t
		}
	}

	if v, ok := d.GetOk("qos_red_probability"); ok || d.HasChange("qos_red_probability") {
		t, err := expandObjectSwitchControllerManagedSwitchQosRedProbability(d, v, "qos_red_probability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-red-probability"] = t
		}
	}

	if v, ok := d.GetOk("route_offload"); ok || d.HasChange("route_offload") {
		t, err := expandObjectSwitchControllerManagedSwitchRouteOffload(d, v, "route_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload"] = t
		}
	}

	if v, ok := d.GetOk("route_offload_mclag"); ok || d.HasChange("route_offload_mclag") {
		t, err := expandObjectSwitchControllerManagedSwitchRouteOffloadMclag(d, v, "route_offload_mclag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload-mclag"] = t
		}
	}

	if v, ok := d.GetOk("route_offload_router"); ok || d.HasChange("route_offload_router") {
		t, err := expandObjectSwitchControllerManagedSwitchRouteOffloadRouter(d, v, "route_offload_router")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload-router"] = t
		}
	}

	if v, ok := d.GetOk("switch_dhcp_opt43_key"); ok || d.HasChange("switch_dhcp_opt43_key") {
		t, err := expandObjectSwitchControllerManagedSwitchSwitchDhcpOpt43Key(d, v, "switch_dhcp_opt43_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-dhcp_opt43_key"] = t
		}
	}

	if v, ok := d.GetOk("switch_id"); ok || d.HasChange("switch_id") {
		t, err := expandObjectSwitchControllerManagedSwitchSwitchId(d, v, "switch_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	if v, ok := d.GetOk("tdr_supported"); ok || d.HasChange("tdr_supported") {
		t, err := expandObjectSwitchControllerManagedSwitchTdrSupported(d, v, "tdr_supported")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tdr-supported"] = t
		}
	}

	return &obj, nil
}
