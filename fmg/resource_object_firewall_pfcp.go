// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure PFCP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallPfcp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallPfcpCreate,
		Read:   resourceObjectFirewallPfcpRead,
		Update: resourceObjectFirewallPfcpUpdate,
		Delete: resourceObjectFirewallPfcpDelete,

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
			"denied_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forwarded_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_reserved_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_freq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_message_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"message_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"min_message_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pfcp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"traffic_count_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallPfcpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallPfcp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallPfcp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFirewallPfcp(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFirewallPfcp(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFirewallPfcp resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFirewallPfcp(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFirewallPfcp resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallPfcpRead(d, m)
}

func resourceObjectFirewallPfcpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallPfcp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallPfcp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallPfcp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallPfcp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallPfcpRead(d, m)
}

func resourceObjectFirewallPfcpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallPfcp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallPfcp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallPfcpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallPfcp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallPfcp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallPfcp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallPfcp resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallPfcpDeniedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpForwardedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpInvalidReservedField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpLogFreq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpMaxMessageLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpMessageFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallPfcpMinMessageLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpMonitorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpPfcpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpTrafficCountLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallPfcpUnknownVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallPfcp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("denied_log", flattenObjectFirewallPfcpDeniedLog(o["denied-log"], d, "denied_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["denied-log"], "ObjectFirewallPfcp-DeniedLog"); ok {
			if err = d.Set("denied_log", vv); err != nil {
				return fmt.Errorf("Error reading denied_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading denied_log: %v", err)
		}
	}

	if err = d.Set("forwarded_log", flattenObjectFirewallPfcpForwardedLog(o["forwarded-log"], d, "forwarded_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["forwarded-log"], "ObjectFirewallPfcp-ForwardedLog"); ok {
			if err = d.Set("forwarded_log", vv); err != nil {
				return fmt.Errorf("Error reading forwarded_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forwarded_log: %v", err)
		}
	}

	if err = d.Set("invalid_reserved_field", flattenObjectFirewallPfcpInvalidReservedField(o["invalid-reserved-field"], d, "invalid_reserved_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-reserved-field"], "ObjectFirewallPfcp-InvalidReservedField"); ok {
			if err = d.Set("invalid_reserved_field", vv); err != nil {
				return fmt.Errorf("Error reading invalid_reserved_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_reserved_field: %v", err)
		}
	}

	if err = d.Set("log_freq", flattenObjectFirewallPfcpLogFreq(o["log-freq"], d, "log_freq")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-freq"], "ObjectFirewallPfcp-LogFreq"); ok {
			if err = d.Set("log_freq", vv); err != nil {
				return fmt.Errorf("Error reading log_freq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_freq: %v", err)
		}
	}

	if err = d.Set("max_message_length", flattenObjectFirewallPfcpMaxMessageLength(o["max-message-length"], d, "max_message_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-message-length"], "ObjectFirewallPfcp-MaxMessageLength"); ok {
			if err = d.Set("max_message_length", vv); err != nil {
				return fmt.Errorf("Error reading max_message_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_message_length: %v", err)
		}
	}

	if err = d.Set("message_filter", flattenObjectFirewallPfcpMessageFilter(o["message-filter"], d, "message_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-filter"], "ObjectFirewallPfcp-MessageFilter"); ok {
			if err = d.Set("message_filter", vv); err != nil {
				return fmt.Errorf("Error reading message_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_filter: %v", err)
		}
	}

	if err = d.Set("min_message_length", flattenObjectFirewallPfcpMinMessageLength(o["min-message-length"], d, "min_message_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-message-length"], "ObjectFirewallPfcp-MinMessageLength"); ok {
			if err = d.Set("min_message_length", vv); err != nil {
				return fmt.Errorf("Error reading min_message_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_message_length: %v", err)
		}
	}

	if err = d.Set("monitor_mode", flattenObjectFirewallPfcpMonitorMode(o["monitor-mode"], d, "monitor_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-mode"], "ObjectFirewallPfcp-MonitorMode"); ok {
			if err = d.Set("monitor_mode", vv); err != nil {
				return fmt.Errorf("Error reading monitor_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_mode: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallPfcpName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallPfcp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pfcp_timeout", flattenObjectFirewallPfcpPfcpTimeout(o["pfcp-timeout"], d, "pfcp_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["pfcp-timeout"], "ObjectFirewallPfcp-PfcpTimeout"); ok {
			if err = d.Set("pfcp_timeout", vv); err != nil {
				return fmt.Errorf("Error reading pfcp_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pfcp_timeout: %v", err)
		}
	}

	if err = d.Set("traffic_count_log", flattenObjectFirewallPfcpTrafficCountLog(o["traffic-count-log"], d, "traffic_count_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-count-log"], "ObjectFirewallPfcp-TrafficCountLog"); ok {
			if err = d.Set("traffic_count_log", vv); err != nil {
				return fmt.Errorf("Error reading traffic_count_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_count_log: %v", err)
		}
	}

	if err = d.Set("unknown_version", flattenObjectFirewallPfcpUnknownVersion(o["unknown-version"], d, "unknown_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-version"], "ObjectFirewallPfcp-UnknownVersion"); ok {
			if err = d.Set("unknown_version", vv); err != nil {
				return fmt.Errorf("Error reading unknown_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_version: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallPfcpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallPfcpDeniedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpForwardedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpInvalidReservedField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpLogFreq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpMaxMessageLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpMessageFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallPfcpMinMessageLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpMonitorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpPfcpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpTrafficCountLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallPfcpUnknownVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallPfcp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("denied_log"); ok || d.HasChange("denied_log") {
		t, err := expandObjectFirewallPfcpDeniedLog(d, v, "denied_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["denied-log"] = t
		}
	}

	if v, ok := d.GetOk("forwarded_log"); ok || d.HasChange("forwarded_log") {
		t, err := expandObjectFirewallPfcpForwardedLog(d, v, "forwarded_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarded-log"] = t
		}
	}

	if v, ok := d.GetOk("invalid_reserved_field"); ok || d.HasChange("invalid_reserved_field") {
		t, err := expandObjectFirewallPfcpInvalidReservedField(d, v, "invalid_reserved_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-reserved-field"] = t
		}
	}

	if v, ok := d.GetOk("log_freq"); ok || d.HasChange("log_freq") {
		t, err := expandObjectFirewallPfcpLogFreq(d, v, "log_freq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-freq"] = t
		}
	}

	if v, ok := d.GetOk("max_message_length"); ok || d.HasChange("max_message_length") {
		t, err := expandObjectFirewallPfcpMaxMessageLength(d, v, "max_message_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-message-length"] = t
		}
	}

	if v, ok := d.GetOk("message_filter"); ok || d.HasChange("message_filter") {
		t, err := expandObjectFirewallPfcpMessageFilter(d, v, "message_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-filter"] = t
		}
	}

	if v, ok := d.GetOk("min_message_length"); ok || d.HasChange("min_message_length") {
		t, err := expandObjectFirewallPfcpMinMessageLength(d, v, "min_message_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-message-length"] = t
		}
	}

	if v, ok := d.GetOk("monitor_mode"); ok || d.HasChange("monitor_mode") {
		t, err := expandObjectFirewallPfcpMonitorMode(d, v, "monitor_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallPfcpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pfcp_timeout"); ok || d.HasChange("pfcp_timeout") {
		t, err := expandObjectFirewallPfcpPfcpTimeout(d, v, "pfcp_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfcp-timeout"] = t
		}
	}

	if v, ok := d.GetOk("traffic_count_log"); ok || d.HasChange("traffic_count_log") {
		t, err := expandObjectFirewallPfcpTrafficCountLog(d, v, "traffic_count_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-count-log"] = t
		}
	}

	if v, ok := d.GetOk("unknown_version"); ok || d.HasChange("unknown_version") {
		t, err := expandObjectFirewallPfcpUnknownVersion(d, v, "unknown_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-version"] = t
		}
	}

	return &obj, nil
}
