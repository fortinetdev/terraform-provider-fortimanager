// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU DSW Queue DTS profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpuDswQueueDtsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuDswQueueDtsProfileCreate,
		Read:   resourceObjectSystemNpuDswQueueDtsProfileRead,
		Update: resourceObjectSystemNpuDswQueueDtsProfileUpdate,
		Delete: resourceObjectSystemNpuDswQueueDtsProfileDelete,

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
			"iport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"oport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"queue_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuDswQueueDtsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuDswQueueDtsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuDswQueueDtsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuDswQueueDtsProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuDswQueueDtsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuDswQueueDtsProfileRead(d, m)
}

func resourceObjectSystemNpuDswQueueDtsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuDswQueueDtsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuDswQueueDtsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuDswQueueDtsProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuDswQueueDtsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuDswQueueDtsProfileRead(d, m)
}

func resourceObjectSystemNpuDswQueueDtsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuDswQueueDtsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuDswQueueDtsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuDswQueueDtsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuDswQueueDtsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuDswQueueDtsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuDswQueueDtsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuDswQueueDtsProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuDswQueueDtsProfileIport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileOport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswQueueDtsProfileQueueSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuDswQueueDtsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("iport", flattenObjectSystemNpuDswQueueDtsProfileIport(o["iport"], d, "iport")); err != nil {
		if vv, ok := fortiAPIPatch(o["iport"], "ObjectSystemNpuDswQueueDtsProfile-Iport"); ok {
			if err = d.Set("iport", vv); err != nil {
				return fmt.Errorf("Error reading iport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iport: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemNpuDswQueueDtsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemNpuDswQueueDtsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("oport", flattenObjectSystemNpuDswQueueDtsProfileOport(o["oport"], d, "oport")); err != nil {
		if vv, ok := fortiAPIPatch(o["oport"], "ObjectSystemNpuDswQueueDtsProfile-Oport"); ok {
			if err = d.Set("oport", vv); err != nil {
				return fmt.Errorf("Error reading oport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oport: %v", err)
		}
	}

	if err = d.Set("profile_id", flattenObjectSystemNpuDswQueueDtsProfileProfileId(o["profile-id"], d, "profile_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-id"], "ObjectSystemNpuDswQueueDtsProfile-ProfileId"); ok {
			if err = d.Set("profile_id", vv); err != nil {
				return fmt.Errorf("Error reading profile_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_id: %v", err)
		}
	}

	if err = d.Set("queue_select", flattenObjectSystemNpuDswQueueDtsProfileQueueSelect(o["queue-select"], d, "queue_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["queue-select"], "ObjectSystemNpuDswQueueDtsProfile-QueueSelect"); ok {
			if err = d.Set("queue_select", vv); err != nil {
				return fmt.Errorf("Error reading queue_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queue_select: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuDswQueueDtsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuDswQueueDtsProfileIport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileOport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswQueueDtsProfileQueueSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuDswQueueDtsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("iport"); ok {
		t, err := expandObjectSystemNpuDswQueueDtsProfileIport(d, v, "iport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iport"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSystemNpuDswQueueDtsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oport"); ok {
		t, err := expandObjectSystemNpuDswQueueDtsProfileOport(d, v, "oport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oport"] = t
		}
	}

	if v, ok := d.GetOk("profile_id"); ok {
		t, err := expandObjectSystemNpuDswQueueDtsProfileProfileId(d, v, "profile_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-id"] = t
		}
	}

	if v, ok := d.GetOk("queue_select"); ok {
		t, err := expandObjectSystemNpuDswQueueDtsProfileQueueSelect(d, v, "queue_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue-select"] = t
		}
	}

	return &obj, nil
}
