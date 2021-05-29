// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure venue name duple.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpVenueName() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpVenueNameCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpVenueNameRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpVenueNameUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpVenueNameDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"value_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"lang": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20AnqpVenueNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpVenueName(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpVenueName resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpVenueName(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpVenueNameRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpVenueNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpVenueName(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpVenueName resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpVenueName(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpVenueNameRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpVenueNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20AnqpVenueName(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpVenueNameRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpVenueName(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpVenueName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpVenueName resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpVenueNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueNameValueList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpVenueNameValueListIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpVenueName-ValueList-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpVenueNameValueListLang(i["lang"], d, pre_append)
			tmp["lang"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpVenueName-ValueList-Lang")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpVenueNameValueListValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpVenueName-ValueList-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpVenueNameValueListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueNameValueListLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueNameValueListValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpVenueName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20AnqpVenueNameName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20AnqpVenueName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("value_list", flattenObjectWirelessControllerHotspot20AnqpVenueNameValueList(o["value-list"], d, "value_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["value-list"], "ObjectWirelessControllerHotspot20AnqpVenueName-ValueList"); ok {
				if err = d.Set("value_list", vv); err != nil {
					return fmt.Errorf("Error reading value_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading value_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("value_list"); ok {
			if err = d.Set("value_list", flattenObjectWirelessControllerHotspot20AnqpVenueNameValueList(o["value-list"], d, "value_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["value-list"], "ObjectWirelessControllerHotspot20AnqpVenueName-ValueList"); ok {
					if err = d.Set("value_list", vv); err != nil {
						return fmt.Errorf("Error reading value_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading value_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpVenueNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpVenueNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueNameValueList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20AnqpVenueNameValueListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lang"], _ = expandObjectWirelessControllerHotspot20AnqpVenueNameValueListLang(d, i["lang"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandObjectWirelessControllerHotspot20AnqpVenueNameValueListValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueNameValueListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueNameValueListLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueNameValueListValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpVenueName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueNameName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value_list"); ok {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueNameValueList(d, v, "value_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-list"] = t
		}
	}

	return &obj, nil
}
