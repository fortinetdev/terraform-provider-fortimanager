// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: LDAP Group Info.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserFssoPollingAdgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserFssoPollingAdgrpCreate,
		Read:   resourceObjectUserFssoPollingAdgrpRead,
		Update: resourceObjectUserFssoPollingAdgrpUpdate,
		Delete: resourceObjectUserFssoPollingAdgrpDelete,

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
			"fsso_polling": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserFssoPollingAdgrpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	fsso_polling := d.Get("fsso_polling").(string)
	paradict["fsso_polling"] = fsso_polling

	obj, err := getObjectObjectUserFssoPollingAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFssoPollingAdgrp resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserFssoPollingAdgrp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFssoPollingAdgrp resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserFssoPollingAdgrpRead(d, m)
}

func resourceObjectUserFssoPollingAdgrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	fsso_polling := d.Get("fsso_polling").(string)
	paradict["fsso_polling"] = fsso_polling

	obj, err := getObjectObjectUserFssoPollingAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFssoPollingAdgrp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserFssoPollingAdgrp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFssoPollingAdgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserFssoPollingAdgrpRead(d, m)
}

func resourceObjectUserFssoPollingAdgrpDelete(d *schema.ResourceData, m interface{}) error {
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

	fsso_polling := d.Get("fsso_polling").(string)
	paradict["fsso_polling"] = fsso_polling

	err = c.DeleteObjectUserFssoPollingAdgrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserFssoPollingAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserFssoPollingAdgrpRead(d *schema.ResourceData, m interface{}) error {
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

	fsso_polling := d.Get("fsso_polling").(string)
	if fsso_polling == "" {
		fsso_polling = importOptionChecking(m.(*FortiClient).Cfg, "fsso_polling")
		if fsso_polling == "" {
			return fmt.Errorf("Parameter fsso_polling is missing")
		}
		if err = d.Set("fsso_polling", fsso_polling); err != nil {
			return fmt.Errorf("Error set params fsso_polling: %v", err)
		}
	}
	paradict["fsso_polling"] = fsso_polling

	o, err := c.ReadObjectUserFssoPollingAdgrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFssoPollingAdgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserFssoPollingAdgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFssoPollingAdgrp resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserFssoPollingAdgrpName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserFssoPollingAdgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectUserFssoPollingAdgrpName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserFssoPollingAdgrp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectUserFssoPollingAdgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserFssoPollingAdgrpName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserFssoPollingAdgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserFssoPollingAdgrpName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
