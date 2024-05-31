// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure peer users.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserPeerCreate,
		Read:   resourceObjectUserPeerRead,
		Update: resourceObjectUserPeerUpdate,
		Delete: resourceObjectUserPeerDelete,

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
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cn_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mandatory_ca_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mfa_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mfa_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"mfa_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mfa_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ocsp_override_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPeer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserPeer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPeer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserPeerRead(d, m)
}

func resourceObjectUserPeerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPeer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserPeer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserPeerRead(d, m)
}

func resourceObjectUserPeerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserPeer resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserPeerCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserPeerCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerCnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerLdapMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserPeerLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerMandatoryCaVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerMfaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerMfaServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserPeerMfaUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerOcspOverrideServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserPeerSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPeerTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ca", flattenObjectUserPeerCa(o["ca"], d, "ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca"], "ObjectUserPeer-Ca"); ok {
			if err = d.Set("ca", vv); err != nil {
				return fmt.Errorf("Error reading ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("cn", flattenObjectUserPeerCn(o["cn"], d, "cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["cn"], "ObjectUserPeer-Cn"); ok {
			if err = d.Set("cn", vv); err != nil {
				return fmt.Errorf("Error reading cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cn: %v", err)
		}
	}

	if err = d.Set("cn_type", flattenObjectUserPeerCnType(o["cn-type"], d, "cn_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["cn-type"], "ObjectUserPeer-CnType"); ok {
			if err = d.Set("cn_type", vv); err != nil {
				return fmt.Errorf("Error reading cn_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cn_type: %v", err)
		}
	}

	if err = d.Set("ldap_mode", flattenObjectUserPeerLdapMode(o["ldap-mode"], d, "ldap_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-mode"], "ObjectUserPeer-LdapMode"); ok {
			if err = d.Set("ldap_mode", vv); err != nil {
				return fmt.Errorf("Error reading ldap_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_mode: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserPeerLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserPeer-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenObjectUserPeerLdapUsername(o["ldap-username"], d, "ldap_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-username"], "ObjectUserPeer-LdapUsername"); ok {
			if err = d.Set("ldap_username", vv); err != nil {
				return fmt.Errorf("Error reading ldap_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("mandatory_ca_verify", flattenObjectUserPeerMandatoryCaVerify(o["mandatory-ca-verify"], d, "mandatory_ca_verify")); err != nil {
		if vv, ok := fortiAPIPatch(o["mandatory-ca-verify"], "ObjectUserPeer-MandatoryCaVerify"); ok {
			if err = d.Set("mandatory_ca_verify", vv); err != nil {
				return fmt.Errorf("Error reading mandatory_ca_verify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mandatory_ca_verify: %v", err)
		}
	}

	if err = d.Set("mfa_mode", flattenObjectUserPeerMfaMode(o["mfa-mode"], d, "mfa_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mfa-mode"], "ObjectUserPeer-MfaMode"); ok {
			if err = d.Set("mfa_mode", vv); err != nil {
				return fmt.Errorf("Error reading mfa_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mfa_mode: %v", err)
		}
	}

	if err = d.Set("mfa_server", flattenObjectUserPeerMfaServer(o["mfa-server"], d, "mfa_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["mfa-server"], "ObjectUserPeer-MfaServer"); ok {
			if err = d.Set("mfa_server", vv); err != nil {
				return fmt.Errorf("Error reading mfa_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mfa_server: %v", err)
		}
	}

	if err = d.Set("mfa_username", flattenObjectUserPeerMfaUsername(o["mfa-username"], d, "mfa_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["mfa-username"], "ObjectUserPeer-MfaUsername"); ok {
			if err = d.Set("mfa_username", vv); err != nil {
				return fmt.Errorf("Error reading mfa_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mfa_username: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserPeerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserPeer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ocsp_override_server", flattenObjectUserPeerOcspOverrideServer(o["ocsp-override-server"], d, "ocsp_override_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocsp-override-server"], "ObjectUserPeer-OcspOverrideServer"); ok {
			if err = d.Set("ocsp_override_server", vv); err != nil {
				return fmt.Errorf("Error reading ocsp_override_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocsp_override_server: %v", err)
		}
	}

	if err = d.Set("subject", flattenObjectUserPeerSubject(o["subject"], d, "subject")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject"], "ObjectUserPeer-Subject"); ok {
			if err = d.Set("subject", vv); err != nil {
				return fmt.Errorf("Error reading subject: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenObjectUserPeerTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "ObjectUserPeer-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	return nil
}

func flattenObjectUserPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserPeerCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserPeerCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerCnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerLdapMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserPeerLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserPeerLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerMandatoryCaVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerMfaMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerMfaPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserPeerMfaServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserPeerMfaUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerOcspOverrideServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserPeerPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserPeerSubject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPeerTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok || d.HasChange("ca") {
		t, err := expandObjectUserPeerCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("cn"); ok || d.HasChange("cn") {
		t, err := expandObjectUserPeerCn(d, v, "cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn"] = t
		}
	}

	if v, ok := d.GetOk("cn_type"); ok || d.HasChange("cn_type") {
		t, err := expandObjectUserPeerCnType(d, v, "cn_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-type"] = t
		}
	}

	if v, ok := d.GetOk("ldap_mode"); ok || d.HasChange("ldap_mode") {
		t, err := expandObjectUserPeerLdapMode(d, v, "ldap_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-mode"] = t
		}
	}

	if v, ok := d.GetOk("ldap_password"); ok || d.HasChange("ldap_password") {
		t, err := expandObjectUserPeerLdapPassword(d, v, "ldap_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserPeerLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("ldap_username"); ok || d.HasChange("ldap_username") {
		t, err := expandObjectUserPeerLdapUsername(d, v, "ldap_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	}

	if v, ok := d.GetOk("mandatory_ca_verify"); ok || d.HasChange("mandatory_ca_verify") {
		t, err := expandObjectUserPeerMandatoryCaVerify(d, v, "mandatory_ca_verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mandatory-ca-verify"] = t
		}
	}

	if v, ok := d.GetOk("mfa_mode"); ok || d.HasChange("mfa_mode") {
		t, err := expandObjectUserPeerMfaMode(d, v, "mfa_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-mode"] = t
		}
	}

	if v, ok := d.GetOk("mfa_password"); ok || d.HasChange("mfa_password") {
		t, err := expandObjectUserPeerMfaPassword(d, v, "mfa_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-password"] = t
		}
	}

	if v, ok := d.GetOk("mfa_server"); ok || d.HasChange("mfa_server") {
		t, err := expandObjectUserPeerMfaServer(d, v, "mfa_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-server"] = t
		}
	}

	if v, ok := d.GetOk("mfa_username"); ok || d.HasChange("mfa_username") {
		t, err := expandObjectUserPeerMfaUsername(d, v, "mfa_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-username"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserPeerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_override_server"); ok || d.HasChange("ocsp_override_server") {
		t, err := expandObjectUserPeerOcspOverrideServer(d, v, "ocsp_override_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-override-server"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandObjectUserPeerPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("subject"); ok || d.HasChange("subject") {
		t, err := expandObjectUserPeerSubject(d, v, "subject")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandObjectUserPeerTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	return &obj, nil
}
