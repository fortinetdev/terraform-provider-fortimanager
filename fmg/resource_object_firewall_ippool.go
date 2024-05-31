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

func resourceObjectFirewallIppool() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallIppoolCreate,
		Read:   resourceObjectFirewallIppoolRead,
		Update: resourceObjectFirewallIppoolUpdate,
		Delete: resourceObjectFirewallIppoolDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"source_endip": &schema.Schema{
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
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
					},
				},
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
			"source_endip": &schema.Schema{
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceObjectFirewallIppoolCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectFirewallIppool(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIppool resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallIppool(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIppool resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallIppoolRead(d, m)
}

func resourceObjectFirewallIppoolUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallIppool(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIppool resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallIppool(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIppool resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallIppoolRead(d, m)
}

func resourceObjectFirewallIppoolDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallIppool(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallIppool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallIppoolRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallIppool(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIppool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallIppool(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIppool resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallIppoolAddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolArpIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallIppoolArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallIppoolBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnClientEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnClientIpv6Shift(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnClientStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnFixedalloc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnOverload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnPortEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnPortStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolCgnSpa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := i["add-nat64-route"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingAddNat64Route(i["add-nat64-route"], d, pre_append)
			tmp["add_nat64_route"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-AddNat64Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_intf"
		if _, ok := i["arp-intf"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingArpIntf(i["arp-intf"], d, pre_append)
			tmp["arp_intf"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-ArpIntf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := i["arp-reply"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingArpReply(i["arp-reply"], d, pre_append)
			tmp["arp_reply"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-ArpReply")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := i["associated-interface"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingAssociatedInterface(i["associated-interface"], d, pre_append)
			tmp["associated_interface"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-AssociatedInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "block_size"
		if _, ok := i["block-size"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingBlockSize(i["block-size"], d, pre_append)
			tmp["block_size"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-BlockSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_block_size"
		if _, ok := i["cgn-block-size"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnBlockSize(i["cgn-block-size"], d, pre_append)
			tmp["cgn_block_size"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnBlockSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_endip"
		if _, ok := i["cgn-client-endip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnClientEndip(i["cgn-client-endip"], d, pre_append)
			tmp["cgn_client_endip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnClientEndip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_ipv6shift"
		if _, ok := i["cgn-client-ipv6shift"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift(i["cgn-client-ipv6shift"], d, pre_append)
			tmp["cgn_client_ipv6shift"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnClientIpv6Shift")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_startip"
		if _, ok := i["cgn-client-startip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnClientStartip(i["cgn-client-startip"], d, pre_append)
			tmp["cgn_client_startip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnClientStartip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_fixedalloc"
		if _, ok := i["cgn-fixedalloc"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnFixedalloc(i["cgn-fixedalloc"], d, pre_append)
			tmp["cgn_fixedalloc"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnFixedalloc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_overload"
		if _, ok := i["cgn-overload"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnOverload(i["cgn-overload"], d, pre_append)
			tmp["cgn_overload"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnOverload")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_end"
		if _, ok := i["cgn-port-end"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnPortEnd(i["cgn-port-end"], d, pre_append)
			tmp["cgn_port_end"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnPortEnd")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_start"
		if _, ok := i["cgn-port-start"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnPortStart(i["cgn-port-start"], d, pre_append)
			tmp["cgn_port_start"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnPortStart")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_spa"
		if _, ok := i["cgn-spa"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingCgnSpa(i["cgn-spa"], d, pre_append)
			tmp["cgn_spa"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-CgnSpa")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := i["endip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingEndip(i["endip"], d, pre_append)
			tmp["endip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Endip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endport"
		if _, ok := i["endport"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingEndport(i["endport"], d, pre_append)
			tmp["endport"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Endport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_ip"
		if _, ok := i["exclude-ip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingExcludeIp(i["exclude-ip"], d, pre_append)
			tmp["exclude_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-ExcludeIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := i["nat64"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingNat64(i["nat64"], d, pre_append)
			tmp["nat64"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Nat64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "num_blocks_per_user"
		if _, ok := i["num-blocks-per-user"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingNumBlocksPerUser(i["num-blocks-per-user"], d, pre_append)
			tmp["num_blocks_per_user"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-NumBlocksPerUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_interim_log"
		if _, ok := i["pba-interim-log"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingPbaInterimLog(i["pba-interim-log"], d, pre_append)
			tmp["pba_interim_log"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-PbaInterimLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_timeout"
		if _, ok := i["pba-timeout"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingPbaTimeout(i["pba-timeout"], d, pre_append)
			tmp["pba_timeout"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-PbaTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_any_host"
		if _, ok := i["permit-any-host"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingPermitAnyHost(i["permit-any-host"], d, pre_append)
			tmp["permit_any_host"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-PermitAnyHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_per_user"
		if _, ok := i["port-per-user"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingPortPerUser(i["port-per-user"], d, pre_append)
			tmp["port_per_user"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-PortPerUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_endip"
		if _, ok := i["source-endip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingSourceEndip(i["source-endip"], d, pre_append)
			tmp["source_endip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-SourceEndip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_startip"
		if _, ok := i["source-startip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingSourceStartip(i["source-startip"], d, pre_append)
			tmp["source_startip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-SourceStartip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := i["startip"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingStartip(i["startip"], d, pre_append)
			tmp["startip"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Startip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startport"
		if _, ok := i["startport"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingStartport(i["startport"], d, pre_append)
			tmp["startport"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Startport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_broadcast_in_ippool"
		if _, ok := i["subnet-broadcast-in-ippool"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool(i["subnet-broadcast-in-ippool"], d, pre_append)
			tmp["subnet_broadcast_in_ippool"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-SubnetBroadcastInIppool")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_clear"
		if _, ok := i["utilization-alarm-clear"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmClear(i["utilization-alarm-clear"], d, pre_append)
			tmp["utilization_alarm_clear"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-UtilizationAlarmClear")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_raise"
		if _, ok := i["utilization-alarm-raise"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise(i["utilization-alarm-raise"], d, pre_append)
			tmp["utilization_alarm_raise"] = fortiAPISubPartPatch(v, "ObjectFirewallIppool-DynamicMapping-UtilizationAlarmRaise")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallIppoolDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallIppoolDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallIppoolDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallIppoolDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallIppoolDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallIppoolDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingAddNat64Route(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingArpIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingAssociatedInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnBlockSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnClientEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnClientStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnFixedalloc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnOverload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnPortEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnPortStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingCgnSpa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingEndport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingExcludeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallIppoolDynamicMappingNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingNumBlocksPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPbaInterimLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPbaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingPortPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSourceEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSourceStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingStartport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmClear(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolEndport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolExcludeIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallIppoolName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolNumBlocksPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolPbaInterimLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolPbaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolPortPerUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolSourceEndip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolSourceStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolStartip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolStartport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolSubnetBroadcastInIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolUtilizationAlarmClear(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIppoolUtilizationAlarmRaise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallIppool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("add_nat64_route", flattenObjectFirewallIppoolAddNat64Route(o["add-nat64-route"], d, "add_nat64_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["add-nat64-route"], "ObjectFirewallIppool-AddNat64Route"); ok {
			if err = d.Set("add_nat64_route", vv); err != nil {
				return fmt.Errorf("Error reading add_nat64_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_nat64_route: %v", err)
		}
	}

	if err = d.Set("arp_intf", flattenObjectFirewallIppoolArpIntf(o["arp-intf"], d, "arp_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-intf"], "ObjectFirewallIppool-ArpIntf"); ok {
			if err = d.Set("arp_intf", vv); err != nil {
				return fmt.Errorf("Error reading arp_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_intf: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenObjectFirewallIppoolArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if vv, ok := fortiAPIPatch(o["arp-reply"], "ObjectFirewallIppool-ArpReply"); ok {
			if err = d.Set("arp_reply", vv); err != nil {
				return fmt.Errorf("Error reading arp_reply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenObjectFirewallIppoolAssociatedInterface(o["associated-interface"], d, "associated_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["associated-interface"], "ObjectFirewallIppool-AssociatedInterface"); ok {
			if err = d.Set("associated_interface", vv); err != nil {
				return fmt.Errorf("Error reading associated_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("block_size", flattenObjectFirewallIppoolBlockSize(o["block-size"], d, "block_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-size"], "ObjectFirewallIppool-BlockSize"); ok {
			if err = d.Set("block_size", vv); err != nil {
				return fmt.Errorf("Error reading block_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_size: %v", err)
		}
	}

	if err = d.Set("cgn_block_size", flattenObjectFirewallIppoolCgnBlockSize(o["cgn-block-size"], d, "cgn_block_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-block-size"], "ObjectFirewallIppool-CgnBlockSize"); ok {
			if err = d.Set("cgn_block_size", vv); err != nil {
				return fmt.Errorf("Error reading cgn_block_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_block_size: %v", err)
		}
	}

	if err = d.Set("cgn_client_endip", flattenObjectFirewallIppoolCgnClientEndip(o["cgn-client-endip"], d, "cgn_client_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-endip"], "ObjectFirewallIppool-CgnClientEndip"); ok {
			if err = d.Set("cgn_client_endip", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_endip: %v", err)
		}
	}

	if err = d.Set("cgn_client_ipv6shift", flattenObjectFirewallIppoolCgnClientIpv6Shift(o["cgn-client-ipv6shift"], d, "cgn_client_ipv6shift")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-ipv6shift"], "ObjectFirewallIppool-CgnClientIpv6Shift"); ok {
			if err = d.Set("cgn_client_ipv6shift", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_ipv6shift: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_ipv6shift: %v", err)
		}
	}

	if err = d.Set("cgn_client_startip", flattenObjectFirewallIppoolCgnClientStartip(o["cgn-client-startip"], d, "cgn_client_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-client-startip"], "ObjectFirewallIppool-CgnClientStartip"); ok {
			if err = d.Set("cgn_client_startip", vv); err != nil {
				return fmt.Errorf("Error reading cgn_client_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_client_startip: %v", err)
		}
	}

	if err = d.Set("cgn_fixedalloc", flattenObjectFirewallIppoolCgnFixedalloc(o["cgn-fixedalloc"], d, "cgn_fixedalloc")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-fixedalloc"], "ObjectFirewallIppool-CgnFixedalloc"); ok {
			if err = d.Set("cgn_fixedalloc", vv); err != nil {
				return fmt.Errorf("Error reading cgn_fixedalloc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_fixedalloc: %v", err)
		}
	}

	if err = d.Set("cgn_overload", flattenObjectFirewallIppoolCgnOverload(o["cgn-overload"], d, "cgn_overload")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-overload"], "ObjectFirewallIppool-CgnOverload"); ok {
			if err = d.Set("cgn_overload", vv); err != nil {
				return fmt.Errorf("Error reading cgn_overload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_overload: %v", err)
		}
	}

	if err = d.Set("cgn_port_end", flattenObjectFirewallIppoolCgnPortEnd(o["cgn-port-end"], d, "cgn_port_end")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-port-end"], "ObjectFirewallIppool-CgnPortEnd"); ok {
			if err = d.Set("cgn_port_end", vv); err != nil {
				return fmt.Errorf("Error reading cgn_port_end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_port_end: %v", err)
		}
	}

	if err = d.Set("cgn_port_start", flattenObjectFirewallIppoolCgnPortStart(o["cgn-port-start"], d, "cgn_port_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-port-start"], "ObjectFirewallIppool-CgnPortStart"); ok {
			if err = d.Set("cgn_port_start", vv); err != nil {
				return fmt.Errorf("Error reading cgn_port_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_port_start: %v", err)
		}
	}

	if err = d.Set("cgn_spa", flattenObjectFirewallIppoolCgnSpa(o["cgn-spa"], d, "cgn_spa")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-spa"], "ObjectFirewallIppool-CgnSpa"); ok {
			if err = d.Set("cgn_spa", vv); err != nil {
				return fmt.Errorf("Error reading cgn_spa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_spa: %v", err)
		}
	}

	if err = d.Set("comments", flattenObjectFirewallIppoolComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectFirewallIppool-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallIppoolDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallIppool-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallIppoolDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallIppool-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("endip", flattenObjectFirewallIppoolEndip(o["endip"], d, "endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["endip"], "ObjectFirewallIppool-Endip"); ok {
			if err = d.Set("endip", vv); err != nil {
				return fmt.Errorf("Error reading endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("endport", flattenObjectFirewallIppoolEndport(o["endport"], d, "endport")); err != nil {
		if vv, ok := fortiAPIPatch(o["endport"], "ObjectFirewallIppool-Endport"); ok {
			if err = d.Set("endport", vv); err != nil {
				return fmt.Errorf("Error reading endport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endport: %v", err)
		}
	}

	if err = d.Set("exclude_ip", flattenObjectFirewallIppoolExcludeIp(o["exclude-ip"], d, "exclude_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-ip"], "ObjectFirewallIppool-ExcludeIp"); ok {
			if err = d.Set("exclude_ip", vv); err != nil {
				return fmt.Errorf("Error reading exclude_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_ip: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallIppoolName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallIppool-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat64", flattenObjectFirewallIppoolNat64(o["nat64"], d, "nat64")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat64"], "ObjectFirewallIppool-Nat64"); ok {
			if err = d.Set("nat64", vv); err != nil {
				return fmt.Errorf("Error reading nat64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("num_blocks_per_user", flattenObjectFirewallIppoolNumBlocksPerUser(o["num-blocks-per-user"], d, "num_blocks_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["num-blocks-per-user"], "ObjectFirewallIppool-NumBlocksPerUser"); ok {
			if err = d.Set("num_blocks_per_user", vv); err != nil {
				return fmt.Errorf("Error reading num_blocks_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading num_blocks_per_user: %v", err)
		}
	}

	if err = d.Set("pba_interim_log", flattenObjectFirewallIppoolPbaInterimLog(o["pba-interim-log"], d, "pba_interim_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-interim-log"], "ObjectFirewallIppool-PbaInterimLog"); ok {
			if err = d.Set("pba_interim_log", vv); err != nil {
				return fmt.Errorf("Error reading pba_interim_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_interim_log: %v", err)
		}
	}

	if err = d.Set("pba_timeout", flattenObjectFirewallIppoolPbaTimeout(o["pba-timeout"], d, "pba_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pba-timeout"], "ObjectFirewallIppool-PbaTimeout"); ok {
			if err = d.Set("pba_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pba_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pba_timeout: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenObjectFirewallIppoolPermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-any-host"], "ObjectFirewallIppool-PermitAnyHost"); ok {
			if err = d.Set("permit_any_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_any_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("port_per_user", flattenObjectFirewallIppoolPortPerUser(o["port-per-user"], d, "port_per_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["port-per-user"], "ObjectFirewallIppool-PortPerUser"); ok {
			if err = d.Set("port_per_user", vv); err != nil {
				return fmt.Errorf("Error reading port_per_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port_per_user: %v", err)
		}
	}

	if err = d.Set("source_endip", flattenObjectFirewallIppoolSourceEndip(o["source-endip"], d, "source_endip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-endip"], "ObjectFirewallIppool-SourceEndip"); ok {
			if err = d.Set("source_endip", vv); err != nil {
				return fmt.Errorf("Error reading source_endip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_endip: %v", err)
		}
	}

	if err = d.Set("source_startip", flattenObjectFirewallIppoolSourceStartip(o["source-startip"], d, "source_startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-startip"], "ObjectFirewallIppool-SourceStartip"); ok {
			if err = d.Set("source_startip", vv); err != nil {
				return fmt.Errorf("Error reading source_startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_startip: %v", err)
		}
	}

	if err = d.Set("startip", flattenObjectFirewallIppoolStartip(o["startip"], d, "startip")); err != nil {
		if vv, ok := fortiAPIPatch(o["startip"], "ObjectFirewallIppool-Startip"); ok {
			if err = d.Set("startip", vv); err != nil {
				return fmt.Errorf("Error reading startip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("startport", flattenObjectFirewallIppoolStartport(o["startport"], d, "startport")); err != nil {
		if vv, ok := fortiAPIPatch(o["startport"], "ObjectFirewallIppool-Startport"); ok {
			if err = d.Set("startport", vv); err != nil {
				return fmt.Errorf("Error reading startport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading startport: %v", err)
		}
	}

	if err = d.Set("subnet_broadcast_in_ippool", flattenObjectFirewallIppoolSubnetBroadcastInIppool(o["subnet-broadcast-in-ippool"], d, "subnet_broadcast_in_ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-broadcast-in-ippool"], "ObjectFirewallIppool-SubnetBroadcastInIppool"); ok {
			if err = d.Set("subnet_broadcast_in_ippool", vv); err != nil {
				return fmt.Errorf("Error reading subnet_broadcast_in_ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_broadcast_in_ippool: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallIppoolType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallIppool-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("utilization_alarm_clear", flattenObjectFirewallIppoolUtilizationAlarmClear(o["utilization-alarm-clear"], d, "utilization_alarm_clear")); err != nil {
		if vv, ok := fortiAPIPatch(o["utilization-alarm-clear"], "ObjectFirewallIppool-UtilizationAlarmClear"); ok {
			if err = d.Set("utilization_alarm_clear", vv); err != nil {
				return fmt.Errorf("Error reading utilization_alarm_clear: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utilization_alarm_clear: %v", err)
		}
	}

	if err = d.Set("utilization_alarm_raise", flattenObjectFirewallIppoolUtilizationAlarmRaise(o["utilization-alarm-raise"], d, "utilization_alarm_raise")); err != nil {
		if vv, ok := fortiAPIPatch(o["utilization-alarm-raise"], "ObjectFirewallIppool-UtilizationAlarmRaise"); ok {
			if err = d.Set("utilization_alarm_raise", vv); err != nil {
				return fmt.Errorf("Error reading utilization_alarm_raise: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utilization_alarm_raise: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallIppoolFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallIppoolAddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolArpIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallIppoolArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallIppoolBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnClientEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnClientIpv6Shift(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnClientStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnFixedalloc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnOverload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnPortEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnPortStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolCgnSpa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallIppoolDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_nat64_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["add-nat64-route"], _ = expandObjectFirewallIppoolDynamicMappingAddNat64Route(d, i["add_nat64_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-intf"], _ = expandObjectFirewallIppoolDynamicMappingArpIntf(d, i["arp_intf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["arp-reply"], _ = expandObjectFirewallIppoolDynamicMappingArpReply(d, i["arp_reply"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "associated_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["associated-interface"], _ = expandObjectFirewallIppoolDynamicMappingAssociatedInterface(d, i["associated_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "block_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["block-size"], _ = expandObjectFirewallIppoolDynamicMappingBlockSize(d, i["block_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_block_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-block-size"], _ = expandObjectFirewallIppoolDynamicMappingCgnBlockSize(d, i["cgn_block_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-client-endip"], _ = expandObjectFirewallIppoolDynamicMappingCgnClientEndip(d, i["cgn_client_endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_ipv6shift"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-client-ipv6shift"], _ = expandObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift(d, i["cgn_client_ipv6shift"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_client_startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-client-startip"], _ = expandObjectFirewallIppoolDynamicMappingCgnClientStartip(d, i["cgn_client_startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_fixedalloc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-fixedalloc"], _ = expandObjectFirewallIppoolDynamicMappingCgnFixedalloc(d, i["cgn_fixedalloc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_overload"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-overload"], _ = expandObjectFirewallIppoolDynamicMappingCgnOverload(d, i["cgn_overload"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_end"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-port-end"], _ = expandObjectFirewallIppoolDynamicMappingCgnPortEnd(d, i["cgn_port_end"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_port_start"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-port-start"], _ = expandObjectFirewallIppoolDynamicMappingCgnPortStart(d, i["cgn_port_start"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cgn_spa"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cgn-spa"], _ = expandObjectFirewallIppoolDynamicMappingCgnSpa(d, i["cgn_spa"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comments"], _ = expandObjectFirewallIppoolDynamicMappingComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endip"], _ = expandObjectFirewallIppoolDynamicMappingEndip(d, i["endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endport"], _ = expandObjectFirewallIppoolDynamicMappingEndport(d, i["endport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclude-ip"], _ = expandObjectFirewallIppoolDynamicMappingExcludeIp(d, i["exclude_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nat64"], _ = expandObjectFirewallIppoolDynamicMappingNat64(d, i["nat64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "num_blocks_per_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["num-blocks-per-user"], _ = expandObjectFirewallIppoolDynamicMappingNumBlocksPerUser(d, i["num_blocks_per_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_interim_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pba-interim-log"], _ = expandObjectFirewallIppoolDynamicMappingPbaInterimLog(d, i["pba_interim_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pba_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pba-timeout"], _ = expandObjectFirewallIppoolDynamicMappingPbaTimeout(d, i["pba_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "permit_any_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["permit-any-host"], _ = expandObjectFirewallIppoolDynamicMappingPermitAnyHost(d, i["permit_any_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_per_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port-per-user"], _ = expandObjectFirewallIppoolDynamicMappingPortPerUser(d, i["port_per_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_endip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-endip"], _ = expandObjectFirewallIppoolDynamicMappingSourceEndip(d, i["source_endip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-startip"], _ = expandObjectFirewallIppoolDynamicMappingSourceStartip(d, i["source_startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["startip"], _ = expandObjectFirewallIppoolDynamicMappingStartip(d, i["startip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "startport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["startport"], _ = expandObjectFirewallIppoolDynamicMappingStartport(d, i["startport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_broadcast_in_ippool"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet-broadcast-in-ippool"], _ = expandObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool(d, i["subnet_broadcast_in_ippool"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallIppoolDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_clear"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utilization-alarm-clear"], _ = expandObjectFirewallIppoolDynamicMappingUtilizationAlarmClear(d, i["utilization_alarm_clear"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "utilization_alarm_raise"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["utilization-alarm-raise"], _ = expandObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise(d, i["utilization_alarm_raise"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallIppoolDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallIppoolDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallIppoolDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallIppoolDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingAddNat64Route(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingArpIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingAssociatedInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnBlockSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnClientEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnClientIpv6Shift(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnClientStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnFixedalloc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnOverload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnPortEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnPortStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingCgnSpa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingEndport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingExcludeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallIppoolDynamicMappingNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingNumBlocksPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPbaInterimLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPbaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingPortPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSourceEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSourceStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingStartport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingSubnetBroadcastInIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingUtilizationAlarmClear(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolDynamicMappingUtilizationAlarmRaise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolEndport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolExcludeIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallIppoolName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolNat64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolNumBlocksPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolPbaInterimLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolPbaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolPortPerUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolSourceEndip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolSourceStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolStartip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolStartport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolSubnetBroadcastInIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolUtilizationAlarmClear(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIppoolUtilizationAlarmRaise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallIppool(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_nat64_route"); ok || d.HasChange("add_nat64_route") {
		t, err := expandObjectFirewallIppoolAddNat64Route(d, v, "add_nat64_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat64-route"] = t
		}
	}

	if v, ok := d.GetOk("arp_intf"); ok || d.HasChange("arp_intf") {
		t, err := expandObjectFirewallIppoolArpIntf(d, v, "arp_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-intf"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok || d.HasChange("arp_reply") {
		t, err := expandObjectFirewallIppoolArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok || d.HasChange("associated_interface") {
		t, err := expandObjectFirewallIppoolAssociatedInterface(d, v, "associated_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOk("block_size"); ok || d.HasChange("block_size") {
		t, err := expandObjectFirewallIppoolBlockSize(d, v, "block_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-size"] = t
		}
	}

	if v, ok := d.GetOk("cgn_block_size"); ok || d.HasChange("cgn_block_size") {
		t, err := expandObjectFirewallIppoolCgnBlockSize(d, v, "cgn_block_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-block-size"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_endip"); ok || d.HasChange("cgn_client_endip") {
		t, err := expandObjectFirewallIppoolCgnClientEndip(d, v, "cgn_client_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-endip"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_ipv6shift"); ok || d.HasChange("cgn_client_ipv6shift") {
		t, err := expandObjectFirewallIppoolCgnClientIpv6Shift(d, v, "cgn_client_ipv6shift")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-ipv6shift"] = t
		}
	}

	if v, ok := d.GetOk("cgn_client_startip"); ok || d.HasChange("cgn_client_startip") {
		t, err := expandObjectFirewallIppoolCgnClientStartip(d, v, "cgn_client_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-client-startip"] = t
		}
	}

	if v, ok := d.GetOk("cgn_fixedalloc"); ok || d.HasChange("cgn_fixedalloc") {
		t, err := expandObjectFirewallIppoolCgnFixedalloc(d, v, "cgn_fixedalloc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-fixedalloc"] = t
		}
	}

	if v, ok := d.GetOk("cgn_overload"); ok || d.HasChange("cgn_overload") {
		t, err := expandObjectFirewallIppoolCgnOverload(d, v, "cgn_overload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-overload"] = t
		}
	}

	if v, ok := d.GetOk("cgn_port_end"); ok || d.HasChange("cgn_port_end") {
		t, err := expandObjectFirewallIppoolCgnPortEnd(d, v, "cgn_port_end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-port-end"] = t
		}
	}

	if v, ok := d.GetOk("cgn_port_start"); ok || d.HasChange("cgn_port_start") {
		t, err := expandObjectFirewallIppoolCgnPortStart(d, v, "cgn_port_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-port-start"] = t
		}
	}

	if v, ok := d.GetOk("cgn_spa"); ok || d.HasChange("cgn_spa") {
		t, err := expandObjectFirewallIppoolCgnSpa(d, v, "cgn_spa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-spa"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectFirewallIppoolComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFirewallIppoolDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok || d.HasChange("endip") {
		t, err := expandObjectFirewallIppoolEndip(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("endport"); ok || d.HasChange("endport") {
		t, err := expandObjectFirewallIppoolEndport(d, v, "endport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endport"] = t
		}
	}

	if v, ok := d.GetOk("exclude_ip"); ok || d.HasChange("exclude_ip") {
		t, err := expandObjectFirewallIppoolExcludeIp(d, v, "exclude_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallIppoolName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok || d.HasChange("nat64") {
		t, err := expandObjectFirewallIppoolNat64(d, v, "nat64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("num_blocks_per_user"); ok || d.HasChange("num_blocks_per_user") {
		t, err := expandObjectFirewallIppoolNumBlocksPerUser(d, v, "num_blocks_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["num-blocks-per-user"] = t
		}
	}

	if v, ok := d.GetOk("pba_interim_log"); ok || d.HasChange("pba_interim_log") {
		t, err := expandObjectFirewallIppoolPbaInterimLog(d, v, "pba_interim_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-interim-log"] = t
		}
	}

	if v, ok := d.GetOk("pba_timeout"); ok || d.HasChange("pba_timeout") {
		t, err := expandObjectFirewallIppoolPbaTimeout(d, v, "pba_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pba-timeout"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok || d.HasChange("permit_any_host") {
		t, err := expandObjectFirewallIppoolPermitAnyHost(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("port_per_user"); ok || d.HasChange("port_per_user") {
		t, err := expandObjectFirewallIppoolPortPerUser(d, v, "port_per_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-per-user"] = t
		}
	}

	if v, ok := d.GetOk("source_endip"); ok || d.HasChange("source_endip") {
		t, err := expandObjectFirewallIppoolSourceEndip(d, v, "source_endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-endip"] = t
		}
	}

	if v, ok := d.GetOk("source_startip"); ok || d.HasChange("source_startip") {
		t, err := expandObjectFirewallIppoolSourceStartip(d, v, "source_startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-startip"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok || d.HasChange("startip") {
		t, err := expandObjectFirewallIppoolStartip(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("startport"); ok || d.HasChange("startport") {
		t, err := expandObjectFirewallIppoolStartport(d, v, "startport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startport"] = t
		}
	}

	if v, ok := d.GetOk("subnet_broadcast_in_ippool"); ok || d.HasChange("subnet_broadcast_in_ippool") {
		t, err := expandObjectFirewallIppoolSubnetBroadcastInIppool(d, v, "subnet_broadcast_in_ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-broadcast-in-ippool"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallIppoolType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("utilization_alarm_clear"); ok || d.HasChange("utilization_alarm_clear") {
		t, err := expandObjectFirewallIppoolUtilizationAlarmClear(d, v, "utilization_alarm_clear")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilization-alarm-clear"] = t
		}
	}

	if v, ok := d.GetOk("utilization_alarm_raise"); ok || d.HasChange("utilization_alarm_raise") {
		t, err := expandObjectFirewallIppoolUtilizationAlarmRaise(d, v, "utilization_alarm_raise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilization-alarm-raise"] = t
		}
	}

	return &obj, nil
}
