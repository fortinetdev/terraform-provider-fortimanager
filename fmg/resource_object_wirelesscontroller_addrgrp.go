// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the MAC address group.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerAddrgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerAddrgrpCreate,
		Read:   resourceObjectWirelessControllerAddrgrpRead,
		Update: resourceObjectWirelessControllerAddrgrpUpdate,
		Delete: resourceObjectWirelessControllerAddrgrpDelete,

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
			"addresses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerAddrgrpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerAddrgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerAddrgrp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerAddrgrp(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerAddrgrp resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectWirelessControllerAddrgrpRead(d, m)
}

func resourceObjectWirelessControllerAddrgrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerAddrgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAddrgrp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerAddrgrp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAddrgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectWirelessControllerAddrgrpRead(d, m)
}

func resourceObjectWirelessControllerAddrgrpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerAddrgrp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerAddrgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerAddrgrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerAddrgrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerAddrgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerAddrgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerAddrgrp resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerAddrgrpAddresses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerAddrgrpDefaultPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerAddrgrpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerAddrgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addresses", flattenObjectWirelessControllerAddrgrpAddresses(o["addresses"], d, "addresses")); err != nil {
		if vv, ok := fortiAPIPatch(o["addresses"], "ObjectWirelessControllerAddrgrp-Addresses"); ok {
			if err = d.Set("addresses", vv); err != nil {
				return fmt.Errorf("Error reading addresses: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addresses: %v", err)
		}
	}

	if err = d.Set("default_policy", flattenObjectWirelessControllerAddrgrpDefaultPolicy(o["default-policy"], d, "default_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-policy"], "ObjectWirelessControllerAddrgrp-DefaultPolicy"); ok {
			if err = d.Set("default_policy", vv); err != nil {
				return fmt.Errorf("Error reading default_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_policy: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerAddrgrpId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerAddrgrp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerAddrgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerAddrgrpAddresses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerAddrgrpDefaultPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAddrgrpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerAddrgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addresses"); ok || d.HasChange("addresses") {
		t, err := expandObjectWirelessControllerAddrgrpAddresses(d, v, "addresses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addresses"] = t
		}
	}

	if v, ok := d.GetOk("default_policy"); ok || d.HasChange("default_policy") {
		t, err := expandObjectWirelessControllerAddrgrpDefaultPolicy(d, v, "default_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-policy"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerAddrgrpId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
