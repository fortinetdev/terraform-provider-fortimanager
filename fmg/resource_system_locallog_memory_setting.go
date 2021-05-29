// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for memory buffer.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogMemorySetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogMemorySettingUpdate,
		Read:   resourceSystemLocallogMemorySettingRead,
		Update: resourceSystemLocallogMemorySettingUpdate,
		Delete: resourceSystemLocallogMemorySettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogMemorySettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogMemorySetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogMemorySetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogMemorySetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogMemorySetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogMemorySetting")

	return resourceSystemLocallogMemorySettingRead(d, m)
}

func resourceSystemLocallogMemorySettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogMemorySetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogMemorySetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogMemorySettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogMemorySetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogMemorySetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogMemorySetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogMemorySetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogMemorySettingDiskfull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "overwrite",
			2: "nolog",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogMemorySettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "emergency",
			1: "alert",
			2: "critical",
			3: "error",
			4: "warning",
			5: "notification",
			6: "information",
			7: "debug",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogMemorySettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemLocallogMemorySetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("diskfull", flattenSystemLocallogMemorySettingDiskfull(o["diskfull"], d, "diskfull")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskfull"], "SystemLocallogMemorySetting-Diskfull"); ok {
			if err = d.Set("diskfull", vv); err != nil {
				return fmt.Errorf("Error reading diskfull: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("severity", flattenSystemLocallogMemorySettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SystemLocallogMemorySetting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLocallogMemorySettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogMemorySetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogMemorySettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogMemorySettingDiskfull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogMemorySettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogMemorySettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogMemorySetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("diskfull"); ok {
		t, err := expandSystemLocallogMemorySettingDiskfull(d, v, "diskfull")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandSystemLocallogMemorySettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemLocallogMemorySettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
