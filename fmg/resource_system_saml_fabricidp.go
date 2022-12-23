// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Authorized identity providers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSamlFabricIdp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSamlFabricIdpCreate,
		Read:   resourceSystemSamlFabricIdpRead,
		Update: resourceSystemSamlFabricIdpUpdate,
		Delete: resourceSystemSamlFabricIdpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dev_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"idp_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
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
			"idp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSamlFabricIdpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSamlFabricIdp(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSamlFabricIdp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSamlFabricIdp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSamlFabricIdp resource: %v", err)
	}

	d.SetId(getStringKey(d, "dev_id"))

	return resourceSystemSamlFabricIdpRead(d, m)
}

func resourceSystemSamlFabricIdpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSamlFabricIdp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSamlFabricIdp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSamlFabricIdp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSamlFabricIdp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "dev_id"))

	return resourceSystemSamlFabricIdpRead(d, m)
}

func resourceSystemSamlFabricIdpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSamlFabricIdp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSamlFabricIdp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSamlFabricIdpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSamlFabricIdp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSamlFabricIdp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSamlFabricIdp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSamlFabricIdp resource from API: %v", err)
	}
	return nil
}

func flattenSystemSamlFabricIdpDevId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlFabricIdpIdpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSamlFabricIdp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dev_id", flattenSystemSamlFabricIdpDevId(o["dev-id"], d, "dev_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-id"], "SystemSamlFabricIdp-DevId"); ok {
			if err = d.Set("dev_id", vv); err != nil {
				return fmt.Errorf("Error reading dev_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_id: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenSystemSamlFabricIdpIdpCert(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "SystemSamlFabricIdp-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenSystemSamlFabricIdpIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "SystemSamlFabricIdp-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlFabricIdpIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "SystemSamlFabricIdp-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlFabricIdpIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "SystemSamlFabricIdp-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("idp_status", flattenSystemSamlFabricIdpIdpStatus(o["idp-status"], d, "idp_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-status"], "SystemSamlFabricIdp-IdpStatus"); ok {
			if err = d.Set("idp_status", vv); err != nil {
				return fmt.Errorf("Error reading idp_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_status: %v", err)
		}
	}

	return nil
}

func flattenSystemSamlFabricIdpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSamlFabricIdpDevId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlFabricIdpIdpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSamlFabricIdp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dev_id"); ok || d.HasChange("dev_id") {
		t, err := expandSystemSamlFabricIdpDevId(d, v, "dev_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandSystemSamlFabricIdpIdpCert(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandSystemSamlFabricIdpIdpEntityId(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandSystemSamlFabricIdpIdpSingleLogoutUrl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandSystemSamlFabricIdpIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_status"); ok || d.HasChange("idp_status") {
		t, err := expandSystemSamlFabricIdpIdpStatus(d, v, "idp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-status"] = t
		}
	}

	return &obj, nil
}
