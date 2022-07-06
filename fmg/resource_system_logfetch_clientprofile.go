// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Log-fetch client profile settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogFetchClientProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogFetchClientProfileCreate,
		Read:   resourceSystemLogFetchClientProfileRead,
		Update: resourceSystemLogFetchClientProfileUpdate,
		Delete: resourceSystemLogFetchClientProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"client_adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_range_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"index_fetch_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"oper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_filter_logic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_filter_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
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
			"peer_cert_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sync_adom_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemLogFetchClientProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogFetchClientProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogFetchClientProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLogFetchClientProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemLogFetchClientProfile resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogFetchClientProfileRead(d, m)
}

func resourceSystemLogFetchClientProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogFetchClientProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFetchClientProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogFetchClientProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFetchClientProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogFetchClientProfileRead(d, m)
}

func resourceSystemLogFetchClientProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogFetchClientProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogFetchClientProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogFetchClientProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogFetchClientProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFetchClientProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogFetchClientProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFetchClientProfile resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogFetchClientProfileClientAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDataRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDataRangeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom"
		if _, ok := i["adom"]; ok {
			v := flattenSystemLogFetchClientProfileDeviceFilterAdom(i["adom"], d, pre_append)
			tmp["adom"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-DeviceFilter-Adom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenSystemLogFetchClientProfileDeviceFilterDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-DeviceFilter-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogFetchClientProfileDeviceFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-DeviceFilter-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSystemLogFetchClientProfileDeviceFilterVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-DeviceFilter-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogFetchClientProfileDeviceFilterAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilterDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilterVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileEndTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogFetchClientProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileIndexFetchLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileLogFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
		if _, ok := i["field"]; ok {
			v := flattenSystemLogFetchClientProfileLogFilterField(i["field"], d, pre_append)
			tmp["field"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-LogFilter-Field")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogFetchClientProfileLogFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-LogFilter-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oper"
		if _, ok := i["oper"]; ok {
			v := flattenSystemLogFetchClientProfileLogFilterOper(i["oper"], d, pre_append)
			tmp["oper"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-LogFilter-Oper")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemLogFetchClientProfileLogFilterValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemLogFetchClientProfile-LogFilter-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogFetchClientProfileLogFilterField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileLogFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileLogFilterOper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileLogFilterValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileLogFilterLogic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileLogFilterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfilePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogFetchClientProfilePeerCertCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileServerAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogFetchClientProfileSyncAdomConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogFetchClientProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("client_adom", flattenSystemLogFetchClientProfileClientAdom(o["client-adom"], d, "client_adom")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-adom"], "SystemLogFetchClientProfile-ClientAdom"); ok {
			if err = d.Set("client_adom", vv); err != nil {
				return fmt.Errorf("Error reading client_adom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_adom: %v", err)
		}
	}

	if err = d.Set("data_range", flattenSystemLogFetchClientProfileDataRange(o["data-range"], d, "data_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-range"], "SystemLogFetchClientProfile-DataRange"); ok {
			if err = d.Set("data_range", vv); err != nil {
				return fmt.Errorf("Error reading data_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_range: %v", err)
		}
	}

	if err = d.Set("data_range_value", flattenSystemLogFetchClientProfileDataRangeValue(o["data-range-value"], d, "data_range_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-range-value"], "SystemLogFetchClientProfile-DataRangeValue"); ok {
			if err = d.Set("data_range_value", vv); err != nil {
				return fmt.Errorf("Error reading data_range_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_range_value: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("device_filter", flattenSystemLogFetchClientProfileDeviceFilter(o["device-filter"], d, "device_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["device-filter"], "SystemLogFetchClientProfile-DeviceFilter"); ok {
				if err = d.Set("device_filter", vv); err != nil {
					return fmt.Errorf("Error reading device_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_filter"); ok {
			if err = d.Set("device_filter", flattenSystemLogFetchClientProfileDeviceFilter(o["device-filter"], d, "device_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["device-filter"], "SystemLogFetchClientProfile-DeviceFilter"); ok {
					if err = d.Set("device_filter", vv); err != nil {
						return fmt.Errorf("Error reading device_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("end_time", flattenSystemLogFetchClientProfileEndTime(o["end-time"], d, "end_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-time"], "SystemLogFetchClientProfile-EndTime"); ok {
			if err = d.Set("end_time", vv); err != nil {
				return fmt.Errorf("Error reading end_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_time: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogFetchClientProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogFetchClientProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("index_fetch_logs", flattenSystemLogFetchClientProfileIndexFetchLogs(o["index-fetch-logs"], d, "index_fetch_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["index-fetch-logs"], "SystemLogFetchClientProfile-IndexFetchLogs"); ok {
			if err = d.Set("index_fetch_logs", vv); err != nil {
				return fmt.Errorf("Error reading index_fetch_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index_fetch_logs: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("log_filter", flattenSystemLogFetchClientProfileLogFilter(o["log-filter"], d, "log_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-filter"], "SystemLogFetchClientProfile-LogFilter"); ok {
				if err = d.Set("log_filter", vv); err != nil {
					return fmt.Errorf("Error reading log_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_filter"); ok {
			if err = d.Set("log_filter", flattenSystemLogFetchClientProfileLogFilter(o["log-filter"], d, "log_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-filter"], "SystemLogFetchClientProfile-LogFilter"); ok {
					if err = d.Set("log_filter", vv); err != nil {
						return fmt.Errorf("Error reading log_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_filter_logic", flattenSystemLogFetchClientProfileLogFilterLogic(o["log-filter-logic"], d, "log_filter_logic")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-filter-logic"], "SystemLogFetchClientProfile-LogFilterLogic"); ok {
			if err = d.Set("log_filter_logic", vv); err != nil {
				return fmt.Errorf("Error reading log_filter_logic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_filter_logic: %v", err)
		}
	}

	if err = d.Set("log_filter_status", flattenSystemLogFetchClientProfileLogFilterStatus(o["log-filter-status"], d, "log_filter_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-filter-status"], "SystemLogFetchClientProfile-LogFilterStatus"); ok {
			if err = d.Set("log_filter_status", vv); err != nil {
				return fmt.Errorf("Error reading log_filter_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_filter_status: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemLogFetchClientProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemLogFetchClientProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer_cert_cn", flattenSystemLogFetchClientProfilePeerCertCn(o["peer-cert-cn"], d, "peer_cert_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-cert-cn"], "SystemLogFetchClientProfile-PeerCertCn"); ok {
			if err = d.Set("peer_cert_cn", vv); err != nil {
				return fmt.Errorf("Error reading peer_cert_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_cert_cn: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemLogFetchClientProfileSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemLogFetchClientProfile-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("server_adom", flattenSystemLogFetchClientProfileServerAdom(o["server-adom"], d, "server_adom")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-adom"], "SystemLogFetchClientProfile-ServerAdom"); ok {
			if err = d.Set("server_adom", vv); err != nil {
				return fmt.Errorf("Error reading server_adom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_adom: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenSystemLogFetchClientProfileServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ip"], "SystemLogFetchClientProfile-ServerIp"); ok {
			if err = d.Set("server_ip", vv); err != nil {
				return fmt.Errorf("Error reading server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("start_time", flattenSystemLogFetchClientProfileStartTime(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "SystemLogFetchClientProfile-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("sync_adom_config", flattenSystemLogFetchClientProfileSyncAdomConfig(o["sync-adom-config"], d, "sync_adom_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-adom-config"], "SystemLogFetchClientProfile-SyncAdomConfig"); ok {
			if err = d.Set("sync_adom_config", vv); err != nil {
				return fmt.Errorf("Error reading sync_adom_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_adom_config: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemLogFetchClientProfileUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "SystemLogFetchClientProfile-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenSystemLogFetchClientProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogFetchClientProfileClientAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDataRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDataRangeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adom"], _ = expandSystemLogFetchClientProfileDeviceFilterAdom(d, i["adom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandSystemLogFetchClientProfileDeviceFilterDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogFetchClientProfileDeviceFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSystemLogFetchClientProfileDeviceFilterVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogFetchClientProfileDeviceFilterAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilterDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilterVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileEndTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogFetchClientProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileIndexFetchLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileLogFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["field"], _ = expandSystemLogFetchClientProfileLogFilterField(d, i["field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLogFetchClientProfileLogFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oper"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oper"], _ = expandSystemLogFetchClientProfileLogFilterOper(d, i["oper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemLogFetchClientProfileLogFilterValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogFetchClientProfileLogFilterField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileLogFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileLogFilterOper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileLogFilterValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileLogFilterLogic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileLogFilterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfilePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogFetchClientProfilePeerCertCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileServerAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogFetchClientProfileSyncAdomConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogFetchClientProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("client_adom"); ok || d.HasChange("client_adom") {
		t, err := expandSystemLogFetchClientProfileClientAdom(d, v, "client_adom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-adom"] = t
		}
	}

	if v, ok := d.GetOk("data_range"); ok || d.HasChange("data_range") {
		t, err := expandSystemLogFetchClientProfileDataRange(d, v, "data_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-range"] = t
		}
	}

	if v, ok := d.GetOk("data_range_value"); ok || d.HasChange("data_range_value") {
		t, err := expandSystemLogFetchClientProfileDataRangeValue(d, v, "data_range_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-range-value"] = t
		}
	}

	if v, ok := d.GetOk("device_filter"); ok || d.HasChange("device_filter") {
		t, err := expandSystemLogFetchClientProfileDeviceFilter(d, v, "device_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-filter"] = t
		}
	}

	if v, ok := d.GetOk("end_time"); ok || d.HasChange("end_time") {
		t, err := expandSystemLogFetchClientProfileEndTime(d, v, "end_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-time"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemLogFetchClientProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("index_fetch_logs"); ok || d.HasChange("index_fetch_logs") {
		t, err := expandSystemLogFetchClientProfileIndexFetchLogs(d, v, "index_fetch_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index-fetch-logs"] = t
		}
	}

	if v, ok := d.GetOk("log_filter"); ok || d.HasChange("log_filter") {
		t, err := expandSystemLogFetchClientProfileLogFilter(d, v, "log_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-filter"] = t
		}
	}

	if v, ok := d.GetOk("log_filter_logic"); ok || d.HasChange("log_filter_logic") {
		t, err := expandSystemLogFetchClientProfileLogFilterLogic(d, v, "log_filter_logic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-filter-logic"] = t
		}
	}

	if v, ok := d.GetOk("log_filter_status"); ok || d.HasChange("log_filter_status") {
		t, err := expandSystemLogFetchClientProfileLogFilterStatus(d, v, "log_filter_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-filter-status"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemLogFetchClientProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemLogFetchClientProfilePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer_cert_cn"); ok || d.HasChange("peer_cert_cn") {
		t, err := expandSystemLogFetchClientProfilePeerCertCn(d, v, "peer_cert_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-cert-cn"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemLogFetchClientProfileSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("server_adom"); ok || d.HasChange("server_adom") {
		t, err := expandSystemLogFetchClientProfileServerAdom(d, v, "server_adom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-adom"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok || d.HasChange("server_ip") {
		t, err := expandSystemLogFetchClientProfileServerIp(d, v, "server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandSystemLogFetchClientProfileStartTime(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("sync_adom_config"); ok || d.HasChange("sync_adom_config") {
		t, err := expandSystemLogFetchClientProfileSyncAdomConfig(d, v, "sync_adom_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-adom-config"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandSystemLogFetchClientProfileUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
