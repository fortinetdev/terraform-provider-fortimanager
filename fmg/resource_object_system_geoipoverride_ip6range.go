// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Table of IPv6 ranges assigned to country.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemGeoipOverrideIp6Range() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemGeoipOverrideIp6RangeCreate,
		Read:   resourceObjectSystemGeoipOverrideIp6RangeRead,
		Update: resourceObjectSystemGeoipOverrideIp6RangeUpdate,
		Delete: resourceObjectSystemGeoipOverrideIp6RangeDelete,

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
			"geoip_override": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemGeoipOverrideIp6RangeCreate(d *schema.ResourceData, m interface{}) error {
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

	geoip_override := d.Get("geoip_override").(string)
	paradict["geoip_override"] = geoip_override

	obj, err := getObjectObjectSystemGeoipOverrideIp6Range(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipOverrideIp6Range resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemGeoipOverrideIp6Range(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipOverrideIp6Range resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemGeoipOverrideIp6RangeRead(d, m)
}

func resourceObjectSystemGeoipOverrideIp6RangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	geoip_override := d.Get("geoip_override").(string)
	paradict["geoip_override"] = geoip_override

	obj, err := getObjectObjectSystemGeoipOverrideIp6Range(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipOverrideIp6Range resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemGeoipOverrideIp6Range(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipOverrideIp6Range resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemGeoipOverrideIp6RangeRead(d, m)
}

func resourceObjectSystemGeoipOverrideIp6RangeDelete(d *schema.ResourceData, m interface{}) error {
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

	geoip_override := d.Get("geoip_override").(string)
	paradict["geoip_override"] = geoip_override

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemGeoipOverrideIp6Range(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemGeoipOverrideIp6Range resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemGeoipOverrideIp6RangeRead(d *schema.ResourceData, m interface{}) error {
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

	geoip_override := d.Get("geoip_override").(string)
	if geoip_override == "" {
		geoip_override = importOptionChecking(m.(*FortiClient).Cfg, "geoip_override")
		if geoip_override == "" {
			return fmt.Errorf("Parameter geoip_override is missing")
		}
		if err = d.Set("geoip_override", geoip_override); err != nil {
			return fmt.Errorf("Error set params geoip_override: %v", err)
		}
	}
	paradict["geoip_override"] = geoip_override

	o, err := c.ReadObjectSystemGeoipOverrideIp6Range(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipOverrideIp6Range resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemGeoipOverrideIp6Range(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipOverrideIp6Range resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemGeoipOverrideIp6RangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIp6RangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIp6RangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemGeoipOverrideIp6Range(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip", flattenObjectSystemGeoipOverrideIp6RangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectSystemGeoipOverrideIp6Range-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemGeoipOverrideIp6RangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemGeoipOverrideIp6Range-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectSystemGeoipOverrideIp6RangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectSystemGeoipOverrideIp6Range-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemGeoipOverrideIp6RangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemGeoipOverrideIp6RangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIp6RangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIp6RangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemGeoipOverrideIp6Range(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectSystemGeoipOverrideIp6RangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemGeoipOverrideIp6RangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectSystemGeoipOverrideIp6RangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
