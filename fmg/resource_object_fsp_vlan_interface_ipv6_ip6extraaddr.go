// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Extra IPv6 address prefixes of interface.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddr() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrCreate,
		Read:   resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrRead,
		Update: resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrUpdate,
		Delete: resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrDelete,

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
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrCreate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanInterfaceIpv6Ip6ExtraAddr(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource: %v", err)
	}

	d.SetId(getStringKey(d, "prefix"))

	return resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrRead(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanInterfaceIpv6Ip6ExtraAddr(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "prefix"))

	return resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrRead(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	wsParams["adom"] = adomv

	err = c.DeleteObjectFspVlanInterfaceIpv6Ip6ExtraAddr(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceIpv6Ip6ExtraAddrRead(d *schema.ResourceData, m interface{}) error {
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
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterfaceIpv6Ip6ExtraAddr(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6Ip6ExtraAddr resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("prefix", flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix4thl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "ObjectFspVlanInterfaceIpv6Ip6ExtraAddr-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceIpv6Ip6ExtraAddrFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceIpv6Ip6ExtraAddr(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandObjectFspVlanInterfaceIpv6Ip6ExtraAddrPrefix4thl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
