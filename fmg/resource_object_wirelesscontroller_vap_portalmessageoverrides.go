// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Individual message overrides.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerVapPortalMessageOverrides() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerVapPortalMessageOverridesUpdate,
		Read:   resourceObjectWirelessControllerVapPortalMessageOverridesRead,
		Update: resourceObjectWirelessControllerVapPortalMessageOverridesUpdate,
		Delete: resourceObjectWirelessControllerVapPortalMessageOverridesDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auth_disclaimer_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_login_failed_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_login_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_reject_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerVapPortalMessageOverridesUpdate(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	obj, err := getObjectObjectWirelessControllerVapPortalMessageOverrides(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapPortalMessageOverrides resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerVapPortalMessageOverrides(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapPortalMessageOverrides resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerVapPortalMessageOverrides")

	return resourceObjectWirelessControllerVapPortalMessageOverridesRead(d, m)
}

func resourceObjectWirelessControllerVapPortalMessageOverridesDelete(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerVapPortalMessageOverrides(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerVapPortalMessageOverrides resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerVapPortalMessageOverridesRead(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	if vap == "" {
		vap = importOptionChecking(m.(*FortiClient).Cfg, "vap")
		if vap == "" {
			return fmt.Errorf("Parameter vap is missing")
		}
		if err = d.Set("vap", vap); err != nil {
			return fmt.Errorf("Error set params vap: %v", err)
		}
	}
	paradict["vap"] = vap

	o, err := c.ReadObjectWirelessControllerVapPortalMessageOverrides(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapPortalMessageOverrides resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerVapPortalMessageOverrides(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapPortalMessageOverrides resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_disclaimer_page", flattenObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(o["auth-disclaimer-page"], d, "auth_disclaimer_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-disclaimer-page"], "ObjectWirelessControllerVapPortalMessageOverrides-AuthDisclaimerPage"); ok {
			if err = d.Set("auth_disclaimer_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_disclaimer_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_disclaimer_page: %v", err)
		}
	}

	if err = d.Set("auth_login_failed_page", flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(o["auth-login-failed-page"], d, "auth_login_failed_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-login-failed-page"], "ObjectWirelessControllerVapPortalMessageOverrides-AuthLoginFailedPage"); ok {
			if err = d.Set("auth_login_failed_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_login_failed_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_login_failed_page: %v", err)
		}
	}

	if err = d.Set("auth_login_page", flattenObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(o["auth-login-page"], d, "auth_login_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-login-page"], "ObjectWirelessControllerVapPortalMessageOverrides-AuthLoginPage"); ok {
			if err = d.Set("auth_login_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_login_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_login_page: %v", err)
		}
	}

	if err = d.Set("auth_reject_page", flattenObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(o["auth-reject-page"], d, "auth_reject_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-reject-page"], "ObjectWirelessControllerVapPortalMessageOverrides-AuthRejectPage"); ok {
			if err = d.Set("auth_reject_page", vv); err != nil {
				return fmt.Errorf("Error reading auth_reject_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_reject_page: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerVapPortalMessageOverridesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_disclaimer_page"); ok || d.HasChange("auth_disclaimer_page") {
		t, err := expandObjectWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage2edl(d, v, "auth_disclaimer_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-disclaimer-page"] = t
		}
	}

	if v, ok := d.GetOk("auth_login_failed_page"); ok || d.HasChange("auth_login_failed_page") {
		t, err := expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage2edl(d, v, "auth_login_failed_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-login-failed-page"] = t
		}
	}

	if v, ok := d.GetOk("auth_login_page"); ok || d.HasChange("auth_login_page") {
		t, err := expandObjectWirelessControllerVapPortalMessageOverridesAuthLoginPage2edl(d, v, "auth_login_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-login-page"] = t
		}
	}

	if v, ok := d.GetOk("auth_reject_page"); ok || d.HasChange("auth_reject_page") {
		t, err := expandObjectWirelessControllerVapPortalMessageOverridesAuthRejectPage2edl(d, v, "auth_reject_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-reject-page"] = t
		}
	}

	return &obj, nil
}
