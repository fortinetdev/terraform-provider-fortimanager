// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Set up DLP rules for this profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpProfileRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpProfileRuleCreate,
		Read:   resourceObjectDlpProfileRuleRead,
		Update: resourceObjectDlpProfileRuleUpdate,
		Delete: resourceObjectDlpProfileRuleDelete,

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
			"profile": &schema.Schema{
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
			"expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
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
			"sensitivity": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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

func resourceObjectDlpProfileRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectDlpProfileRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpProfileRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpProfileRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpProfileRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpProfileRuleRead(d, m)
}

func resourceObjectDlpProfileRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectDlpProfileRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpProfileRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpProfileRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpProfileRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpProfileRuleRead(d, m)
}

func resourceObjectDlpProfileRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectDlpProfileRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpProfileRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpProfileRuleRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectDlpProfileRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpProfileRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpProfileRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpProfileRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpProfileRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleArchive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleFileSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDlpProfileRuleFilterBy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleLabel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDlpProfileRuleMatchPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpProfileRuleSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpProfileRuleSensor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDlpProfileRuleSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpProfileRuleType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpProfileRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectDlpProfileRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectDlpProfileRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("archive", flattenObjectDlpProfileRuleArchive2edl(o["archive"], d, "archive")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive"], "ObjectDlpProfileRule-Archive"); ok {
			if err = d.Set("archive", vv); err != nil {
				return fmt.Errorf("Error reading archive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive: %v", err)
		}
	}

	if err = d.Set("expiry", flattenObjectDlpProfileRuleExpiry2edl(o["expiry"], d, "expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiry"], "ObjectDlpProfileRule-Expiry"); ok {
			if err = d.Set("expiry", vv); err != nil {
				return fmt.Errorf("Error reading expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiry: %v", err)
		}
	}

	if err = d.Set("file_size", flattenObjectDlpProfileRuleFileSize2edl(o["file-size"], d, "file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-size"], "ObjectDlpProfileRule-FileSize"); ok {
			if err = d.Set("file_size", vv); err != nil {
				return fmt.Errorf("Error reading file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_size: %v", err)
		}
	}

	if err = d.Set("file_type", flattenObjectDlpProfileRuleFileType2edl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "ObjectDlpProfileRule-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("filter_by", flattenObjectDlpProfileRuleFilterBy2edl(o["filter-by"], d, "filter_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter-by"], "ObjectDlpProfileRule-FilterBy"); ok {
			if err = d.Set("filter_by", vv); err != nil {
				return fmt.Errorf("Error reading filter_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter_by: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectDlpProfileRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDlpProfileRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("label", flattenObjectDlpProfileRuleLabel2edl(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "ObjectDlpProfileRule-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("match_percentage", flattenObjectDlpProfileRuleMatchPercentage2edl(o["match-percentage"], d, "match_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-percentage"], "ObjectDlpProfileRule-MatchPercentage"); ok {
			if err = d.Set("match_percentage", vv); err != nil {
				return fmt.Errorf("Error reading match_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_percentage: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDlpProfileRuleName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpProfileRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("proto", flattenObjectDlpProfileRuleProto2edl(o["proto"], d, "proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["proto"], "ObjectDlpProfileRule-Proto"); ok {
			if err = d.Set("proto", vv); err != nil {
				return fmt.Errorf("Error reading proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proto: %v", err)
		}
	}

	if err = d.Set("sensitivity", flattenObjectDlpProfileRuleSensitivity2edl(o["sensitivity"], d, "sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensitivity"], "ObjectDlpProfileRule-Sensitivity"); ok {
			if err = d.Set("sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensitivity: %v", err)
		}
	}

	if err = d.Set("sensor", flattenObjectDlpProfileRuleSensor2edl(o["sensor"], d, "sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensor"], "ObjectDlpProfileRule-Sensor"); ok {
			if err = d.Set("sensor", vv); err != nil {
				return fmt.Errorf("Error reading sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensor: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectDlpProfileRuleSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectDlpProfileRule-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectDlpProfileRuleType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectDlpProfileRule-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpProfileRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpProfileRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleArchive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleFileSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDlpProfileRuleFilterBy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleLabel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDlpProfileRuleMatchPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpProfileRuleSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpProfileRuleSensor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDlpProfileRuleSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpProfileRuleType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpProfileRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectDlpProfileRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("archive"); ok || d.HasChange("archive") {
		t, err := expandObjectDlpProfileRuleArchive2edl(d, v, "archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive"] = t
		}
	}

	if v, ok := d.GetOk("expiry"); ok || d.HasChange("expiry") {
		t, err := expandObjectDlpProfileRuleExpiry2edl(d, v, "expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiry"] = t
		}
	}

	if v, ok := d.GetOk("file_size"); ok || d.HasChange("file_size") {
		t, err := expandObjectDlpProfileRuleFileSize2edl(d, v, "file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size"] = t
		}
	}

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandObjectDlpProfileRuleFileType2edl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("filter_by"); ok || d.HasChange("filter_by") {
		t, err := expandObjectDlpProfileRuleFilterBy2edl(d, v, "filter_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-by"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectDlpProfileRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandObjectDlpProfileRuleLabel2edl(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("match_percentage"); ok || d.HasChange("match_percentage") {
		t, err := expandObjectDlpProfileRuleMatchPercentage2edl(d, v, "match_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-percentage"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDlpProfileRuleName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("proto"); ok || d.HasChange("proto") {
		t, err := expandObjectDlpProfileRuleProto2edl(d, v, "proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proto"] = t
		}
	}

	if v, ok := d.GetOk("sensitivity"); ok || d.HasChange("sensitivity") {
		t, err := expandObjectDlpProfileRuleSensitivity2edl(d, v, "sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("sensor"); ok || d.HasChange("sensor") {
		t, err := expandObjectDlpProfileRuleSensor2edl(d, v, "sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensor"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectDlpProfileRuleSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectDlpProfileRuleType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
