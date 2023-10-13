// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP user configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemSnmpUserCreate,
		Read:   resourceSystempSystemSnmpUserRead,
		Update: resourceSystempSystemSnmpUserUpdate,
		Delete: resourceSystempSystemSnmpUserDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_pwd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mib_view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"notify_hosts": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"notify_hosts6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priv_pwd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"queries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdoms": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystempSystemSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemSnmpUser resource while getting object: %v", err)
	}

	_, err = c.CreateSystempSystemSnmpUser(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystempSystemSnmpUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystempSystemSnmpUserRead(d, m)
}

func resourceSystempSystemSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemSnmpUser resource while getting object: %v", err)
	}

	_, err = c.UpdateSystempSystemSnmpUser(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemSnmpUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystempSystemSnmpUserRead(d, m)
}

func resourceSystempSystemSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	err = c.DeleteSystempSystemSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemSnmpUserRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempSystemSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemSnmpUser resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserAuthPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemSnmpUserEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemSnmpUserHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserMibView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemSnmpUserNotifyHosts6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserPrivPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserSourceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserTrapLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserTrapRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserTrapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemSnmpUserVdoms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystempSystemSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_proto", flattenSystempSystemSnmpUserAuthProto(o["auth-proto"], d, "auth_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-proto"], "SystempSystemSnmpUser-AuthProto"); ok {
			if err = d.Set("auth_proto", vv); err != nil {
				return fmt.Errorf("Error reading auth_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("auth_pwd", flattenSystempSystemSnmpUserAuthPwd(o["auth-pwd"], d, "auth_pwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-pwd"], "SystempSystemSnmpUser-AuthPwd"); ok {
			if err = d.Set("auth_pwd", vv); err != nil {
				return fmt.Errorf("Error reading auth_pwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_pwd: %v", err)
		}
	}

	if err = d.Set("events", flattenSystempSystemSnmpUserEvents(o["events"], d, "events")); err != nil {
		if vv, ok := fortiAPIPatch(o["events"], "SystempSystemSnmpUser-Events"); ok {
			if err = d.Set("events", vv); err != nil {
				return fmt.Errorf("Error reading events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("ha_direct", flattenSystempSystemSnmpUserHaDirect(o["ha-direct"], d, "ha_direct")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-direct"], "SystempSystemSnmpUser-HaDirect"); ok {
			if err = d.Set("ha_direct", vv); err != nil {
				return fmt.Errorf("Error reading ha_direct: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("mib_view", flattenSystempSystemSnmpUserMibView(o["mib-view"], d, "mib_view")); err != nil {
		if vv, ok := fortiAPIPatch(o["mib-view"], "SystempSystemSnmpUser-MibView"); ok {
			if err = d.Set("mib_view", vv); err != nil {
				return fmt.Errorf("Error reading mib_view: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mib_view: %v", err)
		}
	}

	if err = d.Set("name", flattenSystempSystemSnmpUserName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystempSystemSnmpUser-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("notify_hosts", flattenSystempSystemSnmpUserNotifyHosts(o["notify-hosts"], d, "notify_hosts")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-hosts"], "SystempSystemSnmpUser-NotifyHosts"); ok {
			if err = d.Set("notify_hosts", vv); err != nil {
				return fmt.Errorf("Error reading notify_hosts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_hosts: %v", err)
		}
	}

	if err = d.Set("notify_hosts6", flattenSystempSystemSnmpUserNotifyHosts6(o["notify-hosts6"], d, "notify_hosts6")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-hosts6"], "SystempSystemSnmpUser-NotifyHosts6"); ok {
			if err = d.Set("notify_hosts6", vv); err != nil {
				return fmt.Errorf("Error reading notify_hosts6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_hosts6: %v", err)
		}
	}

	if err = d.Set("priv_proto", flattenSystempSystemSnmpUserPrivProto(o["priv-proto"], d, "priv_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["priv-proto"], "SystempSystemSnmpUser-PrivProto"); ok {
			if err = d.Set("priv_proto", vv); err != nil {
				return fmt.Errorf("Error reading priv_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("priv_pwd", flattenSystempSystemSnmpUserPrivPwd(o["priv-pwd"], d, "priv_pwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["priv-pwd"], "SystempSystemSnmpUser-PrivPwd"); ok {
			if err = d.Set("priv_pwd", vv); err != nil {
				return fmt.Errorf("Error reading priv_pwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priv_pwd: %v", err)
		}
	}

	if err = d.Set("queries", flattenSystempSystemSnmpUserQueries(o["queries"], d, "queries")); err != nil {
		if vv, ok := fortiAPIPatch(o["queries"], "SystempSystemSnmpUser-Queries"); ok {
			if err = d.Set("queries", vv); err != nil {
				return fmt.Errorf("Error reading queries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("query_port", flattenSystempSystemSnmpUserQueryPort(o["query-port"], d, "query_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-port"], "SystempSystemSnmpUser-QueryPort"); ok {
			if err = d.Set("query_port", vv); err != nil {
				return fmt.Errorf("Error reading query_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("security_level", flattenSystempSystemSnmpUserSecurityLevel(o["security-level"], d, "security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-level"], "SystempSystemSnmpUser-SecurityLevel"); ok {
			if err = d.Set("security_level", vv); err != nil {
				return fmt.Errorf("Error reading security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystempSystemSnmpUserSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystempSystemSnmpUser-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ipv6", flattenSystempSystemSnmpUserSourceIpv6(o["source-ipv6"], d, "source_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ipv6"], "SystempSystemSnmpUser-SourceIpv6"); ok {
			if err = d.Set("source_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading source_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ipv6: %v", err)
		}
	}

	if err = d.Set("status", flattenSystempSystemSnmpUserStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystempSystemSnmpUser-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_lport", flattenSystempSystemSnmpUserTrapLport(o["trap-lport"], d, "trap_lport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-lport"], "SystempSystemSnmpUser-TrapLport"); ok {
			if err = d.Set("trap_lport", vv); err != nil {
				return fmt.Errorf("Error reading trap_lport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_lport: %v", err)
		}
	}

	if err = d.Set("trap_rport", flattenSystempSystemSnmpUserTrapRport(o["trap-rport"], d, "trap_rport")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-rport"], "SystempSystemSnmpUser-TrapRport"); ok {
			if err = d.Set("trap_rport", vv); err != nil {
				return fmt.Errorf("Error reading trap_rport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_rport: %v", err)
		}
	}

	if err = d.Set("trap_status", flattenSystempSystemSnmpUserTrapStatus(o["trap-status"], d, "trap_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["trap-status"], "SystempSystemSnmpUser-TrapStatus"); ok {
			if err = d.Set("trap_status", vv); err != nil {
				return fmt.Errorf("Error reading trap_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trap_status: %v", err)
		}
	}

	if err = d.Set("vdoms", flattenSystempSystemSnmpUserVdoms(o["vdoms"], d, "vdoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdoms"], "SystempSystemSnmpUser-Vdoms"); ok {
			if err = d.Set("vdoms", vv); err != nil {
				return fmt.Errorf("Error reading vdoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdoms: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemSnmpUserEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemSnmpUserHaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserMibView(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserNotifyHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemSnmpUserNotifyHosts6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserQueryPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserSourceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserTrapLport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserTrapRport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserTrapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemSnmpUserVdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystempSystemSnmpUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_proto"); ok || d.HasChange("auth_proto") {
		t, err := expandSystempSystemSnmpUserAuthProto(d, v, "auth_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}

	if v, ok := d.GetOk("auth_pwd"); ok || d.HasChange("auth_pwd") {
		t, err := expandSystempSystemSnmpUserAuthPwd(d, v, "auth_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}

	if v, ok := d.GetOk("events"); ok || d.HasChange("events") {
		t, err := expandSystempSystemSnmpUserEvents(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("ha_direct"); ok || d.HasChange("ha_direct") {
		t, err := expandSystempSystemSnmpUserHaDirect(d, v, "ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-direct"] = t
		}
	}

	if v, ok := d.GetOk("mib_view"); ok || d.HasChange("mib_view") {
		t, err := expandSystempSystemSnmpUserMibView(d, v, "mib_view")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mib-view"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystempSystemSnmpUserName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts"); ok || d.HasChange("notify_hosts") {
		t, err := expandSystempSystemSnmpUserNotifyHosts(d, v, "notify_hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts6"); ok || d.HasChange("notify_hosts6") {
		t, err := expandSystempSystemSnmpUserNotifyHosts6(d, v, "notify_hosts6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts6"] = t
		}
	}

	if v, ok := d.GetOk("priv_proto"); ok || d.HasChange("priv_proto") {
		t, err := expandSystempSystemSnmpUserPrivProto(d, v, "priv_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}

	if v, ok := d.GetOk("priv_pwd"); ok || d.HasChange("priv_pwd") {
		t, err := expandSystempSystemSnmpUserPrivPwd(d, v, "priv_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}

	if v, ok := d.GetOk("queries"); ok || d.HasChange("queries") {
		t, err := expandSystempSystemSnmpUserQueries(d, v, "queries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queries"] = t
		}
	}

	if v, ok := d.GetOk("query_port"); ok || d.HasChange("query_port") {
		t, err := expandSystempSystemSnmpUserQueryPort(d, v, "query_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-port"] = t
		}
	}

	if v, ok := d.GetOk("security_level"); ok || d.HasChange("security_level") {
		t, err := expandSystempSystemSnmpUserSecurityLevel(d, v, "security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystempSystemSnmpUserSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ipv6"); ok || d.HasChange("source_ipv6") {
		t, err := expandSystempSystemSnmpUserSourceIpv6(d, v, "source_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystempSystemSnmpUserStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_lport"); ok || d.HasChange("trap_lport") {
		t, err := expandSystempSystemSnmpUserTrapLport(d, v, "trap_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_rport"); ok || d.HasChange("trap_rport") {
		t, err := expandSystempSystemSnmpUserTrapRport(d, v, "trap_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_status"); ok || d.HasChange("trap_status") {
		t, err := expandSystempSystemSnmpUserTrapStatus(d, v, "trap_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-status"] = t
		}
	}

	if v, ok := d.GetOk("vdoms"); ok || d.HasChange("vdoms") {
		t, err := expandSystempSystemSnmpUserVdoms(d, v, "vdoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdoms"] = t
		}
	}

	return &obj, nil
}
