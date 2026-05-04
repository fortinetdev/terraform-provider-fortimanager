// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Systemp SystemTemplateInterfaceIpRange

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemTemplateInterfaceIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemTemplateInterfaceIpRangeCreate,
		Read:   resourceSystempSystemTemplateInterfaceIpRangeRead,
		Update: resourceSystempSystemTemplateInterfaceIpRangeUpdate,
		Delete: resourceSystempSystemTemplateInterfaceIpRangeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystempSystemTemplateInterfaceIpRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	var_interface := d.Get("interface").(string)
	paradict["devprof"] = devprof
	paradict["interface"] = var_interface

	obj, err := getObjectSystempSystemTemplateInterfaceIpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemTemplateInterfaceIpRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystempSystemTemplateInterfaceIpRange(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystempSystemTemplateInterfaceIpRange(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystempSystemTemplateInterfaceIpRange resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystempSystemTemplateInterfaceIpRange(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystempSystemTemplateInterfaceIpRange resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemTemplateInterfaceIpRangeRead(d, m)
}

func resourceSystempSystemTemplateInterfaceIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	var_interface := d.Get("interface").(string)
	paradict["devprof"] = devprof
	paradict["interface"] = var_interface

	obj, err := getObjectSystempSystemTemplateInterfaceIpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemTemplateInterfaceIpRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempSystemTemplateInterfaceIpRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemTemplateInterfaceIpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemTemplateInterfaceIpRangeRead(d, m)
}

func resourceSystempSystemTemplateInterfaceIpRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	var_interface := d.Get("interface").(string)
	paradict["devprof"] = devprof
	paradict["interface"] = var_interface

	wsParams["adom"] = adomv

	err = c.DeleteSystempSystemTemplateInterfaceIpRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemTemplateInterfaceIpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemTemplateInterfaceIpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	var_interface := d.Get("interface").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["devprof"] = devprof
	paradict["interface"] = var_interface

	o, err := c.ReadSystempSystemTemplateInterfaceIpRange(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystempSystemTemplateInterfaceIpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemTemplateInterfaceIpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemTemplateInterfaceIpRange resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemTemplateInterfaceIpRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemTemplateInterfaceIpRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemTemplateInterfaceIpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip", flattenSystempSystemTemplateInterfaceIpRangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "SystempSystemTemplateInterfaceIpRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystempSystemTemplateInterfaceIpRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystempSystemTemplateInterfaceIpRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenSystempSystemTemplateInterfaceIpRangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "SystempSystemTemplateInterfaceIpRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemTemplateInterfaceIpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemTemplateInterfaceIpRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemTemplateInterfaceIpRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemTemplateInterfaceIpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandSystempSystemTemplateInterfaceIpRangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystempSystemTemplateInterfaceIpRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandSystempSystemTemplateInterfaceIpRangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
