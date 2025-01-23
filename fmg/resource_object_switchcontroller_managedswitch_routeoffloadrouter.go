// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure route offload MCLAG IP address.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerManagedSwitchRouteOffloadRouter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterCreate,
		Read:   resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterRead,
		Update: resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterUpdate,
		Delete: resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"router_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchRouteOffloadRouter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerManagedSwitchRouteOffloadRouter(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterUpdate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchRouteOffloadRouter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerManagedSwitchRouteOffloadRouter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterDelete(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	err = c.DeleteObjectSwitchControllerManagedSwitchRouteOffloadRouter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerManagedSwitchRouteOffloadRouterRead(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadObjectSwitchControllerManagedSwitchRouteOffloadRouter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerManagedSwitchRouteOffloadRouter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchRouteOffloadRouter resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("router_ip", flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(o["router-ip"], d, "router_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-ip"], "ObjectSwitchControllerManagedSwitchRouteOffloadRouter-RouterIp"); ok {
			if err = d.Set("router_ip", vv); err != nil {
				return fmt.Errorf("Error reading router_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_ip: %v", err)
		}
	}

	if err = d.Set("vlan_name", flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(o["vlan-name"], d, "vlan_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-name"], "ObjectSwitchControllerManagedSwitchRouteOffloadRouter-VlanName"); ok {
			if err = d.Set("vlan_name", vv); err != nil {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerManagedSwitchRouteOffloadRouterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("router_ip"); ok || d.HasChange("router_ip") {
		t, err := expandObjectSwitchControllerManagedSwitchRouteOffloadRouterRouterIp2edl(d, v, "router_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-ip"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandObjectSwitchControllerManagedSwitchRouteOffloadRouterVlanName2edl(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	return &obj, nil
}
