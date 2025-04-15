// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Differentiated Services Code Point (DSCP) exceptions.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20QosMapDscpExcept() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20QosMapDscpExceptCreate,
		Read:   resourceObjectWirelessControllerHotspot20QosMapDscpExceptRead,
		Update: resourceObjectWirelessControllerHotspot20QosMapDscpExceptUpdate,
		Delete: resourceObjectWirelessControllerHotspot20QosMapDscpExceptDelete,

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
			"dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20QosMapDscpExceptCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerHotspot20QosMapDscpExcept(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20QosMapDscpExcept resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20QosMapDscpExcept(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20QosMapDscpExceptRead(d, m)
}

func resourceObjectWirelessControllerHotspot20QosMapDscpExceptUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerHotspot20QosMapDscpExcept(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20QosMapDscpExcept resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20QosMapDscpExcept(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20QosMapDscpExceptRead(d, m)
}

func resourceObjectWirelessControllerHotspot20QosMapDscpExceptDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerHotspot20QosMapDscpExcept(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20QosMapDscpExceptRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerHotspot20QosMapDscpExcept(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20QosMapDscpExcept resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20QosMapDscpExcept(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20QosMapDscpExcept resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptDscp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptUp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dscp", flattenObjectWirelessControllerHotspot20QosMapDscpExceptDscp2edl(o["dscp"], d, "dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp"], "ObjectWirelessControllerHotspot20QosMapDscpExcept-Dscp"); ok {
			if err = d.Set("dscp", vv); err != nil {
				return fmt.Errorf("Error reading dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp: %v", err)
		}
	}

	if err = d.Set("index", flattenObjectWirelessControllerHotspot20QosMapDscpExceptIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectWirelessControllerHotspot20QosMapDscpExcept-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("up", flattenObjectWirelessControllerHotspot20QosMapDscpExceptUp2edl(o["up"], d, "up")); err != nil {
		if vv, ok := fortiAPIPatch(o["up"], "ObjectWirelessControllerHotspot20QosMapDscpExcept-Up"); ok {
			if err = d.Set("up", vv); err != nil {
				return fmt.Errorf("Error reading up: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading up: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20QosMapDscpExceptDscp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpExceptIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpExceptUp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp"); ok || d.HasChange("dscp") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpExceptDscp2edl(d, v, "dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpExceptIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("up"); ok || d.HasChange("up") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpExceptUp2edl(d, v, "up")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["up"] = t
		}
	}

	return &obj, nil
}
