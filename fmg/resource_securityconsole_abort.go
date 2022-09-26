// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Abort and cancel a security console task.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityconsoleAbort() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleAbortUpdate,
		Read:   resourceSecurityconsoleAbortRead,
		Update: resourceSecurityconsoleAbortUpdate,
		Delete: resourceSecurityconsoleAbortDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSecurityconsoleAbortUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectSecurityconsoleAbort(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleAbort resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleAbort(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleAbort resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleAbort")

	return resourceSecurityconsoleAbortRead(d, m)
}

func resourceSecurityconsoleAbortDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleAbortRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleAbortAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSecurityconsoleAbort(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSecurityconsoleAbortAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsoleAbort-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	return nil
}

func flattenSecurityconsoleAbortFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleAbortAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityconsoleAbort(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSecurityconsoleAbortAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	return &obj, nil
}
