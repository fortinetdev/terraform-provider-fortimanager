// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU DSW DTS profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuDswDtsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuDswDtsProfileCreate,
		Read:   resourceObjectSystemNpuDswDtsProfileRead,
		Update: resourceObjectSystemNpuDswDtsProfileUpdate,
		Delete: resourceObjectSystemNpuDswDtsProfileDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuDswDtsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuDswDtsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuDswDtsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuDswDtsProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuDswDtsProfile resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "profile_id")))

	return resourceObjectSystemNpuDswDtsProfileRead(d, m)
}

func resourceObjectSystemNpuDswDtsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuDswDtsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuDswDtsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuDswDtsProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuDswDtsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "profile_id")))

	return resourceObjectSystemNpuDswDtsProfileRead(d, m)
}

func resourceObjectSystemNpuDswDtsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectSystemNpuDswDtsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuDswDtsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuDswDtsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSystemNpuDswDtsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuDswDtsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuDswDtsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuDswDtsProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuDswDtsProfileAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileMinLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDswDtsProfileStep(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuDswDtsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectSystemNpuDswDtsProfileAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectSystemNpuDswDtsProfile-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("min_limit", flattenObjectSystemNpuDswDtsProfileMinLimit(o["min-limit"], d, "min_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-limit"], "ObjectSystemNpuDswDtsProfile-MinLimit"); ok {
			if err = d.Set("min_limit", vv); err != nil {
				return fmt.Errorf("Error reading min_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_limit: %v", err)
		}
	}

	if err = d.Set("profile_id", flattenObjectSystemNpuDswDtsProfileProfileId(o["profile-id"], d, "profile_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-id"], "ObjectSystemNpuDswDtsProfile-ProfileId"); ok {
			if err = d.Set("profile_id", vv); err != nil {
				return fmt.Errorf("Error reading profile_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_id: %v", err)
		}
	}

	if err = d.Set("step", flattenObjectSystemNpuDswDtsProfileStep(o["step"], d, "step")); err != nil {
		if vv, ok := fortiAPIPatch(o["step"], "ObjectSystemNpuDswDtsProfile-Step"); ok {
			if err = d.Set("step", vv); err != nil {
				return fmt.Errorf("Error reading step: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading step: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuDswDtsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuDswDtsProfileAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileMinLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDswDtsProfileStep(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuDswDtsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectSystemNpuDswDtsProfileAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("min_limit"); ok || d.HasChange("min_limit") {
		t, err := expandObjectSystemNpuDswDtsProfileMinLimit(d, v, "min_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-limit"] = t
		}
	}

	if v, ok := d.GetOk("profile_id"); ok || d.HasChange("profile_id") {
		t, err := expandObjectSystemNpuDswDtsProfileProfileId(d, v, "profile_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-id"] = t
		}
	}

	if v, ok := d.GetOk("step"); ok || d.HasChange("step") {
		t, err := expandObjectSystemNpuDswDtsProfileStep(d, v, "step")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["step"] = t
		}
	}

	return &obj, nil
}
