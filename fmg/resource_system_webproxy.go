// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system web proxy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemWebProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemWebProxyUpdate,
		Read:   resourceSystemWebProxyRead,
		Update: resourceSystemWebProxyUpdate,
		Delete: resourceSystemWebProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
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
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemWebProxyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemWebProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemWebProxy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemWebProxy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemWebProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemWebProxy")

	return resourceSystemWebProxyRead(d, m)
}

func resourceSystemWebProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemWebProxy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemWebProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemWebProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemWebProxy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemWebProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemWebProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemWebProxy resource from API: %v", err)
	}
	return nil
}

func flattenSystemWebProxyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWebProxyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWebProxyPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemWebProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWebProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWebProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemWebProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenSystemWebProxyAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "SystemWebProxy-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemWebProxyMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemWebProxy-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemWebProxyPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "SystemWebProxy-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemWebProxyPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemWebProxy-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemWebProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemWebProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemWebProxyUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemWebProxy-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenSystemWebProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemWebProxyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWebProxyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWebProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemWebProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWebProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWebProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemWebProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandSystemWebProxyAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemWebProxyMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemWebProxyPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemWebProxyPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemWebProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemWebProxyUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
