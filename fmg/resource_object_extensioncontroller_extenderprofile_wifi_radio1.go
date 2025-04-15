// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Radio-1 config for Wi-Fi 2.4GHz

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileWifiRadio1() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileWifiRadio1Update,
		Read:   resourceObjectExtensionControllerExtenderProfileWifiRadio1Read,
		Update: resourceObjectExtensionControllerExtenderProfileWifiRadio1Update,
		Delete: resourceObjectExtensionControllerExtenderProfileWifiRadio1Delete,

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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"n80211d": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"beacon_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bss_color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bss_color_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"channel": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"extension_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guard_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_ext_vap": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_vaps": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operating_standard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"power_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radio_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileWifiRadio1Update(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtensionControllerExtenderProfileWifiRadio1(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileWifiRadio1 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtensionControllerExtenderProfileWifiRadio1(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileWifiRadio1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileWifiRadio1")

	return resourceObjectExtensionControllerExtenderProfileWifiRadio1Read(d, m)
}

func resourceObjectExtensionControllerExtenderProfileWifiRadio1Delete(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectExtensionControllerExtenderProfileWifiRadio1(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileWifiRadio1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileWifiRadio1Read(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadObjectExtensionControllerExtenderProfileWifiRadio1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileWifiRadio1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileWifiRadio1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileWifiRadio1 resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio180211D3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Band3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Channel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Mode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Status3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileWifiRadio1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("n80211d", flattenObjectExtensionControllerExtenderProfileWifiRadio180211D3rdl(o["80211d"], d, "n80211d")); err != nil {
		if vv, ok := fortiAPIPatch(o["80211d"], "ObjectExtensionControllerExtenderProfileWifiRadio1-80211D"); ok {
			if err = d.Set("n80211d", vv); err != nil {
				return fmt.Errorf("Error reading n80211d: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n80211d: %v", err)
		}
	}

	if err = d.Set("band", flattenObjectExtensionControllerExtenderProfileWifiRadio1Band3rdl(o["band"], d, "band")); err != nil {
		if vv, ok := fortiAPIPatch(o["band"], "ObjectExtensionControllerExtenderProfileWifiRadio1-Band"); ok {
			if err = d.Set("band", vv); err != nil {
				return fmt.Errorf("Error reading band: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading band: %v", err)
		}
	}

	if err = d.Set("bandwidth", flattenObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(o["bandwidth"], d, "bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth"], "ObjectExtensionControllerExtenderProfileWifiRadio1-Bandwidth"); ok {
			if err = d.Set("bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(o["beacon-interval"], d, "beacon_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-interval"], "ObjectExtensionControllerExtenderProfileWifiRadio1-BeaconInterval"); ok {
			if err = d.Set("beacon_interval", vv); err != nil {
				return fmt.Errorf("Error reading beacon_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("bss_color", flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(o["bss-color"], d, "bss_color")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color"], "ObjectExtensionControllerExtenderProfileWifiRadio1-BssColor"); ok {
			if err = d.Set("bss_color", vv); err != nil {
				return fmt.Errorf("Error reading bss_color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color: %v", err)
		}
	}

	if err = d.Set("bss_color_mode", flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(o["bss-color-mode"], d, "bss_color_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["bss-color-mode"], "ObjectExtensionControllerExtenderProfileWifiRadio1-BssColorMode"); ok {
			if err = d.Set("bss_color_mode", vv); err != nil {
				return fmt.Errorf("Error reading bss_color_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bss_color_mode: %v", err)
		}
	}

	if err = d.Set("channel", flattenObjectExtensionControllerExtenderProfileWifiRadio1Channel3rdl(o["channel"], d, "channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["channel"], "ObjectExtensionControllerExtenderProfileWifiRadio1-Channel"); ok {
			if err = d.Set("channel", vv); err != nil {
				return fmt.Errorf("Error reading channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading channel: %v", err)
		}
	}

	if err = d.Set("extension_channel", flattenObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(o["extension-channel"], d, "extension_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-channel"], "ObjectExtensionControllerExtenderProfileWifiRadio1-ExtensionChannel"); ok {
			if err = d.Set("extension_channel", vv); err != nil {
				return fmt.Errorf("Error reading extension_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_channel: %v", err)
		}
	}

	if err = d.Set("guard_interval", flattenObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(o["guard-interval"], d, "guard_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["guard-interval"], "ObjectExtensionControllerExtenderProfileWifiRadio1-GuardInterval"); ok {
			if err = d.Set("guard_interval", vv); err != nil {
				return fmt.Errorf("Error reading guard_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guard_interval: %v", err)
		}
	}

	if err = d.Set("lan_ext_vap", flattenObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(o["lan-ext-vap"], d, "lan_ext_vap")); err != nil {
		if vv, ok := fortiAPIPatch(o["lan-ext-vap"], "ObjectExtensionControllerExtenderProfileWifiRadio1-LanExtVap"); ok {
			if err = d.Set("lan_ext_vap", vv); err != nil {
				return fmt.Errorf("Error reading lan_ext_vap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lan_ext_vap: %v", err)
		}
	}

	if err = d.Set("local_vaps", flattenObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(o["local-vaps"], d, "local_vaps")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-vaps"], "ObjectExtensionControllerExtenderProfileWifiRadio1-LocalVaps"); ok {
			if err = d.Set("local_vaps", vv); err != nil {
				return fmt.Errorf("Error reading local_vaps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_vaps: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenObjectExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(o["max-clients"], d, "max_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-clients"], "ObjectExtensionControllerExtenderProfileWifiRadio1-MaxClients"); ok {
			if err = d.Set("max_clients", vv); err != nil {
				return fmt.Errorf("Error reading max_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("mode", flattenObjectExtensionControllerExtenderProfileWifiRadio1Mode3rdl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "ObjectExtensionControllerExtenderProfileWifiRadio1-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("operating_standard", flattenObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(o["operating-standard"], d, "operating_standard")); err != nil {
		if vv, ok := fortiAPIPatch(o["operating-standard"], "ObjectExtensionControllerExtenderProfileWifiRadio1-OperatingStandard"); ok {
			if err = d.Set("operating_standard", vv); err != nil {
				return fmt.Errorf("Error reading operating_standard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading operating_standard: %v", err)
		}
	}

	if err = d.Set("power_level", flattenObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(o["power-level"], d, "power_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["power-level"], "ObjectExtensionControllerExtenderProfileWifiRadio1-PowerLevel"); ok {
			if err = d.Set("power_level", vv); err != nil {
				return fmt.Errorf("Error reading power_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading power_level: %v", err)
		}
	}

	if err = d.Set("radio_id", flattenObjectExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(o["radio-id"], d, "radio_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["radio-id"], "ObjectExtensionControllerExtenderProfileWifiRadio1-RadioId"); ok {
			if err = d.Set("radio_id", vv); err != nil {
				return fmt.Errorf("Error reading radio_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radio_id: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtensionControllerExtenderProfileWifiRadio1Status3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtensionControllerExtenderProfileWifiRadio1-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileWifiRadio180211D3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Band3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Channel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Mode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Status3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileWifiRadio1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n80211d"); ok || d.HasChange("n80211d") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio180211D3rdl(d, v, "n80211d")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211d"] = t
		}
	}

	if v, ok := d.GetOk("band"); ok || d.HasChange("band") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1Band3rdl(d, v, "band")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["band"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth"); ok || d.HasChange("bandwidth") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth3rdl(d, v, "bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok || d.HasChange("beacon_interval") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval3rdl(d, v, "beacon_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("bss_color"); ok || d.HasChange("bss_color") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1BssColor3rdl(d, v, "bss_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_mode"); ok || d.HasChange("bss_color_mode") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode3rdl(d, v, "bss_color_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-mode"] = t
		}
	}

	if v, ok := d.GetOk("channel"); ok || d.HasChange("channel") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1Channel3rdl(d, v, "channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["channel"] = t
		}
	}

	if v, ok := d.GetOk("extension_channel"); ok || d.HasChange("extension_channel") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel3rdl(d, v, "extension_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-channel"] = t
		}
	}

	if v, ok := d.GetOk("guard_interval"); ok || d.HasChange("guard_interval") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval3rdl(d, v, "guard_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guard-interval"] = t
		}
	}

	if v, ok := d.GetOk("lan_ext_vap"); ok || d.HasChange("lan_ext_vap") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap3rdl(d, v, "lan_ext_vap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-ext-vap"] = t
		}
	}

	if v, ok := d.GetOk("local_vaps"); ok || d.HasChange("local_vaps") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps3rdl(d, v, "local_vaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-vaps"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok || d.HasChange("max_clients") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1MaxClients3rdl(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1Mode3rdl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("operating_standard"); ok || d.HasChange("operating_standard") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard3rdl(d, v, "operating_standard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["operating-standard"] = t
		}
	}

	if v, ok := d.GetOk("power_level"); ok || d.HasChange("power_level") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel3rdl(d, v, "power_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["power-level"] = t
		}
	}

	if v, ok := d.GetOk("radio_id"); ok || d.HasChange("radio_id") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1RadioId3rdl(d, v, "radio_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio1Status3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
