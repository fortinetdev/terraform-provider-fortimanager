// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure integrated FortiLink settings for FortiSwitch.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerFortilinkSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerFortilinkSettingsCreate,
		Read:   resourceObjectSwitchControllerFortilinkSettingsRead,
		Update: resourceObjectSwitchControllerFortilinkSettingsUpdate,
		Delete: resourceObjectSwitchControllerFortilinkSettingsDelete,

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
			"access_vlan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inactive_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_down_flush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bounce_nac_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lan_segment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"member_change": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"nac_lan_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nac_segment_vlans": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"onboarding_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"parent_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerFortilinkSettingsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerFortilinkSettings(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerFortilinkSettings resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerFortilinkSettings(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerFortilinkSettings resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerFortilinkSettingsRead(d, m)
}

func resourceObjectSwitchControllerFortilinkSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerFortilinkSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerFortilinkSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerFortilinkSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerFortilinkSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerFortilinkSettingsRead(d, m)
}

func resourceObjectSwitchControllerFortilinkSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerFortilinkSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerFortilinkSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerFortilinkSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerFortilinkSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerFortilinkSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerFortilinkSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerFortilinkSettings resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerFortilinkSettingsAccessVlanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsInactiveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsLinkDownFlush(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPorts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bounce_nac_port"
	if _, ok := i["bounce-nac-port"]; ok {
		result["bounce_nac_port"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort(i["bounce-nac-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "lan_segment"
	if _, ok := i["lan-segment"]; ok {
		result["lan_segment"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsLanSegment(i["lan-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "member_change"
	if _, ok := i["member-change"]; ok {
		result["member_change"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsMemberChange(i["member-change"], d, pre_append)
	}

	pre_append = pre + ".0." + "nac_lan_interface"
	if _, ok := i["nac-lan-interface"]; ok {
		result["nac_lan_interface"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface(i["nac-lan-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "nac_segment_vlans"
	if _, ok := i["nac-segment-vlans"]; ok {
		result["nac_segment_vlans"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(i["nac-segment-vlans"], d, pre_append)
	}

	pre_append = pre + ".0." + "onboarding_vlan"
	if _, ok := i["onboarding-vlan"]; ok {
		result["onboarding_vlan"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(i["onboarding-vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "parent_key"
	if _, ok := i["parent-key"]; ok {
		result["parent_key"] = flattenObjectSwitchControllerFortilinkSettingsNacPortsParentKey(i["parent-key"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsLanSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsMemberChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsParentKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerFortilinkSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("access_vlan_mode", flattenObjectSwitchControllerFortilinkSettingsAccessVlanMode(o["access-vlan-mode"], d, "access_vlan_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-vlan-mode"], "ObjectSwitchControllerFortilinkSettings-AccessVlanMode"); ok {
			if err = d.Set("access_vlan_mode", vv); err != nil {
				return fmt.Errorf("Error reading access_vlan_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_vlan_mode: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenObjectSwitchControllerFortilinkSettingsFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "ObjectSwitchControllerFortilinkSettings-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("inactive_timer", flattenObjectSwitchControllerFortilinkSettingsInactiveTimer(o["inactive-timer"], d, "inactive_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["inactive-timer"], "ObjectSwitchControllerFortilinkSettings-InactiveTimer"); ok {
			if err = d.Set("inactive_timer", vv); err != nil {
				return fmt.Errorf("Error reading inactive_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inactive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_flush", flattenObjectSwitchControllerFortilinkSettingsLinkDownFlush(o["link-down-flush"], d, "link_down_flush")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-down-flush"], "ObjectSwitchControllerFortilinkSettings-LinkDownFlush"); ok {
			if err = d.Set("link_down_flush", vv); err != nil {
				return fmt.Errorf("Error reading link_down_flush: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_down_flush: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nac_ports", flattenObjectSwitchControllerFortilinkSettingsNacPorts(o["nac-ports"], d, "nac_ports")); err != nil {
			if vv, ok := fortiAPIPatch(o["nac-ports"], "ObjectSwitchControllerFortilinkSettings-NacPorts"); ok {
				if err = d.Set("nac_ports", vv); err != nil {
					return fmt.Errorf("Error reading nac_ports: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nac_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_ports"); ok {
			if err = d.Set("nac_ports", flattenObjectSwitchControllerFortilinkSettingsNacPorts(o["nac-ports"], d, "nac_ports")); err != nil {
				if vv, ok := fortiAPIPatch(o["nac-ports"], "ObjectSwitchControllerFortilinkSettings-NacPorts"); ok {
					if err = d.Set("nac_ports", vv); err != nil {
						return fmt.Errorf("Error reading nac_ports: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nac_ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerFortilinkSettingsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerFortilinkSettings-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerFortilinkSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerFortilinkSettingsAccessVlanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsInactiveTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsLinkDownFlush(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bounce_nac_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bounce-nac-port"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort(d, i["bounce_nac_port"], pre_append)
	}
	pre_append = pre + ".0." + "lan_segment"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lan-segment"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsLanSegment(d, i["lan_segment"], pre_append)
	}
	pre_append = pre + ".0." + "member_change"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["member-change"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsMemberChange(d, i["member_change"], pre_append)
	}
	pre_append = pre + ".0." + "nac_lan_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nac-lan-interface"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface(d, i["nac_lan_interface"], pre_append)
	}
	pre_append = pre + ".0." + "nac_segment_vlans"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["nac-segment-vlans"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d, i["nac_segment_vlans"], pre_append)
	}
	pre_append = pre + ".0." + "onboarding_vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["onboarding-vlan"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(d, i["onboarding_vlan"], pre_append)
	}
	pre_append = pre + ".0." + "parent_key"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["parent-key"], _ = expandObjectSwitchControllerFortilinkSettingsNacPortsParentKey(d, i["parent_key"], pre_append)
	}

	return result, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsLanSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsMemberChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsParentKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerFortilinkSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_vlan_mode"); ok || d.HasChange("access_vlan_mode") {
		t, err := expandObjectSwitchControllerFortilinkSettingsAccessVlanMode(d, v, "access_vlan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-vlan-mode"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandObjectSwitchControllerFortilinkSettingsFortilink(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("inactive_timer"); ok || d.HasChange("inactive_timer") {
		t, err := expandObjectSwitchControllerFortilinkSettingsInactiveTimer(d, v, "inactive_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inactive-timer"] = t
		}
	}

	if v, ok := d.GetOk("link_down_flush"); ok || d.HasChange("link_down_flush") {
		t, err := expandObjectSwitchControllerFortilinkSettingsLinkDownFlush(d, v, "link_down_flush")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-flush"] = t
		}
	}

	if v, ok := d.GetOk("nac_ports"); ok || d.HasChange("nac_ports") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPorts(d, v, "nac_ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-ports"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerFortilinkSettingsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
