// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Per device or ADOM log rate limits.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogRatelimitRatelimits() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogRatelimitRatelimitsCreate,
		Read:   resourceSystemLogRatelimitRatelimitsRead,
		Update: resourceSystemLogRatelimitRatelimitsUpdate,
		Delete: resourceSystemLogRatelimitRatelimitsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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

func resourceSystemLogRatelimitRatelimitsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogRatelimitRatelimits(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogRatelimitRatelimits resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemLogRatelimitRatelimits(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogRatelimitRatelimits resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogRatelimitRatelimitsRead(d, m)
}

func resourceSystemLogRatelimitRatelimitsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogRatelimitRatelimits(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimitRatelimits resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemLogRatelimitRatelimits(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimitRatelimits resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogRatelimitRatelimitsRead(d, m)
}

func resourceSystemLogRatelimitRatelimitsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemLogRatelimitRatelimits(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogRatelimitRatelimits resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogRatelimitRatelimitsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogRatelimitRatelimits(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimitRatelimits resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogRatelimitRatelimits(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimitRatelimits resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogRatelimitRatelimitsFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitRatelimitsRatelimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogRatelimitRatelimits(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("filter", flattenSystemLogRatelimitRatelimitsFilter2edl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "SystemLogRatelimitRatelimits-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenSystemLogRatelimitRatelimitsFilterType2edl(o["filter-type"], d, "filter_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-type"], "SystemLogRatelimitRatelimits-FilterType"); ok {
			if err = d.Set("filter_type", vv); err != nil {
				return fmt.Errorf("Error reading filter_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogRatelimitRatelimitsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogRatelimitRatelimits-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ratelimit", flattenSystemLogRatelimitRatelimitsRatelimit2edl(o["ratelimit"], d, "ratelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["ratelimit"], "SystemLogRatelimitRatelimits-Ratelimit"); ok {
			if err = d.Set("ratelimit", vv); err != nil {
				return fmt.Errorf("Error reading ratelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ratelimit: %v", err)
		}
	}

	return nil
}

func flattenSystemLogRatelimitRatelimitsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogRatelimitRatelimitsFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitRatelimitsRatelimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogRatelimitRatelimits(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandSystemLogRatelimitRatelimitsFilter2edl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok || d.HasChange("filter_type") {
		t, err := expandSystemLogRatelimitRatelimitsFilterType2edl(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLogRatelimitRatelimitsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ratelimit"); ok || d.HasChange("ratelimit") {
		t, err := expandSystemLogRatelimitRatelimitsRatelimit2edl(d, v, "ratelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ratelimit"] = t
		}
	}

	return &obj, nil
}
