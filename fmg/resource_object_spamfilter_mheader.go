// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam MIME header.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSpamfilterMheader() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSpamfilterMheaderCreate,
		Read:   resourceObjectSpamfilterMheaderRead,
		Update: resourceObjectSpamfilterMheaderUpdate,
		Delete: resourceObjectSpamfilterMheaderDelete,

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
						},
						"fieldbody": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fieldname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pattern_type": &schema.Schema{
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

func resourceObjectSpamfilterMheaderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSpamfilterMheader(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSpamfilterMheader resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSpamfilterMheader(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSpamfilterMheader resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSpamfilterMheaderRead(d, m)
}

func resourceObjectSpamfilterMheaderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSpamfilterMheader(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSpamfilterMheader resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSpamfilterMheader(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSpamfilterMheader resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSpamfilterMheaderRead(d, m)
}

func resourceObjectSpamfilterMheaderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSpamfilterMheader(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSpamfilterMheader resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSpamfilterMheaderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSpamfilterMheader(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSpamfilterMheader resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSpamfilterMheader(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSpamfilterMheader resource from API: %v", err)
	}
	return nil
}

func flattenObjectSpamfilterMheaderComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSpamfilterMheaderEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectSpamfilterMheader-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := i["fieldbody"]; ok {
			v := flattenObjectSpamfilterMheaderEntriesFieldbody(i["fieldbody"], d, pre_append)
			tmp["fieldbody"] = fortiAPISubPartPatch(v, "ObjectSpamfilterMheader-Entries-Fieldbody")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := i["fieldname"]; ok {
			v := flattenObjectSpamfilterMheaderEntriesFieldname(i["fieldname"], d, pre_append)
			tmp["fieldname"] = fortiAPISubPartPatch(v, "ObjectSpamfilterMheader-Entries-Fieldname")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSpamfilterMheaderEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSpamfilterMheader-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := i["pattern-type"]; ok {
			v := flattenObjectSpamfilterMheaderEntriesPatternType(i["pattern-type"], d, pre_append)
			tmp["pattern_type"] = fortiAPISubPartPatch(v, "ObjectSpamfilterMheader-Entries-PatternType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSpamfilterMheaderEntriesStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSpamfilterMheader-Entries-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSpamfilterMheaderEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderEntriesFieldbody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderEntriesFieldname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderEntriesPatternType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSpamfilterMheaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSpamfilterMheader(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectSpamfilterMheaderComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectSpamfilterMheader-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectSpamfilterMheaderEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectSpamfilterMheader-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectSpamfilterMheaderEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectSpamfilterMheader-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectSpamfilterMheaderId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSpamfilterMheader-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSpamfilterMheaderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSpamfilterMheader-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSpamfilterMheaderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSpamfilterMheaderComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectSpamfilterMheaderEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldbody"], _ = expandObjectSpamfilterMheaderEntriesFieldbody(d, i["fieldbody"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldname"], _ = expandObjectSpamfilterMheaderEntriesFieldname(d, i["fieldname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSpamfilterMheaderEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern-type"], _ = expandObjectSpamfilterMheaderEntriesPatternType(d, i["pattern_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSpamfilterMheaderEntriesStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSpamfilterMheaderEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderEntriesFieldbody(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderEntriesFieldname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderEntriesPatternType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSpamfilterMheaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSpamfilterMheader(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectSpamfilterMheaderComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectSpamfilterMheaderEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectSpamfilterMheaderId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSpamfilterMheaderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
