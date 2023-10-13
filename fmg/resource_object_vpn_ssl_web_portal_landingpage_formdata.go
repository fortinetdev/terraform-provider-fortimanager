// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Form data.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebPortalLandingPageFormData() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalLandingPageFormDataCreate,
		Read:   resourceObjectVpnSslWebPortalLandingPageFormDataRead,
		Update: resourceObjectVpnSslWebPortalLandingPageFormDataUpdate,
		Delete: resourceObjectVpnSslWebPortalLandingPageFormDataDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectVpnSslWebPortalLandingPageFormDataCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalLandingPageFormData(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalLandingPageFormData resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebPortalLandingPageFormData(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalLandingPageFormDataRead(d, m)
}

func resourceObjectVpnSslWebPortalLandingPageFormDataUpdate(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalLandingPageFormData(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalLandingPageFormData resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortalLandingPageFormData(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalLandingPageFormDataRead(d, m)
}

func resourceObjectVpnSslWebPortalLandingPageFormDataDelete(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	paradict["portal"] = portal

	err = c.DeleteObjectVpnSslWebPortalLandingPageFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalLandingPageFormDataRead(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["portal"] = portal

	o, err := c.ReadObjectVpnSslWebPortalLandingPageFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalLandingPageFormData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortalLandingPageFormData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalLandingPageFormData resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalLandingPageFormDataName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageFormDataValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortalLandingPageFormData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectVpnSslWebPortalLandingPageFormDataName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebPortalLandingPageFormData-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectVpnSslWebPortalLandingPageFormDataValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectVpnSslWebPortalLandingPageFormData-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalLandingPageFormDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalLandingPageFormDataName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageFormDataValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortalLandingPageFormData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebPortalLandingPageFormDataName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectVpnSslWebPortalLandingPageFormDataValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
