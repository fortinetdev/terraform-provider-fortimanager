// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure QoS map set.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20QosMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20QosMapCreate,
		Read:   resourceObjectWirelessControllerHotspot20QosMapRead,
		Update: resourceObjectWirelessControllerHotspot20QosMapUpdate,
		Delete: resourceObjectWirelessControllerHotspot20QosMapDelete,

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
			"dscp_except": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"up": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dscp_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"up": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20QosMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20QosMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20QosMap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20QosMap(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20QosMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20QosMapRead(d, m)
}

func resourceObjectWirelessControllerHotspot20QosMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20QosMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20QosMap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20QosMap(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20QosMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20QosMapRead(d, m)
}

func resourceObjectWirelessControllerHotspot20QosMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20QosMap(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20QosMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20QosMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20QosMap(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20QosMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20QosMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20QosMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExcept(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpExceptDscp(i["dscp"], d, pre_append)
			tmp["dscp"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpExcept-Dscp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpExceptIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpExcept-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := i["up"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpExceptUp(i["up"], d, pre_append)
			tmp["up"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpExcept-Up")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpExceptUp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high"
		if _, ok := i["high"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpRangeHigh(i["high"], d, pre_append)
			tmp["high"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpRange-High")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpRangeIndex(i["index"], d, pre_append)
			tmp["index"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpRange-Index")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "low"
		if _, ok := i["low"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpRangeLow(i["low"], d, pre_append)
			tmp["low"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpRange-Low")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := i["up"]; ok {
			v := flattenObjectWirelessControllerHotspot20QosMapDscpRangeUp(i["up"], d, pre_append)
			tmp["up"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20QosMap-DscpRange-Up")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapDscpRangeUp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20QosMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("dscp_except", flattenObjectWirelessControllerHotspot20QosMapDscpExcept(o["dscp-except"], d, "dscp_except")); err != nil {
			if vv, ok := fortiAPIPatch(o["dscp-except"], "ObjectWirelessControllerHotspot20QosMap-DscpExcept"); ok {
				if err = d.Set("dscp_except", vv); err != nil {
					return fmt.Errorf("Error reading dscp_except: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dscp_except: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_except"); ok {
			if err = d.Set("dscp_except", flattenObjectWirelessControllerHotspot20QosMapDscpExcept(o["dscp-except"], d, "dscp_except")); err != nil {
				if vv, ok := fortiAPIPatch(o["dscp-except"], "ObjectWirelessControllerHotspot20QosMap-DscpExcept"); ok {
					if err = d.Set("dscp_except", vv); err != nil {
						return fmt.Errorf("Error reading dscp_except: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dscp_except: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_range", flattenObjectWirelessControllerHotspot20QosMapDscpRange(o["dscp-range"], d, "dscp_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["dscp-range"], "ObjectWirelessControllerHotspot20QosMap-DscpRange"); ok {
				if err = d.Set("dscp_range", vv); err != nil {
					return fmt.Errorf("Error reading dscp_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dscp_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_range"); ok {
			if err = d.Set("dscp_range", flattenObjectWirelessControllerHotspot20QosMapDscpRange(o["dscp-range"], d, "dscp_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["dscp-range"], "ObjectWirelessControllerHotspot20QosMap-DscpRange"); ok {
					if err = d.Set("dscp_range", vv); err != nil {
						return fmt.Errorf("Error reading dscp_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dscp_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20QosMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20QosMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20QosMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp"], _ = expandObjectWirelessControllerHotspot20QosMapDscpExceptDscp(d, i["dscp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20QosMapDscpExceptIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["up"], _ = expandObjectWirelessControllerHotspot20QosMapDscpExceptUp(d, i["up"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpExceptDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpExceptIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpExceptUp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["high"], _ = expandObjectWirelessControllerHotspot20QosMapDscpRangeHigh(d, i["high"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["index"], _ = expandObjectWirelessControllerHotspot20QosMapDscpRangeIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "low"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["low"], _ = expandObjectWirelessControllerHotspot20QosMapDscpRangeLow(d, i["low"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["up"], _ = expandObjectWirelessControllerHotspot20QosMapDscpRangeUp(d, i["up"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapDscpRangeUp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20QosMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dscp_except"); ok || d.HasChange("dscp_except") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpExcept(d, v, "dscp_except")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-except"] = t
		}
	}

	if v, ok := d.GetOk("dscp_range"); ok || d.HasChange("dscp_range") {
		t, err := expandObjectWirelessControllerHotspot20QosMapDscpRange(d, v, "dscp_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-range"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20QosMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
