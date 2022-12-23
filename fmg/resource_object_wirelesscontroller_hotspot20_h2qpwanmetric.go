// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WAN metrics.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpWanMetric() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpWanMetricCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpWanMetricRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpWanMetricUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpWanMetricDelete,

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
			"downlink_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"downlink_speed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_at_capacity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_measurement_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"symmetric_wan_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uplink_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uplink_speed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20H2QpWanMetricCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpWanMetric(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpWanMetric resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpWanMetric(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpWanMetricRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpWanMetricUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpWanMetric(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpWanMetric resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpWanMetric(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpWanMetricRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpWanMetricDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerHotspot20H2QpWanMetric(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpWanMetricRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpWanMetric(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpWanMetric(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpWanMetric resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricLinkStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricUplinkLoad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpWanMetric(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("downlink_load", flattenObjectWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(o["downlink-load"], d, "downlink_load")); err != nil {
		if vv, ok := fortiAPIPatch(o["downlink-load"], "ObjectWirelessControllerHotspot20H2QpWanMetric-DownlinkLoad"); ok {
			if err = d.Set("downlink_load", vv); err != nil {
				return fmt.Errorf("Error reading downlink_load: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downlink_load: %v", err)
		}
	}

	if err = d.Set("downlink_speed", flattenObjectWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(o["downlink-speed"], d, "downlink_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["downlink-speed"], "ObjectWirelessControllerHotspot20H2QpWanMetric-DownlinkSpeed"); ok {
			if err = d.Set("downlink_speed", vv); err != nil {
				return fmt.Errorf("Error reading downlink_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downlink_speed: %v", err)
		}
	}

	if err = d.Set("link_at_capacity", flattenObjectWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(o["link-at-capacity"], d, "link_at_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-at-capacity"], "ObjectWirelessControllerHotspot20H2QpWanMetric-LinkAtCapacity"); ok {
			if err = d.Set("link_at_capacity", vv); err != nil {
				return fmt.Errorf("Error reading link_at_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_at_capacity: %v", err)
		}
	}

	if err = d.Set("link_status", flattenObjectWirelessControllerHotspot20H2QpWanMetricLinkStatus(o["link-status"], d, "link_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-status"], "ObjectWirelessControllerHotspot20H2QpWanMetric-LinkStatus"); ok {
			if err = d.Set("link_status", vv); err != nil {
				return fmt.Errorf("Error reading link_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_status: %v", err)
		}
	}

	if err = d.Set("load_measurement_duration", flattenObjectWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(o["load-measurement-duration"], d, "load_measurement_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-measurement-duration"], "ObjectWirelessControllerHotspot20H2QpWanMetric-LoadMeasurementDuration"); ok {
			if err = d.Set("load_measurement_duration", vv); err != nil {
				return fmt.Errorf("Error reading load_measurement_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_measurement_duration: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20H2QpWanMetricName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20H2QpWanMetric-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("symmetric_wan_link", flattenObjectWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(o["symmetric-wan-link"], d, "symmetric_wan_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["symmetric-wan-link"], "ObjectWirelessControllerHotspot20H2QpWanMetric-SymmetricWanLink"); ok {
			if err = d.Set("symmetric_wan_link", vv); err != nil {
				return fmt.Errorf("Error reading symmetric_wan_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading symmetric_wan_link: %v", err)
		}
	}

	if err = d.Set("uplink_load", flattenObjectWirelessControllerHotspot20H2QpWanMetricUplinkLoad(o["uplink-load"], d, "uplink_load")); err != nil {
		if vv, ok := fortiAPIPatch(o["uplink-load"], "ObjectWirelessControllerHotspot20H2QpWanMetric-UplinkLoad"); ok {
			if err = d.Set("uplink_load", vv); err != nil {
				return fmt.Errorf("Error reading uplink_load: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uplink_load: %v", err)
		}
	}

	if err = d.Set("uplink_speed", flattenObjectWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(o["uplink-speed"], d, "uplink_speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["uplink-speed"], "ObjectWirelessControllerHotspot20H2QpWanMetric-UplinkSpeed"); ok {
			if err = d.Set("uplink_speed", vv); err != nil {
				return fmt.Errorf("Error reading uplink_speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uplink_speed: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpWanMetricFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricLinkStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricUplinkLoad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpWanMetric(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("downlink_load"); ok || d.HasChange("downlink_load") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(d, v, "downlink_load")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-load"] = t
		}
	}

	if v, ok := d.GetOk("downlink_speed"); ok || d.HasChange("downlink_speed") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(d, v, "downlink_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-speed"] = t
		}
	}

	if v, ok := d.GetOk("link_at_capacity"); ok || d.HasChange("link_at_capacity") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(d, v, "link_at_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-at-capacity"] = t
		}
	}

	if v, ok := d.GetOk("link_status"); ok || d.HasChange("link_status") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricLinkStatus(d, v, "link_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-status"] = t
		}
	}

	if v, ok := d.GetOk("load_measurement_duration"); ok || d.HasChange("load_measurement_duration") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(d, v, "load_measurement_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-measurement-duration"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("symmetric_wan_link"); ok || d.HasChange("symmetric_wan_link") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(d, v, "symmetric_wan_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["symmetric-wan-link"] = t
		}
	}

	if v, ok := d.GetOk("uplink_load"); ok || d.HasChange("uplink_load") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricUplinkLoad(d, v, "uplink_load")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-load"] = t
		}
	}

	if v, ok := d.GetOk("uplink_speed"); ok || d.HasChange("uplink_speed") {
		t, err := expandObjectWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(d, v, "uplink_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-speed"] = t
		}
	}

	return &obj, nil
}
