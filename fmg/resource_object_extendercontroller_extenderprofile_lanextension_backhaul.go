// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: LAN extension backhaul tunnel configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaul() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulCreate,
		Read:   resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulRead,
		Update: resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulUpdate,
		Delete: resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectExtenderControllerExtenderProfileLanExtensionBackhaul(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtenderControllerExtenderProfileLanExtensionBackhaul(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtenderControllerExtenderProfileLanExtensionBackhaul(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileLanExtensionBackhaulRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtenderControllerExtenderProfileLanExtensionBackhaul(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileLanExtensionBackhaul resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtenderControllerExtenderProfileLanExtensionBackhaul-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectExtenderControllerExtenderProfileLanExtensionBackhaul-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("role", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole3rdl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ObjectExtenderControllerExtenderProfileLanExtensionBackhaul-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectExtenderControllerExtenderProfileLanExtensionBackhaul-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole3rdl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
