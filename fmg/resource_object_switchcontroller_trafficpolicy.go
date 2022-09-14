// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch traffic policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSwitchControllerTrafficPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerTrafficPolicyCreate,
		Read:   resourceObjectSwitchControllerTrafficPolicyRead,
		Update: resourceObjectSwitchControllerTrafficPolicyUpdate,
		Delete: resourceObjectSwitchControllerTrafficPolicyDelete,

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
			"cos_queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"guaranteed_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policer_status": &schema.Schema{
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

func resourceObjectSwitchControllerTrafficPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerTrafficPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerTrafficPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerTrafficPolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerTrafficPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerTrafficPolicyRead(d, m)
}

func resourceObjectSwitchControllerTrafficPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerTrafficPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerTrafficPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerTrafficPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerTrafficPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerTrafficPolicyRead(d, m)
}

func resourceObjectSwitchControllerTrafficPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerTrafficPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerTrafficPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerTrafficPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerTrafficPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerTrafficPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerTrafficPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerTrafficPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerTrafficPolicyCosQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyGuaranteedBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyMaximumBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyPolicerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerTrafficPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerTrafficPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cos_queue", flattenObjectSwitchControllerTrafficPolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos-queue"], "ObjectSwitchControllerTrafficPolicy-CosQueue"); ok {
			if err = d.Set("cos_queue", vv); err != nil {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos_queue: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerTrafficPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerTrafficPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenObjectSwitchControllerTrafficPolicyGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-bandwidth"], "ObjectSwitchControllerTrafficPolicy-GuaranteedBandwidth"); ok {
			if err = d.Set("guaranteed_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("guaranteed_burst", flattenObjectSwitchControllerTrafficPolicyGuaranteedBurst(o["guaranteed-burst"], d, "guaranteed_burst")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-burst"], "ObjectSwitchControllerTrafficPolicy-GuaranteedBurst"); ok {
			if err = d.Set("guaranteed_burst", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_burst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_burst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSwitchControllerTrafficPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSwitchControllerTrafficPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("maximum_burst", flattenObjectSwitchControllerTrafficPolicyMaximumBurst(o["maximum-burst"], d, "maximum_burst")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-burst"], "ObjectSwitchControllerTrafficPolicy-MaximumBurst"); ok {
			if err = d.Set("maximum_burst", vv); err != nil {
				return fmt.Errorf("Error reading maximum_burst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_burst: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerTrafficPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerTrafficPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policer_status", flattenObjectSwitchControllerTrafficPolicyPolicerStatus(o["policer-status"], d, "policer_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["policer-status"], "ObjectSwitchControllerTrafficPolicy-PolicerStatus"); ok {
			if err = d.Set("policer_status", vv); err != nil {
				return fmt.Errorf("Error reading policer_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policer_status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSwitchControllerTrafficPolicyType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSwitchControllerTrafficPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerTrafficPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerTrafficPolicyCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyGuaranteedBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyMaximumBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyPolicerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerTrafficPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerTrafficPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {
		t, err := expandObjectSwitchControllerTrafficPolicyCosQueue(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerTrafficPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth"); ok || d.HasChange("guaranteed_bandwidth") {
		t, err := expandObjectSwitchControllerTrafficPolicyGuaranteedBandwidth(d, v, "guaranteed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_burst"); ok || d.HasChange("guaranteed_burst") {
		t, err := expandObjectSwitchControllerTrafficPolicyGuaranteedBurst(d, v, "guaranteed_burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-burst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectSwitchControllerTrafficPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("maximum_burst"); ok || d.HasChange("maximum_burst") {
		t, err := expandObjectSwitchControllerTrafficPolicyMaximumBurst(d, v, "maximum_burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-burst"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerTrafficPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policer_status"); ok || d.HasChange("policer_status") {
		t, err := expandObjectSwitchControllerTrafficPolicyPolicerStatus(d, v, "policer_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policer-status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSwitchControllerTrafficPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
