// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Onetime schedule configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallScheduleOnetime() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallScheduleOnetimeCreate,
		Read:   resourceObjectFirewallScheduleOnetimeRead,
		Update: resourceObjectFirewallScheduleOnetimeUpdate,
		Delete: resourceObjectFirewallScheduleOnetimeDelete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expiration_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
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
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallScheduleOnetimeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallScheduleOnetime(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallScheduleOnetime resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallScheduleOnetime(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallScheduleOnetime resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallScheduleOnetimeRead(d, m)
}

func resourceObjectFirewallScheduleOnetimeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallScheduleOnetime(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallScheduleOnetime resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallScheduleOnetime(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallScheduleOnetime resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallScheduleOnetimeRead(d, m)
}

func resourceObjectFirewallScheduleOnetimeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallScheduleOnetime(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallScheduleOnetime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallScheduleOnetimeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallScheduleOnetime(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallScheduleOnetime resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallScheduleOnetime(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallScheduleOnetime resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallScheduleOnetimeColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleOnetimeEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleOnetimeExpirationDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleOnetimeFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleOnetimeGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleOnetimeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallScheduleOnetimeStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallScheduleOnetime(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("color", flattenObjectFirewallScheduleOnetimeColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallScheduleOnetime-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("end", flattenObjectFirewallScheduleOnetimeEnd(o["end"], d, "end")); err != nil {
		if vv, ok := fortiAPIPatch(o["end"], "ObjectFirewallScheduleOnetime-End"); ok {
			if err = d.Set("end", vv); err != nil {
				return fmt.Errorf("Error reading end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("expiration_days", flattenObjectFirewallScheduleOnetimeExpirationDays(o["expiration-days"], d, "expiration_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiration-days"], "ObjectFirewallScheduleOnetime-ExpirationDays"); ok {
			if err = d.Set("expiration_days", vv); err != nil {
				return fmt.Errorf("Error reading expiration_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiration_days: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectFirewallScheduleOnetimeFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectFirewallScheduleOnetime-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("global_object", flattenObjectFirewallScheduleOnetimeGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "ObjectFirewallScheduleOnetime-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallScheduleOnetimeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallScheduleOnetime-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenObjectFirewallScheduleOnetimeStart(o["start"], d, "start")); err != nil {
		if vv, ok := fortiAPIPatch(o["start"], "ObjectFirewallScheduleOnetime-Start"); ok {
			if err = d.Set("start", vv); err != nil {
				return fmt.Errorf("Error reading start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallScheduleOnetimeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallScheduleOnetimeColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleOnetimeEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleOnetimeExpirationDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleOnetimeFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleOnetimeGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleOnetimeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallScheduleOnetimeStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallScheduleOnetime(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok {
		t, err := expandObjectFirewallScheduleOnetimeColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok {
		t, err := expandObjectFirewallScheduleOnetimeEnd(d, v, "end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("expiration_days"); ok {
		t, err := expandObjectFirewallScheduleOnetimeExpirationDays(d, v, "expiration_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiration-days"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok {
		t, err := expandObjectFirewallScheduleOnetimeFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok {
		t, err := expandObjectFirewallScheduleOnetimeGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallScheduleOnetimeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok {
		t, err := expandObjectFirewallScheduleOnetimeStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	return &obj, nil
}
