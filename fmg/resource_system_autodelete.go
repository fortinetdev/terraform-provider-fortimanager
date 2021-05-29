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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed: true,
			},
		},
	}
}

func resourceSystemAutoDeleteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemAutoDelete(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoDelete resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoDelete(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemAutoDelete(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemAutoDelete(adomv, mkey, nil)
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

func flattenSystemAutoDeleteDlpFilesAutoDeletionSaa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteDlpFilesAutoDeletionRetentionSaa(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteDlpFilesAutoDeletionRunatSaa(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteDlpFilesAutoDeletionStatusSaa(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteDlpFilesAutoDeletionValueSaa(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRetentionSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "days",
			2: "weeks",
			3: "months",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionRunatSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionStatusSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteDlpFilesAutoDeletionValueSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionSaa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteLogAutoDeletionRetentionSaa(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteLogAutoDeletionRunatSaa(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteLogAutoDeletionStatusSaa(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteLogAutoDeletionValueSaa(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteLogAutoDeletionRetentionSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "days",
			2: "weeks",
			3: "months",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionRunatSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionStatusSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteLogAutoDeletionValueSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionSaa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionRetentionSaa(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionRunatSaa(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionStatusSaa(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteQuarantineFilesAutoDeletionValueSaa(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionRetentionSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "days",
			2: "weeks",
			3: "months",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionRunatSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionStatusSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteQuarantineFilesAutoDeletionValueSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionSaa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := i["retention"]; ok {
		result["retention"] = flattenSystemAutoDeleteReportAutoDeletionRetentionSaa(i["retention"], d, pre_append)
	}

	pre_append = pre + ".0." + "runat"
	if _, ok := i["runat"]; ok {
		result["runat"] = flattenSystemAutoDeleteReportAutoDeletionRunatSaa(i["runat"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemAutoDeleteReportAutoDeletionStatusSaa(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "value"
	if _, ok := i["value"]; ok {
		result["value"] = flattenSystemAutoDeleteReportAutoDeletionValueSaa(i["value"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAutoDeleteReportAutoDeletionRetentionSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "days",
			2: "weeks",
			3: "months",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionRunatSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionStatusSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemAutoDeleteReportAutoDeletionValueSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoDeleteStatusFakeSaa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoDelete(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("dlp_files_auto_deletion", flattenSystemAutoDeleteDlpFilesAutoDeletionSaa(o["dlp-files-auto-deletion"], d, "dlp_files_auto_deletion")); err != nil {
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
			if err = d.Set("dlp_files_auto_deletion", flattenSystemAutoDeleteDlpFilesAutoDeletionSaa(o["dlp-files-auto-deletion"], d, "dlp_files_auto_deletion")); err != nil {
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
		if err = d.Set("log_auto_deletion", flattenSystemAutoDeleteLogAutoDeletionSaa(o["log-auto-deletion"], d, "log_auto_deletion")); err != nil {
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
			if err = d.Set("log_auto_deletion", flattenSystemAutoDeleteLogAutoDeletionSaa(o["log-auto-deletion"], d, "log_auto_deletion")); err != nil {
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
		if err = d.Set("quarantine_files_auto_deletion", flattenSystemAutoDeleteQuarantineFilesAutoDeletionSaa(o["quarantine-files-auto-deletion"], d, "quarantine_files_auto_deletion")); err != nil {
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
			if err = d.Set("quarantine_files_auto_deletion", flattenSystemAutoDeleteQuarantineFilesAutoDeletionSaa(o["quarantine-files-auto-deletion"], d, "quarantine_files_auto_deletion")); err != nil {
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
		if err = d.Set("report_auto_deletion", flattenSystemAutoDeleteReportAutoDeletionSaa(o["report-auto-deletion"], d, "report_auto_deletion")); err != nil {
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
			if err = d.Set("report_auto_deletion", flattenSystemAutoDeleteReportAutoDeletionSaa(o["report-auto-deletion"], d, "report_auto_deletion")); err != nil {
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

	if err = d.Set("status_fake", flattenSystemAutoDeleteStatusFakeSaa(o["status-fake"], d, "status_fake")); err != nil {
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

func expandSystemAutoDeleteDlpFilesAutoDeletionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok {
		result["retention"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionRetentionSaa(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok {
		result["runat"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionRunatSaa(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionStatusSaa(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok {
		result["value"], _ = expandSystemAutoDeleteDlpFilesAutoDeletionValueSaa(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRetentionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionRunatSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionStatusSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteDlpFilesAutoDeletionValueSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok {
		result["retention"], _ = expandSystemAutoDeleteLogAutoDeletionRetentionSaa(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok {
		result["runat"], _ = expandSystemAutoDeleteLogAutoDeletionRunatSaa(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemAutoDeleteLogAutoDeletionStatusSaa(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok {
		result["value"], _ = expandSystemAutoDeleteLogAutoDeletionValueSaa(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteLogAutoDeletionRetentionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionRunatSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionStatusSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteLogAutoDeletionValueSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok {
		result["retention"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionRetentionSaa(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok {
		result["runat"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionRunatSaa(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionStatusSaa(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok {
		result["value"], _ = expandSystemAutoDeleteQuarantineFilesAutoDeletionValueSaa(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionRetentionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionRunatSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionStatusSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteQuarantineFilesAutoDeletionValueSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "retention"
	if _, ok := d.GetOk(pre_append); ok {
		result["retention"], _ = expandSystemAutoDeleteReportAutoDeletionRetentionSaa(d, i["retention"], pre_append)
	}
	pre_append = pre + ".0." + "runat"
	if _, ok := d.GetOk(pre_append); ok {
		result["runat"], _ = expandSystemAutoDeleteReportAutoDeletionRunatSaa(d, i["runat"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemAutoDeleteReportAutoDeletionStatusSaa(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "value"
	if _, ok := d.GetOk(pre_append); ok {
		result["value"], _ = expandSystemAutoDeleteReportAutoDeletionValueSaa(d, i["value"], pre_append)
	}

	return result, nil
}

func expandSystemAutoDeleteReportAutoDeletionRetentionSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionRunatSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionStatusSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteReportAutoDeletionValueSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoDeleteStatusFakeSaa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoDelete(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dlp_files_auto_deletion"); ok {
		t, err := expandSystemAutoDeleteDlpFilesAutoDeletionSaa(d, v, "dlp_files_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-files-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("log_auto_deletion"); ok {
		t, err := expandSystemAutoDeleteLogAutoDeletionSaa(d, v, "log_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_files_auto_deletion"); ok {
		t, err := expandSystemAutoDeleteQuarantineFilesAutoDeletionSaa(d, v, "quarantine_files_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-files-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("report_auto_deletion"); ok {
		t, err := expandSystemAutoDeleteReportAutoDeletionSaa(d, v, "report_auto_deletion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-auto-deletion"] = t
		}
	}

	if v, ok := d.GetOk("status_fake"); ok {
		t, err := expandSystemAutoDeleteStatusFakeSaa(d, v, "status_fake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-fake"] = t
		}
	}

	return &obj, nil
}
