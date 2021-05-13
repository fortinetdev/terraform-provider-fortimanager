// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FSSO groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserAdgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserAdgrpCreate,
		Read:   resourceObjectUserAdgrpRead,
		Update: resourceObjectUserAdgrpUpdate,
		Delete: resourceObjectUserAdgrpDelete,

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
			"connector_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserAdgrpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserAdgrp resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserAdgrp(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserAdgrp resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserAdgrpRead(d, m)
}

func resourceObjectUserAdgrpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserAdgrp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserAdgrp(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserAdgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserAdgrpRead(d, m)
}

func resourceObjectUserAdgrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserAdgrp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserAdgrpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserAdgrp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserAdgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserAdgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserAdgrp resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserAdgrpConnectorSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAdgrpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAdgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserAdgrpServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserAdgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("connector_source", flattenObjectUserAdgrpConnectorSource(o["connector-source"], d, "connector_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["connector-source"], "ObjectUserAdgrp-ConnectorSource"); ok {
			if err = d.Set("connector_source", vv); err != nil {
				return fmt.Errorf("Error reading connector_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connector_source: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserAdgrpId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserAdgrp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserAdgrpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserAdgrp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenObjectUserAdgrpServerName(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "ObjectUserAdgrp-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	return nil
}

func flattenObjectUserAdgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserAdgrpConnectorSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAdgrpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAdgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserAdgrpServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserAdgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("connector_source"); ok {
		t, err := expandObjectUserAdgrpConnectorSource(d, v, "connector_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector-source"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectUserAdgrpId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserAdgrpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {
		t, err := expandObjectUserAdgrpServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	return &obj, nil
}
