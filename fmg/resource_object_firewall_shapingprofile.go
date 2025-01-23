// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure shaping profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallShapingProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallShapingProfileCreate,
		Read:   resourceObjectFirewallShapingProfileRead,
		Update: resourceObjectFirewallShapingProfileUpdate,
		Delete: resourceObjectFirewallShapingProfileDelete,

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
			"default_class_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"npu_offloading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"shaping_entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"burst_in_msec": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cburst_in_msec": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"class_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"guaranteed_bandwidth_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"maximum_bandwidth_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"red_probability": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectFirewallShapingProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectFirewallShapingProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShapingProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallShapingProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShapingProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "profile_name"))

	return resourceObjectFirewallShapingProfileRead(d, m)
}

func resourceObjectFirewallShapingProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallShapingProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShapingProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallShapingProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShapingProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "profile_name"))

	return resourceObjectFirewallShapingProfileRead(d, m)
}

func resourceObjectFirewallShapingProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallShapingProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallShapingProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallShapingProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallShapingProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShapingProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallShapingProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShapingProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallShapingProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileDefaultClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectFirewallShapingProfileNpuOffloading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "burst_in_msec"
		if _, ok := i["burst-in-msec"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesBurstInMsec(i["burst-in-msec"], d, pre_append)
			tmp["burst_in_msec"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-BurstInMsec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cburst_in_msec"
		if _, ok := i["cburst-in-msec"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesCburstInMsec(i["cburst-in-msec"], d, pre_append)
			tmp["cburst_in_msec"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-CburstInMsec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := i["guaranteed-bandwidth-percentage"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(i["guaranteed-bandwidth-percentage"], d, pre_append)
			tmp["guaranteed_bandwidth_percentage"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-GuaranteedBandwidthPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit"
		if _, ok := i["limit"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesLimit(i["limit"], d, pre_append)
			tmp["limit"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-Limit")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max"
		if _, ok := i["max"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesMax(i["max"], d, pre_append)
			tmp["max"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-Max")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := i["maximum-bandwidth-percentage"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(i["maximum-bandwidth-percentage"], d, pre_append)
			tmp["maximum_bandwidth_percentage"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-MaximumBandwidthPercentage")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min"
		if _, ok := i["min"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesMin(i["min"], d, pre_append)
			tmp["min"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-Min")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "red_probability"
		if _, ok := i["red-probability"]; ok {
			v := flattenObjectFirewallShapingProfileShapingEntriesRedProbability(i["red-probability"], d, pre_append)
			tmp["red_probability"] = fortiAPISubPartPatch(v, "ObjectFirewallShapingProfile-ShapingEntries-RedProbability")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallShapingProfileShapingEntriesBurstInMsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesCburstInMsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesRedProbability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallShapingProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectFirewallShapingProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallShapingProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("default_class_id", flattenObjectFirewallShapingProfileDefaultClassId(o["default-class-id"], d, "default_class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-class-id"], "ObjectFirewallShapingProfile-DefaultClassId"); ok {
			if err = d.Set("default_class_id", vv); err != nil {
				return fmt.Errorf("Error reading default_class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_class_id: %v", err)
		}
	}

	if err = d.Set("npu_offloading", flattenObjectFirewallShapingProfileNpuOffloading(o["npu-offloading"], d, "npu_offloading")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-offloading"], "ObjectFirewallShapingProfile-NpuOffloading"); ok {
			if err = d.Set("npu_offloading", vv); err != nil {
				return fmt.Errorf("Error reading npu_offloading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_offloading: %v", err)
		}
	}

	if err = d.Set("profile_name", flattenObjectFirewallShapingProfileProfileName(o["profile-name"], d, "profile_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-name"], "ObjectFirewallShapingProfile-ProfileName"); ok {
			if err = d.Set("profile_name", vv); err != nil {
				return fmt.Errorf("Error reading profile_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shaping_entries", flattenObjectFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["shaping-entries"], "ObjectFirewallShapingProfile-ShapingEntries"); ok {
				if err = d.Set("shaping_entries", vv); err != nil {
					return fmt.Errorf("Error reading shaping_entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading shaping_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shaping_entries"); ok {
			if err = d.Set("shaping_entries", flattenObjectFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["shaping-entries"], "ObjectFirewallShapingProfile-ShapingEntries"); ok {
					if err = d.Set("shaping_entries", vv); err != nil {
						return fmt.Errorf("Error reading shaping_entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading shaping_entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenObjectFirewallShapingProfileType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallShapingProfile-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallShapingProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallShapingProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileDefaultClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileNpuOffloading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "burst_in_msec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["burst-in-msec"], _ = expandObjectFirewallShapingProfileShapingEntriesBurstInMsec(d, i["burst_in_msec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cburst_in_msec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cburst-in-msec"], _ = expandObjectFirewallShapingProfileShapingEntriesCburstInMsec(d, i["cburst_in_msec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandObjectFirewallShapingProfileShapingEntriesClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["guaranteed-bandwidth-percentage"], _ = expandObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d, i["guaranteed_bandwidth_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallShapingProfileShapingEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["limit"], _ = expandObjectFirewallShapingProfileShapingEntriesLimit(d, i["limit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["max"], _ = expandObjectFirewallShapingProfileShapingEntriesMax(d, i["max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["maximum-bandwidth-percentage"], _ = expandObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d, i["maximum_bandwidth_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["min"], _ = expandObjectFirewallShapingProfileShapingEntriesMin(d, i["min"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectFirewallShapingProfileShapingEntriesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "red_probability"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["red-probability"], _ = expandObjectFirewallShapingProfileShapingEntriesRedProbability(d, i["red_probability"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallShapingProfileShapingEntriesBurstInMsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesCburstInMsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesRedProbability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallShapingProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallShapingProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("default_class_id"); ok || d.HasChange("default_class_id") {
		t, err := expandObjectFirewallShapingProfileDefaultClassId(d, v, "default_class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-class-id"] = t
		}
	}

	if v, ok := d.GetOk("npu_offloading"); ok || d.HasChange("npu_offloading") {
		t, err := expandObjectFirewallShapingProfileNpuOffloading(d, v, "npu_offloading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offloading"] = t
		}
	}

	if v, ok := d.GetOk("profile_name"); ok || d.HasChange("profile_name") {
		t, err := expandObjectFirewallShapingProfileProfileName(d, v, "profile_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-name"] = t
		}
	}

	if v, ok := d.GetOk("shaping_entries"); ok || d.HasChange("shaping_entries") {
		t, err := expandObjectFirewallShapingProfileShapingEntries(d, v, "shaping_entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaping-entries"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallShapingProfileType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
