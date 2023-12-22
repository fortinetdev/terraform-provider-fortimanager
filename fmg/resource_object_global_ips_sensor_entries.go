// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPS sensor filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectGlobalIpsSensorEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGlobalIpsSensorEntriesCreate,
		Read:   resourceObjectGlobalIpsSensorEntriesRead,
		Update: resourceObjectGlobalIpsSensorEntriesUpdate,
		Delete: resourceObjectGlobalIpsSensorEntriesDelete,

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
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_status": &schema.Schema{
				Type:     schema.TypeString,
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
			"position": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vuln_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
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

func resourceObjectGlobalIpsSensorEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectGlobalIpsSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGlobalIpsSensorEntries resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGlobalIpsSensorEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGlobalIpsSensorEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectGlobalIpsSensorEntriesRead(d, m)
}

func resourceObjectGlobalIpsSensorEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectGlobalIpsSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGlobalIpsSensorEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGlobalIpsSensorEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGlobalIpsSensorEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectGlobalIpsSensorEntriesRead(d, m)
}

func resourceObjectGlobalIpsSensorEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	err = c.DeleteObjectGlobalIpsSensorEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGlobalIpsSensorEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGlobalIpsSensorEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	paradict["sensor"] = sensor

	o, err := c.ReadObjectGlobalIpsSensorEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGlobalIpsSensorEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGlobalIpsSensorEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGlobalIpsSensorEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectGlobalIpsSensorEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesCve2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesDefaultAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesDefaultStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGlobalIpsSensorEntriesExemptIpDstIp2edl(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorEntries-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesExemptIpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorEntries-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenObjectGlobalIpsSensorEntriesExemptIpSrcIp2edl(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "ObjectGlobalIpsSensorEntries-ExemptIp-SrcIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectGlobalIpsSensorEntriesExemptIpDstIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIpSrcIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLastModified2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLocation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLogAttackContext2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesLogPacket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesOs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesPosition2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesRule2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectGlobalIpsSensorEntriesVulnType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectObjectGlobalIpsSensorEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenObjectGlobalIpsSensorEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectGlobalIpsSensorEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenObjectGlobalIpsSensorEntriesApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ObjectGlobalIpsSensorEntries-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("cve", flattenObjectGlobalIpsSensorEntriesCve2edl(o["cve"], d, "cve")); err != nil {
		if vv, ok := fortiAPIPatch(o["cve"], "ObjectGlobalIpsSensorEntries-Cve"); ok {
			if err = d.Set("cve", vv); err != nil {
				return fmt.Errorf("Error reading cve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cve: %v", err)
		}
	}

	if err = d.Set("default_action", flattenObjectGlobalIpsSensorEntriesDefaultAction2edl(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "ObjectGlobalIpsSensorEntries-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if err = d.Set("default_status", flattenObjectGlobalIpsSensorEntriesDefaultStatus2edl(o["default-status"], d, "default_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-status"], "ObjectGlobalIpsSensorEntries-DefaultStatus"); ok {
			if err = d.Set("default_status", vv); err != nil {
				return fmt.Errorf("Error reading default_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exempt_ip", flattenObjectGlobalIpsSensorEntriesExemptIp2edl(o["exempt-ip"], d, "exempt_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["exempt-ip"], "ObjectGlobalIpsSensorEntries-ExemptIp"); ok {
				if err = d.Set("exempt_ip", vv); err != nil {
					return fmt.Errorf("Error reading exempt_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exempt_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exempt_ip"); ok {
			if err = d.Set("exempt_ip", flattenObjectGlobalIpsSensorEntriesExemptIp2edl(o["exempt-ip"], d, "exempt_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["exempt-ip"], "ObjectGlobalIpsSensorEntries-ExemptIp"); ok {
					if err = d.Set("exempt_ip", vv); err != nil {
						return fmt.Errorf("Error reading exempt_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exempt_ip: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectGlobalIpsSensorEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectGlobalIpsSensorEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("last_modified", flattenObjectGlobalIpsSensorEntriesLastModified2edl(o["last-modified"], d, "last_modified")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-modified"], "ObjectGlobalIpsSensorEntries-LastModified"); ok {
			if err = d.Set("last_modified", vv); err != nil {
				return fmt.Errorf("Error reading last_modified: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_modified: %v", err)
		}
	}

	if err = d.Set("location", flattenObjectGlobalIpsSensorEntriesLocation2edl(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "ObjectGlobalIpsSensorEntries-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectGlobalIpsSensorEntriesLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectGlobalIpsSensorEntries-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_attack_context", flattenObjectGlobalIpsSensorEntriesLogAttackContext2edl(o["log-attack-context"], d, "log_attack_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-attack-context"], "ObjectGlobalIpsSensorEntries-LogAttackContext"); ok {
			if err = d.Set("log_attack_context", vv); err != nil {
				return fmt.Errorf("Error reading log_attack_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_attack_context: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenObjectGlobalIpsSensorEntriesLogPacket2edl(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "ObjectGlobalIpsSensorEntries-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("os", flattenObjectGlobalIpsSensorEntriesOs2edl(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "ObjectGlobalIpsSensorEntries-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("position", flattenObjectGlobalIpsSensorEntriesPosition2edl(o["position"], d, "position")); err != nil {
		if vv, ok := fortiAPIPatch(o["position"], "ObjectGlobalIpsSensorEntries-Position"); ok {
			if err = d.Set("position", vv); err != nil {
				return fmt.Errorf("Error reading position: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading position: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectGlobalIpsSensorEntriesProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectGlobalIpsSensorEntries-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenObjectGlobalIpsSensorEntriesQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "ObjectGlobalIpsSensorEntries-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenObjectGlobalIpsSensorEntriesQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "ObjectGlobalIpsSensorEntries-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenObjectGlobalIpsSensorEntriesQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "ObjectGlobalIpsSensorEntries-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("rate_count", flattenObjectGlobalIpsSensorEntriesRateCount2edl(o["rate-count"], d, "rate_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-count"], "ObjectGlobalIpsSensorEntries-RateCount"); ok {
			if err = d.Set("rate_count", vv); err != nil {
				return fmt.Errorf("Error reading rate_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_count: %v", err)
		}
	}

	if err = d.Set("rate_duration", flattenObjectGlobalIpsSensorEntriesRateDuration2edl(o["rate-duration"], d, "rate_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-duration"], "ObjectGlobalIpsSensorEntries-RateDuration"); ok {
			if err = d.Set("rate_duration", vv); err != nil {
				return fmt.Errorf("Error reading rate_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_duration: %v", err)
		}
	}

	if err = d.Set("rate_mode", flattenObjectGlobalIpsSensorEntriesRateMode2edl(o["rate-mode"], d, "rate_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-mode"], "ObjectGlobalIpsSensorEntries-RateMode"); ok {
			if err = d.Set("rate_mode", vv); err != nil {
				return fmt.Errorf("Error reading rate_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_mode: %v", err)
		}
	}

	if err = d.Set("rate_track", flattenObjectGlobalIpsSensorEntriesRateTrack2edl(o["rate-track"], d, "rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-track"], "ObjectGlobalIpsSensorEntries-RateTrack"); ok {
			if err = d.Set("rate_track", vv); err != nil {
				return fmt.Errorf("Error reading rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_track: %v", err)
		}
	}

	if err = d.Set("rule", flattenObjectGlobalIpsSensorEntriesRule2edl(o["rule"], d, "rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule"], "ObjectGlobalIpsSensorEntries-Rule"); ok {
			if err = d.Set("rule", vv); err != nil {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectGlobalIpsSensorEntriesSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectGlobalIpsSensorEntries-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectGlobalIpsSensorEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectGlobalIpsSensorEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tags", flattenObjectGlobalIpsSensorEntriesTags2edl(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "ObjectGlobalIpsSensorEntries-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	if err = d.Set("vuln_type", flattenObjectGlobalIpsSensorEntriesVulnType2edl(o["vuln-type"], d, "vuln_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vuln-type"], "ObjectGlobalIpsSensorEntries-VulnType"); ok {
			if err = d.Set("vuln_type", vv); err != nil {
				return fmt.Errorf("Error reading vuln_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vuln_type: %v", err)
		}
	}

	return nil
}

func flattenObjectGlobalIpsSensorEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGlobalIpsSensorEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesCve2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesDefaultAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesDefaultStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-ip"], _ = expandObjectGlobalIpsSensorEntriesExemptIpDstIp2edl(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectGlobalIpsSensorEntriesExemptIpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandObjectGlobalIpsSensorEntriesExemptIpSrcIp2edl(d, i["src_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpDstIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpSrcIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLastModified2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLocation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLogAttackContext2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesLogPacket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesOs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesPosition2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectGlobalIpsSensorEntriesVulnType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectObjectGlobalIpsSensorEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectGlobalIpsSensorEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandObjectGlobalIpsSensorEntriesApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("cve"); ok || d.HasChange("cve") {
		t, err := expandObjectGlobalIpsSensorEntriesCve2edl(d, v, "cve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cve"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok || d.HasChange("default_action") {
		t, err := expandObjectGlobalIpsSensorEntriesDefaultAction2edl(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("default_status"); ok || d.HasChange("default_status") {
		t, err := expandObjectGlobalIpsSensorEntriesDefaultStatus2edl(d, v, "default_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-status"] = t
		}
	}

	if v, ok := d.GetOk("exempt_ip"); ok || d.HasChange("exempt_ip") {
		t, err := expandObjectGlobalIpsSensorEntriesExemptIp2edl(d, v, "exempt_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectGlobalIpsSensorEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("last_modified"); ok || d.HasChange("last_modified") {
		t, err := expandObjectGlobalIpsSensorEntriesLastModified2edl(d, v, "last_modified")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-modified"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandObjectGlobalIpsSensorEntriesLocation2edl(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectGlobalIpsSensorEntriesLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_attack_context"); ok || d.HasChange("log_attack_context") {
		t, err := expandObjectGlobalIpsSensorEntriesLogAttackContext2edl(d, v, "log_attack_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-attack-context"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandObjectGlobalIpsSensorEntriesLogPacket2edl(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandObjectGlobalIpsSensorEntriesOs2edl(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("position"); ok || d.HasChange("position") {
		t, err := expandObjectGlobalIpsSensorEntriesPosition2edl(d, v, "position")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["position"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectGlobalIpsSensorEntriesProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandObjectGlobalIpsSensorEntriesQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandObjectGlobalIpsSensorEntriesQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandObjectGlobalIpsSensorEntriesQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("rate_count"); ok || d.HasChange("rate_count") {
		t, err := expandObjectGlobalIpsSensorEntriesRateCount2edl(d, v, "rate_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-count"] = t
		}
	}

	if v, ok := d.GetOk("rate_duration"); ok || d.HasChange("rate_duration") {
		t, err := expandObjectGlobalIpsSensorEntriesRateDuration2edl(d, v, "rate_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-duration"] = t
		}
	}

	if v, ok := d.GetOk("rate_mode"); ok || d.HasChange("rate_mode") {
		t, err := expandObjectGlobalIpsSensorEntriesRateMode2edl(d, v, "rate_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-mode"] = t
		}
	}

	if v, ok := d.GetOk("rate_track"); ok || d.HasChange("rate_track") {
		t, err := expandObjectGlobalIpsSensorEntriesRateTrack2edl(d, v, "rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-track"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectGlobalIpsSensorEntriesRule2edl(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectGlobalIpsSensorEntriesSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectGlobalIpsSensorEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandObjectGlobalIpsSensorEntriesTags2edl(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	if v, ok := d.GetOk("vuln_type"); ok || d.HasChange("vuln_type") {
		t, err := expandObjectGlobalIpsSensorEntriesVulnType2edl(d, v, "vuln_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vuln-type"] = t
		}
	}

	return &obj, nil
}
