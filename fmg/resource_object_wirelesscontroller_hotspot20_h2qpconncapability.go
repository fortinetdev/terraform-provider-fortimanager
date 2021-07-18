// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure connection capability.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpConnCapability() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpConnCapabilityCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpConnCapabilityRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpConnCapabilityUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpConnCapabilityDelete,

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
			"esp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ikev2_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ikev2_xx_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"pptp_vpn_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_tcp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_udp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20H2QpConnCapabilityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpConnCapability(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpConnCapability resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpConnCapability(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpConnCapabilityRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpConnCapabilityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpConnCapability(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpConnCapability resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpConnCapability(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpConnCapabilityRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpConnCapabilityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20H2QpConnCapability(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpConnCapabilityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpConnCapability(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpConnCapability(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpConnCapability resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityEspPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityFtpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilitySshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityTlsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpConnCapability(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("esp_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityEspPort(o["esp-port"], d, "esp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["esp-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-EspPort"); ok {
			if err = d.Set("esp_port", vv); err != nil {
				return fmt.Errorf("Error reading esp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading esp_port: %v", err)
		}
	}

	if err = d.Set("ftp_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityFtpPort(o["ftp-port"], d, "ftp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftp-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-FtpPort"); ok {
			if err = d.Set("ftp_port", vv); err != nil {
				return fmt.Errorf("Error reading ftp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftp_port: %v", err)
		}
	}

	if err = d.Set("http_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityHttpPort(o["http-port"], d, "http_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-HttpPort"); ok {
			if err = d.Set("http_port", vv); err != nil {
				return fmt.Errorf("Error reading http_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_port: %v", err)
		}
	}

	if err = d.Set("icmp_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(o["icmp-port"], d, "icmp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-IcmpPort"); ok {
			if err = d.Set("icmp_port", vv); err != nil {
				return fmt.Errorf("Error reading icmp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_port: %v", err)
		}
	}

	if err = d.Set("ikev2_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(o["ikev2-port"], d, "ikev2_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ikev2-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-Ikev2Port"); ok {
			if err = d.Set("ikev2_port", vv); err != nil {
				return fmt.Errorf("Error reading ikev2_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ikev2_port: %v", err)
		}
	}

	if err = d.Set("ikev2_xx_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(o["ikev2-xx-port"], d, "ikev2_xx_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ikev2-xx-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-Ikev2XxPort"); ok {
			if err = d.Set("ikev2_xx_port", vv); err != nil {
				return fmt.Errorf("Error reading ikev2_xx_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ikev2_xx_port: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20H2QpConnCapability-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pptp_vpn_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(o["pptp-vpn-port"], d, "pptp_vpn_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["pptp-vpn-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-PptpVpnPort"); ok {
			if err = d.Set("pptp_vpn_port", vv); err != nil {
				return fmt.Errorf("Error reading pptp_vpn_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pptp_vpn_port: %v", err)
		}
	}

	if err = d.Set("ssh_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilitySshPort(o["ssh-port"], d, "ssh_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-SshPort"); ok {
			if err = d.Set("ssh_port", vv); err != nil {
				return fmt.Errorf("Error reading ssh_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_port: %v", err)
		}
	}

	if err = d.Set("tls_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityTlsPort(o["tls-port"], d, "tls_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-TlsPort"); ok {
			if err = d.Set("tls_port", vv); err != nil {
				return fmt.Errorf("Error reading tls_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_port: %v", err)
		}
	}

	if err = d.Set("voip_tcp_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(o["voip-tcp-port"], d, "voip_tcp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-tcp-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-VoipTcpPort"); ok {
			if err = d.Set("voip_tcp_port", vv); err != nil {
				return fmt.Errorf("Error reading voip_tcp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_tcp_port: %v", err)
		}
	}

	if err = d.Set("voip_udp_port", flattenObjectWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(o["voip-udp-port"], d, "voip_udp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-udp-port"], "ObjectWirelessControllerHotspot20H2QpConnCapability-VoipUdpPort"); ok {
			if err = d.Set("voip_udp_port", vv); err != nil {
				return fmt.Errorf("Error reading voip_udp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_udp_port: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpConnCapabilityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityEspPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityFtpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilitySshPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityTlsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpConnCapability(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("esp_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityEspPort(d, v, "esp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esp-port"] = t
		}
	}

	if v, ok := d.GetOk("ftp_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityFtpPort(d, v, "ftp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-port"] = t
		}
	}

	if v, ok := d.GetOk("http_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityHttpPort(d, v, "http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-port"] = t
		}
	}

	if v, ok := d.GetOk("icmp_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(d, v, "icmp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-port"] = t
		}
	}

	if v, ok := d.GetOk("ikev2_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(d, v, "ikev2_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ikev2-port"] = t
		}
	}

	if v, ok := d.GetOk("ikev2_xx_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(d, v, "ikev2_xx_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ikev2-xx-port"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pptp_vpn_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(d, v, "pptp_vpn_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-vpn-port"] = t
		}
	}

	if v, ok := d.GetOk("ssh_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilitySshPort(d, v, "ssh_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-port"] = t
		}
	}

	if v, ok := d.GetOk("tls_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityTlsPort(d, v, "tls_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-port"] = t
		}
	}

	if v, ok := d.GetOk("voip_tcp_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(d, v, "voip_tcp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-tcp-port"] = t
		}
	}

	if v, ok := d.GetOk("voip_udp_port"); ok {
		t, err := expandObjectWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(d, v, "voip_udp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-udp-port"] = t
		}
	}

	return &obj, nil
}
