// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure APN for GTP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectGtpApn() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGtpApnCreate,
		Read:   resourceObjectGtpApnRead,
		Update: resourceObjectGtpApnUpdate,
		Delete: resourceObjectGtpApnDelete,

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
			"apn": &schema.Schema{
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

func resourceObjectGtpApnCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpApn(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpApn resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGtpApn(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpApn resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpApnRead(d, m)
}

func resourceObjectGtpApnUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpApn(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpApn resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGtpApn(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpApn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpApnRead(d, m)
}

func resourceObjectGtpApnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectGtpApn(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGtpApn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGtpApnRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectGtpApn(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpApn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGtpApn(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpApn resource from API: %v", err)
	}
	return nil
}

func flattenObjectGtpApnApn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpApnName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectGtpApn(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("apn", flattenObjectGtpApnApn(o["apn"], d, "apn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn"], "ObjectGtpApn-Apn"); ok {
			if err = d.Set("apn", vv); err != nil {
				return fmt.Errorf("Error reading apn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectGtpApnName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGtpApn-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectGtpApnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGtpApnApn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpApnName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGtpApn(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apn"); ok {
		t, err := expandObjectGtpApnApn(d, v, "apn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectGtpApnName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
