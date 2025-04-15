// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: TACACS+ server entry configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminTacacs() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminTacacsCreate,
		Read:   resourceSystemAdminTacacsRead,
		Update: resourceSystemAdminTacacsUpdate,
		Delete: resourceSystemAdminTacacsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"authen_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secondary_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tertiary_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminTacacsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAdminTacacs(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminTacacs resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminTacacs(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminTacacs resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminTacacsRead(d, m)
}

func resourceSystemAdminTacacsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAdminTacacs(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminTacacs resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminTacacs(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminTacacs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminTacacsRead(d, m)
}

func resourceSystemAdminTacacsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemAdminTacacs(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminTacacs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminTacacsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemAdminTacacs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminTacacs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminTacacs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminTacacs resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminTacacsAuthenType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTacacsAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTacacsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTacacsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTacacsSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTacacsTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminTacacs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("authen_type", flattenSystemAdminTacacsAuthenType(o["authen-type"], d, "authen_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["authen-type"], "SystemAdminTacacs-AuthenType"); ok {
			if err = d.Set("authen_type", vv); err != nil {
				return fmt.Errorf("Error reading authen_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authen_type: %v", err)
		}
	}

	if err = d.Set("authorization", flattenSystemAdminTacacsAuthorization(o["authorization"], d, "authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization"], "SystemAdminTacacs-Authorization"); ok {
			if err = d.Set("authorization", vv); err != nil {
				return fmt.Errorf("Error reading authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAdminTacacsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAdminTacacs-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAdminTacacsPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemAdminTacacs-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenSystemAdminTacacsSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "SystemAdminTacacs-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemAdminTacacsServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemAdminTacacs-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenSystemAdminTacacsTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "SystemAdminTacacs-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminTacacsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminTacacsAuthenType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTacacsAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTacacsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminTacacsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTacacsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTacacsSecondaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminTacacsSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTacacsTertiaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminTacacsTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminTacacs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authen_type"); ok || d.HasChange("authen_type") {
		t, err := expandSystemAdminTacacsAuthenType(d, v, "authen_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authen-type"] = t
		}
	}

	if v, ok := d.GetOk("authorization"); ok || d.HasChange("authorization") {
		t, err := expandSystemAdminTacacsAuthorization(d, v, "authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystemAdminTacacsKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAdminTacacsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemAdminTacacsPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secondary_key"); ok || d.HasChange("secondary_key") {
		t, err := expandSystemAdminTacacsSecondaryKey(d, v, "secondary_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-key"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandSystemAdminTacacsSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemAdminTacacsServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_key"); ok || d.HasChange("tertiary_key") {
		t, err := expandSystemAdminTacacsTertiaryKey(d, v, "tertiary_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-key"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandSystemAdminTacacsTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	return &obj, nil
}
