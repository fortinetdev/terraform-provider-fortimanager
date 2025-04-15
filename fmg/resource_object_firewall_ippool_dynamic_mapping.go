// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 IP pools.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallIppoolDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallIppoolDynamicMappingCreate,
		Read:   resourceObjectFirewallIppoolDynamicMappingRead,
		Update: resourceObjectFirewallIppoolDynamicMappingUpdate,
		Delete: resourceObjectFirewallIppoolDynamicMappingDelete,

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
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_scope": &schema.Schema{
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
			"add_nat64_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"block_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_block_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cgn_client_endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_client_ipv6shift": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_client_startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_fixedalloc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgn_overload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgn_port_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cgn_port_start": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cgn_spa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_prefix_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"endport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exclude_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"icmp_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_blocks_per_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pba_interim_log": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pba_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"permit_any_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_per_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"privileged_port_use_pba": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_endip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"startport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"subnet_broadcast_in_ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"utilization_alarm_clear": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"utilization_alarm_raise": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectFirewallIppoolDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	ippool := d.Get("ippool").(string)
	paradict["ippool"] = ippool

	obj, err := getObjectObjectFirewallIppoolDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIppoolDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallIppoolDynamicMapping(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIppoolDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectFirewallIppoolDynamicMappingRead(d, m)
}

func resourceObjectFirewallIppoolDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	ippool := d.Get("ippool").(string)
	paradict["ippool"] = ippool

	obj, err := getObjectObjectFirewallIppoolDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIppoolDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallIppoolDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIppoolDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectFirewallIppoolDynamicMappingRead(d, m)
}

func resourceObjectFirewallIppoolDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	ippool := d.Get("ippool").(string)
	paradict["ippool"] = ippool

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallIppoolDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallIppoolDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallIppoolDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	ippool := d.Get("ippool").(string)
	if ippool == "" {
		ippool = importOptionChecking(m.(*FortiClient).Cfg, "ippool")
		if ippool == "" {
			return fmt.Errorf("Parameter ippool is missing")
		}
		if err = d.Set("ippool", ippool); err != nil {
			return fmt.Errorf("Error set params ippool: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["ippool"] = ippool

	o, err := c.ReadObjectFirewallIppoolDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIppoolDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallIppoolDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIppoolDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallIppoolDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallIppoolDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallIppoolDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallIppoolDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallIppoolDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingAddNat64Route2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingArpIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingArpReply2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingAssociatedInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingBlockSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnBlockSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnClientEndip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnClientStartip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnFixedalloc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnOverload2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnPortEnd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnPortStart2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnSpa2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingClientPrefixLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingComments2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingEndip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingEndport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingExcludeIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallIppoolDynamicMappingIcmpSessionQuota2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingNat642edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingNumBlocksPerUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPbaInterimLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPbaTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPermitAnyHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPortPerUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPrivilegedPortUsePba2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSourceEndip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSourcePrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSourceStartip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingStartip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingStartport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingTcpSessionQuota2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingUdpSessionQuota2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmClear2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallIppoolDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectFirewallIppoolDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectFirewallIppoolDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectFirewallIppoolDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectFirewallIppoolDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("add_nat64_route", flattenObjectFirewallIppoolDynamicMappingAddNat64Route2edl(o["add-nat64-route"], d, "add_nat64_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat64-route"], "ObjectFirewallIppoolDynamicMapping-AddNat64Route"); ok {
			if err = d.Set("add_nat64_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat64_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat64_route: %v", err)
		}
	}

	if err = d.Set("arp_intf", flattenObjectFirewallIppoolDynamicMappingArpIntf2edl(o["arp-intf"], d, "arp_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-intf"], "ObjectFirewallIppoolDynamicMapping-ArpIntf"); ok {
			if err = d.Set("arp_intf", vv); err != nil {
				return fmt.Errorf("Error reading arp_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_intf: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenObjectFirewallIppoolDynamicMappingArpReply2edl(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "ObjectFirewallIppoolDynamicMapping-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenObjectFirewallIppoolDynamicMappingAssociatedInterface2edl(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "ObjectFirewallIppoolDynamicMapping-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("block_size", flattenObjectFirewallIppoolDynamicMappingBlockSize2edl(o["block-size"], d, "block_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-size"], "ObjectFirewallIppoolDynamicMapping-BlockSize"); ok {
			if err = d.Set("block_size", vv); err != nil {
				return fmt.Errorf("Error reading block_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_size: %v", err)
		}
	}

	if err = d.Set("cgn_block_size", flattenObjectFirewallIppoolDynamicMappingCgnBlockSize2edl(o["cgn-block-size"], d, "cgn_block_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-block-size"], "ObjectFirewallIppoolDynamicMapping-CgnBlockSize"); ok {
			if err = d.Set("cgn_block_size", vv); err != nil {
				return fmt.Errorf("Error reading cgn_block_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_block_size: %v", err)
		}
	}

	if err = d.Set("cgn_client_endip", flattenObjectFirewallIppoolDynamicMappingCgnClientEndip2edl(o["cgn-client-endip"], d, "cgn_client_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-endip"], "ObjectFirewallIppoolDynamicMapping-CgnClientEndip"); ok {
			if err = d.Set("cgn_client_endip", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_endip: %v", err)
		}
	}

	if err = d.Set("cgn_client_ipv6shift", flattenObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift2edl(o["cgn-client-ipv6shift"], d, "cgn_client_ipv6shift")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-ipv6shift"], "ObjectFirewallIppoolDynamicMapping-CgnClientIpv6Shift"); ok {
			if err = d.Set("cgn_client_ipv6shift", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_ipv6shift: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_ipv6shift: %v", err)
		}
	}

	if err = d.Set("cgn_client_startip", flattenObjectFirewallIppoolDynamicMappingCgnClientStartip2edl(o["cgn-client-startip"], d, "cgn_client_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-startip"], "ObjectFirewallIppoolDynamicMapping-CgnClientStartip"); ok {
			if err = d.Set("cgn_client_startip", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_startip: %v", err)
		}
	}

	if err = d.Set("cgn_fixedalloc", flattenObjectFirewallIppoolDynamicMappingCgnFixedalloc2edl(o["cgn-fixedalloc"], d, "cgn_fixedalloc")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-fixedalloc"], "ObjectFirewallIppoolDynamicMapping-CgnFixedalloc"); ok {
			if err = d.Set("cgn_fixedalloc", vv); err != nil {
				return fmt.Errorf("Error reading cgn_fixedalloc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_fixedalloc: %v", err)
		}
	}

	if err = d.Set("cgn_overload", flattenObjectFirewallIppoolDynamicMappingCgnOverload2edl(o["cgn-overload"], d, "cgn_overload")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-overload"], "ObjectFirewallIppoolDynamicMapping-CgnOverload"); ok {
			if err = d.Set("cgn_overload", vv); err != nil {
				return fmt.Errorf("Error reading cgn_overload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_overload: %v", err)
		}
	}

	if err = d.Set("cgn_port_end", flattenObjectFirewallIppoolDynamicMappingCgnPortEnd2edl(o["cgn-port-end"], d, "cgn_port_end")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-port-end"], "ObjectFirewallIppoolDynamicMapping-CgnPortEnd"); ok {
			if err = d.Set("cgn_port_end", vv); err != nil {
				return fmt.Errorf("Error reading cgn_port_end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_port_end: %v", err)
		}
	}

	if err = d.Set("cgn_port_start", flattenObjectFirewallIppoolDynamicMappingCgnPortStart2edl(o["cgn-port-start"], d, "cgn_port_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-port-start"], "ObjectFirewallIppoolDynamicMapping-CgnPortStart"); ok {
			if err = d.Set("cgn_port_start", vv); err != nil {
				return fmt.Errorf("Error reading cgn_port_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_port_start: %v", err)
		}
	}

	if err = d.Set("cgn_spa", flattenObjectFirewallIppoolDynamicMappingCgnSpa2edl(o["cgn-spa"], d, "cgn_spa")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-spa"], "ObjectFirewallIppoolDynamicMapping-CgnSpa"); ok {
			if err = d.Set("cgn_spa", vv); err != nil {
				return fmt.Errorf("Error reading cgn_spa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_spa: %v", err)
		}
	}

	if err = d.Set("client_prefix_length", flattenObjectFirewallIppoolDynamicMappingClientPrefixLength2edl(o["client-prefix-length"], d, "client_prefix_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-prefix-length"], "ObjectFirewallIppoolDynamicMapping-ClientPrefixLength"); ok {
			if err = d.Set("client_prefix_length", vv); err != nil {
				return fmt.Errorf("Error reading client_prefix_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_prefix_length: %v", err)
		}
	}

	if err = d.Set("comments", flattenObjectFirewallIppoolDynamicMappingComments2edl(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectFirewallIppoolDynamicMapping-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("endip", flattenObjectFirewallIppoolDynamicMappingEndip2edl(o["endip"], d, "endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["endip"], "ObjectFirewallIppoolDynamicMapping-Endip"); ok {
			if err = d.Set("endip", vv); err != nil {
				return fmt.Errorf("Error reading endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("endport", flattenObjectFirewallIppoolDynamicMappingEndport2edl(o["endport"], d, "endport")); err != nil {
		if vv, ok := fortiAPIPatch(o["endport"], "ObjectFirewallIppoolDynamicMapping-Endport"); ok {
			if err = d.Set("endport", vv); err != nil {
				return fmt.Errorf("Error reading endport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endport: %v", err)
		}
	}

	if err = d.Set("exclude_ip", flattenObjectFirewallIppoolDynamicMappingExcludeIp2edl(o["exclude-ip"], d, "exclude_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-ip"], "ObjectFirewallIppoolDynamicMapping-ExcludeIp"); ok {
			if err = d.Set("exclude_ip", vv); err != nil {
				return fmt.Errorf("Error reading exclude_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_ip: %v", err)
		}
	}

	if err = d.Set("icmp_session_quota", flattenObjectFirewallIppoolDynamicMappingIcmpSessionQuota2edl(o["icmp-session-quota"], d, "icmp_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-session-quota"], "ObjectFirewallIppoolDynamicMapping-IcmpSessionQuota"); ok {
			if err = d.Set("icmp_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading icmp_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_session_quota: %v", err)
		}
	}

	if err = d.Set("nat64", flattenObjectFirewallIppoolDynamicMappingNat642edl(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "ObjectFirewallIppoolDynamicMapping-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("num_blocks_per_user", flattenObjectFirewallIppoolDynamicMappingNumBlocksPerUser2edl(o["num-blocks-per-user"], d, "num_blocks_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["num-blocks-per-user"], "ObjectFirewallIppoolDynamicMapping-NumBlocksPerUser"); ok {
			if err = d.Set("num_blocks_per_user", vv); err != nil {
				return fmt.Errorf("Error reading num_blocks_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading num_blocks_per_user: %v", err)
		}
	}

	if err = d.Set("pba_interim_log", flattenObjectFirewallIppoolDynamicMappingPbaInterimLog2edl(o["pba-interim-log"], d, "pba_interim_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-interim-log"], "ObjectFirewallIppoolDynamicMapping-PbaInterimLog"); ok {
			if err = d.Set("pba_interim_log", vv); err != nil {
				return fmt.Errorf("Error reading pba_interim_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_interim_log: %v", err)
		}
	}

	if err = d.Set("pba_timeout", flattenObjectFirewallIppoolDynamicMappingPbaTimeout2edl(o["pba-timeout"], d, "pba_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-timeout"], "ObjectFirewallIppoolDynamicMapping-PbaTimeout"); ok {
			if err = d.Set("pba_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pba_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_timeout: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenObjectFirewallIppoolDynamicMappingPermitAnyHost2edl(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-any-host"], "ObjectFirewallIppoolDynamicMapping-PermitAnyHost"); ok {
			if err = d.Set("permit_any_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_any_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("port_per_user", flattenObjectFirewallIppoolDynamicMappingPortPerUser2edl(o["port-per-user"], d, "port_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-per-user"], "ObjectFirewallIppoolDynamicMapping-PortPerUser"); ok {
			if err = d.Set("port_per_user", vv); err != nil {
				return fmt.Errorf("Error reading port_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_per_user: %v", err)
		}
	}

	if err = d.Set("privileged_port_use_pba", flattenObjectFirewallIppoolDynamicMappingPrivilegedPortUsePba2edl(o["privileged-port-use-pba"], d, "privileged_port_use_pba")); err != nil {
		if vv, ok := fortiAPIPatch(o["privileged-port-use-pba"], "ObjectFirewallIppoolDynamicMapping-PrivilegedPortUsePba"); ok {
			if err = d.Set("privileged_port_use_pba", vv); err != nil {
				return fmt.Errorf("Error reading privileged_port_use_pba: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading privileged_port_use_pba: %v", err)
		}
	}

	if err = d.Set("source_endip", flattenObjectFirewallIppoolDynamicMappingSourceEndip2edl(o["source-endip"], d, "source_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-endip"], "ObjectFirewallIppoolDynamicMapping-SourceEndip"); ok {
			if err = d.Set("source_endip", vv); err != nil {
				return fmt.Errorf("Error reading source_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_endip: %v", err)
		}
	}

	if err = d.Set("source_prefix6", flattenObjectFirewallIppoolDynamicMappingSourcePrefix62edl(o["source-prefix6"], d, "source_prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-prefix6"], "ObjectFirewallIppoolDynamicMapping-SourcePrefix6"); ok {
			if err = d.Set("source_prefix6", vv); err != nil {
				return fmt.Errorf("Error reading source_prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_prefix6: %v", err)
		}
	}

	if err = d.Set("source_startip", flattenObjectFirewallIppoolDynamicMappingSourceStartip2edl(o["source-startip"], d, "source_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-startip"], "ObjectFirewallIppoolDynamicMapping-SourceStartip"); ok {
			if err = d.Set("source_startip", vv); err != nil {
				return fmt.Errorf("Error reading source_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_startip: %v", err)
		}
	}

	if err = d.Set("startip", flattenObjectFirewallIppoolDynamicMappingStartip2edl(o["startip"], d, "startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["startip"], "ObjectFirewallIppoolDynamicMapping-Startip"); ok {
			if err = d.Set("startip", vv); err != nil {
				return fmt.Errorf("Error reading startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("startport", flattenObjectFirewallIppoolDynamicMappingStartport2edl(o["startport"], d, "startport")); err != nil {
		if vv, ok := fortiAPIPatch(o["startport"], "ObjectFirewallIppoolDynamicMapping-Startport"); ok {
			if err = d.Set("startport", vv); err != nil {
				return fmt.Errorf("Error reading startport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startport: %v", err)
		}
	}

	if err = d.Set("subnet_broadcast_in_ippool", flattenObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool2edl(o["subnet-broadcast-in-ippool"], d, "subnet_broadcast_in_ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-broadcast-in-ippool"], "ObjectFirewallIppoolDynamicMapping-SubnetBroadcastInIppool"); ok {
			if err = d.Set("subnet_broadcast_in_ippool", vv); err != nil {
				return fmt.Errorf("Error reading subnet_broadcast_in_ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_broadcast_in_ippool: %v", err)
		}
	}

	if err = d.Set("tcp_session_quota", flattenObjectFirewallIppoolDynamicMappingTcpSessionQuota2edl(o["tcp-session-quota"], d, "tcp_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-session-quota"], "ObjectFirewallIppoolDynamicMapping-TcpSessionQuota"); ok {
			if err = d.Set("tcp_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading tcp_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_session_quota: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallIppoolDynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallIppoolDynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("udp_session_quota", flattenObjectFirewallIppoolDynamicMappingUdpSessionQuota2edl(o["udp-session-quota"], d, "udp_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-session-quota"], "ObjectFirewallIppoolDynamicMapping-UdpSessionQuota"); ok {
			if err = d.Set("udp_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading udp_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_session_quota: %v", err)
		}
	}

	if err = d.Set("utilization_alarm_clear", flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmClear2edl(o["utilization-alarm-clear"], d, "utilization_alarm_clear")); err != nil {
		if vv, ok := fortiAPIPatch(o["utilization-alarm-clear"], "ObjectFirewallIppoolDynamicMapping-UtilizationAlarmClear"); ok {
			if err = d.Set("utilization_alarm_clear", vv); err != nil {
				return fmt.Errorf("Error reading utilization_alarm_clear: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utilization_alarm_clear: %v", err)
		}
	}

	if err = d.Set("utilization_alarm_raise", flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise2edl(o["utilization-alarm-raise"], d, "utilization_alarm_raise")); err != nil {
		if vv, ok := fortiAPIPatch(o["utilization-alarm-raise"], "ObjectFirewallIppoolDynamicMapping-UtilizationAlarmRaise"); ok {
			if err = d.Set("utilization_alarm_raise", vv); err != nil {
				return fmt.Errorf("Error reading utilization_alarm_raise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utilization_alarm_raise: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallIppoolDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallIppoolDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallIppoolDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallIppoolDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallIppoolDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingAddNat64Route2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingArpIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingArpReply2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingAssociatedInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingBlockSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnBlockSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnClientEndip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnClientStartip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnFixedalloc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnOverload2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnPortEnd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnPortStart2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnSpa2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingClientPrefixLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingComments2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingEndip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingEndport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingExcludeIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallIppoolDynamicMappingIcmpSessionQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingNat642edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingNumBlocksPerUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPbaInterimLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPbaTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPermitAnyHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPortPerUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPrivilegedPortUsePba2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSourceEndip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSourcePrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSourceStartip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingStartip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingStartport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingTcpSessionQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingUdpSessionQuota2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingUtilizationAlarmClear2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallIppoolDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectFirewallIppoolDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("add_nat64_route"); ok || d.HasChange("add_nat64_route") {
		t, err := expandObjectFirewallIppoolDynamicMappingAddNat64Route2edl(d, v, "add_nat64_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat64-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_intf"); ok || d.HasChange("arp_intf") {
		t, err := expandObjectFirewallIppoolDynamicMappingArpIntf2edl(d, v, "arp_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-intf"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandObjectFirewallIppoolDynamicMappingArpReply2edl(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok || d.HasChange("associated_interface") {
		t, err := expandObjectFirewallIppoolDynamicMappingAssociatedInterface2edl(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("block_size"); ok || d.HasChange("block_size") {
		t, err := expandObjectFirewallIppoolDynamicMappingBlockSize2edl(d, v, "block_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-size"] = t
		}
	}

	if v, ok := d.GetOk("cgn_block_size"); ok || d.HasChange("cgn_block_size") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnBlockSize2edl(d, v, "cgn_block_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-block-size"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_endip"); ok || d.HasChange("cgn_client_endip") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnClientEndip2edl(d, v, "cgn_client_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-endip"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_ipv6shift"); ok || d.HasChange("cgn_client_ipv6shift") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift2edl(d, v, "cgn_client_ipv6shift")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-ipv6shift"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_startip"); ok || d.HasChange("cgn_client_startip") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnClientStartip2edl(d, v, "cgn_client_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-startip"] = t
		}
	}

	if v, ok := d.GetOk("cgn_fixedalloc"); ok || d.HasChange("cgn_fixedalloc") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnFixedalloc2edl(d, v, "cgn_fixedalloc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-fixedalloc"] = t
		}
	}

	if v, ok := d.GetOk("cgn_overload"); ok || d.HasChange("cgn_overload") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnOverload2edl(d, v, "cgn_overload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-overload"] = t
		}
	}

	if v, ok := d.GetOk("cgn_port_end"); ok || d.HasChange("cgn_port_end") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnPortEnd2edl(d, v, "cgn_port_end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-port-end"] = t
		}
	}

	if v, ok := d.GetOk("cgn_port_start"); ok || d.HasChange("cgn_port_start") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnPortStart2edl(d, v, "cgn_port_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-port-start"] = t
		}
	}

	if v, ok := d.GetOk("cgn_spa"); ok || d.HasChange("cgn_spa") {
		t, err := expandObjectFirewallIppoolDynamicMappingCgnSpa2edl(d, v, "cgn_spa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-spa"] = t
		}
	}

	if v, ok := d.GetOk("client_prefix_length"); ok || d.HasChange("client_prefix_length") {
		t, err := expandObjectFirewallIppoolDynamicMappingClientPrefixLength2edl(d, v, "client_prefix_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-prefix-length"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectFirewallIppoolDynamicMappingComments2edl(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok || d.HasChange("endip") {
		t, err := expandObjectFirewallIppoolDynamicMappingEndip2edl(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("endport"); ok || d.HasChange("endport") {
		t, err := expandObjectFirewallIppoolDynamicMappingEndport2edl(d, v, "endport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endport"] = t
		}
	}

	if v, ok := d.GetOk("exclude_ip"); ok || d.HasChange("exclude_ip") {
		t, err := expandObjectFirewallIppoolDynamicMappingExcludeIp2edl(d, v, "exclude_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-ip"] = t
		}
	}

	if v, ok := d.GetOk("icmp_session_quota"); ok || d.HasChange("icmp_session_quota") {
		t, err := expandObjectFirewallIppoolDynamicMappingIcmpSessionQuota2edl(d, v, "icmp_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandObjectFirewallIppoolDynamicMappingNat642edl(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("num_blocks_per_user"); ok || d.HasChange("num_blocks_per_user") {
		t, err := expandObjectFirewallIppoolDynamicMappingNumBlocksPerUser2edl(d, v, "num_blocks_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["num-blocks-per-user"] = t
		}
	}

	if v, ok := d.GetOk("pba_interim_log"); ok || d.HasChange("pba_interim_log") {
		t, err := expandObjectFirewallIppoolDynamicMappingPbaInterimLog2edl(d, v, "pba_interim_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-interim-log"] = t
		}
	}

	if v, ok := d.GetOk("pba_timeout"); ok || d.HasChange("pba_timeout") {
		t, err := expandObjectFirewallIppoolDynamicMappingPbaTimeout2edl(d, v, "pba_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-timeout"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok || d.HasChange("permit_any_host") {
		t, err := expandObjectFirewallIppoolDynamicMappingPermitAnyHost2edl(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("port_per_user"); ok || d.HasChange("port_per_user") {
		t, err := expandObjectFirewallIppoolDynamicMappingPortPerUser2edl(d, v, "port_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-per-user"] = t
		}
	}

	if v, ok := d.GetOk("privileged_port_use_pba"); ok || d.HasChange("privileged_port_use_pba") {
		t, err := expandObjectFirewallIppoolDynamicMappingPrivilegedPortUsePba2edl(d, v, "privileged_port_use_pba")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["privileged-port-use-pba"] = t
		}
	}

	if v, ok := d.GetOk("source_endip"); ok || d.HasChange("source_endip") {
		t, err := expandObjectFirewallIppoolDynamicMappingSourceEndip2edl(d, v, "source_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-endip"] = t
		}
	}

	if v, ok := d.GetOk("source_prefix6"); ok || d.HasChange("source_prefix6") {
		t, err := expandObjectFirewallIppoolDynamicMappingSourcePrefix62edl(d, v, "source_prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-prefix6"] = t
		}
	}

	if v, ok := d.GetOk("source_startip"); ok || d.HasChange("source_startip") {
		t, err := expandObjectFirewallIppoolDynamicMappingSourceStartip2edl(d, v, "source_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-startip"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok || d.HasChange("startip") {
		t, err := expandObjectFirewallIppoolDynamicMappingStartip2edl(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("startport"); ok || d.HasChange("startport") {
		t, err := expandObjectFirewallIppoolDynamicMappingStartport2edl(d, v, "startport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startport"] = t
		}
	}

	if v, ok := d.GetOk("subnet_broadcast_in_ippool"); ok || d.HasChange("subnet_broadcast_in_ippool") {
		t, err := expandObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool2edl(d, v, "subnet_broadcast_in_ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-broadcast-in-ippool"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_quota"); ok || d.HasChange("tcp_session_quota") {
		t, err := expandObjectFirewallIppoolDynamicMappingTcpSessionQuota2edl(d, v, "tcp_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallIppoolDynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("udp_session_quota"); ok || d.HasChange("udp_session_quota") {
		t, err := expandObjectFirewallIppoolDynamicMappingUdpSessionQuota2edl(d, v, "udp_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("utilization_alarm_clear"); ok || d.HasChange("utilization_alarm_clear") {
		t, err := expandObjectFirewallIppoolDynamicMappingUtilizationAlarmClear2edl(d, v, "utilization_alarm_clear")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilization-alarm-clear"] = t
		}
	}

	if v, ok := d.GetOk("utilization_alarm_raise"); ok || d.HasChange("utilization_alarm_raise") {
		t, err := expandObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise2edl(d, v, "utilization_alarm_raise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilization-alarm-raise"] = t
		}
	}

	return &obj, nil
}
