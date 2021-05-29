// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS 802.1p.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSwitchControllerQosDot1PMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerQosDot1PMapCreate,
		Read:   resourceObjectSwitchControllerQosDot1PMapRead,
		Update: resourceObjectSwitchControllerQosDot1PMapUpdate,
		Delete: resourceObjectSwitchControllerQosDot1PMapDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_pri_tagging": &schema.Schema{
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
			"priority_0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSwitchControllerQosDot1PMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerQosDot1PMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosDot1PMap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerQosDot1PMap(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosDot1PMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosDot1PMapRead(d, m)
}

func resourceObjectSwitchControllerQosDot1PMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerQosDot1PMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosDot1PMap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerQosDot1PMap(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosDot1PMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosDot1PMapRead(d, m)
}

func resourceObjectSwitchControllerQosDot1PMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerQosDot1PMap(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerQosDot1PMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerQosDot1PMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerQosDot1PMap(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosDot1PMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerQosDot1PMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosDot1PMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerQosDot1PMapDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosDot1PMapEgressPriTagging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectSwitchControllerQosDot1PMapPriority7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "queue-0",
			1: "queue-1",
			2: "queue-2",
			3: "queue-3",
			4: "queue-4",
			5: "queue-5",
			6: "queue-6",
			7: "queue-7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectSwitchControllerQosDot1PMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("description", flattenObjectSwitchControllerQosDot1PMapDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerQosDot1PMap-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("egress_pri_tagging", flattenObjectSwitchControllerQosDot1PMapEgressPriTagging(o["egress-pri-tagging"], d, "egress_pri_tagging")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-pri-tagging"], "ObjectSwitchControllerQosDot1PMap-EgressPriTagging"); ok {
			if err = d.Set("egress_pri_tagging", vv); err != nil {
				return fmt.Errorf("Error reading egress_pri_tagging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_pri_tagging: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerQosDot1PMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerQosDot1PMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priority_0", flattenObjectSwitchControllerQosDot1PMapPriority0(o["priority-0"], d, "priority_0")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-0"], "ObjectSwitchControllerQosDot1PMap-Priority0"); ok {
			if err = d.Set("priority_0", vv); err != nil {
				return fmt.Errorf("Error reading priority_0: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_0: %v", err)
		}
	}

	if err = d.Set("priority_1", flattenObjectSwitchControllerQosDot1PMapPriority1(o["priority-1"], d, "priority_1")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-1"], "ObjectSwitchControllerQosDot1PMap-Priority1"); ok {
			if err = d.Set("priority_1", vv); err != nil {
				return fmt.Errorf("Error reading priority_1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_1: %v", err)
		}
	}

	if err = d.Set("priority_2", flattenObjectSwitchControllerQosDot1PMapPriority2(o["priority-2"], d, "priority_2")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-2"], "ObjectSwitchControllerQosDot1PMap-Priority2"); ok {
			if err = d.Set("priority_2", vv); err != nil {
				return fmt.Errorf("Error reading priority_2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_2: %v", err)
		}
	}

	if err = d.Set("priority_3", flattenObjectSwitchControllerQosDot1PMapPriority3(o["priority-3"], d, "priority_3")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-3"], "ObjectSwitchControllerQosDot1PMap-Priority3"); ok {
			if err = d.Set("priority_3", vv); err != nil {
				return fmt.Errorf("Error reading priority_3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_3: %v", err)
		}
	}

	if err = d.Set("priority_4", flattenObjectSwitchControllerQosDot1PMapPriority4(o["priority-4"], d, "priority_4")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-4"], "ObjectSwitchControllerQosDot1PMap-Priority4"); ok {
			if err = d.Set("priority_4", vv); err != nil {
				return fmt.Errorf("Error reading priority_4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_4: %v", err)
		}
	}

	if err = d.Set("priority_5", flattenObjectSwitchControllerQosDot1PMapPriority5(o["priority-5"], d, "priority_5")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-5"], "ObjectSwitchControllerQosDot1PMap-Priority5"); ok {
			if err = d.Set("priority_5", vv); err != nil {
				return fmt.Errorf("Error reading priority_5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_5: %v", err)
		}
	}

	if err = d.Set("priority_6", flattenObjectSwitchControllerQosDot1PMapPriority6(o["priority-6"], d, "priority_6")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-6"], "ObjectSwitchControllerQosDot1PMap-Priority6"); ok {
			if err = d.Set("priority_6", vv); err != nil {
				return fmt.Errorf("Error reading priority_6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_6: %v", err)
		}
	}

	if err = d.Set("priority_7", flattenObjectSwitchControllerQosDot1PMapPriority7(o["priority-7"], d, "priority_7")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-7"], "ObjectSwitchControllerQosDot1PMap-Priority7"); ok {
			if err = d.Set("priority_7", vv); err != nil {
				return fmt.Errorf("Error reading priority_7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_7: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerQosDot1PMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerQosDot1PMapDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapEgressPriTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosDot1PMapPriority7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerQosDot1PMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("egress_pri_tagging"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapEgressPriTagging(d, v, "egress_pri_tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-pri-tagging"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priority_0"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority0(d, v, "priority_0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-0"] = t
		}
	}

	if v, ok := d.GetOk("priority_1"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority1(d, v, "priority_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-1"] = t
		}
	}

	if v, ok := d.GetOk("priority_2"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority2(d, v, "priority_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-2"] = t
		}
	}

	if v, ok := d.GetOk("priority_3"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority3(d, v, "priority_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-3"] = t
		}
	}

	if v, ok := d.GetOk("priority_4"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority4(d, v, "priority_4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-4"] = t
		}
	}

	if v, ok := d.GetOk("priority_5"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority5(d, v, "priority_5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-5"] = t
		}
	}

	if v, ok := d.GetOk("priority_6"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority6(d, v, "priority_6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-6"] = t
		}
	}

	if v, ok := d.GetOk("priority_7"); ok {
		t, err := expandObjectSwitchControllerQosDot1PMapPriority7(d, v, "priority_7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-7"] = t
		}
	}

	return &obj, nil
}
