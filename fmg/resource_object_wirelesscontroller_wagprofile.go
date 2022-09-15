// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure wireless access gateway (WAG) profiles used for tunnels on AP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWagProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWagProfileCreate,
		Read:   resourceObjectWirelessControllerWagProfileRead,
		Update: resourceObjectWirelessControllerWagProfileUpdate,
		Delete: resourceObjectWirelessControllerWagProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ping_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ping_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"return_packet_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wag_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wag_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWagProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerWagProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWagProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerWagProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWagProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerWagProfileRead(d, m)
}

func resourceObjectWirelessControllerWagProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerWagProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWagProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerWagProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWagProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerWagProfileRead(d, m)
}

func resourceObjectWirelessControllerWagProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerWagProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWagProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWagProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerWagProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWagProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWagProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWagProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWagProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfileDhcpIpAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfilePingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfilePingNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfileReturnPacketTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfileTunnelType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfileWagIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWagProfileWagPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWagProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerWagProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerWagProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_addr", flattenObjectWirelessControllerWagProfileDhcpIpAddr(o["dhcp-ip-addr"], d, "dhcp_ip_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ip-addr"], "ObjectWirelessControllerWagProfile-DhcpIpAddr"); ok {
			if err = d.Set("dhcp_ip_addr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ip_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ip_addr: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerWagProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerWagProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ping_interval", flattenObjectWirelessControllerWagProfilePingInterval(o["ping-interval"], d, "ping_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ping-interval"], "ObjectWirelessControllerWagProfile-PingInterval"); ok {
			if err = d.Set("ping_interval", vv); err != nil {
				return fmt.Errorf("Error reading ping_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ping_interval: %v", err)
		}
	}

	if err = d.Set("ping_number", flattenObjectWirelessControllerWagProfilePingNumber(o["ping-number"], d, "ping_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ping-number"], "ObjectWirelessControllerWagProfile-PingNumber"); ok {
			if err = d.Set("ping_number", vv); err != nil {
				return fmt.Errorf("Error reading ping_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ping_number: %v", err)
		}
	}

	if err = d.Set("return_packet_timeout", flattenObjectWirelessControllerWagProfileReturnPacketTimeout(o["return-packet-timeout"], d, "return_packet_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["return-packet-timeout"], "ObjectWirelessControllerWagProfile-ReturnPacketTimeout"); ok {
			if err = d.Set("return_packet_timeout", vv); err != nil {
				return fmt.Errorf("Error reading return_packet_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading return_packet_timeout: %v", err)
		}
	}

	if err = d.Set("tunnel_type", flattenObjectWirelessControllerWagProfileTunnelType(o["tunnel-type"], d, "tunnel_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-type"], "ObjectWirelessControllerWagProfile-TunnelType"); ok {
			if err = d.Set("tunnel_type", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_type: %v", err)
		}
	}

	if err = d.Set("wag_ip", flattenObjectWirelessControllerWagProfileWagIp(o["wag-ip"], d, "wag_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["wag-ip"], "ObjectWirelessControllerWagProfile-WagIp"); ok {
			if err = d.Set("wag_ip", vv); err != nil {
				return fmt.Errorf("Error reading wag_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wag_ip: %v", err)
		}
	}

	if err = d.Set("wag_port", flattenObjectWirelessControllerWagProfileWagPort(o["wag-port"], d, "wag_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["wag-port"], "ObjectWirelessControllerWagProfile-WagPort"); ok {
			if err = d.Set("wag_port", vv); err != nil {
				return fmt.Errorf("Error reading wag_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wag_port: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWagProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWagProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfileDhcpIpAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfilePingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfilePingNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfileReturnPacketTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfileTunnelType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfileWagIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWagProfileWagPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWagProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerWagProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_addr"); ok || d.HasChange("dhcp_ip_addr") {
		t, err := expandObjectWirelessControllerWagProfileDhcpIpAddr(d, v, "dhcp_ip_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ip-addr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerWagProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ping_interval"); ok || d.HasChange("ping_interval") {
		t, err := expandObjectWirelessControllerWagProfilePingInterval(d, v, "ping_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-interval"] = t
		}
	}

	if v, ok := d.GetOk("ping_number"); ok || d.HasChange("ping_number") {
		t, err := expandObjectWirelessControllerWagProfilePingNumber(d, v, "ping_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-number"] = t
		}
	}

	if v, ok := d.GetOk("return_packet_timeout"); ok || d.HasChange("return_packet_timeout") {
		t, err := expandObjectWirelessControllerWagProfileReturnPacketTimeout(d, v, "return_packet_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["return-packet-timeout"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_type"); ok || d.HasChange("tunnel_type") {
		t, err := expandObjectWirelessControllerWagProfileTunnelType(d, v, "tunnel_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-type"] = t
		}
	}

	if v, ok := d.GetOk("wag_ip"); ok || d.HasChange("wag_ip") {
		t, err := expandObjectWirelessControllerWagProfileWagIp(d, v, "wag_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wag-ip"] = t
		}
	}

	if v, ok := d.GetOk("wag_port"); ok || d.HasChange("wag_port") {
		t, err := expandObjectWirelessControllerWagProfileWagPort(d, v, "wag_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wag-port"] = t
		}
	}

	return &obj, nil
}
