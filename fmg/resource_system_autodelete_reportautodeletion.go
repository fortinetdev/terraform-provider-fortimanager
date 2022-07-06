// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automatic deletion policy for reports.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutoDeleteReportAutoDeletion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoDeleteReportAutoDeletionUpdate,
		Read:   resourceSystemAutoDeleteReportAutoDeletionRead,
		Update: resourceSystemAutoDeleteReportAutoDeletionUpdate,
		Delete: resourceSystemAutoDeleteReportAutoDeletionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"retention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"runat": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutoDeleteReportAutoDeletionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAutoDeleteReportAutoDeletion(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteReportAutoDeletion resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoDeleteReportAutoDeletion(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteReportAutoDeletion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoDeleteReportAutoDeletion")

	return resourceSystemAutoDeleteReportAutoDeletionRead(d, m)
}

func resourceSystemAutoDeleteReportAutoDeletionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAutoDeleteReportAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoDeleteReportAutoDeletion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoDeleteReportAutoDeletionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAutoDeleteReportAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteReportAutoDeletion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoDeleteReportAutoDeletion(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteReportAutoDeletion resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoDeleteReportAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDeleteReportAutoDeletion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("retention", flattenSystemAutoDeleteReportAutoDeletionRetention(o["retention"], d, "retention")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention"], "SystemAutoDeleteReportAutoDeletion-Retention"); ok {
			if err = d.Set("retention", vv); err != nil {
				return fmt.Errorf("Error reading retention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention: %v", err)
		}
	}

	if err = d.Set("runat", flattenSystemAutoDeleteReportAutoDeletionRunat(o["runat"], d, "runat")); err != nil {
		if vv, ok := fortiAPIPatch(o["runat"], "SystemAutoDeleteReportAutoDeletion-Runat"); ok {
			if err = d.Set("runat", vv); err != nil {
				return fmt.Errorf("Error reading runat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading runat: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutoDeleteReportAutoDeletionStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutoDeleteReportAutoDeletion-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutoDeleteReportAutoDeletionValue(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutoDeleteReportAutoDeletion-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoDeleteReportAutoDeletionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoDeleteReportAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDeleteReportAutoDeletion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("retention"); ok || d.HasChange("retention") {
		t, err := expandSystemAutoDeleteReportAutoDeletionRetention(d, v, "retention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention"] = t
		}
	}

	if v, ok := d.GetOk("runat"); ok || d.HasChange("runat") {
		t, err := expandSystemAutoDeleteReportAutoDeletionRunat(d, v, "runat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["runat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutoDeleteReportAutoDeletionStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutoDeleteReportAutoDeletionValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
