// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Syslog servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSyslog() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSyslogCreate,
		Read:   resourceSystemSyslogRead,
		Update: resourceSystemSyslogUpdate,
		Delete: resourceSystemSyslogDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSyslogCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSyslog(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSyslog resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSyslog(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemSyslog resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSyslogRead(d, m)
}

func resourceSystemSyslogUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSyslog(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSyslog resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSyslog(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSyslog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSyslogRead(d, m)
}

func resourceSystemSyslogDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSyslog(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSyslog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSyslogRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSyslog(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSyslog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSyslog(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSyslog resource from API: %v", err)
	}
	return nil
}

func flattenSystemSyslogIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSyslogPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSyslog(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenSystemSyslogIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemSyslog-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSyslogName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSyslog-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSyslogPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSyslog-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func flattenSystemSyslogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSyslogIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSyslogPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSyslog(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemSyslogIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSyslogName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSyslogPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	return &obj, nil
}
