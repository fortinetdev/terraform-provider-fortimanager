// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: URL list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListDelete,

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
			"anqp_venue_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	anqp_venue_url := d.Get("anqp_venue_url").(string)
	paradict["anqp_venue_url"] = anqp_venue_url

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpVenueUrlValueList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListUpdate(d *schema.ResourceData, m interface{}) error {
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

	anqp_venue_url := d.Get("anqp_venue_url").(string)
	paradict["anqp_venue_url"] = anqp_venue_url

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpVenueUrlValueList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListDelete(d *schema.ResourceData, m interface{}) error {
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

	anqp_venue_url := d.Get("anqp_venue_url").(string)
	paradict["anqp_venue_url"] = anqp_venue_url

	err = c.DeleteObjectWirelessControllerHotspot20AnqpVenueUrlValueList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlValueListRead(d *schema.ResourceData, m interface{}) error {
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

	anqp_venue_url := d.Get("anqp_venue_url").(string)
	if anqp_venue_url == "" {
		anqp_venue_url = importOptionChecking(m.(*FortiClient).Cfg, "anqp_venue_url")
		if anqp_venue_url == "" {
			return fmt.Errorf("Parameter anqp_venue_url is missing")
		}
		if err = d.Set("anqp_venue_url", anqp_venue_url); err != nil {
			return fmt.Errorf("Error set params anqp_venue_url: %v", err)
		}
	}
	paradict["anqp_venue_url"] = anqp_venue_url

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpVenueUrlValueList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpVenueUrlValueList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("index", flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectWirelessControllerHotspot20AnqpVenueUrlValueList-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("number", flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(o["number"], d, "number")); err != nil {
		if vv, ok := fortiAPIPatch(o["number"], "ObjectWirelessControllerHotspot20AnqpVenueUrlValueList-Number"); ok {
			if err = d.Set("number", vv); err != nil {
				return fmt.Errorf("Error reading number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading number: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectWirelessControllerHotspot20AnqpVenueUrlValueList-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("number"); ok || d.HasChange("number") {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber2edl(d, v, "number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
