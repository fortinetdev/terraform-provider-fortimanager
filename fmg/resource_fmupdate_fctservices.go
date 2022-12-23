// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard to provide services to FortiClient installations.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFctServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFctServicesUpdate,
		Read:   resourceFmupdateFctServicesRead,
		Update: resourceFmupdateFctServicesUpdate,
		Delete: resourceFmupdateFctServicesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFctServicesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFctServices(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFctServices resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFctServices(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFctServices resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFctServices")

	return resourceFmupdateFctServicesRead(d, m)
}

func resourceFmupdateFctServicesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteFmupdateFctServices(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFctServices resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFctServicesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFctServices(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFctServices resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFctServices(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFctServices resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFctServicesPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFctServicesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFctServices(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("port", flattenFmupdateFctServicesPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateFctServices-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFmupdateFctServicesStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateFctServices-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFctServicesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFctServicesPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFctServicesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFctServices(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFmupdateFctServicesPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateFctServicesStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
