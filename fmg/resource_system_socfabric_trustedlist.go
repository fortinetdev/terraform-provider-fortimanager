// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Pre-authorized security fabric nodes

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSocFabricTrustedList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocFabricTrustedListCreate,
		Read:   resourceSystemSocFabricTrustedListRead,
		Update: resourceSystemSocFabricTrustedListUpdate,
		Delete: resourceSystemSocFabricTrustedListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSocFabricTrustedListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSocFabricTrustedList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSocFabricTrustedList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSocFabricTrustedList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSocFabricTrustedList resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemSocFabricTrustedListRead(d, m)
}

func resourceSystemSocFabricTrustedListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSocFabricTrustedList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabricTrustedList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocFabricTrustedList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabricTrustedList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemSocFabricTrustedListRead(d, m)
}

func resourceSystemSocFabricTrustedListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSocFabricTrustedList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocFabricTrustedList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocFabricTrustedListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSocFabricTrustedList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabricTrustedList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocFabricTrustedList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabricTrustedList resource from API: %v", err)
	}
	return nil
}

func flattenSystemSocFabricTrustedListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricTrustedListSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSocFabricTrustedList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSocFabricTrustedListId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSocFabricTrustedList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemSocFabricTrustedListSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemSocFabricTrustedList-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	return nil
}

func flattenSystemSocFabricTrustedListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSocFabricTrustedListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricTrustedListSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSocFabricTrustedList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSocFabricTrustedListId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemSocFabricTrustedListSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	return &obj, nil
}
