// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of aditional SQL skip index fields.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSqlCustomSkipidx() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSqlCustomSkipidxCreate,
		Read:   resourceSystemSqlCustomSkipidxRead,
		Update: resourceSystemSqlCustomSkipidxUpdate,
		Delete: resourceSystemSqlCustomSkipidxDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSqlCustomSkipidxCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSqlCustomSkipidx(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSqlCustomSkipidx resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSqlCustomSkipidx(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemSqlCustomSkipidx resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemSqlCustomSkipidxRead(d, m)
}

func resourceSystemSqlCustomSkipidxUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSqlCustomSkipidx(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSqlCustomSkipidx resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSqlCustomSkipidx(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSqlCustomSkipidx resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemSqlCustomSkipidxRead(d, m)
}

func resourceSystemSqlCustomSkipidxDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSqlCustomSkipidx(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSqlCustomSkipidx resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSqlCustomSkipidxRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSqlCustomSkipidx(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSqlCustomSkipidx resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSqlCustomSkipidx(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSqlCustomSkipidx resource from API: %v", err)
	}
	return nil
}

func flattenSystemSqlCustomSkipidxDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "FortiGate",
			1:  "FortiManager",
			3:  "FortiClient",
			4:  "FortiMail",
			5:  "FortiWeb",
			10: "FortiSandbox",
			15: "FortiProxy",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemSqlCustomSkipidxId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxIndexField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxLogType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "app-ctrl",
			1:  "attack",
			2:  "content",
			3:  "dlp",
			4:  "emailfilter",
			5:  "event",
			6:  "generic",
			7:  "history",
			10: "traffic",
			11: "virus",
			12: "voip",
			13: "webfilter",
			14: "netscan",
			15: "fct-event",
			16: "fct-traffic",
			17: "fct-netscan",
			18: "waf",
			19: "gtp",
			20: "dns",
			21: "ssh",
			22: "ssl",
			23: "file-filter",
			24: "asset",
			25: "protocol",
			26: "siem",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemSqlCustomSkipidx(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device_type", flattenSystemSqlCustomSkipidxDeviceType(o["device-type"], d, "device_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-type"], "SystemSqlCustomSkipidx-DeviceType"); ok {
			if err = d.Set("device_type", vv); err != nil {
				return fmt.Errorf("Error reading device_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemSqlCustomSkipidxId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSqlCustomSkipidx-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("index_field", flattenSystemSqlCustomSkipidxIndexField(o["index-field"], d, "index_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["index-field"], "SystemSqlCustomSkipidx-IndexField"); ok {
			if err = d.Set("index_field", vv); err != nil {
				return fmt.Errorf("Error reading index_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index_field: %v", err)
		}
	}

	if err = d.Set("log_type", flattenSystemSqlCustomSkipidxLogType(o["log-type"], d, "log_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-type"], "SystemSqlCustomSkipidx-LogType"); ok {
			if err = d.Set("log_type", vv); err != nil {
				return fmt.Errorf("Error reading log_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_type: %v", err)
		}
	}

	return nil
}

func flattenSystemSqlCustomSkipidxFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSqlCustomSkipidxDeviceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxIndexField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxLogType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSqlCustomSkipidx(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device_type"); ok {
		t, err := expandSystemSqlCustomSkipidxDeviceType(d, v, "device_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemSqlCustomSkipidxId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("index_field"); ok {
		t, err := expandSystemSqlCustomSkipidxIndexField(d, v, "index_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index-field"] = t
		}
	}

	if v, ok := d.GetOk("log_type"); ok {
		t, err := expandSystemSqlCustomSkipidxLogType(d, v, "log_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-type"] = t
		}
	}

	return &obj, nil
}
