// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Authorized service providers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSamlServiceProviders() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSamlServiceProvidersCreate,
		Read:   resourceSystemSamlServiceProvidersRead,
		Update: resourceSystemSamlServiceProvidersUpdate,
		Delete: resourceSystemSamlServiceProvidersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"idp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"idp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSamlServiceProvidersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSamlServiceProviders(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSamlServiceProviders resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSamlServiceProviders(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSamlServiceProviders resource: %v", err)
	}

	d.SetId(getStringKey(d, "idp_entity_id"))

	return resourceSystemSamlServiceProvidersRead(d, m)
}

func resourceSystemSamlServiceProvidersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSamlServiceProviders(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSamlServiceProviders resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSamlServiceProviders(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSamlServiceProviders resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "idp_entity_id"))

	return resourceSystemSamlServiceProvidersRead(d, m)
}

func resourceSystemSamlServiceProvidersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSamlServiceProviders(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSamlServiceProviders resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSamlServiceProvidersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSamlServiceProviders(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSamlServiceProviders resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSamlServiceProviders(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSamlServiceProviders resource from API: %v", err)
	}
	return nil
}

func flattenSystemSamlServiceProvidersIdpEntityId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleLogoutUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleSignOnUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpAdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpEntityId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleLogoutUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleSignOnUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSamlServiceProviders(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("idp_entity_id", flattenSystemSamlServiceProvidersIdpEntityId2edl(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "SystemSamlServiceProviders-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlServiceProvidersIdpSingleLogoutUrl2edl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "SystemSamlServiceProviders-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlServiceProvidersIdpSingleSignOnUrl2edl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "SystemSamlServiceProviders-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSamlServiceProvidersName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSamlServiceProviders-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("prefix", flattenSystemSamlServiceProvidersPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "SystemSamlServiceProviders-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("sp_adom", flattenSystemSamlServiceProvidersSpAdom2edl(o["sp-adom"], d, "sp_adom")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-adom"], "SystemSamlServiceProviders-SpAdom"); ok {
			if err = d.Set("sp_adom", vv); err != nil {
				return fmt.Errorf("Error reading sp_adom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_adom: %v", err)
		}
	}

	if err = d.Set("sp_cert", flattenSystemSamlServiceProvidersSpCert2edl(o["sp-cert"], d, "sp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-cert"], "SystemSamlServiceProviders-SpCert"); ok {
			if err = d.Set("sp_cert", vv); err != nil {
				return fmt.Errorf("Error reading sp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_cert: %v", err)
		}
	}

	if err = d.Set("sp_entity_id", flattenSystemSamlServiceProvidersSpEntityId2edl(o["sp-entity-id"], d, "sp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-entity-id"], "SystemSamlServiceProviders-SpEntityId"); ok {
			if err = d.Set("sp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading sp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_entity_id: %v", err)
		}
	}

	if err = d.Set("sp_profile", flattenSystemSamlServiceProvidersSpProfile2edl(o["sp-profile"], d, "sp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-profile"], "SystemSamlServiceProviders-SpProfile"); ok {
			if err = d.Set("sp_profile", vv); err != nil {
				return fmt.Errorf("Error reading sp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_profile: %v", err)
		}
	}

	if err = d.Set("sp_single_logout_url", flattenSystemSamlServiceProvidersSpSingleLogoutUrl2edl(o["sp-single-logout-url"], d, "sp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-single-logout-url"], "SystemSamlServiceProviders-SpSingleLogoutUrl"); ok {
			if err = d.Set("sp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading sp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("sp_single_sign_on_url", flattenSystemSamlServiceProvidersSpSingleSignOnUrl2edl(o["sp-single-sign-on-url"], d, "sp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-single-sign-on-url"], "SystemSamlServiceProviders-SpSingleSignOnUrl"); ok {
			if err = d.Set("sp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading sp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_single_sign_on_url: %v", err)
		}
	}

	return nil
}

func flattenSystemSamlServiceProvidersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSamlServiceProvidersIdpEntityId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleLogoutUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleSignOnUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpAdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpEntityId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleLogoutUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleSignOnUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSamlServiceProviders(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandSystemSamlServiceProvidersIdpEntityId2edl(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandSystemSamlServiceProvidersIdpSingleLogoutUrl2edl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandSystemSamlServiceProvidersIdpSingleSignOnUrl2edl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSamlServiceProvidersName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandSystemSamlServiceProvidersPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("sp_adom"); ok || d.HasChange("sp_adom") {
		t, err := expandSystemSamlServiceProvidersSpAdom2edl(d, v, "sp_adom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-adom"] = t
		}
	}

	if v, ok := d.GetOk("sp_cert"); ok || d.HasChange("sp_cert") {
		t, err := expandSystemSamlServiceProvidersSpCert2edl(d, v, "sp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-cert"] = t
		}
	}

	if v, ok := d.GetOk("sp_entity_id"); ok || d.HasChange("sp_entity_id") {
		t, err := expandSystemSamlServiceProvidersSpEntityId2edl(d, v, "sp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("sp_profile"); ok || d.HasChange("sp_profile") {
		t, err := expandSystemSamlServiceProvidersSpProfile2edl(d, v, "sp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-profile"] = t
		}
	}

	if v, ok := d.GetOk("sp_single_logout_url"); ok || d.HasChange("sp_single_logout_url") {
		t, err := expandSystemSamlServiceProvidersSpSingleLogoutUrl2edl(d, v, "sp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("sp_single_sign_on_url"); ok || d.HasChange("sp_single_sign_on_url") {
		t, err := expandSystemSamlServiceProvidersSpSingleSignOnUrl2edl(d, v, "sp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-single-sign-on-url"] = t
		}
	}

	return &obj, nil
}
