// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiToken.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserFortitoken() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserFortitokenCreate,
		Read:   resourceObjectUserFortitokenRead,
		Update: resourceObjectUserFortitokenUpdate,
		Delete: resourceObjectUserFortitokenDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserFortitokenCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserFortitoken(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFortitoken resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserFortitoken(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFortitoken resource: %v", err)
	}

	d.SetId(getStringKey(d, "serial_number"))

	return resourceObjectUserFortitokenRead(d, m)
}

func resourceObjectUserFortitokenUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserFortitoken(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFortitoken resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserFortitoken(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFortitoken resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "serial_number"))

	return resourceObjectUserFortitokenRead(d, m)
}

func resourceObjectUserFortitokenDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserFortitoken(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserFortitoken resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserFortitokenRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserFortitoken(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFortitoken resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserFortitoken(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFortitoken resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserFortitokenComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFortitokenLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFortitokenSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFortitokenStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserFortitoken(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comments", flattenObjectUserFortitokenComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectUserFortitoken-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("license", flattenObjectUserFortitokenLicense(o["license"], d, "license")); err != nil {
		if vv, ok := fortiAPIPatch(o["license"], "ObjectUserFortitoken-License"); ok {
			if err = d.Set("license", vv); err != nil {
				return fmt.Errorf("Error reading license: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading license: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenObjectUserFortitokenSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "ObjectUserFortitoken-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserFortitokenStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserFortitoken-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectUserFortitokenFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserFortitokenComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFortitokenLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFortitokenSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFortitokenStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserFortitoken(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectUserFortitokenComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("license"); ok || d.HasChange("license") {
		t, err := expandObjectUserFortitokenLicense(d, v, "license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandObjectUserFortitokenSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserFortitokenStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
