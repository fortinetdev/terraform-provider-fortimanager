// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Mobile Country Code and Mobile Network Code configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListCreate,
		Read:   resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListRead,
		Update: resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListUpdate,
		Delete: resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListDelete,

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
			"anqp_3gpp_cellular": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"mcc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mnc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	obj, err := getObjectObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListUpdate(d *schema.ResourceData, m interface{}) error {
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

	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	obj, err := getObjectObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListDelete(d *schema.ResourceData, m interface{}) error {
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

	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	err = c.DeleteObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListRead(d *schema.ResourceData, m interface{}) error {
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

	anqp_3gpp_cellular := d.Get("anqp_3gpp_cellular").(string)
	if anqp_3gpp_cellular == "" {
		anqp_3gpp_cellular = importOptionChecking(m.(*FortiClient).Cfg, "anqp_3gpp_cellular")
		if anqp_3gpp_cellular == "" {
			return fmt.Errorf("Parameter anqp_3gpp_cellular is missing")
		}
		if err = d.Set("anqp_3gpp_cellular", anqp_3gpp_cellular); err != nil {
			return fmt.Errorf("Error set params anqp_3gpp_cellular: %v", err)
		}
	}
	paradict["anqp_3gpp_cellular"] = anqp_3gpp_cellular

	o, err := c.ReadObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mcc", flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(o["mcc"], d, "mcc")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcc"], "ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList-Mcc"); ok {
			if err = d.Set("mcc", vv); err != nil {
				return fmt.Errorf("Error reading mcc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcc: %v", err)
		}
	}

	if err = d.Set("mnc", flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(o["mnc"], d, "mnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["mnc"], "ObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList-Mnc"); ok {
			if err = d.Set("mnc", vv); err != nil {
				return fmt.Errorf("Error reading mnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mnc: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mcc"); ok || d.HasChange("mcc") {
		t, err := expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc2edl(d, v, "mcc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcc"] = t
		}
	}

	if v, ok := d.GetOk("mnc"); ok || d.HasChange("mnc") {
		t, err := expandObjectWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc2edl(d, v, "mnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mnc"] = t
		}
	}

	return &obj, nil
}
