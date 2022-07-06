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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed: true,
			},
			"idp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSamlServiceProvidersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSamlServiceProviders(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSamlServiceProviders resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSamlServiceProviders(obj, adomv, nil)

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

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSamlServiceProviders(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSamlServiceProviders resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSamlServiceProviders(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSamlServiceProviders(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSamlServiceProviders(adomv, mkey, nil)
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

func flattenSystemSamlServiceProvidersIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSamlServiceProviders(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("idp_entity_id", flattenSystemSamlServiceProvidersIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "SystemSamlServiceProviders-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "SystemSamlServiceProviders-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "SystemSamlServiceProviders-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSamlServiceProvidersName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSamlServiceProviders-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("prefix", flattenSystemSamlServiceProvidersPrefix(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "SystemSamlServiceProviders-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("sp_cert", flattenSystemSamlServiceProvidersSpCert(o["sp-cert"], d, "sp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-cert"], "SystemSamlServiceProviders-SpCert"); ok {
			if err = d.Set("sp_cert", vv); err != nil {
				return fmt.Errorf("Error reading sp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_cert: %v", err)
		}
	}

	if err = d.Set("sp_entity_id", flattenSystemSamlServiceProvidersSpEntityId(o["sp-entity-id"], d, "sp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-entity-id"], "SystemSamlServiceProviders-SpEntityId"); ok {
			if err = d.Set("sp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading sp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_entity_id: %v", err)
		}
	}

	if err = d.Set("sp_single_logout_url", flattenSystemSamlServiceProvidersSpSingleLogoutUrl(o["sp-single-logout-url"], d, "sp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp-single-logout-url"], "SystemSamlServiceProviders-SpSingleLogoutUrl"); ok {
			if err = d.Set("sp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading sp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("sp_single_sign_on_url", flattenSystemSamlServiceProvidersSpSingleSignOnUrl(o["sp-single-sign-on-url"], d, "sp_single_sign_on_url")); err != nil {
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

func expandSystemSamlServiceProvidersIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSamlServiceProviders(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandSystemSamlServiceProvidersIdpEntityId(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSamlServiceProvidersName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandSystemSamlServiceProvidersPrefix(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("sp_cert"); ok || d.HasChange("sp_cert") {
		t, err := expandSystemSamlServiceProvidersSpCert(d, v, "sp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-cert"] = t
		}
	}

	if v, ok := d.GetOk("sp_entity_id"); ok || d.HasChange("sp_entity_id") {
		t, err := expandSystemSamlServiceProvidersSpEntityId(d, v, "sp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("sp_single_logout_url"); ok || d.HasChange("sp_single_logout_url") {
		t, err := expandSystemSamlServiceProvidersSpSingleLogoutUrl(d, v, "sp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("sp_single_sign_on_url"); ok || d.HasChange("sp_single_sign_on_url") {
		t, err := expandSystemSamlServiceProvidersSpSingleSignOnUrl(d, v, "sp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp-single-sign-on-url"] = t
		}
	}

	return &obj, nil
}
