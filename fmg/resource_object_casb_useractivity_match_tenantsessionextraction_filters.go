// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB user activity session extraction filters.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityMatchTenantSessionExtractionFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersCreate,
		Read:   resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersRead,
		Update: resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersUpdate,
		Delete: resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"place": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersCreate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectObjectCasbUserActivityMatchTenantSessionExtractionFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectCasbUserActivityMatchTenantSessionExtractionFilters(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectCasbUserActivityMatchTenantSessionExtractionFilters(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectCasbUserActivityMatchTenantSessionExtractionFilters(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersRead(d, m)
}

func resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectObjectCasbUserActivityMatchTenantSessionExtractionFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityMatchTenantSessionExtractionFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersRead(d, m)
}

func resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbUserActivityMatchTenantSessionExtractionFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityMatchTenantSessionExtractionFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	if match == "" {
		match = importOptionChecking(m.(*FortiClient).Cfg, "match")
		if match == "" {
			return fmt.Errorf("Parameter match is missing")
		}
		if err = d.Set("match", match); err != nil {
			return fmt.Errorf("Error set params match: %v", err)
		}
	}
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	o, err := c.ReadObjectCasbUserActivityMatchTenantSessionExtractionFilters(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityMatchTenantSessionExtractionFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantSessionExtractionFilters resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityMatchTenantSessionExtractionFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("body_type", flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType4thl(o["body-type"], d, "body_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["body-type"], "ObjectCasbUserActivityMatchTenantSessionExtractionFilters-BodyType"); ok {
			if err = d.Set("body_type", vv); err != nil {
				return fmt.Errorf("Error reading body_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading body_type: %v", err)
		}
	}

	if err = d.Set("cookie_name", flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName4thl(o["cookie-name"], d, "cookie_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["cookie-name"], "ObjectCasbUserActivityMatchTenantSessionExtractionFilters-CookieName"); ok {
			if err = d.Set("cookie_name", vv); err != nil {
				return fmt.Errorf("Error reading cookie_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cookie_name: %v", err)
		}
	}

	if err = d.Set("direction", flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection4thl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "ObjectCasbUserActivityMatchTenantSessionExtractionFilters-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("header_name", flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName4thl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "ObjectCasbUserActivityMatchTenantSessionExtractionFilters-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectCasbUserActivityMatchTenantSessionExtractionFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("place", flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace4thl(o["place"], d, "place")); err != nil {
		if vv, ok := fortiAPIPatch(o["place"], "ObjectCasbUserActivityMatchTenantSessionExtractionFilters-Place"); ok {
			if err = d.Set("place", vv); err != nil {
				return fmt.Errorf("Error reading place: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading place: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityMatchTenantSessionExtractionFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("body_type"); ok || d.HasChange("body_type") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType4thl(d, v, "body_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["body-type"] = t
		}
	}

	if v, ok := d.GetOk("cookie_name"); ok || d.HasChange("cookie_name") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName4thl(d, v, "cookie_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-name"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection4thl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName4thl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("place"); ok || d.HasChange("place") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace4thl(d, v, "place")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["place"] = t
		}
	}

	return &obj, nil
}
