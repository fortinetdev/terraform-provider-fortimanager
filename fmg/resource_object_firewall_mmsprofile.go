// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure MMS profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallMmsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallMmsProfileCreate,
		Read:   resourceObjectFirewallMmsProfileRead,
		Update: resourceObjectFirewallMmsProfileUpdate,
		Delete: resourceObjectFirewallMmsProfileDelete,

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
			"avnotificationtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bwordtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"carrier_endpoint_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"carrier_endpoint_prefix_range_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"carrier_endpoint_prefix_range_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"carrier_endpoint_prefix_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"carrierendpointbwltable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dupe": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action1": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"action2": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"action3": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"block_time1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"block_time2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"block_time3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"window1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"window2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"window3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"extended_utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flood": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action1": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"action2": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"action3": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"block_time1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"block_time2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"block_time3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"limit3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"window1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"window2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"window3": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"mm1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"mm1_addr_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm1_addr_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm1_convert_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm1_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm1_retr_dupe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm1_retrieve_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm1comfortamount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm1comfortinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm1oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"mm3_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm3oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm4": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"mm4_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm4oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm7": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"mm7_addr_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm7_addr_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm7_convert_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm7_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm7comfortamount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm7comfortinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mm7oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mms_antispam_mass_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_av_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_av_oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_av_virus_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_carrier_endpoint_filter_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_checksum_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_checksum_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_notification_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mms_web_content_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mmsbwordthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"notif_msisdn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msisdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"threshold": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
			},
			"notification": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"alert_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"alert_src_msisdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"alert_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bword_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bword_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bword_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"carrier_endpoint_bwl_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"carrier_endpoint_bwl_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"carrier_endpoint_bwl_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"days_allowed": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"detect_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dupe_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dupe_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dupe_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_block_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"file_block_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_block_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"flood_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flood_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"flood_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"from_in_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mms_checksum_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mms_checksum_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mms_checksum_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mmsc_hostname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mmsc_password": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"mmsc_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mmsc_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mmsc_username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rate_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tod_window_duration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tod_window_end": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tod_window_start": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vas_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vasp_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"virus_int": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"virus_int_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"virus_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"outbreak_prevention": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ftgd_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"remove_blocked_const_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFirewallMmsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallMmsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallMmsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallMmsProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallMmsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallMmsProfileRead(d, m)
}

func resourceObjectFirewallMmsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallMmsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallMmsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallMmsProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallMmsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallMmsProfileRead(d, m)
}

func resourceObjectFirewallMmsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallMmsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallMmsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallMmsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallMmsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallMmsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallMmsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallMmsProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallMmsProfileAvnotificationtable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileBwordtable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileCarrierEndpointPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileCarrierEndpointPrefixRangeMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileCarrierEndpointPrefixRangeMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileCarrierEndpointPrefixString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileCarrierendpointbwltable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupe(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action1"
	if _, ok := i["action1"]; ok {
		result["action1"] = flattenObjectFirewallMmsProfileDupeAction1(i["action1"], d, pre_append)
	}

	pre_append = pre + ".0." + "action2"
	if _, ok := i["action2"]; ok {
		result["action2"] = flattenObjectFirewallMmsProfileDupeAction2(i["action2"], d, pre_append)
	}

	pre_append = pre + ".0." + "action3"
	if _, ok := i["action3"]; ok {
		result["action3"] = flattenObjectFirewallMmsProfileDupeAction3(i["action3"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_time1"
	if _, ok := i["block-time1"]; ok {
		result["block_time1"] = flattenObjectFirewallMmsProfileDupeBlockTime1(i["block-time1"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_time2"
	if _, ok := i["block-time2"]; ok {
		result["block_time2"] = flattenObjectFirewallMmsProfileDupeBlockTime2(i["block-time2"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_time3"
	if _, ok := i["block-time3"]; ok {
		result["block_time3"] = flattenObjectFirewallMmsProfileDupeBlockTime3(i["block-time3"], d, pre_append)
	}

	pre_append = pre + ".0." + "limit1"
	if _, ok := i["limit1"]; ok {
		result["limit1"] = flattenObjectFirewallMmsProfileDupeLimit1(i["limit1"], d, pre_append)
	}

	pre_append = pre + ".0." + "limit2"
	if _, ok := i["limit2"]; ok {
		result["limit2"] = flattenObjectFirewallMmsProfileDupeLimit2(i["limit2"], d, pre_append)
	}

	pre_append = pre + ".0." + "limit3"
	if _, ok := i["limit3"]; ok {
		result["limit3"] = flattenObjectFirewallMmsProfileDupeLimit3(i["limit3"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectFirewallMmsProfileDupeProtocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "status1"
	if _, ok := i["status1"]; ok {
		result["status1"] = flattenObjectFirewallMmsProfileDupeStatus1(i["status1"], d, pre_append)
	}

	pre_append = pre + ".0." + "status2"
	if _, ok := i["status2"]; ok {
		result["status2"] = flattenObjectFirewallMmsProfileDupeStatus2(i["status2"], d, pre_append)
	}

	pre_append = pre + ".0." + "status3"
	if _, ok := i["status3"]; ok {
		result["status3"] = flattenObjectFirewallMmsProfileDupeStatus3(i["status3"], d, pre_append)
	}

	pre_append = pre + ".0." + "window1"
	if _, ok := i["window1"]; ok {
		result["window1"] = flattenObjectFirewallMmsProfileDupeWindow1(i["window1"], d, pre_append)
	}

	pre_append = pre + ".0." + "window2"
	if _, ok := i["window2"]; ok {
		result["window2"] = flattenObjectFirewallMmsProfileDupeWindow2(i["window2"], d, pre_append)
	}

	pre_append = pre + ".0." + "window3"
	if _, ok := i["window3"]; ok {
		result["window3"] = flattenObjectFirewallMmsProfileDupeWindow3(i["window3"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallMmsProfileDupeAction1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileDupeAction2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileDupeAction3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileDupeBlockTime1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeBlockTime2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeBlockTime3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeLimit1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeLimit2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeLimit3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeStatus1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeStatus2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeStatus3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeWindow1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeWindow2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileDupeWindow3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileExtendedUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFlood(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action1"
	if _, ok := i["action1"]; ok {
		result["action1"] = flattenObjectFirewallMmsProfileFloodAction1(i["action1"], d, pre_append)
	}

	pre_append = pre + ".0." + "action2"
	if _, ok := i["action2"]; ok {
		result["action2"] = flattenObjectFirewallMmsProfileFloodAction2(i["action2"], d, pre_append)
	}

	pre_append = pre + ".0." + "action3"
	if _, ok := i["action3"]; ok {
		result["action3"] = flattenObjectFirewallMmsProfileFloodAction3(i["action3"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_time1"
	if _, ok := i["block-time1"]; ok {
		result["block_time1"] = flattenObjectFirewallMmsProfileFloodBlockTime1(i["block-time1"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_time2"
	if _, ok := i["block-time2"]; ok {
		result["block_time2"] = flattenObjectFirewallMmsProfileFloodBlockTime2(i["block-time2"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_time3"
	if _, ok := i["block-time3"]; ok {
		result["block_time3"] = flattenObjectFirewallMmsProfileFloodBlockTime3(i["block-time3"], d, pre_append)
	}

	pre_append = pre + ".0." + "limit1"
	if _, ok := i["limit1"]; ok {
		result["limit1"] = flattenObjectFirewallMmsProfileFloodLimit1(i["limit1"], d, pre_append)
	}

	pre_append = pre + ".0." + "limit2"
	if _, ok := i["limit2"]; ok {
		result["limit2"] = flattenObjectFirewallMmsProfileFloodLimit2(i["limit2"], d, pre_append)
	}

	pre_append = pre + ".0." + "limit3"
	if _, ok := i["limit3"]; ok {
		result["limit3"] = flattenObjectFirewallMmsProfileFloodLimit3(i["limit3"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectFirewallMmsProfileFloodProtocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "status1"
	if _, ok := i["status1"]; ok {
		result["status1"] = flattenObjectFirewallMmsProfileFloodStatus1(i["status1"], d, pre_append)
	}

	pre_append = pre + ".0." + "status2"
	if _, ok := i["status2"]; ok {
		result["status2"] = flattenObjectFirewallMmsProfileFloodStatus2(i["status2"], d, pre_append)
	}

	pre_append = pre + ".0." + "status3"
	if _, ok := i["status3"]; ok {
		result["status3"] = flattenObjectFirewallMmsProfileFloodStatus3(i["status3"], d, pre_append)
	}

	pre_append = pre + ".0." + "window1"
	if _, ok := i["window1"]; ok {
		result["window1"] = flattenObjectFirewallMmsProfileFloodWindow1(i["window1"], d, pre_append)
	}

	pre_append = pre + ".0." + "window2"
	if _, ok := i["window2"]; ok {
		result["window2"] = flattenObjectFirewallMmsProfileFloodWindow2(i["window2"], d, pre_append)
	}

	pre_append = pre + ".0." + "window3"
	if _, ok := i["window3"]; ok {
		result["window3"] = flattenObjectFirewallMmsProfileFloodWindow3(i["window3"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallMmsProfileFloodAction1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileFloodAction2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileFloodAction3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileFloodBlockTime1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodBlockTime2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodBlockTime3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodLimit1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodLimit2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodLimit3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodStatus1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodStatus2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodStatus3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodWindow1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodWindow2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileFloodWindow3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileMm1AddrHdr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1AddrSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1ConvertHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1OutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1RetrDupe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1RetrieveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1Comfortamount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1Comfortinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm1Oversizelimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileMm3OutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm3Oversizelimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileMm4OutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm4Oversizelimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileMm7AddrHdr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7AddrSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7ConvertHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7OutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7Comfortamount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7Comfortinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMm7Oversizelimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsAntispamMassLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsAvBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsAvOversizeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsAvVirusLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsCarrierEndpointFilterLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsChecksumLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsChecksumTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsNotificationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsWebContentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileMmsbwordthreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotifMsisdn(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn"
		if _, ok := i["msisdn"]; ok {
			v := flattenObjectFirewallMmsProfileNotifMsisdnMsisdn(i["msisdn"], d, pre_append)
			tmp["msisdn"] = fortiAPISubPartPatch(v, "ObjectFirewallMmsProfile-NotifMsisdn-Msisdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			v := flattenObjectFirewallMmsProfileNotifMsisdnThreshold(i["threshold"], d, pre_append)
			tmp["threshold"] = fortiAPISubPartPatch(v, "ObjectFirewallMmsProfile-NotifMsisdn-Threshold")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallMmsProfileNotifMsisdnMsisdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotifMsisdnThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileNotification(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert_int"
	if _, ok := i["alert-int"]; ok {
		result["alert_int"] = flattenObjectFirewallMmsProfileNotificationAlertInt(i["alert-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "alert_int_mode"
	if _, ok := i["alert-int-mode"]; ok {
		result["alert_int_mode"] = flattenObjectFirewallMmsProfileNotificationAlertIntMode(i["alert-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "alert_src_msisdn"
	if _, ok := i["alert-src-msisdn"]; ok {
		result["alert_src_msisdn"] = flattenObjectFirewallMmsProfileNotificationAlertSrcMsisdn(i["alert-src-msisdn"], d, pre_append)
	}

	pre_append = pre + ".0." + "alert_status"
	if _, ok := i["alert-status"]; ok {
		result["alert_status"] = flattenObjectFirewallMmsProfileNotificationAlertStatus(i["alert-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_int"
	if _, ok := i["bword-int"]; ok {
		result["bword_int"] = flattenObjectFirewallMmsProfileNotificationBwordInt(i["bword-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_int_mode"
	if _, ok := i["bword-int-mode"]; ok {
		result["bword_int_mode"] = flattenObjectFirewallMmsProfileNotificationBwordIntMode(i["bword-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_status"
	if _, ok := i["bword-status"]; ok {
		result["bword_status"] = flattenObjectFirewallMmsProfileNotificationBwordStatus(i["bword-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "carrier_endpoint_bwl_int"
	if _, ok := i["carrier-endpoint-bwl-int"]; ok {
		result["carrier_endpoint_bwl_int"] = flattenObjectFirewallMmsProfileNotificationCarrierEndpointBwlInt(i["carrier-endpoint-bwl-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "carrier_endpoint_bwl_int_mode"
	if _, ok := i["carrier-endpoint-bwl-int-mode"]; ok {
		result["carrier_endpoint_bwl_int_mode"] = flattenObjectFirewallMmsProfileNotificationCarrierEndpointBwlIntMode(i["carrier-endpoint-bwl-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "carrier_endpoint_bwl_status"
	if _, ok := i["carrier-endpoint-bwl-status"]; ok {
		result["carrier_endpoint_bwl_status"] = flattenObjectFirewallMmsProfileNotificationCarrierEndpointBwlStatus(i["carrier-endpoint-bwl-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "days_allowed"
	if _, ok := i["days-allowed"]; ok {
		result["days_allowed"] = flattenObjectFirewallMmsProfileNotificationDaysAllowed(i["days-allowed"], d, pre_append)
	}

	pre_append = pre + ".0." + "detect_server"
	if _, ok := i["detect-server"]; ok {
		result["detect_server"] = flattenObjectFirewallMmsProfileNotificationDetectServer(i["detect-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "dupe_int"
	if _, ok := i["dupe-int"]; ok {
		result["dupe_int"] = flattenObjectFirewallMmsProfileNotificationDupeInt(i["dupe-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "dupe_int_mode"
	if _, ok := i["dupe-int-mode"]; ok {
		result["dupe_int_mode"] = flattenObjectFirewallMmsProfileNotificationDupeIntMode(i["dupe-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "dupe_status"
	if _, ok := i["dupe-status"]; ok {
		result["dupe_status"] = flattenObjectFirewallMmsProfileNotificationDupeStatus(i["dupe-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_block_int"
	if _, ok := i["file-block-int"]; ok {
		result["file_block_int"] = flattenObjectFirewallMmsProfileNotificationFileBlockInt(i["file-block-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_block_int_mode"
	if _, ok := i["file-block-int-mode"]; ok {
		result["file_block_int_mode"] = flattenObjectFirewallMmsProfileNotificationFileBlockIntMode(i["file-block-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_block_status"
	if _, ok := i["file-block-status"]; ok {
		result["file_block_status"] = flattenObjectFirewallMmsProfileNotificationFileBlockStatus(i["file-block-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "flood_int"
	if _, ok := i["flood-int"]; ok {
		result["flood_int"] = flattenObjectFirewallMmsProfileNotificationFloodInt(i["flood-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "flood_int_mode"
	if _, ok := i["flood-int-mode"]; ok {
		result["flood_int_mode"] = flattenObjectFirewallMmsProfileNotificationFloodIntMode(i["flood-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "flood_status"
	if _, ok := i["flood-status"]; ok {
		result["flood_status"] = flattenObjectFirewallMmsProfileNotificationFloodStatus(i["flood-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "from_in_header"
	if _, ok := i["from-in-header"]; ok {
		result["from_in_header"] = flattenObjectFirewallMmsProfileNotificationFromInHeader(i["from-in-header"], d, pre_append)
	}

	pre_append = pre + ".0." + "mms_checksum_int"
	if _, ok := i["mms-checksum-int"]; ok {
		result["mms_checksum_int"] = flattenObjectFirewallMmsProfileNotificationMmsChecksumInt(i["mms-checksum-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "mms_checksum_int_mode"
	if _, ok := i["mms-checksum-int-mode"]; ok {
		result["mms_checksum_int_mode"] = flattenObjectFirewallMmsProfileNotificationMmsChecksumIntMode(i["mms-checksum-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "mms_checksum_status"
	if _, ok := i["mms-checksum-status"]; ok {
		result["mms_checksum_status"] = flattenObjectFirewallMmsProfileNotificationMmsChecksumStatus(i["mms-checksum-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "mmsc_hostname"
	if _, ok := i["mmsc-hostname"]; ok {
		result["mmsc_hostname"] = flattenObjectFirewallMmsProfileNotificationMmscHostname(i["mmsc-hostname"], d, pre_append)
	}

	pre_append = pre + ".0." + "mmsc_password"
	if _, ok := i["mmsc-password"]; ok {
		result["mmsc_password"] = flattenObjectFirewallMmsProfileNotificationMmscPassword(i["mmsc-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "mmsc_port"
	if _, ok := i["mmsc-port"]; ok {
		result["mmsc_port"] = flattenObjectFirewallMmsProfileNotificationMmscPort(i["mmsc-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "mmsc_url"
	if _, ok := i["mmsc-url"]; ok {
		result["mmsc_url"] = flattenObjectFirewallMmsProfileNotificationMmscUrl(i["mmsc-url"], d, pre_append)
	}

	pre_append = pre + ".0." + "mmsc_username"
	if _, ok := i["mmsc-username"]; ok {
		result["mmsc_username"] = flattenObjectFirewallMmsProfileNotificationMmscUsername(i["mmsc-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "msg_protocol"
	if _, ok := i["msg-protocol"]; ok {
		result["msg_protocol"] = flattenObjectFirewallMmsProfileNotificationMsgProtocol(i["msg-protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "msg_type"
	if _, ok := i["msg-type"]; ok {
		result["msg_type"] = flattenObjectFirewallMmsProfileNotificationMsgType(i["msg-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectFirewallMmsProfileNotificationProtocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_limit"
	if _, ok := i["rate-limit"]; ok {
		result["rate_limit"] = flattenObjectFirewallMmsProfileNotificationRateLimit(i["rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "tod_window_duration"
	if _, ok := i["tod-window-duration"]; ok {
		result["tod_window_duration"] = flattenObjectFirewallMmsProfileNotificationTodWindowDuration(i["tod-window-duration"], d, pre_append)
	}

	pre_append = pre + ".0." + "tod_window_end"
	if _, ok := i["tod-window-end"]; ok {
		result["tod_window_end"] = flattenObjectFirewallMmsProfileNotificationTodWindowEnd(i["tod-window-end"], d, pre_append)
	}

	pre_append = pre + ".0." + "tod_window_start"
	if _, ok := i["tod-window-start"]; ok {
		result["tod_window_start"] = flattenObjectFirewallMmsProfileNotificationTodWindowStart(i["tod-window-start"], d, pre_append)
	}

	pre_append = pre + ".0." + "user_domain"
	if _, ok := i["user-domain"]; ok {
		result["user_domain"] = flattenObjectFirewallMmsProfileNotificationUserDomain(i["user-domain"], d, pre_append)
	}

	pre_append = pre + ".0." + "vas_id"
	if _, ok := i["vas-id"]; ok {
		result["vas_id"] = flattenObjectFirewallMmsProfileNotificationVasId(i["vas-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "vasp_id"
	if _, ok := i["vasp-id"]; ok {
		result["vasp_id"] = flattenObjectFirewallMmsProfileNotificationVaspId(i["vasp-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_int"
	if _, ok := i["virus-int"]; ok {
		result["virus_int"] = flattenObjectFirewallMmsProfileNotificationVirusInt(i["virus-int"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_int_mode"
	if _, ok := i["virus-int-mode"]; ok {
		result["virus_int_mode"] = flattenObjectFirewallMmsProfileNotificationVirusIntMode(i["virus-int-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "virus_status"
	if _, ok := i["virus-status"]; ok {
		result["virus_status"] = flattenObjectFirewallMmsProfileNotificationVirusStatus(i["virus-status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallMmsProfileNotificationAlertInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationAlertIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationAlertSrcMsisdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationAlertStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationBwordInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationBwordIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationBwordStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationCarrierEndpointBwlInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationCarrierEndpointBwlIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationCarrierEndpointBwlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationDaysAllowed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileNotificationDetectServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationDupeInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationDupeIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationDupeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFileBlockInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFileBlockIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFileBlockStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFloodInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFloodIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFloodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationFromInHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmsChecksumInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmsChecksumIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmsChecksumStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmscHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmscPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallMmsProfileNotificationMmscPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmscUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMmscUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMsgProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationTodWindowDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationTodWindowEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationTodWindowStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationUserDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationVasId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationVaspId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationVirusInt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationVirusIntMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileNotificationVirusStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenObjectFirewallMmsProfileOutbreakPreventionExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := i["ftgd-service"]; ok {
		result["ftgd_service"] = flattenObjectFirewallMmsProfileOutbreakPreventionFtgdService(i["ftgd-service"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallMmsProfileOutbreakPreventionExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileOutbreakPreventionFtgdService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileRemoveBlockedConstLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallMmsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("avnotificationtable", flattenObjectFirewallMmsProfileAvnotificationtable(o["avnotificationtable"], d, "avnotificationtable")); err != nil {
		if vv, ok := fortiAPIPatch(o["avnotificationtable"], "ObjectFirewallMmsProfile-Avnotificationtable"); ok {
			if err = d.Set("avnotificationtable", vv); err != nil {
				return fmt.Errorf("Error reading avnotificationtable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading avnotificationtable: %v", err)
		}
	}

	if err = d.Set("bwordtable", flattenObjectFirewallMmsProfileBwordtable(o["bwordtable"], d, "bwordtable")); err != nil {
		if vv, ok := fortiAPIPatch(o["bwordtable"], "ObjectFirewallMmsProfile-Bwordtable"); ok {
			if err = d.Set("bwordtable", vv); err != nil {
				return fmt.Errorf("Error reading bwordtable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bwordtable: %v", err)
		}
	}

	if err = d.Set("carrier_endpoint_prefix", flattenObjectFirewallMmsProfileCarrierEndpointPrefix(o["carrier-endpoint-prefix"], d, "carrier_endpoint_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier-endpoint-prefix"], "ObjectFirewallMmsProfile-CarrierEndpointPrefix"); ok {
			if err = d.Set("carrier_endpoint_prefix", vv); err != nil {
				return fmt.Errorf("Error reading carrier_endpoint_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier_endpoint_prefix: %v", err)
		}
	}

	if err = d.Set("carrier_endpoint_prefix_range_max", flattenObjectFirewallMmsProfileCarrierEndpointPrefixRangeMax(o["carrier-endpoint-prefix-range-max"], d, "carrier_endpoint_prefix_range_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier-endpoint-prefix-range-max"], "ObjectFirewallMmsProfile-CarrierEndpointPrefixRangeMax"); ok {
			if err = d.Set("carrier_endpoint_prefix_range_max", vv); err != nil {
				return fmt.Errorf("Error reading carrier_endpoint_prefix_range_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier_endpoint_prefix_range_max: %v", err)
		}
	}

	if err = d.Set("carrier_endpoint_prefix_range_min", flattenObjectFirewallMmsProfileCarrierEndpointPrefixRangeMin(o["carrier-endpoint-prefix-range-min"], d, "carrier_endpoint_prefix_range_min")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier-endpoint-prefix-range-min"], "ObjectFirewallMmsProfile-CarrierEndpointPrefixRangeMin"); ok {
			if err = d.Set("carrier_endpoint_prefix_range_min", vv); err != nil {
				return fmt.Errorf("Error reading carrier_endpoint_prefix_range_min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier_endpoint_prefix_range_min: %v", err)
		}
	}

	if err = d.Set("carrier_endpoint_prefix_string", flattenObjectFirewallMmsProfileCarrierEndpointPrefixString(o["carrier-endpoint-prefix-string"], d, "carrier_endpoint_prefix_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier-endpoint-prefix-string"], "ObjectFirewallMmsProfile-CarrierEndpointPrefixString"); ok {
			if err = d.Set("carrier_endpoint_prefix_string", vv); err != nil {
				return fmt.Errorf("Error reading carrier_endpoint_prefix_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier_endpoint_prefix_string: %v", err)
		}
	}

	if err = d.Set("carrierendpointbwltable", flattenObjectFirewallMmsProfileCarrierendpointbwltable(o["carrierendpointbwltable"], d, "carrierendpointbwltable")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrierendpointbwltable"], "ObjectFirewallMmsProfile-Carrierendpointbwltable"); ok {
			if err = d.Set("carrierendpointbwltable", vv); err != nil {
				return fmt.Errorf("Error reading carrierendpointbwltable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrierendpointbwltable: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallMmsProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallMmsProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dupe", flattenObjectFirewallMmsProfileDupe(o["dupe"], d, "dupe")); err != nil {
			if vv, ok := fortiAPIPatch(o["dupe"], "ObjectFirewallMmsProfile-Dupe"); ok {
				if err = d.Set("dupe", vv); err != nil {
					return fmt.Errorf("Error reading dupe: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dupe: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dupe"); ok {
			if err = d.Set("dupe", flattenObjectFirewallMmsProfileDupe(o["dupe"], d, "dupe")); err != nil {
				if vv, ok := fortiAPIPatch(o["dupe"], "ObjectFirewallMmsProfile-Dupe"); ok {
					if err = d.Set("dupe", vv); err != nil {
						return fmt.Errorf("Error reading dupe: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dupe: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_utm_log", flattenObjectFirewallMmsProfileExtendedUtmLog(o["extended-utm-log"], d, "extended_utm_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-utm-log"], "ObjectFirewallMmsProfile-ExtendedUtmLog"); ok {
			if err = d.Set("extended_utm_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_utm_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_utm_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("flood", flattenObjectFirewallMmsProfileFlood(o["flood"], d, "flood")); err != nil {
			if vv, ok := fortiAPIPatch(o["flood"], "ObjectFirewallMmsProfile-Flood"); ok {
				if err = d.Set("flood", vv); err != nil {
					return fmt.Errorf("Error reading flood: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading flood: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("flood"); ok {
			if err = d.Set("flood", flattenObjectFirewallMmsProfileFlood(o["flood"], d, "flood")); err != nil {
				if vv, ok := fortiAPIPatch(o["flood"], "ObjectFirewallMmsProfile-Flood"); ok {
					if err = d.Set("flood", vv); err != nil {
						return fmt.Errorf("Error reading flood: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading flood: %v", err)
				}
			}
		}
	}

	if err = d.Set("mm1", flattenObjectFirewallMmsProfileMm1(o["mm1"], d, "mm1")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1"], "ObjectFirewallMmsProfile-Mm1"); ok {
			if err = d.Set("mm1", vv); err != nil {
				return fmt.Errorf("Error reading mm1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1: %v", err)
		}
	}

	if err = d.Set("mm1_addr_hdr", flattenObjectFirewallMmsProfileMm1AddrHdr(o["mm1-addr-hdr"], d, "mm1_addr_hdr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1-addr-hdr"], "ObjectFirewallMmsProfile-Mm1AddrHdr"); ok {
			if err = d.Set("mm1_addr_hdr", vv); err != nil {
				return fmt.Errorf("Error reading mm1_addr_hdr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1_addr_hdr: %v", err)
		}
	}

	if err = d.Set("mm1_addr_source", flattenObjectFirewallMmsProfileMm1AddrSource(o["mm1-addr-source"], d, "mm1_addr_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1-addr-source"], "ObjectFirewallMmsProfile-Mm1AddrSource"); ok {
			if err = d.Set("mm1_addr_source", vv); err != nil {
				return fmt.Errorf("Error reading mm1_addr_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1_addr_source: %v", err)
		}
	}

	if err = d.Set("mm1_convert_hex", flattenObjectFirewallMmsProfileMm1ConvertHex(o["mm1-convert-hex"], d, "mm1_convert_hex")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1-convert-hex"], "ObjectFirewallMmsProfile-Mm1ConvertHex"); ok {
			if err = d.Set("mm1_convert_hex", vv); err != nil {
				return fmt.Errorf("Error reading mm1_convert_hex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1_convert_hex: %v", err)
		}
	}

	if err = d.Set("mm1_outbreak_prevention", flattenObjectFirewallMmsProfileMm1OutbreakPrevention(o["mm1-outbreak-prevention"], d, "mm1_outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1-outbreak-prevention"], "ObjectFirewallMmsProfile-Mm1OutbreakPrevention"); ok {
			if err = d.Set("mm1_outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading mm1_outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1_outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("mm1_retr_dupe", flattenObjectFirewallMmsProfileMm1RetrDupe(o["mm1-retr-dupe"], d, "mm1_retr_dupe")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1-retr-dupe"], "ObjectFirewallMmsProfile-Mm1RetrDupe"); ok {
			if err = d.Set("mm1_retr_dupe", vv); err != nil {
				return fmt.Errorf("Error reading mm1_retr_dupe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1_retr_dupe: %v", err)
		}
	}

	if err = d.Set("mm1_retrieve_scan", flattenObjectFirewallMmsProfileMm1RetrieveScan(o["mm1-retrieve-scan"], d, "mm1_retrieve_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1-retrieve-scan"], "ObjectFirewallMmsProfile-Mm1RetrieveScan"); ok {
			if err = d.Set("mm1_retrieve_scan", vv); err != nil {
				return fmt.Errorf("Error reading mm1_retrieve_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1_retrieve_scan: %v", err)
		}
	}

	if err = d.Set("mm1comfortamount", flattenObjectFirewallMmsProfileMm1Comfortamount(o["mm1comfortamount"], d, "mm1comfortamount")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1comfortamount"], "ObjectFirewallMmsProfile-Mm1Comfortamount"); ok {
			if err = d.Set("mm1comfortamount", vv); err != nil {
				return fmt.Errorf("Error reading mm1comfortamount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1comfortamount: %v", err)
		}
	}

	if err = d.Set("mm1comfortinterval", flattenObjectFirewallMmsProfileMm1Comfortinterval(o["mm1comfortinterval"], d, "mm1comfortinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1comfortinterval"], "ObjectFirewallMmsProfile-Mm1Comfortinterval"); ok {
			if err = d.Set("mm1comfortinterval", vv); err != nil {
				return fmt.Errorf("Error reading mm1comfortinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1comfortinterval: %v", err)
		}
	}

	if err = d.Set("mm1oversizelimit", flattenObjectFirewallMmsProfileMm1Oversizelimit(o["mm1oversizelimit"], d, "mm1oversizelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm1oversizelimit"], "ObjectFirewallMmsProfile-Mm1Oversizelimit"); ok {
			if err = d.Set("mm1oversizelimit", vv); err != nil {
				return fmt.Errorf("Error reading mm1oversizelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm1oversizelimit: %v", err)
		}
	}

	if err = d.Set("mm3", flattenObjectFirewallMmsProfileMm3(o["mm3"], d, "mm3")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm3"], "ObjectFirewallMmsProfile-Mm3"); ok {
			if err = d.Set("mm3", vv); err != nil {
				return fmt.Errorf("Error reading mm3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm3: %v", err)
		}
	}

	if err = d.Set("mm3_outbreak_prevention", flattenObjectFirewallMmsProfileMm3OutbreakPrevention(o["mm3-outbreak-prevention"], d, "mm3_outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm3-outbreak-prevention"], "ObjectFirewallMmsProfile-Mm3OutbreakPrevention"); ok {
			if err = d.Set("mm3_outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading mm3_outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm3_outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("mm3oversizelimit", flattenObjectFirewallMmsProfileMm3Oversizelimit(o["mm3oversizelimit"], d, "mm3oversizelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm3oversizelimit"], "ObjectFirewallMmsProfile-Mm3Oversizelimit"); ok {
			if err = d.Set("mm3oversizelimit", vv); err != nil {
				return fmt.Errorf("Error reading mm3oversizelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm3oversizelimit: %v", err)
		}
	}

	if err = d.Set("mm4", flattenObjectFirewallMmsProfileMm4(o["mm4"], d, "mm4")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm4"], "ObjectFirewallMmsProfile-Mm4"); ok {
			if err = d.Set("mm4", vv); err != nil {
				return fmt.Errorf("Error reading mm4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm4: %v", err)
		}
	}

	if err = d.Set("mm4_outbreak_prevention", flattenObjectFirewallMmsProfileMm4OutbreakPrevention(o["mm4-outbreak-prevention"], d, "mm4_outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm4-outbreak-prevention"], "ObjectFirewallMmsProfile-Mm4OutbreakPrevention"); ok {
			if err = d.Set("mm4_outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading mm4_outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm4_outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("mm4oversizelimit", flattenObjectFirewallMmsProfileMm4Oversizelimit(o["mm4oversizelimit"], d, "mm4oversizelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm4oversizelimit"], "ObjectFirewallMmsProfile-Mm4Oversizelimit"); ok {
			if err = d.Set("mm4oversizelimit", vv); err != nil {
				return fmt.Errorf("Error reading mm4oversizelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm4oversizelimit: %v", err)
		}
	}

	if err = d.Set("mm7", flattenObjectFirewallMmsProfileMm7(o["mm7"], d, "mm7")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7"], "ObjectFirewallMmsProfile-Mm7"); ok {
			if err = d.Set("mm7", vv); err != nil {
				return fmt.Errorf("Error reading mm7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7: %v", err)
		}
	}

	if err = d.Set("mm7_addr_hdr", flattenObjectFirewallMmsProfileMm7AddrHdr(o["mm7-addr-hdr"], d, "mm7_addr_hdr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7-addr-hdr"], "ObjectFirewallMmsProfile-Mm7AddrHdr"); ok {
			if err = d.Set("mm7_addr_hdr", vv); err != nil {
				return fmt.Errorf("Error reading mm7_addr_hdr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7_addr_hdr: %v", err)
		}
	}

	if err = d.Set("mm7_addr_source", flattenObjectFirewallMmsProfileMm7AddrSource(o["mm7-addr-source"], d, "mm7_addr_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7-addr-source"], "ObjectFirewallMmsProfile-Mm7AddrSource"); ok {
			if err = d.Set("mm7_addr_source", vv); err != nil {
				return fmt.Errorf("Error reading mm7_addr_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7_addr_source: %v", err)
		}
	}

	if err = d.Set("mm7_convert_hex", flattenObjectFirewallMmsProfileMm7ConvertHex(o["mm7-convert-hex"], d, "mm7_convert_hex")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7-convert-hex"], "ObjectFirewallMmsProfile-Mm7ConvertHex"); ok {
			if err = d.Set("mm7_convert_hex", vv); err != nil {
				return fmt.Errorf("Error reading mm7_convert_hex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7_convert_hex: %v", err)
		}
	}

	if err = d.Set("mm7_outbreak_prevention", flattenObjectFirewallMmsProfileMm7OutbreakPrevention(o["mm7-outbreak-prevention"], d, "mm7_outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7-outbreak-prevention"], "ObjectFirewallMmsProfile-Mm7OutbreakPrevention"); ok {
			if err = d.Set("mm7_outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading mm7_outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7_outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("mm7comfortamount", flattenObjectFirewallMmsProfileMm7Comfortamount(o["mm7comfortamount"], d, "mm7comfortamount")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7comfortamount"], "ObjectFirewallMmsProfile-Mm7Comfortamount"); ok {
			if err = d.Set("mm7comfortamount", vv); err != nil {
				return fmt.Errorf("Error reading mm7comfortamount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7comfortamount: %v", err)
		}
	}

	if err = d.Set("mm7comfortinterval", flattenObjectFirewallMmsProfileMm7Comfortinterval(o["mm7comfortinterval"], d, "mm7comfortinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7comfortinterval"], "ObjectFirewallMmsProfile-Mm7Comfortinterval"); ok {
			if err = d.Set("mm7comfortinterval", vv); err != nil {
				return fmt.Errorf("Error reading mm7comfortinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7comfortinterval: %v", err)
		}
	}

	if err = d.Set("mm7oversizelimit", flattenObjectFirewallMmsProfileMm7Oversizelimit(o["mm7oversizelimit"], d, "mm7oversizelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm7oversizelimit"], "ObjectFirewallMmsProfile-Mm7Oversizelimit"); ok {
			if err = d.Set("mm7oversizelimit", vv); err != nil {
				return fmt.Errorf("Error reading mm7oversizelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm7oversizelimit: %v", err)
		}
	}

	if err = d.Set("mms_antispam_mass_log", flattenObjectFirewallMmsProfileMmsAntispamMassLog(o["mms-antispam-mass-log"], d, "mms_antispam_mass_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-antispam-mass-log"], "ObjectFirewallMmsProfile-MmsAntispamMassLog"); ok {
			if err = d.Set("mms_antispam_mass_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_antispam_mass_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_antispam_mass_log: %v", err)
		}
	}

	if err = d.Set("mms_av_block_log", flattenObjectFirewallMmsProfileMmsAvBlockLog(o["mms-av-block-log"], d, "mms_av_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-av-block-log"], "ObjectFirewallMmsProfile-MmsAvBlockLog"); ok {
			if err = d.Set("mms_av_block_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_av_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_av_block_log: %v", err)
		}
	}

	if err = d.Set("mms_av_oversize_log", flattenObjectFirewallMmsProfileMmsAvOversizeLog(o["mms-av-oversize-log"], d, "mms_av_oversize_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-av-oversize-log"], "ObjectFirewallMmsProfile-MmsAvOversizeLog"); ok {
			if err = d.Set("mms_av_oversize_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_av_oversize_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_av_oversize_log: %v", err)
		}
	}

	if err = d.Set("mms_av_virus_log", flattenObjectFirewallMmsProfileMmsAvVirusLog(o["mms-av-virus-log"], d, "mms_av_virus_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-av-virus-log"], "ObjectFirewallMmsProfile-MmsAvVirusLog"); ok {
			if err = d.Set("mms_av_virus_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_av_virus_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_av_virus_log: %v", err)
		}
	}

	if err = d.Set("mms_carrier_endpoint_filter_log", flattenObjectFirewallMmsProfileMmsCarrierEndpointFilterLog(o["mms-carrier-endpoint-filter-log"], d, "mms_carrier_endpoint_filter_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-carrier-endpoint-filter-log"], "ObjectFirewallMmsProfile-MmsCarrierEndpointFilterLog"); ok {
			if err = d.Set("mms_carrier_endpoint_filter_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_carrier_endpoint_filter_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_carrier_endpoint_filter_log: %v", err)
		}
	}

	if err = d.Set("mms_checksum_log", flattenObjectFirewallMmsProfileMmsChecksumLog(o["mms-checksum-log"], d, "mms_checksum_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-checksum-log"], "ObjectFirewallMmsProfile-MmsChecksumLog"); ok {
			if err = d.Set("mms_checksum_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_checksum_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_checksum_log: %v", err)
		}
	}

	if err = d.Set("mms_checksum_table", flattenObjectFirewallMmsProfileMmsChecksumTable(o["mms-checksum-table"], d, "mms_checksum_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-checksum-table"], "ObjectFirewallMmsProfile-MmsChecksumTable"); ok {
			if err = d.Set("mms_checksum_table", vv); err != nil {
				return fmt.Errorf("Error reading mms_checksum_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_checksum_table: %v", err)
		}
	}

	if err = d.Set("mms_notification_log", flattenObjectFirewallMmsProfileMmsNotificationLog(o["mms-notification-log"], d, "mms_notification_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-notification-log"], "ObjectFirewallMmsProfile-MmsNotificationLog"); ok {
			if err = d.Set("mms_notification_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_notification_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_notification_log: %v", err)
		}
	}

	if err = d.Set("mms_web_content_log", flattenObjectFirewallMmsProfileMmsWebContentLog(o["mms-web-content-log"], d, "mms_web_content_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-web-content-log"], "ObjectFirewallMmsProfile-MmsWebContentLog"); ok {
			if err = d.Set("mms_web_content_log", vv); err != nil {
				return fmt.Errorf("Error reading mms_web_content_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_web_content_log: %v", err)
		}
	}

	if err = d.Set("mmsbwordthreshold", flattenObjectFirewallMmsProfileMmsbwordthreshold(o["mmsbwordthreshold"], d, "mmsbwordthreshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["mmsbwordthreshold"], "ObjectFirewallMmsProfile-Mmsbwordthreshold"); ok {
			if err = d.Set("mmsbwordthreshold", vv); err != nil {
				return fmt.Errorf("Error reading mmsbwordthreshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mmsbwordthreshold: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallMmsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallMmsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("notif_msisdn", flattenObjectFirewallMmsProfileNotifMsisdn(o["notif-msisdn"], d, "notif_msisdn")); err != nil {
			if vv, ok := fortiAPIPatch(o["notif-msisdn"], "ObjectFirewallMmsProfile-NotifMsisdn"); ok {
				if err = d.Set("notif_msisdn", vv); err != nil {
					return fmt.Errorf("Error reading notif_msisdn: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading notif_msisdn: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("notif_msisdn"); ok {
			if err = d.Set("notif_msisdn", flattenObjectFirewallMmsProfileNotifMsisdn(o["notif-msisdn"], d, "notif_msisdn")); err != nil {
				if vv, ok := fortiAPIPatch(o["notif-msisdn"], "ObjectFirewallMmsProfile-NotifMsisdn"); ok {
					if err = d.Set("notif_msisdn", vv); err != nil {
						return fmt.Errorf("Error reading notif_msisdn: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading notif_msisdn: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("notification", flattenObjectFirewallMmsProfileNotification(o["notification"], d, "notification")); err != nil {
			if vv, ok := fortiAPIPatch(o["notification"], "ObjectFirewallMmsProfile-Notification"); ok {
				if err = d.Set("notification", vv); err != nil {
					return fmt.Errorf("Error reading notification: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading notification: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("notification"); ok {
			if err = d.Set("notification", flattenObjectFirewallMmsProfileNotification(o["notification"], d, "notification")); err != nil {
				if vv, ok := fortiAPIPatch(o["notification"], "ObjectFirewallMmsProfile-Notification"); ok {
					if err = d.Set("notification", vv); err != nil {
						return fmt.Errorf("Error reading notification: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading notification: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("outbreak_prevention", flattenObjectFirewallMmsProfileOutbreakPrevention(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
			if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "ObjectFirewallMmsProfile-OutbreakPrevention"); ok {
				if err = d.Set("outbreak_prevention", vv); err != nil {
					return fmt.Errorf("Error reading outbreak_prevention: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading outbreak_prevention: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("outbreak_prevention"); ok {
			if err = d.Set("outbreak_prevention", flattenObjectFirewallMmsProfileOutbreakPrevention(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
				if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "ObjectFirewallMmsProfile-OutbreakPrevention"); ok {
					if err = d.Set("outbreak_prevention", vv); err != nil {
						return fmt.Errorf("Error reading outbreak_prevention: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading outbreak_prevention: %v", err)
				}
			}
		}
	}

	if err = d.Set("remove_blocked_const_length", flattenObjectFirewallMmsProfileRemoveBlockedConstLength(o["remove-blocked-const-length"], d, "remove_blocked_const_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["remove-blocked-const-length"], "ObjectFirewallMmsProfile-RemoveBlockedConstLength"); ok {
			if err = d.Set("remove_blocked_const_length", vv); err != nil {
				return fmt.Errorf("Error reading remove_blocked_const_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remove_blocked_const_length: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectFirewallMmsProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectFirewallMmsProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallMmsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallMmsProfileAvnotificationtable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileBwordtable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileCarrierEndpointPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileCarrierEndpointPrefixRangeMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileCarrierEndpointPrefixRangeMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileCarrierEndpointPrefixString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileCarrierendpointbwltable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action1"], _ = expandObjectFirewallMmsProfileDupeAction1(d, i["action1"], pre_append)
	}
	pre_append = pre + ".0." + "action2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action2"], _ = expandObjectFirewallMmsProfileDupeAction2(d, i["action2"], pre_append)
	}
	pre_append = pre + ".0." + "action3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action3"], _ = expandObjectFirewallMmsProfileDupeAction3(d, i["action3"], pre_append)
	}
	pre_append = pre + ".0." + "block_time1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-time1"], _ = expandObjectFirewallMmsProfileDupeBlockTime1(d, i["block_time1"], pre_append)
	}
	pre_append = pre + ".0." + "block_time2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-time2"], _ = expandObjectFirewallMmsProfileDupeBlockTime2(d, i["block_time2"], pre_append)
	}
	pre_append = pre + ".0." + "block_time3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-time3"], _ = expandObjectFirewallMmsProfileDupeBlockTime3(d, i["block_time3"], pre_append)
	}
	pre_append = pre + ".0." + "limit1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["limit1"], _ = expandObjectFirewallMmsProfileDupeLimit1(d, i["limit1"], pre_append)
	}
	pre_append = pre + ".0." + "limit2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["limit2"], _ = expandObjectFirewallMmsProfileDupeLimit2(d, i["limit2"], pre_append)
	}
	pre_append = pre + ".0." + "limit3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["limit3"], _ = expandObjectFirewallMmsProfileDupeLimit3(d, i["limit3"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectFirewallMmsProfileDupeProtocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "status1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status1"], _ = expandObjectFirewallMmsProfileDupeStatus1(d, i["status1"], pre_append)
	}
	pre_append = pre + ".0." + "status2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status2"], _ = expandObjectFirewallMmsProfileDupeStatus2(d, i["status2"], pre_append)
	}
	pre_append = pre + ".0." + "status3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status3"], _ = expandObjectFirewallMmsProfileDupeStatus3(d, i["status3"], pre_append)
	}
	pre_append = pre + ".0." + "window1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["window1"], _ = expandObjectFirewallMmsProfileDupeWindow1(d, i["window1"], pre_append)
	}
	pre_append = pre + ".0." + "window2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["window2"], _ = expandObjectFirewallMmsProfileDupeWindow2(d, i["window2"], pre_append)
	}
	pre_append = pre + ".0." + "window3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["window3"], _ = expandObjectFirewallMmsProfileDupeWindow3(d, i["window3"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallMmsProfileDupeAction1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileDupeAction2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileDupeAction3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileDupeBlockTime1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeBlockTime2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeBlockTime3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeLimit1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeLimit2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeLimit3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeStatus1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeStatus2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeStatus3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeWindow1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeWindow2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileDupeWindow3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileExtendedUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFlood(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action1"], _ = expandObjectFirewallMmsProfileFloodAction1(d, i["action1"], pre_append)
	}
	pre_append = pre + ".0." + "action2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action2"], _ = expandObjectFirewallMmsProfileFloodAction2(d, i["action2"], pre_append)
	}
	pre_append = pre + ".0." + "action3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action3"], _ = expandObjectFirewallMmsProfileFloodAction3(d, i["action3"], pre_append)
	}
	pre_append = pre + ".0." + "block_time1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-time1"], _ = expandObjectFirewallMmsProfileFloodBlockTime1(d, i["block_time1"], pre_append)
	}
	pre_append = pre + ".0." + "block_time2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-time2"], _ = expandObjectFirewallMmsProfileFloodBlockTime2(d, i["block_time2"], pre_append)
	}
	pre_append = pre + ".0." + "block_time3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["block-time3"], _ = expandObjectFirewallMmsProfileFloodBlockTime3(d, i["block_time3"], pre_append)
	}
	pre_append = pre + ".0." + "limit1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["limit1"], _ = expandObjectFirewallMmsProfileFloodLimit1(d, i["limit1"], pre_append)
	}
	pre_append = pre + ".0." + "limit2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["limit2"], _ = expandObjectFirewallMmsProfileFloodLimit2(d, i["limit2"], pre_append)
	}
	pre_append = pre + ".0." + "limit3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["limit3"], _ = expandObjectFirewallMmsProfileFloodLimit3(d, i["limit3"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectFirewallMmsProfileFloodProtocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "status1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status1"], _ = expandObjectFirewallMmsProfileFloodStatus1(d, i["status1"], pre_append)
	}
	pre_append = pre + ".0." + "status2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status2"], _ = expandObjectFirewallMmsProfileFloodStatus2(d, i["status2"], pre_append)
	}
	pre_append = pre + ".0." + "status3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status3"], _ = expandObjectFirewallMmsProfileFloodStatus3(d, i["status3"], pre_append)
	}
	pre_append = pre + ".0." + "window1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["window1"], _ = expandObjectFirewallMmsProfileFloodWindow1(d, i["window1"], pre_append)
	}
	pre_append = pre + ".0." + "window2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["window2"], _ = expandObjectFirewallMmsProfileFloodWindow2(d, i["window2"], pre_append)
	}
	pre_append = pre + ".0." + "window3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["window3"], _ = expandObjectFirewallMmsProfileFloodWindow3(d, i["window3"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallMmsProfileFloodAction1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileFloodAction2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileFloodAction3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileFloodBlockTime1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodBlockTime2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodBlockTime3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodLimit1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodLimit2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodLimit3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodStatus1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodStatus2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodStatus3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodWindow1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodWindow2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileFloodWindow3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileMm1AddrHdr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1AddrSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1ConvertHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1OutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1RetrDupe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1RetrieveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1Comfortamount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1Comfortinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm1Oversizelimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileMm3OutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm3Oversizelimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileMm4OutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm4Oversizelimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileMm7AddrHdr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7AddrSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7ConvertHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7OutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7Comfortamount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7Comfortinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMm7Oversizelimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsAntispamMassLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsAvBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsAvOversizeLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsAvVirusLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsCarrierEndpointFilterLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsChecksumLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsChecksumTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsNotificationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsWebContentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileMmsbwordthreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotifMsisdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msisdn"], _ = expandObjectFirewallMmsProfileNotifMsisdnMsisdn(d, i["msisdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold"], _ = expandObjectFirewallMmsProfileNotifMsisdnThreshold(d, i["threshold"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallMmsProfileNotifMsisdnMsisdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotifMsisdnThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["alert-int"], _ = expandObjectFirewallMmsProfileNotificationAlertInt(d, i["alert_int"], pre_append)
	}
	pre_append = pre + ".0." + "alert_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["alert-int-mode"], _ = expandObjectFirewallMmsProfileNotificationAlertIntMode(d, i["alert_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "alert_src_msisdn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["alert-src-msisdn"], _ = expandObjectFirewallMmsProfileNotificationAlertSrcMsisdn(d, i["alert_src_msisdn"], pre_append)
	}
	pre_append = pre + ".0." + "alert_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["alert-status"], _ = expandObjectFirewallMmsProfileNotificationAlertStatus(d, i["alert_status"], pre_append)
	}
	pre_append = pre + ".0." + "bword_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-int"], _ = expandObjectFirewallMmsProfileNotificationBwordInt(d, i["bword_int"], pre_append)
	}
	pre_append = pre + ".0." + "bword_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-int-mode"], _ = expandObjectFirewallMmsProfileNotificationBwordIntMode(d, i["bword_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "bword_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bword-status"], _ = expandObjectFirewallMmsProfileNotificationBwordStatus(d, i["bword_status"], pre_append)
	}
	pre_append = pre + ".0." + "carrier_endpoint_bwl_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["carrier-endpoint-bwl-int"], _ = expandObjectFirewallMmsProfileNotificationCarrierEndpointBwlInt(d, i["carrier_endpoint_bwl_int"], pre_append)
	}
	pre_append = pre + ".0." + "carrier_endpoint_bwl_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["carrier-endpoint-bwl-int-mode"], _ = expandObjectFirewallMmsProfileNotificationCarrierEndpointBwlIntMode(d, i["carrier_endpoint_bwl_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "carrier_endpoint_bwl_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["carrier-endpoint-bwl-status"], _ = expandObjectFirewallMmsProfileNotificationCarrierEndpointBwlStatus(d, i["carrier_endpoint_bwl_status"], pre_append)
	}
	pre_append = pre + ".0." + "days_allowed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days-allowed"], _ = expandObjectFirewallMmsProfileNotificationDaysAllowed(d, i["days_allowed"], pre_append)
	}
	pre_append = pre + ".0." + "detect_server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["detect-server"], _ = expandObjectFirewallMmsProfileNotificationDetectServer(d, i["detect_server"], pre_append)
	}
	pre_append = pre + ".0." + "dupe_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dupe-int"], _ = expandObjectFirewallMmsProfileNotificationDupeInt(d, i["dupe_int"], pre_append)
	}
	pre_append = pre + ".0." + "dupe_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dupe-int-mode"], _ = expandObjectFirewallMmsProfileNotificationDupeIntMode(d, i["dupe_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "dupe_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dupe-status"], _ = expandObjectFirewallMmsProfileNotificationDupeStatus(d, i["dupe_status"], pre_append)
	}
	pre_append = pre + ".0." + "file_block_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-block-int"], _ = expandObjectFirewallMmsProfileNotificationFileBlockInt(d, i["file_block_int"], pre_append)
	}
	pre_append = pre + ".0." + "file_block_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-block-int-mode"], _ = expandObjectFirewallMmsProfileNotificationFileBlockIntMode(d, i["file_block_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "file_block_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-block-status"], _ = expandObjectFirewallMmsProfileNotificationFileBlockStatus(d, i["file_block_status"], pre_append)
	}
	pre_append = pre + ".0." + "flood_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["flood-int"], _ = expandObjectFirewallMmsProfileNotificationFloodInt(d, i["flood_int"], pre_append)
	}
	pre_append = pre + ".0." + "flood_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["flood-int-mode"], _ = expandObjectFirewallMmsProfileNotificationFloodIntMode(d, i["flood_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "flood_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["flood-status"], _ = expandObjectFirewallMmsProfileNotificationFloodStatus(d, i["flood_status"], pre_append)
	}
	pre_append = pre + ".0." + "from_in_header"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["from-in-header"], _ = expandObjectFirewallMmsProfileNotificationFromInHeader(d, i["from_in_header"], pre_append)
	}
	pre_append = pre + ".0." + "mms_checksum_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mms-checksum-int"], _ = expandObjectFirewallMmsProfileNotificationMmsChecksumInt(d, i["mms_checksum_int"], pre_append)
	}
	pre_append = pre + ".0." + "mms_checksum_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mms-checksum-int-mode"], _ = expandObjectFirewallMmsProfileNotificationMmsChecksumIntMode(d, i["mms_checksum_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "mms_checksum_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mms-checksum-status"], _ = expandObjectFirewallMmsProfileNotificationMmsChecksumStatus(d, i["mms_checksum_status"], pre_append)
	}
	pre_append = pre + ".0." + "mmsc_hostname"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mmsc-hostname"], _ = expandObjectFirewallMmsProfileNotificationMmscHostname(d, i["mmsc_hostname"], pre_append)
	}
	pre_append = pre + ".0." + "mmsc_password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mmsc-password"], _ = expandObjectFirewallMmsProfileNotificationMmscPassword(d, i["mmsc_password"], pre_append)
	}
	pre_append = pre + ".0." + "mmsc_port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mmsc-port"], _ = expandObjectFirewallMmsProfileNotificationMmscPort(d, i["mmsc_port"], pre_append)
	}
	pre_append = pre + ".0." + "mmsc_url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mmsc-url"], _ = expandObjectFirewallMmsProfileNotificationMmscUrl(d, i["mmsc_url"], pre_append)
	}
	pre_append = pre + ".0." + "mmsc_username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mmsc-username"], _ = expandObjectFirewallMmsProfileNotificationMmscUsername(d, i["mmsc_username"], pre_append)
	}
	pre_append = pre + ".0." + "msg_protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["msg-protocol"], _ = expandObjectFirewallMmsProfileNotificationMsgProtocol(d, i["msg_protocol"], pre_append)
	}
	pre_append = pre + ".0." + "msg_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["msg-type"], _ = expandObjectFirewallMmsProfileNotificationMsgType(d, i["msg_type"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectFirewallMmsProfileNotificationProtocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "rate_limit"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rate-limit"], _ = expandObjectFirewallMmsProfileNotificationRateLimit(d, i["rate_limit"], pre_append)
	}
	pre_append = pre + ".0." + "tod_window_duration"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tod-window-duration"], _ = expandObjectFirewallMmsProfileNotificationTodWindowDuration(d, i["tod_window_duration"], pre_append)
	}
	pre_append = pre + ".0." + "tod_window_end"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tod-window-end"], _ = expandObjectFirewallMmsProfileNotificationTodWindowEnd(d, i["tod_window_end"], pre_append)
	}
	pre_append = pre + ".0." + "tod_window_start"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tod-window-start"], _ = expandObjectFirewallMmsProfileNotificationTodWindowStart(d, i["tod_window_start"], pre_append)
	}
	pre_append = pre + ".0." + "user_domain"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["user-domain"], _ = expandObjectFirewallMmsProfileNotificationUserDomain(d, i["user_domain"], pre_append)
	}
	pre_append = pre + ".0." + "vas_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vas-id"], _ = expandObjectFirewallMmsProfileNotificationVasId(d, i["vas_id"], pre_append)
	}
	pre_append = pre + ".0." + "vasp_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vasp-id"], _ = expandObjectFirewallMmsProfileNotificationVaspId(d, i["vasp_id"], pre_append)
	}
	pre_append = pre + ".0." + "virus_int"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-int"], _ = expandObjectFirewallMmsProfileNotificationVirusInt(d, i["virus_int"], pre_append)
	}
	pre_append = pre + ".0." + "virus_int_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-int-mode"], _ = expandObjectFirewallMmsProfileNotificationVirusIntMode(d, i["virus_int_mode"], pre_append)
	}
	pre_append = pre + ".0." + "virus_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virus-status"], _ = expandObjectFirewallMmsProfileNotificationVirusStatus(d, i["virus_status"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallMmsProfileNotificationAlertInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationAlertIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationAlertSrcMsisdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationAlertStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationBwordInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationBwordIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationBwordStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationCarrierEndpointBwlInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationCarrierEndpointBwlIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationCarrierEndpointBwlStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationDaysAllowed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileNotificationDetectServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationDupeInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationDupeIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationDupeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFileBlockInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFileBlockIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFileBlockStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFloodInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFloodIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFloodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationFromInHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmsChecksumInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmsChecksumIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmsChecksumStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmscHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmscPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallMmsProfileNotificationMmscPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmscUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMmscUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMsgProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationTodWindowDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationTodWindowEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationTodWindowStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationUserDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationVasId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationVaspId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationVirusInt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationVirusIntMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileNotificationVirusStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandObjectFirewallMmsProfileOutbreakPreventionExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ftgd-service"], _ = expandObjectFirewallMmsProfileOutbreakPreventionFtgdService(d, i["ftgd_service"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallMmsProfileOutbreakPreventionExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileOutbreakPreventionFtgdService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileRemoveBlockedConstLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallMmsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("avnotificationtable"); ok || d.HasChange("avnotificationtable") {
		t, err := expandObjectFirewallMmsProfileAvnotificationtable(d, v, "avnotificationtable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["avnotificationtable"] = t
		}
	}

	if v, ok := d.GetOk("bwordtable"); ok || d.HasChange("bwordtable") {
		t, err := expandObjectFirewallMmsProfileBwordtable(d, v, "bwordtable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bwordtable"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix"); ok || d.HasChange("carrier_endpoint_prefix") {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefix(d, v, "carrier_endpoint_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix_range_max"); ok || d.HasChange("carrier_endpoint_prefix_range_max") {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefixRangeMax(d, v, "carrier_endpoint_prefix_range_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix-range-max"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix_range_min"); ok || d.HasChange("carrier_endpoint_prefix_range_min") {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefixRangeMin(d, v, "carrier_endpoint_prefix_range_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix-range-min"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix_string"); ok || d.HasChange("carrier_endpoint_prefix_string") {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefixString(d, v, "carrier_endpoint_prefix_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix-string"] = t
		}
	}

	if v, ok := d.GetOk("carrierendpointbwltable"); ok || d.HasChange("carrierendpointbwltable") {
		t, err := expandObjectFirewallMmsProfileCarrierendpointbwltable(d, v, "carrierendpointbwltable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrierendpointbwltable"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallMmsProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dupe"); ok || d.HasChange("dupe") {
		t, err := expandObjectFirewallMmsProfileDupe(d, v, "dupe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dupe"] = t
		}
	}

	if v, ok := d.GetOk("extended_utm_log"); ok || d.HasChange("extended_utm_log") {
		t, err := expandObjectFirewallMmsProfileExtendedUtmLog(d, v, "extended_utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-utm-log"] = t
		}
	}

	if v, ok := d.GetOk("flood"); ok || d.HasChange("flood") {
		t, err := expandObjectFirewallMmsProfileFlood(d, v, "flood")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flood"] = t
		}
	}

	if v, ok := d.GetOk("mm1"); ok || d.HasChange("mm1") {
		t, err := expandObjectFirewallMmsProfileMm1(d, v, "mm1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1"] = t
		}
	}

	if v, ok := d.GetOk("mm1_addr_hdr"); ok || d.HasChange("mm1_addr_hdr") {
		t, err := expandObjectFirewallMmsProfileMm1AddrHdr(d, v, "mm1_addr_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-addr-hdr"] = t
		}
	}

	if v, ok := d.GetOk("mm1_addr_source"); ok || d.HasChange("mm1_addr_source") {
		t, err := expandObjectFirewallMmsProfileMm1AddrSource(d, v, "mm1_addr_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-addr-source"] = t
		}
	}

	if v, ok := d.GetOk("mm1_convert_hex"); ok || d.HasChange("mm1_convert_hex") {
		t, err := expandObjectFirewallMmsProfileMm1ConvertHex(d, v, "mm1_convert_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-convert-hex"] = t
		}
	}

	if v, ok := d.GetOk("mm1_outbreak_prevention"); ok || d.HasChange("mm1_outbreak_prevention") {
		t, err := expandObjectFirewallMmsProfileMm1OutbreakPrevention(d, v, "mm1_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm1_retr_dupe"); ok || d.HasChange("mm1_retr_dupe") {
		t, err := expandObjectFirewallMmsProfileMm1RetrDupe(d, v, "mm1_retr_dupe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-retr-dupe"] = t
		}
	}

	if v, ok := d.GetOk("mm1_retrieve_scan"); ok || d.HasChange("mm1_retrieve_scan") {
		t, err := expandObjectFirewallMmsProfileMm1RetrieveScan(d, v, "mm1_retrieve_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-retrieve-scan"] = t
		}
	}

	if v, ok := d.GetOk("mm1comfortamount"); ok || d.HasChange("mm1comfortamount") {
		t, err := expandObjectFirewallMmsProfileMm1Comfortamount(d, v, "mm1comfortamount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1comfortamount"] = t
		}
	}

	if v, ok := d.GetOk("mm1comfortinterval"); ok || d.HasChange("mm1comfortinterval") {
		t, err := expandObjectFirewallMmsProfileMm1Comfortinterval(d, v, "mm1comfortinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1comfortinterval"] = t
		}
	}

	if v, ok := d.GetOk("mm1oversizelimit"); ok || d.HasChange("mm1oversizelimit") {
		t, err := expandObjectFirewallMmsProfileMm1Oversizelimit(d, v, "mm1oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mm3"); ok || d.HasChange("mm3") {
		t, err := expandObjectFirewallMmsProfileMm3(d, v, "mm3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3"] = t
		}
	}

	if v, ok := d.GetOk("mm3_outbreak_prevention"); ok || d.HasChange("mm3_outbreak_prevention") {
		t, err := expandObjectFirewallMmsProfileMm3OutbreakPrevention(d, v, "mm3_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm3oversizelimit"); ok || d.HasChange("mm3oversizelimit") {
		t, err := expandObjectFirewallMmsProfileMm3Oversizelimit(d, v, "mm3oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mm4"); ok || d.HasChange("mm4") {
		t, err := expandObjectFirewallMmsProfileMm4(d, v, "mm4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4"] = t
		}
	}

	if v, ok := d.GetOk("mm4_outbreak_prevention"); ok || d.HasChange("mm4_outbreak_prevention") {
		t, err := expandObjectFirewallMmsProfileMm4OutbreakPrevention(d, v, "mm4_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm4oversizelimit"); ok || d.HasChange("mm4oversizelimit") {
		t, err := expandObjectFirewallMmsProfileMm4Oversizelimit(d, v, "mm4oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mm7"); ok || d.HasChange("mm7") {
		t, err := expandObjectFirewallMmsProfileMm7(d, v, "mm7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7"] = t
		}
	}

	if v, ok := d.GetOk("mm7_addr_hdr"); ok || d.HasChange("mm7_addr_hdr") {
		t, err := expandObjectFirewallMmsProfileMm7AddrHdr(d, v, "mm7_addr_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-addr-hdr"] = t
		}
	}

	if v, ok := d.GetOk("mm7_addr_source"); ok || d.HasChange("mm7_addr_source") {
		t, err := expandObjectFirewallMmsProfileMm7AddrSource(d, v, "mm7_addr_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-addr-source"] = t
		}
	}

	if v, ok := d.GetOk("mm7_convert_hex"); ok || d.HasChange("mm7_convert_hex") {
		t, err := expandObjectFirewallMmsProfileMm7ConvertHex(d, v, "mm7_convert_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-convert-hex"] = t
		}
	}

	if v, ok := d.GetOk("mm7_outbreak_prevention"); ok || d.HasChange("mm7_outbreak_prevention") {
		t, err := expandObjectFirewallMmsProfileMm7OutbreakPrevention(d, v, "mm7_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm7comfortamount"); ok || d.HasChange("mm7comfortamount") {
		t, err := expandObjectFirewallMmsProfileMm7Comfortamount(d, v, "mm7comfortamount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7comfortamount"] = t
		}
	}

	if v, ok := d.GetOk("mm7comfortinterval"); ok || d.HasChange("mm7comfortinterval") {
		t, err := expandObjectFirewallMmsProfileMm7Comfortinterval(d, v, "mm7comfortinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7comfortinterval"] = t
		}
	}

	if v, ok := d.GetOk("mm7oversizelimit"); ok || d.HasChange("mm7oversizelimit") {
		t, err := expandObjectFirewallMmsProfileMm7Oversizelimit(d, v, "mm7oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mms_antispam_mass_log"); ok || d.HasChange("mms_antispam_mass_log") {
		t, err := expandObjectFirewallMmsProfileMmsAntispamMassLog(d, v, "mms_antispam_mass_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-antispam-mass-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_av_block_log"); ok || d.HasChange("mms_av_block_log") {
		t, err := expandObjectFirewallMmsProfileMmsAvBlockLog(d, v, "mms_av_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-av-block-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_av_oversize_log"); ok || d.HasChange("mms_av_oversize_log") {
		t, err := expandObjectFirewallMmsProfileMmsAvOversizeLog(d, v, "mms_av_oversize_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-av-oversize-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_av_virus_log"); ok || d.HasChange("mms_av_virus_log") {
		t, err := expandObjectFirewallMmsProfileMmsAvVirusLog(d, v, "mms_av_virus_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-av-virus-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_carrier_endpoint_filter_log"); ok || d.HasChange("mms_carrier_endpoint_filter_log") {
		t, err := expandObjectFirewallMmsProfileMmsCarrierEndpointFilterLog(d, v, "mms_carrier_endpoint_filter_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-carrier-endpoint-filter-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_checksum_log"); ok || d.HasChange("mms_checksum_log") {
		t, err := expandObjectFirewallMmsProfileMmsChecksumLog(d, v, "mms_checksum_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-checksum-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_checksum_table"); ok || d.HasChange("mms_checksum_table") {
		t, err := expandObjectFirewallMmsProfileMmsChecksumTable(d, v, "mms_checksum_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-checksum-table"] = t
		}
	}

	if v, ok := d.GetOk("mms_notification_log"); ok || d.HasChange("mms_notification_log") {
		t, err := expandObjectFirewallMmsProfileMmsNotificationLog(d, v, "mms_notification_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-notification-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_web_content_log"); ok || d.HasChange("mms_web_content_log") {
		t, err := expandObjectFirewallMmsProfileMmsWebContentLog(d, v, "mms_web_content_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-web-content-log"] = t
		}
	}

	if v, ok := d.GetOk("mmsbwordthreshold"); ok || d.HasChange("mmsbwordthreshold") {
		t, err := expandObjectFirewallMmsProfileMmsbwordthreshold(d, v, "mmsbwordthreshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mmsbwordthreshold"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallMmsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notif_msisdn"); ok || d.HasChange("notif_msisdn") {
		t, err := expandObjectFirewallMmsProfileNotifMsisdn(d, v, "notif_msisdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notif-msisdn"] = t
		}
	}

	if v, ok := d.GetOk("notification"); ok || d.HasChange("notification") {
		t, err := expandObjectFirewallMmsProfileNotification(d, v, "notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok || d.HasChange("outbreak_prevention") {
		t, err := expandObjectFirewallMmsProfileOutbreakPrevention(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("remove_blocked_const_length"); ok || d.HasChange("remove_blocked_const_length") {
		t, err := expandObjectFirewallMmsProfileRemoveBlockedConstLength(d, v, "remove_blocked_const_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-blocked-const-length"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectFirewallMmsProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	return &obj, nil
}
