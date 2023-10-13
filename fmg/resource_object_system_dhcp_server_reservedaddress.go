// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Options for the DHCP server to assign IP settings to specific MAC addresses.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemDhcpServerReservedAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemDhcpServerReservedAddressCreate,
		Read:   resourceObjectSystemDhcpServerReservedAddressRead,
		Update: resourceObjectSystemDhcpServerReservedAddressUpdate,
		Delete: resourceObjectSystemDhcpServerReservedAddressDelete,

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
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"circuit_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemDhcpServerReservedAddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	server := d.Get("server").(string)
	paradict["server"] = server

	obj, err := getObjectObjectSystemDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemDhcpServerReservedAddress resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemDhcpServerReservedAddress(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemDhcpServerReservedAddressRead(d, m)
}

func resourceObjectSystemDhcpServerReservedAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	server := d.Get("server").(string)
	paradict["server"] = server

	obj, err := getObjectObjectSystemDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDhcpServerReservedAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemDhcpServerReservedAddress(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDhcpServerReservedAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemDhcpServerReservedAddressRead(d, m)
}

func resourceObjectSystemDhcpServerReservedAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	server := d.Get("server").(string)
	paradict["server"] = server

	err = c.DeleteObjectSystemDhcpServerReservedAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemDhcpServerReservedAddressRead(d *schema.ResourceData, m interface{}) error {
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

	server := d.Get("server").(string)
	if server == "" {
		server = importOptionChecking(m.(*FortiClient).Cfg, "server")
		if server == "" {
			return fmt.Errorf("Parameter server is missing")
		}
		if err = d.Set("server", server); err != nil {
			return fmt.Errorf("Error set params server: %v", err)
		}
	}
	paradict["server"] = server

	o, err := c.ReadObjectSystemDhcpServerReservedAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemDhcpServerReservedAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemDhcpServerReservedAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemDhcpServerReservedAddress resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemDhcpServerReservedAddressAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressCircuitId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressCircuitIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressRemoteId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressRemoteIdType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDhcpServerReservedAddressType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemDhcpServerReservedAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectSystemDhcpServerReservedAddressAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectSystemDhcpServerReservedAddress-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("circuit_id", flattenObjectSystemDhcpServerReservedAddressCircuitId2edl(o["circuit-id"], d, "circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id"], "ObjectSystemDhcpServerReservedAddress-CircuitId"); ok {
			if err = d.Set("circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id: %v", err)
		}
	}

	if err = d.Set("circuit_id_type", flattenObjectSystemDhcpServerReservedAddressCircuitIdType2edl(o["circuit-id-type"], d, "circuit_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id-type"], "ObjectSystemDhcpServerReservedAddress-CircuitIdType"); ok {
			if err = d.Set("circuit_id_type", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id_type: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSystemDhcpServerReservedAddressDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSystemDhcpServerReservedAddress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemDhcpServerReservedAddressId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemDhcpServerReservedAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectSystemDhcpServerReservedAddressIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectSystemDhcpServerReservedAddress-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectSystemDhcpServerReservedAddressMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectSystemDhcpServerReservedAddress-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("remote_id", flattenObjectSystemDhcpServerReservedAddressRemoteId2edl(o["remote-id"], d, "remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id"], "ObjectSystemDhcpServerReservedAddress-RemoteId"); ok {
			if err = d.Set("remote_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id: %v", err)
		}
	}

	if err = d.Set("remote_id_type", flattenObjectSystemDhcpServerReservedAddressRemoteIdType2edl(o["remote-id-type"], d, "remote_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id-type"], "ObjectSystemDhcpServerReservedAddress-RemoteIdType"); ok {
			if err = d.Set("remote_id_type", vv); err != nil {
				return fmt.Errorf("Error reading remote_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id_type: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemDhcpServerReservedAddressType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemDhcpServerReservedAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemDhcpServerReservedAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemDhcpServerReservedAddressAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressCircuitId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressCircuitIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressRemoteId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressRemoteIdType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDhcpServerReservedAddressType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemDhcpServerReservedAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectSystemDhcpServerReservedAddressAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id"); ok || d.HasChange("circuit_id") {
		t, err := expandObjectSystemDhcpServerReservedAddressCircuitId2edl(d, v, "circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id_type"); ok || d.HasChange("circuit_id_type") {
		t, err := expandObjectSystemDhcpServerReservedAddressCircuitIdType2edl(d, v, "circuit_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSystemDhcpServerReservedAddressDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemDhcpServerReservedAddressId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectSystemDhcpServerReservedAddressIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectSystemDhcpServerReservedAddressMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("remote_id"); ok || d.HasChange("remote_id") {
		t, err := expandObjectSystemDhcpServerReservedAddressRemoteId2edl(d, v, "remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_id_type"); ok || d.HasChange("remote_id_type") {
		t, err := expandObjectSystemDhcpServerReservedAddressRemoteIdType2edl(d, v, "remote_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemDhcpServerReservedAddressType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
