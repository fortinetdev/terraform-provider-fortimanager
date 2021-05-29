// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS domain filters.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectDnsfilterDomainFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDnsfilterDomainFilterCreate,
		Read:   resourceObjectDnsfilterDomainFilterRead,
		Update: resourceObjectDnsfilterDomainFilterUpdate,
		Delete: resourceObjectDnsfilterDomainFilterDelete,

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
				Computed: true,
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
						"domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
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

func resourceObjectDnsfilterDomainFilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDnsfilterDomainFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDnsfilterDomainFilter resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDnsfilterDomainFilter(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDnsfilterDomainFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDnsfilterDomainFilterRead(d, m)
}

func resourceObjectDnsfilterDomainFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDnsfilterDomainFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterDomainFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDnsfilterDomainFilter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDnsfilterDomainFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDnsfilterDomainFilterRead(d, m)
}

func resourceObjectDnsfilterDomainFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDnsfilterDomainFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDnsfilterDomainFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDnsfilterDomainFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDnsfilterDomainFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDnsfilterDomainFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDnsfilterDomainFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDnsfilterDomainFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectDnsfilterDomainFilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterDomainFilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDnsfilterDomainFilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectDnsfilterDomainFilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenObjectDnsfilterDomainFilterEntriesDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "ObjectDnsfilterDomainFilter-Entries-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectDnsfilterDomainFilterEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectDnsfilterDomainFilter-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectDnsfilterDomainFilterEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectDnsfilterDomainFilter-Entries-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectDnsfilterDomainFilterEntriesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectDnsfilterDomainFilter-Entries-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDnsfilterDomainFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "block",
			2: "allow",
			3: "monitor",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDnsfilterDomainFilterEntriesDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterDomainFilterEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterDomainFilterEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectDnsfilterDomainFilterEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "simple",
			1: "regex",
			2: "wildcard",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDnsfilterDomainFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDnsfilterDomainFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDnsfilterDomainFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenObjectDnsfilterDomainFilterComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDnsfilterDomainFilter-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectDnsfilterDomainFilterEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectDnsfilterDomainFilter-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectDnsfilterDomainFilterEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectDnsfilterDomainFilter-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectDnsfilterDomainFilterId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDnsfilterDomainFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDnsfilterDomainFilterName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDnsfilterDomainFilter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectDnsfilterDomainFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDnsfilterDomainFilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectDnsfilterDomainFilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandObjectDnsfilterDomainFilterEntriesDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectDnsfilterDomainFilterEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandObjectDnsfilterDomainFilterEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectDnsfilterDomainFilterEntriesType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDnsfilterDomainFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterEntriesDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDnsfilterDomainFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDnsfilterDomainFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectDnsfilterDomainFilterComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandObjectDnsfilterDomainFilterEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectDnsfilterDomainFilterId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectDnsfilterDomainFilterName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
