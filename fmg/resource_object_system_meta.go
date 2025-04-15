// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectSystem Meta

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemMeta() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemMetaCreate,
		Read:   resourceObjectSystemMetaRead,
		Update: resourceObjectSystemMetaUpdate,
		Delete: resourceObjectSystemMetaDelete,

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
			"sys_meta_fields": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fieldlength": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"importance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
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

func resourceObjectSystemMetaCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemMeta(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemMeta resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemMeta(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemMeta resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemMetaRead(d, m)
}

func resourceObjectSystemMetaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemMeta(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemMeta resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemMeta(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemMeta resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemMetaRead(d, m)
}

func resourceObjectSystemMetaDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemMeta(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemMeta resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemMetaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemMeta(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemMeta resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemMeta(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemMeta resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemMetaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemMetaSysMetaFields(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldlength"
		if _, ok := i["fieldlength"]; ok {
			v := flattenObjectSystemMetaSysMetaFieldsFieldlength(i["fieldlength"], d, pre_append)
			tmp["fieldlength"] = fortiAPISubPartPatch(v, "ObjectSystemMeta-SysMetaFields-Fieldlength")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "importance"
		if _, ok := i["importance"]; ok {
			v := flattenObjectSystemMetaSysMetaFieldsImportance(i["importance"], d, pre_append)
			tmp["importance"] = fortiAPISubPartPatch(v, "ObjectSystemMeta-SysMetaFields-Importance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemMetaSysMetaFieldsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemMeta-SysMetaFields-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectSystemMetaSysMetaFieldsFieldlength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemMetaSysMetaFieldsImportance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemMetaSysMetaFieldsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemMeta(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenObjectSystemMetaName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemMeta-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sys_meta_fields", flattenObjectSystemMetaSysMetaFields(o["sys_meta_fields"], d, "sys_meta_fields")); err != nil {
			if vv, ok := fortiAPIPatch(o["sys_meta_fields"], "ObjectSystemMeta-SysMetaFields"); ok {
				if err = d.Set("sys_meta_fields", vv); err != nil {
					return fmt.Errorf("Error reading sys_meta_fields: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sys_meta_fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sys_meta_fields"); ok {
			if err = d.Set("sys_meta_fields", flattenObjectSystemMetaSysMetaFields(o["sys_meta_fields"], d, "sys_meta_fields")); err != nil {
				if vv, ok := fortiAPIPatch(o["sys_meta_fields"], "ObjectSystemMeta-SysMetaFields"); ok {
					if err = d.Set("sys_meta_fields", vv); err != nil {
						return fmt.Errorf("Error reading sys_meta_fields: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sys_meta_fields: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectSystemMetaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemMetaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemMetaSysMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldlength"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fieldlength"], _ = expandObjectSystemMetaSysMetaFieldsFieldlength(d, i["fieldlength"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "importance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["importance"], _ = expandObjectSystemMetaSysMetaFieldsImportance(d, i["importance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSystemMetaSysMetaFieldsName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectSystemMetaSysMetaFieldsFieldlength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemMetaSysMetaFieldsImportance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemMetaSysMetaFieldsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemMeta(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemMetaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sys_meta_fields"); ok || d.HasChange("sys_meta_fields") {
		t, err := expandObjectSystemMetaSysMetaFields(d, v, "sys_meta_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sys_meta_fields"] = t
		}
	}

	return &obj, nil
}
