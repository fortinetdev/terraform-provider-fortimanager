// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Set up DLP filters for this sensor.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpSensorFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpSensorFilterCreate,
		Read:   resourceObjectDlpSensorFilterRead,
		Update: resourceObjectDlpSensorFilterUpdate,
		Delete: resourceObjectDlpSensorFilterDelete,

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
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"fp_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
				Computed: true,
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
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectDlpSensorFilterCreate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectDlpSensorFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpSensorFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectDlpSensorFilter(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpSensorFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpSensorFilterRead(d, m)
}

func resourceObjectDlpSensorFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectDlpSensorFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpSensorFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDlpSensorFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpSensorFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpSensorFilterRead(d, m)
}

func resourceObjectDlpSensorFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	wsParams["adom"] = adomv

	err = c.DeleteObjectDlpSensorFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpSensorFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpSensorFilterRead(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	paradict["sensor"] = sensor

	o, err := c.ReadObjectDlpSensorFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpSensorFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpSensorFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpSensorFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpSensorFilterAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterArchive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterCompanyIdentifier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFileSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDlpSensorFilterFilterBy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterFpSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDlpSensorFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterMatchPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpSensorFilterRegexp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDlpSensorFilterSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpSensorFilterType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpSensorFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectDlpSensorFilterAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectDlpSensorFilter-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("archive", flattenObjectDlpSensorFilterArchive2edl(o["archive"], d, "archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive"], "ObjectDlpSensorFilter-Archive"); ok {
			if err = d.Set("archive", vv); err != nil {
				return fmt.Errorf("Error reading archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive: %v", err)
		}
	}

	if err = d.Set("company_identifier", flattenObjectDlpSensorFilterCompanyIdentifier2edl(o["company-identifier"], d, "company_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["company-identifier"], "ObjectDlpSensorFilter-CompanyIdentifier"); ok {
			if err = d.Set("company_identifier", vv); err != nil {
				return fmt.Errorf("Error reading company_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading company_identifier: %v", err)
		}
	}

	if err = d.Set("expiry", flattenObjectDlpSensorFilterExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "ObjectDlpSensorFilter-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("file_size", flattenObjectDlpSensorFilterFileSize2edl(o["file-size"], d, "file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-size"], "ObjectDlpSensorFilter-FileSize"); ok {
			if err = d.Set("file_size", vv); err != nil {
				return fmt.Errorf("Error reading file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_size: %v", err)
		}
	}

	if err = d.Set("file_type", flattenObjectDlpSensorFilterFileType2edl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "ObjectDlpSensorFilter-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("filter_by", flattenObjectDlpSensorFilterFilterBy2edl(o["filter-by"], d, "filter_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-by"], "ObjectDlpSensorFilter-FilterBy"); ok {
			if err = d.Set("filter_by", vv); err != nil {
				return fmt.Errorf("Error reading filter_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_by: %v", err)
		}
	}

	if err = d.Set("fp_sensitivity", flattenObjectDlpSensorFilterFpSensitivity2edl(o["fp-sensitivity"], d, "fp_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["fp-sensitivity"], "ObjectDlpSensorFilter-FpSensitivity"); ok {
			if err = d.Set("fp_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading fp_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fp_sensitivity: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectDlpSensorFilterId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDlpSensorFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match_percentage", flattenObjectDlpSensorFilterMatchPercentage2edl(o["match-percentage"], d, "match_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-percentage"], "ObjectDlpSensorFilter-MatchPercentage"); ok {
			if err = d.Set("match_percentage", vv); err != nil {
				return fmt.Errorf("Error reading match_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_percentage: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDlpSensorFilterName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpSensorFilter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("proto", flattenObjectDlpSensorFilterProto2edl(o["proto"], d, "proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["proto"], "ObjectDlpSensorFilter-Proto"); ok {
			if err = d.Set("proto", vv); err != nil {
				return fmt.Errorf("Error reading proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("regexp", flattenObjectDlpSensorFilterRegexp2edl(o["regexp"], d, "regexp")); err != nil {
		if vv, ok := fortiAPIPatch(o["regexp"], "ObjectDlpSensorFilter-Regexp"); ok {
			if err = d.Set("regexp", vv); err != nil {
				return fmt.Errorf("Error reading regexp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regexp: %v", err)
		}
	}

	if err = d.Set("sensitivity", flattenObjectDlpSensorFilterSensitivity2edl(o["sensitivity"], d, "sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensitivity"], "ObjectDlpSensorFilter-Sensitivity"); ok {
			if err = d.Set("sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensitivity: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectDlpSensorFilterSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectDlpSensorFilter-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectDlpSensorFilterType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectDlpSensorFilter-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpSensorFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpSensorFilterAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterArchive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterCompanyIdentifier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFileSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDlpSensorFilterFilterBy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterFpSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDlpSensorFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterMatchPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpSensorFilterRegexp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDlpSensorFilterSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpSensorFilterType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpSensorFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectDlpSensorFilterAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("archive"); ok || d.HasChange("archive") {
		t, err := expandObjectDlpSensorFilterArchive2edl(d, v, "archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive"] = t
		}
	}

	if v, ok := d.GetOk("company_identifier"); ok || d.HasChange("company_identifier") {
		t, err := expandObjectDlpSensorFilterCompanyIdentifier2edl(d, v, "company_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company-identifier"] = t
		}
	}

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandObjectDlpSensorFilterExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("file_size"); ok || d.HasChange("file_size") {
		t, err := expandObjectDlpSensorFilterFileSize2edl(d, v, "file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size"] = t
		}
	}

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandObjectDlpSensorFilterFileType2edl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("filter_by"); ok || d.HasChange("filter_by") {
		t, err := expandObjectDlpSensorFilterFilterBy2edl(d, v, "filter_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-by"] = t
		}
	}

	if v, ok := d.GetOk("fp_sensitivity"); ok || d.HasChange("fp_sensitivity") {
		t, err := expandObjectDlpSensorFilterFpSensitivity2edl(d, v, "fp_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectDlpSensorFilterId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match_percentage"); ok || d.HasChange("match_percentage") {
		t, err := expandObjectDlpSensorFilterMatchPercentage2edl(d, v, "match_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-percentage"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDlpSensorFilterName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("proto"); ok || d.HasChange("proto") {
		t, err := expandObjectDlpSensorFilterProto2edl(d, v, "proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proto"] = t
		}
	}

	if v, ok := d.GetOk("regexp"); ok || d.HasChange("regexp") {
		t, err := expandObjectDlpSensorFilterRegexp2edl(d, v, "regexp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regexp"] = t
		}
	}

	if v, ok := d.GetOk("sensitivity"); ok || d.HasChange("sensitivity") {
		t, err := expandObjectDlpSensorFilterSensitivity2edl(d, v, "sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectDlpSensorFilterSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectDlpSensorFilterType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
