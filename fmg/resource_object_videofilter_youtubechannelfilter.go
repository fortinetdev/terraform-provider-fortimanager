// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure YouTube channel filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVideofilterYoutubeChannelFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVideofilterYoutubeChannelFilterCreate,
		Read:   resourceObjectVideofilterYoutubeChannelFilterRead,
		Update: resourceObjectVideofilterYoutubeChannelFilterUpdate,
		Delete: resourceObjectVideofilterYoutubeChannelFilterDelete,

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
			"default_action": &schema.Schema{
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
						"channel_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_category": &schema.Schema{
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

func resourceObjectVideofilterYoutubeChannelFilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVideofilterYoutubeChannelFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVideofilterYoutubeChannelFilter resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVideofilterYoutubeChannelFilter(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVideofilterYoutubeChannelFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVideofilterYoutubeChannelFilterRead(d, m)
}

func resourceObjectVideofilterYoutubeChannelFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVideofilterYoutubeChannelFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterYoutubeChannelFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVideofilterYoutubeChannelFilter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterYoutubeChannelFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVideofilterYoutubeChannelFilterRead(d, m)
}

func resourceObjectVideofilterYoutubeChannelFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVideofilterYoutubeChannelFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVideofilterYoutubeChannelFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVideofilterYoutubeChannelFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVideofilterYoutubeChannelFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterYoutubeChannelFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVideofilterYoutubeChannelFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterYoutubeChannelFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectVideofilterYoutubeChannelFilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVideofilterYoutubeChannelFilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectVideofilterYoutubeChannelFilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := i["channel-id"]; ok {
			v := flattenObjectVideofilterYoutubeChannelFilterEntriesChannelId(i["channel-id"], d, pre_append)
			tmp["channel_id"] = fortiAPISubPartPatch(v, "ObjectVideofilterYoutubeChannelFilter-Entries-ChannelId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectVideofilterYoutubeChannelFilterEntriesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectVideofilterYoutubeChannelFilter-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVideofilterYoutubeChannelFilterEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVideofilterYoutubeChannelFilter-Entries-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVideofilterYoutubeChannelFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterEntriesChannelId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterYoutubeChannelFilterOverrideCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVideofilterYoutubeChannelFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comment", flattenObjectVideofilterYoutubeChannelFilterComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectVideofilterYoutubeChannelFilter-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("default_action", flattenObjectVideofilterYoutubeChannelFilterDefaultAction(o["default-action"], d, "default_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-action"], "ObjectVideofilterYoutubeChannelFilter-DefaultAction"); ok {
			if err = d.Set("default_action", vv); err != nil {
				return fmt.Errorf("Error reading default_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectVideofilterYoutubeChannelFilterEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectVideofilterYoutubeChannelFilter-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectVideofilterYoutubeChannelFilterEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectVideofilterYoutubeChannelFilter-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectVideofilterYoutubeChannelFilterId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectVideofilterYoutubeChannelFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectVideofilterYoutubeChannelFilterLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectVideofilterYoutubeChannelFilter-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVideofilterYoutubeChannelFilterName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVideofilterYoutubeChannelFilter-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_category", flattenObjectVideofilterYoutubeChannelFilterOverrideCategory(o["override-category"], d, "override_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-category"], "ObjectVideofilterYoutubeChannelFilter-OverrideCategory"); ok {
			if err = d.Set("override_category", vv); err != nil {
				return fmt.Errorf("Error reading override_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_category: %v", err)
		}
	}

	return nil
}

func flattenObjectVideofilterYoutubeChannelFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVideofilterYoutubeChannelFilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectVideofilterYoutubeChannelFilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["channel-id"], _ = expandObjectVideofilterYoutubeChannelFilterEntriesChannelId(d, i["channel_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandObjectVideofilterYoutubeChannelFilterEntriesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectVideofilterYoutubeChannelFilterEntriesId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVideofilterYoutubeChannelFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterEntriesChannelId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterYoutubeChannelFilterOverrideCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVideofilterYoutubeChannelFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectVideofilterYoutubeChannelFilterComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok || d.HasChange("default_action") {
		t, err := expandObjectVideofilterYoutubeChannelFilterDefaultAction(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandObjectVideofilterYoutubeChannelFilterEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectVideofilterYoutubeChannelFilterId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectVideofilterYoutubeChannelFilterLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVideofilterYoutubeChannelFilterName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_category"); ok || d.HasChange("override_category") {
		t, err := expandObjectVideofilterYoutubeChannelFilterOverrideCategory(d, v, "override_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-category"] = t
		}
	}

	return &obj, nil
}
