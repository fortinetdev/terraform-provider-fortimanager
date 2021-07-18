// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure shared traffic shaper.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallShaperTrafficShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallShaperTrafficShaperCreate,
		Read:   resourceObjectFirewallShaperTrafficShaperRead,
		Update: resourceObjectFirewallShaperTrafficShaperUpdate,
		Delete: resourceObjectFirewallShaperTrafficShaperDelete,

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
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_marking_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exceed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exceed_class_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exceed_dscp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maximum_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maximum_dscp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"overhead": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"per_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallShaperTrafficShaperCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallShaperTrafficShaper(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShaperTrafficShaper resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallShaperTrafficShaper(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShaperTrafficShaper resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallShaperTrafficShaperRead(d, m)
}

func resourceObjectFirewallShaperTrafficShaperUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallShaperTrafficShaper(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShaperTrafficShaper resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallShaperTrafficShaper(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShaperTrafficShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallShaperTrafficShaperRead(d, m)
}

func resourceObjectFirewallShaperTrafficShaperDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallShaperTrafficShaper(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallShaperTrafficShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallShaperTrafficShaperRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallShaperTrafficShaper(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShaperTrafficShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallShaperTrafficShaper(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShaperTrafficShaper resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallShaperTrafficShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperDscpMarkingMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperExceedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperExceedClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperExceedDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperMaximumBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperMaximumDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperOverhead(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperPerPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperTrafficShaperPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallShaperTrafficShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bandwidth_unit", flattenObjectFirewallShaperTrafficShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-unit"], "ObjectFirewallShaperTrafficShaper-BandwidthUnit"); ok {
			if err = d.Set("bandwidth_unit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenObjectFirewallShaperTrafficShaperDiffserv(o["diffserv"], d, "diffserv")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv"], "ObjectFirewallShaperTrafficShaper-Diffserv"); ok {
			if err = d.Set("diffserv", vv); err != nil {
				return fmt.Errorf("Error reading diffserv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenObjectFirewallShaperTrafficShaperDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "ObjectFirewallShaperTrafficShaper-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dscp_marking_method", flattenObjectFirewallShaperTrafficShaperDscpMarkingMethod(o["dscp-marking-method"], d, "dscp_marking_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-marking-method"], "ObjectFirewallShaperTrafficShaper-DscpMarkingMethod"); ok {
			if err = d.Set("dscp_marking_method", vv); err != nil {
				return fmt.Errorf("Error reading dscp_marking_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_marking_method: %v", err)
		}
	}

	if err = d.Set("exceed_bandwidth", flattenObjectFirewallShaperTrafficShaperExceedBandwidth(o["exceed-bandwidth"], d, "exceed_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["exceed-bandwidth"], "ObjectFirewallShaperTrafficShaper-ExceedBandwidth"); ok {
			if err = d.Set("exceed_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading exceed_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exceed_bandwidth: %v", err)
		}
	}

	if err = d.Set("exceed_class_id", flattenObjectFirewallShaperTrafficShaperExceedClassId(o["exceed-class-id"], d, "exceed_class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["exceed-class-id"], "ObjectFirewallShaperTrafficShaper-ExceedClassId"); ok {
			if err = d.Set("exceed_class_id", vv); err != nil {
				return fmt.Errorf("Error reading exceed_class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exceed_class_id: %v", err)
		}
	}

	if err = d.Set("exceed_dscp", flattenObjectFirewallShaperTrafficShaperExceedDscp(o["exceed-dscp"], d, "exceed_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["exceed-dscp"], "ObjectFirewallShaperTrafficShaper-ExceedDscp"); ok {
			if err = d.Set("exceed_dscp", vv); err != nil {
				return fmt.Errorf("Error reading exceed_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exceed_dscp: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenObjectFirewallShaperTrafficShaperGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["guaranteed-bandwidth"], "ObjectFirewallShaperTrafficShaper-GuaranteedBandwidth"); ok {
			if err = d.Set("guaranteed_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth", flattenObjectFirewallShaperTrafficShaperMaximumBandwidth(o["maximum-bandwidth"], d, "maximum_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-bandwidth"], "ObjectFirewallShaperTrafficShaper-MaximumBandwidth"); ok {
			if err = d.Set("maximum_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
		}
	}

	if err = d.Set("maximum_dscp", flattenObjectFirewallShaperTrafficShaperMaximumDscp(o["maximum-dscp"], d, "maximum_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["maximum-dscp"], "ObjectFirewallShaperTrafficShaper-MaximumDscp"); ok {
			if err = d.Set("maximum_dscp", vv); err != nil {
				return fmt.Errorf("Error reading maximum_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maximum_dscp: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallShaperTrafficShaperName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallShaperTrafficShaper-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overhead", flattenObjectFirewallShaperTrafficShaperOverhead(o["overhead"], d, "overhead")); err != nil {
		if vv, ok := fortiAPIPatch(o["overhead"], "ObjectFirewallShaperTrafficShaper-Overhead"); ok {
			if err = d.Set("overhead", vv); err != nil {
				return fmt.Errorf("Error reading overhead: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overhead: %v", err)
		}
	}

	if err = d.Set("per_policy", flattenObjectFirewallShaperTrafficShaperPerPolicy(o["per-policy"], d, "per_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-policy"], "ObjectFirewallShaperTrafficShaper-PerPolicy"); ok {
			if err = d.Set("per_policy", vv); err != nil {
				return fmt.Errorf("Error reading per_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_policy: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectFirewallShaperTrafficShaperPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectFirewallShaperTrafficShaper-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallShaperTrafficShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallShaperTrafficShaperBandwidthUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperDscpMarkingMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperExceedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperExceedClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperExceedDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperMaximumBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperMaximumDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperOverhead(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperPerPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperTrafficShaperPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallShaperTrafficShaper(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_unit"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperBandwidthUnit(d, v, "bandwidth_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-unit"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperDiffserv(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperDiffservcode(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dscp_marking_method"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperDscpMarkingMethod(d, v, "dscp_marking_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-marking-method"] = t
		}
	}

	if v, ok := d.GetOk("exceed_bandwidth"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperExceedBandwidth(d, v, "exceed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("exceed_class_id"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperExceedClassId(d, v, "exceed_class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-class-id"] = t
		}
	}

	if v, ok := d.GetOk("exceed_dscp"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperExceedDscp(d, v, "exceed_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-dscp"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperGuaranteedBandwidth(d, v, "guaranteed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("maximum_bandwidth"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperMaximumBandwidth(d, v, "maximum_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("maximum_dscp"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperMaximumDscp(d, v, "maximum_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-dscp"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overhead"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperOverhead(d, v, "overhead")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overhead"] = t
		}
	}

	if v, ok := d.GetOk("per_policy"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperPerPolicy(d, v, "per_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandObjectFirewallShaperTrafficShaperPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
