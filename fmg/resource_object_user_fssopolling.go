// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FSSO active directory servers for polling mode.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserFssoPolling() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserFssoPollingCreate,
		Read:   resourceObjectUserFssoPollingRead,
		Update: resourceObjectUserFssoPollingUpdate,
		Delete: resourceObjectUserFssoPollingDelete,

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
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"adgrp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"default_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logon_history": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"polling_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_ntlmv1_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smbv1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectUserFssoPollingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectUserFssoPolling(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFssoPolling resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserFssoPolling(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserFssoPolling resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserFssoPollingRead(d, m)
}

func resourceObjectUserFssoPollingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserFssoPolling(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFssoPolling resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserFssoPolling(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserFssoPolling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserFssoPollingRead(d, m)
}

func resourceObjectUserFssoPollingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectUserFssoPolling(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserFssoPollingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserFssoPolling(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFssoPolling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserFssoPolling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserFssoPolling resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserFssoPollingGuiMeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingAdgrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserFssoPollingAdgrpName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserFssoPolling-Adgrp-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserFssoPollingAdgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingDefaultDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingLogonHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserFssoPollingPollingFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingSmbNtlmv1Auth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingSmbv1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserFssoPollingUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserFssoPolling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_gui_meta", flattenObjectUserFssoPollingGuiMeta(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "ObjectUserFssoPolling-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("adgrp", flattenObjectUserFssoPollingAdgrp(o["adgrp"], d, "adgrp")); err != nil {
			if vv, ok := fortiAPIPatch(o["adgrp"], "ObjectUserFssoPolling-Adgrp"); ok {
				if err = d.Set("adgrp", vv); err != nil {
					return fmt.Errorf("Error reading adgrp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading adgrp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("adgrp"); ok {
			if err = d.Set("adgrp", flattenObjectUserFssoPollingAdgrp(o["adgrp"], d, "adgrp")); err != nil {
				if vv, ok := fortiAPIPatch(o["adgrp"], "ObjectUserFssoPolling-Adgrp"); ok {
					if err = d.Set("adgrp", vv); err != nil {
						return fmt.Errorf("Error reading adgrp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading adgrp: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_domain", flattenObjectUserFssoPollingDefaultDomain(o["default-domain"], d, "default_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-domain"], "ObjectUserFssoPolling-DefaultDomain"); ok {
			if err = d.Set("default_domain", vv); err != nil {
				return fmt.Errorf("Error reading default_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_domain: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserFssoPollingId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserFssoPolling-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserFssoPollingLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserFssoPolling-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("logon_history", flattenObjectUserFssoPollingLogonHistory(o["logon-history"], d, "logon_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-history"], "ObjectUserFssoPolling-LogonHistory"); ok {
			if err = d.Set("logon_history", vv); err != nil {
				return fmt.Errorf("Error reading logon_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_history: %v", err)
		}
	}

	if err = d.Set("polling_frequency", flattenObjectUserFssoPollingPollingFrequency(o["polling-frequency"], d, "polling_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["polling-frequency"], "ObjectUserFssoPolling-PollingFrequency"); ok {
			if err = d.Set("polling_frequency", vv); err != nil {
				return fmt.Errorf("Error reading polling_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading polling_frequency: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserFssoPollingPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserFssoPolling-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserFssoPollingServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserFssoPolling-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("smb_ntlmv1_auth", flattenObjectUserFssoPollingSmbNtlmv1Auth(o["smb-ntlmv1-auth"], d, "smb_ntlmv1_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-ntlmv1-auth"], "ObjectUserFssoPolling-SmbNtlmv1Auth"); ok {
			if err = d.Set("smb_ntlmv1_auth", vv); err != nil {
				return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if err = d.Set("smbv1", flattenObjectUserFssoPollingSmbv1(o["smbv1"], d, "smbv1")); err != nil {
		if vv, ok := fortiAPIPatch(o["smbv1"], "ObjectUserFssoPolling-Smbv1"); ok {
			if err = d.Set("smbv1", vv); err != nil {
				return fmt.Errorf("Error reading smbv1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smbv1: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserFssoPollingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserFssoPolling-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectUserFssoPollingUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectUserFssoPolling-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenObjectUserFssoPollingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserFssoPollingGuiMeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingAdgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectUserFssoPollingAdgrpName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserFssoPollingAdgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingDefaultDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingLogonHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserFssoPollingPollingFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingSmbNtlmv1Auth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingSmbv1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserFssoPollingUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserFssoPolling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandObjectUserFssoPollingGuiMeta(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("adgrp"); ok || d.HasChange("adgrp") {
		t, err := expandObjectUserFssoPollingAdgrp(d, v, "adgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adgrp"] = t
		}
	}

	if v, ok := d.GetOk("default_domain"); ok || d.HasChange("default_domain") {
		t, err := expandObjectUserFssoPollingDefaultDomain(d, v, "default_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-domain"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserFssoPollingId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserFssoPollingLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("logon_history"); ok || d.HasChange("logon_history") {
		t, err := expandObjectUserFssoPollingLogonHistory(d, v, "logon_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-history"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectUserFssoPollingPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("polling_frequency"); ok || d.HasChange("polling_frequency") {
		t, err := expandObjectUserFssoPollingPollingFrequency(d, v, "polling_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polling-frequency"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectUserFssoPollingPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectUserFssoPollingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("smb_ntlmv1_auth"); ok || d.HasChange("smb_ntlmv1_auth") {
		t, err := expandObjectUserFssoPollingSmbNtlmv1Auth(d, v, "smb_ntlmv1_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-ntlmv1-auth"] = t
		}
	}

	if v, ok := d.GetOk("smbv1"); ok || d.HasChange("smbv1") {
		t, err := expandObjectUserFssoPollingSmbv1(d, v, "smbv1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smbv1"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserFssoPollingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectUserFssoPollingUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
