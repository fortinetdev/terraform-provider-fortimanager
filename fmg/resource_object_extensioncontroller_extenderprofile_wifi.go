// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender wifi configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileWifi() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileWifiUpdate,
		Read:   resourceObjectExtensionControllerExtenderProfileWifiRead,
		Update: resourceObjectExtensionControllerExtenderProfileWifiUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileWifiDelete,

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
			"dfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"radio_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileWifiUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtensionControllerExtenderProfileWifi(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileWifi resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerExtenderProfileWifi(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileWifi resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileWifi")

	return resourceObjectExtensionControllerExtenderProfileWifiRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileWifiDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["extender_profile"] = extender_profile

	err = c.DeleteObjectExtensionControllerExtenderProfileWifi(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileWifi resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileWifiRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileWifi(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileWifi resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileWifi(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileWifi resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileWifiDfs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiCountry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio12edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenObjectExtensionControllerExtenderProfileWifiRadio180211D2edl(i["80211d"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1Band2edl(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth"
	if _, ok := i["bandwidth"]; ok {
		result["bandwidth"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth2edl(i["bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval2edl(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColor2edl(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode2edl(i["bss-color-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1Channel2edl(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "extension_channel"
	if _, ok := i["extension-channel"]; ok {
		result["extension_channel"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel2edl(i["extension-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "guard_interval"
	if _, ok := i["guard-interval"]; ok {
		result["guard_interval"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval2edl(i["guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := i["lan-ext-vap"]; ok {
		result["lan_ext_vap"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap2edl(i["lan-ext-vap"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_vaps"
	if _, ok := i["local-vaps"]; ok {
		result["local_vaps"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps2edl(i["local-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1MaxClients2edl(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1Mode2edl(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "operating_standard"
	if _, ok := i["operating-standard"]; ok {
		result["operating_standard"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard2edl(i["operating-standard"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel2edl(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1RadioId2edl(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtensionControllerExtenderProfileWifiRadio1Status2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio180211D2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Band2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Channel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1MaxClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1RadioId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio1Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio22edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenObjectExtensionControllerExtenderProfileWifiRadio280211D2edl(i["80211d"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2Band2edl(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth"
	if _, ok := i["bandwidth"]; ok {
		result["bandwidth"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2Bandwidth2edl(i["bandwidth"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2BeaconInterval2edl(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2BssColor2edl(i["bss-color"], d, pre_append)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2BssColorMode2edl(i["bss-color-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2Channel2edl(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "extension_channel"
	if _, ok := i["extension-channel"]; ok {
		result["extension_channel"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2ExtensionChannel2edl(i["extension-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "guard_interval"
	if _, ok := i["guard-interval"]; ok {
		result["guard_interval"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2GuardInterval2edl(i["guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := i["lan-ext-vap"]; ok {
		result["lan_ext_vap"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2LanExtVap2edl(i["lan-ext-vap"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_vaps"
	if _, ok := i["local-vaps"]; ok {
		result["local_vaps"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2LocalVaps2edl(i["local-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2MaxClients2edl(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2Mode2edl(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "operating_standard"
	if _, ok := i["operating-standard"]; ok {
		result["operating_standard"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2OperatingStandard2edl(i["operating-standard"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2PowerLevel2edl(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2RadioId2edl(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtensionControllerExtenderProfileWifiRadio2Status2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio280211D2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2Band2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2Bandwidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2BeaconInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2BssColor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2BssColorMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2Channel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2ExtensionChannel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2GuardInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2LanExtVap2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2LocalVaps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2MaxClients2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2Mode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2OperatingStandard2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2PowerLevel2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2RadioId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileWifiRadio2Status2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileWifi(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dfs", flattenObjectExtensionControllerExtenderProfileWifiDfs2edl(o["DFS"], d, "dfs")); err != nil {
		if vv, ok := fortiAPIPatch(o["DFS"], "ObjectExtensionControllerExtenderProfileWifi-Dfs"); ok {
			if err = d.Set("dfs", vv); err != nil {
				return fmt.Errorf("Error reading dfs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dfs: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectExtensionControllerExtenderProfileWifiCountry2edl(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectExtensionControllerExtenderProfileWifi-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radio_1", flattenObjectExtensionControllerExtenderProfileWifiRadio12edl(o["radio-1"], d, "radio_1")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-1"], "ObjectExtensionControllerExtenderProfileWifi-Radio1"); ok {
				if err = d.Set("radio_1", vv); err != nil {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenObjectExtensionControllerExtenderProfileWifiRadio12edl(o["radio-1"], d, "radio_1")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-1"], "ObjectExtensionControllerExtenderProfileWifi-Radio1"); ok {
					if err = d.Set("radio_1", vv); err != nil {
						return fmt.Errorf("Error reading radio_1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_2", flattenObjectExtensionControllerExtenderProfileWifiRadio22edl(o["radio-2"], d, "radio_2")); err != nil {
			if vv, ok := fortiAPIPatch(o["radio-2"], "ObjectExtensionControllerExtenderProfileWifi-Radio2"); ok {
				if err = d.Set("radio_2", vv); err != nil {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenObjectExtensionControllerExtenderProfileWifiRadio22edl(o["radio-2"], d, "radio_2")); err != nil {
				if vv, ok := fortiAPIPatch(o["radio-2"], "ObjectExtensionControllerExtenderProfileWifi-Radio2"); ok {
					if err = d.Set("radio_2", vv); err != nil {
						return fmt.Errorf("Error reading radio_2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileWifiFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileWifiDfs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiCountry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["80211d"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio180211D2edl(d, i["n80211d"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1Band2edl(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bandwidth"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth2edl(d, i["bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["beacon-interval"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval2edl(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1BssColor2edl(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color-mode"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode2edl(d, i["bss_color_mode"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1Channel2edl(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "extension_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["extension-channel"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel2edl(d, i["extension_channel"], pre_append)
	}
	pre_append = pre + ".0." + "guard_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["guard-interval"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval2edl(d, i["guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lan-ext-vap"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap2edl(d, i["lan_ext_vap"], pre_append)
	}
	pre_append = pre + ".0." + "local_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-vaps"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps2edl(d, i["local_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-clients"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1MaxClients2edl(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1Mode2edl(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "operating_standard"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["operating-standard"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard2edl(d, i["operating_standard"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel2edl(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1RadioId2edl(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio1Status2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio180211D2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Band2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Bandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1BeaconInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1BssColor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1BssColorMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Channel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1ExtensionChannel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1GuardInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1LanExtVap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1LocalVaps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1MaxClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1OperatingStandard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1PowerLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1RadioId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio1Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["80211d"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio280211D2edl(d, i["n80211d"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["band"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2Band2edl(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bandwidth"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2Bandwidth2edl(d, i["bandwidth"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["beacon-interval"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2BeaconInterval2edl(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2BssColor2edl(d, i["bss_color"], pre_append)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bss-color-mode"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2BssColorMode2edl(d, i["bss_color_mode"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["channel"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2Channel2edl(d, i["channel"], pre_append)
	}
	pre_append = pre + ".0." + "extension_channel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["extension-channel"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2ExtensionChannel2edl(d, i["extension_channel"], pre_append)
	}
	pre_append = pre + ".0." + "guard_interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["guard-interval"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2GuardInterval2edl(d, i["guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lan-ext-vap"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2LanExtVap2edl(d, i["lan_ext_vap"], pre_append)
	}
	pre_append = pre + ".0." + "local_vaps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["local-vaps"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2LocalVaps2edl(d, i["local_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["max-clients"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2MaxClients2edl(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2Mode2edl(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "operating_standard"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["operating-standard"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2OperatingStandard2edl(d, i["operating_standard"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["power-level"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2PowerLevel2edl(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["radio-id"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2RadioId2edl(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtensionControllerExtenderProfileWifiRadio2Status2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio280211D2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2Band2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2Bandwidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2BeaconInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2BssColor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2BssColorMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2Channel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2ExtensionChannel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2GuardInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2LanExtVap2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2LocalVaps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2MaxClients2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2Mode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2OperatingStandard2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2PowerLevel2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2RadioId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileWifiRadio2Status2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileWifi(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dfs"); ok || d.HasChange("dfs") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiDfs2edl(d, v, "dfs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["DFS"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiCountry2edl(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok || d.HasChange("radio_1") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio12edl(d, v, "radio_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok || d.HasChange("radio_2") {
		t, err := expandObjectExtensionControllerExtenderProfileWifiRadio22edl(d, v, "radio_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	return &obj, nil
}
