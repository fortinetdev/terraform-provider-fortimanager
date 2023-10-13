// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Define shaping entries of this shaping profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallShapingProfileShapingEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallShapingProfileShapingEntriesCreate,
		Read:   resourceObjectFirewallShapingProfileShapingEntriesRead,
		Update: resourceObjectFirewallShapingProfileShapingEntriesUpdate,
		Delete: resourceObjectFirewallShapingProfileShapingEntriesDelete,

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
			"shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceObjectFirewallShapingProfileShapingEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	shaping_profile := d.Get("shaping_profile").(string)
	paradict["shaping_profile"] = shaping_profile

	obj, err := getObjectObjectFirewallShapingProfileShapingEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShapingProfileShapingEntries resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallShapingProfileShapingEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShapingProfileShapingEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallShapingProfileShapingEntriesRead(d, m)
}

func resourceObjectFirewallShapingProfileShapingEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	shaping_profile := d.Get("shaping_profile").(string)
	paradict["shaping_profile"] = shaping_profile

	obj, err := getObjectObjectFirewallShapingProfileShapingEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShapingProfileShapingEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallShapingProfileShapingEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShapingProfileShapingEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallShapingProfileShapingEntriesRead(d, m)
}

func resourceObjectFirewallShapingProfileShapingEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	shaping_profile := d.Get("shaping_profile").(string)
	paradict["shaping_profile"] = shaping_profile

	err = c.DeleteObjectFirewallShapingProfileShapingEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallShapingProfileShapingEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallShapingProfileShapingEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	shaping_profile := d.Get("shaping_profile").(string)
	if shaping_profile == "" {
		shaping_profile = importOptionChecking(m.(*FortiClient).Cfg, "shaping_profile")
		if shaping_profile == "" {
			return fmt.Errorf("Parameter shaping_profile is missing")
		}
		if err = d.Set("shaping_profile", shaping_profile); err != nil {
			return fmt.Errorf("Error set params shaping_profile: %v", err)
		}
	}
	paradict["shaping_profile"] = shaping_profile

	o, err := c.ReadObjectFirewallShapingProfileShapingEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShapingProfileShapingEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallShapingProfileShapingEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShapingProfileShapingEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallShapingProfileShapingEntriesBurstInMsec2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesCburstInMsec2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesMax2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesMin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileShapingEntriesRedProbability2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallShapingProfileShapingEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("burst_in_msec", flattenObjectFirewallShapingProfileShapingEntriesBurstInMsec2edl(o["burst-in-msec"], d, "burst_in_msec")); err != nil {
		if vv, ok := fortiAPIPatch(o["burst-in-msec"], "ObjectFirewallShapingProfileShapingEntries-BurstInMsec"); ok {
			if err = d.Set("burst_in_msec", vv); err != nil {
				return fmt.Errorf("Error reading burst_in_msec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading burst_in_msec: %v", err)
		}
	}

	if err = d.Set("cburst_in_msec", flattenObjectFirewallShapingProfileShapingEntriesCburstInMsec2edl(o["cburst-in-msec"], d, "cburst_in_msec")); err != nil {
		if vv, ok := fortiAPIPatch(o["cburst-in-msec"], "ObjectFirewallShapingProfileShapingEntries-CburstInMsec"); ok {
			if err = d.Set("cburst_in_msec", vv); err != nil {
				return fmt.Errorf("Error reading cburst_in_msec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cburst_in_msec: %v", err)
		}
	}

	if err = d.Set("class_id", flattenObjectFirewallShapingProfileShapingEntriesClassId2edl(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "ObjectFirewallShapingProfileShapingEntries-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth_percentage", flattenObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(o["guaranteed-bandwidth-percentage"], d, "guaranteed_bandwidth_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-bandwidth-percentage"], "ObjectFirewallShapingProfileShapingEntries-GuaranteedBandwidthPercentage"); ok {
			if err = d.Set("guaranteed_bandwidth_percentage", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_bandwidth_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_bandwidth_percentage: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallShapingProfileShapingEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallShapingProfileShapingEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("limit", flattenObjectFirewallShapingProfileShapingEntriesLimit2edl(o["limit"], d, "limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit"], "ObjectFirewallShapingProfileShapingEntries-Limit"); ok {
			if err = d.Set("limit", vv); err != nil {
				return fmt.Errorf("Error reading limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit: %v", err)
		}
	}

	if err = d.Set("max", flattenObjectFirewallShapingProfileShapingEntriesMax2edl(o["max"], d, "max")); err != nil {
		if vv, ok := fortiAPIPatch(o["max"], "ObjectFirewallShapingProfileShapingEntries-Max"); ok {
			if err = d.Set("max", vv); err != nil {
				return fmt.Errorf("Error reading max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth_percentage", flattenObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(o["maximum-bandwidth-percentage"], d, "maximum_bandwidth_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-bandwidth-percentage"], "ObjectFirewallShapingProfileShapingEntries-MaximumBandwidthPercentage"); ok {
			if err = d.Set("maximum_bandwidth_percentage", vv); err != nil {
				return fmt.Errorf("Error reading maximum_bandwidth_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_bandwidth_percentage: %v", err)
		}
	}

	if err = d.Set("min", flattenObjectFirewallShapingProfileShapingEntriesMin2edl(o["min"], d, "min")); err != nil {
		if vv, ok := fortiAPIPatch(o["min"], "ObjectFirewallShapingProfileShapingEntries-Min"); ok {
			if err = d.Set("min", vv); err != nil {
				return fmt.Errorf("Error reading min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectFirewallShapingProfileShapingEntriesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectFirewallShapingProfileShapingEntries-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("red_probability", flattenObjectFirewallShapingProfileShapingEntriesRedProbability2edl(o["red-probability"], d, "red_probability")); err != nil {
		if vv, ok := fortiAPIPatch(o["red-probability"], "ObjectFirewallShapingProfileShapingEntries-RedProbability"); ok {
			if err = d.Set("red_probability", vv); err != nil {
				return fmt.Errorf("Error reading red_probability: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading red_probability: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallShapingProfileShapingEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallShapingProfileShapingEntriesBurstInMsec2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesCburstInMsec2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesMax2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesMin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileShapingEntriesRedProbability2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallShapingProfileShapingEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("burst_in_msec"); ok || d.HasChange("burst_in_msec") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesBurstInMsec2edl(d, v, "burst_in_msec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["burst-in-msec"] = t
		}
	}

	if v, ok := d.GetOk("cburst_in_msec"); ok || d.HasChange("cburst_in_msec") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesCburstInMsec2edl(d, v, "cburst_in_msec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cburst-in-msec"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesClassId2edl(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth_percentage"); ok || d.HasChange("guaranteed_bandwidth_percentage") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage2edl(d, v, "guaranteed_bandwidth_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth-percentage"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("limit"); ok || d.HasChange("limit") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesLimit2edl(d, v, "limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit"] = t
		}
	}

	if v, ok := d.GetOk("max"); ok || d.HasChange("max") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesMax2edl(d, v, "max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max"] = t
		}
	}

	if v, ok := d.GetOk("maximum_bandwidth_percentage"); ok || d.HasChange("maximum_bandwidth_percentage") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage2edl(d, v, "maximum_bandwidth_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-bandwidth-percentage"] = t
		}
	}

	if v, ok := d.GetOk("min"); ok || d.HasChange("min") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesMin2edl(d, v, "min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("red_probability"); ok || d.HasChange("red_probability") {
		t, err := expandObjectFirewallShapingProfileShapingEntriesRedProbability2edl(d, v, "red_probability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["red-probability"] = t
		}
	}

	return &obj, nil
}
