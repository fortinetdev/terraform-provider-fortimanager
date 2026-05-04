// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiTelemetry predefined applications.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectTelemetryControllerApplicationPredefine() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectTelemetryControllerApplicationPredefineCreate,
		Read:   resourceObjectTelemetryControllerApplicationPredefineRead,
		Update: resourceObjectTelemetryControllerApplicationPredefineUpdate,
		Delete: resourceObjectTelemetryControllerApplicationPredefineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectTelemetryControllerApplicationPredefineCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectTelemetryControllerApplicationPredefine(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectTelemetryControllerApplicationPredefine resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("app_name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectTelemetryControllerApplicationPredefine(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectTelemetryControllerApplicationPredefine(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectTelemetryControllerApplicationPredefine resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectTelemetryControllerApplicationPredefine(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectTelemetryControllerApplicationPredefine resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "app_name"))

	return resourceObjectTelemetryControllerApplicationPredefineRead(d, m)
}

func resourceObjectTelemetryControllerApplicationPredefineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectTelemetryControllerApplicationPredefine(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerApplicationPredefine resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectTelemetryControllerApplicationPredefine(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerApplicationPredefine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "app_name"))

	return resourceObjectTelemetryControllerApplicationPredefineRead(d, m)
}

func resourceObjectTelemetryControllerApplicationPredefineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectTelemetryControllerApplicationPredefine(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectTelemetryControllerApplicationPredefine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectTelemetryControllerApplicationPredefineRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectTelemetryControllerApplicationPredefine(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectTelemetryControllerApplicationPredefine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectTelemetryControllerApplicationPredefine(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectTelemetryControllerApplicationPredefine resource from API: %v", err)
	}
	return nil
}

func flattenObjectTelemetryControllerApplicationPredefineAppName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerApplicationPredefineComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectTelemetryControllerApplicationPredefine(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("app_name", flattenObjectTelemetryControllerApplicationPredefineAppName(o["app-name"], d, "app_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-name"], "ObjectTelemetryControllerApplicationPredefine-AppName"); ok {
			if err = d.Set("app_name", vv); err != nil {
				return fmt.Errorf("Error reading app_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_name: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectTelemetryControllerApplicationPredefineComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectTelemetryControllerApplicationPredefine-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenObjectTelemetryControllerApplicationPredefineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectTelemetryControllerApplicationPredefineAppName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerApplicationPredefineComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectTelemetryControllerApplicationPredefine(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_name"); ok || d.HasChange("app_name") {
		t, err := expandObjectTelemetryControllerApplicationPredefineAppName(d, v, "app_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectTelemetryControllerApplicationPredefineComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
