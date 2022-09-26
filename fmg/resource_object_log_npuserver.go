// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure all the log servers and create the server groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectLogNpuServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectLogNpuServerUpdate,
		Read:   resourceObjectLogNpuServerRead,
		Update: resourceObjectLogNpuServerUpdate,
		Delete: resourceObjectLogNpuServerDelete,

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
			"log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_processor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netflow_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_gen_event": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_tx_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_user_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_start_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sw_log_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"server_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"template_tx_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"syslog_facility": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"syslog_severity": &schema.Schema{
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

func resourceObjectLogNpuServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectLogNpuServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectLogNpuServer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectLogNpuServer")

	return resourceObjectLogNpuServerRead(d, m)
}

func resourceObjectLogNpuServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectLogNpuServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectLogNpuServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectLogNpuServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectLogNpuServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectLogNpuServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectLogNpuServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectLogNpuServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectLogNpuServerLogProcessingOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerLogProcessorOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerNetflowVerOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupOlna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenObjectLogNpuServerServerGroupGroupNameOlna(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_format"
		if _, ok := i["log-format"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogFormatOlna(i["log-format"], d, pre_append)
			tmp["log_format"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogFormat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_gen_event"
		if _, ok := i["log-gen-event"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogGenEventOlna(i["log-gen-event"], d, pre_append)
			tmp["log_gen_event"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogGenEvent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_mode"
		if _, ok := i["log-mode"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogModeOlna(i["log-mode"], d, pre_append)
			tmp["log_mode"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_tx_mode"
		if _, ok := i["log-tx-mode"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogTxModeOlna(i["log-tx-mode"], d, pre_append)
			tmp["log_tx_mode"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogTxMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_user_info"
		if _, ok := i["log-user-info"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogUserInfoOlna(i["log-user-info"], d, pre_append)
			tmp["log_user_info"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogUserInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_number"
		if _, ok := i["server-number"]; ok {
			v := flattenObjectLogNpuServerServerGroupServerNumberOlna(i["server-number"], d, pre_append)
			tmp["server_number"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-ServerNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_start_id"
		if _, ok := i["server-start-id"]; ok {
			v := flattenObjectLogNpuServerServerGroupServerStartIdOlna(i["server-start-id"], d, pre_append)
			tmp["server_start_id"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-ServerStartId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sw_log_flags"
		if _, ok := i["sw-log-flags"]; ok {
			v := flattenObjectLogNpuServerServerGroupSwLogFlagsOlna(i["sw-log-flags"], d, pre_append)
			tmp["sw_log_flags"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-SwLogFlags")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectLogNpuServerServerGroupGroupNameOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogFormatOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogGenEventOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogModeOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogTxModeOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogUserInfoOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupServerNumberOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupServerStartIdOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupSwLogFlagsOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoOlna(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_port"
		if _, ok := i["dest-port"]; ok {
			v := flattenObjectLogNpuServerServerInfoDestPortOlna(i["dest-port"], d, pre_append)
			tmp["dest_port"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-DestPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectLogNpuServerServerInfoIdOlna(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_family"
		if _, ok := i["ip-family"]; ok {
			v := flattenObjectLogNpuServerServerInfoIpFamilyOlna(i["ip-family"], d, pre_append)
			tmp["ip_family"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-IpFamily")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_server"
		if _, ok := i["ipv4-server"]; ok {
			v := flattenObjectLogNpuServerServerInfoIpv4ServerOlna(i["ipv4-server"], d, pre_append)
			tmp["ipv4_server"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Ipv4Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_server"
		if _, ok := i["ipv6-server"]; ok {
			v := flattenObjectLogNpuServerServerInfoIpv6ServerOlna(i["ipv6-server"], d, pre_append)
			tmp["ipv6_server"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Ipv6Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := i["source-port"]; ok {
			v := flattenObjectLogNpuServerServerInfoSourcePortOlna(i["source-port"], d, pre_append)
			tmp["source_port"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-SourcePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template_tx_timeout"
		if _, ok := i["template-tx-timeout"]; ok {
			v := flattenObjectLogNpuServerServerInfoTemplateTxTimeoutOlna(i["template-tx-timeout"], d, pre_append)
			tmp["template_tx_timeout"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-TemplateTxTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectLogNpuServerServerInfoVdomOlna(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectLogNpuServerServerInfoDestPortOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIdOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpFamilyOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpv4ServerOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpv6ServerOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoSourcePortOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoTemplateTxTimeoutOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoVdomOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerSyslogFacilityOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerSyslogSeverityOlna(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectLogNpuServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("log_processing", flattenObjectLogNpuServerLogProcessingOlna(o["log-processing"], d, "log_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-processing"], "ObjectLogNpuServer-LogProcessing"); ok {
			if err = d.Set("log_processing", vv); err != nil {
				return fmt.Errorf("Error reading log_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_processing: %v", err)
		}
	}

	if err = d.Set("log_processor", flattenObjectLogNpuServerLogProcessorOlna(o["log-processor"], d, "log_processor")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-processor"], "ObjectLogNpuServer-LogProcessor"); ok {
			if err = d.Set("log_processor", vv); err != nil {
				return fmt.Errorf("Error reading log_processor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_processor: %v", err)
		}
	}

	if err = d.Set("netflow_ver", flattenObjectLogNpuServerNetflowVerOlna(o["netflow-ver"], d, "netflow_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["netflow-ver"], "ObjectLogNpuServer-NetflowVer"); ok {
			if err = d.Set("netflow_ver", vv); err != nil {
				return fmt.Errorf("Error reading netflow_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netflow_ver: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_group", flattenObjectLogNpuServerServerGroupOlna(o["server-group"], d, "server_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-group"], "ObjectLogNpuServer-ServerGroup"); ok {
				if err = d.Set("server_group", vv); err != nil {
					return fmt.Errorf("Error reading server_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_group"); ok {
			if err = d.Set("server_group", flattenObjectLogNpuServerServerGroupOlna(o["server-group"], d, "server_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-group"], "ObjectLogNpuServer-ServerGroup"); ok {
					if err = d.Set("server_group", vv); err != nil {
						return fmt.Errorf("Error reading server_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("server_info", flattenObjectLogNpuServerServerInfoOlna(o["server-info"], d, "server_info")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-info"], "ObjectLogNpuServer-ServerInfo"); ok {
				if err = d.Set("server_info", vv); err != nil {
					return fmt.Errorf("Error reading server_info: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_info: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_info"); ok {
			if err = d.Set("server_info", flattenObjectLogNpuServerServerInfoOlna(o["server-info"], d, "server_info")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-info"], "ObjectLogNpuServer-ServerInfo"); ok {
					if err = d.Set("server_info", vv); err != nil {
						return fmt.Errorf("Error reading server_info: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_info: %v", err)
				}
			}
		}
	}

	if err = d.Set("syslog_facility", flattenObjectLogNpuServerSyslogFacilityOlna(o["syslog-facility"], d, "syslog_facility")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-facility"], "ObjectLogNpuServer-SyslogFacility"); ok {
			if err = d.Set("syslog_facility", vv); err != nil {
				return fmt.Errorf("Error reading syslog_facility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_facility: %v", err)
		}
	}

	if err = d.Set("syslog_severity", flattenObjectLogNpuServerSyslogSeverityOlna(o["syslog-severity"], d, "syslog_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-severity"], "ObjectLogNpuServer-SyslogSeverity"); ok {
			if err = d.Set("syslog_severity", vv); err != nil {
				return fmt.Errorf("Error reading syslog_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_severity: %v", err)
		}
	}

	return nil
}

func flattenObjectLogNpuServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectLogNpuServerLogProcessingOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerLogProcessorOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerNetflowVerOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandObjectLogNpuServerServerGroupGroupNameOlna(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-format"], _ = expandObjectLogNpuServerServerGroupLogFormatOlna(d, i["log_format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_gen_event"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-gen-event"], _ = expandObjectLogNpuServerServerGroupLogGenEventOlna(d, i["log_gen_event"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-mode"], _ = expandObjectLogNpuServerServerGroupLogModeOlna(d, i["log_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_tx_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-tx-mode"], _ = expandObjectLogNpuServerServerGroupLogTxModeOlna(d, i["log_tx_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_user_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-user-info"], _ = expandObjectLogNpuServerServerGroupLogUserInfoOlna(d, i["log_user_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-number"], _ = expandObjectLogNpuServerServerGroupServerNumberOlna(d, i["server_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_start_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-start-id"], _ = expandObjectLogNpuServerServerGroupServerStartIdOlna(d, i["server_start_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sw_log_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sw-log-flags"], _ = expandObjectLogNpuServerServerGroupSwLogFlagsOlna(d, i["sw_log_flags"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectLogNpuServerServerGroupGroupNameOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogFormatOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogGenEventOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogModeOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogTxModeOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogUserInfoOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupServerNumberOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupServerStartIdOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupSwLogFlagsOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dest-port"], _ = expandObjectLogNpuServerServerInfoDestPortOlna(d, i["dest_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectLogNpuServerServerInfoIdOlna(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_family"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-family"], _ = expandObjectLogNpuServerServerInfoIpFamilyOlna(d, i["ip_family"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-server"], _ = expandObjectLogNpuServerServerInfoIpv4ServerOlna(d, i["ipv4_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-server"], _ = expandObjectLogNpuServerServerInfoIpv6ServerOlna(d, i["ipv6_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-port"], _ = expandObjectLogNpuServerServerInfoSourcePortOlna(d, i["source_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template_tx_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["template-tx-timeout"], _ = expandObjectLogNpuServerServerInfoTemplateTxTimeoutOlna(d, i["template_tx_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectLogNpuServerServerInfoVdomOlna(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectLogNpuServerServerInfoDestPortOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIdOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpFamilyOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpv4ServerOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpv6ServerOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoSourcePortOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoTemplateTxTimeoutOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoVdomOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerSyslogFacilityOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerSyslogSeverityOlna(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectLogNpuServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_processing"); ok || d.HasChange("log_processing") {
		t, err := expandObjectLogNpuServerLogProcessingOlna(d, v, "log_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-processing"] = t
		}
	}

	if v, ok := d.GetOk("log_processor"); ok || d.HasChange("log_processor") {
		t, err := expandObjectLogNpuServerLogProcessorOlna(d, v, "log_processor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-processor"] = t
		}
	}

	if v, ok := d.GetOk("netflow_ver"); ok || d.HasChange("netflow_ver") {
		t, err := expandObjectLogNpuServerNetflowVerOlna(d, v, "netflow_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netflow-ver"] = t
		}
	}

	if v, ok := d.GetOk("server_group"); ok || d.HasChange("server_group") {
		t, err := expandObjectLogNpuServerServerGroupOlna(d, v, "server_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-group"] = t
		}
	}

	if v, ok := d.GetOk("server_info"); ok || d.HasChange("server_info") {
		t, err := expandObjectLogNpuServerServerInfoOlna(d, v, "server_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-info"] = t
		}
	}

	if v, ok := d.GetOk("syslog_facility"); ok || d.HasChange("syslog_facility") {
		t, err := expandObjectLogNpuServerSyslogFacilityOlna(d, v, "syslog_facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-facility"] = t
		}
	}

	if v, ok := d.GetOk("syslog_severity"); ok || d.HasChange("syslog_severity") {
		t, err := expandObjectLogNpuServerSyslogSeverityOlna(d, v, "syslog_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-severity"] = t
		}
	}

	return &obj, nil
}
