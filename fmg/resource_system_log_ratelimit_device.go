// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device log rate limit.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogRatelimitDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogRatelimitDeviceCreate,
		Read:   resourceSystemLogRatelimitDeviceRead,
		Update: resourceSystemLogRatelimitDeviceUpdate,
		Delete: resourceSystemLogRatelimitDeviceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ratelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemLogRatelimitDeviceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogRatelimitDevice(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogRatelimitDevice resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLogRatelimitDevice(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemLogRatelimitDevice resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogRatelimitDeviceRead(d, m)
}

func resourceSystemLogRatelimitDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogRatelimitDevice(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimitDevice resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogRatelimitDevice(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimitDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogRatelimitDeviceRead(d, m)
}

func resourceSystemLogRatelimitDeviceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLogRatelimitDevice(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogRatelimitDevice resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogRatelimitDeviceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogRatelimitDevice(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimitDevice resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogRatelimitDevice(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimitDevice resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogRatelimitDeviceDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceRatelimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogRatelimitDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenSystemLogRatelimitDeviceDevice2edl(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemLogRatelimitDevice-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenSystemLogRatelimitDeviceFilterType2edl(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "SystemLogRatelimitDevice-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogRatelimitDeviceId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogRatelimitDevice-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ratelimit", flattenSystemLogRatelimitDeviceRatelimit2edl(o["ratelimit"], d, "ratelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["ratelimit"], "SystemLogRatelimitDevice-Ratelimit"); ok {
			if err = d.Set("ratelimit", vv); err != nil {
				return fmt.Errorf("Error reading ratelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ratelimit: %v", err)
		}
	}

	return nil
}

func flattenSystemLogRatelimitDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogRatelimitDeviceDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceRatelimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogRatelimitDevice(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemLogRatelimitDeviceDevice2edl(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandSystemLogRatelimitDeviceFilterType2edl(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLogRatelimitDeviceId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ratelimit"); ok || d.HasChange("ratelimit") {
		t, err := expandSystemLogRatelimitDeviceRatelimit2edl(d, v, "ratelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ratelimit"] = t
		}
	}

	return &obj, nil
}
