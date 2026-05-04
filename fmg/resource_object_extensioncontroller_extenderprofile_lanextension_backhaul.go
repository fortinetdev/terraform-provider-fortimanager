// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: LAN extension backhaul tunnel configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaul() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulCreate,
		Read:   resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulRead,
		Update: resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulDelete,

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"health_check_fail_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"health_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"health_check_probe_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"health_check_probe_tm": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"health_check_recovery_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulCreate(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectExtensionControllerExtenderProfileLanExtensionBackhaul(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectExtensionControllerExtenderProfileLanExtensionBackhaul(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectExtensionControllerExtenderProfileLanExtensionBackhaul(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulUpdate(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtensionControllerExtenderProfileLanExtensionBackhaul(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulDelete(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectExtensionControllerExtenderProfileLanExtensionBackhaul(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileLanExtensionBackhaulRead(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadObjectExtensionControllerExtenderProfileLanExtensionBackhaul(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileLanExtensionBackhaul resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckFailCnt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeCnt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeTm3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckRecoveryCnt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("health_check_fail_cnt", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckFailCnt3rdl(o["health-check-fail-cnt"], d, "health_check_fail_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-fail-cnt"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-HealthCheckFailCnt"); ok {
			if err = d.Set("health_check_fail_cnt", vv); err != nil {
				return fmt.Errorf("Error reading health_check_fail_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_fail_cnt: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckInterval3rdl(o["health-check-interval"], d, "health_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-interval"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-HealthCheckInterval"); ok {
			if err = d.Set("health_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading health_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	if err = d.Set("health_check_probe_cnt", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeCnt3rdl(o["health-check-probe-cnt"], d, "health_check_probe_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-probe-cnt"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-HealthCheckProbeCnt"); ok {
			if err = d.Set("health_check_probe_cnt", vv); err != nil {
				return fmt.Errorf("Error reading health_check_probe_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_probe_cnt: %v", err)
		}
	}

	if err = d.Set("health_check_probe_tm", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeTm3rdl(o["health-check-probe-tm"], d, "health_check_probe_tm")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-probe-tm"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-HealthCheckProbeTm"); ok {
			if err = d.Set("health_check_probe_tm", vv); err != nil {
				return fmt.Errorf("Error reading health_check_probe_tm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_probe_tm: %v", err)
		}
	}

	if err = d.Set("health_check_recovery_cnt", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckRecoveryCnt3rdl(o["health-check-recovery-cnt"], d, "health_check_recovery_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-recovery-cnt"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-HealthCheckRecoveryCnt"); ok {
			if err = d.Set("health_check_recovery_cnt", vv); err != nil {
				return fmt.Errorf("Error reading health_check_recovery_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_recovery_cnt: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("role", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectExtensionControllerExtenderProfileLanExtensionBackhaul-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckFailCnt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeCnt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeTm3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckRecoveryCnt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("health_check_fail_cnt"); ok || d.HasChange("health_check_fail_cnt") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckFailCnt3rdl(d, v, "health_check_fail_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-fail-cnt"] = t
		}
	}

	if v, ok := d.GetOk("health_check_interval"); ok || d.HasChange("health_check_interval") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckInterval3rdl(d, v, "health_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("health_check_probe_cnt"); ok || d.HasChange("health_check_probe_cnt") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeCnt3rdl(d, v, "health_check_probe_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-probe-cnt"] = t
		}
	}

	if v, ok := d.GetOk("health_check_probe_tm"); ok || d.HasChange("health_check_probe_tm") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckProbeTm3rdl(d, v, "health_check_probe_tm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-probe-tm"] = t
		}
	}

	if v, ok := d.GetOk("health_check_recovery_cnt"); ok || d.HasChange("health_check_recovery_cnt") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulHealthCheckRecoveryCnt3rdl(d, v, "health_check_recovery_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-recovery-cnt"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole3rdl(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
