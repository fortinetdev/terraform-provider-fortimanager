// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 of interface.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemInterfaceIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceIpv6Update,
		Read:   resourceSystemInterfaceIpv6Read,
		Update: resourceSystemInterfaceIpv6Update,
		Delete: resourceSystemInterfaceIpv6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip6_autoconf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemInterfaceIpv6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectSystemInterfaceIpv6(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterfaceIpv6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterfaceIpv6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemInterfaceIpv6")

	return resourceSystemInterfaceIpv6Read(d, m)
}

func resourceSystemInterfaceIpv6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	err = c.DeleteSystemInterfaceIpv6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterfaceIpv6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceIpv6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	var_interface := d.Get("interface").(string)
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["interface"] = var_interface

	o, err := c.ReadSystemInterfaceIpv6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterfaceIpv6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterfaceIpv6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceIpv6Ip6Address2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Allowaccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6Autoconf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterfaceIpv6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip6_address", flattenSystemInterfaceIpv6Ip6Address2edl(o["ip6-address"], d, "ip6_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-address"], "SystemInterfaceIpv6-Ip6Address"); ok {
			if err = d.Set("ip6_address", vv); err != nil {
				return fmt.Errorf("Error reading ip6_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("ip6_allowaccess", flattenSystemInterfaceIpv6Ip6Allowaccess2edl(o["ip6-allowaccess"], d, "ip6_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-allowaccess"], "SystemInterfaceIpv6-Ip6Allowaccess"); ok {
			if err = d.Set("ip6_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading ip6_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_allowaccess: %v", err)
		}
	}

	if err = d.Set("ip6_autoconf", flattenSystemInterfaceIpv6Ip6Autoconf2edl(o["ip6-autoconf"], d, "ip6_autoconf")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-autoconf"], "SystemInterfaceIpv6-Ip6Autoconf"); ok {
			if err = d.Set("ip6_autoconf", vv); err != nil {
				return fmt.Errorf("Error reading ip6_autoconf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_autoconf: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceIpv6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceIpv6Ip6Address2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Allowaccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6Autoconf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterfaceIpv6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip6_address"); ok || d.HasChange("ip6_address") {
		t, err := expandSystemInterfaceIpv6Ip6Address2edl(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6_allowaccess"); ok || d.HasChange("ip6_allowaccess") {
		t, err := expandSystemInterfaceIpv6Ip6Allowaccess2edl(d, v, "ip6_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("ip6_autoconf"); ok || d.HasChange("ip6_autoconf") {
		t, err := expandSystemInterfaceIpv6Ip6Autoconf2edl(d, v, "ip6_autoconf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-autoconf"] = t
		}
	}

	return &obj, nil
}
