// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure venue URL.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpVenueUrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpVenueUrlCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpVenueUrlRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpVenueUrlUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpVenueUrlDelete,

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
			},
			"value_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpVenueUrl(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpVenueUrl resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpVenueUrl(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpVenueUrl resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpVenueUrlRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpVenueUrl(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpVenueUrl resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpVenueUrl(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpVenueUrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpVenueUrlRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerHotspot20AnqpVenueUrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpVenueUrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpVenueUrlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpVenueUrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpVenueUrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpVenueUrl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpVenueUrl resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpVenueUrl-ValueList-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "number"
		if _, ok := i["number"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber(i["number"], d, pre_append)
			tmp["number"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpVenueUrl-ValueList-Number")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20AnqpVenueUrl-ValueList-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpVenueUrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20AnqpVenueUrlName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20AnqpVenueUrl-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("value_list", flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueList(o["value-list"], d, "value_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["value-list"], "ObjectWirelessControllerHotspot20AnqpVenueUrl-ValueList"); ok {
				if err = d.Set("value_list", vv); err != nil {
					return fmt.Errorf("Error reading value_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading value_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("value_list"); ok {
			if err = d.Set("value_list", flattenObjectWirelessControllerHotspot20AnqpVenueUrlValueList(o["value-list"], d, "value_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["value-list"], "ObjectWirelessControllerHotspot20AnqpVenueUrl-ValueList"); ok {
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

func flattenObjectWirelessControllerHotspot20AnqpVenueUrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["number"], _ = expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber(d, i["number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpVenueUrlValueListValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpVenueUrl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueUrlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value_list"); ok || d.HasChange("value_list") {
		t, err := expandObjectWirelessControllerHotspot20AnqpVenueUrlValueList(d, v, "value_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-list"] = t
		}
	}

	return &obj, nil
}
