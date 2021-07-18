// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS egress queue policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSwitchControllerQosQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerQosQueuePolicyCreate,
		Read:   resourceObjectSwitchControllerQosQueuePolicyRead,
		Update: resourceObjectSwitchControllerQosQueuePolicyUpdate,
		Delete: resourceObjectSwitchControllerQosQueuePolicyDelete,

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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drop_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ecn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"rate_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
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

func resourceObjectSwitchControllerQosQueuePolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerQosQueuePolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerQosQueuePolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosQueuePolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosQueuePolicyRead(d, m)
}

func resourceObjectSwitchControllerQosQueuePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerQosQueuePolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerQosQueuePolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosQueuePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosQueuePolicyRead(d, m)
}

func resourceObjectSwitchControllerQosQueuePolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerQosQueuePolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerQosQueuePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerQosQueuePolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerQosQueuePolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosQueuePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerQosQueuePolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosQueuePolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueue(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := i["drop-policy"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy(i["drop-policy"], d, pre_append)
			tmp["drop_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-DropPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := i["ecn"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueEcn(i["ecn"], d, pre_append)
			tmp["ecn"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-Ecn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := i["max-rate"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRate(i["max-rate"], d, pre_append)
			tmp["max_rate"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-MaxRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := i["max-rate-percent"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(i["max-rate-percent"], d, pre_append)
			tmp["max_rate_percent"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-MaxRatePercent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := i["min-rate"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRate(i["min-rate"], d, pre_append)
			tmp["min_rate"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-MinRate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := i["min-rate-percent"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent(i["min-rate-percent"], d, pre_append)
			tmp["min_rate_percent"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-MinRatePercent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectSwitchControllerQosQueuePolicyCosQueueWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosQueuePolicy-CosQueue-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueEcn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyRateBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("cos_queue", flattenObjectSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
			if vv, ok := fortiAPIPatch(o["cos-queue"], "ObjectSwitchControllerQosQueuePolicy-CosQueue"); ok {
				if err = d.Set("cos_queue", vv); err != nil {
					return fmt.Errorf("Error reading cos_queue: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cos_queue"); ok {
			if err = d.Set("cos_queue", flattenObjectSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
				if vv, ok := fortiAPIPatch(o["cos-queue"], "ObjectSwitchControllerQosQueuePolicy-CosQueue"); ok {
					if err = d.Set("cos_queue", vv); err != nil {
						return fmt.Errorf("Error reading cos_queue: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cos_queue: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerQosQueuePolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerQosQueuePolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rate_by", flattenObjectSwitchControllerQosQueuePolicyRateBy(o["rate-by"], d, "rate_by")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-by"], "ObjectSwitchControllerQosQueuePolicy-RateBy"); ok {
			if err = d.Set("rate_by", vv); err != nil {
				return fmt.Errorf("Error reading rate_by: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_by: %v", err)
		}
	}

	if err = d.Set("schedule", flattenObjectSwitchControllerQosQueuePolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "ObjectSwitchControllerQosQueuePolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerQosQueuePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerQosQueuePolicyCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["drop-policy"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy(d, i["drop_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ecn"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueEcn(d, i["ecn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-rate"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRate(d, i["max_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-rate-percent"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d, i["max_rate_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-rate"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueMinRate(d, i["min_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-rate-percent"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d, i["min_rate_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandObjectSwitchControllerQosQueuePolicyCosQueueWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueEcn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMinRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyRateBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos_queue"); ok {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueue(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSwitchControllerQosQueuePolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rate_by"); ok {
		t, err := expandObjectSwitchControllerQosQueuePolicyRateBy(d, v, "rate_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-by"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandObjectSwitchControllerQosQueuePolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	return &obj, nil
}
