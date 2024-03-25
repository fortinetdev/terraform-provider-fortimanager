---
subcategory: "System SQL"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_sql"
description: |-
  SQL settings.
---

# fortimanager_system_sql
SQL settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_index`: `fortimanager_system_sql_customindex`
>- `custom_skipidx`: `fortimanager_system_sql_customskipidx`
>- `ts_index_field`: `fortimanager_system_sql_tsindexfield`



## Example Usage

```hcl
resource "fortimanager_system_sql" "trname" {
  password = ["fortinet"]
  server   = "192.168.1.1"
  status   = "enable"
  username = "admin"
}
```

## Argument Reference


The following arguments are supported:


* `background_rebuild` - Disable/Enable rebuild SQL database in the background. disable - Rebuild SQL database not in the background. enable - Rebuild SQL database in the background. Valid values: `disable`, `enable`.

* `compress_table_min_age` - Minimum age in days for SQL tables to be compressed.
* `custom_index` - Custom-Index. The structure of `custom_index` block is documented below.
* `custom_skipidx` - Custom-Skipidx. The structure of `custom_skipidx` block is documented below.
* `database_name` - Database name.
* `database_type` - Database type. mysql - MySQL database. postgres - PostgreSQL local database. Valid values: `mysql`, `postgres`.

* `device_count_high` - Must set to enable if the count of registered devices is greater than 8000. disable - Set to disable if device count is less than 8000. enable - Set to enable if device count is equal to or greater than 8000. Valid values: `disable`, `enable`.

* `event_table_partition_time` - Maximum SQL database table partitioning time range in minute (0 for unlimited) for event logs.
* `fct_table_partition_time` - Maximum SQL database table partitioning time range in minute (0 for unlimited) for FortiClient logs.
* `logtype` - Log type. none - None. app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  voip -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `none`, `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.

* `password` - Password for login remote database.
* `prompt_sql_upgrade` - Prompt to convert log database into SQL database at start time on GUI. disable - Do not prompt to upgrade log database to SQL database at start time on GUI. enable - Prompt to upgrade log database to SQL database at start time on GUI. Valid values: `disable`, `enable`.

* `rebuild_event` - Disable/Enable rebuild event during SQL database rebuilding. disable - Do not rebuild event during SQL database rebuilding. enable - Rebuild event during SQL database rebuilding. Valid values: `disable`, `enable`.

* `rebuild_event_start_time` - Rebuild event starting date and time <hh:mm yyyy/mm/dd>.
* `server` - Database IP or hostname.
* `start_time` - Start date and time <hh:mm yyyy/mm/dd>.
* `status` - SQL database status. disable - Disable SQL database. local - Enable local database. Valid values: `disable`, `local`.

* `text_search_index` - Disable/Enable text search index. disable - Do not create text search index. enable - Create text search index. Valid values: `disable`, `enable`.

* `traffic_table_partition_time` - Maximum SQL database table partitioning time range in minute (0 for unlimited) for traffic logs.
* `ts_index_field` - Ts-Index-Field. The structure of `ts_index_field` block is documented below.
* `username` - User name for login remote database.
* `utm_table_partition_time` - Maximum SQL database table partitioning time range in minute (0 for unlimited) for UTM logs.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_index` block supports:

* `case_sensitive` - Disable/Enable case sensitive index. disable - Build a case insensitive index. enable - Build a case sensitive index. Valid values: `disable`, `enable`.

* `device_type` - Device type. FortiGate - Device type to FortiGate. FortiMail - Device type to FortiMail. FortiWeb - Device type to FortiWeb. Valid values: `FortiGate`, `FortiMail`, `FortiWeb`.

* `id` - Add or Edit log index fields.
* `index_field` - Log field name to be indexed.
* `log_type` - Log type. app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  voip -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.


The `custom_skipidx` block supports:

* `device_type` - Device type. FortiGate - Set device type to FortiGate. FortiManager - Set device type to FortiManager FortiClient - Set device type to FortiClient. FortiMail - Set device type to FortiMail. FortiWeb - Set device type to FortiWeb. FortiSandbox - Set device type to FortiSandbox FortiProxy - Set device type to FortiProxy Valid values: `FortiGate`, `FortiManager`, `FortiClient`, `FortiMail`, `FortiWeb`, `FortiSandbox`, `FortiProxy`.

* `id` - Add or Edit log index fields.
* `index_field` - Field to be added to skip index.
* `log_type` - Log type. app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  voip -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.


The `ts_index_field` block supports:

* `category` - Category of text search index fields.
* `value` - Fields of text search index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Sql can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_sql.labelname SystemSql
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

