// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic Interface

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicInterfaceCreate,
		Read:   resourceObjectDynamicInterfaceRead,
		Update: resourceObjectDynamicInterfaceUpdate,
		Delete: resourceObjectDynamicInterfaceDelete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_mapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"defmap_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"defmap_intrazone_deny": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"defmap_zonemember": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"description": &schema.Schema{
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
						"egress_shaping_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ingress_shaping_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"intrazone_deny": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"local_intf": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
			},
			"egress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"platform_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"egress_shaping_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ingress_shaping_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"intf_zone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"intrazone_deny": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"single_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_only": &schema.Schema{
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

func resourceObjectDynamicInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicInterface resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDynamicInterface(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicInterfaceRead(d, m)
}

func resourceObjectDynamicInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDynamicInterface(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicInterfaceRead(d, m)
}

func resourceObjectDynamicInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDynamicInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDynamicInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicInterface resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicInterfaceColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDefaultMapping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDefmapIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDefmapIntrazoneDeny(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDefmapZonemember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDynamicInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicInterfaceDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "egress_shaping_profile"
		if _, ok := i["egress-shaping-profile"]; ok {
			v := flattenObjectDynamicInterfaceDynamicMappingEgressShapingProfile(i["egress-shaping-profile"], d, pre_append)
			tmp["egress_shaping_profile"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-DynamicMapping-EgressShapingProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_shaping_profile"
		if _, ok := i["ingress-shaping-profile"]; ok {
			v := flattenObjectDynamicInterfaceDynamicMappingIngressShapingProfile(i["ingress-shaping-profile"], d, pre_append)
			tmp["ingress_shaping_profile"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-DynamicMapping-IngressShapingProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intrazone_deny"
		if _, ok := i["intrazone-deny"]; ok {
			v := flattenObjectDynamicInterfaceDynamicMappingIntrazoneDeny(i["intrazone-deny"], d, pre_append)
			tmp["intrazone_deny"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-DynamicMapping-IntrazoneDeny")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_intf"
		if _, ok := i["local-intf"]; ok {
			v := flattenObjectDynamicInterfaceDynamicMappingLocalIntf(i["local-intf"], d, pre_append)
			tmp["local_intf"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-DynamicMapping-LocalIntf")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicInterfaceDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicInterfaceDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicInterfaceDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectDynamicInterfaceDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectDynamicInterfaceDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicInterfaceDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingEgressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingIngressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingIntrazoneDeny(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingLocalIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDynamicInterfaceEgressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceIngressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "egress_shaping_profile"
		if _, ok := i["egress-shaping-profile"]; ok {
			v := flattenObjectDynamicInterfacePlatformMappingEgressShapingProfile(i["egress-shaping-profile"], d, pre_append)
			tmp["egress_shaping_profile"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-PlatformMapping-EgressShapingProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_shaping_profile"
		if _, ok := i["ingress-shaping-profile"]; ok {
			v := flattenObjectDynamicInterfacePlatformMappingIngressShapingProfile(i["ingress-shaping-profile"], d, pre_append)
			tmp["ingress_shaping_profile"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-PlatformMapping-IngressShapingProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intf_zone"
		if _, ok := i["intf-zone"]; ok {
			v := flattenObjectDynamicInterfacePlatformMappingIntfZone(i["intf-zone"], d, pre_append)
			tmp["intf_zone"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-PlatformMapping-IntfZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intrazone_deny"
		if _, ok := i["intrazone-deny"]; ok {
			v := flattenObjectDynamicInterfacePlatformMappingIntrazoneDeny(i["intrazone-deny"], d, pre_append)
			tmp["intrazone_deny"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-PlatformMapping-IntrazoneDeny")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectDynamicInterfacePlatformMappingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicInterface-PlatformMapping-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicInterfacePlatformMappingEgressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingIngressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingIntfZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingIntrazoneDeny(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfacePlatformMappingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceSingleIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceWildcardIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceZoneOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("color", flattenObjectDynamicInterfaceColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectDynamicInterface-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("default_mapping", flattenObjectDynamicInterfaceDefaultMapping(o["default-mapping"], d, "default_mapping")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-mapping"], "ObjectDynamicInterface-DefaultMapping"); ok {
			if err = d.Set("default_mapping", vv); err != nil {
				return fmt.Errorf("Error reading default_mapping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_mapping: %v", err)
		}
	}

	if err = d.Set("defmap_intf", flattenObjectDynamicInterfaceDefmapIntf(o["defmap-intf"], d, "defmap_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["defmap-intf"], "ObjectDynamicInterface-DefmapIntf"); ok {
			if err = d.Set("defmap_intf", vv); err != nil {
				return fmt.Errorf("Error reading defmap_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading defmap_intf: %v", err)
		}
	}

	if err = d.Set("defmap_intrazone_deny", flattenObjectDynamicInterfaceDefmapIntrazoneDeny(o["defmap-intrazone-deny"], d, "defmap_intrazone_deny")); err != nil {
		if vv, ok := fortiAPIPatch(o["defmap-intrazone-deny"], "ObjectDynamicInterface-DefmapIntrazoneDeny"); ok {
			if err = d.Set("defmap_intrazone_deny", vv); err != nil {
				return fmt.Errorf("Error reading defmap_intrazone_deny: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading defmap_intrazone_deny: %v", err)
		}
	}

	if err = d.Set("defmap_zonemember", flattenObjectDynamicInterfaceDefmapZonemember(o["defmap-zonemember"], d, "defmap_zonemember")); err != nil {
		if vv, ok := fortiAPIPatch(o["defmap-zonemember"], "ObjectDynamicInterface-DefmapZonemember"); ok {
			if err = d.Set("defmap_zonemember", vv); err != nil {
				return fmt.Errorf("Error reading defmap_zonemember: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading defmap_zonemember: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectDynamicInterfaceDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectDynamicInterface-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectDynamicInterfaceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicInterface-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectDynamicInterfaceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicInterface-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("egress_shaping_profile", flattenObjectDynamicInterfaceEgressShapingProfile(o["egress-shaping-profile"], d, "egress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-shaping-profile"], "ObjectDynamicInterface-EgressShapingProfile"); ok {
			if err = d.Set("egress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("ingress_shaping_profile", flattenObjectDynamicInterfaceIngressShapingProfile(o["ingress-shaping-profile"], d, "ingress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-shaping-profile"], "ObjectDynamicInterface-IngressShapingProfile"); ok {
			if err = d.Set("ingress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDynamicInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDynamicInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("platform_mapping", flattenObjectDynamicInterfacePlatformMapping(o["platform_mapping"], d, "platform_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["platform_mapping"], "ObjectDynamicInterface-PlatformMapping"); ok {
				if err = d.Set("platform_mapping", vv); err != nil {
					return fmt.Errorf("Error reading platform_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading platform_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("platform_mapping"); ok {
			if err = d.Set("platform_mapping", flattenObjectDynamicInterfacePlatformMapping(o["platform_mapping"], d, "platform_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["platform_mapping"], "ObjectDynamicInterface-PlatformMapping"); ok {
					if err = d.Set("platform_mapping", vv); err != nil {
						return fmt.Errorf("Error reading platform_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading platform_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("single_intf", flattenObjectDynamicInterfaceSingleIntf(o["single-intf"], d, "single_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-intf"], "ObjectDynamicInterface-SingleIntf"); ok {
			if err = d.Set("single_intf", vv); err != nil {
				return fmt.Errorf("Error reading single_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_intf: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenObjectDynamicInterfaceWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "ObjectDynamicInterface-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("wildcard_intf", flattenObjectDynamicInterfaceWildcardIntf(o["wildcard-intf"], d, "wildcard_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-intf"], "ObjectDynamicInterface-WildcardIntf"); ok {
			if err = d.Set("wildcard_intf", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_intf: %v", err)
		}
	}

	if err = d.Set("zone_only", flattenObjectDynamicInterfaceZoneOnly(o["zone-only"], d, "zone_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["zone-only"], "ObjectDynamicInterface-ZoneOnly"); ok {
			if err = d.Set("zone_only", vv); err != nil {
				return fmt.Errorf("Error reading zone_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zone_only: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicInterfaceColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDefaultMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDefmapIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDefmapIntrazoneDeny(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDefmapZonemember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDynamicInterfaceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["_scope"], _ = expandObjectDynamicInterfaceDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "egress_shaping_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["egress-shaping-profile"], _ = expandObjectDynamicInterfaceDynamicMappingEgressShapingProfile(d, i["egress_shaping_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_shaping_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-shaping-profile"], _ = expandObjectDynamicInterfaceDynamicMappingIngressShapingProfile(d, i["ingress_shaping_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intrazone_deny"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intrazone-deny"], _ = expandObjectDynamicInterfaceDynamicMappingIntrazoneDeny(d, i["intrazone_deny"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-intf"], _ = expandObjectDynamicInterfaceDynamicMappingLocalIntf(d, i["local_intf"], pre_append)
		} else {
			tmp["local-intf"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicInterfaceDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectDynamicInterfaceDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectDynamicInterfaceDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicInterfaceDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingEgressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingIngressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingIntrazoneDeny(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingLocalIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDynamicInterfaceEgressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceIngressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "egress_shaping_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["egress-shaping-profile"], _ = expandObjectDynamicInterfacePlatformMappingEgressShapingProfile(d, i["egress_shaping_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_shaping_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-shaping-profile"], _ = expandObjectDynamicInterfacePlatformMappingIngressShapingProfile(d, i["ingress_shaping_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intf_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intf-zone"], _ = expandObjectDynamicInterfacePlatformMappingIntfZone(d, i["intf_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intrazone_deny"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["intrazone-deny"], _ = expandObjectDynamicInterfacePlatformMappingIntrazoneDeny(d, i["intrazone_deny"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectDynamicInterfacePlatformMappingName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicInterfacePlatformMappingEgressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingIngressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingIntfZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingIntrazoneDeny(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfacePlatformMappingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceSingleIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceWildcardIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceZoneOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectDynamicInterfaceColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("default_mapping"); ok || d.HasChange("default_mapping") {
		t, err := expandObjectDynamicInterfaceDefaultMapping(d, v, "default_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-mapping"] = t
		}
	}

	if v, ok := d.GetOk("defmap_intf"); ok || d.HasChange("defmap_intf") {
		t, err := expandObjectDynamicInterfaceDefmapIntf(d, v, "defmap_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defmap-intf"] = t
		}
	}

	if v, ok := d.GetOk("defmap_intrazone_deny"); ok || d.HasChange("defmap_intrazone_deny") {
		t, err := expandObjectDynamicInterfaceDefmapIntrazoneDeny(d, v, "defmap_intrazone_deny")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defmap-intrazone-deny"] = t
		}
	}

	if v, ok := d.GetOk("defmap_zonemember"); ok || d.HasChange("defmap_zonemember") {
		t, err := expandObjectDynamicInterfaceDefmapZonemember(d, v, "defmap_zonemember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defmap-zonemember"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectDynamicInterfaceDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectDynamicInterfaceDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("egress_shaping_profile"); ok || d.HasChange("egress_shaping_profile") {
		t, err := expandObjectDynamicInterfaceEgressShapingProfile(d, v, "egress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("ingress_shaping_profile"); ok || d.HasChange("ingress_shaping_profile") {
		t, err := expandObjectDynamicInterfaceIngressShapingProfile(d, v, "ingress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDynamicInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("platform_mapping"); ok || d.HasChange("platform_mapping") {
		t, err := expandObjectDynamicInterfacePlatformMapping(d, v, "platform_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform_mapping"] = t
		}
	}

	if v, ok := d.GetOk("single_intf"); ok || d.HasChange("single_intf") {
		t, err := expandObjectDynamicInterfaceSingleIntf(d, v, "single_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-intf"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandObjectDynamicInterfaceWildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_intf"); ok || d.HasChange("wildcard_intf") {
		t, err := expandObjectDynamicInterfaceWildcardIntf(d, v, "wildcard_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-intf"] = t
		}
	}

	if v, ok := d.GetOk("zone_only"); ok || d.HasChange("zone_only") {
		t, err := expandObjectDynamicInterfaceZoneOnly(d, v, "zone_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone-only"] = t
		}
	}

	return &obj, nil
}
