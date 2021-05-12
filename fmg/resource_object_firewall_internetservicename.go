// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Define internet service names.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallInternetServiceName() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallInternetServiceNameCreate,
		Read:   resourceObjectFirewallInternetServiceNameRead,
		Update: resourceObjectFirewallInternetServiceNameUpdate,
		Delete: resourceObjectFirewallInternetServiceNameDelete,

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
			"city_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"country_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"region_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallInternetServiceNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallInternetServiceName(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceName resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallInternetServiceName(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallInternetServiceName resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallInternetServiceNameRead(d, m)
}

func resourceObjectFirewallInternetServiceNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallInternetServiceName(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceName resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallInternetServiceName(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallInternetServiceName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallInternetServiceNameRead(d, m)
}

func resourceObjectFirewallInternetServiceNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallInternetServiceName(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallInternetServiceName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallInternetServiceNameRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallInternetServiceName(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallInternetServiceName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallInternetServiceName resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallInternetServiceNameCityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceNameCountryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceNameInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceNameRegionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallInternetServiceNameType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			3:  "default",
			21: "location",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectFirewallInternetServiceName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("city_id", flattenObjectFirewallInternetServiceNameCityId(o["city-id"], d, "city_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["city-id"], "ObjectFirewallInternetServiceName-CityId"); ok {
			if err = d.Set("city_id", vv); err != nil {
				return fmt.Errorf("Error reading city_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city_id: %v", err)
		}
	}

	if err = d.Set("country_id", flattenObjectFirewallInternetServiceNameCountryId(o["country-id"], d, "country_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-id"], "ObjectFirewallInternetServiceName-CountryId"); ok {
			if err = d.Set("country_id", vv); err != nil {
				return fmt.Errorf("Error reading country_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_id: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenObjectFirewallInternetServiceNameInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "ObjectFirewallInternetServiceName-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallInternetServiceNameName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallInternetServiceName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("region_id", flattenObjectFirewallInternetServiceNameRegionId(o["region-id"], d, "region_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-id"], "ObjectFirewallInternetServiceName-RegionId"); ok {
			if err = d.Set("region_id", vv); err != nil {
				return fmt.Errorf("Error reading region_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_id: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallInternetServiceNameType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallInternetServiceName-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallInternetServiceNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallInternetServiceNameCityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceNameCountryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceNameInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceNameRegionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallInternetServiceNameType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallInternetServiceName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("city_id"); ok {
		t, err := expandObjectFirewallInternetServiceNameCityId(d, v, "city_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city-id"] = t
		}
	}

	if v, ok := d.GetOk("country_id"); ok {
		t, err := expandObjectFirewallInternetServiceNameCountryId(d, v, "country_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok {
		t, err := expandObjectFirewallInternetServiceNameInternetServiceId(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallInternetServiceNameName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("region_id"); ok {
		t, err := expandObjectFirewallInternetServiceNameRegionId(d, v, "region_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectFirewallInternetServiceNameType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
