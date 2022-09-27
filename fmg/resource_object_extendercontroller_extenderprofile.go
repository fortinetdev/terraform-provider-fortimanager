// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender extender profile configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileCreate,
		Read:   resourceObjectExtenderControllerExtenderProfileRead,
		Update: resourceObjectExtenderControllerExtenderProfileUpdate,
		Delete: resourceObjectExtenderControllerExtenderProfileDelete,

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
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cellular": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_report": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"signal_threshold": &schema.Schema{
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
						"dataplan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"modem1": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_switch": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dataplan": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"disconnect_period": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"disconnect_threshold": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"signal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"switch_back": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"switch_back_time": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"switch_back_timer": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"conn_status": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"default_sim": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gps": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"modem_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"preferred_carrier": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"redundant_intf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"redundant_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin_code": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"sim2_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim2_pin_code": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
						},
						"modem2": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_switch": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dataplan": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"disconnect_period": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"disconnect_threshold": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"signal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"switch_back": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"switch_back_time": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"switch_back_timer": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"conn_status": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"default_sim": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gps": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"modem_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"preferred_carrier": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"redundant_intf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"redundant_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin_code": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"sim2_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim2_pin_code": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
						},
						"sms_notification": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"data_exhausted": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"fgt_backup_mode_switch": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"low_signal_strength": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"mode_switch": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"os_image_fallback": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"session_disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"system_reboot": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"receiver": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alert": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"phone_number": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
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
				},
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backhaul": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"backhaul_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"backhaul_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link_loadbalance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"login_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectExtenderControllerExtenderProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerExtenderProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerExtenderProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectExtenderControllerExtenderProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerExtenderProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerExtenderProfileRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerExtenderProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtenderControllerExtenderProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerExtenderProfileRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectExtenderControllerExtenderProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectExtenderControllerExtenderProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileBandwidthLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellular(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "controller_report"
	if _, ok := i["controller-report"]; ok {
		result["controller_report"] = flattenObjectExtenderControllerExtenderProfileCellularControllerReport(i["controller-report"], d, pre_append)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtenderControllerExtenderProfileCellularDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1"
	if _, ok := i["modem1"]; ok {
		result["modem1"] = flattenObjectExtenderControllerExtenderProfileCellularModem1(i["modem1"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2"
	if _, ok := i["modem2"]; ok {
		result["modem2"] = flattenObjectExtenderControllerExtenderProfileCellularModem2(i["modem2"], d, pre_append)
	}

	pre_append = pre + ".0." + "sms_notification"
	if _, ok := i["sms-notification"]; ok {
		result["sms_notification"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotification(i["sms-notification"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReport(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenObjectExtenderControllerExtenderProfileCellularControllerReportInterval(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(i["signal-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtenderControllerExtenderProfileCellularControllerReportStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenObjectExtenderControllerExtenderProfileCellularModem1ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenObjectExtenderControllerExtenderProfileCellularModem1DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenObjectExtenderControllerExtenderProfileCellularModem1Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenObjectExtenderControllerExtenderProfileCellularModem1ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenObjectExtenderControllerExtenderProfileCellularModem1PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenObjectExtenderControllerExtenderProfileCellularModem1RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenObjectExtenderControllerExtenderProfileCellularModem1RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenObjectExtenderControllerExtenderProfileCellularModem1Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenObjectExtenderControllerExtenderProfileCellularModem1Sim1PinCode(i["sim1-pin-code"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["sim1_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenObjectExtenderControllerExtenderProfileCellularModem1Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenObjectExtenderControllerExtenderProfileCellularModem1Sim2PinCode(i["sim2-pin-code"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["sim2_pin_code"] = c
		}
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1Sim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem1Sim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenObjectExtenderControllerExtenderProfileCellularModem2ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenObjectExtenderControllerExtenderProfileCellularModem2DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenObjectExtenderControllerExtenderProfileCellularModem2Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenObjectExtenderControllerExtenderProfileCellularModem2ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode(i["sim1-pin-code"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["sim1_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode(i["sim2-pin-code"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["sim2_pin_code"] = c
		}
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotification(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(i["alert"], d, pre_append)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(i["receiver"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileExtension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul"
	if _, ok := i["backhaul"]; ok {
		result["backhaul"] = flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaul(i["backhaul"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {
		result["backhaul_interface"] = flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {
		result["backhaul_ip"] = flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {
		result["ipsec_tunnel"] = flattenObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := i["link-loadbalance"]; ok {
		result["link_loadbalance"] = flattenObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(i["link-loadbalance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaul(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectExtenderControllerExtenderProfileLanExtension-Backhaul-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileLoginPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerExtenderProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("allowaccess", flattenObjectExtenderControllerExtenderProfileAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ObjectExtenderControllerExtenderProfile-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenObjectExtenderControllerExtenderProfileBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-limit"], "ObjectExtenderControllerExtenderProfile-BandwidthLimit"); ok {
			if err = d.Set("bandwidth_limit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cellular", flattenObjectExtenderControllerExtenderProfileCellular(o["cellular"], d, "cellular")); err != nil {
			if vv, ok := fortiAPIPatch(o["cellular"], "ObjectExtenderControllerExtenderProfile-Cellular"); ok {
				if err = d.Set("cellular", vv); err != nil {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cellular: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cellular"); ok {
			if err = d.Set("cellular", flattenObjectExtenderControllerExtenderProfileCellular(o["cellular"], d, "cellular")); err != nil {
				if vv, ok := fortiAPIPatch(o["cellular"], "ObjectExtenderControllerExtenderProfile-Cellular"); ok {
					if err = d.Set("cellular", vv); err != nil {
						return fmt.Errorf("Error reading cellular: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			}
		}
	}

	if err = d.Set("enforce_bandwidth", flattenObjectExtenderControllerExtenderProfileEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-bandwidth"], "ObjectExtenderControllerExtenderProfile-EnforceBandwidth"); ok {
			if err = d.Set("enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("extension", flattenObjectExtenderControllerExtenderProfileExtension(o["extension"], d, "extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension"], "ObjectExtenderControllerExtenderProfile-Extension"); ok {
			if err = d.Set("extension", vv); err != nil {
				return fmt.Errorf("Error reading extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectExtenderControllerExtenderProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectExtenderControllerExtenderProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan_extension", flattenObjectExtenderControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["lan-extension"], "ObjectExtenderControllerExtenderProfile-LanExtension"); ok {
				if err = d.Set("lan_extension", vv); err != nil {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenObjectExtenderControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["lan-extension"], "ObjectExtenderControllerExtenderProfile-LanExtension"); ok {
					if err = d.Set("lan_extension", vv); err != nil {
						return fmt.Errorf("Error reading lan_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("login_password_change", flattenObjectExtenderControllerExtenderProfileLoginPasswordChange(o["login-password-change"], d, "login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-password-change"], "ObjectExtenderControllerExtenderProfile-LoginPasswordChange"); ok {
			if err = d.Set("login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("model", flattenObjectExtenderControllerExtenderProfileModel(o["model"], d, "model")); err != nil {
		if vv, ok := fortiAPIPatch(o["model"], "ObjectExtenderControllerExtenderProfile-Model"); ok {
			if err = d.Set("model", vv); err != nil {
				return fmt.Errorf("Error reading model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtenderControllerExtenderProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtenderControllerExtenderProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileBandwidthLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellular(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "controller_report"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularControllerReport(d, i["controller_report"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["controller-report"] = t
		}
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtenderControllerExtenderProfileCellularDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "modem1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem1(d, i["modem1"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["modem1"] = t
		}
	}
	pre_append = pre + ".0." + "modem2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2(d, i["modem2"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["modem2"] = t
		}
	}
	pre_append = pre + ".0." + "sms_notification"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotification(d, i["sms_notification"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["sms-notification"] = t
		}
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interval"], _ = expandObjectExtenderControllerExtenderProfileCellularControllerReportInterval(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal-threshold"], _ = expandObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(d, i["signal_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtenderControllerExtenderProfileCellularControllerReportStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReportInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReportStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(d, i["alert"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["alert"] = t
		}
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d, i["receiver"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["receiver"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alert"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d, i["backhaul"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["backhaul"] = t
		}
	}
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-interface"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-ip"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-tunnel"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["link-loadbalance"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(d, i["link_loadbalance"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileLoginPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerExtenderProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandObjectExtenderControllerExtenderProfileAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok || d.HasChange("bandwidth_limit") {
		t, err := expandObjectExtenderControllerExtenderProfileBandwidthLimit(d, v, "bandwidth_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("cellular"); ok || d.HasChange("cellular") {
		t, err := expandObjectExtenderControllerExtenderProfileCellular(d, v, "cellular")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cellular"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok || d.HasChange("enforce_bandwidth") {
		t, err := expandObjectExtenderControllerExtenderProfileEnforceBandwidth(d, v, "enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("extension"); ok || d.HasChange("extension") {
		t, err := expandObjectExtenderControllerExtenderProfileExtension(d, v, "extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectExtenderControllerExtenderProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok || d.HasChange("lan_extension") {
		t, err := expandObjectExtenderControllerExtenderProfileLanExtension(d, v, "lan_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok || d.HasChange("login_password") {
		t, err := expandObjectExtenderControllerExtenderProfileLoginPassword(d, v, "login_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok || d.HasChange("login_password_change") {
		t, err := expandObjectExtenderControllerExtenderProfileLoginPasswordChange(d, v, "login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("model"); ok || d.HasChange("model") {
		t, err := expandObjectExtenderControllerExtenderProfileModel(d, v, "model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtenderControllerExtenderProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
