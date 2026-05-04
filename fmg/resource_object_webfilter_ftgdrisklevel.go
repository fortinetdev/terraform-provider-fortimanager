// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter risk level.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterFtgdRiskLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterFtgdRiskLevelCreate,
		Read:   resourceObjectWebfilterFtgdRiskLevelRead,
		Update: resourceObjectWebfilterFtgdRiskLevelUpdate,
		Delete: resourceObjectWebfilterFtgdRiskLevelDelete,

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
			"high": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"low": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterFtgdRiskLevelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterFtgdRiskLevel(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterFtgdRiskLevel resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectWebfilterFtgdRiskLevel(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectWebfilterFtgdRiskLevel(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectWebfilterFtgdRiskLevel resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectWebfilterFtgdRiskLevel(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectWebfilterFtgdRiskLevel resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterFtgdRiskLevelRead(d, m)
}

func resourceObjectWebfilterFtgdRiskLevelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterFtgdRiskLevel(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterFtgdRiskLevel resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterFtgdRiskLevel(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterFtgdRiskLevel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebfilterFtgdRiskLevelRead(d, m)
}

func resourceObjectWebfilterFtgdRiskLevelDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebfilterFtgdRiskLevel(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterFtgdRiskLevel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterFtgdRiskLevelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterFtgdRiskLevel(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectWebfilterFtgdRiskLevel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterFtgdRiskLevel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterFtgdRiskLevel resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterFtgdRiskLevelHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterFtgdRiskLevelLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterFtgdRiskLevelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterFtgdRiskLevel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("high", flattenObjectWebfilterFtgdRiskLevelHigh(o["high"], d, "high")); err != nil {
		if vv, ok := fortiAPIPatch(o["high"], "ObjectWebfilterFtgdRiskLevel-High"); ok {
			if err = d.Set("high", vv); err != nil {
				return fmt.Errorf("Error reading high: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading high: %v", err)
		}
	}

	if err = d.Set("low", flattenObjectWebfilterFtgdRiskLevelLow(o["low"], d, "low")); err != nil {
		if vv, ok := fortiAPIPatch(o["low"], "ObjectWebfilterFtgdRiskLevel-Low"); ok {
			if err = d.Set("low", vv); err != nil {
				return fmt.Errorf("Error reading low: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebfilterFtgdRiskLevelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebfilterFtgdRiskLevel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterFtgdRiskLevelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterFtgdRiskLevelHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterFtgdRiskLevelLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterFtgdRiskLevelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterFtgdRiskLevel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("high"); ok || d.HasChange("high") {
		t, err := expandObjectWebfilterFtgdRiskLevelHigh(d, v, "high")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high"] = t
		}
	}

	if v, ok := d.GetOk("low"); ok || d.HasChange("low") {
		t, err := expandObjectWebfilterFtgdRiskLevelLow(d, v, "low")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebfilterFtgdRiskLevelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
