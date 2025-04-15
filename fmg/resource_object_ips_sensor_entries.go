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

func resourceObjectIpsSensorEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIpsSensorEntriesCreate,
		Read:   resourceObjectIpsSensorEntriesRead,
		Update: resourceObjectIpsSensorEntriesUpdate,
		Delete: resourceObjectIpsSensorEntriesDelete,

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
				Type:     schema.TypeString,
				Optional: true,
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

func resourceObjectIpsSensorEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectIpsSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsSensorEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIpsSensorEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsSensorEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIpsSensorEntriesRead(d, m)
}

func resourceObjectIpsSensorEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	obj, err := getObjectObjectIpsSensorEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsSensorEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIpsSensorEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsSensorEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIpsSensorEntriesRead(d, m)
}

func resourceObjectIpsSensorEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	paradict["sensor"] = sensor

	wsParams["adom"] = adomv

	err = c.DeleteObjectIpsSensorEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIpsSensorEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIpsSensorEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectIpsSensorEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsSensorEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIpsSensorEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsSensorEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectIpsSensorEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesCve2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesDefaultAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesDefaultStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesExemptIp2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectIpsSensorEntriesExemptIpDstIp2edl(i["dst-ip"], d, pre_append)
			tmp["dst_ip"] = fortiAPISubPartPatch(v, "ObjectIpsSensorEntries-ExemptIp-DstIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectIpsSensorEntriesExemptIpId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectIpsSensorEntries-ExemptIp-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {
			v := flattenObjectIpsSensorEntriesExemptIpSrcIp2edl(i["src-ip"], d, pre_append)
			tmp["src_ip"] = fortiAPISubPartPatch(v, "ObjectIpsSensorEntries-ExemptIp-SrcIp")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectIpsSensorEntriesExemptIpDstIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIpsSensorEntriesExemptIpId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesExemptIpSrcIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIpsSensorEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesLastModified2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectIpsSensorEntriesLocation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesLogAttackContext2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesLogPacket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesOs2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRateCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRateDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRateMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesRule2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIpsSensorEntriesSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIpsSensorEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesVulnType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectObjectIpsSensorEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenObjectIpsSensorEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectIpsSensorEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenObjectIpsSensorEntriesApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ObjectIpsSensorEntries-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("cve", flattenObjectIpsSensorEntriesCve2edl(o["cve"], d, "cve")); err != nil {
		if vv, ok := fortiAPIPatch(o["cve"], "ObjectIpsSensorEntries-Cve"); ok {
			if err = d.Set("cve", vv); err != nil {
				return fmt.Errorf("Error reading cve: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cve: %v", err)
		}
	}

	if err = d.Set("default_action", flattenObjectIpsSensorEntriesDefaultAction2edl(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "ObjectIpsSensorEntries-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if err = d.Set("default_status", flattenObjectIpsSensorEntriesDefaultStatus2edl(o["default-status"], d, "default_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-status"], "ObjectIpsSensorEntries-DefaultStatus"); ok {
			if err = d.Set("default_status", vv); err != nil {
				return fmt.Errorf("Error reading default_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("exempt_ip", flattenObjectIpsSensorEntriesExemptIp2edl(o["exempt-ip"], d, "exempt_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["exempt-ip"], "ObjectIpsSensorEntries-ExemptIp"); ok {
				if err = d.Set("exempt_ip", vv); err != nil {
					return fmt.Errorf("Error reading exempt_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading exempt_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exempt_ip"); ok {
			if err = d.Set("exempt_ip", flattenObjectIpsSensorEntriesExemptIp2edl(o["exempt-ip"], d, "exempt_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["exempt-ip"], "ObjectIpsSensorEntries-ExemptIp"); ok {
					if err = d.Set("exempt_ip", vv); err != nil {
						return fmt.Errorf("Error reading exempt_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading exempt_ip: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectIpsSensorEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectIpsSensorEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("last_modified", flattenObjectIpsSensorEntriesLastModified2edl(o["last-modified"], d, "last_modified")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-modified"], "ObjectIpsSensorEntries-LastModified"); ok {
			if err = d.Set("last_modified", vv); err != nil {
				return fmt.Errorf("Error reading last_modified: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_modified: %v", err)
		}
	}

	if err = d.Set("location", flattenObjectIpsSensorEntriesLocation2edl(o["location"], d, "location")); err != nil {
		if vv, ok := fortiAPIPatch(o["location"], "ObjectIpsSensorEntries-Location"); ok {
			if err = d.Set("location", vv); err != nil {
				return fmt.Errorf("Error reading location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectIpsSensorEntriesLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectIpsSensorEntries-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_attack_context", flattenObjectIpsSensorEntriesLogAttackContext2edl(o["log-attack-context"], d, "log_attack_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-attack-context"], "ObjectIpsSensorEntries-LogAttackContext"); ok {
			if err = d.Set("log_attack_context", vv); err != nil {
				return fmt.Errorf("Error reading log_attack_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_attack_context: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenObjectIpsSensorEntriesLogPacket2edl(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "ObjectIpsSensorEntries-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("os", flattenObjectIpsSensorEntriesOs2edl(o["os"], d, "os")); err != nil {
		if vv, ok := fortiAPIPatch(o["os"], "ObjectIpsSensorEntries-Os"); ok {
			if err = d.Set("os", vv); err != nil {
				return fmt.Errorf("Error reading os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectIpsSensorEntriesProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectIpsSensorEntries-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenObjectIpsSensorEntriesQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "ObjectIpsSensorEntries-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenObjectIpsSensorEntriesQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "ObjectIpsSensorEntries-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenObjectIpsSensorEntriesQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "ObjectIpsSensorEntries-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("rate_count", flattenObjectIpsSensorEntriesRateCount2edl(o["rate-count"], d, "rate_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-count"], "ObjectIpsSensorEntries-RateCount"); ok {
			if err = d.Set("rate_count", vv); err != nil {
				return fmt.Errorf("Error reading rate_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_count: %v", err)
		}
	}

	if err = d.Set("rate_duration", flattenObjectIpsSensorEntriesRateDuration2edl(o["rate-duration"], d, "rate_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-duration"], "ObjectIpsSensorEntries-RateDuration"); ok {
			if err = d.Set("rate_duration", vv); err != nil {
				return fmt.Errorf("Error reading rate_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_duration: %v", err)
		}
	}

	if err = d.Set("rate_mode", flattenObjectIpsSensorEntriesRateMode2edl(o["rate-mode"], d, "rate_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-mode"], "ObjectIpsSensorEntries-RateMode"); ok {
			if err = d.Set("rate_mode", vv); err != nil {
				return fmt.Errorf("Error reading rate_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_mode: %v", err)
		}
	}

	if err = d.Set("rate_track", flattenObjectIpsSensorEntriesRateTrack2edl(o["rate-track"], d, "rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-track"], "ObjectIpsSensorEntries-RateTrack"); ok {
			if err = d.Set("rate_track", vv); err != nil {
				return fmt.Errorf("Error reading rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_track: %v", err)
		}
	}

	if err = d.Set("rule", flattenObjectIpsSensorEntriesRule2edl(o["rule"], d, "rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule"], "ObjectIpsSensorEntries-Rule"); ok {
			if err = d.Set("rule", vv); err != nil {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectIpsSensorEntriesSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectIpsSensorEntries-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectIpsSensorEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectIpsSensorEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vuln_type", flattenObjectIpsSensorEntriesVulnType2edl(o["vuln-type"], d, "vuln_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vuln-type"], "ObjectIpsSensorEntries-VulnType"); ok {
			if err = d.Set("vuln_type", vv); err != nil {
				return fmt.Errorf("Error reading vuln_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vuln_type: %v", err)
		}
	}

	return nil
}

func flattenObjectIpsSensorEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIpsSensorEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesCve2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesDefaultAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesDefaultStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesExemptIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["dst-ip"], _ = expandObjectIpsSensorEntriesExemptIpDstIp2edl(d, i["dst_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectIpsSensorEntriesExemptIpId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ip"], _ = expandObjectIpsSensorEntriesExemptIpSrcIp2edl(d, i["src_ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectIpsSensorEntriesExemptIpDstIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIpsSensorEntriesExemptIpId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesExemptIpSrcIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIpsSensorEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesLastModified2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIpsSensorEntriesLocation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesLogAttackContext2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesLogPacket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesOs2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIpsSensorEntriesSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIpsSensorEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesVulnType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectObjectIpsSensorEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectIpsSensorEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandObjectIpsSensorEntriesApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("cve"); ok || d.HasChange("cve") {
		t, err := expandObjectIpsSensorEntriesCve2edl(d, v, "cve")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cve"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok || d.HasChange("default_action") {
		t, err := expandObjectIpsSensorEntriesDefaultAction2edl(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("default_status"); ok || d.HasChange("default_status") {
		t, err := expandObjectIpsSensorEntriesDefaultStatus2edl(d, v, "default_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-status"] = t
		}
	}

	if v, ok := d.GetOk("exempt_ip"); ok || d.HasChange("exempt_ip") {
		t, err := expandObjectIpsSensorEntriesExemptIp2edl(d, v, "exempt_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exempt-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectIpsSensorEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("last_modified"); ok || d.HasChange("last_modified") {
		t, err := expandObjectIpsSensorEntriesLastModified2edl(d, v, "last_modified")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-modified"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok || d.HasChange("location") {
		t, err := expandObjectIpsSensorEntriesLocation2edl(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectIpsSensorEntriesLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_attack_context"); ok || d.HasChange("log_attack_context") {
		t, err := expandObjectIpsSensorEntriesLogAttackContext2edl(d, v, "log_attack_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-attack-context"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandObjectIpsSensorEntriesLogPacket2edl(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok || d.HasChange("os") {
		t, err := expandObjectIpsSensorEntriesOs2edl(d, v, "os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectIpsSensorEntriesProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandObjectIpsSensorEntriesQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandObjectIpsSensorEntriesQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandObjectIpsSensorEntriesQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("rate_count"); ok || d.HasChange("rate_count") {
		t, err := expandObjectIpsSensorEntriesRateCount2edl(d, v, "rate_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-count"] = t
		}
	}

	if v, ok := d.GetOk("rate_duration"); ok || d.HasChange("rate_duration") {
		t, err := expandObjectIpsSensorEntriesRateDuration2edl(d, v, "rate_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-duration"] = t
		}
	}

	if v, ok := d.GetOk("rate_mode"); ok || d.HasChange("rate_mode") {
		t, err := expandObjectIpsSensorEntriesRateMode2edl(d, v, "rate_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-mode"] = t
		}
	}

	if v, ok := d.GetOk("rate_track"); ok || d.HasChange("rate_track") {
		t, err := expandObjectIpsSensorEntriesRateTrack2edl(d, v, "rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-track"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectIpsSensorEntriesRule2edl(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectIpsSensorEntriesSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectIpsSensorEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vuln_type"); ok || d.HasChange("vuln_type") {
		t, err := expandObjectIpsSensorEntriesVulnType2edl(d, v, "vuln_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vuln-type"] = t
		}
	}

	return &obj, nil
}
