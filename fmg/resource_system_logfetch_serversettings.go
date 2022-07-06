// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Log-fetch server settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogFetchServerSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogFetchServerSettingsUpdate,
		Read:   resourceSystemLogFetchServerSettingsRead,
		Update: resourceSystemLogFetchServerSettingsUpdate,
		Delete: resourceSystemLogFetchServerSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_conn_per_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_sessions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"session_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogFetchServerSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogFetchServerSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFetchServerSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogFetchServerSettings(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFetchServerSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogFetchServerSettings")

	return resourceSystemLogFetchServerSettingsRead(d, m)
}

func resourceSystemLogFetchServerSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogFetchServerSettings(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogFetchServerSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogFetchServerSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogFetchServerSettings(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFetchServerSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogFetchServerSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFetchServerSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogFetchServerSettingsMaxConnPerSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchServerSettingsMaxSessions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchServerSettingsSessionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogFetchServerSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_conn_per_session", flattenSystemLogFetchServerSettingsMaxConnPerSession(o["max-conn-per-session"], d, "max_conn_per_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-conn-per-session"], "SystemLogFetchServerSettings-MaxConnPerSession"); ok {
			if err = d.Set("max_conn_per_session", vv); err != nil {
				return fmt.Errorf("Error reading max_conn_per_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_conn_per_session: %v", err)
		}
	}

	if err = d.Set("max_sessions", flattenSystemLogFetchServerSettingsMaxSessions(o["max-sessions"], d, "max_sessions")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-sessions"], "SystemLogFetchServerSettings-MaxSessions"); ok {
			if err = d.Set("max_sessions", vv); err != nil {
				return fmt.Errorf("Error reading max_sessions: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_sessions: %v", err)
		}
	}

	if err = d.Set("session_timeout", flattenSystemLogFetchServerSettingsSessionTimeout(o["session-timeout"], d, "session_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-timeout"], "SystemLogFetchServerSettings-SessionTimeout"); ok {
			if err = d.Set("session_timeout", vv); err != nil {
				return fmt.Errorf("Error reading session_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemLogFetchServerSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogFetchServerSettingsMaxConnPerSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchServerSettingsMaxSessions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchServerSettingsSessionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogFetchServerSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_conn_per_session"); ok || d.HasChange("max_conn_per_session") {
		t, err := expandSystemLogFetchServerSettingsMaxConnPerSession(d, v, "max_conn_per_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-conn-per-session"] = t
		}
	}

	if v, ok := d.GetOk("max_sessions"); ok || d.HasChange("max_sessions") {
		t, err := expandSystemLogFetchServerSettingsMaxSessions(d, v, "max_sessions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-sessions"] = t
		}
	}

	if v, ok := d.GetOk("session_timeout"); ok || d.HasChange("session_timeout") {
		t, err := expandSystemLogFetchServerSettingsSessionTimeout(d, v, "session_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-timeout"] = t
		}
	}

	return &obj, nil
}
