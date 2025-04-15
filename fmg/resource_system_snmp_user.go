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

func resourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpUserCreate,
		Read:   resourceSystemSnmpUserRead,
		Update: resourceSystemSnmpUserUpdate,
		Delete: resourceSystemSnmpUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_pwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"events": &schema.Schema{
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
			"notify_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notify_hosts6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notify_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priv_pwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
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
		},
	}
}

func resourceSystemSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUser resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemSnmpUser(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUser resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSnmpUserRead(d, m)
}

func resourceSystemSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUser resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSnmpUser(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSnmpUserRead(d, m)
}

func resourceSystemSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemSnmpUser(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSnmpUser(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUser resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserNotifyHosts6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserNotifyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_proto", flattenSystemSnmpUserAuthProto(o["auth-proto"], d, "auth_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-proto"], "SystemSnmpUser-AuthProto"); ok {
			if err = d.Set("auth_proto", vv); err != nil {
				return fmt.Errorf("Error reading auth_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("events", flattenSystemSnmpUserEvents(o["events"], d, "events")); err != nil {
		if vv, ok := fortiAPIPatch(o["events"], "SystemSnmpUser-Events"); ok {
			if err = d.Set("events", vv); err != nil {
				return fmt.Errorf("Error reading events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSnmpUserName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSnmpUser-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("notify_hosts", flattenSystemSnmpUserNotifyHosts(o["notify-hosts"], d, "notify_hosts")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-hosts"], "SystemSnmpUser-NotifyHosts"); ok {
			if err = d.Set("notify_hosts", vv); err != nil {
				return fmt.Errorf("Error reading notify_hosts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_hosts: %v", err)
		}
	}

	if err = d.Set("notify_hosts6", flattenSystemSnmpUserNotifyHosts6(o["notify-hosts6"], d, "notify_hosts6")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-hosts6"], "SystemSnmpUser-NotifyHosts6"); ok {
			if err = d.Set("notify_hosts6", vv); err != nil {
				return fmt.Errorf("Error reading notify_hosts6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_hosts6: %v", err)
		}
	}

	if err = d.Set("notify_port", flattenSystemSnmpUserNotifyPort(o["notify-port"], d, "notify_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["notify-port"], "SystemSnmpUser-NotifyPort"); ok {
			if err = d.Set("notify_port", vv); err != nil {
				return fmt.Errorf("Error reading notify_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notify_port: %v", err)
		}
	}

	if err = d.Set("priv_proto", flattenSystemSnmpUserPrivProto(o["priv-proto"], d, "priv_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["priv-proto"], "SystemSnmpUser-PrivProto"); ok {
			if err = d.Set("priv_proto", vv); err != nil {
				return fmt.Errorf("Error reading priv_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("queries", flattenSystemSnmpUserQueries(o["queries"], d, "queries")); err != nil {
		if vv, ok := fortiAPIPatch(o["queries"], "SystemSnmpUser-Queries"); ok {
			if err = d.Set("queries", vv); err != nil {
				return fmt.Errorf("Error reading queries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("query_port", flattenSystemSnmpUserQueryPort(o["query-port"], d, "query_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-port"], "SystemSnmpUser-QueryPort"); ok {
			if err = d.Set("query_port", vv); err != nil {
				return fmt.Errorf("Error reading query_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("security_level", flattenSystemSnmpUserSecurityLevel(o["security-level"], d, "security_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-level"], "SystemSnmpUser-SecurityLevel"); ok {
			if err = d.Set("security_level", vv); err != nil {
				return fmt.Errorf("Error reading security_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpUserEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserNotifyHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserNotifyHosts6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserNotifyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserQueryPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_proto"); ok || d.HasChange("auth_proto") {
		t, err := expandSystemSnmpUserAuthProto(d, v, "auth_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}

	if v, ok := d.GetOk("auth_pwd"); ok || d.HasChange("auth_pwd") {
		t, err := expandSystemSnmpUserAuthPwd(d, v, "auth_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}

	if v, ok := d.GetOk("events"); ok || d.HasChange("events") {
		t, err := expandSystemSnmpUserEvents(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSnmpUserName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts"); ok || d.HasChange("notify_hosts") {
		t, err := expandSystemSnmpUserNotifyHosts(d, v, "notify_hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts6"); ok || d.HasChange("notify_hosts6") {
		t, err := expandSystemSnmpUserNotifyHosts6(d, v, "notify_hosts6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts6"] = t
		}
	}

	if v, ok := d.GetOk("notify_port"); ok || d.HasChange("notify_port") {
		t, err := expandSystemSnmpUserNotifyPort(d, v, "notify_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-port"] = t
		}
	}

	if v, ok := d.GetOk("priv_proto"); ok || d.HasChange("priv_proto") {
		t, err := expandSystemSnmpUserPrivProto(d, v, "priv_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}

	if v, ok := d.GetOk("priv_pwd"); ok || d.HasChange("priv_pwd") {
		t, err := expandSystemSnmpUserPrivPwd(d, v, "priv_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}

	if v, ok := d.GetOk("queries"); ok || d.HasChange("queries") {
		t, err := expandSystemSnmpUserQueries(d, v, "queries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queries"] = t
		}
	}

	if v, ok := d.GetOk("query_port"); ok || d.HasChange("query_port") {
		t, err := expandSystemSnmpUserQueryPort(d, v, "query_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-port"] = t
		}
	}

	if v, ok := d.GetOk("security_level"); ok || d.HasChange("security_level") {
		t, err := expandSystemSnmpUserSecurityLevel(d, v, "security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}

	return &obj, nil
}
