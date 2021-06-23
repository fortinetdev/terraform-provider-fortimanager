// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Script table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmdbScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbScriptCreate,
		Read:   resourceDvmdbScriptRead,
		Update: resourceDvmdbScriptUpdate,
		Delete: resourceDvmdbScriptDelete,

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
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_build": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filter_device": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filter_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_ostype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_osver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modification_time": &schema.Schema{
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
			"script_schedule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"datetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"day_of_week": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"run_on_db": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
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

func resourceDvmdbScriptCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbScript(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbScript resource while getting object: %v", err)
	}

	_, err = c.CreateDvmdbScript(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating DvmdbScript resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbScriptRead(d, m)
}

func resourceDvmdbScriptUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbScript(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbScript resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbScript(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbScript resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbScriptRead(d, m)
}

func resourceDvmdbScriptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteDvmdbScript(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbScript resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbScriptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadDvmdbScript(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbScript resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbScript(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbScript resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbScriptContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterBuild(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterOstype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterOsver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterPlatform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptFilterSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptModificationTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptScriptSchedule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "datetime"
		if _, ok := i["datetime"]; ok {
			v := flattenDvmdbScriptScriptScheduleDatetime(i["datetime"], d, pre_append)
			tmp["datetime"] = fortiAPISubPartPatch(v, "DvmdbScript-ScriptSchedule-Datetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "day_of_week"
		if _, ok := i["day_of_week"]; ok {
			v := flattenDvmdbScriptScriptScheduleDayOfWeek(i["day_of_week"], d, pre_append)
			tmp["day_of_week"] = fortiAPISubPartPatch(v, "DvmdbScript-ScriptSchedule-DayOfWeek")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenDvmdbScriptScriptScheduleDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "DvmdbScript-ScriptSchedule-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenDvmdbScriptScriptScheduleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmdbScript-ScriptSchedule-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "run_on_db"
		if _, ok := i["run_on_db"]; ok {
			v := flattenDvmdbScriptScriptScheduleRunOnDb(i["run_on_db"], d, pre_append)
			tmp["run_on_db"] = fortiAPISubPartPatch(v, "DvmdbScript-ScriptSchedule-RunOnDb")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenDvmdbScriptScriptScheduleType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "DvmdbScript-ScriptSchedule-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenDvmdbScriptScriptScheduleDatetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptScriptScheduleDayOfWeek(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptScriptScheduleDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptScriptScheduleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptScriptScheduleRunOnDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptScriptScheduleType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbScriptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("content", flattenDvmdbScriptContent(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "DvmdbScript-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("desc", flattenDvmdbScriptDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "DvmdbScript-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("filter_build", flattenDvmdbScriptFilterBuild(o["filter_build"], d, "filter_build")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_build"], "DvmdbScript-FilterBuild"); ok {
			if err = d.Set("filter_build", vv); err != nil {
				return fmt.Errorf("Error reading filter_build: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_build: %v", err)
		}
	}

	if err = d.Set("filter_device", flattenDvmdbScriptFilterDevice(o["filter_device"], d, "filter_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_device"], "DvmdbScript-FilterDevice"); ok {
			if err = d.Set("filter_device", vv); err != nil {
				return fmt.Errorf("Error reading filter_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_device: %v", err)
		}
	}

	if err = d.Set("filter_hostname", flattenDvmdbScriptFilterHostname(o["filter_hostname"], d, "filter_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_hostname"], "DvmdbScript-FilterHostname"); ok {
			if err = d.Set("filter_hostname", vv); err != nil {
				return fmt.Errorf("Error reading filter_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_hostname: %v", err)
		}
	}

	if err = d.Set("filter_ostype", flattenDvmdbScriptFilterOstype(o["filter_ostype"], d, "filter_ostype")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_ostype"], "DvmdbScript-FilterOstype"); ok {
			if err = d.Set("filter_ostype", vv); err != nil {
				return fmt.Errorf("Error reading filter_ostype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_ostype: %v", err)
		}
	}

	if err = d.Set("filter_osver", flattenDvmdbScriptFilterOsver(o["filter_osver"], d, "filter_osver")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_osver"], "DvmdbScript-FilterOsver"); ok {
			if err = d.Set("filter_osver", vv); err != nil {
				return fmt.Errorf("Error reading filter_osver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_osver: %v", err)
		}
	}

	if err = d.Set("filter_platform", flattenDvmdbScriptFilterPlatform(o["filter_platform"], d, "filter_platform")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_platform"], "DvmdbScript-FilterPlatform"); ok {
			if err = d.Set("filter_platform", vv); err != nil {
				return fmt.Errorf("Error reading filter_platform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_platform: %v", err)
		}
	}

	if err = d.Set("filter_serial", flattenDvmdbScriptFilterSerial(o["filter_serial"], d, "filter_serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_serial"], "DvmdbScript-FilterSerial"); ok {
			if err = d.Set("filter_serial", vv); err != nil {
				return fmt.Errorf("Error reading filter_serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_serial: %v", err)
		}
	}

	if err = d.Set("modification_time", flattenDvmdbScriptModificationTime(o["modification_time"], d, "modification_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["modification_time"], "DvmdbScript-ModificationTime"); ok {
			if err = d.Set("modification_time", vv); err != nil {
				return fmt.Errorf("Error reading modification_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modification_time: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbScriptName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbScript-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("script_schedule", flattenDvmdbScriptScriptSchedule(o["script_schedule"], d, "script_schedule")); err != nil {
			if vv, ok := fortiAPIPatch(o["script_schedule"], "DvmdbScript-ScriptSchedule"); ok {
				if err = d.Set("script_schedule", vv); err != nil {
					return fmt.Errorf("Error reading script_schedule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading script_schedule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("script_schedule"); ok {
			if err = d.Set("script_schedule", flattenDvmdbScriptScriptSchedule(o["script_schedule"], d, "script_schedule")); err != nil {
				if vv, ok := fortiAPIPatch(o["script_schedule"], "DvmdbScript-ScriptSchedule"); ok {
					if err = d.Set("script_schedule", vv); err != nil {
						return fmt.Errorf("Error reading script_schedule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading script_schedule: %v", err)
				}
			}
		}
	}

	if err = d.Set("target", flattenDvmdbScriptTarget(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "DvmdbScript-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	if err = d.Set("type", flattenDvmdbScriptType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "DvmdbScript-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenDvmdbScriptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbScriptContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterBuild(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterOstype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterOsver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterPlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptFilterSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptModificationTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptScriptSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "datetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["datetime"], _ = expandDvmdbScriptScriptScheduleDatetime(d, i["datetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "day_of_week"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["day_of_week"], _ = expandDvmdbScriptScriptScheduleDayOfWeek(d, i["day_of_week"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device"], _ = expandDvmdbScriptScriptScheduleDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandDvmdbScriptScriptScheduleName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "run_on_db"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["run_on_db"], _ = expandDvmdbScriptScriptScheduleRunOnDb(d, i["run_on_db"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandDvmdbScriptScriptScheduleType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDvmdbScriptScriptScheduleDatetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptScriptScheduleDayOfWeek(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptScriptScheduleDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptScriptScheduleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptScriptScheduleRunOnDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptScriptScheduleType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbScriptType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbScript(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("content"); ok {
		t, err := expandDvmdbScriptContent(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandDvmdbScriptDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("filter_build"); ok {
		t, err := expandDvmdbScriptFilterBuild(d, v, "filter_build")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_build"] = t
		}
	}

	if v, ok := d.GetOk("filter_device"); ok {
		t, err := expandDvmdbScriptFilterDevice(d, v, "filter_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_device"] = t
		}
	}

	if v, ok := d.GetOk("filter_hostname"); ok {
		t, err := expandDvmdbScriptFilterHostname(d, v, "filter_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_hostname"] = t
		}
	}

	if v, ok := d.GetOk("filter_ostype"); ok {
		t, err := expandDvmdbScriptFilterOstype(d, v, "filter_ostype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_ostype"] = t
		}
	}

	if v, ok := d.GetOk("filter_osver"); ok {
		t, err := expandDvmdbScriptFilterOsver(d, v, "filter_osver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_osver"] = t
		}
	}

	if v, ok := d.GetOk("filter_platform"); ok {
		t, err := expandDvmdbScriptFilterPlatform(d, v, "filter_platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_platform"] = t
		}
	}

	if v, ok := d.GetOk("filter_serial"); ok {
		t, err := expandDvmdbScriptFilterSerial(d, v, "filter_serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_serial"] = t
		}
	}

	if v, ok := d.GetOk("modification_time"); ok {
		t, err := expandDvmdbScriptModificationTime(d, v, "modification_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modification_time"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDvmdbScriptName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("script_schedule"); ok {
		t, err := expandDvmdbScriptScriptSchedule(d, v, "script_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script_schedule"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok {
		t, err := expandDvmdbScriptTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandDvmdbScriptType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
