// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DHCP servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanDynamicMappingDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanDynamicMappingDhcpServerUpdate,
		Read:   resourceObjectFspVlanDynamicMappingDhcpServerRead,
		Update: resourceObjectFspVlanDynamicMappingDhcpServerUpdate,
		Delete: resourceObjectFspVlanDynamicMappingDhcpServerDelete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_managed_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conflicted_ip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_keyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ddns_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_update_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_settings_from_fortiipam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticlient_on_net_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lease_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ipsec_lease_hold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac_acl_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"option1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"option2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"option3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"option4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"option5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"option6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"relay_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"circuit_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"circuit_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
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
						"remote_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"remote_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shared_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tftp_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"wifi_ac_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
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

func resourceObjectFspVlanDynamicMappingDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectObjectFspVlanDynamicMappingDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMappingDhcpServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanDynamicMappingDhcpServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMappingDhcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFspVlanDynamicMappingDhcpServer")

	return resourceObjectFspVlanDynamicMappingDhcpServerRead(d, m)
}

func resourceObjectFspVlanDynamicMappingDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	wsParams["adom"] = adomv

	err = c.DeleteObjectFspVlanDynamicMappingDhcpServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanDynamicMappingDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanDynamicMappingDhcpServerRead(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	if dynamic_mapping_name == "" {
		dynamic_mapping_name = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_name")
		if dynamic_mapping_name == "" {
			return fmt.Errorf("Parameter dynamic_mapping_name is missing")
		}
		if err = d.Set("dynamic_mapping_name", dynamic_mapping_name); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_name: %v", err)
		}
	}
	if dynamic_mapping_vdom == "" {
		dynamic_mapping_vdom = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_vdom")
		if dynamic_mapping_vdom == "" {
			return fmt.Errorf("Parameter dynamic_mapping_vdom is missing")
		}
		if err = d.Set("dynamic_mapping_vdom", dynamic_mapping_vdom); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_vdom: %v", err)
		}
	}
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	o, err := c.ReadObjectFspVlanDynamicMappingDhcpServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMappingDhcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanDynamicMappingDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMappingDhcpServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpServerAutoConfiguration3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsAuth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsKey3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsKeyname3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsServerIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsTtl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdate3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDdnsZone3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDefaultGateway3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer33rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsServer43rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDnsService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerDomain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerEnable3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp3rdl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime3rdl(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp3rdl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch3rdl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString3rdl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch3rdl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString3rdl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerFilename3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRange3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp3rdl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime3rdl(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp3rdl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch3rdl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciString3rdl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch3rdl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciString3rdl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-IpRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeUciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpRangeVciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerLeaseTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNetmask3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNextServer3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpServer33rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerNtpService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption33rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption43rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption53rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOption63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptions3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsCode3rdl(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsIp3rdl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsType3rdl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch3rdl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciString3rdl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsValue3rdl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch3rdl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciString3rdl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-Options-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsCode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsUciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerOptionsVciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerRelayAgent3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction3rdl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId3rdl(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType3rdl(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription3rdl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp3rdl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac3rdl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId3rdl(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType3rdl(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType3rdl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerServerType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerSharedSubnet3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerTftpServer3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerTimezone3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerTimezoneOption3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerVciMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerVciString3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAcService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWifiAc33rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWinsServer13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerWinsServer23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanDynamicMappingDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auto_configuration", flattenObjectFspVlanDynamicMappingDhcpServerAutoConfiguration3rdl(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-configuration"], "ObjectFspVlanDynamicMappingDhcpServer-AutoConfiguration"); ok {
			if err = d.Set("auto_configuration", vv); err != nil {
				return fmt.Errorf("Error reading auto_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("auto_managed_status", flattenObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus3rdl(o["auto-managed-status"], d, "auto_managed_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-managed-status"], "ObjectFspVlanDynamicMappingDhcpServer-AutoManagedStatus"); ok {
			if err = d.Set("auto_managed_status", vv); err != nil {
				return fmt.Errorf("Error reading auto_managed_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_managed_status: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout3rdl(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conflicted-ip-timeout"], "ObjectFspVlanDynamicMappingDhcpServer-ConflictedIpTimeout"); ok {
			if err = d.Set("conflicted_ip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenObjectFspVlanDynamicMappingDhcpServerDdnsAuth3rdl(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_key", flattenObjectFspVlanDynamicMappingDhcpServerDdnsKey3rdl(o["ddns-key"], d, "ddns_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-key"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsKey"); ok {
			if err = d.Set("ddns_key", vv); err != nil {
				return fmt.Errorf("Error reading ddns_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_key: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenObjectFspVlanDynamicMappingDhcpServerDdnsKeyname3rdl(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenObjectFspVlanDynamicMappingDhcpServerDdnsServerIp3rdl(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenObjectFspVlanDynamicMappingDhcpServerDdnsTtl3rdl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdate3rdl(o["ddns-update"], d, "ddns_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsUpdate"); ok {
			if err = d.Set("ddns_update", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride3rdl(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update-override"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsUpdateOverride"); ok {
			if err = d.Set("ddns_update_override", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenObjectFspVlanDynamicMappingDhcpServerDdnsZone3rdl(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "ObjectFspVlanDynamicMappingDhcpServer-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenObjectFspVlanDynamicMappingDhcpServerDefaultGateway3rdl(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "ObjectFspVlanDynamicMappingDhcpServer-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_settings_from_fortiipam", flattenObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam3rdl(o["dhcp-settings-from-fortiipam"], d, "dhcp_settings_from_fortiipam")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-settings-from-fortiipam"], "ObjectFspVlanDynamicMappingDhcpServer-DhcpSettingsFromFortiipam"); ok {
			if err = d.Set("dhcp_settings_from_fortiipam", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenObjectFspVlanDynamicMappingDhcpServerDnsServer13rdl(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "ObjectFspVlanDynamicMappingDhcpServer-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenObjectFspVlanDynamicMappingDhcpServerDnsServer23rdl(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "ObjectFspVlanDynamicMappingDhcpServer-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenObjectFspVlanDynamicMappingDhcpServerDnsServer33rdl(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "ObjectFspVlanDynamicMappingDhcpServer-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenObjectFspVlanDynamicMappingDhcpServerDnsServer43rdl(o["dns-server4"], d, "dns_server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server4"], "ObjectFspVlanDynamicMappingDhcpServer-DnsServer4"); ok {
			if err = d.Set("dns_server4", vv); err != nil {
				return fmt.Errorf("Error reading dns_server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenObjectFspVlanDynamicMappingDhcpServerDnsService3rdl(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "ObjectFspVlanDynamicMappingDhcpServer-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectFspVlanDynamicMappingDhcpServerDomain3rdl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectFspVlanDynamicMappingDhcpServer-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("enable", flattenObjectFspVlanDynamicMappingDhcpServerEnable3rdl(o["enable"], d, "enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["enable"], "ObjectFspVlanDynamicMappingDhcpServer-Enable"); ok {
			if err = d.Set("enable", vv); err != nil {
				return fmt.Errorf("Error reading enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enable: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exclude_range", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange3rdl(o["exclude-range"], d, "exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["exclude-range"], "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange"); ok {
				if err = d.Set("exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exclude_range"); ok {
			if err = d.Set("exclude_range", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRange3rdl(o["exclude-range"], d, "exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["exclude-range"], "ObjectFspVlanDynamicMappingDhcpServer-ExcludeRange"); ok {
					if err = d.Set("exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("filename", flattenObjectFspVlanDynamicMappingDhcpServerFilename3rdl(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "ObjectFspVlanDynamicMappingDhcpServer-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus3rdl(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-on-net-status"], "ObjectFspVlanDynamicMappingDhcpServer-ForticlientOnNetStatus"); ok {
			if err = d.Set("forticlient_on_net_status", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFspVlanDynamicMappingDhcpServerId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanDynamicMappingDhcpServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenObjectFspVlanDynamicMappingDhcpServerIpMode3rdl(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "ObjectFspVlanDynamicMappingDhcpServer-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenObjectFspVlanDynamicMappingDhcpServerIpRange3rdl(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectFspVlanDynamicMappingDhcpServer-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenObjectFspVlanDynamicMappingDhcpServerIpRange3rdl(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectFspVlanDynamicMappingDhcpServer-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold3rdl(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-lease-hold"], "ObjectFspVlanDynamicMappingDhcpServer-IpsecLeaseHold"); ok {
			if err = d.Set("ipsec_lease_hold", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenObjectFspVlanDynamicMappingDhcpServerLeaseTime3rdl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "ObjectFspVlanDynamicMappingDhcpServer-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction3rdl(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-acl-default-action"], "ObjectFspVlanDynamicMappingDhcpServer-MacAclDefaultAction"); ok {
			if err = d.Set("mac_acl_default_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("netmask", flattenObjectFspVlanDynamicMappingDhcpServerNetmask3rdl(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "ObjectFspVlanDynamicMappingDhcpServer-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("next_server", flattenObjectFspVlanDynamicMappingDhcpServerNextServer3rdl(o["next-server"], d, "next_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-server"], "ObjectFspVlanDynamicMappingDhcpServer-NextServer"); ok {
			if err = d.Set("next_server", vv); err != nil {
				return fmt.Errorf("Error reading next_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenObjectFspVlanDynamicMappingDhcpServerNtpServer13rdl(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server1"], "ObjectFspVlanDynamicMappingDhcpServer-NtpServer1"); ok {
			if err = d.Set("ntp_server1", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenObjectFspVlanDynamicMappingDhcpServerNtpServer23rdl(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server2"], "ObjectFspVlanDynamicMappingDhcpServer-NtpServer2"); ok {
			if err = d.Set("ntp_server2", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenObjectFspVlanDynamicMappingDhcpServerNtpServer33rdl(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server3"], "ObjectFspVlanDynamicMappingDhcpServer-NtpServer3"); ok {
			if err = d.Set("ntp_server3", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenObjectFspVlanDynamicMappingDhcpServerNtpService3rdl(o["ntp-service"], d, "ntp_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-service"], "ObjectFspVlanDynamicMappingDhcpServer-NtpService"); ok {
			if err = d.Set("ntp_service", vv); err != nil {
				return fmt.Errorf("Error reading ntp_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if err = d.Set("option1", flattenObjectFspVlanDynamicMappingDhcpServerOption13rdl(o["option1"], d, "option1")); err != nil {
		if vv, ok := fortiAPIPatch(o["option1"], "ObjectFspVlanDynamicMappingDhcpServer-Option1"); ok {
			if err = d.Set("option1", vv); err != nil {
				return fmt.Errorf("Error reading option1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option1: %v", err)
		}
	}

	if err = d.Set("option2", flattenObjectFspVlanDynamicMappingDhcpServerOption23rdl(o["option2"], d, "option2")); err != nil {
		if vv, ok := fortiAPIPatch(o["option2"], "ObjectFspVlanDynamicMappingDhcpServer-Option2"); ok {
			if err = d.Set("option2", vv); err != nil {
				return fmt.Errorf("Error reading option2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option2: %v", err)
		}
	}

	if err = d.Set("option3", flattenObjectFspVlanDynamicMappingDhcpServerOption33rdl(o["option3"], d, "option3")); err != nil {
		if vv, ok := fortiAPIPatch(o["option3"], "ObjectFspVlanDynamicMappingDhcpServer-Option3"); ok {
			if err = d.Set("option3", vv); err != nil {
				return fmt.Errorf("Error reading option3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option3: %v", err)
		}
	}

	if err = d.Set("option4", flattenObjectFspVlanDynamicMappingDhcpServerOption43rdl(o["option4"], d, "option4")); err != nil {
		if vv, ok := fortiAPIPatch(o["option4"], "ObjectFspVlanDynamicMappingDhcpServer-Option4"); ok {
			if err = d.Set("option4", vv); err != nil {
				return fmt.Errorf("Error reading option4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option4: %v", err)
		}
	}

	if err = d.Set("option5", flattenObjectFspVlanDynamicMappingDhcpServerOption53rdl(o["option5"], d, "option5")); err != nil {
		if vv, ok := fortiAPIPatch(o["option5"], "ObjectFspVlanDynamicMappingDhcpServer-Option5"); ok {
			if err = d.Set("option5", vv); err != nil {
				return fmt.Errorf("Error reading option5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option5: %v", err)
		}
	}

	if err = d.Set("option6", flattenObjectFspVlanDynamicMappingDhcpServerOption63rdl(o["option6"], d, "option6")); err != nil {
		if vv, ok := fortiAPIPatch(o["option6"], "ObjectFspVlanDynamicMappingDhcpServer-Option6"); ok {
			if err = d.Set("option6", vv); err != nil {
				return fmt.Errorf("Error reading option6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenObjectFspVlanDynamicMappingDhcpServerOptions3rdl(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "ObjectFspVlanDynamicMappingDhcpServer-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenObjectFspVlanDynamicMappingDhcpServerOptions3rdl(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "ObjectFspVlanDynamicMappingDhcpServer-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("relay_agent", flattenObjectFspVlanDynamicMappingDhcpServerRelayAgent3rdl(o["relay-agent"], d, "relay_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-agent"], "ObjectFspVlanDynamicMappingDhcpServer-RelayAgent"); ok {
			if err = d.Set("relay_agent", vv); err != nil {
				return fmt.Errorf("Error reading relay_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_agent: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("reserved_address", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress3rdl(o["reserved-address"], d, "reserved_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["reserved-address"], "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress"); ok {
				if err = d.Set("reserved_address", vv); err != nil {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading reserved_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reserved_address"); ok {
			if err = d.Set("reserved_address", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddress3rdl(o["reserved-address"], d, "reserved_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["reserved-address"], "ObjectFspVlanDynamicMappingDhcpServer-ReservedAddress"); ok {
					if err = d.Set("reserved_address", vv); err != nil {
						return fmt.Errorf("Error reading reserved_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenObjectFspVlanDynamicMappingDhcpServerServerType3rdl(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "ObjectFspVlanDynamicMappingDhcpServer-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("shared_subnet", flattenObjectFspVlanDynamicMappingDhcpServerSharedSubnet3rdl(o["shared-subnet"], d, "shared_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["shared-subnet"], "ObjectFspVlanDynamicMappingDhcpServer-SharedSubnet"); ok {
			if err = d.Set("shared_subnet", vv); err != nil {
				return fmt.Errorf("Error reading shared_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shared_subnet: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFspVlanDynamicMappingDhcpServerStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFspVlanDynamicMappingDhcpServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tftp_server", flattenObjectFspVlanDynamicMappingDhcpServerTftpServer3rdl(o["tftp-server"], d, "tftp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp-server"], "ObjectFspVlanDynamicMappingDhcpServer-TftpServer"); ok {
			if err = d.Set("tftp_server", vv); err != nil {
				return fmt.Errorf("Error reading tftp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp_server: %v", err)
		}
	}

	if err = d.Set("timezone", flattenObjectFspVlanDynamicMappingDhcpServerTimezone3rdl(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "ObjectFspVlanDynamicMappingDhcpServer-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("timezone_option", flattenObjectFspVlanDynamicMappingDhcpServerTimezoneOption3rdl(o["timezone-option"], d, "timezone_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone-option"], "ObjectFspVlanDynamicMappingDhcpServer-TimezoneOption"); ok {
			if err = d.Set("timezone_option", vv); err != nil {
				return fmt.Errorf("Error reading timezone_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenObjectFspVlanDynamicMappingDhcpServerVciMatch3rdl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "ObjectFspVlanDynamicMappingDhcpServer-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenObjectFspVlanDynamicMappingDhcpServerVciString3rdl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "ObjectFspVlanDynamicMappingDhcpServer-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", flattenObjectFspVlanDynamicMappingDhcpServerWifiAcService3rdl(o["wifi-ac-service"], d, "wifi_ac_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac-service"], "ObjectFspVlanDynamicMappingDhcpServer-WifiAcService"); ok {
			if err = d.Set("wifi_ac_service", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenObjectFspVlanDynamicMappingDhcpServerWifiAc13rdl(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac1"], "ObjectFspVlanDynamicMappingDhcpServer-WifiAc1"); ok {
			if err = d.Set("wifi_ac1", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenObjectFspVlanDynamicMappingDhcpServerWifiAc23rdl(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac2"], "ObjectFspVlanDynamicMappingDhcpServer-WifiAc2"); ok {
			if err = d.Set("wifi_ac2", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenObjectFspVlanDynamicMappingDhcpServerWifiAc33rdl(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac3"], "ObjectFspVlanDynamicMappingDhcpServer-WifiAc3"); ok {
			if err = d.Set("wifi_ac3", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenObjectFspVlanDynamicMappingDhcpServerWinsServer13rdl(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "ObjectFspVlanDynamicMappingDhcpServer-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenObjectFspVlanDynamicMappingDhcpServerWinsServer23rdl(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "ObjectFspVlanDynamicMappingDhcpServer-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDynamicMappingDhcpServerAutoConfiguration3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsAuth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsKeyname3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsServerIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsTtl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdate3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDdnsZone3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDefaultGateway3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer33rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsServer43rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDnsService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerDomain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerEnable3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp3rdl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime3rdl(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp3rdl(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch3rdl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString3rdl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch3rdl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString3rdl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerFilename3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp3rdl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime3rdl(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp3rdl(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch3rdl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciString3rdl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch3rdl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciString3rdl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeEndIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeLeaseTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeStartIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeUciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpRangeVciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerLeaseTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNetmask3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNextServer3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpServer33rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerNtpService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption33rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption43rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption53rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOption63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptions3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsCode3rdl(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsIp3rdl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsType3rdl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch3rdl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsUciString3rdl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsValue3rdl(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch3rdl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDynamicMappingDhcpServerOptionsVciString3rdl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsCode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsUciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsUciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsVciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerOptionsVciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerRelayAgent3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction3rdl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId3rdl(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType3rdl(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription3rdl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp3rdl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac3rdl(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId3rdl(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType3rdl(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType3rdl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerServerType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerSharedSubnet3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTftpServer3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTimezone3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerTimezoneOption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerVciMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerVciString3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAcService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWifiAc33rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWinsServer13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerWinsServer23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanDynamicMappingDhcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_configuration"); ok || d.HasChange("auto_configuration") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerAutoConfiguration3rdl(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("auto_managed_status"); ok || d.HasChange("auto_managed_status") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerAutoManagedStatus3rdl(d, v, "auto_managed_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-managed-status"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok || d.HasChange("conflicted_ip_timeout") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerConflictedIpTimeout3rdl(d, v, "conflicted_ip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsAuth3rdl(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsKey3rdl(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsKeyname3rdl(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsServerIp3rdl(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsTtl3rdl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok || d.HasChange("ddns_update") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdate3rdl(d, v, "ddns_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok || d.HasChange("ddns_update_override") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsUpdateOverride3rdl(d, v, "ddns_update_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDdnsZone3rdl(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDefaultGateway3rdl(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_settings_from_fortiipam"); ok || d.HasChange("dhcp_settings_from_fortiipam") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDhcpSettingsFromFortiipam3rdl(d, v, "dhcp_settings_from_fortiipam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-settings-from-fortiipam"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDnsServer13rdl(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDnsServer23rdl(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDnsServer33rdl(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok || d.HasChange("dns_server4") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDnsServer43rdl(d, v, "dns_server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDnsService3rdl(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerDomain3rdl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("enable"); ok || d.HasChange("enable") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerEnable3rdl(d, v, "enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enable"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok || d.HasChange("exclude_range") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRange3rdl(d, v, "exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerFilename3rdl(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok || d.HasChange("forticlient_on_net_status") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerForticlientOnNetStatus3rdl(d, v, "forticlient_on_net_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerIpMode3rdl(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerIpRange3rdl(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok || d.HasChange("ipsec_lease_hold") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerIpsecLeaseHold3rdl(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerLeaseTime3rdl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok || d.HasChange("mac_acl_default_action") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerMacAclDefaultAction3rdl(d, v, "mac_acl_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerNetmask3rdl(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok || d.HasChange("next_server") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerNextServer3rdl(d, v, "next_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok || d.HasChange("ntp_server1") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerNtpServer13rdl(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok || d.HasChange("ntp_server2") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerNtpServer23rdl(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok || d.HasChange("ntp_server3") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerNtpServer33rdl(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok || d.HasChange("ntp_service") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerNtpService3rdl(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("option1"); ok || d.HasChange("option1") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOption13rdl(d, v, "option1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option1"] = t
		}
	}

	if v, ok := d.GetOk("option2"); ok || d.HasChange("option2") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOption23rdl(d, v, "option2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option2"] = t
		}
	}

	if v, ok := d.GetOk("option3"); ok || d.HasChange("option3") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOption33rdl(d, v, "option3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option3"] = t
		}
	}

	if v, ok := d.GetOk("option4"); ok || d.HasChange("option4") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOption43rdl(d, v, "option4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option4"] = t
		}
	}

	if v, ok := d.GetOk("option5"); ok || d.HasChange("option5") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOption53rdl(d, v, "option5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option5"] = t
		}
	}

	if v, ok := d.GetOk("option6"); ok || d.HasChange("option6") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOption63rdl(d, v, "option6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option6"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerOptions3rdl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("relay_agent"); ok || d.HasChange("relay_agent") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerRelayAgent3rdl(d, v, "relay_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-agent"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok || d.HasChange("reserved_address") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddress3rdl(d, v, "reserved_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerServerType3rdl(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("shared_subnet"); ok || d.HasChange("shared_subnet") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerSharedSubnet3rdl(d, v, "shared_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-subnet"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tftp_server"); ok || d.HasChange("tftp_server") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerTftpServer3rdl(d, v, "tftp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerTimezone3rdl(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok || d.HasChange("timezone_option") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerTimezoneOption3rdl(d, v, "timezone_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerVciMatch3rdl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerVciString3rdl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac_service"); ok || d.HasChange("wifi_ac_service") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerWifiAcService3rdl(d, v, "wifi_ac_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac-service"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok || d.HasChange("wifi_ac1") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerWifiAc13rdl(d, v, "wifi_ac1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok || d.HasChange("wifi_ac2") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerWifiAc23rdl(d, v, "wifi_ac2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok || d.HasChange("wifi_ac3") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerWifiAc33rdl(d, v, "wifi_ac3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerWinsServer13rdl(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerWinsServer23rdl(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
