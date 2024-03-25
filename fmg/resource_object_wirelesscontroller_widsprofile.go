// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure wireless intrusion detection system (WIDS) profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWidsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWidsProfileCreate,
		Read:   resourceObjectWirelessControllerWidsProfileRead,
		Update: resourceObjectWirelessControllerWidsProfileUpdate,
		Delete: resourceObjectWirelessControllerWidsProfileDelete,

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
			"ap_auto_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_bgscan_disable_day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_bgscan_disable_end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_bgscan_disable_schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_bgscan_disable_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_bgscan_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_bgscan_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_fgscan_report_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ap_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_scan_channel_list_2g_5g": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_scan_channel_list_6g": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ap_scan_passive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_scan_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"asleap_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"assoc_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"assoc_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_flood_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_flood_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"deauth_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deauth_unknown_src_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_start_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_start_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_start_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_intv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"invalid_mac_oui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_duration_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_duration_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"null_ssid_probe_resp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sensor_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spoofed_deauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weak_wep_iv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_bridge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWidsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWirelessControllerWidsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWidsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerWidsProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerWidsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerWidsProfileRead(d, m)
}

func resourceObjectWirelessControllerWidsProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerWidsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWidsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerWidsProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWidsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerWidsProfileRead(d, m)
}

func resourceObjectWirelessControllerWidsProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerWidsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWidsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWidsProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerWidsProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWidsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWidsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWidsProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWidsProfileApAutoSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanDisableDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWidsProfileApBgscanDisableEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanDisableSchedules(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWidsProfileApBgscanDisableStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApBgscanReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApFgscanReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApScanChannelList2G5G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWidsProfileApScanChannelList6G(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerWidsProfileApScanPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileApScanThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAsleapAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAssocFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAssocFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAssocFrameFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAuthFloodThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAuthFloodTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileAuthFrameFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileDeauthBroadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileDeauthUnknownSrcThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolFailFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolFailIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolFailThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolLogoffFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolLogoffIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolLogoffThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolPreFailFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolPreFailIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolPreFailThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolPreSuccFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolPreSuccIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolPreSuccThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolStartFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolStartIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolStartThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolSuccFlood(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolSuccIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileEapolSuccThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileInvalidMacOui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileLongDurationAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileLongDurationThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileNullSsidProbeResp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileSensorMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileSpoofedDeauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileWeakWepIv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWidsProfileWirelessBridge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWidsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ap_auto_suppress", flattenObjectWirelessControllerWidsProfileApAutoSuppress(o["ap-auto-suppress"], d, "ap_auto_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-auto-suppress"], "ObjectWirelessControllerWidsProfile-ApAutoSuppress"); ok {
			if err = d.Set("ap_auto_suppress", vv); err != nil {
				return fmt.Errorf("Error reading ap_auto_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_auto_suppress: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_day", flattenObjectWirelessControllerWidsProfileApBgscanDisableDay(o["ap-bgscan-disable-day"], d, "ap_bgscan_disable_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-disable-day"], "ObjectWirelessControllerWidsProfile-ApBgscanDisableDay"); ok {
			if err = d.Set("ap_bgscan_disable_day", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_disable_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_disable_day: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_end", flattenObjectWirelessControllerWidsProfileApBgscanDisableEnd(o["ap-bgscan-disable-end"], d, "ap_bgscan_disable_end")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-disable-end"], "ObjectWirelessControllerWidsProfile-ApBgscanDisableEnd"); ok {
			if err = d.Set("ap_bgscan_disable_end", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_disable_end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_disable_end: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_schedules", flattenObjectWirelessControllerWidsProfileApBgscanDisableSchedules(o["ap-bgscan-disable-schedules"], d, "ap_bgscan_disable_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-disable-schedules"], "ObjectWirelessControllerWidsProfile-ApBgscanDisableSchedules"); ok {
			if err = d.Set("ap_bgscan_disable_schedules", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_disable_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_disable_schedules: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_start", flattenObjectWirelessControllerWidsProfileApBgscanDisableStart(o["ap-bgscan-disable-start"], d, "ap_bgscan_disable_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-disable-start"], "ObjectWirelessControllerWidsProfile-ApBgscanDisableStart"); ok {
			if err = d.Set("ap_bgscan_disable_start", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_disable_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_disable_start: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_duration", flattenObjectWirelessControllerWidsProfileApBgscanDuration(o["ap-bgscan-duration"], d, "ap_bgscan_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-duration"], "ObjectWirelessControllerWidsProfile-ApBgscanDuration"); ok {
			if err = d.Set("ap_bgscan_duration", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_duration: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_idle", flattenObjectWirelessControllerWidsProfileApBgscanIdle(o["ap-bgscan-idle"], d, "ap_bgscan_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-idle"], "ObjectWirelessControllerWidsProfile-ApBgscanIdle"); ok {
			if err = d.Set("ap_bgscan_idle", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_idle: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_intv", flattenObjectWirelessControllerWidsProfileApBgscanIntv(o["ap-bgscan-intv"], d, "ap_bgscan_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-intv"], "ObjectWirelessControllerWidsProfile-ApBgscanIntv"); ok {
			if err = d.Set("ap_bgscan_intv", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_intv: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_period", flattenObjectWirelessControllerWidsProfileApBgscanPeriod(o["ap-bgscan-period"], d, "ap_bgscan_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-period"], "ObjectWirelessControllerWidsProfile-ApBgscanPeriod"); ok {
			if err = d.Set("ap_bgscan_period", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_period: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_report_intv", flattenObjectWirelessControllerWidsProfileApBgscanReportIntv(o["ap-bgscan-report-intv"], d, "ap_bgscan_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-bgscan-report-intv"], "ObjectWirelessControllerWidsProfile-ApBgscanReportIntv"); ok {
			if err = d.Set("ap_bgscan_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading ap_bgscan_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_bgscan_report_intv: %v", err)
		}
	}

	if err = d.Set("ap_fgscan_report_intv", flattenObjectWirelessControllerWidsProfileApFgscanReportIntv(o["ap-fgscan-report-intv"], d, "ap_fgscan_report_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-fgscan-report-intv"], "ObjectWirelessControllerWidsProfile-ApFgscanReportIntv"); ok {
			if err = d.Set("ap_fgscan_report_intv", vv); err != nil {
				return fmt.Errorf("Error reading ap_fgscan_report_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_fgscan_report_intv: %v", err)
		}
	}

	if err = d.Set("ap_scan", flattenObjectWirelessControllerWidsProfileApScan(o["ap-scan"], d, "ap_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan"], "ObjectWirelessControllerWidsProfile-ApScan"); ok {
			if err = d.Set("ap_scan", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan: %v", err)
		}
	}

	if err = d.Set("ap_scan_channel_list_2g_5g", flattenObjectWirelessControllerWidsProfileApScanChannelList2G5G(o["ap-scan-channel-list-2G-5G"], d, "ap_scan_channel_list_2g_5g")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-channel-list-2G-5G"], "ObjectWirelessControllerWidsProfile-ApScanChannelList2G5G"); ok {
			if err = d.Set("ap_scan_channel_list_2g_5g", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_channel_list_2g_5g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_channel_list_2g_5g: %v", err)
		}
	}

	if err = d.Set("ap_scan_channel_list_6g", flattenObjectWirelessControllerWidsProfileApScanChannelList6G(o["ap-scan-channel-list-6G"], d, "ap_scan_channel_list_6g")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-channel-list-6G"], "ObjectWirelessControllerWidsProfile-ApScanChannelList6G"); ok {
			if err = d.Set("ap_scan_channel_list_6g", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_channel_list_6g: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_channel_list_6g: %v", err)
		}
	}

	if err = d.Set("ap_scan_passive", flattenObjectWirelessControllerWidsProfileApScanPassive(o["ap-scan-passive"], d, "ap_scan_passive")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-passive"], "ObjectWirelessControllerWidsProfile-ApScanPassive"); ok {
			if err = d.Set("ap_scan_passive", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_passive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_passive: %v", err)
		}
	}

	if err = d.Set("ap_scan_threshold", flattenObjectWirelessControllerWidsProfileApScanThreshold(o["ap-scan-threshold"], d, "ap_scan_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ap-scan-threshold"], "ObjectWirelessControllerWidsProfile-ApScanThreshold"); ok {
			if err = d.Set("ap_scan_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ap_scan_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ap_scan_threshold: %v", err)
		}
	}

	if err = d.Set("asleap_attack", flattenObjectWirelessControllerWidsProfileAsleapAttack(o["asleap-attack"], d, "asleap_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["asleap-attack"], "ObjectWirelessControllerWidsProfile-AsleapAttack"); ok {
			if err = d.Set("asleap_attack", vv); err != nil {
				return fmt.Errorf("Error reading asleap_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading asleap_attack: %v", err)
		}
	}

	if err = d.Set("assoc_flood_thresh", flattenObjectWirelessControllerWidsProfileAssocFloodThresh(o["assoc-flood-thresh"], d, "assoc_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["assoc-flood-thresh"], "ObjectWirelessControllerWidsProfile-AssocFloodThresh"); ok {
			if err = d.Set("assoc_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading assoc_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assoc_flood_thresh: %v", err)
		}
	}

	if err = d.Set("assoc_flood_time", flattenObjectWirelessControllerWidsProfileAssocFloodTime(o["assoc-flood-time"], d, "assoc_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["assoc-flood-time"], "ObjectWirelessControllerWidsProfile-AssocFloodTime"); ok {
			if err = d.Set("assoc_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading assoc_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assoc_flood_time: %v", err)
		}
	}

	if err = d.Set("assoc_frame_flood", flattenObjectWirelessControllerWidsProfileAssocFrameFlood(o["assoc-frame-flood"], d, "assoc_frame_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["assoc-frame-flood"], "ObjectWirelessControllerWidsProfile-AssocFrameFlood"); ok {
			if err = d.Set("assoc_frame_flood", vv); err != nil {
				return fmt.Errorf("Error reading assoc_frame_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assoc_frame_flood: %v", err)
		}
	}

	if err = d.Set("auth_flood_thresh", flattenObjectWirelessControllerWidsProfileAuthFloodThresh(o["auth-flood-thresh"], d, "auth_flood_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-flood-thresh"], "ObjectWirelessControllerWidsProfile-AuthFloodThresh"); ok {
			if err = d.Set("auth_flood_thresh", vv); err != nil {
				return fmt.Errorf("Error reading auth_flood_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_flood_thresh: %v", err)
		}
	}

	if err = d.Set("auth_flood_time", flattenObjectWirelessControllerWidsProfileAuthFloodTime(o["auth-flood-time"], d, "auth_flood_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-flood-time"], "ObjectWirelessControllerWidsProfile-AuthFloodTime"); ok {
			if err = d.Set("auth_flood_time", vv); err != nil {
				return fmt.Errorf("Error reading auth_flood_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_flood_time: %v", err)
		}
	}

	if err = d.Set("auth_frame_flood", flattenObjectWirelessControllerWidsProfileAuthFrameFlood(o["auth-frame-flood"], d, "auth_frame_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-frame-flood"], "ObjectWirelessControllerWidsProfile-AuthFrameFlood"); ok {
			if err = d.Set("auth_frame_flood", vv); err != nil {
				return fmt.Errorf("Error reading auth_frame_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_frame_flood: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWirelessControllerWidsProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerWidsProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("deauth_broadcast", flattenObjectWirelessControllerWidsProfileDeauthBroadcast(o["deauth-broadcast"], d, "deauth_broadcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["deauth-broadcast"], "ObjectWirelessControllerWidsProfile-DeauthBroadcast"); ok {
			if err = d.Set("deauth_broadcast", vv); err != nil {
				return fmt.Errorf("Error reading deauth_broadcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deauth_broadcast: %v", err)
		}
	}

	if err = d.Set("deauth_unknown_src_thresh", flattenObjectWirelessControllerWidsProfileDeauthUnknownSrcThresh(o["deauth-unknown-src-thresh"], d, "deauth_unknown_src_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["deauth-unknown-src-thresh"], "ObjectWirelessControllerWidsProfile-DeauthUnknownSrcThresh"); ok {
			if err = d.Set("deauth_unknown_src_thresh", vv); err != nil {
				return fmt.Errorf("Error reading deauth_unknown_src_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deauth_unknown_src_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_fail_flood", flattenObjectWirelessControllerWidsProfileEapolFailFlood(o["eapol-fail-flood"], d, "eapol_fail_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-fail-flood"], "ObjectWirelessControllerWidsProfile-EapolFailFlood"); ok {
			if err = d.Set("eapol_fail_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_fail_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_fail_flood: %v", err)
		}
	}

	if err = d.Set("eapol_fail_intv", flattenObjectWirelessControllerWidsProfileEapolFailIntv(o["eapol-fail-intv"], d, "eapol_fail_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-fail-intv"], "ObjectWirelessControllerWidsProfile-EapolFailIntv"); ok {
			if err = d.Set("eapol_fail_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_fail_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_fail_intv: %v", err)
		}
	}

	if err = d.Set("eapol_fail_thresh", flattenObjectWirelessControllerWidsProfileEapolFailThresh(o["eapol-fail-thresh"], d, "eapol_fail_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-fail-thresh"], "ObjectWirelessControllerWidsProfile-EapolFailThresh"); ok {
			if err = d.Set("eapol_fail_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_fail_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_fail_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_flood", flattenObjectWirelessControllerWidsProfileEapolLogoffFlood(o["eapol-logoff-flood"], d, "eapol_logoff_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-logoff-flood"], "ObjectWirelessControllerWidsProfile-EapolLogoffFlood"); ok {
			if err = d.Set("eapol_logoff_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_logoff_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_logoff_flood: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_intv", flattenObjectWirelessControllerWidsProfileEapolLogoffIntv(o["eapol-logoff-intv"], d, "eapol_logoff_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-logoff-intv"], "ObjectWirelessControllerWidsProfile-EapolLogoffIntv"); ok {
			if err = d.Set("eapol_logoff_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_logoff_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_logoff_intv: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_thresh", flattenObjectWirelessControllerWidsProfileEapolLogoffThresh(o["eapol-logoff-thresh"], d, "eapol_logoff_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-logoff-thresh"], "ObjectWirelessControllerWidsProfile-EapolLogoffThresh"); ok {
			if err = d.Set("eapol_logoff_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_logoff_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_logoff_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_flood", flattenObjectWirelessControllerWidsProfileEapolPreFailFlood(o["eapol-pre-fail-flood"], d, "eapol_pre_fail_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-fail-flood"], "ObjectWirelessControllerWidsProfile-EapolPreFailFlood"); ok {
			if err = d.Set("eapol_pre_fail_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_fail_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_fail_flood: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_intv", flattenObjectWirelessControllerWidsProfileEapolPreFailIntv(o["eapol-pre-fail-intv"], d, "eapol_pre_fail_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-fail-intv"], "ObjectWirelessControllerWidsProfile-EapolPreFailIntv"); ok {
			if err = d.Set("eapol_pre_fail_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_fail_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_fail_intv: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_thresh", flattenObjectWirelessControllerWidsProfileEapolPreFailThresh(o["eapol-pre-fail-thresh"], d, "eapol_pre_fail_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-fail-thresh"], "ObjectWirelessControllerWidsProfile-EapolPreFailThresh"); ok {
			if err = d.Set("eapol_pre_fail_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_fail_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_fail_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_flood", flattenObjectWirelessControllerWidsProfileEapolPreSuccFlood(o["eapol-pre-succ-flood"], d, "eapol_pre_succ_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-succ-flood"], "ObjectWirelessControllerWidsProfile-EapolPreSuccFlood"); ok {
			if err = d.Set("eapol_pre_succ_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_succ_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_succ_flood: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_intv", flattenObjectWirelessControllerWidsProfileEapolPreSuccIntv(o["eapol-pre-succ-intv"], d, "eapol_pre_succ_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-succ-intv"], "ObjectWirelessControllerWidsProfile-EapolPreSuccIntv"); ok {
			if err = d.Set("eapol_pre_succ_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_succ_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_succ_intv: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_thresh", flattenObjectWirelessControllerWidsProfileEapolPreSuccThresh(o["eapol-pre-succ-thresh"], d, "eapol_pre_succ_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-pre-succ-thresh"], "ObjectWirelessControllerWidsProfile-EapolPreSuccThresh"); ok {
			if err = d.Set("eapol_pre_succ_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_pre_succ_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_pre_succ_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_start_flood", flattenObjectWirelessControllerWidsProfileEapolStartFlood(o["eapol-start-flood"], d, "eapol_start_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-start-flood"], "ObjectWirelessControllerWidsProfile-EapolStartFlood"); ok {
			if err = d.Set("eapol_start_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_start_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_start_flood: %v", err)
		}
	}

	if err = d.Set("eapol_start_intv", flattenObjectWirelessControllerWidsProfileEapolStartIntv(o["eapol-start-intv"], d, "eapol_start_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-start-intv"], "ObjectWirelessControllerWidsProfile-EapolStartIntv"); ok {
			if err = d.Set("eapol_start_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_start_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_start_intv: %v", err)
		}
	}

	if err = d.Set("eapol_start_thresh", flattenObjectWirelessControllerWidsProfileEapolStartThresh(o["eapol-start-thresh"], d, "eapol_start_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-start-thresh"], "ObjectWirelessControllerWidsProfile-EapolStartThresh"); ok {
			if err = d.Set("eapol_start_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_start_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_start_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_succ_flood", flattenObjectWirelessControllerWidsProfileEapolSuccFlood(o["eapol-succ-flood"], d, "eapol_succ_flood")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-succ-flood"], "ObjectWirelessControllerWidsProfile-EapolSuccFlood"); ok {
			if err = d.Set("eapol_succ_flood", vv); err != nil {
				return fmt.Errorf("Error reading eapol_succ_flood: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_succ_flood: %v", err)
		}
	}

	if err = d.Set("eapol_succ_intv", flattenObjectWirelessControllerWidsProfileEapolSuccIntv(o["eapol-succ-intv"], d, "eapol_succ_intv")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-succ-intv"], "ObjectWirelessControllerWidsProfile-EapolSuccIntv"); ok {
			if err = d.Set("eapol_succ_intv", vv); err != nil {
				return fmt.Errorf("Error reading eapol_succ_intv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_succ_intv: %v", err)
		}
	}

	if err = d.Set("eapol_succ_thresh", flattenObjectWirelessControllerWidsProfileEapolSuccThresh(o["eapol-succ-thresh"], d, "eapol_succ_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["eapol-succ-thresh"], "ObjectWirelessControllerWidsProfile-EapolSuccThresh"); ok {
			if err = d.Set("eapol_succ_thresh", vv); err != nil {
				return fmt.Errorf("Error reading eapol_succ_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eapol_succ_thresh: %v", err)
		}
	}

	if err = d.Set("invalid_mac_oui", flattenObjectWirelessControllerWidsProfileInvalidMacOui(o["invalid-mac-oui"], d, "invalid_mac_oui")); err != nil {
		if vv, ok := fortiAPIPatch(o["invalid-mac-oui"], "ObjectWirelessControllerWidsProfile-InvalidMacOui"); ok {
			if err = d.Set("invalid_mac_oui", vv); err != nil {
				return fmt.Errorf("Error reading invalid_mac_oui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading invalid_mac_oui: %v", err)
		}
	}

	if err = d.Set("long_duration_attack", flattenObjectWirelessControllerWidsProfileLongDurationAttack(o["long-duration-attack"], d, "long_duration_attack")); err != nil {
		if vv, ok := fortiAPIPatch(o["long-duration-attack"], "ObjectWirelessControllerWidsProfile-LongDurationAttack"); ok {
			if err = d.Set("long_duration_attack", vv); err != nil {
				return fmt.Errorf("Error reading long_duration_attack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading long_duration_attack: %v", err)
		}
	}

	if err = d.Set("long_duration_thresh", flattenObjectWirelessControllerWidsProfileLongDurationThresh(o["long-duration-thresh"], d, "long_duration_thresh")); err != nil {
		if vv, ok := fortiAPIPatch(o["long-duration-thresh"], "ObjectWirelessControllerWidsProfile-LongDurationThresh"); ok {
			if err = d.Set("long_duration_thresh", vv); err != nil {
				return fmt.Errorf("Error reading long_duration_thresh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading long_duration_thresh: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerWidsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerWidsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("null_ssid_probe_resp", flattenObjectWirelessControllerWidsProfileNullSsidProbeResp(o["null-ssid-probe-resp"], d, "null_ssid_probe_resp")); err != nil {
		if vv, ok := fortiAPIPatch(o["null-ssid-probe-resp"], "ObjectWirelessControllerWidsProfile-NullSsidProbeResp"); ok {
			if err = d.Set("null_ssid_probe_resp", vv); err != nil {
				return fmt.Errorf("Error reading null_ssid_probe_resp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading null_ssid_probe_resp: %v", err)
		}
	}

	if err = d.Set("sensor_mode", flattenObjectWirelessControllerWidsProfileSensorMode(o["sensor-mode"], d, "sensor_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["sensor-mode"], "ObjectWirelessControllerWidsProfile-SensorMode"); ok {
			if err = d.Set("sensor_mode", vv); err != nil {
				return fmt.Errorf("Error reading sensor_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sensor_mode: %v", err)
		}
	}

	if err = d.Set("spoofed_deauth", flattenObjectWirelessControllerWidsProfileSpoofedDeauth(o["spoofed-deauth"], d, "spoofed_deauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["spoofed-deauth"], "ObjectWirelessControllerWidsProfile-SpoofedDeauth"); ok {
			if err = d.Set("spoofed_deauth", vv); err != nil {
				return fmt.Errorf("Error reading spoofed_deauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spoofed_deauth: %v", err)
		}
	}

	if err = d.Set("weak_wep_iv", flattenObjectWirelessControllerWidsProfileWeakWepIv(o["weak-wep-iv"], d, "weak_wep_iv")); err != nil {
		if vv, ok := fortiAPIPatch(o["weak-wep-iv"], "ObjectWirelessControllerWidsProfile-WeakWepIv"); ok {
			if err = d.Set("weak_wep_iv", vv); err != nil {
				return fmt.Errorf("Error reading weak_wep_iv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weak_wep_iv: %v", err)
		}
	}

	if err = d.Set("wireless_bridge", flattenObjectWirelessControllerWidsProfileWirelessBridge(o["wireless-bridge"], d, "wireless_bridge")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-bridge"], "ObjectWirelessControllerWidsProfile-WirelessBridge"); ok {
			if err = d.Set("wireless_bridge", vv); err != nil {
				return fmt.Errorf("Error reading wireless_bridge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_bridge: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWidsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWidsProfileApAutoSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanDisableDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWidsProfileApBgscanDisableEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanDisableSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWidsProfileApBgscanDisableStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApBgscanReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApFgscanReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApScanChannelList2G5G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWidsProfileApScanChannelList6G(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerWidsProfileApScanPassive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileApScanThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAsleapAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAssocFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAssocFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAssocFrameFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAuthFloodThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAuthFloodTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileAuthFrameFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileDeauthBroadcast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileDeauthUnknownSrcThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolFailFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolFailIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolFailThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolLogoffFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolLogoffIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolLogoffThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolPreFailFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolPreFailIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolPreFailThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolPreSuccFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolPreSuccIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolPreSuccThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolStartFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolStartIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolStartThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolSuccFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolSuccIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileEapolSuccThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileInvalidMacOui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileLongDurationAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileLongDurationThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileNullSsidProbeResp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileSensorMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileSpoofedDeauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileWeakWepIv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWidsProfileWirelessBridge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWidsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ap_auto_suppress"); ok || d.HasChange("ap_auto_suppress") {
		t, err := expandObjectWirelessControllerWidsProfileApAutoSuppress(d, v, "ap_auto_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-auto-suppress"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_day"); ok || d.HasChange("ap_bgscan_disable_day") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanDisableDay(d, v, "ap_bgscan_disable_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-day"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_end"); ok || d.HasChange("ap_bgscan_disable_end") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanDisableEnd(d, v, "ap_bgscan_disable_end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-end"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_schedules"); ok || d.HasChange("ap_bgscan_disable_schedules") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanDisableSchedules(d, v, "ap_bgscan_disable_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-schedules"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_start"); ok || d.HasChange("ap_bgscan_disable_start") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanDisableStart(d, v, "ap_bgscan_disable_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-start"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_duration"); ok || d.HasChange("ap_bgscan_duration") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanDuration(d, v, "ap_bgscan_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-duration"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_idle"); ok || d.HasChange("ap_bgscan_idle") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanIdle(d, v, "ap_bgscan_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-idle"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_intv"); ok || d.HasChange("ap_bgscan_intv") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanIntv(d, v, "ap_bgscan_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_period"); ok || d.HasChange("ap_bgscan_period") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanPeriod(d, v, "ap_bgscan_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-period"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_report_intv"); ok || d.HasChange("ap_bgscan_report_intv") {
		t, err := expandObjectWirelessControllerWidsProfileApBgscanReportIntv(d, v, "ap_bgscan_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_fgscan_report_intv"); ok || d.HasChange("ap_fgscan_report_intv") {
		t, err := expandObjectWirelessControllerWidsProfileApFgscanReportIntv(d, v, "ap_fgscan_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-fgscan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan"); ok || d.HasChange("ap_scan") {
		t, err := expandObjectWirelessControllerWidsProfileApScan(d, v, "ap_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_channel_list_2g_5g"); ok || d.HasChange("ap_scan_channel_list_2g_5g") {
		t, err := expandObjectWirelessControllerWidsProfileApScanChannelList2G5G(d, v, "ap_scan_channel_list_2g_5g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-channel-list-2G-5G"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_channel_list_6g"); ok || d.HasChange("ap_scan_channel_list_6g") {
		t, err := expandObjectWirelessControllerWidsProfileApScanChannelList6G(d, v, "ap_scan_channel_list_6g")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-channel-list-6G"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_passive"); ok || d.HasChange("ap_scan_passive") {
		t, err := expandObjectWirelessControllerWidsProfileApScanPassive(d, v, "ap_scan_passive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-passive"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_threshold"); ok || d.HasChange("ap_scan_threshold") {
		t, err := expandObjectWirelessControllerWidsProfileApScanThreshold(d, v, "ap_scan_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-threshold"] = t
		}
	}

	if v, ok := d.GetOk("asleap_attack"); ok || d.HasChange("asleap_attack") {
		t, err := expandObjectWirelessControllerWidsProfileAsleapAttack(d, v, "asleap_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asleap-attack"] = t
		}
	}

	if v, ok := d.GetOk("assoc_flood_thresh"); ok || d.HasChange("assoc_flood_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileAssocFloodThresh(d, v, "assoc_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("assoc_flood_time"); ok || d.HasChange("assoc_flood_time") {
		t, err := expandObjectWirelessControllerWidsProfileAssocFloodTime(d, v, "assoc_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("assoc_frame_flood"); ok || d.HasChange("assoc_frame_flood") {
		t, err := expandObjectWirelessControllerWidsProfileAssocFrameFlood(d, v, "assoc_frame_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-frame-flood"] = t
		}
	}

	if v, ok := d.GetOk("auth_flood_thresh"); ok || d.HasChange("auth_flood_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileAuthFloodThresh(d, v, "auth_flood_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("auth_flood_time"); ok || d.HasChange("auth_flood_time") {
		t, err := expandObjectWirelessControllerWidsProfileAuthFloodTime(d, v, "auth_flood_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("auth_frame_flood"); ok || d.HasChange("auth_frame_flood") {
		t, err := expandObjectWirelessControllerWidsProfileAuthFrameFlood(d, v, "auth_frame_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-frame-flood"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerWidsProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("deauth_broadcast"); ok || d.HasChange("deauth_broadcast") {
		t, err := expandObjectWirelessControllerWidsProfileDeauthBroadcast(d, v, "deauth_broadcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("deauth_unknown_src_thresh"); ok || d.HasChange("deauth_unknown_src_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileDeauthUnknownSrcThresh(d, v, "deauth_unknown_src_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-unknown-src-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_flood"); ok || d.HasChange("eapol_fail_flood") {
		t, err := expandObjectWirelessControllerWidsProfileEapolFailFlood(d, v, "eapol_fail_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_intv"); ok || d.HasChange("eapol_fail_intv") {
		t, err := expandObjectWirelessControllerWidsProfileEapolFailIntv(d, v, "eapol_fail_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_thresh"); ok || d.HasChange("eapol_fail_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileEapolFailThresh(d, v, "eapol_fail_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_flood"); ok || d.HasChange("eapol_logoff_flood") {
		t, err := expandObjectWirelessControllerWidsProfileEapolLogoffFlood(d, v, "eapol_logoff_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_intv"); ok || d.HasChange("eapol_logoff_intv") {
		t, err := expandObjectWirelessControllerWidsProfileEapolLogoffIntv(d, v, "eapol_logoff_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_thresh"); ok || d.HasChange("eapol_logoff_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileEapolLogoffThresh(d, v, "eapol_logoff_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_flood"); ok || d.HasChange("eapol_pre_fail_flood") {
		t, err := expandObjectWirelessControllerWidsProfileEapolPreFailFlood(d, v, "eapol_pre_fail_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_intv"); ok || d.HasChange("eapol_pre_fail_intv") {
		t, err := expandObjectWirelessControllerWidsProfileEapolPreFailIntv(d, v, "eapol_pre_fail_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_thresh"); ok || d.HasChange("eapol_pre_fail_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileEapolPreFailThresh(d, v, "eapol_pre_fail_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_flood"); ok || d.HasChange("eapol_pre_succ_flood") {
		t, err := expandObjectWirelessControllerWidsProfileEapolPreSuccFlood(d, v, "eapol_pre_succ_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_intv"); ok || d.HasChange("eapol_pre_succ_intv") {
		t, err := expandObjectWirelessControllerWidsProfileEapolPreSuccIntv(d, v, "eapol_pre_succ_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_thresh"); ok || d.HasChange("eapol_pre_succ_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileEapolPreSuccThresh(d, v, "eapol_pre_succ_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_flood"); ok || d.HasChange("eapol_start_flood") {
		t, err := expandObjectWirelessControllerWidsProfileEapolStartFlood(d, v, "eapol_start_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_intv"); ok || d.HasChange("eapol_start_intv") {
		t, err := expandObjectWirelessControllerWidsProfileEapolStartIntv(d, v, "eapol_start_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_thresh"); ok || d.HasChange("eapol_start_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileEapolStartThresh(d, v, "eapol_start_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_flood"); ok || d.HasChange("eapol_succ_flood") {
		t, err := expandObjectWirelessControllerWidsProfileEapolSuccFlood(d, v, "eapol_succ_flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_intv"); ok || d.HasChange("eapol_succ_intv") {
		t, err := expandObjectWirelessControllerWidsProfileEapolSuccIntv(d, v, "eapol_succ_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_thresh"); ok || d.HasChange("eapol_succ_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileEapolSuccThresh(d, v, "eapol_succ_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-thresh"] = t
		}
	}

	if v, ok := d.GetOk("invalid_mac_oui"); ok || d.HasChange("invalid_mac_oui") {
		t, err := expandObjectWirelessControllerWidsProfileInvalidMacOui(d, v, "invalid_mac_oui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-mac-oui"] = t
		}
	}

	if v, ok := d.GetOk("long_duration_attack"); ok || d.HasChange("long_duration_attack") {
		t, err := expandObjectWirelessControllerWidsProfileLongDurationAttack(d, v, "long_duration_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-duration-attack"] = t
		}
	}

	if v, ok := d.GetOk("long_duration_thresh"); ok || d.HasChange("long_duration_thresh") {
		t, err := expandObjectWirelessControllerWidsProfileLongDurationThresh(d, v, "long_duration_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-duration-thresh"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerWidsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("null_ssid_probe_resp"); ok || d.HasChange("null_ssid_probe_resp") {
		t, err := expandObjectWirelessControllerWidsProfileNullSsidProbeResp(d, v, "null_ssid_probe_resp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-ssid-probe-resp"] = t
		}
	}

	if v, ok := d.GetOk("sensor_mode"); ok || d.HasChange("sensor_mode") {
		t, err := expandObjectWirelessControllerWidsProfileSensorMode(d, v, "sensor_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensor-mode"] = t
		}
	}

	if v, ok := d.GetOk("spoofed_deauth"); ok || d.HasChange("spoofed_deauth") {
		t, err := expandObjectWirelessControllerWidsProfileSpoofedDeauth(d, v, "spoofed_deauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spoofed-deauth"] = t
		}
	}

	if v, ok := d.GetOk("weak_wep_iv"); ok || d.HasChange("weak_wep_iv") {
		t, err := expandObjectWirelessControllerWidsProfileWeakWepIv(d, v, "weak_wep_iv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weak-wep-iv"] = t
		}
	}

	if v, ok := d.GetOk("wireless_bridge"); ok || d.HasChange("wireless_bridge") {
		t, err := expandObjectWirelessControllerWidsProfileWirelessBridge(d, v, "wireless_bridge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-bridge"] = t
		}
	}

	return &obj, nil
}
