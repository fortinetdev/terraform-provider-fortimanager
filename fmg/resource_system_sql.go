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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
							Computed: true,
						},
						"index_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"index_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
				Computed: true,
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
				Computed: true,
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
				Computed: true,
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
				Computed: true,
			},
			"ts_index_field": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
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
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_table_partition_time": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemSqlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemSql(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSql resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSql(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemSql(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemSql(adomv, mkey, nil)
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

func flattenSystemSqlBackgroundRebuildSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCompressTableMinAgeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexSqa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSqlCustomIndexCaseSensitiveSqa(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			v := flattenSystemSqlCustomIndexDeviceTypeSqa(i["device-type"], d, pre_append)
			tmp["device_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-DeviceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSqlCustomIndexIdSqa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := i["index-field"]; ok {
			v := flattenSystemSqlCustomIndexIndexFieldSqa(i["index-field"], d, pre_append)
			tmp["index_field"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-IndexField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := i["log-type"]; ok {
			v := flattenSystemSqlCustomIndexLogTypeSqa(i["log-type"], d, pre_append)
			tmp["log_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomIndex-LogType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSqlCustomIndexCaseSensitiveSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexDeviceTypeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexIdSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexIndexFieldSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomIndexLogTypeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxSqa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSqlCustomSkipidxDeviceTypeSqa(i["device-type"], d, pre_append)
			tmp["device_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-DeviceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSqlCustomSkipidxIdSqa(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := i["index-field"]; ok {
			v := flattenSystemSqlCustomSkipidxIndexFieldSqa(i["index-field"], d, pre_append)
			tmp["index_field"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-IndexField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := i["log-type"]; ok {
			v := flattenSystemSqlCustomSkipidxLogTypeSqa(i["log-type"], d, pre_append)
			tmp["log_type"] = fortiAPISubPartPatch(v, "SystemSql-CustomSkipidx-LogType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSqlCustomSkipidxDeviceTypeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxIdSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxIndexFieldSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlCustomSkipidxLogTypeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlDatabaseNameSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlDatabaseTypeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlDeviceCountHighSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlEventTablePartitionTimeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlFctTablePartitionTimeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlLogtypeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlPasswordSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlPromptSqlUpgradeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlRebuildEventSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlRebuildEventStartTimeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlServerSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlStartTimeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSqlStatusSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTextSearchIndexSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTrafficTablePartitionTimeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTsIndexFieldSqa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSqlTsIndexFieldCategorySqa(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "SystemSql-TsIndexField-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemSqlTsIndexFieldValueSqa(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemSql-TsIndexField-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSqlTsIndexFieldCategorySqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlTsIndexFieldValueSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlUsernameSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSqlUtmTablePartitionTimeSqa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSql(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("background_rebuild", flattenSystemSqlBackgroundRebuildSqa(o["background-rebuild"], d, "background_rebuild")); err != nil {
		if vv, ok := fortiAPIPatch(o["background-rebuild"], "SystemSql-BackgroundRebuild"); ok {
			if err = d.Set("background_rebuild", vv); err != nil {
				return fmt.Errorf("Error reading background_rebuild: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading background_rebuild: %v", err)
		}
	}

	if err = d.Set("compress_table_min_age", flattenSystemSqlCompressTableMinAgeSqa(o["compress-table-min-age"], d, "compress_table_min_age")); err != nil {
		if vv, ok := fortiAPIPatch(o["compress-table-min-age"], "SystemSql-CompressTableMinAge"); ok {
			if err = d.Set("compress_table_min_age", vv); err != nil {
				return fmt.Errorf("Error reading compress_table_min_age: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compress_table_min_age: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_index", flattenSystemSqlCustomIndexSqa(o["custom-index"], d, "custom_index")); err != nil {
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
			if err = d.Set("custom_index", flattenSystemSqlCustomIndexSqa(o["custom-index"], d, "custom_index")); err != nil {
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
		if err = d.Set("custom_skipidx", flattenSystemSqlCustomSkipidxSqa(o["custom-skipidx"], d, "custom_skipidx")); err != nil {
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
			if err = d.Set("custom_skipidx", flattenSystemSqlCustomSkipidxSqa(o["custom-skipidx"], d, "custom_skipidx")); err != nil {
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

	if err = d.Set("database_name", flattenSystemSqlDatabaseNameSqa(o["database-name"], d, "database_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-name"], "SystemSql-DatabaseName"); ok {
			if err = d.Set("database_name", vv); err != nil {
				return fmt.Errorf("Error reading database_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_name: %v", err)
		}
	}

	if err = d.Set("database_type", flattenSystemSqlDatabaseTypeSqa(o["database-type"], d, "database_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-type"], "SystemSql-DatabaseType"); ok {
			if err = d.Set("database_type", vv); err != nil {
				return fmt.Errorf("Error reading database_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_type: %v", err)
		}
	}

	if err = d.Set("device_count_high", flattenSystemSqlDeviceCountHighSqa(o["device-count-high"], d, "device_count_high")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-count-high"], "SystemSql-DeviceCountHigh"); ok {
			if err = d.Set("device_count_high", vv); err != nil {
				return fmt.Errorf("Error reading device_count_high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_count_high: %v", err)
		}
	}

	if err = d.Set("event_table_partition_time", flattenSystemSqlEventTablePartitionTimeSqa(o["event-table-partition-time"], d, "event_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["event-table-partition-time"], "SystemSql-EventTablePartitionTime"); ok {
			if err = d.Set("event_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading event_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event_table_partition_time: %v", err)
		}
	}

	if err = d.Set("fct_table_partition_time", flattenSystemSqlFctTablePartitionTimeSqa(o["fct-table-partition-time"], d, "fct_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["fct-table-partition-time"], "SystemSql-FctTablePartitionTime"); ok {
			if err = d.Set("fct_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading fct_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fct_table_partition_time: %v", err)
		}
	}

	if err = d.Set("logtype", flattenSystemSqlLogtypeSqa(o["logtype"], d, "logtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtype"], "SystemSql-Logtype"); ok {
			if err = d.Set("logtype", vv); err != nil {
				return fmt.Errorf("Error reading logtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtype: %v", err)
		}
	}

	if err = d.Set("prompt_sql_upgrade", flattenSystemSqlPromptSqlUpgradeSqa(o["prompt-sql-upgrade"], d, "prompt_sql_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["prompt-sql-upgrade"], "SystemSql-PromptSqlUpgrade"); ok {
			if err = d.Set("prompt_sql_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading prompt_sql_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prompt_sql_upgrade: %v", err)
		}
	}

	if err = d.Set("rebuild_event", flattenSystemSqlRebuildEventSqa(o["rebuild-event"], d, "rebuild_event")); err != nil {
		if vv, ok := fortiAPIPatch(o["rebuild-event"], "SystemSql-RebuildEvent"); ok {
			if err = d.Set("rebuild_event", vv); err != nil {
				return fmt.Errorf("Error reading rebuild_event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rebuild_event: %v", err)
		}
	}

	if err = d.Set("rebuild_event_start_time", flattenSystemSqlRebuildEventStartTimeSqa(o["rebuild-event-start-time"], d, "rebuild_event_start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["rebuild-event-start-time"], "SystemSql-RebuildEventStartTime"); ok {
			if err = d.Set("rebuild_event_start_time", vv); err != nil {
				return fmt.Errorf("Error reading rebuild_event_start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rebuild_event_start_time: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSqlServerSqa(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemSql-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("start_time", flattenSystemSqlStartTimeSqa(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "SystemSql-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSqlStatusSqa(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSql-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("text_search_index", flattenSystemSqlTextSearchIndexSqa(o["text-search-index"], d, "text_search_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["text-search-index"], "SystemSql-TextSearchIndex"); ok {
			if err = d.Set("text_search_index", vv); err != nil {
				return fmt.Errorf("Error reading text_search_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading text_search_index: %v", err)
		}
	}

	if err = d.Set("traffic_table_partition_time", flattenSystemSqlTrafficTablePartitionTimeSqa(o["traffic-table-partition-time"], d, "traffic_table_partition_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-table-partition-time"], "SystemSql-TrafficTablePartitionTime"); ok {
			if err = d.Set("traffic_table_partition_time", vv); err != nil {
				return fmt.Errorf("Error reading traffic_table_partition_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_table_partition_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ts_index_field", flattenSystemSqlTsIndexFieldSqa(o["ts-index-field"], d, "ts_index_field")); err != nil {
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
			if err = d.Set("ts_index_field", flattenSystemSqlTsIndexFieldSqa(o["ts-index-field"], d, "ts_index_field")); err != nil {
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

	if err = d.Set("username", flattenSystemSqlUsernameSqa(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemSql-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("utm_table_partition_time", flattenSystemSqlUtmTablePartitionTimeSqa(o["utm-table-partition-time"], d, "utm_table_partition_time")); err != nil {
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

func expandSystemSqlBackgroundRebuildSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCompressTableMinAgeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandSystemSqlCustomIndexCaseSensitiveSqa(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-type"], _ = expandSystemSqlCustomIndexDeviceTypeSqa(d, i["device_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSqlCustomIndexIdSqa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index-field"], _ = expandSystemSqlCustomIndexIndexFieldSqa(d, i["index_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-type"], _ = expandSystemSqlCustomIndexLogTypeSqa(d, i["log_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSqlCustomIndexCaseSensitiveSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexDeviceTypeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexIdSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexIndexFieldSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomIndexLogTypeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-type"], _ = expandSystemSqlCustomSkipidxDeviceTypeSqa(d, i["device_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSqlCustomSkipidxIdSqa(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index-field"], _ = expandSystemSqlCustomSkipidxIndexFieldSqa(d, i["index_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-type"], _ = expandSystemSqlCustomSkipidxLogTypeSqa(d, i["log_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSqlCustomSkipidxDeviceTypeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxIdSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxIndexFieldSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlCustomSkipidxLogTypeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlDatabaseNameSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlDatabaseTypeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlDeviceCountHighSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlEventTablePartitionTimeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlFctTablePartitionTimeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlLogtypeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlPasswordSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlPromptSqlUpgradeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlRebuildEventSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlRebuildEventStartTimeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlServerSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlStartTimeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSqlStatusSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTextSearchIndexSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTrafficTablePartitionTimeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTsIndexFieldSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandSystemSqlTsIndexFieldCategorySqa(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemSqlTsIndexFieldValueSqa(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSqlTsIndexFieldCategorySqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlTsIndexFieldValueSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlUsernameSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSqlUtmTablePartitionTimeSqa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSql(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("background_rebuild"); ok || d.HasChange("background_rebuild") {
		t, err := expandSystemSqlBackgroundRebuildSqa(d, v, "background_rebuild")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["background-rebuild"] = t
		}
	}

	if v, ok := d.GetOk("compress_table_min_age"); ok || d.HasChange("compress_table_min_age") {
		t, err := expandSystemSqlCompressTableMinAgeSqa(d, v, "compress_table_min_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compress-table-min-age"] = t
		}
	}

	if v, ok := d.GetOk("custom_index"); ok || d.HasChange("custom_index") {
		t, err := expandSystemSqlCustomIndexSqa(d, v, "custom_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-index"] = t
		}
	}

	if v, ok := d.GetOk("custom_skipidx"); ok || d.HasChange("custom_skipidx") {
		t, err := expandSystemSqlCustomSkipidxSqa(d, v, "custom_skipidx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-skipidx"] = t
		}
	}

	if v, ok := d.GetOk("database_name"); ok || d.HasChange("database_name") {
		t, err := expandSystemSqlDatabaseNameSqa(d, v, "database_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-name"] = t
		}
	}

	if v, ok := d.GetOk("database_type"); ok || d.HasChange("database_type") {
		t, err := expandSystemSqlDatabaseTypeSqa(d, v, "database_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-type"] = t
		}
	}

	if v, ok := d.GetOk("device_count_high"); ok || d.HasChange("device_count_high") {
		t, err := expandSystemSqlDeviceCountHighSqa(d, v, "device_count_high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-count-high"] = t
		}
	}

	if v, ok := d.GetOk("event_table_partition_time"); ok || d.HasChange("event_table_partition_time") {
		t, err := expandSystemSqlEventTablePartitionTimeSqa(d, v, "event_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-table-partition-time"] = t
		}
	}

	if v, ok := d.GetOk("fct_table_partition_time"); ok || d.HasChange("fct_table_partition_time") {
		t, err := expandSystemSqlFctTablePartitionTimeSqa(d, v, "fct_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fct-table-partition-time"] = t
		}
	}

	if v, ok := d.GetOk("logtype"); ok || d.HasChange("logtype") {
		t, err := expandSystemSqlLogtypeSqa(d, v, "logtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtype"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemSqlPasswordSqa(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("prompt_sql_upgrade"); ok || d.HasChange("prompt_sql_upgrade") {
		t, err := expandSystemSqlPromptSqlUpgradeSqa(d, v, "prompt_sql_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prompt-sql-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("rebuild_event"); ok || d.HasChange("rebuild_event") {
		t, err := expandSystemSqlRebuildEventSqa(d, v, "rebuild_event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rebuild-event"] = t
		}
	}

	if v, ok := d.GetOk("rebuild_event_start_time"); ok || d.HasChange("rebuild_event_start_time") {
		t, err := expandSystemSqlRebuildEventStartTimeSqa(d, v, "rebuild_event_start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rebuild-event-start-time"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemSqlServerSqa(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandSystemSqlStartTimeSqa(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSqlStatusSqa(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("text_search_index"); ok || d.HasChange("text_search_index") {
		t, err := expandSystemSqlTextSearchIndexSqa(d, v, "text_search_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["text-search-index"] = t
		}
	}

	if v, ok := d.GetOk("traffic_table_partition_time"); ok || d.HasChange("traffic_table_partition_time") {
		t, err := expandSystemSqlTrafficTablePartitionTimeSqa(d, v, "traffic_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-table-partition-time"] = t
		}
	}

	if v, ok := d.GetOk("ts_index_field"); ok || d.HasChange("ts_index_field") {
		t, err := expandSystemSqlTsIndexFieldSqa(d, v, "ts_index_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ts-index-field"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemSqlUsernameSqa(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("utm_table_partition_time"); ok || d.HasChange("utm_table_partition_time") {
		t, err := expandSystemSqlUtmTablePartitionTimeSqa(d, v, "utm_table_partition_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-table-partition-time"] = t
		}
	}

	return &obj, nil
}
