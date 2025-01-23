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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expired_password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_change_characters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_lower_case_letter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_upper_case_letter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"reuse_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse_password_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserPasswordPolicy(obj, paradict)

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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserPasswordPolicy(obj, mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectUserPasswordPolicy(mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectUserPasswordPolicy(mkey, paradict)
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

func flattenObjectUserPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyExpiredPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyReusePasswordLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPasswordPolicyWarnDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("expire_days", flattenObjectUserPasswordPolicyExpireDays(o["expire-days"], d, "expire_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-days"], "ObjectUserPasswordPolicy-ExpireDays"); ok {
			if err = d.Set("expire_days", vv); err != nil {
				return fmt.Errorf("Error reading expire_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_days: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenObjectUserPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire-status"], "ObjectUserPasswordPolicy-ExpireStatus"); ok {
			if err = d.Set("expire_status", vv); err != nil {
				return fmt.Errorf("Error reading expire_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire_status: %v", err)
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

	if err = d.Set("min_change_characters", flattenObjectUserPasswordPolicyMinChangeCharacters(o["min-change-characters"], d, "min_change_characters")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-change-characters"], "ObjectUserPasswordPolicy-MinChangeCharacters"); ok {
			if err = d.Set("min_change_characters", vv); err != nil {
				return fmt.Errorf("Error reading min_change_characters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenObjectUserPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-lower-case-letter"], "ObjectUserPasswordPolicy-MinLowerCaseLetter"); ok {
			if err = d.Set("min_lower_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenObjectUserPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-non-alphanumeric"], "ObjectUserPasswordPolicy-MinNonAlphanumeric"); ok {
			if err = d.Set("min_non_alphanumeric", vv); err != nil {
				return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenObjectUserPasswordPolicyMinNumber(o["min-number"], d, "min_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-number"], "ObjectUserPasswordPolicy-MinNumber"); ok {
			if err = d.Set("min_number", vv); err != nil {
				return fmt.Errorf("Error reading min_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenObjectUserPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-upper-case-letter"], "ObjectUserPasswordPolicy-MinUpperCaseLetter"); ok {
			if err = d.Set("min_upper_case_letter", vv); err != nil {
				return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenObjectUserPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-length"], "ObjectUserPasswordPolicy-MinimumLength"); ok {
			if err = d.Set("minimum_length", vv); err != nil {
				return fmt.Errorf("Error reading minimum_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_length: %v", err)
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

	if err = d.Set("reuse_password", flattenObjectUserPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password"], "ObjectUserPasswordPolicy-ReusePassword"); ok {
			if err = d.Set("reuse_password", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	if err = d.Set("reuse_password_limit", flattenObjectUserPasswordPolicyReusePasswordLimit(o["reuse-password-limit"], d, "reuse_password_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["reuse-password-limit"], "ObjectUserPasswordPolicy-ReusePasswordLimit"); ok {
			if err = d.Set("reuse_password_limit", vv); err != nil {
				return fmt.Errorf("Error reading reuse_password_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reuse_password_limit: %v", err)
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

func expandObjectUserPasswordPolicyExpireStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyExpiredPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyMinChangeCharacters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyMinNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyReusePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyReusePasswordLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPasswordPolicyWarnDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("expire_days"); ok || d.HasChange("expire_days") {
		t, err := expandObjectUserPasswordPolicyExpireDays(d, v, "expire_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-days"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok || d.HasChange("expire_status") {
		t, err := expandObjectUserPasswordPolicyExpireStatus(d, v, "expire_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("expired_password_renewal"); ok || d.HasChange("expired_password_renewal") {
		t, err := expandObjectUserPasswordPolicyExpiredPasswordRenewal(d, v, "expired_password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("min_change_characters"); ok || d.HasChange("min_change_characters") {
		t, err := expandObjectUserPasswordPolicyMinChangeCharacters(d, v, "min_change_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-change-characters"] = t
		}
	}

	if v, ok := d.GetOk("min_lower_case_letter"); ok || d.HasChange("min_lower_case_letter") {
		t, err := expandObjectUserPasswordPolicyMinLowerCaseLetter(d, v, "min_lower_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("min_non_alphanumeric"); ok || d.HasChange("min_non_alphanumeric") {
		t, err := expandObjectUserPasswordPolicyMinNonAlphanumeric(d, v, "min_non_alphanumeric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOk("min_number"); ok || d.HasChange("min_number") {
		t, err := expandObjectUserPasswordPolicyMinNumber(d, v, "min_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOk("min_upper_case_letter"); ok || d.HasChange("min_upper_case_letter") {
		t, err := expandObjectUserPasswordPolicyMinUpperCaseLetter(d, v, "min_upper_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok || d.HasChange("minimum_length") {
		t, err := expandObjectUserPasswordPolicyMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserPasswordPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok || d.HasChange("reuse_password") {
		t, err := expandObjectUserPasswordPolicyReusePassword(d, v, "reuse_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password_limit"); ok || d.HasChange("reuse_password_limit") {
		t, err := expandObjectUserPasswordPolicyReusePasswordLimit(d, v, "reuse_password_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password-limit"] = t
		}
	}

	if v, ok := d.GetOk("warn_days"); ok || d.HasChange("warn_days") {
		t, err := expandObjectUserPasswordPolicyWarnDays(d, v, "warn_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-days"] = t
		}
	}

	return &obj, nil
}
