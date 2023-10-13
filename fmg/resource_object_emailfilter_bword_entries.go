// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Spam filter banned word.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectEmailfilterBwordEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterBwordEntriesCreate,
		Read:   resourceObjectEmailfilterBwordEntriesRead,
		Update: resourceObjectEmailfilterBwordEntriesUpdate,
		Delete: resourceObjectEmailfilterBwordEntriesDelete,

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
			"bword": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"score": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"where": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEmailfilterBwordEntriesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	bword := d.Get("bword").(string)
	paradict["bword"] = bword

	obj, err := getObjectObjectEmailfilterBwordEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterBwordEntries resource while getting object: %v", err)
	}

	_, err = c.CreateObjectEmailfilterBwordEntries(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterBwordEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterBwordEntriesRead(d, m)
}

func resourceObjectEmailfilterBwordEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	bword := d.Get("bword").(string)
	paradict["bword"] = bword

	obj, err := getObjectObjectEmailfilterBwordEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterBwordEntries resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectEmailfilterBwordEntries(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterBwordEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterBwordEntriesRead(d, m)
}

func resourceObjectEmailfilterBwordEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	bword := d.Get("bword").(string)
	paradict["bword"] = bword

	err = c.DeleteObjectEmailfilterBwordEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterBwordEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterBwordEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	bword := d.Get("bword").(string)
	if bword == "" {
		bword = importOptionChecking(m.(*FortiClient).Cfg, "bword")
		if bword == "" {
			return fmt.Errorf("Parameter bword is missing")
		}
		if err = d.Set("bword", bword); err != nil {
			return fmt.Errorf("Error set params bword: %v", err)
		}
	}
	paradict["bword"] = bword

	o, err := c.ReadObjectEmailfilterBwordEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterBwordEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterBwordEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterBwordEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterBwordEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesLanguage2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesScore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterBwordEntriesWhere2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterBwordEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectEmailfilterBwordEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectEmailfilterBwordEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectEmailfilterBwordEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectEmailfilterBwordEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("language", flattenObjectEmailfilterBwordEntriesLanguage2edl(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "ObjectEmailfilterBwordEntries-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectEmailfilterBwordEntriesPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectEmailfilterBwordEntries-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenObjectEmailfilterBwordEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "ObjectEmailfilterBwordEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("score", flattenObjectEmailfilterBwordEntriesScore2edl(o["score"], d, "score")); err != nil {
		if vv, ok := fortiAPIPatch(o["score"], "ObjectEmailfilterBwordEntries-Score"); ok {
			if err = d.Set("score", vv); err != nil {
				return fmt.Errorf("Error reading score: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading score: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectEmailfilterBwordEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectEmailfilterBwordEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("where", flattenObjectEmailfilterBwordEntriesWhere2edl(o["where"], d, "where")); err != nil {
		if vv, ok := fortiAPIPatch(o["where"], "ObjectEmailfilterBwordEntries-Where"); ok {
			if err = d.Set("where", vv); err != nil {
				return fmt.Errorf("Error reading where: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading where: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterBwordEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterBwordEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesLanguage2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesScore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterBwordEntriesWhere2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterBwordEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectEmailfilterBwordEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectEmailfilterBwordEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandObjectEmailfilterBwordEntriesLanguage2edl(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectEmailfilterBwordEntriesPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandObjectEmailfilterBwordEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("score"); ok || d.HasChange("score") {
		t, err := expandObjectEmailfilterBwordEntriesScore2edl(d, v, "score")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["score"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectEmailfilterBwordEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("where"); ok || d.HasChange("where") {
		t, err := expandObjectEmailfilterBwordEntriesWhere2edl(d, v, "where")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["where"] = t
		}
	}

	return &obj, nil
}
