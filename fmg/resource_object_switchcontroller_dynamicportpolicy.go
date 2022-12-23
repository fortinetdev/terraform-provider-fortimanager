// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerDynamicPortPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerDynamicPortPolicyCreate,
		Read:   resourceObjectSwitchControllerDynamicPortPolicyRead,
		Update: resourceObjectSwitchControllerDynamicPortPolicyUpdate,
		Delete: resourceObjectSwitchControllerDynamicPortPolicyDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectSwitchControllerDynamicPortPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerDynamicPortPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerDynamicPortPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerDynamicPortPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerDynamicPortPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerDynamicPortPolicyRead(d, m)
}

func resourceObjectSwitchControllerDynamicPortPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerDynamicPortPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDynamicPortPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerDynamicPortPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDynamicPortPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerDynamicPortPolicyRead(d, m)
}

func resourceObjectSwitchControllerDynamicPortPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerDynamicPortPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerDynamicPortPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerDynamicPortPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerDynamicPortPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerDynamicPortPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerDynamicPortPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n802_1x"
		if _, ok := i["802-1x"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicy8021X(i["802-1x"], d, pre_append)
			tmp["n802_1x"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-8021X")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bounce_port_link"
		if _, ok := i["bounce-port-link"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink(i["bounce-port-link"], d, pre_append)
			tmp["bounce_port_link"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-BouncePortLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "family"
		if _, ok := i["family"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyFamily(i["family"], d, pre_append)
			tmp["family"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Family")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := i["hw-vendor"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyHwVendor(i["hw-vendor"], d, pre_append)
			tmp["hw_vendor"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-HwVendor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := i["interface-tags"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags(i["interface-tags"], d, pre_append)
			tmp["interface_tags"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-InterfaceTags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := i["lldp-profile"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile(i["lldp-profile"], d, pre_append)
			tmp["lldp_profile"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-LldpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyMac(i["mac"], d, pre_append)
			tmp["mac"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Mac")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := i["qos-policy"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy(i["qos-policy"], d, pre_append)
			tmp["qos_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-QosPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_policy"
		if _, ok := i["vlan-policy"]; ok {
			v := flattenObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy(i["vlan-policy"], d, pre_append)
			tmp["vlan_policy"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerDynamicPortPolicy-Policy-VlanPolicy")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicy8021X(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenObjectSwitchControllerDynamicPortPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerDynamicPortPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerDynamicPortPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerDynamicPortPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy", flattenObjectSwitchControllerDynamicPortPolicyPolicy(o["policy"], d, "policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["policy"], "ObjectSwitchControllerDynamicPortPolicy-Policy"); ok {
				if err = d.Set("policy", vv); err != nil {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy"); ok {
			if err = d.Set("policy", flattenObjectSwitchControllerDynamicPortPolicyPolicy(o["policy"], d, "policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["policy"], "ObjectSwitchControllerDynamicPortPolicy-Policy"); ok {
					if err = d.Set("policy", vv); err != nil {
						return fmt.Errorf("Error reading policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectSwitchControllerDynamicPortPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerDynamicPortPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n802_1x"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["802-1x"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicy8021X(d, i["n802_1x"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bounce_port_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bounce-port-link"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink(d, i["bounce_port_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "family"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["family"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyFamily(d, i["family"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hw-vendor"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyHwVendor(d, i["hw_vendor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-tags"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d, i["interface_tags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lldp-profile"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile(d, i["lldp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qos-policy"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy(d, i["qos_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-policy"], _ = expandObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy(d, i["vlan_policy"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicy8021X(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyBouncePortLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyLldpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyVlanPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policy"); ok || d.HasChange("policy") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicy(d, v, "policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	return &obj, nil
}
