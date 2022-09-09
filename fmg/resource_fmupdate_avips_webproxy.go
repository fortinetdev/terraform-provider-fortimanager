// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the web proxy for use with FortiGuard antivirus and IPS updates.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateAvIpsWebProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateAvIpsWebProxyUpdate,
		Read:   resourceFmupdateAvIpsWebProxyRead,
		Update: resourceFmupdateAvIpsWebProxyUpdate,
		Delete: resourceFmupdateAvIpsWebProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
			},
		},
	}
}

func resourceFmupdateAvIpsWebProxyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateAvIpsWebProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAvIpsWebProxy resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateAvIpsWebProxy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAvIpsWebProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateAvIpsWebProxy")

	return resourceFmupdateAvIpsWebProxyRead(d, m)
}

func resourceFmupdateAvIpsWebProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateAvIpsWebProxy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateAvIpsWebProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateAvIpsWebProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateAvIpsWebProxy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateAvIpsWebProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateAvIpsWebProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateAvIpsWebProxy resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateAvIpsWebProxyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateAvIpsWebProxyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateAvIpsWebProxyPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateAvIpsWebProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateAvIpsWebProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateAvIpsWebProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateAvIpsWebProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenFmupdateAvIpsWebProxyAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "FmupdateAvIpsWebProxy-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("mode", flattenFmupdateAvIpsWebProxyMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "FmupdateAvIpsWebProxy-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenFmupdateAvIpsWebProxyPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateAvIpsWebProxy-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFmupdateAvIpsWebProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateAvIpsWebProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("username", flattenFmupdateAvIpsWebProxyUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "FmupdateAvIpsWebProxy-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenFmupdateAvIpsWebProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateAvIpsWebProxyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateAvIpsWebProxyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateAvIpsWebProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateAvIpsWebProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateAvIpsWebProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateAvIpsWebProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateAvIpsWebProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandFmupdateAvIpsWebProxyAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandFmupdateAvIpsWebProxyMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandFmupdateAvIpsWebProxyPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFmupdateAvIpsWebProxyPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateAvIpsWebProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandFmupdateAvIpsWebProxyUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
