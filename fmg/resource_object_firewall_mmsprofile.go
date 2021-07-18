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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed: true,
			},
			"bwordtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"carrier_endpoint_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"carrier_endpoint_prefix_range_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"carrier_endpoint_prefix_range_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"carrier_endpoint_prefix_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"carrierendpointbwltable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mm1_addr_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1_addr_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1_convert_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1_retr_dupe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1_retrieve_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm1comfortamount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm1comfortinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm1oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mm3_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm3oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm4": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mm4_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm4oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm7": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mm7_addr_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm7_addr_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm7_convert_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm7_outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mm7comfortamount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm7comfortinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mm7oversizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mms_antispam_mass_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_av_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_av_oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_av_virus_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_carrier_endpoint_filter_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_checksum_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_checksum_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_notification_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_web_content_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mmsbwordthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"notif_msisdn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msisdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"threshold": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"remove_blocked_const_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
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

func flattenObjectFirewallMmsProfileExtendedUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectFirewallMmsProfileRemoveBlockedConstLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallMmsProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallMmsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

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

	if err = d.Set("extended_utm_log", flattenObjectFirewallMmsProfileExtendedUtmLog(o["extended-utm-log"], d, "extended_utm_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-utm-log"], "ObjectFirewallMmsProfile-ExtendedUtmLog"); ok {
			if err = d.Set("extended_utm_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_utm_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_utm_log: %v", err)
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

func expandObjectFirewallMmsProfileExtendedUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msisdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msisdn"], _ = expandObjectFirewallMmsProfileNotifMsisdnMsisdn(d, i["msisdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold"], _ = expandObjectFirewallMmsProfileNotifMsisdnThreshold(d, i["threshold"], pre_append)
		} else {
			tmp["threshold"] = make([]string, 0)
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

func expandObjectFirewallMmsProfileRemoveBlockedConstLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallMmsProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallMmsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("avnotificationtable"); ok {
		t, err := expandObjectFirewallMmsProfileAvnotificationtable(d, v, "avnotificationtable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["avnotificationtable"] = t
		}
	}

	if v, ok := d.GetOk("bwordtable"); ok {
		t, err := expandObjectFirewallMmsProfileBwordtable(d, v, "bwordtable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bwordtable"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix"); ok {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefix(d, v, "carrier_endpoint_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix_range_max"); ok {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefixRangeMax(d, v, "carrier_endpoint_prefix_range_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix-range-max"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix_range_min"); ok {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefixRangeMin(d, v, "carrier_endpoint_prefix_range_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix-range-min"] = t
		}
	}

	if v, ok := d.GetOk("carrier_endpoint_prefix_string"); ok {
		t, err := expandObjectFirewallMmsProfileCarrierEndpointPrefixString(d, v, "carrier_endpoint_prefix_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier-endpoint-prefix-string"] = t
		}
	}

	if v, ok := d.GetOk("carrierendpointbwltable"); ok {
		t, err := expandObjectFirewallMmsProfileCarrierendpointbwltable(d, v, "carrierendpointbwltable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrierendpointbwltable"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectFirewallMmsProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("extended_utm_log"); ok {
		t, err := expandObjectFirewallMmsProfileExtendedUtmLog(d, v, "extended_utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-utm-log"] = t
		}
	}

	if v, ok := d.GetOk("mm1"); ok {
		t, err := expandObjectFirewallMmsProfileMm1(d, v, "mm1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1"] = t
		}
	}

	if v, ok := d.GetOk("mm1_addr_hdr"); ok {
		t, err := expandObjectFirewallMmsProfileMm1AddrHdr(d, v, "mm1_addr_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-addr-hdr"] = t
		}
	}

	if v, ok := d.GetOk("mm1_addr_source"); ok {
		t, err := expandObjectFirewallMmsProfileMm1AddrSource(d, v, "mm1_addr_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-addr-source"] = t
		}
	}

	if v, ok := d.GetOk("mm1_convert_hex"); ok {
		t, err := expandObjectFirewallMmsProfileMm1ConvertHex(d, v, "mm1_convert_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-convert-hex"] = t
		}
	}

	if v, ok := d.GetOk("mm1_outbreak_prevention"); ok {
		t, err := expandObjectFirewallMmsProfileMm1OutbreakPrevention(d, v, "mm1_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm1_retr_dupe"); ok {
		t, err := expandObjectFirewallMmsProfileMm1RetrDupe(d, v, "mm1_retr_dupe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-retr-dupe"] = t
		}
	}

	if v, ok := d.GetOk("mm1_retrieve_scan"); ok {
		t, err := expandObjectFirewallMmsProfileMm1RetrieveScan(d, v, "mm1_retrieve_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1-retrieve-scan"] = t
		}
	}

	if v, ok := d.GetOk("mm1comfortamount"); ok {
		t, err := expandObjectFirewallMmsProfileMm1Comfortamount(d, v, "mm1comfortamount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1comfortamount"] = t
		}
	}

	if v, ok := d.GetOk("mm1comfortinterval"); ok {
		t, err := expandObjectFirewallMmsProfileMm1Comfortinterval(d, v, "mm1comfortinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1comfortinterval"] = t
		}
	}

	if v, ok := d.GetOk("mm1oversizelimit"); ok {
		t, err := expandObjectFirewallMmsProfileMm1Oversizelimit(d, v, "mm1oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm1oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mm3"); ok {
		t, err := expandObjectFirewallMmsProfileMm3(d, v, "mm3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3"] = t
		}
	}

	if v, ok := d.GetOk("mm3_outbreak_prevention"); ok {
		t, err := expandObjectFirewallMmsProfileMm3OutbreakPrevention(d, v, "mm3_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm3oversizelimit"); ok {
		t, err := expandObjectFirewallMmsProfileMm3Oversizelimit(d, v, "mm3oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm3oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mm4"); ok {
		t, err := expandObjectFirewallMmsProfileMm4(d, v, "mm4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4"] = t
		}
	}

	if v, ok := d.GetOk("mm4_outbreak_prevention"); ok {
		t, err := expandObjectFirewallMmsProfileMm4OutbreakPrevention(d, v, "mm4_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm4oversizelimit"); ok {
		t, err := expandObjectFirewallMmsProfileMm4Oversizelimit(d, v, "mm4oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm4oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mm7"); ok {
		t, err := expandObjectFirewallMmsProfileMm7(d, v, "mm7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7"] = t
		}
	}

	if v, ok := d.GetOk("mm7_addr_hdr"); ok {
		t, err := expandObjectFirewallMmsProfileMm7AddrHdr(d, v, "mm7_addr_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-addr-hdr"] = t
		}
	}

	if v, ok := d.GetOk("mm7_addr_source"); ok {
		t, err := expandObjectFirewallMmsProfileMm7AddrSource(d, v, "mm7_addr_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-addr-source"] = t
		}
	}

	if v, ok := d.GetOk("mm7_convert_hex"); ok {
		t, err := expandObjectFirewallMmsProfileMm7ConvertHex(d, v, "mm7_convert_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-convert-hex"] = t
		}
	}

	if v, ok := d.GetOk("mm7_outbreak_prevention"); ok {
		t, err := expandObjectFirewallMmsProfileMm7OutbreakPrevention(d, v, "mm7_outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7-outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("mm7comfortamount"); ok {
		t, err := expandObjectFirewallMmsProfileMm7Comfortamount(d, v, "mm7comfortamount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7comfortamount"] = t
		}
	}

	if v, ok := d.GetOk("mm7comfortinterval"); ok {
		t, err := expandObjectFirewallMmsProfileMm7Comfortinterval(d, v, "mm7comfortinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7comfortinterval"] = t
		}
	}

	if v, ok := d.GetOk("mm7oversizelimit"); ok {
		t, err := expandObjectFirewallMmsProfileMm7Oversizelimit(d, v, "mm7oversizelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm7oversizelimit"] = t
		}
	}

	if v, ok := d.GetOk("mms_antispam_mass_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsAntispamMassLog(d, v, "mms_antispam_mass_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-antispam-mass-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_av_block_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsAvBlockLog(d, v, "mms_av_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-av-block-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_av_oversize_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsAvOversizeLog(d, v, "mms_av_oversize_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-av-oversize-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_av_virus_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsAvVirusLog(d, v, "mms_av_virus_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-av-virus-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_carrier_endpoint_filter_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsCarrierEndpointFilterLog(d, v, "mms_carrier_endpoint_filter_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-carrier-endpoint-filter-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_checksum_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsChecksumLog(d, v, "mms_checksum_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-checksum-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_checksum_table"); ok {
		t, err := expandObjectFirewallMmsProfileMmsChecksumTable(d, v, "mms_checksum_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-checksum-table"] = t
		}
	}

	if v, ok := d.GetOk("mms_notification_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsNotificationLog(d, v, "mms_notification_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-notification-log"] = t
		}
	}

	if v, ok := d.GetOk("mms_web_content_log"); ok {
		t, err := expandObjectFirewallMmsProfileMmsWebContentLog(d, v, "mms_web_content_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-web-content-log"] = t
		}
	}

	if v, ok := d.GetOk("mmsbwordthreshold"); ok {
		t, err := expandObjectFirewallMmsProfileMmsbwordthreshold(d, v, "mmsbwordthreshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mmsbwordthreshold"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallMmsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("notif_msisdn"); ok {
		t, err := expandObjectFirewallMmsProfileNotifMsisdn(d, v, "notif_msisdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notif-msisdn"] = t
		}
	}

	if v, ok := d.GetOk("remove_blocked_const_length"); ok {
		t, err := expandObjectFirewallMmsProfileRemoveBlockedConstLength(d, v, "remove_blocked_const_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-blocked-const-length"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandObjectFirewallMmsProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	return &obj, nil
}
