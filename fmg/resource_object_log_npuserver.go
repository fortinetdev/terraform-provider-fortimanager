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
			"enforce_seq_order": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
							Computed: true,
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
							Computed: true,
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
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"log_transport": &schema.Schema{
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectLogNpuServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectLogNpuServer(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectLogNpuServer(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectLogNpuServer(mkey, paradict)
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

func flattenObjectLogNpuServerEnforceSeqOrder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerLogProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerLogProcessor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerNetflowVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectLogNpuServerServerGroupGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_format"
		if _, ok := i["log-format"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogFormat(i["log-format"], d, pre_append)
			tmp["log_format"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogFormat")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_gen_event"
		if _, ok := i["log-gen-event"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogGenEvent(i["log-gen-event"], d, pre_append)
			tmp["log_gen_event"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogGenEvent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_mode"
		if _, ok := i["log-mode"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogMode(i["log-mode"], d, pre_append)
			tmp["log_mode"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_tx_mode"
		if _, ok := i["log-tx-mode"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogTxMode(i["log-tx-mode"], d, pre_append)
			tmp["log_tx_mode"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogTxMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_user_info"
		if _, ok := i["log-user-info"]; ok {
			v := flattenObjectLogNpuServerServerGroupLogUserInfo(i["log-user-info"], d, pre_append)
			tmp["log_user_info"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-LogUserInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_number"
		if _, ok := i["server-number"]; ok {
			v := flattenObjectLogNpuServerServerGroupServerNumber(i["server-number"], d, pre_append)
			tmp["server_number"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-ServerNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_start_id"
		if _, ok := i["server-start-id"]; ok {
			v := flattenObjectLogNpuServerServerGroupServerStartId(i["server-start-id"], d, pre_append)
			tmp["server_start_id"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-ServerStartId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sw_log_flags"
		if _, ok := i["sw-log-flags"]; ok {
			v := flattenObjectLogNpuServerServerGroupSwLogFlags(i["sw-log-flags"], d, pre_append)
			tmp["sw_log_flags"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerGroup-SwLogFlags")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectLogNpuServerServerGroupGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogGenEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogTxMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupLogUserInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupServerNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupServerStartId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerGroupSwLogFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectLogNpuServerServerInfoDestPort(i["dest-port"], d, pre_append)
			tmp["dest_port"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-DestPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectLogNpuServerServerInfoId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_family"
		if _, ok := i["ip-family"]; ok {
			v := flattenObjectLogNpuServerServerInfoIpFamily(i["ip-family"], d, pre_append)
			tmp["ip_family"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-IpFamily")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_server"
		if _, ok := i["ipv4-server"]; ok {
			v := flattenObjectLogNpuServerServerInfoIpv4Server(i["ipv4-server"], d, pre_append)
			tmp["ipv4_server"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Ipv4Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_server"
		if _, ok := i["ipv6-server"]; ok {
			v := flattenObjectLogNpuServerServerInfoIpv6Server(i["ipv6-server"], d, pre_append)
			tmp["ipv6_server"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Ipv6Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_transport"
		if _, ok := i["log-transport"]; ok {
			v := flattenObjectLogNpuServerServerInfoLogTransport(i["log-transport"], d, pre_append)
			tmp["log_transport"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-LogTransport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := i["source-port"]; ok {
			v := flattenObjectLogNpuServerServerInfoSourcePort(i["source-port"], d, pre_append)
			tmp["source_port"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-SourcePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template_tx_timeout"
		if _, ok := i["template-tx-timeout"]; ok {
			v := flattenObjectLogNpuServerServerInfoTemplateTxTimeout(i["template-tx-timeout"], d, pre_append)
			tmp["template_tx_timeout"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-TemplateTxTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectLogNpuServerServerInfoVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectLogNpuServer-ServerInfo-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectLogNpuServerServerInfoDestPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpv4Server(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpv6Server(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoLogTransport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoTemplateTxTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectLogNpuServerSyslogFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerSyslogSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("enforce_seq_order", flattenObjectLogNpuServerEnforceSeqOrder(o["enforce-seq-order"], d, "enforce_seq_order")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-seq-order"], "ObjectLogNpuServer-EnforceSeqOrder"); ok {
			if err = d.Set("enforce_seq_order", vv); err != nil {
				return fmt.Errorf("Error reading enforce_seq_order: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_seq_order: %v", err)
		}
	}

	if err = d.Set("log_processing", flattenObjectLogNpuServerLogProcessing(o["log-processing"], d, "log_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-processing"], "ObjectLogNpuServer-LogProcessing"); ok {
			if err = d.Set("log_processing", vv); err != nil {
				return fmt.Errorf("Error reading log_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_processing: %v", err)
		}
	}

	if err = d.Set("log_processor", flattenObjectLogNpuServerLogProcessor(o["log-processor"], d, "log_processor")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-processor"], "ObjectLogNpuServer-LogProcessor"); ok {
			if err = d.Set("log_processor", vv); err != nil {
				return fmt.Errorf("Error reading log_processor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_processor: %v", err)
		}
	}

	if err = d.Set("netflow_ver", flattenObjectLogNpuServerNetflowVer(o["netflow-ver"], d, "netflow_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["netflow-ver"], "ObjectLogNpuServer-NetflowVer"); ok {
			if err = d.Set("netflow_ver", vv); err != nil {
				return fmt.Errorf("Error reading netflow_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netflow_ver: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_group", flattenObjectLogNpuServerServerGroup(o["server-group"], d, "server_group")); err != nil {
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
			if err = d.Set("server_group", flattenObjectLogNpuServerServerGroup(o["server-group"], d, "server_group")); err != nil {
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
		if err = d.Set("server_info", flattenObjectLogNpuServerServerInfo(o["server-info"], d, "server_info")); err != nil {
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
			if err = d.Set("server_info", flattenObjectLogNpuServerServerInfo(o["server-info"], d, "server_info")); err != nil {
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

	if err = d.Set("syslog_facility", flattenObjectLogNpuServerSyslogFacility(o["syslog-facility"], d, "syslog_facility")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-facility"], "ObjectLogNpuServer-SyslogFacility"); ok {
			if err = d.Set("syslog_facility", vv); err != nil {
				return fmt.Errorf("Error reading syslog_facility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_facility: %v", err)
		}
	}

	if err = d.Set("syslog_severity", flattenObjectLogNpuServerSyslogSeverity(o["syslog-severity"], d, "syslog_severity")); err != nil {
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

func expandObjectLogNpuServerEnforceSeqOrder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerLogProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerLogProcessor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerNetflowVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["group-name"], _ = expandObjectLogNpuServerServerGroupGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-format"], _ = expandObjectLogNpuServerServerGroupLogFormat(d, i["log_format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_gen_event"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-gen-event"], _ = expandObjectLogNpuServerServerGroupLogGenEvent(d, i["log_gen_event"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-mode"], _ = expandObjectLogNpuServerServerGroupLogMode(d, i["log_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_tx_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-tx-mode"], _ = expandObjectLogNpuServerServerGroupLogTxMode(d, i["log_tx_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_user_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-user-info"], _ = expandObjectLogNpuServerServerGroupLogUserInfo(d, i["log_user_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-number"], _ = expandObjectLogNpuServerServerGroupServerNumber(d, i["server_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_start_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-start-id"], _ = expandObjectLogNpuServerServerGroupServerStartId(d, i["server_start_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sw_log_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sw-log-flags"], _ = expandObjectLogNpuServerServerGroupSwLogFlags(d, i["sw_log_flags"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectLogNpuServerServerGroupGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogGenEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogTxMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupLogUserInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupServerNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupServerStartId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerGroupSwLogFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dest-port"], _ = expandObjectLogNpuServerServerInfoDestPort(d, i["dest_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectLogNpuServerServerInfoId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_family"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-family"], _ = expandObjectLogNpuServerServerInfoIpFamily(d, i["ip_family"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv4-server"], _ = expandObjectLogNpuServerServerInfoIpv4Server(d, i["ipv4_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-server"], _ = expandObjectLogNpuServerServerInfoIpv6Server(d, i["ipv6_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_transport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-transport"], _ = expandObjectLogNpuServerServerInfoLogTransport(d, i["log_transport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-port"], _ = expandObjectLogNpuServerServerInfoSourcePort(d, i["source_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template_tx_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["template-tx-timeout"], _ = expandObjectLogNpuServerServerInfoTemplateTxTimeout(d, i["template_tx_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectLogNpuServerServerInfoVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectLogNpuServerServerInfoDestPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpv4Server(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpv6Server(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoLogTransport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoTemplateTxTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectLogNpuServerSyslogFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerSyslogSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectLogNpuServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("enforce_seq_order"); ok || d.HasChange("enforce_seq_order") {
		t, err := expandObjectLogNpuServerEnforceSeqOrder(d, v, "enforce_seq_order")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-seq-order"] = t
		}
	}

	if v, ok := d.GetOk("log_processing"); ok || d.HasChange("log_processing") {
		t, err := expandObjectLogNpuServerLogProcessing(d, v, "log_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-processing"] = t
		}
	}

	if v, ok := d.GetOk("log_processor"); ok || d.HasChange("log_processor") {
		t, err := expandObjectLogNpuServerLogProcessor(d, v, "log_processor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-processor"] = t
		}
	}

	if v, ok := d.GetOk("netflow_ver"); ok || d.HasChange("netflow_ver") {
		t, err := expandObjectLogNpuServerNetflowVer(d, v, "netflow_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netflow-ver"] = t
		}
	}

	if v, ok := d.GetOk("server_group"); ok || d.HasChange("server_group") {
		t, err := expandObjectLogNpuServerServerGroup(d, v, "server_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-group"] = t
		}
	}

	if v, ok := d.GetOk("server_info"); ok || d.HasChange("server_info") {
		t, err := expandObjectLogNpuServerServerInfo(d, v, "server_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-info"] = t
		}
	}

	if v, ok := d.GetOk("syslog_facility"); ok || d.HasChange("syslog_facility") {
		t, err := expandObjectLogNpuServerSyslogFacility(d, v, "syslog_facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-facility"] = t
		}
	}

	if v, ok := d.GetOk("syslog_severity"); ok || d.HasChange("syslog_severity") {
		t, err := expandObjectLogNpuServerSyslogSeverity(d, v, "syslog_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-severity"] = t
		}
	}

	return &obj, nil
}
