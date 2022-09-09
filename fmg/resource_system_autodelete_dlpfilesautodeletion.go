// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automatic deletion policy for DLP archives.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutoDeleteDlpFilesAutoDeletion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoDeleteDlpFilesAutoDeletionUpdate,
		Read:   resourceSystemAutoDeleteDlpFilesAutoDeletionRead,
		Update: resourceSystemAutoDeleteDlpFilesAutoDeletionUpdate,
		Delete: resourceSystemAutoDeleteDlpFilesAutoDeletionDelete,

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
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemAutoDeleteDlpFilesAutoDeletionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAutoDeleteDlpFilesAutoDeletion(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteDlpFilesAutoDeletion resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoDeleteDlpFilesAutoDeletion(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteDlpFilesAutoDeletion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoDeleteDlpFilesAutoDeletion")

	return resourceSystemAutoDeleteDlpFilesAutoDeletionRead(d, m)
}

func resourceSystemAutoDeleteDlpFilesAutoDeletionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAutoDeleteDlpFilesAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoDeleteDlpFilesAutoDeletion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoDeleteDlpFilesAutoDeletionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAutoDeleteDlpFilesAutoDeletion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteDlpFilesAutoDeletion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoDeleteDlpFilesAutoDeletion(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDeleteDlpFilesAutoDeletion resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDeleteDlpFilesAutoDeletion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("retention", flattenSystemAutoDeleteDlpFilesAutoDeletionRetention(o["retention"], d, "retention")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention"], "SystemAutoDeleteDlpFilesAutoDeletion-Retention"); ok {
			if err = d.Set("retention", vv); err != nil {
				return fmt.Errorf("Error reading retention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention: %v", err)
		}
	}

	if err = d.Set("runat", flattenSystemAutoDeleteDlpFilesAutoDeletionRunat(o["runat"], d, "runat")); err != nil {
		if vv, ok := fortiAPIPatch(o["runat"], "SystemAutoDeleteDlpFilesAutoDeletion-Runat"); ok {
			if err = d.Set("runat", vv); err != nil {
				return fmt.Errorf("Error reading runat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading runat: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutoDeleteDlpFilesAutoDeletionStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutoDeleteDlpFilesAutoDeletion-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutoDeleteDlpFilesAutoDeletionValue(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "SystemAutoDeleteDlpFilesAutoDeletion-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDeleteDlpFilesAutoDeletion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("retention"); ok || d.HasChange("retention") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionRetention(d, v, "retention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention"] = t
		}
	}

	if v, ok := d.GetOk("runat"); ok || d.HasChange("runat") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionRunat(d, v, "runat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["runat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
