// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Mirror action of TCAM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcamMirAct() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamMirActUpdate,
		Read:   resourceObjectSystemNpuNpuTcamMirActRead,
		Update: resourceObjectSystemNpuNpuTcamMirActUpdate,
		Delete: resourceObjectSystemNpuNpuTcamMirActDelete,

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
			"vlif": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpuTcamMirActUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpuTcamMirAct(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamMirAct resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuNpuTcamMirAct(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamMirAct resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpuTcamMirAct")

	return resourceObjectSystemNpuNpuTcamMirActRead(d, m)
}

func resourceObjectSystemNpuNpuTcamMirActDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpuTcamMirAct(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcamMirAct resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamMirActRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpuTcamMirAct(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamMirAct resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcamMirAct(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamMirAct resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamMirActVlif3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcamMirAct(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("vlif", flattenObjectSystemNpuNpuTcamMirActVlif3rdl(o["vlif"], d, "vlif")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlif"], "ObjectSystemNpuNpuTcamMirAct-Vlif"); ok {
			if err = d.Set("vlif", vv); err != nil {
				return fmt.Errorf("Error reading vlif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlif: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamMirActFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamMirActVlif3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcamMirAct(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vlif"); ok || d.HasChange("vlif") {
		t, err := expandObjectSystemNpuNpuTcamMirActVlif3rdl(d, v, "vlif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlif"] = t
		}
	}

	return &obj, nil
}
