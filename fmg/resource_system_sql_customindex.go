// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of SQL index fields.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSqlCustomIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSqlCustomIndexCreate,
		Read:   resourceSystemSqlCustomIndexRead,
		Update: resourceSystemSqlCustomIndexUpdate,
		Delete: resourceSystemSqlCustomIndexDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"index_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSqlCustomIndexCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSqlCustomIndex(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSqlCustomIndex resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemSqlCustomIndex(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSqlCustomIndex resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSqlCustomIndexRead(d, m)
}

func resourceSystemSqlCustomIndexUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSqlCustomIndex(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSqlCustomIndex resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSqlCustomIndex(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSqlCustomIndex resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSqlCustomIndexRead(d, m)
}

func resourceSystemSqlCustomIndexDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemSqlCustomIndex(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSqlCustomIndex resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSqlCustomIndexRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSqlCustomIndex(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSqlCustomIndex resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSqlCustomIndex(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSqlCustomIndex resource from API: %v", err)
	}
	return nil
}

func flattenSystemSqlCustomIndexCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexDeviceType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexIndexField2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexLogType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSqlCustomIndex(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("case_sensitive", flattenSystemSqlCustomIndexCaseSensitive2edl(o["case-sensitive"], d, "case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitive"], "SystemSqlCustomIndex-CaseSensitive"); ok {
			if err = d.Set("case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}

	if err = d.Set("device_type", flattenSystemSqlCustomIndexDeviceType2edl(o["device-type"], d, "device_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-type"], "SystemSqlCustomIndex-DeviceType"); ok {
			if err = d.Set("device_type", vv); err != nil {
				return fmt.Errorf("Error reading device_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemSqlCustomIndexId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSqlCustomIndex-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("index_field", flattenSystemSqlCustomIndexIndexField2edl(o["index-field"], d, "index_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["index-field"], "SystemSqlCustomIndex-IndexField"); ok {
			if err = d.Set("index_field", vv); err != nil {
				return fmt.Errorf("Error reading index_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index_field: %v", err)
		}
	}

	if err = d.Set("log_type", flattenSystemSqlCustomIndexLogType2edl(o["log-type"], d, "log_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-type"], "SystemSqlCustomIndex-LogType"); ok {
			if err = d.Set("log_type", vv); err != nil {
				return fmt.Errorf("Error reading log_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_type: %v", err)
		}
	}

	return nil
}

func flattenSystemSqlCustomIndexFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSqlCustomIndexCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexDeviceType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexIndexField2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexLogType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSqlCustomIndex(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("case_sensitive"); ok || d.HasChange("case_sensitive") {
		t, err := expandSystemSqlCustomIndexCaseSensitive2edl(d, v, "case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("device_type"); ok || d.HasChange("device_type") {
		t, err := expandSystemSqlCustomIndexDeviceType2edl(d, v, "device_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSqlCustomIndexId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("index_field"); ok || d.HasChange("index_field") {
		t, err := expandSystemSqlCustomIndexIndexField2edl(d, v, "index_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index-field"] = t
		}
	}

	if v, ok := d.GetOk("log_type"); ok || d.HasChange("log_type") {
		t, err := expandSystemSqlCustomIndexLogType2edl(d, v, "log_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-type"] = t
		}
	}

	return &obj, nil
}
