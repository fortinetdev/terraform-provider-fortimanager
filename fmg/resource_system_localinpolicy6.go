// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 local in policy configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocalInPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocalInPolicy6Create,
		Read:   resourceSystemLocalInPolicy6Read,
		Update: resourceSystemLocalInPolicy6Update,
		Delete: resourceSystemLocalInPolicy6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocalInPolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy6 resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLocalInPolicy6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLocalInPolicy6Read(d, m)
}

func resourceSystemLocalInPolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocalInPolicy6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLocalInPolicy6Read(d, m)
}

func resourceSystemLocalInPolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLocalInPolicy6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocalInPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocalInPolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLocalInPolicy6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocalInPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocalInPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Dport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Intf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicy6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocalInPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemLocalInPolicy6Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemLocalInPolicy6-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("dport", flattenSystemLocalInPolicy6Dport(o["dport"], d, "dport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy6-Dport"); ok {
			if err = d.Set("dport", vv); err != nil {
				return fmt.Errorf("Error reading dport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dport: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemLocalInPolicy6Dst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy6-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLocalInPolicy6Id(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLocalInPolicy6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("intf", flattenSystemLocalInPolicy6Intf(o["intf"], d, "intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy6-Intf"); ok {
			if err = d.Set("intf", vv); err != nil {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemLocalInPolicy6Protocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemLocalInPolicy6-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("src", flattenSystemLocalInPolicy6Src(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy6-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	return nil
}

func flattenSystemLocalInPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocalInPolicy6Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Dport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Dst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Intf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicy6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocalInPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemLocalInPolicy6Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("dport"); ok || d.HasChange("dport") {
		t, err := expandSystemLocalInPolicy6Dport(d, v, "dport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemLocalInPolicy6Dst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLocalInPolicy6Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok || d.HasChange("intf") {
		t, err := expandSystemLocalInPolicy6Intf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemLocalInPolicy6Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandSystemLocalInPolicy6Src(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	return &obj, nil
}
