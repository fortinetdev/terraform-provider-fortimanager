// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpSysinfoUpdate,
		Read:   resourceSystemSnmpSysinfoRead,
		Update: resourceSystemSnmpSysinfoUpdate,
		Delete: resourceSystemSnmpSysinfoDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"contact_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"engine_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_cpu_high_exclude_nice_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_low_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSnmpSysinfo(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpSysinfo(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSnmpSysinfo")

	return resourceSystemSnmpSysinfoRead(d, m)
}

func resourceSystemSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSnmpSysinfo(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSnmpSysinfo(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpSysinfo(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapCpuHighExcludeNiceThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("contact_info", flattenSystemSnmpSysinfoContactInfo(o["contact_info"], d, "contact_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact_info"], "SystemSnmpSysinfo-ContactInfo"); ok {
			if err = d.Set("contact_info", vv); err != nil {
				return fmt.Errorf("Error reading contact_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemSnmpSysinfoDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemSnmpSysinfo-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenSystemSnmpSysinfoEngineId(o["engine-id"], d, "engine_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["engine-id"], "SystemSnmpSysinfo-EngineId"); ok {
			if err = d.Set("engine_id", vv); err != nil {
				return fmt.Errorf("Error reading engine_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("location", flattenSystemSnmpSysinfoLocation(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "SystemSnmpSysinfo-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSnmpSysinfoStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSnmpSysinfo-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_cpu_high_exclude_nice_threshold", flattenSystemSnmpSysinfoTrapCpuHighExcludeNiceThreshold(o["trap-cpu-high-exclude-nice-threshold"], d, "trap_cpu_high_exclude_nice_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-cpu-high-exclude-nice-threshold"], "SystemSnmpSysinfo-TrapCpuHighExcludeNiceThreshold"); ok {
			if err = d.Set("trap_cpu_high_exclude_nice_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_cpu_high_exclude_nice_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_cpu_high_exclude_nice_threshold: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", flattenSystemSnmpSysinfoTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-high-cpu-threshold"], "SystemSnmpSysinfo-TrapHighCpuThreshold"); ok {
			if err = d.Set("trap_high_cpu_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSystemSnmpSysinfoTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-low-memory-threshold"], "SystemSnmpSysinfo-TrapLowMemoryThreshold"); ok {
			if err = d.Set("trap_low_memory_threshold", vv); err != nil {
				return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapCpuHighExcludeNiceThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpSysinfo(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("contact_info"); ok {
		t, err := expandSystemSnmpSysinfoContactInfo(d, v, "contact_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact_info"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemSnmpSysinfoDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("engine_id"); ok {
		t, err := expandSystemSnmpSysinfoEngineId(d, v, "engine_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {
		t, err := expandSystemSnmpSysinfoLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpSysinfoStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_cpu_high_exclude_nice_threshold"); ok {
		t, err := expandSystemSnmpSysinfoTrapCpuHighExcludeNiceThreshold(d, v, "trap_cpu_high_exclude_nice_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-cpu-high-exclude-nice-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok {
		t, err := expandSystemSnmpSysinfoTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_low_memory_threshold"); ok {
		t, err := expandSystemSnmpSysinfoTrapLowMemoryThreshold(d, v, "trap_low_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-low-memory-threshold"] = t
		}
	}

	return &obj, nil
}
