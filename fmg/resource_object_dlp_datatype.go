// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure predefined data type used by DLP blocking.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDlpDataType() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpDataTypeCreate,
		Read:   resourceObjectDlpDataTypeRead,
		Update: resourceObjectDlpDataTypeUpdate,
		Delete: resourceObjectDlpDataTypeDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"look_ahead": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"look_back": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"match_around": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_transformed_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectDlpDataTypeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectDlpDataType(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpDataType resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpDataType(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpDataType resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpDataTypeRead(d, m)
}

func resourceObjectDlpDataTypeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectDlpDataType(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpDataType resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpDataType(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpDataType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectDlpDataTypeRead(d, m)
}

func resourceObjectDlpDataTypeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectDlpDataType(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpDataType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpDataTypeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectDlpDataType(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpDataType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpDataType(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpDataType resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpDataTypeComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeLookAhead(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeLookBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeMatchAround(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeTransform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpDataTypeVerifyTransformedPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpDataType(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectDlpDataTypeComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDlpDataType-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("look_ahead", flattenObjectDlpDataTypeLookAhead(o["look-ahead"], d, "look_ahead")); err != nil {
		if vv, ok := fortiAPIPatch(o["look-ahead"], "ObjectDlpDataType-LookAhead"); ok {
			if err = d.Set("look_ahead", vv); err != nil {
				return fmt.Errorf("Error reading look_ahead: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading look_ahead: %v", err)
		}
	}

	if err = d.Set("look_back", flattenObjectDlpDataTypeLookBack(o["look-back"], d, "look_back")); err != nil {
		if vv, ok := fortiAPIPatch(o["look-back"], "ObjectDlpDataType-LookBack"); ok {
			if err = d.Set("look_back", vv); err != nil {
				return fmt.Errorf("Error reading look_back: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading look_back: %v", err)
		}
	}

	if err = d.Set("match_around", flattenObjectDlpDataTypeMatchAround(o["match-around"], d, "match_around")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-around"], "ObjectDlpDataType-MatchAround"); ok {
			if err = d.Set("match_around", vv); err != nil {
				return fmt.Errorf("Error reading match_around: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_around: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDlpDataTypeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpDataType-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pattern", flattenObjectDlpDataTypePattern(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "ObjectDlpDataType-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("transform", flattenObjectDlpDataTypeTransform(o["transform"], d, "transform")); err != nil {
		if vv, ok := fortiAPIPatch(o["transform"], "ObjectDlpDataType-Transform"); ok {
			if err = d.Set("transform", vv); err != nil {
				return fmt.Errorf("Error reading transform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transform: %v", err)
		}
	}

	if err = d.Set("verify", flattenObjectDlpDataTypeVerify(o["verify"], d, "verify")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify"], "ObjectDlpDataType-Verify"); ok {
			if err = d.Set("verify", vv); err != nil {
				return fmt.Errorf("Error reading verify: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify: %v", err)
		}
	}

	if err = d.Set("verify_transformed_pattern", flattenObjectDlpDataTypeVerifyTransformedPattern(o["verify-transformed-pattern"], d, "verify_transformed_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-transformed-pattern"], "ObjectDlpDataType-VerifyTransformedPattern"); ok {
			if err = d.Set("verify_transformed_pattern", vv); err != nil {
				return fmt.Errorf("Error reading verify_transformed_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_transformed_pattern: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpDataTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpDataTypeComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeLookAhead(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeLookBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeMatchAround(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeTransform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpDataTypeVerifyTransformedPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpDataType(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectDlpDataTypeComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("look_ahead"); ok || d.HasChange("look_ahead") {
		t, err := expandObjectDlpDataTypeLookAhead(d, v, "look_ahead")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["look-ahead"] = t
		}
	}

	if v, ok := d.GetOk("look_back"); ok || d.HasChange("look_back") {
		t, err := expandObjectDlpDataTypeLookBack(d, v, "look_back")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["look-back"] = t
		}
	}

	if v, ok := d.GetOk("match_around"); ok || d.HasChange("match_around") {
		t, err := expandObjectDlpDataTypeMatchAround(d, v, "match_around")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-around"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectDlpDataTypeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandObjectDlpDataTypePattern(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("transform"); ok || d.HasChange("transform") {
		t, err := expandObjectDlpDataTypeTransform(d, v, "transform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transform"] = t
		}
	}

	if v, ok := d.GetOk("verify"); ok || d.HasChange("verify") {
		t, err := expandObjectDlpDataTypeVerify(d, v, "verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify"] = t
		}
	}

	if v, ok := d.GetOk("verify_transformed_pattern"); ok || d.HasChange("verify_transformed_pattern") {
		t, err := expandObjectDlpDataTypeVerifyTransformedPattern(d, v, "verify_transformed_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-transformed-pattern"] = t
		}
	}

	return &obj, nil
}
