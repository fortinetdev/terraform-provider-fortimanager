// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Automatic deletion policy for logs, reports, archived, and quarantined files.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoDelete() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoDeleteUpdate,
		Read:   resourceSystemAutoDeleteRead,
		Update: resourceSystemAutoDeleteUpdate,
		Delete: resourceSystemAutoDeleteDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dlp_files_auto_deletion": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"runat": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_auto_deletion": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"runat": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"quarantine_files_auto_deletion": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"runat": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"report_auto_deletion": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"runat": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status_fake": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemAutoDeleteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAutoDelete(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDelete resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoDelete(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDelete resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoDelete")

	return resourceSystemAutoDeleteRead(d, m)
}

func resourceSystemAutoDeleteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemAutoDelete(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoDelete resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoDeleteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemAutoDelete(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDelete resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoDelete(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoDelete resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoDeleteDlpFilesAutoDeletion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteDlpFilesAutoDeletionRetention(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteDlpFilesAutoDeletionRunat(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteDlpFilesAutoDeletionStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteDlpFilesAutoDeletionValue(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteLogAutoDeletionRetention(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteLogAutoDeletionRunat(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteLogAutoDeletionStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteLogAutoDeletionValue(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteLogAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionRetention(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionRunat(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionValue(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteReportAutoDeletionRetention(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteReportAutoDeletionRunat(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteReportAutoDeletionStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteReportAutoDeletionValue(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteReportAutoDeletionRetention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionRunat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteStatusFake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDelete(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("dlp_files_auto_deletion", flattenSystemAutoDeleteDlpFilesAutoDeletion(o["dlp-files-auto-deletion"], d, "dlp_files_auto_deletion")); err != nil {
			if vv, ok := fortiAPIPatch(o["dlp-files-auto-deletion"], "SystemAutoDelete-DlpFilesAutoDeletion"); ok {
				if err = d.Set("dlp_files_auto_deletion", vv); err != nil {
					return fmt.Errorf("Error reading dlp_files_auto_deletion: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dlp_files_auto_deletion: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dlp_files_auto_deletion"); ok {
			if err = d.Set("dlp_files_auto_deletion", flattenSystemAutoDeleteDlpFilesAutoDeletion(o["dlp-files-auto-deletion"], d, "dlp_files_auto_deletion")); err != nil {
				if vv, ok := fortiAPIPatch(o["dlp-files-auto-deletion"], "SystemAutoDelete-DlpFilesAutoDeletion"); ok {
					if err = d.Set("dlp_files_auto_deletion", vv); err != nil {
						return fmt.Errorf("Error reading dlp_files_auto_deletion: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dlp_files_auto_deletion: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("log_auto_deletion", flattenSystemAutoDeleteLogAutoDeletion(o["log-auto-deletion"], d, "log_auto_deletion")); err != nil {
			if vv, ok := fortiAPIPatch(o["log-auto-deletion"], "SystemAutoDelete-LogAutoDeletion"); ok {
				if err = d.Set("log_auto_deletion", vv); err != nil {
					return fmt.Errorf("Error reading log_auto_deletion: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading log_auto_deletion: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_auto_deletion"); ok {
			if err = d.Set("log_auto_deletion", flattenSystemAutoDeleteLogAutoDeletion(o["log-auto-deletion"], d, "log_auto_deletion")); err != nil {
				if vv, ok := fortiAPIPatch(o["log-auto-deletion"], "SystemAutoDelete-LogAutoDeletion"); ok {
					if err = d.Set("log_auto_deletion", vv); err != nil {
						return fmt.Errorf("Error reading log_auto_deletion: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading log_auto_deletion: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("quarantine_files_auto_deletion", flattenSystemAutoDeleteQuarantineFilesAutoDeletion(o["quarantine-files-auto-deletion"], d, "quarantine_files_auto_deletion")); err != nil {
			if vv, ok := fortiAPIPatch(o["quarantine-files-auto-deletion"], "SystemAutoDelete-QuarantineFilesAutoDeletion"); ok {
				if err = d.Set("quarantine_files_auto_deletion", vv); err != nil {
					return fmt.Errorf("Error reading quarantine_files_auto_deletion: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading quarantine_files_auto_deletion: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quarantine_files_auto_deletion"); ok {
			if err = d.Set("quarantine_files_auto_deletion", flattenSystemAutoDeleteQuarantineFilesAutoDeletion(o["quarantine-files-auto-deletion"], d, "quarantine_files_auto_deletion")); err != nil {
				if vv, ok := fortiAPIPatch(o["quarantine-files-auto-deletion"], "SystemAutoDelete-QuarantineFilesAutoDeletion"); ok {
					if err = d.Set("quarantine_files_auto_deletion", vv); err != nil {
						return fmt.Errorf("Error reading quarantine_files_auto_deletion: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading quarantine_files_auto_deletion: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("report_auto_deletion", flattenSystemAutoDeleteReportAutoDeletion(o["report-auto-deletion"], d, "report_auto_deletion")); err != nil {
			if vv, ok := fortiAPIPatch(o["report-auto-deletion"], "SystemAutoDelete-ReportAutoDeletion"); ok {
				if err = d.Set("report_auto_deletion", vv); err != nil {
					return fmt.Errorf("Error reading report_auto_deletion: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading report_auto_deletion: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("report_auto_deletion"); ok {
			if err = d.Set("report_auto_deletion", flattenSystemAutoDeleteReportAutoDeletion(o["report-auto-deletion"], d, "report_auto_deletion")); err != nil {
				if vv, ok := fortiAPIPatch(o["report-auto-deletion"], "SystemAutoDelete-ReportAutoDeletion"); ok {
					if err = d.Set("report_auto_deletion", vv); err != nil {
						return fmt.Errorf("Error reading report_auto_deletion: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading report_auto_deletion: %v", err)
				}
			}
		}
	}

	if err = d.Set("status_fake", flattenSystemAutoDeleteStatusFake(o["status-fake"], d, "status_fake")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-fake"], "SystemAutoDelete-StatusFake"); ok {
			if err = d.Set("status_fake", vv); err != nil {
				return fmt.Errorf("Error reading status_fake: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_fake: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoDeleteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoDeleteDlpFilesAutoDeletion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retention"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionRetention(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["runat"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionRunat(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["value"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionValue(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retention"], _ = expandSystemAutoDeleteLogAutoDeletionRetention(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["runat"], _ = expandSystemAutoDeleteLogAutoDeletionRunat(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemAutoDeleteLogAutoDeletionStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["value"], _ = expandSystemAutoDeleteLogAutoDeletionValue(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteLogAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retention"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionRetention(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["runat"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionRunat(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["value"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionValue(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retention"], _ = expandSystemAutoDeleteReportAutoDeletionRetention(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["runat"], _ = expandSystemAutoDeleteReportAutoDeletionRunat(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemAutoDeleteReportAutoDeletionStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["value"], _ = expandSystemAutoDeleteReportAutoDeletionValue(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteReportAutoDeletionRetention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionRunat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteStatusFake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDelete(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dlp_files_auto_deletion"); ok || d.HasChange("dlp_files_auto_deletion") {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletion(d, v, "dlp_files_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-files-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("log_auto_deletion"); ok || d.HasChange("log_auto_deletion") {
		t, err := expandSystemAutoDeleteLogAutoDeletion(d, v, "log_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_files_auto_deletion"); ok || d.HasChange("quarantine_files_auto_deletion") {
		t, err := expandSystemAutoDeleteQuarantineFilesAutoDeletion(d, v, "quarantine_files_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-files-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("report_auto_deletion"); ok || d.HasChange("report_auto_deletion") {
		t, err := expandSystemAutoDeleteReportAutoDeletion(d, v, "report_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("status_fake"); ok || d.HasChange("status_fake") {
		t, err := expandSystemAutoDeleteStatusFake(d, v, "status_fake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-fake"] = t
		}
	}

	return &obj, nil
}
