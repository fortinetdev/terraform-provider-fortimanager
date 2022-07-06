// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure disk space available for use by the Upgrade Manager.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateDiskQuota() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateDiskQuotaUpdate,
		Read:   resourceFmupdateDiskQuotaRead,
		Update: resourceFmupdateDiskQuotaUpdate,
		Delete: resourceFmupdateDiskQuotaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateDiskQuotaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateDiskQuota(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateDiskQuota resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateDiskQuota(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateDiskQuota resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateDiskQuota")

	return resourceFmupdateDiskQuotaRead(d, m)
}

func resourceFmupdateDiskQuotaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateDiskQuota(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateDiskQuota resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateDiskQuotaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateDiskQuota(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateDiskQuota resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateDiskQuota(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateDiskQuota resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateDiskQuotaValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateDiskQuota(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("value", flattenFmupdateDiskQuotaValue(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "FmupdateDiskQuota-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenFmupdateDiskQuotaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateDiskQuotaValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateDiskQuota(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandFmupdateDiskQuotaValue(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
