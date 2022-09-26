// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Password policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyUpdate,
		Read:   resourceSystemPasswordPolicyRead,
		Update: resourceSystemPasswordPolicyUpdate,
		Delete: resourceSystemPasswordPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"change_4_characters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"must_contain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystemPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPasswordPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemPasswordPolicy")

	return resourceSystemPasswordPolicyRead(d, m)
}

func resourceSystemPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemPasswordPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemPasswordPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemPasswordPolicyChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMustContain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPasswordPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("change_4_characters", flattenSystemPasswordPolicyChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["change-4-characters"], "SystemPasswordPolicy-Change4Characters"); ok {
			if err = d.Set("change_4_characters", vv); err != nil {
				return fmt.Errorf("Error reading change_4_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire", flattenSystemPasswordPolicyExpire(o["expire"], d, "expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire"], "SystemPasswordPolicy-Expire"); ok {
			if err = d.Set("expire", vv); err != nil {
				return fmt.Errorf("Error reading expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenSystemPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-length"], "SystemPasswordPolicy-MinimumLength"); ok {
			if err = d.Set("minimum_length", vv); err != nil {
				return fmt.Errorf("Error reading minimum_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("must_contain", flattenSystemPasswordPolicyMustContain(o["must-contain"], d, "must_contain")); err != nil {
		if vv, ok := fortiAPIPatch(o["must-contain"], "SystemPasswordPolicy-MustContain"); ok {
			if err = d.Set("must_contain", vv); err != nil {
				return fmt.Errorf("Error reading must_contain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading must_contain: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemPasswordPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemPasswordPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPasswordPolicyChange4Characters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMustContain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPasswordPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("change_4_characters"); ok || d.HasChange("change_4_characters") {
		t, err := expandSystemPasswordPolicyChange4Characters(d, v, "change_4_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-4-characters"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok || d.HasChange("expire") {
		t, err := expandSystemPasswordPolicyExpire(d, v, "expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok || d.HasChange("minimum_length") {
		t, err := expandSystemPasswordPolicyMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOk("must_contain"); ok || d.HasChange("must_contain") {
		t, err := expandSystemPasswordPolicyMustContain(d, v, "must_contain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["must-contain"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemPasswordPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
