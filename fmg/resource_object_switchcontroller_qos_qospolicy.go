// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSwitchControllerQosQosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerQosQosPolicyCreate,
		Read:   resourceObjectSwitchControllerQosQosPolicyRead,
		Update: resourceObjectSwitchControllerQosQosPolicyUpdate,
		Delete: resourceObjectSwitchControllerQosQosPolicyDelete,

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
			"default_cos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"queue_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_dot1p_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip_dscp_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerQosQosPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerQosQosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerQosQosPolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosQosPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosQosPolicyRead(d, m)
}

func resourceObjectSwitchControllerQosQosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerQosQosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerQosQosPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosQosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosQosPolicyRead(d, m)
}

func resourceObjectSwitchControllerQosQosPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerQosQosPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerQosQosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerQosQosPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerQosQosPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosQosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerQosQosPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosQosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerQosQosPolicyDefaultCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQosPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQosPolicyQueuePolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQosPolicyTrustDot1PMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosQosPolicyTrustIpDscpMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerQosQosPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("default_cos", flattenObjectSwitchControllerQosQosPolicyDefaultCos(o["default-cos"], d, "default_cos")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-cos"], "ObjectSwitchControllerQosQosPolicy-DefaultCos"); ok {
			if err = d.Set("default_cos", vv); err != nil {
				return fmt.Errorf("Error reading default_cos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_cos: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerQosQosPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerQosQosPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("queue_policy", flattenObjectSwitchControllerQosQosPolicyQueuePolicy(o["queue-policy"], d, "queue_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["queue-policy"], "ObjectSwitchControllerQosQosPolicy-QueuePolicy"); ok {
			if err = d.Set("queue_policy", vv); err != nil {
				return fmt.Errorf("Error reading queue_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading queue_policy: %v", err)
		}
	}

	if err = d.Set("trust_dot1p_map", flattenObjectSwitchControllerQosQosPolicyTrustDot1PMap(o["trust-dot1p-map"], d, "trust_dot1p_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-dot1p-map"], "ObjectSwitchControllerQosQosPolicy-TrustDot1PMap"); ok {
			if err = d.Set("trust_dot1p_map", vv); err != nil {
				return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
		}
	}

	if err = d.Set("trust_ip_dscp_map", flattenObjectSwitchControllerQosQosPolicyTrustIpDscpMap(o["trust-ip-dscp-map"], d, "trust_ip_dscp_map")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ip-dscp-map"], "ObjectSwitchControllerQosQosPolicy-TrustIpDscpMap"); ok {
			if err = d.Set("trust_ip_dscp_map", vv); err != nil {
				return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerQosQosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerQosQosPolicyDefaultCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQosPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQosPolicyQueuePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQosPolicyTrustDot1PMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosQosPolicyTrustIpDscpMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerQosQosPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_cos"); ok {
		t, err := expandObjectSwitchControllerQosQosPolicyDefaultCos(d, v, "default_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cos"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSwitchControllerQosQosPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("queue_policy"); ok {
		t, err := expandObjectSwitchControllerQosQosPolicyQueuePolicy(d, v, "queue_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue-policy"] = t
		}
	}

	if v, ok := d.GetOk("trust_dot1p_map"); ok {
		t, err := expandObjectSwitchControllerQosQosPolicyTrustDot1PMap(d, v, "trust_dot1p_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-dot1p-map"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_dscp_map"); ok {
		t, err := expandObjectSwitchControllerQosQosPolicyTrustIpDscpMap(d, v, "trust_ip_dscp_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-dscp-map"] = t
		}
	}

	return &obj, nil
}
