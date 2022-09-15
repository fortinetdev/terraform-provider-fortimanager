// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic MulticastInterface

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicMulticastInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicMulticastInterfaceCreate,
		Read:   resourceObjectDynamicMulticastInterfaceRead,
		Update: resourceObjectDynamicMulticastInterfaceUpdate,
		Delete: resourceObjectDynamicMulticastInterfaceDelete,

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
			"default_mapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"defmap_intf": &schema.Schema{
				Type:     schema.TypeString,
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
						"local_intf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectDynamicMulticastInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicMulticastInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicMulticastInterface resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDynamicMulticastInterface(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicMulticastInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicMulticastInterfaceRead(d, m)
}

func resourceObjectDynamicMulticastInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicMulticastInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicMulticastInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDynamicMulticastInterface(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicMulticastInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicMulticastInterfaceRead(d, m)
}

func resourceObjectDynamicMulticastInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDynamicMulticastInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicMulticastInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicMulticastInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDynamicMulticastInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicMulticastInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicMulticastInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicMulticastInterface resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicMulticastInterfaceDefaultMapping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceDefmapIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicMulticastInterfaceDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectDynamicMulticastInterface-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_intf"
		if _, ok := i["local-intf"]; ok {
			v := flattenObjectDynamicMulticastInterfaceDynamicMappingLocalIntf(i["local-intf"], d, pre_append)
			tmp["local_intf"] = fortiAPISubPartPatch(v, "ObjectDynamicMulticastInterface-DynamicMapping-LocalIntf")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicMulticastInterfaceDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicMulticastInterfaceDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicMulticastInterfaceDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectDynamicMulticastInterfaceDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectDynamicMulticastInterfaceDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicMulticastInterfaceDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceDynamicMappingLocalIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicMulticastInterfaceZoneOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicMulticastInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("default_mapping", flattenObjectDynamicMulticastInterfaceDefaultMapping(o["default-mapping"], d, "default_mapping")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-mapping"], "ObjectDynamicMulticastInterface-DefaultMapping"); ok {
			if err = d.Set("default_mapping", vv); err != nil {
				return fmt.Errorf("Error reading default_mapping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_mapping: %v", err)
		}
	}

	if err = d.Set("defmap_intf", flattenObjectDynamicMulticastInterfaceDefmapIntf(o["defmap-intf"], d, "defmap_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["defmap-intf"], "ObjectDynamicMulticastInterface-DefmapIntf"); ok {
			if err = d.Set("defmap_intf", vv); err != nil {
				return fmt.Errorf("Error reading defmap_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading defmap_intf: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectDynamicMulticastInterfaceDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectDynamicMulticastInterface-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectDynamicMulticastInterfaceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicMulticastInterface-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectDynamicMulticastInterfaceDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicMulticastInterface-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectDynamicMulticastInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDynamicMulticastInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("zone_only", flattenObjectDynamicMulticastInterfaceZoneOnly(o["zone-only"], d, "zone_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["zone-only"], "ObjectDynamicMulticastInterface-ZoneOnly"); ok {
			if err = d.Set("zone_only", vv); err != nil {
				return fmt.Errorf("Error reading zone_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading zone_only: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicMulticastInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicMulticastInterfaceDefaultMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceDefmapIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_scope"], _ = expandObjectDynamicMulticastInterfaceDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["local-intf"], _ = expandObjectDynamicMulticastInterfaceDynamicMappingLocalIntf(d, i["local_intf"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicMulticastInterfaceDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectDynamicMulticastInterfaceDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectDynamicMulticastInterfaceDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicMulticastInterfaceDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceDynamicMappingLocalIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicMulticastInterfaceZoneOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicMulticastInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_mapping"); ok || d.HasChange("default_mapping") {
		t, err := expandObjectDynamicMulticastInterfaceDefaultMapping(d, v, "default_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-mapping"] = t
		}
	}

	if v, ok := d.GetOk("defmap_intf"); ok || d.HasChange("defmap_intf") {
		t, err := expandObjectDynamicMulticastInterfaceDefmapIntf(d, v, "defmap_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defmap-intf"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectDynamicMulticastInterfaceDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectDynamicMulticastInterfaceDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDynamicMulticastInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("zone_only"); ok || d.HasChange("zone_only") {
		t, err := expandObjectDynamicMulticastInterfaceZoneOnly(d, v, "zone_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone-only"] = t
		}
	}

	return &obj, nil
}
