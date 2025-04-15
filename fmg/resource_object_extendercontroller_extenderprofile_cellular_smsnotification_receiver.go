// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SMS notification receiver list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverCreate,
		Read:   resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverRead,
		Update: resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverUpdate,
		Delete: resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverDelete,

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alert": &schema.Schema{
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
			"phone_number": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverCreate(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverUpdate(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverDelete(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverRead(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("alert", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(o["alert"], d, "alert")); err != nil {
		if vv, ok := fortiAPIPatch(o["alert"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver-Alert"); ok {
			if err = d.Set("alert", vv); err != nil {
				return fmt.Errorf("Error reading alert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alert: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("phone_number", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(o["phone-number"], d, "phone_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["phone-number"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver-PhoneNumber"); ok {
			if err = d.Set("phone_number", vv); err != nil {
				return fmt.Errorf("Error reading phone_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phone_number: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alert"); ok || d.HasChange("alert") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(d, v, "alert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("phone_number"); ok || d.HasChange("phone_number") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(d, v, "phone_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phone-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
