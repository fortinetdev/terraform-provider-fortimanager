// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFirewall ShapingProfileClasses

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallShapingProfileClasses() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallShapingProfileClassesCreate,
		Read:   resourceObjectFirewallShapingProfileClassesRead,
		Update: resourceObjectFirewallShapingProfileClassesUpdate,
		Delete: resourceObjectFirewallShapingProfileClassesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"class_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallShapingProfileClassesCreate(d *schema.ResourceData, m interface{}) error {
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

	shaping_profile := d.Get("shaping_profile").(string)
	paradict["shaping_profile"] = shaping_profile

	obj, err := getObjectObjectFirewallShapingProfileClasses(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShapingProfileClasses resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFirewallShapingProfileClasses(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFirewallShapingProfileClasses(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFirewallShapingProfileClasses resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFirewallShapingProfileClasses(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFirewallShapingProfileClasses resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallShapingProfileClassesRead(d, m)
}

func resourceObjectFirewallShapingProfileClassesUpdate(d *schema.ResourceData, m interface{}) error {
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

	shaping_profile := d.Get("shaping_profile").(string)
	paradict["shaping_profile"] = shaping_profile

	obj, err := getObjectObjectFirewallShapingProfileClasses(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShapingProfileClasses resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallShapingProfileClasses(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShapingProfileClasses resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallShapingProfileClassesRead(d, m)
}

func resourceObjectFirewallShapingProfileClassesDelete(d *schema.ResourceData, m interface{}) error {
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

	shaping_profile := d.Get("shaping_profile").(string)
	paradict["shaping_profile"] = shaping_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallShapingProfileClasses(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallShapingProfileClasses resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallShapingProfileClassesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallShapingProfileClasses(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallShapingProfileClasses resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallShapingProfileClasses(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShapingProfileClasses resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallShapingProfileClassesClassId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileClassesGuaranteedBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileClassesMaximumBandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileClassesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShapingProfileClassesPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallShapingProfileClasses(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("class_id", flattenObjectFirewallShapingProfileClassesClassId2edl(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "ObjectFirewallShapingProfileClasses-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenObjectFirewallShapingProfileClassesGuaranteedBandwidth2edl(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-bandwidth"], "ObjectFirewallShapingProfileClasses-GuaranteedBandwidth"); ok {
			if err = d.Set("guaranteed_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth", flattenObjectFirewallShapingProfileClassesMaximumBandwidth2edl(o["maximum-bandwidth"], d, "maximum_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-bandwidth"], "ObjectFirewallShapingProfileClasses-MaximumBandwidth"); ok {
			if err = d.Set("maximum_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallShapingProfileClassesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallShapingProfileClasses-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectFirewallShapingProfileClassesPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectFirewallShapingProfileClasses-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallShapingProfileClassesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallShapingProfileClassesClassId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileClassesGuaranteedBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileClassesMaximumBandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileClassesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShapingProfileClassesPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallShapingProfileClasses(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandObjectFirewallShapingProfileClassesClassId2edl(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth"); ok || d.HasChange("guaranteed_bandwidth") {
		t, err := expandObjectFirewallShapingProfileClassesGuaranteedBandwidth2edl(d, v, "guaranteed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("maximum_bandwidth"); ok || d.HasChange("maximum_bandwidth") {
		t, err := expandObjectFirewallShapingProfileClassesMaximumBandwidth2edl(d, v, "maximum_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallShapingProfileClassesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectFirewallShapingProfileClassesPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
