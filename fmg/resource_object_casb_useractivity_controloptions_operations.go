// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB control option operations.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityControlOptionsOperations() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityControlOptionsOperationsCreate,
		Read:   resourceObjectCasbUserActivityControlOptionsOperationsRead,
		Update: resourceObjectCasbUserActivityControlOptionsOperationsUpdate,
		Delete: resourceObjectCasbUserActivityControlOptionsOperationsDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"control_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"search_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value_from_input": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"values": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCasbUserActivityControlOptionsOperationsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	user_activity := d.Get("user_activity").(string)
	control_options := d.Get("control_options").(string)
	paradict["user_activity"] = user_activity
	paradict["control_options"] = control_options

	obj, err := getObjectObjectCasbUserActivityControlOptionsOperations(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityControlOptionsOperations resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCasbUserActivityControlOptionsOperations(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityControlOptionsOperations resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbUserActivityControlOptionsOperationsRead(d, m)
}

func resourceObjectCasbUserActivityControlOptionsOperationsUpdate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	control_options := d.Get("control_options").(string)
	paradict["user_activity"] = user_activity
	paradict["control_options"] = control_options

	obj, err := getObjectObjectCasbUserActivityControlOptionsOperations(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityControlOptionsOperations resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCasbUserActivityControlOptionsOperations(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityControlOptionsOperations resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbUserActivityControlOptionsOperationsRead(d, m)
}

func resourceObjectCasbUserActivityControlOptionsOperationsDelete(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	control_options := d.Get("control_options").(string)
	paradict["user_activity"] = user_activity
	paradict["control_options"] = control_options

	err = c.DeleteObjectCasbUserActivityControlOptionsOperations(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityControlOptionsOperations resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityControlOptionsOperationsRead(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	control_options := d.Get("control_options").(string)
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	if control_options == "" {
		control_options = importOptionChecking(m.(*FortiClient).Cfg, "control_options")
		if control_options == "" {
			return fmt.Errorf("Parameter control_options is missing")
		}
		if err = d.Set("control_options", control_options); err != nil {
			return fmt.Errorf("Error set params control_options: %v", err)
		}
	}
	paradict["user_activity"] = user_activity
	paradict["control_options"] = control_options

	o, err := c.ReadObjectCasbUserActivityControlOptionsOperations(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityControlOptionsOperations resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityControlOptionsOperations(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityControlOptionsOperations resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityControlOptionsOperationsAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsCaseSensitive3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsSearchKey3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsSearchPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsTarget3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsValueFromInput3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsValues3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectCasbUserActivityControlOptionsOperations(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectCasbUserActivityControlOptionsOperationsAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectCasbUserActivityControlOptionsOperations-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("case_sensitive", flattenObjectCasbUserActivityControlOptionsOperationsCaseSensitive3rdl(o["case-sensitive"], d, "case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitive"], "ObjectCasbUserActivityControlOptionsOperations-CaseSensitive"); ok {
			if err = d.Set("case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}

	if err = d.Set("direction", flattenObjectCasbUserActivityControlOptionsOperationsDirection3rdl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "ObjectCasbUserActivityControlOptionsOperations-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("header_name", flattenObjectCasbUserActivityControlOptionsOperationsHeaderName3rdl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "ObjectCasbUserActivityControlOptionsOperations-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbUserActivityControlOptionsOperationsName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbUserActivityControlOptionsOperations-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("search_key", flattenObjectCasbUserActivityControlOptionsOperationsSearchKey3rdl(o["search-key"], d, "search_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-key"], "ObjectCasbUserActivityControlOptionsOperations-SearchKey"); ok {
			if err = d.Set("search_key", vv); err != nil {
				return fmt.Errorf("Error reading search_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_key: %v", err)
		}
	}

	if err = d.Set("search_pattern", flattenObjectCasbUserActivityControlOptionsOperationsSearchPattern3rdl(o["search-pattern"], d, "search_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["search-pattern"], "ObjectCasbUserActivityControlOptionsOperations-SearchPattern"); ok {
			if err = d.Set("search_pattern", vv); err != nil {
				return fmt.Errorf("Error reading search_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading search_pattern: %v", err)
		}
	}

	if err = d.Set("target", flattenObjectCasbUserActivityControlOptionsOperationsTarget3rdl(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "ObjectCasbUserActivityControlOptionsOperations-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	if err = d.Set("value_from_input", flattenObjectCasbUserActivityControlOptionsOperationsValueFromInput3rdl(o["value-from-input"], d, "value_from_input")); err != nil {
		if vv, ok := fortiAPIPatch(o["value-from-input"], "ObjectCasbUserActivityControlOptionsOperations-ValueFromInput"); ok {
			if err = d.Set("value_from_input", vv); err != nil {
				return fmt.Errorf("Error reading value_from_input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value_from_input: %v", err)
		}
	}

	if err = d.Set("values", flattenObjectCasbUserActivityControlOptionsOperationsValues3rdl(o["values"], d, "values")); err != nil {
		if vv, ok := fortiAPIPatch(o["values"], "ObjectCasbUserActivityControlOptionsOperations-Values"); ok {
			if err = d.Set("values", vv); err != nil {
				return fmt.Errorf("Error reading values: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading values: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityControlOptionsOperationsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityControlOptionsOperationsAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsCaseSensitive3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsSearchKey3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsSearchPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsTarget3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsValueFromInput3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsValues3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectCasbUserActivityControlOptionsOperations(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitive"); ok || d.HasChange("case_sensitive") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsCaseSensitive3rdl(d, v, "case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsDirection3rdl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsHeaderName3rdl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("search_key"); ok || d.HasChange("search_key") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsSearchKey3rdl(d, v, "search_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-key"] = t
		}
	}

	if v, ok := d.GetOk("search_pattern"); ok || d.HasChange("search_pattern") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsSearchPattern3rdl(d, v, "search_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-pattern"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsTarget3rdl(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("value_from_input"); ok || d.HasChange("value_from_input") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsValueFromInput3rdl(d, v, "value_from_input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-from-input"] = t
		}
	}

	if v, ok := d.GetOk("values"); ok || d.HasChange("values") {
		t, err := expandObjectCasbUserActivityControlOptionsOperationsValues3rdl(d, v, "values")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["values"] = t
		}
	}

	return &obj, nil
}
