// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 address templates.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAddress6Template() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAddress6TemplateCreate,
		Read:   resourceObjectFirewallAddress6TemplateRead,
		Update: resourceObjectFirewallAddress6TemplateUpdate,
		Delete: resourceObjectFirewallAddress6TemplateDelete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bits": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"exclusive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"values": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"subnet_segment_count": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectFirewallAddress6TemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallAddress6Template(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6Template resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallAddress6Template(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6Template resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddress6TemplateRead(d, m)
}

func resourceObjectFirewallAddress6TemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallAddress6Template(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6Template resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallAddress6Template(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6Template resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddress6TemplateRead(d, m)
}

func resourceObjectFirewallAddress6TemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallAddress6Template(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAddress6Template resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAddress6TemplateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallAddress6Template(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6Template resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAddress6Template(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6Template resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAddress6TemplateImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := i["bits"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentBits(i["bits"], d, pre_append)
			tmp["bits"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6Template-SubnetSegment-Bits")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := i["exclusive"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentExclusive(i["exclusive"], d, pre_append)
			tmp["exclusive"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6Template-SubnetSegment-Exclusive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6Template-SubnetSegment-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6Template-SubnetSegment-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := i["values"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentValues(i["values"], d, pre_append)
			tmp["values"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6Template-SubnetSegment-Values")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValues(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentValuesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6TemplateSubnetSegment-Values-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentValuesValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6TemplateSubnetSegment-Values-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAddress6Template(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenObjectFirewallAddress6TemplateImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "ObjectFirewallAddress6Template-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectFirewallAddress6TemplateFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectFirewallAddress6Template-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("ip6", flattenObjectFirewallAddress6TemplateIp6(o["ip6"], d, "ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6"], "ObjectFirewallAddress6Template-Ip6"); ok {
			if err = d.Set("ip6", vv); err != nil {
				return fmt.Errorf("Error reading ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallAddress6TemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAddress6Template-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("subnet_segment", flattenObjectFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
			if vv, ok := fortiAPIPatch(o["subnet-segment"], "ObjectFirewallAddress6Template-SubnetSegment"); ok {
				if err = d.Set("subnet_segment", vv); err != nil {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading subnet_segment: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("subnet_segment"); ok {
			if err = d.Set("subnet_segment", flattenObjectFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
				if vv, ok := fortiAPIPatch(o["subnet-segment"], "ObjectFirewallAddress6Template-SubnetSegment"); ok {
					if err = d.Set("subnet_segment", vv); err != nil {
						return fmt.Errorf("Error reading subnet_segment: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading subnet_segment: %v", err)
				}
			}
		}
	}

	if err = d.Set("subnet_segment_count", flattenObjectFirewallAddress6TemplateSubnetSegmentCount(o["subnet-segment-count"], d, "subnet_segment_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-segment-count"], "ObjectFirewallAddress6Template-SubnetSegmentCount"); ok {
			if err = d.Set("subnet_segment_count", vv); err != nil {
				return fmt.Errorf("Error reading subnet_segment_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_segment_count: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAddress6TemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAddress6TemplateImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bits"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentBits(d, i["bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["exclusive"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentExclusive(d, i["exclusive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectFirewallAddress6TemplateSubnetSegmentValues(d, i["values"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["values"] = t
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentExclusive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentValuesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentValuesValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValuesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValuesValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAddress6Template(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandObjectFirewallAddress6TemplateImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectFirewallAddress6TemplateFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok || d.HasChange("ip6") {
		t, err := expandObjectFirewallAddress6TemplateIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallAddress6TemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment"); ok || d.HasChange("subnet_segment") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegment(d, v, "subnet_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment_count"); ok || d.HasChange("subnet_segment_count") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentCount(d, v, "subnet_segment_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment-count"] = t
		}
	}

	return &obj, nil
}
