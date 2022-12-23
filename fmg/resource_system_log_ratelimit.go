// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Logging rate limit.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogRatelimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogRatelimitUpdate,
		Read:   resourceSystemLogRatelimitRead,
		Update: resourceSystemLogRatelimitUpdate,
		Delete: resourceSystemLogRatelimitDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ratelimit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"device_ratelimit_default": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ratelimits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ratelimit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"system_ratelimit": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemLogRatelimitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogRatelimit(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimit resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogRatelimit(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogRatelimit")

	return resourceSystemLogRatelimitRead(d, m)
}

func resourceSystemLogRatelimitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLogRatelimit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogRatelimit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogRatelimitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogRatelimit(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogRatelimit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimit resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogRatelimitDeviceSLR(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenSystemLogRatelimitDeviceDeviceSLR(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenSystemLogRatelimitDeviceFilterTypeSLR(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogRatelimitDeviceIdSLR(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ratelimit"
		if _, ok := i["ratelimit"]; ok {
			v := flattenSystemLogRatelimitDeviceRatelimitSLR(i["ratelimit"], d, pre_append)
			tmp["ratelimit"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-Ratelimit")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogRatelimitDeviceDeviceSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceFilterTypeSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceIdSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceRatelimitSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceRatelimitDefaultSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitModeSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsSLR(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenSystemLogRatelimitRatelimitsFilterSLR(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Ratelimits-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenSystemLogRatelimitRatelimitsFilterTypeSLR(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Ratelimits-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogRatelimitRatelimitsIdSLR(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Ratelimits-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ratelimit"
		if _, ok := i["ratelimit"]; ok {
			v := flattenSystemLogRatelimitRatelimitsRatelimitSLR(i["ratelimit"], d, pre_append)
			tmp["ratelimit"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Ratelimits-Ratelimit")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogRatelimitRatelimitsFilterSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsFilterTypeSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsIdSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsRatelimitSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitSystemRatelimitSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogRatelimit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("device", flattenSystemLogRatelimitDeviceSLR(o["device"], d, "device")); err != nil {
			if vv, ok := fortiAPIPatch(o["device"], "SystemLogRatelimit-Device"); ok {
				if err = d.Set("device", vv); err != nil {
					return fmt.Errorf("Error reading device: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device"); ok {
			if err = d.Set("device", flattenSystemLogRatelimitDeviceSLR(o["device"], d, "device")); err != nil {
				if vv, ok := fortiAPIPatch(o["device"], "SystemLogRatelimit-Device"); ok {
					if err = d.Set("device", vv); err != nil {
						return fmt.Errorf("Error reading device: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device: %v", err)
				}
			}
		}
	}

	if err = d.Set("device_ratelimit_default", flattenSystemLogRatelimitDeviceRatelimitDefaultSLR(o["device-ratelimit-default"], d, "device_ratelimit_default")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-ratelimit-default"], "SystemLogRatelimit-DeviceRatelimitDefault"); ok {
			if err = d.Set("device_ratelimit_default", vv); err != nil {
				return fmt.Errorf("Error reading device_ratelimit_default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_ratelimit_default: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemLogRatelimitModeSLR(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemLogRatelimit-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ratelimits", flattenSystemLogRatelimitRatelimitsSLR(o["ratelimits"], d, "ratelimits")); err != nil {
			if vv, ok := fortiAPIPatch(o["ratelimits"], "SystemLogRatelimit-Ratelimits"); ok {
				if err = d.Set("ratelimits", vv); err != nil {
					return fmt.Errorf("Error reading ratelimits: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ratelimits: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ratelimits"); ok {
			if err = d.Set("ratelimits", flattenSystemLogRatelimitRatelimitsSLR(o["ratelimits"], d, "ratelimits")); err != nil {
				if vv, ok := fortiAPIPatch(o["ratelimits"], "SystemLogRatelimit-Ratelimits"); ok {
					if err = d.Set("ratelimits", vv); err != nil {
						return fmt.Errorf("Error reading ratelimits: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ratelimits: %v", err)
				}
			}
		}
	}

	if err = d.Set("system_ratelimit", flattenSystemLogRatelimitSystemRatelimitSLR(o["system-ratelimit"], d, "system_ratelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-ratelimit"], "SystemLogRatelimit-SystemRatelimit"); ok {
			if err = d.Set("system_ratelimit", vv); err != nil {
				return fmt.Errorf("Error reading system_ratelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_ratelimit: %v", err)
		}
	}

	return nil
}

func flattenSystemLogRatelimitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogRatelimitDeviceSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandSystemLogRatelimitDeviceDeviceSLR(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandSystemLogRatelimitDeviceFilterTypeSLR(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogRatelimitDeviceIdSLR(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ratelimit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ratelimit"], _ = expandSystemLogRatelimitDeviceRatelimitSLR(d, i["ratelimit"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogRatelimitDeviceDeviceSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceFilterTypeSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceIdSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceRatelimitSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceRatelimitDefaultSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitModeSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandSystemLogRatelimitRatelimitsFilterSLR(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-type"], _ = expandSystemLogRatelimitRatelimitsFilterTypeSLR(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogRatelimitRatelimitsIdSLR(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ratelimit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ratelimit"], _ = expandSystemLogRatelimitRatelimitsRatelimitSLR(d, i["ratelimit"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogRatelimitRatelimitsFilterSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsFilterTypeSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsIdSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsRatelimitSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitSystemRatelimitSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogRatelimit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemLogRatelimitDeviceSLR(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("device_ratelimit_default"); ok || d.HasChange("device_ratelimit_default") {
		t, err := expandSystemLogRatelimitDeviceRatelimitDefaultSLR(d, v, "device_ratelimit_default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-ratelimit-default"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemLogRatelimitModeSLR(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("ratelimits"); ok || d.HasChange("ratelimits") {
		t, err := expandSystemLogRatelimitRatelimitsSLR(d, v, "ratelimits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ratelimits"] = t
		}
	}

	if v, ok := d.GetOk("system_ratelimit"); ok || d.HasChange("system_ratelimit") {
		t, err := expandSystemLogRatelimitSystemRatelimitSLR(d, v, "system_ratelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-ratelimit"] = t
		}
	}

	return &obj, nil
}
