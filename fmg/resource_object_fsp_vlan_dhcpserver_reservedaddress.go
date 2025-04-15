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

func resourceObjectFspVlanDhcpServerReservedAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanDhcpServerReservedAddressCreate,
		Read:   resourceObjectFspVlanDhcpServerReservedAddressRead,
		Update: resourceObjectFspVlanDhcpServerReservedAddressUpdate,
		Delete: resourceObjectFspVlanDhcpServerReservedAddressDelete,

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

func resourceObjectFspVlanDhcpServerReservedAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFspVlanDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDhcpServerReservedAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanDhcpServerReservedAddress(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanDhcpServerReservedAddressRead(d, m)
}

func resourceObjectFspVlanDhcpServerReservedAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFspVlanDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDhcpServerReservedAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanDhcpServerReservedAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDhcpServerReservedAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanDhcpServerReservedAddressRead(d, m)
}

func resourceObjectFspVlanDhcpServerReservedAddressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFspVlanDhcpServerReservedAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanDhcpServerReservedAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFspVlanDhcpServerReservedAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDhcpServerReservedAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanDhcpServerReservedAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDhcpServerReservedAddress resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDhcpServerReservedAddressAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressCircuitId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressCircuitIdType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressDescription3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressMac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressRemoteId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressRemoteIdType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDhcpServerReservedAddressType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanDhcpServerReservedAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectFspVlanDhcpServerReservedAddressAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectFspVlanDhcpServerReservedAddress-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("circuit_id", flattenObjectFspVlanDhcpServerReservedAddressCircuitId3rdl(o["circuit-id"], d, "circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id"], "ObjectFspVlanDhcpServerReservedAddress-CircuitId"); ok {
			if err = d.Set("circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id: %v", err)
		}
	}

	if err = d.Set("circuit_id_type", flattenObjectFspVlanDhcpServerReservedAddressCircuitIdType3rdl(o["circuit-id-type"], d, "circuit_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id-type"], "ObjectFspVlanDhcpServerReservedAddress-CircuitIdType"); ok {
			if err = d.Set("circuit_id_type", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id_type: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectFspVlanDhcpServerReservedAddressDescription3rdl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectFspVlanDhcpServerReservedAddress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFspVlanDhcpServerReservedAddressId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanDhcpServerReservedAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFspVlanDhcpServerReservedAddressIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFspVlanDhcpServerReservedAddress-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectFspVlanDhcpServerReservedAddressMac3rdl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectFspVlanDhcpServerReservedAddress-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("remote_id", flattenObjectFspVlanDhcpServerReservedAddressRemoteId3rdl(o["remote-id"], d, "remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id"], "ObjectFspVlanDhcpServerReservedAddress-RemoteId"); ok {
			if err = d.Set("remote_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id: %v", err)
		}
	}

	if err = d.Set("remote_id_type", flattenObjectFspVlanDhcpServerReservedAddressRemoteIdType3rdl(o["remote-id-type"], d, "remote_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id-type"], "ObjectFspVlanDhcpServerReservedAddress-RemoteIdType"); ok {
			if err = d.Set("remote_id_type", vv); err != nil {
				return fmt.Errorf("Error reading remote_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id_type: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFspVlanDhcpServerReservedAddressType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFspVlanDhcpServerReservedAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanDhcpServerReservedAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDhcpServerReservedAddressAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressCircuitId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressCircuitIdType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressDescription3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressMac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressRemoteId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressRemoteIdType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDhcpServerReservedAddressType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanDhcpServerReservedAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id"); ok || d.HasChange("circuit_id") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressCircuitId3rdl(d, v, "circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id_type"); ok || d.HasChange("circuit_id_type") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressCircuitIdType3rdl(d, v, "circuit_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressDescription3rdl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressMac3rdl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("remote_id"); ok || d.HasChange("remote_id") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressRemoteId3rdl(d, v, "remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_id_type"); ok || d.HasChange("remote_id_type") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressRemoteIdType3rdl(d, v, "remote_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFspVlanDhcpServerReservedAddressType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
