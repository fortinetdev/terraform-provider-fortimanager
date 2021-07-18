// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure user password policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserPasswordPolicyCreate,
		Read:   resourceObjectUserPasswordPolicyRead,
		Update: resourceObjectUserPasswordPolicyUpdate,
		Delete: resourceObjectUserPasswordPolicyDelete,

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
			"expire_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"expired_password_renewal": &schema.Schema{
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
			"warn_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserPasswordPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserPasswordPolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPasswordPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserPasswordPolicyRead(d, m)
}

func resourceObjectUserPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserPasswordPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserPasswordPolicyRead(d, m)
}

func resourceObjectUserPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserPasswordPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserPasswordPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserPasswordPolicyExpireDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyExpiredPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyWarnDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("expire_days", flattenObjectUserPasswordPolicyExpireDays(o["expire-days"], d, "expire_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-days"], "ObjectUserPasswordPolicy-ExpireDays"); ok {
			if err = d.Set("expire_days", vv); err != nil {
				return fmt.Errorf("Error reading expire_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_days: %v", err)
		}
	}

	if err = d.Set("expired_password_renewal", flattenObjectUserPasswordPolicyExpiredPasswordRenewal(o["expired-password-renewal"], d, "expired_password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["expired-password-renewal"], "ObjectUserPasswordPolicy-ExpiredPasswordRenewal"); ok {
			if err = d.Set("expired_password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading expired_password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expired_password_renewal: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserPasswordPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserPasswordPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("warn_days", flattenObjectUserPasswordPolicyWarnDays(o["warn-days"], d, "warn_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["warn-days"], "ObjectUserPasswordPolicy-WarnDays"); ok {
			if err = d.Set("warn_days", vv); err != nil {
				return fmt.Errorf("Error reading warn_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading warn_days: %v", err)
		}
	}

	return nil
}

func flattenObjectUserPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserPasswordPolicyExpireDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyExpiredPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyWarnDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("expire_days"); ok {
		t, err := expandObjectUserPasswordPolicyExpireDays(d, v, "expire_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-days"] = t
		}
	}

	if v, ok := d.GetOk("expired_password_renewal"); ok {
		t, err := expandObjectUserPasswordPolicyExpiredPasswordRenewal(d, v, "expired_password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserPasswordPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("warn_days"); ok {
		t, err := expandObjectUserPasswordPolicyWarnDays(d, v, "warn_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-days"] = t
		}
	}

	return &obj, nil
}
