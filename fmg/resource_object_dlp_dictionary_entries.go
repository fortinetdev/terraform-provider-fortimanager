// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DLP dictionary entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpDictionaryEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpDictionaryEntriesCreate,
		Read:   resourceObjectDlpDictionaryEntriesRead,
		Update: resourceObjectDlpDictionaryEntriesUpdate,
		Delete: resourceObjectDlpDictionaryEntriesDelete,

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
			"dictionary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ignore_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repeat": &schema.Schema{
				Type:     schema.TypeString,
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
			},
		},
	}
}

func resourceObjectDlpDictionaryEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	dictionary := d.Get("dictionary").(string)
	paradict["dictionary"] = dictionary

	obj, err := getObjectObjectDlpDictionaryEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpDictionaryEntries resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpDictionaryEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpDictionaryEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpDictionaryEntriesRead(d, m)
}

func resourceObjectDlpDictionaryEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	dictionary := d.Get("dictionary").(string)
	paradict["dictionary"] = dictionary

	obj, err := getObjectObjectDlpDictionaryEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpDictionaryEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpDictionaryEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpDictionaryEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpDictionaryEntriesRead(d, m)
}

func resourceObjectDlpDictionaryEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	dictionary := d.Get("dictionary").(string)
	paradict["dictionary"] = dictionary

	err = c.DeleteObjectDlpDictionaryEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpDictionaryEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpDictionaryEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	dictionary := d.Get("dictionary").(string)
	if dictionary == "" {
		dictionary = importOptionChecking(m.(*FortiClient).Cfg, "dictionary")
		if dictionary == "" {
			return fmt.Errorf("Parameter dictionary is missing")
		}
		if err = d.Set("dictionary", dictionary); err != nil {
			return fmt.Errorf("Error set params dictionary: %v", err)
		}
	}
	paradict["dictionary"] = dictionary

	o, err := c.ReadObjectDlpDictionaryEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpDictionaryEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpDictionaryEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpDictionaryEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpDictionaryEntriesComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDictionaryEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDictionaryEntriesIgnoreCase2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDictionaryEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDictionaryEntriesRepeat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDictionaryEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDictionaryEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectDlpDictionaryEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectDlpDictionaryEntriesComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDlpDictionaryEntries-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectDlpDictionaryEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDlpDictionaryEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ignore_case", flattenObjectDlpDictionaryEntriesIgnoreCase2edl(o["ignore-case"], d, "ignore_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-case"], "ObjectDlpDictionaryEntries-IgnoreCase"); ok {
			if err = d.Set("ignore_case", vv); err != nil {
				return fmt.Errorf("Error reading ignore_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_case: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectDlpDictionaryEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectDlpDictionaryEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("repeat", flattenObjectDlpDictionaryEntriesRepeat2edl(o["repeat"], d, "repeat")); err != nil {
		if vv, ok := fortiAPIPatch(o["repeat"], "ObjectDlpDictionaryEntries-Repeat"); ok {
			if err = d.Set("repeat", vv); err != nil {
				return fmt.Errorf("Error reading repeat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectDlpDictionaryEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectDlpDictionaryEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectDlpDictionaryEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectDlpDictionaryEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpDictionaryEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpDictionaryEntriesComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDictionaryEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDictionaryEntriesIgnoreCase2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDictionaryEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDictionaryEntriesRepeat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDictionaryEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDictionaryEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectDlpDictionaryEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectDlpDictionaryEntriesComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectDlpDictionaryEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ignore_case"); ok || d.HasChange("ignore_case") {
		t, err := expandObjectDlpDictionaryEntriesIgnoreCase2edl(d, v, "ignore_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-case"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectDlpDictionaryEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("repeat"); ok || d.HasChange("repeat") {
		t, err := expandObjectDlpDictionaryEntriesRepeat2edl(d, v, "repeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["repeat"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectDlpDictionaryEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectDlpDictionaryEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
