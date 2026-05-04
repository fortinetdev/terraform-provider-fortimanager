// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg Script

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFmgScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgScriptCreate,
		Read:   resourceObjectFmgScriptRead,
		Update: resourceObjectFmgScriptUpdate,
		Delete: resourceObjectFmgScriptDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			},
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_build": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filter_device": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filter_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_ostype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filter_osver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"filter_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"member": &schema.Schema{
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
			"schedule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"datetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"day_of_week": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"run_on_db": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"timestamp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
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

func resourceObjectFmgScriptCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFmgScript(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgScript resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFmgScript(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFmgScript(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFmgScript resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFmgScript(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFmgScript resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFmgScriptRead(d, m)
}

func resourceObjectFmgScriptUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFmgScript(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgScript resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFmgScript(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgScript resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFmgScriptRead(d, m)
}

func resourceObjectFmgScriptDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFmgScript(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgScript resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgScriptRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFmgScript(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFmgScript resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgScript(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgScript resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgScriptContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterBuild(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterOstype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterOsver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterPlatform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptFilterSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFmgScriptName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptSchedule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFmgScriptScheduleDatetime(i["datetime"], d, pre_append)
			tmp["datetime"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-Datetime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "day_of_week"
		if _, ok := i["day-of-week"]; ok {
			v := flattenObjectFmgScriptScheduleDayOfWeek(i["day-of-week"], d, pre_append)
			tmp["day_of_week"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-DayOfWeek")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenObjectFmgScriptScheduleDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "run_on_db"
		if _, ok := i["run-on-db"]; ok {
			v := flattenObjectFmgScriptScheduleRunOnDb(i["run-on-db"], d, pre_append)
			tmp["run_on_db"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-RunOnDb")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timestamp"
		if _, ok := i["timestamp"]; ok {
			v := flattenObjectFmgScriptScheduleTimestamp(i["timestamp"], d, pre_append)
			tmp["timestamp"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-Timestamp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFmgScriptScheduleType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenObjectFmgScriptScheduleUser(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "ObjectFmgScript-Schedule-User")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFmgScriptScheduleDatetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleDayOfWeek(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleRunOnDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFmgScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("content", flattenObjectFmgScriptContent(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "ObjectFmgScript-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("desc", flattenObjectFmgScriptDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "ObjectFmgScript-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("filter_build", flattenObjectFmgScriptFilterBuild(o["filter_build"], d, "filter_build")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_build"], "ObjectFmgScript-FilterBuild"); ok {
			if err = d.Set("filter_build", vv); err != nil {
				return fmt.Errorf("Error reading filter_build: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_build: %v", err)
		}
	}

	if err = d.Set("filter_device", flattenObjectFmgScriptFilterDevice(o["filter_device"], d, "filter_device")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_device"], "ObjectFmgScript-FilterDevice"); ok {
			if err = d.Set("filter_device", vv); err != nil {
				return fmt.Errorf("Error reading filter_device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_device: %v", err)
		}
	}

	if err = d.Set("filter_hostname", flattenObjectFmgScriptFilterHostname(o["filter_hostname"], d, "filter_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_hostname"], "ObjectFmgScript-FilterHostname"); ok {
			if err = d.Set("filter_hostname", vv); err != nil {
				return fmt.Errorf("Error reading filter_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_hostname: %v", err)
		}
	}

	if err = d.Set("filter_ostype", flattenObjectFmgScriptFilterOstype(o["filter_ostype"], d, "filter_ostype")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_ostype"], "ObjectFmgScript-FilterOstype"); ok {
			if err = d.Set("filter_ostype", vv); err != nil {
				return fmt.Errorf("Error reading filter_ostype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_ostype: %v", err)
		}
	}

	if err = d.Set("filter_osver", flattenObjectFmgScriptFilterOsver(o["filter_osver"], d, "filter_osver")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_osver"], "ObjectFmgScript-FilterOsver"); ok {
			if err = d.Set("filter_osver", vv); err != nil {
				return fmt.Errorf("Error reading filter_osver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_osver: %v", err)
		}
	}

	if err = d.Set("filter_platform", flattenObjectFmgScriptFilterPlatform(o["filter_platform"], d, "filter_platform")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_platform"], "ObjectFmgScript-FilterPlatform"); ok {
			if err = d.Set("filter_platform", vv); err != nil {
				return fmt.Errorf("Error reading filter_platform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_platform: %v", err)
		}
	}

	if err = d.Set("filter_serial", flattenObjectFmgScriptFilterSerial(o["filter_serial"], d, "filter_serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter_serial"], "ObjectFmgScript-FilterSerial"); ok {
			if err = d.Set("filter_serial", vv); err != nil {
				return fmt.Errorf("Error reading filter_serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_serial: %v", err)
		}
	}

	if err = d.Set("member", flattenObjectFmgScriptMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectFmgScript-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFmgScriptName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFmgScript-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("schedule", flattenObjectFmgScriptSchedule(o["schedule"], d, "schedule")); err != nil {
			if vv, ok := fortiAPIPatch(o["schedule"], "ObjectFmgScript-Schedule"); ok {
				if err = d.Set("schedule", vv); err != nil {
					return fmt.Errorf("Error reading schedule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("schedule"); ok {
			if err = d.Set("schedule", flattenObjectFmgScriptSchedule(o["schedule"], d, "schedule")); err != nil {
				if vv, ok := fortiAPIPatch(o["schedule"], "ObjectFmgScript-Schedule"); ok {
					if err = d.Set("schedule", vv); err != nil {
						return fmt.Errorf("Error reading schedule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading schedule: %v", err)
				}
			}
		}
	}

	if err = d.Set("target", flattenObjectFmgScriptTarget(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "ObjectFmgScript-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFmgScriptType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFmgScript-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgScriptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgScriptContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterBuild(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterOstype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterOsver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterPlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptFilterSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFmgScriptName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "datetime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["datetime"], _ = expandObjectFmgScriptScheduleDatetime(d, i["datetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "day_of_week"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["day-of-week"], _ = expandObjectFmgScriptScheduleDayOfWeek(d, i["day_of_week"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandObjectFmgScriptScheduleDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "run_on_db"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["run-on-db"], _ = expandObjectFmgScriptScheduleRunOnDb(d, i["run_on_db"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timestamp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["timestamp"], _ = expandObjectFmgScriptScheduleTimestamp(d, i["timestamp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFmgScriptScheduleType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandObjectFmgScriptScheduleUser(d, i["user"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFmgScriptScheduleDatetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleDayOfWeek(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleRunOnDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFmgScript(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("content"); ok || d.HasChange("content") {
		t, err := expandObjectFmgScriptContent(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok || d.HasChange("desc") {
		t, err := expandObjectFmgScriptDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("filter_build"); ok || d.HasChange("filter_build") {
		t, err := expandObjectFmgScriptFilterBuild(d, v, "filter_build")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_build"] = t
		}
	}

	if v, ok := d.GetOk("filter_device"); ok || d.HasChange("filter_device") {
		t, err := expandObjectFmgScriptFilterDevice(d, v, "filter_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_device"] = t
		}
	}

	if v, ok := d.GetOk("filter_hostname"); ok || d.HasChange("filter_hostname") {
		t, err := expandObjectFmgScriptFilterHostname(d, v, "filter_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_hostname"] = t
		}
	}

	if v, ok := d.GetOk("filter_ostype"); ok || d.HasChange("filter_ostype") {
		t, err := expandObjectFmgScriptFilterOstype(d, v, "filter_ostype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_ostype"] = t
		}
	}

	if v, ok := d.GetOk("filter_osver"); ok || d.HasChange("filter_osver") {
		t, err := expandObjectFmgScriptFilterOsver(d, v, "filter_osver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_osver"] = t
		}
	}

	if v, ok := d.GetOk("filter_platform"); ok || d.HasChange("filter_platform") {
		t, err := expandObjectFmgScriptFilterPlatform(d, v, "filter_platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_platform"] = t
		}
	}

	if v, ok := d.GetOk("filter_serial"); ok || d.HasChange("filter_serial") {
		t, err := expandObjectFmgScriptFilterSerial(d, v, "filter_serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter_serial"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandObjectFmgScriptMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFmgScriptName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandObjectFmgScriptSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectFmgScriptTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFmgScriptType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
