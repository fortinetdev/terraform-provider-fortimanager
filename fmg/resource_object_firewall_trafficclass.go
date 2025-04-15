// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure names for shaping classes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallTrafficClass() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallTrafficClassCreate,
		Read:   resourceObjectFirewallTrafficClassRead,
		Update: resourceObjectFirewallTrafficClassUpdate,
		Delete: resourceObjectFirewallTrafficClassDelete,

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
			"class_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"class_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallTrafficClassCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallTrafficClass(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallTrafficClass resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallTrafficClass(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallTrafficClass resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "class_id")))

	return resourceObjectFirewallTrafficClassRead(d, m)
}

func resourceObjectFirewallTrafficClassUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallTrafficClass(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallTrafficClass resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallTrafficClass(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallTrafficClass resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "class_id")))

	return resourceObjectFirewallTrafficClassRead(d, m)
}

func resourceObjectFirewallTrafficClassDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallTrafficClass(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallTrafficClass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallTrafficClassRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallTrafficClass(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallTrafficClass resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallTrafficClass(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallTrafficClass resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallTrafficClassClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallTrafficClassClassName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallTrafficClass(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("class_id", flattenObjectFirewallTrafficClassClassId(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "ObjectFirewallTrafficClass-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("class_name", flattenObjectFirewallTrafficClassClassName(o["class-name"], d, "class_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-name"], "ObjectFirewallTrafficClass-ClassName"); ok {
			if err = d.Set("class_name", vv); err != nil {
				return fmt.Errorf("Error reading class_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_name: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallTrafficClassFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallTrafficClassClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallTrafficClassClassName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallTrafficClass(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandObjectFirewallTrafficClassClassId(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("class_name"); ok || d.HasChange("class_name") {
		t, err := expandObjectFirewallTrafficClassClassName(d, v, "class_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-name"] = t
		}
	}

	return &obj, nil
}
