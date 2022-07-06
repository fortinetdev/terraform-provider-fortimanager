// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerArrpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerArrpProfileCreate,
		Read:   resourceObjectWirelessControllerArrpProfileRead,
		Update: resourceObjectWirelessControllerArrpProfileUpdate,
		Delete: resourceObjectWirelessControllerArrpProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"darrp_optimize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"darrp_optimize_schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"include_dfs_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_weather_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"override_darrp_optimize": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selection_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_channel_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_noise_floor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_rx_errors": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"threshold_spectral_rssi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold_tx_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_channel_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_dfs_channel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_managed_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_noise_floor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_rogue_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_spectral_rssi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weight_weather_channel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerArrpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerArrpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerArrpProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerArrpProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerArrpProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerArrpProfileRead(d, m)
}

func resourceObjectWirelessControllerArrpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerArrpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerArrpProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerArrpProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerArrpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerArrpProfileRead(d, m)
}

func resourceObjectWirelessControllerArrpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerArrpProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerArrpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerArrpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerArrpProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerArrpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerArrpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerArrpProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerArrpProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileDarrpOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileDarrpOptimizeSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerArrpProfileIncludeDfsChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileIncludeWeatherChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileMonitorPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileOverrideDarrpOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileSelectionPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileThresholdAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileThresholdChannelLoad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileThresholdNoiseFloor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileThresholdRxErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileThresholdSpectralRssi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileThresholdTxRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightChannelLoad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightDfsChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightManagedAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightNoiseFloor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightRogueAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightSpectralRssi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerArrpProfileWeightWeatherChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerArrpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerArrpProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerArrpProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("darrp_optimize", flattenObjectWirelessControllerArrpProfileDarrpOptimize(o["darrp-optimize"], d, "darrp_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp-optimize"], "ObjectWirelessControllerArrpProfile-DarrpOptimize"); ok {
			if err = d.Set("darrp_optimize", vv); err != nil {
				return fmt.Errorf("Error reading darrp_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp_optimize: %v", err)
		}
	}

	if err = d.Set("darrp_optimize_schedules", flattenObjectWirelessControllerArrpProfileDarrpOptimizeSchedules(o["darrp-optimize-schedules"], d, "darrp_optimize_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["darrp-optimize-schedules"], "ObjectWirelessControllerArrpProfile-DarrpOptimizeSchedules"); ok {
			if err = d.Set("darrp_optimize_schedules", vv); err != nil {
				return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
		}
	}

	if err = d.Set("include_dfs_channel", flattenObjectWirelessControllerArrpProfileIncludeDfsChannel(o["include-dfs-channel"], d, "include_dfs_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-dfs-channel"], "ObjectWirelessControllerArrpProfile-IncludeDfsChannel"); ok {
			if err = d.Set("include_dfs_channel", vv); err != nil {
				return fmt.Errorf("Error reading include_dfs_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_dfs_channel: %v", err)
		}
	}

	if err = d.Set("include_weather_channel", flattenObjectWirelessControllerArrpProfileIncludeWeatherChannel(o["include-weather-channel"], d, "include_weather_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-weather-channel"], "ObjectWirelessControllerArrpProfile-IncludeWeatherChannel"); ok {
			if err = d.Set("include_weather_channel", vv); err != nil {
				return fmt.Errorf("Error reading include_weather_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_weather_channel: %v", err)
		}
	}

	if err = d.Set("monitor_period", flattenObjectWirelessControllerArrpProfileMonitorPeriod(o["monitor-period"], d, "monitor_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor-period"], "ObjectWirelessControllerArrpProfile-MonitorPeriod"); ok {
			if err = d.Set("monitor_period", vv); err != nil {
				return fmt.Errorf("Error reading monitor_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor_period: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerArrpProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerArrpProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_darrp_optimize", flattenObjectWirelessControllerArrpProfileOverrideDarrpOptimize(o["override-darrp-optimize"], d, "override_darrp_optimize")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-darrp-optimize"], "ObjectWirelessControllerArrpProfile-OverrideDarrpOptimize"); ok {
			if err = d.Set("override_darrp_optimize", vv); err != nil {
				return fmt.Errorf("Error reading override_darrp_optimize: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_darrp_optimize: %v", err)
		}
	}

	if err = d.Set("selection_period", flattenObjectWirelessControllerArrpProfileSelectionPeriod(o["selection-period"], d, "selection_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["selection-period"], "ObjectWirelessControllerArrpProfile-SelectionPeriod"); ok {
			if err = d.Set("selection_period", vv); err != nil {
				return fmt.Errorf("Error reading selection_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selection_period: %v", err)
		}
	}

	if err = d.Set("threshold_ap", flattenObjectWirelessControllerArrpProfileThresholdAp(o["threshold-ap"], d, "threshold_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-ap"], "ObjectWirelessControllerArrpProfile-ThresholdAp"); ok {
			if err = d.Set("threshold_ap", vv); err != nil {
				return fmt.Errorf("Error reading threshold_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_ap: %v", err)
		}
	}

	if err = d.Set("threshold_channel_load", flattenObjectWirelessControllerArrpProfileThresholdChannelLoad(o["threshold-channel-load"], d, "threshold_channel_load")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-channel-load"], "ObjectWirelessControllerArrpProfile-ThresholdChannelLoad"); ok {
			if err = d.Set("threshold_channel_load", vv); err != nil {
				return fmt.Errorf("Error reading threshold_channel_load: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_channel_load: %v", err)
		}
	}

	if err = d.Set("threshold_noise_floor", flattenObjectWirelessControllerArrpProfileThresholdNoiseFloor(o["threshold-noise-floor"], d, "threshold_noise_floor")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-noise-floor"], "ObjectWirelessControllerArrpProfile-ThresholdNoiseFloor"); ok {
			if err = d.Set("threshold_noise_floor", vv); err != nil {
				return fmt.Errorf("Error reading threshold_noise_floor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_noise_floor: %v", err)
		}
	}

	if err = d.Set("threshold_rx_errors", flattenObjectWirelessControllerArrpProfileThresholdRxErrors(o["threshold-rx-errors"], d, "threshold_rx_errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-rx-errors"], "ObjectWirelessControllerArrpProfile-ThresholdRxErrors"); ok {
			if err = d.Set("threshold_rx_errors", vv); err != nil {
				return fmt.Errorf("Error reading threshold_rx_errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_rx_errors: %v", err)
		}
	}

	if err = d.Set("threshold_spectral_rssi", flattenObjectWirelessControllerArrpProfileThresholdSpectralRssi(o["threshold-spectral-rssi"], d, "threshold_spectral_rssi")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-spectral-rssi"], "ObjectWirelessControllerArrpProfile-ThresholdSpectralRssi"); ok {
			if err = d.Set("threshold_spectral_rssi", vv); err != nil {
				return fmt.Errorf("Error reading threshold_spectral_rssi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_spectral_rssi: %v", err)
		}
	}

	if err = d.Set("threshold_tx_retries", flattenObjectWirelessControllerArrpProfileThresholdTxRetries(o["threshold-tx-retries"], d, "threshold_tx_retries")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold-tx-retries"], "ObjectWirelessControllerArrpProfile-ThresholdTxRetries"); ok {
			if err = d.Set("threshold_tx_retries", vv); err != nil {
				return fmt.Errorf("Error reading threshold_tx_retries: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold_tx_retries: %v", err)
		}
	}

	if err = d.Set("weight_channel_load", flattenObjectWirelessControllerArrpProfileWeightChannelLoad(o["weight-channel-load"], d, "weight_channel_load")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-channel-load"], "ObjectWirelessControllerArrpProfile-WeightChannelLoad"); ok {
			if err = d.Set("weight_channel_load", vv); err != nil {
				return fmt.Errorf("Error reading weight_channel_load: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_channel_load: %v", err)
		}
	}

	if err = d.Set("weight_dfs_channel", flattenObjectWirelessControllerArrpProfileWeightDfsChannel(o["weight-dfs-channel"], d, "weight_dfs_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-dfs-channel"], "ObjectWirelessControllerArrpProfile-WeightDfsChannel"); ok {
			if err = d.Set("weight_dfs_channel", vv); err != nil {
				return fmt.Errorf("Error reading weight_dfs_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_dfs_channel: %v", err)
		}
	}

	if err = d.Set("weight_managed_ap", flattenObjectWirelessControllerArrpProfileWeightManagedAp(o["weight-managed-ap"], d, "weight_managed_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-managed-ap"], "ObjectWirelessControllerArrpProfile-WeightManagedAp"); ok {
			if err = d.Set("weight_managed_ap", vv); err != nil {
				return fmt.Errorf("Error reading weight_managed_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_managed_ap: %v", err)
		}
	}

	if err = d.Set("weight_noise_floor", flattenObjectWirelessControllerArrpProfileWeightNoiseFloor(o["weight-noise-floor"], d, "weight_noise_floor")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-noise-floor"], "ObjectWirelessControllerArrpProfile-WeightNoiseFloor"); ok {
			if err = d.Set("weight_noise_floor", vv); err != nil {
				return fmt.Errorf("Error reading weight_noise_floor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_noise_floor: %v", err)
		}
	}

	if err = d.Set("weight_rogue_ap", flattenObjectWirelessControllerArrpProfileWeightRogueAp(o["weight-rogue-ap"], d, "weight_rogue_ap")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-rogue-ap"], "ObjectWirelessControllerArrpProfile-WeightRogueAp"); ok {
			if err = d.Set("weight_rogue_ap", vv); err != nil {
				return fmt.Errorf("Error reading weight_rogue_ap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_rogue_ap: %v", err)
		}
	}

	if err = d.Set("weight_spectral_rssi", flattenObjectWirelessControllerArrpProfileWeightSpectralRssi(o["weight-spectral-rssi"], d, "weight_spectral_rssi")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-spectral-rssi"], "ObjectWirelessControllerArrpProfile-WeightSpectralRssi"); ok {
			if err = d.Set("weight_spectral_rssi", vv); err != nil {
				return fmt.Errorf("Error reading weight_spectral_rssi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_spectral_rssi: %v", err)
		}
	}

	if err = d.Set("weight_weather_channel", flattenObjectWirelessControllerArrpProfileWeightWeatherChannel(o["weight-weather-channel"], d, "weight_weather_channel")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight-weather-channel"], "ObjectWirelessControllerArrpProfile-WeightWeatherChannel"); ok {
			if err = d.Set("weight_weather_channel", vv); err != nil {
				return fmt.Errorf("Error reading weight_weather_channel: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight_weather_channel: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerArrpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerArrpProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileDarrpOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileDarrpOptimizeSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerArrpProfileIncludeDfsChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileIncludeWeatherChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileMonitorPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileOverrideDarrpOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileSelectionPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileThresholdAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileThresholdChannelLoad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileThresholdNoiseFloor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileThresholdRxErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileThresholdSpectralRssi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileThresholdTxRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightChannelLoad(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightDfsChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightManagedAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightNoiseFloor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightRogueAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightSpectralRssi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerArrpProfileWeightWeatherChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerArrpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerArrpProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("darrp_optimize"); ok || d.HasChange("darrp_optimize") {
		t, err := expandObjectWirelessControllerArrpProfileDarrpOptimize(d, v, "darrp_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize"] = t
		}
	}

	if v, ok := d.GetOk("darrp_optimize_schedules"); ok || d.HasChange("darrp_optimize_schedules") {
		t, err := expandObjectWirelessControllerArrpProfileDarrpOptimizeSchedules(d, v, "darrp_optimize_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize-schedules"] = t
		}
	}

	if v, ok := d.GetOk("include_dfs_channel"); ok || d.HasChange("include_dfs_channel") {
		t, err := expandObjectWirelessControllerArrpProfileIncludeDfsChannel(d, v, "include_dfs_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-dfs-channel"] = t
		}
	}

	if v, ok := d.GetOk("include_weather_channel"); ok || d.HasChange("include_weather_channel") {
		t, err := expandObjectWirelessControllerArrpProfileIncludeWeatherChannel(d, v, "include_weather_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-weather-channel"] = t
		}
	}

	if v, ok := d.GetOk("monitor_period"); ok || d.HasChange("monitor_period") {
		t, err := expandObjectWirelessControllerArrpProfileMonitorPeriod(d, v, "monitor_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-period"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerArrpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_darrp_optimize"); ok || d.HasChange("override_darrp_optimize") {
		t, err := expandObjectWirelessControllerArrpProfileOverrideDarrpOptimize(d, v, "override_darrp_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-darrp-optimize"] = t
		}
	}

	if v, ok := d.GetOk("selection_period"); ok || d.HasChange("selection_period") {
		t, err := expandObjectWirelessControllerArrpProfileSelectionPeriod(d, v, "selection_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selection-period"] = t
		}
	}

	if v, ok := d.GetOk("threshold_ap"); ok || d.HasChange("threshold_ap") {
		t, err := expandObjectWirelessControllerArrpProfileThresholdAp(d, v, "threshold_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-ap"] = t
		}
	}

	if v, ok := d.GetOk("threshold_channel_load"); ok || d.HasChange("threshold_channel_load") {
		t, err := expandObjectWirelessControllerArrpProfileThresholdChannelLoad(d, v, "threshold_channel_load")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-channel-load"] = t
		}
	}

	if v, ok := d.GetOk("threshold_noise_floor"); ok || d.HasChange("threshold_noise_floor") {
		t, err := expandObjectWirelessControllerArrpProfileThresholdNoiseFloor(d, v, "threshold_noise_floor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-noise-floor"] = t
		}
	}

	if v, ok := d.GetOk("threshold_rx_errors"); ok || d.HasChange("threshold_rx_errors") {
		t, err := expandObjectWirelessControllerArrpProfileThresholdRxErrors(d, v, "threshold_rx_errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-rx-errors"] = t
		}
	}

	if v, ok := d.GetOk("threshold_spectral_rssi"); ok || d.HasChange("threshold_spectral_rssi") {
		t, err := expandObjectWirelessControllerArrpProfileThresholdSpectralRssi(d, v, "threshold_spectral_rssi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-spectral-rssi"] = t
		}
	}

	if v, ok := d.GetOk("threshold_tx_retries"); ok || d.HasChange("threshold_tx_retries") {
		t, err := expandObjectWirelessControllerArrpProfileThresholdTxRetries(d, v, "threshold_tx_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-tx-retries"] = t
		}
	}

	if v, ok := d.GetOk("weight_channel_load"); ok || d.HasChange("weight_channel_load") {
		t, err := expandObjectWirelessControllerArrpProfileWeightChannelLoad(d, v, "weight_channel_load")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-channel-load"] = t
		}
	}

	if v, ok := d.GetOk("weight_dfs_channel"); ok || d.HasChange("weight_dfs_channel") {
		t, err := expandObjectWirelessControllerArrpProfileWeightDfsChannel(d, v, "weight_dfs_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-dfs-channel"] = t
		}
	}

	if v, ok := d.GetOk("weight_managed_ap"); ok || d.HasChange("weight_managed_ap") {
		t, err := expandObjectWirelessControllerArrpProfileWeightManagedAp(d, v, "weight_managed_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-managed-ap"] = t
		}
	}

	if v, ok := d.GetOk("weight_noise_floor"); ok || d.HasChange("weight_noise_floor") {
		t, err := expandObjectWirelessControllerArrpProfileWeightNoiseFloor(d, v, "weight_noise_floor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-noise-floor"] = t
		}
	}

	if v, ok := d.GetOk("weight_rogue_ap"); ok || d.HasChange("weight_rogue_ap") {
		t, err := expandObjectWirelessControllerArrpProfileWeightRogueAp(d, v, "weight_rogue_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-rogue-ap"] = t
		}
	}

	if v, ok := d.GetOk("weight_spectral_rssi"); ok || d.HasChange("weight_spectral_rssi") {
		t, err := expandObjectWirelessControllerArrpProfileWeightSpectralRssi(d, v, "weight_spectral_rssi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-spectral-rssi"] = t
		}
	}

	if v, ok := d.GetOk("weight_weather_channel"); ok || d.HasChange("weight_weather_channel") {
		t, err := expandObjectWirelessControllerArrpProfileWeightWeatherChannel(d, v, "weight_weather_channel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-weather-channel"] = t
		}
	}

	return &obj, nil
}
