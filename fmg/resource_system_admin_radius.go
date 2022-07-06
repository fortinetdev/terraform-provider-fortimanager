// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure radius.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAdminRadius() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminRadiusCreate,
		Read:   resourceSystemAdminRadiusRead,
		Update: resourceSystemAdminRadiusUpdate,
		Delete: resourceSystemAdminRadiusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auth_type": &schema.Schema{
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
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secondary_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAdminRadiusCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAdminRadius(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminRadius resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdminRadius(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdminRadius resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminRadiusRead(d, m)
}

func resourceSystemAdminRadiusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAdminRadius(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminRadius resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminRadius(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminRadius resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminRadiusRead(d, m)
}

func resourceSystemAdminRadiusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAdminRadius(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminRadius resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminRadiusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAdminRadius(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminRadius resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminRadius(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminRadius resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminRadiusAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRadiusName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRadiusNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRadiusSecondarySecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminRadiusSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminRadius(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_type", flattenSystemAdminRadiusAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "SystemAdminRadius-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAdminRadiusName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAdminRadius-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenSystemAdminRadiusNasIp(o["nas-ip"], d, "nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-ip"], "SystemAdminRadius-NasIp"); ok {
			if err = d.Set("nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAdminRadiusPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemAdminRadius-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secondary_secret", flattenSystemAdminRadiusSecondarySecret(o["secondary-secret"], d, "secondary_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-secret"], "SystemAdminRadius-SecondarySecret"); ok {
			if err = d.Set("secondary_secret", vv); err != nil {
				return fmt.Errorf("Error reading secondary_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_secret: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenSystemAdminRadiusSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "SystemAdminRadius-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("secret", flattenSystemAdminRadiusSecret(o["secret"], d, "secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["secret"], "SystemAdminRadius-Secret"); ok {
			if err = d.Set("secret", vv); err != nil {
				return fmt.Errorf("Error reading secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secret: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemAdminRadiusServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemAdminRadius-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminRadiusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminRadiusAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRadiusName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRadiusNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRadiusSecondarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminRadiusSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminRadius(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandSystemAdminRadiusAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAdminRadiusName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok || d.HasChange("nas_ip") {
		t, err := expandSystemAdminRadiusNasIp(d, v, "nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemAdminRadiusPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secondary_secret"); ok || d.HasChange("secondary_secret") {
		t, err := expandSystemAdminRadiusSecondarySecret(d, v, "secondary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandSystemAdminRadiusSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok || d.HasChange("secret") {
		t, err := expandSystemAdminRadiusSecret(d, v, "secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemAdminRadiusServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	return &obj, nil
}
