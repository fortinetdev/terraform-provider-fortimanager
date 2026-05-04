// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DLP exact-data-match column types.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpExactDataMatchColumns() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpExactDataMatchColumnsCreate,
		Read:   resourceObjectDlpExactDataMatchColumnsRead,
		Update: resourceObjectDlpExactDataMatchColumnsUpdate,
		Delete: resourceObjectDlpExactDataMatchColumnsDelete,

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
			"exact_data_match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"optional": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectDlpExactDataMatchColumnsCreate(d *schema.ResourceData, m interface{}) error {
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

	exact_data_match := d.Get("exact_data_match").(string)
	paradict["exact_data_match"] = exact_data_match

	obj, err := getObjectObjectDlpExactDataMatchColumns(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpExactDataMatchColumns resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("index")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectDlpExactDataMatchColumns(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectDlpExactDataMatchColumns(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectDlpExactDataMatchColumns resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectDlpExactDataMatchColumns(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectDlpExactDataMatchColumns resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectDlpExactDataMatchColumnsRead(d, m)
}

func resourceObjectDlpExactDataMatchColumnsUpdate(d *schema.ResourceData, m interface{}) error {
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

	exact_data_match := d.Get("exact_data_match").(string)
	paradict["exact_data_match"] = exact_data_match

	obj, err := getObjectObjectDlpExactDataMatchColumns(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpExactDataMatchColumns resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDlpExactDataMatchColumns(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpExactDataMatchColumns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectDlpExactDataMatchColumnsRead(d, m)
}

func resourceObjectDlpExactDataMatchColumnsDelete(d *schema.ResourceData, m interface{}) error {
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

	exact_data_match := d.Get("exact_data_match").(string)
	paradict["exact_data_match"] = exact_data_match

	wsParams["adom"] = adomv

	err = c.DeleteObjectDlpExactDataMatchColumns(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpExactDataMatchColumns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpExactDataMatchColumnsRead(d *schema.ResourceData, m interface{}) error {
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

	exact_data_match := d.Get("exact_data_match").(string)
	if exact_data_match == "" {
		exact_data_match = importOptionChecking(m.(*FortiClient).Cfg, "exact_data_match")
		if exact_data_match == "" {
			return fmt.Errorf("Parameter exact_data_match is missing")
		}
		if err = d.Set("exact_data_match", exact_data_match); err != nil {
			return fmt.Errorf("Error set params exact_data_match: %v", err)
		}
	}
	paradict["exact_data_match"] = exact_data_match

	o, err := c.ReadObjectDlpExactDataMatchColumns(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectDlpExactDataMatchColumns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpExactDataMatchColumns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpExactDataMatchColumns resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpExactDataMatchColumnsIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpExactDataMatchColumnsOptional2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpExactDataMatchColumnsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectDlpExactDataMatchColumns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("index", flattenObjectDlpExactDataMatchColumnsIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectDlpExactDataMatchColumns-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("optional", flattenObjectDlpExactDataMatchColumnsOptional2edl(o["optional"], d, "optional")); err != nil {
		if vv, ok := fortiAPIPatch(o["optional"], "ObjectDlpExactDataMatchColumns-Optional"); ok {
			if err = d.Set("optional", vv); err != nil {
				return fmt.Errorf("Error reading optional: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optional: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectDlpExactDataMatchColumnsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectDlpExactDataMatchColumns-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpExactDataMatchColumnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpExactDataMatchColumnsIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpExactDataMatchColumnsOptional2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpExactDataMatchColumnsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectDlpExactDataMatchColumns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectDlpExactDataMatchColumnsIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("optional"); ok || d.HasChange("optional") {
		t, err := expandObjectDlpExactDataMatchColumnsOptional2edl(d, v, "optional")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectDlpExactDataMatchColumnsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
