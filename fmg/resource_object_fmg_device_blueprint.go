// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg DeviceBlueprint

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFmgDeviceBlueprint() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgDeviceBlueprintCreate,
		Read:   resourceObjectFmgDeviceBlueprintRead,
		Update: resourceObjectFmgDeviceBlueprintUpdate,
		Delete: resourceObjectFmgDeviceBlueprintDelete,

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
			"cliprofs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dev_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefer_img_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prerun_cliprof": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"prov_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"templates": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceObjectFmgDeviceBlueprintCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFmgDeviceBlueprint(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgDeviceBlueprint resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFmgDeviceBlueprint(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgDeviceBlueprint resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFmgDeviceBlueprintRead(d, m)
}

func resourceObjectFmgDeviceBlueprintUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFmgDeviceBlueprint(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgDeviceBlueprint resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFmgDeviceBlueprint(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgDeviceBlueprint resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFmgDeviceBlueprintRead(d, m)
}

func resourceObjectFmgDeviceBlueprintDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFmgDeviceBlueprint(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgDeviceBlueprint resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgDeviceBlueprintRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFmgDeviceBlueprint(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgDeviceBlueprint resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgDeviceBlueprint(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgDeviceBlueprint resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgDeviceBlueprintCliprofs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFmgDeviceBlueprintDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintDevGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFmgDeviceBlueprintFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintPkg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintPlatform(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintPreferImgVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintPrerunCliprof(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFmgDeviceBlueprintProvType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintTemplateGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgDeviceBlueprintTemplates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFmgDeviceBlueprint(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cliprofs", flattenObjectFmgDeviceBlueprintCliprofs(o["cliprofs"], d, "cliprofs")); err != nil {
		if vv, ok := fortiAPIPatch(o["cliprofs"], "ObjectFmgDeviceBlueprint-Cliprofs"); ok {
			if err = d.Set("cliprofs", vv); err != nil {
				return fmt.Errorf("Error reading cliprofs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cliprofs: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectFmgDeviceBlueprintDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectFmgDeviceBlueprint-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("dev_group", flattenObjectFmgDeviceBlueprintDevGroup(o["dev-group"], d, "dev_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-group"], "ObjectFmgDeviceBlueprint-DevGroup"); ok {
			if err = d.Set("dev_group", vv); err != nil {
				return fmt.Errorf("Error reading dev_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_group: %v", err)
		}
	}

	if err = d.Set("folder", flattenObjectFmgDeviceBlueprintFolder(o["folder"], d, "folder")); err != nil {
		if vv, ok := fortiAPIPatch(o["folder"], "ObjectFmgDeviceBlueprint-Folder"); ok {
			if err = d.Set("folder", vv); err != nil {
				return fmt.Errorf("Error reading folder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFmgDeviceBlueprintName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFmgDeviceBlueprint-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pkg", flattenObjectFmgDeviceBlueprintPkg(o["pkg"], d, "pkg")); err != nil {
		if vv, ok := fortiAPIPatch(o["pkg"], "ObjectFmgDeviceBlueprint-Pkg"); ok {
			if err = d.Set("pkg", vv); err != nil {
				return fmt.Errorf("Error reading pkg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pkg: %v", err)
		}
	}

	if err = d.Set("platform", flattenObjectFmgDeviceBlueprintPlatform(o["platform"], d, "platform")); err != nil {
		if vv, ok := fortiAPIPatch(o["platform"], "ObjectFmgDeviceBlueprint-Platform"); ok {
			if err = d.Set("platform", vv); err != nil {
				return fmt.Errorf("Error reading platform: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading platform: %v", err)
		}
	}

	if err = d.Set("prefer_img_ver", flattenObjectFmgDeviceBlueprintPreferImgVer(o["prefer-img-ver"], d, "prefer_img_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefer-img-ver"], "ObjectFmgDeviceBlueprint-PreferImgVer"); ok {
			if err = d.Set("prefer_img_ver", vv); err != nil {
				return fmt.Errorf("Error reading prefer_img_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefer_img_ver: %v", err)
		}
	}

	if err = d.Set("prerun_cliprof", flattenObjectFmgDeviceBlueprintPrerunCliprof(o["prerun-cliprof"], d, "prerun_cliprof")); err != nil {
		if vv, ok := fortiAPIPatch(o["prerun-cliprof"], "ObjectFmgDeviceBlueprint-PrerunCliprof"); ok {
			if err = d.Set("prerun_cliprof", vv); err != nil {
				return fmt.Errorf("Error reading prerun_cliprof: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prerun_cliprof: %v", err)
		}
	}

	if err = d.Set("prov_type", flattenObjectFmgDeviceBlueprintProvType(o["prov-type"], d, "prov_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["prov-type"], "ObjectFmgDeviceBlueprint-ProvType"); ok {
			if err = d.Set("prov_type", vv); err != nil {
				return fmt.Errorf("Error reading prov_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prov_type: %v", err)
		}
	}

	if err = d.Set("template_group", flattenObjectFmgDeviceBlueprintTemplateGroup(o["template-group"], d, "template_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["template-group"], "ObjectFmgDeviceBlueprint-TemplateGroup"); ok {
			if err = d.Set("template_group", vv); err != nil {
				return fmt.Errorf("Error reading template_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template_group: %v", err)
		}
	}

	if err = d.Set("templates", flattenObjectFmgDeviceBlueprintTemplates(o["templates"], d, "templates")); err != nil {
		if vv, ok := fortiAPIPatch(o["templates"], "ObjectFmgDeviceBlueprint-Templates"); ok {
			if err = d.Set("templates", vv); err != nil {
				return fmt.Errorf("Error reading templates: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading templates: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgDeviceBlueprintFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgDeviceBlueprintCliprofs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFmgDeviceBlueprintDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintDevGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFmgDeviceBlueprintFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintPkg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintPlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintPreferImgVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintPrerunCliprof(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFmgDeviceBlueprintProvType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintTemplateGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgDeviceBlueprintTemplates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFmgDeviceBlueprint(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cliprofs"); ok || d.HasChange("cliprofs") {
		t, err := expandObjectFmgDeviceBlueprintCliprofs(d, v, "cliprofs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cliprofs"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectFmgDeviceBlueprintDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("dev_group"); ok || d.HasChange("dev_group") {
		t, err := expandObjectFmgDeviceBlueprintDevGroup(d, v, "dev_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-group"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok || d.HasChange("folder") {
		t, err := expandObjectFmgDeviceBlueprintFolder(d, v, "folder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFmgDeviceBlueprintName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pkg"); ok || d.HasChange("pkg") {
		t, err := expandObjectFmgDeviceBlueprintPkg(d, v, "pkg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pkg"] = t
		}
	}

	if v, ok := d.GetOk("platform"); ok || d.HasChange("platform") {
		t, err := expandObjectFmgDeviceBlueprintPlatform(d, v, "platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform"] = t
		}
	}

	if v, ok := d.GetOk("prefer_img_ver"); ok || d.HasChange("prefer_img_ver") {
		t, err := expandObjectFmgDeviceBlueprintPreferImgVer(d, v, "prefer_img_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer-img-ver"] = t
		}
	}

	if v, ok := d.GetOk("prerun_cliprof"); ok || d.HasChange("prerun_cliprof") {
		t, err := expandObjectFmgDeviceBlueprintPrerunCliprof(d, v, "prerun_cliprof")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prerun-cliprof"] = t
		}
	}

	if v, ok := d.GetOk("prov_type"); ok || d.HasChange("prov_type") {
		t, err := expandObjectFmgDeviceBlueprintProvType(d, v, "prov_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prov-type"] = t
		}
	}

	if v, ok := d.GetOk("template_group"); ok || d.HasChange("template_group") {
		t, err := expandObjectFmgDeviceBlueprintTemplateGroup(d, v, "template_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-group"] = t
		}
	}

	if v, ok := d.GetOk("templates"); ok || d.HasChange("templates") {
		t, err := expandObjectFmgDeviceBlueprintTemplates(d, v, "templates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["templates"] = t
		}
	}

	return &obj, nil
}
