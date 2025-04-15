// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: COS queue configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerQosQueuePolicyCosQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerQosQueuePolicyCosQueueCreate,
		Read:   resourceObjectSwitchControllerQosQueuePolicyCosQueueRead,
		Update: resourceObjectSwitchControllerQosQueuePolicyCosQueueUpdate,
		Delete: resourceObjectSwitchControllerQosQueuePolicyCosQueueDelete,

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
			"queue_policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			},
			"max_rate_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_rate_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerQosQueuePolicyCosQueueCreate(d *schema.ResourceData, m interface{}) error {
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

	queue_policy := d.Get("queue_policy").(string)
	paradict["queue_policy"] = queue_policy

	obj, err := getObjectObjectSwitchControllerQosQueuePolicyCosQueue(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosQueuePolicyCosQueue resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerQosQueuePolicyCosQueue(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosQueuePolicyCosQueue resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosQueuePolicyCosQueueRead(d, m)
}

func resourceObjectSwitchControllerQosQueuePolicyCosQueueUpdate(d *schema.ResourceData, m interface{}) error {
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

	queue_policy := d.Get("queue_policy").(string)
	paradict["queue_policy"] = queue_policy

	obj, err := getObjectObjectSwitchControllerQosQueuePolicyCosQueue(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosQueuePolicyCosQueue resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerQosQueuePolicyCosQueue(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosQueuePolicyCosQueue resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosQueuePolicyCosQueueRead(d, m)
}

func resourceObjectSwitchControllerQosQueuePolicyCosQueueDelete(d *schema.ResourceData, m interface{}) error {
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

	queue_policy := d.Get("queue_policy").(string)
	paradict["queue_policy"] = queue_policy

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerQosQueuePolicyCosQueue(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerQosQueuePolicyCosQueue resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerQosQueuePolicyCosQueueRead(d *schema.ResourceData, m interface{}) error {
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

	queue_policy := d.Get("queue_policy").(string)
	if queue_policy == "" {
		queue_policy = importOptionChecking(m.(*FortiClient).Cfg, "queue_policy")
		if queue_policy == "" {
			return fmt.Errorf("Parameter queue_policy is missing")
		}
		if err = d.Set("queue_policy", queue_policy); err != nil {
			return fmt.Errorf("Error set params queue_policy: %v", err)
		}
	}
	paradict["queue_policy"] = queue_policy

	o, err := c.ReadObjectSwitchControllerQosQueuePolicyCosQueue(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosQueuePolicyCosQueue resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerQosQueuePolicyCosQueue(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosQueuePolicyCosQueue resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueEcn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerQosQueuePolicyCosQueue(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("description", flattenObjectSwitchControllerQosQueuePolicyCosQueueDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerQosQueuePolicyCosQueue-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("drop_policy", flattenObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy2edl(o["drop-policy"], d, "drop_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop-policy"], "ObjectSwitchControllerQosQueuePolicyCosQueue-DropPolicy"); ok {
			if err = d.Set("drop_policy", vv); err != nil {
				return fmt.Errorf("Error reading drop_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop_policy: %v", err)
		}
	}

	if err = d.Set("ecn", flattenObjectSwitchControllerQosQueuePolicyCosQueueEcn2edl(o["ecn"], d, "ecn")); err != nil {
		if vv, ok := fortiAPIPatch(o["ecn"], "ObjectSwitchControllerQosQueuePolicyCosQueue-Ecn"); ok {
			if err = d.Set("ecn", vv); err != nil {
				return fmt.Errorf("Error reading ecn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ecn: %v", err)
		}
	}

	if err = d.Set("max_rate", flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRate2edl(o["max-rate"], d, "max_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-rate"], "ObjectSwitchControllerQosQueuePolicyCosQueue-MaxRate"); ok {
			if err = d.Set("max_rate", vv); err != nil {
				return fmt.Errorf("Error reading max_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_rate: %v", err)
		}
	}

	if err = d.Set("max_rate_percent", flattenObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent2edl(o["max-rate-percent"], d, "max_rate_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-rate-percent"], "ObjectSwitchControllerQosQueuePolicyCosQueue-MaxRatePercent"); ok {
			if err = d.Set("max_rate_percent", vv); err != nil {
				return fmt.Errorf("Error reading max_rate_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_rate_percent: %v", err)
		}
	}

	if err = d.Set("min_rate", flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRate2edl(o["min-rate"], d, "min_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-rate"], "ObjectSwitchControllerQosQueuePolicyCosQueue-MinRate"); ok {
			if err = d.Set("min_rate", vv); err != nil {
				return fmt.Errorf("Error reading min_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_rate: %v", err)
		}
	}

	if err = d.Set("min_rate_percent", flattenObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent2edl(o["min-rate-percent"], d, "min_rate_percent")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-rate-percent"], "ObjectSwitchControllerQosQueuePolicyCosQueue-MinRatePercent"); ok {
			if err = d.Set("min_rate_percent", vv); err != nil {
				return fmt.Errorf("Error reading min_rate_percent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_rate_percent: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerQosQueuePolicyCosQueueName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerQosQueuePolicyCosQueue-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectSwitchControllerQosQueuePolicyCosQueueWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectSwitchControllerQosQueuePolicyCosQueue-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerQosQueuePolicyCosQueueFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueEcn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMinRate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQueuePolicyCosQueueWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerQosQueuePolicyCosQueue(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("drop_policy"); ok || d.HasChange("drop_policy") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueDropPolicy2edl(d, v, "drop_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-policy"] = t
		}
	}

	if v, ok := d.GetOk("ecn"); ok || d.HasChange("ecn") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueEcn2edl(d, v, "ecn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ecn"] = t
		}
	}

	if v, ok := d.GetOk("max_rate"); ok || d.HasChange("max_rate") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRate2edl(d, v, "max_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-rate"] = t
		}
	}

	if v, ok := d.GetOk("max_rate_percent"); ok || d.HasChange("max_rate_percent") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueMaxRatePercent2edl(d, v, "max_rate_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-rate-percent"] = t
		}
	}

	if v, ok := d.GetOk("min_rate"); ok || d.HasChange("min_rate") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueMinRate2edl(d, v, "min_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-rate"] = t
		}
	}

	if v, ok := d.GetOk("min_rate_percent"); ok || d.HasChange("min_rate_percent") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueMinRatePercent2edl(d, v, "min_rate_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-rate-percent"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectSwitchControllerQosQueuePolicyCosQueueWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
