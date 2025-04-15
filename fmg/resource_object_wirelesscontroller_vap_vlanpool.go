// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VLAN pool.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerVapVlanPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerVapVlanPoolCreate,
		Read:   resourceObjectWirelessControllerVapVlanPoolRead,
		Update: resourceObjectWirelessControllerVapVlanPoolUpdate,
		Delete: resourceObjectWirelessControllerVapVlanPoolDelete,

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
			"_wtp_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerVapVlanPoolCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerVapVlanPool(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapVlanPool resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerVapVlanPool(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapVlanPool resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerVapVlanPoolRead(d, m)
}

func resourceObjectWirelessControllerVapVlanPoolUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerVapVlanPool(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapVlanPool resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerVapVlanPool(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapVlanPool resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerVapVlanPoolRead(d, m)
}

func resourceObjectWirelessControllerVapVlanPoolDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerVapVlanPool(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerVapVlanPool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerVapVlanPoolRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerVapVlanPool(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapVlanPool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerVapVlanPool(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapVlanPool resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerVapVlanPoolWtpGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanPoolId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerVapVlanPool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("_wtp_group", flattenObjectWirelessControllerVapVlanPoolWtpGroup2edl(o["_wtp-group"], d, "_wtp_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["_wtp-group"], "ObjectWirelessControllerVapVlanPool-WtpGroup"); ok {
			if err = d.Set("_wtp_group", vv); err != nil {
				return fmt.Errorf("Error reading _wtp_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _wtp_group: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerVapVlanPoolId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerVapVlanPool-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerVapVlanPoolFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerVapVlanPoolWtpGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanPoolId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerVapVlanPool(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_wtp_group"); ok || d.HasChange("_wtp_group") {
		t, err := expandObjectWirelessControllerVapVlanPoolWtpGroup2edl(d, v, "_wtp_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_wtp-group"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerVapVlanPoolId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
