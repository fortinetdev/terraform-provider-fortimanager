// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch DHCP snooping static clients.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientCreate,
		Read:   resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientRead,
		Update: resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientUpdate,
		Delete: resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientCreate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientUpdate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientDelete(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ip", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("vlan", flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "ObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandObjectSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
