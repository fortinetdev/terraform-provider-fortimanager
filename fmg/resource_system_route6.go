// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Routing table configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemRoute6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemRoute6Create,
		Read:   resourceSystemRoute6Read,
		Update: resourceSystemRoute6Update,
		Delete: resourceSystemRoute6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prio": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemRoute6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemRoute6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemRoute6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemRoute6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemRoute6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "prio")))

	return resourceSystemRoute6Read(d, m)
}

func resourceSystemRoute6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemRoute6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemRoute6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemRoute6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemRoute6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "prio")))

	return resourceSystemRoute6Read(d, m)
}

func resourceSystemRoute6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemRoute6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemRoute6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemRoute6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemRoute6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemRoute6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemRoute6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemRoute6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemRoute6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemRoute6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemRoute6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemRoute6Prio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemRoute6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenSystemRoute6Device(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemRoute6-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemRoute6Dst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemRoute6-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemRoute6Gateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystemRoute6-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("prio", flattenSystemRoute6Prio(o["prio"], d, "prio")); err != nil {
		if vv, ok := fortiAPIPatch(o["prio"], "SystemRoute6-Prio"); ok {
			if err = d.Set("prio", vv); err != nil {
				return fmt.Errorf("Error reading prio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prio: %v", err)
		}
	}

	return nil
}

func flattenSystemRoute6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemRoute6Device(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemRoute6Dst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemRoute6Gateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemRoute6Prio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemRoute6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemRoute6Device(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemRoute6Dst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystemRoute6Gateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("prio"); ok || d.HasChange("prio") {
		t, err := expandSystemRoute6Prio(d, v, "prio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prio"] = t
		}
	}

	return &obj, nil
}
