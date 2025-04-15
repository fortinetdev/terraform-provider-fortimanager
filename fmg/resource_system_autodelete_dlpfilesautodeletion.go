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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

func resourceSystemAutoDeleteDlpFilesAutoDeletionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAutoDeleteDlpFilesAutoDeletion(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDeleteDlpFilesAutoDeletion resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAutoDeleteDlpFilesAutoDeletion(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemAutoDeleteDlpFilesAutoDeletion(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemAutoDeleteDlpFilesAutoDeletion(mkey, paradict)
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

func flattenSystemAutoDeleteDlpFilesAutoDeletionRetention2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRunat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDeleteDlpFilesAutoDeletion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("retention", flattenSystemAutoDeleteDlpFilesAutoDeletionRetention2edl(o["retention"], d, "retention")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention"], "SystemAutoDeleteDlpFilesAutoDeletion-Retention"); ok {
			if err = d.Set("retention", vv); err != nil {
				return fmt.Errorf("Error reading retention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention: %v", err)
		}
	}

	if err = d.Set("runat", flattenSystemAutoDeleteDlpFilesAutoDeletionRunat2edl(o["runat"], d, "runat")); err != nil {
		if vv, ok := fortiAPIPatch(o["runat"], "SystemAutoDeleteDlpFilesAutoDeletion-Runat"); ok {
			if err = d.Set("runat", vv); err != nil {
				return fmt.Errorf("Error reading runat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading runat: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutoDeleteDlpFilesAutoDeletionStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemAutoDeleteDlpFilesAutoDeletion-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("value", flattenSystemAutoDeleteDlpFilesAutoDeletionValue2edl(o["value"], d, "value")); err != nil {
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

func expandSystemAutoDeleteDlpFilesAutoDeletionRetention2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRunat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDeleteDlpFilesAutoDeletion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("retention"); ok || d.HasChange("retention") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionRetention2edl(d, v, "retention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention"] = t
		}
	}

	if v, ok := d.GetOk("runat"); ok || d.HasChange("runat") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionRunat2edl(d, v, "runat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["runat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
