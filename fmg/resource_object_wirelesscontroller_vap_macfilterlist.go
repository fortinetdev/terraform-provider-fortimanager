// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create a list of MAC addresses for MAC address filtering.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerVapMacFilterList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerVapMacFilterListCreate,
		Read:   resourceObjectWirelessControllerVapMacFilterListRead,
		Update: resourceObjectWirelessControllerVapMacFilterListUpdate,
		Delete: resourceObjectWirelessControllerVapMacFilterListDelete,

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
			"vap": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerVapMacFilterListCreate(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	obj, err := getObjectObjectWirelessControllerVapMacFilterList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapMacFilterList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerVapMacFilterList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapMacFilterList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerVapMacFilterListRead(d, m)
}

func resourceObjectWirelessControllerVapMacFilterListUpdate(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	obj, err := getObjectObjectWirelessControllerVapMacFilterList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapMacFilterList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerVapMacFilterList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapMacFilterList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerVapMacFilterListRead(d, m)
}

func resourceObjectWirelessControllerVapMacFilterListDelete(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	paradict["vap"] = vap

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerVapMacFilterList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerVapMacFilterList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerVapMacFilterListRead(d *schema.ResourceData, m interface{}) error {
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

	vap := d.Get("vap").(string)
	if vap == "" {
		vap = importOptionChecking(m.(*FortiClient).Cfg, "vap")
		if vap == "" {
			return fmt.Errorf("Parameter vap is missing")
		}
		if err = d.Set("vap", vap); err != nil {
			return fmt.Errorf("Error set params vap: %v", err)
		}
	}
	paradict["vap"] = vap

	o, err := c.ReadObjectWirelessControllerVapMacFilterList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapMacFilterList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerVapMacFilterList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapMacFilterList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerVapMacFilterListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilterListMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapMacFilterListMacFilterPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerVapMacFilterList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerVapMacFilterListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerVapMacFilterList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectWirelessControllerVapMacFilterListMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectWirelessControllerVapMacFilterList-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("mac_filter_policy", flattenObjectWirelessControllerVapMacFilterListMacFilterPolicy2edl(o["mac-filter-policy"], d, "mac_filter_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-filter-policy"], "ObjectWirelessControllerVapMacFilterList-MacFilterPolicy"); ok {
			if err = d.Set("mac_filter_policy", vv); err != nil {
				return fmt.Errorf("Error reading mac_filter_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_filter_policy: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerVapMacFilterListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerVapMacFilterListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilterListMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapMacFilterListMacFilterPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerVapMacFilterList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerVapMacFilterListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectWirelessControllerVapMacFilterListMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_policy"); ok || d.HasChange("mac_filter_policy") {
		t, err := expandObjectWirelessControllerVapMacFilterListMacFilterPolicy2edl(d, v, "mac_filter_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-policy"] = t
		}
	}

	return &obj, nil
}
