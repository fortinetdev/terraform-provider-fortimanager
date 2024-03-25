// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure port to NPU group mapping.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuPortNpuMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuPortNpuMapCreate,
		Read:   resourceObjectSystemNpuPortNpuMapRead,
		Update: resourceObjectSystemNpuPortNpuMapUpdate,
		Delete: resourceObjectSystemNpuPortNpuMapDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"npu_group_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuPortNpuMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuPortNpuMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuPortNpuMap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuPortNpuMap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuPortNpuMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "interface"))

	return resourceObjectSystemNpuPortNpuMapRead(d, m)
}

func resourceObjectSystemNpuPortNpuMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuPortNpuMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPortNpuMap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuPortNpuMap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuPortNpuMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "interface"))

	return resourceObjectSystemNpuPortNpuMapRead(d, m)
}

func resourceObjectSystemNpuPortNpuMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuPortNpuMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuPortNpuMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuPortNpuMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuPortNpuMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPortNpuMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuPortNpuMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuPortNpuMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuPortNpuMapInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuPortNpuMapNpuGroupIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuPortNpuMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("interface", flattenObjectSystemNpuPortNpuMapInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectSystemNpuPortNpuMap-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("npu_group_index", flattenObjectSystemNpuPortNpuMapNpuGroupIndex2edl(o["npu-group-index"], d, "npu_group_index")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-group-index"], "ObjectSystemNpuPortNpuMap-NpuGroupIndex"); ok {
			if err = d.Set("npu_group_index", vv); err != nil {
				return fmt.Errorf("Error reading npu_group_index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_group_index: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuPortNpuMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuPortNpuMapInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuPortNpuMapNpuGroupIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuPortNpuMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectSystemNpuPortNpuMapInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("npu_group_index"); ok || d.HasChange("npu_group_index") {
		t, err := expandObjectSystemNpuPortNpuMapNpuGroupIndex2edl(d, v, "npu_group_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-group-index"] = t
		}
	}

	return &obj, nil
}
