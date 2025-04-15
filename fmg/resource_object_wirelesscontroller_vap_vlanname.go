// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Table for mapping VLAN name to VLAN ID.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerVapVlanName() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerVapVlanNameCreate,
		Read:   resourceObjectWirelessControllerVapVlanNameRead,
		Update: resourceObjectWirelessControllerVapVlanNameUpdate,
		Delete: resourceObjectWirelessControllerVapVlanNameDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerVapVlanNameCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerVapVlanName(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapVlanName resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerVapVlanName(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerVapVlanName resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerVapVlanNameRead(d, m)
}

func resourceObjectWirelessControllerVapVlanNameUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerVapVlanName(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapVlanName resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerVapVlanName(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerVapVlanName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerVapVlanNameRead(d, m)
}

func resourceObjectWirelessControllerVapVlanNameDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerVapVlanName(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerVapVlanName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerVapVlanNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerVapVlanName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapVlanName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerVapVlanName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerVapVlanName resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerVapVlanNameName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerVapVlanNameVlanId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerVapVlanName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectWirelessControllerVapVlanNameName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerVapVlanName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlan_id", flattenObjectWirelessControllerVapVlanNameVlanId2edl(o["vlan-id"], d, "vlan_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-id"], "ObjectWirelessControllerVapVlanName-VlanId"); ok {
			if err = d.Set("vlan_id", vv); err != nil {
				return fmt.Errorf("Error reading vlan_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_id: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerVapVlanNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerVapVlanNameName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerVapVlanNameVlanId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerVapVlanName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerVapVlanNameName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_id"); ok || d.HasChange("vlan_id") {
		t, err := expandObjectWirelessControllerVapVlanNameVlanId2edl(d, v, "vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-id"] = t
		}
	}

	return &obj, nil
}
