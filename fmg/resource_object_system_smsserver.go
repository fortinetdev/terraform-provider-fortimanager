// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SMS server for sending SMS messages to support user authentication.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemSmsServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSmsServerCreate,
		Read:   resourceObjectSystemSmsServerRead,
		Update: resourceObjectSystemSmsServerUpdate,
		Delete: resourceObjectSystemSmsServerDelete,

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
			"mail_server": &schema.Schema{
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
		},
	}
}

func resourceObjectSystemSmsServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemSmsServer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSmsServer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemSmsServer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSmsServer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSmsServerRead(d, m)
}

func resourceObjectSystemSmsServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemSmsServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSmsServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemSmsServer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSmsServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSmsServerRead(d, m)
}

func resourceObjectSystemSmsServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemSmsServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSmsServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSmsServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemSmsServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSmsServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSmsServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSmsServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSmsServerMailServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSmsServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSmsServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("mail_server", flattenObjectSystemSmsServerMailServer(o["mail-server"], d, "mail_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["mail-server"], "ObjectSystemSmsServer-MailServer"); ok {
			if err = d.Set("mail_server", vv); err != nil {
				return fmt.Errorf("Error reading mail_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mail_server: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemSmsServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemSmsServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSmsServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSmsServerMailServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSmsServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSmsServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mail_server"); ok || d.HasChange("mail_server") {
		t, err := expandObjectSystemSmsServerMailServer(d, v, "mail_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemSmsServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
