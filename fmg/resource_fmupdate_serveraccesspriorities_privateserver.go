// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure multiple FortiManager units and private servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateServerAccessPrioritiesPrivateServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateServerAccessPrioritiesPrivateServerCreate,
		Read:   resourceFmupdateServerAccessPrioritiesPrivateServerRead,
		Update: resourceFmupdateServerAccessPrioritiesPrivateServerUpdate,
		Delete: resourceFmupdateServerAccessPrioritiesPrivateServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
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
			"time_zone": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateServerAccessPrioritiesPrivateServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateServerAccessPrioritiesPrivateServer(d)
	if err != nil {
		return fmt.Errorf("Error creating FmupdateServerAccessPrioritiesPrivateServer resource while getting object: %v", err)
	}

	_, err = c.CreateFmupdateServerAccessPrioritiesPrivateServer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating FmupdateServerAccessPrioritiesPrivateServer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFmupdateServerAccessPrioritiesPrivateServerRead(d, m)
}

func resourceFmupdateServerAccessPrioritiesPrivateServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateServerAccessPrioritiesPrivateServer(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateServerAccessPrioritiesPrivateServer resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateServerAccessPrioritiesPrivateServer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateServerAccessPrioritiesPrivateServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFmupdateServerAccessPrioritiesPrivateServerRead(d, m)
}

func resourceFmupdateServerAccessPrioritiesPrivateServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateServerAccessPrioritiesPrivateServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateServerAccessPrioritiesPrivateServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateServerAccessPrioritiesPrivateServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateServerAccessPrioritiesPrivateServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateServerAccessPrioritiesPrivateServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateServerAccessPrioritiesPrivateServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateServerAccessPrioritiesPrivateServer resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateServerAccessPrioritiesPrivateServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesPrivateServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesPrivateServerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateServerAccessPrioritiesPrivateServerTimeZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateServerAccessPrioritiesPrivateServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFmupdateServerAccessPrioritiesPrivateServerId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FmupdateServerAccessPrioritiesPrivateServer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFmupdateServerAccessPrioritiesPrivateServerIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FmupdateServerAccessPrioritiesPrivateServer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFmupdateServerAccessPrioritiesPrivateServerIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "FmupdateServerAccessPrioritiesPrivateServer-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("time_zone", flattenFmupdateServerAccessPrioritiesPrivateServerTimeZone(o["time_zone"], d, "time_zone")); err != nil {
		if vv, ok := fortiAPIPatch(o["time_zone"], "FmupdateServerAccessPrioritiesPrivateServer-TimeZone"); ok {
			if err = d.Set("time_zone", vv); err != nil {
				return fmt.Errorf("Error reading time_zone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time_zone: %v", err)
		}
	}

	return nil
}

func flattenFmupdateServerAccessPrioritiesPrivateServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateServerAccessPrioritiesPrivateServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateServerAccessPrioritiesPrivateServerTimeZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateServerAccessPrioritiesPrivateServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandFmupdateServerAccessPrioritiesPrivateServerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFmupdateServerAccessPrioritiesPrivateServerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandFmupdateServerAccessPrioritiesPrivateServerIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("time_zone"); ok {
		t, err := expandFmupdateServerAccessPrioritiesPrivateServerTimeZone(d, v, "time_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time_zone"] = t
		}
	}

	return &obj, nil
}
