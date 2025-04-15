// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Managed-switch port list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerManagedSwitchPorts() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerManagedSwitchPortsCreate,
		Read:   resourceObjectSwitchControllerManagedSwitchPortsRead,
		Update: resourceObjectSwitchControllerManagedSwitchPortsUpdate,
		Delete: resourceObjectSwitchControllerManagedSwitchPortsDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectSwitchControllerManagedSwitchPortsCreate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchPorts(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchPorts resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerManagedSwitchPorts(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchPorts resource: %v", err)
	}

	d.SetId(getStringKey(d, "port_name"))

	return resourceObjectSwitchControllerManagedSwitchPortsRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchPortsUpdate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchPorts(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchPorts resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerManagedSwitchPorts(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchPorts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "port_name"))

	return resourceObjectSwitchControllerManagedSwitchPortsRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchPortsDelete(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerManagedSwitchPorts(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerManagedSwitchPorts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerManagedSwitchPortsRead(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadObjectSwitchControllerManagedSwitchPorts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchPorts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerManagedSwitchPorts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchPorts resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerManagedSwitchPortsAccessMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAclGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerManagedSwitchPortsAggregatorMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAllowedVlans2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectSwitchControllerManagedSwitchPortsAllowedVlansAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsArpInspectionTrust2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsAuthenticatedPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsBundle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId2edl(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId2edl(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {
			v := flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName2edl(i["vlan-name"], d, pre_append)
			tmp["vlan_name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override-VlanName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDhcpSnooping2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDiscardMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsDslProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsEdgePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsEncryptedPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFecCapable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFecState2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlapguard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFlowControl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsFortiswitchAcls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpSnooping2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsInterfaceTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectSwitchControllerManagedSwitchPortsIpSourceGuard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLacpSpeed2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLearningLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLinkStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLldpProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLldpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLoopGuard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMaxBundle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMclag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMclagIclPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMediaType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerManagedSwitchPortsMinBundle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsP2PPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPacketSampleRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPacketSampler2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPauseMeter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPauseMeterResume2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeMaxPower2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePortMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePortPower2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePortPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeStandard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPoeStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortOwner2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsPtpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsQosPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsRpvstPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSampleDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSflowCounterInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSflowSampleRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsSflowSampler2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStickyMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpRootGuard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsStpState2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsTrunkMember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchPortsUntaggedVlans2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectSwitchControllerManagedSwitchPortsVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerManagedSwitchPorts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("access_mode", flattenObjectSwitchControllerManagedSwitchPortsAccessMode2edl(o["access-mode"], d, "access_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-mode"], "ObjectSwitchControllerManagedSwitchPorts-AccessMode"); ok {
			if err = d.Set("access_mode", vv); err != nil {
				return fmt.Errorf("Error reading access_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_mode: %v", err)
		}
	}

	if err = d.Set("acl_group", flattenObjectSwitchControllerManagedSwitchPortsAclGroup2edl(o["acl-group"], d, "acl_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["acl-group"], "ObjectSwitchControllerManagedSwitchPorts-AclGroup"); ok {
			if err = d.Set("acl_group", vv); err != nil {
				return fmt.Errorf("Error reading acl_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acl_group: %v", err)
		}
	}

	if err = d.Set("aggregator_mode", flattenObjectSwitchControllerManagedSwitchPortsAggregatorMode2edl(o["aggregator-mode"], d, "aggregator_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregator-mode"], "ObjectSwitchControllerManagedSwitchPorts-AggregatorMode"); ok {
			if err = d.Set("aggregator_mode", vv); err != nil {
				return fmt.Errorf("Error reading aggregator_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregator_mode: %v", err)
		}
	}

	if err = d.Set("allowed_vlans", flattenObjectSwitchControllerManagedSwitchPortsAllowedVlans2edl(o["allowed-vlans"], d, "allowed_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-vlans"], "ObjectSwitchControllerManagedSwitchPorts-AllowedVlans"); ok {
			if err = d.Set("allowed_vlans", vv); err != nil {
				return fmt.Errorf("Error reading allowed_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_vlans: %v", err)
		}
	}

	if err = d.Set("allowed_vlans_all", flattenObjectSwitchControllerManagedSwitchPortsAllowedVlansAll2edl(o["allowed-vlans-all"], d, "allowed_vlans_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-vlans-all"], "ObjectSwitchControllerManagedSwitchPorts-AllowedVlansAll"); ok {
			if err = d.Set("allowed_vlans_all", vv); err != nil {
				return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
		}
	}

	if err = d.Set("arp_inspection_trust", flattenObjectSwitchControllerManagedSwitchPortsArpInspectionTrust2edl(o["arp-inspection-trust"], d, "arp_inspection_trust")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-inspection-trust"], "ObjectSwitchControllerManagedSwitchPorts-ArpInspectionTrust"); ok {
			if err = d.Set("arp_inspection_trust", vv); err != nil {
				return fmt.Errorf("Error reading arp_inspection_trust: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_inspection_trust: %v", err)
		}
	}

	if err = d.Set("authenticated_port", flattenObjectSwitchControllerManagedSwitchPortsAuthenticatedPort2edl(o["authenticated-port"], d, "authenticated_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["authenticated-port"], "ObjectSwitchControllerManagedSwitchPorts-AuthenticatedPort"); ok {
			if err = d.Set("authenticated_port", vv); err != nil {
				return fmt.Errorf("Error reading authenticated_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authenticated_port: %v", err)
		}
	}

	if err = d.Set("bundle", flattenObjectSwitchControllerManagedSwitchPortsBundle2edl(o["bundle"], d, "bundle")); err != nil {
		if vv, ok := fortiAPIPatch(o["bundle"], "ObjectSwitchControllerManagedSwitchPorts-Bundle"); ok {
			if err = d.Set("bundle", vv); err != nil {
				return fmt.Errorf("Error reading bundle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bundle: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerManagedSwitchPortsDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerManagedSwitchPorts-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_snoop_option82_override", flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override2edl(o["dhcp-snoop-option82-override"], d, "dhcp_snoop_option82_override")); err != nil {
			if vv, ok := fortiAPIPatch(o["dhcp-snoop-option82-override"], "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override"); ok {
				if err = d.Set("dhcp_snoop_option82_override", vv); err != nil {
					return fmt.Errorf("Error reading dhcp_snoop_option82_override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dhcp_snoop_option82_override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_snoop_option82_override"); ok {
			if err = d.Set("dhcp_snoop_option82_override", flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override2edl(o["dhcp-snoop-option82-override"], d, "dhcp_snoop_option82_override")); err != nil {
				if vv, ok := fortiAPIPatch(o["dhcp-snoop-option82-override"], "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Override"); ok {
					if err = d.Set("dhcp_snoop_option82_override", vv); err != nil {
						return fmt.Errorf("Error reading dhcp_snoop_option82_override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dhcp_snoop_option82_override: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp_snoop_option82_trust", flattenObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust2edl(o["dhcp-snoop-option82-trust"], d, "dhcp_snoop_option82_trust")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-snoop-option82-trust"], "ObjectSwitchControllerManagedSwitchPorts-DhcpSnoopOption82Trust"); ok {
			if err = d.Set("dhcp_snoop_option82_trust", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_snoop_option82_trust: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_snoop_option82_trust: %v", err)
		}
	}

	if err = d.Set("dhcp_snooping", flattenObjectSwitchControllerManagedSwitchPortsDhcpSnooping2edl(o["dhcp-snooping"], d, "dhcp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-snooping"], "ObjectSwitchControllerManagedSwitchPorts-DhcpSnooping"); ok {
			if err = d.Set("dhcp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_snooping: %v", err)
		}
	}

	if err = d.Set("discard_mode", flattenObjectSwitchControllerManagedSwitchPortsDiscardMode2edl(o["discard-mode"], d, "discard_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["discard-mode"], "ObjectSwitchControllerManagedSwitchPorts-DiscardMode"); ok {
			if err = d.Set("discard_mode", vv); err != nil {
				return fmt.Errorf("Error reading discard_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discard_mode: %v", err)
		}
	}

	if err = d.Set("dsl_profile", flattenObjectSwitchControllerManagedSwitchPortsDslProfile2edl(o["dsl-profile"], d, "dsl_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dsl-profile"], "ObjectSwitchControllerManagedSwitchPorts-DslProfile"); ok {
			if err = d.Set("dsl_profile", vv); err != nil {
				return fmt.Errorf("Error reading dsl_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dsl_profile: %v", err)
		}
	}

	if err = d.Set("edge_port", flattenObjectSwitchControllerManagedSwitchPortsEdgePort2edl(o["edge-port"], d, "edge_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["edge-port"], "ObjectSwitchControllerManagedSwitchPorts-EdgePort"); ok {
			if err = d.Set("edge_port", vv); err != nil {
				return fmt.Errorf("Error reading edge_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading edge_port: %v", err)
		}
	}

	if err = d.Set("encrypted_port", flattenObjectSwitchControllerManagedSwitchPortsEncryptedPort2edl(o["encrypted-port"], d, "encrypted_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["encrypted-port"], "ObjectSwitchControllerManagedSwitchPorts-EncryptedPort"); ok {
			if err = d.Set("encrypted_port", vv); err != nil {
				return fmt.Errorf("Error reading encrypted_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encrypted_port: %v", err)
		}
	}

	if err = d.Set("fec_capable", flattenObjectSwitchControllerManagedSwitchPortsFecCapable2edl(o["fec-capable"], d, "fec_capable")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-capable"], "ObjectSwitchControllerManagedSwitchPorts-FecCapable"); ok {
			if err = d.Set("fec_capable", vv); err != nil {
				return fmt.Errorf("Error reading fec_capable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_capable: %v", err)
		}
	}

	if err = d.Set("fec_state", flattenObjectSwitchControllerManagedSwitchPortsFecState2edl(o["fec-state"], d, "fec_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["fec-state"], "ObjectSwitchControllerManagedSwitchPorts-FecState"); ok {
			if err = d.Set("fec_state", vv); err != nil {
				return fmt.Errorf("Error reading fec_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fec_state: %v", err)
		}
	}

	if err = d.Set("flap_duration", flattenObjectSwitchControllerManagedSwitchPortsFlapDuration2edl(o["flap-duration"], d, "flap_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["flap-duration"], "ObjectSwitchControllerManagedSwitchPorts-FlapDuration"); ok {
			if err = d.Set("flap_duration", vv); err != nil {
				return fmt.Errorf("Error reading flap_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flap_duration: %v", err)
		}
	}

	if err = d.Set("flap_rate", flattenObjectSwitchControllerManagedSwitchPortsFlapRate2edl(o["flap-rate"], d, "flap_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["flap-rate"], "ObjectSwitchControllerManagedSwitchPorts-FlapRate"); ok {
			if err = d.Set("flap_rate", vv); err != nil {
				return fmt.Errorf("Error reading flap_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flap_rate: %v", err)
		}
	}

	if err = d.Set("flap_timeout", flattenObjectSwitchControllerManagedSwitchPortsFlapTimeout2edl(o["flap-timeout"], d, "flap_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["flap-timeout"], "ObjectSwitchControllerManagedSwitchPorts-FlapTimeout"); ok {
			if err = d.Set("flap_timeout", vv); err != nil {
				return fmt.Errorf("Error reading flap_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flap_timeout: %v", err)
		}
	}

	if err = d.Set("flapguard", flattenObjectSwitchControllerManagedSwitchPortsFlapguard2edl(o["flapguard"], d, "flapguard")); err != nil {
		if vv, ok := fortiAPIPatch(o["flapguard"], "ObjectSwitchControllerManagedSwitchPorts-Flapguard"); ok {
			if err = d.Set("flapguard", vv); err != nil {
				return fmt.Errorf("Error reading flapguard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flapguard: %v", err)
		}
	}

	if err = d.Set("flow_control", flattenObjectSwitchControllerManagedSwitchPortsFlowControl2edl(o["flow-control"], d, "flow_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["flow-control"], "ObjectSwitchControllerManagedSwitchPorts-FlowControl"); ok {
			if err = d.Set("flow_control", vv); err != nil {
				return fmt.Errorf("Error reading flow_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flow_control: %v", err)
		}
	}

	if err = d.Set("fortiswitch_acls", flattenObjectSwitchControllerManagedSwitchPortsFortiswitchAcls2edl(o["fortiswitch-acls"], d, "fortiswitch_acls")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiswitch-acls"], "ObjectSwitchControllerManagedSwitchPorts-FortiswitchAcls"); ok {
			if err = d.Set("fortiswitch_acls", vv); err != nil {
				return fmt.Errorf("Error reading fortiswitch_acls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiswitch_acls: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_flood_reports", flattenObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports2edl(o["igmp-snooping-flood-reports"], d, "igmp_snooping_flood_reports")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-snooping-flood-reports"], "ObjectSwitchControllerManagedSwitchPorts-IgmpSnoopingFloodReports"); ok {
			if err = d.Set("igmp_snooping_flood_reports", vv); err != nil {
				return fmt.Errorf("Error reading igmp_snooping_flood_reports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_snooping_flood_reports: %v", err)
		}
	}

	if err = d.Set("igmp_snooping", flattenObjectSwitchControllerManagedSwitchPortsIgmpSnooping2edl(o["igmp-snooping"], d, "igmp_snooping")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-snooping"], "ObjectSwitchControllerManagedSwitchPorts-IgmpSnooping"); ok {
			if err = d.Set("igmp_snooping", vv); err != nil {
				return fmt.Errorf("Error reading igmp_snooping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_snooping: %v", err)
		}
	}

	if err = d.Set("igmps_flood_reports", flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports2edl(o["igmps-flood-reports"], d, "igmps_flood_reports")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmps-flood-reports"], "ObjectSwitchControllerManagedSwitchPorts-IgmpsFloodReports"); ok {
			if err = d.Set("igmps_flood_reports", vv); err != nil {
				return fmt.Errorf("Error reading igmps_flood_reports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmps_flood_reports: %v", err)
		}
	}

	if err = d.Set("igmps_flood_traffic", flattenObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic2edl(o["igmps-flood-traffic"], d, "igmps_flood_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmps-flood-traffic"], "ObjectSwitchControllerManagedSwitchPorts-IgmpsFloodTraffic"); ok {
			if err = d.Set("igmps_flood_traffic", vv); err != nil {
				return fmt.Errorf("Error reading igmps_flood_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmps_flood_traffic: %v", err)
		}
	}

	if err = d.Set("interface_tags", flattenObjectSwitchControllerManagedSwitchPortsInterfaceTags2edl(o["interface-tags"], d, "interface_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-tags"], "ObjectSwitchControllerManagedSwitchPorts-InterfaceTags"); ok {
			if err = d.Set("interface_tags", vv); err != nil {
				return fmt.Errorf("Error reading interface_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_tags: %v", err)
		}
	}

	if err = d.Set("ip_source_guard", flattenObjectSwitchControllerManagedSwitchPortsIpSourceGuard2edl(o["ip-source-guard"], d, "ip_source_guard")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-source-guard"], "ObjectSwitchControllerManagedSwitchPorts-IpSourceGuard"); ok {
			if err = d.Set("ip_source_guard", vv); err != nil {
				return fmt.Errorf("Error reading ip_source_guard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_source_guard: %v", err)
		}
	}

	if err = d.Set("isl_peer_device_sn", flattenObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn2edl(o["isl-peer-device-sn"], d, "isl_peer_device_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["isl-peer-device-sn"], "ObjectSwitchControllerManagedSwitchPorts-IslPeerDeviceSn"); ok {
			if err = d.Set("isl_peer_device_sn", vv); err != nil {
				return fmt.Errorf("Error reading isl_peer_device_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isl_peer_device_sn: %v", err)
		}
	}

	if err = d.Set("lacp_speed", flattenObjectSwitchControllerManagedSwitchPortsLacpSpeed2edl(o["lacp-speed"], d, "lacp_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["lacp-speed"], "ObjectSwitchControllerManagedSwitchPorts-LacpSpeed"); ok {
			if err = d.Set("lacp_speed", vv); err != nil {
				return fmt.Errorf("Error reading lacp_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lacp_speed: %v", err)
		}
	}

	if err = d.Set("learning_limit", flattenObjectSwitchControllerManagedSwitchPortsLearningLimit2edl(o["learning-limit"], d, "learning_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["learning-limit"], "ObjectSwitchControllerManagedSwitchPorts-LearningLimit"); ok {
			if err = d.Set("learning_limit", vv); err != nil {
				return fmt.Errorf("Error reading learning_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learning_limit: %v", err)
		}
	}

	if err = d.Set("link_status", flattenObjectSwitchControllerManagedSwitchPortsLinkStatus2edl(o["link-status"], d, "link_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-status"], "ObjectSwitchControllerManagedSwitchPorts-LinkStatus"); ok {
			if err = d.Set("link_status", vv); err != nil {
				return fmt.Errorf("Error reading link_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_status: %v", err)
		}
	}

	if err = d.Set("lldp_profile", flattenObjectSwitchControllerManagedSwitchPortsLldpProfile2edl(o["lldp-profile"], d, "lldp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-profile"], "ObjectSwitchControllerManagedSwitchPorts-LldpProfile"); ok {
			if err = d.Set("lldp_profile", vv); err != nil {
				return fmt.Errorf("Error reading lldp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_profile: %v", err)
		}
	}

	if err = d.Set("lldp_status", flattenObjectSwitchControllerManagedSwitchPortsLldpStatus2edl(o["lldp-status"], d, "lldp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-status"], "ObjectSwitchControllerManagedSwitchPorts-LldpStatus"); ok {
			if err = d.Set("lldp_status", vv); err != nil {
				return fmt.Errorf("Error reading lldp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_status: %v", err)
		}
	}

	if err = d.Set("loop_guard", flattenObjectSwitchControllerManagedSwitchPortsLoopGuard2edl(o["loop-guard"], d, "loop_guard")); err != nil {
		if vv, ok := fortiAPIPatch(o["loop-guard"], "ObjectSwitchControllerManagedSwitchPorts-LoopGuard"); ok {
			if err = d.Set("loop_guard", vv); err != nil {
				return fmt.Errorf("Error reading loop_guard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loop_guard: %v", err)
		}
	}

	if err = d.Set("loop_guard_timeout", flattenObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout2edl(o["loop-guard-timeout"], d, "loop_guard_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["loop-guard-timeout"], "ObjectSwitchControllerManagedSwitchPorts-LoopGuardTimeout"); ok {
			if err = d.Set("loop_guard_timeout", vv); err != nil {
				return fmt.Errorf("Error reading loop_guard_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loop_guard_timeout: %v", err)
		}
	}

	if err = d.Set("matched_dpp_intf_tags", flattenObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags2edl(o["matched-dpp-intf-tags"], d, "matched_dpp_intf_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["matched-dpp-intf-tags"], "ObjectSwitchControllerManagedSwitchPorts-MatchedDppIntfTags"); ok {
			if err = d.Set("matched_dpp_intf_tags", vv); err != nil {
				return fmt.Errorf("Error reading matched_dpp_intf_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading matched_dpp_intf_tags: %v", err)
		}
	}

	if err = d.Set("matched_dpp_policy", flattenObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy2edl(o["matched-dpp-policy"], d, "matched_dpp_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["matched-dpp-policy"], "ObjectSwitchControllerManagedSwitchPorts-MatchedDppPolicy"); ok {
			if err = d.Set("matched_dpp_policy", vv); err != nil {
				return fmt.Errorf("Error reading matched_dpp_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading matched_dpp_policy: %v", err)
		}
	}

	if err = d.Set("max_bundle", flattenObjectSwitchControllerManagedSwitchPortsMaxBundle2edl(o["max-bundle"], d, "max_bundle")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-bundle"], "ObjectSwitchControllerManagedSwitchPorts-MaxBundle"); ok {
			if err = d.Set("max_bundle", vv); err != nil {
				return fmt.Errorf("Error reading max_bundle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_bundle: %v", err)
		}
	}

	if err = d.Set("mcast_snooping_flood_traffic", flattenObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic2edl(o["mcast-snooping-flood-traffic"], d, "mcast_snooping_flood_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcast-snooping-flood-traffic"], "ObjectSwitchControllerManagedSwitchPorts-McastSnoopingFloodTraffic"); ok {
			if err = d.Set("mcast_snooping_flood_traffic", vv); err != nil {
				return fmt.Errorf("Error reading mcast_snooping_flood_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcast_snooping_flood_traffic: %v", err)
		}
	}

	if err = d.Set("mclag", flattenObjectSwitchControllerManagedSwitchPortsMclag2edl(o["mclag"], d, "mclag")); err != nil {
		if vv, ok := fortiAPIPatch(o["mclag"], "ObjectSwitchControllerManagedSwitchPorts-Mclag"); ok {
			if err = d.Set("mclag", vv); err != nil {
				return fmt.Errorf("Error reading mclag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mclag: %v", err)
		}
	}

	if err = d.Set("mclag_icl_port", flattenObjectSwitchControllerManagedSwitchPortsMclagIclPort2edl(o["mclag-icl-port"], d, "mclag_icl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["mclag-icl-port"], "ObjectSwitchControllerManagedSwitchPorts-MclagIclPort"); ok {
			if err = d.Set("mclag_icl_port", vv); err != nil {
				return fmt.Errorf("Error reading mclag_icl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mclag_icl_port: %v", err)
		}
	}

	if err = d.Set("media_type", flattenObjectSwitchControllerManagedSwitchPortsMediaType2edl(o["media-type"], d, "media_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["media-type"], "ObjectSwitchControllerManagedSwitchPorts-MediaType"); ok {
			if err = d.Set("media_type", vv); err != nil {
				return fmt.Errorf("Error reading media_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading media_type: %v", err)
		}
	}

	if err = d.Set("member_withdrawal_behavior", flattenObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior2edl(o["member-withdrawal-behavior"], d, "member_withdrawal_behavior")); err != nil {
		if vv, ok := fortiAPIPatch(o["member-withdrawal-behavior"], "ObjectSwitchControllerManagedSwitchPorts-MemberWithdrawalBehavior"); ok {
			if err = d.Set("member_withdrawal_behavior", vv); err != nil {
				return fmt.Errorf("Error reading member_withdrawal_behavior: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member_withdrawal_behavior: %v", err)
		}
	}

	if err = d.Set("members", flattenObjectSwitchControllerManagedSwitchPortsMembers2edl(o["members"], d, "members")); err != nil {
		if vv, ok := fortiAPIPatch(o["members"], "ObjectSwitchControllerManagedSwitchPorts-Members"); ok {
			if err = d.Set("members", vv); err != nil {
				return fmt.Errorf("Error reading members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading members: %v", err)
		}
	}

	if err = d.Set("min_bundle", flattenObjectSwitchControllerManagedSwitchPortsMinBundle2edl(o["min-bundle"], d, "min_bundle")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-bundle"], "ObjectSwitchControllerManagedSwitchPorts-MinBundle"); ok {
			if err = d.Set("min_bundle", vv); err != nil {
				return fmt.Errorf("Error reading min_bundle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_bundle: %v", err)
		}
	}

	if err = d.Set("mode", flattenObjectSwitchControllerManagedSwitchPortsMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "ObjectSwitchControllerManagedSwitchPorts-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("p2p_port", flattenObjectSwitchControllerManagedSwitchPortsP2PPort2edl(o["p2p-port"], d, "p2p_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["p2p-port"], "ObjectSwitchControllerManagedSwitchPorts-P2PPort"); ok {
			if err = d.Set("p2p_port", vv); err != nil {
				return fmt.Errorf("Error reading p2p_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading p2p_port: %v", err)
		}
	}

	if err = d.Set("packet_sample_rate", flattenObjectSwitchControllerManagedSwitchPortsPacketSampleRate2edl(o["packet-sample-rate"], d, "packet_sample_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-sample-rate"], "ObjectSwitchControllerManagedSwitchPorts-PacketSampleRate"); ok {
			if err = d.Set("packet_sample_rate", vv); err != nil {
				return fmt.Errorf("Error reading packet_sample_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_sample_rate: %v", err)
		}
	}

	if err = d.Set("packet_sampler", flattenObjectSwitchControllerManagedSwitchPortsPacketSampler2edl(o["packet-sampler"], d, "packet_sampler")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-sampler"], "ObjectSwitchControllerManagedSwitchPorts-PacketSampler"); ok {
			if err = d.Set("packet_sampler", vv); err != nil {
				return fmt.Errorf("Error reading packet_sampler: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_sampler: %v", err)
		}
	}

	if err = d.Set("pause_meter", flattenObjectSwitchControllerManagedSwitchPortsPauseMeter2edl(o["pause-meter"], d, "pause_meter")); err != nil {
		if vv, ok := fortiAPIPatch(o["pause-meter"], "ObjectSwitchControllerManagedSwitchPorts-PauseMeter"); ok {
			if err = d.Set("pause_meter", vv); err != nil {
				return fmt.Errorf("Error reading pause_meter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pause_meter: %v", err)
		}
	}

	if err = d.Set("pause_meter_resume", flattenObjectSwitchControllerManagedSwitchPortsPauseMeterResume2edl(o["pause-meter-resume"], d, "pause_meter_resume")); err != nil {
		if vv, ok := fortiAPIPatch(o["pause-meter-resume"], "ObjectSwitchControllerManagedSwitchPorts-PauseMeterResume"); ok {
			if err = d.Set("pause_meter_resume", vv); err != nil {
				return fmt.Errorf("Error reading pause_meter_resume: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pause_meter_resume: %v", err)
		}
	}

	if err = d.Set("poe_max_power", flattenObjectSwitchControllerManagedSwitchPortsPoeMaxPower2edl(o["poe-max-power"], d, "poe_max_power")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-max-power"], "ObjectSwitchControllerManagedSwitchPorts-PoeMaxPower"); ok {
			if err = d.Set("poe_max_power", vv); err != nil {
				return fmt.Errorf("Error reading poe_max_power: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_max_power: %v", err)
		}
	}

	if err = d.Set("poe_mode_bt_cabable", flattenObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable2edl(o["poe-mode-bt-cabable"], d, "poe_mode_bt_cabable")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-mode-bt-cabable"], "ObjectSwitchControllerManagedSwitchPorts-PoeModeBtCabable"); ok {
			if err = d.Set("poe_mode_bt_cabable", vv); err != nil {
				return fmt.Errorf("Error reading poe_mode_bt_cabable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_mode_bt_cabable: %v", err)
		}
	}

	if err = d.Set("poe_port_mode", flattenObjectSwitchControllerManagedSwitchPortsPoePortMode2edl(o["poe-port-mode"], d, "poe_port_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-port-mode"], "ObjectSwitchControllerManagedSwitchPorts-PoePortMode"); ok {
			if err = d.Set("poe_port_mode", vv); err != nil {
				return fmt.Errorf("Error reading poe_port_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_port_mode: %v", err)
		}
	}

	if err = d.Set("poe_port_power", flattenObjectSwitchControllerManagedSwitchPortsPoePortPower2edl(o["poe-port-power"], d, "poe_port_power")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-port-power"], "ObjectSwitchControllerManagedSwitchPorts-PoePortPower"); ok {
			if err = d.Set("poe_port_power", vv); err != nil {
				return fmt.Errorf("Error reading poe_port_power: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_port_power: %v", err)
		}
	}

	if err = d.Set("poe_port_priority", flattenObjectSwitchControllerManagedSwitchPortsPoePortPriority2edl(o["poe-port-priority"], d, "poe_port_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-port-priority"], "ObjectSwitchControllerManagedSwitchPorts-PoePortPriority"); ok {
			if err = d.Set("poe_port_priority", vv); err != nil {
				return fmt.Errorf("Error reading poe_port_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_port_priority: %v", err)
		}
	}

	if err = d.Set("poe_pre_standard_detection", flattenObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection2edl(o["poe-pre-standard-detection"], d, "poe_pre_standard_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-pre-standard-detection"], "ObjectSwitchControllerManagedSwitchPorts-PoePreStandardDetection"); ok {
			if err = d.Set("poe_pre_standard_detection", vv); err != nil {
				return fmt.Errorf("Error reading poe_pre_standard_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_pre_standard_detection: %v", err)
		}
	}

	if err = d.Set("poe_standard", flattenObjectSwitchControllerManagedSwitchPortsPoeStandard2edl(o["poe-standard"], d, "poe_standard")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-standard"], "ObjectSwitchControllerManagedSwitchPorts-PoeStandard"); ok {
			if err = d.Set("poe_standard", vv); err != nil {
				return fmt.Errorf("Error reading poe_standard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_standard: %v", err)
		}
	}

	if err = d.Set("poe_status", flattenObjectSwitchControllerManagedSwitchPortsPoeStatus2edl(o["poe-status"], d, "poe_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["poe-status"], "ObjectSwitchControllerManagedSwitchPorts-PoeStatus"); ok {
			if err = d.Set("poe_status", vv); err != nil {
				return fmt.Errorf("Error reading poe_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poe_status: %v", err)
		}
	}

	if err = d.Set("port_name", flattenObjectSwitchControllerManagedSwitchPortsPortName2edl(o["port-name"], d, "port_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-name"], "ObjectSwitchControllerManagedSwitchPorts-PortName"); ok {
			if err = d.Set("port_name", vv); err != nil {
				return fmt.Errorf("Error reading port_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_name: %v", err)
		}
	}

	if err = d.Set("port_owner", flattenObjectSwitchControllerManagedSwitchPortsPortOwner2edl(o["port-owner"], d, "port_owner")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-owner"], "ObjectSwitchControllerManagedSwitchPorts-PortOwner"); ok {
			if err = d.Set("port_owner", vv); err != nil {
				return fmt.Errorf("Error reading port_owner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_owner: %v", err)
		}
	}

	if err = d.Set("port_policy", flattenObjectSwitchControllerManagedSwitchPortsPortPolicy2edl(o["port-policy"], d, "port_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-policy"], "ObjectSwitchControllerManagedSwitchPorts-PortPolicy"); ok {
			if err = d.Set("port_policy", vv); err != nil {
				return fmt.Errorf("Error reading port_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_policy: %v", err)
		}
	}

	if err = d.Set("port_security_policy", flattenObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy2edl(o["port-security-policy"], d, "port_security_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-security-policy"], "ObjectSwitchControllerManagedSwitchPorts-PortSecurityPolicy"); ok {
			if err = d.Set("port_security_policy", vv); err != nil {
				return fmt.Errorf("Error reading port_security_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_security_policy: %v", err)
		}
	}

	if err = d.Set("port_selection_criteria", flattenObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria2edl(o["port-selection-criteria"], d, "port_selection_criteria")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-selection-criteria"], "ObjectSwitchControllerManagedSwitchPorts-PortSelectionCriteria"); ok {
			if err = d.Set("port_selection_criteria", vv); err != nil {
				return fmt.Errorf("Error reading port_selection_criteria: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_selection_criteria: %v", err)
		}
	}

	if err = d.Set("ptp_status", flattenObjectSwitchControllerManagedSwitchPortsPtpStatus2edl(o["ptp-status"], d, "ptp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ptp-status"], "ObjectSwitchControllerManagedSwitchPorts-PtpStatus"); ok {
			if err = d.Set("ptp_status", vv); err != nil {
				return fmt.Errorf("Error reading ptp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ptp_status: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenObjectSwitchControllerManagedSwitchPortsQosPolicy2edl(o["qos-policy"], d, "qos_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-policy"], "ObjectSwitchControllerManagedSwitchPorts-QosPolicy"); ok {
			if err = d.Set("qos_policy", vv); err != nil {
				return fmt.Errorf("Error reading qos_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("restricted_auth_port", flattenObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort2edl(o["restricted-auth-port"], d, "restricted_auth_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted-auth-port"], "ObjectSwitchControllerManagedSwitchPorts-RestrictedAuthPort"); ok {
			if err = d.Set("restricted_auth_port", vv); err != nil {
				return fmt.Errorf("Error reading restricted_auth_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_auth_port: %v", err)
		}
	}

	if err = d.Set("rpvst_port", flattenObjectSwitchControllerManagedSwitchPortsRpvstPort2edl(o["rpvst-port"], d, "rpvst_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpvst-port"], "ObjectSwitchControllerManagedSwitchPorts-RpvstPort"); ok {
			if err = d.Set("rpvst_port", vv); err != nil {
				return fmt.Errorf("Error reading rpvst_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpvst_port: %v", err)
		}
	}

	if err = d.Set("sample_direction", flattenObjectSwitchControllerManagedSwitchPortsSampleDirection2edl(o["sample-direction"], d, "sample_direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["sample-direction"], "ObjectSwitchControllerManagedSwitchPorts-SampleDirection"); ok {
			if err = d.Set("sample_direction", vv); err != nil {
				return fmt.Errorf("Error reading sample_direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sample_direction: %v", err)
		}
	}

	if err = d.Set("sflow_counter_interval", flattenObjectSwitchControllerManagedSwitchPortsSflowCounterInterval2edl(o["sflow-counter-interval"], d, "sflow_counter_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sflow-counter-interval"], "ObjectSwitchControllerManagedSwitchPorts-SflowCounterInterval"); ok {
			if err = d.Set("sflow_counter_interval", vv); err != nil {
				return fmt.Errorf("Error reading sflow_counter_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sflow_counter_interval: %v", err)
		}
	}

	if err = d.Set("sflow_sample_rate", flattenObjectSwitchControllerManagedSwitchPortsSflowSampleRate2edl(o["sflow-sample-rate"], d, "sflow_sample_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["sflow-sample-rate"], "ObjectSwitchControllerManagedSwitchPorts-SflowSampleRate"); ok {
			if err = d.Set("sflow_sample_rate", vv); err != nil {
				return fmt.Errorf("Error reading sflow_sample_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sflow_sample_rate: %v", err)
		}
	}

	if err = d.Set("sflow_sampler", flattenObjectSwitchControllerManagedSwitchPortsSflowSampler2edl(o["sflow-sampler"], d, "sflow_sampler")); err != nil {
		if vv, ok := fortiAPIPatch(o["sflow-sampler"], "ObjectSwitchControllerManagedSwitchPorts-SflowSampler"); ok {
			if err = d.Set("sflow_sampler", vv); err != nil {
				return fmt.Errorf("Error reading sflow_sampler: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sflow_sampler: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSwitchControllerManagedSwitchPortsStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSwitchControllerManagedSwitchPorts-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sticky_mac", flattenObjectSwitchControllerManagedSwitchPortsStickyMac2edl(o["sticky-mac"], d, "sticky_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["sticky-mac"], "ObjectSwitchControllerManagedSwitchPorts-StickyMac"); ok {
			if err = d.Set("sticky_mac", vv); err != nil {
				return fmt.Errorf("Error reading sticky_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sticky_mac: %v", err)
		}
	}

	if err = d.Set("stp_bpdu_guard", flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuard2edl(o["stp-bpdu-guard"], d, "stp_bpdu_guard")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp-bpdu-guard"], "ObjectSwitchControllerManagedSwitchPorts-StpBpduGuard"); ok {
			if err = d.Set("stp_bpdu_guard", vv); err != nil {
				return fmt.Errorf("Error reading stp_bpdu_guard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp_bpdu_guard: %v", err)
		}
	}

	if err = d.Set("stp_bpdu_guard_timeout", flattenObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout2edl(o["stp-bpdu-guard-timeout"], d, "stp_bpdu_guard_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp-bpdu-guard-timeout"], "ObjectSwitchControllerManagedSwitchPorts-StpBpduGuardTimeout"); ok {
			if err = d.Set("stp_bpdu_guard_timeout", vv); err != nil {
				return fmt.Errorf("Error reading stp_bpdu_guard_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp_bpdu_guard_timeout: %v", err)
		}
	}

	if err = d.Set("stp_root_guard", flattenObjectSwitchControllerManagedSwitchPortsStpRootGuard2edl(o["stp-root-guard"], d, "stp_root_guard")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp-root-guard"], "ObjectSwitchControllerManagedSwitchPorts-StpRootGuard"); ok {
			if err = d.Set("stp_root_guard", vv); err != nil {
				return fmt.Errorf("Error reading stp_root_guard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp_root_guard: %v", err)
		}
	}

	if err = d.Set("stp_state", flattenObjectSwitchControllerManagedSwitchPortsStpState2edl(o["stp-state"], d, "stp_state")); err != nil {
		if vv, ok := fortiAPIPatch(o["stp-state"], "ObjectSwitchControllerManagedSwitchPorts-StpState"); ok {
			if err = d.Set("stp_state", vv); err != nil {
				return fmt.Errorf("Error reading stp_state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stp_state: %v", err)
		}
	}

	if err = d.Set("trunk_member", flattenObjectSwitchControllerManagedSwitchPortsTrunkMember2edl(o["trunk-member"], d, "trunk_member")); err != nil {
		if vv, ok := fortiAPIPatch(o["trunk-member"], "ObjectSwitchControllerManagedSwitchPorts-TrunkMember"); ok {
			if err = d.Set("trunk_member", vv); err != nil {
				return fmt.Errorf("Error reading trunk_member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trunk_member: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSwitchControllerManagedSwitchPortsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSwitchControllerManagedSwitchPorts-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("untagged_vlans", flattenObjectSwitchControllerManagedSwitchPortsUntaggedVlans2edl(o["untagged-vlans"], d, "untagged_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["untagged-vlans"], "ObjectSwitchControllerManagedSwitchPorts-UntaggedVlans"); ok {
			if err = d.Set("untagged_vlans", vv); err != nil {
				return fmt.Errorf("Error reading untagged_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untagged_vlans: %v", err)
		}
	}

	if err = d.Set("vlan", flattenObjectSwitchControllerManagedSwitchPortsVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "ObjectSwitchControllerManagedSwitchPorts-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerManagedSwitchPortsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerManagedSwitchPortsAccessMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAclGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerManagedSwitchPortsAggregatorMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAllowedVlans2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSwitchControllerManagedSwitchPortsAllowedVlansAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsArpInspectionTrust2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsAuthenticatedPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsBundle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["circuit-id"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId2edl(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId2edl(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-name"], _ = expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName2edl(d, i["vlan_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDhcpSnooping2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDiscardMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsDslProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsEdgePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsEncryptedPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFecCapable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFecState2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlapguard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFlowControl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsFortiswitchAcls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpSnooping2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsInterfaceTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSwitchControllerManagedSwitchPortsIpSourceGuard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLacpSpeed2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLearningLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLinkStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLldpProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLldpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLoopGuard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMaxBundle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMclag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMclagIclPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMediaType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerManagedSwitchPortsMinBundle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsP2PPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPacketSampleRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPacketSampler2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPauseMeter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPauseMeterResume2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeMaxPower2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePortMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePortPower2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePortPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeStandard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPoeStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortOwner2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsPtpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsQosPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsRpvstPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSampleDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSflowCounterInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSflowSampleRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsSflowSampler2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStickyMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpBpduGuard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpRootGuard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsStpState2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsTrunkMember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchPortsUntaggedVlans2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectSwitchControllerManagedSwitchPortsVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerManagedSwitchPorts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_mode"); ok || d.HasChange("access_mode") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsAccessMode2edl(d, v, "access_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-mode"] = t
		}
	}

	if v, ok := d.GetOk("acl_group"); ok || d.HasChange("acl_group") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsAclGroup2edl(d, v, "acl_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acl-group"] = t
		}
	}

	if v, ok := d.GetOk("aggregator_mode"); ok || d.HasChange("aggregator_mode") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsAggregatorMode2edl(d, v, "aggregator_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregator-mode"] = t
		}
	}

	if v, ok := d.GetOk("allowed_vlans"); ok || d.HasChange("allowed_vlans") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsAllowedVlans2edl(d, v, "allowed_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans"] = t
		}
	}

	if v, ok := d.GetOk("allowed_vlans_all"); ok || d.HasChange("allowed_vlans_all") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsAllowedVlansAll2edl(d, v, "allowed_vlans_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans-all"] = t
		}
	}

	if v, ok := d.GetOk("arp_inspection_trust"); ok || d.HasChange("arp_inspection_trust") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsArpInspectionTrust2edl(d, v, "arp_inspection_trust")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-inspection-trust"] = t
		}
	}

	if v, ok := d.GetOk("authenticated_port"); ok || d.HasChange("authenticated_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsAuthenticatedPort2edl(d, v, "authenticated_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authenticated-port"] = t
		}
	}

	if v, ok := d.GetOk("bundle"); ok || d.HasChange("bundle") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsBundle2edl(d, v, "bundle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bundle"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_option82_override"); ok || d.HasChange("dhcp_snoop_option82_override") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override2edl(d, v, "dhcp_snoop_option82_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-option82-override"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_option82_trust"); ok || d.HasChange("dhcp_snoop_option82_trust") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust2edl(d, v, "dhcp_snoop_option82_trust")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-option82-trust"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping"); ok || d.HasChange("dhcp_snooping") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsDhcpSnooping2edl(d, v, "dhcp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("discard_mode"); ok || d.HasChange("discard_mode") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsDiscardMode2edl(d, v, "discard_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discard-mode"] = t
		}
	}

	if v, ok := d.GetOk("dsl_profile"); ok || d.HasChange("dsl_profile") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsDslProfile2edl(d, v, "dsl_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsl-profile"] = t
		}
	}

	if v, ok := d.GetOk("edge_port"); ok || d.HasChange("edge_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsEdgePort2edl(d, v, "edge_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["edge-port"] = t
		}
	}

	if v, ok := d.GetOk("encrypted_port"); ok || d.HasChange("encrypted_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsEncryptedPort2edl(d, v, "encrypted_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypted-port"] = t
		}
	}

	if v, ok := d.GetOk("fec_capable"); ok || d.HasChange("fec_capable") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFecCapable2edl(d, v, "fec_capable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-capable"] = t
		}
	}

	if v, ok := d.GetOk("fec_state"); ok || d.HasChange("fec_state") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFecState2edl(d, v, "fec_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-state"] = t
		}
	}

	if v, ok := d.GetOk("flap_duration"); ok || d.HasChange("flap_duration") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFlapDuration2edl(d, v, "flap_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-duration"] = t
		}
	}

	if v, ok := d.GetOk("flap_rate"); ok || d.HasChange("flap_rate") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFlapRate2edl(d, v, "flap_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-rate"] = t
		}
	}

	if v, ok := d.GetOk("flap_timeout"); ok || d.HasChange("flap_timeout") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFlapTimeout2edl(d, v, "flap_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-timeout"] = t
		}
	}

	if v, ok := d.GetOk("flapguard"); ok || d.HasChange("flapguard") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFlapguard2edl(d, v, "flapguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flapguard"] = t
		}
	}

	if v, ok := d.GetOk("flow_control"); ok || d.HasChange("flow_control") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFlowControl2edl(d, v, "flow_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-control"] = t
		}
	}

	if v, ok := d.GetOk("fortiswitch_acls"); ok || d.HasChange("fortiswitch_acls") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsFortiswitchAcls2edl(d, v, "fortiswitch_acls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiswitch-acls"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_flood_reports"); ok || d.HasChange("igmp_snooping_flood_reports") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports2edl(d, v, "igmp_snooping_flood_reports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-flood-reports"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok || d.HasChange("igmp_snooping") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsIgmpSnooping2edl(d, v, "igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("igmps_flood_reports"); ok || d.HasChange("igmps_flood_reports") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodReports2edl(d, v, "igmps_flood_reports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmps-flood-reports"] = t
		}
	}

	if v, ok := d.GetOk("igmps_flood_traffic"); ok || d.HasChange("igmps_flood_traffic") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsIgmpsFloodTraffic2edl(d, v, "igmps_flood_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmps-flood-traffic"] = t
		}
	}

	if v, ok := d.GetOk("interface_tags"); ok || d.HasChange("interface_tags") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsInterfaceTags2edl(d, v, "interface_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-tags"] = t
		}
	}

	if v, ok := d.GetOk("ip_source_guard"); ok || d.HasChange("ip_source_guard") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsIpSourceGuard2edl(d, v, "ip_source_guard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-source-guard"] = t
		}
	}

	if v, ok := d.GetOk("isl_peer_device_sn"); ok || d.HasChange("isl_peer_device_sn") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsIslPeerDeviceSn2edl(d, v, "isl_peer_device_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isl-peer-device-sn"] = t
		}
	}

	if v, ok := d.GetOk("lacp_speed"); ok || d.HasChange("lacp_speed") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLacpSpeed2edl(d, v, "lacp_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-speed"] = t
		}
	}

	if v, ok := d.GetOk("learning_limit"); ok || d.HasChange("learning_limit") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLearningLimit2edl(d, v, "learning_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-limit"] = t
		}
	}

	if v, ok := d.GetOk("link_status"); ok || d.HasChange("link_status") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLinkStatus2edl(d, v, "link_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-status"] = t
		}
	}

	if v, ok := d.GetOk("lldp_profile"); ok || d.HasChange("lldp_profile") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLldpProfile2edl(d, v, "lldp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-profile"] = t
		}
	}

	if v, ok := d.GetOk("lldp_status"); ok || d.HasChange("lldp_status") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLldpStatus2edl(d, v, "lldp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-status"] = t
		}
	}

	if v, ok := d.GetOk("loop_guard"); ok || d.HasChange("loop_guard") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLoopGuard2edl(d, v, "loop_guard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loop-guard"] = t
		}
	}

	if v, ok := d.GetOk("loop_guard_timeout"); ok || d.HasChange("loop_guard_timeout") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsLoopGuardTimeout2edl(d, v, "loop_guard_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loop-guard-timeout"] = t
		}
	}

	if v, ok := d.GetOk("matched_dpp_intf_tags"); ok || d.HasChange("matched_dpp_intf_tags") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMatchedDppIntfTags2edl(d, v, "matched_dpp_intf_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["matched-dpp-intf-tags"] = t
		}
	}

	if v, ok := d.GetOk("matched_dpp_policy"); ok || d.HasChange("matched_dpp_policy") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMatchedDppPolicy2edl(d, v, "matched_dpp_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["matched-dpp-policy"] = t
		}
	}

	if v, ok := d.GetOk("max_bundle"); ok || d.HasChange("max_bundle") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMaxBundle2edl(d, v, "max_bundle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-bundle"] = t
		}
	}

	if v, ok := d.GetOk("mcast_snooping_flood_traffic"); ok || d.HasChange("mcast_snooping_flood_traffic") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic2edl(d, v, "mcast_snooping_flood_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-snooping-flood-traffic"] = t
		}
	}

	if v, ok := d.GetOk("mclag"); ok || d.HasChange("mclag") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMclag2edl(d, v, "mclag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag"] = t
		}
	}

	if v, ok := d.GetOk("mclag_icl_port"); ok || d.HasChange("mclag_icl_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMclagIclPort2edl(d, v, "mclag_icl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag-icl-port"] = t
		}
	}

	if v, ok := d.GetOk("media_type"); ok || d.HasChange("media_type") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMediaType2edl(d, v, "media_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["media-type"] = t
		}
	}

	if v, ok := d.GetOk("member_withdrawal_behavior"); ok || d.HasChange("member_withdrawal_behavior") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior2edl(d, v, "member_withdrawal_behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-withdrawal-behavior"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMembers2edl(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("min_bundle"); ok || d.HasChange("min_bundle") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMinBundle2edl(d, v, "min_bundle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-bundle"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("p2p_port"); ok || d.HasChange("p2p_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsP2PPort2edl(d, v, "p2p_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-port"] = t
		}
	}

	if v, ok := d.GetOk("packet_sample_rate"); ok || d.HasChange("packet_sample_rate") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPacketSampleRate2edl(d, v, "packet_sample_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("packet_sampler"); ok || d.HasChange("packet_sampler") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPacketSampler2edl(d, v, "packet_sampler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-sampler"] = t
		}
	}

	if v, ok := d.GetOk("pause_meter"); ok || d.HasChange("pause_meter") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPauseMeter2edl(d, v, "pause_meter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pause-meter"] = t
		}
	}

	if v, ok := d.GetOk("pause_meter_resume"); ok || d.HasChange("pause_meter_resume") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPauseMeterResume2edl(d, v, "pause_meter_resume")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pause-meter-resume"] = t
		}
	}

	if v, ok := d.GetOk("poe_max_power"); ok || d.HasChange("poe_max_power") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoeMaxPower2edl(d, v, "poe_max_power")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-max-power"] = t
		}
	}

	if v, ok := d.GetOk("poe_mode_bt_cabable"); ok || d.HasChange("poe_mode_bt_cabable") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoeModeBtCabable2edl(d, v, "poe_mode_bt_cabable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-mode-bt-cabable"] = t
		}
	}

	if v, ok := d.GetOk("poe_port_mode"); ok || d.HasChange("poe_port_mode") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoePortMode2edl(d, v, "poe_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("poe_port_power"); ok || d.HasChange("poe_port_power") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoePortPower2edl(d, v, "poe_port_power")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-port-power"] = t
		}
	}

	if v, ok := d.GetOk("poe_port_priority"); ok || d.HasChange("poe_port_priority") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoePortPriority2edl(d, v, "poe_port_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-port-priority"] = t
		}
	}

	if v, ok := d.GetOk("poe_pre_standard_detection"); ok || d.HasChange("poe_pre_standard_detection") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoePreStandardDetection2edl(d, v, "poe_pre_standard_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-pre-standard-detection"] = t
		}
	}

	if v, ok := d.GetOk("poe_standard"); ok || d.HasChange("poe_standard") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoeStandard2edl(d, v, "poe_standard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-standard"] = t
		}
	}

	if v, ok := d.GetOk("poe_status"); ok || d.HasChange("poe_status") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPoeStatus2edl(d, v, "poe_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-status"] = t
		}
	}

	if v, ok := d.GetOk("port_name"); ok || d.HasChange("port_name") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPortName2edl(d, v, "port_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-name"] = t
		}
	}

	if v, ok := d.GetOk("port_owner"); ok || d.HasChange("port_owner") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPortOwner2edl(d, v, "port_owner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-owner"] = t
		}
	}

	if v, ok := d.GetOk("port_policy"); ok || d.HasChange("port_policy") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPortPolicy2edl(d, v, "port_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-policy"] = t
		}
	}

	if v, ok := d.GetOk("port_security_policy"); ok || d.HasChange("port_security_policy") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPortSecurityPolicy2edl(d, v, "port_security_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-security-policy"] = t
		}
	}

	if v, ok := d.GetOk("port_selection_criteria"); ok || d.HasChange("port_selection_criteria") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPortSelectionCriteria2edl(d, v, "port_selection_criteria")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-selection-criteria"] = t
		}
	}

	if v, ok := d.GetOk("ptp_status"); ok || d.HasChange("ptp_status") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsPtpStatus2edl(d, v, "ptp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-status"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok || d.HasChange("qos_policy") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsQosPolicy2edl(d, v, "qos_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("restricted_auth_port"); ok || d.HasChange("restricted_auth_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsRestrictedAuthPort2edl(d, v, "restricted_auth_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted-auth-port"] = t
		}
	}

	if v, ok := d.GetOk("rpvst_port"); ok || d.HasChange("rpvst_port") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsRpvstPort2edl(d, v, "rpvst_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpvst-port"] = t
		}
	}

	if v, ok := d.GetOk("sample_direction"); ok || d.HasChange("sample_direction") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsSampleDirection2edl(d, v, "sample_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-direction"] = t
		}
	}

	if v, ok := d.GetOk("sflow_counter_interval"); ok || d.HasChange("sflow_counter_interval") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsSflowCounterInterval2edl(d, v, "sflow_counter_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflow-counter-interval"] = t
		}
	}

	if v, ok := d.GetOk("sflow_sample_rate"); ok || d.HasChange("sflow_sample_rate") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsSflowSampleRate2edl(d, v, "sflow_sample_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflow-sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("sflow_sampler"); ok || d.HasChange("sflow_sampler") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsSflowSampler2edl(d, v, "sflow_sampler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflow-sampler"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sticky_mac"); ok || d.HasChange("sticky_mac") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsStickyMac2edl(d, v, "sticky_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-mac"] = t
		}
	}

	if v, ok := d.GetOk("stp_bpdu_guard"); ok || d.HasChange("stp_bpdu_guard") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsStpBpduGuard2edl(d, v, "stp_bpdu_guard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-bpdu-guard"] = t
		}
	}

	if v, ok := d.GetOk("stp_bpdu_guard_timeout"); ok || d.HasChange("stp_bpdu_guard_timeout") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsStpBpduGuardTimeout2edl(d, v, "stp_bpdu_guard_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-bpdu-guard-timeout"] = t
		}
	}

	if v, ok := d.GetOk("stp_root_guard"); ok || d.HasChange("stp_root_guard") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsStpRootGuard2edl(d, v, "stp_root_guard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-root-guard"] = t
		}
	}

	if v, ok := d.GetOk("stp_state"); ok || d.HasChange("stp_state") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsStpState2edl(d, v, "stp_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-state"] = t
		}
	}

	if v, ok := d.GetOk("trunk_member"); ok || d.HasChange("trunk_member") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsTrunkMember2edl(d, v, "trunk_member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trunk-member"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("untagged_vlans"); ok || d.HasChange("untagged_vlans") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsUntaggedVlans2edl(d, v, "untagged_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandObjectSwitchControllerManagedSwitchPortsVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
