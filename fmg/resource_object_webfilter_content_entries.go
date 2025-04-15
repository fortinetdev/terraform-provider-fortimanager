// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure banned word entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterContentEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterContentEntriesCreate,
		Read:   resourceObjectWebfilterContentEntriesRead,
		Update: resourceObjectWebfilterContentEntriesUpdate,
		Delete: resourceObjectWebfilterContentEntriesDelete,

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
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
		},
	}
}

func resourceObjectWebfilterContentEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	content := d.Get("content").(string)
	paradict["content"] = content

	obj, err := getObjectObjectWebfilterContentEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterContentEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebfilterContentEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterContentEntries resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterContentEntriesRead(d, m)
}

func resourceObjectWebfilterContentEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	content := d.Get("content").(string)
	paradict["content"] = content

	obj, err := getObjectObjectWebfilterContentEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterContentEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterContentEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterContentEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterContentEntriesRead(d, m)
}

func resourceObjectWebfilterContentEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	content := d.Get("content").(string)
	paradict["content"] = content

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebfilterContentEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterContentEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterContentEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	content := d.Get("content").(string)
	if content == "" {
		content = importOptionChecking(m.(*FortiClient).Cfg, "content")
		if content == "" {
			return fmt.Errorf("Parameter content is missing")
		}
		if err = d.Set("content", content); err != nil {
			return fmt.Errorf("Error set params content: %v", err)
		}
	}
	paradict["content"] = content

	o, err := c.ReadObjectWebfilterContentEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterContentEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterContentEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterContentEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterContentEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterContentEntriesLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterContentEntriesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterContentEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterContentEntriesScore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterContentEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterContentEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectWebfilterContentEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWebfilterContentEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("lang", flattenObjectWebfilterContentEntriesLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "ObjectWebfilterContentEntries-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebfilterContentEntriesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebfilterContentEntries-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenObjectWebfilterContentEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "ObjectWebfilterContentEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("score", flattenObjectWebfilterContentEntriesScore2edl(o["score"], d, "score")); err != nil {
		if vv, ok := fortiAPIPatch(o["score"], "ObjectWebfilterContentEntries-Score"); ok {
			if err = d.Set("score", vv); err != nil {
				return fmt.Errorf("Error reading score: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading score: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebfilterContentEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebfilterContentEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterContentEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterContentEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterContentEntriesLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterContentEntriesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterContentEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterContentEntriesScore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterContentEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterContentEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWebfilterContentEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandObjectWebfilterContentEntriesLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebfilterContentEntriesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandObjectWebfilterContentEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("score"); ok || d.HasChange("score") {
		t, err := expandObjectWebfilterContentEntriesScore2edl(d, v, "score")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["score"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectWebfilterContentEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
