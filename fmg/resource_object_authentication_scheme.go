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
			"domain_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"kerberos_keytab": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceObjectAuthenticationSchemeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectAuthenticationScheme resource while getting object: %v", err)
	}

	_, err = c.CreateObjectAuthenticationScheme(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectAuthenticationScheme resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectAuthenticationSchemeRead(d, m)
}

func resourceObjectAuthenticationSchemeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAuthenticationScheme resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectAuthenticationScheme(obj, adomv, mkey, nil)
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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectAuthenticationScheme(adomv, mkey, nil)
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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectAuthenticationScheme(adomv, mkey, nil)
	if err != nil {
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

func flattenObjectAuthenticationSchemeDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeEmsDeviceOwner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeFssoGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeKerberosKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
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
	return v
}

func flattenObjectAuthenticationSchemeSamlTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeSshCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeUserCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAuthenticationSchemeUserDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectAuthenticationScheme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
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

	return nil
}

func flattenObjectAuthenticationSchemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAuthenticationSchemeDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeEmsDeviceOwner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeFssoGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeKerberosKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
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
	return v, nil
}

func expandObjectAuthenticationSchemeSamlTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeSshCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeUserCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAuthenticationSchemeUserDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectObjectAuthenticationScheme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain_controller"); ok || d.HasChange("domain_controller") {
		t, err := expandObjectAuthenticationSchemeDomainController(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
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

	return &obj, nil
}
