// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: VRRP Proxy ARP configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceVrrpProxyArp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceVrrpProxyArpCreate,
		Read:   resourceObjectFspVlanInterfaceVrrpProxyArpRead,
		Update: resourceObjectFspVlanInterfaceVrrpProxyArpUpdate,
		Delete: resourceObjectFspVlanInterfaceVrrpProxyArpDelete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vrrp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFspVlanInterfaceVrrpProxyArpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vlan := d.Get("vlan").(string)
	vrrp := d.Get("vrrp").(string)
	paradict["vlan"] = vlan
	paradict["vrrp"] = vrrp

	obj, err := getObjectObjectFspVlanInterfaceVrrpProxyArp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceVrrpProxyArp resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFspVlanInterfaceVrrpProxyArp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceVrrpProxyArp resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanInterfaceVrrpProxyArpRead(d, m)
}

func resourceObjectFspVlanInterfaceVrrpProxyArpUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	vrrp := d.Get("vrrp").(string)
	paradict["vlan"] = vlan
	paradict["vrrp"] = vrrp

	obj, err := getObjectObjectFspVlanInterfaceVrrpProxyArp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceVrrpProxyArp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlanInterfaceVrrpProxyArp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceVrrpProxyArp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanInterfaceVrrpProxyArpRead(d, m)
}

func resourceObjectFspVlanInterfaceVrrpProxyArpDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	vrrp := d.Get("vrrp").(string)
	paradict["vlan"] = vlan
	paradict["vrrp"] = vrrp

	err = c.DeleteObjectFspVlanInterfaceVrrpProxyArp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceVrrpProxyArp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceVrrpProxyArpRead(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	vrrp := d.Get("vrrp").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	if vrrp == "" {
		vrrp = importOptionChecking(m.(*FortiClient).Cfg, "vrrp")
		if vrrp == "" {
			return fmt.Errorf("Parameter vrrp is missing")
		}
		if err = d.Set("vrrp", vrrp); err != nil {
			return fmt.Errorf("Error set params vrrp: %v", err)
		}
	}
	paradict["vlan"] = vlan
	paradict["vrrp"] = vrrp

	o, err := c.ReadObjectFspVlanInterfaceVrrpProxyArp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceVrrpProxyArp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceVrrpProxyArp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceVrrpProxyArp resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceVrrpProxyArpId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceVrrpProxyArpIp4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceVrrpProxyArp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectFspVlanInterfaceVrrpProxyArpId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanInterfaceVrrpProxyArp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFspVlanInterfaceVrrpProxyArpIp4thl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFspVlanInterfaceVrrpProxyArp-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceVrrpProxyArpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceVrrpProxyArpId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceVrrpProxyArpIp4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceVrrpProxyArp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanInterfaceVrrpProxyArpId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFspVlanInterfaceVrrpProxyArpIp4thl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
