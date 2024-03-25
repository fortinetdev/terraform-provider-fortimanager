// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU interface to CPU core mapping.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuPortCpuMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuPortCpuMapCreate,
		Read:   resourceObjectSystemNpuPortCpuMapRead,
		Update: resourceObjectSystemNpuPortCpuMapUpdate,
		Delete: resourceObjectSystemNpuPortCpuMapDelete,

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
			"cpu_core": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuPortCpuMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuPortCpuMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuPortCpuMap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuPortCpuMap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuPortCpuMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "interface"))

	return resourceObjectSystemNpuPortCpuMapRead(d, m)
}

func resourceObjectSystemNpuPortCpuMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuPortCpuMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPortCpuMap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuPortCpuMap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPortCpuMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "interface"))

	return resourceObjectSystemNpuPortCpuMapRead(d, m)
}

func resourceObjectSystemNpuPortCpuMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuPortCpuMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuPortCpuMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuPortCpuMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuPortCpuMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPortCpuMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuPortCpuMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPortCpuMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuPortCpuMapCpuCore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortCpuMapInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuPortCpuMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cpu_core", flattenObjectSystemNpuPortCpuMapCpuCore2edl(o["cpu-core"], d, "cpu_core")); err != nil {
		if vv, ok := fortiAPIPatch(o["cpu-core"], "ObjectSystemNpuPortCpuMap-CpuCore"); ok {
			if err = d.Set("cpu_core", vv); err != nil {
				return fmt.Errorf("Error reading cpu_core: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cpu_core: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectSystemNpuPortCpuMapInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectSystemNpuPortCpuMap-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuPortCpuMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuPortCpuMapCpuCore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortCpuMapInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuPortCpuMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cpu_core"); ok || d.HasChange("cpu_core") {
		t, err := expandObjectSystemNpuPortCpuMapCpuCore2edl(d, v, "cpu_core")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-core"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectSystemNpuPortCpuMapInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
