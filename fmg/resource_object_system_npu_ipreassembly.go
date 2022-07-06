// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IP reassebmly engine configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpuIpReassembly() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuIpReassemblyUpdate,
		Read:   resourceObjectSystemNpuIpReassemblyRead,
		Update: resourceObjectSystemNpuIpReassemblyUpdate,
		Delete: resourceObjectSystemNpuIpReassemblyDelete,

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
			"max_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuIpReassemblyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuIpReassembly(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuIpReassembly resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuIpReassembly(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuIpReassembly resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuIpReassembly")

	return resourceObjectSystemNpuIpReassemblyRead(d, m)
}

func resourceObjectSystemNpuIpReassemblyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuIpReassembly(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuIpReassembly resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuIpReassemblyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuIpReassembly(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuIpReassembly resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuIpReassembly(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuIpReassembly resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuIpReassemblyMaxTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpReassemblyMinTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIpReassemblyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuIpReassembly(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("max_timeout", flattenObjectSystemNpuIpReassemblyMaxTimeout(o["max-timeout"], d, "max_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-timeout"], "ObjectSystemNpuIpReassembly-MaxTimeout"); ok {
			if err = d.Set("max_timeout", vv); err != nil {
				return fmt.Errorf("Error reading max_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_timeout: %v", err)
		}
	}

	if err = d.Set("min_timeout", flattenObjectSystemNpuIpReassemblyMinTimeout(o["min-timeout"], d, "min_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["min-timeout"], "ObjectSystemNpuIpReassembly-MinTimeout"); ok {
			if err = d.Set("min_timeout", vv); err != nil {
				return fmt.Errorf("Error reading min_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min_timeout: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSystemNpuIpReassemblyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSystemNpuIpReassembly-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuIpReassemblyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuIpReassemblyMaxTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpReassemblyMinTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIpReassemblyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuIpReassembly(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_timeout"); ok || d.HasChange("max_timeout") {
		t, err := expandObjectSystemNpuIpReassemblyMaxTimeout(d, v, "max_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-timeout"] = t
		}
	}

	if v, ok := d.GetOk("min_timeout"); ok || d.HasChange("min_timeout") {
		t, err := expandObjectSystemNpuIpReassemblyMinTimeout(d, v, "min_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-timeout"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSystemNpuIpReassemblyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
