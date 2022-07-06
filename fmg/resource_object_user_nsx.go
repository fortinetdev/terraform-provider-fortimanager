// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser Nsx

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserNsx() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserNsxCreate,
		Read:   resourceObjectUserNsxRead,
		Update: resourceObjectUserNsxUpdate,
		Delete: resourceObjectUserNsxDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fmgip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmgpasswd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"fmguser": &schema.Schema{
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
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserNsxCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserNsx(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserNsx resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserNsx(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserNsx resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserNsxRead(d, m)
}

func resourceObjectUserNsxUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserNsx(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserNsx resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserNsx(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserNsx resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserNsxRead(d, m)
}

func resourceObjectUserNsxDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserNsx(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserNsx resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserNsxRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserNsx(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserNsx resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserNsx(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserNsx resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserNsxFmgip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxFmgpasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserNsxFmguser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserNsxServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserNsxStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserNsxUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserNsx(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fmgip", flattenObjectUserNsxFmgip(o["fmgip"], d, "fmgip")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgip"], "ObjectUserNsx-Fmgip"); ok {
			if err = d.Set("fmgip", vv); err != nil {
				return fmt.Errorf("Error reading fmgip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgip: %v", err)
		}
	}

	if err = d.Set("fmguser", flattenObjectUserNsxFmguser(o["fmguser"], d, "fmguser")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmguser"], "ObjectUserNsx-Fmguser"); ok {
			if err = d.Set("fmguser", vv); err != nil {
				return fmt.Errorf("Error reading fmguser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmguser: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserNsxName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserNsx-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserNsxServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserNsx-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("service_id", flattenObjectUserNsxServiceId(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "ObjectUserNsx-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserNsxStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserNsx-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectUserNsxUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectUserNsx-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenObjectUserNsxFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserNsxFmgip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxFmgpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserNsxFmguser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserNsxServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserNsxStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserNsxUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserNsx(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgip"); ok || d.HasChange("fmgip") {
		t, err := expandObjectUserNsxFmgip(d, v, "fmgip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgip"] = t
		}
	}

	if v, ok := d.GetOk("fmgpasswd"); ok || d.HasChange("fmgpasswd") {
		t, err := expandObjectUserNsxFmgpasswd(d, v, "fmgpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgpasswd"] = t
		}
	}

	if v, ok := d.GetOk("fmguser"); ok || d.HasChange("fmguser") {
		t, err := expandObjectUserNsxFmguser(d, v, "fmguser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmguser"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserNsxName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserNsxPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserNsxServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandObjectUserNsxServiceId(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserNsxStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectUserNsxUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
