// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Show Internet Service application.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceUpdate,
		Read:   resourceObjectFirewallInternetServiceRead,
		Update: resourceObjectFirewallInternetServiceUpdate,
		Delete: resourceObjectFirewallInternetServiceDelete,

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
			"city": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_range_number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"extra_ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icon_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"jitter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
			},
			"packetloss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reputation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"singularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sld_id": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectFirewallInternetServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallInternetService(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetService resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallInternetService(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallInternetService")

	return resourceObjectFirewallInternetServiceRead(d, m)
}

func resourceObjectFirewallInternetServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallInternetService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallInternetService(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetService resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceCityOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServiceCountryOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServiceDatabaseOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceDirectionOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryOfia(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallInternetServiceEntryIdOfia(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallInternetService-Entry-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_number"
		if _, ok := i["ip-number"]; ok {
			v := flattenObjectFirewallInternetServiceEntryIpNumberOfia(i["ip-number"], d, pre_append)
			tmp["ip_number"] = fortiAPISubPartPatch(v, "ObjectFirewallInternetService-Entry-IpNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_range_number"
		if _, ok := i["ip-range-number"]; ok {
			v := flattenObjectFirewallInternetServiceEntryIpRangeNumberOfia(i["ip-range-number"], d, pre_append)
			tmp["ip_range_number"] = fortiAPISubPartPatch(v, "ObjectFirewallInternetService-Entry-IpRangeNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectFirewallInternetServiceEntryPortOfia(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectFirewallInternetService-Entry-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectFirewallInternetServiceEntryProtocolOfia(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectFirewallInternetService-Entry-Protocol")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallInternetServiceEntryIdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryIpNumberOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryIpRangeNumberOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceEntryPortOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServiceEntryProtocolOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceExtraIpRangeNumberOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIconIdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIpNumberOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceIpRangeNumberOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceJitterThresholdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceLatencyThresholdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceNameOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceOffsetOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceObsoleteOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceRegionOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallInternetServicePacketlossThresholdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceReputationOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceSingularityOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceSldIdOfia(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallInternetService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("city", flattenObjectFirewallInternetServiceCityOfia(o["city"], d, "city")); err != nil {
		if vv, ok := fortiAPIPatch(o["city"], "ObjectFirewallInternetService-City"); ok {
			if err = d.Set("city", vv); err != nil {
				return fmt.Errorf("Error reading city: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectFirewallInternetServiceCountryOfia(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectFirewallInternetService-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("database", flattenObjectFirewallInternetServiceDatabaseOfia(o["database"], d, "database")); err != nil {
		if vv, ok := fortiAPIPatch(o["database"], "ObjectFirewallInternetService-Database"); ok {
			if err = d.Set("database", vv); err != nil {
				return fmt.Errorf("Error reading database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("direction", flattenObjectFirewallInternetServiceDirectionOfia(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "ObjectFirewallInternetService-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entry", flattenObjectFirewallInternetServiceEntryOfia(o["entry"], d, "entry")); err != nil {
			if vv, ok := fortiAPIPatch(o["entry"], "ObjectFirewallInternetService-Entry"); ok {
				if err = d.Set("entry", vv); err != nil {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entry"); ok {
			if err = d.Set("entry", flattenObjectFirewallInternetServiceEntryOfia(o["entry"], d, "entry")); err != nil {
				if vv, ok := fortiAPIPatch(o["entry"], "ObjectFirewallInternetService-Entry"); ok {
					if err = d.Set("entry", vv); err != nil {
						return fmt.Errorf("Error reading entry: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			}
		}
	}

	if err = d.Set("extra_ip_range_number", flattenObjectFirewallInternetServiceExtraIpRangeNumberOfia(o["extra-ip-range-number"], d, "extra_ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["extra-ip-range-number"], "ObjectFirewallInternetService-ExtraIpRangeNumber"); ok {
			if err = d.Set("extra_ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
		}
	}

	if err = d.Set("icon_id", flattenObjectFirewallInternetServiceIconIdOfia(o["icon-id"], d, "icon_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon-id"], "ObjectFirewallInternetService-IconId"); ok {
			if err = d.Set("icon_id", vv); err != nil {
				return fmt.Errorf("Error reading icon_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallInternetServiceIdOfia(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallInternetService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_number", flattenObjectFirewallInternetServiceIpNumberOfia(o["ip-number"], d, "ip_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-number"], "ObjectFirewallInternetService-IpNumber"); ok {
			if err = d.Set("ip_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("ip_range_number", flattenObjectFirewallInternetServiceIpRangeNumberOfia(o["ip-range-number"], d, "ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-range-number"], "ObjectFirewallInternetService-IpRangeNumber"); ok {
			if err = d.Set("ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenObjectFirewallInternetServiceJitterThresholdOfia(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "ObjectFirewallInternetService-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenObjectFirewallInternetServiceLatencyThresholdOfia(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "ObjectFirewallInternetService-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallInternetServiceNameOfia(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallInternetService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("offset", flattenObjectFirewallInternetServiceOffsetOfia(o["offset"], d, "offset")); err != nil {
		if vv, ok := fortiAPIPatch(o["offset"], "ObjectFirewallInternetService-Offset"); ok {
			if err = d.Set("offset", vv); err != nil {
				return fmt.Errorf("Error reading offset: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading offset: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenObjectFirewallInternetServiceObsoleteOfia(o["obsolete"], d, "obsolete")); err != nil {
		if vv, ok := fortiAPIPatch(o["obsolete"], "ObjectFirewallInternetService-Obsolete"); ok {
			if err = d.Set("obsolete", vv); err != nil {
				return fmt.Errorf("Error reading obsolete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("region", flattenObjectFirewallInternetServiceRegionOfia(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "ObjectFirewallInternetService-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("packetloss_threshold", flattenObjectFirewallInternetServicePacketlossThresholdOfia(o["packetloss-threshold"], d, "packetloss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packetloss-threshold"], "ObjectFirewallInternetService-PacketlossThreshold"); ok {
			if err = d.Set("packetloss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packetloss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packetloss_threshold: %v", err)
		}
	}

	if err = d.Set("reputation", flattenObjectFirewallInternetServiceReputationOfia(o["reputation"], d, "reputation")); err != nil {
		if vv, ok := fortiAPIPatch(o["reputation"], "ObjectFirewallInternetService-Reputation"); ok {
			if err = d.Set("reputation", vv); err != nil {
				return fmt.Errorf("Error reading reputation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reputation: %v", err)
		}
	}

	if err = d.Set("singularity", flattenObjectFirewallInternetServiceSingularityOfia(o["singularity"], d, "singularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["singularity"], "ObjectFirewallInternetService-Singularity"); ok {
			if err = d.Set("singularity", vv); err != nil {
				return fmt.Errorf("Error reading singularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading singularity: %v", err)
		}
	}

	if err = d.Set("sld_id", flattenObjectFirewallInternetServiceSldIdOfia(o["sld-id"], d, "sld_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sld-id"], "ObjectFirewallInternetService-SldId"); ok {
			if err = d.Set("sld_id", vv); err != nil {
				return fmt.Errorf("Error reading sld_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sld_id: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceCityOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServiceCountryOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServiceDatabaseOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceDirectionOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandObjectFirewallInternetServiceEntryIdOfia(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-number"], _ = expandObjectFirewallInternetServiceEntryIpNumberOfia(d, i["ip_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_range_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-range-number"], _ = expandObjectFirewallInternetServiceEntryIpRangeNumberOfia(d, i["ip_range_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectFirewallInternetServiceEntryPortOfia(d, i["port"], pre_append)
		} else {
			tmp["port"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandObjectFirewallInternetServiceEntryProtocolOfia(d, i["protocol"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallInternetServiceEntryIdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryIpNumberOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryIpRangeNumberOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceEntryPortOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServiceEntryProtocolOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceExtraIpRangeNumberOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIconIdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIpNumberOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceIpRangeNumberOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceJitterThresholdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceLatencyThresholdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceNameOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceOffsetOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceObsoleteOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceRegionOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallInternetServicePacketlossThresholdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceReputationOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceSingularityOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceSldIdOfia(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("city"); ok || d.HasChange("city") {
		t, err := expandObjectFirewallInternetServiceCityOfia(d, v, "city")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandObjectFirewallInternetServiceCountryOfia(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok || d.HasChange("database") {
		t, err := expandObjectFirewallInternetServiceDatabaseOfia(d, v, "database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandObjectFirewallInternetServiceDirectionOfia(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("entry"); ok || d.HasChange("entry") {
		t, err := expandObjectFirewallInternetServiceEntryOfia(d, v, "entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	if v, ok := d.GetOk("extra_ip_range_number"); ok || d.HasChange("extra_ip_range_number") {
		t, err := expandObjectFirewallInternetServiceExtraIpRangeNumberOfia(d, v, "extra_ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("icon_id"); ok || d.HasChange("icon_id") {
		t, err := expandObjectFirewallInternetServiceIconIdOfia(d, v, "icon_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-id"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectFirewallInternetServiceIdOfia(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_number"); ok || d.HasChange("ip_number") {
		t, err := expandObjectFirewallInternetServiceIpNumberOfia(d, v, "ip_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-number"] = t
		}
	}

	if v, ok := d.GetOk("ip_range_number"); ok || d.HasChange("ip_range_number") {
		t, err := expandObjectFirewallInternetServiceIpRangeNumberOfia(d, v, "ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandObjectFirewallInternetServiceJitterThresholdOfia(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandObjectFirewallInternetServiceLatencyThresholdOfia(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallInternetServiceNameOfia(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("offset"); ok || d.HasChange("offset") {
		t, err := expandObjectFirewallInternetServiceOffsetOfia(d, v, "offset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["offset"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok || d.HasChange("obsolete") {
		t, err := expandObjectFirewallInternetServiceObsoleteOfia(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok || d.HasChange("region") {
		t, err := expandObjectFirewallInternetServiceRegionOfia(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("packetloss_threshold"); ok || d.HasChange("packetloss_threshold") {
		t, err := expandObjectFirewallInternetServicePacketlossThresholdOfia(d, v, "packetloss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packetloss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("reputation"); ok || d.HasChange("reputation") {
		t, err := expandObjectFirewallInternetServiceReputationOfia(d, v, "reputation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation"] = t
		}
	}

	if v, ok := d.GetOk("singularity"); ok || d.HasChange("singularity") {
		t, err := expandObjectFirewallInternetServiceSingularityOfia(d, v, "singularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["singularity"] = t
		}
	}

	if v, ok := d.GetOk("sld_id"); ok || d.HasChange("sld_id") {
		t, err := expandObjectFirewallInternetServiceSldIdOfia(d, v, "sld_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sld-id"] = t
		}
	}

	return &obj, nil
}
