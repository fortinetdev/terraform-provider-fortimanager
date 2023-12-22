// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 subnet segments.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAddress6TemplateSubnetSegment() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAddress6TemplateSubnetSegmentCreate,
		Read:   resourceObjectFirewallAddress6TemplateSubnetSegmentRead,
		Update: resourceObjectFirewallAddress6TemplateSubnetSegmentUpdate,
		Delete: resourceObjectFirewallAddress6TemplateSubnetSegmentDelete,

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
			"address6_template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"exclusive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	address6_template := d.Get("address6_template").(string)
	paradict["address6_template"] = address6_template

	obj, err := getObjectObjectFirewallAddress6TemplateSubnetSegment(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6TemplateSubnetSegment resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallAddress6TemplateSubnetSegment(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallAddress6TemplateSubnetSegmentRead(d, m)
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentUpdate(d *schema.ResourceData, m interface{}) error {
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

	address6_template := d.Get("address6_template").(string)
	paradict["address6_template"] = address6_template

	obj, err := getObjectObjectFirewallAddress6TemplateSubnetSegment(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6TemplateSubnetSegment resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallAddress6TemplateSubnetSegment(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallAddress6TemplateSubnetSegmentRead(d, m)
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentDelete(d *schema.ResourceData, m interface{}) error {
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

	address6_template := d.Get("address6_template").(string)
	paradict["address6_template"] = address6_template

	err = c.DeleteObjectFirewallAddress6TemplateSubnetSegment(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentRead(d *schema.ResourceData, m interface{}) error {
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

	address6_template := d.Get("address6_template").(string)
	if address6_template == "" {
		address6_template = importOptionChecking(m.(*FortiClient).Cfg, "address6_template")
		if address6_template == "" {
			return fmt.Errorf("Parameter address6_template is missing")
		}
		if err = d.Set("address6_template", address6_template); err != nil {
			return fmt.Errorf("Error set params address6_template: %v", err)
		}
	}
	paradict["address6_template"] = address6_template

	o, err := c.ReadObjectFirewallAddress6TemplateSubnetSegment(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6TemplateSubnetSegment resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAddress6TemplateSubnetSegment(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6TemplateSubnetSegment resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentBits2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentExclusive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValues2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentValuesName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6TemplateSubnetSegment-Values-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectFirewallAddress6TemplateSubnetSegmentValuesValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectFirewallAddress6TemplateSubnetSegment-Values-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("bits", flattenObjectFirewallAddress6TemplateSubnetSegmentBits2edl(o["bits"], d, "bits")); err != nil {
		if vv, ok := fortiAPIPatch(o["bits"], "ObjectFirewallAddress6TemplateSubnetSegment-Bits"); ok {
			if err = d.Set("bits", vv); err != nil {
				return fmt.Errorf("Error reading bits: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bits: %v", err)
		}
	}

	if err = d.Set("exclusive", flattenObjectFirewallAddress6TemplateSubnetSegmentExclusive2edl(o["exclusive"], d, "exclusive")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclusive"], "ObjectFirewallAddress6TemplateSubnetSegment-Exclusive"); ok {
			if err = d.Set("exclusive", vv); err != nil {
				return fmt.Errorf("Error reading exclusive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclusive: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallAddress6TemplateSubnetSegmentId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallAddress6TemplateSubnetSegment-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallAddress6TemplateSubnetSegmentName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAddress6TemplateSubnetSegment-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("values", flattenObjectFirewallAddress6TemplateSubnetSegmentValues2edl(o["values"], d, "values")); err != nil {
			if vv, ok := fortiAPIPatch(o["values"], "ObjectFirewallAddress6TemplateSubnetSegment-Values"); ok {
				if err = d.Set("values", vv); err != nil {
					return fmt.Errorf("Error reading values: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading values: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("values"); ok {
			if err = d.Set("values", flattenObjectFirewallAddress6TemplateSubnetSegmentValues2edl(o["values"], d, "values")); err != nil {
				if vv, ok := fortiAPIPatch(o["values"], "ObjectFirewallAddress6TemplateSubnetSegment-Values"); ok {
					if err = d.Set("values", vv); err != nil {
						return fmt.Errorf("Error reading values: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading values: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAddress6TemplateSubnetSegmentBits2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentExclusive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValues2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentValuesName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectFirewallAddress6TemplateSubnetSegmentValuesValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValuesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValuesValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bits"); ok || d.HasChange("bits") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentBits2edl(d, v, "bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bits"] = t
		}
	}

	if v, ok := d.GetOk("exclusive"); ok || d.HasChange("exclusive") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentExclusive2edl(d, v, "exclusive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusive"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("values"); ok || d.HasChange("values") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentValues2edl(d, v, "values")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["values"] = t
		}
	}

	return &obj, nil
}
