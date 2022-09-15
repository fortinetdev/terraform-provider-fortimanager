// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure driver HA scan for SSE.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuSseHaScan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuSseHaScanUpdate,
		Read:   resourceObjectSystemNpuSseHaScanRead,
		Update: resourceObjectSystemNpuSseHaScanUpdate,
		Delete: resourceObjectSystemNpuSseHaScanDelete,

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
			"gap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_session_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuSseHaScanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuSseHaScan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuSseHaScan resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuSseHaScan(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuSseHaScan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuSseHaScan")

	return resourceObjectSystemNpuSseHaScanRead(d, m)
}

func resourceObjectSystemNpuSseHaScanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuSseHaScan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuSseHaScan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuSseHaScanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuSseHaScan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuSseHaScan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuSseHaScan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuSseHaScan resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuSseHaScanGap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseHaScanMaxSessionCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSseHaScanMinDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuSseHaScan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("gap", flattenObjectSystemNpuSseHaScanGap(o["gap"], d, "gap")); err != nil {
		if vv, ok := fortiAPIPatch(o["gap"], "ObjectSystemNpuSseHaScan-Gap"); ok {
			if err = d.Set("gap", vv); err != nil {
				return fmt.Errorf("Error reading gap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gap: %v", err)
		}
	}

	if err = d.Set("max_session_cnt", flattenObjectSystemNpuSseHaScanMaxSessionCnt(o["max-session-cnt"], d, "max_session_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-session-cnt"], "ObjectSystemNpuSseHaScan-MaxSessionCnt"); ok {
			if err = d.Set("max_session_cnt", vv); err != nil {
				return fmt.Errorf("Error reading max_session_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_session_cnt: %v", err)
		}
	}

	if err = d.Set("min_duration", flattenObjectSystemNpuSseHaScanMinDuration(o["min-duration"], d, "min_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-duration"], "ObjectSystemNpuSseHaScan-MinDuration"); ok {
			if err = d.Set("min_duration", vv); err != nil {
				return fmt.Errorf("Error reading min_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_duration: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuSseHaScanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuSseHaScanGap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseHaScanMaxSessionCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSseHaScanMinDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuSseHaScan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("gap"); ok || d.HasChange("gap") {
		t, err := expandObjectSystemNpuSseHaScanGap(d, v, "gap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gap"] = t
		}
	}

	if v, ok := d.GetOk("max_session_cnt"); ok || d.HasChange("max_session_cnt") {
		t, err := expandObjectSystemNpuSseHaScanMaxSessionCnt(d, v, "max_session_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-session-cnt"] = t
		}
	}

	if v, ok := d.GetOk("min_duration"); ok || d.HasChange("min_duration") {
		t, err := expandObjectSystemNpuSseHaScanMinDuration(d, v, "min_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-duration"] = t
		}
	}

	return &obj, nil
}
