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

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverCreate,
		Read:   resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead,
		Update: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverDelete,

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

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("alert", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(o["alert"], d, "alert")); err != nil {
		if vv, ok := fortiAPIPatch(o["alert"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver-Alert"); ok {
			if err = d.Set("alert", vv); err != nil {
				return fmt.Errorf("Error reading alert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alert: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("phone_number", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(o["phone-number"], d, "phone_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["phone-number"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver-PhoneNumber"); ok {
			if err = d.Set("phone_number", vv); err != nil {
				return fmt.Errorf("Error reading phone_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phone_number: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alert"); ok || d.HasChange("alert") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert4thl(d, v, "alert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("phone_number"); ok || d.HasChange("phone_number") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber4thl(d, v, "phone_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phone-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus4thl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
