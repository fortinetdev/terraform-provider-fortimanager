// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam IP trust.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSpamfilterIptrust() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSpamfilterIptrustCreate,
		Read:   resourceObjectSpamfilterIptrustRead,
		Update: resourceObjectSpamfilterIptrustUpdate,
		Delete: resourceObjectSpamfilterIptrustDelete,

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
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip4_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip6_subnet": &schema.Schema{
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectSpamfilterIptrustCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSpamfilterIptrust(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSpamfilterIptrust resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSpamfilterIptrust(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSpamfilterIptrust resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSpamfilterIptrustRead(d, m)
}

func resourceObjectSpamfilterIptrustUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSpamfilterIptrust(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSpamfilterIptrust resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSpamfilterIptrust(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSpamfilterIptrust resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSpamfilterIptrustRead(d, m)
}

func resourceObjectSpamfilterIptrustDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSpamfilterIptrust(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSpamfilterIptrust resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSpamfilterIptrustRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSpamfilterIptrust(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSpamfilterIptrust resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSpamfilterIptrust(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSpamfilterIptrust resource from API: %v", err)
	}
	return nil
}

func flattenObjectSpamfilterIptrustComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSpamfilterIptrustEntriesAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "ObjectSpamfilterIptrust-Entries-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSpamfilterIptrustEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSpamfilterIptrust-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := i["ip4-subnet"]; ok {
			v := flattenObjectSpamfilterIptrustEntriesIp4Subnet(i["ip4-subnet"], d, pre_append)
			tmp["ip4_subnet"] = fortiAPISubPartPatch(v, "ObjectSpamfilterIptrust-Entries-Ip4Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := i["ip6-subnet"]; ok {
			v := flattenObjectSpamfilterIptrustEntriesIp6Subnet(i["ip6-subnet"], d, pre_append)
			tmp["ip6_subnet"] = fortiAPISubPartPatch(v, "ObjectSpamfilterIptrust-Entries-Ip6Subnet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSpamfilterIptrustEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSpamfilterIptrust-Entries-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSpamfilterIptrustEntriesAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustEntriesIp4Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustEntriesIp6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterIptrustName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSpamfilterIptrust(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectSpamfilterIptrustComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectSpamfilterIptrust-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectSpamfilterIptrustEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectSpamfilterIptrust-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectSpamfilterIptrustEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectSpamfilterIptrust-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectSpamfilterIptrustId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSpamfilterIptrust-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSpamfilterIptrustName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSpamfilterIptrust-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSpamfilterIptrustFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSpamfilterIptrustComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandObjectSpamfilterIptrustEntriesAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSpamfilterIptrustEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip4-subnet"], _ = expandObjectSpamfilterIptrustEntriesIp4Subnet(d, i["ip4_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6-subnet"], _ = expandObjectSpamfilterIptrustEntriesIp6Subnet(d, i["ip6_subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSpamfilterIptrustEntriesStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSpamfilterIptrustEntriesAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustEntriesIp4Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustEntriesIp6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterIptrustName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSpamfilterIptrust(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectSpamfilterIptrustComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectSpamfilterIptrustEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectSpamfilterIptrustId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSpamfilterIptrustName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
