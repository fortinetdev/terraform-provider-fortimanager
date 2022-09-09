// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DLP sensors.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectDlpSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpSensorCreate,
		Read:   resourceObjectDlpSensorRead,
		Update: resourceObjectDlpSensorUpdate,
		Delete: resourceObjectDlpSensorDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dlp_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"archive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"company_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"file_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter_by": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fp_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"proto": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"regexp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"flow_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"full_archive_proto": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"nac_quar_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary_proto": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceObjectDlpSensorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDlpSensor(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpSensor resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpSensor(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpSensor resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpSensorRead(d, m)
}

func resourceObjectDlpSensorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDlpSensor(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpSensor resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpSensor(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpSensorRead(d, m)
}

func resourceObjectDlpSensorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDlpSensor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpSensorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDlpSensor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpSensor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpSensor resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpSensorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorDlpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectDlpSensorFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := i["archive"]; ok {
			v := flattenObjectDlpSensorFilterArchive(i["archive"], d, pre_append)
			tmp["archive"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Archive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := i["company-identifier"]; ok {
			v := flattenObjectDlpSensorFilterCompanyIdentifier(i["company-identifier"], d, pre_append)
			tmp["company_identifier"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-CompanyIdentifier")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := i["expiry"]; ok {
			v := flattenObjectDlpSensorFilterExpiry(i["expiry"], d, pre_append)
			tmp["expiry"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Expiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := i["file-size"]; ok {
			v := flattenObjectDlpSensorFilterFileSize(i["file-size"], d, pre_append)
			tmp["file_size"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-FileSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenObjectDlpSensorFilterFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := i["filter-by"]; ok {
			v := flattenObjectDlpSensorFilterFilterBy(i["filter-by"], d, pre_append)
			tmp["filter_by"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-FilterBy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fp_sensitivity"
		if _, ok := i["fp-sensitivity"]; ok {
			v := flattenObjectDlpSensorFilterFpSensitivity(i["fp-sensitivity"], d, pre_append)
			tmp["fp_sensitivity"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-FpSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDlpSensorFilterId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := i["match-percentage"]; ok {
			v := flattenObjectDlpSensorFilterMatchPercentage(i["match-percentage"], d, pre_append)
			tmp["match_percentage"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-MatchPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectDlpSensorFilterName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := i["proto"]; ok {
			v := flattenObjectDlpSensorFilterProto(i["proto"], d, pre_append)
			tmp["proto"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Proto")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {
			v := flattenObjectDlpSensorFilterRegexp(i["regexp"], d, pre_append)
			tmp["regexp"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Regexp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := i["sensitivity"]; ok {
			v := flattenObjectDlpSensorFilterSensitivity(i["sensitivity"], d, pre_append)
			tmp["sensitivity"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Sensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectDlpSensorFilterSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectDlpSensorFilterType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectDlpSensor-Filter-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDlpSensorFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterCompanyIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFilterBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFpSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterMatchPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpSensorFilterRegexp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFlowBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFullArchiveProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpSensorNacQuarLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorSummaryProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectDlpSensor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectDlpSensorComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDlpSensor-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dlp_log", flattenObjectDlpSensorDlpLog(o["dlp-log"], d, "dlp_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-log"], "ObjectDlpSensor-DlpLog"); ok {
			if err = d.Set("dlp_log", vv); err != nil {
				return fmt.Errorf("Error reading dlp_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_log: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenObjectDlpSensorExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectDlpSensor-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenObjectDlpSensorFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "ObjectDlpSensor-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filter", flattenObjectDlpSensorFilter(o["filter"], d, "filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["filter"], "ObjectDlpSensor-Filter"); ok {
				if err = d.Set("filter", vv); err != nil {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter"); ok {
			if err = d.Set("filter", flattenObjectDlpSensorFilter(o["filter"], d, "filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["filter"], "ObjectDlpSensor-Filter"); ok {
					if err = d.Set("filter", vv); err != nil {
						return fmt.Errorf("Error reading filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("flow_based", flattenObjectDlpSensorFlowBased(o["flow-based"], d, "flow_based")); err != nil {
		if vv, ok := fortiAPIPatch(o["flow-based"], "ObjectDlpSensor-FlowBased"); ok {
			if err = d.Set("flow_based", vv); err != nil {
				return fmt.Errorf("Error reading flow_based: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flow_based: %v", err)
		}
	}

	if err = d.Set("full_archive_proto", flattenObjectDlpSensorFullArchiveProto(o["full-archive-proto"], d, "full_archive_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["full-archive-proto"], "ObjectDlpSensor-FullArchiveProto"); ok {
			if err = d.Set("full_archive_proto", vv); err != nil {
				return fmt.Errorf("Error reading full_archive_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading full_archive_proto: %v", err)
		}
	}

	if err = d.Set("nac_quar_log", flattenObjectDlpSensorNacQuarLog(o["nac-quar-log"], d, "nac_quar_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-quar-log"], "ObjectDlpSensor-NacQuarLog"); ok {
			if err = d.Set("nac_quar_log", vv); err != nil {
				return fmt.Errorf("Error reading nac_quar_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_quar_log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDlpSensorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpSensor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectDlpSensorOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectDlpSensor-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectDlpSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectDlpSensor-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("summary_proto", flattenObjectDlpSensorSummaryProto(o["summary-proto"], d, "summary_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["summary-proto"], "ObjectDlpSensor-SummaryProto"); ok {
			if err = d.Set("summary_proto", vv); err != nil {
				return fmt.Errorf("Error reading summary_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading summary_proto: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpSensorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorDlpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectDlpSensorFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["archive"], _ = expandObjectDlpSensorFilterArchive(d, i["archive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["company-identifier"], _ = expandObjectDlpSensorFilterCompanyIdentifier(d, i["company_identifier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["expiry"], _ = expandObjectDlpSensorFilterExpiry(d, i["expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-size"], _ = expandObjectDlpSensorFilterFileSize(d, i["file_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandObjectDlpSensorFilterFileType(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter-by"], _ = expandObjectDlpSensorFilterFilterBy(d, i["filter_by"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fp_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fp-sensitivity"], _ = expandObjectDlpSensorFilterFpSensitivity(d, i["fp_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectDlpSensorFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-percentage"], _ = expandObjectDlpSensorFilterMatchPercentage(d, i["match_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectDlpSensorFilterName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["proto"], _ = expandObjectDlpSensorFilterProto(d, i["proto"], pre_append)
		} else {
			tmp["proto"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["regexp"], _ = expandObjectDlpSensorFilterRegexp(d, i["regexp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sensitivity"], _ = expandObjectDlpSensorFilterSensitivity(d, i["sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandObjectDlpSensorFilterSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectDlpSensorFilterType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDlpSensorFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterCompanyIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFilterBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFpSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterMatchPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpSensorFilterRegexp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFlowBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFullArchiveProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpSensorNacQuarLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorSummaryProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectDlpSensor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectDlpSensorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dlp_log"); ok || d.HasChange("dlp_log") {
		t, err := expandObjectDlpSensorDlpLog(d, v, "dlp_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandObjectDlpSensorExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandObjectDlpSensorFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandObjectDlpSensorFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("flow_based"); ok || d.HasChange("flow_based") {
		t, err := expandObjectDlpSensorFlowBased(d, v, "flow_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-based"] = t
		}
	}

	if v, ok := d.GetOk("full_archive_proto"); ok || d.HasChange("full_archive_proto") {
		t, err := expandObjectDlpSensorFullArchiveProto(d, v, "full_archive_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-archive-proto"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar_log"); ok || d.HasChange("nac_quar_log") {
		t, err := expandObjectDlpSensorNacQuarLog(d, v, "nac_quar_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar-log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDlpSensorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectDlpSensorOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectDlpSensorReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("summary_proto"); ok || d.HasChange("summary_proto") {
		t, err := expandObjectDlpSensorSummaryProto(d, v, "summary_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-proto"] = t
		}
	}

	return &obj, nil
}
