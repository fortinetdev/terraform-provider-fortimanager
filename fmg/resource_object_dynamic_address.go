// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic Address

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicAddressCreate,
		Read:   resourceObjectDynamicAddressRead,
		Update: resourceObjectDynamicAddressUpdate,
		Delete: resourceObjectDynamicAddressDelete,

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
			"default": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_addr_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectDynamicAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectDynamicAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectDynamicAddress(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicAddress resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicAddressRead(d, m)
}

func resourceObjectDynamicAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectDynamicAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDynamicAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDynamicAddressRead(d, m)
}

func resourceObjectDynamicAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectDynamicAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectDynamicAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicAddress resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicAddressDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDynamicAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicAddressDynamicAddrMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicAddressDynamicAddrMappingAddr(i["addr"], d, pre_append)
			tmp["addr"] = fortiAPISubPartPatch(v, "ObjectDynamicAddress-DynamicAddrMapping-Addr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDynamicAddressDynamicAddrMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDynamicAddress-DynamicAddrMapping-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectDynamicAddressDynamicAddrMappingAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicAddressDynamicAddrMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("default", flattenObjectDynamicAddressDefault(o["default"], d, "default")); err != nil {
		if vv, ok := fortiAPIPatch(o["default"], "ObjectDynamicAddress-Default"); ok {
			if err = d.Set("default", vv); err != nil {
				return fmt.Errorf("Error reading default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectDynamicAddressDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectDynamicAddress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_addr_mapping", flattenObjectDynamicAddressDynamicAddrMapping(o["dynamic_addr_mapping"], d, "dynamic_addr_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_addr_mapping"], "ObjectDynamicAddress-DynamicAddrMapping"); ok {
				if err = d.Set("dynamic_addr_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_addr_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_addr_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_addr_mapping"); ok {
			if err = d.Set("dynamic_addr_mapping", flattenObjectDynamicAddressDynamicAddrMapping(o["dynamic_addr_mapping"], d, "dynamic_addr_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_addr_mapping"], "ObjectDynamicAddress-DynamicAddrMapping"); ok {
					if err = d.Set("dynamic_addr_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_addr_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_addr_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectDynamicAddressName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDynamicAddress-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicAddressDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDynamicAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicAddressDynamicAddrMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr"], _ = expandObjectDynamicAddressDynamicAddrMappingAddr(d, i["addr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectDynamicAddressDynamicAddrMappingId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectDynamicAddressDynamicAddrMappingAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicAddressDynamicAddrMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default"); ok || d.HasChange("default") {
		t, err := expandObjectDynamicAddressDefault(d, v, "default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectDynamicAddressDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_addr_mapping"); ok || d.HasChange("dynamic_addr_mapping") {
		t, err := expandObjectDynamicAddressDynamicAddrMapping(d, v, "dynamic_addr_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_addr_mapping"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDynamicAddressName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
