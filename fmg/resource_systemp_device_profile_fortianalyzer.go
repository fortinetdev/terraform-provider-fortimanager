// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Systemp DeviceProfileFortianalyzer

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempDeviceProfileFortianalyzer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempDeviceProfileFortianalyzerUpdate,
		Read:   resourceSystempDeviceProfileFortianalyzerRead,
		Update: resourceSystempDeviceProfileFortianalyzerUpdate,
		Delete: resourceSystempDeviceProfileFortianalyzerDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"managed_sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_sn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystempDeviceProfileFortianalyzerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	obj, err := getObjectSystempDeviceProfileFortianalyzer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempDeviceProfileFortianalyzer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempDeviceProfileFortianalyzer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempDeviceProfileFortianalyzer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempDeviceProfileFortianalyzer")

	return resourceSystempDeviceProfileFortianalyzerRead(d, m)
}

func resourceSystempDeviceProfileFortianalyzerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	wsParams["adom"] = adomv

	err = c.DeleteSystempDeviceProfileFortianalyzer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempDeviceProfileFortianalyzer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempDeviceProfileFortianalyzerRead(d *schema.ResourceData, m interface{}) error {
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
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempDeviceProfileFortianalyzer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempDeviceProfileFortianalyzer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempDeviceProfileFortianalyzer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempDeviceProfileFortianalyzer resource from API: %v", err)
	}
	return nil
}

func flattenSystempDeviceProfileFortianalyzerManagedSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortianalyzerTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortianalyzerTargetIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortianalyzerTargetSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystempDeviceProfileFortianalyzer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("managed_sn", flattenSystempDeviceProfileFortianalyzerManagedSn(o["managed-sn"], d, "managed_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["managed-sn"], "SystempDeviceProfileFortianalyzer-ManagedSn"); ok {
			if err = d.Set("managed_sn", vv); err != nil {
				return fmt.Errorf("Error reading managed_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading managed_sn: %v", err)
		}
	}

	if err = d.Set("target", flattenSystempDeviceProfileFortianalyzerTarget(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "SystempDeviceProfileFortianalyzer-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	if err = d.Set("target_ip", flattenSystempDeviceProfileFortianalyzerTargetIp(o["target-ip"], d, "target_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-ip"], "SystempDeviceProfileFortianalyzer-TargetIp"); ok {
			if err = d.Set("target_ip", vv); err != nil {
				return fmt.Errorf("Error reading target_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_ip: %v", err)
		}
	}

	if err = d.Set("target_sn", flattenSystempDeviceProfileFortianalyzerTargetSn(o["target-sn"], d, "target_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-sn"], "SystempDeviceProfileFortianalyzer-TargetSn"); ok {
			if err = d.Set("target_sn", vv); err != nil {
				return fmt.Errorf("Error reading target_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_sn: %v", err)
		}
	}

	return nil
}

func flattenSystempDeviceProfileFortianalyzerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempDeviceProfileFortianalyzerManagedSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortianalyzerTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortianalyzerTargetIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortianalyzerTargetSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystempDeviceProfileFortianalyzer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("managed_sn"); ok || d.HasChange("managed_sn") {
		t, err := expandSystempDeviceProfileFortianalyzerManagedSn(d, v, "managed_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["managed-sn"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandSystempDeviceProfileFortianalyzerTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("target_ip"); ok || d.HasChange("target_ip") {
		t, err := expandSystempDeviceProfileFortianalyzerTargetIp(d, v, "target_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-ip"] = t
		}
	}

	if v, ok := d.GetOk("target_sn"); ok || d.HasChange("target_sn") {
		t, err := expandSystempDeviceProfileFortianalyzerTargetSn(d, v, "target_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-sn"] = t
		}
	}

	return &obj, nil
}
