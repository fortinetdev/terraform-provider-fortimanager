// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser Guardicore

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserGuardicore() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserGuardicoreCreate,
		Read:   resourceObjectUserGuardicoreRead,
		Update: resourceObjectUserGuardicoreUpdate,
		Delete: resourceObjectUserGuardicoreDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
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
			"scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upd_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserGuardicoreCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserGuardicore(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGuardicore resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserGuardicore(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserGuardicore(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserGuardicore resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserGuardicore(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserGuardicore resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserGuardicoreRead(d, m)
}

func resourceObjectUserGuardicoreUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserGuardicore(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGuardicore resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserGuardicore(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGuardicore resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserGuardicoreRead(d, m)
}

func resourceObjectUserGuardicoreDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserGuardicore(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserGuardicore resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserGuardicoreRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectUserGuardicore(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserGuardicore resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserGuardicore(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGuardicore resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserGuardicoreAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicoreFilters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserGuardicoreName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicorePageSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicoreScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicoreServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicoreStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicoreUpdInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGuardicoreUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserGuardicore(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("alias", flattenObjectUserGuardicoreAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "ObjectUserGuardicore-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("filters", flattenObjectUserGuardicoreFilters(o["filters"], d, "filters")); err != nil {
		if vv, ok := fortiAPIPatch(o["filters"], "ObjectUserGuardicore-Filters"); ok {
			if err = d.Set("filters", vv); err != nil {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filters: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserGuardicoreName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserGuardicore-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("page_size", flattenObjectUserGuardicorePageSize(o["page_size"], d, "page_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["page_size"], "ObjectUserGuardicore-PageSize"); ok {
			if err = d.Set("page_size", vv); err != nil {
				return fmt.Errorf("Error reading page_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading page_size: %v", err)
		}
	}

	if err = d.Set("scheme", flattenObjectUserGuardicoreScheme(o["scheme"], d, "scheme")); err != nil {
		if vv, ok := fortiAPIPatch(o["scheme"], "ObjectUserGuardicore-Scheme"); ok {
			if err = d.Set("scheme", vv); err != nil {
				return fmt.Errorf("Error reading scheme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scheme: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserGuardicoreServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserGuardicore-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserGuardicoreStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserGuardicore-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upd_interval", flattenObjectUserGuardicoreUpdInterval(o["upd_interval"], d, "upd_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["upd_interval"], "ObjectUserGuardicore-UpdInterval"); ok {
			if err = d.Set("upd_interval", vv); err != nil {
				return fmt.Errorf("Error reading upd_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upd_interval: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectUserGuardicoreUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectUserGuardicore-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenObjectUserGuardicoreFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserGuardicoreAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicoreFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGuardicoreName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicorePageSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicorePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserGuardicoreScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicoreServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicoreStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicoreUpdInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGuardicoreUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserGuardicore(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandObjectUserGuardicoreAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandObjectUserGuardicoreFilters(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserGuardicoreName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("page_size"); ok || d.HasChange("page_size") {
		t, err := expandObjectUserGuardicorePageSize(d, v, "page_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page_size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserGuardicorePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("scheme"); ok || d.HasChange("scheme") {
		t, err := expandObjectUserGuardicoreScheme(d, v, "scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scheme"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserGuardicoreServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserGuardicoreStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upd_interval"); ok || d.HasChange("upd_interval") {
		t, err := expandObjectUserGuardicoreUpdInterval(d, v, "upd_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upd_interval"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectUserGuardicoreUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
