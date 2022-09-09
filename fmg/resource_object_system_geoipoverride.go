// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemGeoipOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemGeoipOverrideCreate,
		Read:   resourceObjectSystemGeoipOverrideRead,
		Update: resourceObjectSystemGeoipOverrideUpdate,
		Delete: resourceObjectSystemGeoipOverrideDelete,

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
			"country_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip6_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
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

func resourceObjectSystemGeoipOverrideCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemGeoipOverride(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipOverride resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemGeoipOverride(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemGeoipOverride resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemGeoipOverrideRead(d, m)
}

func resourceObjectSystemGeoipOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemGeoipOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemGeoipOverride(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemGeoipOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemGeoipOverrideRead(d, m)
}

func resourceObjectSystemGeoipOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemGeoipOverride(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemGeoipOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemGeoipOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemGeoipOverride(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemGeoipOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemGeoipOverride resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemGeoipOverrideCountryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectSystemGeoipOverrideIpRangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectSystemGeoipOverride-IpRange-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemGeoipOverrideIpRangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemGeoipOverride-IpRange-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectSystemGeoipOverrideIpRangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectSystemGeoipOverride-IpRange-StartIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemGeoipOverrideIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIp6Range(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			v := flattenObjectSystemGeoipOverrideIp6RangeEndIp(i["end-ip"], d, pre_append)
			tmp["end_ip"] = fortiAPISubPartPatch(v, "ObjectSystemGeoipOverride-Ip6Range-EndIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectSystemGeoipOverrideIp6RangeId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectSystemGeoipOverride-Ip6Range-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			v := flattenObjectSystemGeoipOverrideIp6RangeStartIp(i["start-ip"], d, pre_append)
			tmp["start_ip"] = fortiAPISubPartPatch(v, "ObjectSystemGeoipOverride-Ip6Range-StartIp")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemGeoipOverrideIp6RangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIp6RangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideIp6RangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemGeoipOverrideName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemGeoipOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("country_id", flattenObjectSystemGeoipOverrideCountryId(o["country-id"], d, "country_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-id"], "ObjectSystemGeoipOverride-CountryId"); ok {
			if err = d.Set("country_id", vv); err != nil {
				return fmt.Errorf("Error reading country_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_id: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSystemGeoipOverrideDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSystemGeoipOverride-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenObjectSystemGeoipOverrideIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectSystemGeoipOverride-IpRange"); ok {
				if err = d.Set("ip_range", vv); err != nil {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenObjectSystemGeoipOverrideIpRange(o["ip-range"], d, "ip_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-range"], "ObjectSystemGeoipOverride-IpRange"); ok {
					if err = d.Set("ip_range", vv); err != nil {
						return fmt.Errorf("Error reading ip_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip6_range", flattenObjectSystemGeoipOverrideIp6Range(o["ip6-range"], d, "ip6_range")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip6-range"], "ObjectSystemGeoipOverride-Ip6Range"); ok {
				if err = d.Set("ip6_range", vv); err != nil {
					return fmt.Errorf("Error reading ip6_range: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip6_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_range"); ok {
			if err = d.Set("ip6_range", flattenObjectSystemGeoipOverrideIp6Range(o["ip6-range"], d, "ip6_range")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip6-range"], "ObjectSystemGeoipOverride-Ip6Range"); ok {
					if err = d.Set("ip6_range", vv); err != nil {
						return fmt.Errorf("Error reading ip6_range: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip6_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSystemGeoipOverrideName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemGeoipOverride-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemGeoipOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemGeoipOverrideCountryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectSystemGeoipOverrideIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemGeoipOverrideIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectSystemGeoipOverrideIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemGeoipOverrideIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIp6Range(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-ip"], _ = expandObjectSystemGeoipOverrideIp6RangeEndIp(d, i["end_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectSystemGeoipOverrideIp6RangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-ip"], _ = expandObjectSystemGeoipOverrideIp6RangeStartIp(d, i["start_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemGeoipOverrideIp6RangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIp6RangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideIp6RangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemGeoipOverrideName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemGeoipOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("country_id"); ok || d.HasChange("country_id") {
		t, err := expandObjectSystemGeoipOverrideCountryId(d, v, "country_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSystemGeoipOverrideDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandObjectSystemGeoipOverrideIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ip6_range"); ok || d.HasChange("ip6_range") {
		t, err := expandObjectSystemGeoipOverrideIp6Range(d, v, "ip6_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-range"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemGeoipOverrideName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
