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

func resourceObjectFspVlanDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanDhcpServerUpdate,
		Read:   resourceObjectFspVlanDhcpServerRead,
		Update: resourceObjectFspVlanDhcpServerUpdate,
		Delete: resourceObjectFspVlanDhcpServerDelete,

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

func resourceObjectFspVlanDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDhcpServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlanDhcpServer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDhcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFspVlanDhcpServer")

	return resourceObjectFspVlanDhcpServerRead(d, m)
}

func resourceObjectFspVlanDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["vlan"] = vlan

	err = c.DeleteObjectFspVlanDhcpServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanDhcpServerRead(d *schema.ResourceData, m interface{}) error {
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
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanDhcpServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDhcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDhcpServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDhcpServerAutoConfiguration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerAutoManagedStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerConflictedIpTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsAuth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFspVlanDhcpServerDdnsKeyname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsServerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsUpdate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsUpdateOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDdnsZone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDefaultGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDhcpSettingsFromFortiipam2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsServer42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDnsService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerEnable2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerExcludeRangeEndIp2edl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeLeaseTime2edl(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeStartIp2edl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeUciMatch2edl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeUciString2edl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeVciMatch2edl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerExcludeRangeVciString2edl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ExcludeRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerExcludeRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerExcludeRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerExcludeRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerFilename2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerForticlientOnNetStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRange2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerIpRangeEndIp2edl(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := i["lease-time"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeLeaseTime2edl(i["lease-time"], d, pre_append)
			tmp["lease_time"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-LeaseTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeStartIp2edl(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeUciMatch2edl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeUciString2edl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeVciMatch2edl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerIpRangeVciString2edl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-IpRange-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerIpRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerIpRangeVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerIpRangeVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerIpsecLeaseHold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerLeaseTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerMacAclDefaultAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNetmask2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNextServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpServer32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerNtpService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOption12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOption22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOption32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOption42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOption52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOption62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptions2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerOptionsCode2edl(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := i["uci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsUciMatch2edl(i["uci-match"], d, pre_append)
			tmp["uci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-UciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := i["uci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsUciString2edl(i["uci-string"], d, pre_append)
			tmp["uci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-UciString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-Value")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := i["vci-match"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsVciMatch2edl(i["vci-match"], d, pre_append)
			tmp["vci_match"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-VciMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			v := flattenObjectFspVlanDhcpServerOptionsVciString2edl(i["vci-string"], d, pre_append)
			tmp["vci_string"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-Options-VciString")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerOptionsCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOptionsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsUciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsUciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerOptionsValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerOptionsVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerRelayAgent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddress2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFspVlanDhcpServerReservedAddressAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressCircuitId2edl(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressCircuitIdType2edl(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressDescription2edl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressIp2edl(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressMac2edl(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressRemoteId2edl(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressRemoteIdType2edl(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFspVlanDhcpServerReservedAddressType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFspVlanDhcpServer-ReservedAddress-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFspVlanDhcpServerReservedAddressAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressCircuitIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressRemoteId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressRemoteIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerServerType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerSharedSubnet2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerTftpServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerTimezone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerTimezoneOption2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerVciMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerVciString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDhcpServerWifiAcService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWifiAc12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWifiAc22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWifiAc32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWinsServer12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerWinsServer22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auto_configuration", flattenObjectFspVlanDhcpServerAutoConfiguration2edl(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-configuration"], "ObjectFspVlanDhcpServer-AutoConfiguration"); ok {
			if err = d.Set("auto_configuration", vv); err != nil {
				return fmt.Errorf("Error reading auto_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("auto_managed_status", flattenObjectFspVlanDhcpServerAutoManagedStatus2edl(o["auto-managed-status"], d, "auto_managed_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-managed-status"], "ObjectFspVlanDhcpServer-AutoManagedStatus"); ok {
			if err = d.Set("auto_managed_status", vv); err != nil {
				return fmt.Errorf("Error reading auto_managed_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_managed_status: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenObjectFspVlanDhcpServerConflictedIpTimeout2edl(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conflicted-ip-timeout"], "ObjectFspVlanDhcpServer-ConflictedIpTimeout"); ok {
			if err = d.Set("conflicted_ip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenObjectFspVlanDhcpServerDdnsAuth2edl(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "ObjectFspVlanDhcpServer-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_key", flattenObjectFspVlanDhcpServerDdnsKey2edl(o["ddns-key"], d, "ddns_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-key"], "ObjectFspVlanDhcpServer-DdnsKey"); ok {
			if err = d.Set("ddns_key", vv); err != nil {
				return fmt.Errorf("Error reading ddns_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_key: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenObjectFspVlanDhcpServerDdnsKeyname2edl(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "ObjectFspVlanDhcpServer-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenObjectFspVlanDhcpServerDdnsServerIp2edl(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "ObjectFspVlanDhcpServer-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenObjectFspVlanDhcpServerDdnsTtl2edl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "ObjectFspVlanDhcpServer-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenObjectFspVlanDhcpServerDdnsUpdate2edl(o["ddns-update"], d, "ddns_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update"], "ObjectFspVlanDhcpServer-DdnsUpdate"); ok {
			if err = d.Set("ddns_update", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenObjectFspVlanDhcpServerDdnsUpdateOverride2edl(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update-override"], "ObjectFspVlanDhcpServer-DdnsUpdateOverride"); ok {
			if err = d.Set("ddns_update_override", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenObjectFspVlanDhcpServerDdnsZone2edl(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "ObjectFspVlanDhcpServer-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenObjectFspVlanDhcpServerDefaultGateway2edl(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "ObjectFspVlanDhcpServer-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_settings_from_fortiipam", flattenObjectFspVlanDhcpServerDhcpSettingsFromFortiipam2edl(o["dhcp-settings-from-fortiipam"], d, "dhcp_settings_from_fortiipam")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-settings-from-fortiipam"], "ObjectFspVlanDhcpServer-DhcpSettingsFromFortiipam"); ok {
			if err = d.Set("dhcp_settings_from_fortiipam", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenObjectFspVlanDhcpServerDnsServer12edl(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "ObjectFspVlanDhcpServer-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenObjectFspVlanDhcpServerDnsServer22edl(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "ObjectFspVlanDhcpServer-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenObjectFspVlanDhcpServerDnsServer32edl(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "ObjectFspVlanDhcpServer-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenObjectFspVlanDhcpServerDnsServer42edl(o["dns-server4"], d, "dns_server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server4"], "ObjectFspVlanDhcpServer-DnsServer4"); ok {
			if err = d.Set("dns_server4", vv); err != nil {
				return fmt.Errorf("Error reading dns_server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenObjectFspVlanDhcpServerDnsService2edl(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "ObjectFspVlanDhcpServer-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectFspVlanDhcpServerDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectFspVlanDhcpServer-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("enable", flattenObjectFspVlanDhcpServerEnable2edl(o["enable"], d, "enable")); err != nil {
		if vv, ok := fortiAPIPatch(o["enable"], "ObjectFspVlanDhcpServer-Enable"); ok {
			if err = d.Set("enable", vv); err != nil {
				return fmt.Errorf("Error reading enable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enable: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exclude_range", flattenObjectFspVlanDhcpServerExcludeRange2edl(o["exclude-range"], d, "exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["exclude-range"], "ObjectFspVlanDhcpServer-ExcludeRange"); ok {
				if err = d.Set("exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exclude_range"); ok {
			if err = d.Set("exclude_range", flattenObjectFspVlanDhcpServerExcludeRange2edl(o["exclude-range"], d, "exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["exclude-range"], "ObjectFspVlanDhcpServer-ExcludeRange"); ok {
					if err = d.Set("exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("filename", flattenObjectFspVlanDhcpServerFilename2edl(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "ObjectFspVlanDhcpServer-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenObjectFspVlanDhcpServerForticlientOnNetStatus2edl(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-on-net-status"], "ObjectFspVlanDhcpServer-ForticlientOnNetStatus"); ok {
			if err = d.Set("forticlient_on_net_status", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFspVlanDhcpServerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanDhcpServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenObjectFspVlanDhcpServerIpMode2edl(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "ObjectFspVlanDhcpServer-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenObjectFspVlanDhcpServerIpRange2edl(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectFspVlanDhcpServer-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenObjectFspVlanDhcpServerIpRange2edl(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectFspVlanDhcpServer-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenObjectFspVlanDhcpServerIpsecLeaseHold2edl(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-lease-hold"], "ObjectFspVlanDhcpServer-IpsecLeaseHold"); ok {
			if err = d.Set("ipsec_lease_hold", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenObjectFspVlanDhcpServerLeaseTime2edl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "ObjectFspVlanDhcpServer-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenObjectFspVlanDhcpServerMacAclDefaultAction2edl(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-acl-default-action"], "ObjectFspVlanDhcpServer-MacAclDefaultAction"); ok {
			if err = d.Set("mac_acl_default_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("netmask", flattenObjectFspVlanDhcpServerNetmask2edl(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "ObjectFspVlanDhcpServer-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("next_server", flattenObjectFspVlanDhcpServerNextServer2edl(o["next-server"], d, "next_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-server"], "ObjectFspVlanDhcpServer-NextServer"); ok {
			if err = d.Set("next_server", vv); err != nil {
				return fmt.Errorf("Error reading next_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenObjectFspVlanDhcpServerNtpServer12edl(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server1"], "ObjectFspVlanDhcpServer-NtpServer1"); ok {
			if err = d.Set("ntp_server1", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenObjectFspVlanDhcpServerNtpServer22edl(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server2"], "ObjectFspVlanDhcpServer-NtpServer2"); ok {
			if err = d.Set("ntp_server2", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenObjectFspVlanDhcpServerNtpServer32edl(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server3"], "ObjectFspVlanDhcpServer-NtpServer3"); ok {
			if err = d.Set("ntp_server3", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenObjectFspVlanDhcpServerNtpService2edl(o["ntp-service"], d, "ntp_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-service"], "ObjectFspVlanDhcpServer-NtpService"); ok {
			if err = d.Set("ntp_service", vv); err != nil {
				return fmt.Errorf("Error reading ntp_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if err = d.Set("option1", flattenObjectFspVlanDhcpServerOption12edl(o["option1"], d, "option1")); err != nil {
		if vv, ok := fortiAPIPatch(o["option1"], "ObjectFspVlanDhcpServer-Option1"); ok {
			if err = d.Set("option1", vv); err != nil {
				return fmt.Errorf("Error reading option1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option1: %v", err)
		}
	}

	if err = d.Set("option2", flattenObjectFspVlanDhcpServerOption22edl(o["option2"], d, "option2")); err != nil {
		if vv, ok := fortiAPIPatch(o["option2"], "ObjectFspVlanDhcpServer-Option2"); ok {
			if err = d.Set("option2", vv); err != nil {
				return fmt.Errorf("Error reading option2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option2: %v", err)
		}
	}

	if err = d.Set("option3", flattenObjectFspVlanDhcpServerOption32edl(o["option3"], d, "option3")); err != nil {
		if vv, ok := fortiAPIPatch(o["option3"], "ObjectFspVlanDhcpServer-Option3"); ok {
			if err = d.Set("option3", vv); err != nil {
				return fmt.Errorf("Error reading option3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option3: %v", err)
		}
	}

	if err = d.Set("option4", flattenObjectFspVlanDhcpServerOption42edl(o["option4"], d, "option4")); err != nil {
		if vv, ok := fortiAPIPatch(o["option4"], "ObjectFspVlanDhcpServer-Option4"); ok {
			if err = d.Set("option4", vv); err != nil {
				return fmt.Errorf("Error reading option4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option4: %v", err)
		}
	}

	if err = d.Set("option5", flattenObjectFspVlanDhcpServerOption52edl(o["option5"], d, "option5")); err != nil {
		if vv, ok := fortiAPIPatch(o["option5"], "ObjectFspVlanDhcpServer-Option5"); ok {
			if err = d.Set("option5", vv); err != nil {
				return fmt.Errorf("Error reading option5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option5: %v", err)
		}
	}

	if err = d.Set("option6", flattenObjectFspVlanDhcpServerOption62edl(o["option6"], d, "option6")); err != nil {
		if vv, ok := fortiAPIPatch(o["option6"], "ObjectFspVlanDhcpServer-Option6"); ok {
			if err = d.Set("option6", vv); err != nil {
				return fmt.Errorf("Error reading option6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenObjectFspVlanDhcpServerOptions2edl(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "ObjectFspVlanDhcpServer-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenObjectFspVlanDhcpServerOptions2edl(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "ObjectFspVlanDhcpServer-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("relay_agent", flattenObjectFspVlanDhcpServerRelayAgent2edl(o["relay-agent"], d, "relay_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["relay-agent"], "ObjectFspVlanDhcpServer-RelayAgent"); ok {
			if err = d.Set("relay_agent", vv); err != nil {
				return fmt.Errorf("Error reading relay_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading relay_agent: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("reserved_address", flattenObjectFspVlanDhcpServerReservedAddress2edl(o["reserved-address"], d, "reserved_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["reserved-address"], "ObjectFspVlanDhcpServer-ReservedAddress"); ok {
				if err = d.Set("reserved_address", vv); err != nil {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading reserved_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reserved_address"); ok {
			if err = d.Set("reserved_address", flattenObjectFspVlanDhcpServerReservedAddress2edl(o["reserved-address"], d, "reserved_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["reserved-address"], "ObjectFspVlanDhcpServer-ReservedAddress"); ok {
					if err = d.Set("reserved_address", vv); err != nil {
						return fmt.Errorf("Error reading reserved_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenObjectFspVlanDhcpServerServerType2edl(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "ObjectFspVlanDhcpServer-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("shared_subnet", flattenObjectFspVlanDhcpServerSharedSubnet2edl(o["shared-subnet"], d, "shared_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["shared-subnet"], "ObjectFspVlanDhcpServer-SharedSubnet"); ok {
			if err = d.Set("shared_subnet", vv); err != nil {
				return fmt.Errorf("Error reading shared_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shared_subnet: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFspVlanDhcpServerStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFspVlanDhcpServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tftp_server", flattenObjectFspVlanDhcpServerTftpServer2edl(o["tftp-server"], d, "tftp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp-server"], "ObjectFspVlanDhcpServer-TftpServer"); ok {
			if err = d.Set("tftp_server", vv); err != nil {
				return fmt.Errorf("Error reading tftp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp_server: %v", err)
		}
	}

	if err = d.Set("timezone", flattenObjectFspVlanDhcpServerTimezone2edl(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "ObjectFspVlanDhcpServer-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("timezone_option", flattenObjectFspVlanDhcpServerTimezoneOption2edl(o["timezone-option"], d, "timezone_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone-option"], "ObjectFspVlanDhcpServer-TimezoneOption"); ok {
			if err = d.Set("timezone_option", vv); err != nil {
				return fmt.Errorf("Error reading timezone_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenObjectFspVlanDhcpServerVciMatch2edl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "ObjectFspVlanDhcpServer-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenObjectFspVlanDhcpServerVciString2edl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "ObjectFspVlanDhcpServer-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", flattenObjectFspVlanDhcpServerWifiAcService2edl(o["wifi-ac-service"], d, "wifi_ac_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac-service"], "ObjectFspVlanDhcpServer-WifiAcService"); ok {
			if err = d.Set("wifi_ac_service", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenObjectFspVlanDhcpServerWifiAc12edl(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac1"], "ObjectFspVlanDhcpServer-WifiAc1"); ok {
			if err = d.Set("wifi_ac1", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenObjectFspVlanDhcpServerWifiAc22edl(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac2"], "ObjectFspVlanDhcpServer-WifiAc2"); ok {
			if err = d.Set("wifi_ac2", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenObjectFspVlanDhcpServerWifiAc32edl(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac3"], "ObjectFspVlanDhcpServer-WifiAc3"); ok {
			if err = d.Set("wifi_ac3", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenObjectFspVlanDhcpServerWinsServer12edl(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "ObjectFspVlanDhcpServer-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenObjectFspVlanDhcpServerWinsServer22edl(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "ObjectFspVlanDhcpServer-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDhcpServerAutoConfiguration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerAutoManagedStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerConflictedIpTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsAuth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFspVlanDhcpServerDdnsKeyname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsServerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsUpdate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsUpdateOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDdnsZone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDefaultGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDhcpSettingsFromFortiipam2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsServer42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDnsService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerEnable2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandObjectFspVlanDhcpServerExcludeRangeEndIp2edl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerExcludeRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandObjectFspVlanDhcpServerExcludeRangeLeaseTime2edl(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDhcpServerExcludeRangeStartIp2edl(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDhcpServerExcludeRangeUciMatch2edl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDhcpServerExcludeRangeUciString2edl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDhcpServerExcludeRangeVciMatch2edl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDhcpServerExcludeRangeVciString2edl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerExcludeRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerExcludeRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerFilename2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerForticlientOnNetStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandObjectFspVlanDhcpServerIpRangeEndIp2edl(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerIpRangeId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lease-time"], _ = expandObjectFspVlanDhcpServerIpRangeLeaseTime2edl(d, i["lease_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFspVlanDhcpServerIpRangeStartIp2edl(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDhcpServerIpRangeUciMatch2edl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDhcpServerIpRangeUciString2edl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDhcpServerIpRangeVciMatch2edl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDhcpServerIpRangeVciString2edl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerIpRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerIpRangeVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerIpRangeVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerIpsecLeaseHold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerLeaseTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerMacAclDefaultAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNetmask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNextServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpServer32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerNtpService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOption12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOption22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOption32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOption42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOption52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOption62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["code"], _ = expandObjectFspVlanDhcpServerOptionsCode2edl(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerOptionsId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDhcpServerOptionsIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDhcpServerOptionsType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-match"], _ = expandObjectFspVlanDhcpServerOptionsUciMatch2edl(d, i["uci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uci-string"], _ = expandObjectFspVlanDhcpServerOptionsUciString2edl(d, i["uci_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFspVlanDhcpServerOptionsValue2edl(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-match"], _ = expandObjectFspVlanDhcpServerOptionsVciMatch2edl(d, i["vci_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vci-string"], _ = expandObjectFspVlanDhcpServerOptionsVciString2edl(d, i["vci_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerOptionsCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOptionsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsUciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsUciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerOptionsValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerOptionsVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerRelayAgent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectFspVlanDhcpServerReservedAddressAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectFspVlanDhcpServerReservedAddressCircuitId2edl(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandObjectFspVlanDhcpServerReservedAddressCircuitIdType2edl(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectFspVlanDhcpServerReservedAddressDescription2edl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFspVlanDhcpServerReservedAddressId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFspVlanDhcpServerReservedAddressIp2edl(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectFspVlanDhcpServerReservedAddressMac2edl(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectFspVlanDhcpServerReservedAddressRemoteId2edl(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandObjectFspVlanDhcpServerReservedAddressRemoteIdType2edl(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFspVlanDhcpServerReservedAddressType2edl(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFspVlanDhcpServerReservedAddressAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressCircuitIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressRemoteId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressRemoteIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerServerType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerSharedSubnet2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerTftpServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerTimezone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerTimezoneOption2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerVciMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerVciString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDhcpServerWifiAcService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWifiAc12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWifiAc22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWifiAc32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWinsServer12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerWinsServer22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanDhcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_configuration"); ok || d.HasChange("auto_configuration") {
		t, err := expandObjectFspVlanDhcpServerAutoConfiguration2edl(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("auto_managed_status"); ok || d.HasChange("auto_managed_status") {
		t, err := expandObjectFspVlanDhcpServerAutoManagedStatus2edl(d, v, "auto_managed_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-managed-status"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok || d.HasChange("conflicted_ip_timeout") {
		t, err := expandObjectFspVlanDhcpServerConflictedIpTimeout2edl(d, v, "conflicted_ip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandObjectFspVlanDhcpServerDdnsAuth2edl(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandObjectFspVlanDhcpServerDdnsKey2edl(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandObjectFspVlanDhcpServerDdnsKeyname2edl(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandObjectFspVlanDhcpServerDdnsServerIp2edl(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandObjectFspVlanDhcpServerDdnsTtl2edl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok || d.HasChange("ddns_update") {
		t, err := expandObjectFspVlanDhcpServerDdnsUpdate2edl(d, v, "ddns_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok || d.HasChange("ddns_update_override") {
		t, err := expandObjectFspVlanDhcpServerDdnsUpdateOverride2edl(d, v, "ddns_update_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandObjectFspVlanDhcpServerDdnsZone2edl(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandObjectFspVlanDhcpServerDefaultGateway2edl(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_settings_from_fortiipam"); ok || d.HasChange("dhcp_settings_from_fortiipam") {
		t, err := expandObjectFspVlanDhcpServerDhcpSettingsFromFortiipam2edl(d, v, "dhcp_settings_from_fortiipam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-settings-from-fortiipam"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandObjectFspVlanDhcpServerDnsServer12edl(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandObjectFspVlanDhcpServerDnsServer22edl(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandObjectFspVlanDhcpServerDnsServer32edl(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok || d.HasChange("dns_server4") {
		t, err := expandObjectFspVlanDhcpServerDnsServer42edl(d, v, "dns_server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandObjectFspVlanDhcpServerDnsService2edl(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectFspVlanDhcpServerDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("enable"); ok || d.HasChange("enable") {
		t, err := expandObjectFspVlanDhcpServerEnable2edl(d, v, "enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enable"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok || d.HasChange("exclude_range") {
		t, err := expandObjectFspVlanDhcpServerExcludeRange2edl(d, v, "exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandObjectFspVlanDhcpServerFilename2edl(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok || d.HasChange("forticlient_on_net_status") {
		t, err := expandObjectFspVlanDhcpServerForticlientOnNetStatus2edl(d, v, "forticlient_on_net_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanDhcpServerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandObjectFspVlanDhcpServerIpMode2edl(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandObjectFspVlanDhcpServerIpRange2edl(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok || d.HasChange("ipsec_lease_hold") {
		t, err := expandObjectFspVlanDhcpServerIpsecLeaseHold2edl(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandObjectFspVlanDhcpServerLeaseTime2edl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok || d.HasChange("mac_acl_default_action") {
		t, err := expandObjectFspVlanDhcpServerMacAclDefaultAction2edl(d, v, "mac_acl_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandObjectFspVlanDhcpServerNetmask2edl(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok || d.HasChange("next_server") {
		t, err := expandObjectFspVlanDhcpServerNextServer2edl(d, v, "next_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok || d.HasChange("ntp_server1") {
		t, err := expandObjectFspVlanDhcpServerNtpServer12edl(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok || d.HasChange("ntp_server2") {
		t, err := expandObjectFspVlanDhcpServerNtpServer22edl(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok || d.HasChange("ntp_server3") {
		t, err := expandObjectFspVlanDhcpServerNtpServer32edl(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok || d.HasChange("ntp_service") {
		t, err := expandObjectFspVlanDhcpServerNtpService2edl(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("option1"); ok || d.HasChange("option1") {
		t, err := expandObjectFspVlanDhcpServerOption12edl(d, v, "option1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option1"] = t
		}
	}

	if v, ok := d.GetOk("option2"); ok || d.HasChange("option2") {
		t, err := expandObjectFspVlanDhcpServerOption22edl(d, v, "option2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option2"] = t
		}
	}

	if v, ok := d.GetOk("option3"); ok || d.HasChange("option3") {
		t, err := expandObjectFspVlanDhcpServerOption32edl(d, v, "option3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option3"] = t
		}
	}

	if v, ok := d.GetOk("option4"); ok || d.HasChange("option4") {
		t, err := expandObjectFspVlanDhcpServerOption42edl(d, v, "option4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option4"] = t
		}
	}

	if v, ok := d.GetOk("option5"); ok || d.HasChange("option5") {
		t, err := expandObjectFspVlanDhcpServerOption52edl(d, v, "option5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option5"] = t
		}
	}

	if v, ok := d.GetOk("option6"); ok || d.HasChange("option6") {
		t, err := expandObjectFspVlanDhcpServerOption62edl(d, v, "option6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option6"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectFspVlanDhcpServerOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("relay_agent"); ok || d.HasChange("relay_agent") {
		t, err := expandObjectFspVlanDhcpServerRelayAgent2edl(d, v, "relay_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-agent"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok || d.HasChange("reserved_address") {
		t, err := expandObjectFspVlanDhcpServerReservedAddress2edl(d, v, "reserved_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandObjectFspVlanDhcpServerServerType2edl(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("shared_subnet"); ok || d.HasChange("shared_subnet") {
		t, err := expandObjectFspVlanDhcpServerSharedSubnet2edl(d, v, "shared_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-subnet"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFspVlanDhcpServerStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tftp_server"); ok || d.HasChange("tftp_server") {
		t, err := expandObjectFspVlanDhcpServerTftpServer2edl(d, v, "tftp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandObjectFspVlanDhcpServerTimezone2edl(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok || d.HasChange("timezone_option") {
		t, err := expandObjectFspVlanDhcpServerTimezoneOption2edl(d, v, "timezone_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandObjectFspVlanDhcpServerVciMatch2edl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandObjectFspVlanDhcpServerVciString2edl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac_service"); ok || d.HasChange("wifi_ac_service") {
		t, err := expandObjectFspVlanDhcpServerWifiAcService2edl(d, v, "wifi_ac_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac-service"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok || d.HasChange("wifi_ac1") {
		t, err := expandObjectFspVlanDhcpServerWifiAc12edl(d, v, "wifi_ac1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok || d.HasChange("wifi_ac2") {
		t, err := expandObjectFspVlanDhcpServerWifiAc22edl(d, v, "wifi_ac2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok || d.HasChange("wifi_ac3") {
		t, err := expandObjectFspVlanDhcpServerWifiAc32edl(d, v, "wifi_ac3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandObjectFspVlanDhcpServerWinsServer12edl(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandObjectFspVlanDhcpServerWinsServer22edl(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
