// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic VirtualWanLinkMembers

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectDynamicVirtualWanLinkMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicVirtualWanLinkMembersCreate,
		Read:   resourceObjectDynamicVirtualWanLinkMembersRead,
		Update: resourceObjectDynamicVirtualWanLinkMembersUpdate,
		Delete: resourceObjectDynamicVirtualWanLinkMembersDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detect_failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detect_http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detect_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detect_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_recoverytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"detect_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
										Computed: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"detect_failtime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"detect_http_get": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detect_http_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detect_http_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"detect_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"detect_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detect_recoverytime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"detect_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detect_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ingress_spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"volume_ratio": &schema.Schema{
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
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"volume_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
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

func resourceObjectDynamicVirtualWanLinkMembersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicVirtualWanLinkMembers(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicVirtualWanLinkMembers resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDynamicVirtualWanLinkMembers(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicVirtualWanLinkMembers resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicVirtualWanLinkMembersRead(d, m)
}

func resourceObjectDynamicVirtualWanLinkMembersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicVirtualWanLinkMembers(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicVirtualWanLinkMembers resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDynamicVirtualWanLinkMembers(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicVirtualWanLinkMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicVirtualWanLinkMembersRead(d, m)
}

func resourceObjectDynamicVirtualWanLinkMembersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDynamicVirtualWanLinkMembers(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicVirtualWanLinkMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicVirtualWanLinkMembersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDynamicVirtualWanLinkMembers(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicVirtualWanLinkMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicVirtualWanLinkMembers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicVirtualWanLinkMembers resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicVirtualWanLinkMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "ping",
			2: "tcp-echo",
			4: "udp-echo",
			8: "http",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDetectTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_failtime"
		if _, ok := i["detect-failtime"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectFailtime(i["detect-failtime"], d, pre_append)
			tmp["detect_failtime"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectFailtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_http_get"
		if _, ok := i["detect-http-get"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpGet(i["detect-http-get"], d, pre_append)
			tmp["detect_http_get"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectHttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_http_match"
		if _, ok := i["detect-http-match"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpMatch(i["detect-http-match"], d, pre_append)
			tmp["detect_http_match"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectHttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_http_port"
		if _, ok := i["detect-http-port"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpPort(i["detect-http-port"], d, pre_append)
			tmp["detect_http_port"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectHttpPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_interval"
		if _, ok := i["detect-interval"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectInterval(i["detect-interval"], d, pre_append)
			tmp["detect_interval"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_protocol"
		if _, ok := i["detect-protocol"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectProtocol(i["detect-protocol"], d, pre_append)
			tmp["detect_protocol"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectProtocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_recoverytime"
		if _, ok := i["detect-recoverytime"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectRecoverytime(i["detect-recoverytime"], d, pre_append)
			tmp["detect_recoverytime"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectRecoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_server"
		if _, ok := i["detect-server"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectServer(i["detect-server"], d, pre_append)
			tmp["detect_server"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_timeout"
		if _, ok := i["detect-timeout"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectTimeout(i["detect-timeout"], d, pre_append)
			tmp["detect_timeout"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-DetectTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingGateway6(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingIngressSpilloverThreshold(i["ingress-spillover-threshold"], d, pre_append)
			tmp["ingress_spillover_threshold"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-IngressSpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingSpilloverThreshold(i["spillover-threshold"], d, pre_append)
			tmp["spillover_threshold"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-SpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingVolumeRatio(i["volume-ratio"], d, pre_append)
			tmp["volume_ratio"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-VolumeRatio")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembers-DynamicMapping-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembersDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectDynamicVirtualWanLinkMembersDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkMembersDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "ping",
			2: "tcp-echo",
			4: "udp-echo",
			8: "http",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingDetectTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersDynamicMappingWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicVirtualWanLinkMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenObjectDynamicVirtualWanLinkMembersComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDynamicVirtualWanLinkMembers-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", flattenObjectDynamicVirtualWanLinkMembersCost(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "ObjectDynamicVirtualWanLinkMembers-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("detect_failtime", flattenObjectDynamicVirtualWanLinkMembersDetectFailtime(o["detect-failtime"], d, "detect_failtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-failtime"], "ObjectDynamicVirtualWanLinkMembers-DetectFailtime"); ok {
			if err = d.Set("detect_failtime", vv); err != nil {
				return fmt.Errorf("Error reading detect_failtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_failtime: %v", err)
		}
	}

	if err = d.Set("detect_http_get", flattenObjectDynamicVirtualWanLinkMembersDetectHttpGet(o["detect-http-get"], d, "detect_http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-http-get"], "ObjectDynamicVirtualWanLinkMembers-DetectHttpGet"); ok {
			if err = d.Set("detect_http_get", vv); err != nil {
				return fmt.Errorf("Error reading detect_http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_http_get: %v", err)
		}
	}

	if err = d.Set("detect_http_match", flattenObjectDynamicVirtualWanLinkMembersDetectHttpMatch(o["detect-http-match"], d, "detect_http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-http-match"], "ObjectDynamicVirtualWanLinkMembers-DetectHttpMatch"); ok {
			if err = d.Set("detect_http_match", vv); err != nil {
				return fmt.Errorf("Error reading detect_http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_http_match: %v", err)
		}
	}

	if err = d.Set("detect_http_port", flattenObjectDynamicVirtualWanLinkMembersDetectHttpPort(o["detect-http-port"], d, "detect_http_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-http-port"], "ObjectDynamicVirtualWanLinkMembers-DetectHttpPort"); ok {
			if err = d.Set("detect_http_port", vv); err != nil {
				return fmt.Errorf("Error reading detect_http_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_http_port: %v", err)
		}
	}

	if err = d.Set("detect_interval", flattenObjectDynamicVirtualWanLinkMembersDetectInterval(o["detect-interval"], d, "detect_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-interval"], "ObjectDynamicVirtualWanLinkMembers-DetectInterval"); ok {
			if err = d.Set("detect_interval", vv); err != nil {
				return fmt.Errorf("Error reading detect_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_interval: %v", err)
		}
	}

	if err = d.Set("detect_protocol", flattenObjectDynamicVirtualWanLinkMembersDetectProtocol(o["detect-protocol"], d, "detect_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-protocol"], "ObjectDynamicVirtualWanLinkMembers-DetectProtocol"); ok {
			if err = d.Set("detect_protocol", vv); err != nil {
				return fmt.Errorf("Error reading detect_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_protocol: %v", err)
		}
	}

	if err = d.Set("detect_recoverytime", flattenObjectDynamicVirtualWanLinkMembersDetectRecoverytime(o["detect-recoverytime"], d, "detect_recoverytime")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-recoverytime"], "ObjectDynamicVirtualWanLinkMembers-DetectRecoverytime"); ok {
			if err = d.Set("detect_recoverytime", vv); err != nil {
				return fmt.Errorf("Error reading detect_recoverytime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_recoverytime: %v", err)
		}
	}

	if err = d.Set("detect_server", flattenObjectDynamicVirtualWanLinkMembersDetectServer(o["detect-server"], d, "detect_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-server"], "ObjectDynamicVirtualWanLinkMembers-DetectServer"); ok {
			if err = d.Set("detect_server", vv); err != nil {
				return fmt.Errorf("Error reading detect_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_server: %v", err)
		}
	}

	if err = d.Set("detect_timeout", flattenObjectDynamicVirtualWanLinkMembersDetectTimeout(o["detect-timeout"], d, "detect_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["detect-timeout"], "ObjectDynamicVirtualWanLinkMembers-DetectTimeout"); ok {
			if err = d.Set("detect_timeout", vv); err != nil {
				return fmt.Errorf("Error reading detect_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detect_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectDynamicVirtualWanLinkMembersDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicVirtualWanLinkMembers-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectDynamicVirtualWanLinkMembersDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicVirtualWanLinkMembers-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("gateway", flattenObjectDynamicVirtualWanLinkMembersGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "ObjectDynamicVirtualWanLinkMembers-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenObjectDynamicVirtualWanLinkMembersGateway6(o["gateway6"], d, "gateway6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway6"], "ObjectDynamicVirtualWanLinkMembers-Gateway6"); ok {
			if err = d.Set("gateway6", vv); err != nil {
				return fmt.Errorf("Error reading gateway6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenObjectDynamicVirtualWanLinkMembersIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-spillover-threshold"], "ObjectDynamicVirtualWanLinkMembers-IngressSpilloverThreshold"); ok {
			if err = d.Set("ingress_spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectDynamicVirtualWanLinkMembersInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectDynamicVirtualWanLinkMembers-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDynamicVirtualWanLinkMembersName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDynamicVirtualWanLinkMembers-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectDynamicVirtualWanLinkMembersPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectDynamicVirtualWanLinkMembers-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("source", flattenObjectDynamicVirtualWanLinkMembersSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "ObjectDynamicVirtualWanLinkMembers-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenObjectDynamicVirtualWanLinkMembersSource6(o["source6"], d, "source6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source6"], "ObjectDynamicVirtualWanLinkMembers-Source6"); ok {
			if err = d.Set("source6", vv); err != nil {
				return fmt.Errorf("Error reading source6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenObjectDynamicVirtualWanLinkMembersSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spillover-threshold"], "ObjectDynamicVirtualWanLinkMembers-SpilloverThreshold"); ok {
			if err = d.Set("spillover_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spillover_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectDynamicVirtualWanLinkMembersStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectDynamicVirtualWanLinkMembers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("volume_ratio", flattenObjectDynamicVirtualWanLinkMembersVolumeRatio(o["volume-ratio"], d, "volume_ratio")); err != nil {
		if vv, ok := fortiAPIPatch(o["volume-ratio"], "ObjectDynamicVirtualWanLinkMembers-VolumeRatio"); ok {
			if err = d.Set("volume_ratio", vv); err != nil {
				return fmt.Errorf("Error reading volume_ratio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectDynamicVirtualWanLinkMembersWeight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectDynamicVirtualWanLinkMembers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicVirtualWanLinkMembersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicVirtualWanLinkMembersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDetectTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["_scope"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_failtime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-failtime"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectFailtime(d, i["detect_failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_http_get"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-http-get"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpGet(d, i["detect_http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_http_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-http-match"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpMatch(d, i["detect_http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_http_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-http-port"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpPort(d, i["detect_http_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-interval"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectInterval(d, i["detect_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-protocol"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectProtocol(d, i["detect_protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_recoverytime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-recoverytime"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectRecoverytime(d, i["detect_recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-server"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectServer(d, i["detect_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detect-timeout"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectTimeout(d, i["detect_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway6"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingGateway6(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ingress-spillover-threshold"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source6"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["spillover-threshold"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingSpilloverThreshold(d, i["spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["volume-ratio"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingVolumeRatio(d, i["volume_ratio"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandObjectDynamicVirtualWanLinkMembersDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingDetectTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersDynamicMappingWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkMembersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicVirtualWanLinkMembers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersCost(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("detect_failtime"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectFailtime(d, v, "detect_failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-failtime"] = t
		}
	}

	if v, ok := d.GetOk("detect_http_get"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectHttpGet(d, v, "detect_http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-http-get"] = t
		}
	}

	if v, ok := d.GetOk("detect_http_match"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectHttpMatch(d, v, "detect_http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-http-match"] = t
		}
	}

	if v, ok := d.GetOk("detect_http_port"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectHttpPort(d, v, "detect_http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-http-port"] = t
		}
	}

	if v, ok := d.GetOk("detect_interval"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectInterval(d, v, "detect_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-interval"] = t
		}
	}

	if v, ok := d.GetOk("detect_protocol"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectProtocol(d, v, "detect_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-protocol"] = t
		}
	}

	if v, ok := d.GetOk("detect_recoverytime"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectRecoverytime(d, v, "detect_recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("detect_server"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectServer(d, v, "detect_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-server"] = t
		}
	}

	if v, ok := d.GetOk("detect_timeout"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDetectTimeout(d, v, "detect_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detect-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersGateway6(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersIngressSpilloverThreshold(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersSource6(d, v, "source6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersSpilloverThreshold(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("volume_ratio"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersVolumeRatio(d, v, "volume_ratio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandObjectDynamicVirtualWanLinkMembersWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
