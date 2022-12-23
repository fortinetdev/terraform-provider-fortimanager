// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Interface sniffer.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSniffer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnifferCreate,
		Read:   resourceSystemSnifferRead,
		Update: resourceSystemSnifferUpdate,
		Delete: resourceSystemSnifferDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_packet_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"non_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
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

func resourceSystemSnifferCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSniffer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSniffer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSniffer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemSniffer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnifferRead(d, m)
}

func resourceSystemSnifferUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemSniffer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSniffer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSniffer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemSniffer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnifferRead(d, m)
}

func resourceSystemSnifferDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemSniffer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnifferRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemSniffer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSniffer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSniffer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSniffer resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnifferHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferMaxPacketCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferNonIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnifferVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSniffer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("host", flattenSystemSnifferHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "SystemSniffer-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemSnifferId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSniffer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemSnifferInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemSniffer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystemSnifferIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "SystemSniffer-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("max_packet_count", flattenSystemSnifferMaxPacketCount(o["max-packet-count"], d, "max_packet_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-packet-count"], "SystemSniffer-MaxPacketCount"); ok {
			if err = d.Set("max_packet_count", vv); err != nil {
				return fmt.Errorf("Error reading max_packet_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_packet_count: %v", err)
		}
	}

	if err = d.Set("non_ip", flattenSystemSnifferNonIp(o["non-ip"], d, "non_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["non-ip"], "SystemSniffer-NonIp"); ok {
			if err = d.Set("non_ip", vv); err != nil {
				return fmt.Errorf("Error reading non_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading non_ip: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemSnifferPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemSniffer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemSnifferProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemSniffer-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSystemSnifferVlan(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SystemSniffer-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSystemSnifferFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnifferHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferMaxPacketCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferNonIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSniffer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandSystemSnifferHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSnifferId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemSnifferInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandSystemSnifferIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("max_packet_count"); ok || d.HasChange("max_packet_count") {
		t, err := expandSystemSnifferMaxPacketCount(d, v, "max_packet_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-packet-count"] = t
		}
	}

	if v, ok := d.GetOk("non_ip"); ok || d.HasChange("non_ip") {
		t, err := expandSystemSnifferNonIp(d, v, "non_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["non-ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemSnifferPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemSnifferProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSystemSnifferVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
