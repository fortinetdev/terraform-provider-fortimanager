// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpFpSensitivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpFpSensitivityCreate,
		Read:   resourceObjectDlpFpSensitivityRead,
		Update: resourceObjectDlpFpSensitivityUpdate,
		Delete: resourceObjectDlpFpSensitivityDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectDlpFpSensitivityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDlpFpSensitivity(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpFpSensitivity resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpFpSensitivity(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpFpSensitivity resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpFpSensitivityRead(d, m)
}

func resourceObjectDlpFpSensitivityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDlpFpSensitivity(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpFpSensitivity resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpFpSensitivity(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpFpSensitivity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpFpSensitivityRead(d, m)
}

func resourceObjectDlpFpSensitivityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDlpFpSensitivity(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpFpSensitivity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpFpSensitivityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDlpFpSensitivity(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpFpSensitivity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpFpSensitivity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpFpSensitivity resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpFpSensitivityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpFpSensitivity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectDlpFpSensitivityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpFpSensitivity-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpFpSensitivityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpFpSensitivityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpFpSensitivity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDlpFpSensitivityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
