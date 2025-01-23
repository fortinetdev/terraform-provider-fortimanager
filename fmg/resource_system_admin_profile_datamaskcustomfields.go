// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Customized datamask fields.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminProfileDatamaskCustomFields() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminProfileDatamaskCustomFieldsCreate,
		Read:   resourceSystemAdminProfileDatamaskCustomFieldsRead,
		Update: resourceSystemAdminProfileDatamaskCustomFieldsUpdate,
		Delete: resourceSystemAdminProfileDatamaskCustomFieldsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"field_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"field_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"field_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"field_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAdminProfileDatamaskCustomFieldsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectSystemAdminProfileDatamaskCustomFields(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfileDatamaskCustomFields resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdminProfileDatamaskCustomFields(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfileDatamaskCustomFields resource: %v", err)
	}

	d.SetId(getStringKey(d, "field_name"))

	return resourceSystemAdminProfileDatamaskCustomFieldsRead(d, m)
}

func resourceSystemAdminProfileDatamaskCustomFieldsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectSystemAdminProfileDatamaskCustomFields(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfileDatamaskCustomFields resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminProfileDatamaskCustomFields(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfileDatamaskCustomFields resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "field_name"))

	return resourceSystemAdminProfileDatamaskCustomFieldsRead(d, m)
}

func resourceSystemAdminProfileDatamaskCustomFieldsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteSystemAdminProfileDatamaskCustomFields(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminProfileDatamaskCustomFields resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminProfileDatamaskCustomFieldsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadSystemAdminProfileDatamaskCustomFields(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfileDatamaskCustomFields resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminProfileDatamaskCustomFields(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfileDatamaskCustomFields resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminProfileDatamaskCustomFieldsFieldType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminProfileDatamaskCustomFields(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("field_category", flattenSystemAdminProfileDatamaskCustomFieldsFieldCategory2edl(o["field-category"], d, "field_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["field-category"], "SystemAdminProfileDatamaskCustomFields-FieldCategory"); ok {
			if err = d.Set("field_category", vv); err != nil {
				return fmt.Errorf("Error reading field_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading field_category: %v", err)
		}
	}

	if err = d.Set("field_name", flattenSystemAdminProfileDatamaskCustomFieldsFieldName2edl(o["field-name"], d, "field_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["field-name"], "SystemAdminProfileDatamaskCustomFields-FieldName"); ok {
			if err = d.Set("field_name", vv); err != nil {
				return fmt.Errorf("Error reading field_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading field_name: %v", err)
		}
	}

	if err = d.Set("field_status", flattenSystemAdminProfileDatamaskCustomFieldsFieldStatus2edl(o["field-status"], d, "field_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["field-status"], "SystemAdminProfileDatamaskCustomFields-FieldStatus"); ok {
			if err = d.Set("field_status", vv); err != nil {
				return fmt.Errorf("Error reading field_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading field_status: %v", err)
		}
	}

	if err = d.Set("field_type", flattenSystemAdminProfileDatamaskCustomFieldsFieldType2edl(o["field-type"], d, "field_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["field-type"], "SystemAdminProfileDatamaskCustomFields-FieldType"); ok {
			if err = d.Set("field_type", vv); err != nil {
				return fmt.Errorf("Error reading field_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading field_type: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminProfileDatamaskCustomFieldsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminProfileDatamaskCustomFieldsFieldType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminProfileDatamaskCustomFields(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("field_category"); ok || d.HasChange("field_category") {
		t, err := expandSystemAdminProfileDatamaskCustomFieldsFieldCategory2edl(d, v, "field_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field-category"] = t
		}
	}

	if v, ok := d.GetOk("field_name"); ok || d.HasChange("field_name") {
		t, err := expandSystemAdminProfileDatamaskCustomFieldsFieldName2edl(d, v, "field_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field-name"] = t
		}
	}

	if v, ok := d.GetOk("field_status"); ok || d.HasChange("field_status") {
		t, err := expandSystemAdminProfileDatamaskCustomFieldsFieldStatus2edl(d, v, "field_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field-status"] = t
		}
	}

	if v, ok := d.GetOk("field_type"); ok || d.HasChange("field_type") {
		t, err := expandSystemAdminProfileDatamaskCustomFieldsFieldType2edl(d, v, "field_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field-type"] = t
		}
	}

	return &obj, nil
}
