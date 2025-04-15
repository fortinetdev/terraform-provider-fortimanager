// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Spam filter mime header content.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectEmailfilterMheaderEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterMheaderEntriesCreate,
		Read:   resourceObjectEmailfilterMheaderEntriesRead,
		Update: resourceObjectEmailfilterMheaderEntriesUpdate,
		Delete: resourceObjectEmailfilterMheaderEntriesDelete,

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
			"mheader": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"pattern_type": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectEmailfilterMheaderEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	mheader := d.Get("mheader").(string)
	paradict["mheader"] = mheader

	obj, err := getObjectObjectEmailfilterMheaderEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterMheaderEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectEmailfilterMheaderEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterMheaderEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterMheaderEntriesRead(d, m)
}

func resourceObjectEmailfilterMheaderEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	mheader := d.Get("mheader").(string)
	paradict["mheader"] = mheader

	obj, err := getObjectObjectEmailfilterMheaderEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterMheaderEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectEmailfilterMheaderEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterMheaderEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectEmailfilterMheaderEntriesRead(d, m)
}

func resourceObjectEmailfilterMheaderEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	mheader := d.Get("mheader").(string)
	paradict["mheader"] = mheader

	wsParams["adom"] = adomv

	err = c.DeleteObjectEmailfilterMheaderEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterMheaderEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterMheaderEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	mheader := d.Get("mheader").(string)
	if mheader == "" {
		mheader = importOptionChecking(m.(*FortiClient).Cfg, "mheader")
		if mheader == "" {
			return fmt.Errorf("Parameter mheader is missing")
		}
		if err = d.Set("mheader", mheader); err != nil {
			return fmt.Errorf("Error set params mheader: %v", err)
		}
	}
	paradict["mheader"] = mheader

	o, err := c.ReadObjectEmailfilterMheaderEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterMheaderEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterMheaderEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterMheaderEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterMheaderEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterMheaderEntriesFieldbody2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterMheaderEntriesFieldname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterMheaderEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterMheaderEntriesPatternType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterMheaderEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterMheaderEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectEmailfilterMheaderEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectEmailfilterMheaderEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fieldbody", flattenObjectEmailfilterMheaderEntriesFieldbody2edl(o["fieldbody"], d, "fieldbody")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldbody"], "ObjectEmailfilterMheaderEntries-Fieldbody"); ok {
			if err = d.Set("fieldbody", vv); err != nil {
				return fmt.Errorf("Error reading fieldbody: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldbody: %v", err)
		}
	}

	if err = d.Set("fieldname", flattenObjectEmailfilterMheaderEntriesFieldname2edl(o["fieldname"], d, "fieldname")); err != nil {
		if vv, ok := fortiAPIPatch(o["fieldname"], "ObjectEmailfilterMheaderEntries-Fieldname"); ok {
			if err = d.Set("fieldname", vv); err != nil {
				return fmt.Errorf("Error reading fieldname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fieldname: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectEmailfilterMheaderEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectEmailfilterMheaderEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("pattern_type", flattenObjectEmailfilterMheaderEntriesPatternType2edl(o["pattern-type"], d, "pattern_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern-type"], "ObjectEmailfilterMheaderEntries-PatternType"); ok {
			if err = d.Set("pattern_type", vv); err != nil {
				return fmt.Errorf("Error reading pattern_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern_type: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectEmailfilterMheaderEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectEmailfilterMheaderEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterMheaderEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterMheaderEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterMheaderEntriesFieldbody2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterMheaderEntriesFieldname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterMheaderEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterMheaderEntriesPatternType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterMheaderEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterMheaderEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectEmailfilterMheaderEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fieldbody"); ok || d.HasChange("fieldbody") {
		t, err := expandObjectEmailfilterMheaderEntriesFieldbody2edl(d, v, "fieldbody")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldbody"] = t
		}
	}

	if v, ok := d.GetOk("fieldname"); ok || d.HasChange("fieldname") {
		t, err := expandObjectEmailfilterMheaderEntriesFieldname2edl(d, v, "fieldname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fieldname"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectEmailfilterMheaderEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("pattern_type"); ok || d.HasChange("pattern_type") {
		t, err := expandObjectEmailfilterMheaderEntriesPatternType2edl(d, v, "pattern_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectEmailfilterMheaderEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
