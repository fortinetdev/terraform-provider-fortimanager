// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automatic deletion policy for device logs.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutoDeleteLogAutoDeletion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoDeleteLogAutoDeletionUpdate,
		Read:   resourceSystemAutoDeleteLogAutoDeletionRead,
		Update: resourceSystemAutoDeleteLogAutoDeletionUpdate,
		Delete: resourceSystemAutoDeleteLogAutoDeletionDelete,

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

func resourceSystemAutoDeleteLogAutoDeletionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAutoDeleteLogAutoDeletion(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteLogAutoDeletion resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoDeleteLogAutoDeletion(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteLogAutoDeletion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoDeleteLogAutoDeletion")

	return resourceSystemAutoDeleteLogAutoDeletionRead(d, m)
}

func resourceSystemAutoDeleteLogAutoDeletionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAutoDeleteLogAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoDeleteLogAutoDeletion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoDeleteLogAutoDeletionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAutoDeleteLogAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteLogAutoDeletion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoDeleteLogAutoDeletion(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteLogAutoDeletion resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoDeleteLogAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDeleteLogAutoDeletion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("retention", flattenSystemAutoDeleteLogAutoDeletionRetention(o["retention"], d, "retention")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention"], "SystemAutoDeleteLogAutoDeletion-Retention"); ok {
			if err = d.Set("retention", vv); err != nil {
				return fmt.Errorf("Error reading retention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention: %v", err)
		}
	}

	if err = d.Set("runat", flattenSystemAutoDeleteLogAutoDeletionRunat(o["runat"], d, "runat")); err != nil {
		if vv, ok := fortiAPIPatch(o["runat"], "SystemAutoDeleteLogAutoDeletion-Runat"); ok {
			if err = d.Set("runat", vv); err != nil {
				return fmt.Errorf("Error reading runat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading runat: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutoDeleteLogAutoDeletionStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutoDeleteLogAutoDeletion-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutoDeleteLogAutoDeletionValue(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutoDeleteLogAutoDeletion-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoDeleteLogAutoDeletionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoDeleteLogAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDeleteLogAutoDeletion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("retention"); ok || d.HasChange("retention") {
		t, err := expandSystemAutoDeleteLogAutoDeletionRetention(d, v, "retention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention"] = t
		}
	}

	if v, ok := d.GetOk("runat"); ok || d.HasChange("runat") {
		t, err := expandSystemAutoDeleteLogAutoDeletionRunat(d, v, "runat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["runat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutoDeleteLogAutoDeletionStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutoDeleteLogAutoDeletionValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
