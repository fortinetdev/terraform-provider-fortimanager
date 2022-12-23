// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemReportSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReportSettingUpdate,
		Read:   resourceSystemReportSettingRead,
		Update: resourceSystemReportSettingUpdate,
		Delete: resourceSystemReportSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aggregate_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capwap_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"capwap_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude_capwap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hcache_lossless": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_rpt_pdf_rows": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_table_rows": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"report_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"template_auto_install": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"week_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemReportSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemReportSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemReportSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemReportSetting")

	return resourceSystemReportSettingRead(d, m)
}

func resourceSystemReportSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemReportSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReportSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReportSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemReportSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReportSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemReportSettingAggregateReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingCapwapPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingCapwapService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingExcludeCapwap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingHcacheLossless(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingLdapCacheTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingMaxRptPdfRows(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingMaxTableRows(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingReportPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingTemplateAutoInstall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportSettingWeekStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReportSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aggregate_report", flattenSystemReportSettingAggregateReport(o["aggregate-report"], d, "aggregate_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggregate-report"], "SystemReportSetting-AggregateReport"); ok {
			if err = d.Set("aggregate_report", vv); err != nil {
				return fmt.Errorf("Error reading aggregate_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggregate_report: %v", err)
		}
	}

	if err = d.Set("capwap_port", flattenSystemReportSettingCapwapPort(o["capwap-port"], d, "capwap_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["capwap-port"], "SystemReportSetting-CapwapPort"); ok {
			if err = d.Set("capwap_port", vv); err != nil {
				return fmt.Errorf("Error reading capwap_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capwap_port: %v", err)
		}
	}

	if err = d.Set("capwap_service", flattenSystemReportSettingCapwapService(o["capwap-service"], d, "capwap_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["capwap-service"], "SystemReportSetting-CapwapService"); ok {
			if err = d.Set("capwap_service", vv); err != nil {
				return fmt.Errorf("Error reading capwap_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capwap_service: %v", err)
		}
	}

	if err = d.Set("exclude_capwap", flattenSystemReportSettingExcludeCapwap(o["exclude-capwap"], d, "exclude_capwap")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-capwap"], "SystemReportSetting-ExcludeCapwap"); ok {
			if err = d.Set("exclude_capwap", vv); err != nil {
				return fmt.Errorf("Error reading exclude_capwap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_capwap: %v", err)
		}
	}

	if err = d.Set("hcache_lossless", flattenSystemReportSettingHcacheLossless(o["hcache-lossless"], d, "hcache_lossless")); err != nil {
		if vv, ok := fortiAPIPatch(o["hcache-lossless"], "SystemReportSetting-HcacheLossless"); ok {
			if err = d.Set("hcache_lossless", vv); err != nil {
				return fmt.Errorf("Error reading hcache_lossless: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hcache_lossless: %v", err)
		}
	}

	if err = d.Set("ldap_cache_timeout", flattenSystemReportSettingLdapCacheTimeout(o["ldap-cache-timeout"], d, "ldap_cache_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-cache-timeout"], "SystemReportSetting-LdapCacheTimeout"); ok {
			if err = d.Set("ldap_cache_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ldap_cache_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_cache_timeout: %v", err)
		}
	}

	if err = d.Set("max_rpt_pdf_rows", flattenSystemReportSettingMaxRptPdfRows(o["max-rpt-pdf-rows"], d, "max_rpt_pdf_rows")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-rpt-pdf-rows"], "SystemReportSetting-MaxRptPdfRows"); ok {
			if err = d.Set("max_rpt_pdf_rows", vv); err != nil {
				return fmt.Errorf("Error reading max_rpt_pdf_rows: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_rpt_pdf_rows: %v", err)
		}
	}

	if err = d.Set("max_table_rows", flattenSystemReportSettingMaxTableRows(o["max-table-rows"], d, "max_table_rows")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-table-rows"], "SystemReportSetting-MaxTableRows"); ok {
			if err = d.Set("max_table_rows", vv); err != nil {
				return fmt.Errorf("Error reading max_table_rows: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_table_rows: %v", err)
		}
	}

	if err = d.Set("report_priority", flattenSystemReportSettingReportPriority(o["report-priority"], d, "report_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["report-priority"], "SystemReportSetting-ReportPriority"); ok {
			if err = d.Set("report_priority", vv); err != nil {
				return fmt.Errorf("Error reading report_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report_priority: %v", err)
		}
	}

	if err = d.Set("template_auto_install", flattenSystemReportSettingTemplateAutoInstall(o["template-auto-install"], d, "template_auto_install")); err != nil {
		if vv, ok := fortiAPIPatch(o["template-auto-install"], "SystemReportSetting-TemplateAutoInstall"); ok {
			if err = d.Set("template_auto_install", vv); err != nil {
				return fmt.Errorf("Error reading template_auto_install: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template_auto_install: %v", err)
		}
	}

	if err = d.Set("week_start", flattenSystemReportSettingWeekStart(o["week-start"], d, "week_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["week-start"], "SystemReportSetting-WeekStart"); ok {
			if err = d.Set("week_start", vv); err != nil {
				return fmt.Errorf("Error reading week_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading week_start: %v", err)
		}
	}

	return nil
}

func flattenSystemReportSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReportSettingAggregateReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingCapwapPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingCapwapService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingExcludeCapwap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingHcacheLossless(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingLdapCacheTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingMaxRptPdfRows(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingMaxTableRows(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingReportPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingTemplateAutoInstall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportSettingWeekStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReportSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggregate_report"); ok || d.HasChange("aggregate_report") {
		t, err := expandSystemReportSettingAggregateReport(d, v, "aggregate_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-report"] = t
		}
	}

	if v, ok := d.GetOk("capwap_port"); ok || d.HasChange("capwap_port") {
		t, err := expandSystemReportSettingCapwapPort(d, v, "capwap_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capwap-port"] = t
		}
	}

	if v, ok := d.GetOk("capwap_service"); ok || d.HasChange("capwap_service") {
		t, err := expandSystemReportSettingCapwapService(d, v, "capwap_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capwap-service"] = t
		}
	}

	if v, ok := d.GetOk("exclude_capwap"); ok || d.HasChange("exclude_capwap") {
		t, err := expandSystemReportSettingExcludeCapwap(d, v, "exclude_capwap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-capwap"] = t
		}
	}

	if v, ok := d.GetOk("hcache_lossless"); ok || d.HasChange("hcache_lossless") {
		t, err := expandSystemReportSettingHcacheLossless(d, v, "hcache_lossless")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hcache-lossless"] = t
		}
	}

	if v, ok := d.GetOk("ldap_cache_timeout"); ok || d.HasChange("ldap_cache_timeout") {
		t, err := expandSystemReportSettingLdapCacheTimeout(d, v, "ldap_cache_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-cache-timeout"] = t
		}
	}

	if v, ok := d.GetOk("max_rpt_pdf_rows"); ok || d.HasChange("max_rpt_pdf_rows") {
		t, err := expandSystemReportSettingMaxRptPdfRows(d, v, "max_rpt_pdf_rows")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-rpt-pdf-rows"] = t
		}
	}

	if v, ok := d.GetOk("max_table_rows"); ok || d.HasChange("max_table_rows") {
		t, err := expandSystemReportSettingMaxTableRows(d, v, "max_table_rows")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-table-rows"] = t
		}
	}

	if v, ok := d.GetOk("report_priority"); ok || d.HasChange("report_priority") {
		t, err := expandSystemReportSettingReportPriority(d, v, "report_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-priority"] = t
		}
	}

	if v, ok := d.GetOk("template_auto_install"); ok || d.HasChange("template_auto_install") {
		t, err := expandSystemReportSettingTemplateAutoInstall(d, v, "template_auto_install")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-auto-install"] = t
		}
	}

	if v, ok := d.GetOk("week_start"); ok || d.HasChange("week_start") {
		t, err := expandSystemReportSettingWeekStart(d, v, "week_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["week-start"] = t
		}
	}

	return &obj, nil
}
