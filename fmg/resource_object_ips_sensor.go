// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPS sensor.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectIpsSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIpsSensorCreate,
		Read:   resourceObjectIpsSensorRead,
		Update: resourceObjectIpsSensorUpdate,
		Delete: resourceObjectIpsSensorDelete,

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
			"block_malicious_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"cve": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"exempt_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_attack_context": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rate_duration": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rate_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectIpsSensorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectIpsSensor(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsSensor resource while getting object: %v", err)
	}

	_, err = c.CreateObjectIpsSensor(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsSensor resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIpsSensorRead(d, m)
}

func resourceObjectIpsSensorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectIpsSensor(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsSensor resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectIpsSensor(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIpsSensorRead(d, m)
}

func resourceObjectIpsSensorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectIpsSensor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIpsSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIpsSensorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectIpsSensor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIpsSensor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsSensor resource from API: %v", err)
	}
	return nil
}

func flattenObjectIpsSensorBlockMaliciousUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectIpsSensorEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenObjectIpsSensorEntriesApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := i["cve"]; ok {
			v := flattenObjectIpsSensorEntriesCve(i["cve"], d, pre_append)
			tmp["cve"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Cve")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := i["exempt-ip"]; ok {
			v := flattenObjectIpsSensorEntriesExemptIp(i["exempt-ip"], d, pre_append)
			tmp["exempt_ip"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-ExemptIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectIpsSensorEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := i["location"]; ok {
			v := flattenObjectIpsSensorEntriesLocation(i["location"], d, pre_append)
			tmp["location"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Location")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectIpsSensorEntriesLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := i["log-attack-context"]; ok {
			v := flattenObjectIpsSensorEntriesLogAttackContext(i["log-attack-context"], d, pre_append)
			tmp["log_attack_context"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-LogAttackContext")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenObjectIpsSensorEntriesLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {
			v := flattenObjectIpsSensorEntriesOs(i["os"], d, pre_append)
			tmp["os"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Os")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectIpsSensorEntriesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenObjectIpsSensorEntriesQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenObjectIpsSensorEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenObjectIpsSensorEntriesQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {
			v := flattenObjectIpsSensorEntriesRateCount(i["rate-count"], d, pre_append)
			tmp["rate_count"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-RateCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {
			v := flattenObjectIpsSensorEntriesRateDuration(i["rate-duration"], d, pre_append)
			tmp["rate_duration"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-RateDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {
			v := flattenObjectIpsSensorEntriesRateMode(i["rate-mode"], d, pre_append)
			tmp["rate_mode"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-RateMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {
			v := flattenObjectIpsSensorEntriesRateTrack(i["rate-track"], d, pre_append)
			tmp["rate_track"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-RateTrack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := i["rule"]; ok {
			v := flattenObjectIpsSensorEntriesRule(i["rule"], d, pre_append)
			tmp["rule"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Rule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectIpsSensorEntriesSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectIpsSensorEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectIpsSensor-Entries-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectIpsSensorEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "pass",
			2: "block",
			4: "reset",
			5: "default",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesCve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesExemptIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := i["dst-ip"]; ok {
			v := flattenObjectIpsSensorEntriesExemptIpDstIp(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "ObjectIpsSensorEntries-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectIpsSensorEntriesExemptIpId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectIpsSensorEntries-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenObjectIpsSensorEntriesExemptIpSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "ObjectIpsSensorEntries-ExemptIp-SrcIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectIpsSensorEntriesExemptIpDstIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesExemptIpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesExemptIpSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesLogAttackContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "attacker",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesRateCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			8: "periodical",
			9: "continuous",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "src-ip",
			2: "dest-ip",
			3: "dhcp-client-mac",
			4: "dns-domain",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorEntriesRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
			3: "default",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectIpsSensorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "block",
			2: "monitor",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectIpsSensor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("block_malicious_url", flattenObjectIpsSensorBlockMaliciousUrl(o["block-malicious-url"], d, "block_malicious_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-malicious-url"], "ObjectIpsSensor-BlockMaliciousUrl"); ok {
			if err = d.Set("block_malicious_url", vv); err != nil {
				return fmt.Errorf("Error reading block_malicious_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_malicious_url: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectIpsSensorComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectIpsSensor-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectIpsSensorEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectIpsSensor-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectIpsSensorEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectIpsSensor-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenObjectIpsSensorExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectIpsSensor-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectIpsSensorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectIpsSensor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectIpsSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectIpsSensor-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenObjectIpsSensorScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "ObjectIpsSensor-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	return nil
}

func flattenObjectIpsSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIpsSensorBlockMaliciousUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectIpsSensorEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["application"], _ = expandObjectIpsSensorEntriesApplication(d, i["application"], pre_append)
		} else {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cve"], _ = expandObjectIpsSensorEntriesCve(d, i["cve"], pre_append)
		} else {
			tmp["cve"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["exempt-ip"], _ = expandObjectIpsSensorEntriesExemptIp(d, i["exempt_ip"], pre_append)
		} else {
			tmp["exempt-ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectIpsSensorEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["location"], _ = expandObjectIpsSensorEntriesLocation(d, i["location"], pre_append)
		} else {
			tmp["location"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandObjectIpsSensorEntriesLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log-attack-context"], _ = expandObjectIpsSensorEntriesLogAttackContext(d, i["log_attack_context"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log-packet"], _ = expandObjectIpsSensorEntriesLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["os"], _ = expandObjectIpsSensorEntriesOs(d, i["os"], pre_append)
		} else {
			tmp["os"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandObjectIpsSensorEntriesProtocol(d, i["protocol"], pre_append)
		} else {
			tmp["protocol"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine"], _ = expandObjectIpsSensorEntriesQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-expiry"], _ = expandObjectIpsSensorEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-log"], _ = expandObjectIpsSensorEntriesQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-count"], _ = expandObjectIpsSensorEntriesRateCount(d, i["rate_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-duration"], _ = expandObjectIpsSensorEntriesRateDuration(d, i["rate_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-mode"], _ = expandObjectIpsSensorEntriesRateMode(d, i["rate_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-track"], _ = expandObjectIpsSensorEntriesRateTrack(d, i["rate_track"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rule"], _ = expandObjectIpsSensorEntriesRule(d, i["rule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandObjectIpsSensorEntriesSeverity(d, i["severity"], pre_append)
		} else {
			tmp["severity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandObjectIpsSensorEntriesStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectIpsSensorEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesCve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesExemptIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-ip"], _ = expandObjectIpsSensorEntriesExemptIpDstIp(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectIpsSensorEntriesExemptIpId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-ip"], _ = expandObjectIpsSensorEntriesExemptIpSrcIp(d, i["src_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectIpsSensorEntriesExemptIpDstIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesExemptIpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesExemptIpSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesLogAttackContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIpsSensor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_malicious_url"); ok {
		t, err := expandObjectIpsSensorBlockMaliciousUrl(d, v, "block_malicious_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malicious-url"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectIpsSensorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandObjectIpsSensorEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandObjectIpsSensorExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectIpsSensorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandObjectIpsSensorReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandObjectIpsSensorScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	return &obj, nil
}
