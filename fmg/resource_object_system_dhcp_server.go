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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemDhcpServerCreate,
		Read:   resourceObjectSystemDhcpServerRead,
		Update: resourceObjectSystemDhcpServerUpdate,
		Delete: resourceObjectSystemDhcpServerDelete,

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
			"auto_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_managed_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"conflicted_ip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"ddns_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ddns_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_update_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_settings_from_fortiipam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
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
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
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
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
			},
			"mac_acl_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_service": &schema.Schema{
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
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"reserved_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tftp_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timezone_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"wifi_ac_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_ac1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_ac2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_ac3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server2": &schema.Schema{
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

func resourceObjectSystemDhcpServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemDhcpServer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemDhcpServer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemDhcpServer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemDhcpServerRead(d, m)
}

func resourceObjectSystemDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDhcpServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemDhcpServer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDhcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemDhcpServerRead(d, m)
}

func resourceObjectSystemDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemDhcpServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemDhcpServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemDhcpServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemDhcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemDhcpServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemDhcpServerAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerAutoManagedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDhcpSettingsFromFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemDhcpServerExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemDhcpServerExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectSystemDhcpServerExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ExcludeRange-StartIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemDhcpServerExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemDhcpServerIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemDhcpServerIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectSystemDhcpServerIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-IpRange-StartIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemDhcpServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemDhcpServerOptionsCode(i["code"], d, pre_append)
			tmp["code"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-Options-Code")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemDhcpServerOptionsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-Options-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectSystemDhcpServerOptionsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-Options-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemDhcpServerOptionsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-Options-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectSystemDhcpServerOptionsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-Options-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemDhcpServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDhcpServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemDhcpServerReservedAddressAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressCircuitId(i["circuit-id"], d, pre_append)
			tmp["circuit_id"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-CircuitId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressCircuitIdType(i["circuit-id-type"], d, pre_append)
			tmp["circuit_id_type"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-CircuitIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressRemoteId(i["remote-id"], d, pre_append)
			tmp["remote_id"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-RemoteId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressRemoteIdType(i["remote-id-type"], d, pre_append)
			tmp["remote_id_type"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-RemoteIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemDhcpServerReservedAddressType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemDhcpServer-ReservedAddress-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemDhcpServerReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDhcpServerTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDhcpServerWifiAcService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auto_configuration", flattenObjectSystemDhcpServerAutoConfiguration(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-configuration"], "ObjectSystemDhcpServer-AutoConfiguration"); ok {
			if err = d.Set("auto_configuration", vv); err != nil {
				return fmt.Errorf("Error reading auto_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("auto_managed_status", flattenObjectSystemDhcpServerAutoManagedStatus(o["auto-managed-status"], d, "auto_managed_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-managed-status"], "ObjectSystemDhcpServer-AutoManagedStatus"); ok {
			if err = d.Set("auto_managed_status", vv); err != nil {
				return fmt.Errorf("Error reading auto_managed_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_managed_status: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenObjectSystemDhcpServerConflictedIpTimeout(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["conflicted-ip-timeout"], "ObjectSystemDhcpServer-ConflictedIpTimeout"); ok {
			if err = d.Set("conflicted_ip_timeout", vv); err != nil {
				return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenObjectSystemDhcpServerDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-auth"], "ObjectSystemDhcpServer-DdnsAuth"); ok {
			if err = d.Set("ddns_auth", vv); err != nil {
				return fmt.Errorf("Error reading ddns_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_key", flattenObjectSystemDhcpServerDdnsKey(o["ddns-key"], d, "ddns_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-key"], "ObjectSystemDhcpServer-DdnsKey"); ok {
			if err = d.Set("ddns_key", vv); err != nil {
				return fmt.Errorf("Error reading ddns_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_key: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenObjectSystemDhcpServerDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-keyname"], "ObjectSystemDhcpServer-DdnsKeyname"); ok {
			if err = d.Set("ddns_keyname", vv); err != nil {
				return fmt.Errorf("Error reading ddns_keyname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenObjectSystemDhcpServerDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-server-ip"], "ObjectSystemDhcpServer-DdnsServerIp"); ok {
			if err = d.Set("ddns_server_ip", vv); err != nil {
				return fmt.Errorf("Error reading ddns_server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenObjectSystemDhcpServerDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-ttl"], "ObjectSystemDhcpServer-DdnsTtl"); ok {
			if err = d.Set("ddns_ttl", vv); err != nil {
				return fmt.Errorf("Error reading ddns_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenObjectSystemDhcpServerDdnsUpdate(o["ddns-update"], d, "ddns_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update"], "ObjectSystemDhcpServer-DdnsUpdate"); ok {
			if err = d.Set("ddns_update", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenObjectSystemDhcpServerDdnsUpdateOverride(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-update-override"], "ObjectSystemDhcpServer-DdnsUpdateOverride"); ok {
			if err = d.Set("ddns_update_override", vv); err != nil {
				return fmt.Errorf("Error reading ddns_update_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenObjectSystemDhcpServerDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddns-zone"], "ObjectSystemDhcpServer-DdnsZone"); ok {
			if err = d.Set("ddns_zone", vv); err != nil {
				return fmt.Errorf("Error reading ddns_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenObjectSystemDhcpServerDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "ObjectSystemDhcpServer-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_settings_from_fortiipam", flattenObjectSystemDhcpServerDhcpSettingsFromFortiipam(o["dhcp-settings-from-fortiipam"], d, "dhcp_settings_from_fortiipam")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-settings-from-fortiipam"], "ObjectSystemDhcpServer-DhcpSettingsFromFortiipam"); ok {
			if err = d.Set("dhcp_settings_from_fortiipam", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_settings_from_fortiipam: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenObjectSystemDhcpServerDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "ObjectSystemDhcpServer-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenObjectSystemDhcpServerDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "ObjectSystemDhcpServer-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenObjectSystemDhcpServerDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server3"], "ObjectSystemDhcpServer-DnsServer3"); ok {
			if err = d.Set("dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenObjectSystemDhcpServerDnsServer4(o["dns-server4"], d, "dns_server4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server4"], "ObjectSystemDhcpServer-DnsServer4"); ok {
			if err = d.Set("dns_server4", vv); err != nil {
				return fmt.Errorf("Error reading dns_server4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenObjectSystemDhcpServerDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "ObjectSystemDhcpServer-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectSystemDhcpServerDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectSystemDhcpServer-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exclude_range", flattenObjectSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["exclude-range"], "ObjectSystemDhcpServer-ExcludeRange"); ok {
				if err = d.Set("exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exclude_range"); ok {
			if err = d.Set("exclude_range", flattenObjectSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["exclude-range"], "ObjectSystemDhcpServer-ExcludeRange"); ok {
					if err = d.Set("exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("filename", flattenObjectSystemDhcpServerFilename(o["filename"], d, "filename")); err != nil {
		if vv, ok := fortiAPIPatch(o["filename"], "ObjectSystemDhcpServer-Filename"); ok {
			if err = d.Set("filename", vv); err != nil {
				return fmt.Errorf("Error reading filename: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenObjectSystemDhcpServerForticlientOnNetStatus(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-on-net-status"], "ObjectSystemDhcpServer-ForticlientOnNetStatus"); ok {
			if err = d.Set("forticlient_on_net_status", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemDhcpServerId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemDhcpServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectSystemDhcpServerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectSystemDhcpServer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenObjectSystemDhcpServerIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "ObjectSystemDhcpServer-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenObjectSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectSystemDhcpServer-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenObjectSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectSystemDhcpServer-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenObjectSystemDhcpServerIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-lease-hold"], "ObjectSystemDhcpServer-IpsecLeaseHold"); ok {
			if err = d.Set("ipsec_lease_hold", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenObjectSystemDhcpServerLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "ObjectSystemDhcpServer-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenObjectSystemDhcpServerMacAclDefaultAction(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-acl-default-action"], "ObjectSystemDhcpServer-MacAclDefaultAction"); ok {
			if err = d.Set("mac_acl_default_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("netmask", flattenObjectSystemDhcpServerNetmask(o["netmask"], d, "netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["netmask"], "ObjectSystemDhcpServer-Netmask"); ok {
			if err = d.Set("netmask", vv); err != nil {
				return fmt.Errorf("Error reading netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("next_server", flattenObjectSystemDhcpServerNextServer(o["next-server"], d, "next_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-server"], "ObjectSystemDhcpServer-NextServer"); ok {
			if err = d.Set("next_server", vv); err != nil {
				return fmt.Errorf("Error reading next_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenObjectSystemDhcpServerNtpServer1(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server1"], "ObjectSystemDhcpServer-NtpServer1"); ok {
			if err = d.Set("ntp_server1", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenObjectSystemDhcpServerNtpServer2(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server2"], "ObjectSystemDhcpServer-NtpServer2"); ok {
			if err = d.Set("ntp_server2", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenObjectSystemDhcpServerNtpServer3(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-server3"], "ObjectSystemDhcpServer-NtpServer3"); ok {
			if err = d.Set("ntp_server3", vv); err != nil {
				return fmt.Errorf("Error reading ntp_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenObjectSystemDhcpServerNtpService(o["ntp-service"], d, "ntp_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntp-service"], "ObjectSystemDhcpServer-NtpService"); ok {
			if err = d.Set("ntp_service", vv); err != nil {
				return fmt.Errorf("Error reading ntp_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("options", flattenObjectSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
			if vv, ok := fortiAPIPatch(o["options"], "ObjectSystemDhcpServer-Options"); ok {
				if err = d.Set("options", vv); err != nil {
					return fmt.Errorf("Error reading options: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenObjectSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
				if vv, ok := fortiAPIPatch(o["options"], "ObjectSystemDhcpServer-Options"); ok {
					if err = d.Set("options", vv); err != nil {
						return fmt.Errorf("Error reading options: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("reserved_address", flattenObjectSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["reserved-address"], "ObjectSystemDhcpServer-ReservedAddress"); ok {
				if err = d.Set("reserved_address", vv); err != nil {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading reserved_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reserved_address"); ok {
			if err = d.Set("reserved_address", flattenObjectSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["reserved-address"], "ObjectSystemDhcpServer-ReservedAddress"); ok {
					if err = d.Set("reserved_address", vv); err != nil {
						return fmt.Errorf("Error reading reserved_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenObjectSystemDhcpServerServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "ObjectSystemDhcpServer-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSystemDhcpServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSystemDhcpServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tftp_server", flattenObjectSystemDhcpServerTftpServer(o["tftp-server"], d, "tftp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tftp-server"], "ObjectSystemDhcpServer-TftpServer"); ok {
			if err = d.Set("tftp_server", vv); err != nil {
				return fmt.Errorf("Error reading tftp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tftp_server: %v", err)
		}
	}

	if err = d.Set("timezone", flattenObjectSystemDhcpServerTimezone(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "ObjectSystemDhcpServer-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("timezone_option", flattenObjectSystemDhcpServerTimezoneOption(o["timezone-option"], d, "timezone_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone-option"], "ObjectSystemDhcpServer-TimezoneOption"); ok {
			if err = d.Set("timezone_option", vv); err != nil {
				return fmt.Errorf("Error reading timezone_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenObjectSystemDhcpServerVciMatch(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "ObjectSystemDhcpServer-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenObjectSystemDhcpServerVciString(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "ObjectSystemDhcpServer-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", flattenObjectSystemDhcpServerWifiAcService(o["wifi-ac-service"], d, "wifi_ac_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac-service"], "ObjectSystemDhcpServer-WifiAcService"); ok {
			if err = d.Set("wifi_ac_service", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenObjectSystemDhcpServerWifiAc1(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac1"], "ObjectSystemDhcpServer-WifiAc1"); ok {
			if err = d.Set("wifi_ac1", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenObjectSystemDhcpServerWifiAc2(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac2"], "ObjectSystemDhcpServer-WifiAc2"); ok {
			if err = d.Set("wifi_ac2", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenObjectSystemDhcpServerWifiAc3(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi-ac3"], "ObjectSystemDhcpServer-WifiAc3"); ok {
			if err = d.Set("wifi_ac3", vv); err != nil {
				return fmt.Errorf("Error reading wifi_ac3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenObjectSystemDhcpServerWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "ObjectSystemDhcpServer-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenObjectSystemDhcpServerWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "ObjectSystemDhcpServer-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemDhcpServerAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerAutoManagedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDhcpSettingsFromFortiipam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDnsServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectSystemDhcpServerExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemDhcpServerExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectSystemDhcpServerExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemDhcpServerExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectSystemDhcpServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemDhcpServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectSystemDhcpServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemDhcpServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerNextServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerNtpServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerNtpServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerNtpServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerNtpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["code"], _ = expandObjectSystemDhcpServerOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemDhcpServerOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectSystemDhcpServerOptionsIp(d, i["ip"], pre_append)
		} else {
			tmp["ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemDhcpServerOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectSystemDhcpServerOptionsValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemDhcpServerOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerOptionsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDhcpServerOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectSystemDhcpServerReservedAddressAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id"], _ = expandObjectSystemDhcpServerReservedAddressCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-id-type"], _ = expandObjectSystemDhcpServerReservedAddressCircuitIdType(d, i["circuit_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectSystemDhcpServerReservedAddressDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemDhcpServerReservedAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectSystemDhcpServerReservedAddressIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectSystemDhcpServerReservedAddressMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id"], _ = expandObjectSystemDhcpServerReservedAddressRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-id-type"], _ = expandObjectSystemDhcpServerReservedAddressRemoteIdType(d, i["remote_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemDhcpServerReservedAddressType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemDhcpServerReservedAddressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressCircuitIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressRemoteIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDhcpServerTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerTimezoneOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDhcpServerWifiAcService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerWifiAc1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerWifiAc2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerWifiAc3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemDhcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_configuration"); ok || d.HasChange("auto_configuration") {
		t, err := expandObjectSystemDhcpServerAutoConfiguration(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("auto_managed_status"); ok || d.HasChange("auto_managed_status") {
		t, err := expandObjectSystemDhcpServerAutoManagedStatus(d, v, "auto_managed_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-managed-status"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok || d.HasChange("conflicted_ip_timeout") {
		t, err := expandObjectSystemDhcpServerConflictedIpTimeout(d, v, "conflicted_ip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok || d.HasChange("ddns_auth") {
		t, err := expandObjectSystemDhcpServerDdnsAuth(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok || d.HasChange("ddns_key") {
		t, err := expandObjectSystemDhcpServerDdnsKey(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok || d.HasChange("ddns_keyname") {
		t, err := expandObjectSystemDhcpServerDdnsKeyname(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok || d.HasChange("ddns_server_ip") {
		t, err := expandObjectSystemDhcpServerDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok || d.HasChange("ddns_ttl") {
		t, err := expandObjectSystemDhcpServerDdnsTtl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok || d.HasChange("ddns_update") {
		t, err := expandObjectSystemDhcpServerDdnsUpdate(d, v, "ddns_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok || d.HasChange("ddns_update_override") {
		t, err := expandObjectSystemDhcpServerDdnsUpdateOverride(d, v, "ddns_update_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok || d.HasChange("ddns_zone") {
		t, err := expandObjectSystemDhcpServerDdnsZone(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandObjectSystemDhcpServerDefaultGateway(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_settings_from_fortiipam"); ok || d.HasChange("dhcp_settings_from_fortiipam") {
		t, err := expandObjectSystemDhcpServerDhcpSettingsFromFortiipam(d, v, "dhcp_settings_from_fortiipam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-settings-from-fortiipam"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandObjectSystemDhcpServerDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandObjectSystemDhcpServerDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok || d.HasChange("dns_server3") {
		t, err := expandObjectSystemDhcpServerDnsServer3(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok || d.HasChange("dns_server4") {
		t, err := expandObjectSystemDhcpServerDnsServer4(d, v, "dns_server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandObjectSystemDhcpServerDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectSystemDhcpServerDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok || d.HasChange("exclude_range") {
		t, err := expandObjectSystemDhcpServerExcludeRange(d, v, "exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok || d.HasChange("filename") {
		t, err := expandObjectSystemDhcpServerFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok || d.HasChange("forticlient_on_net_status") {
		t, err := expandObjectSystemDhcpServerForticlientOnNetStatus(d, v, "forticlient_on_net_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectSystemDhcpServerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectSystemDhcpServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandObjectSystemDhcpServerIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandObjectSystemDhcpServerIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok || d.HasChange("ipsec_lease_hold") {
		t, err := expandObjectSystemDhcpServerIpsecLeaseHold(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandObjectSystemDhcpServerLeaseTime(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok || d.HasChange("mac_acl_default_action") {
		t, err := expandObjectSystemDhcpServerMacAclDefaultAction(d, v, "mac_acl_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok || d.HasChange("netmask") {
		t, err := expandObjectSystemDhcpServerNetmask(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok || d.HasChange("next_server") {
		t, err := expandObjectSystemDhcpServerNextServer(d, v, "next_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok || d.HasChange("ntp_server1") {
		t, err := expandObjectSystemDhcpServerNtpServer1(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok || d.HasChange("ntp_server2") {
		t, err := expandObjectSystemDhcpServerNtpServer2(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok || d.HasChange("ntp_server3") {
		t, err := expandObjectSystemDhcpServerNtpServer3(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok || d.HasChange("ntp_service") {
		t, err := expandObjectSystemDhcpServerNtpService(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectSystemDhcpServerOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok || d.HasChange("reserved_address") {
		t, err := expandObjectSystemDhcpServerReservedAddress(d, v, "reserved_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandObjectSystemDhcpServerServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSystemDhcpServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tftp_server"); ok || d.HasChange("tftp_server") {
		t, err := expandObjectSystemDhcpServerTftpServer(d, v, "tftp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandObjectSystemDhcpServerTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok || d.HasChange("timezone_option") {
		t, err := expandObjectSystemDhcpServerTimezoneOption(d, v, "timezone_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandObjectSystemDhcpServerVciMatch(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandObjectSystemDhcpServerVciString(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac_service"); ok || d.HasChange("wifi_ac_service") {
		t, err := expandObjectSystemDhcpServerWifiAcService(d, v, "wifi_ac_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac-service"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok || d.HasChange("wifi_ac1") {
		t, err := expandObjectSystemDhcpServerWifiAc1(d, v, "wifi_ac1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok || d.HasChange("wifi_ac2") {
		t, err := expandObjectSystemDhcpServerWifiAc2(d, v, "wifi_ac2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok || d.HasChange("wifi_ac3") {
		t, err := expandObjectSystemDhcpServerWifiAc3(d, v, "wifi_ac3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandObjectSystemDhcpServerWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandObjectSystemDhcpServerWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
