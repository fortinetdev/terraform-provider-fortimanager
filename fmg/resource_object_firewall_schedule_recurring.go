// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Recurring schedule configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallScheduleRecurring() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallScheduleRecurringCreate,
		Read:   resourceObjectFirewallScheduleRecurringRead,
		Update: resourceObjectFirewallScheduleRecurringUpdate,
		Delete: resourceObjectFirewallScheduleRecurringDelete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallScheduleRecurringCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallScheduleRecurring(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallScheduleRecurring resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallScheduleRecurring(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallScheduleRecurring resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallScheduleRecurringRead(d, m)
}

func resourceObjectFirewallScheduleRecurringUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallScheduleRecurring(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallScheduleRecurring resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallScheduleRecurring(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallScheduleRecurring resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallScheduleRecurringRead(d, m)
}

func resourceObjectFirewallScheduleRecurringDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallScheduleRecurring(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallScheduleRecurring resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallScheduleRecurringRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallScheduleRecurring(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallScheduleRecurring resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallScheduleRecurring(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallScheduleRecurring resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallScheduleRecurringColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleRecurringDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallScheduleRecurringEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleRecurringFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleRecurringGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleRecurringName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleRecurringStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallScheduleRecurring(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("color", flattenObjectFirewallScheduleRecurringColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallScheduleRecurring-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("day", flattenObjectFirewallScheduleRecurringDay(o["day"], d, "day")); err != nil {
		if vv, ok := fortiAPIPatch(o["day"], "ObjectFirewallScheduleRecurring-Day"); ok {
			if err = d.Set("day", vv); err != nil {
				return fmt.Errorf("Error reading day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("end", flattenObjectFirewallScheduleRecurringEnd(o["end"], d, "end")); err != nil {
		if vv, ok := fortiAPIPatch(o["end"], "ObjectFirewallScheduleRecurring-End"); ok {
			if err = d.Set("end", vv); err != nil {
				return fmt.Errorf("Error reading end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectFirewallScheduleRecurringFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectFirewallScheduleRecurring-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("global_object", flattenObjectFirewallScheduleRecurringGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "ObjectFirewallScheduleRecurring-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallScheduleRecurringName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallScheduleRecurring-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenObjectFirewallScheduleRecurringStart(o["start"], d, "start")); err != nil {
		if vv, ok := fortiAPIPatch(o["start"], "ObjectFirewallScheduleRecurring-Start"); ok {
			if err = d.Set("start", vv); err != nil {
				return fmt.Errorf("Error reading start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallScheduleRecurringFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallScheduleRecurringColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleRecurringDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallScheduleRecurringEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleRecurringFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleRecurringGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleRecurringName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleRecurringStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallScheduleRecurring(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallScheduleRecurringColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok || d.HasChange("day") {
		t, err := expandObjectFirewallScheduleRecurringDay(d, v, "day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok || d.HasChange("end") {
		t, err := expandObjectFirewallScheduleRecurringEnd(d, v, "end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectFirewallScheduleRecurringFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandObjectFirewallScheduleRecurringGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallScheduleRecurringName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok || d.HasChange("start") {
		t, err := expandObjectFirewallScheduleRecurringStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	return &obj, nil
}
