// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automatic deletion policy for quarantined files.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoDeleteQuarantineFilesAutoDeletion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoDeleteQuarantineFilesAutoDeletionUpdate,
		Read:   resourceSystemAutoDeleteQuarantineFilesAutoDeletionRead,
		Update: resourceSystemAutoDeleteQuarantineFilesAutoDeletionUpdate,
		Delete: resourceSystemAutoDeleteQuarantineFilesAutoDeletionDelete,

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

func resourceSystemAutoDeleteQuarantineFilesAutoDeletionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAutoDeleteQuarantineFilesAutoDeletion(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteQuarantineFilesAutoDeletion resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoDeleteQuarantineFilesAutoDeletion(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteQuarantineFilesAutoDeletion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoDeleteQuarantineFilesAutoDeletion")

	return resourceSystemAutoDeleteQuarantineFilesAutoDeletionRead(d, m)
}

func resourceSystemAutoDeleteQuarantineFilesAutoDeletionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAutoDeleteQuarantineFilesAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoDeleteQuarantineFilesAutoDeletion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoDeleteQuarantineFilesAutoDeletionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAutoDeleteQuarantineFilesAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteQuarantineFilesAutoDeletion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoDeleteQuarantineFilesAutoDeletion(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteQuarantineFilesAutoDeletion resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDeleteQuarantineFilesAutoDeletion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("retention", flattenSystemAutoDeleteQuarantineFilesAutoDeletionRetention(o["retention"], d, "retention")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention"], "SystemAutoDeleteQuarantineFilesAutoDeletion-Retention"); ok {
			if err = d.Set("retention", vv); err != nil {
				return fmt.Errorf("Error reading retention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention: %v", err)
		}
	}

	if err = d.Set("runat", flattenSystemAutoDeleteQuarantineFilesAutoDeletionRunat(o["runat"], d, "runat")); err != nil {
		if vv, ok := fortiAPIPatch(o["runat"], "SystemAutoDeleteQuarantineFilesAutoDeletion-Runat"); ok {
			if err = d.Set("runat", vv); err != nil {
				return fmt.Errorf("Error reading runat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading runat: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutoDeleteQuarantineFilesAutoDeletionStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutoDeleteQuarantineFilesAutoDeletion-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutoDeleteQuarantineFilesAutoDeletionValue(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutoDeleteQuarantineFilesAutoDeletion-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDeleteQuarantineFilesAutoDeletion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("retention"); ok || d.HasChange("retention") {
		t, err := expandSystemAutoDeleteQuarantineFilesAutoDeletionRetention(d, v, "retention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention"] = t
		}
	}

	if v, ok := d.GetOk("runat"); ok || d.HasChange("runat") {
		t, err := expandSystemAutoDeleteQuarantineFilesAutoDeletionRunat(d, v, "runat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["runat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutoDeleteQuarantineFilesAutoDeletionStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutoDeleteQuarantineFilesAutoDeletionValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
