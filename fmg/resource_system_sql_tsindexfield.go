// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of SQL text search index fields.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSqlTsIndexField() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSqlTsIndexFieldCreate,
		Read:   resourceSystemSqlTsIndexFieldRead,
		Update: resourceSystemSqlTsIndexFieldUpdate,
		Delete: resourceSystemSqlTsIndexFieldDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSqlTsIndexFieldCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSqlTsIndexField(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSqlTsIndexField resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSqlTsIndexField(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemSqlTsIndexField resource: %v", err)
	}

	d.SetId(getStringKey(d, "category"))

	return resourceSystemSqlTsIndexFieldRead(d, m)
}

func resourceSystemSqlTsIndexFieldUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSqlTsIndexField(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSqlTsIndexField resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSqlTsIndexField(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSqlTsIndexField resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "category"))

	return resourceSystemSqlTsIndexFieldRead(d, m)
}

func resourceSystemSqlTsIndexFieldDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSqlTsIndexField(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSqlTsIndexField resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSqlTsIndexFieldRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSqlTsIndexField(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSqlTsIndexField resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSqlTsIndexField(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSqlTsIndexField resource from API: %v", err)
	}
	return nil
}

func flattenSystemSqlTsIndexFieldCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTsIndexFieldValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSqlTsIndexField(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenSystemSqlTsIndexFieldCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "SystemSqlTsIndexField-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemSqlTsIndexFieldValue(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemSqlTsIndexField-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemSqlTsIndexFieldFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSqlTsIndexFieldCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTsIndexFieldValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSqlTsIndexField(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok {
		t, err := expandSystemSqlTsIndexFieldCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok {
		t, err := expandSystemSqlTsIndexFieldValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
