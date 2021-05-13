// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Log based alert settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogAlertUpdate,
		Read:   resourceSystemLogAlertRead,
		Update: resourceSystemLogAlertUpdate,
		Delete: resourceSystemLogAlertDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_alert_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogAlertUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogAlert(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogAlert resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogAlert(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogAlert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogAlert")

	return resourceSystemLogAlertRead(d, m)
}

func resourceSystemLogAlertDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogAlert(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogAlert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogAlertRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogAlert(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogAlert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogAlert(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogAlert resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogAlertMaxAlertCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogAlert(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_alert_count", flattenSystemLogAlertMaxAlertCount(o["max-alert-count"], d, "max_alert_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-alert-count"], "SystemLogAlert-MaxAlertCount"); ok {
			if err = d.Set("max_alert_count", vv); err != nil {
				return fmt.Errorf("Error reading max_alert_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_alert_count: %v", err)
		}
	}

	return nil
}

func flattenSystemLogAlertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogAlertMaxAlertCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogAlert(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_alert_count"); ok {
		t, err := expandSystemLogAlertMaxAlertCount(d, v, "max_alert_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-alert-count"] = t
		}
	}

	return &obj, nil
}
