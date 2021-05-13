// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic VirtualWanLinkNeighbor

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectDynamicVirtualWanLinkNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicVirtualWanLinkNeighborCreate,
		Read:   resourceObjectDynamicVirtualWanLinkNeighborRead,
		Update: resourceObjectDynamicVirtualWanLinkNeighborUpdate,
		Delete: resourceObjectDynamicVirtualWanLinkNeighborDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
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
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip": &schema.Schema{
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
			"role": &schema.Schema{
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

func resourceObjectDynamicVirtualWanLinkNeighborCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicVirtualWanLinkNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicVirtualWanLinkNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDynamicVirtualWanLinkNeighbor(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicVirtualWanLinkNeighbor resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicVirtualWanLinkNeighborRead(d, m)
}

func resourceObjectDynamicVirtualWanLinkNeighborUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDynamicVirtualWanLinkNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicVirtualWanLinkNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDynamicVirtualWanLinkNeighbor(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicVirtualWanLinkNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicVirtualWanLinkNeighborRead(d, m)
}

func resourceObjectDynamicVirtualWanLinkNeighborDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDynamicVirtualWanLinkNeighbor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicVirtualWanLinkNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicVirtualWanLinkNeighborRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDynamicVirtualWanLinkNeighbor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicVirtualWanLinkNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicVirtualWanLinkNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicVirtualWanLinkNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicVirtualWanLinkNeighborDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkNeighbor-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkNeighbor-DynamicMapping-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkNeighbor-DynamicMapping-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkNeighbor-DynamicMapping-Role")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkNeighborDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectDynamicVirtualWanLinkNeighborDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborDynamicMappingRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "primary",
			2: "secondary",
			3: "standalone",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicVirtualWanLinkNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "primary",
			2: "secondary",
			3: "standalone",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectDynamicVirtualWanLinkNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenObjectDynamicVirtualWanLinkNeighborDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectDynamicVirtualWanLinkNeighbor-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectDynamicVirtualWanLinkNeighborDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicVirtualWanLinkNeighbor-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectDynamicVirtualWanLinkNeighborDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectDynamicVirtualWanLinkNeighbor-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip", flattenObjectDynamicVirtualWanLinkNeighborIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectDynamicVirtualWanLinkNeighbor-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDynamicVirtualWanLinkNeighborName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDynamicVirtualWanLinkNeighbor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("role", flattenObjectDynamicVirtualWanLinkNeighborRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ObjectDynamicVirtualWanLinkNeighbor-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicVirtualWanLinkNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicVirtualWanLinkNeighborDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["_scope"], _ = expandObjectDynamicVirtualWanLinkNeighborDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandObjectDynamicVirtualWanLinkNeighborDynamicMappingDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandObjectDynamicVirtualWanLinkNeighborDynamicMappingIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandObjectDynamicVirtualWanLinkNeighborDynamicMappingRole(d, i["role"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMappingDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMappingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborDynamicMappingRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicVirtualWanLinkNeighborRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicVirtualWanLinkNeighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandObjectDynamicVirtualWanLinkNeighborDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok {
		t, err := expandObjectDynamicVirtualWanLinkNeighborDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandObjectDynamicVirtualWanLinkNeighborIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectDynamicVirtualWanLinkNeighborName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {
		t, err := expandObjectDynamicVirtualWanLinkNeighborRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	return &obj, nil
}
