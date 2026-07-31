// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Authentication Schemes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectAuthenticationScheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectAuthenticationSchemeCreate,
		Read:   resourceObjectAuthenticationSchemeRead,
		Update: resourceObjectAuthenticationSchemeUpdate,
		Delete: resourceObjectAuthenticationSchemeDelete,

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
			"cert_http_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digest_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"digest_rfc2069": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_idp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ems_device_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kerberos_keytab": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"method": &schema.Schema{
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
			"negotiate_ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"require_tfa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssh_ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_database": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_user_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bearer_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bearer_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bearer_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captcha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captcha_secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captcha_site_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"captcha_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oidc_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oidc_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"search_all_ldap_databases": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectAuthenticationSchemeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectAuthenticationScheme resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectAuthenticationScheme(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectAuthenticationScheme(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectAuthenticationScheme resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectAuthenticationScheme(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectAuthenticationScheme resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectAuthenticationSchemeRead(d, m)
}

func resourceObjectAuthenticationSchemeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAuthenticationScheme resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectAuthenticationScheme(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAuthenticationScheme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectAuthenticationSchemeRead(d, m)
}

func resourceObjectAuthenticationSchemeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectAuthenticationScheme(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectAuthenticationScheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectAuthenticationSchemeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectAuthenticationScheme(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectAuthenticationScheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectAuthenticationScheme(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAuthenticationScheme resource from API: %v", err)
	}
	return nil
}

func flattenObjectAuthenticationSchemeCertHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeDigestAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAuthenticationSchemeDigestRfc2069(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAuthenticationSchemeExternalIdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAuthenticationSchemeEmsDeviceOwner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAuthenticationSchemeFssoGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeKerberosKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAuthenticationSchemeMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAuthenticationSchemeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeNegotiateNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeRequireTfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAuthenticationSchemeSamlTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeSshCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectAuthenticationSchemeUserCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeUserDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAuthenticationSchemeAuthUserHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeBearerFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeBearerHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeBearerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeCaptcha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeCaptchaSecretKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeCaptchaSiteKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeCaptchaVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeOidcServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectAuthenticationSchemeOidcTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeSearchAllLdapDatabases(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectAuthenticationScheme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cert_http_header", flattenObjectAuthenticationSchemeCertHttpHeader(o["cert-http-header"], d, "cert_http_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert-http-header"], "ObjectAuthenticationScheme-CertHttpHeader"); ok {
			if err = d.Set("cert_http_header", vv); err != nil {
				return fmt.Errorf("Error reading cert_http_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert_http_header: %v", err)
		}
	}

	if err = d.Set("digest_algo", flattenObjectAuthenticationSchemeDigestAlgo(o["digest-algo"], d, "digest_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-algo"], "ObjectAuthenticationScheme-DigestAlgo"); ok {
			if err = d.Set("digest_algo", vv); err != nil {
				return fmt.Errorf("Error reading digest_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_algo: %v", err)
		}
	}

	if err = d.Set("digest_rfc2069", flattenObjectAuthenticationSchemeDigestRfc2069(o["digest-rfc2069"], d, "digest_rfc2069")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-rfc2069"], "ObjectAuthenticationScheme-DigestRfc2069"); ok {
			if err = d.Set("digest_rfc2069", vv); err != nil {
				return fmt.Errorf("Error reading digest_rfc2069: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_rfc2069: %v", err)
		}
	}

	if err = d.Set("domain_controller", flattenObjectAuthenticationSchemeDomainController(o["domain-controller"], d, "domain_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-controller"], "ObjectAuthenticationScheme-DomainController"); ok {
			if err = d.Set("domain_controller", vv); err != nil {
				return fmt.Errorf("Error reading domain_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if err = d.Set("external_idp", flattenObjectAuthenticationSchemeExternalIdp(o["external-idp"], d, "external_idp")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-idp"], "ObjectAuthenticationScheme-ExternalIdp"); ok {
			if err = d.Set("external_idp", vv); err != nil {
				return fmt.Errorf("Error reading external_idp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_idp: %v", err)
		}
	}

	if err = d.Set("ems_device_owner", flattenObjectAuthenticationSchemeEmsDeviceOwner(o["ems-device-owner"], d, "ems_device_owner")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-device-owner"], "ObjectAuthenticationScheme-EmsDeviceOwner"); ok {
			if err = d.Set("ems_device_owner", vv); err != nil {
				return fmt.Errorf("Error reading ems_device_owner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_device_owner: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenObjectAuthenticationSchemeFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-agent-for-ntlm"], "ObjectAuthenticationScheme-FssoAgentForNtlm"); ok {
			if err = d.Set("fsso_agent_for_ntlm", vv); err != nil {
				return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("fsso_guest", flattenObjectAuthenticationSchemeFssoGuest(o["fsso-guest"], d, "fsso_guest")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-guest"], "ObjectAuthenticationScheme-FssoGuest"); ok {
			if err = d.Set("fsso_guest", vv); err != nil {
				return fmt.Errorf("Error reading fsso_guest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_guest: %v", err)
		}
	}

	if err = d.Set("group_attr_type", flattenObjectAuthenticationSchemeGroupAttrType(o["group-attr-type"], d, "group_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-attr-type"], "ObjectAuthenticationScheme-GroupAttrType"); ok {
			if err = d.Set("group_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading group_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_attr_type: %v", err)
		}
	}

	if err = d.Set("kerberos_keytab", flattenObjectAuthenticationSchemeKerberosKeytab(o["kerberos-keytab"], d, "kerberos_keytab")); err != nil {
		if vv, ok := fortiAPIPatch(o["kerberos-keytab"], "ObjectAuthenticationScheme-KerberosKeytab"); ok {
			if err = d.Set("kerberos_keytab", vv); err != nil {
				return fmt.Errorf("Error reading kerberos_keytab: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading kerberos_keytab: %v", err)
		}
	}

	if err = d.Set("method", flattenObjectAuthenticationSchemeMethod(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "ObjectAuthenticationScheme-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectAuthenticationSchemeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectAuthenticationScheme-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("negotiate_ntlm", flattenObjectAuthenticationSchemeNegotiateNtlm(o["negotiate-ntlm"], d, "negotiate_ntlm")); err != nil {
		if vv, ok := fortiAPIPatch(o["negotiate-ntlm"], "ObjectAuthenticationScheme-NegotiateNtlm"); ok {
			if err = d.Set("negotiate_ntlm", vv); err != nil {
				return fmt.Errorf("Error reading negotiate_ntlm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negotiate_ntlm: %v", err)
		}
	}

	if err = d.Set("require_tfa", flattenObjectAuthenticationSchemeRequireTfa(o["require-tfa"], d, "require_tfa")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-tfa"], "ObjectAuthenticationScheme-RequireTfa"); ok {
			if err = d.Set("require_tfa", vv); err != nil {
				return fmt.Errorf("Error reading require_tfa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_tfa: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenObjectAuthenticationSchemeSamlServer(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "ObjectAuthenticationScheme-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("saml_timeout", flattenObjectAuthenticationSchemeSamlTimeout(o["saml-timeout"], d, "saml_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-timeout"], "ObjectAuthenticationScheme-SamlTimeout"); ok {
			if err = d.Set("saml_timeout", vv); err != nil {
				return fmt.Errorf("Error reading saml_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_timeout: %v", err)
		}
	}

	if err = d.Set("ssh_ca", flattenObjectAuthenticationSchemeSshCa(o["ssh-ca"], d, "ssh_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-ca"], "ObjectAuthenticationScheme-SshCa"); ok {
			if err = d.Set("ssh_ca", vv); err != nil {
				return fmt.Errorf("Error reading ssh_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_ca: %v", err)
		}
	}

	if err = d.Set("user_cert", flattenObjectAuthenticationSchemeUserCert(o["user-cert"], d, "user_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-cert"], "ObjectAuthenticationScheme-UserCert"); ok {
			if err = d.Set("user_cert", vv); err != nil {
				return fmt.Errorf("Error reading user_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_cert: %v", err)
		}
	}

	if err = d.Set("user_database", flattenObjectAuthenticationSchemeUserDatabase(o["user-database"], d, "user_database")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-database"], "ObjectAuthenticationScheme-UserDatabase"); ok {
			if err = d.Set("user_database", vv); err != nil {
				return fmt.Errorf("Error reading user_database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_database: %v", err)
		}
	}

	if err = d.Set("auth_user_header", flattenObjectAuthenticationSchemeAuthUserHeader(o["auth-user-header"], d, "auth_user_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-user-header"], "ObjectAuthenticationScheme-AuthUserHeader"); ok {
			if err = d.Set("auth_user_header", vv); err != nil {
				return fmt.Errorf("Error reading auth_user_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_user_header: %v", err)
		}
	}

	if err = d.Set("bearer_format", flattenObjectAuthenticationSchemeBearerFormat(o["bearer-format"], d, "bearer_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["bearer-format"], "ObjectAuthenticationScheme-BearerFormat"); ok {
			if err = d.Set("bearer_format", vv); err != nil {
				return fmt.Errorf("Error reading bearer_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bearer_format: %v", err)
		}
	}

	if err = d.Set("bearer_header", flattenObjectAuthenticationSchemeBearerHeader(o["bearer-header"], d, "bearer_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["bearer-header"], "ObjectAuthenticationScheme-BearerHeader"); ok {
			if err = d.Set("bearer_header", vv); err != nil {
				return fmt.Errorf("Error reading bearer_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bearer_header: %v", err)
		}
	}

	if err = d.Set("bearer_type", flattenObjectAuthenticationSchemeBearerType(o["bearer-type"], d, "bearer_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["bearer-type"], "ObjectAuthenticationScheme-BearerType"); ok {
			if err = d.Set("bearer_type", vv); err != nil {
				return fmt.Errorf("Error reading bearer_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bearer_type: %v", err)
		}
	}

	if err = d.Set("captcha", flattenObjectAuthenticationSchemeCaptcha(o["captcha"], d, "captcha")); err != nil {
		if vv, ok := fortiAPIPatch(o["captcha"], "ObjectAuthenticationScheme-Captcha"); ok {
			if err = d.Set("captcha", vv); err != nil {
				return fmt.Errorf("Error reading captcha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captcha: %v", err)
		}
	}

	if err = d.Set("captcha_secret_key", flattenObjectAuthenticationSchemeCaptchaSecretKey(o["captcha-secret-key"], d, "captcha_secret_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["captcha-secret-key"], "ObjectAuthenticationScheme-CaptchaSecretKey"); ok {
			if err = d.Set("captcha_secret_key", vv); err != nil {
				return fmt.Errorf("Error reading captcha_secret_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captcha_secret_key: %v", err)
		}
	}

	if err = d.Set("captcha_site_key", flattenObjectAuthenticationSchemeCaptchaSiteKey(o["captcha-site-key"], d, "captcha_site_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["captcha-site-key"], "ObjectAuthenticationScheme-CaptchaSiteKey"); ok {
			if err = d.Set("captcha_site_key", vv); err != nil {
				return fmt.Errorf("Error reading captcha_site_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captcha_site_key: %v", err)
		}
	}

	if err = d.Set("captcha_vendor", flattenObjectAuthenticationSchemeCaptchaVendor(o["captcha-vendor"], d, "captcha_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["captcha-vendor"], "ObjectAuthenticationScheme-CaptchaVendor"); ok {
			if err = d.Set("captcha_vendor", vv); err != nil {
				return fmt.Errorf("Error reading captcha_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captcha_vendor: %v", err)
		}
	}

	if err = d.Set("oidc_server", flattenObjectAuthenticationSchemeOidcServer(o["oidc-server"], d, "oidc_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["oidc-server"], "ObjectAuthenticationScheme-OidcServer"); ok {
			if err = d.Set("oidc_server", vv); err != nil {
				return fmt.Errorf("Error reading oidc_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oidc_server: %v", err)
		}
	}

	if err = d.Set("oidc_timeout", flattenObjectAuthenticationSchemeOidcTimeout(o["oidc-timeout"], d, "oidc_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["oidc-timeout"], "ObjectAuthenticationScheme-OidcTimeout"); ok {
			if err = d.Set("oidc_timeout", vv); err != nil {
				return fmt.Errorf("Error reading oidc_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oidc_timeout: %v", err)
		}
	}

	if err = d.Set("search_all_ldap_databases", flattenObjectAuthenticationSchemeSearchAllLdapDatabases(o["search-all-ldap-databases"], d, "search_all_ldap_databases")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-all-ldap-databases"], "ObjectAuthenticationScheme-SearchAllLdapDatabases"); ok {
			if err = d.Set("search_all_ldap_databases", vv); err != nil {
				return fmt.Errorf("Error reading search_all_ldap_databases: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_all_ldap_databases: %v", err)
		}
	}

	return nil
}

func flattenObjectAuthenticationSchemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAuthenticationSchemeCertHttpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeDigestAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAuthenticationSchemeDigestRfc2069(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAuthenticationSchemeExternalIdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAuthenticationSchemeEmsDeviceOwner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAuthenticationSchemeFssoGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeGroupAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeKerberosKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAuthenticationSchemeMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAuthenticationSchemeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeNegotiateNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeRequireTfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeSamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAuthenticationSchemeSamlTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeSshCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectAuthenticationSchemeUserCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeUserDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAuthenticationSchemeAuthUserHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeBearerFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeBearerHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeBearerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeCaptcha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeCaptchaSecretKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeCaptchaSiteKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeCaptchaVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeOidcServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectAuthenticationSchemeOidcTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeSearchAllLdapDatabases(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectAuthenticationScheme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert_http_header"); ok || d.HasChange("cert_http_header") {
		t, err := expandObjectAuthenticationSchemeCertHttpHeader(d, v, "cert_http_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-http-header"] = t
		}
	}

	if v, ok := d.GetOk("digest_algo"); ok || d.HasChange("digest_algo") {
		t, err := expandObjectAuthenticationSchemeDigestAlgo(d, v, "digest_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-algo"] = t
		}
	}

	if v, ok := d.GetOk("digest_rfc2069"); ok || d.HasChange("digest_rfc2069") {
		t, err := expandObjectAuthenticationSchemeDigestRfc2069(d, v, "digest_rfc2069")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-rfc2069"] = t
		}
	}

	if v, ok := d.GetOk("domain_controller"); ok || d.HasChange("domain_controller") {
		t, err := expandObjectAuthenticationSchemeDomainController(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	}

	if v, ok := d.GetOk("external_idp"); ok || d.HasChange("external_idp") {
		t, err := expandObjectAuthenticationSchemeExternalIdp(d, v, "external_idp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-idp"] = t
		}
	}

	if v, ok := d.GetOk("ems_device_owner"); ok || d.HasChange("ems_device_owner") {
		t, err := expandObjectAuthenticationSchemeEmsDeviceOwner(d, v, "ems_device_owner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-device-owner"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok || d.HasChange("fsso_agent_for_ntlm") {
		t, err := expandObjectAuthenticationSchemeFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("fsso_guest"); ok || d.HasChange("fsso_guest") {
		t, err := expandObjectAuthenticationSchemeFssoGuest(d, v, "fsso_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-guest"] = t
		}
	}

	if v, ok := d.GetOk("group_attr_type"); ok || d.HasChange("group_attr_type") {
		t, err := expandObjectAuthenticationSchemeGroupAttrType(d, v, "group_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("kerberos_keytab"); ok || d.HasChange("kerberos_keytab") {
		t, err := expandObjectAuthenticationSchemeKerberosKeytab(d, v, "kerberos_keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kerberos-keytab"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandObjectAuthenticationSchemeMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectAuthenticationSchemeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_ntlm"); ok || d.HasChange("negotiate_ntlm") {
		t, err := expandObjectAuthenticationSchemeNegotiateNtlm(d, v, "negotiate_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("require_tfa"); ok || d.HasChange("require_tfa") {
		t, err := expandObjectAuthenticationSchemeRequireTfa(d, v, "require_tfa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-tfa"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandObjectAuthenticationSchemeSamlServer(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("saml_timeout"); ok || d.HasChange("saml_timeout") {
		t, err := expandObjectAuthenticationSchemeSamlTimeout(d, v, "saml_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssh_ca"); ok || d.HasChange("ssh_ca") {
		t, err := expandObjectAuthenticationSchemeSshCa(d, v, "ssh_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-ca"] = t
		}
	}

	if v, ok := d.GetOk("user_cert"); ok || d.HasChange("user_cert") {
		t, err := expandObjectAuthenticationSchemeUserCert(d, v, "user_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_database"); ok || d.HasChange("user_database") {
		t, err := expandObjectAuthenticationSchemeUserDatabase(d, v, "user_database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-database"] = t
		}
	}

	if v, ok := d.GetOk("auth_user_header"); ok || d.HasChange("auth_user_header") {
		t, err := expandObjectAuthenticationSchemeAuthUserHeader(d, v, "auth_user_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-user-header"] = t
		}
	}

	if v, ok := d.GetOk("bearer_format"); ok || d.HasChange("bearer_format") {
		t, err := expandObjectAuthenticationSchemeBearerFormat(d, v, "bearer_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bearer-format"] = t
		}
	}

	if v, ok := d.GetOk("bearer_header"); ok || d.HasChange("bearer_header") {
		t, err := expandObjectAuthenticationSchemeBearerHeader(d, v, "bearer_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bearer-header"] = t
		}
	}

	if v, ok := d.GetOk("bearer_type"); ok || d.HasChange("bearer_type") {
		t, err := expandObjectAuthenticationSchemeBearerType(d, v, "bearer_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bearer-type"] = t
		}
	}

	if v, ok := d.GetOk("captcha"); ok || d.HasChange("captcha") {
		t, err := expandObjectAuthenticationSchemeCaptcha(d, v, "captcha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captcha"] = t
		}
	}

	if v, ok := d.GetOk("captcha_secret_key"); ok || d.HasChange("captcha_secret_key") {
		t, err := expandObjectAuthenticationSchemeCaptchaSecretKey(d, v, "captcha_secret_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captcha-secret-key"] = t
		}
	}

	if v, ok := d.GetOk("captcha_site_key"); ok || d.HasChange("captcha_site_key") {
		t, err := expandObjectAuthenticationSchemeCaptchaSiteKey(d, v, "captcha_site_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captcha-site-key"] = t
		}
	}

	if v, ok := d.GetOk("captcha_vendor"); ok || d.HasChange("captcha_vendor") {
		t, err := expandObjectAuthenticationSchemeCaptchaVendor(d, v, "captcha_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captcha-vendor"] = t
		}
	}

	if v, ok := d.GetOk("oidc_server"); ok || d.HasChange("oidc_server") {
		t, err := expandObjectAuthenticationSchemeOidcServer(d, v, "oidc_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oidc-server"] = t
		}
	}

	if v, ok := d.GetOk("oidc_timeout"); ok || d.HasChange("oidc_timeout") {
		t, err := expandObjectAuthenticationSchemeOidcTimeout(d, v, "oidc_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oidc-timeout"] = t
		}
	}

	if v, ok := d.GetOk("search_all_ldap_databases"); ok || d.HasChange("search_all_ldap_databases") {
		t, err := expandObjectAuthenticationSchemeSearchAllLdapDatabases(d, v, "search_all_ldap_databases")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-all-ldap-databases"] = t
		}
	}

	return &obj, nil
}
