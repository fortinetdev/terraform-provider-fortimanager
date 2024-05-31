// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure URL filter lists.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterUrlfilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterUrlfilterCreate,
		Read:   resourceObjectWebfilterUrlfilterRead,
		Update: resourceObjectWebfilterUrlfilterUpdate,
		Delete: resourceObjectWebfilterUrlfilterDelete,

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
							Computed: true,
						},
						"antiphish_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_address_family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"referrer_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"web_proxy_profile": &schema.Schema{
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
			"ip_addr_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip4_mapped_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"one_arm_ips_urlfilter": &schema.Schema{
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

func resourceObjectWebfilterUrlfilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWebfilterUrlfilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterUrlfilter resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebfilterUrlfilter(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterUrlfilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterUrlfilterRead(d, m)
}

func resourceObjectWebfilterUrlfilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterUrlfilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterUrlfilter resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterUrlfilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterUrlfilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebfilterUrlfilterRead(d, m)
}

func resourceObjectWebfilterUrlfilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterUrlfilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterUrlfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterUrlfilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterUrlfilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterUrlfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterUrlfilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterUrlfilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterUrlfilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWebfilterUrlfilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish_action"
		if _, ok := i["antiphish-action"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesAntiphishAction(i["antiphish-action"], d, pre_append)
			tmp["antiphish_action"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-AntiphishAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := i["dns-address-family"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesDnsAddressFamily(i["dns-address-family"], d, pre_append)
			tmp["dns_address_family"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-DnsAddressFamily")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := i["exempt"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesExempt(i["exempt"], d, pre_append)
			tmp["exempt"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-Exempt")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := i["referrer-host"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesReferrerHost(i["referrer-host"], d, pre_append)
			tmp["referrer_host"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-ReferrerHost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := i["web-proxy-profile"]; ok {
			v := flattenObjectWebfilterUrlfilterEntriesWebProxyProfile(i["web-proxy-profile"], d, pre_append)
			tmp["web_proxy_profile"] = fortiAPISubPartPatch(v, "ObjectWebfilterUrlfilter-Entries-WebProxyProfile")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWebfilterUrlfilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesAntiphishAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesDnsAddressFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterUrlfilterEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesReferrerHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterEntriesWebProxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWebfilterUrlfilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterIpAddrBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterIp4MappedIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterUrlfilterOneArmIpsUrlfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterUrlfilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectWebfilterUrlfilterComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWebfilterUrlfilter-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectWebfilterUrlfilterEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectWebfilterUrlfilter-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectWebfilterUrlfilterEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectWebfilterUrlfilter-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectWebfilterUrlfilterId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWebfilterUrlfilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_addr_block", flattenObjectWebfilterUrlfilterIpAddrBlock(o["ip-addr-block"], d, "ip_addr_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-addr-block"], "ObjectWebfilterUrlfilter-IpAddrBlock"); ok {
			if err = d.Set("ip_addr_block", vv); err != nil {
				return fmt.Errorf("Error reading ip_addr_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_addr_block: %v", err)
		}
	}

	if err = d.Set("ip4_mapped_ip6", flattenObjectWebfilterUrlfilterIp4MappedIp6(o["ip4-mapped-ip6"], d, "ip4_mapped_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-mapped-ip6"], "ObjectWebfilterUrlfilter-Ip4MappedIp6"); ok {
			if err = d.Set("ip4_mapped_ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip4_mapped_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_mapped_ip6: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebfilterUrlfilterName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebfilterUrlfilter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("one_arm_ips_urlfilter", flattenObjectWebfilterUrlfilterOneArmIpsUrlfilter(o["one-arm-ips-urlfilter"], d, "one_arm_ips_urlfilter")); err != nil {
		if vv, ok := fortiAPIPatch(o["one-arm-ips-urlfilter"], "ObjectWebfilterUrlfilter-OneArmIpsUrlfilter"); ok {
			if err = d.Set("one_arm_ips_urlfilter", vv); err != nil {
				return fmt.Errorf("Error reading one_arm_ips_urlfilter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading one_arm_ips_urlfilter: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterUrlfilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterUrlfilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectWebfilterUrlfilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["antiphish-action"], _ = expandObjectWebfilterUrlfilterEntriesAntiphishAction(d, i["antiphish_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-address-family"], _ = expandObjectWebfilterUrlfilterEntriesDnsAddressFamily(d, i["dns_address_family"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exempt"], _ = expandObjectWebfilterUrlfilterEntriesExempt(d, i["exempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectWebfilterUrlfilterEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["referrer-host"], _ = expandObjectWebfilterUrlfilterEntriesReferrerHost(d, i["referrer_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectWebfilterUrlfilterEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectWebfilterUrlfilterEntriesType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandObjectWebfilterUrlfilterEntriesUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["web-proxy-profile"], _ = expandObjectWebfilterUrlfilterEntriesWebProxyProfile(d, i["web_proxy_profile"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWebfilterUrlfilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesAntiphishAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesDnsAddressFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterUrlfilterEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesReferrerHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterEntriesWebProxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebfilterUrlfilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterIpAddrBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterIp4MappedIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterUrlfilterOneArmIpsUrlfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterUrlfilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWebfilterUrlfilterComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectWebfilterUrlfilterEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWebfilterUrlfilterId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_block"); ok || d.HasChange("ip_addr_block") {
		t, err := expandObjectWebfilterUrlfilterIpAddrBlock(d, v, "ip_addr_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-block"] = t
		}
	}

	if v, ok := d.GetOk("ip4_mapped_ip6"); ok || d.HasChange("ip4_mapped_ip6") {
		t, err := expandObjectWebfilterUrlfilterIp4MappedIp6(d, v, "ip4_mapped_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-mapped-ip6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebfilterUrlfilterName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("one_arm_ips_urlfilter"); ok || d.HasChange("one_arm_ips_urlfilter") {
		t, err := expandObjectWebfilterUrlfilterOneArmIpsUrlfilter(d, v, "one_arm_ips_urlfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["one-arm-ips-urlfilter"] = t
		}
	}

	return &obj, nil
}
