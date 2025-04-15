// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Priority action of TCAM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcamPriAct() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamPriActUpdate,
		Read:   resourceObjectSystemNpuNpuTcamPriActRead,
		Update: resourceObjectSystemNpuNpuTcamPriActUpdate,
		Delete: resourceObjectSystemNpuNpuTcamPriActDelete,

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
			"npu_tcam": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpuTcamPriActUpdate(d *schema.ResourceData, m interface{}) error {
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

	npu_tcam := d.Get("npu_tcam").(string)
	paradict["npu_tcam"] = npu_tcam

	obj, err := getObjectObjectSystemNpuNpuTcamPriAct(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamPriAct resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuNpuTcamPriAct(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamPriAct resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpuTcamPriAct")

	return resourceObjectSystemNpuNpuTcamPriActRead(d, m)
}

func resourceObjectSystemNpuNpuTcamPriActDelete(d *schema.ResourceData, m interface{}) error {
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

	npu_tcam := d.Get("npu_tcam").(string)
	paradict["npu_tcam"] = npu_tcam

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemNpuNpuTcamPriAct(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcamPriAct resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamPriActRead(d *schema.ResourceData, m interface{}) error {
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

	npu_tcam := d.Get("npu_tcam").(string)
	if npu_tcam == "" {
		npu_tcam = importOptionChecking(m.(*FortiClient).Cfg, "npu_tcam")
		if npu_tcam == "" {
			return fmt.Errorf("Parameter npu_tcam is missing")
		}
		if err = d.Set("npu_tcam", npu_tcam); err != nil {
			return fmt.Errorf("Error set params npu_tcam: %v", err)
		}
	}
	paradict["npu_tcam"] = npu_tcam

	o, err := c.ReadObjectSystemNpuNpuTcamPriAct(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamPriAct resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcamPriAct(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamPriAct resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamPriActPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamPriActWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcamPriAct(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("priority", flattenObjectSystemNpuNpuTcamPriActPriority3rdl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectSystemNpuNpuTcamPriAct-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectSystemNpuNpuTcamPriActWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectSystemNpuNpuTcamPriAct-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamPriActFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamPriActPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamPriActWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcamPriAct(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectSystemNpuNpuTcamPriActPriority3rdl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectSystemNpuNpuTcamPriActWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
