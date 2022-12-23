// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure strict/loose server override.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateServerOverrideStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateServerOverrideStatusUpdate,
		Read:   resourceFmupdateServerOverrideStatusRead,
		Update: resourceFmupdateServerOverrideStatusUpdate,
		Delete: resourceFmupdateServerOverrideStatusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateServerOverrideStatusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateServerOverrideStatus(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateServerOverrideStatus resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateServerOverrideStatus(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateServerOverrideStatus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateServerOverrideStatus")

	return resourceFmupdateServerOverrideStatusRead(d, m)
}

func resourceFmupdateServerOverrideStatusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteFmupdateServerOverrideStatus(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateServerOverrideStatus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateServerOverrideStatusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateServerOverrideStatus(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateServerOverrideStatus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateServerOverrideStatus(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateServerOverrideStatus resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateServerOverrideStatusMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateServerOverrideStatus(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mode", flattenFmupdateServerOverrideStatusMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "FmupdateServerOverrideStatus-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func flattenFmupdateServerOverrideStatusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateServerOverrideStatusMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateServerOverrideStatus(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandFmupdateServerOverrideStatusMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	return &obj, nil
}
