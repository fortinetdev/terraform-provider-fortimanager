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

func resourceObjectFspVlanDynamicMappingDhcpServerReservedAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressCreate,
		Read:   resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressRead,
		Update: resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressUpdate,
		Delete: resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressDelete,

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
			"dynamic_mapping_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_vdom": &schema.Schema{
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

func resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressCreate(d *schema.ResourceData, m interface{}) error {
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
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectObjectFspVlanDynamicMappingDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanDynamicMappingDhcpServerReservedAddress(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressRead(d, m)
}

func resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressUpdate(d *schema.ResourceData, m interface{}) error {
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
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectObjectFspVlanDynamicMappingDhcpServerReservedAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanDynamicMappingDhcpServerReservedAddress(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressRead(d, m)
}

func resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressDelete(d *schema.ResourceData, m interface{}) error {
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
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	wsParams["adom"] = adomv

	err = c.DeleteObjectFspVlanDynamicMappingDhcpServerReservedAddress(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanDynamicMappingDhcpServerReservedAddressRead(d *schema.ResourceData, m interface{}) error {
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
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	if dynamic_mapping_name == "" {
		dynamic_mapping_name = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_name")
		if dynamic_mapping_name == "" {
			return fmt.Errorf("Parameter dynamic_mapping_name is missing")
		}
		if err = d.Set("dynamic_mapping_name", dynamic_mapping_name); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_name: %v", err)
		}
	}
	if dynamic_mapping_vdom == "" {
		dynamic_mapping_vdom = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_vdom")
		if dynamic_mapping_vdom == "" {
			return fmt.Errorf("Parameter dynamic_mapping_vdom is missing")
		}
		if err = d.Set("dynamic_mapping_vdom", dynamic_mapping_vdom); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_vdom: %v", err)
		}
	}
	paradict["vlan"] = vlan
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	o, err := c.ReadObjectFspVlanDynamicMappingDhcpServerReservedAddress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanDynamicMappingDhcpServerReservedAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMappingDhcpServerReservedAddress resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanDynamicMappingDhcpServerReservedAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressAction4thl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("circuit_id", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId4thl(o["circuit-id"], d, "circuit_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-CircuitId"); ok {
			if err = d.Set("circuit_id", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id: %v", err)
		}
	}

	if err = d.Set("circuit_id_type", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType4thl(o["circuit-id-type"], d, "circuit_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-id-type"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-CircuitIdType"); ok {
			if err = d.Set("circuit_id_type", vv); err != nil {
				return fmt.Errorf("Error reading circuit_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_id_type: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription4thl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressIp4thl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressMac4thl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("remote_id", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId4thl(o["remote-id"], d, "remote_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-RemoteId"); ok {
			if err = d.Set("remote_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id: %v", err)
		}
	}

	if err = d.Set("remote_id_type", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType4thl(o["remote-id-type"], d, "remote_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-id-type"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-RemoteIdType"); ok {
			if err = d.Set("remote_id_type", vv); err != nil {
				return fmt.Errorf("Error reading remote_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_id_type: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressType4thl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFspVlanDynamicMappingDhcpServerReservedAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpServerReservedAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanDynamicMappingDhcpServerReservedAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressAction4thl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id"); ok || d.HasChange("circuit_id") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitId4thl(d, v, "circuit_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id"] = t
		}
	}

	if v, ok := d.GetOk("circuit_id_type"); ok || d.HasChange("circuit_id_type") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressCircuitIdType4thl(d, v, "circuit_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-id-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressDescription4thl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressIp4thl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressMac4thl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("remote_id"); ok || d.HasChange("remote_id") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteId4thl(d, v, "remote_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_id_type"); ok || d.HasChange("remote_id_type") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressRemoteIdType4thl(d, v, "remote_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-id-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerReservedAddressType4thl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
