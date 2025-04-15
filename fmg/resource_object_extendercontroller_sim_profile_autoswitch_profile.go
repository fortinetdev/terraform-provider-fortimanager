// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectExtenderController SimProfileAutoSwitchProfile

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerSimProfileAutoSwitchProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerSimProfileAutoSwitchProfileUpdate,
		Read:   resourceObjectExtenderControllerSimProfileAutoSwitchProfileRead,
		Update: resourceObjectExtenderControllerSimProfileAutoSwitchProfileUpdate,
		Delete: resourceObjectExtenderControllerSimProfileAutoSwitchProfileDelete,

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
			"sim_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dataplan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disconnect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disconnect_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disconnect_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_back": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_back_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_back_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtenderControllerSimProfileAutoSwitchProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	sim_profile := d.Get("sim_profile").(string)
	paradict["sim_profile"] = sim_profile

	obj, err := getObjectObjectExtenderControllerSimProfileAutoSwitchProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerSimProfileAutoSwitchProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtenderControllerSimProfileAutoSwitchProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerSimProfileAutoSwitchProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtenderControllerSimProfileAutoSwitchProfile")

	return resourceObjectExtenderControllerSimProfileAutoSwitchProfileRead(d, m)
}

func resourceObjectExtenderControllerSimProfileAutoSwitchProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	sim_profile := d.Get("sim_profile").(string)
	paradict["sim_profile"] = sim_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectExtenderControllerSimProfileAutoSwitchProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerSimProfileAutoSwitchProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerSimProfileAutoSwitchProfileRead(d *schema.ResourceData, m interface{}) error {
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

	sim_profile := d.Get("sim_profile").(string)
	if sim_profile == "" {
		sim_profile = importOptionChecking(m.(*FortiClient).Cfg, "sim_profile")
		if sim_profile == "" {
			return fmt.Errorf("Parameter sim_profile is missing")
		}
		if err = d.Set("sim_profile", sim_profile); err != nil {
			return fmt.Errorf("Error set params sim_profile: %v", err)
		}
	}
	paradict["sim_profile"] = sim_profile

	o, err := c.ReadObjectExtenderControllerSimProfileAutoSwitchProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerSimProfileAutoSwitchProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerSimProfileAutoSwitchProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerSimProfileAutoSwitchProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSignal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerSimProfileAutoSwitchProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dataplan", flattenObjectExtenderControllerSimProfileAutoSwitchProfileDataplan2edl(o["dataplan"], d, "dataplan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dataplan"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-Dataplan"); ok {
			if err = d.Set("dataplan", vv); err != nil {
				return fmt.Errorf("Error reading dataplan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dataplan: %v", err)
		}
	}

	if err = d.Set("disconnect", flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect2edl(o["disconnect"], d, "disconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-Disconnect"); ok {
			if err = d.Set("disconnect", vv); err != nil {
				return fmt.Errorf("Error reading disconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect: %v", err)
		}
	}

	if err = d.Set("disconnect_period", flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod2edl(o["disconnect-period"], d, "disconnect_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect-period"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-DisconnectPeriod"); ok {
			if err = d.Set("disconnect_period", vv); err != nil {
				return fmt.Errorf("Error reading disconnect_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect_period: %v", err)
		}
	}

	if err = d.Set("disconnect_threshold", flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold2edl(o["disconnect-threshold"], d, "disconnect_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["disconnect-threshold"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-DisconnectThreshold"); ok {
			if err = d.Set("disconnect_threshold", vv); err != nil {
				return fmt.Errorf("Error reading disconnect_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disconnect_threshold: %v", err)
		}
	}

	if err = d.Set("signal", flattenObjectExtenderControllerSimProfileAutoSwitchProfileSignal2edl(o["signal"], d, "signal")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-Signal"); ok {
			if err = d.Set("signal", vv); err != nil {
				return fmt.Errorf("Error reading signal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtenderControllerSimProfileAutoSwitchProfileStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("switch_back", flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack2edl(o["switch-back"], d, "switch_back")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-back"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-SwitchBack"); ok {
			if err = d.Set("switch_back", vv); err != nil {
				return fmt.Errorf("Error reading switch_back: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_back: %v", err)
		}
	}

	if err = d.Set("switch_back_time", flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime2edl(o["switch-back-time"], d, "switch_back_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-back-time"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-SwitchBackTime"); ok {
			if err = d.Set("switch_back_time", vv); err != nil {
				return fmt.Errorf("Error reading switch_back_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_back_time: %v", err)
		}
	}

	if err = d.Set("switch_back_timer", flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer2edl(o["switch-back-timer"], d, "switch_back_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-back-timer"], "ObjectExtenderControllerSimProfileAutoSwitchProfile-SwitchBackTimer"); ok {
			if err = d.Set("switch_back_timer", vv); err != nil {
				return fmt.Errorf("Error reading switch_back_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_back_timer: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSignal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerSimProfileAutoSwitchProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dataplan"); ok || d.HasChange("dataplan") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileDataplan2edl(d, v, "dataplan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dataplan"] = t
		}
	}

	if v, ok := d.GetOk("disconnect"); ok || d.HasChange("disconnect") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect2edl(d, v, "disconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_period"); ok || d.HasChange("disconnect_period") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod2edl(d, v, "disconnect_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-period"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_threshold"); ok || d.HasChange("disconnect_threshold") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold2edl(d, v, "disconnect_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-threshold"] = t
		}
	}

	if v, ok := d.GetOk("signal"); ok || d.HasChange("signal") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileSignal2edl(d, v, "signal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("switch_back"); ok || d.HasChange("switch_back") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack2edl(d, v, "switch_back")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-back"] = t
		}
	}

	if v, ok := d.GetOk("switch_back_time"); ok || d.HasChange("switch_back_time") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime2edl(d, v, "switch_back_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-back-time"] = t
		}
	}

	if v, ok := d.GetOk("switch_back_timer"); ok || d.HasChange("switch_back_timer") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer2edl(d, v, "switch_back_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-back-timer"] = t
		}
	}

	return &obj, nil
}
