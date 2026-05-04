// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SCIM client entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserScim() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserScimCreate,
		Read:   resourceObjectUserScimRead,
		Update: resourceObjectUserScimUpdate,
		Delete: resourceObjectUserScimDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cascade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_authentication_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_secret_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"token_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserScimCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserScim(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserScim resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserScim(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserScim(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserScim resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserScim(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserScim resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserScimRead(d, m)
}

func resourceObjectUserScimUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserScim(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserScim resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserScim(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserScim resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserScimRead(d, m)
}

func resourceObjectUserScimDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserScim(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserScim resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserScimRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserScim(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserScim resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserScim(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserScim resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserScimAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimBaseUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimCascade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserScimClientAuthenticationMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimClientIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimClientSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserScimTokenCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectUserScim(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_method", flattenObjectUserScimAuthMethod(o["auth-method"], d, "auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-method"], "ObjectUserScim-AuthMethod"); ok {
			if err = d.Set("auth_method", vv); err != nil {
				return fmt.Errorf("Error reading auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("base_url", flattenObjectUserScimBaseUrl(o["base-url"], d, "base_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["base-url"], "ObjectUserScim-BaseUrl"); ok {
			if err = d.Set("base_url", vv); err != nil {
				return fmt.Errorf("Error reading base_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base_url: %v", err)
		}
	}

	if err = d.Set("cascade", flattenObjectUserScimCascade(o["cascade"], d, "cascade")); err != nil {
		if vv, ok := fortiAPIPatch(o["cascade"], "ObjectUserScim-Cascade"); ok {
			if err = d.Set("cascade", vv); err != nil {
				return fmt.Errorf("Error reading cascade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cascade: %v", err)
		}
	}

	if err = d.Set("certificate", flattenObjectUserScimCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ObjectUserScim-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("client_authentication_method", flattenObjectUserScimClientAuthenticationMethod(o["client-authentication-method"], d, "client_authentication_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-authentication-method"], "ObjectUserScim-ClientAuthenticationMethod"); ok {
			if err = d.Set("client_authentication_method", vv); err != nil {
				return fmt.Errorf("Error reading client_authentication_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_authentication_method: %v", err)
		}
	}

	if err = d.Set("client_identity_check", flattenObjectUserScimClientIdentityCheck(o["client-identity-check"], d, "client_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-identity-check"], "ObjectUserScim-ClientIdentityCheck"); ok {
			if err = d.Set("client_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading client_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_identity_check: %v", err)
		}
	}

	if err = d.Set("client_secret_token", flattenObjectUserScimClientSecretToken(o["client-secret-token"], d, "client_secret_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-secret-token"], "ObjectUserScim-ClientSecretToken"); ok {
			if err = d.Set("client_secret_token", vv); err != nil {
				return fmt.Errorf("Error reading client_secret_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_secret_token: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserScimId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserScim-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserScimName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserScim-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserScimStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserScim-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("token_certificate", flattenObjectUserScimTokenCertificate(o["token-certificate"], d, "token_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["token-certificate"], "ObjectUserScim-TokenCertificate"); ok {
			if err = d.Set("token_certificate", vv); err != nil {
				return fmt.Errorf("Error reading token_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading token_certificate: %v", err)
		}
	}

	return nil
}

func flattenObjectUserScimFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserScimAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimBaseUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimCascade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserScimClientAuthenticationMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimClientIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimClientSecretToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserScimStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserScimTokenCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectUserScim(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_method"); ok || d.HasChange("auth_method") {
		t, err := expandObjectUserScimAuthMethod(d, v, "auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("base_url"); ok || d.HasChange("base_url") {
		t, err := expandObjectUserScimBaseUrl(d, v, "base_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base-url"] = t
		}
	}

	if v, ok := d.GetOk("cascade"); ok || d.HasChange("cascade") {
		t, err := expandObjectUserScimCascade(d, v, "cascade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cascade"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandObjectUserScimCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("client_authentication_method"); ok || d.HasChange("client_authentication_method") {
		t, err := expandObjectUserScimClientAuthenticationMethod(d, v, "client_authentication_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-authentication-method"] = t
		}
	}

	if v, ok := d.GetOk("client_identity_check"); ok || d.HasChange("client_identity_check") {
		t, err := expandObjectUserScimClientIdentityCheck(d, v, "client_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("client_secret_token"); ok || d.HasChange("client_secret_token") {
		t, err := expandObjectUserScimClientSecretToken(d, v, "client_secret_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret-token"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserScimId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserScimName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok || d.HasChange("secret") {
		t, err := expandObjectUserScimSecret(d, v, "secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserScimStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("token_certificate"); ok || d.HasChange("token_certificate") {
		t, err := expandObjectUserScimTokenCertificate(d, v, "token_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-certificate"] = t
		}
	}

	return &obj, nil
}
