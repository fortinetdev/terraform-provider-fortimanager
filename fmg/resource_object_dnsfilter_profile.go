// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS domain filter profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectDnsfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDnsfilterProfileCreate,
		Read:   resourceObjectDnsfilterProfileRead,
		Update: resourceObjectDnsfilterProfileUpdate,
		Delete: resourceObjectDnsfilterProfileDelete,

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
			"block_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_botnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_translation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"netmask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src6": &schema.Schema{
							Type:     schema.TypeString,
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
			"domain_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_filter_table": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"external_ip_blocklist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftgd_dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_all_domain": &schema.Schema{
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
			"redirect_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_portal6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"safe_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_domain_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_ftgd_err_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_restrict": &schema.Schema{
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

func resourceObjectDnsfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDnsfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDnsfilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDnsfilterProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDnsfilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDnsfilterProfileRead(d, m)
}

func resourceObjectDnsfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDnsfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDnsfilterProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDnsfilterProfileRead(d, m)
}

func resourceObjectDnsfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDnsfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDnsfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDnsfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDnsfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDnsfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDnsfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDnsfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectDnsfilterProfileBlockAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "block",
			1: "redirect",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDnsfilterProfileBlockBotnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationDst6(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "netmask"
		if _, ok := i["netmask"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationNetmask(i["netmask"], d, pre_append)
			tmp["netmask"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Netmask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationSrc(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationSrc6(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectDnsfilterProfileDnsTranslationStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfile-DnsTranslation-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDnsfilterProfileDnsTranslationAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			7: "ipv4",
			8: "ipv6",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationDst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationSrc6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileDnsTranslationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileDomainFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_filter_table"
	if _, ok := i["domain-filter-table"]; ok {
		result["domain_filter_table"] = flattenObjectDnsfilterProfileDomainFilterDomainFilterTable(i["domain-filter-table"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectDnsfilterProfileDomainFilterDomainFilterTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileExternalIpBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileFtgdDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenObjectDnsfilterProfileFtgdDnsFilters(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectDnsfilterProfileFtgdDnsOptions(i["options"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectDnsfilterProfileFtgdDnsFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDnsfilterProfileFtgdDnsFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfileFtgdDns-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectDnsfilterProfileFtgdDnsFiltersCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfileFtgdDns-Filters-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDnsfilterProfileFtgdDnsFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfileFtgdDns-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectDnsfilterProfileFtgdDnsFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectDnsfilterProfileFtgdDns-Filters-Log")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDnsfilterProfileFtgdDnsFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			3: "monitor",
			6: "block",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDnsfilterProfileFtgdDnsFiltersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileFtgdDnsFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileFtgdDnsFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileFtgdDnsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:    "error-allow",
			1024: "ftgd-disable",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectDnsfilterProfileLogAllDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileRedirectPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileRedirectPortal6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterProfileSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileSdnsDomainLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileSdnsFtgdErrLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterProfileYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			6:  "strict",
			26: "moderate",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectDnsfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("block_action", flattenObjectDnsfilterProfileBlockAction(o["block-action"], d, "block_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-action"], "ObjectDnsfilterProfile-BlockAction"); ok {
			if err = d.Set("block_action", vv); err != nil {
				return fmt.Errorf("Error reading block_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_action: %v", err)
		}
	}

	if err = d.Set("block_botnet", flattenObjectDnsfilterProfileBlockBotnet(o["block-botnet"], d, "block_botnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-botnet"], "ObjectDnsfilterProfile-BlockBotnet"); ok {
			if err = d.Set("block_botnet", vv); err != nil {
				return fmt.Errorf("Error reading block_botnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_botnet: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectDnsfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDnsfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dns_translation", flattenObjectDnsfilterProfileDnsTranslation(o["dns-translation"], d, "dns_translation")); err != nil {
			if vv, ok := fortiAPIPatch(o["dns-translation"], "ObjectDnsfilterProfile-DnsTranslation"); ok {
				if err = d.Set("dns_translation", vv); err != nil {
					return fmt.Errorf("Error reading dns_translation: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dns_translation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns_translation"); ok {
			if err = d.Set("dns_translation", flattenObjectDnsfilterProfileDnsTranslation(o["dns-translation"], d, "dns_translation")); err != nil {
				if vv, ok := fortiAPIPatch(o["dns-translation"], "ObjectDnsfilterProfile-DnsTranslation"); ok {
					if err = d.Set("dns_translation", vv); err != nil {
						return fmt.Errorf("Error reading dns_translation: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dns_translation: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("domain_filter", flattenObjectDnsfilterProfileDomainFilter(o["domain-filter"], d, "domain_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["domain-filter"], "ObjectDnsfilterProfile-DomainFilter"); ok {
				if err = d.Set("domain_filter", vv); err != nil {
					return fmt.Errorf("Error reading domain_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading domain_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("domain_filter"); ok {
			if err = d.Set("domain_filter", flattenObjectDnsfilterProfileDomainFilter(o["domain-filter"], d, "domain_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["domain-filter"], "ObjectDnsfilterProfile-DomainFilter"); ok {
					if err = d.Set("domain_filter", vv); err != nil {
						return fmt.Errorf("Error reading domain_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading domain_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("external_ip_blocklist", flattenObjectDnsfilterProfileExternalIpBlocklist(o["external-ip-blocklist"], d, "external_ip_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-ip-blocklist"], "ObjectDnsfilterProfile-ExternalIpBlocklist"); ok {
			if err = d.Set("external_ip_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_ip_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_ip_blocklist: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftgd_dns", flattenObjectDnsfilterProfileFtgdDns(o["ftgd-dns"], d, "ftgd_dns")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftgd-dns"], "ObjectDnsfilterProfile-FtgdDns"); ok {
				if err = d.Set("ftgd_dns", vv); err != nil {
					return fmt.Errorf("Error reading ftgd_dns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftgd_dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_dns"); ok {
			if err = d.Set("ftgd_dns", flattenObjectDnsfilterProfileFtgdDns(o["ftgd-dns"], d, "ftgd_dns")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftgd-dns"], "ObjectDnsfilterProfile-FtgdDns"); ok {
					if err = d.Set("ftgd_dns", vv); err != nil {
						return fmt.Errorf("Error reading ftgd_dns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftgd_dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_all_domain", flattenObjectDnsfilterProfileLogAllDomain(o["log-all-domain"], d, "log_all_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-all-domain"], "ObjectDnsfilterProfile-LogAllDomain"); ok {
			if err = d.Set("log_all_domain", vv); err != nil {
				return fmt.Errorf("Error reading log_all_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_all_domain: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDnsfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDnsfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("redirect_portal", flattenObjectDnsfilterProfileRedirectPortal(o["redirect-portal"], d, "redirect_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-portal"], "ObjectDnsfilterProfile-RedirectPortal"); ok {
			if err = d.Set("redirect_portal", vv); err != nil {
				return fmt.Errorf("Error reading redirect_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_portal: %v", err)
		}
	}

	if err = d.Set("redirect_portal6", flattenObjectDnsfilterProfileRedirectPortal6(o["redirect-portal6"], d, "redirect_portal6")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-portal6"], "ObjectDnsfilterProfile-RedirectPortal6"); ok {
			if err = d.Set("redirect_portal6", vv); err != nil {
				return fmt.Errorf("Error reading redirect_portal6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_portal6: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenObjectDnsfilterProfileSafeSearch(o["safe-search"], d, "safe_search")); err != nil {
		if vv, ok := fortiAPIPatch(o["safe-search"], "ObjectDnsfilterProfile-SafeSearch"); ok {
			if err = d.Set("safe_search", vv); err != nil {
				return fmt.Errorf("Error reading safe_search: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("sdns_domain_log", flattenObjectDnsfilterProfileSdnsDomainLog(o["sdns-domain-log"], d, "sdns_domain_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-domain-log"], "ObjectDnsfilterProfile-SdnsDomainLog"); ok {
			if err = d.Set("sdns_domain_log", vv); err != nil {
				return fmt.Errorf("Error reading sdns_domain_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_domain_log: %v", err)
		}
	}

	if err = d.Set("sdns_ftgd_err_log", flattenObjectDnsfilterProfileSdnsFtgdErrLog(o["sdns-ftgd-err-log"], d, "sdns_ftgd_err_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdns-ftgd-err-log"], "ObjectDnsfilterProfile-SdnsFtgdErrLog"); ok {
			if err = d.Set("sdns_ftgd_err_log", vv); err != nil {
				return fmt.Errorf("Error reading sdns_ftgd_err_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdns_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenObjectDnsfilterProfileYoutubeRestrict(o["youtube-restrict"], d, "youtube_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-restrict"], "ObjectDnsfilterProfile-YoutubeRestrict"); ok {
			if err = d.Set("youtube_restrict", vv); err != nil {
				return fmt.Errorf("Error reading youtube_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func flattenObjectDnsfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDnsfilterProfileBlockAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileBlockBotnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandObjectDnsfilterProfileDnsTranslationAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandObjectDnsfilterProfileDnsTranslationDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst6"], _ = expandObjectDnsfilterProfileDnsTranslationDst6(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectDnsfilterProfileDnsTranslationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "netmask"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["netmask"], _ = expandObjectDnsfilterProfileDnsTranslationNetmask(d, i["netmask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandObjectDnsfilterProfileDnsTranslationPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src"], _ = expandObjectDnsfilterProfileDnsTranslationSrc(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src6"], _ = expandObjectDnsfilterProfileDnsTranslationSrc6(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandObjectDnsfilterProfileDnsTranslationStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDnsfilterProfileDnsTranslationAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationDst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationSrc6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDnsTranslationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileDomainFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_filter_table"
	if _, ok := d.GetOk(pre_append); ok {
		result["domain-filter-table"], _ = expandObjectDnsfilterProfileDomainFilterDomainFilterTable(d, i["domain_filter_table"], pre_append)
	}

	return result, nil
}

func expandObjectDnsfilterProfileDomainFilterDomainFilterTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileExternalIpBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileFtgdDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandObjectDnsfilterProfileFtgdDnsFilters(d, i["filters"], pre_append)
	} else {
		result["filters"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectDnsfilterProfileFtgdDnsOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectDnsfilterProfileFtgdDnsFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectDnsfilterProfileFtgdDnsFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandObjectDnsfilterProfileFtgdDnsFiltersCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectDnsfilterProfileFtgdDnsFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandObjectDnsfilterProfileFtgdDnsFiltersLog(d, i["log"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDnsfilterProfileFtgdDnsFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileFtgdDnsFiltersCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileFtgdDnsFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileFtgdDnsFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileFtgdDnsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDnsfilterProfileLogAllDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileRedirectPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileRedirectPortal6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileSdnsDomainLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileSdnsFtgdErrLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterProfileYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDnsfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_action"); ok {
		t, err := expandObjectDnsfilterProfileBlockAction(d, v, "block_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-action"] = t
		}
	}

	if v, ok := d.GetOk("block_botnet"); ok {
		t, err := expandObjectDnsfilterProfileBlockBotnet(d, v, "block_botnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-botnet"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectDnsfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns_translation"); ok {
		t, err := expandObjectDnsfilterProfileDnsTranslation(d, v, "dns_translation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-translation"] = t
		}
	}

	if v, ok := d.GetOk("domain_filter"); ok {
		t, err := expandObjectDnsfilterProfileDomainFilter(d, v, "domain_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-filter"] = t
		}
	}

	if v, ok := d.GetOk("external_ip_blocklist"); ok {
		t, err := expandObjectDnsfilterProfileExternalIpBlocklist(d, v, "external_ip_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_dns"); ok {
		t, err := expandObjectDnsfilterProfileFtgdDns(d, v, "ftgd_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-dns"] = t
		}
	}

	if v, ok := d.GetOk("log_all_domain"); ok {
		t, err := expandObjectDnsfilterProfileLogAllDomain(d, v, "log_all_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-domain"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectDnsfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("redirect_portal"); ok {
		t, err := expandObjectDnsfilterProfileRedirectPortal(d, v, "redirect_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-portal"] = t
		}
	}

	if v, ok := d.GetOk("redirect_portal6"); ok {
		t, err := expandObjectDnsfilterProfileRedirectPortal6(d, v, "redirect_portal6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-portal6"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok {
		t, err := expandObjectDnsfilterProfileSafeSearch(d, v, "safe_search")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("sdns_domain_log"); ok {
		t, err := expandObjectDnsfilterProfileSdnsDomainLog(d, v, "sdns_domain_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("sdns_ftgd_err_log"); ok {
		t, err := expandObjectDnsfilterProfileSdnsFtgdErrLog(d, v, "sdns_ftgd_err_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok {
		t, err := expandObjectDnsfilterProfileYoutubeRestrict(d, v, "youtube_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	return &obj, nil
}
