// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SQL settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSql() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSqlUpdate,
		Read:   resourceSystemSqlRead,
		Update: resourceSystemSqlUpdate,
		Delete: resourceSystemSqlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"background_rebuild": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compress_table_min_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_index": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"index_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"custom_skipidx": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"index_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"database_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_count_high": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_table_partition_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fct_table_partition_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"logtype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"prompt_sql_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rebuild_event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rebuild_event_start_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"text_search_index": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_table_partition_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ts_index_field": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"utm_table_partition_time": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemSqlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSql(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSql resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemSql(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSql resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSql")

	return resourceSystemSqlRead(d, m)
}

func resourceSystemSqlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemSql(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSql resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSqlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSql(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSql resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSql(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSql resource from API: %v", err)
	}
	return nil
}

func flattenSystemSqlBackgroundRebuild(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCompressTableMinAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndex(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenSystemSqlCustomIndexCaseSensitive(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			v := flattenSystemSqlCustomIndexDeviceType(i["device-type"], d, pre_append)
			tmp["device_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-DeviceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSqlCustomIndexId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := i["index-field"]; ok {
			v := flattenSystemSqlCustomIndexIndexField(i["index-field"], d, pre_append)
			tmp["index_field"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-IndexField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := i["log-type"]; ok {
			v := flattenSystemSqlCustomIndexLogType(i["log-type"], d, pre_append)
			tmp["log_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-LogType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSqlCustomIndexCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexIndexField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexLogType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidx(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			v := flattenSystemSqlCustomSkipidxDeviceType(i["device-type"], d, pre_append)
			tmp["device_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-DeviceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSqlCustomSkipidxId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := i["index-field"]; ok {
			v := flattenSystemSqlCustomSkipidxIndexField(i["index-field"], d, pre_append)
			tmp["index_field"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-IndexField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := i["log-type"]; ok {
			v := flattenSystemSqlCustomSkipidxLogType(i["log-type"], d, pre_append)
			tmp["log_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-LogType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSqlCustomSkipidxDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxIndexField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxLogType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlDatabaseName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlDatabaseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlDeviceCountHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlEventTablePartitionTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlFctTablePartitionTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlLogtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlPromptSqlUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlRebuildEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlRebuildEventStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTextSearchIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTrafficTablePartitionTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTsIndexField(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenSystemSqlTsIndexFieldCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "SystemSql-TsIndexField-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemSqlTsIndexFieldValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemSql-TsIndexField-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSqlTsIndexFieldCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTsIndexFieldValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlUtmTablePartitionTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSql(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("background_rebuild", flattenSystemSqlBackgroundRebuild(o["background-rebuild"], d, "background_rebuild")); err != nil {
		if vv, ok := fortiAPIPatch(o["background-rebuild"], "SystemSql-BackgroundRebuild"); ok {
			if err = d.Set("background_rebuild", vv); err != nil {
				return fmt.Errorf("Error reading background_rebuild: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading background_rebuild: %v", err)
		}
	}

	if err = d.Set("compress_table_min_age", flattenSystemSqlCompressTableMinAge(o["compress-table-min-age"], d, "compress_table_min_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["compress-table-min-age"], "SystemSql-CompressTableMinAge"); ok {
			if err = d.Set("compress_table_min_age", vv); err != nil {
				return fmt.Errorf("Error reading compress_table_min_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compress_table_min_age: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_index", flattenSystemSqlCustomIndex(o["custom-index"], d, "custom_index")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-index"], "SystemSql-CustomIndex"); ok {
				if err = d.Set("custom_index", vv); err != nil {
					return fmt.Errorf("Error reading custom_index: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_index: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_index"); ok {
			if err = d.Set("custom_index", flattenSystemSqlCustomIndex(o["custom-index"], d, "custom_index")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-index"], "SystemSql-CustomIndex"); ok {
					if err = d.Set("custom_index", vv); err != nil {
						return fmt.Errorf("Error reading custom_index: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_index: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("custom_skipidx", flattenSystemSqlCustomSkipidx(o["custom-skipidx"], d, "custom_skipidx")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-skipidx"], "SystemSql-CustomSkipidx"); ok {
				if err = d.Set("custom_skipidx", vv); err != nil {
					return fmt.Errorf("Error reading custom_skipidx: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_skipidx: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_skipidx"); ok {
			if err = d.Set("custom_skipidx", flattenSystemSqlCustomSkipidx(o["custom-skipidx"], d, "custom_skipidx")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-skipidx"], "SystemSql-CustomSkipidx"); ok {
					if err = d.Set("custom_skipidx", vv); err != nil {
						return fmt.Errorf("Error reading custom_skipidx: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_skipidx: %v", err)
				}
			}
		}
	}

	if err = d.Set("database_name", flattenSystemSqlDatabaseName(o["database-name"], d, "database_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-name"], "SystemSql-DatabaseName"); ok {
			if err = d.Set("database_name", vv); err != nil {
				return fmt.Errorf("Error reading database_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_name: %v", err)
		}
	}

	if err = d.Set("database_type", flattenSystemSqlDatabaseType(o["database-type"], d, "database_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-type"], "SystemSql-DatabaseType"); ok {
			if err = d.Set("database_type", vv); err != nil {
				return fmt.Errorf("Error reading database_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_type: %v", err)
		}
	}

	if err = d.Set("device_count_high", flattenSystemSqlDeviceCountHigh(o["device-count-high"], d, "device_count_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-count-high"], "SystemSql-DeviceCountHigh"); ok {
			if err = d.Set("device_count_high", vv); err != nil {
				return fmt.Errorf("Error reading device_count_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_count_high: %v", err)
		}
	}

	if err = d.Set("event_table_partition_time", flattenSystemSqlEventTablePartitionTime(o["event-table-partition-time"], d, "event_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["event-table-partition-time"], "SystemSql-EventTablePartitionTime"); ok {
			if err = d.Set("event_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading event_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event_table_partition_time: %v", err)
		}
	}

	if err = d.Set("fct_table_partition_time", flattenSystemSqlFctTablePartitionTime(o["fct-table-partition-time"], d, "fct_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["fct-table-partition-time"], "SystemSql-FctTablePartitionTime"); ok {
			if err = d.Set("fct_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading fct_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fct_table_partition_time: %v", err)
		}
	}

	if err = d.Set("logtype", flattenSystemSqlLogtype(o["logtype"], d, "logtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtype"], "SystemSql-Logtype"); ok {
			if err = d.Set("logtype", vv); err != nil {
				return fmt.Errorf("Error reading logtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtype: %v", err)
		}
	}

	if err = d.Set("prompt_sql_upgrade", flattenSystemSqlPromptSqlUpgrade(o["prompt-sql-upgrade"], d, "prompt_sql_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["prompt-sql-upgrade"], "SystemSql-PromptSqlUpgrade"); ok {
			if err = d.Set("prompt_sql_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading prompt_sql_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prompt_sql_upgrade: %v", err)
		}
	}

	if err = d.Set("rebuild_event", flattenSystemSqlRebuildEvent(o["rebuild-event"], d, "rebuild_event")); err != nil {
		if vv, ok := fortiAPIPatch(o["rebuild-event"], "SystemSql-RebuildEvent"); ok {
			if err = d.Set("rebuild_event", vv); err != nil {
				return fmt.Errorf("Error reading rebuild_event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rebuild_event: %v", err)
		}
	}

	if err = d.Set("rebuild_event_start_time", flattenSystemSqlRebuildEventStartTime(o["rebuild-event-start-time"], d, "rebuild_event_start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["rebuild-event-start-time"], "SystemSql-RebuildEventStartTime"); ok {
			if err = d.Set("rebuild_event_start_time", vv); err != nil {
				return fmt.Errorf("Error reading rebuild_event_start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rebuild_event_start_time: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSqlServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemSql-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("start_time", flattenSystemSqlStartTime(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "SystemSql-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSqlStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSql-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("text_search_index", flattenSystemSqlTextSearchIndex(o["text-search-index"], d, "text_search_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["text-search-index"], "SystemSql-TextSearchIndex"); ok {
			if err = d.Set("text_search_index", vv); err != nil {
				return fmt.Errorf("Error reading text_search_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading text_search_index: %v", err)
		}
	}

	if err = d.Set("traffic_table_partition_time", flattenSystemSqlTrafficTablePartitionTime(o["traffic-table-partition-time"], d, "traffic_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-table-partition-time"], "SystemSql-TrafficTablePartitionTime"); ok {
			if err = d.Set("traffic_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading traffic_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_table_partition_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ts_index_field", flattenSystemSqlTsIndexField(o["ts-index-field"], d, "ts_index_field")); err != nil {
			if vv, ok := fortiAPIPatch(o["ts-index-field"], "SystemSql-TsIndexField"); ok {
				if err = d.Set("ts_index_field", vv); err != nil {
					return fmt.Errorf("Error reading ts_index_field: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ts_index_field: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ts_index_field"); ok {
			if err = d.Set("ts_index_field", flattenSystemSqlTsIndexField(o["ts-index-field"], d, "ts_index_field")); err != nil {
				if vv, ok := fortiAPIPatch(o["ts-index-field"], "SystemSql-TsIndexField"); ok {
					if err = d.Set("ts_index_field", vv); err != nil {
						return fmt.Errorf("Error reading ts_index_field: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ts_index_field: %v", err)
				}
			}
		}
	}

	if err = d.Set("username", flattenSystemSqlUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemSql-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("utm_table_partition_time", flattenSystemSqlUtmTablePartitionTime(o["utm-table-partition-time"], d, "utm_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-table-partition-time"], "SystemSql-UtmTablePartitionTime"); ok {
			if err = d.Set("utm_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading utm_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_table_partition_time: %v", err)
		}
	}

	return nil
}

func flattenSystemSqlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSqlBackgroundRebuild(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCompressTableMinAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandSystemSqlCustomIndexCaseSensitive(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-type"], _ = expandSystemSqlCustomIndexDeviceType(d, i["device_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSqlCustomIndexId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index-field"], _ = expandSystemSqlCustomIndexIndexField(d, i["index_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-type"], _ = expandSystemSqlCustomIndexLogType(d, i["log_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSqlCustomIndexCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexDeviceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexIndexField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexLogType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-type"], _ = expandSystemSqlCustomSkipidxDeviceType(d, i["device_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSqlCustomSkipidxId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index-field"], _ = expandSystemSqlCustomSkipidxIndexField(d, i["index_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-type"], _ = expandSystemSqlCustomSkipidxLogType(d, i["log_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSqlCustomSkipidxDeviceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxIndexField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxLogType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlDatabaseName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlDatabaseType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlDeviceCountHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlEventTablePartitionTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlFctTablePartitionTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlLogtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlPromptSqlUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlRebuildEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlRebuildEventStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTextSearchIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTrafficTablePartitionTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTsIndexField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandSystemSqlTsIndexFieldCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemSqlTsIndexFieldValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSqlTsIndexFieldCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTsIndexFieldValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlUtmTablePartitionTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSql(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("background_rebuild"); ok || d.HasChange("background_rebuild") {
		t, err := expandSystemSqlBackgroundRebuild(d, v, "background_rebuild")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["background-rebuild"] = t
		}
	}

	if v, ok := d.GetOk("compress_table_min_age"); ok || d.HasChange("compress_table_min_age") {
		t, err := expandSystemSqlCompressTableMinAge(d, v, "compress_table_min_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compress-table-min-age"] = t
		}
	}

	if v, ok := d.GetOk("custom_index"); ok || d.HasChange("custom_index") {
		t, err := expandSystemSqlCustomIndex(d, v, "custom_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-index"] = t
		}
	}

	if v, ok := d.GetOk("custom_skipidx"); ok || d.HasChange("custom_skipidx") {
		t, err := expandSystemSqlCustomSkipidx(d, v, "custom_skipidx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-skipidx"] = t
		}
	}

	if v, ok := d.GetOk("database_name"); ok || d.HasChange("database_name") {
		t, err := expandSystemSqlDatabaseName(d, v, "database_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-name"] = t
		}
	}

	if v, ok := d.GetOk("database_type"); ok || d.HasChange("database_type") {
		t, err := expandSystemSqlDatabaseType(d, v, "database_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-type"] = t
		}
	}

	if v, ok := d.GetOk("device_count_high"); ok || d.HasChange("device_count_high") {
		t, err := expandSystemSqlDeviceCountHigh(d, v, "device_count_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-count-high"] = t
		}
	}

	if v, ok := d.GetOk("event_table_partition_time"); ok || d.HasChange("event_table_partition_time") {
		t, err := expandSystemSqlEventTablePartitionTime(d, v, "event_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-table-partition-time"] = t
		}
	}

	if v, ok := d.GetOk("fct_table_partition_time"); ok || d.HasChange("fct_table_partition_time") {
		t, err := expandSystemSqlFctTablePartitionTime(d, v, "fct_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fct-table-partition-time"] = t
		}
	}

	if v, ok := d.GetOk("logtype"); ok || d.HasChange("logtype") {
		t, err := expandSystemSqlLogtype(d, v, "logtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtype"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemSqlPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("prompt_sql_upgrade"); ok || d.HasChange("prompt_sql_upgrade") {
		t, err := expandSystemSqlPromptSqlUpgrade(d, v, "prompt_sql_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prompt-sql-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("rebuild_event"); ok || d.HasChange("rebuild_event") {
		t, err := expandSystemSqlRebuildEvent(d, v, "rebuild_event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rebuild-event"] = t
		}
	}

	if v, ok := d.GetOk("rebuild_event_start_time"); ok || d.HasChange("rebuild_event_start_time") {
		t, err := expandSystemSqlRebuildEventStartTime(d, v, "rebuild_event_start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rebuild-event-start-time"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemSqlServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandSystemSqlStartTime(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSqlStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("text_search_index"); ok || d.HasChange("text_search_index") {
		t, err := expandSystemSqlTextSearchIndex(d, v, "text_search_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["text-search-index"] = t
		}
	}

	if v, ok := d.GetOk("traffic_table_partition_time"); ok || d.HasChange("traffic_table_partition_time") {
		t, err := expandSystemSqlTrafficTablePartitionTime(d, v, "traffic_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-table-partition-time"] = t
		}
	}

	if v, ok := d.GetOk("ts_index_field"); ok || d.HasChange("ts_index_field") {
		t, err := expandSystemSqlTsIndexField(d, v, "ts_index_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ts-index-field"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemSqlUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("utm_table_partition_time"); ok || d.HasChange("utm_table_partition_time") {
		t, err := expandSystemSqlUtmTablePartitionTime(d, v, "utm_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-table-partition-time"] = t
		}
	}

	return &obj, nil
}
