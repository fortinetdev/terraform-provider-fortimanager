// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure port using NPU or Intel-NIC.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpuPortPathOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuPortPathOptionUpdate,
		Read:   resourceObjectSystemNpuPortPathOptionRead,
		Update: resourceObjectSystemNpuPortPathOptionUpdate,
		Delete: resourceObjectSystemNpuPortPathOptionDelete,

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
			"ports_using_npu": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuPortPathOptionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuPortPathOption(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPortPathOption resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuPortPathOption(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPortPathOption resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuPortPathOption")

	return resourceObjectSystemNpuPortPathOptionRead(d, m)
}

func resourceObjectSystemNpuPortPathOptionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuPortPathOption(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuPortPathOption resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuPortPathOptionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuPortPathOption(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPortPathOption resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuPortPathOption(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPortPathOption resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuPortPathOptionPortsUsingNpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectSystemNpuPortPathOption(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ports_using_npu", flattenObjectSystemNpuPortPathOptionPortsUsingNpu(o["ports-using-npu"], d, "ports_using_npu")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports-using-npu"], "ObjectSystemNpuPortPathOption-PortsUsingNpu"); ok {
			if err = d.Set("ports_using_npu", vv); err != nil {
				return fmt.Errorf("Error reading ports_using_npu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports_using_npu: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuPortPathOptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuPortPathOptionPortsUsingNpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectSystemNpuPortPathOption(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ports_using_npu"); ok {
		t, err := expandObjectSystemNpuPortPathOptionPortsUsingNpu(d, v, "ports_using_npu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports-using-npu"] = t
		}
	}

	return &obj, nil
}
