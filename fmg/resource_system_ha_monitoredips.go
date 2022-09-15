// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Monitored IP addresses.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaMonitoredIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaMonitoredIpsCreate,
		Read:   resourceSystemHaMonitoredIpsRead,
		Update: resourceSystemHaMonitoredIpsUpdate,
		Delete: resourceSystemHaMonitoredIpsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemHaMonitoredIpsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaMonitoredIps(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaMonitoredIps resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaMonitoredIps(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemHaMonitoredIps resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaMonitoredIpsRead(d, m)
}

func resourceSystemHaMonitoredIpsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaMonitoredIps(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitoredIps resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaMonitoredIps(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitoredIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaMonitoredIpsRead(d, m)
}

func resourceSystemHaMonitoredIpsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemHaMonitoredIps(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaMonitoredIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaMonitoredIpsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemHaMonitoredIps(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitoredIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaMonitoredIps(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitoredIps resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaMonitoredIpsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitoredIpsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaMonitoredIps(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemHaMonitoredIpsId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemHaMonitoredIps-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemHaMonitoredIpsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemHaMonitoredIps-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemHaMonitoredIpsIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemHaMonitoredIps-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenSystemHaMonitoredIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaMonitoredIpsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitoredIpsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaMonitoredIps(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemHaMonitoredIpsId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemHaMonitoredIpsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemHaMonitoredIpsIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
