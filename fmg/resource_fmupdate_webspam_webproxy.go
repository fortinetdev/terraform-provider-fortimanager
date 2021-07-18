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

func resourceFmupdateWebSpamWebProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateWebSpamWebProxyUpdate,
		Read:   resourceFmupdateWebSpamWebProxyRead,
		Update: resourceFmupdateWebSpamWebProxyUpdate,
		Delete: resourceFmupdateWebSpamWebProxyDelete,

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
				Computed: true,
			},
		},
	}
}

func resourceFmupdateWebSpamWebProxyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateWebSpamWebProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamWebProxy resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateWebSpamWebProxy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateWebSpamWebProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateWebSpamWebProxy")

	return resourceFmupdateWebSpamWebProxyRead(d, m)
}

func resourceFmupdateWebSpamWebProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateWebSpamWebProxy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateWebSpamWebProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateWebSpamWebProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateWebSpamWebProxy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamWebProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateWebSpamWebProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateWebSpamWebProxy resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateWebSpamWebProxyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamWebProxyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamWebProxyPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFmupdateWebSpamWebProxyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamWebProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateWebSpamWebProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateWebSpamWebProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenFmupdateWebSpamWebProxyAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "FmupdateWebSpamWebProxy-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("mode", flattenFmupdateWebSpamWebProxyMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "FmupdateWebSpamWebProxy-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenFmupdateWebSpamWebProxyPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateWebSpamWebProxy-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFmupdateWebSpamWebProxyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateWebSpamWebProxy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("username", flattenFmupdateWebSpamWebProxyUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "FmupdateWebSpamWebProxy-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenFmupdateWebSpamWebProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateWebSpamWebProxyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamWebProxyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamWebProxyPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFmupdateWebSpamWebProxyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamWebProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateWebSpamWebProxyUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateWebSpamWebProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok {
		t, err := expandFmupdateWebSpamWebProxyAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandFmupdateWebSpamWebProxyMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandFmupdateWebSpamWebProxyPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandFmupdateWebSpamWebProxyPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFmupdateWebSpamWebProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandFmupdateWebSpamWebProxyUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
