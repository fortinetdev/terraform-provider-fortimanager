package fortimanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDeviceDnsDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDeviceDnsDatabaseCreate,
		Read:   resourceObjectDeviceDnsDatabaseRead,
		Update: resourceObjectDeviceDnsDatabaseUpdate,
		Delete: resourceObjectDeviceDnsDatabaseDelete,

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
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"allow_transfer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"authoritative": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"preference": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
			},
			"ip_master": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forwarder6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"primary_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rr_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"view": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceObjectDeviceDnsDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	cfg := m.(*FortiClient).Cfg
	c.Retries = 1

	mkey, err := getObjectDeviceDnsDatabaseId(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDeviceDnsDatabase resource while getting ID: %v", err)
	}
	vdom := d.Get("vdom").(string)
	device := d.Get("device").(string)

	_, err = adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDeviceDnsDatabase(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDeviceDnsDatabase resource while getting object: %v", err)
	}
	objJSON, err := json.Marshal(obj)

	url := getObjectDeviceDnsDatabaseUrl(vdom, device, "")

	log.Printf("REQUEST BODY: %s", objJSON)

	str := fmt.Sprintf("{\"method\":\"add\",\"params\":[{\"url\":\"%s\", \"data\": %s}]}", url, string(objJSON))
	buf := bytes.NewBufferString(str)

	req := c.NewRequest("POST", "/jsonrpc", nil, buf)
	err = req.Send()

	if err != nil || req.HTTPResponse == nil {
		return fmt.Errorf("Error creating ObjectDeviceDnsDatabase resource: %v", err)
	}

	defer req.HTTPResponse.Body.Close()
	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		return fmt.Errorf("Error creating ObjectDeviceDnsDatabase resource: %v", err)
	}

	log.Printf("RESPONSE BODY: %s", body)

	d.SetId(mkey)

	return resourceObjectDeviceDnsDatabaseRead(d, m)
}

func resourceObjectDeviceDnsDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom, device, name, err := parseObjectDeviceDnsDatabaseId(d)

	if err != nil {
		return fmt.Errorf("Error updating ObjectDeviceDnsDatabase resource while getting ID: %v", err)
	}

	cfg := m.(*FortiClient).Cfg
	_, err = adomChecking(cfg, d)

	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDeviceDnsDatabase(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDeviceDnsDatabase resource while getting object: %v", err)
	}

	objJSON, err := json.Marshal(obj)

	if err != nil {
		return fmt.Errorf("Error updating ObjectDeviceDnsDatabase resource while parsing object: %v", err)
	}

	url := getObjectDeviceDnsDatabaseUrl(vdom, device, name)
	str := fmt.Sprintf("{\"method\":\"update\",\"params\":[{\"url\":\"%s\", \"data\": %s}]}", url, string(objJSON))
	buf := bytes.NewBufferString(str)

	req := c.NewRequest("POST", "/jsonrpc", nil, buf)
	err = req.Send()

	if err != nil || req.HTTPResponse == nil {
		return fmt.Errorf("Error updating ObjectDeviceDnsDatabase resource: %v", err)
	}

	defer req.HTTPResponse.Body.Close()
	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		return fmt.Errorf("Error updating ObjectDeviceDnsDatabase resource: %v", err)
	}

	d.SetId(mkey)

	return resourceObjectDeviceDnsDatabaseRead(d, m)
}

func resourceObjectDeviceDnsDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom, device, name, err := parseObjectDeviceDnsDatabaseId(d)

	if err != nil {
		return fmt.Errorf("Error deleting ObjectDeviceDnsDatabase resource while getting ID: %v", err)
	}

	cfg := m.(*FortiClient).Cfg
	_, err = adomChecking(cfg, d)

	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	url := getObjectDeviceDnsDatabaseUrl(vdom, device, "")
	str := fmt.Sprintf("{\"method\":\"delete\",\"params\":[{\"url\":\"%s\",\"confirm\":1,\"filter\":[\"name\",\"in\",%q]}]}", url, name)
	buf := bytes.NewBufferString(str)

	req := c.NewRequest("POST", "/jsonrpc", nil, buf)
	err = req.Send()

	if err != nil || req.HTTPResponse == nil {
		return fmt.Errorf("Error deleting ObjectDeviceDnsDatabase resource: %v", err)
	}

	defer req.HTTPResponse.Body.Close()
	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		return fmt.Errorf("Error deleting ObjectDeviceDnsDatabase resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDeviceDnsDatabaseRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom, device, name, err := parseObjectDeviceDnsDatabaseId(d)

	if err != nil {
		return fmt.Errorf("Error reading ObjectDeviceDnsDatabase resource while getting ID: %v", err)
	}

	cfg := m.(*FortiClient).Cfg
	_, err = adomChecking(cfg, d)

	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	url := getObjectDeviceDnsDatabaseUrl(vdom, device, "")
	str := fmt.Sprintf("{\"method\":\"get\",\"params\":[{\"url\":\"%s\"}]}", url)
	buf := bytes.NewBufferString(str)

	req := c.NewRequest("POST", "/jsonrpc", nil, buf)
	err = req.Send()

	if err != nil || req.HTTPResponse == nil {
		return fmt.Errorf("Error reading ObjectDeviceDnsDatabase resource: %v", err)
	}

	defer req.HTTPResponse.Body.Close()
	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		return fmt.Errorf("Error reading ObjectDeviceDnsDatabase resource: %v", err)
	}

	result := make(map[string]interface{})
	json.Unmarshal([]byte(string(body)), &result)

	o, err := getObjectDeviceDnsDatabaseResultData(result, name)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDeviceDnsDatabase resource:%v\n%s", err, string(body))
	}

	err = refreshObjectObjectDeviceDnsDatabase(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDeviceDnsDatabase resource: %v\n%s", err, string(body))
	}
	return nil
}

func getObjectDeviceDnsDatabaseUrl(vdom, device, name string) string {
	base := fmt.Sprintf("pm/config/device/%s/vdom/%s/system/dns-database", device, vdom)

	if name != "" {
		return base + "/" + name
	}

	return base
}

func parseObjectDeviceDnsDatabaseId(d *schema.ResourceData) (string, string, string, error) {
	parts := strings.Split(d.Id(), "/")

	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("Expected ID in format vdom/device/dns-database")
	}

	return parts[0], parts[1], parts[2], nil
}

func getObjectDeviceDnsDatabaseId(d *schema.ResourceData) (string, error) {
	vdom, ok := d.GetOk("vdom")
	if !ok {
		return "", fmt.Errorf("Expected vdom")
	}
	device, ok := d.GetOk("device")
	if !ok {
		return "", fmt.Errorf("Expected device")
	}
	name, ok := d.GetOk("name")
	if !ok {
		return "", fmt.Errorf("Expected name")
	}

	return fmt.Sprintf("%s/%s/%s", vdom, device, name), nil
}

func getObjectDeviceDnsDatabaseResultData(o map[string]interface{}, name string) (map[string]interface{}, error) {

	result, ok := o["result"].([]interface{})
	if !ok || len(result) == 0 {
		return nil, fmt.Errorf("missing result")
	}

	result0, ok := result[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid result[0]")
	}

	data, ok := result0["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("missing data")
	}

	for i, v := range data {
		obj, ok := v.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid data[%d]", i)
		}

		if objName, ok := obj["name"].(string); ok && objName == name {
			return obj, nil
		}
	}

	return nil, fmt.Errorf("dns database with name %q not found", name)
}

