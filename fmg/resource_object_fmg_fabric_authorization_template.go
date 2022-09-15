// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg FabricAuthorizationTemplate

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFmgFabricAuthorizationTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgFabricAuthorizationTemplateCreate,
		Read:   resourceObjectFmgFabricAuthorizationTemplateRead,
		Update: resourceObjectFmgFabricAuthorizationTemplateUpdate,
		Delete: resourceObjectFmgFabricAuthorizationTemplateDelete,

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
			"extender_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"platforms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"extension_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortilink": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wireless_controller": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectFmgFabricAuthorizationTemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFmgFabricAuthorizationTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgFabricAuthorizationTemplate resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFmgFabricAuthorizationTemplate(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgFabricAuthorizationTemplate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFmgFabricAuthorizationTemplateRead(d, m)
}

func resourceObjectFmgFabricAuthorizationTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFmgFabricAuthorizationTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgFabricAuthorizationTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFmgFabricAuthorizationTemplate(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgFabricAuthorizationTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFmgFabricAuthorizationTemplateRead(d, m)
}

func resourceObjectFmgFabricAuthorizationTemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFmgFabricAuthorizationTemplate(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgFabricAuthorizationTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgFabricAuthorizationTemplateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFmgFabricAuthorizationTemplate(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgFabricAuthorizationTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgFabricAuthorizationTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgFabricAuthorizationTemplate resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgFabricAuthorizationTemplateDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplateExtenderController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatforms(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "count"
		if _, ok := i["count"]; ok {
			v := flattenObjectFmgFabricAuthorizationTemplatePlatformsCount(i["count"], d, pre_append)
			tmp["count"] = fortiAPISubPartPatch(v, "ObjectFmgFabricAuthorizationTemplate-Platforms-Count")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extension_type"
		if _, ok := i["extension-type"]; ok {
			v := flattenObjectFmgFabricAuthorizationTemplatePlatformsExtensionType(i["extension-type"], d, pre_append)
			tmp["extension_type"] = fortiAPISubPartPatch(v, "ObjectFmgFabricAuthorizationTemplate-Platforms-ExtensionType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink"
		if _, ok := i["fortilink"]; ok {
			v := flattenObjectFmgFabricAuthorizationTemplatePlatformsFortilink(i["fortilink"], d, pre_append)
			tmp["fortilink"] = fortiAPISubPartPatch(v, "ObjectFmgFabricAuthorizationTemplate-Platforms-Fortilink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectFmgFabricAuthorizationTemplatePlatformsPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectFmgFabricAuthorizationTemplate-Platforms-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectFmgFabricAuthorizationTemplatePlatformsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectFmgFabricAuthorizationTemplate-Platforms-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplateSwitchController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplateWirelessController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFmgFabricAuthorizationTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenObjectFmgFabricAuthorizationTemplateDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectFmgFabricAuthorizationTemplate-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("extender_controller", flattenObjectFmgFabricAuthorizationTemplateExtenderController(o["extender-controller"], d, "extender_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["extender-controller"], "ObjectFmgFabricAuthorizationTemplate-ExtenderController"); ok {
			if err = d.Set("extender_controller", vv); err != nil {
				return fmt.Errorf("Error reading extender_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extender_controller: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFmgFabricAuthorizationTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFmgFabricAuthorizationTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("platforms", flattenObjectFmgFabricAuthorizationTemplatePlatforms(o["platforms"], d, "platforms")); err != nil {
			if vv, ok := fortiAPIPatch(o["platforms"], "ObjectFmgFabricAuthorizationTemplate-Platforms"); ok {
				if err = d.Set("platforms", vv); err != nil {
					return fmt.Errorf("Error reading platforms: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading platforms: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("platforms"); ok {
			if err = d.Set("platforms", flattenObjectFmgFabricAuthorizationTemplatePlatforms(o["platforms"], d, "platforms")); err != nil {
				if vv, ok := fortiAPIPatch(o["platforms"], "ObjectFmgFabricAuthorizationTemplate-Platforms"); ok {
					if err = d.Set("platforms", vv); err != nil {
						return fmt.Errorf("Error reading platforms: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading platforms: %v", err)
				}
			}
		}
	}

	if err = d.Set("switch_controller", flattenObjectFmgFabricAuthorizationTemplateSwitchController(o["switch-controller"], d, "switch_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller"], "ObjectFmgFabricAuthorizationTemplate-SwitchController"); ok {
			if err = d.Set("switch_controller", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	if err = d.Set("wireless_controller", flattenObjectFmgFabricAuthorizationTemplateWirelessController(o["wireless-controller"], d, "wireless_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-controller"], "ObjectFmgFabricAuthorizationTemplate-WirelessController"); ok {
			if err = d.Set("wireless_controller", vv); err != nil {
				return fmt.Errorf("Error reading wireless_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_controller: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgFabricAuthorizationTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgFabricAuthorizationTemplateDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplateExtenderController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatforms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["count"], _ = expandObjectFmgFabricAuthorizationTemplatePlatformsCount(d, i["count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extension_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extension-type"], _ = expandObjectFmgFabricAuthorizationTemplatePlatformsExtensionType(d, i["extension_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortilink"], _ = expandObjectFmgFabricAuthorizationTemplatePlatformsFortilink(d, i["fortilink"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectFmgFabricAuthorizationTemplatePlatformsPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectFmgFabricAuthorizationTemplatePlatformsType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsExtensionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplateSwitchController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplateWirelessController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFmgFabricAuthorizationTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectFmgFabricAuthorizationTemplateDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("extender_controller"); ok || d.HasChange("extender_controller") {
		t, err := expandObjectFmgFabricAuthorizationTemplateExtenderController(d, v, "extender_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extender-controller"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFmgFabricAuthorizationTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("platforms"); ok || d.HasChange("platforms") {
		t, err := expandObjectFmgFabricAuthorizationTemplatePlatforms(d, v, "platforms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platforms"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok || d.HasChange("switch_controller") {
		t, err := expandObjectFmgFabricAuthorizationTemplateSwitchController(d, v, "switch_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller"] = t
		}
	}

	if v, ok := d.GetOk("wireless_controller"); ok || d.HasChange("wireless_controller") {
		t, err := expandObjectFmgFabricAuthorizationTemplateWirelessController(d, v, "wireless_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-controller"] = t
		}
	}

	return &obj, nil
}
