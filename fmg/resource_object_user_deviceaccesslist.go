// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure device access control lists.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserDeviceAccessList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserDeviceAccessListCreate,
		Read:   resourceObjectUserDeviceAccessListRead,
		Update: resourceObjectUserDeviceAccessListUpdate,
		Delete: resourceObjectUserDeviceAccessListDelete,

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
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectUserDeviceAccessListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserDeviceAccessList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserDeviceAccessList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserDeviceAccessList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserDeviceAccessList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserDeviceAccessListRead(d, m)
}

func resourceObjectUserDeviceAccessListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserDeviceAccessList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserDeviceAccessList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserDeviceAccessList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserDeviceAccessList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserDeviceAccessListRead(d, m)
}

func resourceObjectUserDeviceAccessListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserDeviceAccessList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserDeviceAccessList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserDeviceAccessListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserDeviceAccessList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserDeviceAccessList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserDeviceAccessList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserDeviceAccessList resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserDeviceAccessListDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDeviceAccessListDeviceList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserDeviceAccessListDeviceListAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectUserDeviceAccessList-DeviceList-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenObjectUserDeviceAccessListDeviceListDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "ObjectUserDeviceAccessList-DeviceList-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserDeviceAccessListDeviceListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserDeviceAccessList-DeviceList-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserDeviceAccessListDeviceListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDeviceAccessListDeviceListDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDeviceAccessListDeviceListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserDeviceAccessListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserDeviceAccessList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("default_action", flattenObjectUserDeviceAccessListDefaultAction(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "ObjectUserDeviceAccessList-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("device_list", flattenObjectUserDeviceAccessListDeviceList(o["device-list"], d, "device_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["device-list"], "ObjectUserDeviceAccessList-DeviceList"); ok {
				if err = d.Set("device_list", vv); err != nil {
					return fmt.Errorf("Error reading device_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_list"); ok {
			if err = d.Set("device_list", flattenObjectUserDeviceAccessListDeviceList(o["device-list"], d, "device_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["device-list"], "ObjectUserDeviceAccessList-DeviceList"); ok {
					if err = d.Set("device_list", vv); err != nil {
						return fmt.Errorf("Error reading device_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectUserDeviceAccessListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserDeviceAccessList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectUserDeviceAccessListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserDeviceAccessListDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDeviceAccessListDeviceList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectUserDeviceAccessListDeviceListAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device"], _ = expandObjectUserDeviceAccessListDeviceListDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectUserDeviceAccessListDeviceListId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserDeviceAccessListDeviceListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDeviceAccessListDeviceListDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDeviceAccessListDeviceListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserDeviceAccessListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserDeviceAccessList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_action"); ok {
		t, err := expandObjectUserDeviceAccessListDefaultAction(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("device_list"); ok {
		t, err := expandObjectUserDeviceAccessListDeviceList(d, v, "device_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserDeviceAccessListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
