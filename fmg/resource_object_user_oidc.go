// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser Oidc

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserOidc() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserOidcCreate,
		Read:   resourceObjectUserOidcRead,
		Update: resourceObjectUserOidcUpdate,
		Delete: resourceObjectUserOidcDelete,

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
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clock_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"discovery_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_hint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icon_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"jwks_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
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
			"private_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"token_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_issuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserOidcCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserOidc(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserOidc resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserOidc(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserOidc(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserOidc resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserOidc(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserOidc resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserOidcRead(d, m)
}

func resourceObjectUserOidcUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserOidc(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserOidc resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserOidc(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserOidc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserOidcRead(d, m)
}

func resourceObjectUserOidcDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserOidc(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserOidc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserOidcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserOidc(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserOidc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserOidc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserOidc resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserOidcAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcAuthorizationUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcClientSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcDiscoveryUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcDisplayName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcDomainHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcGroupAttrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcIconUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcIssuer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcJwksUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserOidcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserOidcTokenUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcUserAttrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcUserRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcVerifyCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserOidcVerifyIssuer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserOidc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_method", flattenObjectUserOidcAuthMethod(o["auth-method"], d, "auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-method"], "ObjectUserOidc-AuthMethod"); ok {
			if err = d.Set("auth_method", vv); err != nil {
				return fmt.Errorf("Error reading auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectUserOidcAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectUserOidc-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("authorization_url", flattenObjectUserOidcAuthorizationUrl(o["authorization-url"], d, "authorization_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization-url"], "ObjectUserOidc-AuthorizationUrl"); ok {
			if err = d.Set("authorization_url", vv); err != nil {
				return fmt.Errorf("Error reading authorization_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization_url: %v", err)
		}
	}

	if err = d.Set("client_id", flattenObjectUserOidcClientId(o["client-id"], d, "client_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-id"], "ObjectUserOidc-ClientId"); ok {
			if err = d.Set("client_id", vv); err != nil {
				return fmt.Errorf("Error reading client_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("client_secret", flattenObjectUserOidcClientSecret(o["client-secret"], d, "client_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-secret"], "ObjectUserOidc-ClientSecret"); ok {
			if err = d.Set("client_secret", vv); err != nil {
				return fmt.Errorf("Error reading client_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_secret: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", flattenObjectUserOidcClockTolerance(o["clock-tolerance"], d, "clock_tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["clock-tolerance"], "ObjectUserOidc-ClockTolerance"); ok {
			if err = d.Set("clock_tolerance", vv); err != nil {
				return fmt.Errorf("Error reading clock_tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("discovery_url", flattenObjectUserOidcDiscoveryUrl(o["discovery-url"], d, "discovery_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["discovery-url"], "ObjectUserOidc-DiscoveryUrl"); ok {
			if err = d.Set("discovery_url", vv); err != nil {
				return fmt.Errorf("Error reading discovery_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discovery_url: %v", err)
		}
	}

	if err = d.Set("display_name", flattenObjectUserOidcDisplayName(o["display-name"], d, "display_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-name"], "ObjectUserOidc-DisplayName"); ok {
			if err = d.Set("display_name", vv); err != nil {
				return fmt.Errorf("Error reading display_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_name: %v", err)
		}
	}

	if err = d.Set("domain_hint", flattenObjectUserOidcDomainHint(o["domain-hint"], d, "domain_hint")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-hint"], "ObjectUserOidc-DomainHint"); ok {
			if err = d.Set("domain_hint", vv); err != nil {
				return fmt.Errorf("Error reading domain_hint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_hint: %v", err)
		}
	}

	if err = d.Set("group_attr_name", flattenObjectUserOidcGroupAttrName(o["group-attr-name"], d, "group_attr_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-attr-name"], "ObjectUserOidc-GroupAttrName"); ok {
			if err = d.Set("group_attr_name", vv); err != nil {
				return fmt.Errorf("Error reading group_attr_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_attr_name: %v", err)
		}
	}

	if err = d.Set("icon_url", flattenObjectUserOidcIconUrl(o["icon-url"], d, "icon_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon-url"], "ObjectUserOidc-IconUrl"); ok {
			if err = d.Set("icon_url", vv); err != nil {
				return fmt.Errorf("Error reading icon_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon_url: %v", err)
		}
	}

	if err = d.Set("issuer", flattenObjectUserOidcIssuer(o["issuer"], d, "issuer")); err != nil {
		if vv, ok := fortiAPIPatch(o["issuer"], "ObjectUserOidc-Issuer"); ok {
			if err = d.Set("issuer", vv); err != nil {
				return fmt.Errorf("Error reading issuer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("jwks_uri", flattenObjectUserOidcJwksUri(o["jwks-uri"], d, "jwks_uri")); err != nil {
		if vv, ok := fortiAPIPatch(o["jwks-uri"], "ObjectUserOidc-JwksUri"); ok {
			if err = d.Set("jwks_uri", vv); err != nil {
				return fmt.Errorf("Error reading jwks_uri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jwks_uri: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserOidcLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserOidc-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserOidcName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserOidc-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("private_key", flattenObjectUserOidcPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "ObjectUserOidc-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("token_url", flattenObjectUserOidcTokenUrl(o["token-url"], d, "token_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["token-url"], "ObjectUserOidc-TokenUrl"); ok {
			if err = d.Set("token_url", vv); err != nil {
				return fmt.Errorf("Error reading token_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading token_url: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserOidcType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserOidc-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user_attr_name", flattenObjectUserOidcUserAttrName(o["user-attr-name"], d, "user_attr_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-attr-name"], "ObjectUserOidc-UserAttrName"); ok {
			if err = d.Set("user_attr_name", vv); err != nil {
				return fmt.Errorf("Error reading user_attr_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_attr_name: %v", err)
		}
	}

	if err = d.Set("user_regex", flattenObjectUserOidcUserRegex(o["user-regex"], d, "user_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-regex"], "ObjectUserOidc-UserRegex"); ok {
			if err = d.Set("user_regex", vv); err != nil {
				return fmt.Errorf("Error reading user_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_regex: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenObjectUserOidcVerifyCert(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "ObjectUserOidc-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("verify_issuer", flattenObjectUserOidcVerifyIssuer(o["verify-issuer"], d, "verify_issuer")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-issuer"], "ObjectUserOidc-VerifyIssuer"); ok {
			if err = d.Set("verify_issuer", vv); err != nil {
				return fmt.Errorf("Error reading verify_issuer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_issuer: %v", err)
		}
	}

	return nil
}

func flattenObjectUserOidcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserOidcAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcAuthorizationUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcClientSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcClockTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcDiscoveryUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcDisplayName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcDomainHint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcGroupAttrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcIconUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcIssuer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcJwksUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserOidcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserOidcTokenUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcUserAttrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcUserRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcVerifyCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserOidcVerifyIssuer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserOidc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_method"); ok || d.HasChange("auth_method") {
		t, err := expandObjectUserOidcAuthMethod(d, v, "auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandObjectUserOidcAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("authorization_url"); ok || d.HasChange("authorization_url") {
		t, err := expandObjectUserOidcAuthorizationUrl(d, v, "authorization_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization-url"] = t
		}
	}

	if v, ok := d.GetOk("client_id"); ok || d.HasChange("client_id") {
		t, err := expandObjectUserOidcClientId(d, v, "client_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}

	if v, ok := d.GetOk("client_secret"); ok || d.HasChange("client_secret") {
		t, err := expandObjectUserOidcClientSecret(d, v, "client_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}

	if v, ok := d.GetOk("clock_tolerance"); ok || d.HasChange("clock_tolerance") {
		t, err := expandObjectUserOidcClockTolerance(d, v, "clock_tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("discovery_url"); ok || d.HasChange("discovery_url") {
		t, err := expandObjectUserOidcDiscoveryUrl(d, v, "discovery_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovery-url"] = t
		}
	}

	if v, ok := d.GetOk("display_name"); ok || d.HasChange("display_name") {
		t, err := expandObjectUserOidcDisplayName(d, v, "display_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-name"] = t
		}
	}

	if v, ok := d.GetOk("domain_hint"); ok || d.HasChange("domain_hint") {
		t, err := expandObjectUserOidcDomainHint(d, v, "domain_hint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-hint"] = t
		}
	}

	if v, ok := d.GetOk("group_attr_name"); ok || d.HasChange("group_attr_name") {
		t, err := expandObjectUserOidcGroupAttrName(d, v, "group_attr_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("icon_url"); ok || d.HasChange("icon_url") {
		t, err := expandObjectUserOidcIconUrl(d, v, "icon_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-url"] = t
		}
	}

	if v, ok := d.GetOk("issuer"); ok || d.HasChange("issuer") {
		t, err := expandObjectUserOidcIssuer(d, v, "issuer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["issuer"] = t
		}
	}

	if v, ok := d.GetOk("jwks_uri"); ok || d.HasChange("jwks_uri") {
		t, err := expandObjectUserOidcJwksUri(d, v, "jwks_uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jwks-uri"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserOidcLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserOidcName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok || d.HasChange("private_key") {
		t, err := expandObjectUserOidcPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("token_url"); ok || d.HasChange("token_url") {
		t, err := expandObjectUserOidcTokenUrl(d, v, "token_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-url"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserOidcType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user_attr_name"); ok || d.HasChange("user_attr_name") {
		t, err := expandObjectUserOidcUserAttrName(d, v, "user_attr_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("user_regex"); ok || d.HasChange("user_regex") {
		t, err := expandObjectUserOidcUserRegex(d, v, "user_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-regex"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandObjectUserOidcVerifyCert(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("verify_issuer"); ok || d.HasChange("verify_issuer") {
		t, err := expandObjectUserOidcVerifyIssuer(d, v, "verify_issuer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-issuer"] = t
		}
	}

	return &obj, nil
}
