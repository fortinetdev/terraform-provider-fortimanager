// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Subnet segment values.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAddress6TemplateSubnetSegmentValues() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAddress6TemplateSubnetSegmentValuesCreate,
		Read:   resourceObjectFirewallAddress6TemplateSubnetSegmentValuesRead,
		Update: resourceObjectFirewallAddress6TemplateSubnetSegmentValuesUpdate,
		Delete: resourceObjectFirewallAddress6TemplateSubnetSegmentValuesDelete,

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
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentValuesCreate(d *schema.ResourceData, m interface{}) error {
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

	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	obj, err := getObjectObjectFirewallAddress6TemplateSubnetSegmentValues(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6TemplateSubnetSegmentValues resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallAddress6TemplateSubnetSegmentValues(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddress6TemplateSubnetSegmentValuesRead(d, m)
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentValuesUpdate(d *schema.ResourceData, m interface{}) error {
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

	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	obj, err := getObjectObjectFirewallAddress6TemplateSubnetSegmentValues(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6TemplateSubnetSegmentValues resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallAddress6TemplateSubnetSegmentValues(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAddress6TemplateSubnetSegmentValuesRead(d, m)
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentValuesDelete(d *schema.ResourceData, m interface{}) error {
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

	address6_template := d.Get("address6_template").(string)
	subnet_segment := d.Get("subnet_segment").(string)
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallAddress6TemplateSubnetSegmentValues(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAddress6TemplateSubnetSegmentValuesRead(d *schema.ResourceData, m interface{}) error {
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
	subnet_segment := d.Get("subnet_segment").(string)
	if address6_template == "" {
		address6_template = importOptionChecking(m.(*FortiClient).Cfg, "address6_template")
		if address6_template == "" {
			return fmt.Errorf("Parameter address6_template is missing")
		}
		if err = d.Set("address6_template", address6_template); err != nil {
			return fmt.Errorf("Error set params address6_template: %v", err)
		}
	}
	if subnet_segment == "" {
		subnet_segment = importOptionChecking(m.(*FortiClient).Cfg, "subnet_segment")
		if subnet_segment == "" {
			return fmt.Errorf("Parameter subnet_segment is missing")
		}
		if err = d.Set("subnet_segment", subnet_segment); err != nil {
			return fmt.Errorf("Error set params subnet_segment: %v", err)
		}
	}
	paradict["address6_template"] = address6_template
	paradict["subnet_segment"] = subnet_segment

	o, err := c.ReadObjectFirewallAddress6TemplateSubnetSegmentValues(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6TemplateSubnetSegmentValues resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAddress6TemplateSubnetSegmentValues(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAddress6TemplateSubnetSegmentValues resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectFirewallAddress6TemplateSubnetSegmentValuesName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAddress6TemplateSubnetSegmentValues-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectFirewallAddress6TemplateSubnetSegmentValues-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAddress6TemplateSubnetSegmentValuesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValuesName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentValuesName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectFirewallAddress6TemplateSubnetSegmentValuesValue3rdl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
