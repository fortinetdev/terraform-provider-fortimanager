// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter local categories.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWebfilterFtgdLocalCat() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterFtgdLocalCatCreate,
		Read:   resourceObjectWebfilterFtgdLocalCatRead,
		Update: resourceObjectWebfilterFtgdLocalCatUpdate,
		Delete: resourceObjectWebfilterFtgdLocalCatDelete,

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
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
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

func resourceObjectWebfilterFtgdLocalCatCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebfilterFtgdLocalCat(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterFtgdLocalCat resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebfilterFtgdLocalCat(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterFtgdLocalCat resource: %v", err)
	}

	d.SetId(getStringKey(d, "desc"))

	return resourceObjectWebfilterFtgdLocalCatRead(d, m)
}

func resourceObjectWebfilterFtgdLocalCatUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebfilterFtgdLocalCat(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterFtgdLocalCat resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebfilterFtgdLocalCat(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterFtgdLocalCat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "desc"))

	return resourceObjectWebfilterFtgdLocalCatRead(d, m)
}

func resourceObjectWebfilterFtgdLocalCatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWebfilterFtgdLocalCat(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterFtgdLocalCat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterFtgdLocalCatRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWebfilterFtgdLocalCat(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterFtgdLocalCat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterFtgdLocalCat(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterFtgdLocalCat resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterFtgdLocalCatDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterFtgdLocalCatId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterFtgdLocalCatStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectWebfilterFtgdLocalCat(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("desc", flattenObjectWebfilterFtgdLocalCatDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "ObjectWebfilterFtgdLocalCat-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectWebfilterFtgdLocalCatId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWebfilterFtgdLocalCat-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebfilterFtgdLocalCatStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebfilterFtgdLocalCat-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterFtgdLocalCatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterFtgdLocalCatDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterFtgdLocalCatId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterFtgdLocalCatStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterFtgdLocalCat(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandObjectWebfilterFtgdLocalCatDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectWebfilterFtgdLocalCatId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandObjectWebfilterFtgdLocalCatStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
