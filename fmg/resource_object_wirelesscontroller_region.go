// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiAP regions (for floor plans and maps).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerRegion() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerRegionCreate,
		Read:   resourceObjectWirelessControllerRegionRead,
		Update: resourceObjectWirelessControllerRegionUpdate,
		Delete: resourceObjectWirelessControllerRegionDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grayscale": &schema.Schema{
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
			"opacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerRegionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerRegion(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerRegion resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerRegion(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerRegion resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerRegionRead(d, m)
}

func resourceObjectWirelessControllerRegionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerRegion(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerRegion resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerRegion(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerRegion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerRegionRead(d, m)
}

func resourceObjectWirelessControllerRegionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerRegion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerRegion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerRegionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerRegion(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerRegion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerRegion(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerRegion resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerRegionComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerRegionGrayscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerRegionImageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "gif",
			2: "jpeg",
			3: "png",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerRegionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerRegionOpacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerRegion(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comments", flattenObjectWirelessControllerRegionComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectWirelessControllerRegion-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("grayscale", flattenObjectWirelessControllerRegionGrayscale(o["grayscale"], d, "grayscale")); err != nil {
		if vv, ok := fortiAPIPatch(o["grayscale"], "ObjectWirelessControllerRegion-Grayscale"); ok {
			if err = d.Set("grayscale", vv); err != nil {
				return fmt.Errorf("Error reading grayscale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading grayscale: %v", err)
		}
	}

	if err = d.Set("image_type", flattenObjectWirelessControllerRegionImageType(o["image-type"], d, "image_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-type"], "ObjectWirelessControllerRegion-ImageType"); ok {
			if err = d.Set("image_type", vv); err != nil {
				return fmt.Errorf("Error reading image_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerRegionName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerRegion-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("opacity", flattenObjectWirelessControllerRegionOpacity(o["opacity"], d, "opacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["opacity"], "ObjectWirelessControllerRegion-Opacity"); ok {
			if err = d.Set("opacity", vv); err != nil {
				return fmt.Errorf("Error reading opacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading opacity: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerRegionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerRegionComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerRegionGrayscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerRegionImageType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerRegionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerRegionOpacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerRegion(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandObjectWirelessControllerRegionComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("grayscale"); ok {
		t, err := expandObjectWirelessControllerRegionGrayscale(d, v, "grayscale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grayscale"] = t
		}
	}

	if v, ok := d.GetOk("image_type"); ok {
		t, err := expandObjectWirelessControllerRegionImageType(d, v, "image_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerRegionName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("opacity"); ok {
		t, err := expandObjectWirelessControllerRegionOpacity(d, v, "opacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opacity"] = t
		}
	}

	return &obj, nil
}
