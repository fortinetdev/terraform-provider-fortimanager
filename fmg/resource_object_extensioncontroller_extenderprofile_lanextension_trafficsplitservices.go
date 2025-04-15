// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Config FortiExtender traffic split interface for LAN extension.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesCreate,
		Read:   resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead,
		Update: resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesDelete,

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
			"address": &schema.Schema{
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
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vsdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address", flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("service", flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("vsdb", flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(o["vsdb"], d, "vsdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["vsdb"], "ObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices-Vsdb"); ok {
			if err = d.Set("vsdb", vv); err != nil {
				return fmt.Errorf("Error reading vsdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vsdb: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService3rdl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("vsdb"); ok || d.HasChange("vsdb") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb3rdl(d, v, "vsdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vsdb"] = t
		}
	}

	return &obj, nil
}
