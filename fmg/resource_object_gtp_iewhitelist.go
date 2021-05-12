// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IE white list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectGtpIeWhiteList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGtpIeWhiteListCreate,
		Read:   resourceObjectGtpIeWhiteListRead,
		Update: resourceObjectGtpIeWhiteListUpdate,
		Delete: resourceObjectGtpIeWhiteListDelete,

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
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ie": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectGtpIeWhiteListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpIeWhiteList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpIeWhiteList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGtpIeWhiteList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpIeWhiteList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpIeWhiteListRead(d, m)
}

func resourceObjectGtpIeWhiteListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectGtpIeWhiteList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpIeWhiteList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGtpIeWhiteList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpIeWhiteList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpIeWhiteListRead(d, m)
}

func resourceObjectGtpIeWhiteListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectGtpIeWhiteList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGtpIeWhiteList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGtpIeWhiteListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectGtpIeWhiteList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpIeWhiteList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGtpIeWhiteList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpIeWhiteList resource from API: %v", err)
	}
	return nil
}

func flattenObjectGtpIeWhiteListEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectGtpIeWhiteListEntriesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectGtpIeWhiteList-Entries-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ie"
		if _, ok := i["ie"]; ok {
			v := flattenObjectGtpIeWhiteListEntriesIe(i["ie"], d, pre_append)
			tmp["ie"] = fortiAPISubPartPatch(v, "ObjectGtpIeWhiteList-Entries-Ie")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := i["message"]; ok {
			v := flattenObjectGtpIeWhiteListEntriesMessage(i["message"], d, pre_append)
			tmp["message"] = fortiAPISubPartPatch(v, "ObjectGtpIeWhiteList-Entries-Message")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectGtpIeWhiteListEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpIeWhiteListEntriesIe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpIeWhiteListEntriesMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpIeWhiteListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectGtpIeWhiteList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("entries", flattenObjectGtpIeWhiteListEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectGtpIeWhiteList-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectGtpIeWhiteListEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectGtpIeWhiteList-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectGtpIeWhiteListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGtpIeWhiteList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectGtpIeWhiteListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGtpIeWhiteListEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectGtpIeWhiteListEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ie"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ie"], _ = expandObjectGtpIeWhiteListEntriesIe(d, i["ie"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "message"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["message"], _ = expandObjectGtpIeWhiteListEntriesMessage(d, i["message"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectGtpIeWhiteListEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpIeWhiteListEntriesIe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpIeWhiteListEntriesMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpIeWhiteListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGtpIeWhiteList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandObjectGtpIeWhiteListEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectGtpIeWhiteListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
