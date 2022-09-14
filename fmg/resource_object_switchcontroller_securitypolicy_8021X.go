// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure 802.1x MAC Authentication Bypass (MAB) policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSwitchControllerSecurityPolicy8021X() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerSecurityPolicy8021XCreate,
		Read:   resourceObjectSwitchControllerSecurityPolicy8021XRead,
		Update: resourceObjectSwitchControllerSecurityPolicy8021XUpdate,
		Delete: resourceObjectSwitchControllerSecurityPolicy8021XDelete,

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
			"auth_fail_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_fail_vlan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_fail_vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authserver_timeout_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authserver_timeout_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authserver_timeout_vlanid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap_auto_untagged_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_passthru": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"framevid_apply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_auth_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"guest_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_vlan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guest_vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"open_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_timeout_overwrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerSecurityPolicy8021XCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerSecurityPolicy8021X(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerSecurityPolicy8021X resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerSecurityPolicy8021X(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSecurityPolicy8021XRead(d, m)
}

func resourceObjectSwitchControllerSecurityPolicy8021XUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerSecurityPolicy8021X(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicy8021X resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerSecurityPolicy8021X(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSecurityPolicy8021XRead(d, m)
}

func resourceObjectSwitchControllerSecurityPolicy8021XDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerSecurityPolicy8021X(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerSecurityPolicy8021XRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerSecurityPolicy8021X(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerSecurityPolicy8021X(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerSecurityPolicy8021X resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerSecurityPolicy8021XAuthFailVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XAuthFailVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XAuthFailVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XEapPassthru(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XFramevidApply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XGuestAuthDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XGuestVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XGuestVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XGuestVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XOpenAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerSecurityPolicy8021XUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_fail_vlan", flattenObjectSwitchControllerSecurityPolicy8021XAuthFailVlan(o["auth-fail-vlan"], d, "auth_fail_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-fail-vlan"], "ObjectSwitchControllerSecurityPolicy8021X-AuthFailVlan"); ok {
			if err = d.Set("auth_fail_vlan", vv); err != nil {
				return fmt.Errorf("Error reading auth_fail_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_fail_vlan: %v", err)
		}
	}

	if err = d.Set("auth_fail_vlan_id", flattenObjectSwitchControllerSecurityPolicy8021XAuthFailVlanId(o["auth-fail-vlan-id"], d, "auth_fail_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-fail-vlan-id"], "ObjectSwitchControllerSecurityPolicy8021X-AuthFailVlanId"); ok {
			if err = d.Set("auth_fail_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading auth_fail_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_fail_vlan_id: %v", err)
		}
	}

	if err = d.Set("auth_fail_vlanid", flattenObjectSwitchControllerSecurityPolicy8021XAuthFailVlanid(o["auth-fail-vlanid"], d, "auth_fail_vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-fail-vlanid"], "ObjectSwitchControllerSecurityPolicy8021X-AuthFailVlanid"); ok {
			if err = d.Set("auth_fail_vlanid", vv); err != nil {
				return fmt.Errorf("Error reading auth_fail_vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_fail_vlanid: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_period", flattenObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(o["authserver-timeout-period"], d, "authserver_timeout_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-period"], "ObjectSwitchControllerSecurityPolicy8021X-AuthserverTimeoutPeriod"); ok {
			if err = d.Set("authserver_timeout_period", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_period: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_vlan", flattenObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(o["authserver-timeout-vlan"], d, "authserver_timeout_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-vlan"], "ObjectSwitchControllerSecurityPolicy8021X-AuthserverTimeoutVlan"); ok {
			if err = d.Set("authserver_timeout_vlan", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_vlan: %v", err)
		}
	}

	if err = d.Set("authserver_timeout_vlanid", flattenObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(o["authserver-timeout-vlanid"], d, "authserver_timeout_vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["authserver-timeout-vlanid"], "ObjectSwitchControllerSecurityPolicy8021X-AuthserverTimeoutVlanid"); ok {
			if err = d.Set("authserver_timeout_vlanid", vv); err != nil {
				return fmt.Errorf("Error reading authserver_timeout_vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authserver_timeout_vlanid: %v", err)
		}
	}

	if err = d.Set("eap_auto_untagged_vlans", flattenObjectSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(o["eap-auto-untagged-vlans"], d, "eap_auto_untagged_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-auto-untagged-vlans"], "ObjectSwitchControllerSecurityPolicy8021X-EapAutoUntaggedVlans"); ok {
			if err = d.Set("eap_auto_untagged_vlans", vv); err != nil {
				return fmt.Errorf("Error reading eap_auto_untagged_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_auto_untagged_vlans: %v", err)
		}
	}

	if err = d.Set("eap_passthru", flattenObjectSwitchControllerSecurityPolicy8021XEapPassthru(o["eap-passthru"], d, "eap_passthru")); err != nil {
		if vv, ok := fortiAPIPatch(o["eap-passthru"], "ObjectSwitchControllerSecurityPolicy8021X-EapPassthru"); ok {
			if err = d.Set("eap_passthru", vv); err != nil {
				return fmt.Errorf("Error reading eap_passthru: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eap_passthru: %v", err)
		}
	}

	if err = d.Set("framevid_apply", flattenObjectSwitchControllerSecurityPolicy8021XFramevidApply(o["framevid-apply"], d, "framevid_apply")); err != nil {
		if vv, ok := fortiAPIPatch(o["framevid-apply"], "ObjectSwitchControllerSecurityPolicy8021X-FramevidApply"); ok {
			if err = d.Set("framevid_apply", vv); err != nil {
				return fmt.Errorf("Error reading framevid_apply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading framevid_apply: %v", err)
		}
	}

	if err = d.Set("guest_auth_delay", flattenObjectSwitchControllerSecurityPolicy8021XGuestAuthDelay(o["guest-auth-delay"], d, "guest_auth_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-auth-delay"], "ObjectSwitchControllerSecurityPolicy8021X-GuestAuthDelay"); ok {
			if err = d.Set("guest_auth_delay", vv); err != nil {
				return fmt.Errorf("Error reading guest_auth_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_auth_delay: %v", err)
		}
	}

	if err = d.Set("guest_vlan", flattenObjectSwitchControllerSecurityPolicy8021XGuestVlan(o["guest-vlan"], d, "guest_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-vlan"], "ObjectSwitchControllerSecurityPolicy8021X-GuestVlan"); ok {
			if err = d.Set("guest_vlan", vv); err != nil {
				return fmt.Errorf("Error reading guest_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_vlan: %v", err)
		}
	}

	if err = d.Set("guest_vlan_id", flattenObjectSwitchControllerSecurityPolicy8021XGuestVlanId(o["guest-vlan-id"], d, "guest_vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-vlan-id"], "ObjectSwitchControllerSecurityPolicy8021X-GuestVlanId"); ok {
			if err = d.Set("guest_vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading guest_vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_vlan_id: %v", err)
		}
	}

	if err = d.Set("guest_vlanid", flattenObjectSwitchControllerSecurityPolicy8021XGuestVlanid(o["guest-vlanid"], d, "guest_vlanid")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-vlanid"], "ObjectSwitchControllerSecurityPolicy8021X-GuestVlanid"); ok {
			if err = d.Set("guest_vlanid", vv); err != nil {
				return fmt.Errorf("Error reading guest_vlanid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_vlanid: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenObjectSwitchControllerSecurityPolicy8021XMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-auth-bypass"], "ObjectSwitchControllerSecurityPolicy8021X-MacAuthBypass"); ok {
			if err = d.Set("mac_auth_bypass", vv); err != nil {
				return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerSecurityPolicy8021XName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerSecurityPolicy8021X-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("open_auth", flattenObjectSwitchControllerSecurityPolicy8021XOpenAuth(o["open-auth"], d, "open_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["open-auth"], "ObjectSwitchControllerSecurityPolicy8021X-OpenAuth"); ok {
			if err = d.Set("open_auth", vv); err != nil {
				return fmt.Errorf("Error reading open_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading open_auth: %v", err)
		}
	}

	if err = d.Set("policy_type", flattenObjectSwitchControllerSecurityPolicy8021XPolicyType(o["policy-type"], d, "policy_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-type"], "ObjectSwitchControllerSecurityPolicy8021X-PolicyType"); ok {
			if err = d.Set("policy_type", vv); err != nil {
				return fmt.Errorf("Error reading policy_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_type: %v", err)
		}
	}

	if err = d.Set("radius_timeout_overwrite", flattenObjectSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(o["radius-timeout-overwrite"], d, "radius_timeout_overwrite")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-timeout-overwrite"], "ObjectSwitchControllerSecurityPolicy8021X-RadiusTimeoutOverwrite"); ok {
			if err = d.Set("radius_timeout_overwrite", vv); err != nil {
				return fmt.Errorf("Error reading radius_timeout_overwrite: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_timeout_overwrite: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenObjectSwitchControllerSecurityPolicy8021XSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "ObjectSwitchControllerSecurityPolicy8021X-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("user_group", flattenObjectSwitchControllerSecurityPolicy8021XUserGroup(o["user-group"], d, "user_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group"], "ObjectSwitchControllerSecurityPolicy8021X-UserGroup"); ok {
			if err = d.Set("user_group", vv); err != nil {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerSecurityPolicy8021XFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerSecurityPolicy8021XAuthFailVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XAuthFailVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XAuthFailVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XEapPassthru(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XFramevidApply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XGuestAuthDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XGuestVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XGuestVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XGuestVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XOpenAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerSecurityPolicy8021XUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_fail_vlan"); ok || d.HasChange("auth_fail_vlan") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XAuthFailVlan(d, v, "auth_fail_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlan"] = t
		}
	}

	if v, ok := d.GetOk("auth_fail_vlan_id"); ok || d.HasChange("auth_fail_vlan_id") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XAuthFailVlanId(d, v, "auth_fail_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("auth_fail_vlanid"); ok || d.HasChange("auth_fail_vlanid") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XAuthFailVlanid(d, v, "auth_fail_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_period"); ok || d.HasChange("authserver_timeout_period") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutPeriod(d, v, "authserver_timeout_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-period"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_vlan"); ok || d.HasChange("authserver_timeout_vlan") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlan(d, v, "authserver_timeout_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-vlan"] = t
		}
	}

	if v, ok := d.GetOk("authserver_timeout_vlanid"); ok || d.HasChange("authserver_timeout_vlanid") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XAuthserverTimeoutVlanid(d, v, "authserver_timeout_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authserver-timeout-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("eap_auto_untagged_vlans"); ok || d.HasChange("eap_auto_untagged_vlans") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XEapAutoUntaggedVlans(d, v, "eap_auto_untagged_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-auto-untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("eap_passthru"); ok || d.HasChange("eap_passthru") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XEapPassthru(d, v, "eap_passthru")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-passthru"] = t
		}
	}

	if v, ok := d.GetOk("framevid_apply"); ok || d.HasChange("framevid_apply") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XFramevidApply(d, v, "framevid_apply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["framevid-apply"] = t
		}
	}

	if v, ok := d.GetOk("guest_auth_delay"); ok || d.HasChange("guest_auth_delay") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XGuestAuthDelay(d, v, "guest_auth_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-auth-delay"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlan"); ok || d.HasChange("guest_vlan") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XGuestVlan(d, v, "guest_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlan"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlan_id"); ok || d.HasChange("guest_vlan_id") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XGuestVlanId(d, v, "guest_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlanid"); ok || d.HasChange("guest_vlanid") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XGuestVlanid(d, v, "guest_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok || d.HasChange("mac_auth_bypass") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XMacAuthBypass(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("open_auth"); ok || d.HasChange("open_auth") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XOpenAuth(d, v, "open_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-auth"] = t
		}
	}

	if v, ok := d.GetOk("policy_type"); ok || d.HasChange("policy_type") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XPolicyType(d, v, "policy_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-type"] = t
		}
	}

	if v, ok := d.GetOk("radius_timeout_overwrite"); ok || d.HasChange("radius_timeout_overwrite") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(d, v, "radius_timeout_overwrite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-timeout-overwrite"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok || d.HasChange("user_group") {
		t, err := expandObjectSwitchControllerSecurityPolicy8021XUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	return &obj, nil
}
