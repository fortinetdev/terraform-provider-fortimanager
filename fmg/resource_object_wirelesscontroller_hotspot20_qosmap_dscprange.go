// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Differentiated Services Code Point (DSCP) ranges.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20QosMapDscpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20QosMapDscpRangeCreate,
		Read:   resourceObjectWirelessControllerHotspot20QosMapDscpRangeRead,
		Update: resourceObjectWirelessControllerHotspot20QosMapDscpRangeUpdate,
		Delete: resourceObjectWirelessControllerHotspot20QosMapDscpRangeDelete,

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
			"qos_map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20QosMapDscpRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	qos_map := d.Get("qos_map").(string)
	paradict["qos_map"] = qos_map

	obj, err := getObjectObjectWirelessControllerHotspot20QosMapDscpRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20QosMapDscpRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20QosMapDscpRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20QosMapDscpRangeRead(d, m)
}

func resourceObjectWirelessControllerHotspot20QosMapDscpRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	qos_map := d.Get("qos_map").(string)
	paradict["qos_map"] = qos_map

	obj, err := getObjectObjectWirelessControllerHotspot20QosMapDscpRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20QosMapDscpRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20QosMapDscpRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20QosMapDscpRangeRead(d, m)
}

func resourceObjectWirelessControllerHotspot20QosMapDscpRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	qos_map := d.Get("qos_map").(string)
	paradict["qos_map"] = qos_map

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerHotspot20QosMapDscpRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20QosMapDscpRangeRead(d *schema.ResourceData, m interface{}) error {
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

	qos_map := d.Get("qos_map").(string)
	if qos_map == "" {
		qos_map = importOptionChecking(m.(*FortiClient).Cfg, "qos_map")
		if qos_map == "" {
			return fmt.Errorf("Parameter qos_map is missing")
		}
		if err = d.Set("qos_map", qos_map); err != nil {
			return fmt.Errorf("Error set params qos_map: %v", err)
		}
	}
	paradict["qos_map"] = qos_map

	o, err := c.ReadObjectWirelessControllerHotspot20QosMapDscpRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20QosMapDscpRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20QosMapDscpRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20QosMapDscpRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeHigh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeLow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeUp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("high", flattenObjectWirelessControllerHotspot20QosMapDscpRangeHigh2edl(o["high"], d, "high")); err != nil {
		if vv, ok := fortiAPIPatch(o["high"], "ObjectWirelessControllerHotspot20QosMapDscpRange-High"); ok {
			if err = d.Set("high", vv); err != nil {
				return fmt.Errorf("Error reading high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high: %v", err)
		}
	}

	if err = d.Set("index", flattenObjectWirelessControllerHotspot20QosMapDscpRangeIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectWirelessControllerHotspot20QosMapDscpRange-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("low", flattenObjectWirelessControllerHotspot20QosMapDscpRangeLow2edl(o["low"], d, "low")); err != nil {
		if vv, ok := fortiAPIPatch(o["low"], "ObjectWirelessControllerHotspot20QosMapDscpRange-Low"); ok {
			if err = d.Set("low", vv); err != nil {
				return fmt.Errorf("Error reading low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low: %v", err)
		}
	}

	if err = d.Set("up", flattenObjectWirelessControllerHotspot20QosMapDscpRangeUp2edl(o["up"], d, "up")); err != nil {
		if vv, ok := fortiAPIPatch(o["up"], "ObjectWirelessControllerHotspot20QosMapDscpRange-Up"); ok {
			if err = d.Set("up", vv); err != nil {
				return fmt.Errorf("Error reading up: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading up: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeHigh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeLow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeUp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("high"); ok || d.HasChange("high") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpRangeHigh2edl(d, v, "high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpRangeIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("low"); ok || d.HasChange("low") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpRangeLow2edl(d, v, "low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low"] = t
		}
	}

	if v, ok := d.GetOk("up"); ok || d.HasChange("up") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpRangeUp2edl(d, v, "up")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["up"] = t
		}
	}

	return &obj, nil
}
