// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Table of IP ranges assigned to country.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemGeoipOverrideIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemGeoipOverrideIpRangeCreate,
		Read:   resourceObjectSystemGeoipOverrideIpRangeRead,
		Update: resourceObjectSystemGeoipOverrideIpRangeUpdate,
		Delete: resourceObjectSystemGeoipOverrideIpRangeDelete,

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

func resourceObjectSystemGeoipOverrideIpRangeCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["geoip_override"] = geoip_override

	obj, err := getObjectObjectSystemGeoipOverrideIpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipOverrideIpRange resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemGeoipOverrideIpRange(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipOverrideIpRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemGeoipOverrideIpRangeRead(d, m)
}

func resourceObjectSystemGeoipOverrideIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["geoip_override"] = geoip_override

	obj, err := getObjectObjectSystemGeoipOverrideIpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipOverrideIpRange resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemGeoipOverrideIpRange(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipOverrideIpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemGeoipOverrideIpRangeRead(d, m)
}

func resourceObjectSystemGeoipOverrideIpRangeDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["geoip_override"] = geoip_override

	err = c.DeleteObjectSystemGeoipOverrideIpRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemGeoipOverrideIpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemGeoipOverrideIpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemGeoipOverrideIpRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipOverrideIpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemGeoipOverrideIpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipOverrideIpRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemGeoipOverrideIpRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIpRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIpRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemGeoipOverrideIpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip", flattenObjectSystemGeoipOverrideIpRangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectSystemGeoipOverrideIpRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemGeoipOverrideIpRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemGeoipOverrideIpRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectSystemGeoipOverrideIpRangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectSystemGeoipOverrideIpRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemGeoipOverrideIpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemGeoipOverrideIpRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIpRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIpRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemGeoipOverrideIpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectSystemGeoipOverrideIpRangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemGeoipOverrideIpRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectSystemGeoipOverrideIpRangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
