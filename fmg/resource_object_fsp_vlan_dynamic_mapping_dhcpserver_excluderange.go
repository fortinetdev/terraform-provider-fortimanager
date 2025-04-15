// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Exclude one or more ranges of IP addresses from being assigned to clients.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanDynamicMappingDhcpServerExcludeRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeCreate,
		Read:   resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeRead,
		Update: resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeUpdate,
		Delete: resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeDelete,

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
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFspVlanDynamicMappingDhcpServerExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanDynamicMappingDhcpServerExcludeRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeRead(d, m)
}

func resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFspVlanDynamicMappingDhcpServerExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanDynamicMappingDhcpServerExcludeRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeRead(d, m)
}

func resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFspVlanDynamicMappingDhcpServerExcludeRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanDynamicMappingDhcpServerExcludeRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFspVlanDynamicMappingDhcpServerExcludeRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanDynamicMappingDhcpServerExcludeRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanDynamicMappingDhcpServerExcludeRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFspVlanDynamicMappingDhcpServerExcludeRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp4thl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime4thl(o["lease-time"], d, "lease_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["lease-time"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-LeaseTime"); ok {
			if err = d.Set("lease_time", vv); err != nil {
				return fmt.Errorf("Error reading lease_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp4thl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("uci_match", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch4thl(o["uci-match"], d, "uci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-match"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-UciMatch"); ok {
			if err = d.Set("uci_match", vv); err != nil {
				return fmt.Errorf("Error reading uci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_match: %v", err)
		}
	}

	if err = d.Set("uci_string", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString4thl(o["uci-string"], d, "uci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["uci-string"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-UciString"); ok {
			if err = d.Set("uci_string", vv); err != nil {
				return fmt.Errorf("Error reading uci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uci_string: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch4thl(o["vci-match"], d, "vci_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-match"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-VciMatch"); ok {
			if err = d.Set("vci_match", vv); err != nil {
				return fmt.Errorf("Error reading vci_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString4thl(o["vci-string"], d, "vci_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["vci-string"], "ObjectFspVlanDynamicMappingDhcpServerExcludeRange-VciString"); ok {
			if err = d.Set("vci_string", vv); err != nil {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vci_string: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanDynamicMappingDhcpServerExcludeRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFspVlanDynamicMappingDhcpServerExcludeRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeEndIp4thl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok || d.HasChange("lease_time") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeLeaseTime4thl(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeStartIp4thl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("uci_match"); ok || d.HasChange("uci_match") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciMatch4thl(d, v, "uci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-match"] = t
		}
	}

	if v, ok := d.GetOk("uci_string"); ok || d.HasChange("uci_string") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeUciString4thl(d, v, "uci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uci-string"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok || d.HasChange("vci_match") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciMatch4thl(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandObjectFspVlanDynamicMappingDhcpServerExcludeRangeVciString4thl(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	return &obj, nil
}
