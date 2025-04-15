// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Application list entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectApplicationListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectApplicationListEntriesCreate,
		Read:   resourceObjectApplicationListEntriesRead,
		Update: resourceObjectApplicationListEntriesUpdate,
		Delete: resourceObjectApplicationListEntriesDelete,

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
			"list": &schema.Schema{
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
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"behavior": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclusion": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"protocols": &schema.Schema{
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
			"risk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceObjectApplicationListEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	paradict["list"] = list

	obj, err := getObjectObjectApplicationListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectApplicationListEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectApplicationListEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListEntriesRead(d, m)
}

func resourceObjectApplicationListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	paradict["list"] = list

	obj, err := getObjectObjectApplicationListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectApplicationListEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectApplicationListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectApplicationListEntriesRead(d, m)
}

func resourceObjectApplicationListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	paradict["list"] = list

	wsParams["adom"] = adomv

	err = c.DeleteObjectApplicationListEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectApplicationListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectApplicationListEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	list := d.Get("list").(string)
	if list == "" {
		list = importOptionChecking(m.(*FortiClient).Cfg, "list")
		if list == "" {
			return fmt.Errorf("Parameter list is missing")
		}
		if err = d.Set("list", list); err != nil {
			return fmt.Errorf("Error set params list: %v", err)
		}
	}
	paradict["list"] = list

	o, err := c.ReadObjectApplicationListEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectApplicationListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectApplicationListEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectApplicationListEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesBehavior2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectApplicationListEntriesExclusion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesLogPacket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParameters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectApplicationListEntriesParametersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntries-Parameters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembers2edl(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntries-Parameters-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectApplicationListEntriesParametersValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntries-Parameters-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectApplicationListEntriesParametersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembers2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectApplicationListEntriesParametersMembersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectApplicationListEntriesParametersMembersValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectApplicationListEntriesParameters-Members-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectApplicationListEntriesParametersMembersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersMembersValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesParametersValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesPerIpShaper2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectApplicationListEntriesPopularity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesProtocols2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesRisk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesSessionTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectApplicationListEntriesShaper2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectApplicationListEntriesShaperReverse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectApplicationListEntriesSubCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectApplicationListEntriesTechnology2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectApplicationListEntriesVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectApplicationListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenObjectApplicationListEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectApplicationListEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenObjectApplicationListEntriesApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ObjectApplicationListEntries-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("behavior", flattenObjectApplicationListEntriesBehavior2edl(o["behavior"], d, "behavior")); err != nil {
		if vv, ok := fortiAPIPatch(o["behavior"], "ObjectApplicationListEntries-Behavior"); ok {
			if err = d.Set("behavior", vv); err != nil {
				return fmt.Errorf("Error reading behavior: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectApplicationListEntriesCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectApplicationListEntries-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("exclusion", flattenObjectApplicationListEntriesExclusion2edl(o["exclusion"], d, "exclusion")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclusion"], "ObjectApplicationListEntries-Exclusion"); ok {
			if err = d.Set("exclusion", vv); err != nil {
				return fmt.Errorf("Error reading exclusion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclusion: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectApplicationListEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectApplicationListEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectApplicationListEntriesLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectApplicationListEntries-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenObjectApplicationListEntriesLogPacket2edl(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "ObjectApplicationListEntries-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("parameters", flattenObjectApplicationListEntriesParameters2edl(o["parameters"], d, "parameters")); err != nil {
			if vv, ok := fortiAPIPatch(o["parameters"], "ObjectApplicationListEntries-Parameters"); ok {
				if err = d.Set("parameters", vv); err != nil {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading parameters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameters"); ok {
			if err = d.Set("parameters", flattenObjectApplicationListEntriesParameters2edl(o["parameters"], d, "parameters")); err != nil {
				if vv, ok := fortiAPIPatch(o["parameters"], "ObjectApplicationListEntries-Parameters"); ok {
					if err = d.Set("parameters", vv); err != nil {
						return fmt.Errorf("Error reading parameters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			}
		}
	}

	if err = d.Set("per_ip_shaper", flattenObjectApplicationListEntriesPerIpShaper2edl(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "ObjectApplicationListEntries-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("popularity", flattenObjectApplicationListEntriesPopularity2edl(o["popularity"], d, "popularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["popularity"], "ObjectApplicationListEntries-Popularity"); ok {
			if err = d.Set("popularity", vv); err != nil {
				return fmt.Errorf("Error reading popularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	if err = d.Set("protocols", flattenObjectApplicationListEntriesProtocols2edl(o["protocols"], d, "protocols")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocols"], "ObjectApplicationListEntries-Protocols"); ok {
			if err = d.Set("protocols", vv); err != nil {
				return fmt.Errorf("Error reading protocols: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocols: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenObjectApplicationListEntriesQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "ObjectApplicationListEntries-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenObjectApplicationListEntriesQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "ObjectApplicationListEntries-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenObjectApplicationListEntriesQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "ObjectApplicationListEntries-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("rate_count", flattenObjectApplicationListEntriesRateCount2edl(o["rate-count"], d, "rate_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-count"], "ObjectApplicationListEntries-RateCount"); ok {
			if err = d.Set("rate_count", vv); err != nil {
				return fmt.Errorf("Error reading rate_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_count: %v", err)
		}
	}

	if err = d.Set("rate_duration", flattenObjectApplicationListEntriesRateDuration2edl(o["rate-duration"], d, "rate_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-duration"], "ObjectApplicationListEntries-RateDuration"); ok {
			if err = d.Set("rate_duration", vv); err != nil {
				return fmt.Errorf("Error reading rate_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_duration: %v", err)
		}
	}

	if err = d.Set("rate_mode", flattenObjectApplicationListEntriesRateMode2edl(o["rate-mode"], d, "rate_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-mode"], "ObjectApplicationListEntries-RateMode"); ok {
			if err = d.Set("rate_mode", vv); err != nil {
				return fmt.Errorf("Error reading rate_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_mode: %v", err)
		}
	}

	if err = d.Set("rate_track", flattenObjectApplicationListEntriesRateTrack2edl(o["rate-track"], d, "rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-track"], "ObjectApplicationListEntries-RateTrack"); ok {
			if err = d.Set("rate_track", vv); err != nil {
				return fmt.Errorf("Error reading rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_track: %v", err)
		}
	}

	if err = d.Set("risk", flattenObjectApplicationListEntriesRisk2edl(o["risk"], d, "risk")); err != nil {
		if vv, ok := fortiAPIPatch(o["risk"], "ObjectApplicationListEntries-Risk"); ok {
			if err = d.Set("risk", vv); err != nil {
				return fmt.Errorf("Error reading risk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading risk: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenObjectApplicationListEntriesSessionTtl2edl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "ObjectApplicationListEntries-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("shaper", flattenObjectApplicationListEntriesShaper2edl(o["shaper"], d, "shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["shaper"], "ObjectApplicationListEntries-Shaper"); ok {
			if err = d.Set("shaper", vv); err != nil {
				return fmt.Errorf("Error reading shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shaper: %v", err)
		}
	}

	if err = d.Set("shaper_reverse", flattenObjectApplicationListEntriesShaperReverse2edl(o["shaper-reverse"], d, "shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["shaper-reverse"], "ObjectApplicationListEntries-ShaperReverse"); ok {
			if err = d.Set("shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shaper_reverse: %v", err)
		}
	}

	if err = d.Set("sub_category", flattenObjectApplicationListEntriesSubCategory2edl(o["sub-category"], d, "sub_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-category"], "ObjectApplicationListEntries-SubCategory"); ok {
			if err = d.Set("sub_category", vv); err != nil {
				return fmt.Errorf("Error reading sub_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_category: %v", err)
		}
	}

	if err = d.Set("technology", flattenObjectApplicationListEntriesTechnology2edl(o["technology"], d, "technology")); err != nil {
		if vv, ok := fortiAPIPatch(o["technology"], "ObjectApplicationListEntries-Technology"); ok {
			if err = d.Set("technology", vv); err != nil {
				return fmt.Errorf("Error reading technology: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("vendor", flattenObjectApplicationListEntriesVendor2edl(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "ObjectApplicationListEntries-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenObjectApplicationListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectApplicationListEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesBehavior2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectApplicationListEntriesExclusion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesLogPacket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParameters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectApplicationListEntriesParametersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectApplicationListEntriesParametersMembers2edl(d, i["members"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["members"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectApplicationListEntriesParametersValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListEntriesParametersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectApplicationListEntriesParametersMembersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectApplicationListEntriesParametersMembersName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectApplicationListEntriesParametersMembersValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectApplicationListEntriesParametersMembersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersMembersValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesParametersValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesPerIpShaper2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectApplicationListEntriesPopularity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesProtocols2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesRisk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesSessionTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectApplicationListEntriesShaper2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectApplicationListEntriesShaperReverse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectApplicationListEntriesSubCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesTechnology2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectApplicationListEntriesVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectApplicationListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectApplicationListEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandObjectApplicationListEntriesApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("behavior"); ok || d.HasChange("behavior") {
		t, err := expandObjectApplicationListEntriesBehavior2edl(d, v, "behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandObjectApplicationListEntriesCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("exclusion"); ok || d.HasChange("exclusion") {
		t, err := expandObjectApplicationListEntriesExclusion2edl(d, v, "exclusion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusion"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectApplicationListEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectApplicationListEntriesLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandObjectApplicationListEntriesLogPacket2edl(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("parameters"); ok || d.HasChange("parameters") {
		t, err := expandObjectApplicationListEntriesParameters2edl(d, v, "parameters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameters"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandObjectApplicationListEntriesPerIpShaper2edl(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("popularity"); ok || d.HasChange("popularity") {
		t, err := expandObjectApplicationListEntriesPopularity2edl(d, v, "popularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	if v, ok := d.GetOk("protocols"); ok || d.HasChange("protocols") {
		t, err := expandObjectApplicationListEntriesProtocols2edl(d, v, "protocols")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocols"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandObjectApplicationListEntriesQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandObjectApplicationListEntriesQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandObjectApplicationListEntriesQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("rate_count"); ok || d.HasChange("rate_count") {
		t, err := expandObjectApplicationListEntriesRateCount2edl(d, v, "rate_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-count"] = t
		}
	}

	if v, ok := d.GetOk("rate_duration"); ok || d.HasChange("rate_duration") {
		t, err := expandObjectApplicationListEntriesRateDuration2edl(d, v, "rate_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-duration"] = t
		}
	}

	if v, ok := d.GetOk("rate_mode"); ok || d.HasChange("rate_mode") {
		t, err := expandObjectApplicationListEntriesRateMode2edl(d, v, "rate_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-mode"] = t
		}
	}

	if v, ok := d.GetOk("rate_track"); ok || d.HasChange("rate_track") {
		t, err := expandObjectApplicationListEntriesRateTrack2edl(d, v, "rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-track"] = t
		}
	}

	if v, ok := d.GetOk("risk"); ok || d.HasChange("risk") {
		t, err := expandObjectApplicationListEntriesRisk2edl(d, v, "risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandObjectApplicationListEntriesSessionTtl2edl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("shaper"); ok || d.HasChange("shaper") {
		t, err := expandObjectApplicationListEntriesShaper2edl(d, v, "shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaper"] = t
		}
	}

	if v, ok := d.GetOk("shaper_reverse"); ok || d.HasChange("shaper_reverse") {
		t, err := expandObjectApplicationListEntriesShaperReverse2edl(d, v, "shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("sub_category"); ok || d.HasChange("sub_category") {
		t, err := expandObjectApplicationListEntriesSubCategory2edl(d, v, "sub_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-category"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok || d.HasChange("technology") {
		t, err := expandObjectApplicationListEntriesTechnology2edl(d, v, "technology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandObjectApplicationListEntriesVendor2edl(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
