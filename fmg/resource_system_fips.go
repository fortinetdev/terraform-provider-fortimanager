// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for FIPS-CC mode.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFips() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFipsUpdate,
		Read:   resourceSystemFipsRead,
		Update: resourceSystemFipsUpdate,
		Delete: resourceSystemFipsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"entropy_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"re_seed_interval": &schema.Schema{
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

func resourceSystemFipsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemFips(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFips resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFips(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemFips resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFips")

	return resourceSystemFipsRead(d, m)
}

func resourceSystemFipsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemFips(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFips resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFipsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemFips(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFips resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFips(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFips resource from API: %v", err)
	}
	return nil
}

func flattenSystemFipsEntropyToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsReSeedInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFips(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("entropy_token", flattenSystemFipsEntropyToken(o["entropy-token"], d, "entropy_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["entropy-token"], "SystemFips-EntropyToken"); ok {
			if err = d.Set("entropy_token", vv); err != nil {
				return fmt.Errorf("Error reading entropy_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entropy_token: %v", err)
		}
	}

	if err = d.Set("re_seed_interval", flattenSystemFipsReSeedInterval(o["re-seed-interval"], d, "re_seed_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["re-seed-interval"], "SystemFips-ReSeedInterval"); ok {
			if err = d.Set("re_seed_interval", vv); err != nil {
				return fmt.Errorf("Error reading re_seed_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading re_seed_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFipsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFips-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemFipsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFipsEntropyToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsReSeedInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFips(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("entropy_token"); ok || d.HasChange("entropy_token") {
		t, err := expandSystemFipsEntropyToken(d, v, "entropy_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entropy-token"] = t
		}
	}

	if v, ok := d.GetOk("re_seed_interval"); ok || d.HasChange("re_seed_interval") {
		t, err := expandSystemFipsReSeedInterval(d, v, "re_seed_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["re-seed-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemFipsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
