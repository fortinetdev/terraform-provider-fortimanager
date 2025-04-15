// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VideoFilter FortiGuard category.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVideofilterProfileFortiguardCategoryFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVideofilterProfileFortiguardCategoryFiltersCreate,
		Read:   resourceObjectVideofilterProfileFortiguardCategoryFiltersRead,
		Update: resourceObjectVideofilterProfileFortiguardCategoryFiltersUpdate,
		Delete: resourceObjectVideofilterProfileFortiguardCategoryFiltersDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVideofilterProfileFortiguardCategoryFiltersCreate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectVideofilterProfileFortiguardCategoryFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVideofilterProfileFortiguardCategoryFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectVideofilterProfileFortiguardCategoryFilters(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVideofilterProfileFortiguardCategoryFiltersRead(d, m)
}

func resourceObjectVideofilterProfileFortiguardCategoryFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectVideofilterProfileFortiguardCategoryFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterProfileFortiguardCategoryFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVideofilterProfileFortiguardCategoryFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVideofilterProfileFortiguardCategoryFiltersRead(d, m)
}

func resourceObjectVideofilterProfileFortiguardCategoryFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectVideofilterProfileFortiguardCategoryFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVideofilterProfileFortiguardCategoryFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectVideofilterProfileFortiguardCategoryFilters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterProfileFortiguardCategoryFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVideofilterProfileFortiguardCategoryFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterProfileFortiguardCategoryFilters resource from API: %v", err)
	}
	return nil
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersLog3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectVideofilterProfileFortiguardCategoryFiltersAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectVideofilterProfileFortiguardCategoryFilters-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("category_id", flattenObjectVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(o["category-id"], d, "category_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["category-id"], "ObjectVideofilterProfileFortiguardCategoryFilters-CategoryId"); ok {
			if err = d.Set("category_id", vv); err != nil {
				return fmt.Errorf("Error reading category_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category_id: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectVideofilterProfileFortiguardCategoryFiltersId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectVideofilterProfileFortiguardCategoryFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenObjectVideofilterProfileFortiguardCategoryFiltersLog3rdl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ObjectVideofilterProfileFortiguardCategoryFilters-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	return nil
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersLog3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectVideofilterProfileFortiguardCategoryFiltersAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("category_id"); ok || d.HasChange("category_id") {
		t, err := expandObjectVideofilterProfileFortiguardCategoryFiltersCategoryId3rdl(d, v, "category_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category-id"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectVideofilterProfileFortiguardCategoryFiltersId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandObjectVideofilterProfileFortiguardCategoryFiltersLog3rdl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	return &obj, nil
}
