// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectSystem MetaSysMetaFields

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemMetaSysMetaFields() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemMetaSysMetaFieldsCreate,
		Read:   resourceObjectSystemMetaSysMetaFieldsRead,
		Update: resourceObjectSystemMetaSysMetaFieldsUpdate,
		Delete: resourceObjectSystemMetaSysMetaFieldsDelete,

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
			"meta": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fieldlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"importance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemMetaSysMetaFieldsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	meta := d.Get("meta").(string)
	paradict["meta"] = meta

	obj, err := getObjectObjectSystemMetaSysMetaFields(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemMetaSysMetaFields resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemMetaSysMetaFields(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemMetaSysMetaFields resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemMetaSysMetaFieldsRead(d, m)
}

func resourceObjectSystemMetaSysMetaFieldsUpdate(d *schema.ResourceData, m interface{}) error {
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

	meta := d.Get("meta").(string)
	paradict["meta"] = meta

	obj, err := getObjectObjectSystemMetaSysMetaFields(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemMetaSysMetaFields resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemMetaSysMetaFields(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemMetaSysMetaFields resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemMetaSysMetaFieldsRead(d, m)
}

func resourceObjectSystemMetaSysMetaFieldsDelete(d *schema.ResourceData, m interface{}) error {
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

	meta := d.Get("meta").(string)
	paradict["meta"] = meta

	err = c.DeleteObjectSystemMetaSysMetaFields(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemMetaSysMetaFields resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemMetaSysMetaFieldsRead(d *schema.ResourceData, m interface{}) error {
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

	meta := d.Get("meta").(string)
	if meta == "" {
		meta = importOptionChecking(m.(*FortiClient).Cfg, "meta")
		if meta == "" {
			return fmt.Errorf("Parameter meta is missing")
		}
		if err = d.Set("meta", meta); err != nil {
			return fmt.Errorf("Error set params meta: %v", err)
		}
	}
	paradict["meta"] = meta

	o, err := c.ReadObjectSystemMetaSysMetaFields(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemMetaSysMetaFields resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemMetaSysMetaFields(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemMetaSysMetaFields resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemMetaSysMetaFieldsFieldlength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemMetaSysMetaFieldsImportance2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemMetaSysMetaFieldsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemMetaSysMetaFields(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fieldlength", flattenObjectSystemMetaSysMetaFieldsFieldlength2edl(o["fieldlength"], d, "fieldlength")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldlength"], "ObjectSystemMetaSysMetaFields-Fieldlength"); ok {
			if err = d.Set("fieldlength", vv); err != nil {
				return fmt.Errorf("Error reading fieldlength: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldlength: %v", err)
		}
	}

	if err = d.Set("importance", flattenObjectSystemMetaSysMetaFieldsImportance2edl(o["importance"], d, "importance")); err != nil {
		if vv, ok := fortiAPIPatch(o["importance"], "ObjectSystemMetaSysMetaFields-Importance"); ok {
			if err = d.Set("importance", vv); err != nil {
				return fmt.Errorf("Error reading importance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading importance: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemMetaSysMetaFieldsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemMetaSysMetaFields-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemMetaSysMetaFieldsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemMetaSysMetaFieldsFieldlength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemMetaSysMetaFieldsImportance2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemMetaSysMetaFieldsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemMetaSysMetaFields(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fieldlength"); ok || d.HasChange("fieldlength") {
		t, err := expandObjectSystemMetaSysMetaFieldsFieldlength2edl(d, v, "fieldlength")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldlength"] = t
		}
	}

	if v, ok := d.GetOk("importance"); ok || d.HasChange("importance") {
		t, err := expandObjectSystemMetaSysMetaFieldsImportance2edl(d, v, "importance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["importance"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemMetaSysMetaFieldsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
