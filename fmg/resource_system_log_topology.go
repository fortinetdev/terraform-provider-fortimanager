// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Logging topology settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogTopology() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogTopologyUpdate,
		Read:   resourceSystemLogTopologyRead,
		Update: resourceSystemLogTopologyUpdate,
		Delete: resourceSystemLogTopologyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_depth_share": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogTopologyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogTopology(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogTopology resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogTopology(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogTopology resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogTopology")

	return resourceSystemLogTopologyRead(d, m)
}

func resourceSystemLogTopologyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogTopology(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogTopology resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogTopologyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogTopology(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogTopology resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogTopology(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogTopology resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogTopologyMaxDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogTopologyMaxDepthShare(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogTopology(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_depth", flattenSystemLogTopologyMaxDepth(o["max-depth"], d, "max_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-depth"], "SystemLogTopology-MaxDepth"); ok {
			if err = d.Set("max_depth", vv); err != nil {
				return fmt.Errorf("Error reading max_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_depth: %v", err)
		}
	}

	if err = d.Set("max_depth_share", flattenSystemLogTopologyMaxDepthShare(o["max-depth-share"], d, "max_depth_share")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-depth-share"], "SystemLogTopology-MaxDepthShare"); ok {
			if err = d.Set("max_depth_share", vv); err != nil {
				return fmt.Errorf("Error reading max_depth_share: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_depth_share: %v", err)
		}
	}

	return nil
}

func flattenSystemLogTopologyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogTopologyMaxDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogTopologyMaxDepthShare(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogTopology(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_depth"); ok || d.HasChange("max_depth") {
		t, err := expandSystemLogTopologyMaxDepth(d, v, "max_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-depth"] = t
		}
	}

	if v, ok := d.GetOk("max_depth_share"); ok || d.HasChange("max_depth_share") {
		t, err := expandSystemLogTopologyMaxDepthShare(d, v, "max_depth_share")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-depth-share"] = t
		}
	}

	return &obj, nil
}
