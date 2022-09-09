// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure application control lists.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectApplicationList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectApplicationListCreate,
		Read:   resourceObjectApplicationListRead,
		Update: resourceObjectApplicationListUpdate,
		Delete: resourceObjectApplicationListDelete,

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
			"app_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"control_default_network_services": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"deep_app_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_network_services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"services": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"violation_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"enforce_default_app_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
						},
						"behavior": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclusion": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"members": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"per_ip_shaper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"popularity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"protocols": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rate_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rate_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rate_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"risk": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
						},
						"session_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"shaper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shaper_reverse": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sub_category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
						},
						"technology": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"vendor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_inclusion_ssl_di_sigs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"other_application_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"other_application_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"p2p_block_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"p2p_black_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unknown_application_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unknown_application_log": &schema.Schema{
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

func resourceObjectApplicationListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectApplicationList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectApplicationList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectApplicationListRead(d, m)
}

func resourceObjectApplicationListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectApplicationList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectApplicationList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectApplicationListRead(d, m)
}

func resourceObjectApplicationListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectApplicationList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectApplicationList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectApplicationListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectApplicationList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectApplicationList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationList resource from API: %v", err)
	}
	return nil
}

func flattenObjectApplicationListAppReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListControlDefaultNetworkServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListDeepAppInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListDefaultNetworkServices(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectApplicationListDefaultNetworkServicesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationList-DefaultNetworkServices-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectApplicationListDefaultNetworkServicesPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectApplicationList-DefaultNetworkServices-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := i["services"]; ok {
			v := flattenObjectApplicationListDefaultNetworkServicesServices(i["services"], d, pre_append)
			tmp["services"] = fortiAPISubPartPatch(v, "ObjectApplicationList-DefaultNetworkServices-Services")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "violation_action"
		if _, ok := i["violation-action"]; ok {
			v := flattenObjectApplicationListDefaultNetworkServicesViolationAction(i["violation-action"], d, pre_append)
			tmp["violation_action"] = fortiAPISubPartPatch(v, "ObjectApplicationList-DefaultNetworkServices-ViolationAction")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectApplicationListDefaultNetworkServicesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListDefaultNetworkServicesPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListDefaultNetworkServicesServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListDefaultNetworkServicesViolationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEnforceDefaultAppPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectApplicationListEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenObjectApplicationListEntriesApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := i["behavior"]; ok {
			v := flattenObjectApplicationListEntriesBehavior(i["behavior"], d, pre_append)
			tmp["behavior"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Behavior")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectApplicationListEntriesCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusion"
		if _, ok := i["exclusion"]; ok {
			v := flattenObjectApplicationListEntriesExclusion(i["exclusion"], d, pre_append)
			tmp["exclusion"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Exclusion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectApplicationListEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectApplicationListEntriesLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenObjectApplicationListEntriesLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := i["parameters"]; ok {
			v := flattenObjectApplicationListEntriesParameters(i["parameters"], d, pre_append)
			tmp["parameters"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Parameters")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := i["per-ip-shaper"]; ok {
			v := flattenObjectApplicationListEntriesPerIpShaper(i["per-ip-shaper"], d, pre_append)
			tmp["per_ip_shaper"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-PerIpShaper")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := i["popularity"]; ok {
			v := flattenObjectApplicationListEntriesPopularity(i["popularity"], d, pre_append)
			tmp["popularity"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Popularity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := i["protocols"]; ok {
			v := flattenObjectApplicationListEntriesProtocols(i["protocols"], d, pre_append)
			tmp["protocols"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Protocols")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenObjectApplicationListEntriesQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenObjectApplicationListEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenObjectApplicationListEntriesQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {
			v := flattenObjectApplicationListEntriesRateCount(i["rate-count"], d, pre_append)
			tmp["rate_count"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-RateCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {
			v := flattenObjectApplicationListEntriesRateDuration(i["rate-duration"], d, pre_append)
			tmp["rate_duration"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-RateDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {
			v := flattenObjectApplicationListEntriesRateMode(i["rate-mode"], d, pre_append)
			tmp["rate_mode"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-RateMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {
			v := flattenObjectApplicationListEntriesRateTrack(i["rate-track"], d, pre_append)
			tmp["rate_track"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-RateTrack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := i["risk"]; ok {
			v := flattenObjectApplicationListEntriesRisk(i["risk"], d, pre_append)
			tmp["risk"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Risk")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := i["session-ttl"]; ok {
			v := flattenObjectApplicationListEntriesSessionTtl(i["session-ttl"], d, pre_append)
			tmp["session_ttl"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-SessionTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := i["shaper"]; ok {
			v := flattenObjectApplicationListEntriesShaper(i["shaper"], d, pre_append)
			tmp["shaper"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Shaper")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := i["shaper-reverse"]; ok {
			v := flattenObjectApplicationListEntriesShaperReverse(i["shaper-reverse"], d, pre_append)
			tmp["shaper_reverse"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-ShaperReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := i["sub-category"]; ok {
			v := flattenObjectApplicationListEntriesSubCategory(i["sub-category"], d, pre_append)
			tmp["sub_category"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-SubCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := i["technology"]; ok {
			v := flattenObjectApplicationListEntriesTechnology(i["technology"], d, pre_append)
			tmp["technology"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Technology")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := i["vendor"]; ok {
			v := flattenObjectApplicationListEntriesVendor(i["vendor"], d, pre_append)
			tmp["vendor"] = fortiAPISubPartPatch(v, "ObjectApplicationList-Entries-Vendor")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectApplicationListEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesExclusion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParameters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectApplicationListEntriesParametersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntries-Parameters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntries-Parameters-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectApplicationListEntriesParametersValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntries-Parameters-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectApplicationListEntriesParametersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectApplicationListEntriesParametersMembersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesPopularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesSubCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListForceInclusionSslDiSigs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListOtherApplicationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListOtherApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListP2PBlockList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListP2PBlackList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListUnknownApplicationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListUnknownApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectApplicationList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("app_replacemsg", flattenObjectApplicationListAppReplacemsg(o["app-replacemsg"], d, "app_replacemsg")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-replacemsg"], "ObjectApplicationList-AppReplacemsg"); ok {
			if err = d.Set("app_replacemsg", vv); err != nil {
				return fmt.Errorf("Error reading app_replacemsg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_replacemsg: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectApplicationListComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectApplicationList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("control_default_network_services", flattenObjectApplicationListControlDefaultNetworkServices(o["control-default-network-services"], d, "control_default_network_services")); err != nil {
		if vv, ok := fortiAPIPatch(o["control-default-network-services"], "ObjectApplicationList-ControlDefaultNetworkServices"); ok {
			if err = d.Set("control_default_network_services", vv); err != nil {
				return fmt.Errorf("Error reading control_default_network_services: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading control_default_network_services: %v", err)
		}
	}

	if err = d.Set("deep_app_inspection", flattenObjectApplicationListDeepAppInspection(o["deep-app-inspection"], d, "deep_app_inspection")); err != nil {
		if vv, ok := fortiAPIPatch(o["deep-app-inspection"], "ObjectApplicationList-DeepAppInspection"); ok {
			if err = d.Set("deep_app_inspection", vv); err != nil {
				return fmt.Errorf("Error reading deep_app_inspection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading deep_app_inspection: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("default_network_services", flattenObjectApplicationListDefaultNetworkServices(o["default-network-services"], d, "default_network_services")); err != nil {
			if vv, ok := fortiAPIPatch(o["default-network-services"], "ObjectApplicationList-DefaultNetworkServices"); ok {
				if err = d.Set("default_network_services", vv); err != nil {
					return fmt.Errorf("Error reading default_network_services: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading default_network_services: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("default_network_services"); ok {
			if err = d.Set("default_network_services", flattenObjectApplicationListDefaultNetworkServices(o["default-network-services"], d, "default_network_services")); err != nil {
				if vv, ok := fortiAPIPatch(o["default-network-services"], "ObjectApplicationList-DefaultNetworkServices"); ok {
					if err = d.Set("default_network_services", vv); err != nil {
						return fmt.Errorf("Error reading default_network_services: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading default_network_services: %v", err)
				}
			}
		}
	}

	if err = d.Set("enforce_default_app_port", flattenObjectApplicationListEnforceDefaultAppPort(o["enforce-default-app-port"], d, "enforce_default_app_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-default-app-port"], "ObjectApplicationList-EnforceDefaultAppPort"); ok {
			if err = d.Set("enforce_default_app_port", vv); err != nil {
				return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectApplicationListEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectApplicationList-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectApplicationListEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectApplicationList-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenObjectApplicationListExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectApplicationList-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("force_inclusion_ssl_di_sigs", flattenObjectApplicationListForceInclusionSslDiSigs(o["force-inclusion-ssl-di-sigs"], d, "force_inclusion_ssl_di_sigs")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-inclusion-ssl-di-sigs"], "ObjectApplicationList-ForceInclusionSslDiSigs"); ok {
			if err = d.Set("force_inclusion_ssl_di_sigs", vv); err != nil {
				return fmt.Errorf("Error reading force_inclusion_ssl_di_sigs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_inclusion_ssl_di_sigs: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectApplicationListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectApplicationList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectApplicationListOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectApplicationList-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("other_application_action", flattenObjectApplicationListOtherApplicationAction(o["other-application-action"], d, "other_application_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["other-application-action"], "ObjectApplicationList-OtherApplicationAction"); ok {
			if err = d.Set("other_application_action", vv); err != nil {
				return fmt.Errorf("Error reading other_application_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading other_application_action: %v", err)
		}
	}

	if err = d.Set("other_application_log", flattenObjectApplicationListOtherApplicationLog(o["other-application-log"], d, "other_application_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["other-application-log"], "ObjectApplicationList-OtherApplicationLog"); ok {
			if err = d.Set("other_application_log", vv); err != nil {
				return fmt.Errorf("Error reading other_application_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading other_application_log: %v", err)
		}
	}

	if err = d.Set("p2p_block_list", flattenObjectApplicationListP2PBlockList(o["p2p-block-list"], d, "p2p_block_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["p2p-block-list"], "ObjectApplicationList-P2PBlockList"); ok {
			if err = d.Set("p2p_block_list", vv); err != nil {
				return fmt.Errorf("Error reading p2p_block_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading p2p_block_list: %v", err)
		}
	}

	if err = d.Set("p2p_black_list", flattenObjectApplicationListP2PBlackList(o["p2p-black-list"], d, "p2p_black_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["p2p-black-list"], "ObjectApplicationList-P2PBlackList"); ok {
			if err = d.Set("p2p_black_list", vv); err != nil {
				return fmt.Errorf("Error reading p2p_black_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading p2p_black_list: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectApplicationListReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectApplicationList-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("unknown_application_action", flattenObjectApplicationListUnknownApplicationAction(o["unknown-application-action"], d, "unknown_application_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-application-action"], "ObjectApplicationList-UnknownApplicationAction"); ok {
			if err = d.Set("unknown_application_action", vv); err != nil {
				return fmt.Errorf("Error reading unknown_application_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_application_action: %v", err)
		}
	}

	if err = d.Set("unknown_application_log", flattenObjectApplicationListUnknownApplicationLog(o["unknown-application-log"], d, "unknown_application_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-application-log"], "ObjectApplicationList-UnknownApplicationLog"); ok {
			if err = d.Set("unknown_application_log", vv); err != nil {
				return fmt.Errorf("Error reading unknown_application_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_application_log: %v", err)
		}
	}

	return nil
}

func flattenObjectApplicationListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectApplicationListAppReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListControlDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListDeepAppInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectApplicationListDefaultNetworkServicesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectApplicationListDefaultNetworkServicesPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["services"], _ = expandObjectApplicationListDefaultNetworkServicesServices(d, i["services"], pre_append)
		} else {
			tmp["services"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "violation_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["violation-action"], _ = expandObjectApplicationListDefaultNetworkServicesViolationAction(d, i["violation_action"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListDefaultNetworkServicesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListDefaultNetworkServicesPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListDefaultNetworkServicesServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListDefaultNetworkServicesViolationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEnforceDefaultAppPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectApplicationListEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandObjectApplicationListEntriesApplication(d, i["application"], pre_append)
		} else {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["behavior"], _ = expandObjectApplicationListEntriesBehavior(d, i["behavior"], pre_append)
		} else {
			tmp["behavior"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectApplicationListEntriesCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusion"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclusion"], _ = expandObjectApplicationListEntriesExclusion(d, i["exclusion"], pre_append)
		} else {
			tmp["exclusion"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectApplicationListEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectApplicationListEntriesLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-packet"], _ = expandObjectApplicationListEntriesLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["parameters"], _ = expandObjectApplicationListEntriesParameters(d, i["parameters"], pre_append)
		} else {
			tmp["parameters"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["per-ip-shaper"], _ = expandObjectApplicationListEntriesPerIpShaper(d, i["per_ip_shaper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["popularity"], _ = expandObjectApplicationListEntriesPopularity(d, i["popularity"], pre_append)
		} else {
			tmp["popularity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocols"], _ = expandObjectApplicationListEntriesProtocols(d, i["protocols"], pre_append)
		} else {
			tmp["protocols"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandObjectApplicationListEntriesQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandObjectApplicationListEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandObjectApplicationListEntriesQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-count"], _ = expandObjectApplicationListEntriesRateCount(d, i["rate_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-duration"], _ = expandObjectApplicationListEntriesRateDuration(d, i["rate_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-mode"], _ = expandObjectApplicationListEntriesRateMode(d, i["rate_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-track"], _ = expandObjectApplicationListEntriesRateTrack(d, i["rate_track"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["risk"], _ = expandObjectApplicationListEntriesRisk(d, i["risk"], pre_append)
		} else {
			tmp["risk"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["session-ttl"], _ = expandObjectApplicationListEntriesSessionTtl(d, i["session_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shaper"], _ = expandObjectApplicationListEntriesShaper(d, i["shaper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shaper-reverse"], _ = expandObjectApplicationListEntriesShaperReverse(d, i["shaper_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sub-category"], _ = expandObjectApplicationListEntriesSubCategory(d, i["sub_category"], pre_append)
		} else {
			tmp["sub-category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["technology"], _ = expandObjectApplicationListEntriesTechnology(d, i["technology"], pre_append)
		} else {
			tmp["technology"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vendor"], _ = expandObjectApplicationListEntriesVendor(d, i["vendor"], pre_append)
		} else {
			tmp["vendor"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesExclusion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectApplicationListEntriesParametersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandObjectApplicationListEntriesParametersMembers(d, i["members"], pre_append)
		} else {
			tmp["members"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectApplicationListEntriesParametersValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListEntriesParametersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectApplicationListEntriesParametersMembersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectApplicationListEntriesParametersMembersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectApplicationListEntriesParametersMembersValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListEntriesParametersMembersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesPopularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesSubCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListForceInclusionSslDiSigs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListOtherApplicationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListOtherApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListP2PBlockList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListP2PBlackList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListUnknownApplicationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListUnknownApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectApplicationList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_replacemsg"); ok || d.HasChange("app_replacemsg") {
		t, err := expandObjectApplicationListAppReplacemsg(d, v, "app_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectApplicationListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("control_default_network_services"); ok || d.HasChange("control_default_network_services") {
		t, err := expandObjectApplicationListControlDefaultNetworkServices(d, v, "control_default_network_services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-default-network-services"] = t
		}
	}

	if v, ok := d.GetOk("deep_app_inspection"); ok || d.HasChange("deep_app_inspection") {
		t, err := expandObjectApplicationListDeepAppInspection(d, v, "deep_app_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-inspection"] = t
		}
	}

	if v, ok := d.GetOk("default_network_services"); ok || d.HasChange("default_network_services") {
		t, err := expandObjectApplicationListDefaultNetworkServices(d, v, "default_network_services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-network-services"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok || d.HasChange("enforce_default_app_port") {
		t, err := expandObjectApplicationListEnforceDefaultAppPort(d, v, "enforce_default_app_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectApplicationListEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandObjectApplicationListExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("force_inclusion_ssl_di_sigs"); ok || d.HasChange("force_inclusion_ssl_di_sigs") {
		t, err := expandObjectApplicationListForceInclusionSslDiSigs(d, v, "force_inclusion_ssl_di_sigs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-inclusion-ssl-di-sigs"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectApplicationListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectApplicationListOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("other_application_action"); ok || d.HasChange("other_application_action") {
		t, err := expandObjectApplicationListOtherApplicationAction(d, v, "other_application_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-action"] = t
		}
	}

	if v, ok := d.GetOk("other_application_log"); ok || d.HasChange("other_application_log") {
		t, err := expandObjectApplicationListOtherApplicationLog(d, v, "other_application_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-log"] = t
		}
	}

	if v, ok := d.GetOk("p2p_block_list"); ok || d.HasChange("p2p_block_list") {
		t, err := expandObjectApplicationListP2PBlockList(d, v, "p2p_block_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-block-list"] = t
		}
	}

	if v, ok := d.GetOk("p2p_black_list"); ok || d.HasChange("p2p_black_list") {
		t, err := expandObjectApplicationListP2PBlackList(d, v, "p2p_black_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-black-list"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectApplicationListReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_action"); ok || d.HasChange("unknown_application_action") {
		t, err := expandObjectApplicationListUnknownApplicationAction(d, v, "unknown_application_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-action"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_log"); ok || d.HasChange("unknown_application_log") {
		t, err := expandObjectApplicationListUnknownApplicationLog(d, v, "unknown_application_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-log"] = t
		}
	}

	return &obj, nil
}
