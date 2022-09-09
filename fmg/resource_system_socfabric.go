// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SOC Fabric.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSocFabric() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSocFabricUpdate,
		Read:   resourceSystemSocFabricRead,
		Update: resourceSystemSocFabricUpdate,
		Delete: resourceSystemSocFabricDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"supervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSocFabricUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSocFabric(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabric resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSocFabric(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemSocFabric resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSocFabric")

	return resourceSystemSocFabricRead(d, m)
}

func resourceSystemSocFabricDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSocFabric(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSocFabric resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSocFabricRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSocFabric(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabric resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSocFabric(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSocFabric resource from API: %v", err)
	}
	return nil
}

func flattenSystemSocFabricName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricPsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSocFabricRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSocFabricSupervisor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSocFabric(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemSocFabricName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSocFabric-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSocFabricPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSocFabric-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("psk", flattenSystemSocFabricPsk(o["psk"], d, "psk")); err != nil {
		if vv, ok := fortiAPIPatch(o["psk"], "SystemSocFabric-Psk"); ok {
			if err = d.Set("psk", vv); err != nil {
				return fmt.Errorf("Error reading psk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading psk: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSocFabricRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystemSocFabric-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemSocFabricSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemSocFabric-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSocFabricStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSocFabric-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("supervisor", flattenSystemSocFabricSupervisor(o["supervisor"], d, "supervisor")); err != nil {
		if vv, ok := fortiAPIPatch(o["supervisor"], "SystemSocFabric-Supervisor"); ok {
			if err = d.Set("supervisor", vv); err != nil {
				return fmt.Errorf("Error reading supervisor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading supervisor: %v", err)
		}
	}

	return nil
}

func flattenSystemSocFabricFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSocFabricName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricPsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSocFabricRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSocFabricSupervisor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSocFabric(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSocFabricName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSocFabricPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("psk"); ok || d.HasChange("psk") {
		t, err := expandSystemSocFabricPsk(d, v, "psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystemSocFabricRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemSocFabricSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSocFabricStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("supervisor"); ok || d.HasChange("supervisor") {
		t, err := expandSystemSocFabricSupervisor(d, v, "supervisor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supervisor"] = t
		}
	}

	return &obj, nil
}
