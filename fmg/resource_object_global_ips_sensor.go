// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
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

func resourceObjectGlobalIpsSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGlobalIpsSensorCreate,
		Read:   resourceObjectGlobalIpsSensorRead,
		Update: resourceObjectGlobalIpsSensorUpdate,
		Delete: resourceObjectGlobalIpsSensorDelete,

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
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"cve": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exempt_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"last_modified": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_attack_context": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"position": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
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
						"rule": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"vuln_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
						},
					},
				},
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
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
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"applicationreal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"locationreal": &schema.Schema{
							Type:     schema.TypeString,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"osreal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"protocolreal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"severityreal": &schema.Schema{
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
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exempt_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceObjectGlobalIpsSensorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGlobalIpsSensor(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGlobalIpsSensor resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGlobalIpsSensor(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGlobalIpsSensor resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGlobalIpsSensorRead(d, m)
}

func resourceObjectGlobalIpsSensorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGlobalIpsSensor(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGlobalIpsSensor resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGlobalIpsSensor(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGlobalIpsSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGlobalIpsSensorRead(d, m)
}

func resourceObjectGlobalIpsSensorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectGlobalIpsSensor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGlobalIpsSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGlobalIpsSensorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectGlobalIpsSensor(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGlobalIpsSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGlobalIpsSensor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGlobalIpsSensor resource from API: %v", err)
	}
	return nil
}

func flattenObjectGlobalIpsSensorBlockMaliciousUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGlobalIpsSensorEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := i["cve"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesCve(i["cve"], d, pre_append)
			tmp["cve"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Cve")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_action"
		if _, ok := i["default-action"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesDefaultAction(i["default-action"], d, pre_append)
			tmp["default_action"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-DefaultAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_status"
		if _, ok := i["default-status"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesDefaultStatus(i["default-status"], d, pre_append)
			tmp["default_status"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-DefaultStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := i["exempt-ip"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesExemptIp(i["exempt-ip"], d, pre_append)
			tmp["exempt_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-ExemptIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_modified"
		if _, ok := i["last-modified"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesLastModified(i["last-modified"], d, pre_append)
			tmp["last_modified"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-LastModified")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := i["location"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesLocation(i["location"], d, pre_append)
			tmp["location"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Location")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := i["log-attack-context"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesLogAttackContext(i["log-attack-context"], d, pre_append)
			tmp["log_attack_context"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-LogAttackContext")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesOs(i["os"], d, pre_append)
			tmp["os"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Os")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "position"
		if _, ok := i["position"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesPosition(i["position"], d, pre_append)
			tmp["position"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Position")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesRateCount(i["rate-count"], d, pre_append)
			tmp["rate_count"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-RateCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesRateDuration(i["rate-duration"], d, pre_append)
			tmp["rate_duration"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-RateDuration")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesRateMode(i["rate-mode"], d, pre_append)
			tmp["rate_mode"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-RateMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesRateTrack(i["rate-track"], d, pre_append)
			tmp["rate_track"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-RateTrack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := i["rule"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesRule(i["rule"], d, pre_append)
			tmp["rule"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Rule")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-Tags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vuln_type"
		if _, ok := i["vuln-type"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesVulnType(i["vuln-type"], d, pre_append)
			tmp["vuln_type"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Entries-VulnType")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectGlobalIpsSensorEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesCve(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesDefaultStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGlobalIpsSensorEntriesExemptIpDstIp(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorEntries-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesExemptIpId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorEntries-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesExemptIpSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorEntries-ExemptIp-SrcIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectGlobalIpsSensorEntriesExemptIpDstIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIpSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLastModified(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLogAttackContext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesPosition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesVulnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectGlobalIpsSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGlobalIpsSensorFilterAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			v := flattenObjectGlobalIpsSensorFilterApplication(i["application"], d, pre_append)
			tmp["application"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Application")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "applicationreal"
		if _, ok := i["application(real)"]; ok {
			v := flattenObjectGlobalIpsSensorFilterApplicationReal(i["application(real)"], d, pre_append)
			tmp["applicationreal"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-ApplicationReal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := i["location"]; ok {
			v := flattenObjectGlobalIpsSensorFilterLocation(i["location"], d, pre_append)
			tmp["location"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Location")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "locationreal"
		if _, ok := i["location(real)"]; ok {
			v := flattenObjectGlobalIpsSensorFilterLocationReal(i["location(real)"], d, pre_append)
			tmp["locationreal"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-LocationReal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectGlobalIpsSensorFilterLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenObjectGlobalIpsSensorFilterLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectGlobalIpsSensorFilterName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {
			v := flattenObjectGlobalIpsSensorFilterOs(i["os"], d, pre_append)
			tmp["os"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Os")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osreal"
		if _, ok := i["os(real)"]; ok {
			v := flattenObjectGlobalIpsSensorFilterOsReal(i["os(real)"], d, pre_append)
			tmp["osreal"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-OsReal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectGlobalIpsSensorFilterProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocolreal"
		if _, ok := i["protocol(real)"]; ok {
			v := flattenObjectGlobalIpsSensorFilterProtocolReal(i["protocol(real)"], d, pre_append)
			tmp["protocolreal"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-ProtocolReal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenObjectGlobalIpsSensorFilterQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenObjectGlobalIpsSensorFilterQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenObjectGlobalIpsSensorFilterQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenObjectGlobalIpsSensorFilterSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severityreal"
		if _, ok := i["severity(real)"]; ok {
			v := flattenObjectGlobalIpsSensorFilterSeverityReal(i["severity(real)"], d, pre_append)
			tmp["severityreal"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-SeverityReal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectGlobalIpsSensorFilterStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Filter-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectGlobalIpsSensorFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorFilterApplicationReal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorFilterLocationReal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorFilterOsReal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorFilterProtocolReal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorFilterSeverityReal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorFilterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGlobalIpsSensorOverrideAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := i["exempt-ip"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideExemptIp(i["exempt-ip"], d, pre_append)
			tmp["exempt_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-ExemptIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideLogPacket(i["log-packet"], d, pre_append)
			tmp["log_packet"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-LogPacket")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := i["rule-id"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideRuleId(i["rule-id"], d, pre_append)
			tmp["rule_id"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-RuleId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensor-Override-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectGlobalIpsSensorOverrideAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideExemptIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGlobalIpsSensorOverrideExemptIpDstIp(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorOverride-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideExemptIpId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorOverride-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenObjectGlobalIpsSensorOverrideExemptIpSrcIp(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorOverride-ExemptIp-SrcIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectGlobalIpsSensorOverrideExemptIpDstIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideExemptIpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideExemptIpSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectGlobalIpsSensor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("block_malicious_url", flattenObjectGlobalIpsSensorBlockMaliciousUrl(o["block-malicious-url"], d, "block_malicious_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-malicious-url"], "ObjectGlobalIpsSensor-BlockMaliciousUrl"); ok {
			if err = d.Set("block_malicious_url", vv); err != nil {
				return fmt.Errorf("Error reading block_malicious_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_malicious_url: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectGlobalIpsSensorComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectGlobalIpsSensor-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectGlobalIpsSensorEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectGlobalIpsSensor-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectGlobalIpsSensorEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectGlobalIpsSensor-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("extended_log", flattenObjectGlobalIpsSensorExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "ObjectGlobalIpsSensor-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filter", flattenObjectGlobalIpsSensorFilter(o["filter"], d, "filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["filter"], "ObjectGlobalIpsSensor-Filter"); ok {
				if err = d.Set("filter", vv); err != nil {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter"); ok {
			if err = d.Set("filter", flattenObjectGlobalIpsSensorFilter(o["filter"], d, "filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["filter"], "ObjectGlobalIpsSensor-Filter"); ok {
					if err = d.Set("filter", vv); err != nil {
						return fmt.Errorf("Error reading filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("log", flattenObjectGlobalIpsSensorLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectGlobalIpsSensor-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectGlobalIpsSensorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGlobalIpsSensor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("override", flattenObjectGlobalIpsSensorOverride(o["override"], d, "override")); err != nil {
			if vv, ok := fortiAPIPatch(o["override"], "ObjectGlobalIpsSensor-Override"); ok {
				if err = d.Set("override", vv); err != nil {
					return fmt.Errorf("Error reading override: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("override"); ok {
			if err = d.Set("override", flattenObjectGlobalIpsSensorOverride(o["override"], d, "override")); err != nil {
				if vv, ok := fortiAPIPatch(o["override"], "ObjectGlobalIpsSensor-Override"); ok {
					if err = d.Set("override", vv); err != nil {
						return fmt.Errorf("Error reading override: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading override: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectGlobalIpsSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectGlobalIpsSensor-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenObjectGlobalIpsSensorScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-botnet-connections"], "ObjectGlobalIpsSensor-ScanBotnetConnections"); ok {
			if err = d.Set("scan_botnet_connections", vv); err != nil {
				return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	return nil
}

func flattenObjectGlobalIpsSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGlobalIpsSensorBlockMaliciousUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectGlobalIpsSensorEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandObjectGlobalIpsSensorEntriesApplication(d, i["application"], pre_append)
		} else {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cve"], _ = expandObjectGlobalIpsSensorEntriesCve(d, i["cve"], pre_append)
		} else {
			tmp["cve"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-action"], _ = expandObjectGlobalIpsSensorEntriesDefaultAction(d, i["default_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default-status"], _ = expandObjectGlobalIpsSensorEntriesDefaultStatus(d, i["default_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exempt-ip"], _ = expandObjectGlobalIpsSensorEntriesExemptIp(d, i["exempt_ip"], pre_append)
		} else {
			tmp["exempt-ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectGlobalIpsSensorEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_modified"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["last-modified"], _ = expandObjectGlobalIpsSensorEntriesLastModified(d, i["last_modified"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["location"], _ = expandObjectGlobalIpsSensorEntriesLocation(d, i["location"], pre_append)
		} else {
			tmp["location"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectGlobalIpsSensorEntriesLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-attack-context"], _ = expandObjectGlobalIpsSensorEntriesLogAttackContext(d, i["log_attack_context"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-packet"], _ = expandObjectGlobalIpsSensorEntriesLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os"], _ = expandObjectGlobalIpsSensorEntriesOs(d, i["os"], pre_append)
		} else {
			tmp["os"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "position"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["position"], _ = expandObjectGlobalIpsSensorEntriesPosition(d, i["position"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectGlobalIpsSensorEntriesProtocol(d, i["protocol"], pre_append)
		} else {
			tmp["protocol"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandObjectGlobalIpsSensorEntriesQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandObjectGlobalIpsSensorEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandObjectGlobalIpsSensorEntriesQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-count"], _ = expandObjectGlobalIpsSensorEntriesRateCount(d, i["rate_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-duration"], _ = expandObjectGlobalIpsSensorEntriesRateDuration(d, i["rate_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-mode"], _ = expandObjectGlobalIpsSensorEntriesRateMode(d, i["rate_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rate-track"], _ = expandObjectGlobalIpsSensorEntriesRateTrack(d, i["rate_track"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule"], _ = expandObjectGlobalIpsSensorEntriesRule(d, i["rule"], pre_append)
		} else {
			tmp["rule"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandObjectGlobalIpsSensorEntriesSeverity(d, i["severity"], pre_append)
		} else {
			tmp["severity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectGlobalIpsSensorEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandObjectGlobalIpsSensorEntriesTags(d, i["tags"], pre_append)
		} else {
			tmp["tags"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vuln_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vuln-type"], _ = expandObjectGlobalIpsSensorEntriesVulnType(d, i["vuln_type"], pre_append)
		} else {
			tmp["vuln-type"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectGlobalIpsSensorEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesCve(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesDefaultStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-ip"], _ = expandObjectGlobalIpsSensorEntriesExemptIpDstIp(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectGlobalIpsSensorEntriesExemptIpId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandObjectGlobalIpsSensorEntriesExemptIpSrcIp(d, i["src_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpDstIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLastModified(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLogAttackContext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesPosition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesVulnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectGlobalIpsSensorFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application"], _ = expandObjectGlobalIpsSensorFilterApplication(d, i["application"], pre_append)
		} else {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "applicationreal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["application(real)"], _ = expandObjectGlobalIpsSensorFilterApplicationReal(d, i["applicationreal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["location"], _ = expandObjectGlobalIpsSensorFilterLocation(d, i["location"], pre_append)
		} else {
			tmp["location"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "locationreal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["location(real)"], _ = expandObjectGlobalIpsSensorFilterLocationReal(d, i["locationreal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectGlobalIpsSensorFilterLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-packet"], _ = expandObjectGlobalIpsSensorFilterLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectGlobalIpsSensorFilterName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os"], _ = expandObjectGlobalIpsSensorFilterOs(d, i["os"], pre_append)
		} else {
			tmp["os"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osreal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os(real)"], _ = expandObjectGlobalIpsSensorFilterOsReal(d, i["osreal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectGlobalIpsSensorFilterProtocol(d, i["protocol"], pre_append)
		} else {
			tmp["protocol"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocolreal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol(real)"], _ = expandObjectGlobalIpsSensorFilterProtocolReal(d, i["protocolreal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandObjectGlobalIpsSensorFilterQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandObjectGlobalIpsSensorFilterQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandObjectGlobalIpsSensorFilterQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandObjectGlobalIpsSensorFilterSeverity(d, i["severity"], pre_append)
		} else {
			tmp["severity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severityreal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity(real)"], _ = expandObjectGlobalIpsSensorFilterSeverityReal(d, i["severityreal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectGlobalIpsSensorFilterStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectGlobalIpsSensorFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorFilterApplicationReal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorFilterLocationReal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorFilterOsReal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorFilterProtocolReal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorFilterSeverityReal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorFilterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectGlobalIpsSensorOverrideAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exempt-ip"], _ = expandObjectGlobalIpsSensorOverrideExemptIp(d, i["exempt_ip"], pre_append)
		} else {
			tmp["exempt-ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectGlobalIpsSensorOverrideLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log-packet"], _ = expandObjectGlobalIpsSensorOverrideLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandObjectGlobalIpsSensorOverrideQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandObjectGlobalIpsSensorOverrideQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandObjectGlobalIpsSensorOverrideQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rule-id"], _ = expandObjectGlobalIpsSensorOverrideRuleId(d, i["rule_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectGlobalIpsSensorOverrideStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectGlobalIpsSensorOverrideAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideExemptIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-ip"], _ = expandObjectGlobalIpsSensorOverrideExemptIpDstIp(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectGlobalIpsSensorOverrideExemptIpId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandObjectGlobalIpsSensorOverrideExemptIpSrcIp(d, i["src_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectGlobalIpsSensorOverrideExemptIpDstIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideExemptIpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideExemptIpSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGlobalIpsSensor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_malicious_url"); ok || d.HasChange("block_malicious_url") {
		t, err := expandObjectGlobalIpsSensorBlockMaliciousUrl(d, v, "block_malicious_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malicious-url"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectGlobalIpsSensorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectGlobalIpsSensorEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandObjectGlobalIpsSensorExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandObjectGlobalIpsSensorFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectGlobalIpsSensorLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectGlobalIpsSensorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok || d.HasChange("override") {
		t, err := expandObjectGlobalIpsSensorOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandObjectGlobalIpsSensorReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok || d.HasChange("scan_botnet_connections") {
		t, err := expandObjectGlobalIpsSensorScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	return &obj, nil
}
