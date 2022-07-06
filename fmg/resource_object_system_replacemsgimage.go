// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure replacement message images.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemReplacemsgImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemReplacemsgImageCreate,
		Read:   resourceObjectSystemReplacemsgImageRead,
		Update: resourceObjectSystemReplacemsgImageUpdate,
		Delete: resourceObjectSystemReplacemsgImageDelete,

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
			"image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemReplacemsgImageCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemReplacemsgImage(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemReplacemsgImage resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemReplacemsgImage(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemReplacemsgImage resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemReplacemsgImageRead(d, m)
}

func resourceObjectSystemReplacemsgImageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemReplacemsgImage(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemReplacemsgImage resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemReplacemsgImage(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemReplacemsgImage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemReplacemsgImageRead(d, m)
}

func resourceObjectSystemReplacemsgImageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemReplacemsgImage(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemReplacemsgImage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemReplacemsgImageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemReplacemsgImage(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemReplacemsgImage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemReplacemsgImage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemReplacemsgImage resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemReplacemsgImageImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgImageImageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemReplacemsgImageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemReplacemsgImage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("image_base64", flattenObjectSystemReplacemsgImageImageBase64(o["image-base64"], d, "image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-base64"], "ObjectSystemReplacemsgImage-ImageBase64"); ok {
			if err = d.Set("image_base64", vv); err != nil {
				return fmt.Errorf("Error reading image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_base64: %v", err)
		}
	}

	if err = d.Set("image_type", flattenObjectSystemReplacemsgImageImageType(o["image-type"], d, "image_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-type"], "ObjectSystemReplacemsgImage-ImageType"); ok {
			if err = d.Set("image_type", vv); err != nil {
				return fmt.Errorf("Error reading image_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemReplacemsgImageName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemReplacemsgImage-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemReplacemsgImageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemReplacemsgImageImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgImageImageType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemReplacemsgImageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemReplacemsgImage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("image_base64"); ok || d.HasChange("image_base64") {
		t, err := expandObjectSystemReplacemsgImageImageBase64(d, v, "image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-base64"] = t
		}
	}

	if v, ok := d.GetOk("image_type"); ok || d.HasChange("image_type") {
		t, err := expandObjectSystemReplacemsgImageImageType(d, v, "image_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemReplacemsgImageName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
