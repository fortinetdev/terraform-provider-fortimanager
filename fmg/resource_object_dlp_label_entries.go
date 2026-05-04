// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DLP label entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpLabelEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpLabelEntriesCreate,
		Read:   resourceObjectDlpLabelEntriesRead,
		Update: resourceObjectDlpLabelEntriesUpdate,
		Delete: resourceObjectDlpLabelEntriesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fortidata_label_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"mpip_label_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectDlpLabelEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	label := d.Get("label").(string)
	paradict["label"] = label

	obj, err := getObjectObjectDlpLabelEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpLabelEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectDlpLabelEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectDlpLabelEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectDlpLabelEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectDlpLabelEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectDlpLabelEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpLabelEntriesRead(d, m)
}

func resourceObjectDlpLabelEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	label := d.Get("label").(string)
	paradict["label"] = label

	obj, err := getObjectObjectDlpLabelEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpLabelEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDlpLabelEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpLabelEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpLabelEntriesRead(d, m)
}

func resourceObjectDlpLabelEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	label := d.Get("label").(string)
	paradict["label"] = label

	wsParams["adom"] = adomv

	err = c.DeleteObjectDlpLabelEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpLabelEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpLabelEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	label := d.Get("label").(string)
	if label == "" {
		label = importOptionChecking(m.(*FortiClient).Cfg, "label")
		if label == "" {
			return fmt.Errorf("Parameter label is missing")
		}
		if err = d.Set("label", label); err != nil {
			return fmt.Errorf("Error set params label: %v", err)
		}
	}
	paradict["label"] = label

	o, err := c.ReadObjectDlpLabelEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectDlpLabelEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpLabelEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpLabelEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpLabelEntriesFortidataLabelName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelEntriesGuid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpLabelEntriesMpipLabelName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpLabelEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fortidata_label_name", flattenObjectDlpLabelEntriesFortidataLabelName2edl(o["fortidata-label-name"], d, "fortidata_label_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortidata-label-name"], "ObjectDlpLabelEntries-FortidataLabelName"); ok {
			if err = d.Set("fortidata_label_name", vv); err != nil {
				return fmt.Errorf("Error reading fortidata_label_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortidata_label_name: %v", err)
		}
	}

	if err = d.Set("guid", flattenObjectDlpLabelEntriesGuid2edl(o["guid"], d, "guid")); err != nil {
		if vv, ok := fortiAPIPatch(o["guid"], "ObjectDlpLabelEntries-Guid"); ok {
			if err = d.Set("guid", vv); err != nil {
				return fmt.Errorf("Error reading guid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guid: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectDlpLabelEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDlpLabelEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mpip_label_name", flattenObjectDlpLabelEntriesMpipLabelName2edl(o["mpip-label-name"], d, "mpip_label_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpip-label-name"], "ObjectDlpLabelEntries-MpipLabelName"); ok {
			if err = d.Set("mpip_label_name", vv); err != nil {
				return fmt.Errorf("Error reading mpip_label_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpip_label_name: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpLabelEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpLabelEntriesFortidataLabelName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelEntriesGuid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpLabelEntriesMpipLabelName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpLabelEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fortidata_label_name"); ok || d.HasChange("fortidata_label_name") {
		t, err := expandObjectDlpLabelEntriesFortidataLabelName2edl(d, v, "fortidata_label_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortidata-label-name"] = t
		}
	}

	if v, ok := d.GetOk("guid"); ok || d.HasChange("guid") {
		t, err := expandObjectDlpLabelEntriesGuid2edl(d, v, "guid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guid"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectDlpLabelEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mpip_label_name"); ok || d.HasChange("mpip_label_name") {
		t, err := expandObjectDlpLabelEntriesMpipLabelName2edl(d, v, "mpip_label_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpip-label-name"] = t
		}
	}

	return &obj, nil
}
