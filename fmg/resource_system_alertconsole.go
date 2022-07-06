// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Alert console.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAlertConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAlertConsoleUpdate,
		Read:   resourceSystemAlertConsoleRead,
		Update: resourceSystemAlertConsoleUpdate,
		Delete: resourceSystemAlertConsoleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity_level": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAlertConsoleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAlertConsole(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertConsole resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAlertConsole(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertConsole resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAlertConsole")

	return resourceSystemAlertConsoleRead(d, m)
}

func resourceSystemAlertConsoleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAlertConsole(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAlertConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAlertConsoleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAlertConsole(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertConsole resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlertConsole(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertConsole resource from API: %v", err)
	}
	return nil
}

func flattenSystemAlertConsolePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertConsoleSeverityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemAlertConsole(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("period", flattenSystemAlertConsolePeriod(o["period"], d, "period")); err != nil {
		if vv, ok := fortiAPIPatch(o["period"], "SystemAlertConsole-Period"); ok {
			if err = d.Set("period", vv); err != nil {
				return fmt.Errorf("Error reading period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading period: %v", err)
		}
	}

	if err = d.Set("severity_level", flattenSystemAlertConsoleSeverityLevel(o["severity-level"], d, "severity_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity-level"], "SystemAlertConsole-SeverityLevel"); ok {
			if err = d.Set("severity_level", vv); err != nil {
				return fmt.Errorf("Error reading severity_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity_level: %v", err)
		}
	}

	return nil
}

func flattenSystemAlertConsoleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAlertConsolePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertConsoleSeverityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemAlertConsole(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("period"); ok || d.HasChange("period") {
		t, err := expandSystemAlertConsolePeriod(d, v, "period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["period"] = t
		}
	}

	if v, ok := d.GetOk("severity_level"); ok || d.HasChange("severity_level") {
		t, err := expandSystemAlertConsoleSeverityLevel(d, v, "severity_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity-level"] = t
		}
	}

	return &obj, nil
}
