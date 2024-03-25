// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NPU DoS configurations.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuDosOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuDosOptionsUpdate,
		Read:   resourceObjectSystemNpuDosOptionsRead,
		Update: resourceObjectSystemNpuDosOptionsUpdate,
		Delete: resourceObjectSystemNpuDosOptionsDelete,

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
			"npu_dos_meter_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"npu_dos_synproxy_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"npu_dos_tpe_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuDosOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuDosOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuDosOptions resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuDosOptions(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuDosOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuDosOptions")

	return resourceObjectSystemNpuDosOptionsRead(d, m)
}

func resourceObjectSystemNpuDosOptionsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuDosOptions(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuDosOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuDosOptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuDosOptions(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuDosOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuDosOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuDosOptions resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuDosOptionsNpuDosMeterMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsNpuDosSynproxyMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuDosOptionsNpuDosTpeMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuDosOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("npu_dos_meter_mode", flattenObjectSystemNpuDosOptionsNpuDosMeterMode2edl(o["npu-dos-meter-mode"], d, "npu_dos_meter_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-dos-meter-mode"], "ObjectSystemNpuDosOptions-NpuDosMeterMode"); ok {
			if err = d.Set("npu_dos_meter_mode", vv); err != nil {
				return fmt.Errorf("Error reading npu_dos_meter_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_dos_meter_mode: %v", err)
		}
	}

	if err = d.Set("npu_dos_synproxy_mode", flattenObjectSystemNpuDosOptionsNpuDosSynproxyMode2edl(o["npu-dos-synproxy-mode"], d, "npu_dos_synproxy_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-dos-synproxy-mode"], "ObjectSystemNpuDosOptions-NpuDosSynproxyMode"); ok {
			if err = d.Set("npu_dos_synproxy_mode", vv); err != nil {
				return fmt.Errorf("Error reading npu_dos_synproxy_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_dos_synproxy_mode: %v", err)
		}
	}

	if err = d.Set("npu_dos_tpe_mode", flattenObjectSystemNpuDosOptionsNpuDosTpeMode2edl(o["npu-dos-tpe-mode"], d, "npu_dos_tpe_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-dos-tpe-mode"], "ObjectSystemNpuDosOptions-NpuDosTpeMode"); ok {
			if err = d.Set("npu_dos_tpe_mode", vv); err != nil {
				return fmt.Errorf("Error reading npu_dos_tpe_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_dos_tpe_mode: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuDosOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuDosOptionsNpuDosMeterMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsNpuDosSynproxyMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuDosOptionsNpuDosTpeMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuDosOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("npu_dos_meter_mode"); ok || d.HasChange("npu_dos_meter_mode") {
		t, err := expandObjectSystemNpuDosOptionsNpuDosMeterMode2edl(d, v, "npu_dos_meter_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-dos-meter-mode"] = t
		}
	}

	if v, ok := d.GetOk("npu_dos_synproxy_mode"); ok || d.HasChange("npu_dos_synproxy_mode") {
		t, err := expandObjectSystemNpuDosOptionsNpuDosSynproxyMode2edl(d, v, "npu_dos_synproxy_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-dos-synproxy-mode"] = t
		}
	}

	if v, ok := d.GetOk("npu_dos_tpe_mode"); ok || d.HasChange("npu_dos_tpe_mode") {
		t, err := expandObjectSystemNpuDosOptionsNpuDosTpeMode2edl(d, v, "npu_dos_tpe_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-dos-tpe-mode"] = t
		}
	}

	return &obj, nil
}
