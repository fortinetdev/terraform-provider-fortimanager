// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectSystem GeoipCountry

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemGeoipCountry() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemGeoipCountryCreate,
		Read:   resourceObjectSystemGeoipCountryRead,
		Update: resourceObjectSystemGeoipCountryUpdate,
		Delete: resourceObjectSystemGeoipCountryDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemGeoipCountryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemGeoipCountry(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipCountry resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemGeoipCountry(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipCountry resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectSystemGeoipCountryRead(d, m)
}

func resourceObjectSystemGeoipCountryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemGeoipCountry(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipCountry resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemGeoipCountry(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipCountry resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceObjectSystemGeoipCountryRead(d, m)
}

func resourceObjectSystemGeoipCountryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemGeoipCountry(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemGeoipCountry resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemGeoipCountryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemGeoipCountry(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipCountry resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemGeoipCountry(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipCountry resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemGeoipCountryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipCountryName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemGeoipCountry(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenObjectSystemGeoipCountryId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemGeoipCountry-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemGeoipCountryName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemGeoipCountry-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemGeoipCountryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemGeoipCountryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipCountryName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemGeoipCountry(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectSystemGeoipCountryId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSystemGeoipCountryName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
