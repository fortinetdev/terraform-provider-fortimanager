// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Alert events.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAlertEvent() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAlertEventCreate,
		Read:   resourceSystemAlertEventRead,
		Update: resourceSystemAlertEventUpdate,
		Delete: resourceSystemAlertEventDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"alert_destination": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"smtp_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"syslog_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"to": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"enable_generic_text": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"enable_severity_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"event_time_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"generic_text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_events": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity_level_comp": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"severity_level_logs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAlertEventCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAlertEvent(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAlertEvent resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAlertEvent(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemAlertEvent resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemAlertEventRead(d, m)
}

func resourceSystemAlertEventUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAlertEvent(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertEvent resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAlertEvent(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertEvent resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemAlertEventRead(d, m)
}

func resourceSystemAlertEventDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAlertEvent(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAlertEvent resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAlertEventRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAlertEvent(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertEvent resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlertEvent(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertEvent resource from API: %v", err)
	}
	return nil
}

func flattenSystemAlertEventAlertDestination(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := i["from"]; ok {
			v := flattenSystemAlertEventAlertDestinationFrom(i["from"], d, pre_append)
			tmp["from"] = fortiAPISubPartPatch(v, "SystemAlertEvent-AlertDestination-From")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtp_name"
		if _, ok := i["smtp-name"]; ok {
			v := flattenSystemAlertEventAlertDestinationSmtpName(i["smtp-name"], d, pre_append)
			tmp["smtp_name"] = fortiAPISubPartPatch(v, "SystemAlertEvent-AlertDestination-SmtpName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "snmp_name"
		if _, ok := i["snmp-name"]; ok {
			v := flattenSystemAlertEventAlertDestinationSnmpName(i["snmp-name"], d, pre_append)
			tmp["snmp_name"] = fortiAPISubPartPatch(v, "SystemAlertEvent-AlertDestination-SnmpName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syslog_name"
		if _, ok := i["syslog-name"]; ok {
			v := flattenSystemAlertEventAlertDestinationSyslogName(i["syslog-name"], d, pre_append)
			tmp["syslog_name"] = fortiAPISubPartPatch(v, "SystemAlertEvent-AlertDestination-SyslogName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to"
		if _, ok := i["to"]; ok {
			v := flattenSystemAlertEventAlertDestinationTo(i["to"], d, pre_append)
			tmp["to"] = fortiAPISubPartPatch(v, "SystemAlertEvent-AlertDestination-To")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemAlertEventAlertDestinationType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemAlertEvent-AlertDestination-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAlertEventAlertDestinationFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventAlertDestinationSmtpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventAlertDestinationSnmpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventAlertDestinationSyslogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventAlertDestinationTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventAlertDestinationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "mail",
			1: "snmp",
			2: "syslog",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventEnableGenericText(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "enable",
			2: "disable",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventEnableSeverityFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "enable",
			2: "disable",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventEventTimePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "0.5",
			1:   "1",
			3:   "3",
			6:   "6",
			12:  "12",
			24:  "24",
			72:  "72",
			168: "168",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventGenericText(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAlertEventNumEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "1",
			5:   "5",
			10:  "10",
			50:  "50",
			100: "100",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventSeverityFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "high",
			1: "medium-high",
			2: "medium",
			3: "medium-low",
			4: "low",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventSeverityLevelComp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: ">=",
			2: "=",
			4: "<=",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenSystemAlertEventSeverityLevelLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:   "no-check",
			2:   "information",
			4:   "notify",
			8:   "warning",
			16:  "error",
			32:  "critical",
			64:  "alert",
			128: "emergency",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemAlertEvent(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("alert_destination", flattenSystemAlertEventAlertDestination(o["alert-destination"], d, "alert_destination")); err != nil {
			if vv, ok := fortiAPIPatch(o["alert-destination"], "SystemAlertEvent-AlertDestination"); ok {
				if err = d.Set("alert_destination", vv); err != nil {
					return fmt.Errorf("Error reading alert_destination: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading alert_destination: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alert_destination"); ok {
			if err = d.Set("alert_destination", flattenSystemAlertEventAlertDestination(o["alert-destination"], d, "alert_destination")); err != nil {
				if vv, ok := fortiAPIPatch(o["alert-destination"], "SystemAlertEvent-AlertDestination"); ok {
					if err = d.Set("alert_destination", vv); err != nil {
						return fmt.Errorf("Error reading alert_destination: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading alert_destination: %v", err)
				}
			}
		}
	}

	if err = d.Set("enable_generic_text", flattenSystemAlertEventEnableGenericText(o["enable-generic-text"], d, "enable_generic_text")); err != nil {
		if vv, ok := fortiAPIPatch(o["enable-generic-text"], "SystemAlertEvent-EnableGenericText"); ok {
			if err = d.Set("enable_generic_text", vv); err != nil {
				return fmt.Errorf("Error reading enable_generic_text: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enable_generic_text: %v", err)
		}
	}

	if err = d.Set("enable_severity_filter", flattenSystemAlertEventEnableSeverityFilter(o["enable-severity-filter"], d, "enable_severity_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["enable-severity-filter"], "SystemAlertEvent-EnableSeverityFilter"); ok {
			if err = d.Set("enable_severity_filter", vv); err != nil {
				return fmt.Errorf("Error reading enable_severity_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enable_severity_filter: %v", err)
		}
	}

	if err = d.Set("event_time_period", flattenSystemAlertEventEventTimePeriod(o["event-time-period"], d, "event_time_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["event-time-period"], "SystemAlertEvent-EventTimePeriod"); ok {
			if err = d.Set("event_time_period", vv); err != nil {
				return fmt.Errorf("Error reading event_time_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event_time_period: %v", err)
		}
	}

	if err = d.Set("generic_text", flattenSystemAlertEventGenericText(o["generic-text"], d, "generic_text")); err != nil {
		if vv, ok := fortiAPIPatch(o["generic-text"], "SystemAlertEvent-GenericText"); ok {
			if err = d.Set("generic_text", vv); err != nil {
				return fmt.Errorf("Error reading generic_text: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading generic_text: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAlertEventName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAlertEvent-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("num_events", flattenSystemAlertEventNumEvents(o["num-events"], d, "num_events")); err != nil {
		if vv, ok := fortiAPIPatch(o["num-events"], "SystemAlertEvent-NumEvents"); ok {
			if err = d.Set("num_events", vv); err != nil {
				return fmt.Errorf("Error reading num_events: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading num_events: %v", err)
		}
	}

	if err = d.Set("severity_filter", flattenSystemAlertEventSeverityFilter(o["severity-filter"], d, "severity_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity-filter"], "SystemAlertEvent-SeverityFilter"); ok {
			if err = d.Set("severity_filter", vv); err != nil {
				return fmt.Errorf("Error reading severity_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity_filter: %v", err)
		}
	}

	if err = d.Set("severity_level_comp", flattenSystemAlertEventSeverityLevelComp(o["severity-level-comp"], d, "severity_level_comp")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity-level-comp"], "SystemAlertEvent-SeverityLevelComp"); ok {
			if err = d.Set("severity_level_comp", vv); err != nil {
				return fmt.Errorf("Error reading severity_level_comp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity_level_comp: %v", err)
		}
	}

	if err = d.Set("severity_level_logs", flattenSystemAlertEventSeverityLevelLogs(o["severity-level-logs"], d, "severity_level_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity-level-logs"], "SystemAlertEvent-SeverityLevelLogs"); ok {
			if err = d.Set("severity_level_logs", vv); err != nil {
				return fmt.Errorf("Error reading severity_level_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity_level_logs: %v", err)
		}
	}

	return nil
}

func flattenSystemAlertEventFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAlertEventAlertDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["from"], _ = expandSystemAlertEventAlertDestinationFrom(d, i["from"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtp_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["smtp-name"], _ = expandSystemAlertEventAlertDestinationSmtpName(d, i["smtp_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "snmp_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["snmp-name"], _ = expandSystemAlertEventAlertDestinationSnmpName(d, i["snmp_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syslog_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["syslog-name"], _ = expandSystemAlertEventAlertDestinationSyslogName(d, i["syslog_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["to"], _ = expandSystemAlertEventAlertDestinationTo(d, i["to"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemAlertEventAlertDestinationType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAlertEventAlertDestinationFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventAlertDestinationSmtpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventAlertDestinationSnmpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventAlertDestinationSyslogName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventAlertDestinationTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventAlertDestinationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventEnableGenericText(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAlertEventEnableSeverityFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAlertEventEventTimePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventGenericText(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventNumEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventSeverityFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertEventSeverityLevelComp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAlertEventSeverityLevelLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemAlertEvent(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alert_destination"); ok {
		t, err := expandSystemAlertEventAlertDestination(d, v, "alert_destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert-destination"] = t
		}
	}

	if v, ok := d.GetOk("enable_generic_text"); ok {
		t, err := expandSystemAlertEventEnableGenericText(d, v, "enable_generic_text")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enable-generic-text"] = t
		}
	}

	if v, ok := d.GetOk("enable_severity_filter"); ok {
		t, err := expandSystemAlertEventEnableSeverityFilter(d, v, "enable_severity_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enable-severity-filter"] = t
		}
	}

	if v, ok := d.GetOk("event_time_period"); ok {
		t, err := expandSystemAlertEventEventTimePeriod(d, v, "event_time_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-time-period"] = t
		}
	}

	if v, ok := d.GetOk("generic_text"); ok {
		t, err := expandSystemAlertEventGenericText(d, v, "generic_text")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["generic-text"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAlertEventName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("num_events"); ok {
		t, err := expandSystemAlertEventNumEvents(d, v, "num_events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["num-events"] = t
		}
	}

	if v, ok := d.GetOk("severity_filter"); ok {
		t, err := expandSystemAlertEventSeverityFilter(d, v, "severity_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity-filter"] = t
		}
	}

	if v, ok := d.GetOk("severity_level_comp"); ok {
		t, err := expandSystemAlertEventSeverityLevelComp(d, v, "severity_level_comp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity-level-comp"] = t
		}
	}

	if v, ok := d.GetOk("severity_level_logs"); ok {
		t, err := expandSystemAlertEventSeverityLevelLogs(d, v, "severity_level_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity-level-logs"] = t
		}
	}

	return &obj, nil
}