---
subcategory: "System SQL"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_sql_customskipidx"
description: |-
  List of aditional SQL skip index fields.
---

# fortimanager_system_sql_customskipidx
List of aditional SQL skip index fields.

~> This resource is a sub resource for variable `custom_skipidx` of resource `fortimanager_system_sql`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_sql_customskipidx" "trname" {
  device_type = "FortiGate"
  fosid       = "2"
  index_field = "srcip"
  log_type    = "content"
}
```

## Argument Reference


The following arguments are supported:


* `device_type` - Device type. FortiGate - Set device type to FortiGate. FortiManager - Set device type to FortiManager FortiClient - Set device type to FortiClient. FortiMail - Set device type to FortiMail. FortiWeb - Set device type to FortiWeb. FortiSandbox - Set device type to FortiSandbox FortiProxy - Set device type to FortiProxy Valid values: `FortiGate`, `FortiManager`, `FortiClient`, `FortiMail`, `FortiWeb`, `FortiSandbox`, `FortiProxy`.

* `fosid` - Add or Edit log index fields.
* `index_field` - Field to be added to skip index.
* `log_type` - Log type. app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  voip -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SqlCustomSkipidx can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_sql_customskipidx.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

