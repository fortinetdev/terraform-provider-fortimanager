// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 firewall addresses.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallAddress6() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAddress6Create,
		Read:   resourceObjectFirewallAddress6Read,
		Update: resourceObjectFirewallAddress6Update,
		Delete: resourceObjectFirewallAddress6Delete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_image_base64": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"cache_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"country": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fabric_object": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"global_object": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"macaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet_segment": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
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
						"tags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"visibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"net_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"obj_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
			},
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"visibility": &schema.Schema{
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

func resourceObjectFirewallAddress6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallAddress6(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallAddress6(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddress6Read(d, m)
}

func resourceObjectFirewallAddress6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallAddress6(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallAddress6(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddress6Read(d, m)
}

func resourceObjectFirewallAddress6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallAddress6(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAddress6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallAddress6(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAddress6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAddress6ImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6CacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Country(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := i["_image-base64"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingImageBase64(i["_image-base64"], d, pre_append)
			tmp["_image_base64"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-ImageBase64")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := i["cache-ttl"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingCacheTtl(i["cache-ttl"], d, pre_append)
			tmp["cache_ttl"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-CacheTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := i["color"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingColor(i["color"], d, pre_append)
			tmp["color"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Color")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := i["country"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingCountry(i["country"], d, pre_append)
			tmp["country"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Country")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := i["end-mac"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingEndMac(i["end-mac"], d, pre_append)
			tmp["end_mac"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-EndMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := i["global-object"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingGlobalObject(i["global-object"], d, pre_append)
			tmp["global_object"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-GlobalObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := i["host-type"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingHostType(i["host-type"], d, pre_append)
			tmp["host_type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-HostType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingIp6(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := i["macaddr"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingMacaddr(i["macaddr"], d, pre_append)
			tmp["macaddr"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Macaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-ObjId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := i["sdn"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingSdn(i["sdn"], d, pre_append)
			tmp["sdn"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Sdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-StartIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := i["start-mac"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingStartMac(i["start-mac"], d, pre_append)
			tmp["start_mac"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-StartMac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_segment"
		if _, ok := i["subnet-segment"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingSubnetSegment(i["subnet-segment"], d, pre_append)
			tmp["subnet_segment"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-SubnetSegment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Tags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := i["template"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingTemplate(i["template"], d, pre_append)
			tmp["template"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Template")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := i["visibility"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingVisibility(i["visibility"], d, pre_append)
			tmp["visibility"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-DynamicMapping-Visibility")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6DynamicMappingImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddress6DynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6DynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6DynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6DynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingEndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingHostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddress6DynamicMappingObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingStartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingSubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddress6DynamicMappingSubnetSegmentName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6DynamicMapping-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingSubnetSegmentType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6DynamicMapping-SubnetSegment-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFirewallAddress6DynamicMappingSubnetSegmentValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6DynamicMapping-SubnetSegment-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6DynamicMappingSubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingSubnetSegmentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingSubnetSegmentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6DynamicMappingVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6EndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6FabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Fqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6GlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Host(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6HostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Ip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6List(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddress6ListIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-List-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := i["net-id"]; ok {
			v := flattenObjectFirewallAddress6ListNetId(i["net-id"], d, pre_append)
			tmp["net_id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-List-NetId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := i["obj-id"]; ok {
			v := flattenObjectFirewallAddress6ListObjId(i["obj-id"], d, pre_append)
			tmp["obj_id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-List-ObjId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6ListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6ListNetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6ListObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Macaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6ObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Sdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6StartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6SubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddress6SubnetSegmentName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFirewallAddress6SubnetSegmentType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-SubnetSegment-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFirewallAddress6SubnetSegmentValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-SubnetSegment-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6SubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6SubnetSegmentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6SubnetSegmentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Tagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectFirewallAddress6TaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallAddress6TaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenObjectFirewallAddress6TaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6-Tagging-Tags")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6TaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallAddress6Template(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6Visibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAddress6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenObjectFirewallAddress6ImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "ObjectFirewallAddress6-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("cache_ttl", flattenObjectFirewallAddress6CacheTtl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-ttl"], "ObjectFirewallAddress6-CacheTtl"); ok {
			if err = d.Set("cache_ttl", vv); err != nil {
				return fmt.Errorf("Error reading cache_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallAddress6Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallAddress6-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallAddress6Comment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallAddress6-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectFirewallAddress6Country(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectFirewallAddress6-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectFirewallAddress6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallAddress6-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectFirewallAddress6DynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectFirewallAddress6-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("end_ip", flattenObjectFirewallAddress6EndIp(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectFirewallAddress6-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenObjectFirewallAddress6EndMac(o["end-mac"], d, "end_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-mac"], "ObjectFirewallAddress6-EndMac"); ok {
			if err = d.Set("end_mac", vv); err != nil {
				return fmt.Errorf("Error reading end_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectFirewallAddress6FabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectFirewallAddress6-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenObjectFirewallAddress6Fqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "ObjectFirewallAddress6-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("global_object", flattenObjectFirewallAddress6GlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "ObjectFirewallAddress6-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("host", flattenObjectFirewallAddress6Host(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectFirewallAddress6-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_type", flattenObjectFirewallAddress6HostType(o["host-type"], d, "host_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-type"], "ObjectFirewallAddress6-HostType"); ok {
			if err = d.Set("host_type", vv); err != nil {
				return fmt.Errorf("Error reading host_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("ip6", flattenObjectFirewallAddress6Ip6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "ObjectFirewallAddress6-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("list", flattenObjectFirewallAddress6List(o["list"], d, "list")); err != nil {
			if vv, ok := fortiAPIPatch(o["list"], "ObjectFirewallAddress6-List"); ok {
				if err = d.Set("list", vv); err != nil {
					return fmt.Errorf("Error reading list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list"); ok {
			if err = d.Set("list", flattenObjectFirewallAddress6List(o["list"], d, "list")); err != nil {
				if vv, ok := fortiAPIPatch(o["list"], "ObjectFirewallAddress6-List"); ok {
					if err = d.Set("list", vv); err != nil {
						return fmt.Errorf("Error reading list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading list: %v", err)
				}
			}
		}
	}

	if err = d.Set("macaddr", flattenObjectFirewallAddress6Macaddr(o["macaddr"], d, "macaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["macaddr"], "ObjectFirewallAddress6-Macaddr"); ok {
			if err = d.Set("macaddr", vv); err != nil {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallAddress6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAddress6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenObjectFirewallAddress6ObjId(o["obj-id"], d, "obj_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["obj-id"], "ObjectFirewallAddress6-ObjId"); ok {
			if err = d.Set("obj_id", vv); err != nil {
				return fmt.Errorf("Error reading obj_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("sdn", flattenObjectFirewallAddress6Sdn(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "ObjectFirewallAddress6-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectFirewallAddress6StartIp(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectFirewallAddress6-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenObjectFirewallAddress6StartMac(o["start-mac"], d, "start_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-mac"], "ObjectFirewallAddress6-StartMac"); ok {
			if err = d.Set("start_mac", vv); err != nil {
				return fmt.Errorf("Error reading start_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("subnet_segment", flattenObjectFirewallAddress6SubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["subnet-segment"], "ObjectFirewallAddress6-SubnetSegment"); ok {
				if err = d.Set("subnet_segment", vv); err != nil {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subnet_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subnet_segment"); ok {
			if err = d.Set("subnet_segment", flattenObjectFirewallAddress6SubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["subnet-segment"], "ObjectFirewallAddress6-SubnetSegment"); ok {
					if err = d.Set("subnet_segment", vv); err != nil {
						return fmt.Errorf("Error reading subnet_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenObjectFirewallAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "ObjectFirewallAddress6-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenObjectFirewallAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "ObjectFirewallAddress6-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("template", flattenObjectFirewallAddress6Template(o["template"], d, "template")); err != nil {
		if vv, ok := fortiAPIPatch(o["template"], "ObjectFirewallAddress6-Template"); ok {
			if err = d.Set("template", vv); err != nil {
				return fmt.Errorf("Error reading template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallAddress6Type(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallAddress6-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallAddress6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallAddress6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("visibility", flattenObjectFirewallAddress6Visibility(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "ObjectFirewallAddress6-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAddress6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAddress6ImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6CacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Country(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_image_base64"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_image-base64"], _ = expandObjectFirewallAddress6DynamicMappingImageBase64(d, i["_image_base64"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_scope"], _ = expandObjectFirewallAddress6DynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cache_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cache-ttl"], _ = expandObjectFirewallAddress6DynamicMappingCacheTtl(d, i["cache_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color"], _ = expandObjectFirewallAddress6DynamicMappingColor(d, i["color"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectFirewallAddress6DynamicMappingComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["country"], _ = expandObjectFirewallAddress6DynamicMappingCountry(d, i["country"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectFirewallAddress6DynamicMappingEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-mac"], _ = expandObjectFirewallAddress6DynamicMappingEndMac(d, i["end_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object"], _ = expandObjectFirewallAddress6DynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fqdn"], _ = expandObjectFirewallAddress6DynamicMappingFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "global_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["global-object"], _ = expandObjectFirewallAddress6DynamicMappingGlobalObject(d, i["global_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandObjectFirewallAddress6DynamicMappingHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host-type"], _ = expandObjectFirewallAddress6DynamicMappingHostType(d, i["host_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandObjectFirewallAddress6DynamicMappingIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["macaddr"], _ = expandObjectFirewallAddress6DynamicMappingMacaddr(d, i["macaddr"], pre_append)
		} else {
			tmp["macaddr"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-id"], _ = expandObjectFirewallAddress6DynamicMappingObjId(d, i["obj_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sdn"], _ = expandObjectFirewallAddress6DynamicMappingSdn(d, i["sdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectFirewallAddress6DynamicMappingStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-mac"], _ = expandObjectFirewallAddress6DynamicMappingStartMac(d, i["start_mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet_segment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subnet-segment"], _ = expandObjectFirewallAddress6DynamicMappingSubnetSegment(d, i["subnet_segment"], pre_append)
		} else {
			tmp["subnet-segment"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandObjectFirewallAddress6DynamicMappingTags(d, i["tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["template"], _ = expandObjectFirewallAddress6DynamicMappingTemplate(d, i["template"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallAddress6DynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandObjectFirewallAddress6DynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visibility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["visibility"], _ = expandObjectFirewallAddress6DynamicMappingVisibility(d, i["visibility"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6DynamicMappingImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallAddress6DynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectFirewallAddress6DynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6DynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingEndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingHostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingMacaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallAddress6DynamicMappingObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingStartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingSubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallAddress6DynamicMappingSubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallAddress6DynamicMappingSubnetSegmentType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFirewallAddress6DynamicMappingSubnetSegmentValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6DynamicMappingSubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingSubnetSegmentType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingSubnetSegmentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6DynamicMappingVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6EndMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6FabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Fqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6GlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Host(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6HostType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Ip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6List(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectFirewallAddress6ListIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["net-id"], _ = expandObjectFirewallAddress6ListNetId(d, i["net_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "obj_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["obj-id"], _ = expandObjectFirewallAddress6ListObjId(d, i["obj_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6ListIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6ListNetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6ListObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Macaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallAddress6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6ObjId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Sdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6StartMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6SubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallAddress6SubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFirewallAddress6SubnetSegmentType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFirewallAddress6SubnetSegmentValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6SubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6SubnetSegmentType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6SubnetSegmentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Tagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectFirewallAddress6TaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallAddress6TaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandObjectFirewallAddress6TaggingTags(d, i["tags"], pre_append)
		} else {
			tmp["tags"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6TaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallAddress6Template(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6Visibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAddress6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandObjectFirewallAddress6ImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("cache_ttl"); ok || d.HasChange("cache_ttl") {
		t, err := expandObjectFirewallAddress6CacheTtl(d, v, "cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallAddress6Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallAddress6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandObjectFirewallAddress6Country(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectFirewallAddress6DynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectFirewallAddress6EndIp(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok || d.HasChange("end_mac") {
		t, err := expandObjectFirewallAddress6EndMac(d, v, "end_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectFirewallAddress6FabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandObjectFirewallAddress6Fqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandObjectFirewallAddress6GlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectFirewallAddress6Host(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_type"); ok || d.HasChange("host_type") {
		t, err := expandObjectFirewallAddress6HostType(d, v, "host_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-type"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandObjectFirewallAddress6Ip6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("list"); ok || d.HasChange("list") {
		t, err := expandObjectFirewallAddress6List(d, v, "list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandObjectFirewallAddress6Macaddr(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallAddress6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok || d.HasChange("obj_id") {
		t, err := expandObjectFirewallAddress6ObjId(d, v, "obj_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandObjectFirewallAddress6Sdn(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectFirewallAddress6StartIp(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok || d.HasChange("start_mac") {
		t, err := expandObjectFirewallAddress6StartMac(d, v, "start_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment"); ok || d.HasChange("subnet_segment") {
		t, err := expandObjectFirewallAddress6SubnetSegment(d, v, "subnet_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandObjectFirewallAddress6Tagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("template"); ok || d.HasChange("template") {
		t, err := expandObjectFirewallAddress6Template(d, v, "template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallAddress6Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallAddress6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok || d.HasChange("visibility") {
		t, err := expandObjectFirewallAddress6Visibility(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	return &obj, nil
}
