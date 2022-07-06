// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure local users.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserLocalCreate,
		Read:   resourceObjectUserLocalRead,
		Update: resourceObjectUserLocalUpdate,
		Delete: resourceObjectUserLocalDelete,

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
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortitoken": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"passwd_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tacacs_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_insensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workstation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserLocalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLocal resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserLocal(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLocal resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserLocalRead(d, m)
}

func resourceObjectUserLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLocal resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserLocal(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserLocalRead(d, m)
}

func resourceObjectUserLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserLocal(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserLocal(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLocal resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserLocalAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalPasswdPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalPpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalPpkSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUsernameCaseInsensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUsernameCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUsernameSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalWorkstation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_concurrent_override", flattenObjectUserLocalAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "ObjectUserLocal-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenObjectUserLocalAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "ObjectUserLocal-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenObjectUserLocalAuthtimeout(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "ObjectUserLocal-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("email_to", flattenObjectUserLocalEmailTo(o["email-to"], d, "email_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-to"], "ObjectUserLocal-EmailTo"); ok {
			if err = d.Set("email_to", vv); err != nil {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenObjectUserLocalFortitoken(o["fortitoken"], d, "fortitoken")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken"], "ObjectUserLocal-Fortitoken"); ok {
			if err = d.Set("fortitoken", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserLocalId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserLocal-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserLocalLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserLocal-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("passwd_policy", flattenObjectUserLocalPasswdPolicy(o["passwd-policy"], d, "passwd_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["passwd-policy"], "ObjectUserLocal-PasswdPolicy"); ok {
			if err = d.Set("passwd_policy", vv); err != nil {
				return fmt.Errorf("Error reading passwd_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passwd_policy: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenObjectUserLocalPpkIdentity(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk-identity"], "ObjectUserLocal-PpkIdentity"); ok {
			if err = d.Set("ppk_identity", vv); err != nil {
				return fmt.Errorf("Error reading ppk_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenObjectUserLocalRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "ObjectUserLocal-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenObjectUserLocalSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "ObjectUserLocal-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenObjectUserLocalSmsPhone(o["sms-phone"], d, "sms_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-phone"], "ObjectUserLocal-SmsPhone"); ok {
			if err = d.Set("sms_phone", vv); err != nil {
				return fmt.Errorf("Error reading sms_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenObjectUserLocalSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "ObjectUserLocal-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserLocalStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserLocal-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tacacs_server", flattenObjectUserLocalTacacsServer(o["tacacs+-server"], d, "tacacs_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tacacs+-server"], "ObjectUserLocal-TacacsServer"); ok {
			if err = d.Set("tacacs_server", vv); err != nil {
				return fmt.Errorf("Error reading tacacs_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tacacs_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenObjectUserLocalTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "ObjectUserLocal-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenObjectUserLocalTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "ObjectUserLocal-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenObjectUserLocalTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "ObjectUserLocal-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserLocalType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserLocal-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username_case_insensitivity", flattenObjectUserLocalUsernameCaseInsensitivity(o["username-case-insensitivity"], d, "username_case_insensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-insensitivity"], "ObjectUserLocal-UsernameCaseInsensitivity"); ok {
			if err = d.Set("username_case_insensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
		}
	}

	if err = d.Set("username_case_sensitivity", flattenObjectUserLocalUsernameCaseSensitivity(o["username-case-sensitivity"], d, "username_case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitivity"], "ObjectUserLocal-UsernameCaseSensitivity"); ok {
			if err = d.Set("username_case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
		}
	}

	if err = d.Set("username_sensitivity", flattenObjectUserLocalUsernameSensitivity(o["username-sensitivity"], d, "username_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-sensitivity"], "ObjectUserLocal-UsernameSensitivity"); ok {
			if err = d.Set("username_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_sensitivity: %v", err)
		}
	}

	if err = d.Set("workstation", flattenObjectUserLocalWorkstation(o["workstation"], d, "workstation")); err != nil {
		if vv, ok := fortiAPIPatch(o["workstation"], "ObjectUserLocal-Workstation"); ok {
			if err = d.Set("workstation", vv); err != nil {
				return fmt.Errorf("Error reading workstation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workstation: %v", err)
		}
	}

	return nil
}

func flattenObjectUserLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserLocalAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalFortitoken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalPasswdPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalPpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalPpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalSmsPhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUsernameCaseInsensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUsernameCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUsernameSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalWorkstation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandObjectUserLocalAuthConcurrentOverride(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandObjectUserLocalAuthConcurrentValue(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandObjectUserLocalAuthtimeout(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandObjectUserLocalEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok || d.HasChange("fortitoken") {
		t, err := expandObjectUserLocalFortitoken(d, v, "fortitoken")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectUserLocalId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserLocalLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandObjectUserLocalPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("passwd_policy"); ok || d.HasChange("passwd_policy") {
		t, err := expandObjectUserLocalPasswdPolicy(d, v, "passwd_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-policy"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok || d.HasChange("ppk_identity") {
		t, err := expandObjectUserLocalPpkIdentity(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok || d.HasChange("ppk_secret") {
		t, err := expandObjectUserLocalPpkSecret(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandObjectUserLocalRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandObjectUserLocalSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok || d.HasChange("sms_phone") {
		t, err := expandObjectUserLocalSmsPhone(d, v, "sms_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandObjectUserLocalSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserLocalStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_server"); ok || d.HasChange("tacacs_server") {
		t, err := expandObjectUserLocalTacacsServer(d, v, "tacacs_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs+-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandObjectUserLocalTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandObjectUserLocalTwoFactorAuthentication(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandObjectUserLocalTwoFactorNotification(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserLocalType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username_case_insensitivity"); ok || d.HasChange("username_case_insensitivity") {
		t, err := expandObjectUserLocalUsernameCaseInsensitivity(d, v, "username_case_insensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-insensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitivity"); ok || d.HasChange("username_case_sensitivity") {
		t, err := expandObjectUserLocalUsernameCaseSensitivity(d, v, "username_case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_sensitivity"); ok || d.HasChange("username_sensitivity") {
		t, err := expandObjectUserLocalUsernameSensitivity(d, v, "username_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("workstation"); ok || d.HasChange("workstation") {
		t, err := expandObjectUserLocalWorkstation(d, v, "workstation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workstation"] = t
		}
	}

	return &obj, nil
}
