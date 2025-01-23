// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg FabricAuthorizationTemplatePlatforms

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFmgFabricAuthorizationTemplatePlatforms() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgFabricAuthorizationTemplatePlatformsCreate,
		Read:   resourceObjectFmgFabricAuthorizationTemplatePlatformsRead,
		Update: resourceObjectFmgFabricAuthorizationTemplatePlatformsUpdate,
		Delete: resourceObjectFmgFabricAuthorizationTemplatePlatformsDelete,

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
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fmgcount": &schema.Schema{
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
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFmgFabricAuthorizationTemplatePlatformsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	template := d.Get("template").(string)
	paradict["template"] = template

	obj, err := getObjectObjectFmgFabricAuthorizationTemplatePlatforms(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgFabricAuthorizationTemplatePlatforms resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFmgFabricAuthorizationTemplatePlatforms(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgFabricAuthorizationTemplatePlatforms resource: %v", err)
	}

	d.SetId(getStringKey(d, "prefix"))

	return resourceObjectFmgFabricAuthorizationTemplatePlatformsRead(d, m)
}

func resourceObjectFmgFabricAuthorizationTemplatePlatformsUpdate(d *schema.ResourceData, m interface{}) error {
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

	template := d.Get("template").(string)
	paradict["template"] = template

	obj, err := getObjectObjectFmgFabricAuthorizationTemplatePlatforms(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgFabricAuthorizationTemplatePlatforms resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFmgFabricAuthorizationTemplatePlatforms(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgFabricAuthorizationTemplatePlatforms resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "prefix"))

	return resourceObjectFmgFabricAuthorizationTemplatePlatformsRead(d, m)
}

func resourceObjectFmgFabricAuthorizationTemplatePlatformsDelete(d *schema.ResourceData, m interface{}) error {
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

	template := d.Get("template").(string)
	paradict["template"] = template

	err = c.DeleteObjectFmgFabricAuthorizationTemplatePlatforms(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgFabricAuthorizationTemplatePlatforms resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgFabricAuthorizationTemplatePlatformsRead(d *schema.ResourceData, m interface{}) error {
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

	template := d.Get("template").(string)
	if template == "" {
		template = importOptionChecking(m.(*FortiClient).Cfg, "template")
		if template == "" {
			return fmt.Errorf("Parameter template is missing")
		}
		if err = d.Set("template", template); err != nil {
			return fmt.Errorf("Error set params template: %v", err)
		}
	}
	paradict["template"] = template

	o, err := c.ReadObjectFmgFabricAuthorizationTemplatePlatforms(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgFabricAuthorizationTemplatePlatforms resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgFabricAuthorizationTemplatePlatforms(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgFabricAuthorizationTemplatePlatforms resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsExtensionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsFortilink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFmgFabricAuthorizationTemplatePlatforms(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fmgcount", flattenObjectFmgFabricAuthorizationTemplatePlatformsCount2edl(o["count"], d, "fmgcount")); err != nil {
		if vv, ok := fortiAPIPatch(o["count"], "ObjectFmgFabricAuthorizationTemplatePlatforms-Count"); ok {
			if err = d.Set("fmgcount", vv); err != nil {
				return fmt.Errorf("Error reading fmgcount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgcount: %v", err)
		}
	}

	if err = d.Set("extension_type", flattenObjectFmgFabricAuthorizationTemplatePlatformsExtensionType2edl(o["extension-type"], d, "extension_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-type"], "ObjectFmgFabricAuthorizationTemplatePlatforms-ExtensionType"); ok {
			if err = d.Set("extension_type", vv); err != nil {
				return fmt.Errorf("Error reading extension_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_type: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenObjectFmgFabricAuthorizationTemplatePlatformsFortilink2edl(o["fortilink"], d, "fortilink")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortilink"], "ObjectFmgFabricAuthorizationTemplatePlatforms-Fortilink"); ok {
			if err = d.Set("fortilink", vv); err != nil {
				return fmt.Errorf("Error reading fortilink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("prefix", flattenObjectFmgFabricAuthorizationTemplatePlatformsPrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "ObjectFmgFabricAuthorizationTemplatePlatforms-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFmgFabricAuthorizationTemplatePlatformsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFmgFabricAuthorizationTemplatePlatforms-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgFabricAuthorizationTemplatePlatformsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsExtensionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsFortilink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgFabricAuthorizationTemplatePlatformsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFmgFabricAuthorizationTemplatePlatforms(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgcount"); ok || d.HasChange("fmgcount") {
		t, err := expandObjectFmgFabricAuthorizationTemplatePlatformsCount2edl(d, v, "fmgcount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("extension_type"); ok || d.HasChange("extension_type") {
		t, err := expandObjectFmgFabricAuthorizationTemplatePlatformsExtensionType2edl(d, v, "extension_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-type"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok || d.HasChange("fortilink") {
		t, err := expandObjectFmgFabricAuthorizationTemplatePlatformsFortilink2edl(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandObjectFmgFabricAuthorizationTemplatePlatformsPrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFmgFabricAuthorizationTemplatePlatformsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
