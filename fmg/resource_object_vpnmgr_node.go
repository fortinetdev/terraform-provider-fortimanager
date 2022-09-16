// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VPN node for VPN Manager. Must specify vpntable and scope member.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnmgrNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnmgrNodeCreate,
		Read:   resourceObjectVpnmgrNodeRead,
		Update: resourceObjectVpnmgrNodeUpdate,
		Delete: resourceObjectVpnmgrNodeDelete,

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
			"scopemember": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"add_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_ip_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authpasswd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"authusr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authusrgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"automatic_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_ra_giaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_mode": &schema.Schema{
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
			"encapsulation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_interface_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extgw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extgw_hubip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extgw_p2_per_net": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extgwip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hub_public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hub_iface": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"iface": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Computed: true,
			},
			"ipv4_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_exclude_range": &schema.Schema{
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
			"ipv4_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_split_exclude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_split_include": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2tp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode_cfg_ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"net_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"peergrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protected_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"seq": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spoke_zone": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"summary_addr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"seq": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"tunnel_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unity_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usrgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_interface_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vpn_zone": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"vpntable": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"xauthtype": &schema.Schema{
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

func resourceObjectVpnmgrNodeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnmgrNode(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNode resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnmgrNode(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNode resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVpnmgrNodeRead(d, m)
}

func resourceObjectVpnmgrNodeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnmgrNode(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNode resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnmgrNode(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNode resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVpnmgrNodeRead(d, m)
}

func resourceObjectVpnmgrNodeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnmgrNode(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnmgrNode resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnmgrNodeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnmgrNode(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNode resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnmgrNode(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNode resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnmgrNodeScopeMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnmgrNodeScopeMemberName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-ScopeMember-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectVpnmgrNodeScopeMemberVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-ScopeMember-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnmgrNodeScopeMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeScopeMemberVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAddRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAssignIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAssignIpFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAuthpasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodeAuthusr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAuthusrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeAutomaticRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeDhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeDhcpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeDnsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeEncapsulation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeExchangeInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeExtgw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeExtgwHubip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeExtgwP2PerNet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeExtgwip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeHubPublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeHubIface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodeIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnmgrNodeIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVpnmgrNodeIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectVpnmgrNodeIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-IpRange-StartIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnmgrNodeIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnmgrNodeIpv4ExcludeRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-Ipv4ExcludeRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVpnmgrNodeIpv4ExcludeRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-Ipv4ExcludeRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectVpnmgrNodeIpv4ExcludeRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-Ipv4ExcludeRange-StartIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4Netmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeL2Tp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeLocalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeModeCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeModeCfgIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeNetDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodePeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodePeergrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodePeerid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodePeertype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeProtectedSubnet(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := i["addr"]; ok {
			v := flattenObjectVpnmgrNodeProtectedSubnetAddr(i["addr"], d, pre_append)
			tmp["addr"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-ProtectedSubnet-Addr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectVpnmgrNodeProtectedSubnetSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-ProtectedSubnet-Seq")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnmgrNodeProtectedSubnetAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeProtectedSubnetSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodePublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeRouteOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeSpokeZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodeSummaryAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := i["addr"]; ok {
			v := flattenObjectVpnmgrNodeSummaryAddrAddr(i["addr"], d, pre_append)
			tmp["addr"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-SummaryAddr-Addr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectVpnmgrNodeSummaryAddrPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-SummaryAddr-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := i["seq"]; ok {
			v := flattenObjectVpnmgrNodeSummaryAddrSeq(i["seq"], d, pre_append)
			tmp["seq"] = fortiAPISubPartPatch(v, "ObjectVpnmgrNode-SummaryAddr-Seq")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnmgrNodeSummaryAddrAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeSummaryAddrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeSummaryAddrSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeTunnelSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeUnitySupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeUsrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeVpnInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeVpnZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodeVpntable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnmgrNodeXauthtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnmgrNode(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("scopemember", flattenObjectVpnmgrNodeScopeMember(o["scope member"], d, "scopemember")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope member"], "ObjectVpnmgrNode-ScopeMember"); ok {
				if err = d.Set("scopemember", vv); err != nil {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scopemember: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scopemember"); ok {
			if err = d.Set("scopemember", flattenObjectVpnmgrNodeScopeMember(o["scope member"], d, "scopemember")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope member"], "ObjectVpnmgrNode-ScopeMember"); ok {
					if err = d.Set("scopemember", vv); err != nil {
						return fmt.Errorf("Error reading scopemember: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			}
		}
	}

	if err = d.Set("add_route", flattenObjectVpnmgrNodeAddRoute(o["add-route"], d, "add_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-route"], "ObjectVpnmgrNode-AddRoute"); ok {
			if err = d.Set("add_route", vv); err != nil {
				return fmt.Errorf("Error reading add_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("assign_ip", flattenObjectVpnmgrNodeAssignIp(o["assign-ip"], d, "assign_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip"], "ObjectVpnmgrNode-AssignIp"); ok {
			if err = d.Set("assign_ip", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("assign_ip_from", flattenObjectVpnmgrNodeAssignIpFrom(o["assign-ip-from"], d, "assign_ip_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["assign-ip-from"], "ObjectVpnmgrNode-AssignIpFrom"); ok {
			if err = d.Set("assign_ip_from", vv); err != nil {
				return fmt.Errorf("Error reading assign_ip_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assign_ip_from: %v", err)
		}
	}

	if err = d.Set("authusr", flattenObjectVpnmgrNodeAuthusr(o["authusr"], d, "authusr")); err != nil {
		if vv, ok := fortiAPIPatch(o["authusr"], "ObjectVpnmgrNode-Authusr"); ok {
			if err = d.Set("authusr", vv); err != nil {
				return fmt.Errorf("Error reading authusr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authusr: %v", err)
		}
	}

	if err = d.Set("authusrgrp", flattenObjectVpnmgrNodeAuthusrgrp(o["authusrgrp"], d, "authusrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["authusrgrp"], "ObjectVpnmgrNode-Authusrgrp"); ok {
			if err = d.Set("authusrgrp", vv); err != nil {
				return fmt.Errorf("Error reading authusrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authusrgrp: %v", err)
		}
	}

	if err = d.Set("auto_configuration", flattenObjectVpnmgrNodeAutoConfiguration(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-configuration"], "ObjectVpnmgrNode-AutoConfiguration"); ok {
			if err = d.Set("auto_configuration", vv); err != nil {
				return fmt.Errorf("Error reading auto_configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("automatic_routing", flattenObjectVpnmgrNodeAutomaticRouting(o["automatic_routing"], d, "automatic_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["automatic_routing"], "ObjectVpnmgrNode-AutomaticRouting"); ok {
			if err = d.Set("automatic_routing", vv); err != nil {
				return fmt.Errorf("Error reading automatic_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading automatic_routing: %v", err)
		}
	}

	if err = d.Set("banner", flattenObjectVpnmgrNodeBanner(o["banner"], d, "banner")); err != nil {
		if vv, ok := fortiAPIPatch(o["banner"], "ObjectVpnmgrNode-Banner"); ok {
			if err = d.Set("banner", vv); err != nil {
				return fmt.Errorf("Error reading banner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading banner: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenObjectVpnmgrNodeDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-gateway"], "ObjectVpnmgrNode-DefaultGateway"); ok {
			if err = d.Set("default_gateway", vv); err != nil {
				return fmt.Errorf("Error reading default_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenObjectVpnmgrNodeDhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ra-giaddr"], "ObjectVpnmgrNode-DhcpRaGiaddr"); ok {
			if err = d.Set("dhcp_ra_giaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp_server", flattenObjectVpnmgrNodeDhcpServer(o["dhcp-server"], d, "dhcp_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-server"], "ObjectVpnmgrNode-DhcpServer"); ok {
			if err = d.Set("dhcp_server", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_server: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenObjectVpnmgrNodeDnsMode(o["dns-mode"], d, "dns_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-mode"], "ObjectVpnmgrNode-DnsMode"); ok {
			if err = d.Set("dns_mode", vv); err != nil {
				return fmt.Errorf("Error reading dns_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenObjectVpnmgrNodeDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-service"], "ObjectVpnmgrNode-DnsService"); ok {
			if err = d.Set("dns_service", vv); err != nil {
				return fmt.Errorf("Error reading dns_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectVpnmgrNodeDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectVpnmgrNode-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenObjectVpnmgrNodeEncapsulation(o["encapsulation"], d, "encapsulation")); err != nil {
		if vv, ok := fortiAPIPatch(o["encapsulation"], "ObjectVpnmgrNode-Encapsulation"); ok {
			if err = d.Set("encapsulation", vv); err != nil {
				return fmt.Errorf("Error reading encapsulation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("exchange_interface_ip", flattenObjectVpnmgrNodeExchangeInterfaceIp(o["exchange-interface-ip"], d, "exchange_interface_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["exchange-interface-ip"], "ObjectVpnmgrNode-ExchangeInterfaceIp"); ok {
			if err = d.Set("exchange_interface_ip", vv); err != nil {
				return fmt.Errorf("Error reading exchange_interface_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exchange_interface_ip: %v", err)
		}
	}

	if err = d.Set("extgw", flattenObjectVpnmgrNodeExtgw(o["extgw"], d, "extgw")); err != nil {
		if vv, ok := fortiAPIPatch(o["extgw"], "ObjectVpnmgrNode-Extgw"); ok {
			if err = d.Set("extgw", vv); err != nil {
				return fmt.Errorf("Error reading extgw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extgw: %v", err)
		}
	}

	if err = d.Set("extgw_hubip", flattenObjectVpnmgrNodeExtgwHubip(o["extgw_hubip"], d, "extgw_hubip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extgw_hubip"], "ObjectVpnmgrNode-ExtgwHubip"); ok {
			if err = d.Set("extgw_hubip", vv); err != nil {
				return fmt.Errorf("Error reading extgw_hubip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extgw_hubip: %v", err)
		}
	}

	if err = d.Set("extgw_p2_per_net", flattenObjectVpnmgrNodeExtgwP2PerNet(o["extgw_p2_per_net"], d, "extgw_p2_per_net")); err != nil {
		if vv, ok := fortiAPIPatch(o["extgw_p2_per_net"], "ObjectVpnmgrNode-ExtgwP2PerNet"); ok {
			if err = d.Set("extgw_p2_per_net", vv); err != nil {
				return fmt.Errorf("Error reading extgw_p2_per_net: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extgw_p2_per_net: %v", err)
		}
	}

	if err = d.Set("extgwip", flattenObjectVpnmgrNodeExtgwip(o["extgwip"], d, "extgwip")); err != nil {
		if vv, ok := fortiAPIPatch(o["extgwip"], "ObjectVpnmgrNode-Extgwip"); ok {
			if err = d.Set("extgwip", vv); err != nil {
				return fmt.Errorf("Error reading extgwip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extgwip: %v", err)
		}
	}

	if err = d.Set("hub_public_ip", flattenObjectVpnmgrNodeHubPublicIp(o["hub-public-ip"], d, "hub_public_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["hub-public-ip"], "ObjectVpnmgrNode-HubPublicIp"); ok {
			if err = d.Set("hub_public_ip", vv); err != nil {
				return fmt.Errorf("Error reading hub_public_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hub_public_ip: %v", err)
		}
	}

	if err = d.Set("hub_iface", flattenObjectVpnmgrNodeHubIface(o["hub_iface"], d, "hub_iface")); err != nil {
		if vv, ok := fortiAPIPatch(o["hub_iface"], "ObjectVpnmgrNode-HubIface"); ok {
			if err = d.Set("hub_iface", vv); err != nil {
				return fmt.Errorf("Error reading hub_iface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hub_iface: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectVpnmgrNodeId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectVpnmgrNode-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("iface", flattenObjectVpnmgrNodeIface(o["iface"], d, "iface")); err != nil {
		if vv, ok := fortiAPIPatch(o["iface"], "ObjectVpnmgrNode-Iface"); ok {
			if err = d.Set("iface", vv); err != nil {
				return fmt.Errorf("Error reading iface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenObjectVpnmgrNodeIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectVpnmgrNode-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenObjectVpnmgrNodeIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectVpnmgrNode-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenObjectVpnmgrNodeIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-lease-hold"], "ObjectVpnmgrNode-IpsecLeaseHold"); ok {
			if err = d.Set("ipsec_lease_hold", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server1", flattenObjectVpnmgrNodeIpv4DnsServer1(o["ipv4-dns-server1"], d, "ipv4_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server1"], "ObjectVpnmgrNode-Ipv4DnsServer1"); ok {
			if err = d.Set("ipv4_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server2", flattenObjectVpnmgrNodeIpv4DnsServer2(o["ipv4-dns-server2"], d, "ipv4_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server2"], "ObjectVpnmgrNode-Ipv4DnsServer2"); ok {
			if err = d.Set("ipv4_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server3", flattenObjectVpnmgrNodeIpv4DnsServer3(o["ipv4-dns-server3"], d, "ipv4_dns_server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-dns-server3"], "ObjectVpnmgrNode-Ipv4DnsServer3"); ok {
			if err = d.Set("ipv4_dns_server3", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenObjectVpnmgrNodeIpv4EndIp(o["ipv4-end-ip"], d, "ipv4_end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-end-ip"], "ObjectVpnmgrNode-Ipv4EndIp"); ok {
			if err = d.Set("ipv4_end_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv4_exclude_range", flattenObjectVpnmgrNodeIpv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv4-exclude-range"], "ObjectVpnmgrNode-Ipv4ExcludeRange"); ok {
				if err = d.Set("ipv4_exclude_range", vv); err != nil {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv4_exclude_range"); ok {
			if err = d.Set("ipv4_exclude_range", flattenObjectVpnmgrNodeIpv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv4-exclude-range"], "ObjectVpnmgrNode-Ipv4ExcludeRange"); ok {
					if err = d.Set("ipv4_exclude_range", vv); err != nil {
						return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_name", flattenObjectVpnmgrNodeIpv4Name(o["ipv4-name"], d, "ipv4_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-name"], "ObjectVpnmgrNode-Ipv4Name"); ok {
			if err = d.Set("ipv4_name", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_name: %v", err)
		}
	}

	if err = d.Set("ipv4_netmask", flattenObjectVpnmgrNodeIpv4Netmask(o["ipv4-netmask"], d, "ipv4_netmask")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-netmask"], "ObjectVpnmgrNode-Ipv4Netmask"); ok {
			if err = d.Set("ipv4_netmask", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_netmask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("ipv4_split_exclude", flattenObjectVpnmgrNodeIpv4SplitExclude(o["ipv4-split-exclude"], d, "ipv4_split_exclude")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-split-exclude"], "ObjectVpnmgrNode-Ipv4SplitExclude"); ok {
			if err = d.Set("ipv4_split_exclude", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv4_split_include", flattenObjectVpnmgrNodeIpv4SplitInclude(o["ipv4-split-include"], d, "ipv4_split_include")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-split-include"], "ObjectVpnmgrNode-Ipv4SplitInclude"); ok {
			if err = d.Set("ipv4_split_include", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_split_include: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_split_include: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenObjectVpnmgrNodeIpv4StartIp(o["ipv4-start-ip"], d, "ipv4_start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-start-ip"], "ObjectVpnmgrNode-Ipv4StartIp"); ok {
			if err = d.Set("ipv4_start_ip", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server1", flattenObjectVpnmgrNodeIpv4WinsServer1(o["ipv4-wins-server1"], d, "ipv4_wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-wins-server1"], "ObjectVpnmgrNode-Ipv4WinsServer1"); ok {
			if err = d.Set("ipv4_wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server2", flattenObjectVpnmgrNodeIpv4WinsServer2(o["ipv4-wins-server2"], d, "ipv4_wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-wins-server2"], "ObjectVpnmgrNode-Ipv4WinsServer2"); ok {
			if err = d.Set("ipv4_wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
		}
	}

	if err = d.Set("l2tp", flattenObjectVpnmgrNodeL2Tp(o["l2tp"], d, "l2tp")); err != nil {
		if vv, ok := fortiAPIPatch(o["l2tp"], "ObjectVpnmgrNode-L2Tp"); ok {
			if err = d.Set("l2tp", vv); err != nil {
				return fmt.Errorf("Error reading l2tp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l2tp: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenObjectVpnmgrNodeLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw"], "ObjectVpnmgrNode-LocalGw"); ok {
			if err = d.Set("local_gw", vv); err != nil {
				return fmt.Errorf("Error reading local_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("localid", flattenObjectVpnmgrNodeLocalid(o["localid"], d, "localid")); err != nil {
		if vv, ok := fortiAPIPatch(o["localid"], "ObjectVpnmgrNode-Localid"); ok {
			if err = d.Set("localid", vv); err != nil {
				return fmt.Errorf("Error reading localid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("mode_cfg", flattenObjectVpnmgrNodeModeCfg(o["mode-cfg"], d, "mode_cfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-cfg"], "ObjectVpnmgrNode-ModeCfg"); ok {
			if err = d.Set("mode_cfg", vv); err != nil {
				return fmt.Errorf("Error reading mode_cfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_cfg: %v", err)
		}
	}

	if err = d.Set("mode_cfg_ip_version", flattenObjectVpnmgrNodeModeCfgIpVersion(o["mode-cfg-ip-version"], d, "mode_cfg_ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-cfg-ip-version"], "ObjectVpnmgrNode-ModeCfgIpVersion"); ok {
			if err = d.Set("mode_cfg_ip_version", vv); err != nil {
				return fmt.Errorf("Error reading mode_cfg_ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_cfg_ip_version: %v", err)
		}
	}

	if err = d.Set("net_device", flattenObjectVpnmgrNodeNetDevice(o["net-device"], d, "net_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["net-device"], "ObjectVpnmgrNode-NetDevice"); ok {
			if err = d.Set("net_device", vv); err != nil {
				return fmt.Errorf("Error reading net_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading net_device: %v", err)
		}
	}

	if err = d.Set("peer", flattenObjectVpnmgrNodePeer(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "ObjectVpnmgrNode-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peergrp", flattenObjectVpnmgrNodePeergrp(o["peergrp"], d, "peergrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["peergrp"], "ObjectVpnmgrNode-Peergrp"); ok {
			if err = d.Set("peergrp", vv); err != nil {
				return fmt.Errorf("Error reading peergrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peergrp: %v", err)
		}
	}

	if err = d.Set("peerid", flattenObjectVpnmgrNodePeerid(o["peerid"], d, "peerid")); err != nil {
		if vv, ok := fortiAPIPatch(o["peerid"], "ObjectVpnmgrNode-Peerid"); ok {
			if err = d.Set("peerid", vv); err != nil {
				return fmt.Errorf("Error reading peerid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peerid: %v", err)
		}
	}

	if err = d.Set("peertype", flattenObjectVpnmgrNodePeertype(o["peertype"], d, "peertype")); err != nil {
		if vv, ok := fortiAPIPatch(o["peertype"], "ObjectVpnmgrNode-Peertype"); ok {
			if err = d.Set("peertype", vv); err != nil {
				return fmt.Errorf("Error reading peertype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peertype: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("protected_subnet", flattenObjectVpnmgrNodeProtectedSubnet(o["protected_subnet"], d, "protected_subnet")); err != nil {
			if vv, ok := fortiAPIPatch(o["protected_subnet"], "ObjectVpnmgrNode-ProtectedSubnet"); ok {
				if err = d.Set("protected_subnet", vv); err != nil {
					return fmt.Errorf("Error reading protected_subnet: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading protected_subnet: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("protected_subnet"); ok {
			if err = d.Set("protected_subnet", flattenObjectVpnmgrNodeProtectedSubnet(o["protected_subnet"], d, "protected_subnet")); err != nil {
				if vv, ok := fortiAPIPatch(o["protected_subnet"], "ObjectVpnmgrNode-ProtectedSubnet"); ok {
					if err = d.Set("protected_subnet", vv); err != nil {
						return fmt.Errorf("Error reading protected_subnet: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading protected_subnet: %v", err)
				}
			}
		}
	}

	if err = d.Set("public_ip", flattenObjectVpnmgrNodePublicIp(o["public-ip"], d, "public_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["public-ip"], "ObjectVpnmgrNode-PublicIp"); ok {
			if err = d.Set("public_ip", vv); err != nil {
				return fmt.Errorf("Error reading public_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading public_ip: %v", err)
		}
	}

	if err = d.Set("role", flattenObjectVpnmgrNodeRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ObjectVpnmgrNode-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("route_overlap", flattenObjectVpnmgrNodeRouteOverlap(o["route-overlap"], d, "route_overlap")); err != nil {
		if vv, ok := fortiAPIPatch(o["route-overlap"], "ObjectVpnmgrNode-RouteOverlap"); ok {
			if err = d.Set("route_overlap", vv); err != nil {
				return fmt.Errorf("Error reading route_overlap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route_overlap: %v", err)
		}
	}

	if err = d.Set("spoke_zone", flattenObjectVpnmgrNodeSpokeZone(o["spoke-zone"], d, "spoke_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["spoke-zone"], "ObjectVpnmgrNode-SpokeZone"); ok {
			if err = d.Set("spoke_zone", vv); err != nil {
				return fmt.Errorf("Error reading spoke_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spoke_zone: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("summary_addr", flattenObjectVpnmgrNodeSummaryAddr(o["summary_addr"], d, "summary_addr")); err != nil {
			if vv, ok := fortiAPIPatch(o["summary_addr"], "ObjectVpnmgrNode-SummaryAddr"); ok {
				if err = d.Set("summary_addr", vv); err != nil {
					return fmt.Errorf("Error reading summary_addr: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading summary_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_addr"); ok {
			if err = d.Set("summary_addr", flattenObjectVpnmgrNodeSummaryAddr(o["summary_addr"], d, "summary_addr")); err != nil {
				if vv, ok := fortiAPIPatch(o["summary_addr"], "ObjectVpnmgrNode-SummaryAddr"); ok {
					if err = d.Set("summary_addr", vv); err != nil {
						return fmt.Errorf("Error reading summary_addr: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading summary_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("tunnel_search", flattenObjectVpnmgrNodeTunnelSearch(o["tunnel-search"], d, "tunnel_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-search"], "ObjectVpnmgrNode-TunnelSearch"); ok {
			if err = d.Set("tunnel_search", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_search: %v", err)
		}
	}

	if err = d.Set("unity_support", flattenObjectVpnmgrNodeUnitySupport(o["unity-support"], d, "unity_support")); err != nil {
		if vv, ok := fortiAPIPatch(o["unity-support"], "ObjectVpnmgrNode-UnitySupport"); ok {
			if err = d.Set("unity_support", vv); err != nil {
				return fmt.Errorf("Error reading unity_support: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unity_support: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenObjectVpnmgrNodeUsrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["usrgrp"], "ObjectVpnmgrNode-Usrgrp"); ok {
			if err = d.Set("usrgrp", vv); err != nil {
				return fmt.Errorf("Error reading usrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("vpn_interface_priority", flattenObjectVpnmgrNodeVpnInterfacePriority(o["vpn-interface-priority"], d, "vpn_interface_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-interface-priority"], "ObjectVpnmgrNode-VpnInterfacePriority"); ok {
			if err = d.Set("vpn_interface_priority", vv); err != nil {
				return fmt.Errorf("Error reading vpn_interface_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_interface_priority: %v", err)
		}
	}

	if err = d.Set("vpn_zone", flattenObjectVpnmgrNodeVpnZone(o["vpn-zone"], d, "vpn_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn-zone"], "ObjectVpnmgrNode-VpnZone"); ok {
			if err = d.Set("vpn_zone", vv); err != nil {
				return fmt.Errorf("Error reading vpn_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_zone: %v", err)
		}
	}

	if err = d.Set("vpntable", flattenObjectVpnmgrNodeVpntable(o["vpntable"], d, "vpntable")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpntable"], "ObjectVpnmgrNode-Vpntable"); ok {
			if err = d.Set("vpntable", vv); err != nil {
				return fmt.Errorf("Error reading vpntable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpntable: %v", err)
		}
	}

	if err = d.Set("xauthtype", flattenObjectVpnmgrNodeXauthtype(o["xauthtype"], d, "xauthtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["xauthtype"], "ObjectVpnmgrNode-Xauthtype"); ok {
			if err = d.Set("xauthtype", vv); err != nil {
				return fmt.Errorf("Error reading xauthtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xauthtype: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnmgrNodeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnmgrNodeScopeMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnmgrNodeScopeMemberName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectVpnmgrNodeScopeMemberVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnmgrNodeScopeMemberName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeScopeMemberVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAddRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAssignIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAssignIpFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAuthpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnmgrNodeAuthusr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAuthusrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeAutomaticRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeBanner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeDhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeDhcpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeDnsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeEncapsulation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeExchangeInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeExtgw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeExtgwHubip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeExtgwP2PerNet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeExtgwip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeHubPublicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeHubIface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectVpnmgrNodeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectVpnmgrNodeIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandObjectVpnmgrNodeIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectVpnmgrNodeIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectVpnmgrNodeIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnmgrNodeIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["end-ip"], _ = expandObjectVpnmgrNodeIpv4ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectVpnmgrNodeIpv4ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectVpnmgrNodeIpv4ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnmgrNodeIpv4ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4Netmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeL2Tp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeLocalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeModeCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeModeCfgIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeNetDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodePeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectVpnmgrNodePeergrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodePeerid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodePeertype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeProtectedSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr"], _ = expandObjectVpnmgrNodeProtectedSubnetAddr(d, i["addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectVpnmgrNodeProtectedSubnetSeq(d, i["seq"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnmgrNodeProtectedSubnetAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeProtectedSubnetSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodePublicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeRouteOverlap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeSpokeZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectVpnmgrNodeSummaryAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr"], _ = expandObjectVpnmgrNodeSummaryAddrAddr(d, i["addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectVpnmgrNodeSummaryAddrPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq"], _ = expandObjectVpnmgrNodeSummaryAddrSeq(d, i["seq"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnmgrNodeSummaryAddrAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeSummaryAddrPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeSummaryAddrSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeTunnelSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeUnitySupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeUsrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeVpnInterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeVpnZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectVpnmgrNodeVpntable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectVpnmgrNodeXauthtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnmgrNode(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("scopemember"); ok || d.HasChange("scopemember") {
		t, err := expandObjectVpnmgrNodeScopeMember(d, v, "scopemember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope member"] = t
		}
	}

	if v, ok := d.GetOk("add_route"); ok || d.HasChange("add_route") {
		t, err := expandObjectVpnmgrNodeAddRoute(d, v, "add_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip"); ok || d.HasChange("assign_ip") {
		t, err := expandObjectVpnmgrNodeAssignIp(d, v, "assign_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip_from"); ok || d.HasChange("assign_ip_from") {
		t, err := expandObjectVpnmgrNodeAssignIpFrom(d, v, "assign_ip_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip-from"] = t
		}
	}

	if v, ok := d.GetOk("authpasswd"); ok || d.HasChange("authpasswd") {
		t, err := expandObjectVpnmgrNodeAuthpasswd(d, v, "authpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authpasswd"] = t
		}
	}

	if v, ok := d.GetOk("authusr"); ok || d.HasChange("authusr") {
		t, err := expandObjectVpnmgrNodeAuthusr(d, v, "authusr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusr"] = t
		}
	}

	if v, ok := d.GetOk("authusrgrp"); ok || d.HasChange("authusrgrp") {
		t, err := expandObjectVpnmgrNodeAuthusrgrp(d, v, "authusrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusrgrp"] = t
		}
	}

	if v, ok := d.GetOk("auto_configuration"); ok || d.HasChange("auto_configuration") {
		t, err := expandObjectVpnmgrNodeAutoConfiguration(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("automatic_routing"); ok || d.HasChange("automatic_routing") {
		t, err := expandObjectVpnmgrNodeAutomaticRouting(d, v, "automatic_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["automatic_routing"] = t
		}
	}

	if v, ok := d.GetOk("banner"); ok || d.HasChange("banner") {
		t, err := expandObjectVpnmgrNodeBanner(d, v, "banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok || d.HasChange("default_gateway") {
		t, err := expandObjectVpnmgrNodeDefaultGateway(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok || d.HasChange("dhcp_ra_giaddr") {
		t, err := expandObjectVpnmgrNodeDhcpRaGiaddr(d, v, "dhcp_ra_giaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok || d.HasChange("dhcp_server") {
		t, err := expandObjectVpnmgrNodeDhcpServer(d, v, "dhcp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok || d.HasChange("dns_mode") {
		t, err := expandObjectVpnmgrNodeDnsMode(d, v, "dns_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok || d.HasChange("dns_service") {
		t, err := expandObjectVpnmgrNodeDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectVpnmgrNodeDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok || d.HasChange("encapsulation") {
		t, err := expandObjectVpnmgrNodeEncapsulation(d, v, "encapsulation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("exchange_interface_ip"); ok || d.HasChange("exchange_interface_ip") {
		t, err := expandObjectVpnmgrNodeExchangeInterfaceIp(d, v, "exchange_interface_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-interface-ip"] = t
		}
	}

	if v, ok := d.GetOk("extgw"); ok || d.HasChange("extgw") {
		t, err := expandObjectVpnmgrNodeExtgw(d, v, "extgw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extgw"] = t
		}
	}

	if v, ok := d.GetOk("extgw_hubip"); ok || d.HasChange("extgw_hubip") {
		t, err := expandObjectVpnmgrNodeExtgwHubip(d, v, "extgw_hubip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extgw_hubip"] = t
		}
	}

	if v, ok := d.GetOk("extgw_p2_per_net"); ok || d.HasChange("extgw_p2_per_net") {
		t, err := expandObjectVpnmgrNodeExtgwP2PerNet(d, v, "extgw_p2_per_net")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extgw_p2_per_net"] = t
		}
	}

	if v, ok := d.GetOk("extgwip"); ok || d.HasChange("extgwip") {
		t, err := expandObjectVpnmgrNodeExtgwip(d, v, "extgwip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extgwip"] = t
		}
	}

	if v, ok := d.GetOk("hub_public_ip"); ok || d.HasChange("hub_public_ip") {
		t, err := expandObjectVpnmgrNodeHubPublicIp(d, v, "hub_public_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hub-public-ip"] = t
		}
	}

	if v, ok := d.GetOk("hub_iface"); ok || d.HasChange("hub_iface") {
		t, err := expandObjectVpnmgrNodeHubIface(d, v, "hub_iface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hub_iface"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectVpnmgrNodeId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("iface"); ok || d.HasChange("iface") {
		t, err := expandObjectVpnmgrNodeIface(d, v, "iface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iface"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandObjectVpnmgrNodeIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok || d.HasChange("ipsec_lease_hold") {
		t, err := expandObjectVpnmgrNodeIpsecLeaseHold(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server1"); ok || d.HasChange("ipv4_dns_server1") {
		t, err := expandObjectVpnmgrNodeIpv4DnsServer1(d, v, "ipv4_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server2"); ok || d.HasChange("ipv4_dns_server2") {
		t, err := expandObjectVpnmgrNodeIpv4DnsServer2(d, v, "ipv4_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server3"); ok || d.HasChange("ipv4_dns_server3") {
		t, err := expandObjectVpnmgrNodeIpv4DnsServer3(d, v, "ipv4_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok || d.HasChange("ipv4_end_ip") {
		t, err := expandObjectVpnmgrNodeIpv4EndIp(d, v, "ipv4_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_exclude_range"); ok || d.HasChange("ipv4_exclude_range") {
		t, err := expandObjectVpnmgrNodeIpv4ExcludeRange(d, v, "ipv4_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_name"); ok || d.HasChange("ipv4_name") {
		t, err := expandObjectVpnmgrNodeIpv4Name(d, v, "ipv4_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-name"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_netmask"); ok || d.HasChange("ipv4_netmask") {
		t, err := expandObjectVpnmgrNodeIpv4Netmask(d, v, "ipv4_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_exclude"); ok || d.HasChange("ipv4_split_exclude") {
		t, err := expandObjectVpnmgrNodeIpv4SplitExclude(d, v, "ipv4_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_include"); ok || d.HasChange("ipv4_split_include") {
		t, err := expandObjectVpnmgrNodeIpv4SplitInclude(d, v, "ipv4_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-include"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok || d.HasChange("ipv4_start_ip") {
		t, err := expandObjectVpnmgrNodeIpv4StartIp(d, v, "ipv4_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server1"); ok || d.HasChange("ipv4_wins_server1") {
		t, err := expandObjectVpnmgrNodeIpv4WinsServer1(d, v, "ipv4_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server2"); ok || d.HasChange("ipv4_wins_server2") {
		t, err := expandObjectVpnmgrNodeIpv4WinsServer2(d, v, "ipv4_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("l2tp"); ok || d.HasChange("l2tp") {
		t, err := expandObjectVpnmgrNodeL2Tp(d, v, "l2tp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tp"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok || d.HasChange("local_gw") {
		t, err := expandObjectVpnmgrNodeLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("localid"); ok || d.HasChange("localid") {
		t, err := expandObjectVpnmgrNodeLocalid(d, v, "localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg"); ok || d.HasChange("mode_cfg") {
		t, err := expandObjectVpnmgrNodeModeCfg(d, v, "mode_cfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg_ip_version"); ok || d.HasChange("mode_cfg_ip_version") {
		t, err := expandObjectVpnmgrNodeModeCfgIpVersion(d, v, "mode_cfg_ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg-ip-version"] = t
		}
	}

	if v, ok := d.GetOk("net_device"); ok || d.HasChange("net_device") {
		t, err := expandObjectVpnmgrNodeNetDevice(d, v, "net_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["net-device"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandObjectVpnmgrNodePeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("peergrp"); ok || d.HasChange("peergrp") {
		t, err := expandObjectVpnmgrNodePeergrp(d, v, "peergrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peergrp"] = t
		}
	}

	if v, ok := d.GetOk("peerid"); ok || d.HasChange("peerid") {
		t, err := expandObjectVpnmgrNodePeerid(d, v, "peerid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerid"] = t
		}
	}

	if v, ok := d.GetOk("peertype"); ok || d.HasChange("peertype") {
		t, err := expandObjectVpnmgrNodePeertype(d, v, "peertype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peertype"] = t
		}
	}

	if v, ok := d.GetOk("protected_subnet"); ok || d.HasChange("protected_subnet") {
		t, err := expandObjectVpnmgrNodeProtectedSubnet(d, v, "protected_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protected_subnet"] = t
		}
	}

	if v, ok := d.GetOk("public_ip"); ok || d.HasChange("public_ip") {
		t, err := expandObjectVpnmgrNodePublicIp(d, v, "public_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-ip"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandObjectVpnmgrNodeRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("route_overlap"); ok || d.HasChange("route_overlap") {
		t, err := expandObjectVpnmgrNodeRouteOverlap(d, v, "route_overlap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-overlap"] = t
		}
	}

	if v, ok := d.GetOk("spoke_zone"); ok || d.HasChange("spoke_zone") {
		t, err := expandObjectVpnmgrNodeSpokeZone(d, v, "spoke_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spoke-zone"] = t
		}
	}

	if v, ok := d.GetOk("summary_addr"); ok || d.HasChange("summary_addr") {
		t, err := expandObjectVpnmgrNodeSummaryAddr(d, v, "summary_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary_addr"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_search"); ok || d.HasChange("tunnel_search") {
		t, err := expandObjectVpnmgrNodeTunnelSearch(d, v, "tunnel_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-search"] = t
		}
	}

	if v, ok := d.GetOk("unity_support"); ok || d.HasChange("unity_support") {
		t, err := expandObjectVpnmgrNodeUnitySupport(d, v, "unity_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unity-support"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok || d.HasChange("usrgrp") {
		t, err := expandObjectVpnmgrNodeUsrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	if v, ok := d.GetOk("vpn_interface_priority"); ok || d.HasChange("vpn_interface_priority") {
		t, err := expandObjectVpnmgrNodeVpnInterfacePriority(d, v, "vpn_interface_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-interface-priority"] = t
		}
	}

	if v, ok := d.GetOk("vpn_zone"); ok || d.HasChange("vpn_zone") {
		t, err := expandObjectVpnmgrNodeVpnZone(d, v, "vpn_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-zone"] = t
		}
	}

	if v, ok := d.GetOk("vpntable"); ok || d.HasChange("vpntable") {
		t, err := expandObjectVpnmgrNodeVpntable(d, v, "vpntable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntable"] = t
		}
	}

	if v, ok := d.GetOk("xauthtype"); ok || d.HasChange("xauthtype") {
		t, err := expandObjectVpnmgrNodeXauthtype(d, v, "xauthtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xauthtype"] = t
		}
	}

	return &obj, nil
}
