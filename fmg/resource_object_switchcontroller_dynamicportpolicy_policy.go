// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port policies with matching criteria and actions.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerDynamicPortPolicyPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerDynamicPortPolicyPolicyCreate,
		Read:   resourceObjectSwitchControllerDynamicPortPolicyPolicyRead,
		Update: resourceObjectSwitchControllerDynamicPortPolicyPolicyUpdate,
		Delete: resourceObjectSwitchControllerDynamicPortPolicyPolicyDelete,

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
			"dynamic_port_policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"n802_1x": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bounce_port_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"lldp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"qos_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	paradict["dynamic_port_policy"] = dynamic_port_policy

	obj, err := getObjectObjectSwitchControllerDynamicPortPolicyPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerDynamicPortPolicyPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerDynamicPortPolicyPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerDynamicPortPolicyPolicyRead(d, m)
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	paradict["dynamic_port_policy"] = dynamic_port_policy

	obj, err := getObjectObjectSwitchControllerDynamicPortPolicyPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDynamicPortPolicyPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerDynamicPortPolicyPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerDynamicPortPolicyPolicyRead(d, m)
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	paradict["dynamic_port_policy"] = dynamic_port_policy

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerDynamicPortPolicyPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	if dynamic_port_policy == "" {
		dynamic_port_policy = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_port_policy")
		if dynamic_port_policy == "" {
			return fmt.Errorf("Parameter dynamic_port_policy is missing")
		}
		if err = d.Set("dynamic_port_policy", dynamic_port_policy); err != nil {
			return fmt.Errorf("Error set params dynamic_port_policy: %v", err)
		}
	}
	paradict["dynamic_port_policy"] = dynamic_port_policy

	o, err := c.ReadObjectSwitchControllerDynamicPortPolicyPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicyPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerDynamicPortPolicyPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicyPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicy8021X2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyFamily2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("n802_1x", flattenObjectSwitchControllerDynamicPortPolicyPolicy8021X2edl(o["802-1x"], d, "n802_1x")); err != nil {
		if vv, ok := fortiAPIPatch(o["802-1x"], "ObjectSwitchControllerDynamicPortPolicyPolicy-8021X"); ok {
			if err = d.Set("n802_1x", vv); err != nil {
				return fmt.Errorf("Error reading n802_1x: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n802_1x: %v", err)
		}
	}

	if err = d.Set("bounce_port_link", flattenObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(o["bounce-port-link"], d, "bounce_port_link")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-port-link"], "ObjectSwitchControllerDynamicPortPolicyPolicy-BouncePortLink"); ok {
			if err = d.Set("bounce_port_link", vv); err != nil {
				return fmt.Errorf("Error reading bounce_port_link: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_port_link: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectSwitchControllerDynamicPortPolicyPolicyCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerDynamicPortPolicyPolicyDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("family", flattenObjectSwitchControllerDynamicPortPolicyPolicyFamily2edl(o["family"], d, "family")); err != nil {
		if vv, ok := fortiAPIPatch(o["family"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Family"); ok {
			if err = d.Set("family", vv); err != nil {
				return fmt.Errorf("Error reading family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading family: %v", err)
		}
	}

	if err = d.Set("host", flattenObjectSwitchControllerDynamicPortPolicyPolicyHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenObjectSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(o["hw-vendor"], d, "hw_vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw-vendor"], "ObjectSwitchControllerDynamicPortPolicyPolicy-HwVendor"); ok {
			if err = d.Set("hw_vendor", vv); err != nil {
				return fmt.Errorf("Error reading hw_vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("interface_tags", flattenObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(o["interface-tags"], d, "interface_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-tags"], "ObjectSwitchControllerDynamicPortPolicyPolicy-InterfaceTags"); ok {
			if err = d.Set("interface_tags", vv); err != nil {
				return fmt.Errorf("Error reading interface_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_tags: %v", err)
		}
	}

	if err = d.Set("lldp_profile", flattenObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(o["lldp-profile"], d, "lldp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["lldp-profile"], "ObjectSwitchControllerDynamicPortPolicyPolicy-LldpProfile"); ok {
			if err = d.Set("lldp_profile", vv); err != nil {
				return fmt.Errorf("Error reading lldp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lldp_profile: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectSwitchControllerDynamicPortPolicyPolicyMac2edl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerDynamicPortPolicyPolicyName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(o["qos-policy"], d, "qos_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-policy"], "ObjectSwitchControllerDynamicPortPolicyPolicy-QosPolicy"); ok {
			if err = d.Set("qos_policy", vv); err != nil {
				return fmt.Errorf("Error reading qos_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSwitchControllerDynamicPortPolicyPolicyStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSwitchControllerDynamicPortPolicyPolicyType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSwitchControllerDynamicPortPolicyPolicy-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vlan_policy", flattenObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(o["vlan-policy"], d, "vlan_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-policy"], "ObjectSwitchControllerDynamicPortPolicyPolicy-VlanPolicy"); ok {
			if err = d.Set("vlan_policy", vv); err != nil {
				return fmt.Errorf("Error reading vlan_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_policy: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerDynamicPortPolicyPolicy8021X2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyFamily2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n802_1x"); ok || d.HasChange("n802_1x") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicy8021X2edl(d, v, "n802_1x")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802-1x"] = t
		}
	}

	if v, ok := d.GetOk("bounce_port_link"); ok || d.HasChange("bounce_port_link") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink2edl(d, v, "bounce_port_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-port-link"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("family"); ok || d.HasChange("family") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyFamily2edl(d, v, "family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["family"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok || d.HasChange("hw_vendor") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyHwVendor2edl(d, v, "hw_vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("interface_tags"); ok || d.HasChange("interface_tags") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags2edl(d, v, "interface_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-tags"] = t
		}
	}

	if v, ok := d.GetOk("lldp_profile"); ok || d.HasChange("lldp_profile") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile2edl(d, v, "lldp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-profile"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyMac2edl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok || d.HasChange("qos_policy") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy2edl(d, v, "qos_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vlan_policy"); ok || d.HasChange("vlan_policy") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy2edl(d, v, "vlan_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-policy"] = t
		}
	}

	return &obj, nil
}
