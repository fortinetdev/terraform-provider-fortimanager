// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS databases.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemDnsDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemDnsDatabaseCreate,
		Read:   resourceObjectSystemDnsDatabaseRead,
		Update: resourceObjectSystemDnsDatabaseUpdate,
		Delete: resourceObjectSystemDnsDatabaseDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"allow_transfer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"authoritative": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"canonical_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preference": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forwarder": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_master": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forwarder6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"primary_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
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

func resourceObjectSystemDnsDatabaseCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemDnsDatabase(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemDnsDatabase resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectSystemDnsDatabase(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectSystemDnsDatabase(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectSystemDnsDatabase resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectSystemDnsDatabase(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectSystemDnsDatabase resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemDnsDatabaseRead(d, m)
}

func resourceObjectSystemDnsDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemDnsDatabase(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDnsDatabase resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemDnsDatabase(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDnsDatabase resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemDnsDatabaseRead(d, m)
}

func resourceObjectSystemDnsDatabaseDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemDnsDatabase(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemDnsDatabase resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemDnsDatabaseRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemDnsDatabase(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectSystemDnsDatabase resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemDnsDatabase(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemDnsDatabase resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemDnsDatabaseAllowTransfer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDnsDatabaseAuthoritative(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseContact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntry(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "canonical_name"
		if _, ok := i["canonical-name"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryCanonicalName(i["canonical-name"], d, pre_append)
			tmp["canonical_name"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-CanonicalName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryHostname(i["hostname"], d, pre_append)
			tmp["hostname"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Hostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := i["ipv6"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryIpv6(i["ipv6"], d, pre_append)
			tmp["ipv6"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Ipv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if _, ok := i["preference"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryPreference(i["preference"], d, pre_append)
			tmp["preference"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Preference")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := i["ttl"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryTtl(i["ttl"], d, pre_append)
			tmp["ttl"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Ttl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectSystemDnsDatabaseDnsEntryType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectSystemDnsDatabase-DnsEntry-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemDnsDatabaseDnsEntryCanonicalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseForwarder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDnsDatabaseIpMaster(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseForwarder6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDnsDatabaseInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseIpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabasePrimaryName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseRrMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemDnsDatabaseSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemDnsDatabase(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("allow_transfer", flattenObjectSystemDnsDatabaseAllowTransfer(o["allow-transfer"], d, "allow_transfer")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-transfer"], "ObjectSystemDnsDatabase-AllowTransfer"); ok {
			if err = d.Set("allow_transfer", vv); err != nil {
				return fmt.Errorf("Error reading allow_transfer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_transfer: %v", err)
		}
	}

	if err = d.Set("authoritative", flattenObjectSystemDnsDatabaseAuthoritative(o["authoritative"], d, "authoritative")); err != nil {
		if vv, ok := fortiAPIPatch(o["authoritative"], "ObjectSystemDnsDatabase-Authoritative"); ok {
			if err = d.Set("authoritative", vv); err != nil {
				return fmt.Errorf("Error reading authoritative: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authoritative: %v", err)
		}
	}

	if err = d.Set("contact", flattenObjectSystemDnsDatabaseContact(o["contact"], d, "contact")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact"], "ObjectSystemDnsDatabase-Contact"); ok {
			if err = d.Set("contact", vv); err != nil {
				return fmt.Errorf("Error reading contact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dns_entry", flattenObjectSystemDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry")); err != nil {
			if vv, ok := fortiAPIPatch(o["dns-entry"], "ObjectSystemDnsDatabase-DnsEntry"); ok {
				if err = d.Set("dns_entry", vv); err != nil {
					return fmt.Errorf("Error reading dns_entry: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dns_entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns_entry"); ok {
			if err = d.Set("dns_entry", flattenObjectSystemDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry")); err != nil {
				if vv, ok := fortiAPIPatch(o["dns-entry"], "ObjectSystemDnsDatabase-DnsEntry"); ok {
					if err = d.Set("dns_entry", vv); err != nil {
						return fmt.Errorf("Error reading dns_entry: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dns_entry: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain", flattenObjectSystemDnsDatabaseDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectSystemDnsDatabase-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("forwarder", flattenObjectSystemDnsDatabaseForwarder(o["forwarder"], d, "forwarder")); err != nil {
		if vv, ok := fortiAPIPatch(o["forwarder"], "ObjectSystemDnsDatabase-Forwarder"); ok {
			if err = d.Set("forwarder", vv); err != nil {
				return fmt.Errorf("Error reading forwarder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forwarder: %v", err)
		}
	}

	if err = d.Set("ip_master", flattenObjectSystemDnsDatabaseIpMaster(o["ip-master"], d, "ip_master")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-master"], "ObjectSystemDnsDatabase-IpMaster"); ok {
			if err = d.Set("ip_master", vv); err != nil {
				return fmt.Errorf("Error reading ip_master: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_master: %v", err)
		}
	}

	if err = d.Set("forwarder6", flattenObjectSystemDnsDatabaseForwarder6(o["forwarder6"], d, "forwarder6")); err != nil {
		if vv, ok := fortiAPIPatch(o["forwarder6"], "ObjectSystemDnsDatabase-Forwarder6"); ok {
			if err = d.Set("forwarder6", vv); err != nil {
				return fmt.Errorf("Error reading forwarder6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forwarder6: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectSystemDnsDatabaseInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectSystemDnsDatabase-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectSystemDnsDatabaseInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectSystemDnsDatabase-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip_primary", flattenObjectSystemDnsDatabaseIpPrimary(o["ip-primary"], d, "ip_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-primary"], "ObjectSystemDnsDatabase-IpPrimary"); ok {
			if err = d.Set("ip_primary", vv); err != nil {
				return fmt.Errorf("Error reading ip_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemDnsDatabaseName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemDnsDatabase-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("primary_name", flattenObjectSystemDnsDatabasePrimaryName(o["primary-name"], d, "primary_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-name"], "ObjectSystemDnsDatabase-PrimaryName"); ok {
			if err = d.Set("primary_name", vv); err != nil {
				return fmt.Errorf("Error reading primary_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_name: %v", err)
		}
	}

	if err = d.Set("rr_max", flattenObjectSystemDnsDatabaseRrMax(o["rr-max"], d, "rr_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-max"], "ObjectSystemDnsDatabase-RrMax"); ok {
			if err = d.Set("rr_max", vv); err != nil {
				return fmt.Errorf("Error reading rr_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_max: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectSystemDnsDatabaseSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectSystemDnsDatabase-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenObjectSystemDnsDatabaseSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "ObjectSystemDnsDatabase-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenObjectSystemDnsDatabaseSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "ObjectSystemDnsDatabase-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSystemDnsDatabaseStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSystemDnsDatabase-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ttl", flattenObjectSystemDnsDatabaseTtl(o["ttl"], d, "ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl"], "ObjectSystemDnsDatabase-Ttl"); ok {
			if err = d.Set("ttl", vv); err != nil {
				return fmt.Errorf("Error reading ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemDnsDatabaseType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemDnsDatabase-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("view", flattenObjectSystemDnsDatabaseView(o["view"], d, "view")); err != nil {
		if vv, ok := fortiAPIPatch(o["view"], "ObjectSystemDnsDatabase-View"); ok {
			if err = d.Set("view", vv); err != nil {
				return fmt.Errorf("Error reading view: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading view: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenObjectSystemDnsDatabaseVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "ObjectSystemDnsDatabase-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemDnsDatabaseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemDnsDatabaseAllowTransfer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDnsDatabaseAuthoritative(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseContact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "canonical_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["canonical-name"], _ = expandObjectSystemDnsDatabaseDnsEntryCanonicalName(d, i["canonical_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hostname"], _ = expandObjectSystemDnsDatabaseDnsEntryHostname(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemDnsDatabaseDnsEntryId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectSystemDnsDatabaseDnsEntryIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6"], _ = expandObjectSystemDnsDatabaseDnsEntryIpv6(d, i["ipv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preference"], _ = expandObjectSystemDnsDatabaseDnsEntryPreference(d, i["preference"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSystemDnsDatabaseDnsEntryStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ttl"], _ = expandObjectSystemDnsDatabaseDnsEntryTtl(d, i["ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectSystemDnsDatabaseDnsEntryType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemDnsDatabaseDnsEntryCanonicalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryPreference(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseForwarder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDnsDatabaseIpMaster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseForwarder6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDnsDatabaseInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseIpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabasePrimaryName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseRrMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemDnsDatabaseSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseView(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemDnsDatabase(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_transfer"); ok || d.HasChange("allow_transfer") {
		t, err := expandObjectSystemDnsDatabaseAllowTransfer(d, v, "allow_transfer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-transfer"] = t
		}
	}

	if v, ok := d.GetOk("authoritative"); ok || d.HasChange("authoritative") {
		t, err := expandObjectSystemDnsDatabaseAuthoritative(d, v, "authoritative")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authoritative"] = t
		}
	}

	if v, ok := d.GetOk("contact"); ok || d.HasChange("contact") {
		t, err := expandObjectSystemDnsDatabaseContact(d, v, "contact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact"] = t
		}
	}

	if v, ok := d.GetOk("dns_entry"); ok || d.HasChange("dns_entry") {
		t, err := expandObjectSystemDnsDatabaseDnsEntry(d, v, "dns_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-entry"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectSystemDnsDatabaseDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("forwarder"); ok || d.HasChange("forwarder") {
		t, err := expandObjectSystemDnsDatabaseForwarder(d, v, "forwarder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarder"] = t
		}
	}

	if v, ok := d.GetOk("ip_master"); ok || d.HasChange("ip_master") {
		t, err := expandObjectSystemDnsDatabaseIpMaster(d, v, "ip_master")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-master"] = t
		}
	}

	if v, ok := d.GetOk("forwarder6"); ok || d.HasChange("forwarder6") {
		t, err := expandObjectSystemDnsDatabaseForwarder6(d, v, "forwarder6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarder6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectSystemDnsDatabaseInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectSystemDnsDatabaseInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip_primary"); ok || d.HasChange("ip_primary") {
		t, err := expandObjectSystemDnsDatabaseIpPrimary(d, v, "ip_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-primary"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemDnsDatabaseName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("primary_name"); ok || d.HasChange("primary_name") {
		t, err := expandObjectSystemDnsDatabasePrimaryName(d, v, "primary_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-name"] = t
		}
	}

	if v, ok := d.GetOk("rr_max"); ok || d.HasChange("rr_max") {
		t, err := expandObjectSystemDnsDatabaseRrMax(d, v, "rr_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-max"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectSystemDnsDatabaseSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandObjectSystemDnsDatabaseSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandObjectSystemDnsDatabaseSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSystemDnsDatabaseStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok || d.HasChange("ttl") {
		t, err := expandObjectSystemDnsDatabaseTtl(d, v, "ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemDnsDatabaseType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("view"); ok || d.HasChange("view") {
		t, err := expandObjectSystemDnsDatabaseView(d, v, "view")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["view"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandObjectSystemDnsDatabaseVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
