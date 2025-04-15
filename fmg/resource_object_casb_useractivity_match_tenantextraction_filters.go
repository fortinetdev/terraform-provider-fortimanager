// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB user activity tenant extraction filters.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityMatchTenantExtractionFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityMatchTenantExtractionFiltersCreate,
		Read:   resourceObjectCasbUserActivityMatchTenantExtractionFiltersRead,
		Update: resourceObjectCasbUserActivityMatchTenantExtractionFiltersUpdate,
		Delete: resourceObjectCasbUserActivityMatchTenantExtractionFiltersDelete,

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
				Computed: true,
			},
			"place": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectCasbUserActivityMatchTenantExtractionFiltersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbUserActivityMatchTenantExtractionFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatchTenantExtractionFilters resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	v, err := c.CreateObjectCasbUserActivityMatchTenantExtractionFilters(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceObjectCasbUserActivityMatchTenantExtractionFiltersRead(d, m)
		} else {
			return fmt.Errorf("Error creating ObjectCasbUserActivityMatchTenantExtractionFilters resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchTenantExtractionFiltersRead(d, m)
}

func resourceObjectCasbUserActivityMatchTenantExtractionFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbUserActivityMatchTenantExtractionFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantExtractionFilters resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityMatchTenantExtractionFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchTenantExtractionFiltersRead(d, m)
}

func resourceObjectCasbUserActivityMatchTenantExtractionFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCasbUserActivityMatchTenantExtractionFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityMatchTenantExtractionFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCasbUserActivityMatchTenantExtractionFilters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantExtractionFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityMatchTenantExtractionFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantExtractionFilters resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersDirection4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersPlace4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityMatchTenantExtractionFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("body_type", flattenObjectCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(o["body-type"], d, "body_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["body-type"], "ObjectCasbUserActivityMatchTenantExtractionFilters-BodyType"); ok {
			if err = d.Set("body_type", vv); err != nil {
				return fmt.Errorf("Error reading body_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading body_type: %v", err)
		}
	}

	if err = d.Set("direction", flattenObjectCasbUserActivityMatchTenantExtractionFiltersDirection4thl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "ObjectCasbUserActivityMatchTenantExtractionFilters-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("header_name", flattenObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "ObjectCasbUserActivityMatchTenantExtractionFilters-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectCasbUserActivityMatchTenantExtractionFiltersId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectCasbUserActivityMatchTenantExtractionFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("place", flattenObjectCasbUserActivityMatchTenantExtractionFiltersPlace4thl(o["place"], d, "place")); err != nil {
		if vv, ok := fortiAPIPatch(o["place"], "ObjectCasbUserActivityMatchTenantExtractionFilters-Place"); ok {
			if err = d.Set("place", vv); err != nil {
				return fmt.Errorf("Error reading place: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading place: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersDirection4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersPlace4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityMatchTenantExtractionFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("body_type"); ok || d.HasChange("body_type") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFiltersBodyType4thl(d, v, "body_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["body-type"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFiltersDirection4thl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName4thl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFiltersId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("place"); ok || d.HasChange("place") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFiltersPlace4thl(d, v, "place")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["place"] = t
		}
	}

	return &obj, nil
}
