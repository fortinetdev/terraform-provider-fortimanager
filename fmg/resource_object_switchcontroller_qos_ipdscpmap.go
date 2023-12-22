// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS IP precedence/DSCP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerQosIpDscpMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerQosIpDscpMapCreate,
		Read:   resourceObjectSwitchControllerQosIpDscpMapRead,
		Update: resourceObjectSwitchControllerQosIpDscpMapUpdate,
		Delete: resourceObjectSwitchControllerQosIpDscpMapDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"map": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos_queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"diffserv": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip_precedence": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectSwitchControllerQosIpDscpMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerQosIpDscpMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosIpDscpMap resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerQosIpDscpMap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosIpDscpMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosIpDscpMapRead(d, m)
}

func resourceObjectSwitchControllerQosIpDscpMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerQosIpDscpMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosIpDscpMap resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerQosIpDscpMap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosIpDscpMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosIpDscpMapRead(d, m)
}

func resourceObjectSwitchControllerQosIpDscpMapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerQosIpDscpMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerQosIpDscpMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerQosIpDscpMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerQosIpDscpMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosIpDscpMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerQosIpDscpMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosIpDscpMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerQosIpDscpMapDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosIpDscpMapMap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := i["cos-queue"]; ok {
			v := flattenObjectSwitchControllerQosIpDscpMapMapCosQueue(i["cos-queue"], d, pre_append)
			tmp["cos_queue"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosIpDscpMap-Map-CosQueue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := i["diffserv"]; ok {
			v := flattenObjectSwitchControllerQosIpDscpMapMapDiffserv(i["diffserv"], d, pre_append)
			tmp["diffserv"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosIpDscpMap-Map-Diffserv")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := i["ip-precedence"]; ok {
			v := flattenObjectSwitchControllerQosIpDscpMapMapIpPrecedence(i["ip-precedence"], d, pre_append)
			tmp["ip_precedence"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosIpDscpMap-Map-IpPrecedence")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSwitchControllerQosIpDscpMapMapName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosIpDscpMap-Map-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectSwitchControllerQosIpDscpMapMapValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerQosIpDscpMap-Map-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerQosIpDscpMapMapCosQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosIpDscpMapMapDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerQosIpDscpMapMapIpPrecedence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerQosIpDscpMapMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosIpDscpMapMapValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosIpDscpMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerQosIpDscpMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenObjectSwitchControllerQosIpDscpMapDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerQosIpDscpMap-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("map", flattenObjectSwitchControllerQosIpDscpMapMap(o["map"], d, "map")); err != nil {
			if vv, ok := fortiAPIPatch(o["map"], "ObjectSwitchControllerQosIpDscpMap-Map"); ok {
				if err = d.Set("map", vv); err != nil {
					return fmt.Errorf("Error reading map: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading map: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("map"); ok {
			if err = d.Set("map", flattenObjectSwitchControllerQosIpDscpMapMap(o["map"], d, "map")); err != nil {
				if vv, ok := fortiAPIPatch(o["map"], "ObjectSwitchControllerQosIpDscpMap-Map"); ok {
					if err = d.Set("map", vv); err != nil {
						return fmt.Errorf("Error reading map: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading map: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerQosIpDscpMapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerQosIpDscpMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerQosIpDscpMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerQosIpDscpMapDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosIpDscpMapMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos-queue"], _ = expandObjectSwitchControllerQosIpDscpMapMapCosQueue(d, i["cos_queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffserv"], _ = expandObjectSwitchControllerQosIpDscpMapMapDiffserv(d, i["diffserv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-precedence"], _ = expandObjectSwitchControllerQosIpDscpMapMapIpPrecedence(d, i["ip_precedence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSwitchControllerQosIpDscpMapMapName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectSwitchControllerQosIpDscpMapMapValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerQosIpDscpMapMapCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosIpDscpMapMapDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerQosIpDscpMapMapIpPrecedence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerQosIpDscpMapMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosIpDscpMapMapValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosIpDscpMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerQosIpDscpMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerQosIpDscpMapDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("map"); ok || d.HasChange("map") {
		t, err := expandObjectSwitchControllerQosIpDscpMapMap(d, v, "map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["map"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerQosIpDscpMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