func flattenObjectDeviceDnsDatabaseAllowTransfer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDeviceDnsDatabaseAuthoritative(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseContact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntry(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDeviceDnsDatabaseDnsEntryCanonicalName(i["canonical-name"], d, pre_append)
			tmp["canonical_name"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-CanonicalName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryHostname(i["hostname"], d, pre_append)
			tmp["hostname"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Hostname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := i["ipv6"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryIpv6(i["ipv6"], d, pre_append)
			tmp["ipv6"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Ipv6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if _, ok := i["preference"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryPreference(i["preference"], d, pre_append)
			tmp["preference"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Preference")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := i["ttl"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryTtl(i["ttl"], d, pre_append)
			tmp["ttl"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Ttl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectDeviceDnsDatabaseDnsEntryType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectDeviceDnsDatabase-DnsEntry-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectDeviceDnsDatabaseDnsEntryCanonicalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDnsEntryType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseForwarder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDeviceDnsDatabaseIpMaster(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseForwarder6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDeviceDnsDatabaseInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseIpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabasePrimaryName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseRrMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectDeviceDnsDatabaseSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDeviceDnsDatabaseVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDeviceDnsDatabase(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("allow_transfer", flattenObjectDeviceDnsDatabaseAllowTransfer(o["allow-transfer"], d, "allow_transfer")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-transfer"], "ObjectDeviceDnsDatabase-AllowTransfer"); ok {
			if err = d.Set("allow_transfer", vv); err != nil {
				return fmt.Errorf("Error reading allow_transfer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_transfer: %v", err)
		}
	}

	if err = d.Set("authoritative", flattenObjectDeviceDnsDatabaseAuthoritative(o["authoritative"], d, "authoritative")); err != nil {
		if vv, ok := fortiAPIPatch(o["authoritative"], "ObjectDeviceDnsDatabase-Authoritative"); ok {
			if err = d.Set("authoritative", vv); err != nil {
				return fmt.Errorf("Error reading authoritative: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authoritative: %v", err)
		}
	}

	if err = d.Set("contact", flattenObjectDeviceDnsDatabaseContact(o["contact"], d, "contact")); err != nil {
		if vv, ok := fortiAPIPatch(o["contact"], "ObjectDeviceDnsDatabase-Contact"); ok {
			if err = d.Set("contact", vv); err != nil {
				return fmt.Errorf("Error reading contact: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading contact: %v", err)
		}
	}

	if isImportTable() {
		log.Printf("Tu OK")

		if err = d.Set("dns_entry", flattenObjectDeviceDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry")); err != nil {
			if vv, ok := fortiAPIPatch(o["dns-entry"], "ObjectDeviceDnsDatabase-DnsEntry"); ok {
				if err = d.Set("dns_entry", vv); err != nil {
					return fmt.Errorf("Error reading dns_entry: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dns_entry: %v", err)
			}
		}
	} else {
		log.Printf("Tu NOK")

		if _, ok := d.GetOk("dns_entry"); ok {
			if err = d.Set("dns_entry", flattenObjectDeviceDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry")); err != nil {
				if vv, ok := fortiAPIPatch(o["dns-entry"], "ObjectDeviceDnsDatabase-DnsEntry"); ok {
					if err = d.Set("dns_entry", vv); err != nil {
						return fmt.Errorf("Error reading dns_entry: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dns_entry: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain", flattenObjectDeviceDnsDatabaseDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectDeviceDnsDatabase-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("forwarder", flattenObjectDeviceDnsDatabaseForwarder(o["forwarder"], d, "forwarder")); err != nil {
		if vv, ok := fortiAPIPatch(o["forwarder"], "ObjectDeviceDnsDatabase-Forwarder"); ok {
			if err = d.Set("forwarder", vv); err != nil {
				return fmt.Errorf("Error reading forwarder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forwarder: %v", err)
		}
	}

	if err = d.Set("ip_master", flattenObjectDeviceDnsDatabaseIpMaster(o["ip-master"], d, "ip_master")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-master"], "ObjectDeviceDnsDatabase-IpMaster"); ok {
			if err = d.Set("ip_master", vv); err != nil {
				return fmt.Errorf("Error reading ip_master: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_master: %v", err)
		}
	}

	if err = d.Set("forwarder6", flattenObjectDeviceDnsDatabaseForwarder6(o["forwarder6"], d, "forwarder6")); err != nil {
		if vv, ok := fortiAPIPatch(o["forwarder6"], "ObjectDeviceDnsDatabase-Forwarder6"); ok {
			if err = d.Set("forwarder6", vv); err != nil {
				return fmt.Errorf("Error reading forwarder6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forwarder6: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectDeviceDnsDatabaseInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectDeviceDnsDatabase-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectDeviceDnsDatabaseInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectDeviceDnsDatabase-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("ip_primary", flattenObjectDeviceDnsDatabaseIpPrimary(o["ip-primary"], d, "ip_primary")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-primary"], "ObjectDeviceDnsDatabase-IpPrimary"); ok {
			if err = d.Set("ip_primary", vv); err != nil {
				return fmt.Errorf("Error reading ip_primary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDeviceDnsDatabaseName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDeviceDnsDatabase-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("primary_name", flattenObjectDeviceDnsDatabasePrimaryName(o["primary-name"], d, "primary_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-name"], "ObjectDeviceDnsDatabase-PrimaryName"); ok {
			if err = d.Set("primary_name", vv); err != nil {
				return fmt.Errorf("Error reading primary_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_name: %v", err)
		}
	}

	if err = d.Set("rr_max", flattenObjectDeviceDnsDatabaseRrMax(o["rr-max"], d, "rr_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["rr-max"], "ObjectDeviceDnsDatabase-RrMax"); ok {
			if err = d.Set("rr_max", vv); err != nil {
				return fmt.Errorf("Error reading rr_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rr_max: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectDeviceDnsDatabaseSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectDeviceDnsDatabase-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenObjectDeviceDnsDatabaseSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "ObjectDeviceDnsDatabase-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenObjectDeviceDnsDatabaseSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "ObjectDeviceDnsDatabase-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectDeviceDnsDatabaseStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectDeviceDnsDatabase-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ttl", flattenObjectDeviceDnsDatabaseTtl(o["ttl"], d, "ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl"], "ObjectDeviceDnsDatabase-Ttl"); ok {
			if err = d.Set("ttl", vv); err != nil {
				return fmt.Errorf("Error reading ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectDeviceDnsDatabaseType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectDeviceDnsDatabase-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("view", flattenObjectDeviceDnsDatabaseView(o["view"], d, "view")); err != nil {
		if vv, ok := fortiAPIPatch(o["view"], "ObjectDeviceDnsDatabase-View"); ok {
			if err = d.Set("view", vv); err != nil {
				return fmt.Errorf("Error reading view: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading view: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenObjectDeviceDnsDatabaseVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "ObjectDeviceDnsDatabase-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func expandObjectDeviceDnsDatabaseAllowTransfer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDeviceDnsDatabaseAuthoritative(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseContact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["canonical-name"], _ = expandObjectDeviceDnsDatabaseDnsEntryCanonicalName(d, i["canonical_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hostname"], _ = expandObjectDeviceDnsDatabaseDnsEntryHostname(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectDeviceDnsDatabaseDnsEntryId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandObjectDeviceDnsDatabaseDnsEntryIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6"], _ = expandObjectDeviceDnsDatabaseDnsEntryIpv6(d, i["ipv6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preference"], _ = expandObjectDeviceDnsDatabaseDnsEntryPreference(d, i["preference"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectDeviceDnsDatabaseDnsEntryStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ttl"], _ = expandObjectDeviceDnsDatabaseDnsEntryTtl(d, i["ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectDeviceDnsDatabaseDnsEntryType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryCanonicalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryPreference(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDnsEntryType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseForwarder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDeviceDnsDatabaseIpMaster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseForwarder6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDeviceDnsDatabaseInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseIpPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabasePrimaryName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseRrMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectDeviceDnsDatabaseSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseView(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDeviceDnsDatabaseVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDeviceDnsDatabase(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_transfer"); ok {
		t, err := expandObjectDeviceDnsDatabaseAllowTransfer(d, v, "allow_transfer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-transfer"] = t
		}
	}

	if v, ok := d.GetOk("authoritative"); ok {
		t, err := expandObjectDeviceDnsDatabaseAuthoritative(d, v, "authoritative")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authoritative"] = t
		}
	}

	if v, ok := d.GetOk("contact"); ok {
		t, err := expandObjectDeviceDnsDatabaseContact(d, v, "contact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact"] = t
		}
	}

	if v, ok := d.GetOk("dns_entry"); ok {
		t, err := expandObjectDeviceDnsDatabaseDnsEntry(d, v, "dns_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-entry"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandObjectDeviceDnsDatabaseDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("forwarder"); ok {
		t, err := expandObjectDeviceDnsDatabaseForwarder(d, v, "forwarder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarder"] = t
		}
	}

	if v, ok := d.GetOk("ip_master"); ok {
		t, err := expandObjectDeviceDnsDatabaseIpMaster(d, v, "ip_master")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-master"] = t
		}
	}

	if v, ok := d.GetOk("forwarder6"); ok {
		t, err := expandObjectDeviceDnsDatabaseForwarder6(d, v, "forwarder6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarder6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectDeviceDnsDatabaseInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandObjectDeviceDnsDatabaseInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("ip_primary"); ok {
		t, err := expandObjectDeviceDnsDatabaseIpPrimary(d, v, "ip_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-primary"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectDeviceDnsDatabaseName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("primary_name"); ok {
		t, err := expandObjectDeviceDnsDatabasePrimaryName(d, v, "primary_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-name"] = t
		}
	}

	if v, ok := d.GetOk("rr_max"); ok {
		t, err := expandObjectDeviceDnsDatabaseRrMax(d, v, "rr_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-max"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandObjectDeviceDnsDatabaseSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok {
		t, err := expandObjectDeviceDnsDatabaseSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandObjectDeviceDnsDatabaseSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandObjectDeviceDnsDatabaseStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok {
		t, err := expandObjectDeviceDnsDatabaseTtl(d, v, "ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectDeviceDnsDatabaseType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("view"); ok {
		t, err := expandObjectDeviceDnsDatabaseView(d, v, "view")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["view"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok {
		t, err := expandObjectDeviceDnsDatabaseVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
