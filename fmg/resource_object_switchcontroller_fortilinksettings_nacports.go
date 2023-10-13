// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NAC specific configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerFortilinkSettingsNacPorts() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerFortilinkSettingsNacPortsUpdate,
		Read:   resourceObjectSwitchControllerFortilinkSettingsNacPortsRead,
		Update: resourceObjectSwitchControllerFortilinkSettingsNacPortsUpdate,
		Delete: resourceObjectSwitchControllerFortilinkSettingsNacPortsDelete,

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
			"fortilink_settings": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
	}
}

func resourceObjectSwitchControllerFortilinkSettingsNacPortsUpdate(d *schema.ResourceData, m interface{}) error {
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

	fortilink_settings := d.Get("fortilink_settings").(string)
	paradict["fortilink_settings"] = fortilink_settings

	obj, err := getObjectObjectSwitchControllerFortilinkSettingsNacPorts(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerFortilinkSettingsNacPorts resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerFortilinkSettingsNacPorts(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerFortilinkSettingsNacPorts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSwitchControllerFortilinkSettingsNacPorts")

	return resourceObjectSwitchControllerFortilinkSettingsNacPortsRead(d, m)
}

func resourceObjectSwitchControllerFortilinkSettingsNacPortsDelete(d *schema.ResourceData, m interface{}) error {
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

	fortilink_settings := d.Get("fortilink_settings").(string)
	paradict["fortilink_settings"] = fortilink_settings

	err = c.DeleteObjectSwitchControllerFortilinkSettingsNacPorts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerFortilinkSettingsNacPorts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerFortilinkSettingsNacPortsRead(d *schema.ResourceData, m interface{}) error {
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

	fortilink_settings := d.Get("fortilink_settings").(string)
	if fortilink_settings == "" {
		fortilink_settings = importOptionChecking(m.(*FortiClient).Cfg, "fortilink_settings")
		if fortilink_settings == "" {
			return fmt.Errorf("Parameter fortilink_settings is missing")
		}
		if err = d.Set("fortilink_settings", fortilink_settings); err != nil {
			return fmt.Errorf("Error set params fortilink_settings: %v", err)
		}
	}
	paradict["fortilink_settings"] = fortilink_settings

	o, err := c.ReadObjectSwitchControllerFortilinkSettingsNacPorts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerFortilinkSettingsNacPorts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerFortilinkSettingsNacPorts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerFortilinkSettingsNacPorts resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsParentKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("bounce_nac_port", flattenObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(o["bounce-nac-port"], d, "bounce_nac_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["bounce-nac-port"], "ObjectSwitchControllerFortilinkSettingsNacPorts-BounceNacPort"); ok {
			if err = d.Set("bounce_nac_port", vv); err != nil {
				return fmt.Errorf("Error reading bounce_nac_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bounce_nac_port: %v", err)
		}
	}

	if err = d.Set("lan_segment", flattenObjectSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(o["lan-segment"], d, "lan_segment")); err != nil {
		if vv, ok := fortiAPIPatch(o["lan-segment"], "ObjectSwitchControllerFortilinkSettingsNacPorts-LanSegment"); ok {
			if err = d.Set("lan_segment", vv); err != nil {
				return fmt.Errorf("Error reading lan_segment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lan_segment: %v", err)
		}
	}

	if err = d.Set("member_change", flattenObjectSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(o["member-change"], d, "member_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["member-change"], "ObjectSwitchControllerFortilinkSettingsNacPorts-MemberChange"); ok {
			if err = d.Set("member_change", vv); err != nil {
				return fmt.Errorf("Error reading member_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member_change: %v", err)
		}
	}

	if err = d.Set("nac_lan_interface", flattenObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(o["nac-lan-interface"], d, "nac_lan_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-lan-interface"], "ObjectSwitchControllerFortilinkSettingsNacPorts-NacLanInterface"); ok {
			if err = d.Set("nac_lan_interface", vv); err != nil {
				return fmt.Errorf("Error reading nac_lan_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_lan_interface: %v", err)
		}
	}

	if err = d.Set("nac_segment_vlans", flattenObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(o["nac-segment-vlans"], d, "nac_segment_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["nac-segment-vlans"], "ObjectSwitchControllerFortilinkSettingsNacPorts-NacSegmentVlans"); ok {
			if err = d.Set("nac_segment_vlans", vv); err != nil {
				return fmt.Errorf("Error reading nac_segment_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nac_segment_vlans: %v", err)
		}
	}

	if err = d.Set("onboarding_vlan", flattenObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(o["onboarding-vlan"], d, "onboarding_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["onboarding-vlan"], "ObjectSwitchControllerFortilinkSettingsNacPorts-OnboardingVlan"); ok {
			if err = d.Set("onboarding_vlan", vv); err != nil {
				return fmt.Errorf("Error reading onboarding_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading onboarding_vlan: %v", err)
		}
	}

	if err = d.Set("parent_key", flattenObjectSwitchControllerFortilinkSettingsNacPortsParentKey2edl(o["parent-key"], d, "parent_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["parent-key"], "ObjectSwitchControllerFortilinkSettingsNacPorts-ParentKey"); ok {
			if err = d.Set("parent_key", vv); err != nil {
				return fmt.Errorf("Error reading parent_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading parent_key: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerFortilinkSettingsNacPortsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerFortilinkSettingsNacPortsParentKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bounce_nac_port"); ok || d.HasChange("bounce_nac_port") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsBounceNacPort2edl(d, v, "bounce_nac_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-nac-port"] = t
		}
	}

	if v, ok := d.GetOk("lan_segment"); ok || d.HasChange("lan_segment") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsLanSegment2edl(d, v, "lan_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-segment"] = t
		}
	}

	if v, ok := d.GetOk("member_change"); ok || d.HasChange("member_change") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsMemberChange2edl(d, v, "member_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-change"] = t
		}
	}

	if v, ok := d.GetOk("nac_lan_interface"); ok || d.HasChange("nac_lan_interface") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsNacLanInterface2edl(d, v, "nac_lan_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-lan-interface"] = t
		}
	}

	if v, ok := d.GetOk("nac_segment_vlans"); ok || d.HasChange("nac_segment_vlans") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans2edl(d, v, "nac_segment_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-segment-vlans"] = t
		}
	}

	if v, ok := d.GetOk("onboarding_vlan"); ok || d.HasChange("onboarding_vlan") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsOnboardingVlan2edl(d, v, "onboarding_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onboarding-vlan"] = t
		}
	}

	if v, ok := d.GetOk("parent_key"); ok || d.HasChange("parent_key") {
		t, err := expandObjectSwitchControllerFortilinkSettingsNacPortsParentKey2edl(d, v, "parent_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parent-key"] = t
		}
	}

	return &obj, nil
}
