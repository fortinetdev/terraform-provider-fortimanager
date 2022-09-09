// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Peer.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemHaPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaPeerCreate,
		Read:   resourceSystemHaPeerRead,
		Update: resourceSystemHaPeerUpdate,
		Delete: resourceSystemHaPeerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaPeer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaPeer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemHaPeer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaPeerRead(d, m)
}

func resourceSystemHaPeerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaPeer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaPeer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaPeerRead(d, m)
}

func resourceSystemHaPeerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemHaPeer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaPeerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemHaPeer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaPeer resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaPeerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPeerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemHaPeerId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemHaPeer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemHaPeerIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemHaPeer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenSystemHaPeerIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "SystemHaPeer-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemHaPeerSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "SystemHaPeer-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemHaPeerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemHaPeer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemHaPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaPeerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPeerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemHaPeerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemHaPeerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandSystemHaPeerIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandSystemHaPeerSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemHaPeerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
